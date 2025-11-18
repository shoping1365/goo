package models

import (
	"time"
)

// MobileAppHeader مدل هدرهای موبایل و اپلیکیشن
type MobileAppHeader struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	Name              string    `json:"name" gorm:"size:255;not null;unique"`
	Description       string    `json:"description" gorm:"type:text"`
	Platform          string    `json:"platform" gorm:"size:50;default:'mobile'"` // mobile, app, both
	PageSelection     string    `json:"page_selection" gorm:"size:50;default:'all'"`
	SpecificPages     string    `json:"specific_pages" gorm:"type:text"`
	ExcludedPages     string    `json:"excluded_pages" gorm:"type:text"`
	HeaderType        string    `json:"header_type" gorm:"size:50;default:'default'"` // default, minimal, full
	LogoURL           string    `json:"logo_url" gorm:"size:500"`
	LogoAlt           string    `json:"logo_alt" gorm:"size:255"`
	ShowSearch        bool      `json:"show_search" gorm:"default:true"`
	ShowCart          bool      `json:"show_cart" gorm:"default:true"`
	ShowUserMenu      bool      `json:"show_user_menu" gorm:"default:true"`
	ShowNotifications bool      `json:"show_notifications" gorm:"default:true"`
	ShowMenuButton    bool      `json:"show_menu_button" gorm:"default:true"`
	BackgroundColor   string    `json:"background_color" gorm:"size:50;default:'#ffffff'"`
	TextColor         string    `json:"text_color" gorm:"size:50;default:'#000000'"`
	TopImageURL       string    `json:"top_image_url" gorm:"size:500"`
	TopImageAlt       string    `json:"top_image_alt" gorm:"size:255"`
	BottomImageURL    string    `json:"bottom_image_url" gorm:"size:500"`
	BottomImageAlt    string    `json:"bottom_image_alt" gorm:"size:255"`
	IsActive          bool      `json:"is_active" gorm:"default:true"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// TableName نام جدول برای MobileAppHeader
func (MobileAppHeader) TableName() string {
	return "mobile_app_headers"
}

// MobileAppHeaderSettingsRequest مدل درخواست تنظیمات هدر موبایل و اپلیکیشن
type MobileAppHeaderSettingsRequest struct {
	LogoURL         string `json:"mobile_app_header_logo_url" binding:"omitempty"`
	LogoAlt         string `json:"mobile_app_header_logo_alt" binding:"omitempty"`
	ShowSearch      *bool  `json:"mobile_app_header_show_search"`
	ShowCart        *bool  `json:"mobile_app_header_show_cart"`
	ShowUserMenu    *bool  `json:"mobile_app_header_show_user_menu"`
	PhoneNumber     string `json:"mobile_app_header_phone_number" binding:"omitempty"`
	Email           string `json:"mobile_app_header_email" binding:"omitempty,email"`
	ShowSocialLinks *bool  `json:"mobile_app_header_show_social_links"`
	FacebookURL     string `json:"mobile_app_header_facebook_url" binding:"omitempty,url"`
	InstagramURL    string `json:"mobile_app_header_instagram_url" binding:"omitempty,url"`
	TelegramURL     string `json:"mobile_app_header_telegram_url" binding:"omitempty,url"`
	WhatsappURL     string `json:"mobile_app_header_whatsapp_url" binding:"omitempty,url"`
	Sticky          *bool  `json:"mobile_app_header_sticky"`
	BackgroundColor string `json:"mobile_app_header_background_color" binding:"omitempty"`
	TextColor       string `json:"mobile_app_header_text_color" binding:"omitempty"`
	Height          string `json:"mobile_app_header_height" binding:"omitempty"`
	Platform        string `json:"mobile_app_header_platform" binding:"omitempty"`
}

// MobileAppHeaderSettingsResponse مدل پاسخ تنظیمات هدر موبایل و اپلیکیشن
type MobileAppHeaderSettingsResponse struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

// MobileAppHeaderLogoUploadResponse مدل پاسخ آپلود لوگو هدر موبایل و اپلیکیشن
type MobileAppHeaderLogoUploadResponse struct {
	Success bool `json:"success"`
	Data    struct {
		LogoURL string `json:"logo_url"`
	} `json:"data"`
	Message string `json:"message"`
}
