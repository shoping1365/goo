package router

import (
	"my-go-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterPaymentGatewayRoutes مسیرهای مربوط به مدیریت درگاه‌های پرداخت را ثبت می‌کند
func RegisterPaymentGatewayRoutes(r *gin.Engine, handler *handlers.PaymentGatewayHandler) {
	api := r.Group("/api/payment-gateways")
	{
		// عملیات CRUD اصلی
		api.POST("", handler.Create)              // ثبت درگاه جدید
		api.GET("", handler.GetAll)               // لیست همه درگاه‌ها
		api.GET("/active", handler.GetActive)     // لیست درگاه‌های فعال
		api.GET("/filters", handler.GetByFilters) // لیست با فیلترهای پیشرفته
		api.GET("/:id", handler.GetByID)          // دریافت یک درگاه
		api.PUT("/:id", handler.Update)           // ویرایش درگاه
		api.DELETE("/:id", handler.Delete)        // حذف درگاه

		// عملیات وضعیت
		api.PUT("/:id/status", handler.UpdateStatus) // تغییر وضعیت درگاه

		// آمار و گزارشات
		api.GET("/stats", handler.GetStats)               // آمار کلی
		api.GET("/top", handler.GetTopGateways)           // درگاه‌های برتر
		api.POST("/reset-stats", handler.ResetTodayStats) // ریست آمار امروز

		// تست و بررسی
		api.POST("/:id/test-connection", handler.TestConnection) // تست اتصال
	}
}
