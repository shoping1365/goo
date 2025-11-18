package models

import "time"

// ProductWarehouseStock نگهدارنده موجودی هر محصول در هر انبار
// مقدار quantity کل موجودی ثبت‌شده و reserved مربوط به رزروهای موقت است.
type ProductWarehouseStock struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ProductID   uint      `json:"product_id" gorm:"index;not null"`
	WarehouseID uint      `json:"warehouse_id" gorm:"index;not null"`
	Quantity    int       `json:"quantity"`
	Reserved    int       `json:"reserved"`
	MinQty      int       `json:"min_qty"`
	MaxQty      int       `json:"max_qty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
