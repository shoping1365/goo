package queries

import (
	"time"
)

// ==================== Product Queries ====================

// ProductFilter فیلتر برای جستجو در محصولات
type ProductFilter struct {
	Name          string     `json:"name"`
	SKU           string     `json:"sku"`
	CategoryID    *uint      `json:"category_id"`
	BrandID       *uint      `json:"brand_id"`
	Status        string     `json:"status"`
	IsActive      *bool      `json:"is_active"`
	IsFeatured    *bool      `json:"is_featured"`
	IsNew         *bool      `json:"is_new"`
	IsOnSale      *bool      `json:"is_on_sale"`
	MinPrice      *float64   `json:"min_price"`
	MaxPrice      *float64   `json:"max_price"`
	InStock       *bool      `json:"in_stock"`
	LowStock      *bool      `json:"low_stock"`
	Tags          []string   `json:"tags"`
	CreatedAfter  *time.Time `json:"created_after"`
	CreatedBefore *time.Time `json:"created_before"`
	UpdatedAfter  *time.Time `json:"updated_after"`
	UpdatedBefore *time.Time `json:"updated_before"`
}

// GetProductByIDQuery کوئری دریافت محصول بر اساس ID
type GetProductByIDQuery struct {
	ID uint `json:"id" binding:"required"`
}

// GetProductBySlugQuery کوئری دریافت محصول بر اساس slug
type GetProductBySlugQuery struct {
	Slug string `json:"slug" binding:"required"`
}

// GetProductBySKUQuery کوئری دریافت محصول بر اساس SKU
type GetProductBySKUQuery struct {
	SKU string `json:"sku" binding:"required"`
}

// GetProductsQuery کوئری دریافت لیست محصولات
type GetProductsQuery struct {
	Page      int           `json:"page" binding:"required,min=1"`
	PageSize  int           `json:"page_size" binding:"required,min=1,max=100"`
	Filter    ProductFilter `json:"filter"`
	SortBy    string        `json:"sort_by"`
	SortOrder string        `json:"sort_order" binding:"oneof=asc desc"`
}

// GetProductsByCategoryQuery کوئری دریافت محصولات یک دسته‌بندی
type GetProductsByCategoryQuery struct {
	CategoryID      uint      `json:"category_id" binding:"required"`
	Page            int       `json:"page" binding:"required,min=1"`
	PageSize        int       `json:"page_size" binding:"required,min=1,max=100"`
	SortBy          string    `json:"sort_by"`
	SortOrder       string    `json:"sort_order" binding:"oneof=asc desc"`
	CursorID        uint      `json:"cursor_id"`         // برای keyset
	CursorCreatedAt time.Time `json:"cursor_created_at"` // برای keyset
}

// GetProductsByBrandQuery کوئری دریافت محصولات یک برند
type GetProductsByBrandQuery struct {
	BrandID         uint      `json:"brand_id" binding:"required"`
	Page            int       `json:"page" binding:"required,min=1"`
	PageSize        int       `json:"page_size" binding:"required,min=1,max=100"`
	SortBy          string    `json:"sort_by"`
	SortOrder       string    `json:"sort_order" binding:"oneof=asc desc"`
	CursorID        uint      `json:"cursor_id"`
	CursorCreatedAt time.Time `json:"cursor_created_at"`
}

// GetPublishedProductsQuery کوئری دریافت محصولات منتشر شده
type GetPublishedProductsQuery struct {
	Page            int       `json:"page" binding:"required,min=1"`
	PageSize        int       `json:"page_size" binding:"required,min=1,max=100"`
	SortBy          string    `json:"sort_by"`
	SortOrder       string    `json:"sort_order" binding:"oneof=asc desc"`
	CursorID        uint      `json:"cursor_id"`
	CursorCreatedAt time.Time `json:"cursor_created_at"`
}

// GetNewProductsQuery کوئری دریافت محصولات جدید
type GetNewProductsQuery struct {
	Limit int `json:"limit" binding:"required,min=1,max=50"`
}

// GetDiscountedProductsQuery کوئری دریافت محصولات تخفیف‌دار
type GetDiscountedProductsQuery struct {
	Limit int `json:"limit" binding:"required,min=1,max=50"`
}

// GetFeaturedProductsQuery کوئری دریافت محصولات ویژه
type GetFeaturedProductsQuery struct {
	Limit int `json:"limit" binding:"required,min=1,max=50"`
}

