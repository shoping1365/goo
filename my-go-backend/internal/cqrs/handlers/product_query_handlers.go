package handlers

import (
	"context"
	"my-go-backend/internal/cqrs/queries"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
)

// ProductQueryHandler پردازنده کوئری‌های مربوط به محصولات و موجودیت‌های مرتبط
// این handler مسئول اجرای تمام کوئری‌های read (خواندن و جستجو) است
type ProductQueryHandler struct {
	uowFactory unitofwork.UnitOfWorkFactory
}

// NewProductQueryHandler ایجاد نمونه جدید از ProductQueryHandler
func NewProductQueryHandler(uowFactory unitofwork.UnitOfWorkFactory) *ProductQueryHandler {
	return &ProductQueryHandler{
		uowFactory: uowFactory,
	}
}

// ==================== Product Queries ====================

// HandleGetProductByID پردازش کوئری دریافت محصول بر اساس ID
func (h *ProductQueryHandler) HandleGetProductByID(ctx context.Context, query queries.GetProductByIDQuery) (*models.Product, error) {
	uow := h.uowFactory.Create()
	return uow.ProductRepository().GetByID(ctx, query.ID)
}

// HandleGetProductBySlug پردازش کوئری دریافت محصول بر اساس slug
func (h *ProductQueryHandler) HandleGetProductBySlug(ctx context.Context, query queries.GetProductBySlugQuery) (*models.Product, error) {
	uow := h.uowFactory.Create()
	return uow.ProductRepository().GetBySlug(ctx, query.Slug)
}

// HandleGetProductBySKU پردازش کوئری دریافت محصول بر اساس SKU
func (h *ProductQueryHandler) HandleGetProductBySKU(ctx context.Context, query queries.GetProductBySKUQuery) (*models.Product, error) {
	uow := h.uowFactory.Create()
	return uow.ProductRepository().GetBySKU(ctx, query.SKU)
}

// HandleGetProducts پردازش کوئری دریافت لیست محصولات با فیلتر
func (h *ProductQueryHandler) HandleGetProducts(ctx context.Context, query queries.GetProductsQuery) ([]models.Product, int64, error) {
	// TODO: پیاده‌سازی GetProducts در ProductRepository
	return nil, 0, nil
}

