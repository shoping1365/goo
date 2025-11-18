package repository

import (
	"context"
	"my-go-backend/internal/models"
	"time"
)

// ==================== Generic Repository Interfaces ====================

// BaseRepository اینترفیس پایه برای تمام repository ها
type BaseRepository[T any] interface {
	// Create ایجاد رکورد جدید
	Create(ctx context.Context, entity *T) error

	// GetByID دریافت رکورد بر اساس ID
	GetByID(ctx context.Context, id uint) (*T, error)

	// GetAll دریافت تمام رکوردها
	GetAll(ctx context.Context) ([]T, error)

	// Update به‌روزرسانی رکورد
	Update(ctx context.Context, entity *T) error

	// Delete حذف رکورد
	Delete(ctx context.Context, id uint) error

	// Count تعداد رکوردها
	Count(ctx context.Context) (int64, error)
}

// SearchableRepository اینترفیس برای repository های قابل جستجو
type SearchableRepository[T any] interface {
	BaseRepository[T]

	// Search جستجو در رکوردها
	Search(ctx context.Context, query string, limit, offset int) ([]T, error)

	// Filter فیلتر کردن رکوردها
	Filter(ctx context.Context, filters map[string]interface{}, limit, offset int) ([]T, error)
}

// ==================== Brand Repository ====================

// BrandRepositoryInterface اینترفیس repository برای برندها
type BrandRepositoryInterface interface {
	BaseRepository[models.Brand]

	// GetBySlug دریافت برند بر اساس slug
	GetBySlug(ctx context.Context, slug string) (*models.Brand, error)

	// GetPublished دریافت برندهای منتشر شده
	GetPublished(ctx context.Context) ([]models.Brand, error)

	// GetByHomePage دریافت برندهای نمایش داده شده در صفحه اصلی
	GetByHomePage(ctx context.Context) ([]models.Brand, error)

	// GetByMenu دریافت برندهای نمایش داده شده در منو
	GetByMenu(ctx context.Context) ([]models.Brand, error)

	// SearchByName جستجو بر اساس نام
	SearchByName(ctx context.Context, name string) ([]models.Brand, error)
}

// ==================== Category Repository ====================

// CategoryRepositoryInterface اینترفیس repository برای دسته‌بندی‌ها
type CategoryRepositoryInterface interface {
	BaseRepository[models.Category]

	// GetBySlug دریافت دسته‌بندی بر اساس slug
	GetBySlug(ctx context.Context, slug string) (*models.Category, error)

	// GetByParentID دریافت زیردسته‌ها
	GetByParentID(ctx context.Context, parentID *uint) ([]models.Category, error)

	// GetRootCategories دریافت دسته‌های اصلی
	GetRootCategories(ctx context.Context) ([]models.Category, error)

	// GetPublished دریافت دسته‌های منتشر شده
	GetPublished(ctx context.Context) ([]models.Category, error)

	// GetByHomePage دریافت دسته‌های نمایش داده شده در صفحه اصلی
	GetByHomePage(ctx context.Context) ([]models.Category, error)

	// GetByMenu دریافت دسته‌های نمایش داده شده در منو
	GetByMenu(ctx context.Context) ([]models.Category, error)

	// GetHierarchy دریافت ساختار سلسله مراتبی
	GetHierarchy(ctx context.Context) ([]models.Category, error)

	// GetWithProductCount دریافت دسته‌ها با تعداد محصولات
	GetWithProductCount(ctx context.Context) ([]models.Category, error)
}

// ==================== Product Repository ====================

