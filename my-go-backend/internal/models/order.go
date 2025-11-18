package models

import (
	"time"

	"gorm.io/gorm"
)

// Order سفارش
type Order struct {
	ID                uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderNumber       string     `json:"order_number" gorm:"uniqueIndex;not null"`
	UserID            *uint      `json:"user_id"`
	ShippingAddressID uint       `json:"shipping_address_id"`
	IsPaid            bool       `json:"is_paid" gorm:"default:false"`
	FinalAmount       float64    `json:"final_amount"`
	TrackingCode      string     `json:"tracking_code"`
	Status            string     `json:"status" gorm:"type:enum('pending', 'processing', 'shipped', 'delivered', 'cancelled', 'refunded');not null;default:'pending'"`
	TotalAmount       float64    `json:"total_amount" gorm:"not null"`
	SubtotalAmount    float64    `json:"subtotal_amount" gorm:"not null"`
	TaxAmount         float64    `json:"tax_amount" gorm:"default:0"`
	ShippingAmount    float64    `json:"shipping_amount" gorm:"default:0"`
	DiscountAmount    float64    `json:"discount_amount" gorm:"default:0"`
	PaymentStatus     string     `json:"payment_status" gorm:"type:enum('pending', 'paid', 'failed', 'refunded');not null;default:'pending'"`
	PaymentMethod     string     `json:"payment_method"`
	ShippingAddress   string     `json:"shipping_address" gorm:"type:text"`
	BillingAddress    string     `json:"billing_address" gorm:"type:text"`
	Notes             string     `json:"notes" gorm:"type:text"`
	ReservedUntil     *time.Time `json:"reserved_until"`
	UpdatedAt         time.Time  `json:"updated_at"`
	CreatedAt         time.Time  `json:"created_at"`

	// روابط
	User         *User                  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OrderItems   []OrderItem            `json:"order_items" gorm:"foreignKey:OrderID"`
	Reservations []InventoryReservation `json:"reservations" gorm:"foreignKey:OrderID"`
}

// OrderItem جزئیات سفارش
type OrderItem struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID      uint      `json:"order_id" gorm:"not null"`
	ProductID    uint      `json:"product_id" gorm:"not null"`
	ProductName  string    `json:"product_name" gorm:"not null"`
	ProductImage string    `json:"product_image"`
	ProductSKU   string    `json:"product_sku"`
	UnitPrice    float64   `json:"unit_price" gorm:"not null"`
	Quantity     int       `json:"quantity" gorm:"not null"`
	TotalPrice   float64   `json:"total_price" gorm:"not null"`
	FinalPrice   float64   `json:"final_price"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`

	// روابط
	Order   Order   `json:"order" gorm:"foreignKey:OrderID"`
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}

// متدهای مفید برای Order
func (o *Order) IsPending() bool {
	return o.Status == "pending"
}

func (o *Order) IsPaidStatus() bool {
	return o.PaymentStatus == "paid"
}

func (o *Order) IsCancelled() bool {
	return o.Status == "cancelled"
}

func (o *Order) CanReserveInventory() bool {
	return o.IsPending() && !o.IsPaidStatus()
}

func (o *Order) ShouldReleaseReservation() bool {
	if o.ReservedUntil == nil {
		return false
	}
	return time.Now().After(*o.ReservedUntil) && !o.IsPaidStatus()
}

// متدهای مفید برای OrderItem
func (oi *OrderItem) CalculateTotal() float64 {
	return oi.UnitPrice * float64(oi.Quantity)
}

// hooks برای GORM
func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	return
}

func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	o.UpdatedAt = time.Now()
	return
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	oi.CreatedAt = time.Now()
	oi.UpdatedAt = time.Now()
	// TotalPrice قبلاً در handler محاسبه شده است
	return
}

func (oi *OrderItem) BeforeUpdate(tx *gorm.DB) (err error) {
	oi.UpdatedAt = time.Now()
	oi.TotalPrice = oi.CalculateTotal()
	return
}

// OrderStatus constants
const (
	OrderStatusPending    = "pending"
	OrderStatusProcessing = "processing"
	OrderStatusShipped    = "shipped"
	OrderStatusDelivered  = "delivered"
	OrderStatusCancelled  = "cancelled"
	OrderStatusRefunded   = "refunded"
)

// PaymentStatus constants
const (
	PaymentStatusPending  = "pending"
	PaymentStatusPaid     = "paid"
	PaymentStatusFailed   = "failed"
	PaymentStatusRefunded = "refunded"
)
