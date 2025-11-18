package models

import (
	"time"

	"gorm.io/gorm"
)

// StockNotification مدل اعلان‌های موجودی محصولات
type StockNotification struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	UserID       uint           `json:"user_id" gorm:"not null;index"`
	ProductID    uint           `json:"product_id" gorm:"not null;index"`
	Email        string         `json:"email" gorm:"type:varchar(255);not null"`
	Phone        string         `json:"phone" gorm:"type:varchar(20)"`
	Status       string         `json:"status" gorm:"type:varchar(20);default:'pending';index"` // pending, sent, cancelled
	SentAt       *time.Time     `json:"sent_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// روابط
	User    User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

// TableName نام جدول
func (StockNotification) TableName() string {
	return "stock_notifications"
}
