package repository

import (
	"my-go-backend/internal/models"
	"gorm.io/gorm"
)

type CustomHeaderRepository struct {
	db *gorm.DB
}

func NewCustomHeaderRepository(db *gorm.DB) *CustomHeaderRepository {
	return &CustomHeaderRepository{db: db}
}

// GetAllHeaders دریافت تمام هدرها با لایه‌هایشان
func (r *CustomHeaderRepository) GetAllHeaders() ([]models.Header, error) {
	var headers []models.Header
	err := r.db.Preload("Layers").Where("deleted_at IS NULL").Find(&headers).Error
	return headers, err
}

// GetHeaderByID دریافت هدر بر اساس ID
func (r *CustomHeaderRepository) GetHeaderByID(id uint) (*models.Header, error) {
	var header models.Header
	err := r.db.Preload("Layers").Where("id = ? AND deleted_at IS NULL", id).First(&header).Error
	if err != nil {
		return nil, err
	}
	return &header, nil
}

// CreateHeader ایجاد هدر جدید
func (r *CustomHeaderRepository) CreateHeader(header *models.Header) error {
	return r.db.Create(header).Error
}

// UpdateHeader به‌روزرسانی هدر
func (r *CustomHeaderRepository) UpdateHeader(header *models.Header) error {
	return r.db.Save(header).Error
}

// DeleteHeader حذف هدر (soft delete)
func (r *CustomHeaderRepository) DeleteHeader(id uint) error {
	return r.db.Delete(&models.Header{}, id).Error
}

// GetActiveHeaders دریافت هدرهای فعال
func (r *CustomHeaderRepository) GetActiveHeaders() ([]models.Header, error) {
	var headers []models.Header
	err := r.db.Preload("Layers").Where("is_active = ? AND deleted_at IS NULL", true).Find(&headers).Error
	return headers, err
}

// HeaderExists بررسی وجود هدر بر اساس نام
func (r *CustomHeaderRepository) HeaderExists(name string, excludeID ...uint) (bool, error) {
	var count int64
	query := r.db.Model(&models.Header{}).Where("name = ? AND deleted_at IS NULL", name)
	
	if len(excludeID) > 0 {
		query = query.Where("id != ?", excludeID[0])
	}
	
	err := query.Count(&count).Error
	return count > 0, err
} 