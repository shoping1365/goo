package bus

import (
	"context"
	"fmt"
	"my-go-backend/internal/cqrs/handlers"
	"my-go-backend/internal/cqrs/queries"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
)

// QueryBus رابط برای ارسال کوئری‌ها به handler مناسب
// این bus مسئول routing کوئری‌ها به handler صحیح است
type QueryBus interface {
	// Media File Queries
	GetMediaFileByID(ctx context.Context, id uint) (*models.MediaFile, error)
	GetMediaFileByPath(ctx context.Context, path string) (*models.MediaFile, error)
	GetMediaFilesByUser(ctx context.Context, userID uint) ([]models.MediaFile, error)
	GetMediaFilesByType(ctx context.Context, fileType string) ([]models.MediaFile, error)
	GetMediaFilesByCategory(ctx context.Context, category string) ([]models.MediaFile, error)
	GetMediaFilesPaged(ctx context.Context, page, pageSize int, filter queries.MediaFileFilter) ([]models.MediaFile, int64, error)
	SearchMediaFiles(ctx context.Context, query string) ([]models.MediaFile, error)
	GetRecentUploads(ctx context.Context, limit int) ([]models.MediaFile, error)
	GetMediaFileStats(ctx context.Context, userID uint) (*repository.MediaFileStats, error)
	GetMediaFileWithVariants(ctx context.Context, id uint) (*models.MediaFile, error)

	// Media Version Queries
	GetMediaVersionByID(ctx context.Context, id uint) (*models.MediaVersion, error)
	GetMediaVersionsByMediaID(ctx context.Context, mediaID uint) ([]models.MediaVersion, error)
	GetActiveMediaVersion(ctx context.Context, mediaID uint) (*models.MediaVersion, error)
	GetMediaBackups(ctx context.Context, mediaID uint) ([]models.MediaVersion, error)
	GetMediaVersionsByFormat(ctx context.Context, mediaID uint, format string) ([]models.MediaVersion, error)
	GetMediaVersionsByQuality(ctx context.Context, mediaID uint, quality string) ([]models.MediaVersion, error)
	GetMediaVersionHistory(ctx context.Context, mediaID uint) ([]models.MediaVersion, error)
	GetTotalSizeByMedia(ctx context.Context, mediaID uint) (uint64, error)

	// Compression Job Queries
	GetCompressionJobByID(ctx context.Context, id uint) (*models.CompressionJob, error)
	GetCompressionJobsByMedia(ctx context.Context, mediaID uint) ([]models.CompressionJob, error)
	GetCompressionJobsByStatus(ctx context.Context, status string) ([]models.CompressionJob, error)
	GetPendingCompressionJobs(ctx context.Context, limit int) ([]models.CompressionJob, error)

	// Compression Setting Queries
	GetCompressionSettingByID(ctx context.Context, id uint) (*models.CompressionSetting, error)
	GetCompressionSettingsByScope(ctx context.Context, scope string) ([]models.CompressionSetting, error)
	GetCompressionSettingsByUser(ctx context.Context, userID uint) ([]models.CompressionSetting, error)
	GetCompressionSettingsByMedia(ctx context.Context, mediaID uint) ([]models.CompressionSetting, error)

	// Compression Stats Queries
	GetCompressionStatsByType(ctx context.Context, statType string) ([]models.CompressionStat, error)
	GetCompressionStatsByMedia(ctx context.Context, mediaID uint) ([]models.CompressionStat, error)
	GetCompressionStatsByUser(ctx context.Context, userID uint) ([]models.CompressionStat, error)

	// Report Queries
	GetMediaStorageReport(ctx context.Context, groupBy string, limit int) ([]handlers.MediaStorageReport, error)

	// ==================== Product Queries ====================
	GetProductByID(ctx context.Context, id uint) (*models.Product, error)
	GetProductBySlug(ctx context.Context, slug string) (*models.Product, error)
	GetProductBySKU(ctx context.Context, sku string) (*models.Product, error)
	GetProducts(ctx context.Context, page, pageSize int, filter queries.ProductFilter) ([]models.Product, int64, error)
	GetProductsByCategory(ctx context.Context, categoryID uint, page, pageSize int) ([]models.Product, int64, error)
	GetProductsByBrand(ctx context.Context, brandID uint, page, pageSize int) ([]models.Product, int64, error)
	GetPublishedProducts(ctx context.Context, page, pageSize int) ([]models.Product, int64, error)
	GetNewProducts(ctx context.Context, limit int) ([]models.Product, error)
	GetDiscountedProducts(ctx context.Context, page, pageSize int) ([]models.Product, int64, error)
	GetFeaturedProducts(ctx context.Context, limit int) ([]models.Product, error)
	SearchProducts(ctx context.Context, searchTerm string, page, pageSize int) ([]models.Product, int64, error)
	GetProductsByPriceRange(ctx context.Context, minPrice, maxPrice float64, page, pageSize int) ([]models.Product, int64, error)
	GetRelatedProducts(ctx context.Context, productID uint, limit int) ([]models.Product, error)
	GetLowStockProducts(ctx context.Context, threshold int) ([]models.Product, error)
	GetProductStats(ctx context.Context) (map[string]interface{}, error)

	// ==================== Category Queries ====================
	GetCategoryByID(ctx context.Context, id uint) (*models.Category, error)
	GetCategoryBySlug(ctx context.Context, slug string) (*models.Category, error)
	GetCategories(ctx context.Context, page, pageSize int, filter queries.CategoryFilter) ([]models.Category, int64, error)
	GetRootCategories(ctx context.Context) ([]models.Category, error)
	GetCategoryHierarchy(ctx context.Context) ([]models.Category, error)
	GetCategoriesWithProductCount(ctx context.Context) ([]models.Category, error)
	GetFeaturedCategories(ctx context.Context, limit int) ([]models.Category, error)

	// ==================== Brand Queries ====================
	GetBrandByID(ctx context.Context, id uint) (*models.Brand, error)
	GetBrandBySlug(ctx context.Context, slug string) (*models.Brand, error)
	GetBrands(ctx context.Context, page, pageSize int, filter queries.BrandFilter) ([]models.Brand, int64, error)
	GetPublishedBrands(ctx context.Context) ([]models.Brand, error)
	GetFeaturedBrands(ctx context.Context, limit int) ([]models.Brand, error)
	SearchBrands(ctx context.Context, searchTerm string) ([]models.Brand, error)

	// ==================== Review Queries ====================
	GetReviewByID(ctx context.Context, id uint) (*models.Review, error)
	GetReviews(ctx context.Context, page, pageSize int, filter queries.ReviewFilter) ([]models.Review, int64, error)
	GetReviewsByProduct(ctx context.Context, productID uint, page, pageSize int) ([]models.Review, int64, error)
	GetReviewsByUser(ctx context.Context, userID uint, page, pageSize int) ([]models.Review, int64, error)
	GetPublishedReviews(ctx context.Context, page, pageSize int) ([]models.Review, int64, error)
	GetPendingReviews(ctx context.Context, page, pageSize int) ([]models.Review, int64, error)
	GetAverageRating(ctx context.Context, productID uint) (float64, error)
	GetReviewCount(ctx context.Context, productID uint) (int64, error)

	// ==================== Wishlist Queries ====================
	GetWishlistByUser(ctx context.Context, userID uint) ([]models.Wishlist, error)
	GetWishlistProducts(ctx context.Context, userID uint) ([]models.Product, error)
	IsInWishlist(ctx context.Context, userID, productID uint) (bool, error)
	GetWishlistCount(ctx context.Context, userID uint) (int64, error)

	// ==================== Stock Notification Queries ====================
	GetStockNotifications(ctx context.Context, page, pageSize int, filter queries.StockNotificationFilter) ([]models.StockNotification, int64, error)
	GetStockNotificationsByProduct(ctx context.Context, productID uint) ([]models.StockNotification, error)
	GetStockNotificationsByUser(ctx context.Context, userID uint) ([]models.StockNotification, error)
	GetPendingStockNotifications(ctx context.Context) ([]models.StockNotification, error)

	// ==================== Product QA Queries ====================
	GetProductQA(ctx context.Context, page, pageSize int, filter queries.ProductQAFilter) ([]models.ProductQA, int64, error)
	GetQuestionsByProduct(ctx context.Context, productID uint, page, pageSize int) ([]models.ProductQA, int64, error)
	GetAnswersByQuestion(ctx context.Context, questionID uint) ([]models.ProductQA, error)
	GetPublishedQA(ctx context.Context, page, pageSize int) ([]models.ProductQA, int64, error)
	GetPendingQA(ctx context.Context, page, pageSize int) ([]models.ProductQA, int64, error)
	GetQuestionCount(ctx context.Context, productID uint) (int64, error)
	GetAnswerCount(ctx context.Context, questionID uint) (int64, error)
}

