package repository

import (
	"context"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository ایجاد نمونه جدید از repository دسته‌بندی
// اتصال دیتابیس با PrepareStmt برای بهبود کارایی کوئری‌های پرتکرار مقداردهی می‌شود.
func NewCategoryRepository(db *gorm.DB) CategoryRepositoryInterface {
	return &CategoryRepository{db: db.Session(&gorm.Session{PrepareStmt: true})}
}

// Create ایجاد دسته‌بندی جدید
func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

// GetByID دریافت دسته‌بندی بر اساس ID
func (r *CategoryRepository) GetByID(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetBySlug دریافت دسته‌بندی بر اساس slug
func (r *CategoryRepository) GetBySlug(ctx context.Context, slug string) (*models.Category, error) {
	var category models.Category
	err := r.db.WithContext(ctx).Where("slug = ?", slug).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetAll دریافت تمام دسته‌بندی‌ها
func (r *CategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Order("\"order\" ASC, name ASC").Find(&categories).Error
	return categories, err
}

// GetByParentID دریافت زیردسته‌ها
func (r *CategoryRepository) GetByParentID(ctx context.Context, parentID *uint) ([]models.Category, error) {
	var categories []models.Category
	query := r.db.WithContext(ctx).Where("parent_id = ?", parentID)
	err := query.Order("\"order\" ASC, name ASC").Find(&categories).Error
	return categories, err
}

// GetRootCategories دریافت دسته‌های اصلی
func (r *CategoryRepository) GetRootCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Where("parent_id IS NULL").Order("\"order\" ASC, name ASC").Find(&categories).Error
	return categories, err
}

// GetPublished دریافت دسته‌های منتشر شده
func (r *CategoryRepository) GetPublished(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Where("published = ?", true).Order("\"order\" ASC, name ASC").Find(&categories).Error
	return categories, err
}

// GetByHomePage دریافت دسته‌های نمایش داده شده در صفحه اصلی
func (r *CategoryRepository) GetByHomePage(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Where("published = ? AND show_on_home = ?", true, true).Order("\"order\" ASC, name ASC").Find(&categories).Error
	return categories, err
}

// GetByMenu دریافت دسته‌های نمایش داده شده در منو
func (r *CategoryRepository) GetByMenu(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Where("published = ? AND show_in_menu = ?", true, true).Order("\"order\" ASC, name ASC").Find(&categories).Error
	return categories, err
}

// GetHierarchy دریافت ساختار سلسله مراتبی
func (r *CategoryRepository) GetHierarchy(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).Where("parent_id IS NULL").Order("\"order\" ASC, name ASC").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	// بارگذاری زیردسته‌ها به صورت بازگشتی
	for i := range categories {
		if err := r.loadChildren(ctx, &categories[i]); err != nil {
			return nil, err
		}
	}

	return categories, nil
}

// loadChildren بارگذاری زیردسته‌ها به صورت بازگشتی
func (r *CategoryRepository) loadChildren(ctx context.Context, category *models.Category) error {
	var children []models.Category
	err := r.db.WithContext(ctx).Where("parent_id = ?", category.ID).Order("\"order\" ASC, name ASC").Find(&children).Error
	if err != nil {
		return err
	}

	// بارگذاری زیردسته‌های زیردسته‌ها
	for i := range children {
		if err := r.loadChildren(ctx, &children[i]); err != nil {
			return err
		}
	}

	return nil
}

// GetWithProductCount دریافت دسته‌ها با تعداد محصولات
func (r *CategoryRepository) GetWithProductCount(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.WithContext(ctx).
		Select("categories.*, COUNT(products.id) as product_count").
		Joins("LEFT JOIN products ON categories.id = products.category_id").
		Group("categories.id").
		Order("categories.\"order\" ASC, categories.name ASC").
		Find(&categories).Error
	return categories, err
}

// Update به‌روزرسانی دسته‌بندی
func (r *CategoryRepository) Update(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

// Delete حذف دسته‌بندی
func (r *CategoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Category{}, id).Error
}

// Count تعداد دسته‌بندی‌ها
func (r *CategoryRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Category{}).Count(&count).Error
	return count, err
}
