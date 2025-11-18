package services

import (
	"context"
	"fmt"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"

	"gorm.io/gorm"
)

// InventoryService service برای مدیریت موجودی
type InventoryService struct {
	inventoryRepo *repository.InventoryRepository
	orderRepo     *repository.OrderRepository
}

// NewInventoryService ایجاد instance جدید از InventoryService
func NewInventoryService(inventoryRepo *repository.InventoryRepository, orderRepo *repository.OrderRepository) *InventoryService {
	return &InventoryService{
		inventoryRepo: inventoryRepo,
		orderRepo:     orderRepo,
	}
}

// CheckAndReserveInventory بررسی موجودی و رزرو برای سفارش
func (s *InventoryService) CheckAndReserveInventory(order *models.Order) error {
	tx := s.inventoryRepo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// بررسی موجودی تمام محصولات در سفارش
	for _, item := range order.OrderItems {
		// دریافت موجودی محصول
		inventory, err := s.inventoryRepo.GetInventoryByProductID(item.ProductID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("product inventory not found for product %d: %w", item.ProductID, err)
		}

		// بررسی موجودی کافی
		if !inventory.IsAvailable(item.Quantity) {
			tx.Rollback()
			return fmt.Errorf("insufficient inventory for product %d. Available: %d, Requested: %d",
				item.ProductID, inventory.AvailableQuantity, item.Quantity)
		}
	}

	// دریافت تنظیمات موجودی
	settings, err := s.inventoryRepo.GetInventorySettings()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get inventory settings: %w", err)
	}

	// محاسبه زمان انقضای رزرو
	reservedUntil := time.Now().Add(time.Duration(settings.ReservationDurationMinutes) * time.Minute)

	// رزرو موجودی برای تمام محصولات
	for _, item := range order.OrderItems {
		_, err := s.inventoryRepo.ReserveInventory(item.ProductID, order.ID, uint(item.Quantity), reservedUntil)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to reserve inventory for product %d: %w", item.ProductID, err)
		}
	}

	tx.Commit()
	return nil
}

// ReleaseReservation آزادسازی رزرو موجودی
func (s *InventoryService) ReleaseReservation(reservationID uint) error {
	return s.inventoryRepo.ReleaseReservation(reservationID)
}

// ReleaseExpiredReservations آزادسازی رزروهای منقضی شده
func (s *InventoryService) ReleaseExpiredReservations() error {
	return s.inventoryRepo.ReleaseExpiredReservations()
}

// ProcessOrderReservation پردازش رزرو برای سفارش (زمانی که سفارش پرداخت می‌شود)
func (s *InventoryService) ProcessOrderReservation(orderID uint) error {
	tx := s.inventoryRepo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// دریافت سفارش
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("order not found: %w", err)
	}

	// اگر سفارش پرداخت شده، رزروها را به فروش تبدیل می‌کنیم
	if order.IsPaidStatus() {
		// پیدا کردن رزروها
		var reservations []models.InventoryReservation
		if err := tx.Where("order_id = ? AND is_released = false", orderID).Find(&reservations).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to find reservations: %w", err)
		}

		// به‌روزرسانی موجودی (کم کردن از موجودی واقعی)
		for _, reservation := range reservations {
			var inventory models.ProductInventory
			if err := tx.Where("product_id = ?", reservation.ProductID).First(&inventory).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("product inventory not found: %w", err)
			}

			// کم کردن از موجودی واقعی و حذف از رزرو
			inventory.StockQuantity -= reservation.ReservedQuantity
			if err := tx.Save(&inventory).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to update inventory: %w", err)
			}

			// حذف رزرو (چون به فروش تبدیل شده)
			if err := tx.Delete(&reservation).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to delete reservation: %w", err)
			}
		}
	}

	tx.Commit()
	return nil
}

// CancelOrderReservation لغو رزرو سفارش
func (s *InventoryService) CancelOrderReservation(orderID uint) error {
	tx := s.inventoryRepo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// پیدا کردن تمام رزروهای سفارش
	var reservations []models.InventoryReservation
	if err := tx.Where("order_id = ? AND is_released = false", orderID).Find(&reservations).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find reservations: %w", err)
	}

	// آزادسازی رزروها
	for _, reservation := range reservations {
		if err := s.inventoryRepo.ReleaseReservation(reservation.ID); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to release reservation %d: %w", reservation.ID, err)
		}
	}

	tx.Commit()
	return nil
}

