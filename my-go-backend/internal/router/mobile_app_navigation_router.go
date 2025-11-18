package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/services"
)

// RegisterMobileAppNavigationRoutes ثبت مسیرهای مربوط به ناوبری موبایل و اپلیکیشن
func RegisterMobileAppNavigationRoutes(public *gin.RouterGroup, admin *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	// ایجاد service و handler
	settingService := services.NewSettingService(db, redis)
	navigationHandler := handlers.NewMobileAppNavigationHandler(settingService)

	// گروه مسیرهای عمومی ناوبری (بدون احراز هویت)
	navigationPublic := public.Group("/mobile-app-navigation-settings")
	{
		// دریافت تنظیمات ناوبری برای نمایش در frontend
		navigationPublic.GET("", navigationHandler.GetMobileAppNavigationSettings)
	}

	// گروه مسیرهای admin ناوبری (نیاز به احراز هویت)
	navigationAdmin := admin.Group("/mobile-app-navigation-settings")
	{
		navigationAdmin.GET("", middleware.Auth(),
			middleware.RequirePermission("mobile_app_navigation"),
			navigationHandler.ListMobileAppNavigations)

		navigationAdmin.POST("", middleware.Auth(),
			middleware.RequirePermission("mobile_app_navigation"),
			navigationHandler.CreateMobileAppNavigation)

		navigationAdmin.GET("/:id", middleware.Auth(),
			middleware.RequirePermission("mobile_app_navigation"),
			navigationHandler.GetMobileAppNavigationByID)

		navigationAdmin.PUT("/:id", middleware.Auth(),
			middleware.RequirePermission("mobile_app_navigation"),
			navigationHandler.UpdateMobileAppNavigationByID)

		navigationAdmin.DELETE("/:id", middleware.Auth(),
			middleware.RequirePermission("mobile_app_navigation"),
			navigationHandler.DeleteMobileAppNavigationByID)

		// مسیرهای مربوط به لایه‌ها
		navigationAdmin.GET("/:id/layers", middleware.Auth(),
			middleware.RequirePermission("mobile_app_navigation"),
			navigationHandler.GetMobileAppNavigationLayers)

		navigationAdmin.PUT("/:id/layers", middleware.Auth(),
			middleware.RequirePermission("mobile_app_navigation"),
			navigationHandler.UpdateMobileAppNavigationLayers)

		// مسیر آپلود لوگو
		navigationAdmin.POST("/upload-logo", middleware.Auth(),
			middleware.RequirePermission("mobile_app_navigation"),
			navigationHandler.UploadMobileAppNavigationLogo)
	}
}