// ProductRepositoryInterface اینترفیس repository برای محصولات
type ProductRepositoryInterface interface {
	BaseRepository[models.Product]

	// GetBySlug دریافت محصول بر اساس slug
	GetBySlug(ctx context.Context, slug string) (*models.Product, error)

	// GetBySKU دریافت محصول بر اساس SKU
	GetBySKU(ctx context.Context, sku string) (*models.Product, error)

	// GetByCategory دریافت محصولات یک دسته‌بندی
	GetByCategory(ctx context.Context, categoryID uint, limit, offset int) ([]models.Product, error)

	// GetByBrand دریافت محصولات یک برند
	GetByBrand(ctx context.Context, brandID uint, limit, offset int) ([]models.Product, error)

	// GetPublished دریافت محصولات منتشر شده
	GetPublished(ctx context.Context, limit, offset int) ([]models.Product, error)

	// Keyset Pagination – دریافت محصولات با کلید (created_at,id) برای سرعت بالا
	GetPublishedAfter(ctx context.Context, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error)
	GetByCategoryAfter(ctx context.Context, categoryID uint, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error)
	GetByBrandAfter(ctx context.Context, brandID uint, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error)

	// GetNewProducts دریافت محصولات جدید
	GetNewProducts(ctx context.Context, limit int) ([]models.Product, error)

	// GetDiscountedProducts دریافت محصولات تخفیف‌دار
	GetDiscountedProducts(ctx context.Context, limit int) ([]models.Product, error)

	// SearchByName جستجو بر اساس نام
	SearchByName(ctx context.Context, name string, limit, offset int) ([]models.Product, error)

	// GetByPriceRange دریافت محصولات در محدوده قیمت
	GetByPriceRange(ctx context.Context, minPrice, maxPrice float64, limit, offset int) ([]models.Product, error)

	// GetRelatedProducts دریافت محصولات مرتبط
	GetRelatedProducts(ctx context.Context, productID uint, limit int) ([]models.Product, error)

	// UpdateStock به‌روزرسانی موجودی
	UpdateStock(ctx context.Context, productID uint, quantity int) error

	// GetLowStockProducts دریافت محصولات با موجودی کم
	GetLowStockProducts(ctx context.Context) ([]models.Product, error)
}

// ==================== Attribute Repository ====================

// AttributeRepositoryInterface اینترفیس repository برای ویژگی‌ها
type AttributeRepositoryInterface interface {
	BaseRepository[models.Attribute]

	// GetByCode دریافت ویژگی بر اساس کد
	GetByCode(ctx context.Context, code string) (*models.Attribute, error)

	// GetActive دریافت ویژگی‌های فعال
	GetActive(ctx context.Context) ([]models.Attribute, error)

	// GetFilterable دریافت ویژگی‌های قابل فیلتر
	GetFilterable(ctx context.Context) ([]models.Attribute, error)

	// GetRequired دریافت ویژگی‌های اجباری
	GetRequired(ctx context.Context) ([]models.Attribute, error)

	// GetWithValues دریافت ویژگی‌ها با مقادیر
	GetWithValues(ctx context.Context) ([]models.Attribute, error)

	// GetByDataType دریافت ویژگی‌ها بر اساس نوع داده
	GetByDataType(ctx context.Context, dataType string) ([]models.Attribute, error)
}

// ==================== Attribute Group Repository ====================

// AttributeGroupRepositoryInterface اینترفیس repository برای گروه‌های ویژگی
type AttributeGroupRepositoryInterface interface {
	BaseRepository[models.AttributeGroup]

	// GetByCategory دریافت گروه‌های ویژگی یک دسته‌بندی
	GetByCategory(ctx context.Context, categoryID uint) ([]models.AttributeGroup, error)

	// GetWithAttributes دریافت گروه‌ها با ویژگی‌ها
	GetWithAttributes(ctx context.Context) ([]models.AttributeGroup, error)

	// GetByCategoryWithAttributes دریافت گروه‌های یک دسته‌بندی با ویژگی‌ها
	GetByCategoryWithAttributes(ctx context.Context, categoryID uint) ([]models.AttributeGroup, error)
}

// ==================== Review Repository ====================

