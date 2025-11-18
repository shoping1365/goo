package models

import (
	"time"

	"gorm.io/gorm"
)

// ProductQA مدل پرسش و پاسخ محصولات
type ProductQA struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ProductID   uint           `json:"product_id" gorm:"not null;index"`
	UserID      uint           `json:"user_id" gorm:"not null;index"`
	ParentID    *uint          `json:"parent_id" gorm:"index"` // برای پاسخ‌ها
	Question    string         `json:"question" gorm:"type:text;not null"`
	Answer      string         `json:"answer" gorm:"type:text"`
	Status      string         `json:"status" gorm:"type:varchar(20);default:'pending';index"` // pending, approved, rejected
	IsAnonymous bool           `json:"is_anonymous" gorm:"default:false"`
	HelpfulCount int           `json:"helpful_count" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// روابط
	Product Product     `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	User    User        `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Parent  *ProductQA  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies []ProductQA `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
}

// TableName نام جدول
func (ProductQA) TableName() string {
	return "product_qa"
}
