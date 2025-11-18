package models

import (
	"time"
)

// MobileAppFooterLayer مدل لایه‌های فوتر موبایل و اپلیکیشن
type MobileAppFooterLayer struct {
	ID                uint            `json:"id" gorm:"primaryKey"`
	MobileAppFooterID uint            `json:"mobile_app_footer_id" gorm:"not null"`
	MobileAppFooter   MobileAppFooter `json:"mobile_app_footer" gorm:"foreignKey:MobileAppFooterID"`
	Name              string          `json:"name" gorm:"size:255;not null"`
	Width             int             `json:"width" gorm:"default:100"`
	Height            int             `json:"height" gorm:"default:200"`
	RowCount          int             `json:"row_count" gorm:"default:1"`
	Color             string          `json:"color" gorm:"size:50;default:'#1f2937'"`
	Opacity           float64         `json:"opacity" gorm:"default:1.0"`
	ShowSeparator     bool            `json:"show_separator" gorm:"default:false"`
	SeparatorType     string          `json:"separator_type" gorm:"size:50;default:'line'"`
	SeparatorColor    string          `json:"separator_color" gorm:"size:50;default:'#ffffff'"`
	SeparatorOpacity  float64         `json:"separator_opacity" gorm:"default:0.2"`
	SeparatorWidth    int             `json:"separator_width" gorm:"default:1"`
	Items             string          `json:"items" gorm:"type:text"` // JSON array as string
	// فیلدهای جدید برای پدینگ
	PaddingLeft   int       `json:"padding_left" gorm:"default:20"`
	PaddingRight  int       `json:"padding_right" gorm:"default:20"`
	PaddingTop    int       `json:"padding_top" gorm:"default:20"`
	PaddingBottom int       `json:"padding_bottom" gorm:"default:20"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// MobileAppFooter مدل فوترهای موبایل و اپلیکیشن
type MobileAppFooter struct {
	ID            uint                   `json:"id" gorm:"primaryKey"`
	Name          string                 `json:"name" gorm:"size:255;not null;unique"`
	Description   string                 `json:"description" gorm:"type:text"`
	Platform      string                 `json:"platform" gorm:"size:50;default:'mobile'"` // mobile, app, both
	PageSelection string                 `json:"page_selection" gorm:"size:50;default:'all'"`
	SpecificPages string                 `json:"specific_pages" gorm:"type:text"`
	ExcludedPages string                 `json:"excluded_pages" gorm:"type:text"`
	Layers        []MobileAppFooterLayer `json:"layers" gorm:"foreignKey:MobileAppFooterID;constraint:OnDelete:CASCADE"`
	IsActive      bool                   `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

// TableName نام جدول برای MobileAppFooter
func (MobileAppFooter) TableName() string {
	return "mobile_app_footers"
}

// TableName نام جدول برای MobileAppFooterLayer
func (MobileAppFooterLayer) TableName() string {
	return "mobile_app_footer_layers"
}

// MobileAppFooterSettingsRequest مدل درخواست تنظیمات فوتر موبایل و اپلیکیشن
type MobileAppFooterSettingsRequest struct {
	LogoURL         string `json:"mobile_app_footer_logo_url" binding:"omitempty"`
	LogoAlt         string `json:"mobile_app_footer_logo_alt" binding:"omitempty"`
	ShowCopyright   *bool  `json:"mobile_app_footer_show_copyright"`
	ShowSocialLinks *bool  `json:"mobile_app_footer_show_social_links"`
	PhoneNumber     string `json:"mobile_app_footer_phone_number" binding:"omitempty"`
	Email           string `json:"mobile_app_footer_email" binding:"omitempty,email"`
	Address         string `json:"mobile_app_footer_address" binding:"omitempty"`
	FacebookURL     string `json:"mobile_app_footer_facebook_url" binding:"omitempty,url"`
	InstagramURL    string `json:"mobile_app_footer_instagram_url" binding:"omitempty,url"`
	TelegramURL     string `json:"mobile_app_footer_telegram_url" binding:"omitempty,url"`
	WhatsappURL     string `json:"mobile_app_footer_whatsapp_url" binding:"omitempty,url"`
	BackgroundColor string `json:"mobile_app_footer_background_color" binding:"omitempty"`
	TextColor       string `json:"mobile_app_footer_text_color" binding:"omitempty"`
	Height          string `json:"mobile_app_footer_height" binding:"omitempty"`
	Platform        string `json:"mobile_app_footer_platform" binding:"omitempty"`
}

// MobileAppFooterSettingsResponse مدل پاسخ تنظیمات فوتر موبایل و اپلیکیشن
type MobileAppFooterSettingsResponse struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

// MobileAppFooterLogoUploadResponse مدل پاسخ آپلود لوگو فوتر موبایل و اپلیکیشن
type MobileAppFooterLogoUploadResponse struct {
	Success bool `json:"success"`
	Data    struct {
		LogoURL string `json:"logo_url"`
	} `json:"data"`
	Message string `json:"message"`
}
