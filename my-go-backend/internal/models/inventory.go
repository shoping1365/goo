package models

import (
	"time"
	"gorm.io/gorm"
)

// ProductInventory موجودی محصولات
type ProductInventory struct {
	ID                   uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID            uint      `json:"product_id" gorm:"uniqueIndex;not null"`
	StockQuantity        int       `json:"stock_quantity" gorm:"not null;default:0"`
	ReservedQuantity     int       `json:"reserved_quantity" gorm:"not null;default:0"`
	AvailableQuantity    int       `json:"available_quantity" gorm:"->"`
	LowStockThreshold    int       `json:"low_stock_threshold" gorm:"not null;default:10"`
	IsTracked            bool      `json:"is_tracked" gorm:"not null;default:true"`
	UpdatedAt            time.Time `json:"updated_at"`
	CreatedAt            time.Time `json:"created_at"`

	// روابط
	Product              Product              `json:"product" gorm:"foreignKey:ProductID"`
	Reservations         []InventoryReservation `json:"reservations" gorm:"foreignKey:ProductID"`
}

// TableName نام جدول را مشخص می‌کند
func (ProductInventory) TableName() string {
	return "product_inventory"
}

// InventoryReservation رزرو موجودی
type InventoryReservation struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID        uint      `json:"product_id" gorm:"not null"`
	OrderID          uint      `json:"order_id" gorm:"not null"`
	ReservedQuantity int       `json:"reserved_quantity" gorm:"not null"`
	ReservedUntil    time.Time `json:"reserved_until" gorm:"not null"`
	IsReleased       bool      `json:"is_released" gorm:"not null;default:false"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedAt        time.Time `json:"created_at"`

	// روابط
	Product          Product `json:"product" gorm:"foreignKey:ProductID"`
	Order            Order   `json:"order" gorm:"foreignKey:OrderID"`
}

// TableName نام جدول را مشخص می‌کند
func (InventoryReservation) TableName() string {
	return "inventory_reservations"
}

// InventorySettings تنظیمات سیستم موجودی
type InventorySettings struct {
	ID                            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ReservationDurationMinutes    int       `json:"reservation_duration_minutes" gorm:"not null;default:30"`
	AutoReleaseEnabled            bool      `json:"auto_release_enabled" gorm:"not null;default:true"`
	LowStockNotificationEnabled   bool      `json:"low_stock_notification_enabled" gorm:"not null;default:true"`
	OutOfStockNotificationEnabled bool      `json:"out_of_stock_notification_enabled" gorm:"not null;default:true"`
	UpdatedAt                     time.Time `json:"updated_at"`
	CreatedAt                     time.Time `json:"created_at"`
}

// TableName نام جدول را مشخص می‌کند
func (InventorySettings) TableName() string {
	return "inventory_settings"
}

// متدهای مفید برای ProductInventory
func (pi *ProductInventory) IsAvailable(quantity int) bool {
	return pi.AvailableQuantity >= quantity
}

func (pi *ProductInventory) IsLowStock() bool {
	return pi.AvailableQuantity <= pi.LowStockThreshold
}

func (pi *ProductInventory) IsOutOfStock() bool {
	return pi.AvailableQuantity <= 0
}

// متدهای مفید برای InventoryReservation
func (ir *InventoryReservation) IsExpired() bool {
	return time.Now().After(ir.ReservedUntil)
}

func (ir *InventoryReservation) CanBeReleased() bool {
	return !ir.IsReleased && ir.IsExpired()
}

// hooks برای GORM
func (pi *ProductInventory) BeforeCreate(tx *gorm.DB) (err error) {
	pi.CreatedAt = time.Now()
	pi.UpdatedAt = time.Now()
	return
}

func (pi *ProductInventory) BeforeUpdate(tx *gorm.DB) (err error) {
	pi.UpdatedAt = time.Now()
	return
}

func (ir *InventoryReservation) BeforeCreate(tx *gorm.DB) (err error) {
	ir.CreatedAt = time.Now()
	ir.UpdatedAt = time.Now()
	return
}

func (ir *InventoryReservation) BeforeUpdate(tx *gorm.DB) (err error) {
	ir.UpdatedAt = time.Now()
	return
}

func (is *InventorySettings) BeforeCreate(tx *gorm.DB) (err error) {
	is.CreatedAt = time.Now()
	is.UpdatedAt = time.Now()
	return
}

func (is *InventorySettings) BeforeUpdate(tx *gorm.DB) (err error) {
	is.UpdatedAt = time.Now()
	return
}
