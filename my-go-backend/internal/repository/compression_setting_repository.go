package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// CompressionSettingRepositoryInterface defines the contract for compression setting data access
// اینترفیس ریپازیتوری تنظیمات فشرده‌سازی برای تست‌پذیری و جداسازی لایه‌ها
// Methods: CRUD + custom queries

type CompressionSettingRepositoryInterface interface {
	Create(ctx context.Context, setting *models.CompressionSetting) error
	Update(ctx context.Context, setting *models.CompressionSetting) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*models.CompressionSetting, error)
	GetByScope(ctx context.Context, scope string) ([]models.CompressionSetting, error)
	GetByUser(ctx context.Context, userID uint) ([]models.CompressionSetting, error)
	GetByMedia(ctx context.Context, mediaID uint) ([]models.CompressionSetting, error)
	GetByJob(ctx context.Context, jobID uint) ([]models.CompressionSetting, error)
}

// CompressionSettingRepository is the GORM implementation of CompressionSettingRepositoryInterface
// پیاده‌سازی ریپازیتوری تنظیمات فشرده‌سازی با GORM

type CompressionSettingRepository struct {
	DB *gorm.DB
}

// Create inserts a new compression setting
func (r *CompressionSettingRepository) Create(ctx context.Context, setting *models.CompressionSetting) error {
	return r.DB.WithContext(ctx).Create(setting).Error
}

// Update updates an existing compression setting
func (r *CompressionSettingRepository) Update(ctx context.Context, setting *models.CompressionSetting) error {
	return r.DB.WithContext(ctx).Save(setting).Error
}

// Delete removes a compression setting by ID
func (r *CompressionSettingRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.CompressionSetting{}, id).Error
}

// GetByID returns a compression setting by its ID
func (r *CompressionSettingRepository) GetByID(ctx context.Context, id uint) (*models.CompressionSetting, error) {
	var setting models.CompressionSetting
	err := r.DB.WithContext(ctx).First(&setting, id).Error
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

// GetByScope returns all settings for a given scope
func (r *CompressionSettingRepository) GetByScope(ctx context.Context, scope string) ([]models.CompressionSetting, error) {
	var settings []models.CompressionSetting
	err := r.DB.WithContext(ctx).Where("scope = ?", scope).Find(&settings).Error
	return settings, err
}

// GetByUser returns all settings for a given user
func (r *CompressionSettingRepository) GetByUser(ctx context.Context, userID uint) ([]models.CompressionSetting, error) {
	var settings []models.CompressionSetting
	err := r.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&settings).Error
	return settings, err
}

// GetByMedia returns all settings for a given media file
func (r *CompressionSettingRepository) GetByMedia(ctx context.Context, mediaID uint) ([]models.CompressionSetting, error) {
	var settings []models.CompressionSetting
	err := r.DB.WithContext(ctx).Where("media_id = ?", mediaID).Find(&settings).Error
	return settings, err
}

// GetByJob returns all settings for a given job
func (r *CompressionSettingRepository) GetByJob(ctx context.Context, jobID uint) ([]models.CompressionSetting, error) {
	var settings []models.CompressionSetting
	err := r.DB.WithContext(ctx).Where("job_id = ?", jobID).Find(&settings).Error
	return settings, err
}
