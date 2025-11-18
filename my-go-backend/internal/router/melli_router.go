package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupMelliRoutes تنظیم مسیرهای مربوط به درگاه ملی
func SetupMelliRoutes(r *gin.Engine, melliHandler *handlers.MelliHandler) {
	// گروه مسیرهای ملی
	melliGroup := r.Group("/api/payments/melli")
	{
		// مسیرهای عمومی (بدون احراز هویت)
		melliGroup.POST("/create", melliHandler.CreateMelliPayment)
		melliGroup.POST("/verify", melliHandler.VerifyMelliPayment)
		melliGroup.POST("/refund", melliHandler.RefundMelliPayment)
		melliGroup.POST("/callback", melliHandler.ProcessMelliCallback)
	}

	// گروه مسیرهای ادمین ملی (با احراز هویت)
	adminMelliGroup := r.Group("/api/admin/payment-gateways/melli")
	adminMelliGroup.Use(middleware.Auth(), middleware.Admin())
	{
		adminMelliGroup.GET("", melliHandler.GetMelliSettings)
		adminMelliGroup.PUT("", melliHandler.UpdateMelliSettings)
		adminMelliGroup.POST("/test", melliHandler.TestMelliConnection)
		adminMelliGroup.GET("/transactions", melliHandler.GetMelliTransactions)
	}
}
