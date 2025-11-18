package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// BrandRepository پیاده‌سازی repository برای برندها
type BrandRepository struct {
	db *gorm.DB
}

// NewBrandRepository ایجاد نمونه جدید از repository برند
func NewBrandRepository(db *gorm.DB) BrandRepositoryInterface {
	return &BrandRepository{db: db}
}

// Create ایجاد برند جدید
func (r *BrandRepository) Create(ctx context.Context, brand *models.Brand) error {
	return r.db.WithContext(ctx).Create(brand).Error
}

// GetByID دریافت برند بر اساس ID
func (r *BrandRepository) GetByID(ctx context.Context, id uint) (*models.Brand, error) {
	var brand models.Brand
	err := r.db.WithContext(ctx).First(&brand, id).Error
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

// GetBySlug دریافت برند بر اساس slug
func (r *BrandRepository) GetBySlug(ctx context.Context, slug string) (*models.Brand, error) {
	var brand models.Brand
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&brand).Error
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

// GetAll دریافت تمام برندها
func (r *BrandRepository) GetAll(ctx context.Context) ([]models.Brand, error) {
	var brands []models.Brand
	err := r.db.WithContext(ctx).Order("name ASC").Find(&brands).Error
	return brands, err
}

// GetPublished دریافت برندهای منتشر شده
func (r *BrandRepository) GetPublished(ctx context.Context) ([]models.Brand, error) {
	var brands []models.Brand
	err := r.db.WithContext(ctx).Where("published = ?", true).Order("name ASC").Find(&brands).Error
	return brands, err
}

// GetByHomePage دریافت برندهای نمایش داده شده در صفحه اصلی
func (r *BrandRepository) GetByHomePage(ctx context.Context) ([]models.Brand, error) {
	var brands []models.Brand
	err := r.db.WithContext(ctx).Where("published = ? AND show_on_home = ?", true, true).Order("name ASC").Find(&brands).Error
	return brands, err
}

// GetByMenu دریافت برندهای نمایش داده شده در منو
func (r *BrandRepository) GetByMenu(ctx context.Context) ([]models.Brand, error) {
	var brands []models.Brand
	err := r.db.WithContext(ctx).Where("published = ? AND show_in_menu = ?", true, true).Order("name ASC").Find(&brands).Error
	return brands, err
}

// SearchByName جستجو بر اساس نام
func (r *BrandRepository) SearchByName(ctx context.Context, name string) ([]models.Brand, error) {
	var brands []models.Brand
	err := r.db.WithContext(ctx).Where("name ILIKE ? OR official_name ILIKE ?", "%"+name+"%", "%"+name+"%").Order("name ASC").Find(&brands).Error
	return brands, err
}

// Update به‌روزرسانی برند
func (r *BrandRepository) Update(ctx context.Context, brand *models.Brand) error {
	return r.db.WithContext(ctx).Save(brand).Error
}

// Delete حذف برند
func (r *BrandRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Brand{}, id).Error
}

// Count تعداد برندها
func (r *BrandRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Brand{}).Count(&count).Error
	return count, err
}
