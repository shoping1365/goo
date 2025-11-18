package services

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"gorm.io/gorm"
)

/*
سرویس احراز هویت یکپارچه
-----------------------------------------------------------------
این سرویس گسترش سیستم JWT موجود است و ویژگی‌های زیر را اضافه می‌کند:
- ورود/ثبت‌نام یکپارچه با موبایل
- انتخاب روش احراز (OTP یا پسورد)
- مدیریت کدهای OTP
- تشخیص خودکار کاربران موجود/جدید
- قفل حساب بر اساس تلاش‌های ناموفق
*/

type UnifiedAuthService struct {
	db          *gorm.DB
	authService *AuthService // سرویس JWT موجود
	smsService  *SMSService  // سرویس SMS موجود
}

// NewUnifiedAuthService ایجاد نمونه جدید از سرویس احراز هویت یکپارچه
func NewUnifiedAuthService(db *gorm.DB, authService *AuthService, smsService *SMSService) *UnifiedAuthService {
	return &UnifiedAuthService{
		db:          db,
		authService: authService,
		smsService:  smsService,
	}
}

// getRoleIDByName گرفتن ID نقش بر اساس نام آن (مثلاً developer)
func (s *UnifiedAuthService) getRoleIDByName(name string) uint {
	type roleIDRow struct{ ID uint }
	var row roleIDRow
	_ = s.db.Table("roles").Select("id").Where("name = ?", name).Take(&row).Error
	return row.ID
}