// ReviewRepositoryInterface اینترفیس repository برای نظرات
type ReviewRepositoryInterface interface {
	BaseRepository[models.Review]

	// GetByProduct دریافت نظرات یک محصول
	GetByProduct(ctx context.Context, productID uint, limit, offset int) ([]models.Review, error)

	// GetByCustomer دریافت نظرات یک مشتری
	GetByCustomer(ctx context.Context, customerID uint, limit, offset int) ([]models.Review, error)

	// GetPublished دریافت نظرات منتشر شده
	GetPublished(ctx context.Context, limit, offset int) ([]models.Review, error)

	// GetPending دریافت نظرات در انتظار تایید
	GetPending(ctx context.Context, limit, offset int) ([]models.Review, error)

	// GetByRating دریافت نظرات بر اساس امتیاز
	GetByRating(ctx context.Context, productID uint, rating int, limit, offset int) ([]models.Review, error)

	// GetAverageRating دریافت میانگین امتیاز محصول
	GetAverageRating(ctx context.Context, productID uint) (float64, error)

	// GetReviewCount دریافت تعداد نظرات محصول
	GetReviewCount(ctx context.Context, productID uint) (int64, error)

	// UpdateStatus به‌روزرسانی وضعیت نظر
	UpdateStatus(ctx context.Context, reviewID uint, status string) error

	// IncrementHelpfulCount افزایش تعداد مفید بودن
	IncrementHelpfulCount(ctx context.Context, reviewID uint) error

	// HasHelpfulVote بررسی اینکه کاربر قبلاً برای این نظر رأی داده است یا خیر
	HasHelpfulVote(ctx context.Context, reviewID uint, userID uint) (bool, error)

	// CreateHelpfulVote ایجاد رکورد رأی مفید برای جلوگیری از رأی تکراری
	CreateHelpfulVote(ctx context.Context, reviewID uint, userID uint) error
}

// ==================== Wishlist Repository ====================

// WishlistRepositoryInterface اینترفیس repository برای لیست علاقمندی‌ها
type WishlistRepositoryInterface interface {
	BaseRepository[models.Wishlist]

	// GetByUser دریافت لیست علاقمندی‌های کاربر
	GetByUser(ctx context.Context, userID uint) (*models.Wishlist, error)

	// AddProduct افزودن محصول به لیست علاقمندی‌ها
	AddProduct(ctx context.Context, userID, productID uint) error

	// RemoveProduct حذف محصول از لیست علاقمندی‌ها
	RemoveProduct(ctx context.Context, userID, productID uint) error

	// IsInWishlist بررسی وجود محصول در لیست علاقمندی‌ها
	IsInWishlist(ctx context.Context, userID, productID uint) (bool, error)

	// GetWishlistProducts دریافت محصولات لیست علاقمندی‌ها
	GetWishlistProducts(ctx context.Context, userID uint) ([]models.Product, error)

	// GetWishlistCount دریافت تعداد محصولات در لیست علاقمندی‌ها
	GetWishlistCount(ctx context.Context, userID uint) (int64, error)
}

// ==================== Stock Notification Repository ====================

// StockNotificationRepositoryInterface اینترفیس repository برای اعلان‌های موجودی
type StockNotificationRepositoryInterface interface {
	BaseRepository[models.StockNotification]

	// GetByProduct دریافت اعلان‌های یک محصول
	GetByProduct(ctx context.Context, productID uint) ([]models.StockNotification, error)

	// GetByUser دریافت اعلان‌های یک کاربر
	GetByUser(ctx context.Context, userID uint) ([]models.StockNotification, error)

	// GetPending دریافت اعلان‌های در انتظار
	GetPending(ctx context.Context) ([]models.StockNotification, error)

	// GetPendingByProduct دریافت اعلان‌های در انتظار یک محصول
	GetPendingByProduct(ctx context.Context, productID uint) ([]models.StockNotification, error)

	// GetSent دریافت اعلان‌های ارسال شده
	GetSent(ctx context.Context) ([]models.StockNotification, error)

	// MarkAsSent علامت‌گذاری به عنوان ارسال شده
	MarkAsSent(ctx context.Context, notificationID uint) error

	// DeleteByProduct حذف اعلان‌های یک محصول
	DeleteByProduct(ctx context.Context, productID uint) error

	// DeleteByUser حذف اعلان‌های یک کاربر
	DeleteByUser(ctx context.Context, userID uint) error
}

// ==================== Product QA Repository ====================

