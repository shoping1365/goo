package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// AttributeRepository پیاده‌سازی repository برای ویژگی‌ها
type AttributeRepository struct {
	db *gorm.DB
}

// NewAttributeRepository ایجاد نمونه جدید از repository ویژگی
func NewAttributeRepository(db *gorm.DB) AttributeRepositoryInterface {
	return &AttributeRepository{db: db}
}

// Create ایجاد ویژگی جدید
func (r *AttributeRepository) Create(ctx context.Context, attribute *models.Attribute) error {
	return r.db.WithContext(ctx).Create(attribute).Error
}

// GetByID دریافت ویژگی بر اساس ID
func (r *AttributeRepository) GetByID(ctx context.Context, id uint) (*models.Attribute, error) {
	var attribute models.Attribute
	err := r.db.WithContext(ctx).Preload("Values").First(&attribute, id).Error
	if err != nil {
		return nil, err
	}
	return &attribute, nil
}

// GetByCode دریافت ویژگی بر اساس کد
func (r *AttributeRepository) GetByCode(ctx context.Context, code string) (*models.Attribute, error) {
	var attribute models.Attribute
	err := r.db.WithContext(ctx).Preload("Values").Where("code = ?", code).First(&attribute).Error
	if err != nil {
		return nil, err
	}
	return &attribute, nil
}

// GetAll دریافت تمام ویژگی‌ها
func (r *AttributeRepository) GetAll(ctx context.Context) ([]models.Attribute, error) {
	var attributes []models.Attribute
	err := r.db.WithContext(ctx).Preload("Values").Order("sort_order ASC, name ASC").Find(&attributes).Error
	return attributes, err
}

// GetActive دریافت ویژگی‌های فعال
func (r *AttributeRepository) GetActive(ctx context.Context) ([]models.Attribute, error) {
	var attributes []models.Attribute
	err := r.db.WithContext(ctx).Preload("Values").Where("is_active = ?", true).Order("sort_order ASC, name ASC").Find(&attributes).Error
	return attributes, err
}

// GetFilterable دریافت ویژگی‌های قابل فیلتر
func (r *AttributeRepository) GetFilterable(ctx context.Context) ([]models.Attribute, error) {
	var attributes []models.Attribute
	err := r.db.WithContext(ctx).Preload("Values").Where("is_active = ? AND is_filterable = ?", true, true).Order("sort_order ASC, name ASC").Find(&attributes).Error
	return attributes, err
}

// GetRequired دریافت ویژگی‌های اجباری
func (r *AttributeRepository) GetRequired(ctx context.Context) ([]models.Attribute, error) {
	var attributes []models.Attribute
	err := r.db.WithContext(ctx).Preload("Values").Where("is_active = ? AND is_required = ?", true, true).Order("sort_order ASC, name ASC").Find(&attributes).Error
	return attributes, err
}

// GetWithValues دریافت ویژگی‌ها با مقادیر
func (r *AttributeRepository) GetWithValues(ctx context.Context) ([]models.Attribute, error) {
	var attributes []models.Attribute
	err := r.db.WithContext(ctx).Preload("Values").Order("sort_order ASC, name ASC").Find(&attributes).Error
	return attributes, err
}

// GetByDataType دریافت ویژگی‌ها بر اساس نوع داده
func (r *AttributeRepository) GetByDataType(ctx context.Context, dataType string) ([]models.Attribute, error) {
	var attributes []models.Attribute
	err := r.db.WithContext(ctx).Preload("Values").Where("data_type = ? AND is_active = ?", dataType, true).Order("sort_order ASC, name ASC").Find(&attributes).Error
	return attributes, err
}

// Update به‌روزرسانی ویژگی
func (r *AttributeRepository) Update(ctx context.Context, attribute *models.Attribute) error {
	return r.db.WithContext(ctx).Save(attribute).Error
}

// Delete حذف ویژگی
func (r *AttributeRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Attribute{}, id).Error
}

// Count تعداد ویژگی‌ها
func (r *AttributeRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Attribute{}).Count(&count).Error
	return count, err
}
