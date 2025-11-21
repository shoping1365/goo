package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"my-go-backend/internal/cqrs/commands"
	"my-go-backend/internal/cqrs/queries"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"strings"
	"time"

	"github.com/lib/pq"
)

// ProductService سرویس برای مدیریت منطق کسب و کار محصولات
type ProductService struct {
	uowFactory unitofwork.UnitOfWorkFactory
}

// NewProductService ایجاد نمونه جدید از سرویس محصول
func NewProductService(uowFactory unitofwork.UnitOfWorkFactory) *ProductService {
	return &ProductService{
		uowFactory: uowFactory,
	}
}

// CreateProduct ایجاد محصول جدید با اعتبارسنجی و منطق کسب و کار
func (s *ProductService) CreateProduct(ctx context.Context, cmd commands.CreateProductCommand) (*models.Product, error) {
	// اعتبارسنجی داده‌ها
	if err := s.validateCreateProduct(cmd); err != nil {
		return nil, err
	}

	uow := s.uowFactory.Create()

	// شروع تراکنش
	if err := uow.BeginTx(ctx); err != nil {
		return nil, fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			if rollbackErr := uow.Rollback(); rollbackErr != nil {
				log.Printf("خطا در rollback تراکنش: %v", rollbackErr)
			}
		}
	}()

	// بررسی تکراری نبودن SKU
	if cmd.SKU != "" {
		existingProduct, err := uow.ProductRepository().GetBySKU(ctx, cmd.SKU)
		if err == nil && existingProduct != nil {
			return nil, fmt.Errorf("sku تکراری است: %s", cmd.SKU)
		}
	}

	// بررسی تکراری نبودن Slug
	if cmd.Slug != "" {
		existingProduct, err := uow.ProductRepository().GetBySlug(ctx, cmd.Slug)
		if err == nil && existingProduct != nil {
			return nil, fmt.Errorf("slug تکراری است: %s", cmd.Slug)
		}
	}

	// بررسی وجود دسته‌بندی
	_, err := uow.CategoryRepository().GetByID(ctx, cmd.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("دسته‌بندی با شناسه %d یافت نشد", cmd.CategoryID)
	}

	// بررسی وجود برند
	if cmd.BrandID != nil {
		_, err := uow.BrandRepository().GetByID(ctx, *cmd.BrandID)
		if err != nil {
			return nil, fmt.Errorf("برند با شناسه %d یافت نشد", *cmd.BrandID)
		}
	}

	// محاسبه قیمت فروش
	salePrice := cmd.Price
	if cmd.SalePrice != nil {
		salePrice = *cmd.SalePrice
	}

	// محاسبه درصد تخفیف
	discountPercent := 0.0
	if cmd.Price > 0 && salePrice < cmd.Price {
		discountPercent = ((cmd.Price - salePrice) / cmd.Price) * 100
	}

	// تبدیل Tags از []string به string
	tagsString := ""
	if len(cmd.Tags) > 0 {
		tagsString = strings.Join(cmd.Tags, ",")
	}

	// تبدیل Images از []string به []ProductImage
	var productImages []models.ProductImage
	for i, imageURL := range cmd.Images {
		productImage := models.ProductImage{
			ImageURL:  imageURL,
			SortOrder: i + 1,
			CreatedAt: time.Now(),
		}
		productImages = append(productImages, productImage)
	}

	// ایجاد محصول
	product := &models.Product{
		UUID:                s.generateUUID(),
		Slug:                cmd.Slug,
		SKU:                 cmd.SKU,
		Name:                cmd.Name,
		Price:               cmd.Price,
		SalePrice:           salePrice,
		Cost:                cmd.CostPrice,
		DiscountPercent:     discountPercent,
		DiscountAmount:      cmd.Price - salePrice,
		Description:         cmd.Description,
		FullDescription:     cmd.ShortDescription,
		Status:              cmd.Status,
		IsNew:               cmd.IsNew,
		IsActive:            cmd.IsActive,
		IsFeatured:          cmd.IsFeatured,
		IsOnSale:            cmd.IsOnSale,
		Weight:              cmd.Weight,
		StockQuantity:       cmd.StockQuantity,
		MinStockQuantity:    0, // مقدار پیش‌فرض
		MaxStockQuantity:    0, // مقدار پیش‌فرض
		StockStatus:         s.calculateStockStatus(cmd.StockQuantity, 0),
		TrackInventory:      true,  // مقدار پیش‌فرض
		ShowStockToCustomer: true,  // مقدار پیش‌فرض
		DisableBuyButton:    false, // مقدار پیش‌فرض
		CallForPrice:        false, // مقدار پیش‌فرض
		SeoTitle:            cmd.MetaTitle,
		MetaDescription:     cmd.MetaDescription,
		MetaKeywords:        cmd.MetaKeywords,
		CategoryID:          &cmd.CategoryID, // تبدیل به *uint
		BrandID:             cmd.BrandID,
		Tags:                tagsString,
		Images:              productImages,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	// ذخیره محصول با Retry روی خطای یکتا بودن SKU برای سرعت (بدون COUNT)
	const maxCreateAttempts = 5
	var createErr error
	for attempt := 1; attempt <= maxCreateAttempts; attempt++ {
		// اگر SKU مشخص نشده بود، اجازه بده BeforeCreate آن را بسازد
		if cmd.SKU == "" {
			product.SKU = ""
		}
		createErr = uow.ProductRepository().Create(ctx, product)
		if createErr == nil {
			break
		}
		var pgErr *pq.Error
		if errors.As(createErr, &pgErr) && pgErr.Code == "23505" {
			// برخورد یکتا؛ تلاش مجدد با SKU تازه
			if attempt < maxCreateAttempts {
				continue
			}
		}
		// سایر خطاها یا اتمام تلاش‌ها
		return nil, fmt.Errorf("خطا در ایجاد محصول: %w", createErr)
	}

	// Commit تراکنش
	if err := uow.Commit(); err != nil {
		return nil, fmt.Errorf("خطا در ذخیره تراکنش: %w", err)
	}

	committed = true
	return product, nil
}

