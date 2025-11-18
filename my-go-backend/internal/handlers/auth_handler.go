package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
)

type AuthHandler struct {
	db              *gorm.DB
	authService     *services.AuthService
	settingsService *services.SettingsService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		db:              db,
		authService:     services.NewAuthService(db),
		settingsService: services.NewSettingsService(db),
	}
}

// SendOtp ارسال کد تایید به شماره موبایل
func (h *AuthHandler) SendOtp(c *gin.Context) {
	var req models.SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "داده‌های ورودی نامعتبر است"})
		return
	}

	// بررسی تنظیمات OTP
	settings, err := h.settingsService.GetAuthSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت تنظیمات"})
		return
	}

	if !settings.MobileAuthEnabled {
		c.JSON(http.StatusBadRequest, gin.H{"error": "احراز هویت با موبایل غیرفعال است"})
		return
	}

	// بررسی محدودیت ارسال مجدد
	lastOtp, err := h.authService.GetLastOtp(req.Mobile, "login")
	if err == nil && lastOtp != nil {
		cooldown := time.Duration(settings.OtpResendCooldown) * time.Second
		if time.Since(lastOtp.CreatedAt) < cooldown {
			remaining := cooldown - time.Since(lastOtp.CreatedAt)
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": fmt.Sprintf("لطفاً %d ثانیه دیگر تلاش کنید", int(remaining.Seconds())),
			})
			return
		}
	}

	// تولید کد OTP
	code := h.generateOtpCode(settings.OtpLength)
	expiresAt := time.Now().Add(time.Duration(settings.OtpExpiryMinutes) * time.Minute)

	// ذخیره کد در دیتابیس
	otp := &models.OTPCode{
		Mobile:    req.Mobile,
		Code:      code,
		Purpose:   "login", // یا req.Type اگر در مدل SendOTPRequest وجود داشته باشد
		ExpiresAt: expiresAt,
		IPAddress: c.ClientIP(),
		UserAgent: c.GetHeader("User-Agent"),
	}

	if err := h.db.Create(otp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ذخیره کد تایید"})
		return
	}

	// ارسال پیامک (در اینجا فقط لاگ می‌کنیم)
	fmt.Printf("کد تایید %s برای شماره %s ارسال شد\n", code, req.Mobile)

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"message":    "کد تایید ارسال شد",
		"expires_in": settings.OtpExpiryMinutes * 60,
	})
}

// VerifyOtp تایید کد OTP
func (h *AuthHandler) VerifyOtp(c *gin.Context) {
	var req models.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "داده‌های ورودی نامعتبر است"})
		return
	}

	// بررسی کد OTP
	_, err := h.authService.VerifyOtp(req.Mobile, req.Code, "login")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// بررسی یا ایجاد کاربر
	user, err := h.authService.GetOrCreateUser(req.Mobile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ایجاد کاربر"})
		return
	}

	// بررسی بلاک بودن کاربر
	if user.Status == "blocked" || user.Status == "suspended" {
		c.JSON(http.StatusForbidden, gin.H{"error": "حساب کاربری شما بلاک شده است"})
		return
	}

	// تولید توکن JWT
	settings, _ := h.settingsService.GetAuthSettings()
	token, refreshToken, err := h.authService.GenerateTokens(user, settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در تولید توکن"})
		return
	}

	if err := h.authService.CreateSessionWithDeviceInfo(user, token, refreshToken, settings, c, "otp"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ایجاد نشست"})
		return
	}

	// تنظیم کوکی‌های HttpOnly
	h.setAuthCookies(c, token, refreshToken, settings)

	// ثبت تلاش ورود موفق
	h.authService.LogLoginAttempt(user.ID, req.Mobile, "", true, "otp", "", c.ClientIP(), c.GetHeader("User-Agent"))

	// بروزرسانی آخرین ورود
	h.db.Model(user).Updates(map[string]interface{}{
		"last_login_at":      time.Now(),
		"failed_login_count": 0,
		"locked_until":       nil,
	})

	c.JSON(http.StatusOK, models.AuthResponse{
		Success:      true,
		Message:      "ورود موفقیت‌آمیز",
		User:         user,
		Token:        "", // توکن در کوکی قرار می‌گیرد
		RefreshToken: "", // توکن در کوکی قرار می‌گیرد
	})
}

