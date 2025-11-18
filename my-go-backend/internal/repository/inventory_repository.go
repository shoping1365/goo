package repository

import (
	"fmt"
	"time"

	"my-go-backend/internal/models"
	"gorm.io/gorm"
)

// InventoryRepository repository برای مدیریت موجودی
type InventoryRepository struct {
	db *gorm.DB
}

// NewInventoryRepository ایجاد instance جدید از InventoryRepository
func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

// GetDB دریافت instance دیتابیس
func (r *InventoryRepository) GetDB() *gorm.DB {
	return r.db
}

// GetInventoryByProductID دریافت موجودی یک محصول
func (r *InventoryRepository) GetInventoryByProductID(productID uint) (*models.ProductInventory, error) {
	var inventory models.ProductInventory
	err := r.db.Where("product_id = ?", productID).First(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// CreateOrUpdateInventory ایجاد یا به‌روزرسانی موجودی محصول
func (r *InventoryRepository) CreateOrUpdateInventory(inventory *models.ProductInventory) error {
	return r.db.Save(inventory).Error
}

// CheckAvailability بررسی موجودی محصول
func (r *InventoryRepository) CheckAvailability(productID uint, quantity int) (bool, error) {
	var inventory models.ProductInventory
	err := r.db.Where("product_id = ?", productID).First(&inventory).Error
	if err != nil {
		return false, err
	}
	return inventory.IsAvailable(quantity), nil
}

// ReserveInventory رزرو موجودی برای سفارش
func (r *InventoryRepository) ReserveInventory(productID, orderID, quantity uint, reservedUntil time.Time) (*models.InventoryReservation, error) {
	tx := r.db.Begin()

	// بررسی موجودی
	var inventory models.ProductInventory
	if err := tx.Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("product inventory not found")
	}

	if !inventory.IsAvailable(int(quantity)) {
		tx.Rollback()
		return nil, fmt.Errorf("insufficient inventory")
	}

	// ایجاد رزرو
	reservation := &models.InventoryReservation{
		ProductID:        productID,
		OrderID:          orderID,
		ReservedQuantity: int(quantity),
		ReservedUntil:    reservedUntil,
		IsReleased:       false,
	}

	if err := tx.Create(reservation).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create reservation: %w", err)
	}

	// به‌روزرسانی موجودی
	inventory.ReservedQuantity += int(quantity)
	if err := tx.Save(&inventory).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update inventory: %w", err)
	}

	tx.Commit()
	return reservation, nil
}

// ReleaseReservation آزادسازی رزرو موجودی
func (r *InventoryRepository) ReleaseReservation(reservationID uint) error {
	tx := r.db.Begin()

	// پیدا کردن رزرو
	var reservation models.InventoryReservation
	if err := tx.Where("id = ?", reservationID).First(&reservation).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("reservation not found")
	}

	// بررسی اینکه رزرو قبلاً آزاد نشده باشد
	if reservation.IsReleased {
		tx.Rollback()
		return fmt.Errorf("reservation already released")
	}

	// به‌روزرسانی رزرو
	reservation.IsReleased = true
	if err := tx.Save(&reservation).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update reservation: %w", err)
	}

	// به‌روزرسانی موجودی
	var inventory models.ProductInventory
	if err := tx.Where("product_id = ?", reservation.ProductID).First(&inventory).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("product inventory not found")
	}

	inventory.ReservedQuantity -= reservation.ReservedQuantity
	if err := tx.Save(&inventory).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update inventory: %w", err)
	}

	tx.Commit()
	return nil
}

// ReleaseExpiredReservations آزادسازی رزروهای منقضی شده
func (r *InventoryRepository) ReleaseExpiredReservations() error {
	tx := r.db.Begin()

	// پیدا کردن رزروهای منقضی شده
	var expiredReservations []models.InventoryReservation
	if err := tx.Where("is_released = false AND reserved_until < ?", time.Now()).Find(&expiredReservations).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find expired reservations: %w", err)
	}

	for _, reservation := range expiredReservations {
		// آزادسازی رزرو
		reservation.IsReleased = true
		if err := tx.Save(&reservation).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to release reservation %d: %w", reservation.ID, err)
		}

		// به‌روزرسانی موجودی
		var inventory models.ProductInventory
		if err := tx.Where("product_id = ?", reservation.ProductID).First(&inventory).Error; err != nil {
			continue // اگر موجودی پیدا نشد، ادامه می‌دهیم
		}

		inventory.ReservedQuantity -= reservation.ReservedQuantity
		if err := tx.Save(&inventory).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update inventory for product %d: %w", reservation.ProductID, err)
		}
	}

	tx.Commit()
	return nil
}

