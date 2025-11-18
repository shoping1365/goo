package models

import (
	"time"

	"gorm.io/gorm"
)

// CategoryBrandPage represents a dedicated landing page for the combination of one category and one brand.
// It allows unique SEO metadata, banner, video and long description content.
type CategoryBrandPage struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	CategoryID uint   `json:"category_id" gorm:"not null;index"`
	BrandID    uint   `json:"brand_id" gorm:"not null;index"`
	Slug       string `json:"slug" gorm:"type:varchar(180);not null;uniqueIndex"`

	Title       string `json:"title" gorm:"type:varchar(150)"`
	Description string `json:"description" gorm:"type:text"`
	BannerURL   string `json:"banner_url" gorm:"type:varchar(255)"`
	VideoURL    string `json:"video_url" gorm:"type:varchar(255)"`

	MetaTitle       string `json:"meta_title" gorm:"type:varchar(150)"`
	MetaKeywords    string `json:"meta_keywords" gorm:"type:varchar(255)"`
	MetaDescription string `json:"meta_description" gorm:"type:varchar(255)"`

	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
	Brand    Brand    `json:"brand" gorm:"foreignKey:BrandID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