// SearchProductsQuery کوئری جستجو در محصولات
type SearchProductsQuery struct {
	Query     string `json:"query" binding:"required"`
	Page      int    `json:"page" binding:"required,min=1"`
	PageSize  int    `json:"page_size" binding:"required,min=1,max=100"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" binding:"oneof=asc desc"`
}

// GetProductsByPriceRangeQuery کوئری دریافت محصولات در محدوده قیمت
type GetProductsByPriceRangeQuery struct {
	MinPrice  float64 `json:"min_price" binding:"required,min=0"`
	MaxPrice  float64 `json:"max_price" binding:"required,min=0"`
	Page      int     `json:"page" binding:"required,min=1"`
	PageSize  int     `json:"page_size" binding:"required,min=1,max=100"`
	SortBy    string  `json:"sort_by"`
	SortOrder string  `json:"sort_order" binding:"oneof=asc desc"`
}

// GetRelatedProductsQuery کوئری دریافت محصولات مرتبط
type GetRelatedProductsQuery struct {
	ProductID uint `json:"product_id" binding:"required"`
	Limit     int  `json:"limit" binding:"required,min=1,max=20"`
}

// GetLowStockProductsQuery کوئری دریافت محصولات با موجودی کم
type GetLowStockProductsQuery struct {
	Threshold int `json:"threshold" binding:"required,min=1"`
	Page      int `json:"page" binding:"required,min=1"`
	PageSize  int `json:"page_size" binding:"required,min=1,max=100"`
}

// GetProductStatsQuery کوئری دریافت آمار محصولات
type GetProductStatsQuery struct {
	CategoryID *uint      `json:"category_id"`
	BrandID    *uint      `json:"brand_id"`
	FromDate   *time.Time `json:"from_date"`
	ToDate     *time.Time `json:"to_date"`
}

// ==================== Category Queries ====================

// CategoryFilter فیلتر برای جستجو در دسته‌بندی‌ها
type CategoryFilter struct {
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	ParentID       *uint  `json:"parent_id"`
	IsActive       *bool  `json:"is_active"`
	IsFeatured     *bool  `json:"is_featured"`
	ShowInHomePage *bool  `json:"show_in_home_page"`
	ShowInMenu     *bool  `json:"show_in_menu"`
}

// GetCategoryByIDQuery کوئری دریافت دسته‌بندی بر اساس ID
type GetCategoryByIDQuery struct {
	ID uint `json:"id" binding:"required"`
}

// GetCategoryBySlugQuery کوئری دریافت دسته‌بندی بر اساس slug
type GetCategoryBySlugQuery struct {
	Slug string `json:"slug" binding:"required"`
}

// GetCategoriesQuery کوئری دریافت لیست دسته‌بندی‌ها
type GetCategoriesQuery struct {
	Page      int            `json:"page" binding:"required,min=1"`
	PageSize  int            `json:"page_size" binding:"required,min=1,max=100"`
	Filter    CategoryFilter `json:"filter"`
	SortBy    string         `json:"sort_by"`
	SortOrder string         `json:"sort_order" binding:"oneof=asc desc"`
}

// GetRootCategoriesQuery کوئری دریافت دسته‌های اصلی
type GetRootCategoriesQuery struct {
	IsActive *bool `json:"is_active"`
}

// GetCategoryHierarchyQuery کوئری دریافت ساختار سلسله مراتبی دسته‌بندی‌ها
type GetCategoryHierarchyQuery struct {
	CategoryID *uint `json:"category_id"`
}

// GetCategoriesWithProductCountQuery کوئری دریافت دسته‌بندی‌ها با تعداد محصولات
type GetCategoriesWithProductCountQuery struct {
	IsActive *bool `json:"is_active"`
}

// GetFeaturedCategoriesQuery کوئری دریافت دسته‌بندی‌های ویژه
type GetFeaturedCategoriesQuery struct {
	Limit int `json:"limit" binding:"required,min=1,max=20"`
}

// ==================== Brand Queries ====================

// BrandFilter فیلتر برای جستجو در برندها
type BrandFilter struct {
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	IsActive       *bool  `json:"is_active"`
	IsFeatured     *bool  `json:"is_featured"`
	ShowInHomePage *bool  `json:"show_in_home_page"`
	ShowInMenu     *bool  `json:"show_in_menu"`
}

// GetBrandByIDQuery کوئری دریافت برند بر اساس ID
type GetBrandByIDQuery struct {
	ID uint `json:"id" binding:"required"`
}

// GetBrandBySlugQuery کوئری دریافت برند بر اساس slug
type GetBrandBySlugQuery struct {
	Slug string `json:"slug" binding:"required"`
}

