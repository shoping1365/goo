package services_test

import (
	"errors"
	"testing"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/tests/mocks"
)

func TestOrderService_CreateOrder(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockOrderRepo := &mocks.MockOrderRepository{}
		mockInventoryRepo := &mocks.MockInventoryRepository{}

		service := services.NewOrderService(mockOrderRepo, mockInventoryRepo)

		userID := uint(1)
		order := &models.Order{
			UserID: &userID,
			OrderItems: []models.OrderItem{
				{ProductID: 1, Quantity: 2},
			},
		}

		mockInventoryRepo.CheckAndReserveInventoryFunc = func(o *models.Order) error {
			if o != order {
				return errors.New("unexpected order")
			}
			return nil
		}

		mockOrderRepo.CreateOrderFunc = func(o *models.Order) error {
			if o != order {
				return errors.New("unexpected order")
			}
			return nil
		}

		err := service.CreateOrder(order)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("InventoryReservationFailed", func(t *testing.T) {
		mockOrderRepo := &mocks.MockOrderRepository{}
		mockInventoryRepo := &mocks.MockInventoryRepository{}

		service := services.NewOrderService(mockOrderRepo, mockInventoryRepo)

		userID := uint(1)
		order := &models.Order{
			UserID: &userID,
			OrderItems: []models.OrderItem{
				{ProductID: 1, Quantity: 2},
			},
		}

		mockInventoryRepo.CheckAndReserveInventoryFunc = func(o *models.Order) error {
			return errors.New("insufficient inventory")
		}

		err := service.CreateOrder(order)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if err.Error() != "failed to reserve inventory: insufficient inventory" {
			t.Errorf("expected error message 'failed to reserve inventory: insufficient inventory', got '%v'", err.Error())
		}
	})
}

func TestOrderService_UpdatePaymentStatus(t *testing.T) {
	t.Run("Success_Paid", func(t *testing.T) {
		mockOrderRepo := &mocks.MockOrderRepository{}
		mockInventoryRepo := &mocks.MockInventoryRepository{}

		service := services.NewOrderService(mockOrderRepo, mockInventoryRepo)

		orderID := uint(1)
		status := models.PaymentStatusPaid

		mockInventoryRepo.ProcessOrderReservationFunc = func(id uint) error {
			if id != orderID {
				return errors.New("unexpected order ID")
			}
			return nil
		}

		mockOrderRepo.UpdatePaymentStatusFunc = func(id uint, s string) error {
			if id != orderID || s != status {
				return errors.New("unexpected arguments")
			}
			return nil
		}

		err := service.UpdatePaymentStatus(orderID, status)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Success_Pending", func(t *testing.T) {
		mockOrderRepo := &mocks.MockOrderRepository{}
		mockInventoryRepo := &mocks.MockInventoryRepository{}

		service := services.NewOrderService(mockOrderRepo, mockInventoryRepo)

		orderID := uint(1)
		status := models.PaymentStatusPending

		// ProcessOrderReservation should NOT be called

		mockOrderRepo.UpdatePaymentStatusFunc = func(id uint, s string) error {
			if id != orderID || s != status {
				return errors.New("unexpected arguments")
			}
			return nil
		}

		err := service.UpdatePaymentStatus(orderID, status)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}

func TestOrderService_CancelOrder(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockOrderRepo := &mocks.MockOrderRepository{}
		mockInventoryRepo := &mocks.MockInventoryRepository{}

		service := services.NewOrderService(mockOrderRepo, mockInventoryRepo)

		orderID := uint(1)

		mockInventoryRepo.CancelOrderReservationFunc = func(id uint) error {
			if id != orderID {
				return errors.New("unexpected order ID")
			}
			return nil
		}

		mockOrderRepo.CancelOrderFunc = func(id uint) error {
			if id != orderID {
				return errors.New("unexpected order ID")
			}
			return nil
		}

		err := service.CancelOrder(orderID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}
