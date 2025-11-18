package models

import (
	"time"

	"gorm.io/gorm"
)

// Attribute represents a product specification (feature) that can be assigned to products.
// The structure matches columns created in migration 011 (attributes table) and is extended
// to include a few extra fields required by the new admin UI such as IsRequired, IsFilterable
// and SortOrder. Make sure corresponding migrations are executed before using these columns.
//
// Fields:
//   - ID: Primary key
//   - Name: Attribute name (unique)
//   - DisplayName: Title that will be shown on the storefront
//   - Code: Unique code useful for integrations
//   - DataType: The data type of the value (text, number, boolean, date, ...)
//   - Unit: Unit of measurement if the value is numeric (cm, gr, V, ...)
//   - SortOrder: Display order in lists (optional)
//   - IsRequired: Whether the attribute is mandatory for product creation
//   - IsFilterable: Whether the attribute can be used inside product list filters
//   - IsActive: Shows/hides the attribute completely
//   - HasOptions: Indicates whether the attribute has options
//   - CreatedAt / UpdatedAt / DeletedAt: Audit columns
//
// NOTE: Attribute has a one-to-many relation with AttributeValue.

type Attribute struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"type:varchar(100);not null;uniqueIndex"`
	DisplayName  string `json:"display_name" gorm:"type:varchar(150)"`
	Code         string `json:"code" gorm:"type:varchar(50);uniqueIndex"`
	DataType     string `json:"data_type" gorm:"type:varchar(20);not null"`
	Unit         string `json:"unit" gorm:"type:varchar(30)"`
	SortOrder    uint   `json:"sort_order" gorm:"default:0;index"`
	IsRequired   bool   `json:"is_required" gorm:"default:false"`
	IsFilterable bool   `json:"is_filterable" gorm:"default:false"`
	IsActive     bool   `json:"is_active" gorm:"default:true;index"`
	HasOptions   bool   `json:"has_options" gorm:"default:false;index"`
	// در برخی طرح‌های قدیمی DB، ستون attribute_group_id به‌صورت NOT NULL تعریف شده بود.
	// برای سازگاری، این فیلد را اختیاری تعریف می‌کنیم تا ایجاد ویژگی بدون گروه ممکن باشد.
	AttributeGroupID *uint          `json:"attribute_group_id,omitempty" gorm:"index"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	Values []AttributeValue `json:"values,omitempty" gorm:"foreignKey:AttributeID"`
}

// AttributeValue represents a possible option/value for an attribute (e.g. Red, XL, 128GB).
//
// NOTE: The UI currently supports an optional colour for a value; this can be stored inside
// the Meta field in JSON format (e.g. {"color":"#ff0000"}).

type AttributeValue struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	AttributeID  uint           `json:"attribute_id" gorm:"not null;index"`
	Value        string         `json:"value" gorm:"type:varchar(100);not null;index"`
	TextValue    *string        `json:"text_value,omitempty" gorm:"type:varchar(255)"`
	NumericValue *float64       `json:"numeric_value,omitempty"`
	BoolValue    *bool          `json:"bool_value,omitempty"`
	DateValue    *time.Time     `json:"date_value,omitempty"`
	Slug         string         `json:"slug" gorm:"type:varchar(100);index"`
	SortOrder    uint           `json:"sort_order" gorm:"default:0;index"`
	Meta         string         `json:"meta" gorm:"type:jsonb"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relations
	Attribute Attribute `json:"-" gorm:"foreignKey:AttributeID"`
}
