package models

import (
	"time"

	"gorm.io/gorm"
)

// AttributeGroup represents a collection of related attributes that belong to a specific product category.
// It is mapped to the attribute_groups table which is created in migration 009.
//
// Fields:
//   - ID:          Primary key
//   - Name:        Human-readable name of the group (e.g., "مشخصات فنی گوشی")
//   - CategoryID:  FK to categories table – each group belongs to exactly one category
//   - Description: Optional description shown to admins
//   - CreatedAt/UpdatedAt/DeletedAt: Audit columns managed by GORM
//
// Relations:
//   - Category:   Lazy relationship to Category struct
//   - Attributes: Many-to-many through AttributeGroupAttribute join model
//
// NOTE: Make sure migrations 009 & 010 are executed before using this model.

type AttributeGroup struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"type:varchar(100);not null;index"`
	CategoryID  uint           `json:"category_id" gorm:"not null;index"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	Category   Category                  `json:"category" gorm:"foreignKey:CategoryID"`
	Attributes []AttributeGroupAttribute `json:"attributes" gorm:"foreignKey:AttributeGroupID"`
}

// AttributeGroupAttribute is the join entity between AttributeGroup and Attribute with
// additional metadata (sort order, control type, etc.). It maps to the table created in
// migration 010.

type AttributeGroupAttribute struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	AttributeGroupID uint           `json:"attribute_group_id" gorm:"not null;index"`
	AttributeID      uint           `json:"attribute_id" gorm:"not null;index"`
	SortOrder        uint           `json:"sort_order" gorm:"default:1;index"`
	Type             string         `json:"type" gorm:"type:varchar(30);not null"`         // select | custom_text | multi_select
	ControlType      string         `json:"control_type" gorm:"type:varchar(30);not null"` // dropdown_single | textbox_single | ...
	HasFilter        bool           `json:"has_filter" gorm:"default:false;index"`
	IsKey            bool           `json:"is_key" gorm:"default:false;index"`
	ShowOnProduct    bool           `json:"show_on_product" gorm:"default:true;index"`
	IsRequired       bool           `json:"is_required" gorm:"default:false;index"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	Attribute      Attribute      `json:"attribute" gorm:"foreignKey:AttributeID"`
	AttributeGroup AttributeGroup `json:"-" gorm:"foreignKey:AttributeGroupID"`
}