// UpdateProduct به‌روزرسانی محصول با اعتبارسنجی و منطق کسب و کار
func (s *ProductService) UpdateProduct(ctx context.Context, cmd commands.UpdateProductCommand) error {
	uow := s.uowFactory.Create()

	// شروع تراکنش
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			if rollbackErr := uow.Rollback(); rollbackErr != nil {
				log.Printf("خطا در rollback تراکنش: %v", rollbackErr)
			}
		}
	}()

	// دریافت محصول موجود
	product, err := uow.ProductRepository().GetByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("محصول با شناسه %d یافت نشد", cmd.ID)
	}

	// اعتبارسنجی تغییرات
	if err := s.validateUpdateProduct(cmd, product); err != nil {
		return err
	}

	// بررسی تکراری نبودن SKU
	if cmd.SKU != "" && cmd.SKU != product.SKU {
		existingProduct, err := uow.ProductRepository().GetBySKU(ctx, cmd.SKU)
		if err == nil && existingProduct != nil && existingProduct.ID != int(cmd.ID) {
			return fmt.Errorf("SKU تکراری است: %s", cmd.SKU)
		}
	}

	// بررسی تکراری نبودن Slug
	if cmd.Slug != "" && cmd.Slug != product.Slug {
		existingProduct, err := uow.ProductRepository().GetBySlug(ctx, cmd.Slug)
		if err == nil && existingProduct != nil && existingProduct.ID != int(cmd.ID) {
			return fmt.Errorf("Slug تکراری است: %s", cmd.Slug)
		}
	}

	// بررسی وجود دسته‌بندی
	if product.CategoryID == nil || cmd.CategoryID != *product.CategoryID {
		_, err := uow.CategoryRepository().GetByID(ctx, cmd.CategoryID)
		if err != nil {
			return fmt.Errorf("دسته‌بندی با شناسه %d یافت نشد", cmd.CategoryID)
		}
	}

	// بررسی وجود برند
	if cmd.BrandID != nil && (product.BrandID == nil || *cmd.BrandID != *product.BrandID) {
		_, err := uow.BrandRepository().GetByID(ctx, *cmd.BrandID)
		if err != nil {
			return fmt.Errorf("برند با شناسه %d یافت نشد", *cmd.BrandID)
		}
	}

	// محاسبه قیمت فروش
	salePrice := cmd.Price
	if cmd.SalePrice != nil {
		salePrice = *cmd.SalePrice
	}

	// محاسبه درصد تخفیف
	discountPercent := 0.0
	if cmd.Price > 0 && salePrice < cmd.Price {
		discountPercent = ((cmd.Price - salePrice) / cmd.Price) * 100
	}

	// تبدیل Tags از []string به string
	tagsString := ""
	if len(cmd.Tags) > 0 {
		tagsString = strings.Join(cmd.Tags, ",")
	}

	// تبدیل Images از []string به []ProductImage
	var productImages []models.ProductImage
	for i, imageURL := range cmd.Images {
		productImage := models.ProductImage{
			ImageURL:  imageURL,
			SortOrder: i + 1,
			CreatedAt: time.Now(),
		}
		productImages = append(productImages, productImage)
	}

	// به‌روزرسانی فیلدهای محصول
	product.Slug = cmd.Slug
	product.SKU = cmd.SKU
	product.Name = cmd.Name
	product.Price = cmd.Price
	product.SalePrice = salePrice
	product.Cost = cmd.CostPrice
	product.DiscountPercent = discountPercent
	product.DiscountAmount = cmd.Price - salePrice
	product.Description = cmd.Description
	product.FullDescription = cmd.ShortDescription
	product.Status = cmd.Status
	product.IsNew = cmd.IsNew
	product.IsActive = cmd.IsActive
	product.IsFeatured = cmd.IsFeatured
	product.IsOnSale = cmd.IsOnSale
	product.Weight = cmd.Weight
	product.StockQuantity = cmd.StockQuantity
	product.StockStatus = s.calculateStockStatus(cmd.StockQuantity, product.MinStockQuantity)
	product.SeoTitle = cmd.MetaTitle
	product.MetaDescription = cmd.MetaDescription
	product.MetaKeywords = cmd.MetaKeywords
	product.CategoryID = &cmd.CategoryID
	product.BrandID = cmd.BrandID
	product.Tags = tagsString
	product.Images = productImages
	product.UpdatedAt = time.Now()

	// ذخیره تغییرات
	err = uow.ProductRepository().Update(ctx, product)
	if err != nil {
		return fmt.Errorf("خطا در به‌روزرسانی محصول: %w", err)
	}

	// Commit تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در ذخیره تراکنش: %w", err)
	}

	committed = true
	return nil
}