// CheckMobileAndSendOTP بررسی موبایل و ارسال کد تایید
func (s *UnifiedAuthService) CheckMobileAndSendOTP(mobile, ip, userAgent string) (*models.AuthResponse, error) {
	// شناسایی ورودی (شماره موبایل یا dev)
	identifier := mobile

	// اگر ورودی dev است، اجازه فقط برای نقش توسعه‌دهنده
	if identifier == "dev" {
		var devUser models.User
		if err := s.db.Where("username = ?", "dev").First(&devUser).Error; err != nil {
			return &models.AuthResponse{Success: false, Message: "کاربر dev یافت نشد"}, nil
		}
		// احراز نقش توسعه‌دهنده به‌صورت داینامیک نه عدد ثابت
		// اجازه دسترسی dev اگر role_id=100 (developer) باشد
		devRoleID := s.getRoleIDByName("developer")
		if !((devRoleID != 0 && devUser.RoleID == devRoleID) || devUser.RoleID == 100) {
			return &models.AuthResponse{Success: false, Message: "دسترسی نامعتبر"}, nil
		}
		mobile = devUser.Mobile // موبایل واقعی برای ارسال OTP
	}

	// در صورتی که ورودی dev نباشد، باید شماره موبایل معتبر باشد
	if identifier != "dev" && !utils.IsValidMobile(identifier) {
		return &models.AuthResponse{Success: false, Message: "ورود فقط با شماره موبایل مجاز است"}, nil
	}

	// بررسی قفل بودن حساب
	if isLocked, reason := s.isAccountLocked(mobile); isLocked {
		return &models.AuthResponse{
			Success: false,
			Message: fmt.Sprintf("حساب شما قفل شده است. دلیل: %s", reason),
		}, nil
	}

	// بررسی محدودیت ارسال OTP
	if err := s.checkOTPRateLimit(mobile); err != nil {
		return &models.AuthResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	// بررسی وجود کاربر
	var user models.User
	var userExists bool
	if identifier == "dev" {
		userExists = s.db.Where("username = ?", identifier).First(&user).Error == nil
	} else {
		// جستجو با شماره موبایل یا username
		userExists = s.db.Where("mobile = ? OR username = ?", identifier, identifier).First(&user).Error == nil
	}

	// تولید و ارسال کد OTP
	_, err := s.generateAndSendOTP(mobile, "login", ip, userAgent)
	if err != nil {
		s.logLoginAttempt(mobile, "otp_request", false, err.Error(), ip, userAgent)
		return &models.AuthResponse{
			Success: false,
			Message: "خطا در ارسال کد تایید",
		}, err
	}

	s.logLoginAttempt(mobile, "otp_request", true, "", ip, userAgent)

	response := &models.AuthResponse{
		Success:    true,
		Message:    "کد تایید ارسال شد",
		RequireOTP: true,
		UserExists: userExists,
	}

	// اگر کاربر وجود دارد و پسورد دارد، گزینه ورود با پسورد را نمایش دهیم
	if userExists && user.PasswordHash != "" {
		response.RequirePass = true
		response.Message = "کد تایید ارسال شد. می‌توانید با کد تایید یا پسورد وارد شوید"
	}

	return response, nil
}

// SendOtp alias برای CheckMobileAndSendOTP (سازگاری با handler)
func (s *UnifiedAuthService) SendOtp(mobile, ip, userAgent string) (*models.AuthResponse, error) {
	return s.CheckMobileAndSendOTP(mobile, ip, userAgent)
}

// VerifyOTPAndLogin تایید کد OTP و ورود
func (s *UnifiedAuthService) VerifyOTPAndLogin(mobile, code, ip, userAgent string) (*models.AuthResponse, error) {
	identifier := mobile

	// اگر ورودی dev است، فقط نقش توسعه‌دهنده مجاز است
	if identifier == "dev" {
		var devUser models.User
		if err := s.db.Where("username = ?", "dev").First(&devUser).Error; err != nil {
			return &models.AuthResponse{Success: false, Message: "کاربر dev یافت نشد"}, nil
		}
		devRoleID := s.getRoleIDByName("developer")
		if devRoleID != 0 && devUser.RoleID != devRoleID {
			return &models.AuthResponse{Success: false, Message: "دسترسی نامعتبر"}, nil
		}
		mobile = devUser.Mobile
	}

	// در صورتی که ورودی dev نباشد، باید شماره موبایل معتبر باشد
	if identifier != "dev" && !utils.IsValidMobile(identifier) {
		return &models.AuthResponse{Success: false, Message: "ورود فقط با شماره موبایل مجاز است"}, nil
	}

	// بررسی قفل بودن حساب
	if isLocked, reason := s.isAccountLocked(mobile); isLocked {
		return &models.AuthResponse{
			Success: false,
			Message: fmt.Sprintf("حساب شما قفل شده است. دلیل: %s", reason),
		}, nil
	}

	// تایید کد OTP
	otpCode, err := s.verifyOTP(mobile, code, "login")
	if err != nil {
		s.logLoginAttempt(mobile, "otp_verify", false, err.Error(), ip, userAgent)
		s.handleFailedAttempt(mobile, ip, userAgent)
		return &models.AuthResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	// بررسی وجود کاربر
	var user models.User
	if identifier == "dev" {
		err = s.db.Where("username = ?", identifier).First(&user).Error
	} else {
		// جستجو با شماره موبایل یا username
		err = s.db.Where("mobile = ? OR username = ?", identifier, identifier).First(&user).Error
	}

	if err != nil {
		// کاربر وجود ندارد - ثبت‌نام خودکار
		user, err = s.createNewUser(mobile)
		if err != nil {
			s.logLoginAttempt(mobile, "auto_register", false, err.Error(), ip, userAgent)
			return &models.AuthResponse{
				Success: false,
				Message: "خطا در ثبت‌نام خودکار",
			}, err
		}
	}

	// علامت‌گذاری کد OTP به عنوان استفاده شده
	s.markOTPAsUsed(otpCode.ID)

	// دریافت تنظیمات JWT
	authSettings, err := s.authService.GetAuthSettings()
	if err != nil {
		s.logLoginAttempt(mobile, "auth_settings", false, err.Error(), ip, userAgent)
		return &models.AuthResponse{
			Success: false,
			Message: "خطا در دریافت تنظیمات احراز هویت",
		}, err
	}

	// تولید توکن‌های JWT (از سیستم موجود)
	accessToken, refreshToken, err := s.authService.GenerateTokens(&user, authSettings)
	if err != nil {
		s.logLoginAttempt(mobile, "token_generation", false, err.Error(), ip, userAgent)
		return &models.AuthResponse{
			Success: false,
			Message: "خطا در تولید توکن احراز هویت",
		}, err
	}

	// TODO: ایجاد session با device info (نیاز به gin.Context دارد)
	// فعلاً فقط لاگ می‌کنیم

	s.logLoginAttempt(mobile, "otp_verify", true, "", ip, userAgent)

	return &models.AuthResponse{
		Success:      true,
		Message:      "ورود موفقیت‌آمیز بود",
		Token:        accessToken,
		RefreshToken: refreshToken,
		User:         &user,
	}, nil
}

// LoginWithPassword ورود با پسورد
func (s *UnifiedAuthService) LoginWithPassword(identifier, password, ip, userAgent string) (*models.AuthResponse, error) {
	// محدودیت: فقط شماره موبایل یا dev
	if identifier != "dev" && !utils.IsValidMobile(identifier) {
		return &models.AuthResponse{Success: false, Message: "ورود فقط با شماره موبایل مجاز است"}, nil
	}

	// اگر dev باشد، اطمینان از نقش توسعه‌دهنده
	if identifier == "dev" {
		var devUser models.User
		if err := s.db.Where("username = ?", "dev").First(&devUser).Error; err != nil {
			return &models.AuthResponse{Success: false, Message: "کاربر dev یافت نشد"}, nil
		}
		devRoleID := s.getRoleIDByName("developer")
		if devRoleID != 0 && devUser.RoleID != devRoleID {
			return &models.AuthResponse{Success: false, Message: "دسترسی نامعتبر"}, nil
		}
	}

	mobile := identifier
	// بررسی قفل بودن حساب
	if isLocked, reason := s.isAccountLocked(identifier); isLocked {
		return &models.AuthResponse{
			Success: false,
			Message: fmt.Sprintf("حساب شما قفل شده است. دلیل: %s", reason),
		}, nil
	}

	// بررسی وجود کاربر
	var user models.User
	err := s.db.Where("mobile = ? OR username = ?", identifier, identifier).First(&user).Error
	if err != nil {
		s.logLoginAttempt(mobile, "password", false, "کاربر یافت نشد", ip, userAgent)
		s.handleFailedAttempt(mobile, ip, userAgent)
		return &models.AuthResponse{
			Success: false,
			Message: "شماره موبایل یا پسورد اشتباه است",
		}, nil
	}

	// بررسی پسورد
	if !utils.CheckPassword(password, user.PasswordHash) {
		s.logLoginAttempt(mobile, "password", false, "پسورد اشتباه", ip, userAgent)
		s.handleFailedAttempt(mobile, ip, userAgent)
		return &models.AuthResponse{
			Success: false,
			Message: "شماره موبایل یا پسورد اشتباه است",
		}, nil
	}

	// دریافت تنظیمات JWT
	authSettings, err := s.authService.GetAuthSettings()
	if err != nil {
		s.logLoginAttempt(mobile, "auth_settings", false, err.Error(), ip, userAgent)
		return &models.AuthResponse{
			Success: false,
			Message: "خطا در دریافت تنظیمات احراز هویت",
		}, err
	}

	// تولید توکن‌های JWT (از سیستم موجود)
	accessToken, refreshToken, err := s.authService.GenerateTokens(&user, authSettings)
	if err != nil {
		s.logLoginAttempt(mobile, "token_generation", false, err.Error(), ip, userAgent)
		return &models.AuthResponse{
			Success: false,
			Message: "خطا در تولید توکن احراز هویت",
		}, err
	}

	// به‌روزرسانی اطلاعات ورود کاربر
	s.updateUserLoginInfo(&user, ip)

	s.logLoginAttempt(mobile, "password", true, "", ip, userAgent)

	return &models.AuthResponse{
		Success:      true,
		Message:      "ورود موفقیت‌آمیز بود",
		Token:        accessToken,
		RefreshToken: refreshToken,
		User:         &user,
	}, nil
}

// generateAndSendOTP تولید و ارسال کد OTP
func (s *UnifiedAuthService) generateAndSendOTP(mobile, purpose, ip, userAgent string) (*models.OTPCode, error) {
	// دریافت تنظیمات OTP
	settings, err := s.getAuthSettings()
	if err != nil {
		return nil, err
	}

	// تولید کد تصادفی
	var codeLength int
	if codeLengthValue, exists := settings["auth.otp.length"]; exists && codeLengthValue != nil {
		if codeLengthInt, ok := codeLengthValue.(int); ok {
			codeLength = codeLengthInt
		}
	}
	if codeLength == 0 {
		codeLength = 5
	}

	code, err := s.generateRandomCode(codeLength)
	if err != nil {
		return nil, err
	}

	// محاسبه زمان انقضا
	var expiryMinutes int
	if expiryMinutesValue, exists := settings["auth.otp.expiry_minutes"]; exists && expiryMinutesValue != nil {
		if expiryMinutesInt, ok := expiryMinutesValue.(int); ok {
			expiryMinutes = expiryMinutesInt
		}
	}
	if expiryMinutes == 0 {
		expiryMinutes = 5
	}
	expiresAt := time.Now().Add(time.Duration(expiryMinutes) * time.Minute)

	// ذخیره کد در دیتابیس
	otpCode := models.OTPCode{
		Mobile:    mobile,
		Code:      code,
		Purpose:   purpose,
		ExpiresAt: expiresAt,
		IPAddress: ip,
		UserAgent: userAgent,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.db.Create(&otpCode).Error; err != nil {
		return nil, err
	}

	// ارسال SMS
	err = s.sendOTPSMS(mobile, code)
	if err != nil {
		return nil, err
	}

	return &otpCode, nil
}

// verifyOTP تایید کد OTP
func (s *UnifiedAuthService) verifyOTP(mobile, code, purpose string) (*models.OTPCode, error) {
	var otpCode models.OTPCode

	// جستجوی کد معتبر
	err := s.db.Where(
		"mobile = ? AND code = ? AND purpose = ? AND is_used = ? AND expires_at > ?",
		mobile, code, purpose, false, time.Now(),
	).First(&otpCode).Error

	if err != nil {
		return nil, fmt.Errorf("کد تایید نامعتبر یا منقضی شده است")
	}

	return &otpCode, nil
}

// markOTPAsUsed علامت‌گذاری کد OTP به عنوان استفاده شده
func (s *UnifiedAuthService) markOTPAsUsed(otpID uint) error {
	now := time.Now()
	return s.db.Model(&models.OTPCode{}).Where("id = ?", otpID).Updates(map[string]interface{}{
		"is_used": true,
		"used_at": &now,
	}).Error
}

// createNewUser ایجاد کاربر جدید (ثبت‌نام خودکار)
func (s *UnifiedAuthService) createNewUser(mobile string) (models.User, error) {
	now := time.Now()
	user := models.User{
		Username:       mobile, // استفاده از موبایل به عنوان نام کاربری
		Mobile:         mobile,
		Name:           mobile, // نام پیش‌فرض (کاربر بعداً می‌تواند تغییر دهد)
		RoleID:         1,      // customer role (default)
		Status:         "active",
		MobileVerified: true, // چون با OTP تایید شده
		RegisteredAt:   now,
		LastLoginAt:    &now,
		UpdatedAt:      now,
	}

	// تولید ایمیل فیک (یا خالی) چون ایمیل در مدل required است
	user.Email = fmt.Sprintf("%s@temp.local", mobile)

	if err := s.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// updateUserLoginInfo به‌روزرسانی اطلاعات ورود کاربر
func (s *UnifiedAuthService) updateUserLoginInfo(user *models.User, ip string) {
	now := time.Now()
	s.db.Model(user).Updates(map[string]interface{}{
		"last_login_at":      &now,
		"failed_login_count": 0, // بازنشانی شمارنده تلاش‌های ناموفق
	})
}

// generateRandomCode تولید کد تصادفی
func (s *UnifiedAuthService) generateRandomCode(length int) (string, error) {
	const digits = "0123456789"
	code := make([]byte, length)

	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		code[i] = digits[num.Int64()]
	}

	return string(code), nil
}

// sendOTPSMS ارسال SMS حاوی کد OTP
func (s *UnifiedAuthService) sendOTPSMS(mobile, code string) error {
	// الگوی استاندارد برای OTP را بر اساس اولویت درگاه‌ها پیدا می‌کنیم
	pattern, err := s.findPatternByPurpose("auth_otp")
	if err != nil {
		return err
	}

	// آماده‌سازی درخواست ارسال SMS با درگاه و کد پترن واقعی
	request := models.SMSSendRequest{
		PatternCode: pattern.PatternCode,
		Mobile:      mobile,
		GatewayID:   pattern.GatewayID,
		Scope:       pattern.Scope,   // اضافه کردن scope
		Feature:     pattern.Feature, // اضافه کردن feature
		Variables: map[string]string{
			"code": code,
		},
	}

	if _, err := s.smsService.SendSMS(request); err != nil {
		return fmt.Errorf("خطا در ارسال SMS: %v", err)
	}

	s.logSMSPatternUsage(pattern.ID, mobile, "auth_otp", map[string]string{"code": code})
	return nil
}

// findPatternByPurpose یافتن پترن مناسب بر اساس هدف/کاربرد (بدون نیاز به تنظیم دستی کد پترن)
// این تابع با توجه به اولویت گیت‌وی‌های فعال، اولین پترن فعال منطبق را برمی‌گرداند
func (s *UnifiedAuthService) findPatternByPurpose(purpose string) (*models.SMSPattern, error) {
	// گیت‌وی‌های فعال به ترتیب اولویت
	var gateways []models.SMSGateway
	if err := s.db.Where("is_active = ?", true).Order("priority ASC").Find(&gateways).Error; err != nil {
		return nil, fmt.Errorf("خطا در دریافت درگاه‌ها: %v", err)
	}
	if len(gateways) == 0 {
		return nil, fmt.Errorf("هیچ درگاه فعالی یافت نشد")
	}

	// در هر درگاه، به دنبال پترن با feature و وضعیت active بگرد
	for _, gw := range gateways {
		var pattern models.SMSPattern
		if err := s.db.Where("gateway_id = ? AND status = 'active' AND feature = ?", gw.ID, purpose).First(&pattern).Error; err == nil {
			// اگر pattern_code خالی باشد، قابل استفاده نیست
			if pattern.PatternCode == "" {
				continue
			}
			return &pattern, nil
		}
	}

	return nil, fmt.Errorf("هیچ پترن فعالی برای هدف %s یافت نشد", purpose)
}

// logSMSPatternUsage ثبت استفاده از پترن SMS
func (s *UnifiedAuthService) logSMSPatternUsage(patternID uint, mobile, purpose string, variables map[string]string) {
	// حذف شد - لاگ استفاده از پترن‌ها از طریق سرویس SMS انجام می‌شود
}

// دریافت تنظیمات احراز هویت
func (s *UnifiedAuthService) getAuthSettings() (map[string]interface{}, error) {
	// استفاده از جدول settings موجود
	var settings []models.Setting
	if err := s.db.Where("category = ? AND is_public = ?", "auth", false).Find(&settings).Error; err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	for _, setting := range settings {
		switch setting.Key {
		case "auth.otp.length", "auth.otp.expiry_minutes", "auth.otp.resend_cooldown", "auth.max_login_attempts", "auth.lockout_duration_minutes", "auth.session_timeout_minutes":
			result[setting.Key] = utils.ParseInt(setting.Value)
		case "auth.otp.enabled":
			result[setting.Key] = setting.Value == "true"
		default:
			result[setting.Key] = setting.Value
		}
	}

	return result, nil
}

// سایر متدهای کمکی...

// isAccountLocked بررسی قفل بودن حساب
func (s *UnifiedAuthService) isAccountLocked(mobile string) (bool, string) {
	var user models.User
	err := s.db.Where("mobile = ?", mobile).First(&user).Error
	if err != nil {
		return false, ""
	}

	// بررسی قفل موقت
	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return true, "حساب شما به دلیل تلاش‌های ناموفق متعدد قفل شده است"
	}

	// بررسی وضعیت حساب
	if user.Status == "blocked" || user.Status == "suspended" {
		reason := "حساب شما مسدود شده است"
		return true, reason
	}

	return false, ""
}

// checkOTPRateLimit بررسی محدودیت ارسال OTP
func (s *UnifiedAuthService) checkOTPRateLimit(mobile string) error {
	settings, _ := s.getAuthSettings()

	// بررسی وجود تنظیم cooldown و استفاده از مقدار پیش‌فرض در صورت عدم وجود
	var cooldown int
	if cooldownValue, exists := settings["auth.otp.resend_cooldown"]; exists && cooldownValue != nil {
		if cooldownInt, ok := cooldownValue.(int); ok {
			cooldown = cooldownInt
		}
	}
	if cooldown == 0 {
		cooldown = 60 // مقدار پیش‌فرض 60 ثانیه
	}

	var lastOTP models.OTPCode
	err := s.db.Where("mobile = ?", mobile).Order("created_at DESC").First(&lastOTP).Error
	if err == nil {
		if time.Since(lastOTP.CreatedAt) < time.Duration(cooldown)*time.Second {
			remaining := cooldown - int(time.Since(lastOTP.CreatedAt).Seconds())
			return fmt.Errorf("لطفاً %d ثانیه صبر کنید تا بتوانید کد جدید درخواست کنید", remaining)
		}
	}

	return nil
}

// logLoginAttempt ثبت تلاش ورود (استفاده از AuthEvent)
func (s *UnifiedAuthService) logLoginAttempt(mobile, attemptType string, isSuccessful bool, failureReason, ip, userAgent string) {
	s.authService.LogLoginAttempt(0, mobile, "", isSuccessful, attemptType, failureReason, ip, userAgent)
}

// handleFailedAttempt مدیریت تلاش‌های ناموفق
func (s *UnifiedAuthService) handleFailedAttempt(mobile, ip, userAgent string) {
	settings, _ := s.getAuthSettings()

	// بررسی وجود تنظیم max_attempts و استفاده از مقدار پیش‌فرض در صورت عدم وجود
	var maxAttempts int
	if maxAttemptsValue, exists := settings["auth.max_login_attempts"]; exists && maxAttemptsValue != nil {
		if maxAttemptsInt, ok := maxAttemptsValue.(int); ok {
			maxAttempts = maxAttemptsInt
		}
	}
	if maxAttempts == 0 {
		maxAttempts = 5 // مقدار پیش‌فرض 5 تلاش
	}

	// شمارش تلاش‌های ناموفق در 24 ساعت گذشته
	var failedCount int64
	s.db.Model(&models.LoginAttempt{}).Where(
		"mobile = ? AND is_successful = ? AND created_at > ?",
		mobile, false, time.Now().Add(-24*time.Hour),
	).Count(&failedCount)

	if failedCount >= int64(maxAttempts) {
		// قفل کردن حساب
		var lockDuration int
		if lockDurationValue, exists := settings["auth.lockout_duration_minutes"]; exists && lockDurationValue != nil {
			if lockDurationInt, ok := lockDurationValue.(int); ok {
				lockDuration = lockDurationInt
			}
		}
		if lockDuration == 0 {
			lockDuration = 15 // مقدار پیش‌فرض 15 دقیقه
		}

		lockedUntil := time.Now().Add(time.Duration(lockDuration) * time.Minute)

		// بروزرسانی اطلاعات کاربر
		var user models.User
		if err := s.db.Where("mobile = ?", mobile).First(&user).Error; err == nil {
			user.LockedUntil = &lockedUntil
			user.FailedLoginCount = int(failedCount) + 1
			s.db.Save(&user)
		}
	}
}
