package models

import (
	"time"

	"gorm.io/gorm"
)

// Category represents a product category in the catalogue
// If you need nesting, ParentID points to the parent category (nullable).
type Category struct {
	ID              uint    `json:"id" gorm:"primaryKey"`
	Name            string  `json:"name" gorm:"type:varchar(100);not null;uniqueIndex"`
	NameEn          string  `json:"name_en" gorm:"type:varchar(150)"`
	Slug            string  `json:"slug" gorm:"type:varchar(120);not null;uniqueIndex"`
	ParentID        *uint   `json:"parent_id" gorm:"index"`
	ParentSlug      *string `json:"parent_slug" gorm:"type:varchar(120);index"` // فیلد جدید برای ذخیره slug والد
	Description     string  `json:"description" gorm:"type:text"`
	ImageURL        string  `json:"image_url" gorm:"type:varchar(255)"`
	Icon            string  `json:"icon" gorm:"type:varchar(100)"`
	Order           uint    `json:"order" gorm:"default:1;index"`
	Published       bool    `json:"published" gorm:"default:true;index"`
	ShowOnHome      bool    `json:"show_on_home" gorm:"default:false;index"`
	ShowInMenu      bool    `json:"show_in_menu" gorm:"default:false;index"`
	BannerURL       string  `json:"banner_url" gorm:"type:varchar(255)"`
	NoticeMessage   string  `json:"notice_message" gorm:"type:text"`
	MetaTitle       string  `json:"meta_title" gorm:"type:varchar(150)"`
	MetaKeywords    string  `json:"meta_keywords" gorm:"type:varchar(255)"`
	MetaDescription string  `json:"meta_description" gorm:"type:varchar(255)"`

	// Field to hold the count of products, not stored in DB
	ProductCount int64 `json:"product_count" gorm:"-"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
