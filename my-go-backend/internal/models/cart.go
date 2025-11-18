package models

import (
	"time"

	"gorm.io/gorm"
)

// Cart مدل سبد خرید
type Cart struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    *uint          `json:"user_id,omitempty" gorm:"index"`
	SessionID string         `json:"session_id" gorm:"index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// ارتباط با آیتم‌های سبد خرید
	Items []CartItem `json:"items,omitempty" gorm:"foreignKey:CartID"`
}

// CartItem مدل آیتم سبد خرید
type CartItem struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CartID         uint           `json:"cart_id" gorm:"not null;index"`
	ProductID      uint           `json:"product_id" gorm:"not null;index"`
	Quantity       int            `json:"quantity" gorm:"not null;default:1"`
	IsNextPurchase bool           `json:"is_next" gorm:"column:is_next_purchase;default:false"`
	UnitPrice      float64        `json:"unit_price" gorm:"not null;column:unit_price"`
	DiscountAmount float64        `json:"discount_amount" gorm:"column:discount_amount"`
	FinalPrice     float64        `json:"final_price" gorm:"not null;column:final_price"`
	CreatedAt      time.Time      `json:"created_at" gorm:"column:added_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// ارتباط با محصول
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

// CartResponse پاسخ API سبد خرید
type CartResponse struct {
	ID         uint       `json:"id"`
	UserID     *uint      `json:"user_id,omitempty"`
	SessionID  string     `json:"session_id"`
	Items      []CartItem `json:"items"`
	TotalItems int        `json:"total_items"`
	TotalPrice float64    `json:"total_price"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// AddToCartRequest درخواست افزودن به سبد خرید
type AddToCartRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,min=1"`
}

// UpdateCartItemRequest درخواست به‌روزرسانی آیتم سبد خرید
type UpdateCartItemRequest struct {
	CartItemID uint `json:"cart_item_id" binding:"required"`
	Quantity   int  `json:"quantity" binding:"required,min=1"`
}

// RemoveFromCartRequest درخواست حذف از سبد خرید
type RemoveFromCartRequest struct {
	CartItemID uint `json:"cart_item_id" binding:"required"`
}

// CartItemResponse پاسخ آیتم سبد خرید
type CartItemResponse struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Name      string    `json:"name"`
	SKU       string    `json:"sku"`
	Image     string    `json:"image"`
	Price     float64   `json:"price"`
	Features  []string  `json:"features"`
	CreatedAt time.Time `json:"created_at"`
}