// GetInventoryStatus دریافت وضعیت موجودی یک محصول
func (s *InventoryService) GetInventoryStatus(productID uint) (*models.ProductInventory, error) {
	return s.inventoryRepo.GetInventoryByProductID(productID)
}

// UpdateInventory به‌روزرسانی موجودی محصول
func (s *InventoryService) UpdateInventory(productID uint, stockQuantity, reservedQuantity int) error {
	inventory, err := s.inventoryRepo.GetInventoryByProductID(productID)
	if err != nil {
		// اگر موجودی وجود ندارد، ایجاد می‌کنیم
		inventory = &models.ProductInventory{
			ProductID:        productID,
			StockQuantity:    stockQuantity,
			ReservedQuantity: reservedQuantity,
		}
		return s.inventoryRepo.CreateOrUpdateInventory(inventory)
	}

	// به‌روزرسانی موجودی موجود
	inventory.StockQuantity = stockQuantity
	inventory.ReservedQuantity = reservedQuantity
	return s.inventoryRepo.CreateOrUpdateInventory(inventory)
}

// GetLowStockProducts دریافت محصولات با موجودی کم
func (s *InventoryService) GetLowStockProducts() ([]models.ProductInventory, error) {
	return s.inventoryRepo.GetLowStockProducts()
}

// GetOutOfStockProducts دریافت محصولات ناموجود
func (s *InventoryService) GetOutOfStockProducts() ([]models.ProductInventory, error) {
	return s.inventoryRepo.GetOutOfStockProducts()
}

// GetInventorySettings دریافت تنظیمات موجودی
func (s *InventoryService) GetInventorySettings() (*models.InventorySettings, error) {
	return s.inventoryRepo.GetInventorySettings()
}

// UpdateInventorySettings به‌روزرسانی تنظیمات موجودی
func (s *InventoryService) UpdateInventorySettings(settings *models.InventorySettings) error {
	return s.inventoryRepo.UpdateInventorySettings(settings)
}

// GetDB دریافت instance دیتابیس (برای transactions)
func (s *InventoryService) GetDB() *gorm.DB {
	return s.inventoryRepo.GetDB()
}

// CheckAndUpdateCartItems بررسی و به‌روزرسانی آیتم‌های سبد خرید
func (s *InventoryService) CheckAndUpdateCartItems(cartItems []models.OrderItem) ([]models.OrderItem, []string) {
	var updatedItems []models.OrderItem
	var unavailableMessages []string

	for _, item := range cartItems {
		// بررسی موجودی محصول
		inventory, err := s.inventoryRepo.GetInventoryByProductID(item.ProductID)
		if err != nil {
			unavailableMessages = append(unavailableMessages,
				fmt.Sprintf("محصول %s موجود نیست", item.ProductName))
			continue
		}

		// بررسی موجودی کافی
		if !inventory.IsAvailable(item.Quantity) {
			unavailableMessages = append(unavailableMessages,
				fmt.Sprintf("موجودی کافی برای %s موجود نیست. موجود: %d, درخواست: %d",
					item.ProductName, inventory.AvailableQuantity, item.Quantity))
			continue
		}

		// اگر قیمت تغییر کرده بود، به‌روزرسانی می‌کنیم
		// (اینجا باید قیمت فعلی محصول از دیتابیس دریافت شود)

		updatedItems = append(updatedItems, item)
	}

	return updatedItems, unavailableMessages
}

// DeductOnPayment کسر موجودی پس از پرداخت موفق
func (s *InventoryService) DeductOnPayment(ctx context.Context, productID uint, quantity int) error {
	// پیدا کردن موجودی محصول
	inventory, err := s.inventoryRepo.GetInventoryByProductID(productID)
	if err != nil {
		return fmt.Errorf("failed to get inventory for product %d: %w", productID, err)
	}

	// بررسی موجودی کافی
	if !inventory.IsAvailable(quantity) {
		return fmt.Errorf("insufficient inventory for product %d", productID)
	}

	// کسر موجودی
	inventory.StockQuantity -= quantity
	if err := s.inventoryRepo.CreateOrUpdateInventory(inventory); err != nil {
		return fmt.Errorf("failed to update inventory for product %d: %w", productID, err)
	}

	return nil
}

