package models

import (
	"time"

	"gorm.io/gorm"
)

// PostCategory مدل دسته‌بندی نوشته‌ها
type PostCategory struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:100;not null;uniqueIndex"`
	Slug        string         `json:"slug" gorm:"size:120;not null;uniqueIndex"`
	Description string         `json:"description" gorm:"type:text"`
	IsActive    bool           `json:"is_active" gorm:"default:true;index"`
	Order       uint           `json:"order" gorm:"default:1;index"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// روابط
	Posts []Post `json:"posts,omitempty" gorm:"foreignKey:CategoryID"`
}

// TableName نام جدول
func (PostCategory) TableName() string {
	return "post_categories"
}
