package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// AttributeGroupRepository پیاده‌سازی repository برای گروه‌های ویژگی
type AttributeGroupRepository struct {
	db *gorm.DB
}

// NewAttributeGroupRepository ایجاد نمونه جدید از repository گروه ویژگی
func NewAttributeGroupRepository(db *gorm.DB) AttributeGroupRepositoryInterface {
	return &AttributeGroupRepository{db: db}
}

// Create ایجاد گروه ویژگی جدید
func (r *AttributeGroupRepository) Create(ctx context.Context, group *models.AttributeGroup) error {
	return r.db.WithContext(ctx).Create(group).Error
}

// GetByID دریافت گروه ویژگی بر اساس ID
func (r *AttributeGroupRepository) GetByID(ctx context.Context, id uint) (*models.AttributeGroup, error) {
	var group models.AttributeGroup
	err := r.db.WithContext(ctx).Preload("Category").Preload("Attributes.Attribute").First(&group, id).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// GetAll دریافت تمام گروه‌های ویژگی
func (r *AttributeGroupRepository) GetAll(ctx context.Context) ([]models.AttributeGroup, error) {
	var groups []models.AttributeGroup
	err := r.db.WithContext(ctx).Preload("Category").Preload("Attributes.Attribute").Order("name ASC").Find(&groups).Error
	return groups, err
}

// GetByCategory دریافت گروه‌های ویژگی یک دسته‌بندی
func (r *AttributeGroupRepository) GetByCategory(ctx context.Context, categoryID uint) ([]models.AttributeGroup, error) {
	var groups []models.AttributeGroup
	err := r.db.WithContext(ctx).Preload("Category").Preload("Attributes.Attribute").Where("category_id = ?", categoryID).Order("name ASC").Find(&groups).Error
	return groups, err
}

// GetWithAttributes دریافت گروه‌ها با ویژگی‌ها
func (r *AttributeGroupRepository) GetWithAttributes(ctx context.Context) ([]models.AttributeGroup, error) {
	var groups []models.AttributeGroup
	err := r.db.WithContext(ctx).Preload("Category").Preload("Attributes.Attribute").Order("name ASC").Find(&groups).Error
	return groups, err
}

// GetByCategoryWithAttributes دریافت گروه‌های یک دسته‌بندی با ویژگی‌ها
func (r *AttributeGroupRepository) GetByCategoryWithAttributes(ctx context.Context, categoryID uint) ([]models.AttributeGroup, error) {
	var groups []models.AttributeGroup
	err := r.db.WithContext(ctx).Preload("Category").Preload("Attributes.Attribute").Where("category_id = ?", categoryID).Order("name ASC").Find(&groups).Error
	return groups, err
}

// Update به‌روزرسانی گروه ویژگی
func (r *AttributeGroupRepository) Update(ctx context.Context, group *models.AttributeGroup) error {
	return r.db.WithContext(ctx).Save(group).Error
}

// Delete حذف گروه ویژگی
func (r *AttributeGroupRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.AttributeGroup{}, id).Error
}

// Count تعداد گروه‌های ویژگی
func (r *AttributeGroupRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.AttributeGroup{}).Count(&count).Error
	return count, err
}
