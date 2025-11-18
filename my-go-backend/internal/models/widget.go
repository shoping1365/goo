package models

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
)

// Widget مدل ابزارک - هر ابزارک یک نمونه منحصر به فرد با محتوای اختصاصی است
type Widget struct {
	ID           uint           `json:"id" gorm:"primaryKey"`                                  // کلید اصلی
	Code         string         `json:"code" gorm:"type:varchar(100);uniqueIndex;not null"`    // کد یونیک برای استفاده در صفحات
	Title        string         `json:"title" gorm:"type:varchar(255);not null"`               // عنوان ابزارک
	Description  string         `json:"description" gorm:"type:text"`                          // توضیحات ابزارک
	Type         string         `json:"type" gorm:"type:varchar(50);not null;index"`           // نوع ابزارک (slider, banner, product-grid, ...)
	Status       string         `json:"status" gorm:"type:varchar(20);default:'active';index"` // وضعیت (active, inactive, draft)
	SortOrder    int            `json:"sort_order" gorm:"column:sort_order;default:1;index"`   // ترتیب نمایش
	Image        string         `json:"image" gorm:"type:varchar(500)"`                        // تصویر پیش‌نمایش ابزارک
	Link         string         `json:"link" gorm:"type:varchar(500)"`                         // لینک ابزارک (در صورت نیاز)
	Page         string         `json:"page" gorm:"type:varchar(50);default:'home';index"`     // صفحه‌ای که ابزارک در آن نمایش داده می‌شود
	ShowOnMobile bool           `json:"show_on_mobile" gorm:"default:true"`                    // نمایش در موبایل
	Config       datatypes.JSON `json:"config" gorm:"type:jsonb"`                              // تنظیمات و محتوای اختصاصی ابزارک (JSON)
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

// WidgetConfig ساختار پایه برای تنظیمات ابزارک
type WidgetConfig struct {
	// تنظیمات عمومی
	BackgroundColor string `json:"background_color,omitempty"`
	Padding         string `json:"padding,omitempty"`
	Margin          string `json:"margin,omitempty"`
	CustomClass     string `json:"custom_class,omitempty"`
}

// SliderConfig تنظیمات اسلایدر
type SliderConfig struct {
	WidgetConfig
	SliderCount    int          `json:"slider_count"`
	WideBg         bool         `json:"wide_bg"`
	BgColor        string       `json:"bg_color"`
	BannerPosition string       `json:"banner_position"` // left, right
	DisplayOrder   string       `json:"display_order"`   // asc, desc
	Height         int          `json:"height"`
	Slides         []SlideItem  `json:"slides"`
	SideBanners    []BannerItem `json:"side_banners"`
}

// BannerConfig تنظیمات بنر
type BannerConfig struct {
	WidgetConfig
	Layout  string       `json:"layout"` // single, double, triple, quad
	Banners []BannerItem `json:"banners"`
}

// ProductConfig تنظیمات محصولات
type ProductConfig struct {
	WidgetConfig
	CategoryID   uint   `json:"category_id,omitempty"`
	ProductCount int    `json:"product_count"`
	Layout       string `json:"layout"` // grid, list, carousel
	ShowPrice    bool   `json:"show_price"`
	ShowRating   bool   `json:"show_rating"`
	ShowDiscount bool   `json:"show_discount"`
}

// SlideItem آیتم اسلایدر
type SlideItem struct {
	ID          uint   `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image"`
	Link        string `json:"link,omitempty"`
	Order       int    `json:"order"`
	Status      string `json:"status"` // active, inactive
}

// BannerItem آیتم بنر
type BannerItem struct {
	ID          uint   `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image"`
	Link        string `json:"link,omitempty"`
	Order       int    `json:"order"`
	Status      string `json:"status"` // active, inactive
}

// ProductCarouselConfig تنظیمات کاروسل محصولات
type ProductCarouselConfig struct {
	WidgetConfig
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	CategoryID    *uint  `json:"category_id,omitempty"`
	ProductCount  int    `json:"product_count"`
	SlidesPerView int    `json:"slides_per_view"`
	ShowPrice     bool   `json:"show_price"`
	ShowRating    bool   `json:"show_rating"`
	ShowDiscount  bool   `json:"show_discount"`
	Autoplay      bool   `json:"autoplay"`
	AutoplayDelay int    `json:"autoplay_delay"`
	WideBg        bool   `json:"wide_bg"`
	BgColor       string `json:"bg_color"`
}

// CategoryConfig تنظیمات دسته‌بندی
type CategoryConfig struct {
	WidgetConfig
	Categories       []CategoryItem `json:"categories"`
	Layout           string         `json:"layout"` // grid, list
	ShowProductCount bool           `json:"show_product_count"`
	Columns          int            `json:"columns"` // تعداد ستون‌ها
	ShowTitle        bool           `json:"show_title"`
	Title            string         `json:"title"`
	TitleColor       string         `json:"title_color"`
	CardStyle        string         `json:"card_style"`     // default, rounded, shadow
	ImageSize        string         `json:"image_size"`     // small, medium, large
	TextAlignment    string         `json:"text_alignment"` // left, center, right
}

// CategoryItem آیتم دسته‌بندی
type CategoryItem struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	Image        string `json:"image"`
	Link         string `json:"link"`
	ProductCount int    `json:"product_count,omitempty"`
}

// تابع تولید کد یونیک
func (w *Widget) GenerateUniqueCode() string {
	typeMap := map[string]string{
		"main-slider-side-banner": "SLIDER_MAIN",
		"single-slider-side":      "SLIDER_SINGLE",
		"full-slider":             "SLIDER_FULL",

		"products-box-bg": "PRODUCT_BG",
		"full-banner":     "BANNER_FULL",

		"double-banner":    "BANNER_DOUBLE",
		"triple-banner":    "BANNER_TRIPLE",
		"quad-banner":      "BANNER_QUAD",
		"penta-banner":     "BANNER_PENTA",
		"brands-slider":    "BRANDS_SLIDER",
		"services-slider":  "SERVICES_SLIDER",
		"category-box":     "CATEGORY_BOX",
		"product-carousel": "PRODUCT_CAROUSEL",
		"blog-posts":       "BLOG_POSTS",
	}

	prefix := typeMap[w.Type]
	if prefix == "" {
		prefix = "WIDGET"
	}

	// در اینجا باید شماره یونیک تولید کنید (می‌تواند از database count استفاده کند)
	return prefix + "_001" // برای الان ثابت، بعداً dynamic می‌شود
}

// تابع validate برای مدل
func (w *Widget) Validate() error {
	if w.Title == "" {
		return fmt.Errorf("عنوان ابزارک الزامی است")
	}
	if w.Type == "" {
		return fmt.Errorf("نوع ابزارک الزامی است")
	}
	validTypes := []string{
		"main-slider-side-banner", "single-slider-side", "full-slider",
		"products-box-bg", "full-banner",
		"double-banner", "triple-banner", "quad-banner", "penta-banner",
		"brands-slider", "services-slider", "category-box", "blog-posts",
		"product-carousel",
	}

	isValidType := false
	for _, validType := range validTypes {
		if w.Type == validType {
			isValidType = true
			break
		}
	}

	if !isValidType {
		return fmt.Errorf("نوع ابزارک نامعتبر است")
	}

	return nil
}