// queryBusImpl پیاده‌سازی QueryBus
type queryBusImpl struct {
	mediaHandler   *handlers.MediaQueryHandler
	productHandler *handlers.ProductQueryHandler
}

// NewQueryBus ایجاد نمونه جدید از QueryBus
func NewQueryBus(mediaHandler *handlers.MediaQueryHandler, productHandler *handlers.ProductQueryHandler) QueryBus {
	return &queryBusImpl{
		mediaHandler:   mediaHandler,
		productHandler: productHandler,
	}
}

// GetMediaFileByID دریافت فایل رسانه بر اساس ID
func (b *queryBusImpl) GetMediaFileByID(ctx context.Context, id uint) (*models.MediaFile, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه فایل نامعتبر است")
	}
	query := queries.GetMediaFileByIDQuery{ID: id}
	return b.mediaHandler.HandleGetMediaFileByID(ctx, query)
}

// GetMediaFileByPath دریافت فایل رسانه بر اساس مسیر
func (b *queryBusImpl) GetMediaFileByPath(ctx context.Context, path string) (*models.MediaFile, error) {
	if path == "" {
		return nil, fmt.Errorf("مسیر فایل نمی‌تواند خالی باشد")
	}
	query := queries.GetMediaFileByPathQuery{Path: path}
	return b.mediaHandler.HandleGetMediaFileByPath(ctx, query)
}

// GetMediaFilesByUser دریافت فایل‌های کاربر
func (b *queryBusImpl) GetMediaFilesByUser(ctx context.Context, userID uint) ([]models.MediaFile, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	query := queries.GetMediaFilesByUserQuery{UserID: userID}
	return b.mediaHandler.HandleGetMediaFilesByUser(ctx, query)
}

