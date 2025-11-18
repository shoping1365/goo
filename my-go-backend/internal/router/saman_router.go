package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupSamanRoutes تنظیم مسیرهای درگاه پرداخت سامان
func SetupSamanRoutes(r *gin.Engine, samanHandler *handlers.SamanHandler) {
	// گروه مسیرهای سامان
	samanGroup := r.Group("/api/saman")
	{
		// مسیرهای عمومی (بدون احراز هویت)
		samanGroup.POST("/payment", samanHandler.CreateSamanPayment)    // ایجاد پرداخت
		samanGroup.POST("/verify", samanHandler.VerifySamanPayment)     // تایید پرداخت
		samanGroup.POST("/callback", samanHandler.ProcessSamanCallback) // پردازش callback
	}

	// گروه مسیرهای مدیریتی سامان
	adminSamanGroup := r.Group("/api/admin/saman")
	adminSamanGroup.Use(middleware.Auth(), middleware.Admin())
	{
		// مدیریت تنظیمات
		adminSamanGroup.GET("/settings", samanHandler.GetSamanSettings)            // دریافت تنظیمات
		adminSamanGroup.PUT("/settings", samanHandler.UpdateSamanSettings)         // به‌روزرسانی تنظیمات
		adminSamanGroup.POST("/test-connection", samanHandler.TestSamanConnection) // تست اتصال

		// مدیریت تراکنش‌ها
		adminSamanGroup.GET("/transactions", samanHandler.GetSamanTransactions) // لیست تراکنش‌ها
		adminSamanGroup.POST("/refund", samanHandler.RefundSamanPayment)        // بازگشت وجه
	}
}
