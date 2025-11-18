package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// PaymentRouter روتر عملیات پرداخت
// شامل مسیرهای ایجاد، تایید و مدیریت تراکنش‌های پرداخت

func SetupPaymentRoutes(r *gin.Engine, paymentHandler *handlers.PaymentHandler, mellatHandler *handlers.MellatHandler) {
	// گروه مسیرهای پرداخت
	paymentGroup := r.Group("/api/payments")
	{
		// مسیرهای عمومی (بدون احراز هویت)
		paymentGroup.POST("/create", paymentHandler.CreatePayment)                // ایجاد پرداخت جدید
		paymentGroup.POST("/verify", paymentHandler.VerifyPayment)                // تایید پرداخت
		paymentGroup.GET("/transaction/:id", paymentHandler.GetTransaction)       // دریافت اطلاعات تراکنش
		paymentGroup.POST("/calculate", paymentHandler.CalculatePaymentBreakdown) // محاسبه تفصیلی مبلغ و کارمزد

		// مسیرهای مدیریتی (نیاز به احراز هویت ادمین)
		adminGroup := paymentGroup.Group("/admin")
		adminGroup.Use(middleware.Auth(), middleware.Admin())
		{
			adminGroup.GET("/transactions", paymentHandler.GetTransactions)           // لیست تراکنش‌ها
			adminGroup.GET("/gateway/:id/test", paymentHandler.TestGatewayConnection) // تست اتصال درگاه
		}

		// مسیرهای قابلیت‌های درگاه‌ها (جداگانه برای جلوگیری از تداخل)
		capabilitiesGroup := paymentGroup.Group("/gateway-capabilities")
		capabilitiesGroup.Use(middleware.Auth(), middleware.Admin())
		{
			capabilitiesGroup.GET("/:type", paymentHandler.GetGatewayCapabilities) // قابلیت‌های درگاه
		}
	}

	// مسیرهای callback درگاه‌های پرداخت
	callbackGroup := r.Group("/api/payment-callbacks")
	{
		callbackGroup.POST("/zarinpal", paymentHandler.VerifyPayment) // callback زرین‌پال
		// TODO: اضافه کردن callback سایر درگاه‌ها
		// callbackGroup.POST("/mellat", paymentHandler.VerifyPayment)
		// callbackGroup.POST("/parsian", paymentHandler.VerifyPayment)
		// callbackGroup.POST("/paypal", paymentHandler.VerifyPayment)
		// callbackGroup.POST("/stripe", paymentHandler.VerifyPayment)
	}

	// Setup Mellat routes
	SetupMellatRoutes(r, mellatHandler)
}
