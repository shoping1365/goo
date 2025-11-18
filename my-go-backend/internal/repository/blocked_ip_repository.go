package repository

import (
	"context"
	"errors"
	"time"

	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// BlockedIPRepositoryInterface قرارداد عملیات دیتابیس برای IP های مسدود شده
// اینترفیس صرفاً شامل عملیات CRUD و کوئری است و هیچ منطق تجاری در آن قرار نمی‌گیرد
type BlockedIPRepositoryInterface interface {
	// FindActiveByIP یافتن رکورد فعال برای IP مشخص
	FindActiveByIP(ctx context.Context, ip string) (*models.BlockedIP, error)
	// CreateOrReactivate ایجاد یا فعال‌سازی مجدد مسدودیت IP
	CreateOrReactivate(ctx context.Context, ip, reason, blockedBy string, expiresAt *time.Time) error
	// Deactivate غیرفعال‌کردن مسدودیت IP
	Deactivate(ctx context.Context, ip string) error
	// ListActive لیست IPهای فعال مسدود شده (برای شمارش سریع)
	ListActive(ctx context.Context, page, limit int) ([]models.BlockedIP, int64, error)
}

// BlockedIPRepository پیاده‌سازی GORM برای BlockedIPRepositoryInterface
type BlockedIPRepository struct {
	DB *gorm.DB
}

// FindActiveByIP یافتن رکورد فعال
func (r *BlockedIPRepository) FindActiveByIP(ctx context.Context, ip string) (*models.BlockedIP, error) {
	var rec models.BlockedIP
	if err := r.DB.WithContext(ctx).Where("ip_address = ? AND is_active = true", ip).First(&rec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rec, nil
}

// CreateOrReactivate ایجاد یا فعال‌سازی مجدد
func (r *BlockedIPRepository) CreateOrReactivate(ctx context.Context, ip, reason, blockedBy string, expiresAt *time.Time) error {
	var existing models.BlockedIP
	tx := r.DB.WithContext(ctx)
	if err := tx.Where("ip_address = ?", ip).First(&existing).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			rec := models.BlockedIP{
				IPAddress: ip,
				Reason:    reason,
				BlockedBy: blockedBy,
				BlockedAt: time.Now(),
				ExpiresAt: expiresAt,
				IsActive:  true,
			}
			return tx.Create(&rec).Error
		}
		return err
	}
	// فعال‌سازی مجدد
	updates := map[string]interface{}{
		"reason":     reason,
		"blocked_by": blockedBy,
		"blocked_at": time.Now(),
		"expires_at": expiresAt,
		"is_active":  true,
	}
	return tx.Model(&existing).Where("ip_address = ?", ip).Updates(updates).Error
}

// Deactivate غیرفعال‌کردن مسدودیت
func (r *BlockedIPRepository) Deactivate(ctx context.Context, ip string) error {
	return r.DB.WithContext(ctx).Model(&models.BlockedIP{}).Where("ip_address = ?", ip).Update("is_active", false).Error
}

// ListActive لیست IPهای فعال (با صفحه‌بندی)
func (r *BlockedIPRepository) ListActive(ctx context.Context, page, limit int) ([]models.BlockedIP, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 1000 {
		limit = 50
	}
	q := r.DB.WithContext(ctx).Model(&models.BlockedIP{}).Where("is_active = ?", true)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []models.BlockedIP
	if err := q.Order("blocked_at DESC").Offset((page - 1) * limit).Limit(limit).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}