// DeleteProduct حذف محصول با بررسی وابستگی‌ها
func (s *ProductService) DeleteProduct(ctx context.Context, cmd commands.DeleteProductCommand) error {
	uow := s.uowFactory.Create()

	// شروع تراکنش
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			if rollbackErr := uow.Rollback(); rollbackErr != nil {
				log.Printf("خطا در rollback تراکنش: %v", rollbackErr)
			}
		}
	}()

	// بررسی وجود محصول
	_, err := uow.ProductRepository().GetByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("محصول با شناسه %d یافت نشد", cmd.ID)
	}

	// بررسی وابستگی‌ها (نظرات، سوالات، اعلان‌های موجودی)
	reviews, err := uow.ReviewRepository().GetByProduct(ctx, cmd.ID, 1, 1)
	if err == nil && len(reviews) > 0 {
		return fmt.Errorf("محصول دارای نظرات است و قابل حذف نیست")
	}

	notifications, err := uow.StockNotificationRepository().GetByProduct(ctx, cmd.ID)
	if err == nil && len(notifications) > 0 {
		return fmt.Errorf("محصول دارای اعلان‌های موجودی است و قابل حذف نیست")
	}

	// حذف محصول
	err = uow.ProductRepository().Delete(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("خطا در حذف محصول: %w", err)
	}

	// Commit تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در ذخیره تراکنش: %w", err)
	}

	committed = true
	return nil
}