// GetMediaFilesByType دریافت فایل‌ها بر اساس نوع
func (b *queryBusImpl) GetMediaFilesByType(ctx context.Context, fileType string) ([]models.MediaFile, error) {
	if fileType == "" {
		return nil, fmt.Errorf("نوع فایل نمی‌تواند خالی باشد")
	}
	query := queries.GetMediaFilesByTypeQuery{FileType: fileType}
	return b.mediaHandler.HandleGetMediaFilesByType(ctx, query)
}

// GetMediaFilesByCategory دریافت فایل‌ها بر اساس دسته‌بندی
func (b *queryBusImpl) GetMediaFilesByCategory(ctx context.Context, category string) ([]models.MediaFile, error) {
	query := queries.GetMediaFilesByCategoryQuery{Category: category}
	return b.mediaHandler.HandleGetMediaFilesByCategory(ctx, query)
}

// GetMediaFilesPaged دریافت فایل‌ها به صورت صفحه‌بندی شده
func (b *queryBusImpl) GetMediaFilesPaged(ctx context.Context, page, pageSize int, filter queries.MediaFileFilter) ([]models.MediaFile, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	query := queries.GetMediaFilesPagedQuery{
		Page:     page,
		PageSize: pageSize,
		Filter:   filter,
	}
	return b.mediaHandler.HandleGetMediaFilesPaged(ctx, query)
}

// SearchMediaFiles جستجو در فایل‌های رسانه
func (b *queryBusImpl) SearchMediaFiles(ctx context.Context, searchQuery string) ([]models.MediaFile, error) {
	if searchQuery == "" {
		return nil, fmt.Errorf("عبارت جستجو نمی‌تواند خالی باشد")
	}
	query := queries.SearchMediaFilesQuery{Query: searchQuery}
	return b.mediaHandler.HandleSearchMediaFiles(ctx, query)
}

// GetRecentUploads دریافت آخرین آپلودها
func (b *queryBusImpl) GetRecentUploads(ctx context.Context, limit int) ([]models.MediaFile, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	query := queries.GetRecentUploadsQuery{Limit: limit}
	return b.mediaHandler.HandleGetRecentUploads(ctx, query)
}

// GetMediaFileStats دریافت آمار فایل‌های کاربر
func (b *queryBusImpl) GetMediaFileStats(ctx context.Context, userID uint) (*repository.MediaFileStats, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	query := queries.GetMediaFileStatsQuery{UserID: userID}
	return b.mediaHandler.HandleGetMediaFileStats(ctx, query)
}

// GetMediaFileWithVariants دریافت فایل با واریانت‌ها
func (b *queryBusImpl) GetMediaFileWithVariants(ctx context.Context, id uint) (*models.MediaFile, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه فایل نامعتبر است")
	}
	query := queries.GetMediaFileWithVariantsQuery{ID: id}
	return b.mediaHandler.HandleGetMediaFileWithVariants(ctx, query)
}

// GetMediaVersionByID دریافت نسخه بر اساس ID
func (b *queryBusImpl) GetMediaVersionByID(ctx context.Context, id uint) (*models.MediaVersion, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه نسخه نامعتبر است")
	}
	query := queries.GetMediaVersionByIDQuery{ID: id}
	return b.mediaHandler.HandleGetMediaVersionByID(ctx, query)
}

// GetMediaVersionsByMediaID دریافت نسخه‌های یک رسانه
func (b *queryBusImpl) GetMediaVersionsByMediaID(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	query := queries.GetMediaVersionsByMediaIDQuery{MediaID: mediaID}
	return b.mediaHandler.HandleGetMediaVersionsByMediaID(ctx, query)
}

// GetActiveMediaVersion دریافت نسخه فعال
func (b *queryBusImpl) GetActiveMediaVersion(ctx context.Context, mediaID uint) (*models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	query := queries.GetActiveMediaVersionQuery{MediaID: mediaID}
	return b.mediaHandler.HandleGetActiveMediaVersion(ctx, query)
}

// GetMediaBackups دریافت بکاپ‌ها
func (b *queryBusImpl) GetMediaBackups(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	query := queries.GetMediaBackupsQuery{MediaID: mediaID}
	return b.mediaHandler.HandleGetMediaBackups(ctx, query)
}

// GetMediaVersionsByFormat دریافت نسخه‌ها بر اساس فرمت
func (b *queryBusImpl) GetMediaVersionsByFormat(ctx context.Context, mediaID uint, format string) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	if format == "" {
		return nil, fmt.Errorf("فرمت نمی‌تواند خالی باشد")
	}
	query := queries.GetMediaVersionsByFormatQuery{MediaID: mediaID, Format: format}
	return b.mediaHandler.HandleGetMediaVersionsByFormat(ctx, query)
}

// GetMediaVersionsByQuality دریافت نسخه‌ها بر اساس کیفیت
func (b *queryBusImpl) GetMediaVersionsByQuality(ctx context.Context, mediaID uint, quality string) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	query := queries.GetMediaVersionsByQualityQuery{MediaID: mediaID, Quality: quality}
	return b.mediaHandler.HandleGetMediaVersionsByQuality(ctx, query)
}

