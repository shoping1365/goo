package models

import (
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// Product represents a product in the system
type Product struct {
	ID     int    `json:"id"`
	UUID   string `json:"uuid" gorm:"type:uuid;uniqueIndex;not null"`
	Slug   string `json:"slug" gorm:"type:varchar(255);uniqueIndex;not null"`
	SKU    string `json:"sku" gorm:"type:varchar(100);uniqueIndex"` // Stock Keeping Unit, auto-generated
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
	Image  string `json:"image"`
	// ImageURL: لینک تصویر اصلی از جدول product_images (خواندنی)
	ImageURL            string     `json:"image_url,omitempty" gorm:"->;column:image_url"`
	Price               float64    `json:"price"`
	OldPrice            float64    `json:"old_price"`
	Cost                float64    `json:"cost"`
	Profit              float64    `json:"profit"`
	DiscountPercent     float64    `json:"discount_percent"`
	DiscountAmount      float64    `json:"discount_amount"`
	SalePrice           float64    `json:"sale_price"`
	SaleStartAt         *time.Time `json:"sale_start_at"`
	SaleEndAt           *time.Time `json:"sale_end_at"`
	WholesalePrice      float64    `json:"wholesale_price"`
	WholesaleSalePrice  float64    `json:"wholesale_sale_price"`
	Description         string     `json:"description"`
	FullDescription     string     `json:"full_description"`
	Status              string     `json:"status"`
	IsNew               bool       `json:"is_new"`
	Weight              float64    `json:"weight"`
	Length              float64    `json:"length"`
	Width               float64    `json:"width"`
	Height              float64    `json:"height"`
	ShippingCost        float64    `json:"shipping_cost"`
	ShippingTime        int        `json:"shipping_time"`
	StockQuantity       int        `json:"stock_quantity"`
	MinStockQuantity    int        `json:"min_stock_quantity"`
	MaxStockQuantity    int        `json:"max_stock_quantity"`
	StockStatus         string     `json:"stock_status"`
	TrackInventory      bool       `json:"track_inventory"`
	ShowStockToCustomer bool       `json:"show_stock_to_customer"`
	DisableBuyButton    bool       `json:"disable_buy_button"`
	AllowReservation    bool       `json:"allow_reservation" gorm:"default:false"`
	CallForPrice        bool       `json:"call_for_price"`
	SeoTitle            string     `json:"seo_title"`
	MetaDescription     string     `json:"meta_description"`
	MetaKeywords        string     `json:"meta_keywords"`
	CategoryID          *uint      `json:"category_id"`
	BrandID             *uint      `json:"brand_id"`
	IsActive            bool       `json:"is_active" gorm:"default:true"`
	IsFeatured          bool       `json:"is_featured" gorm:"default:false"`
	IsOnSale            bool       `json:"is_on_sale" gorm:"default:false"`
	MetaTitle           string     `json:"meta_title"`
	Tags                string     `json:"tags"`
	URL                 string     `json:"url" gorm:"type:varchar(500)"` // URL کامل محصول
	UpdatedBy           *uint      `json:"updated_by"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	// حذف فیلد قدیمی video_url؛ ویدیوها در جدول product_videos نگهداری می‌شوند
	// VideoURL string `json:"video_url" gorm:"column:video_url"`
	DefaultWarehouseID *uint `json:"default_warehouse_id" gorm:"index"`

	// Variant image display settings
	ShowVariantImageInList         bool `json:"show_variant_image_in_list" gorm:"default:false"`
	ChangeMainImageOnVariantSelect bool `json:"change_main_image_on_variant_select" gorm:"default:false"`
	ShowAllVariantImagesInGallery  bool `json:"show_all_variant_images_in_gallery" gorm:"default:false"`
	AutoCompressVariantImages      bool `json:"auto_compress_variant_images" gorm:"default:false"`

	// Relations
	Category           *Category              `json:"category,omitempty"`
	Brand              *Brand                 `json:"brand,omitempty"`
	Images             []ProductImage         `json:"images,omitempty"`
	Variants           []ProductVariant       `json:"variants,omitempty"`
	Specifications     []ProductSpecification `json:"specifications,omitempty"`
	RelatedProducts    []Product              `json:"related_products,omitempty" gorm:"-"`
	Videos             []ProductVideo         `json:"videos,omitempty"`
	ComplementProducts []Product              `json:"complement_products,omitempty" gorm:"-"`
	// فروش ویژه مرحله‌ای
	SpecialOffers []ProductSpecialOffer `json:"special_offers,omitempty" gorm:"foreignKey:ProductID"`
}

// ProductSpecialOffer پله‌های فروش ویژه: قیمت و تعداد مجاز
type ProductSpecialOffer struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ProductImage represents a product image
type ProductImage struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	ImageURL  string    `json:"image_url"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
}

// ProductVariant represents a product variant
type ProductVariant struct {
	ID              int       `json:"id"`
	ProductID       int       `json:"product_id"`
	Name            string    `json:"name"`
	Value           string    `json:"value"`
	PriceAdjustment float64   `json:"price_adjustment"`
	StockQuantity   int       `json:"stock_quantity"`
	CreatedAt       time.Time `json:"created_at"`
}

// ProductSpecification represents a product specification
type ProductSpecification struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

// ProductVideo represents a product video
type ProductVideo struct {
	ID            int       `json:"id"`
	ProductID     int       `json:"product_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	VideoURL      string    `json:"video_url"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	ShowInGallery bool      `json:"show_in_gallery"`
	Autoplay      bool      `json:"autoplay"`
	ShowControls  bool      `json:"show_controls"`
	SortOrder     int       `json:"sort_order"`
	LazyLoad      bool      `json:"lazy_load"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// rand.Seed is deprecated in Go 1.20+, the global random generator is automatically seeded
// func init() {
//	rand.Seed(time.Now().UnixNano())
// }

// BeforeCreate GORM hook – تولید تصادفی SKU و اطمینان از عدم تکراری بودن
// منطق: حداکثر N بار تلاش برای ساخت SKU تصادفیِ یکتا؛ در صورت شکست، یک مقدار fallback می‌سازد
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	// اگر کاربر/سیستم از بیرون SKU ست کرده باشد، همان استفاده می‌شود
	if p.SKU != "" {
		return nil
	}

	// تلاش برای ساخت SKU تصادفی یکتا (۴ تا ۶ رقم)
	const maxAttempts = 20
	for i := 0; i < maxAttempts; i++ {
		candidate := strconv.Itoa(1000 + rand.Intn(999000))
		var count int64
		if err := tx.Model(&Product{}).Where("sku = ?", candidate).Count(&count).Error; err != nil {
			return err
		}
		if count == 0 {
			p.SKU = candidate
			return nil
		}
	}

	// در صورت بسیار نادرِ برخوردهای متوالی، یک مقدار fallback بساز
	// (استفاده از زمان برای کاهش احتمال تکرار؛ محدود به ۶ رقم)
	p.SKU = strconv.Itoa(int(time.Now().UnixNano() % 1_000_000))
	return nil
}

// AfterCreate پس از ایجاد محصول، رویداد همگام‌سازی جستجو را صف می‌کند.
func (p *Product) AfterCreate(tx *gorm.DB) error {
	return enqueueSearchEvent(tx, SearchResourceProduct, uint64(p.ID), SearchOpUpsert)
}

// AfterUpdate پس از به‌روزرسانی محصول، رویداد همگام‌سازی جستجو را صف می‌کند.
func (p *Product) AfterUpdate(tx *gorm.DB) error {
	return enqueueSearchEvent(tx, SearchResourceProduct, uint64(p.ID), SearchOpUpsert)
}

// AfterDelete پس از حذف محصول، حذف از ایندکس را صف می‌کند.
func (p *Product) AfterDelete(tx *gorm.DB) error {
	return enqueueSearchEvent(tx, SearchResourceProduct, uint64(p.ID), SearchOpDelete)
}
