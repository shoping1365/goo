package models

import "time"

// InventoryMovement ثبت رویدادهای ورود/خروج/انتقال/تعدیل موجودی برای ممیزی
type InventoryMovement struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ProductID   uint      `json:"product_id" gorm:"index;not null"`
	WarehouseID uint      `json:"warehouse_id" gorm:"index;not null"`
	Delta       int       `json:"delta"`                              // مثبت=ورود، منفی=خروج
	Type        string    `json:"type" gorm:"type:varchar(32);index"` // in|out|transfer_in|transfer_out|adjust
	RefOrderID  *uint     `json:"ref_order_id" gorm:"index"`
	Notes       string    `json:"notes" gorm:"type:varchar(500)"`
	CreatedBy   *uint     `json:"created_by" gorm:"index"`
	CreatedAt   time.Time `json:"created_at"`
}