// GetMediaVersionHistory دریافت تاریخچه نسخه‌ها
func (b *queryBusImpl) GetMediaVersionHistory(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	query := queries.GetMediaVersionHistoryQuery{MediaID: mediaID}
	return b.mediaHandler.HandleGetMediaVersionHistory(ctx, query)
}

// GetTotalSizeByMedia محاسبه حجم کل
func (b *queryBusImpl) GetTotalSizeByMedia(ctx context.Context, mediaID uint) (uint64, error) {
	if mediaID == 0 {
		return 0, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	query := queries.GetTotalSizeByMediaQuery{MediaID: mediaID}
	return b.mediaHandler.HandleGetTotalSizeByMedia(ctx, query)
}

// GetCompressionJobByID دریافت job بر اساس ID
func (b *queryBusImpl) GetCompressionJobByID(ctx context.Context, id uint) (*models.CompressionJob, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه job نامعتبر است")
	}
	query := queries.GetCompressionJobByIDQuery{ID: id}
	return b.mediaHandler.HandleGetCompressionJobByID(ctx, query)
}

// GetCompressionJobsByMedia دریافت job های یک رسانه
func (b *queryBusImpl) GetCompressionJobsByMedia(ctx context.Context, mediaID uint) ([]models.CompressionJob, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	query := queries.GetCompressionJobsByMediaQuery{MediaID: mediaID}
	return b.mediaHandler.HandleGetCompressionJobsByMedia(ctx, query)
}

// GetCompressionJobsByStatus دریافت job ها بر اساس وضعیت
func (b *queryBusImpl) GetCompressionJobsByStatus(ctx context.Context, status string) ([]models.CompressionJob, error) {
	if status == "" {
		return nil, fmt.Errorf("وضعیت نمی‌تواند خالی باشد")
	}
	query := queries.GetCompressionJobsByStatusQuery{Status: status}
	return b.mediaHandler.HandleGetCompressionJobsByStatus(ctx, query)
}

// GetPendingCompressionJobs دریافت job های در انتظار
func (b *queryBusImpl) GetPendingCompressionJobs(ctx context.Context, limit int) ([]models.CompressionJob, error) {
	if limit <= 0 {
		limit = 10
	}
	query := queries.GetPendingCompressionJobsQuery{Limit: limit}
	return b.mediaHandler.HandleGetPendingCompressionJobs(ctx, query)
}

// GetCompressionSettingByID دریافت تنظیمات بر اساس ID
func (b *queryBusImpl) GetCompressionSettingByID(ctx context.Context, id uint) (*models.CompressionSetting, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه تنظیمات نامعتبر است")
	}
	query := queries.GetCompressionSettingByIDQuery{ID: id}
	return b.mediaHandler.HandleGetCompressionSettingByID(ctx, query)
}

// GetCompressionSettingsByScope دریافت تنظیمات بر اساس scope
func (b *queryBusImpl) GetCompressionSettingsByScope(ctx context.Context, scope string) ([]models.CompressionSetting, error) {
	if scope == "" {
		return nil, fmt.Errorf("scope نمی‌تواند خالی باشد")
	}
	query := queries.GetCompressionSettingsByScopeQuery{Scope: scope}
	return b.mediaHandler.HandleGetCompressionSettingsByScope(ctx, query)
}

// GetCompressionSettingsByUser دریافت تنظیمات کاربر
func (b *queryBusImpl) GetCompressionSettingsByUser(ctx context.Context, userID uint) ([]models.CompressionSetting, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	query := queries.GetCompressionSettingsByUserQuery{UserID: userID}
	return b.mediaHandler.HandleGetCompressionSettingsByUser(ctx, query)
}

// GetCompressionSettingsByMedia دریافت تنظیمات رسانه
func (b *queryBusImpl) GetCompressionSettingsByMedia(ctx context.Context, mediaID uint) ([]models.CompressionSetting, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	query := queries.GetCompressionSettingsByMediaQuery{MediaID: mediaID}
	return b.mediaHandler.HandleGetCompressionSettingsByMedia(ctx, query)
}

// GetCompressionStatsByType دریافت آمار بر اساس نوع
func (b *queryBusImpl) GetCompressionStatsByType(ctx context.Context, statType string) ([]models.CompressionStat, error) {
	if statType == "" {
		return nil, fmt.Errorf("نوع آمار نمی‌تواند خالی باشد")
	}
	query := queries.GetCompressionStatsByTypeQuery{StatType: statType}
	return b.mediaHandler.HandleGetCompressionStatsByType(ctx, query)
}

// GetCompressionStatsByMedia دریافت آمار رسانه
func (b *queryBusImpl) GetCompressionStatsByMedia(ctx context.Context, mediaID uint) ([]models.CompressionStat, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	query := queries.GetCompressionStatsByMediaQuery{MediaID: mediaID}
	return b.mediaHandler.HandleGetCompressionStatsByMedia(ctx, query)
}

// GetCompressionStatsByUser دریافت آمار کاربر
func (b *queryBusImpl) GetCompressionStatsByUser(ctx context.Context, userID uint) ([]models.CompressionStat, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	query := queries.GetCompressionStatsByUserQuery{UserID: userID}
	return b.mediaHandler.HandleGetCompressionStatsByUser(ctx, query)
}

