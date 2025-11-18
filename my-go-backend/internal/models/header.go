package models

import (
	"time"
)

// HeaderLayer مدل لایه‌های هدر
type HeaderLayer struct {
	ID               uint    `json:"id" gorm:"primaryKey"`
	HeaderID         uint    `json:"header_id" gorm:"not null"`
	Header           Header  `json:"header" gorm:"foreignKey:HeaderID"`
	Name             string  `json:"name" gorm:"size:255;not null"`
	Width            int     `json:"width" gorm:"default:100"`
	Height           int     `json:"height" gorm:"default:50"`
	RowCount         int     `json:"row_count" gorm:"default:1"`
	Color            string  `json:"color" gorm:"size:50;default:'#ffffff'"`
	Opacity          float64 `json:"opacity" gorm:"default:1.0"`
	ShowSeparator    bool    `json:"showSeparator" gorm:"default:false"`
	SeparatorType    string  `json:"separatorType" gorm:"size:50;default:'line'"`
	SeparatorColor   string  `json:"separatorColor" gorm:"size:50;default:'#000000'"`
	SeparatorOpacity float64 `json:"separatorOpacity" gorm:"default:1.0"`
	SeparatorWidth   int     `json:"separatorWidth" gorm:"default:1"`

	EnableBorder   bool   `json:"enable_border" gorm:"column:enable_border;default:false"`
	BorderPosition string `json:"border_position" gorm:"column:border_position;size:50;default:'all'"`
	BorderColor    string `json:"border_color" gorm:"column:border_color;size:50;default:'#e5e7eb'"`
	BorderWidth    int    `json:"border_width" gorm:"column:border_width;default:1"`
	BorderStyle    string `json:"border_style" gorm:"column:border_style;size:50;default:'solid'"`

	EnableShadow    bool   `json:"enable_shadow" gorm:"column:enable_shadow;default:false"`
	ShadowIntensity string `json:"shadow_intensity" gorm:"column:shadow_intensity;size:50;default:'md'"`
	ShadowDirection string `json:"shadow_direction" gorm:"column:shadow_direction;size:50;default:'top'"`

	Direction        string `json:"direction" gorm:"column:direction;size:10;default:'rtl'"`
	MobileResponsive bool   `json:"mobile_responsive" gorm:"column:mobile_responsive;default:true"`
	TabletResponsive bool   `json:"tablet_responsive" gorm:"column:tablet_responsive;default:true"`

	Items string `json:"items" gorm:"type:text"` // JSON array as string
	// StyleSettings حاوی تنظیمات کامل UI مانند مرز و سایه است
	StyleSettings string `json:"style_settings" gorm:"column:style_settings;type:jsonb;default:'{}'"`
	// فیلدهای جدید برای پدینگ
	PaddingLeft   int       `json:"padding_left" gorm:"column:padding_left;default:20"`
	PaddingRight  int       `json:"padding_right" gorm:"column:padding_right;default:20"`
	PaddingTop    int       `json:"padding_top" gorm:"column:padding_top;default:10"`
	PaddingBottom int       `json:"padding_bottom" gorm:"column:padding_bottom;default:10"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Header مدل هدرهای سفارشی
type Header struct {
	ID            uint          `json:"id" gorm:"primaryKey"`
	Name          string        `json:"name" gorm:"size:255;not null;unique"`
	Description   string        `json:"description" gorm:"type:text"`
	PageSelection string        `json:"page_selection" gorm:"size:50;default:'all'"`
	SpecificPages string        `json:"specific_pages" gorm:"type:text"`
	ExcludedPages string        `json:"excluded_pages" gorm:"type:text"`
	Layers        []HeaderLayer `json:"layers" gorm:"foreignKey:HeaderID;constraint:OnDelete:CASCADE"`
	IsActive      bool          `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

// TableName نام جدول برای Header
func (Header) TableName() string {
	return "headers"
}

// TableName نام جدول برای HeaderLayer
func (HeaderLayer) TableName() string {
	return "header_layers"
}

// HeaderSettingsRequest مدل درخواست تنظیمات هدر
type HeaderSettingsRequest struct {
	LogoURL         string `json:"header_logo_url" binding:"omitempty"`
	LogoAlt         string `json:"header_logo_alt" binding:"omitempty"`
	ShowSearch      *bool  `json:"header_show_search"`
	ShowCart        *bool  `json:"header_show_cart"`
	ShowUserMenu    *bool  `json:"header_show_user_menu"`
	PhoneNumber     string `json:"header_phone_number" binding:"omitempty"`
	Email           string `json:"header_email" binding:"omitempty,email"`
	ShowSocialLinks *bool  `json:"header_show_social_links"`
	FacebookURL     string `json:"header_facebook_url" binding:"omitempty,url"`
	InstagramURL    string `json:"header_instagram_url" binding:"omitempty,url"`
	TelegramURL     string `json:"header_telegram_url" binding:"omitempty,url"`
	WhatsappURL     string `json:"header_whatsapp_url" binding:"omitempty,url"`
	Sticky          *bool  `json:"header_sticky"`
	BackgroundColor string `json:"header_background_color" binding:"omitempty"`
	TextColor       string `json:"header_text_color" binding:"omitempty"`
	Height          string `json:"header_height" binding:"omitempty"`
}

// HeaderSettingsResponse مدل پاسخ تنظیمات هدر
type HeaderSettingsResponse struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

// HeaderLogoUploadResponse مدل پاسخ آپلود لوگو
type HeaderLogoUploadResponse struct {
	Success bool `json:"success"`
	Data    struct {
		LogoURL string `json:"logo_url"`
	} `json:"data"`
	Message string `json:"message"`
}