// Login ورود با یوزرنیم و پسورد
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "داده‌های ورودی نامعتبر است"})
		return
	}

	// بررسی تنظیمات
	settings, err := h.settingsService.GetAuthSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت تنظیمات"})
		return
	}

	// اضافه کردن لاگ برای دیباگ تنظیمات
	fmt.Printf("🔧 Auth settings - UsernameAuthEnabled: %v, MobileAuthEnabled: %v\n", settings.UsernameAuthEnabled, settings.MobileAuthEnabled)

	if !settings.UsernameAuthEnabled {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ورود با یوزرنیم غیرفعال است"})
		return
	}

	// پیدا کردن کاربر
	var user models.User
	identifier := req.Identifier
	if identifier == "" {
		identifier = req.Username // fallback برای backward compatibility
	}
	query := h.db.Where("username = ? OR mobile = ?", identifier, identifier)
	if err := query.First(&user).Error; err != nil {
		h.logFailedLoginAttempt(nil, identifier, identifier, "password", "کاربر یافت نشد", c)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "یوزرنیم یا پسورد اشتباه است"})
		return
	}

	// بررسی بلاک بودن کاربر
	if user.Status == "blocked" || user.Status == "suspended" {
		c.JSON(http.StatusForbidden, gin.H{"error": "حساب کاربری شما بلاک شده است"})
		return
	}

	// بررسی قفل حساب
	if user.LockedUntil != nil && time.Now().Before(*user.LockedUntil) {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "حساب شما قفل شده است"})
		return
	}

	// بررسی پسورد
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		// افزایش تعداد تلاش‌های ناموفق
		failedCount := user.FailedLoginCount + 1
		updates := map[string]interface{}{"failed_login_count": failedCount}

		// بررسی بلاک خودکار
		if failedCount >= settings.AutoBlockFailedLogins {
			blockUntil := time.Now().Add(time.Duration(settings.AutoBlockDurationHours) * time.Hour)
			updates["locked_until"] = blockUntil
		}

		h.db.Model(&user).Updates(updates)

		h.logFailedLoginAttempt(&user, user.Mobile, user.Username, "password", "پسورد اشتباه", c)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "یوزرنیم یا پسورد اشتباه است"})
		return
	}

	// تولید توکن
	token, refreshToken, err := h.authService.GenerateTokens(&user, settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در تولید توکن"})
		return
	}

	if err := h.authService.CreateSessionWithDeviceInfo(&user, token, refreshToken, settings, c, "password"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ایجاد نشست"})
		return
	}

	// تنظیم کوکی‌های HttpOnly
	h.setAuthCookies(c, token, refreshToken, settings)

	// اضافه کردن لاگ برای دیباگ
	fmt.Printf("🔐 Cookies set - User ID: %d\n", user.ID)

	// ثبت تلاش ورود موفق
	h.authService.LogLoginAttempt(user.ID, user.Mobile, user.Username, true, "password", "", c.ClientIP(), c.GetHeader("User-Agent"))

	// بروزرسانی آخرین ورود
	h.db.Model(&user).Updates(map[string]interface{}{
		"last_login_at":      time.Now(),
		"failed_login_count": 0,
		"locked_until":       nil,
	})

	// دریافت نام نقش برای response
	var roleName string
	if user.RoleID > 0 {
		var r models.Role
		if err := h.db.First(&r, user.RoleID).Error; err == nil {
			roleName = r.Name
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"user":          user,
		"role_id":       user.RoleID,
		"role_name":     roleName,
		"access_token":  token,
		"expires_in":    int64(settings.JwtExpiryHours * 3600),
	})
}

// Register ثبت‌نام کاربر جدید
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "داده‌های ورودی نامعتبر است"})
		return
	}

	// بررسی تنظیمات
	settings, err := h.settingsService.GetAuthSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت تنظیمات"})
		return
	}

	// بررسی وجود کاربر
	var existingUser models.User
	if err := h.db.Where("mobile = ? OR email = ?", req.Mobile, req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "کاربر با این شماره موبایل یا ایمیل قبلاً ثبت‌نام کرده است"})
		return
	}

	// ایجاد کاربر جدید
	user := &models.User{
		Name:         req.Name,
		Mobile:       req.Mobile,
		Email:        req.Email,
		Username:     req.Mobile, // استفاده از موبایل به عنوان username
		RoleID:       1,          // default customer role
		Status:       "active",
		RegisteredAt: time.Now(),
		UpdatedAt:    time.Now(),
	}

	// هش کردن پسورد (اگر ارائه شده)
	if req.Password != "" {
		if len(req.Password) < settings.MinPasswordLength {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("پسورد باید حداقل %d کاراکتر باشد", settings.MinPasswordLength)})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در رمزگذاری پسورد"})
			return
		}
		user.PasswordHash = string(hashedPassword)
	}

	if err := h.db.Create(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ایجاد کاربر"})
		return
	}

	// تولید توکن‌ها
	token, refreshToken, err := h.authService.GenerateTokens(user, settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در تولید توکن"})
		return
	}

	if err := h.authService.CreateSessionWithDeviceInfo(user, token, refreshToken, settings, c, "register"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ایجاد نشست"})
		return
	}

	// تنظیم کوکی‌های HttpOnly
	h.setAuthCookies(c, token, refreshToken, settings)

	// ثبت فعالیت
	h.authService.LogUserActivity(user.ID, "register", "ثبت‌نام جدید", c.ClientIP(), c.GetHeader("User-Agent"))

	c.JSON(http.StatusCreated, models.AuthResponse{
		Success:      true,
		Message:      "ثبت‌نام موفقیت‌آمیز",
		User:         user,
		Token:        "", // توکن در کوکی قرار می‌گیرد
		RefreshToken: "", // توکن در کوکی قرار می‌گیرد
	})
}

