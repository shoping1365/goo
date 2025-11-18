package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// SetupDeviceManagementRoutes - مسیرهای مدیریت دستگاه‌ها
func SetupDeviceManagementRoutes(router *gin.Engine, db *gorm.DB) {
	deviceHandler := handlers.NewDeviceManagementHandler(db)

	// گروه API که نیاز به احراز هویت دارد
	authenticated := router.Group("/api/user")
	authenticated.Use(middleware.Auth())
	{
		// مسیرهای مدیریت دستگاه‌ها
		devices := authenticated.Group("/devices")
		{
			// GET /api/user/devices - لیست دستگاه‌های لاگین شده
			devices.GET("", deviceHandler.GetUserDevices)

			// DELETE /api/user/devices/:id - خروج از دستگاه خاص
			devices.DELETE("/:id", deviceHandler.LogoutDevice)

			// POST /api/user/devices/logout-all - خروج از همه دستگاه‌ها
			devices.POST("/logout-all", deviceHandler.LogoutAllDevices)

			// POST /api/user/devices/:id/trust - اعتماد به دستگاه
			devices.POST("/:id/trust", deviceHandler.TrustDevice)

			// DELETE /api/user/devices/:id/trust - حذف اعتماد از دستگاه
			devices.DELETE("/:id/trust", deviceHandler.UntrustDevice)
		}
	}
}
