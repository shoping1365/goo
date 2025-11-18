package models

import (
	"time"
)

// FooterLayer مدل لایه‌های فوتر
type FooterLayer struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	FooterID uint    `json:"footer_id" gorm:"not null"`
	Footer   Footer  `json:"footer" gorm:"foreignKey:FooterID"`
	Name     string  `json:"name" gorm:"size:255;not null"`
	Width    int     `json:"width" gorm:"default:100"`
	Height   int     `json:"height" gorm:"default:200"`
	RowCount int     `json:"row_count" gorm:"default:1"`
	Color    string  `json:"color" gorm:"size:50;default:'#1f2937'"`
	Opacity  float64 `json:"opacity" gorm:"default:1.0"`

	// Separator settings (قدیمی - برای سازگاری با داده‌های موجود)
	ShowSeparator    bool    `json:"showSeparator" gorm:"default:false"`
	SeparatorType    string  `json:"separatorType" gorm:"size:50;default:'solid'"`
	SeparatorColor   string  `json:"separatorColor" gorm:"size:50;default:'#ffffff'"`
	SeparatorOpacity float64 `json:"separatorOpacity" gorm:"default:0.2"`
	SeparatorWidth   int     `json:"separatorWidth" gorm:"default:1"`

	Items string `json:"items" gorm:"type:text"` // JSON array as string

	// تنظیمات استایل شامل: border, shadow, و سایر تنظیمات UI
	// ساختار JSON: {"border": {...}, "shadow": {...}, "layout": {...}}
	StyleSettings string `json:"styleSettings" gorm:"type:jsonb;default:'{}'"`

	// فیلدهای جدید برای پدینگ
	PaddingLeft   int       `json:"padding_left" gorm:"default:20"`
	PaddingRight  int       `json:"padding_right" gorm:"default:20"`
	PaddingTop    int       `json:"padding_top" gorm:"default:20"`
	PaddingBottom int       `json:"padding_bottom" gorm:"default:20"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Footer مدل فوترهای سفارشی
type Footer struct {
	ID            uint          `json:"id" gorm:"primaryKey"`
	Name          string        `json:"name" gorm:"size:255;not null;unique"`
	Description   string        `json:"description" gorm:"type:text"`
	PageSelection string        `json:"page_selection" gorm:"size:50;default:'all'"`
	SpecificPages string        `json:"specific_pages" gorm:"type:text"`
	ExcludedPages string        `json:"excluded_pages" gorm:"type:text"`
	Layers        []FooterLayer `json:"layers" gorm:"foreignKey:FooterID;constraint:OnDelete:CASCADE"`
	IsActive      bool          `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

// TableName نام جدول برای Footer
func (Footer) TableName() string {
	return "footers"
}

// TableName نام جدول برای FooterLayer
func (FooterLayer) TableName() string {
	return "footer_layers"
}

// FooterSettingsRequest مدل درخواست تنظیمات فوتر
type FooterSettingsRequest struct {
	LogoURL         string `json:"footer_logo_url" binding:"omitempty"`
	LogoAlt         string `json:"footer_logo_alt" binding:"omitempty"`
	ShowCopyright   *bool  `json:"footer_show_copyright"`
	ShowSocialLinks *bool  `json:"footer_show_social_links"`
	PhoneNumber     string `json:"footer_phone_number" binding:"omitempty"`
	Email           string `json:"footer_email" binding:"omitempty,email"`
	Address         string `json:"footer_address" binding:"omitempty"`
	FacebookURL     string `json:"footer_facebook_url" binding:"omitempty,url"`
	InstagramURL    string `json:"footer_instagram_url" binding:"omitempty,url"`
	TelegramURL     string `json:"footer_telegram_url" binding:"omitempty,url"`
	WhatsappURL     string `json:"footer_whatsapp_url" binding:"omitempty,url"`
	BackgroundColor string `json:"footer_background_color" binding:"omitempty"`
	TextColor       string `json:"footer_text_color" binding:"omitempty"`
	Height          string `json:"footer_height" binding:"omitempty"`
}

// FooterSettingsResponse مدل پاسخ تنظیمات فوتر
type FooterSettingsResponse struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

// FooterLogoUploadResponse مدل پاسخ آپلود لوگو
type FooterLogoUploadResponse struct {
	Success bool `json:"success"`
	Data    struct {
		LogoURL string `json:"logo_url"`
	} `json:"data"`
	Message string `json:"message"`
}
