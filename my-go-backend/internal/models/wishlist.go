package models

import (
	"time"

	"gorm.io/gorm"
)

// Wishlist مدل لیست علاقمندی‌های کاربر
type Wishlist struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// روابط
	User     User              `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Items    []WishlistItem    `json:"items,omitempty" gorm:"foreignKey:WishlistID"`
}

// WishlistItem آیتم‌های لیست علاقمندی‌ها
type WishlistItem struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	WishlistID uint           `json:"wishlist_id" gorm:"not null;index"`
	ProductID  uint           `json:"product_id" gorm:"not null;index"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	// روابط
	Wishlist Wishlist `json:"wishlist,omitempty" gorm:"foreignKey:WishlistID"`
	Product  Product  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

// TableName نام جدول
func (Wishlist) TableName() string {
	return "wishlists"
}

func (WishlistItem) TableName() string {
	return "wishlist_items"
}
