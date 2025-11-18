package models

import (
	"time"

	"gorm.io/datatypes"
)

// حذف شد - از جدول settings موجود استفاده می‌کنیم

// AuthSettings - ساختار تنظیمات احراز هویت (برای JWT و سیستم موجود)
type AuthSettings struct {
	// تنظیمات OTP
	MobileAuthEnabled bool `json:"mobile_auth_enabled"`
	OtpLength         int  `json:"otp_length"`
	OtpExpiryMinutes  int  `json:"otp_expiry_minutes"`
	MaxOtpAttempts    int  `json:"max_otp_attempts"`
	OtpResendCooldown int  `json:"otp_resend_cooldown"`

	// تنظیمات JWT
	JwtExpiryHours         int    `json:"jwt_expiry_hours"`
	RefreshTokenExpiryDays int    `json:"refresh_token_expiry_days"`
	JwtSecret              string `json:"jwt_secret"`
	MaxConcurrentSessions  int    `json:"max_concurrent_sessions"`

	// تنظیمات ورود با یوزرنیم
	UsernameAuthEnabled   bool `json:"username_auth_enabled"`
	MinPasswordLength     int  `json:"min_password_length"`
	MaxLoginAttempts      int  `json:"max_login_attempts"`
	AccountLockoutMinutes int  `json:"account_lockout_minutes"`
	PasswordExpiryDays    int  `json:"password_expiry_days"`

	// تنظیمات امنیتی
	TwoFactorEnabled            bool `json:"two_factor_enabled"`
	SuspiciousActivityDetection bool `json:"suspicious_activity_detection"`
	SessionTimeoutMinutes       int  `json:"session_timeout_minutes"`
	LoginHistoryRetentionDays   int  `json:"login_history_retention_days"`
	AutoBlockFailedLogins       int  `json:"auto_block_failed_logins"`
	AutoBlockDurationHours      int  `json:"auto_block_duration_hours"`

	// تنظیمات نقش‌ها
	DefaultUserRole          string `json:"default_user_role"`
	EmailVerificationEnabled bool   `json:"email_verification_enabled"`
	PhoneVerificationEnabled bool   `json:"phone_verification_enabled"`
}

// OTPCode - کدهای تایید OTP
type OTPCode struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Mobile    string     `json:"mobile" gorm:"size:20;index;not null"` // شماره موبایل
	Code      string     `json:"-" gorm:"size:10;not null"`            // کد تایید (مخفی در JSON)
	Purpose   string     `json:"purpose" gorm:"size:50;not null"`      // هدف (login, register, password_reset)
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null;index"`     // زمان انقضا
	IsUsed    bool       `json:"is_used" gorm:"default:false;index"`   // استفاده شده؟
	UsedAt    *time.Time `json:"used_at" gorm:"index"`                 // زمان استفاده
	IPAddress string     `json:"ip_address" gorm:"size:45"`            // IP تلاش
	UserAgent string     `json:"user_agent" gorm:"size:500"`           // User Agent
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// حذف شد - از جدول sessions موجود که بروزرسانی شده استفاده می‌کنیم

// LoginAttempt - تلاش‌های ورود
type LoginAttempt struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Mobile        string    `json:"mobile" gorm:"size:20;index;not null"`     // شماره موبایل
	AttemptType   string    `json:"attempt_type" gorm:"size:50;not null"`     // نوع تلاش (otp_request, otp_verify, password)
	IsSuccessful  bool      `json:"is_successful" gorm:"default:false;index"` // موفق بوده؟
	FailureReason string    `json:"failure_reason" gorm:"size:255"`           // دلیل شکست
	AttemptIP     string    `json:"attempt_ip" gorm:"size:45;index"`          // IP تلاش
	UserAgent     string    `json:"user_agent" gorm:"size:500"`               // User Agent
	CreatedAt     time.Time `json:"created_at" gorm:"index"`                  // زمان تلاش
	UpdatedAt     time.Time `json:"updated_at"`                               // زمان آخرین ویرایش
}

// حذف شد - از فیلد locked_until در جدول users استفاده می‌کنیم

// حذف شد - استفاده از پترن‌ها از طریق سرویس SMS انجام می‌شود

// درخواست‌های API

// SendOTPRequest - درخواست ارسال کد تایید
type SendOTPRequest struct {
	Mobile string `json:"mobile" binding:"required,len=11" example:"09123456789"` // شماره موبایل
}

// VerifyOTPRequest - درخواست تایید کد
type VerifyOTPRequest struct {
	Mobile string `json:"mobile" binding:"required,len=11" example:"09123456789"` // شماره موبایل
	Code   string `json:"code" binding:"required,min=4,max=6" example:"12345"`    // کد تایید
}

// UnifiedLoginRequest - درخواست ورود یکپارچه
type UnifiedLoginRequest struct {
	Mobile     string `json:"mobile" binding:"required,len=11" example:"09123456789"` // شماره موبایل
	AuthMethod string `json:"auth_method" binding:"required" example:"otp"`           // روش احراز (otp, password)
	Password   string `json:"password,omitempty" example:"password123"`               // پسورد (در صورت انتخاب password)
	Code       string `json:"code,omitempty" example:"12345"`                         // کد تایید (در صورت انتخاب otp)
}

