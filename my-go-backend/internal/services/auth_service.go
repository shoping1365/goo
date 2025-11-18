package services

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

// GetLastOtp دریافت آخرین کد OTP ارسالی
func (s *AuthService) GetLastOtp(mobile, otpType string) (*models.OTPCode, error) {
	var otp models.OTPCode
	err := s.db.Where("mobile = ? AND purpose = ? AND is_used = false AND expires_at > ?",
		mobile, otpType, time.Now()).
		Order("created_at DESC").
		First(&otp).Error

	if err != nil {
		return nil, err
	}
	return &otp, nil
}

// VerifyOtp تایید کد OTP
func (s *AuthService) VerifyOtp(mobile, code, otpType string) (*models.OTPCode, error) {
	var otp models.OTPCode
	err := s.db.Where("mobile = ? AND code = ? AND purpose = ? AND is_used = false AND expires_at > ?",
		mobile, code, otpType, time.Now()).
		First(&otp).Error

	if err != nil {
		return nil, errors.New("کد تایید نامعتبر یا منقضی شده است")
	}

	// علامت‌گذاری به عنوان استفاده شده
	s.db.Model(&otp).Update("is_used", true)
	s.db.Model(&otp).Update("used_at", time.Now())

	return &otp, nil
}

// GetOrCreateUser دریافت یا ایجاد کاربر
func (s *AuthService) GetOrCreateUser(mobile string) (*models.User, error) {
	var user models.User
	err := s.db.Where("mobile = ?", mobile).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// ایجاد کاربر جدید
		user = models.User{
			Mobile:       mobile,
			RoleID:       1, // default customer role
			Status:       "active",
			RegisteredAt: time.Now(),
		}

		if err := s.db.Create(&user).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

// GenerateTokens تولید توکن JWT و توکن تازه‌سازی
func (s *AuthService) GenerateTokens(user *models.User, settings *models.AuthSettings) (string, string, error) {
	// اطمینان از اینکه مدت انقضا صفر یا منفی نیست
	if settings.JwtExpiryHours <= 0 {
		settings.JwtExpiryHours = 24
	}
	// تولید توکن JWT (Access Token)
	expiresAt := time.Now().Add(time.Duration(settings.JwtExpiryHours) * time.Hour)

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"mobile":  user.Mobile,
		"role_id": user.RoleID,
		"exp":     expiresAt.Unix(),
		"iat":     time.Now().Unix(),
		"type":    "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(settings.JwtSecret))
	if err != nil {
		return "", "", err
	}

	// تولید توکن تازه‌سازی (فقط string برمی‌گردانیم، session در CreateSessionWithDeviceInfo ذخیره میشه)
	refreshToken := s.generateRefreshToken()

	return tokenString, refreshToken, nil
}

// ValidateAccessToken تایید توکن دسترسی
func (s *AuthService) ValidateAccessToken(tokenString string) (*jwt.Token, error) {
	settings, _ := s.GetAuthSettings()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// بررسی الگوریتم
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("الگوریتم امضای نامعتبر")
		}
		return []byte(settings.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// بررسی نوع توکن
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if tokenType, exists := claims["type"]; !exists || tokenType != "access" {
			return nil, errors.New("نوع توکن نامعتبر")
		}
		return token, nil
	}

	return nil, errors.New("توکن نامعتبر")
}

// ValidateRefreshToken تایید توکن تازه‌سازی از UserSession
func (s *AuthService) ValidateRefreshToken(tokenString string) (*models.UserSession, error) {
	var session models.UserSession
	err := s.db.Where("refresh_token = ? AND is_active = true AND expires_at > ?",
		tokenString, time.Now()).
		First(&session).Error

	if err != nil {
		return nil, errors.New("توکن تازه‌سازی نامعتبر یا منقضی شده است")
	}

	return &session, nil
}

// RevokeRefreshToken باطل کردن توکن تازه‌سازی
func (s *AuthService) RevokeRefreshToken(tokenString string) error {
	return s.db.Model(&models.UserSession{}).
		Where("refresh_token = ?", tokenString).
		Update("is_active", false).Error
}

