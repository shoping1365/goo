package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

/*
ShopSettingsRouter
روتر تنظیمات فروشگاه

این روتر شامل endpoint های زیر است:
- GET /api/admin/shop-settings: دریافت تنظیمات فروشگاه (نیاز به احراز هویت ادمین)
- PUT /api/admin/shop-settings: بروزرسانی تنظیمات فروشگاه (نیاز به احراز هویت ادمین)
- GET /api/shop-settings: دریافت تنظیمات عمومی فروشگاه (عمومی)
*/

// SetupShopSettingsRoutes تنظیم مسیرهای تنظیمات فروشگاه
func SetupShopSettingsRoutes(admin *gin.RouterGroup, public *gin.RouterGroup, shopSettingsHandler *handlers.ShopSettingsHandler) {
	// گروه مسیرهای ادمین
	// هم‌راستا با کتابخانه رسانه: از Auth() استفاده می‌کنیم تا هدرهای dev (X-User-*) نیز پذیرفته شوند
	admin.Use(middleware.Auth())
	admin.Use(middleware.RequirePermission("admin_panel_access"))

	{
		// دریافت تنظیمات فروشگاه (برای ادمین)
		admin.GET("/shop-settings", shopSettingsHandler.GetShopSettings)

		// بروزرسانی تنظیمات فروشگاه
		admin.PUT("/shop-settings", shopSettingsHandler.UpdateShopSettings)
	}

	// مسیر عمومی برای دریافت تنظیمات عمومی فروشگاه
	public.GET("/shop-settings", shopSettingsHandler.GetPublicShopSettings)
}
