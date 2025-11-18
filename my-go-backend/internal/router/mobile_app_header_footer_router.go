package router

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/services"
)

// RegisterMobileAppHeaderRoutes ثبت مسیرهای مربوط به تنظیمات هدر موبایل و اپلیکیشن
func RegisterMobileAppHeaderRoutes(public *gin.RouterGroup, admin *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	// ایجاد service و handler
	settingService := services.NewSettingService(db, redis)
	mobileAppHeaderHandler := handlers.NewMobileAppHeaderHandler(settingService)

	// گروه مسیرهای عمومی هدر موبایل و اپلیکیشن (بدون احراز هویت)
	mobileAppHeaderPublic := public.Group("/mobile-app-header-settings")
	{
		// دریافت تنظیمات هدر موبایل و اپلیکیشن برای نمایش در frontend
		mobileAppHeaderPublic.GET("", mobileAppHeaderHandler.GetMobileAppHeaderSettings)
	}

	// گروه مسیرهای admin هدر موبایل و اپلیکیشن (نیاز به احراز هویت)
	mobileAppHeaderAdmin := admin.Group("/mobile-app-header-settings")
	{
		mobileAppHeaderAdmin.GET("", middleware.Auth(), middleware.RequirePermission("mobile_app_header_manage"), mobileAppHeaderHandler.ListMobileAppHeaders)
		mobileAppHeaderAdmin.POST("", middleware.Auth(), middleware.RequirePermission("mobile_app_header_manage"), mobileAppHeaderHandler.CreateMobileAppHeader)
		mobileAppHeaderAdmin.GET(":id", middleware.Auth(), middleware.RequirePermission("mobile_app_header_manage"), mobileAppHeaderHandler.GetMobileAppHeaderByID)
		mobileAppHeaderAdmin.PUT(":id", middleware.Auth(), middleware.RequirePermission("mobile_app_header_manage"), mobileAppHeaderHandler.UpdateMobileAppHeaderByID)
		mobileAppHeaderAdmin.DELETE(":id", middleware.Auth(), middleware.RequirePermission("mobile_app_header_manage"), mobileAppHeaderHandler.DeleteMobileAppHeaderByID)

		// آپلود لوگو هدر موبایل و اپلیکیشن
		mobileAppHeaderAdmin.POST("upload-logo", middleware.Auth(), middleware.RequirePermission("mobile_app_header_manage"), mobileAppHeaderHandler.UploadMobileAppHeaderLogo)
	}
}

// RegisterMobileAppFooterRoutes ثبت مسیرهای مربوط به تنظیمات فوتر موبایل و اپلیکیشن
func RegisterMobileAppFooterRoutes(public *gin.RouterGroup, admin *gin.RouterGroup, db *gorm.DB, redis *redis.Client) {
	// ایجاد service و handler
	settingService := services.NewSettingService(db, redis)
	mobileAppFooterHandler := handlers.NewMobileAppFooterHandler(settingService)

	// گروه مسیرهای عمومی فوتر موبایل و اپلیکیشن (بدون احراز هویت)
	mobileAppFooterPublic := public.Group("/mobile-app-footer-settings")
	{
		// دریافت تنظیمات فوتر موبایل و اپلیکیشن برای نمایش در frontend
		mobileAppFooterPublic.GET("", mobileAppFooterHandler.GetMobileAppFooterSettings)
	}

	// گروه مسیرهای admin فوتر موبایل و اپلیکیشن (نیاز به احراز هویت)
	mobileAppFooterAdmin := admin.Group("/mobile-app-footer-settings")
	{
		mobileAppFooterAdmin.GET("", middleware.Auth(), middleware.RequirePermission("mobile_app_footer_manage"), mobileAppFooterHandler.ListMobileAppFooters)
		mobileAppFooterAdmin.POST("", middleware.Auth(), middleware.RequirePermission("mobile_app_footer_manage"), mobileAppFooterHandler.CreateMobileAppFooter)
		mobileAppFooterAdmin.GET(":id", middleware.Auth(), middleware.RequirePermission("mobile_app_footer_manage"), mobileAppFooterHandler.GetMobileAppFooterByID)
		mobileAppFooterAdmin.PUT(":id", middleware.Auth(), middleware.RequirePermission("mobile_app_footer_manage"), mobileAppFooterHandler.UpdateMobileAppFooterByID)
		mobileAppFooterAdmin.DELETE(":id", middleware.Auth(), middleware.RequirePermission("mobile_app_footer_manage"), mobileAppFooterHandler.DeleteMobileAppFooterByID)

		// مدیریت لایه‌های فوتر موبایل و اپلیکیشن
		mobileAppFooterAdmin.GET(":id/layers", middleware.Auth(), middleware.RequirePermission("mobile_app_footer_manage"), mobileAppFooterHandler.GetMobileAppFooterLayers)
		mobileAppFooterAdmin.PUT(":id/layers", middleware.Auth(), middleware.RequirePermission("mobile_app_footer_manage"), mobileAppFooterHandler.UpdateMobileAppFooterLayers)

		// آپلود لوگو فوتر موبایل و اپلیکیشن
		mobileAppFooterAdmin.POST("upload-logo", middleware.Auth(), middleware.RequirePermission("mobile_app_footer_manage"), mobileAppFooterHandler.UploadMobileAppFooterLogo)
	}
}