// LoginPasswordRequest - درخواست ورود با پسورد بدون نیاز به تعیین auth_method
// همهٔ کاربران (مشتری یا ادمین) می‌توانند از این مدل استفاده کنند.
type LoginPasswordRequest struct {
	Mobile   string `json:"identifier" binding:"required" example:"09123456789"` // شناسه ورود (موبایل یا username)
	Password string `json:"password" binding:"required" example:"password123"`   // پسورد
}

// AuthResponse - پاسخ احراز هویت
type AuthResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	RequireOTP   bool   `json:"require_otp,omitempty"`   // نیاز به کد تایید؟
	RequirePass  bool   `json:"require_pass,omitempty"`  // نیاز به پسورد؟
	UserExists   bool   `json:"user_exists,omitempty"`   // کاربر وجود دارد؟
	Token        string `json:"token,omitempty"`         // توکن احراز هویت
	RefreshToken string `json:"refresh_token,omitempty"` // توکن تجدید
	User         *User  `json:"user,omitempty"`          // اطلاعات کاربر
}

// LoginRequest - درخواست ورود
type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"` // یوزرنیم یا موبایل
	Username   string `json:"username,omitempty"`            // برای backward compatibility
	Password   string `json:"password" binding:"required"`
}

// RegisterRequest - درخواست ثبت‌نام
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Mobile   string `json:"mobile" binding:"required,len=11"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// SetPasswordRequest - درخواست تنظیم پسورد
type SetPasswordRequest struct {
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

// TableName مشخص کردن نام جدول‌ها
func (OTPCode) TableName() string {
	return "otp_codes"
}

func (LoginAttempt) TableName() string {
	return "login_attempts"
}

// RefreshToken - توکن‌های تازه‌سازی
type RefreshToken struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Token     string     `json:"-" gorm:"size:255;uniqueIndex;not null"` // توکن (مخفی در JSON)
	UserID    uint       `json:"user_id" gorm:"not null;index"`          // کاربر
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null;index"`       // زمان انقضا
	Revoked   bool       `json:"revoked" gorm:"default:false;index"`     // لغو شده؟
	RevokedAt *time.Time `json:"revoked_at" gorm:"index"`                // زمان لغو
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// روابط
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// UserActivity - فعالیت‌های کاربر
type UserActivity struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`         // کاربر
	Action    string    `json:"action" gorm:"size:100;not null;index"` // نوع فعالیت
	Details   string    `json:"details" gorm:"type:text"`              // جزئیات
	IPAddress string    `json:"ip_address" gorm:"size:45;index"`       // IP
	UserAgent string    `json:"user_agent" gorm:"size:500"`            // User Agent
	CreatedAt time.Time `json:"created_at" gorm:"index"`               // زمان فعالیت

	// روابط
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

// UserAnalytics - آمار کاربران
type UserAnalytics struct {
	ID               uint       `json:"id" gorm:"primaryKey"`
	UserID           uint       `json:"user_id" gorm:"not null;index"`
	TotalLogins      int64      `json:"total_logins" gorm:"default:0"`
	SuccessfulLogins int64      `json:"successful_logins" gorm:"default:0"`
	FailedLogins     int64      `json:"failed_logins" gorm:"default:0"`
	SuccessRate      float64    `json:"success_rate" gorm:"default:0"`
	LastLoginAt      *time.Time `json:"last_login_at"`
	ActiveSessions   int64      `json:"active_sessions" gorm:"default:0"`
	TrustedDevices   int64      `json:"trusted_devices" gorm:"default:0"`
	SecurityEvents   int64      `json:"security_events" gorm:"default:0"`
	UnresolvedEvents int64      `json:"unresolved_events" gorm:"default:0"`
	AccountLockCount int        `json:"account_lock_count" gorm:"default:0"`
	PasswordChanges  int        `json:"password_changes" gorm:"default:0"`
	ProfileUpdates   int        `json:"profile_updates" gorm:"default:0"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`

	// روابط
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// AuthUserSession - جلسات کاربران
type AuthUserSession struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	UserID       uint       `json:"user_id" gorm:"not null;index"`
	SessionToken string     `json:"-" gorm:"size:255;uniqueIndex;not null"`
	IPAddress    string     `json:"ip_address" gorm:"size:45;index"`
	UserAgent    string     `json:"user_agent" gorm:"size:500"`
	DeviceInfo   string     `json:"device_info" gorm:"type:text"`
	LoginAt      time.Time  `json:"login_at" gorm:"not null;index"`
	LastActivity *time.Time `json:"last_activity"`
	ExpiresAt    time.Time  `json:"expires_at" gorm:"not null;index"`
	IsActive     bool       `json:"is_active" gorm:"default:true;index"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	// روابط
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// UserDevice - دستگاه‌های کاربر
type UserDevice struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	UserID     uint       `json:"user_id" gorm:"not null;index"`
	DeviceID   string     `json:"device_id" gorm:"size:255;uniqueIndex;not null"`
	DeviceName string     `json:"device_name" gorm:"size:255"`
	DeviceType string     `json:"device_type" gorm:"size:50"` // mobile, tablet, desktop
	OS         string     `json:"os" gorm:"size:100"`
	Browser    string     `json:"browser" gorm:"size:100"`
	IPAddress  string     `json:"ip_address" gorm:"size:45;index"`
	UserAgent  string     `json:"user_agent" gorm:"size:500"`
	IsTrusted  bool       `json:"is_trusted" gorm:"default:false"`
	LastUsedAt *time.Time `json:"last_used_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`

	// روابط
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// UserSession - جلسات کاربر با ردیابی کامل دستگاه
// این ساختار جایگزین AuthUserSession شده است
// NOTE: از datatypes.JSON برای فیلدهای JSON استفاده می‌کنیم

