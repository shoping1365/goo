package models

import (
	"time"

	"gorm.io/gorm"
)

// DiscountNotification مدل اعلان‌های تخفیف محصولات
// این مدل برای ثبت درخواست‌های «خبرم کن در زمان تخفیف» استفاده می‌شود.
type DiscountNotification struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	UserID         uint           `json:"user_id" gorm:"not null;index"`
	ProductID      uint           `json:"product_id" gorm:"not null;index"`
	Email          string         `json:"email" gorm:"type:varchar(255);not null"`
	Phone          string         `json:"phone" gorm:"type:varchar(20)"`
	ThresholdType  string         `json:"threshold_type" gorm:"type:varchar(20);default:'any';index"` // any | percent | amount
	ThresholdValue float64        `json:"threshold_value" gorm:"type:numeric(12,2);default:0"`        // درصد یا مبلغ
	Status         string         `json:"status" gorm:"type:varchar(20);default:'pending';index"`     // pending | sent | cancelled
	SentAt         *time.Time     `json:"sent_at"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// روابط
	User    User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (DiscountNotification) TableName() string { return "discount_notifications" }