// HandleGetProductsByCategory پردازش کوئری دریافت محصولات بر اساس دسته‌بندی
func (h *ProductQueryHandler) HandleGetProductsByCategory(ctx context.Context, query queries.GetProductsByCategoryQuery) ([]models.Product, int64, error) {
	uow := h.uowFactory.Create()
	// Keyset: اگر cursor ارائه شده بود، از مسیر سریع استفاده کن
	if !query.CursorCreatedAt.IsZero() && query.CursorID > 0 {
		items, err := uow.ProductRepository().GetByCategoryAfter(ctx, query.CategoryID, query.PageSize, query.CursorCreatedAt, query.CursorID)
		if err != nil {
			return nil, 0, err
		}
		return items, 0, nil
	}
	products, err := uow.ProductRepository().GetByCategory(ctx, query.CategoryID, query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return products, 0, nil
}

// HandleGetProductsByBrand پردازش کوئری دریافت محصولات بر اساس برند
func (h *ProductQueryHandler) HandleGetProductsByBrand(ctx context.Context, query queries.GetProductsByBrandQuery) ([]models.Product, int64, error) {
	uow := h.uowFactory.Create()
	if !query.CursorCreatedAt.IsZero() && query.CursorID > 0 {
		items, err := uow.ProductRepository().GetByBrandAfter(ctx, query.BrandID, query.PageSize, query.CursorCreatedAt, query.CursorID)
		if err != nil {
			return nil, 0, err
		}
		return items, 0, nil
	}
	products, err := uow.ProductRepository().GetByBrand(ctx, query.BrandID, query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return products, 0, nil
}

// HandleGetPublishedProducts پردازش کوئری دریافت محصولات منتشر شده
func (h *ProductQueryHandler) HandleGetPublishedProducts(ctx context.Context, query queries.GetPublishedProductsQuery) ([]models.Product, int64, error) {
	uow := h.uowFactory.Create()
	if !query.CursorCreatedAt.IsZero() && query.CursorID > 0 {
		items, err := uow.ProductRepository().GetPublishedAfter(ctx, query.PageSize, query.CursorCreatedAt, query.CursorID)
		if err != nil {
			return nil, 0, err
		}
		return items, 0, nil
	}
	products, err := uow.ProductRepository().GetPublished(ctx, query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return products, 0, nil
}

// HandleGetNewProducts پردازش کوئری دریافت محصولات جدید
func (h *ProductQueryHandler) HandleGetNewProducts(ctx context.Context, query queries.GetNewProductsQuery) ([]models.Product, error) {
	uow := h.uowFactory.Create()
	return uow.ProductRepository().GetNewProducts(ctx, query.Limit)
}

// HandleGetDiscountedProducts پردازش کوئری دریافت محصولات تخفیف‌دار
func (h *ProductQueryHandler) HandleGetDiscountedProducts(ctx context.Context, query queries.GetDiscountedProductsQuery) ([]models.Product, int64, error) {
	uow := h.uowFactory.Create()
	products, err := uow.ProductRepository().GetDiscountedProducts(ctx, query.Limit)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return products, 0, nil
}

// HandleGetFeaturedProducts پردازش کوئری دریافت محصولات ویژه
func (h *ProductQueryHandler) HandleGetFeaturedProducts(ctx context.Context, query queries.GetFeaturedProductsQuery) ([]models.Product, error) {
	// TODO: پیاده‌سازی GetFeatured در ProductRepository
	return nil, nil
}

// HandleSearchProducts پردازش کوئری جستجو در محصولات
func (h *ProductQueryHandler) HandleSearchProducts(ctx context.Context, query queries.SearchProductsQuery) ([]models.Product, int64, error) {
	// TODO: پیاده‌سازی Search در ProductRepository
	return nil, 0, nil
}

// HandleGetProductsByPriceRange پردازش کوئری دریافت محصولات بر اساس محدوده قیمت
func (h *ProductQueryHandler) HandleGetProductsByPriceRange(ctx context.Context, query queries.GetProductsByPriceRangeQuery) ([]models.Product, int64, error) {
	// TODO: پیاده‌سازی GetByPriceRange در ProductRepository
	return nil, 0, nil
}

// HandleGetRelatedProducts پردازش کوئری دریافت محصولات مرتبط
func (h *ProductQueryHandler) HandleGetRelatedProducts(ctx context.Context, query queries.GetRelatedProductsQuery) ([]models.Product, error) {
	uow := h.uowFactory.Create()
	return uow.ProductRepository().GetRelatedProducts(ctx, query.ProductID, query.Limit)
}

// HandleGetLowStockProducts پردازش کوئری دریافت محصولات با موجودی کم
func (h *ProductQueryHandler) HandleGetLowStockProducts(ctx context.Context, query queries.GetLowStockProductsQuery) ([]models.Product, error) {
	uow := h.uowFactory.Create()
	return uow.ProductRepository().GetLowStockProducts(ctx)
}

// HandleGetProductStats پردازش کوئری دریافت آمار محصولات
func (h *ProductQueryHandler) HandleGetProductStats(ctx context.Context, query queries.GetProductStatsQuery) (map[string]interface{}, error) {
	// TODO: پیاده‌سازی GetStats در ProductRepository
	return nil, nil
}

// ==================== Category Queries ====================

// HandleGetCategoryByID پردازش کوئری دریافت دسته‌بندی بر اساس ID
func (h *ProductQueryHandler) HandleGetCategoryByID(ctx context.Context, query queries.GetCategoryByIDQuery) (*models.Category, error) {
	uow := h.uowFactory.Create()
	return uow.CategoryRepository().GetByID(ctx, query.ID)
}

// HandleGetCategoryBySlug پردازش کوئری دریافت دسته‌بندی بر اساس slug
func (h *ProductQueryHandler) HandleGetCategoryBySlug(ctx context.Context, query queries.GetCategoryBySlugQuery) (*models.Category, error) {
	uow := h.uowFactory.Create()
	return uow.CategoryRepository().GetBySlug(ctx, query.Slug)
}

// HandleGetCategories پردازش کوئری دریافت لیست دسته‌بندی‌ها
func (h *ProductQueryHandler) HandleGetCategories(ctx context.Context, query queries.GetCategoriesQuery) ([]models.Category, int64, error) {
	// TODO: پیاده‌سازی GetCategories در CategoryRepository
	return nil, 0, nil
}

// HandleGetRootCategories پردازش کوئری دریافت دسته‌بندی‌های اصلی
func (h *ProductQueryHandler) HandleGetRootCategories(ctx context.Context, query queries.GetRootCategoriesQuery) ([]models.Category, error) {
	uow := h.uowFactory.Create()
	return uow.CategoryRepository().GetRootCategories(ctx)
}

// HandleGetCategoryHierarchy پردازش کوئری دریافت سلسله مراتب دسته‌بندی‌ها
func (h *ProductQueryHandler) HandleGetCategoryHierarchy(ctx context.Context, query queries.GetCategoryHierarchyQuery) ([]models.Category, error) {
	uow := h.uowFactory.Create()
	return uow.CategoryRepository().GetHierarchy(ctx)
}

// HandleGetCategoriesWithProductCount پردازش کوئری دریافت دسته‌بندی‌ها با تعداد محصولات
func (h *ProductQueryHandler) HandleGetCategoriesWithProductCount(ctx context.Context, query queries.GetCategoriesWithProductCountQuery) ([]models.Category, error) {
	uow := h.uowFactory.Create()
	return uow.CategoryRepository().GetWithProductCount(ctx)
}

// HandleGetFeaturedCategories پردازش کوئری دریافت دسته‌بندی‌های ویژه
func (h *ProductQueryHandler) HandleGetFeaturedCategories(ctx context.Context, query queries.GetFeaturedCategoriesQuery) ([]models.Category, error) {
	// TODO: پیاده‌سازی GetFeatured در CategoryRepository
	return nil, nil
}

// ==================== Brand Queries ====================

// HandleGetBrandByID پردازش کوئری دریافت برند بر اساس ID
func (h *ProductQueryHandler) HandleGetBrandByID(ctx context.Context, query queries.GetBrandByIDQuery) (*models.Brand, error) {
	uow := h.uowFactory.Create()
	return uow.BrandRepository().GetByID(ctx, query.ID)
}

// HandleGetBrandBySlug پردازش کوئری دریافت برند بر اساس slug
func (h *ProductQueryHandler) HandleGetBrandBySlug(ctx context.Context, query queries.GetBrandBySlugQuery) (*models.Brand, error) {
	uow := h.uowFactory.Create()
	return uow.BrandRepository().GetBySlug(ctx, query.Slug)
}

// HandleGetBrands پردازش کوئری دریافت لیست برندها
func (h *ProductQueryHandler) HandleGetBrands(ctx context.Context, query queries.GetBrandsQuery) ([]models.Brand, int64, error) {
	// TODO: پیاده‌سازی GetBrands در BrandRepository
	return nil, 0, nil
}

// HandleGetPublishedBrands پردازش کوئری دریافت برندهای منتشر شده
func (h *ProductQueryHandler) HandleGetPublishedBrands(ctx context.Context, query queries.GetPublishedBrandsQuery) ([]models.Brand, error) {
	uow := h.uowFactory.Create()
	return uow.BrandRepository().GetPublished(ctx)
}

// HandleGetFeaturedBrands پردازش کوئری دریافت برندهای ویژه
func (h *ProductQueryHandler) HandleGetFeaturedBrands(ctx context.Context, query queries.GetFeaturedBrandsQuery) ([]models.Brand, error) {
	// TODO: پیاده‌سازی GetFeatured در BrandRepository
	return nil, nil
}

// HandleSearchBrands پردازش کوئری جستجو در برندها
func (h *ProductQueryHandler) HandleSearchBrands(ctx context.Context, query queries.SearchBrandsQuery) ([]models.Brand, error) {
	// TODO: پیاده‌سازی Search در BrandRepository
	return nil, nil
}

// ==================== Review Queries ====================

// HandleGetReviewByID پردازش کوئری دریافت نظر بر اساس ID
func (h *ProductQueryHandler) HandleGetReviewByID(ctx context.Context, query queries.GetReviewByIDQuery) (*models.Review, error) {
	uow := h.uowFactory.Create()
	return uow.ReviewRepository().GetByID(ctx, query.ID)
}

// HandleGetReviews پردازش کوئری دریافت لیست نظرات
func (h *ProductQueryHandler) HandleGetReviews(ctx context.Context, query queries.GetReviewsQuery) ([]models.Review, int64, error) {
	// TODO: پیاده‌سازی GetReviews در ReviewRepository
	return nil, 0, nil
}

// HandleGetReviewsByProduct پردازش کوئری دریافت نظرات محصول
func (h *ProductQueryHandler) HandleGetReviewsByProduct(ctx context.Context, query queries.GetReviewsByProductQuery) ([]models.Review, int64, error) {
	uow := h.uowFactory.Create()
	reviews, err := uow.ReviewRepository().GetByProduct(ctx, query.ProductID, query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return reviews, 0, nil
}

// HandleGetReviewsByUser پردازش کوئری دریافت نظرات کاربر
func (h *ProductQueryHandler) HandleGetReviewsByUser(ctx context.Context, query queries.GetReviewsByUserQuery) ([]models.Review, int64, error) {
	uow := h.uowFactory.Create()
	reviews, err := uow.ReviewRepository().GetByCustomer(ctx, query.UserID, query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return reviews, 0, nil
}

// HandleGetPublishedReviews پردازش کوئری دریافت نظرات منتشر شده
func (h *ProductQueryHandler) HandleGetPublishedReviews(ctx context.Context, query queries.GetPublishedReviewsQuery) ([]models.Review, int64, error) {
	uow := h.uowFactory.Create()
	reviews, err := uow.ReviewRepository().GetPublished(ctx, query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return reviews, 0, nil
}

// HandleGetPendingReviews پردازش کوئری دریافت نظرات در انتظار تایید
func (h *ProductQueryHandler) HandleGetPendingReviews(ctx context.Context, query queries.GetPendingReviewsQuery) ([]models.Review, int64, error) {
	uow := h.uowFactory.Create()
	reviews, err := uow.ReviewRepository().GetPending(ctx, query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return reviews, 0, nil
}

// HandleGetAverageRating پردازش کوئری دریافت میانگین امتیاز محصول
func (h *ProductQueryHandler) HandleGetAverageRating(ctx context.Context, query queries.GetAverageRatingQuery) (float64, error) {
	uow := h.uowFactory.Create()
	return uow.ReviewRepository().GetAverageRating(ctx, query.ProductID)
}

// HandleGetReviewCount پردازش کوئری دریافت تعداد نظرات
func (h *ProductQueryHandler) HandleGetReviewCount(ctx context.Context, query queries.GetReviewCountQuery) (int64, error) {
	uow := h.uowFactory.Create()
	return uow.ReviewRepository().GetReviewCount(ctx, query.ProductID)
}

// ==================== Wishlist Queries ====================

// HandleGetWishlistByUser پردازش کوئری دریافت لیست علاقه‌مندی کاربر
func (h *ProductQueryHandler) HandleGetWishlistByUser(ctx context.Context, query queries.GetWishlistByUserQuery) ([]models.Wishlist, error) {
	uow := h.uowFactory.Create()
	_, err := uow.WishlistRepository().GetByUser(ctx, query.UserID)
	if err != nil {
		return nil, err
	}
	// TODO: تبدیل به آرایه
	return []models.Wishlist{}, nil
}

// HandleGetWishlistProducts پردازش کوئری دریافت محصولات لیست علاقه‌مندی
func (h *ProductQueryHandler) HandleGetWishlistProducts(ctx context.Context, query queries.GetWishlistProductsQuery) ([]models.Product, error) {
	uow := h.uowFactory.Create()
	return uow.WishlistRepository().GetWishlistProducts(ctx, query.UserID)
}

// HandleIsInWishlist پردازش کوئری بررسی وجود محصول در لیست علاقه‌مندی
func (h *ProductQueryHandler) HandleIsInWishlist(ctx context.Context, query queries.IsInWishlistQuery) (bool, error) {
	uow := h.uowFactory.Create()
	return uow.WishlistRepository().IsInWishlist(ctx, query.UserID, query.ProductID)
}

// HandleGetWishlistCount پردازش کوئری دریافت تعداد محصولات در لیست علاقه‌مندی
func (h *ProductQueryHandler) HandleGetWishlistCount(ctx context.Context, query queries.GetWishlistCountQuery) (int64, error) {
	uow := h.uowFactory.Create()
	return uow.WishlistRepository().GetWishlistCount(ctx, query.UserID)
}

// ==================== Stock Notification Queries ====================

// HandleGetStockNotifications پردازش کوئری دریافت لیست اعلان‌های موجودی
func (h *ProductQueryHandler) HandleGetStockNotifications(ctx context.Context, query queries.GetStockNotificationsQuery) ([]models.StockNotification, int64, error) {
	// TODO: پیاده‌سازی GetNotifications در StockNotificationRepository
	return nil, 0, nil
}

// HandleGetStockNotificationsByProduct پردازش کوئری دریافت اعلان‌های موجودی محصول
func (h *ProductQueryHandler) HandleGetStockNotificationsByProduct(ctx context.Context, query queries.GetStockNotificationsByProductQuery) ([]models.StockNotification, error) {
	uow := h.uowFactory.Create()
	return uow.StockNotificationRepository().GetByProduct(ctx, query.ProductID)
}

// HandleGetStockNotificationsByUser پردازش کوئری دریافت اعلان‌های موجودی کاربر
func (h *ProductQueryHandler) HandleGetStockNotificationsByUser(ctx context.Context, query queries.GetStockNotificationsByUserQuery) ([]models.StockNotification, error) {
	uow := h.uowFactory.Create()
	return uow.StockNotificationRepository().GetByUser(ctx, query.UserID)
}

// HandleGetPendingStockNotifications پردازش کوئری دریافت اعلان‌های موجودی در انتظار
func (h *ProductQueryHandler) HandleGetPendingStockNotifications(ctx context.Context, query queries.GetPendingStockNotificationsQuery) ([]models.StockNotification, error) {
	uow := h.uowFactory.Create()
	return uow.StockNotificationRepository().GetPending(ctx)
}

// ==================== Product QA Queries ====================

// HandleGetProductQA پردازش کوئری دریافت پرسش و پاسخ محصول
func (h *ProductQueryHandler) HandleGetProductQA(ctx context.Context, query queries.GetProductQAQuery) ([]models.ProductQA, int64, error) {
	// TODO: پیاده‌سازی GetQA در ProductQARepository
	return nil, 0, nil
}

// HandleGetQuestionsByProduct پردازش کوئری دریافت سوالات محصول
func (h *ProductQueryHandler) HandleGetQuestionsByProduct(ctx context.Context, query queries.GetQuestionsByProductQuery) ([]models.ProductQA, int64, error) {
	uow := h.uowFactory.Create()
	qa, err := uow.ProductQARepository().GetQuestions(ctx, query.ProductID, query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return qa, 0, nil
}

// HandleGetAnswersByQuestion پردازش کوئری دریافت پاسخ‌های سوال
func (h *ProductQueryHandler) HandleGetAnswersByQuestion(ctx context.Context, query queries.GetAnswersByQuestionQuery) ([]models.ProductQA, error) {
	uow := h.uowFactory.Create()
	return uow.ProductQARepository().GetAnswers(ctx, query.QuestionID, 0, 0)
}

// HandleGetPublishedQA پردازش کوئری دریافت پرسش و پاسخ‌های منتشر شده
func (h *ProductQueryHandler) HandleGetPublishedQA(ctx context.Context, query queries.GetPublishedQAQuery) ([]models.ProductQA, int64, error) {
	uow := h.uowFactory.Create()
	qa, err := uow.ProductQARepository().GetPublished(ctx, query.ProductID, query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return qa, 0, nil
}

// HandleGetPendingQA پردازش کوئری دریافت پرسش و پاسخ‌های در انتظار تایید
func (h *ProductQueryHandler) HandleGetPendingQA(ctx context.Context, query queries.GetPendingQAQuery) ([]models.ProductQA, int64, error) {
	uow := h.uowFactory.Create()
	qa, err := uow.ProductQARepository().GetPending(ctx, query.PageSize, (query.Page-1)*query.PageSize)
	if err != nil {
		return nil, 0, err
	}
	// TODO: محاسبه تعداد کل
	return qa, 0, nil
}

// HandleGetQuestionCount پردازش کوئری دریافت تعداد سوالات
func (h *ProductQueryHandler) HandleGetQuestionCount(ctx context.Context, query queries.GetQuestionCountQuery) (int64, error) {
	uow := h.uowFactory.Create()
	return uow.ProductQARepository().GetQuestionCount(ctx, query.ProductID)
}

// HandleGetAnswerCount پردازش کوئری دریافت تعداد پاسخ‌ها
func (h *ProductQueryHandler) HandleGetAnswerCount(ctx context.Context, query queries.GetAnswerCountQuery) (int64, error) {
	uow := h.uowFactory.Create()
	return uow.ProductQARepository().GetAnswerCount(ctx, query.QuestionID)
}
