package services_test

import (
	"testing"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
)

func TestPaymentServiceFactory_CreateService(t *testing.T) {
	factory := services.NewPaymentServiceFactory()

	t.Run("Zarinpal", func(t *testing.T) {
		gateway := &models.PaymentGateway{Type: "zarinpal"}
		service, err := factory.CreateService(gateway)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if service == nil {
			t.Error("expected service, got nil")
		}
	})

	t.Run("Unsupported", func(t *testing.T) {
		gateway := &models.PaymentGateway{Type: "unsupported"}
		service, err := factory.CreateService(gateway)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if service != nil {
			t.Error("expected nil service, got service")
		}
	})

	t.Run("NilGateway", func(t *testing.T) {
		service, err := factory.CreateService(nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if service != nil {
			t.Error("expected nil service, got service")
		}
	})
}

func TestPaymentServiceFactory_ValidateGatewayConfig(t *testing.T) {
	factory := services.NewPaymentServiceFactory()

	t.Run("Valid", func(t *testing.T) {
		gateway := &models.PaymentGateway{
			Name:       "Zarinpal",
			Type:       "zarinpal",
			MerchantId: "123456789",
			ApiEndpoints: models.PaymentGatewayEndpoints{
				Payment: "https://api.zarinpal.com/pg/v4/payment/request.json",
			},
		}
		err := factory.ValidateGatewayConfig(gateway)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("InvalidType", func(t *testing.T) {
		gateway := &models.PaymentGateway{Name: "Test", Type: "invalid"}
		err := factory.ValidateGatewayConfig(gateway)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("MissingName", func(t *testing.T) {
		gateway := &models.PaymentGateway{Type: "zarinpal"}
		err := factory.ValidateGatewayConfig(gateway)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
