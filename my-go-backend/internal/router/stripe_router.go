package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupStripeRoutes تنظیم مسیرهای درگاه پرداخت استرایپ
func SetupStripeRoutes(r *gin.Engine, stripeHandler *handlers.StripeHandler) {
	// گروه مسیرهای استرایپ
	stripeGroup := r.Group("/api/stripe")
	{
		// مسیرهای عمومی (بدون احراز هویت)
		stripeGroup.POST("/payment", stripeHandler.CreateStripePayment)   // ایجاد پرداخت
		stripeGroup.POST("/verify", stripeHandler.VerifyStripePayment)    // تایید پرداخت
		stripeGroup.GET("/callback", stripeHandler.ProcessStripeCallback) // پردازش callback
		stripeGroup.POST("/webhook", stripeHandler.ProcessStripeWebhook)  // پردازش webhook
	}

	// گروه مسیرهای مدیریتی استرایپ
	adminStripeGroup := r.Group("/api/admin/stripe")
	adminStripeGroup.Use(middleware.Auth(), middleware.Admin())
	{
		// مدیریت تنظیمات
		adminStripeGroup.GET("/settings", stripeHandler.GetStripeSettings)            // دریافت تنظیمات
		adminStripeGroup.PUT("/settings", stripeHandler.UpdateStripeSettings)         // به‌روزرسانی تنظیمات
		adminStripeGroup.POST("/test-connection", stripeHandler.TestStripeConnection) // تست اتصال

		// مدیریت تراکنش‌ها
		adminStripeGroup.GET("/transactions", stripeHandler.GetStripeTransactions) // لیست تراکنش‌ها
		adminStripeGroup.POST("/refund", stripeHandler.RefundStripePayment)        // بازگشت وجه
	}
}