// GetBrandsQuery کوئری دریافت لیست برندها
type GetBrandsQuery struct {
	Page      int         `json:"page" binding:"required,min=1"`
	PageSize  int         `json:"page_size" binding:"required,min=1,max=100"`
	Filter    BrandFilter `json:"filter"`
	SortBy    string      `json:"sort_by"`
	SortOrder string      `json:"sort_order" binding:"oneof=asc desc"`
}

// GetPublishedBrandsQuery کوئری دریافت برندهای منتشر شده
type GetPublishedBrandsQuery struct {
	Limit int `json:"limit" binding:"required,min=1,max=50"`
}

// GetFeaturedBrandsQuery کوئری دریافت برندهای ویژه
type GetFeaturedBrandsQuery struct {
	Limit int `json:"limit" binding:"required,min=1,max=20"`
}

// SearchBrandsQuery کوئری جستجو در برندها
type SearchBrandsQuery struct {
	Query string `json:"query" binding:"required"`
	Limit int    `json:"limit" binding:"required,min=1,max=20"`
}

// ==================== Review Queries ====================

// ReviewFilter فیلتر برای جستجو در نظرات
type ReviewFilter struct {
	ProductID   *uint  `json:"product_id"`
	UserID      *uint  `json:"user_id"`
	Status      string `json:"status"`
	Rating      *int   `json:"rating"`
	IsAnonymous *bool  `json:"is_anonymous"`
}

// GetReviewByIDQuery کوئری دریافت نظر بر اساس ID
type GetReviewByIDQuery struct {
	ID uint `json:"id" binding:"required"`
}

// GetReviewsQuery کوئری دریافت لیست نظرات
type GetReviewsQuery struct {
	Page      int          `json:"page" binding:"required,min=1"`
	PageSize  int          `json:"page_size" binding:"required,min=1,max=100"`
	Filter    ReviewFilter `json:"filter"`
	SortBy    string       `json:"sort_by"`
	SortOrder string       `json:"sort_order" binding:"oneof=asc desc"`
}

// GetReviewsByProductQuery کوئری دریافت نظرات یک محصول
type GetReviewsByProductQuery struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Page      int    `json:"page" binding:"required,min=1"`
	PageSize  int    `json:"page_size" binding:"required,min=1,max=100"`
	Status    string `json:"status"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" binding:"oneof=asc desc"`
}

// GetReviewsByUserQuery کوئری دریافت نظرات یک کاربر
type GetReviewsByUserQuery struct {
	UserID    uint   `json:"user_id" binding:"required"`
	Page      int    `json:"page" binding:"required,min=1"`
	PageSize  int    `json:"page_size" binding:"required,min=1,max=100"`
	Status    string `json:"status"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" binding:"oneof=asc desc"`
}

// GetPublishedReviewsQuery کوئری دریافت نظرات منتشر شده
type GetPublishedReviewsQuery struct {
	Page      int    `json:"page" binding:"required,min=1"`
	PageSize  int    `json:"page_size" binding:"required,min=1,max=100"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" binding:"oneof=asc desc"`
}

// GetPendingReviewsQuery کوئری دریافت نظرات در انتظار تایید
type GetPendingReviewsQuery struct {
	Page      int    `json:"page" binding:"required,min=1"`
	PageSize  int    `json:"page_size" binding:"required,min=1,max=100"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" binding:"oneof=asc desc"`
}

// GetAverageRatingQuery کوئری دریافت میانگین امتیاز محصول
type GetAverageRatingQuery struct {
	ProductID uint `json:"product_id" binding:"required"`
}

// GetReviewCountQuery کوئری دریافت تعداد نظرات محصول
type GetReviewCountQuery struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Status    string `json:"status"`
}

// ==================== Wishlist Queries ====================

// GetWishlistByUserQuery کوئری دریافت لیست علاقمندی‌های کاربر
type GetWishlistByUserQuery struct {
	UserID uint `json:"user_id" binding:"required"`
}

// GetWishlistProductsQuery کوئری دریافت محصولات لیست علاقمندی‌ها
type GetWishlistProductsQuery struct {
	UserID    uint   `json:"user_id" binding:"required"`
	Page      int    `json:"page" binding:"required,min=1"`
	PageSize  int    `json:"page_size" binding:"required,min=1,max=100"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" binding:"oneof=asc desc"`
}

// IsInWishlistQuery کوئری بررسی وجود محصول در لیست علاقمندی‌ها
type IsInWishlistQuery struct {
	UserID    uint `json:"user_id" binding:"required"`
	ProductID uint `json:"product_id" binding:"required"`
}

