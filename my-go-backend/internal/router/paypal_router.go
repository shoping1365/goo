package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupPayPalRoutes تنظیم مسیرهای درگاه پرداخت پی‌پال
func SetupPayPalRoutes(r *gin.Engine, paypalHandler *handlers.PayPalHandler) {
	// گروه مسیرهای پی‌پال
	paypalGroup := r.Group("/api/paypal")
	{
		// مسیرهای عمومی (بدون احراز هویت)
		paypalGroup.POST("/payment", paypalHandler.CreatePayPalPayment)   // ایجاد پرداخت
		paypalGroup.POST("/verify", paypalHandler.VerifyPayPalPayment)    // تایید پرداخت
		paypalGroup.GET("/callback", paypalHandler.ProcessPayPalCallback) // پردازش callback
	}

	// گروه مسیرهای مدیریتی پی‌پال
	adminPayPalGroup := r.Group("/api/admin/paypal")
	adminPayPalGroup.Use(middleware.Auth(), middleware.Admin())
	{
		// مدیریت تنظیمات
		adminPayPalGroup.GET("/settings", paypalHandler.GetPayPalSettings)            // دریافت تنظیمات
		adminPayPalGroup.PUT("/settings", paypalHandler.UpdatePayPalSettings)         // به‌روزرسانی تنظیمات
		adminPayPalGroup.POST("/test-connection", paypalHandler.TestPayPalConnection) // تست اتصال

		// مدیریت تراکنش‌ها
		adminPayPalGroup.GET("/transactions", paypalHandler.GetPayPalTransactions) // لیست تراکنش‌ها
		adminPayPalGroup.POST("/refund", paypalHandler.RefundPayPalPayment)        // بازگشت وجه
	}
}
