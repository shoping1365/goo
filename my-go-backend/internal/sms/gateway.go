package sms

import (
	"my-go-backend/internal/models"
)

// SmsGateway رابط مشترک بین تمام درگاه‌های پیامکی
// هر درگاه باید این متدها را پیاده‌سازی کند تا بتوان آن را در GatewayManager ثبت کرد.
// تمام کامنت‌ها به فارسی جهت رعایت قانون مستندسازی

type SmsGateway interface {
	// Send ارسال پیام با پترن مشخص
	Send(pattern models.SMSPattern, req models.SMSSendRequest) (*models.SMSSendResponse, error)

	// HealthCheck بررسی وضعیت سلامت درگاه (اعتبارسنجی کلید، موجودی و ...)
	HealthCheck() error

	// Record شی مدل مرتبط با درگاه جهت دسترسی به فیلدها (ID،‌ اولویت و ...)
	Record() models.SMSGateway
}
