package services

import (
	"fmt"
	"math/rand"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
)

// OrderService service برای مدیریت سفارش‌ها
type OrderService struct {
	orderRepo     repository.OrderRepositoryInterface
	inventoryRepo repository.InventoryRepositoryInterface
}

// NewOrderService ایجاد instance جدید از OrderService
func NewOrderService(orderRepo repository.OrderRepositoryInterface, inventoryRepo repository.InventoryRepositoryInterface) *OrderService {
	return &OrderService{
		orderRepo:     orderRepo,
		inventoryRepo: inventoryRepo,
	}
}

// CreateOrder ایجاد سفارش جدید
func (s *OrderService) CreateOrder(order *models.Order) error {
	// بررسی و رزرو موجودی
	if err := s.inventoryRepo.CheckAndReserveInventory(order); err != nil {
		return fmt.Errorf("failed to reserve inventory: %w", err)
	}

	// ایجاد سفارش
	return s.orderRepo.CreateOrder(order)
}

// GetOrderByID دریافت سفارش بر اساس ID
func (s *OrderService) GetOrderByID(orderID uint) (*models.Order, error) {
	return s.orderRepo.GetOrderByID(orderID)
}

// GetOrderByOrderNumber دریافت سفارش بر اساس شماره سفارش
func (s *OrderService) GetOrderByOrderNumber(orderNumber string) (*models.Order, error) {
	return s.orderRepo.GetOrderByOrderNumber(orderNumber)
}

// GetUserOrders دریافت سفارش‌های یک کاربر
func (s *OrderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.orderRepo.GetUserOrders(userID)
}

// UpdateOrder به‌روزرسانی سفارش
func (s *OrderService) UpdateOrder(order *models.Order) error {
	return s.orderRepo.UpdateOrder(order)
}

// UpdateOrderStatus به‌روزرسانی وضعیت سفارش
func (s *OrderService) UpdateOrderStatus(orderID uint, status string) error {
	return s.orderRepo.UpdateOrderStatus(orderID, status)
}

// UpdatePaymentStatus به‌روزرسانی وضعیت پرداخت
func (s *OrderService) UpdatePaymentStatus(orderID uint, status string) error {
	// اگر پرداخت انجام شده، موجودی را به فروش تبدیل می‌کنیم
	if status == models.PaymentStatusPaid {
		if err := s.inventoryRepo.ProcessOrderReservation(orderID); err != nil {
			return fmt.Errorf("failed to process order reservation: %w", err)
		}
	}

	return s.orderRepo.UpdatePaymentStatus(orderID, status)
}

// CancelOrder لغو سفارش
func (s *OrderService) CancelOrder(orderID uint) error {
	// لغو رزرو موجودی
	if err := s.inventoryRepo.CancelOrderReservation(orderID); err != nil {
		return fmt.Errorf("failed to cancel order reservation: %w", err)
	}

	// لغو سفارش
	return s.orderRepo.CancelOrder(orderID)
}

// ProcessExpiredOrders پردازش سفارش‌های منقضی شده
func (s *OrderService) ProcessExpiredOrders() error {
	// آزادسازی رزروهای منقضی شده
	if err := s.inventoryRepo.ReleaseExpiredReservations(); err != nil {
		return fmt.Errorf("failed to release expired reservations: %w", err)
	}

	// لغو سفارش‌های منقضی شده
	return s.orderRepo.ReleaseExpiredReservations()
}

// GetPendingOrders دریافت سفارش‌های در انتظار
func (s *OrderService) GetPendingOrders() ([]models.Order, error) {
	return s.orderRepo.GetPendingOrders()
}

// GetExpiredReservations دریافت رزروهای منقضی شده
func (s *OrderService) GetExpiredReservations() ([]models.InventoryReservation, error) {
	return s.orderRepo.GetExpiredReservations()
}

// CalculateOrderTotals محاسبه مجموع‌های سفارش
func (s *OrderService) CalculateOrderTotals(order *models.Order) {
	subtotal := 0.0

	for _, item := range order.OrderItems {
		item.TotalPrice = item.CalculateTotal()
		subtotal += item.TotalPrice
	}

	order.SubtotalAmount = subtotal
	order.TotalAmount = subtotal + order.TaxAmount + order.ShippingAmount - order.DiscountAmount
}

// ValidateOrderItem اعتبارسنجی آیتم سفارش
func (s *OrderService) ValidateOrderItem(item *models.OrderItem) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}

	if item.UnitPrice < 0 {
		return fmt.Errorf("unit price cannot be negative")
	}

	return nil
}

// GetOrderSummary دریافت خلاصه سفارش
func (s *OrderService) GetOrderSummary(orderID uint) (map[string]interface{}, error) {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	summary := map[string]interface{}{
		"order_id":        order.ID,
		"order_number":    order.OrderNumber,
		"user_id":         order.UserID,
		"status":          order.Status,
		"payment_status":  order.PaymentStatus,
		"total_amount":    order.TotalAmount,
		"subtotal_amount": order.SubtotalAmount,
		"tax_amount":      order.TaxAmount,
		"shipping_amount": order.ShippingAmount,
		"discount_amount": order.DiscountAmount,
		"items_count":     len(order.OrderItems),
		"created_at":      order.CreatedAt,
		"updated_at":      order.UpdatedAt,
	}

	return summary, nil
}

// CanUserAccessOrder بررسی دسترسی کاربر به سفارش
func (s *OrderService) CanUserAccessOrder(orderID uint, userID uint) (bool, error) {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return false, err
	}

	// اگر userID nil باشد، سفارش برای کاربر مهمان است
	if order.UserID == nil {
		return false, nil
	}

	return *order.UserID == userID, nil
}

// GenerateOrderNumber تولید شماره سفارش جدید
func (s *OrderService) GenerateOrderNumber() string {
	random := fmt.Sprintf("%06d", rand.Intn(1000000))
	return fmt.Sprintf("ORD-%s", random)
}

// IsOrderExpired بررسی انقضای سفارش
func (s *OrderService) IsOrderExpired(order *models.Order) bool {
	if order.ReservedUntil == nil {
		return false
	}
	return time.Now().After(*order.ReservedUntil)
}

// ShouldReleaseReservation بررسی اینکه آیا رزرو سفارش باید آزاد شود
func (s *OrderService) ShouldReleaseReservation(order *models.Order) bool {
	return s.IsOrderExpired(order) && !order.IsPaidStatus()
}
