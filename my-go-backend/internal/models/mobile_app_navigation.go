package models

import (
	"time"

	"gorm.io/gorm"
)

// MobileAppNavigation مدل ناوبری موبایل و اپلیکیشن
type MobileAppNavigation struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name" gorm:"size:255;not null" binding:"required"`
	Description     string         `json:"description" gorm:"type:text"`
	Platform        string         `json:"platform" gorm:"size:50;not null" binding:"required"` // mobile, app, both
	BackgroundColor string         `json:"background_color" gorm:"size:50;default:'#f8fafc'"`    // رنگ پس‌زمینه
	TextColor       string         `json:"text_color" gorm:"size:50;default:'#000000'"`         // رنگ متن
	NavigationItems string         `json:"navigation_items" gorm:"type:text"`                   // JSON array of navigation items
	PageSelection   string         `json:"page_selection" gorm:"size:50;not null"`               // all, specific, excluded
	SpecificPages   string         `json:"specific_pages" gorm:"type:text"`                     // JSON array of page paths
	ExcludedPages   string         `json:"excluded_pages" gorm:"type:text"`                     // JSON array of page paths
	IsActive        bool           `json:"is_active" gorm:"default:true"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// روابط
	Layers []MobileAppNavigationLayer `json:"layers" gorm:"foreignKey:MobileAppNavigationID;constraint:OnDelete:CASCADE"`
}

// MobileAppNavigationLayer مدل لایه‌های ناوبری موبایل و اپلیکیشن
type MobileAppNavigationLayer struct {
	ID                    uint      `json:"id" gorm:"primaryKey"`
	MobileAppNavigationID uint      `json:"mobile_app_navigation_id" gorm:"not null"`
	Name                  string    `json:"name" gorm:"size:255;not null"`
	Type                  string    `json:"type" gorm:"size:50;not null"` // menu, button, link, etc.
	Position              string    `json:"position" gorm:"size:50"`      // top, bottom, left, right, center
	Items                 string    `json:"items" gorm:"type:text"`       // JSON array of navigation items
	Style                 string    `json:"style" gorm:"type:text"`       // JSON object for styling
	IsActive              bool      `json:"is_active" gorm:"default:true"`
	SortOrder             int       `json:"sort_order" gorm:"default:0"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`

	// روابط
	MobileAppNavigation MobileAppNavigation `json:"mobile_app_navigation" gorm:"foreignKey:MobileAppNavigationID"`
}

// MobileAppNavigationItem ساختار آیتم ناوبری
type MobileAppNavigationItem struct {
	ID        string                    `json:"id"`
	Label     string                    `json:"label"`
	URL       string                    `json:"url"`
	Icon      string                    `json:"icon"`
	Target    string                    `json:"target"` // _self, _blank
	IsActive  bool                      `json:"is_active"`
	SortOrder int                       `json:"sort_order"`
	Children  []MobileAppNavigationItem `json:"children,omitempty"`
}

// MobileAppNavigationLogoUploadResponse پاسخ آپلود لوگو ناوبری
type MobileAppNavigationLogoUploadResponse struct {
	Success bool `json:"success"`
	Data    struct {
		LogoURL string `json:"logo_url"`
	} `json:"data"`
	Message string `json:"message"`
}
