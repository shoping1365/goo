package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Schema مدل اسکیمای پیش‌فرض برای مقالات و سایر محتواها
type Schema struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(255);not null;index"`
	Type        string `json:"type" gorm:"type:varchar(50);not null;index"` // Article, Product, Organization, etc.
	Description string `json:"description" gorm:"type:text"`
	IsTemplate  bool   `json:"is_template" gorm:"default:true;index"` // آیا این اسکیما تمپلیت است؟
	IsActive    bool   `json:"is_active" gorm:"default:true;index"`

	// فیلدهای عمومی اسکیما
	Title         string `json:"title" gorm:"type:varchar(255)"`
	Slug          string `json:"slug" gorm:"type:varchar(255)"`
	Excerpt       string `json:"excerpt" gorm:"type:text"`
	Content       string `json:"content" gorm:"type:text"`
	SiteName      string `json:"site_name" gorm:"type:varchar(255)"`
	FeaturedImage string `json:"featured_image" gorm:"type:varchar(500)"`
	ArticleID     string `json:"article_id" gorm:"type:varchar(100)"`
	ArticleURL    string `json:"article_url" gorm:"type:varchar(500)"`

	// فیلدهای مخصوص مقاله
	Author      string     `json:"author" gorm:"type:varchar(255)"`
	PublishedAt *time.Time `json:"published_at"`

	// فیلدهای مخصوص محصول
	Price    float64 `json:"price" gorm:"type:decimal(10,2)"`
	Currency string  `json:"currency" gorm:"type:varchar(10);default:'IRR'"`

	// فیلدهای مخصوص سازمان
	Address   string `json:"address" gorm:"type:varchar(500)"`
	Telephone string `json:"telephone" gorm:"type:varchar(50)"`

	// فیلدهای متا
	MetaTitle       string `json:"meta_title" gorm:"type:varchar(255)"`
	MetaKeywords    string `json:"meta_keywords" gorm:"type:varchar(255)"`
	MetaDescription string `json:"meta_description" gorm:"type:varchar(500)"`

	// فیلدهای Open Graph
	OgTitle       string `json:"og_title" gorm:"type:varchar(255)"`
	OgDescription string `json:"og_description" gorm:"type:varchar(500)"`
	OgImage       string `json:"og_image" gorm:"type:varchar(500)"`
	OgType        string `json:"og_type" gorm:"type:varchar(50)"`
	OgSiteName    string `json:"og_site_name" gorm:"type:varchar(100)"`

	// فیلدهای اضافی (JSON)
	ExtraFields datatypes.JSON `json:"extra_fields" gorm:"type:jsonb"`

	// فیلدهای سیستمی
	CreatedBy uint           `json:"created_by" gorm:"index"`
	UpdatedBy uint           `json:"updated_by" gorm:"index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// روابط
	Creator User `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	Updater User `json:"updater,omitempty" gorm:"foreignKey:UpdatedBy"`
}

// TableName نام جدول
func (Schema) TableName() string {
	return "schemas"
}

// SchemaTemplateRequest درخواست ایجاد تمپلیت اسکیما
type SchemaTemplateRequest struct {
	Name        string                 `json:"name" binding:"required"`
	Type        string                 `json:"type" binding:"required"`
	Description string                 `json:"description"`
	Fields      map[string]interface{} `json:"fields"`
}

// SchemaTemplateResponse پاسخ تمپلیت اسکیما
type SchemaTemplateResponse struct {
	ID          uint                   `json:"id"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Description string                 `json:"description"`
	Fields      map[string]interface{} `json:"fields"`
	IsActive    bool                   `json:"is_active"`
	CreatedAt   time.Time              `json:"created_at"`
}
