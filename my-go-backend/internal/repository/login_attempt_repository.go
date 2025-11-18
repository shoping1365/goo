package repository

import (
	"context"
	"time"

	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// LoginAttemptFilter فیلترهای جستجو برای تاریخچه ورود
type LoginAttemptFilter struct {
	// جستجو بر اساس موبایل (جستجوی جزئی)
	Search string
	// فیلتر موفق/ناموفق
	IsSuccessful *bool
	// روش تلاش (otp_request, otp_verify, password)
	AttemptType string
	// از تاریخ (RFC3339 یا هر فرمت قابل تبدیل توسط DB)
	DateFrom string
	// فیلتر بر اساس IP
	IP string
}

// LoginAttemptRepositoryInterface
// اینترفیس ریپازیتوری تاریخچه ورود کاربران
type LoginAttemptRepositoryInterface interface {
	// List دریافت لیست تلاش‌های ورود با صفحه‌بندی و فیلتر به‌همراه مجموع کل
	List(ctx context.Context, filter LoginAttemptFilter, page, limit int) ([]models.LoginAttempt, int64, error)
	// CountSuccessful تعداد ورودهای موفق
	CountSuccessful(ctx context.Context) (int64, error)
	// CountFailed تعداد ورودهای ناموفق
	CountFailed(ctx context.Context) (int64, error)
	// CountDistinctFailedIPsSince تعداد IP های یکتای ناموفق از تاریخ مشخص
	CountDistinctFailedIPsSince(ctx context.Context, since time.Time) (int64, error)
}

// LoginAttemptRepository پیاده‌سازی GORM برای ریپازیتوری تلاش‌های ورود
type LoginAttemptRepository struct {
	DB *gorm.DB
}

// List دریافت لیست تلاش‌های ورود با فیلتر و صفحه‌بندی
func (r *LoginAttemptRepository) List(ctx context.Context, filter LoginAttemptFilter, page, limit int) ([]models.LoginAttempt, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 1000 {
		limit = 20
	}

	q := r.DB.WithContext(ctx).Model(&models.LoginAttempt{})

	if filter.Search != "" {
		q = q.Where("mobile ILIKE ?", "%"+filter.Search+"%")
	}
	if filter.IsSuccessful != nil {
		q = q.Where("is_successful = ?", *filter.IsSuccessful)
	}
	if filter.AttemptType != "" {
		q = q.Where("attempt_type = ?", filter.AttemptType)
	}
	if filter.DateFrom != "" {
		// به DB اجازه می‌دهیم تبدیل را انجام دهد
		q = q.Where("created_at >= ?", filter.DateFrom)
	}
	if filter.IP != "" {
		q = q.Where("attempt_ip = ?", filter.IP)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var attempts []models.LoginAttempt
	if err := q.Order("created_at DESC").Offset((page - 1) * limit).Limit(limit).Find(&attempts).Error; err != nil {
		return nil, 0, err
	}

	return attempts, total, nil
}

// CountSuccessful شمارش ورودهای موفق
func (r *LoginAttemptRepository) CountSuccessful(ctx context.Context) (int64, error) {
	var count int64
	err := r.DB.WithContext(ctx).Model(&models.LoginAttempt{}).Where("is_successful = ?", true).Count(&count).Error
	return count, err
}

// CountFailed شمارش ورودهای ناموفق
func (r *LoginAttemptRepository) CountFailed(ctx context.Context) (int64, error) {
	var count int64
	err := r.DB.WithContext(ctx).Model(&models.LoginAttempt{}).Where("is_successful = ?", false).Count(&count).Error
	return count, err
}

// CountDistinctFailedIPsSince تعداد IP های یکتای ناموفق از تاریخ مشخص
func (r *LoginAttemptRepository) CountDistinctFailedIPsSince(ctx context.Context, since time.Time) (int64, error) {
	var count int64
	err := r.DB.WithContext(ctx).
		Model(&models.LoginAttempt{}).
		Where("is_successful = ?", false).
		Where("created_at >= ?", since).
		Distinct("attempt_ip").
		Count(&count).Error
	return count, err
}