// RefreshToken تازه‌سازی توکن دسترسی
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// دریافت refresh token از کوکی
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "توکن تازه‌سازی یافت نشد"})
		return
	}

	// تایید refresh token
	refreshTokenModel, err := h.authService.ValidateRefreshToken(refreshToken)
	if err != nil {
		// پاک کردن کوکی‌های نامعتبر
		h.clearAuthCookies(c)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "توکن تازه‌سازی نامعتبر است"})
		return
	}

	// دریافت کاربر
	var user models.User
	if err := h.db.First(&user, refreshTokenModel.UserID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر یافت نشد"})
		return
	}

	// بررسی بلاک بودن کاربر
	if user.Status == "blocked" || user.Status == "suspended" {
		h.clearAuthCookies(c)
		c.JSON(http.StatusForbidden, gin.H{"error": "حساب کاربری شما بلاک شده است"})
		return
	}

	// تولید توکن‌های جدید
	settings, _ := h.settingsService.GetAuthSettings()
	newToken, newRefreshToken, err := h.authService.GenerateTokens(&user, settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در تولید توکن"})
		return
	}

	if err := h.authService.RotateSessionTokens(refreshTokenModel, newToken, newRefreshToken, settings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بروزرسانی نشست"})
		return
	}

	// تنظیم کوکی‌های جدید
	h.setAuthCookies(c, newToken, newRefreshToken, settings)

	c.JSON(http.StatusOK, gin.H{
		"message":    "توکن با موفقیت تازه‌سازی شد",
		"expires_in": int64(settings.JwtExpiryHours * 3600),
	})
}

// Me دریافت اطلاعات کاربر فعلی
func (h *AuthHandler) Me(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}

	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "کاربر یافت نشد"})
		return
	}

	// دریافت نام نقش برای response
	var roleName string
	if user.RoleID > 0 {
		var r models.Role
		if err := h.db.First(&r, user.RoleID).Error; err == nil {
			roleName = r.Name
		}
	}

	// اضافه کردن role به user object برای سازگاری با frontend
	userResponse := gin.H{
		"id":       user.ID,
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
		"mobile":   user.Mobile,
		"role_id":  user.RoleID,
		"role":     roleName,
		"status":   user.Status,
	}

	c.JSON(http.StatusOK, gin.H{"authenticated": true, "user": userResponse, "role_id": user.RoleID, "role_name": roleName, "role": roleName})
}

// Logout خروج کاربر
func (h *AuthHandler) Logout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if exists {
		if sessionToken, hasSession := c.Get("session_token"); hasSession {
			if tokenStr, ok := sessionToken.(string); ok && tokenStr != "" {
				if err := h.authService.DeactivateSessionByToken(tokenStr); err != nil {
					fmt.Printf("failed to deactivate session: %v\n", err)
				}
			}
		} else {
			h.authService.RevokeAllUserTokens(userID.(uint))
		}

		h.authService.LogUserActivity(userID.(uint), "logout", "خروج از سیستم", c.ClientIP(), c.GetHeader("User-Agent"))
	}

	// پاک کردن کوکی‌ها
	h.clearAuthCookies(c)

	c.JSON(http.StatusOK, gin.H{"message": "با موفقیت خارج شدید"})
}

// setAuthCookies تنظیم کوکی‌های احراز هویت
func (h *AuthHandler) setAuthCookies(c *gin.Context, accessToken, refreshToken string, settings *models.AuthSettings) {
	// تنظیم SameSite mode برای امنیت و جلوگیری از CSRF
	c.SetSameSite(http.SameSiteLaxMode)

	// کوکی Access Token
	c.SetCookie(
		"access_token",
		accessToken,
		int((time.Duration(settings.JwtExpiryHours) * time.Hour).Seconds()),
		"/",
		"",
		false, // Secure = false برای localhost
		true,  // HttpOnly = true
	)

	// کوکی Refresh Token
	c.SetCookie(
		"refresh_token",
		refreshToken,
		int((time.Duration(settings.RefreshTokenExpiryDays) * 24 * time.Hour).Seconds()),
		"/",
		"",
		false, // Secure = false برای localhost
		true,  // HttpOnly = true
	)
}

