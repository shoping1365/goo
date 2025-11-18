package repository

import (
	"context"
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository ایجاد نمونه جدید از repository محصول
func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{db: db.Session(&gorm.Session{PrepareStmt: true})}
}

// Create ایجاد محصول جدید
func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

// GetByID دریافت محصول بر اساس ID
func (r *ProductRepository) GetByID(ctx context.Context, id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Preload("Specifications", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "name", "value") }).
		First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetBySlug دریافت محصول بر اساس slug
func (r *ProductRepository) GetBySlug(ctx context.Context, slug string) (*models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Preload("Specifications", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "name", "value") }).
		Where("slug = ?", slug).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetBySKU دریافت محصول بر اساس SKU
func (r *ProductRepository) GetBySKU(ctx context.Context, sku string) (*models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Preload("Specifications", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "name", "value") }).
		Where("sku = ?", sku).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAll دریافت تمام محصولات
func (r *ProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Order("created_at DESC").Find(&products).Error
	return products, err
}

// GetByCategory دریافت محصولات یک دسته‌بندی
func (r *ProductRepository) GetByCategory(ctx context.Context, categoryID uint, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Where("category_id = ?", categoryID)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&products).Error
	return products, err
}

// GetByBrand دریافت محصولات یک برند
func (r *ProductRepository) GetByBrand(ctx context.Context, brandID uint, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Where("brand_id = ?", brandID)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&products).Error
	return products, err
}

// GetPublished دریافت محصولات منتشر شده
func (r *ProductRepository) GetPublished(ctx context.Context, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Where("status = ?", "published")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&products).Error
	return products, err
}

// GetPublishedAfter Keyset: (created_at DESC, id DESC)
func (r *ProductRepository) GetPublishedAfter(ctx context.Context, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error) {
	var list []models.Product
	qb := r.db.WithContext(ctx).
		Where("status = ?", "published").
		Order("created_at DESC, id DESC").
		Limit(limit)
	if !cursorCreatedAt.IsZero() && cursorID > 0 {
		qb = qb.Where("(created_at, id) < (?, ?)", cursorCreatedAt, cursorID)
	}
	if err := qb.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// GetByCategoryAfter Keyset: (created_at DESC, id DESC)
func (r *ProductRepository) GetByCategoryAfter(ctx context.Context, categoryID uint, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error) {
	var list []models.Product
	qb := r.db.WithContext(ctx).
		Where("category_id = ? AND status = ?", categoryID, "published").
		Order("created_at DESC, id DESC").
		Limit(limit)
	if !cursorCreatedAt.IsZero() && cursorID > 0 {
		qb = qb.Where("(created_at, id) < (?, ?)", cursorCreatedAt, cursorID)
	}
	if err := qb.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// GetByBrandAfter Keyset: (created_at DESC, id DESC)
func (r *ProductRepository) GetByBrandAfter(ctx context.Context, brandID uint, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error) {
	var list []models.Product
	qb := r.db.WithContext(ctx).
		Where("brand_id = ? AND status = ?", brandID, "published").
		Order("created_at DESC, id DESC").
		Limit(limit)
	if !cursorCreatedAt.IsZero() && cursorID > 0 {
		qb = qb.Where("(created_at, id) < (?, ?)", cursorCreatedAt, cursorID)
	}
	if err := qb.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// GetNewProducts دریافت محصولات جدید
func (r *ProductRepository) GetNewProducts(ctx context.Context, limit int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Where("status = ? AND is_new = ?", "published", true)
	if limit > 0 {
		query = query.Limit(limit)
	}
	err := query.Order("created_at DESC").Find(&products).Error
	return products, err
}

// GetDiscountedProducts دریافت محصولات تخفیف‌دار
func (r *ProductRepository) GetDiscountedProducts(ctx context.Context, limit int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Where("status = ? AND discount_percent > ?", "published", 0)
	if limit > 0 {
		query = query.Limit(limit)
	}
	err := query.Order("discount_percent DESC").Find(&products).Error
	return products, err
}

// SearchByName جستجو بر اساس نام
func (r *ProductRepository) SearchByName(ctx context.Context, name string, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Where("name ILIKE ? OR name_en ILIKE ?", "%"+name+"%", "%"+name+"%")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&products).Error
	return products, err
}

// GetByPriceRange دریافت محصولات در محدوده قیمت
func (r *ProductRepository) GetByPriceRange(ctx context.Context, minPrice, maxPrice float64, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Where("price BETWEEN ? AND ?", minPrice, maxPrice)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("price ASC").Find(&products).Error
	return products, err
}

// GetRelatedProducts دریافت محصولات مرتبط
func (r *ProductRepository) GetRelatedProducts(ctx context.Context, productID uint, limit int) ([]models.Product, error) {
	var product models.Product
	err := r.db.WithContext(ctx).First(&product, productID).Error
	if err != nil {
		return nil, err
	}

	var products []models.Product
	query := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Images", func(db *gorm.DB) *gorm.DB { return db.Select("id", "product_id", "image_url", "sort_order") }).
		Where("id != ? AND category_id = ? AND status = ?", productID, product.CategoryID, "published")
	if limit > 0 {
		query = query.Limit(limit)
	}
	err = query.Order("created_at DESC").Find(&products).Error
	return products, err
}

// UpdateStock به‌روزرسانی موجودی
func (r *ProductRepository) UpdateStock(ctx context.Context, productID uint, quantity int) error {
	return r.db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", productID).Update("stock_quantity", quantity).Error
}

// GetLowStockProducts دریافت محصولات با موجودی کم
func (r *ProductRepository) GetLowStockProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).
		Preload("Category", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Where("stock_quantity <= min_stock_quantity AND track_inventory = ?", true).
		Order("stock_quantity ASC").Find(&products).Error
	return products, err
}

// Update به‌روزرسانی محصول
func (r *ProductRepository) Update(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

// Delete حذف محصول
func (r *ProductRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Product{}, id).Error
}

// Count تعداد محصولات
func (r *ProductRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Product{}).Count(&count).Error
	return count, err
}
