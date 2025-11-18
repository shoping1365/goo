package services

import (
	"strconv"

	"gorm.io/gorm"

	"my-go-backend/internal/models"
)

type SettingsService struct {
	db *gorm.DB
}

func NewSettingsService(db *gorm.DB) *SettingsService {
	return &SettingsService{db: db}
}

// GetAuthSettings دریافت تنظیمات احراز هویت
func (s *SettingsService) GetAuthSettings() (*models.AuthSettings, error) {
	var settings []models.Setting
	err := s.db.Where("category = ?", "auth").Find(&settings).Error
	if err != nil {
		return nil, err
	}

	authSettings := &models.AuthSettings{}

	for _, setting := range settings {
		switch setting.Key {
		case "mobile_auth_enabled":
			authSettings.MobileAuthEnabled = setting.Value == "true"
		case "otp_length":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.OtpLength = val
			}
		case "otp_expiry_minutes":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.OtpExpiryMinutes = val
			}
		case "max_otp_attempts":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.MaxOtpAttempts = val
			}
		case "otp_resend_cooldown":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.OtpResendCooldown = val
			}
		case "jwt_expiry_hours":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.JwtExpiryHours = val
			}
		case "refresh_token_expiry_days":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.RefreshTokenExpiryDays = val
			}
		case "jwt_secret":
			authSettings.JwtSecret = setting.Value
		case "max_concurrent_sessions":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.MaxConcurrentSessions = val
			}
		case "username_auth_enabled":
			authSettings.UsernameAuthEnabled = setting.Value == "true"
		case "min_password_length":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.MinPasswordLength = val
			}
		case "max_login_attempts":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.MaxLoginAttempts = val
			}
		case "account_lockout_minutes":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.AccountLockoutMinutes = val
			}
		case "password_expiry_days":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.PasswordExpiryDays = val
			}
		case "two_factor_enabled":
			authSettings.TwoFactorEnabled = setting.Value == "true"
		case "suspicious_activity_detection":
			authSettings.SuspiciousActivityDetection = setting.Value == "true"
		case "session_timeout_minutes":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.SessionTimeoutMinutes = val
			}
		case "login_history_retention_days":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.LoginHistoryRetentionDays = val
			}
		case "auto_block_failed_logins":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.AutoBlockFailedLogins = val
			}
		case "auto_block_duration_hours":
			if val, err := strconv.Atoi(setting.Value); err == nil {
				authSettings.AutoBlockDurationHours = val
			}
		case "default_user_role":
			authSettings.DefaultUserRole = setting.Value
		case "email_verification_enabled":
			authSettings.EmailVerificationEnabled = setting.Value == "true"
		case "phone_verification_enabled":
			authSettings.PhoneVerificationEnabled = setting.Value == "true"
		}
	}

	// تنظیم مقادیر پیش‌فرض اگر خالی باشند
	if authSettings.OtpLength == 0 {
		authSettings.OtpLength = 5
	}
	if authSettings.OtpExpiryMinutes == 0 {
		authSettings.OtpExpiryMinutes = 5
	}
	if authSettings.MaxOtpAttempts == 0 {
		authSettings.MaxOtpAttempts = 3
	}
	if authSettings.OtpResendCooldown == 0 {
		authSettings.OtpResendCooldown = 60
	}
	if authSettings.JwtExpiryHours == 0 {
		authSettings.JwtExpiryHours = 24
	}
	if authSettings.RefreshTokenExpiryDays == 0 {
		authSettings.RefreshTokenExpiryDays = 30
	}
	if authSettings.JwtSecret == "" {
		authSettings.JwtSecret = "your-super-secret-jwt-key-change-this-in-production"
	}
	if authSettings.MaxConcurrentSessions == 0 {
		authSettings.MaxConcurrentSessions = 5
	}
	if authSettings.MinPasswordLength == 0 {
		authSettings.MinPasswordLength = 8
	}
	if authSettings.MaxLoginAttempts == 0 {
		authSettings.MaxLoginAttempts = 5
	}
	if authSettings.AccountLockoutMinutes == 0 {
		authSettings.AccountLockoutMinutes = 15
	}
	if authSettings.PasswordExpiryDays == 0 {
		authSettings.PasswordExpiryDays = 90
	}
	if authSettings.SessionTimeoutMinutes == 0 {
		authSettings.SessionTimeoutMinutes = 30
	}
	if authSettings.LoginHistoryRetentionDays == 0 {
		authSettings.LoginHistoryRetentionDays = 365
	}
	if authSettings.AutoBlockFailedLogins == 0 {
		authSettings.AutoBlockFailedLogins = 10
	}
	if authSettings.AutoBlockDurationHours == 0 {
		authSettings.AutoBlockDurationHours = 24
	}
	if authSettings.DefaultUserRole == "" {
		authSettings.DefaultUserRole = "user"
	}

	return authSettings, nil
}
