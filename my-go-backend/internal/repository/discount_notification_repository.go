package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

	"my-go-backend/internal/models"
)

// DiscountNotificationRepositoryInterface رابط ریپازیتوری اعلان‌های تخفیف
type DiscountNotificationRepositoryInterface interface {
	Create(ctx context.Context, notification *models.DiscountNotification) error
	GetByID(ctx context.Context, id uint) (*models.DiscountNotification, error)
	GetAll(ctx context.Context) ([]models.DiscountNotification, error)
	GetPending(ctx context.Context) ([]models.DiscountNotification, error)
	GetByProduct(ctx context.Context, productID uint) ([]models.DiscountNotification, error)
	GetPendingByProduct(ctx context.Context, productID uint) ([]models.DiscountNotification, error)
	GetByUser(ctx context.Context, userID uint) ([]models.DiscountNotification, error)
	MarkAsSent(ctx context.Context, notificationID uint) error
	Update(ctx context.Context, notification *models.DiscountNotification) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context) (int64, error)
}

type DiscountNotificationRepository struct {
	db *gorm.DB
}

func NewDiscountNotificationRepository(db *gorm.DB) DiscountNotificationRepositoryInterface {
	return &DiscountNotificationRepository{db: db}
}

func (r *DiscountNotificationRepository) Create(ctx context.Context, notification *models.DiscountNotification) error {
	return r.db.WithContext(ctx).Create(notification).Error
}

func (r *DiscountNotificationRepository) GetByID(ctx context.Context, id uint) (*models.DiscountNotification, error) {
	var n models.DiscountNotification
	if err := r.db.WithContext(ctx).Preload("User").Preload("Product").First(&n, id).Error; err != nil {
		return nil, err
	}
	return &n, nil
}

func (r *DiscountNotificationRepository) GetAll(ctx context.Context) ([]models.DiscountNotification, error) {
	var list []models.DiscountNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *DiscountNotificationRepository) GetPending(ctx context.Context) ([]models.DiscountNotification, error) {
	var list []models.DiscountNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").Where("status = ?", "pending").Order("created_at ASC").Find(&list).Error
	return list, err
}

func (r *DiscountNotificationRepository) GetByProduct(ctx context.Context, productID uint) ([]models.DiscountNotification, error) {
	var list []models.DiscountNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").Where("product_id = ?", productID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *DiscountNotificationRepository) GetPendingByProduct(ctx context.Context, productID uint) ([]models.DiscountNotification, error) {
	var list []models.DiscountNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").
		Where("product_id = ? AND status = ?", productID, "pending").
		Order("created_at ASC").Find(&list).Error
	return list, err
}

func (r *DiscountNotificationRepository) GetByUser(ctx context.Context, userID uint) ([]models.DiscountNotification, error) {
	var list []models.DiscountNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").Where("user_id = ?", userID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *DiscountNotificationRepository) MarkAsSent(ctx context.Context, notificationID uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&models.DiscountNotification{}).
		Where("id = ?", notificationID).
		Updates(map[string]any{"status": "sent", "sent_at": &now}).Error
}

func (r *DiscountNotificationRepository) Update(ctx context.Context, notification *models.DiscountNotification) error {
	return r.db.WithContext(ctx).Save(notification).Error
}

func (r *DiscountNotificationRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.DiscountNotification{}, id).Error
}

func (r *DiscountNotificationRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.DiscountNotification{}).Count(&count).Error
	return count, err
}