// UpdateProductStock به‌روزرسانی موجودی محصول
func (s *ProductService) UpdateProductStock(ctx context.Context, cmd commands.UpdateProductStockCommand) error {
	uow := s.uowFactory.Create()

	// شروع تراکنش
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			if rollbackErr := uow.Rollback(); rollbackErr != nil {
				log.Printf("خطا در rollback تراکنش: %v", rollbackErr)
			}
		}
	}()

	// دریافت محصول
	product, err := uow.ProductRepository().GetByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("محصول با شناسه %d یافت نشد", cmd.ID)
	}

	// به‌روزرسانی موجودی
	product.StockQuantity = cmd.Quantity
	product.StockStatus = s.calculateStockStatus(cmd.Quantity, product.MinStockQuantity)
	product.UpdatedAt = time.Now()

	// بررسی اعلان‌های موجودی
	if cmd.Quantity > 0 {
		notifications, err := uow.StockNotificationRepository().GetByProduct(ctx, cmd.ID)
		if err == nil && len(notifications) > 0 {
			// ارسال اعلان‌ها
			for _, notification := range notifications {
				notification.Status = "sent"
				notification.SentAt = &time.Time{}
				uow.StockNotificationRepository().Update(ctx, &notification)
			}
		}
	}

	// ذخیره تغییرات
	err = uow.ProductRepository().Update(ctx, product)
	if err != nil {
		return fmt.Errorf("خطا در به‌روزرسانی موجودی: %w", err)
	}

	// Commit تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در ذخیره تراکنش: %w", err)
	}

	committed = true
	return nil
}

// PublishProduct انتشار محصول
func (s *ProductService) PublishProduct(ctx context.Context, cmd commands.PublishProductCommand) error {
	uow := s.uowFactory.Create()

	// شروع تراکنش
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			if rollbackErr := uow.Rollback(); rollbackErr != nil {
				log.Printf("خطا در rollback تراکنش: %v", rollbackErr)
			}
		}
	}()

	// دریافت محصول
	product, err := uow.ProductRepository().GetByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("محصول با شناسه %d یافت نشد", cmd.ID)
	}

	// بررسی شرایط انتشار
	if product.Status == "draft" {
		if product.Name == "" {
			return fmt.Errorf("نام محصول الزامی است")
		}
		if product.Price <= 0 {
			return fmt.Errorf("قیمت محصول باید بزرگتر از صفر باشد")
		}
		if product.StockQuantity < 0 {
			return fmt.Errorf("موجودی محصول نمی‌تواند منفی باشد")
		}
	}

	// انتشار محصول
	product.Status = "published"
	product.UpdatedAt = time.Now()
	if cmd.PublishedBy != 0 {
		product.UpdatedBy = &cmd.PublishedBy
	}

	// ذخیره تغییرات
	err = uow.ProductRepository().Update(ctx, product)
	if err != nil {
		return fmt.Errorf("خطا در انتشار محصول: %w", err)
	}

	// Commit تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در ذخیره تراکنش: %w", err)
	}

	committed = true
	return nil
}

// ArchiveProduct آرشیو محصول
func (s *ProductService) ArchiveProduct(ctx context.Context, cmd commands.ArchiveProductCommand) error {
	uow := s.uowFactory.Create()

	// شروع تراکنش
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			if rollbackErr := uow.Rollback(); rollbackErr != nil {
				log.Printf("خطا در rollback تراکنش: %v", rollbackErr)
			}
		}
	}()

	// دریافت محصول
	product, err := uow.ProductRepository().GetByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("محصول با شناسه %d یافت نشد", cmd.ID)
	}

	// آرشیو محصول
	product.Status = "archived"
	product.UpdatedAt = time.Now()
	if cmd.ArchivedBy != 0 {
		product.UpdatedBy = &cmd.ArchivedBy
	}

	// ذخیره تغییرات
	err = uow.ProductRepository().Update(ctx, product)
	if err != nil {
		return fmt.Errorf("خطا در آرشیو محصول: %w", err)
	}

	// Commit تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در ذخیره تراکنش: %w", err)
	}

	committed = true
	return nil
}

// BulkUpdateProductStatus به‌روزرسانی دسته‌ای وضعیت محصولات
func (s *ProductService) BulkUpdateProductStatus(ctx context.Context, cmd commands.BulkUpdateProductStatusCommand) error {
	uow := s.uowFactory.Create()

	// شروع تراکنش
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			if rollbackErr := uow.Rollback(); rollbackErr != nil {
				log.Printf("خطا در rollback تراکنش: %v", rollbackErr)
			}
		}
	}()

	// اعتبارسنجی وضعیت
	validStatuses := []string{"draft", "published", "archived", "out_of_stock"}
	isValidStatus := false
	for _, status := range validStatuses {
		if cmd.Status == status {
			isValidStatus = true
			break
		}
	}
	if !isValidStatus {
		return fmt.Errorf("وضعیت نامعتبر است: %s", cmd.Status)
	}

	// به‌روزرسانی محصولات
	for _, productID := range cmd.ProductIDs {
		product, err := uow.ProductRepository().GetByID(ctx, productID)
		if err != nil {
			continue // رد کردن محصولات ناموجود
		}

		product.Status = cmd.Status
		product.UpdatedAt = time.Now()
		if cmd.UpdatedBy != 0 {
			product.UpdatedBy = &cmd.UpdatedBy
		}

		uow.ProductRepository().Update(ctx, product)
	}

	// Commit تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در ذخیره تراکنش: %w", err)
	}

	committed = true
	return nil
}