type UserSession struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	UserID            uint           `json:"user_id" gorm:"not null;index"`
	SessionToken      string         `json:"-" gorm:"size:255;uniqueIndex;not null"`
	RefreshToken      string         `json:"-" gorm:"size:255;uniqueIndex;not null"`
	DeviceFingerprint string         `json:"device_fingerprint" gorm:"size:255;index"`
	DeviceName        string         `json:"device_name" gorm:"size:255"`
	Platform          string         `json:"platform" gorm:"size:100"`
	Browser           string         `json:"browser" gorm:"size:100"`
	BrowserVersion    string         `json:"browser_version" gorm:"size:50"`
	OSName            string         `json:"os_name" gorm:"size:100"`
	OSVersion         string         `json:"os_version" gorm:"size:50"`
	IPAddress         string         `json:"ip_address" gorm:"size:45;index"`
	City              string         `json:"city" gorm:"size:100"`
	Country           string         `json:"country" gorm:"size:100"`
	Timezone          string         `json:"timezone" gorm:"size:100"`
	UserAgent         string         `json:"user_agent" gorm:"size:500"`
	DeviceInfo        datatypes.JSON `json:"device_info" gorm:"type:json"`
	AuthMethod        string         `json:"auth_method" gorm:"size:50;index"`
	IsActive          bool           `json:"is_active" gorm:"default:true;index"`
	IsTrusted         bool           `json:"is_trusted" gorm:"default:false;index"`
	IsCurrent         bool           `json:"is_current" gorm:"default:false;index"`
	LoginAt           time.Time      `json:"login_at" gorm:"index"`
	LastActivity      *time.Time     `json:"last_activity"`
	ExpiresAt         time.Time      `json:"expires_at" gorm:"index"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`

	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}

// AuthEvent - رویدادهای احراز هویت و امنیت

type AuthEvent struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    *uint          `json:"user_id" gorm:"index"`
	Mobile    string         `json:"mobile" gorm:"size:20;index"`
	EventType string         `json:"event_type" gorm:"size:100;not null;index"`
	Status    string         `json:"status" gorm:"size:20;not null;index"`   // success, failed
	Severity  string         `json:"severity" gorm:"size:20;not null;index"` // info, warning, critical
	IPAddress string         `json:"ip_address" gorm:"size:45;index"`
	UserAgent string         `json:"user_agent" gorm:"size:500"`
	Metadata  datatypes.JSON `json:"metadata" gorm:"type:json"`
	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (AuthEvent) TableName() string {
	return "auth_events"
}

// SecurityEvent - رویدادهای امنیتی
type SecurityEvent struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	UserID      *uint      `json:"user_id" gorm:"index"`                      // null برای رویدادهای عمومی
	EventType   string     `json:"event_type" gorm:"size:100;not null;index"` // login_failed, suspicious_activity, etc.
	Severity    string     `json:"severity" gorm:"size:20;not null;index"`    // low, medium, high, critical
	Description string     `json:"description" gorm:"type:text"`
	IPAddress   string     `json:"ip_address" gorm:"size:45;index"`
	UserAgent   string     `json:"user_agent" gorm:"size:500"`
	Metadata    string     `json:"metadata" gorm:"type:json"` // اطلاعات اضافی
	Resolved    bool       `json:"resolved" gorm:"default:false;index"`
	ResolvedAt  *time.Time `json:"resolved_at"`
	ResolvedBy  *uint      `json:"resolved_by" gorm:"index"`
	CreatedAt   time.Time  `json:"created_at" gorm:"index"`

	// روابط
	User           *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	ResolvedByUser *User `json:"resolved_by_user,omitempty" gorm:"foreignKey:ResolvedBy"`
}

func (UserActivity) TableName() string {
	return "user_activities"
}

func (UserAnalytics) TableName() string {
	return "user_analytics"
}

func (AuthUserSession) TableName() string {
	return "auth_user_sessions"
}

func (UserDevice) TableName() string {
	return "user_devices"
}

func (SecurityEvent) TableName() string {
	return "security_events"
}
