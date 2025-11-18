package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name        string    `json:"name"`
	Email       string    `json:"email" gorm:"unique"`
	Phone       string    `json:"phone"`
	JoinDate    time.Time `json:"join_date"`
	Reviews     []Review  `json:"reviews"`
	ReviewCount int       `json:"review_count" gorm:"-"`
}

type Review struct {
	gorm.Model
	CustomerID   uint           `json:"customer_id"`
	Customer     Customer       `json:"customer" gorm:"foreignKey:CustomerID"`
	ProductID    uint           `json:"product_id"`
	Product      Product        `json:"product" gorm:"foreignKey:ProductID"`
	Rating       int            `json:"rating" gorm:"check:rating >= 1 AND rating <= 5"`
	Title        string         `json:"title" gorm:"type:varchar(255)"`
	Comment      string         `json:"comment"`
	Pros         string         `json:"pros" gorm:"type:text"`
	Cons         string         `json:"cons" gorm:"type:text"`
	Status       string         `json:"status" gorm:"type:varchar(20);default:'pending'"`
	Images       datatypes.JSON `json:"images" gorm:"type:jsonb"`
	AdminReply   string         `json:"admin_reply"`
	IPAddress    string         `json:"ip_address"`
	UserAgent    string         `json:"user_agent" gorm:"type:text"`
	DeviceInfo   datatypes.JSON `json:"device_info" gorm:"type:jsonb"`
	HasPurchased bool           `json:"has_purchased"`
	HelpfulCount int            `json:"helpful_count" gorm:"default:0"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type ReviewImage struct {
	gorm.Model
	ReviewID uint
	Review   Review
	URL      string `gorm:"size:255;not null"`
}