// GetMediaStorageReport دریافت گزارش فضای ذخیره‌سازی
func (b *queryBusImpl) GetMediaStorageReport(ctx context.Context, groupBy string, limit int) ([]handlers.MediaStorageReport, error) {
	if groupBy == "" {
		groupBy = "all"
	}
	if limit <= 0 {
		limit = 10
	}
	query := queries.GetMediaStorageReportQuery{GroupBy: groupBy, Limit: limit}
	return b.mediaHandler.HandleGetMediaStorageReport(ctx, query)
}

// ==================== Product Queries Implementation ====================

// GetProductByID دریافت محصول بر اساس ID
func (b *queryBusImpl) GetProductByID(ctx context.Context, id uint) (*models.Product, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه محصول نامعتبر است")
	}
	query := queries.GetProductByIDQuery{ID: id}
	return b.productHandler.HandleGetProductByID(ctx, query)
}

// GetProductBySlug دریافت محصول بر اساس slug
func (b *queryBusImpl) GetProductBySlug(ctx context.Context, slug string) (*models.Product, error) {
	if slug == "" {
		return nil, fmt.Errorf("slug محصول نمی‌تواند خالی باشد")
	}
	query := queries.GetProductBySlugQuery{Slug: slug}
	return b.productHandler.HandleGetProductBySlug(ctx, query)
}

// GetProductBySKU دریافت محصول بر اساس SKU
func (b *queryBusImpl) GetProductBySKU(ctx context.Context, sku string) (*models.Product, error) {
	if sku == "" {
		return nil, fmt.Errorf("SKU محصول نمی‌تواند خالی باشد")
	}
	query := queries.GetProductBySKUQuery{SKU: sku}
	return b.productHandler.HandleGetProductBySKU(ctx, query)
}

// GetProducts دریافت لیست محصولات با فیلتر
func (b *queryBusImpl) GetProducts(ctx context.Context, page, pageSize int, filter queries.ProductFilter) ([]models.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetProductsQuery{
		Page:     page,
		PageSize: pageSize,
		Filter:   filter,
	}
	return b.productHandler.HandleGetProducts(ctx, query)
}

