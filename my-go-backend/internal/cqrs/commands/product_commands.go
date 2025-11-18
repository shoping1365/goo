package commands

import (
	"time"
)

// ==================== Product Commands ====================

// CreateProductCommand دستور ایجاد محصول جدید
type CreateProductCommand struct {
	Name             string                 `json:"name" binding:"required"`
	Slug             string                 `json:"slug" binding:"required"`
	SKU              string                 `json:"sku" binding:"required"`
	Description      string                 `json:"description"`
	ShortDescription string                 `json:"short_description"`
	Price            float64                `json:"price" binding:"required,min=0"`
	SalePrice        *float64               `json:"sale_price"`
	CostPrice        float64                `json:"cost_price"`
	StockQuantity    int                    `json:"stock_quantity" binding:"required,min=0"`
	Weight           float64                `json:"weight"`
	Dimensions       string                 `json:"dimensions"`
	BrandID          *uint                  `json:"brand_id"`
	CategoryID       uint                   `json:"category_id" binding:"required"`
	Status           string                 `json:"status" binding:"required,oneof=draft published archived"`
	IsActive         bool                   `json:"is_active"`
	IsFeatured       bool                   `json:"is_featured"`
	IsNew            bool                   `json:"is_new"`
	IsOnSale         bool                   `json:"is_on_sale"`
	MetaTitle        string                 `json:"meta_title"`
	MetaDescription  string                 `json:"meta_description"`
	MetaKeywords     string                 `json:"meta_keywords"`
	Tags             []string               `json:"tags"`
	AttributeValues  map[string]interface{} `json:"attribute_values"`
	Images           []string               `json:"images"`
	CreatedBy        uint                   `json:"created_by"`
}

// UpdateProductCommand دستور به‌روزرسانی محصول
type UpdateProductCommand struct {
	ID               uint                   `json:"id" binding:"required"`
	Name             string                 `json:"name"`
	Slug             string                 `json:"slug"`
	SKU              string                 `json:"sku"`
	Description      string                 `json:"description"`
	ShortDescription string                 `json:"short_description"`
	Price            float64                `json:"price"`
	SalePrice        *float64               `json:"sale_price"`
	CostPrice        float64                `json:"cost_price"`
	StockQuantity    int                    `json:"stock_quantity"`
	Weight           float64                `json:"weight"`
	Dimensions       string                 `json:"dimensions"`
	BrandID          *uint                  `json:"brand_id"`
	CategoryID       uint                   `json:"category_id"`
	Status           string                 `json:"status"`
	IsActive         bool                   `json:"is_active"`
	IsFeatured       bool                   `json:"is_featured"`
	IsNew            bool                   `json:"is_new"`
	IsOnSale         bool                   `json:"is_on_sale"`
	MetaTitle        string                 `json:"meta_title"`
	MetaDescription  string                 `json:"meta_description"`
	MetaKeywords     string                 `json:"meta_keywords"`
	Tags             []string               `json:"tags"`
	AttributeValues  map[string]interface{} `json:"attribute_values"`
	Images           []string               `json:"images"`
	UpdatedBy        uint                   `json:"updated_by"`
}

// DeleteProductCommand دستور حذف محصول
type DeleteProductCommand struct {
	ID        uint `json:"id" binding:"required"`
	DeletedBy uint `json:"deleted_by"`
}