// clearAuthCookies پاک کردن کوکی‌های احراز هویت
func (h *AuthHandler) clearAuthCookies(c *gin.Context) {
	// تنظیم SameSite برای consistency
	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
}

// generateOtpCode تولید کد OTP تصادفی
func (h *AuthHandler) generateOtpCode(length int) string {
	digits := "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = digits[rand.Intn(len(digits))]
	}
	return string(code)
}

// logFailedLoginAttempt ثبت تلاش ناموفق ورود با جزئیات دستگاه
func (h *AuthHandler) logFailedLoginAttempt(user *models.User, mobile, username, attemptType, failureReason string, c *gin.Context) {
	// Parse User-Agent and get device info
	userAgent := c.GetHeader("User-Agent")
	deviceInfo := utils.ParseUserAgent(userAgent)

	// Get client IP
	clientIP := c.ClientIP()

	// Determine device type
	var deviceType string
	if deviceInfo.IsMobile {
		deviceType = "Mobile"
	} else if deviceInfo.IsTablet {
		deviceType = "Tablet"
	} else {
		deviceType = "Desktop"
	}

	// Create enhanced metadata with device info
	metadata := map[string]interface{}{
		"mobile":         mobile,
		"username":       username,
		"method":         attemptType,
		"failure_reason": failureReason,
		"device_info": map[string]interface{}{
			"browser":         deviceInfo.Browser,
			"browser_version": deviceInfo.BrowserVersion,
			"os":              deviceInfo.OSName,
			"os_version":      deviceInfo.OSVersion,
			"device_type":     deviceType,
			"device_model":    deviceInfo.DeviceName,
			"platform":        deviceInfo.Platform,
		},
	}

	// Log to AuthEvent with enhanced metadata
	var userID uint
	if user != nil {
		userID = user.ID
	}
	h.authService.LogLoginAttempt(userID, mobile, username, false, attemptType, failureReason, clientIP, userAgent)

	// Also create a separate detailed failed login record in auth_events
	metadataJSON, _ := json.Marshal(metadata)

	event := &models.AuthEvent{
		UserID: func() *uint {
			if userID > 0 {
				return &userID
			} else {
				return nil
			}
		}(),
		Mobile:    mobile,
		EventType: "failed_login_attempt",
		Status:    "failed",
		Severity:  "warning",
		IPAddress: clientIP,
		UserAgent: userAgent,
		Metadata:  datatypes.JSON(metadataJSON),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	h.db.Create(event)

	// Cleanup old failed attempts (keep only last 100 per user)
	h.cleanupOldFailedAttempts(userID)
}

// cleanupOldFailedAttempts حذف رکوردهای قدیمی تلاش‌های ناموفق
func (h *AuthHandler) cleanupOldFailedAttempts(userID uint) {
	// حذف رکوردهای قدیمی‌تر از 100 رکورد برای کاربر
	var count int64
	h.db.Model(&models.AuthEvent{}).Where("user_id = ? AND event_type = 'failed_login_attempt'", userID).Count(&count)

	if count > 100 {
		// پیدا کردن ID رکوردهای قدیمی
		var oldRecords []struct {
			ID uint `json:"id"`
		}
		h.db.Model(&models.AuthEvent{}).
			Select("id").
			Where("user_id = ? AND event_type = 'failed_login_attempt'", userID).
			Order("created_at ASC").
			Limit(int(count - 100)).
			Find(&oldRecords)

		// حذف رکوردهای قدیمی
		if len(oldRecords) > 0 {
			ids := make([]uint, len(oldRecords))
			for i, record := range oldRecords {
				ids[i] = record.ID
			}
			h.db.Where("id IN ?", ids).Delete(&models.AuthEvent{})
		}
	}
}

// CheckUser بررسی وجود کاربر با شماره موبایل
func (h *AuthHandler) CheckUser(c *gin.Context) {
	var req struct {
		Mobile string `json:"mobile" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شماره موبایل الزامی است"})
		return
	}

	// بررسی وجود کاربر
	var user models.User
	userExists := h.db.Where("mobile = ?", req.Mobile).First(&user).Error == nil

	response := gin.H{
		"exists":      userExists,
		"hasPassword": false,
		"isActive":    true,
		"role":        "user",
		"registered":  userExists,
	}

	if userExists {
		response["hasPassword"] = user.PasswordHash != ""
		response["isActive"] = user.Status == "active"
		response["role_id"] = user.RoleID
	}

	c.JSON(http.StatusOK, response)
}

// FixAuthSettings موقت - برای دیباگ تنظیمات
func (h *AuthHandler) FixAuthSettings(c *gin.Context) {
	settings, err := h.settingsService.GetAuthSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت تنظیمات"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"settings": settings,
		"message":  "تنظیمات احراز هویت",
	})
}
