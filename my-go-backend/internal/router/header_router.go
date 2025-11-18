package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/services"
)

// RegisterHeaderRoutes ثبت مسیرهای مربوط به تنظیمات هدر
func RegisterHeaderRoutes(public *gin.RouterGroup, admin *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	// ایجاد service و handler
	settingService := services.NewSettingService(db, redis)
	headerHandler := handlers.NewHeaderHandler(settingService)

	// گروه مسیرهای عمومی هدر (بدون احراز هویت)
	headerPublic := public.Group("/header-settings")
	{
		// دریافت تنظیمات هدر برای نمایش در frontend
		// استفاده از middleware برای مدیریت شامل کردن و مستثنی کردن هدر و لایه‌ها
		headerPublic.GET("", middleware.HeaderLayerInclusionExclusionMiddleware(db), headerHandler.GetHeaderSettings)
	}

	// گروه مسیرهای admin هدر (نیاز به احراز هویت)
	headerAdmin := admin.Group("/header-settings")
	{
		headerAdmin.GET("", middleware.Auth(), middleware.RequirePermission("header_manage"), headerHandler.ListHeaders)
		headerAdmin.POST("", middleware.Auth(), middleware.RequirePermission("header_manage"), headerHandler.CreateHeader)
		headerAdmin.GET(":id", middleware.Auth(), middleware.RequirePermission("header_manage"), headerHandler.GetHeaderByID)
		headerAdmin.PUT(":id", middleware.Auth(), middleware.RequirePermission("header_manage"), headerHandler.UpdateHeaderByID)
		headerAdmin.DELETE(":id", middleware.Auth(), middleware.RequirePermission("header_manage"), headerHandler.DeleteHeaderByID)

		// مدیریت لایه‌های هدر
		headerAdmin.GET(":id/layers", middleware.Auth(), middleware.RequirePermission("header_manage"), headerHandler.GetHeaderLayers)
		headerAdmin.PUT(":id/layers", middleware.Auth(), middleware.RequirePermission("header_manage"), headerHandler.UpdateHeaderLayers)
	}
}