// GetProductsByCategory دریافت محصولات بر اساس دسته‌بندی
func (b *queryBusImpl) GetProductsByCategory(ctx context.Context, categoryID uint, page, pageSize int) ([]models.Product, int64, error) {
	if categoryID == 0 {
		return nil, 0, fmt.Errorf("شناسه دسته‌بندی نامعتبر است")
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetProductsByCategoryQuery{
		CategoryID: categoryID,
		Page:       page,
		PageSize:   pageSize,
	}
	return b.productHandler.HandleGetProductsByCategory(ctx, query)
}

// GetProductsByBrand دریافت محصولات بر اساس برند
func (b *queryBusImpl) GetProductsByBrand(ctx context.Context, brandID uint, page, pageSize int) ([]models.Product, int64, error) {
	if brandID == 0 {
		return nil, 0, fmt.Errorf("شناسه برند نامعتبر است")
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetProductsByBrandQuery{
		BrandID:  brandID,
		Page:     page,
		PageSize: pageSize,
	}
	return b.productHandler.HandleGetProductsByBrand(ctx, query)
}

// GetPublishedProducts دریافت محصولات منتشر شده
func (b *queryBusImpl) GetPublishedProducts(ctx context.Context, page, pageSize int) ([]models.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetPublishedProductsQuery{
		Page:     page,
		PageSize: pageSize,
	}
	return b.productHandler.HandleGetPublishedProducts(ctx, query)
}

// GetNewProducts دریافت محصولات جدید
func (b *queryBusImpl) GetNewProducts(ctx context.Context, limit int) ([]models.Product, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	query := queries.GetNewProductsQuery{Limit: limit}
	return b.productHandler.HandleGetNewProducts(ctx, query)
}

// GetDiscountedProducts دریافت محصولات تخفیف‌دار
func (b *queryBusImpl) GetDiscountedProducts(ctx context.Context, page, pageSize int) ([]models.Product, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetDiscountedProductsQuery{
		Limit: pageSize,
	}
	return b.productHandler.HandleGetDiscountedProducts(ctx, query)
}

// GetFeaturedProducts دریافت محصولات ویژه
func (b *queryBusImpl) GetFeaturedProducts(ctx context.Context, limit int) ([]models.Product, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	query := queries.GetFeaturedProductsQuery{Limit: limit}
	return b.productHandler.HandleGetFeaturedProducts(ctx, query)
}

// SearchProducts جستجو در محصولات
func (b *queryBusImpl) SearchProducts(ctx context.Context, searchTerm string, page, pageSize int) ([]models.Product, int64, error) {
	if searchTerm == "" {
		return nil, 0, fmt.Errorf("عبارت جستجو نمی‌تواند خالی باشد")
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.SearchProductsQuery{
		Query:    searchTerm,
		Page:     page,
		PageSize: pageSize,
	}
	return b.productHandler.HandleSearchProducts(ctx, query)
}

// GetProductsByPriceRange دریافت محصولات بر اساس محدوده قیمت
func (b *queryBusImpl) GetProductsByPriceRange(ctx context.Context, minPrice, maxPrice float64, page, pageSize int) ([]models.Product, int64, error) {
	if minPrice < 0 || maxPrice < 0 || minPrice > maxPrice {
		return nil, 0, fmt.Errorf("محدوده قیمت نامعتبر است")
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetProductsByPriceRangeQuery{
		MinPrice: minPrice,
		MaxPrice: maxPrice,
		Page:     page,
		PageSize: pageSize,
	}
	return b.productHandler.HandleGetProductsByPriceRange(ctx, query)
}

// GetRelatedProducts دریافت محصولات مرتبط
func (b *queryBusImpl) GetRelatedProducts(ctx context.Context, productID uint, limit int) ([]models.Product, error) {
	if productID == 0 {
		return nil, fmt.Errorf("شناسه محصول نامعتبر است")
	}
	if limit <= 0 || limit > 50 {
		limit = 10
	}
	query := queries.GetRelatedProductsQuery{
		ProductID: productID,
		Limit:     limit,
	}
	return b.productHandler.HandleGetRelatedProducts(ctx, query)
}

// GetLowStockProducts دریافت محصولات با موجودی کم
func (b *queryBusImpl) GetLowStockProducts(ctx context.Context, threshold int) ([]models.Product, error) {
	if threshold < 0 {
		threshold = 10
	}
	query := queries.GetLowStockProductsQuery{Threshold: threshold}
	return b.productHandler.HandleGetLowStockProducts(ctx, query)
}

// GetProductStats دریافت آمار محصولات
func (b *queryBusImpl) GetProductStats(ctx context.Context) (map[string]interface{}, error) {
	query := queries.GetProductStatsQuery{}
	return b.productHandler.HandleGetProductStats(ctx, query)
}

// ==================== Category Queries Implementation ====================

// GetCategoryByID دریافت دسته‌بندی بر اساس ID
func (b *queryBusImpl) GetCategoryByID(ctx context.Context, id uint) (*models.Category, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه دسته‌بندی نامعتبر است")
	}
	query := queries.GetCategoryByIDQuery{ID: id}
	return b.productHandler.HandleGetCategoryByID(ctx, query)
}

// GetCategoryBySlug دریافت دسته‌بندی بر اساس slug
func (b *queryBusImpl) GetCategoryBySlug(ctx context.Context, slug string) (*models.Category, error) {
	if slug == "" {
		return nil, fmt.Errorf("slug دسته‌بندی نمی‌تواند خالی باشد")
	}
	query := queries.GetCategoryBySlugQuery{Slug: slug}
	return b.productHandler.HandleGetCategoryBySlug(ctx, query)
}

// GetCategories دریافت لیست دسته‌بندی‌ها
func (b *queryBusImpl) GetCategories(ctx context.Context, page, pageSize int, filter queries.CategoryFilter) ([]models.Category, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetCategoriesQuery{
		Page:     page,
		PageSize: pageSize,
		Filter:   filter,
	}
	return b.productHandler.HandleGetCategories(ctx, query)
}

// GetRootCategories دریافت دسته‌بندی‌های اصلی
func (b *queryBusImpl) GetRootCategories(ctx context.Context) ([]models.Category, error) {
	query := queries.GetRootCategoriesQuery{}
	return b.productHandler.HandleGetRootCategories(ctx, query)
}

// GetCategoryHierarchy دریافت سلسله مراتب دسته‌بندی‌ها
func (b *queryBusImpl) GetCategoryHierarchy(ctx context.Context) ([]models.Category, error) {
	query := queries.GetCategoryHierarchyQuery{}
	return b.productHandler.HandleGetCategoryHierarchy(ctx, query)
}

// GetCategoriesWithProductCount دریافت دسته‌بندی‌ها با تعداد محصولات
func (b *queryBusImpl) GetCategoriesWithProductCount(ctx context.Context) ([]models.Category, error) {
	query := queries.GetCategoriesWithProductCountQuery{}
	return b.productHandler.HandleGetCategoriesWithProductCount(ctx, query)
}

// GetFeaturedCategories دریافت دسته‌بندی‌های ویژه
func (b *queryBusImpl) GetFeaturedCategories(ctx context.Context, limit int) ([]models.Category, error) {
	if limit <= 0 || limit > 50 {
		limit = 10
	}
	return b.productHandler.HandleGetFeaturedCategories(ctx, queries.GetFeaturedCategoriesQuery{Limit: limit})
}

// ==================== Brand Queries Implementation ====================

// GetBrandByID دریافت برند بر اساس ID
func (b *queryBusImpl) GetBrandByID(ctx context.Context, id uint) (*models.Brand, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه برند نامعتبر است")
	}
	query := queries.GetBrandByIDQuery{ID: id}
	return b.productHandler.HandleGetBrandByID(ctx, query)
}

// GetBrandBySlug دریافت برند بر اساس slug
func (b *queryBusImpl) GetBrandBySlug(ctx context.Context, slug string) (*models.Brand, error) {
	if slug == "" {
		return nil, fmt.Errorf("slug برند نمی‌تواند خالی باشد")
	}
	query := queries.GetBrandBySlugQuery{Slug: slug}
	return b.productHandler.HandleGetBrandBySlug(ctx, query)
}

// GetBrands دریافت لیست برندها
func (b *queryBusImpl) GetBrands(ctx context.Context, page, pageSize int, filter queries.BrandFilter) ([]models.Brand, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetBrandsQuery{
		Page:     page,
		PageSize: pageSize,
		Filter:   filter,
	}
	return b.productHandler.HandleGetBrands(ctx, query)
}

// GetPublishedBrands دریافت برندهای منتشر شده
func (b *queryBusImpl) GetPublishedBrands(ctx context.Context) ([]models.Brand, error) {
	query := queries.GetPublishedBrandsQuery{}
	return b.productHandler.HandleGetPublishedBrands(ctx, query)
}

// GetFeaturedBrands دریافت برندهای ویژه
func (b *queryBusImpl) GetFeaturedBrands(ctx context.Context, limit int) ([]models.Brand, error) {
	if limit <= 0 || limit > 50 {
		limit = 10
	}
	query := queries.GetFeaturedBrandsQuery{Limit: limit}
	return b.productHandler.HandleGetFeaturedBrands(ctx, query)
}

// SearchBrands جستجو در برندها
func (b *queryBusImpl) SearchBrands(ctx context.Context, searchTerm string) ([]models.Brand, error) {
	if searchTerm == "" {
		return nil, fmt.Errorf("عبارت جستجو نمی‌تواند خالی باشد")
	}
	query := queries.SearchBrandsQuery{Query: searchTerm, Limit: 20}
	return b.productHandler.HandleSearchBrands(ctx, query)
}

// ==================== Review Queries Implementation ====================

// GetReviewByID دریافت نظر بر اساس ID
func (b *queryBusImpl) GetReviewByID(ctx context.Context, id uint) (*models.Review, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه نظر نامعتبر است")
	}
	query := queries.GetReviewByIDQuery{ID: id}
	return b.productHandler.HandleGetReviewByID(ctx, query)
}

// GetReviews دریافت لیست نظرات
func (b *queryBusImpl) GetReviews(ctx context.Context, page, pageSize int, filter queries.ReviewFilter) ([]models.Review, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetReviewsQuery{
		Page:     page,
		PageSize: pageSize,
		Filter:   filter,
	}
	return b.productHandler.HandleGetReviews(ctx, query)
}

// GetReviewsByProduct دریافت نظرات محصول
func (b *queryBusImpl) GetReviewsByProduct(ctx context.Context, productID uint, page, pageSize int) ([]models.Review, int64, error) {
	if productID == 0 {
		return nil, 0, fmt.Errorf("شناسه محصول نامعتبر است")
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetReviewsByProductQuery{
		ProductID: productID,
		Page:      page,
		PageSize:  pageSize,
	}
	return b.productHandler.HandleGetReviewsByProduct(ctx, query)
}

// GetReviewsByUser دریافت نظرات کاربر
func (b *queryBusImpl) GetReviewsByUser(ctx context.Context, userID uint, page, pageSize int) ([]models.Review, int64, error) {
	if userID == 0 {
		return nil, 0, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetReviewsByUserQuery{
		UserID:   userID,
		Page:     page,
		PageSize: pageSize,
	}
	return b.productHandler.HandleGetReviewsByUser(ctx, query)
}

// GetPublishedReviews دریافت نظرات منتشر شده
func (b *queryBusImpl) GetPublishedReviews(ctx context.Context, page, pageSize int) ([]models.Review, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetPublishedReviewsQuery{
		Page:     page,
		PageSize: pageSize,
	}
	return b.productHandler.HandleGetPublishedReviews(ctx, query)
}

// GetPendingReviews دریافت نظرات در انتظار تایید
func (b *queryBusImpl) GetPendingReviews(ctx context.Context, page, pageSize int) ([]models.Review, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetPendingReviewsQuery{
		Page:     page,
		PageSize: pageSize,
	}
	return b.productHandler.HandleGetPendingReviews(ctx, query)
}

// GetAverageRating دریافت میانگین امتیاز محصول
func (b *queryBusImpl) GetAverageRating(ctx context.Context, productID uint) (float64, error) {
	if productID == 0 {
		return 0, fmt.Errorf("شناسه محصول نامعتبر است")
	}
	query := queries.GetAverageRatingQuery{ProductID: productID}
	return b.productHandler.HandleGetAverageRating(ctx, query)
}

// GetReviewCount دریافت تعداد نظرات
func (b *queryBusImpl) GetReviewCount(ctx context.Context, productID uint) (int64, error) {
	if productID == 0 {
		return 0, fmt.Errorf("شناسه محصول نامعتبر است")
	}
	query := queries.GetReviewCountQuery{ProductID: productID}
	return b.productHandler.HandleGetReviewCount(ctx, query)
}

// ==================== Wishlist Queries Implementation ====================

// GetWishlistByUser دریافت لیست علاقه‌مندی کاربر
func (b *queryBusImpl) GetWishlistByUser(ctx context.Context, userID uint) ([]models.Wishlist, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	query := queries.GetWishlistByUserQuery{UserID: userID}
	return b.productHandler.HandleGetWishlistByUser(ctx, query)
}

// GetWishlistProducts دریافت محصولات لیست علاقه‌مندی
func (b *queryBusImpl) GetWishlistProducts(ctx context.Context, userID uint) ([]models.Product, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	query := queries.GetWishlistProductsQuery{UserID: userID}
	return b.productHandler.HandleGetWishlistProducts(ctx, query)
}

// IsInWishlist بررسی وجود محصول در لیست علاقه‌مندی
func (b *queryBusImpl) IsInWishlist(ctx context.Context, userID, productID uint) (bool, error) {
	if userID == 0 {
		return false, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	if productID == 0 {
		return false, fmt.Errorf("شناسه محصول نامعتبر است")
	}
	query := queries.IsInWishlistQuery{
		UserID:    userID,
		ProductID: productID,
	}
	return b.productHandler.HandleIsInWishlist(ctx, query)
}

// GetWishlistCount دریافت تعداد محصولات در لیست علاقه‌مندی
func (b *queryBusImpl) GetWishlistCount(ctx context.Context, userID uint) (int64, error) {
	if userID == 0 {
		return 0, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	query := queries.GetWishlistCountQuery{UserID: userID}
	return b.productHandler.HandleGetWishlistCount(ctx, query)
}

// ==================== Stock Notification Queries Implementation ====================

// GetStockNotifications دریافت لیست اعلان‌های موجودی
func (b *queryBusImpl) GetStockNotifications(ctx context.Context, page, pageSize int, filter queries.StockNotificationFilter) ([]models.StockNotification, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetStockNotificationsQuery{
		Page:     page,
		PageSize: pageSize,
		Filter:   filter,
	}
	return b.productHandler.HandleGetStockNotifications(ctx, query)
}

// GetStockNotificationsByProduct دریافت اعلان‌های موجودی محصول
func (b *queryBusImpl) GetStockNotificationsByProduct(ctx context.Context, productID uint) ([]models.StockNotification, error) {
	if productID == 0 {
		return nil, fmt.Errorf("شناسه محصول نامعتبر است")
	}
	query := queries.GetStockNotificationsByProductQuery{ProductID: productID}
	return b.productHandler.HandleGetStockNotificationsByProduct(ctx, query)
}

// GetStockNotificationsByUser دریافت اعلان‌های موجودی کاربر
func (b *queryBusImpl) GetStockNotificationsByUser(ctx context.Context, userID uint) ([]models.StockNotification, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	query := queries.GetStockNotificationsByUserQuery{UserID: userID}
	return b.productHandler.HandleGetStockNotificationsByUser(ctx, query)
}

// GetPendingStockNotifications دریافت اعلان‌های موجودی در انتظار
func (b *queryBusImpl) GetPendingStockNotifications(ctx context.Context) ([]models.StockNotification, error) {
	query := queries.GetPendingStockNotificationsQuery{}
	return b.productHandler.HandleGetPendingStockNotifications(ctx, query)
}

// ==================== Product QA Queries Implementation ====================

// GetProductQA دریافت پرسش و پاسخ محصول
func (b *queryBusImpl) GetProductQA(ctx context.Context, page, pageSize int, filter queries.ProductQAFilter) ([]models.ProductQA, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetProductQAQuery{
		Page:     page,
		PageSize: pageSize,
		Filter:   filter,
	}
	return b.productHandler.HandleGetProductQA(ctx, query)
}

// GetQuestionsByProduct دریافت سوالات محصول
func (b *queryBusImpl) GetQuestionsByProduct(ctx context.Context, productID uint, page, pageSize int) ([]models.ProductQA, int64, error) {
	if productID == 0 {
		return nil, 0, fmt.Errorf("شناسه محصول نامعتبر است")
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetQuestionsByProductQuery{
		ProductID: productID,
		Page:      page,
		PageSize:  pageSize,
	}
	return b.productHandler.HandleGetQuestionsByProduct(ctx, query)
}

// GetAnswersByQuestion دریافت پاسخ‌های سوال
func (b *queryBusImpl) GetAnswersByQuestion(ctx context.Context, questionID uint) ([]models.ProductQA, error) {
	if questionID == 0 {
		return nil, fmt.Errorf("شناسه سوال نامعتبر است")
	}
	query := queries.GetAnswersByQuestionQuery{QuestionID: questionID}
	return b.productHandler.HandleGetAnswersByQuestion(ctx, query)
}

// GetPublishedQA دریافت پرسش و پاسخ‌های منتشر شده
func (b *queryBusImpl) GetPublishedQA(ctx context.Context, page, pageSize int) ([]models.ProductQA, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetPublishedQAQuery{
		Page:     page,
		PageSize: pageSize,
	}
	return b.productHandler.HandleGetPublishedQA(ctx, query)
}

// GetPendingQA دریافت پرسش و پاسخ‌های در انتظار تایید
func (b *queryBusImpl) GetPendingQA(ctx context.Context, page, pageSize int) ([]models.ProductQA, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	query := queries.GetPendingQAQuery{
		Page:     page,
		PageSize: pageSize,
	}
	return b.productHandler.HandleGetPendingQA(ctx, query)
}

// GetQuestionCount دریافت تعداد سوالات
func (b *queryBusImpl) GetQuestionCount(ctx context.Context, productID uint) (int64, error) {
	if productID == 0 {
		return 0, fmt.Errorf("شناسه محصول نامعتبر است")
	}
	query := queries.GetQuestionCountQuery{ProductID: productID}
	return b.productHandler.HandleGetQuestionCount(ctx, query)
}

// GetAnswerCount دریافت تعداد پاسخ‌ها
func (b *queryBusImpl) GetAnswerCount(ctx context.Context, questionID uint) (int64, error) {
	if questionID == 0 {
		return 0, fmt.Errorf("شناسه سوال نامعتبر است")
	}
	query := queries.GetAnswerCountQuery{QuestionID: questionID}
	return b.productHandler.HandleGetAnswerCount(ctx, query)
}
