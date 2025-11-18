package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupGiftCardRoutes(router *gin.Engine, handler *handlers.GiftCardHandler) {
	// گروه مسیرهای گیفت کارت
	giftCardGroup := router.Group("/api/admin/gift-cards")
	{
		// اعمال middleware احراز هویت و مجوز
		// giftCardGroup.Use(middleware.AuthMiddleware()) // احراز هویت غیرفعال شده است
		giftCardGroup.Use(middleware.RequirePermission("admin_panel_access"))

		// داشبورد و آمار
		giftCardGroup.GET("/dashboard/stats", handler.GetDashboardStats)
		giftCardGroup.GET("/dashboard/chart-data", handler.GetChartData)
		giftCardGroup.GET("/dashboard/recent-activities", handler.GetRecentActivities)

		// مدیریت گیفت کارت‌ها
		giftCardGroup.GET("", handler.GetAllGiftCards)
		giftCardGroup.POST("", handler.CreateGiftCard)
		giftCardGroup.GET("/:id", handler.GetGiftCard)
		giftCardGroup.PUT("/:id", handler.UpdateGiftCard)
		giftCardGroup.DELETE("/:id", handler.DeleteGiftCard)

		// جستجو و فیلتر
		giftCardGroup.GET("/search", handler.SearchGiftCards)
		giftCardGroup.GET("/code/:code", handler.GetGiftCardByCode)

		// عملیات روی گیفت کارت
		giftCardGroup.POST("/:id/activate", handler.ActivateGiftCard)
		giftCardGroup.POST("/:id/deactivate", handler.DeactivateGiftCard)
		giftCardGroup.POST("/use", handler.UseGiftCard)

		// تراکنش‌ها
		giftCardGroup.GET("/:id/transactions", handler.GetTransactions)
	}

	// مسیرهای عمومی (بدون احراز هویت ادمین)
	publicGiftCardGroup := router.Group("/api/gift-cards")
	{
		// بررسی گیفت کارت بر اساس کد (برای مشتریان)
		publicGiftCardGroup.GET("/check/:code", handler.GetGiftCardByCode)
	}
}
