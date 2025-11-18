package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupPermissionRoutes تنظیم مسیرهای مربوط به مدیریت دسترسی‌ها
// این تابع مسیرهای مربوط به مدیریت مجوزها را با محافظت امنیتی تعریف می‌کند
func SetupPermissionRoutes(r *gin.Engine, db *gorm.DB) {
	permissionHandler := handlers.NewPermissionHandler(db)

	// گروه مسیرهای مربوط به دسترسی‌ها - با پیشوند /api
	permissionRoutes := r.Group("/api/permissions")
	permissionRoutes.Use(middleware.Auth())
	permissionRoutes.Use(middleware.RequirePermission("permissions_manage"))
	{
		// مسیرهای GET
		permissionRoutes.GET("", permissionHandler.ListPermissions)                  // لیست دسترسی‌ها
		permissionRoutes.GET("/:id", permissionHandler.GetPermission)                // دریافت دسترسی خاص
		permissionRoutes.GET("/modules", permissionHandler.GetPermissionModules)     // لیست ماژول‌ها
		permissionRoutes.GET("/actions", permissionHandler.GetPermissionActions)     // لیست عملیات‌ها
		permissionRoutes.GET("/resources", permissionHandler.GetPermissionResources) // لیست منابع

		// مسیرهای POST
		permissionRoutes.POST("", permissionHandler.CreatePermission)

		// مسیرهای PUT
		permissionRoutes.PUT("/:id", permissionHandler.UpdatePermission)

		// مسیرهای DELETE
		permissionRoutes.DELETE("/:id", permissionHandler.DeletePermission)
	}

	// گروه مسیرهای مربوط به دسترسی‌ها - بدون پیشوند /api (برای محیط توسعه)
	permissionRoutesNoPrefixDev := r.Group("/permissions")
	permissionRoutesNoPrefixDev.Use(middleware.Auth())
	permissionRoutesNoPrefixDev.Use(middleware.RequirePermission("permissions_manage"))
	{
		// مسیرهای GET
		permissionRoutesNoPrefixDev.GET("", permissionHandler.ListPermissions)                  // لیست دسترسی‌ها
		permissionRoutesNoPrefixDev.GET("/:id", permissionHandler.GetPermission)                // دریافت دسترسی خاص
		permissionRoutesNoPrefixDev.GET("/modules", permissionHandler.GetPermissionModules)     // لیست ماژول‌ها
		permissionRoutesNoPrefixDev.GET("/actions", permissionHandler.GetPermissionActions)     // لیست عملیات‌ها
		permissionRoutesNoPrefixDev.GET("/resources", permissionHandler.GetPermissionResources) // لیست منابع

		// مسیرهای POST
		permissionRoutesNoPrefixDev.POST("", permissionHandler.CreatePermission)

		// مسیرهای PUT
		permissionRoutesNoPrefixDev.PUT("/:id", permissionHandler.UpdatePermission)

		// مسیرهای DELETE
		permissionRoutesNoPrefixDev.DELETE("/:id", permissionHandler.DeletePermission)
	}
}
