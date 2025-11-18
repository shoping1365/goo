package router

import (
	"my-go-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterSMSGatewayRoutes مسیرهای مربوط به مدیریت درگاه پیامک را ثبت می‌کند
func RegisterSMSGatewayRoutes(r *gin.Engine, handler *handlers.SMSGatewayHandler) {
	// اضافه کردن مسیرهای اضافی که در admin_router تعریف نشده‌اند
	r.GET("/api/sms-gateways/messages", handler.GetSMSMessages) // دریافت پیامک‌های ارسالی
	r.GET("/api/sms-gateways/:id/outbox", handler.GetOutbox)    // دریافت پیامک‌های ارسالی (outbox) از IPPanel
	r.GET("/api/ippanel-outbox", handlers.GetIPPanelOutbox)     // دریافت پیامک‌های خروجی IPPanel

	// مسیرهای بدون پیشوند /api برای سازگاری
	r.GET("/sms-gateways", handler.List)                          // سازگاری با درخواست‌های بدون /api
	r.GET("/sms-gateways/:id", handler.GetByID)                   // دریافت درگاه بدون /api
	r.GET("/sms-gateways/:id/balance", handler.GetGatewayBalance) // موجودی درگاه بدون /api
	r.POST("/sms-gateways", handler.Create)                       // ایجاد درگاه بدون /api
	// تست اتصال و ارسال تستی بدون پیشوند
	r.POST("/sms-gateways/:id/test-connection", handler.TestConnection)
	r.POST("/sms-gateways/:id/send-test", handler.SendTestSMS) // ارسال پیامک تستی بدون /api
	r.PUT("/sms-gateways/:id", handler.Update)                 // ویرایش درگاه بدون /api
	r.DELETE("/sms-gateways/:id", handler.Delete)              // حذف درگاه بدون /api
}
