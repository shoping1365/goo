package router

import (
	"my-go-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

/*
APISettingsRouter
روتر تنظیمات API

این روتر شامل endpoint های زیر است:
- GET /api/admin/api-settings: دریافت تنظیمات API (نیاز به احراز هویت ادمین)
- PUT /api/admin/api-settings: بروزرسانی تنظیمات API (نیاز به احراز هویت ادمین)
- POST /api/admin/api-settings/test-openai: تست اتصال OpenAI (نیاز به احراز هویت ادمین)
*/

// SetupAPISettingsRoutes تنظیم مسیرهای تنظیمات API
func SetupAPISettingsRoutes(admin *gin.RouterGroup, apiSettingsHandler *handlers.APISettingsHandler) {
	// گروه مسیرهای ادمین (احراز هویت غیرفعال شده است)
	// admin.Use(middleware.AuthMiddleware())
	// admin.Use(middleware.RequirePermission("admin_panel_access"))

	{
		// دریافت تنظیمات API (برای ادمین)
		admin.GET("/api-settings", apiSettingsHandler.GetAPISettings)

		// بروزرسانی تنظیمات API
		admin.PUT("/api-settings", apiSettingsHandler.UpdateAPISettings)

		// تست اتصال OpenAI
		admin.POST("/api-settings/test-openai", apiSettingsHandler.TestOpenAIConnection)

		// دریافت آمار استفاده OpenAI
		admin.POST("/api-settings/fetch-usage", apiSettingsHandler.FetchOpenAIUsageData)
	}
}