// GetWishlistCountQuery کوئری دریافت تعداد محصولات در لیست علاقمندی‌ها
type GetWishlistCountQuery struct {
	UserID uint `json:"user_id" binding:"required"`
}

// ==================== Stock Notification Queries ====================

// StockNotificationFilter فیلتر برای جستجو در اعلان‌های موجودی
type StockNotificationFilter struct {
	ProductID *uint  `json:"product_id"`
	UserID    *uint  `json:"user_id"`
	Status    string `json:"status"`
	Email     string `json:"email"`
}

// GetStockNotificationsQuery کوئری دریافت لیست اعلان‌های موجودی
type GetStockNotificationsQuery struct {
	Page      int                     `json:"page" binding:"required,min=1"`
	PageSize  int                     `json:"page_size" binding:"required,min=1,max=100"`
	Filter    StockNotificationFilter `json:"filter"`
	SortBy    string                  `json:"sort_by"`
	SortOrder string                  `json:"sort_order" binding:"oneof=asc desc"`
}

// GetStockNotificationsByProductQuery کوئری دریافت اعلان‌های یک محصول
type GetStockNotificationsByProductQuery struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Status    string `json:"status"`
}

// GetStockNotificationsByUserQuery کوئری دریافت اعلان‌های یک کاربر
type GetStockNotificationsByUserQuery struct {
	UserID uint   `json:"user_id" binding:"required"`
	Status string `json:"status"`
}

// GetPendingStockNotificationsQuery کوئری دریافت اعلان‌های در انتظار
type GetPendingStockNotificationsQuery struct {
	Page      int    `json:"page" binding:"required,min=1"`
	PageSize  int    `json:"page_size" binding:"required,min=1,max=100"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" binding:"oneof=asc desc"`
}

// ==================== Product QA Queries ====================

// ProductQAFilter فیلتر برای جستجو در پرسش و پاسخ‌ها
type ProductQAFilter struct {
	ProductID   *uint  `json:"product_id"`
	UserID      *uint  `json:"user_id"`
	Status      string `json:"status"`
	IsAnonymous *bool  `json:"is_anonymous"`
	ParentID    *uint  `json:"parent_id"` // برای تفکیک سوالات و پاسخ‌ها
}

// GetProductQAQuery کوئری دریافت لیست پرسش و پاسخ‌ها
type GetProductQAQuery struct {
	Page      int             `json:"page" binding:"required,min=1"`
	PageSize  int             `json:"page_size" binding:"required,min=1,max=100"`
	Filter    ProductQAFilter `json:"filter"`
	SortBy    string          `json:"sort_by"`
	SortOrder string          `json:"sort_order" binding:"oneof=asc desc"`
}

// GetQuestionsByProductQuery کوئری دریافت سوالات یک محصول
type GetQuestionsByProductQuery struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Page      int    `json:"page" binding:"required,min=1"`
	PageSize  int    `json:"page_size" binding:"required,min=1,max=100"`
	Status    string `json:"status"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" binding:"oneof=asc desc"`
}

// GetAnswersByQuestionQuery کوئری دریافت پاسخ‌های یک سوال
type GetAnswersByQuestionQuery struct {
	QuestionID uint   `json:"question_id" binding:"required"`
	Page       int    `json:"page" binding:"required,min=1"`
	PageSize   int    `json:"page_size" binding:"required,min=1,max=100"`
	Status     string `json:"status"`
	SortBy     string `json:"sort_by"`
	SortOrder  string `json:"sort_order" binding:"oneof=asc desc"`
}

// GetPublishedQAQuery کوئری دریافت پرسش و پاسخ‌های منتشر شده
type GetPublishedQAQuery struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Page      int    `json:"page" binding:"required,min=1"`
	PageSize  int    `json:"page_size" binding:"required,min=1,max=100"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" binding:"oneof=asc desc"`
}

// GetPendingQAQuery کوئری دریافت پرسش و پاسخ‌های در انتظار تایید
type GetPendingQAQuery struct {
	Page      int    `json:"page" binding:"required,min=1"`
	PageSize  int    `json:"page_size" binding:"required,min=1,max=100"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order" binding:"oneof=asc desc"`
}

// GetQuestionCountQuery کوئری دریافت تعداد سوالات محصول
type GetQuestionCountQuery struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Status    string `json:"status"`
}

// GetAnswerCountQuery کوئری دریافت تعداد پاسخ‌های سوال
type GetAnswerCountQuery struct {
	QuestionID uint   `json:"question_id" binding:"required"`
	Status     string `json:"status"`
}
