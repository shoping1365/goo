package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupMellatRoutes تنظیم مسیرهای مربوط به درگاه ملت
func SetupMellatRoutes(r *gin.Engine, mellatHandler *handlers.MellatHandler) {
	// گروه مسیرهای ملت
	mellatGroup := r.Group("/api/payments/mellat")
	{
		// مسیرهای عمومی (بدون احراز هویت)
		mellatGroup.POST("/create", mellatHandler.CreateMellatPayment)
		mellatGroup.POST("/verify", mellatHandler.VerifyMellatPayment)
		mellatGroup.POST("/refund", mellatHandler.RefundMellatPayment)
		mellatGroup.POST("/callback", mellatHandler.ProcessMellatCallback)
	}

	// گروه مسیرهای ادمین ملت (با احراز هویت)
	adminMellatGroup := r.Group("/api/admin/payment-gateways/mellat")
	adminMellatGroup.Use(middleware.Auth(), middleware.Admin())
	{
		adminMellatGroup.GET("", mellatHandler.GetMellatSettings)
		adminMellatGroup.PUT("", mellatHandler.UpdateMellatSettings)
		adminMellatGroup.POST("/test", mellatHandler.TestMellatConnection)
		adminMellatGroup.GET("/transactions", mellatHandler.GetMellatTransactions)
	}
}
