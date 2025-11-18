package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// CompressionJobRepositoryInterface defines the contract for compression job data access
// اینترفیس ریپازیتوری صف و لاگ عملیات فشرده‌سازی برای تست‌پذیری و جداسازی لایه‌ها
// Methods: CRUD + custom queries

type CompressionJobRepositoryInterface interface {
	Create(ctx context.Context, job *models.CompressionJob) error
	Update(ctx context.Context, job *models.CompressionJob) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*models.CompressionJob, error)
	GetByMediaID(ctx context.Context, mediaID uint) ([]models.CompressionJob, error)
	GetByStatus(ctx context.Context, status string) ([]models.CompressionJob, error)
	ListByBatch(ctx context.Context, batchID uint) ([]models.CompressionJob, error)
}

// CompressionJobRepository is the GORM implementation of CompressionJobRepositoryInterface
// پیاده‌سازی ریپازیتوری صف و لاگ عملیات فشرده‌سازی با GORM

type CompressionJobRepository struct {
	DB *gorm.DB
}

// Create inserts a new compression job
func (r *CompressionJobRepository) Create(ctx context.Context, job *models.CompressionJob) error {
	return r.DB.WithContext(ctx).Create(job).Error
}

// Update updates an existing compression job
func (r *CompressionJobRepository) Update(ctx context.Context, job *models.CompressionJob) error {
	return r.DB.WithContext(ctx).Save(job).Error
}

// Delete removes a compression job by ID
func (r *CompressionJobRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.CompressionJob{}, id).Error
}

// GetByID returns a compression job by its ID
func (r *CompressionJobRepository) GetByID(ctx context.Context, id uint) (*models.CompressionJob, error) {
	var job models.CompressionJob
	err := r.DB.WithContext(ctx).First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

// GetByMediaID returns all jobs for a given media file
func (r *CompressionJobRepository) GetByMediaID(ctx context.Context, mediaID uint) ([]models.CompressionJob, error) {
	var jobs []models.CompressionJob
	err := r.DB.WithContext(ctx).Where("media_id = ?", mediaID).Find(&jobs).Error
	return jobs, err
}

// GetByStatus returns all jobs with a given status
func (r *CompressionJobRepository) GetByStatus(ctx context.Context, status string) ([]models.CompressionJob, error) {
	var jobs []models.CompressionJob
	err := r.DB.WithContext(ctx).Where("status = ?", status).Find(&jobs).Error
	return jobs, err
}

// ListByBatch returns all jobs in a given batch
func (r *CompressionJobRepository) ListByBatch(ctx context.Context, batchID uint) ([]models.CompressionJob, error) {
	var jobs []models.CompressionJob
	err := r.DB.WithContext(ctx).Where("batch_id = ?", batchID).Find(&jobs).Error
	return jobs, err
}
