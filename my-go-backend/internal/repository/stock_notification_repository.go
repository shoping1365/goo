package repository

import (
	"context"
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

// StockNotificationRepository پیاده‌سازی repository برای اعلان‌های موجودی
type StockNotificationRepository struct {
	db *gorm.DB
}

// NewStockNotificationRepository ایجاد نمونه جدید از repository اعلان موجودی
func NewStockNotificationRepository(db *gorm.DB) StockNotificationRepositoryInterface {
	return &StockNotificationRepository{db: db}
}

// Create ایجاد اعلان موجودی جدید
func (r *StockNotificationRepository) Create(ctx context.Context, notification *models.StockNotification) error {
	return r.db.WithContext(ctx).Create(notification).Error
}

// GetByID دریافت اعلان موجودی بر اساس ID
func (r *StockNotificationRepository) GetByID(ctx context.Context, id uint) (*models.StockNotification, error) {
	var notification models.StockNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").First(&notification, id).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

// GetAll دریافت تمام اعلان‌های موجودی
func (r *StockNotificationRepository) GetAll(ctx context.Context) ([]models.StockNotification, error) {
	var notifications []models.StockNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

// GetByProduct دریافت اعلان‌های یک محصول
func (r *StockNotificationRepository) GetByProduct(ctx context.Context, productID uint) ([]models.StockNotification, error) {
	var notifications []models.StockNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").Where("product_id = ?", productID).Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

// GetByUser دریافت اعلان‌های یک کاربر
func (r *StockNotificationRepository) GetByUser(ctx context.Context, userID uint) ([]models.StockNotification, error) {
	var notifications []models.StockNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").Where("user_id = ?", userID).Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

// GetPending دریافت اعلان‌های در انتظار
func (r *StockNotificationRepository) GetPending(ctx context.Context) ([]models.StockNotification, error) {
	var notifications []models.StockNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").Where("status = ?", "pending").Order("created_at ASC").Find(&notifications).Error
	return notifications, err
}

// GetPendingByProduct دریافت اعلان‌های در انتظار یک محصول خاص
func (r *StockNotificationRepository) GetPendingByProduct(ctx context.Context, productID uint) ([]models.StockNotification, error) {
	var notifications []models.StockNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").
		Where("product_id = ? AND status = ?", productID, "pending").
		Order("created_at ASC").Find(&notifications).Error
	return notifications, err
}

// GetSent دریافت اعلان‌های ارسال شده
func (r *StockNotificationRepository) GetSent(ctx context.Context) ([]models.StockNotification, error) {
	var notifications []models.StockNotification
	err := r.db.WithContext(ctx).Preload("User").Preload("Product").Where("status = ?", "sent").Order("sent_at DESC").Find(&notifications).Error
	return notifications, err
}

// MarkAsSent علامت‌گذاری به عنوان ارسال شده
func (r *StockNotificationRepository) MarkAsSent(ctx context.Context, notificationID uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&models.StockNotification{}).
		Where("id = ?", notificationID).
		Updates(map[string]interface{}{
			"status":  "sent",
			"sent_at": &now,
		}).Error
}

// DeleteByProduct حذف اعلان‌های یک محصول
func (r *StockNotificationRepository) DeleteByProduct(ctx context.Context, productID uint) error {
	return r.db.WithContext(ctx).Where("product_id = ?", productID).Delete(&models.StockNotification{}).Error
}

// DeleteByUser حذف اعلان‌های یک کاربر
func (r *StockNotificationRepository) DeleteByUser(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&models.StockNotification{}).Error
}

// Update به‌روزرسانی اعلان موجودی
func (r *StockNotificationRepository) Update(ctx context.Context, notification *models.StockNotification) error {
	return r.db.WithContext(ctx).Save(notification).Error
}

// Delete حذف اعلان موجودی
func (r *StockNotificationRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.StockNotification{}, id).Error
}

// Count تعداد اعلان‌های موجودی
func (r *StockNotificationRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.StockNotification{}).Count(&count).Error
	return count, err
}