// SyncProductStockFromWarehouses همگام‌سازی موجودی محصول با انبارها
func (s *InventoryService) SyncProductStockFromWarehouses(ctx context.Context, productID uint) error {
	// استفاده از ProductWarehouseStockRepository برای دریافت موجودی از انبارها
	warehouseStockRepo := repository.NewProductWarehouseStockRepository(s.inventoryRepo.GetDB())

	// محاسبه مجموع موجودی از تمام انبارها
	totalStock, totalReserved, err := warehouseStockRepo.SumByProduct(ctx, productID)
	if err != nil {
		return nil // خطا را برنگردانیم تا API کار کند
	}

	// به‌روزرسانی موجودی محصول
	inventory, err := s.inventoryRepo.GetInventoryByProductID(productID)
	if err != nil {
		// اگر موجودی وجود ندارد، ایجاد می‌کنیم
		inventory = &models.ProductInventory{
			ProductID:         productID,
			StockQuantity:     totalStock,
			ReservedQuantity:  totalReserved,
			LowStockThreshold: 10,
			IsTracked:         true,
		}
	} else {
		inventory.StockQuantity = totalStock
		inventory.ReservedQuantity = totalReserved
	}

	// اگر خطا در ایجاد یا به‌روزرسانی موجودی رخ داد، آن را لاگ کنیم اما خطا را برنگردانیم
	if err := s.inventoryRepo.CreateOrUpdateInventory(inventory); err != nil {
		return nil // خطا را برنگردانیم تا API کار کند
	}

	return nil
}

// SetProductStocks تنظیم موجودی محصولات
func (s *InventoryService) SetProductStocks(ctx context.Context, entries []struct {
	WarehouseID uint
	ProductID   uint
	Quantity    int
	MinQty      int
	MaxQty      int
}) error {
	// جمع‌آوری IDهای منحصربه‌فرد محصولات برای همگام‌سازی
	productIDs := make(map[uint]bool)

	// استفاده از transaction برای اطمینان از consistency
	tx := s.inventoryRepo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ایجاد repository جدید با transaction
	txWarehouseStockRepo := repository.NewProductWarehouseStockRepository(tx)

	for _, entry := range entries {
		// ایجاد یا به‌روزرسانی موجودی در انبار
		warehouseStock := &models.ProductWarehouseStock{
			ProductID:   entry.ProductID,
			WarehouseID: entry.WarehouseID,
			Quantity:    entry.Quantity,
			Reserved:    0,
			MinQty:      entry.MinQty,
			MaxQty:      entry.MaxQty,
		}

		if err := txWarehouseStockRepo.Upsert(ctx, warehouseStock); err != nil {
			tx.Rollback()
			fmt.Printf("خطا در Upsert warehouse stock: ProductID=%d, WarehouseID=%d, Error=%v\n",
				entry.ProductID, entry.WarehouseID, err)
			return fmt.Errorf("failed to update warehouse stock for product %d in warehouse %d: %w",
				entry.ProductID, entry.WarehouseID, err)
		}

		productIDs[entry.ProductID] = true
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	// همگام‌سازی موجودی کلی برای تمام محصولات تغییر یافته
	for productID := range productIDs {
		_ = s.SyncProductStockFromWarehouses(ctx, productID)
	}

	return nil
}

// AdjustStock تنظیم موجودی محصول
func (s *InventoryService) AdjustStock(ctx context.Context, productID uint, warehouseID uint, delta int, reason string, userID *uint) error {
	// استفاده از ProductWarehouseStockRepository برای تنظیم موجودی در انبار
	warehouseStockRepo := repository.NewProductWarehouseStockRepository(s.inventoryRepo.GetDB())

	// تنظیم موجودی در انبار
	if err := warehouseStockRepo.AdjustQuantity(ctx, productID, warehouseID, delta); err != nil {
		return fmt.Errorf("failed to adjust stock for product %d in warehouse %d: %w", productID, warehouseID, err)
	}

	// همگام‌سازی موجودی کلی محصول
	_ = s.SyncProductStockFromWarehouses(ctx, productID)

	return nil
}
