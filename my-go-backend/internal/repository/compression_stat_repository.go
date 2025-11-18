package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// CompressionStatRepositoryInterface defines the contract for compression stat data access
// اینترفیس ریپازیتوری آمار و گزارش‌های فشرده‌سازی برای تست‌پذیری و جداسازی لایه‌ها
// Methods: CRUD + custom queries

type CompressionStatRepositoryInterface interface {
	Create(ctx context.Context, stat *models.CompressionStat) error
	Update(ctx context.Context, stat *models.CompressionStat) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*models.CompressionStat, error)
	GetByType(ctx context.Context, statType string) ([]models.CompressionStat, error)
	GetByMedia(ctx context.Context, mediaID uint) ([]models.CompressionStat, error)
	GetByUser(ctx context.Context, userID uint) ([]models.CompressionStat, error)
	GetByFormat(ctx context.Context, format string) ([]models.CompressionStat, error)
	GetByPeriod(ctx context.Context, period string) ([]models.CompressionStat, error)
}

// CompressionStatRepository is the GORM implementation of CompressionStatRepositoryInterface
// پیاده‌سازی ریپازیتوری آمار و گزارش‌های فشرده‌سازی با GORM

type CompressionStatRepository struct {
	DB *gorm.DB
}

// Create inserts a new compression stat
func (r *CompressionStatRepository) Create(ctx context.Context, stat *models.CompressionStat) error {
	return r.DB.WithContext(ctx).Create(stat).Error
}

// Update updates an existing compression stat
func (r *CompressionStatRepository) Update(ctx context.Context, stat *models.CompressionStat) error {
	return r.DB.WithContext(ctx).Save(stat).Error
}

// Delete removes a compression stat by ID
func (r *CompressionStatRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.CompressionStat{}, id).Error
}

// GetByID returns a compression stat by its ID
func (r *CompressionStatRepository) GetByID(ctx context.Context, id uint) (*models.CompressionStat, error) {
	var stat models.CompressionStat
	err := r.DB.WithContext(ctx).First(&stat, id).Error
	if err != nil {
		return nil, err
	}
	return &stat, nil
}

// GetByType returns all stats for a given type
func (r *CompressionStatRepository) GetByType(ctx context.Context, statType string) ([]models.CompressionStat, error) {
	var stats []models.CompressionStat
	err := r.DB.WithContext(ctx).Where("stat_type = ?", statType).Find(&stats).Error
	return stats, err
}

// GetByMedia returns all stats for a given media file
func (r *CompressionStatRepository) GetByMedia(ctx context.Context, mediaID uint) ([]models.CompressionStat, error) {
	var stats []models.CompressionStat
	err := r.DB.WithContext(ctx).Where("media_id = ?", mediaID).Find(&stats).Error
	return stats, err
}

// GetByUser returns all stats for a given user
func (r *CompressionStatRepository) GetByUser(ctx context.Context, userID uint) ([]models.CompressionStat, error) {
	var stats []models.CompressionStat
	err := r.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&stats).Error
	return stats, err
}

// GetByFormat returns all stats for a given format
func (r *CompressionStatRepository) GetByFormat(ctx context.Context, format string) ([]models.CompressionStat, error) {
	var stats []models.CompressionStat
	err := r.DB.WithContext(ctx).Where("format = ?", format).Find(&stats).Error
	return stats, err
}

// GetByPeriod returns all stats for a given period
func (r *CompressionStatRepository) GetByPeriod(ctx context.Context, period string) ([]models.CompressionStat, error) {
	var stats []models.CompressionStat
	err := r.DB.WithContext(ctx).Where("period = ?", period).Find(&stats).Error
	return stats, err
}