// GetProducts دریافت محصولات با فیلتر و صفحه‌بندی
func (s *ProductService) GetProducts(ctx context.Context, query queries.GetProductsQuery) ([]models.Product, int64, error) {
	uow := s.uowFactory.Create()

	// اعتبارسنجی پارامترها
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 || query.PageSize > 100 {
		query.PageSize = 20
	}

	// تبدیل فیلتر
	filter := s.convertProductFilter(query.Filter)

	// TODO: پیاده‌سازی فیلتر و صفحه‌بندی در repository
	// فعلاً از GetAll استفاده می‌کنیم
	allProducts, err := uow.ProductRepository().GetAll(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("خطا در دریافت محصولات: %w", err)
	}

	// فیلتر کردن در حافظه (موقت)
	var filteredProducts []models.Product
	for _, product := range allProducts {
		if s.matchesFilter(product, filter) {
			filteredProducts = append(filteredProducts, product)
		}
	}

	// صفحه‌بندی
	total := int64(len(filteredProducts))
	start := (query.Page - 1) * query.PageSize
	end := start + query.PageSize
	if start >= len(filteredProducts) {
		return []models.Product{}, total, nil
	}
	if end > len(filteredProducts) {
		end = len(filteredProducts)
	}

	products := filteredProducts[start:end]

	return products, total, nil
}

// GetProductByID دریافت محصول بر اساس ID
func (s *ProductService) GetProductByID(ctx context.Context, query queries.GetProductByIDQuery) (*models.Product, error) {
	uow := s.uowFactory.Create()

	product, err := uow.ProductRepository().GetByID(ctx, query.ID)
	if err != nil {
		return nil, fmt.Errorf("محصول با شناسه %d یافت نشد", query.ID)
	}

	return product, nil
}

// GetProductBySlug دریافت محصول بر اساس Slug
func (s *ProductService) GetProductBySlug(ctx context.Context, query queries.GetProductBySlugQuery) (*models.Product, error) {
	uow := s.uowFactory.Create()

	product, err := uow.ProductRepository().GetBySlug(ctx, query.Slug)
	if err != nil {
		return nil, fmt.Errorf("محصول با slug %s یافت نشد", query.Slug)
	}

	return product, nil
}

// GetProductBySKU دریافت محصول بر اساس SKU
func (s *ProductService) GetProductBySKU(ctx context.Context, query queries.GetProductBySKUQuery) (*models.Product, error) {
	uow := s.uowFactory.Create()

	product, err := uow.ProductRepository().GetBySKU(ctx, query.SKU)
	if err != nil {
		return nil, fmt.Errorf("محصول با SKU %s یافت نشد", query.SKU)
	}

	return product, nil
}

// ==================== Helper Methods ====================

// validateCreateProduct اعتبارسنجی داده‌های ایجاد محصول
func (s *ProductService) validateCreateProduct(cmd commands.CreateProductCommand) error {
	if cmd.Name == "" {
		return fmt.Errorf("نام محصول الزامی است")
	}
	if cmd.Price < 0 {
		return fmt.Errorf("قیمت محصول نمی‌تواند منفی باشد")
	}
	if cmd.StockQuantity < 0 {
		return fmt.Errorf("موجودی محصول نمی‌تواند منفی باشد")
	}
	if cmd.Slug == "" {
		return fmt.Errorf("Slug محصول الزامی است")
	}
	return nil
}

// validateUpdateProduct اعتبارسنجی داده‌های به‌روزرسانی محصول
func (s *ProductService) validateUpdateProduct(cmd commands.UpdateProductCommand, product *models.Product) error {
	if cmd.Name != "" && len(cmd.Name) < 2 {
		return fmt.Errorf("نام محصول باید حداقل 2 کاراکتر باشد")
	}
	if cmd.Price < 0 {
		return fmt.Errorf("قیمت محصول نمی‌تواند منفی باشد")
	}
	if cmd.StockQuantity < 0 {
		return fmt.Errorf("موجودی محصول نمی‌تواند منفی باشد")
	}
	return nil
}

