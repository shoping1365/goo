package services

import (
	"errors"
	"fmt"

	"my-go-backend/internal/models"
)

// PaymentServiceFactory فکتوری برای ایجاد سرویس‌های درگاه پرداخت
// این فکتوری بر اساس نوع درگاه، سرویس مناسب را برمی‌گرداند

// PaymentServiceInterface رابط مشترک برای تمام سرویس‌های درگاه پرداخت
type PaymentServiceInterface interface {
	// CreatePayment ایجاد درخواست پرداخت جدید
	CreatePayment(amount int, callbackURL, description, email, mobile string) (*models.PaymentTransaction, error)

	// VerifyPayment تایید پرداخت
	VerifyPayment(amount int, authority string) (bool, string, error)

	// GetPaymentURL دریافت آدرس پرداخت
	GetPaymentURL(transaction *models.PaymentTransaction, callbackURL string) string

	// TestConnection تست اتصال به درگاه
	TestConnection() error

	// GetBalance دریافت موجودی حساب (در صورت پشتیبانی)
	GetBalance() (float64, error)

	// RefundPayment بازگشت وجه (در صورت پشتیبانی)
	RefundPayment(refID string, amount int) error
}

// PaymentServiceFactory فکتوری اصلی برای ایجاد سرویس‌های درگاه پرداخت
type PaymentServiceFactory struct{}

// NewPaymentServiceFactory سازنده فکتوری سرویس پرداخت
func NewPaymentServiceFactory() *PaymentServiceFactory {
	return &PaymentServiceFactory{}
}

// CreateService ایجاد سرویس درگاه پرداخت بر اساس نوع درگاه
func (f *PaymentServiceFactory) CreateService(gateway *models.PaymentGateway) (PaymentServiceInterface, error) {
	if gateway == nil {
		return nil, errors.New("درگاه پرداخت نمی‌تواند خالی باشد")
	}

	switch gateway.Type {
	case "zarinpal":
		return NewZarinpalService(gateway), nil
	case "saman":
		return NewSamanService(gateway), nil
	case "mellat":
		return NewMellatService(gateway), nil
	case "melli":
		return NewMelliService(gateway), nil
	case "parsian":
		return NewParsianService(gateway), nil
	case "paypal":
		return NewPayPalService(gateway), nil
	case "stripe":
		return NewStripeService(gateway), nil
	default:
		return nil, fmt.Errorf("نوع درگاه پرداخت '%s' پشتیبانی نمی‌شود", gateway.Type)
	}
}

// GetSupportedGateways دریافت لیست درگاه‌های پشتیبانی شده
func (f *PaymentServiceFactory) GetSupportedGateways() []string {
	return []string{
		"zarinpal",
		"saman",
		"mellat",
		"parsian",
		"paypal",
		"stripe",
	}
}

// ValidateGatewayConfig اعتبارسنجی تنظیمات درگاه پرداخت
func (f *PaymentServiceFactory) ValidateGatewayConfig(gateway *models.PaymentGateway) error {
	if gateway == nil {
		return errors.New("درگاه پرداخت نمی‌تواند خالی باشد")
	}

	// بررسی نام درگاه
	if gateway.Name == "" {
		return errors.New("نام درگاه الزامی است")
	}

	// بررسی نوع درگاه
	supportedGateways := f.GetSupportedGateways()
	isSupported := false
	for _, supported := range supportedGateways {
		if supported == gateway.Type {
			isSupported = true
			break
		}
	}
	if !isSupported {
		return fmt.Errorf("نوع درگاه '%s' پشتیبانی نمی‌شود", gateway.Type)
	}

	// بررسی مرچنت ID
	if gateway.MerchantId == "" {
		return errors.New("مرچنت ID الزامی است")
	}

	// بررسی آدرس‌های API
	if gateway.ApiEndpoints.Payment == "" {
		return errors.New("آدرس API پرداخت الزامی است")
	}

	return nil
}

// TestGatewayConnection تست اتصال درگاه پرداخت
func (f *PaymentServiceFactory) TestGatewayConnection(gateway *models.PaymentGateway) error {
	service, err := f.CreateService(gateway)
	if err != nil {
		return fmt.Errorf("خطا در ایجاد سرویس: %v", err)
	}

	return service.TestConnection()
}

// GetGatewayCapabilities دریافت قابلیت‌های درگاه پرداخت
func (f *PaymentServiceFactory) GetGatewayCapabilities(gatewayType string) map[string]bool {
	capabilities := map[string]bool{
		"payment":      true,  // ایجاد پرداخت
		"verification": true,  // تایید پرداخت
		"refund":       false, // بازگشت وجه
		"balance":      false, // دریافت موجودی
		"sandbox":      true,  // حالت تست
	}

	switch gatewayType {
	case "zarinpal":
		capabilities["refund"] = false
		capabilities["balance"] = false
	case "saman":
		capabilities["refund"] = true
		capabilities["balance"] = false
	case "mellat":
		capabilities["refund"] = true
		capabilities["balance"] = true
	case "parsian":
		capabilities["refund"] = true
		capabilities["balance"] = false
	case "paypal":
		capabilities["refund"] = true
		capabilities["balance"] = true
	case "stripe":
		capabilities["refund"] = true
		capabilities["balance"] = true
	}

	return capabilities
}