// UpdateProductStockCommand دستور به‌روزرسانی موجودی محصول
type UpdateProductStockCommand struct {
	ID        uint `json:"id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
	UpdatedBy uint `json:"updated_by"`
}

// PublishProductCommand دستور انتشار محصول
type PublishProductCommand struct {
	ID          uint `json:"id" binding:"required"`
	PublishedBy uint `json:"published_by"`
}

// ArchiveProductCommand دستور آرشیو محصول
type ArchiveProductCommand struct {
	ID         uint `json:"id" binding:"required"`
	ArchivedBy uint `json:"archived_by"`
}

// BulkUpdateProductStatusCommand دستور به‌روزرسانی وضعیت دسته‌ای محصولات
type BulkUpdateProductStatusCommand struct {
	ProductIDs []uint `json:"product_ids" binding:"required,min=1"`
	Status     string `json:"status" binding:"required,oneof=draft published archived"`
	UpdatedBy  uint   `json:"updated_by"`
}

// ==================== Category Commands ====================

// CreateCategoryCommand دستور ایجاد دسته‌بندی جدید
type CreateCategoryCommand struct {
	Name            string `json:"name" binding:"required"`
	Slug            string `json:"slug" binding:"required"`
	Description     string `json:"description"`
	ParentID        *uint  `json:"parent_id"`
	Image           string `json:"image"`
	IsActive        bool   `json:"is_active"`
	IsFeatured      bool   `json:"is_featured"`
	ShowInHomePage  bool   `json:"show_in_home_page"`
	ShowInMenu      bool   `json:"show_in_menu"`
	SortOrder       int    `json:"sort_order"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeywords    string `json:"meta_keywords"`
	CreatedBy       uint   `json:"created_by"`
}

// UpdateCategoryCommand دستور به‌روزرسانی دسته‌بندی
type UpdateCategoryCommand struct {
	ID              uint   `json:"id" binding:"required"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	ParentID        *uint  `json:"parent_id"`
	Image           string `json:"image"`
	IsActive        bool   `json:"is_active"`
	IsFeatured      bool   `json:"is_featured"`
	ShowInHomePage  bool   `json:"show_in_home_page"`
	ShowInMenu      bool   `json:"show_in_menu"`
	SortOrder       int    `json:"sort_order"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeywords    string `json:"meta_keywords"`
	UpdatedBy       uint   `json:"updated_by"`
}

// DeleteCategoryCommand دستور حذف دسته‌بندی
type DeleteCategoryCommand struct {
	ID        uint `json:"id" binding:"required"`
	DeletedBy uint `json:"deleted_by"`
}

// ReorderCategoriesCommand دستور تغییر ترتیب دسته‌بندی‌ها
type ReorderCategoriesCommand struct {
	CategoryOrders []struct {
		ID    uint `json:"id"`
		Order int  `json:"order"`
	} `json:"category_orders" binding:"required,min=1"`
	UpdatedBy uint `json:"updated_by"`
}

// ==================== Brand Commands ====================

