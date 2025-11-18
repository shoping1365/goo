package models

import (
	"time"

	"gorm.io/gorm"
)

// Brand represents a product brand in the catalogue.
// NOTE: The structure is aligned with the brands table created in migration 003.
// If you add new columns, remember to create a dedicated migration/ensure script.
//
// Fields:
//   - ID: Primary key
//   - Name: Brand name (unique)
//   - Slug: SEO-friendly slug (unique)
//   - Description: Optional brand description
//   - Published: If the brand is visible to customers
//   - ShowOnHome: Whether the brand should appear on home page widgets
//   - ShowInMenu: Whether the brand should be visible inside the main navigation menu
//   - MetaTitle/MetaKeywords/MetaDescription: SEO metadata
//   - ImageURL / VideoURL: Media assets associated with the brand
//   - CreatedAt / UpdatedAt / DeletedAt: Audit columns
//
// TODO: A display order column might be required by the front-end. For now it's
// omitted because the underlying table does not have this field. Add it later
// via an explicit migration once the DB schema is updated.

type Brand struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name" gorm:"type:varchar(100);not null;uniqueIndex"`
	Slug            string         `json:"slug" gorm:"type:varchar(120);not null;uniqueIndex"`
	OfficialName    string         `json:"official_name" gorm:"type:varchar(150)"`
	Description     string         `json:"description" gorm:"type:text"`
	Published       bool           `json:"published" gorm:"default:false;index"`
	ShowOnHome      bool           `json:"show_on_home" gorm:"default:false;index"`
	ShowInMenu      bool           `json:"show_in_menu" gorm:"default:false;index"`
	MetaTitle       string         `json:"meta_title" gorm:"type:varchar(150)"`
	MetaKeywords    string         `json:"meta_keywords" gorm:"type:varchar(255)"`
	MetaDescription string         `json:"meta_description" gorm:"type:varchar(255)"`
	ImageURL        string         `json:"image_url" gorm:"type:varchar(255)"`
	VideoURL        string         `json:"video_url" gorm:"type:varchar(255)"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
