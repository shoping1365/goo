package repository

import (
	"context"
	"database/sql"
	"fmt"
	"my-go-backend/internal/database"
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

// OptimizedProductRepository پیاده‌سازی بهینه‌شده repository برای محصولات با استفاده از Prepared Statements
type OptimizedProductRepository struct {
	db                   *gorm.DB
	preparedStatementMgr *database.PreparedStatementManager
}

// NewOptimizedProductRepository ایجاد نمونه جدید از repository محصول بهینه‌شده
func NewOptimizedProductRepository(db *gorm.DB, preparedStatementMgr *database.PreparedStatementManager) ProductRepositoryInterface {
	return &OptimizedProductRepository{
		db:                   db,
		preparedStatementMgr: preparedStatementMgr,
	}
}

// Create ایجاد محصول جدید
func (r *OptimizedProductRepository) Create(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

// GetByID دریافت محصول بر اساس ID با استفاده از Prepared Statement
func (r *OptimizedProductRepository) GetByID(ctx context.Context, id uint) (*models.Product, error) {
	stmt, err := r.preparedStatementMgr.GetStatement("product_get_by_id")
	if err != nil {
		// Fallback به GORM در صورت عدم دسترسی به Prepared Statement
		var product models.Product
		err := r.db.WithContext(ctx).Preload("Category").Preload("Images").Preload("Specifications").First(&product, id).Error
		if err != nil {
			return nil, err
		}
		return &product, nil
	}

	row := database.ExecutePreparedQueryRow(ctx, stmt, id)

	var product models.Product
	var createdAt, updatedAt time.Time
	var deletedAt sql.NullTime
	var categoryID, brandID, updatedBy sql.NullInt64

	err = row.Scan(
		&product.ID,
		&product.UUID,
		&product.Slug,
		&product.SKU,
		&product.Name,
		&product.NameEn,
		&product.Image,
		&product.Price,
		&product.OldPrice,
		&product.Cost,
		&product.Profit,
		&product.DiscountPercent,
		&product.DiscountAmount,
		&product.SalePrice,
		&product.WholesalePrice,
		&product.WholesaleSalePrice,
		&product.Description,
		&product.FullDescription,
		&product.Status,
		&product.IsNew,
		&product.IsActive,
		&product.IsFeatured,
		&product.IsOnSale,
		&product.Weight,
		&product.Length,
		&product.Width,
		&product.Height,
		&product.ShippingCost,
		&product.ShippingTime,
		&product.StockQuantity,
		&product.MinStockQuantity,
		&product.MaxStockQuantity,
		&product.StockStatus,
		&product.TrackInventory,
		&product.ShowStockToCustomer,
		&product.DisableBuyButton,
		&product.CallForPrice,
		&product.SeoTitle,
		&product.MetaDescription,
		&product.MetaKeywords,
		&categoryID,
		&brandID,
		&product.MetaTitle,
		&product.MetaDescription,
		&product.MetaKeywords,
		&product.Tags,
		&product.Images,
		&updatedBy,
		&createdAt,
		&updatedAt,
		// &product.VideoURL,
		&deletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("خطا در اسکن نتیجه: %w", err)
	}

	product.CreatedAt = createdAt
	product.UpdatedAt = updatedAt

	if categoryID.Valid {
		categoryIDUint := uint(categoryID.Int64)
		product.CategoryID = &categoryIDUint
	}
	if brandID.Valid {
		brandIDUint := uint(brandID.Int64)
		product.BrandID = &brandIDUint
	}
	if updatedBy.Valid {
		updatedByUint := uint(updatedBy.Int64)
		product.UpdatedBy = &updatedByUint
	}

	return &product, nil
}

// GetBySlug دریافت محصول بر اساس slug با استفاده از Prepared Statement
func (r *OptimizedProductRepository) GetBySlug(ctx context.Context, slug string) (*models.Product, error) {
	stmt, err := r.preparedStatementMgr.GetStatement("product_get_by_slug")
	if err != nil {
		// Fallback به GORM
		var product models.Product
		err := r.db.WithContext(ctx).Preload("Category").Preload("Images").Preload("Specifications").Where("slug = ?", slug).First(&product).Error
		if err != nil {
			return nil, err
		}
		return &product, nil
	}

	row := database.ExecutePreparedQueryRow(ctx, stmt, slug)

	var product models.Product
	var createdAt, updatedAt time.Time
	var deletedAt sql.NullTime
	var categoryID, brandID, updatedBy sql.NullInt64

	err = row.Scan(
		&product.ID,
		&product.UUID,
		&product.Slug,
		&product.SKU,
		&product.Name,
		&product.NameEn,
		&product.Image,
		&product.Price,
		&product.OldPrice,
		&product.Cost,
		&product.Profit,
		&product.DiscountPercent,
		&product.DiscountAmount,
		&product.SalePrice,
		&product.WholesalePrice,
		&product.WholesaleSalePrice,
		&product.Description,
		&product.FullDescription,
		&product.Status,
		&product.IsNew,
		&product.IsActive,
		&product.IsFeatured,
		&product.IsOnSale,
		&product.Weight,
		&product.Length,
		&product.Width,
		&product.Height,
		&product.ShippingCost,
		&product.ShippingTime,
		&product.StockQuantity,
		&product.MinStockQuantity,
		&product.MaxStockQuantity,
		&product.StockStatus,
		&product.TrackInventory,
		&product.ShowStockToCustomer,
		&product.DisableBuyButton,
		&product.CallForPrice,
		&product.SeoTitle,
		&product.MetaDescription,
		&product.MetaKeywords,
		&categoryID,
		&brandID,
		&product.MetaTitle,
		&product.MetaDescription,
		&product.MetaKeywords,
		&product.Tags,
		&product.Images,
		&updatedBy,
		&createdAt,
		&updatedAt,
		// &product.VideoURL,
		&deletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("خطا در اسکن نتیجه: %w", err)
	}

	product.CreatedAt = createdAt
	product.UpdatedAt = updatedAt

	if categoryID.Valid {
		categoryIDUint := uint(categoryID.Int64)
		product.CategoryID = &categoryIDUint
	}
	if brandID.Valid {
		brandIDUint := uint(brandID.Int64)
		product.BrandID = &brandIDUint
	}
	if updatedBy.Valid {
		updatedByUint := uint(updatedBy.Int64)
		product.UpdatedBy = &updatedByUint
	}

	return &product, nil
}

// GetBySKU دریافت محصول بر اساس SKU با استفاده از Prepared Statement
func (r *OptimizedProductRepository) GetBySKU(ctx context.Context, sku string) (*models.Product, error) {
	stmt, err := r.preparedStatementMgr.GetStatement("product_get_by_sku")
	if err != nil {
		// Fallback به GORM
		var product models.Product
		err := r.db.WithContext(ctx).Preload("Category").Preload("Images").Preload("Specifications").Where("sku = ?", sku).First(&product).Error
		if err != nil {
			return nil, err
		}
		return &product, nil
	}

	row := database.ExecutePreparedQueryRow(ctx, stmt, sku)

	var product models.Product
	var createdAt, updatedAt time.Time
	var deletedAt sql.NullTime
	var categoryID, brandID, updatedBy sql.NullInt64

	err = row.Scan(
		&product.ID,
		&product.UUID,
		&product.Slug,
		&product.SKU,
		&product.Name,
		&product.NameEn,
		&product.Image,
		&product.Price,
		&product.OldPrice,
		&product.Cost,
		&product.Profit,
		&product.DiscountPercent,
		&product.DiscountAmount,
		&product.SalePrice,
		&product.WholesalePrice,
		&product.WholesaleSalePrice,
		&product.Description,
		&product.FullDescription,
		&product.Status,
		&product.IsNew,
		&product.IsActive,
		&product.IsFeatured,
		&product.IsOnSale,
		&product.Weight,
		&product.Length,
		&product.Width,
		&product.Height,
		&product.ShippingCost,
		&product.ShippingTime,
		&product.StockQuantity,
		&product.MinStockQuantity,
		&product.MaxStockQuantity,
		&product.StockStatus,
		&product.TrackInventory,
		&product.ShowStockToCustomer,
		&product.DisableBuyButton,
		&product.CallForPrice,
		&product.SeoTitle,
		&product.MetaDescription,
		&product.MetaKeywords,
		&categoryID,
		&brandID,
		&product.MetaTitle,
		&product.MetaDescription,
		&product.MetaKeywords,
		&product.Tags,
		&product.Images,
		&updatedBy,
		&createdAt,
		&updatedAt,
		// &product.VideoURL,
		&deletedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("خطا در اسکن نتیجه: %w", err)
	}

	product.CreatedAt = createdAt
	product.UpdatedAt = updatedAt

	if categoryID.Valid {
		categoryIDUint := uint(categoryID.Int64)
		product.CategoryID = &categoryIDUint
	}
	if brandID.Valid {
		brandIDUint := uint(brandID.Int64)
		product.BrandID = &brandIDUint
	}
	if updatedBy.Valid {
		updatedByUint := uint(updatedBy.Int64)
		product.UpdatedBy = &updatedByUint
	}

	return &product, nil
}

// GetPublished دریافت محصولات منتشر شده با استفاده از Prepared Statement
func (r *OptimizedProductRepository) GetPublished(ctx context.Context, limit, offset int) ([]models.Product, error) {
	stmt, err := r.preparedStatementMgr.GetStatement("product_get_published")
	if err != nil {
		// Fallback به GORM
		var products []models.Product
		query := r.db.WithContext(ctx).Preload("Category").Preload("Images").Where("status = ?", "published")
		if limit > 0 {
			query = query.Limit(limit)
		}
		if offset > 0 {
			query = query.Offset(offset)
		}
		err := query.Order("created_at DESC").Find(&products).Error
		return products, err
	}

	rows, err := database.ExecutePreparedQuery(ctx, stmt, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("خطا در اجرای کوئری: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var createdAt, updatedAt time.Time
		var deletedAt sql.NullTime
		var categoryID, brandID, updatedBy sql.NullInt64

		err := rows.Scan(
			&product.ID,
			&product.UUID,
			&product.Slug,
			&product.SKU,
			&product.Name,
			&product.NameEn,
			&product.Image,
			&product.Price,
			&product.OldPrice,
			&product.Cost,
			&product.Profit,
			&product.DiscountPercent,
			&product.DiscountAmount,
			&product.SalePrice,
			&product.WholesalePrice,
			&product.WholesaleSalePrice,
			&product.Description,
			&product.FullDescription,
			&product.Status,
			&product.IsNew,
			&product.IsActive,
			&product.IsFeatured,
			&product.IsOnSale,
			&product.Weight,
			&product.Length,
			&product.Width,
			&product.Height,
			&product.ShippingCost,
			&product.ShippingTime,
			&product.StockQuantity,
			&product.MinStockQuantity,
			&product.MaxStockQuantity,
			&product.StockStatus,
			&product.TrackInventory,
			&product.ShowStockToCustomer,
			&product.DisableBuyButton,
			&product.CallForPrice,
			&product.SeoTitle,
			&product.MetaDescription,
			&product.MetaKeywords,
			&categoryID,
			&brandID,
			&product.MetaTitle,
			&product.MetaDescription,
			&product.MetaKeywords,
			&product.Tags,
			&product.Images,
			&updatedBy,
			&createdAt,
			&updatedAt,
			// &product.VideoURL,
			&deletedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("خطا در اسکن نتیجه: %w", err)
		}

		product.CreatedAt = createdAt
		product.UpdatedAt = updatedAt

		if categoryID.Valid {
			categoryIDUint := uint(categoryID.Int64)
			product.CategoryID = &categoryIDUint
		}
		if brandID.Valid {
			brandIDUint := uint(brandID.Int64)
			product.BrandID = &brandIDUint
		}
		if updatedBy.Valid {
			updatedByUint := uint(updatedBy.Int64)
			product.UpdatedBy = &updatedByUint
		}

		products = append(products, product)
	}

	return products, nil
}

// GetByCategory دریافت محصولات یک دسته‌بندی با استفاده از Prepared Statement
func (r *OptimizedProductRepository) GetByCategory(ctx context.Context, categoryID uint, limit, offset int) ([]models.Product, error) {
	stmt, err := r.preparedStatementMgr.GetStatement("product_get_by_category")
	if err != nil {
		// Fallback به GORM
		var products []models.Product
		query := r.db.WithContext(ctx).Preload("Category").Preload("Images").Where("category_id = ?", categoryID)
		if limit > 0 {
			query = query.Limit(limit)
		}
		if offset > 0 {
			query = query.Offset(offset)
		}
		err := query.Order("created_at DESC").Find(&products).Error
		return products, err
	}

	rows, err := database.ExecutePreparedQuery(ctx, stmt, categoryID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("خطا در اجرای کوئری: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var createdAt, updatedAt time.Time
		var deletedAt sql.NullTime
		var categoryID, brandID, updatedBy sql.NullInt64

		err := rows.Scan(
			&product.ID,
			&product.UUID,
			&product.Slug,
			&product.SKU,
			&product.Name,
			&product.NameEn,
			&product.Image,
			&product.Price,
			&product.OldPrice,
			&product.Cost,
			&product.Profit,
			&product.DiscountPercent,
			&product.DiscountAmount,
			&product.SalePrice,
			&product.WholesalePrice,
			&product.WholesaleSalePrice,
			&product.Description,
			&product.FullDescription,
			&product.Status,
			&product.IsNew,
			&product.IsActive,
			&product.IsFeatured,
			&product.IsOnSale,
			&product.Weight,
			&product.Length,
			&product.Width,
			&product.Height,
			&product.ShippingCost,
			&product.ShippingTime,
			&product.StockQuantity,
			&product.MinStockQuantity,
			&product.MaxStockQuantity,
			&product.StockStatus,
			&product.TrackInventory,
			&product.ShowStockToCustomer,
			&product.DisableBuyButton,
			&product.CallForPrice,
			&product.SeoTitle,
			&product.MetaDescription,
			&product.MetaKeywords,
			&categoryID,
			&brandID,
			&product.MetaTitle,
			&product.MetaDescription,
			&product.MetaKeywords,
			&product.Tags,
			&product.Images,
			&updatedBy,
			&createdAt,
			&updatedAt,
			// &product.VideoURL,
			&deletedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("خطا در اسکن نتیجه: %w", err)
		}

		product.CreatedAt = createdAt
		product.UpdatedAt = updatedAt

		if categoryID.Valid {
			categoryIDUint := uint(categoryID.Int64)
			product.CategoryID = &categoryIDUint
		}
		if brandID.Valid {
			brandIDUint := uint(brandID.Int64)
			product.BrandID = &brandIDUint
		}
		if updatedBy.Valid {
			updatedByUint := uint(updatedBy.Int64)
			product.UpdatedBy = &updatedByUint
		}

		products = append(products, product)
	}

	return products, nil
}

// GetByBrand دریافت محصولات یک برند با استفاده از Prepared Statement
func (r *OptimizedProductRepository) GetByBrand(ctx context.Context, brandID uint, limit, offset int) ([]models.Product, error) {
	stmt, err := r.preparedStatementMgr.GetStatement("product_get_by_brand")
	if err != nil {
		// Fallback به GORM
		var products []models.Product
		query := r.db.WithContext(ctx).Preload("Category").Preload("Images").Where("brand_id = ?", brandID)
		if limit > 0 {
			query = query.Limit(limit)
		}
		if offset > 0 {
			query = query.Offset(offset)
		}
		err := query.Order("created_at DESC").Find(&products).Error
		return products, err
	}

	rows, err := database.ExecutePreparedQuery(ctx, stmt, brandID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("خطا در اجرای کوئری: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var createdAt, updatedAt time.Time
		var deletedAt sql.NullTime
		var categoryID, brandID, updatedBy sql.NullInt64

		err := rows.Scan(
			&product.ID,
			&product.UUID,
			&product.Slug,
			&product.SKU,
			&product.Name,
			&product.NameEn,
			&product.Image,
			&product.Price,
			&product.OldPrice,
			&product.Cost,
			&product.Profit,
			&product.DiscountPercent,
			&product.DiscountAmount,
			&product.SalePrice,
			&product.WholesalePrice,
			&product.WholesaleSalePrice,
			&product.Description,
			&product.FullDescription,
			&product.Status,
			&product.IsNew,
			&product.IsActive,
			&product.IsFeatured,
			&product.IsOnSale,
			&product.Weight,
			&product.Length,
			&product.Width,
			&product.Height,
			&product.ShippingCost,
			&product.ShippingTime,
			&product.StockQuantity,
			&product.MinStockQuantity,
			&product.MaxStockQuantity,
			&product.StockStatus,
			&product.TrackInventory,
			&product.ShowStockToCustomer,
			&product.DisableBuyButton,
			&product.CallForPrice,
			&product.SeoTitle,
			&product.MetaDescription,
			&product.MetaKeywords,
			&categoryID,
			&brandID,
			&product.MetaTitle,
			&product.MetaDescription,
			&product.MetaKeywords,
			&product.Tags,
			&product.Images,
			&updatedBy,
			&createdAt,
			&updatedAt,
			// &product.VideoURL,
			&deletedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("خطا در اسکن نتیجه: %w", err)
		}

		product.CreatedAt = createdAt
		product.UpdatedAt = updatedAt

		if categoryID.Valid {
			categoryIDUint := uint(categoryID.Int64)
			product.CategoryID = &categoryIDUint
		}
		if brandID.Valid {
			brandIDUint := uint(brandID.Int64)
			product.BrandID = &brandIDUint
		}
		if updatedBy.Valid {
			updatedByUint := uint(updatedBy.Int64)
			product.UpdatedBy = &updatedByUint
		}

		products = append(products, product)
	}

	return products, nil
}

// GetByBrandAfter Keyset: (created_at DESC, id DESC)
func (r *OptimizedProductRepository) GetByBrandAfter(ctx context.Context, brandID uint, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error) {
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

// GetByCategoryAfter Keyset: (created_at DESC, id DESC)
func (r *OptimizedProductRepository) GetByCategoryAfter(ctx context.Context, categoryID uint, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error) {
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

// GetPublishedAfter Keyset: (created_at DESC, id DESC)
func (r *OptimizedProductRepository) GetPublishedAfter(ctx context.Context, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error) {
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

// GetNewProducts دریافت محصولات جدید (fallback به GORM)
func (r *OptimizedProductRepository) GetNewProducts(ctx context.Context, limit int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).Preload("Category").Preload("Images").Where("is_new = ?", true).Limit(limit).Order("created_at DESC").Find(&products).Error
	return products, err
}

// GetDiscountedProducts دریافت محصولات تخفیف‌دار (fallback به GORM)
func (r *OptimizedProductRepository) GetDiscountedProducts(ctx context.Context, limit int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).Preload("Category").Preload("Images").Where("sale_price < price").Limit(limit).Order("(price - sale_price) DESC").Find(&products).Error
	return products, err
}

// SearchByName جستجو در محصولات بر اساس نام (fallback به GORM)
func (r *OptimizedProductRepository) SearchByName(ctx context.Context, name string, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx).Preload("Category").Preload("Images").Where("name ILIKE ? OR description ILIKE ?", "%"+name+"%", "%"+name+"%")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("created_at DESC").Find(&products).Error
	return products, err
}

// GetByPriceRange دریافت محصولات بر اساس محدوده قیمت (fallback به GORM)
func (r *OptimizedProductRepository) GetByPriceRange(ctx context.Context, minPrice, maxPrice float64, limit, offset int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.WithContext(ctx).Preload("Category").Preload("Images").Where("price BETWEEN ? AND ?", minPrice, maxPrice)
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Order("price ASC").Find(&products).Error
	return products, err
}

// GetRelatedProducts دریافت محصولات مرتبط (fallback به GORM)
func (r *OptimizedProductRepository) GetRelatedProducts(ctx context.Context, productID uint, limit int) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).Preload("Category").Preload("Images").Where("category_id = (SELECT category_id FROM products WHERE id = ?)", productID).Where("id != ?", productID).Limit(limit).Order("created_at DESC").Find(&products).Error
	return products, err
}

// GetLowStockProducts دریافت محصولات با موجودی کم (fallback به GORM)
func (r *OptimizedProductRepository) GetLowStockProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).Preload("Category").Preload("Images").Where("stock_quantity <= min_stock_quantity").Order("stock_quantity ASC").Find(&products).Error
	return products, err
}

// Update به‌روزرسانی محصول (fallback به GORM)
func (r *OptimizedProductRepository) Update(ctx context.Context, product *models.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

// Delete حذف محصول (fallback به GORM)
func (r *OptimizedProductRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Product{}, id).Error
}

// Count شمارش محصولات (fallback به GORM)
func (r *OptimizedProductRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Product{}).Count(&count).Error
	return count, err
}

// GetAll دریافت تمام محصولات (fallback به GORM)
func (r *OptimizedProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.db.WithContext(ctx).Preload("Category").Preload("Images").Order("created_at DESC").Find(&products).Error
	return products, err
}

// UpdateStock به‌روزرسانی موجودی محصول با استفاده از Prepared Statement
func (r *OptimizedProductRepository) UpdateStock(ctx context.Context, productID uint, quantity int) error {
	stmt, err := r.preparedStatementMgr.GetStatement("product_update_stock")
	if err != nil {
		// Fallback به GORM
		return r.db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", productID).Update("stock_quantity", quantity).Error
	}

	// محاسبه وضعیت موجودی
	stockStatus := "in_stock"
	if quantity <= 0 {
		stockStatus = "out_of_stock"
	}

	_, err = database.ExecutePreparedExec(ctx, stmt, quantity, stockStatus, productID)
	return err
}