// CreateBrandCommand دستور ایجاد برند جدید
type CreateBrandCommand struct {
	Name            string `json:"name" binding:"required"`
	Slug            string `json:"slug" binding:"required"`
	Description     string `json:"description"`
	Logo            string `json:"logo"`
	Website         string `json:"website"`
	IsActive        bool   `json:"is_active"`
	IsFeatured      bool   `json:"is_featured"`
	ShowInHomePage  bool   `json:"show_in_home_page"`
	ShowInMenu      bool   `json:"show_in_menu"`
	SortOrder       int    `json:"sort_order"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeywords    string `json:"meta_keywords"`
	CreatedBy       uint   `json:"created_by"`
}

// UpdateBrandCommand دستور به‌روزرسانی برند
type UpdateBrandCommand struct {
	ID              uint   `json:"id" binding:"required"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	Logo            string `json:"logo"`
	Website         string `json:"website"`
	IsActive        bool   `json:"is_active"`
	IsFeatured      bool   `json:"is_featured"`
	ShowInHomePage  bool   `json:"show_in_home_page"`
	ShowInMenu      bool   `json:"show_in_menu"`
	SortOrder       int    `json:"sort_order"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	MetaKeywords    string `json:"meta_keywords"`
	UpdatedBy       uint   `json:"updated_by"`
}

// DeleteBrandCommand دستور حذف برند
type DeleteBrandCommand struct {
	ID        uint `json:"id" binding:"required"`
	DeletedBy uint `json:"deleted_by"`
}

// ==================== Review Commands ====================

// CreateReviewCommand دستور ایجاد نظر جدید
type CreateReviewCommand struct {
	ProductID   uint   `json:"product_id" binding:"required"`
	UserID      uint   `json:"user_id" binding:"required"`
	Rating      int    `json:"rating" binding:"required,min=1,max=5"`
	Title       string `json:"title"`
	Comment     string `json:"comment"`
	IsAnonymous bool   `json:"is_anonymous"`
	Status      string `json:"status" binding:"required,oneof=pending approved rejected"`
}

// UpdateReviewCommand دستور به‌روزرسانی نظر
type UpdateReviewCommand struct {
	ID          uint   `json:"id" binding:"required"`
	Rating      int    `json:"rating" binding:"min=1,max=5"`
	Title       string `json:"title"`
	Comment     string `json:"comment"`
	IsAnonymous bool   `json:"is_anonymous"`
	Status      string `json:"status" binding:"oneof=pending approved rejected"`
	UpdatedBy   uint   `json:"updated_by"`
}

// DeleteReviewCommand دستور حذف نظر
type DeleteReviewCommand struct {
	ID        uint `json:"id" binding:"required"`
	DeletedBy uint `json:"deleted_by"`
}

// ApproveReviewCommand دستور تایید نظر
type ApproveReviewCommand struct {
	ID         uint `json:"id" binding:"required"`
	ApprovedBy uint `json:"approved_by"`
}

// RejectReviewCommand دستور رد نظر
type RejectReviewCommand struct {
	ID         uint   `json:"id" binding:"required"`
	Reason     string `json:"reason"`
	RejectedBy uint   `json:"rejected_by"`
}

// ==================== Wishlist Commands ====================

// AddToWishlistCommand دستور افزودن محصول به لیست علاقمندی‌ها
type AddToWishlistCommand struct {
	UserID    uint `json:"user_id" binding:"required"`
	ProductID uint `json:"product_id" binding:"required"`
}

// RemoveFromWishlistCommand دستور حذف محصول از لیست علاقمندی‌ها
type RemoveFromWishlistCommand struct {
	UserID    uint `json:"user_id" binding:"required"`
	ProductID uint `json:"product_id" binding:"required"`
}

// ClearWishlistCommand دستور پاک کردن لیست علاقمندی‌ها
type ClearWishlistCommand struct {
	UserID uint `json:"user_id" binding:"required"`
}

// ==================== Stock Notification Commands ====================

// CreateStockNotificationCommand دستور ایجاد اعلان موجودی
type CreateStockNotificationCommand struct {
	UserID    uint   `json:"user_id" binding:"required"`
	ProductID uint   `json:"product_id" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
}

// MarkNotificationAsSentCommand دستور علامت‌گذاری اعلان به عنوان ارسال شده
type MarkNotificationAsSentCommand struct {
	ID     uint      `json:"id" binding:"required"`
	SentAt time.Time `json:"sent_at"`
	SentBy uint      `json:"sent_by"`
}

// DeleteStockNotificationCommand دستور حذف اعلان موجودی
type DeleteStockNotificationCommand struct {
	ID        uint `json:"id" binding:"required"`
	DeletedBy uint `json:"deleted_by"`
}

// ==================== Product QA Commands ====================

// CreateQuestionCommand دستور ایجاد سوال جدید
type CreateQuestionCommand struct {
	ProductID   uint   `json:"product_id" binding:"required"`
	UserID      uint   `json:"user_id" binding:"required"`
	Question    string `json:"question" binding:"required"`
	IsAnonymous bool   `json:"is_anonymous"`
	Status      string `json:"status" binding:"required,oneof=pending approved rejected"`
}

// CreateAnswerCommand دستور ایجاد پاسخ جدید
type CreateAnswerCommand struct {
	QuestionID  uint   `json:"question_id" binding:"required"`
	UserID      uint   `json:"user_id" binding:"required"`
	Answer      string `json:"answer" binding:"required"`
	IsAnonymous bool   `json:"is_anonymous"`
	Status      string `json:"status" binding:"required,oneof=pending approved rejected"`
}

// UpdateQAStatusCommand دستور به‌روزرسانی وضعیت پرسش و پاسخ
type UpdateQAStatusCommand struct {
	ID        uint   `json:"id" binding:"required"`
	Status    string `json:"status" binding:"required,oneof=pending approved rejected"`
	UpdatedBy uint   `json:"updated_by"`
}

// DeleteQACommand دستور حذف پرسش و پاسخ
type DeleteQACommand struct {
	ID        uint `json:"id" binding:"required"`
	DeletedBy uint `json:"deleted_by"`
}
