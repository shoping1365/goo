package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupParsianRoutes تنظیم مسیرهای درگاه پرداخت پارسیان
func SetupParsianRoutes(r *gin.Engine, parsianHandler *handlers.ParsianHandler) {
	// گروه مسیرهای پارسیان
	parsianGroup := r.Group("/api/parsian")
	{
		// مسیرهای عمومی (بدون احراز هویت)
		parsianGroup.POST("/payment", parsianHandler.CreateParsianPayment)    // ایجاد پرداخت
		parsianGroup.POST("/verify", parsianHandler.VerifyParsianPayment)     // تایید پرداخت
		parsianGroup.POST("/callback", parsianHandler.ProcessParsianCallback) // پردازش callback
	}

	// گروه مسیرهای مدیریتی پارسیان
	adminParsianGroup := r.Group("/api/admin/parsian")
	adminParsianGroup.Use(middleware.Auth(), middleware.Admin())
	{
		// مدیریت تنظیمات
		adminParsianGroup.GET("/settings", parsianHandler.GetParsianSettings)            // دریافت تنظیمات
		adminParsianGroup.PUT("/settings", parsianHandler.UpdateParsianSettings)         // به‌روزرسانی تنظیمات
		adminParsianGroup.POST("/test-connection", parsianHandler.TestParsianConnection) // تست اتصال

		// مدیریت تراکنش‌ها
		adminParsianGroup.GET("/transactions", parsianHandler.GetParsianTransactions) // لیست تراکنش‌ها
		adminParsianGroup.POST("/refund", parsianHandler.RefundParsianPayment)        // بازگشت وجه
	}
}