// GetLowStockProducts دریافت محصولات با موجودی کم
func (r *InventoryRepository) GetLowStockProducts() ([]models.ProductInventory, error) {
	var inventories []models.ProductInventory
	err := r.db.Where("is_tracked = true AND available_quantity <= low_stock_threshold").Find(&inventories).Error
	return inventories, err
}

// GetOutOfStockProducts دریافت محصولات ناموجود
func (r *InventoryRepository) GetOutOfStockProducts() ([]models.ProductInventory, error) {
	var inventories []models.ProductInventory
	err := r.db.Where("is_tracked = true AND available_quantity <= 0").Find(&inventories).Error
	return inventories, err
}

// GetInventorySettings دریافت تنظیمات موجودی
func (r *InventoryRepository) GetInventorySettings() (*models.InventorySettings, error) {
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

// UpdateInventorySettings به‌روزرسانی تنظیمات موجودی
func (r *InventoryRepository) UpdateInventorySettings(settings *models.InventorySettings) error {
	return r.db.Save(settings).Error
}

// CheckAndReserveInventory بررسی موجودی و رزرو برای سفارش
func (r *InventoryRepository) CheckAndReserveInventory(order *models.Order) error {
	tx := r.db.Begin()

	// پردازش هر آیتم سفارش
	for _, item := range order.OrderItems {
		// بررسی موجودی
		var inventory models.ProductInventory
		if err := tx.Where("product_id = ?", item.ProductID).First(&inventory).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("product inventory not found for product %d: %w", item.ProductID, err)
		}

		if !inventory.IsAvailable(item.Quantity) {
			tx.Rollback()
			return fmt.Errorf("insufficient inventory for product %d", item.ProductID)
		}

		// ایجاد رزرو
		settings, err := r.GetInventorySettings()
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to get inventory settings: %w", err)
		}

		reservedUntil := time.Now().Add(time.Duration(settings.ReservationDurationMinutes) * time.Minute)

		reservation := &models.InventoryReservation{
			ProductID:        item.ProductID,
			OrderID:          order.ID,
			ReservedQuantity: item.Quantity,
			ReservedUntil:    reservedUntil,
			IsReleased:       false,
		}

		if err := tx.Create(reservation).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create reservation for product %d: %w", item.ProductID, err)
		}

		// به‌روزرسانی موجودی
		inventory.ReservedQuantity += item.Quantity
		if err := tx.Save(&inventory).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update inventory for product %d: %w", item.ProductID, err)
		}
	}

	tx.Commit()
	return nil
}

// ProcessOrderReservation پردازش رزرو سفارش (تبدیل به فروش)
func (r *InventoryRepository) ProcessOrderReservation(orderID uint) error {
	tx := r.db.Begin()

	// پیدا کردن رزروهای سفارش
	var reservations []models.InventoryReservation
	if err := tx.Where("order_id = ? AND is_released = false", orderID).Find(&reservations).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find reservations for order %d: %w", orderID, err)
	}

	for _, reservation := range reservations {
		// به‌روزرسانی رزرو
		reservation.IsReleased = true
		if err := tx.Save(&reservation).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update reservation %d: %w", reservation.ID, err)
		}

		// به‌روزرسانی موجودی
		var inventory models.ProductInventory
		if err := tx.Where("product_id = ?", reservation.ProductID).First(&inventory).Error; err != nil {
			continue // اگر موجودی پیدا نشد، ادامه می‌دهیم
		}

		inventory.ReservedQuantity -= reservation.ReservedQuantity
		inventory.StockQuantity -= reservation.ReservedQuantity
		if err := tx.Save(&inventory).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update inventory for product %d: %w", reservation.ProductID, err)
		}
	}

	tx.Commit()
	return nil
}

// CancelOrderReservation لغو رزرو سفارش
func (r *InventoryRepository) CancelOrderReservation(orderID uint) error {
	tx := r.db.Begin()

	// پیدا کردن رزروهای سفارش
	var reservations []models.InventoryReservation
	if err := tx.Where("order_id = ? AND is_released = false", orderID).Find(&reservations).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find reservations for order %d: %w", orderID, err)
	}

	for _, reservation := range reservations {
		// به‌روزرسانی رزرو
		reservation.IsReleased = true
		if err := tx.Save(&reservation).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update reservation %d: %w", reservation.ID, err)
		}

		// به‌روزرسانی موجودی
		var inventory models.ProductInventory
		if err := tx.Where("product_id = ?", reservation.ProductID).First(&inventory).Error; err != nil {
			continue // اگر موجودی پیدا نشد، ادامه می‌دهیم
		}

		inventory.ReservedQuantity -= reservation.ReservedQuantity
		if err := tx.Save(&inventory).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update inventory for product %d: %w", reservation.ProductID, err)
		}
	}

	tx.Commit()
	return nil
}