// RevokeAllUserTokens باطل کردن تمام session های کاربر
func (s *AuthService) RevokeAllUserTokens(userID uint) error {
	return s.db.Model(&models.UserSession{}).
		Where("user_id = ?", userID).
		Update("is_active", false).Error
}

// LogLoginAttempt ثبت تلاش ورود (استفاده از AuthEvent)
func (s *AuthService) LogLoginAttempt(userID uint, mobile, username string, success bool, method, reason, ipAddress, userAgent string) {
	status := "success"
	if !success {
		status = "failed"
	}

	// تولید metadata
	metadata := map[string]interface{}{
		"mobile":   mobile,
		"username": username,
		"method":   method,
	}
	if reason != "" {
		metadata["failure_reason"] = reason
	}

	metadataJSON, err := json.Marshal(metadata)
	if err != nil {
		// در صورت خطا در marshal، metadata خالی استفاده می‌کنیم
		metadataJSON = []byte("{}")
	}

	event := &models.AuthEvent{
		UserID: func() *uint {
			if userID > 0 {
				return &userID
			} else {
				return nil
			}
		}(),
		Mobile:    mobile,
		EventType: "login_attempt",
		Status:    status,
		Severity: func() string {
			if success {
				return "info"
			} else {
				return "warning"
			}
		}(),
		IPAddress: ipAddress,
		UserAgent: userAgent,
		Metadata:  datatypes.JSON(metadataJSON),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.db.Create(event).Error; err != nil {
		// لاگ کردن خطا در صورت عدم موفقیت در ثبت event
		fmt.Printf("Failed to create auth event: %v\n", err)
	}
}

// LogUserActivity ثبت فعالیت کاربر (استفاده از AuthEvent)
func (s *AuthService) LogUserActivity(userID uint, action, details, ipAddress, userAgent string) {
	activity := &models.AuthEvent{
		UserID:    &userID,
		EventType: action,
		Status:    "success",
		Severity:  "info",
		IPAddress: ipAddress,
		UserAgent: userAgent,
		Metadata:  datatypes.JSON(fmt.Sprintf(`{"details": "%s"}`, details)),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.db.Create(activity).Error; err != nil {
		// لاگ کردن خطا در صورت عدم موفقیت در ثبت activity
		fmt.Printf("Failed to create user activity: %v\n", err)
	}
}

// LogAuthEvent ثبت رخداد امنیتی با metadata انعطاف‌پذیر
func (s *AuthService) LogAuthEvent(userID uint, eventType string, metadata interface{}, ipAddress, userAgent string) {
	// تبدیل metadata به JSON
	var metadataJSON datatypes.JSON
	if metadata != nil {
		if jsonBytes, err := json.Marshal(metadata); err == nil {
			metadataJSON = datatypes.JSON(jsonBytes)
		} else {
			metadataJSON = datatypes.JSON(`{}`)
		}
	} else {
		metadataJSON = datatypes.JSON(`{}`)
	}

	event := &models.AuthEvent{
		UserID:    &userID,
		EventType: eventType,
		Status:    "success",
		Severity:  "info",
		IPAddress: ipAddress,
		UserAgent: userAgent,
		Metadata:  metadataJSON,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.db.Create(event).Error; err != nil {
		// لاگ کردن خطا در صورت عدم موفقیت در ثبت event
		fmt.Printf("Failed to create auth event: %v\n", err)
	}
}

// CreateSessionWithDeviceInfo ایجاد session کامل با device tracking
func (s *AuthService) CreateSessionWithDeviceInfo(user *models.User, sessionToken, refreshToken string, settings *models.AuthSettings, c *gin.Context, authMethod string) error {
	if user == nil {
		return errors.New("user context is required for session creation")
	}
	if settings == nil {
		return errors.New("auth settings are required for session creation")
	}

	expiryDays := settings.RefreshTokenExpiryDays
	if expiryDays <= 0 {
		expiryDays = 30
	}

	now := time.Now()
	sessionExpiry := now.Add(time.Duration(expiryDays) * 24 * time.Hour)

	// تشخیص دستگاه
	deviceInfo := utils.DetectDevice(c)
	fingerprint := utils.GetClientFingerprint(c)

	// تولید اطلاعات device info JSON
	deviceInfoJSON, err := json.Marshal(map[string]interface{}{
		"screen_resolution": c.GetHeader("X-Screen-Resolution"),
		"language":          c.GetHeader("Accept-Language"),
		"timezone":          c.GetHeader("X-Timezone"),
		"connection_type":   c.GetHeader("X-Connection-Type"),
	})
	if err != nil {
		deviceInfoJSON = []byte("{}")
	}

	// ایجاد session جدید
	session := &models.UserSession{
		UserID:            user.ID,
		SessionToken:      sessionToken,
		RefreshToken:      refreshToken,
		DeviceFingerprint: fingerprint,
		DeviceName:        deviceInfo.DeviceName,
		Platform:          deviceInfo.Platform,
		Browser:           deviceInfo.Browser,
		BrowserVersion:    deviceInfo.BrowserVersion,
		OSName:            deviceInfo.OSName,
		OSVersion:         deviceInfo.OSVersion,
		IPAddress:         c.ClientIP(),
		UserAgent:         c.GetHeader("User-Agent"),
		City:              "",
		Country:           "",
		Timezone:          c.GetHeader("X-Timezone"),
		DeviceInfo:        datatypes.JSON(deviceInfoJSON),
		AuthMethod:        authMethod,
		IsActive:          true,
		IsTrusted:         false,
		IsCurrent:         true,
		LoginAt:           now,
		LastActivity:      &now,
		ExpiresAt:         sessionExpiry,
	}

	if err := s.db.Model(&models.UserSession{}).
		Where("user_id = ? AND is_current = ?", user.ID, true).
		Update("is_current", false).Error; err != nil {
		fmt.Printf("Failed to deactivate previous sessions: %v\n", err)
	}

	return s.db.Create(session).Error
}

// RotateSessionTokens بروزرسانی session موجود با توکن‌های تازه صادرشده
func (s *AuthService) RotateSessionTokens(session *models.UserSession, newSessionToken, newRefreshToken string, settings *models.AuthSettings) error {
	if session == nil {
		return errors.New("session is required for rotation")
	}
	if settings == nil {
		return errors.New("auth settings are required for rotation")
	}

	expiryDays := settings.RefreshTokenExpiryDays
	if expiryDays <= 0 {
		expiryDays = 30
	}

	updates := map[string]interface{}{
		"session_token": newSessionToken,
		"refresh_token": newRefreshToken,
		"last_activity": time.Now(),
		"expires_at":    time.Now().Add(time.Duration(expiryDays) * 24 * time.Hour),
		"is_active":     true,
		"is_current":    true,
	}

	if err := s.db.Model(&models.UserSession{}).
		Where("id = ?", session.ID).
		Updates(updates).Error; err != nil {
		return err
	}

	return s.db.Model(&models.UserSession{}).
		Where("user_id = ? AND id <> ? AND is_current = ?", session.UserID, session.ID, true).
		Update("is_current", false).Error
}

// DeactivateSessionByToken غیرفعال‌سازی session بر اساس توکن دسترسی مرتبط
func (s *AuthService) DeactivateSessionByToken(sessionToken string) error {
	if sessionToken == "" {
		return errors.New("session token is required")
	}

	return s.db.Model(&models.UserSession{}).
		Where("session_token = ?", sessionToken).
		Updates(map[string]interface{}{
			"is_active":  false,
			"is_current": false,
		}).Error
}

// TouchSessionActivity بروزرسانی LastActivity برای session فعال
func (s *AuthService) TouchSessionActivity(sessionID uint) {
	if sessionID == 0 {
		return
	}

	_ = s.db.Model(&models.UserSession{}).
		Where("id = ?", sessionID).
		Updates(map[string]interface{}{
			"last_activity": time.Now(),
		}).Error
}

// generateRefreshToken تولید توکن تازه‌سازی تصادفی
func (s *AuthService) generateRefreshToken() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		// در صورت خطا، یک توکن پیش‌فرض برمی‌گردانیم
		return "default-refresh-token"
	}
	return hex.EncodeToString(bytes)
}

// GetAuthSettings دریافت تنظیمات احراز هویت از جدول settings
func (s *AuthService) GetAuthSettings() (*models.AuthSettings, error) {
	// مقادیر پیش‌فرض (اگر در DB موجود نباشد)
	settings := &models.AuthSettings{
		MobileAuthEnabled:      true,
		OtpLength:              5,
		OtpExpiryMinutes:       5,
		MaxOtpAttempts:         3,
		OtpResendCooldown:      60,
		JwtExpiryHours:         24,
		RefreshTokenExpiryDays: 30,
		JwtSecret: func() string {
			if secret := os.Getenv("JWT_SECRET"); secret != "" {
				return secret
			}
			return "your-super-secret-jwt-key-change-this-in-production"
		}(),
		MaxConcurrentSessions:       5,
		UsernameAuthEnabled:         true,
		MinPasswordLength:           8,
		MaxLoginAttempts:            5,
		AccountLockoutMinutes:       15,
		PasswordExpiryDays:          90,
		TwoFactorEnabled:            false,
		SuspiciousActivityDetection: true,
		SessionTimeoutMinutes:       30,
		LoginHistoryRetentionDays:   365,
		AutoBlockFailedLogins:       10,
		AutoBlockDurationHours:      24,
		DefaultUserRole:             "customer",
		EmailVerificationEnabled:    false,
		PhoneVerificationEnabled:    true,
	}

	// بارگذاری از DB
	var rows []models.Setting
	if err := s.db.Where("category = ?", "auth").Find(&rows).Error; err != nil {
		return settings, err
	}

	for _, row := range rows {
		switch row.Key {
		case "mobile_auth_enabled":
			settings.MobileAuthEnabled = row.Value == "true"
		case "otp_length":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.OtpLength = v
			}
		case "otp_expiry_minutes":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.OtpExpiryMinutes = v
			}
		case "max_otp_attempts":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.MaxOtpAttempts = v
			}
		case "otp_resend_cooldown":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.OtpResendCooldown = v
			}
		case "jwt_expiry_hours":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.JwtExpiryHours = v
			}
		case "refresh_token_expiry_days":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.RefreshTokenExpiryDays = v
			}
		case "jwt_secret":
			// JWT secret از متغیر محیطی اولویت دارد (استاندارد صنعت)
			// فقط اگر متغیر محیطی تنظیم نشده باشد، از دیتابیس استفاده می‌کنیم
			if os.Getenv("JWT_SECRET") == "" {
				settings.JwtSecret = row.Value
			}
		case "max_concurrent_sessions":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.MaxConcurrentSessions = v
			}
		case "username_auth_enabled":
			settings.UsernameAuthEnabled = row.Value == "true"
		case "min_password_length":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.MinPasswordLength = v
			}
		case "max_login_attempts":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.MaxLoginAttempts = v
			}
		case "account_lockout_minutes":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.AccountLockoutMinutes = v
			}
		case "password_expiry_days":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.PasswordExpiryDays = v
			}
		case "two_factor_enabled":
			settings.TwoFactorEnabled = row.Value == "true"
		case "suspicious_activity_detection":
			settings.SuspiciousActivityDetection = row.Value == "true"
		case "session_timeout_minutes":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.SessionTimeoutMinutes = v
			}
		case "login_history_retention_days":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.LoginHistoryRetentionDays = v
			}
		case "auto_block_failed_logins":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.AutoBlockFailedLogins = v
			}
		case "auto_block_duration_hours":
			if v, err := strconv.Atoi(row.Value); err == nil {
				settings.AutoBlockDurationHours = v
			}
		case "default_user_role":
			settings.DefaultUserRole = row.Value
		case "email_verification_enabled":
			settings.EmailVerificationEnabled = row.Value == "true"
		case "phone_verification_enabled":
			settings.PhoneVerificationEnabled = row.Value == "true"
		}
	}

	// اطمینان از اینکه مدت اعتبار توکن صفر یا منفی نباشد
	if settings.JwtExpiryHours <= 0 {
		settings.JwtExpiryHours = 24
	}

	return settings, nil
}
