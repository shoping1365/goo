package repository

import (
	"fmt"
	"math/rand"
	"time"

	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// OrderRepository repository برای مدیریت سفارش‌ها
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository ایجاد instance جدید از OrderRepository
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// GetDB دریافت instance دیتابیس
func (r *OrderRepository) GetDB() *gorm.DB {
	return r.db
}

// CreateOrder ایجاد سفارش جدید
func (r *OrderRepository) CreateOrder(order *models.Order) error {
	tx := r.db.Begin()

	// تولید شماره سفارش
	order.OrderNumber = r.generateOrderNumber()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	// اگر سفارش پرداخت نشده است، زمان رزرو را تنظیم می‌کنیم
	if order.PaymentStatus == models.PaymentStatusPending {
		settings, err := r.getInventorySettings()
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to get inventory settings: %w", err)
		}
		reservedUntil := time.Now().Add(time.Duration(settings.ReservationDurationMinutes) * time.Minute)
		order.ReservedUntil = &reservedUntil
	}

	// ذخیره سفارش
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create order: %w", err)
	}

	// ذخیره جزئیات سفارش
	for _, item := range order.OrderItems {
		item.OrderID = order.ID
		// CreatedAt و UpdatedAt در BeforeCreate hook تنظیم می‌شوند
		// TotalPrice قبلاً در handler محاسبه شده است

		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create order item: %w", err)
		}
	}

	tx.Commit()
	return nil
}

// GetOrderByID دریافت سفارش بر اساس ID
func (r *OrderRepository) GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("OrderItems").Preload("User").First(&order, orderID).Error
	return &order, err
}

// GetOrderByOrderNumber دریافت سفارش بر اساس شماره سفارش
func (r *OrderRepository) GetOrderByOrderNumber(orderNumber string) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("OrderItems").Preload("User").Where("order_number = ?", orderNumber).First(&order).Error
	return &order, err
}

// GetUserOrders دریافت سفارش‌های یک کاربر
func (r *OrderRepository) GetUserOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&orders).Error
	return orders, err
}

// UpdateOrder به‌روزرسانی سفارش
func (r *OrderRepository) UpdateOrder(order *models.Order) error {
	order.UpdatedAt = time.Now()
	return r.db.Save(order).Error
}

// UpdateOrderStatus به‌روزرسانی وضعیت سفارش
func (r *OrderRepository) UpdateOrderStatus(orderID uint, status string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

// UpdatePaymentStatus به‌روزرسانی وضعیت پرداخت
func (r *OrderRepository) UpdatePaymentStatus(orderID uint, status string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("payment_status", status).Error
}

// CancelOrder لغو سفارش
func (r *OrderRepository) CancelOrder(orderID uint) error {
	tx := r.db.Begin()

	// دریافت سفارش
	var order models.Order
	if err := tx.First(&order, orderID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("order not found")
	}

	// بررسی اینکه سفارش قابل لغو باشد
	if !order.CanReserveInventory() {
		tx.Rollback()
		return fmt.Errorf("order cannot be cancelled")
	}

	// به‌روزرسانی وضعیت سفارش
	order.Status = models.OrderStatusCancelled
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update order status: %w", err)
	}

	// آزادسازی رزروها
	if err := r.releaseOrderReservations(tx, orderID); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to release reservations: %w", err)
	}

	tx.Commit()
	return nil
}

// releaseOrderReservations آزادسازی رزروهای سفارش
func (r *OrderRepository) releaseOrderReservations(tx *gorm.DB, orderID uint) error {
	// پیدا کردن تمام رزروهای سفارش
	var reservations []models.InventoryReservation
	if err := tx.Where("order_id = ? AND is_released = false", orderID).Find(&reservations).Error; err != nil {
		return err
	}

	for _, reservation := range reservations {
		// آزادسازی رزرو
		reservation.IsReleased = true
		if err := tx.Save(&reservation).Error; err != nil {
			return err
		}

		// به‌روزرسانی موجودی
		var inventory models.ProductInventory
		if err := tx.Where("product_id = ?", reservation.ProductID).First(&inventory).Error; err != nil {
			continue // اگر موجودی پیدا نشد، ادامه می‌دهیم
		}

		inventory.ReservedQuantity -= reservation.ReservedQuantity
		if err := tx.Save(&inventory).Error; err != nil {
			return err
		}
	}

	return nil
}

// ReleaseExpiredReservations آزادسازی رزروهای منقضی شده
func (r *OrderRepository) ReleaseExpiredReservations() error {
	tx := r.db.Begin()

	// پیدا کردن سفارش‌های با رزرو منقضی شده
	var expiredOrders []models.Order
	if err := tx.Where("reserved_until < ? AND payment_status = ? AND status = ?",
		time.Now(), models.PaymentStatusPending, models.OrderStatusPending).Find(&expiredOrders).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find expired orders: %w", err)
	}

	for _, order := range expiredOrders {
		// لغو سفارش
		order.Status = models.OrderStatusCancelled
		if err := tx.Save(&order).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to cancel order %d: %w", order.ID, err)
		}

		// آزادسازی رزروها
		if err := r.releaseOrderReservations(tx, order.ID); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to release reservations for order %d: %w", order.ID, err)
		}
	}

	tx.Commit()
	return nil
}

// GetPendingOrders دریافت سفارش‌های در انتظار
func (r *OrderRepository) GetPendingOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("status = ? AND payment_status = ?", models.OrderStatusPending, models.PaymentStatusPending).Find(&orders).Error
	return orders, err
}

// GetExpiredReservations دریافت رزروهای منقضی شده
func (r *OrderRepository) GetExpiredReservations() ([]models.InventoryReservation, error) {
	var reservations []models.InventoryReservation
	err := r.db.Where("is_released = false AND reserved_until < ?", time.Now()).Find(&reservations).Error
	return reservations, err
}

// generateOrderNumber تولید شماره سفارش منحصر به فرد
func (r *OrderRepository) generateOrderNumber() string {
	random := fmt.Sprintf("%06d", rand.Intn(1000000))
	return fmt.Sprintf("ORD-%s", random)
}

// getInventorySettings دریافت تنظیمات موجودی (متد helper)
func (r *OrderRepository) getInventorySettings() (*models.InventorySettings, error) {
	var settings models.InventorySettings
	err := r.db.First(&settings).Error
	if err != nil {
		// اگر تنظیمات وجود ندارد، تنظیمات پیش‌فرض ایجاد می‌کنیم
		settings = models.InventorySettings{
			ReservationDurationMinutes:    30,
			AutoReleaseEnabled:            true,
			LowStockNotificationEnabled:   true,
			OutOfStockNotificationEnabled: true,
		}
		if createErr := r.db.Create(&settings).Error; createErr != nil {
			return nil, createErr
		}
	}
	return &settings, err
}