// ProductQARepositoryInterface اینترفیس repository برای پرسش و پاسخ محصولات
type ProductQARepositoryInterface interface {
	BaseRepository[models.ProductQA]

	// GetByProduct دریافت پرسش و پاسخ‌های یک محصول
	GetByProduct(ctx context.Context, productID uint, limit, offset int) ([]models.ProductQA, error)

	// GetByUser دریافت پرسش و پاسخ‌های یک کاربر
	GetByUser(ctx context.Context, userID uint, limit, offset int) ([]models.ProductQA, error)

	// GetQuestions دریافت سوالات
	GetQuestions(ctx context.Context, productID uint, limit, offset int) ([]models.ProductQA, error)

	// GetAnswers دریافت پاسخ‌ها
	GetAnswers(ctx context.Context, questionID uint, limit, offset int) ([]models.ProductQA, error)

	// GetPublished دریافت پرسش و پاسخ‌های منتشر شده
	GetPublished(ctx context.Context, productID uint, limit, offset int) ([]models.ProductQA, error)

	// GetPending دریافت پرسش و پاسخ‌های در انتظار تایید
	GetPending(ctx context.Context, limit, offset int) ([]models.ProductQA, error)

	// UpdateStatus به‌روزرسانی وضعیت
	UpdateStatus(ctx context.Context, qaID uint, status string) error

	// GetQuestionCount دریافت تعداد سوالات محصول
	GetQuestionCount(ctx context.Context, productID uint) (int64, error)

	// GetAnswerCount دریافت تعداد پاسخ‌های سوال
	GetAnswerCount(ctx context.Context, questionID uint) (int64, error)
}

// ==================== Menu Repository ====================