// calculateStockStatus محاسبه وضعیت موجودی
func (s *ProductService) calculateStockStatus(quantity, minQuantity int) string {
	if quantity <= 0 {
		return "out_of_stock"
	}
	if quantity <= minQuantity {
		return "low_stock"
	}
	return "in_stock"
}

// generateUUID تولید UUID منحصر به فرد
func (s *ProductService) generateUUID() string {
	// پیاده‌سازی ساده برای تولید UUID
	// در محیط تولید باید از کتابخانه UUID استفاده شود
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// convertProductFilter تبدیل فیلتر کوئری به فیلتر repository
func (s *ProductService) convertProductFilter(filter queries.ProductFilter) map[string]interface{} {
	result := make(map[string]interface{})

	if filter.Name != "" {
		result["name"] = filter.Name
	}
	if filter.SKU != "" {
		result["sku"] = filter.SKU
	}
	if filter.CategoryID != nil {
		result["category_id"] = *filter.CategoryID
	}
	if filter.BrandID != nil {
		result["brand_id"] = *filter.BrandID
	}
	if filter.Status != "" {
		result["status"] = filter.Status
	}
	if filter.IsActive != nil {
		result["is_active"] = *filter.IsActive
	}
	if filter.IsFeatured != nil {
		result["is_featured"] = *filter.IsFeatured
	}
	if filter.IsNew != nil {
		result["is_new"] = *filter.IsNew
	}
	if filter.IsOnSale != nil {
		result["is_on_sale"] = *filter.IsOnSale
	}
	if filter.MinPrice != nil {
		result["min_price"] = *filter.MinPrice
	}
	if filter.MaxPrice != nil {
		result["max_price"] = *filter.MaxPrice
	}
	if filter.InStock != nil {
		result["in_stock"] = *filter.InStock
	}
	if filter.LowStock != nil {
		result["low_stock"] = *filter.LowStock
	}

	return result
}

// matchesFilter بررسی تطبیق محصول با فیلتر
func (s *ProductService) matchesFilter(product models.Product, filter map[string]interface{}) bool {
	for key, value := range filter {
		switch key {
		case "name":
			if name, ok := value.(string); ok && name != "" {
				if !strings.Contains(strings.ToLower(product.Name), strings.ToLower(name)) {
					return false
				}
			}
		case "sku":
			if sku, ok := value.(string); ok && sku != "" {
				if !strings.Contains(strings.ToLower(product.SKU), strings.ToLower(sku)) {
					return false
				}
			}
		case "category_id":
			if categoryID, ok := value.(uint); ok {
				if product.CategoryID == nil || *product.CategoryID != categoryID {
					return false
				}
			}
		case "brand_id":
			if brandID, ok := value.(uint); ok {
				if product.BrandID == nil || *product.BrandID != brandID {
					return false
				}
			}
		case "status":
			if status, ok := value.(string); ok && status != "" {
				if product.Status != status {
					return false
				}
			}
		case "is_active":
			if isActive, ok := value.(bool); ok {
				if product.IsActive != isActive {
					return false
				}
			}
		case "is_featured":
			if isFeatured, ok := value.(bool); ok {
				if product.IsFeatured != isFeatured {
					return false
				}
			}
		case "is_new":
			if isNew, ok := value.(bool); ok {
				if product.IsNew != isNew {
					return false
				}
			}
		case "is_on_sale":
			if isOnSale, ok := value.(bool); ok {
				if product.IsOnSale != isOnSale {
					return false
				}
			}
		case "min_price":
			if minPrice, ok := value.(float64); ok {
				if product.Price < minPrice {
					return false
				}
			}
		case "max_price":
			if maxPrice, ok := value.(float64); ok {
				if product.Price > maxPrice {
					return false
				}
			}
		case "in_stock":
			if inStock, ok := value.(bool); ok {
				hasStock := product.StockQuantity > 0
				if hasStock != inStock {
					return false
				}
			}
		case "low_stock":
			if lowStock, ok := value.(bool); ok {
				isLowStock := product.StockQuantity <= product.MinStockQuantity && product.StockQuantity > 0
				if isLowStock != lowStock {
					return false
				}
			}
		}
	}
	return true
}
