package models

import (
	"time"

	"gorm.io/gorm"
)

// ProductAttributeValue maps to product_attribute_values table (enhanced schema).
type ProductAttributeValue struct {
	ID               uint  `json:"id" gorm:"primaryKey"`
	ProductID        uint  `json:"product_id" gorm:"not null;index:idx_product_attribute_option,priority:1"`
	AttributeID      uint  `json:"attribute_id" gorm:"not null;index:idx_product_attribute_option,priority:2"`
	AttributeValueID *uint `json:"attribute_value_id" gorm:"column:attribute_value_id;index:idx_product_attribute_option,priority:3"`

	ValueText    *string    `json:"value_text" gorm:"type:text"`
	ValueNumeric *float64   `json:"value_numeric"`
	ValueBool    *bool      `json:"value_bool"`
	ValueDate    *time.Time `json:"value_date"`

	SortOrder uint           `json:"sort_order" gorm:"default:0;index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
