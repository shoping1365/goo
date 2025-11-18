package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/services"
)

// RegisterFooterRoutes ثبت مسیرهای مربوط به تنظیمات فوتر
func RegisterFooterRoutes(public *gin.RouterGroup, admin *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	// ایجاد service و handler
	settingService := services.NewSettingService(db, redis)
	footerHandler := handlers.NewFooterHandler(settingService)

	// گروه مسیرهای عمومی فوتر (بدون احراز هویت)
	footerPublic := public.Group("/footer-settings")
	{
		// دریافت تنظیمات فوتر برای نمایش در frontend
		// استفاده از middleware برای مدیریت شامل کردن و مستثنی کردن فوتر و لایه‌ها
		footerPublic.GET("", middleware.FooterLayerInclusionExclusionMiddleware(db), footerHandler.GetFooterSettings)
	}

	// گروه مسیرهای admin فوتر (نیاز به احراز هویت)
	footerAdmin := admin.Group("/footer-settings")
	{
		footerAdmin.GET("", middleware.Auth(), middleware.RequirePermission("footer_manage"), footerHandler.ListFooters)
		footerAdmin.POST("", middleware.Auth(), middleware.RequirePermission("footer_manage"), footerHandler.CreateFooter)
		footerAdmin.GET(":id", middleware.Auth(), middleware.RequirePermission("footer_manage"), footerHandler.GetFooterByID)
		footerAdmin.PUT(":id", middleware.Auth(), middleware.RequirePermission("footer_manage"), footerHandler.UpdateFooterByID)
		footerAdmin.DELETE(":id", middleware.Auth(), middleware.RequirePermission("footer_manage"), footerHandler.DeleteFooterByID)

		// مدیریت لایه‌های فوتر
		footerAdmin.GET(":id/layers", middleware.Auth(), middleware.RequirePermission("footer_manage"), footerHandler.GetFooterLayers)
		footerAdmin.PUT(":id/layers", middleware.Auth(), middleware.RequirePermission("footer_manage"), footerHandler.UpdateFooterLayers)
	}
}