// MenuRepositoryInterface اینترفیس repository برای منوها
type MenuRepositoryInterface interface {
	// ==================== Menu Operations ====================

	// CreateMenu ایجاد منوی جدید
	CreateMenu(ctx context.Context, menu *models.Menu) error

	// GetMenuByID دریافت منو بر اساس ID
	GetMenuByID(ctx context.Context, id uint) (*models.Menu, error)

	// GetMenuBySlug دریافت منو بر اساس slug
	GetMenuBySlug(ctx context.Context, slug string) (*models.Menu, error)

	// GetAllMenus دریافت تمام منوها
	GetAllMenus(ctx context.Context) ([]models.Menu, error)

	// GetEnabledMenus دریافت منوهای فعال
	GetEnabledMenus(ctx context.Context) ([]models.Menu, error)

	// UpdateMenu به‌روزرسانی منو
	UpdateMenu(ctx context.Context, menu *models.Menu) error

	// DeleteMenu حذف منو
	DeleteMenu(ctx context.Context, id uint) error

	// ==================== MenuItem Operations ====================

	// CreateMenuItem ایجاد آیتم منوی جدید
	CreateMenuItem(ctx context.Context, item *models.MenuItem) error

	// GetMenuItemByID دریافت آیتم منو بر اساس ID
	GetMenuItemByID(ctx context.Context, id uint) (*models.MenuItem, error)

	// GetMenuItemsByMenuID دریافت آیتم‌های منو بر اساس ID منو
	GetMenuItemsByMenuID(ctx context.Context, menuID uint) ([]models.MenuItem, error)

	// UpdateMenuItem به‌روزرسانی آیتم منو
	UpdateMenuItem(ctx context.Context, item *models.MenuItem) error

	// DeleteMenuItem حذف آیتم منو
	DeleteMenuItem(ctx context.Context, id uint) error

	// ReorderMenuItems تغییر ترتیب آیتم‌های منو
	ReorderMenuItems(ctx context.Context, menuID uint, itemOrders []struct {
		ID    uint `json:"id"`
		Order int  `json:"order"`
	}) error

	// ==================== MenuColumn Operations ====================

	// CreateMenuColumn ایجاد ستون منوی جدید
	CreateMenuColumn(ctx context.Context, column *models.MenuColumn) error

	// GetMenuColumnByID دریافت ستون منو بر اساس ID
	GetMenuColumnByID(ctx context.Context, id uint) (*models.MenuColumn, error)

	// GetMenuColumnsByMenuID دریافت ستون‌های منو بر اساس ID منو
	GetMenuColumnsByMenuID(ctx context.Context, menuID uint) ([]models.MenuColumn, error)

	// UpdateMenuColumn به‌روزرسانی ستون منو
	UpdateMenuColumn(ctx context.Context, column *models.MenuColumn) error

	// DeleteMenuColumn حذف ستون منو
	DeleteMenuColumn(ctx context.Context, id uint) error

	// ReorderMenuColumns تغییر ترتیب ستون‌های منو
	ReorderMenuColumns(ctx context.Context, menuID uint, columnOrders []struct {
		ID    uint `json:"id"`
		Order int  `json:"order"`
	}) error

	// ==================== MenuLocation Operations ====================

	// CreateMenuLocation ایجاد موقعیت منوی جدید
	CreateMenuLocation(ctx context.Context, location *models.MenuLocation) error

	// GetMenuLocationByID دریافت موقعیت منو بر اساس ID
	GetMenuLocationByID(ctx context.Context, id uint) (*models.MenuLocation, error)

	// GetMenuLocationBySlug دریافت موقعیت منو بر اساس slug
	GetMenuLocationBySlug(ctx context.Context, slug string) (*models.MenuLocation, error)

	// GetAllMenuLocations دریافت تمام موقعیت‌های منو
	GetAllMenuLocations(ctx context.Context) ([]models.MenuLocation, error)

	// UpdateMenuLocation به‌روزرسانی موقعیت منو
	UpdateMenuLocation(ctx context.Context, location *models.MenuLocation) error

	// DeleteMenuLocation حذف موقعیت منو
	DeleteMenuLocation(ctx context.Context, id uint) error

	// ==================== MenuAssignment Operations ====================

	// CreateMenuAssignment ایجاد انتساب منوی جدید
	CreateMenuAssignment(ctx context.Context, assignment *models.MenuAssignment) error

	// GetMenuAssignmentByID دریافت انتساب منو بر اساس ID
	GetMenuAssignmentByID(ctx context.Context, id uint) (*models.MenuAssignment, error)

	// GetMenuAssignmentsByLocationID دریافت انتساب‌های یک موقعیت
	GetMenuAssignmentsByLocationID(ctx context.Context, locationID uint) ([]models.MenuAssignment, error)

	// GetMenuAssignmentsByMenuID دریافت انتساب‌های یک منو
	GetMenuAssignmentsByMenuID(ctx context.Context, menuID uint) ([]models.MenuAssignment, error)

	// UpdateMenuAssignment به‌روزرسانی انتساب منو
	UpdateMenuAssignment(ctx context.Context, assignment *models.MenuAssignment) error

	// DeleteMenuAssignment حذف انتساب منو
	DeleteMenuAssignment(ctx context.Context, id uint) error

	// AssignMenuToLocation انتساب منو به موقعیت
	AssignMenuToLocation(ctx context.Context, menuID, locationID uint, order int) error

	// UnassignMenuFromLocation حذف انتساب منو از موقعیت
	UnassignMenuFromLocation(ctx context.Context, menuID, locationID uint) error

	// GetMenusByLocation دریافت منوهای یک موقعیت
	GetMenusByLocation(ctx context.Context, locationSlug string) ([]models.Menu, error)

	// ==================== Content Operations ====================

	// GetMenuContentPages دریافت صفحات قابل استفاده در منو
	GetMenuContentPages(ctx context.Context) ([]struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Slug  string `json:"slug"`
	}, error)

	// GetMenuContentPosts دریافت پست‌های قابل استفاده در منو
	GetMenuContentPosts(ctx context.Context) ([]struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Slug  string `json:"slug"`
	}, error)

	// GetMenuContentCategories دریافت دسته‌بندی‌های قابل استفاده در منو
	GetMenuContentCategories(ctx context.Context) ([]struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}, error)

	// GetMenuContentProductCategories دریافت دسته‌بندی‌های محصولات قابل استفاده در منو
	GetMenuContentProductCategories(ctx context.Context) ([]struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}, error)
}
