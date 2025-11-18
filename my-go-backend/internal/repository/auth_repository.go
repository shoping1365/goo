package repository

import (
	"context"
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

// AuthRepository - رابط برای عملیات احراز هویت
type AuthRepository interface {
	// User queries
	FindUserByPhone(ctx context.Context, phone string) (*models.User, error)
	FindUserByID(ctx context.Context, userID uint) (*models.User, error)
	CreateUserFromOTP(ctx context.Context, phone string) (*models.User, error)
	UpdateUserLastLogin(ctx context.Context, userID uint, ip string) error

	// OTP queries
	CreateOTP(ctx context.Context, phone string, code string, purpose string, expiresAt time.Time) error
	FindValidOTP(ctx context.Context, phone string) (*models.OTPCode, error)
	MarkOTPAsUsed(ctx context.Context, otpID uint) error
	IncrementOTPAttempts(ctx context.Context, otpID uint) error

	// Session queries
	CreateSession(ctx context.Context, userID uint, refreshTokenHash string, expiresAt time.Time) error
	FindSessionByTokenHash(ctx context.Context, tokenHash string) (*models.Session, error)
	DeleteSession(ctx context.Context, sessionID uint) error
	DeleteAllUserSessions(ctx context.Context, userID uint) error

	// Permission queries
	GetUserPermissions(ctx context.Context, userID uint) ([]string, error)
}

// authRepositoryImpl - پیاده‌سازی AuthRepository
type authRepositoryImpl struct {
	db *gorm.DB
}

// NewAuthRepository - ایجاد instance جدید
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{
		db: db.Session(&gorm.Session{PrepareStmt: true}),
	}
}

// ============ User Queries ============

// FindUserByPhone - پیدا کردن کاربر با شماره موبایل
func (r *authRepositoryImpl) FindUserByPhone(ctx context.Context, phone string) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("mobile = ? AND status = ?", phone, "active").First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

// FindUserByID - پیدا کردن کاربر با ID
func (r *authRepositoryImpl) FindUserByID(ctx context.Context, userID uint) (*models.User, error) {
	var user models.User
	result := r.db.WithContext(ctx).Where("id = ? AND status = ?", userID, "active").First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

// CreateUserFromOTP - ایجاد کاربر جدید از OTP
func (r *authRepositoryImpl) CreateUserFromOTP(ctx context.Context, phone string) (*models.User, error) {
	user := &models.User{
		Mobile:   phone,
		Username: phone, // استفاده از موبایل به عنوان username موقتی
		Status:   "active",
		RoleID:   1, // Default user role
	}

	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// UpdateUserLastLogin - به‌روزرسانی آخرین ورود
func (r *authRepositoryImpl) UpdateUserLastLogin(ctx context.Context, userID uint, ip string) error {
	return r.db.WithContext(ctx).Model(&models.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"last_login_at": time.Now(),
			"last_login_ip": ip,
		}).Error
}

// ============ OTP Queries ============

// CreateOTP - ایجاد OTP جدید
func (r *authRepositoryImpl) CreateOTP(ctx context.Context, phone string, code string, purpose string, expiresAt time.Time) error {
	otp := &models.OTPCode{
		Mobile:    phone,
		Code:      code,
		Purpose:   purpose,
		ExpiresAt: expiresAt,
		IsUsed:    false,
	}

	return r.db.WithContext(ctx).Create(otp).Error
}

// FindValidOTP - پیدا کردن OTP معتبر آخرین
func (r *authRepositoryImpl) FindValidOTP(ctx context.Context, phone string) (*models.OTPCode, error) {
	var otp models.OTPCode
	result := r.db.WithContext(ctx).
		Where("mobile = ? AND is_used = false AND expires_at > ?", phone, time.Now()).
		Order("created_at DESC").
		First(&otp)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &otp, nil
}

// MarkOTPAsUsed - علامت‌گذاری OTP به عنوان استفاده شده
func (r *authRepositoryImpl) MarkOTPAsUsed(ctx context.Context, otpID uint) error {
	return r.db.WithContext(ctx).Model(&models.OTPCode{}).
		Where("id = ?", otpID).
		Updates(map[string]interface{}{
			"is_used": true,
			"used_at": time.Now(),
		}).Error
}

// IncrementOTPAttempts - افزایش تعداد تلاش‌های OTP
func (r *authRepositoryImpl) IncrementOTPAttempts(ctx context.Context, otpID uint) error {
	return r.db.WithContext(ctx).Model(&models.OTPCode{}).
		Where("id = ?", otpID).
		Update("attempts", gorm.Expr("attempts + 1")).Error
}

// ============ Session Queries ============

// CreateSession - ایجاد نشست جدید
func (r *authRepositoryImpl) CreateSession(ctx context.Context, userID uint, refreshTokenHash string, expiresAt time.Time) error {
	session := &models.Session{
		UserID:    userID,
		Token:     refreshTokenHash,
		ExpiresAt: expiresAt,
		IsActive:  true,
	}

	return r.db.WithContext(ctx).Create(session).Error
}

// FindSessionByTokenHash - پیدا کردن نشست با token hash
func (r *authRepositoryImpl) FindSessionByTokenHash(ctx context.Context, tokenHash string) (*models.Session, error) {
	var session models.Session
	result := r.db.WithContext(ctx).
		Where("token = ? AND is_active = true AND expires_at > ?", tokenHash, time.Now()).
		First(&session)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &session, nil
}

// DeleteSession - حذف نشست
func (r *authRepositoryImpl) DeleteSession(ctx context.Context, sessionID uint) error {
	return r.db.WithContext(ctx).Model(&models.Session{}).
		Where("id = ?", sessionID).
		Update("is_active", false).Error
}

// DeleteAllUserSessions - حذف تمام نشست‌های کاربر
func (r *authRepositoryImpl) DeleteAllUserSessions(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).Model(&models.Session{}).
		Where("user_id = ?", userID).
		Update("is_active", false).Error
}

// ============ Permission Queries ============

// GetUserPermissions - دریافت مجوزهای کاربر
func (r *authRepositoryImpl) GetUserPermissions(ctx context.Context, userID uint) ([]string, error) {
	var permissions []string

	err := r.db.WithContext(ctx).
		Table("permissions").
		Joins("INNER JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Joins("INNER JOIN roles ON roles.id = role_permissions.role_id").
		Joins("INNER JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ? AND permissions.is_active = true", userID).
		Select("DISTINCT permissions.name").
		Pluck("name", &permissions).
		Error

	if err != nil {
		return nil, err
	}

	return permissions, nil
}
