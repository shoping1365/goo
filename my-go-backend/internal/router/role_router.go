package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoleRoutes تنظیم مسیرهای مربوط به مدیریت نقش‌ها
func SetupRoleRoutes(r *gin.Engine, db *gorm.DB) {
	roleHandler := handlers.NewRoleHandler(db)

	// گروه مسیرهای مربوط به نقش‌ها - با پیشوند /api
	roleRoutes := r.Group("/api/roles")
	{
		// اعمال middleware احراز هویت و دسترسی
		roleRoutes.Use(middleware.Auth())
		roleRoutes.Use(middleware.RequirePermission("role.read"))

		// مسیرهای GET
		roleRoutes.GET("", roleHandler.ListRoles)                          // لیست نقش‌ها
		roleRoutes.GET("/statistics", roleHandler.GetRoleStatistics)       // آمار کلی نقش‌ها
		roleRoutes.GET("/:id", roleHandler.GetRole)                        // دریافت نقش خاص
		roleRoutes.GET("/:id/permissions", roleHandler.GetRolePermissions) // دریافت دسترسی‌های نقش

		// مسیرهای POST (نیاز به دسترسی create)
		roleRoutes.POST("", middleware.RequirePermission("role.create"), roleHandler.CreateRole)

		// مسیرهای PUT (نیاز به دسترسی update)
		roleRoutes.PUT("/:id", middleware.RequirePermission("role.update"), roleHandler.UpdateRole)
		roleRoutes.PUT("/:id/permissions", middleware.RequirePermission("role.update"), roleHandler.UpdateRolePermissions) // به‌روزرسانی مجوزهای نقش

		// مسیرهای DELETE (نیاز به دسترسی delete)
		roleRoutes.DELETE("/:id", middleware.RequirePermission("role.delete"), roleHandler.DeleteRole)
	}

	// گروه مسیرهای مربوط به نقش‌ها - بدون پیشوند /api (برای محیط توسعه)
	roleRoutesNoPrefixDev := r.Group("/roles")
	{
		// اعمال middleware احراز هویت و دسترسی
		roleRoutesNoPrefixDev.Use(middleware.Auth())
		roleRoutesNoPrefixDev.Use(middleware.RequirePermission("role.read"))

		// مسیرهای GET
		roleRoutesNoPrefixDev.GET("", roleHandler.ListRoles)                          // لیست نقش‌ها
		roleRoutesNoPrefixDev.GET("/statistics", roleHandler.GetRoleStatistics)       // آمار کلی نقش‌ها
		roleRoutesNoPrefixDev.GET("/:id", roleHandler.GetRole)                        // دریافت نقش خاص
		roleRoutesNoPrefixDev.GET("/:id/permissions", roleHandler.GetRolePermissions) // دریافت دسترسی‌های نقش

		// مسیرهای POST (نیاز به دسترسی create)
		roleRoutesNoPrefixDev.POST("", middleware.RequirePermission("role.create"), roleHandler.CreateRole)

		// مسیرهای PUT (نیاز به دسترسی update)
		roleRoutesNoPrefixDev.PUT("/:id", middleware.RequirePermission("role.update"), roleHandler.UpdateRole)
		roleRoutesNoPrefixDev.PUT("/:id/permissions", middleware.RequirePermission("role.update"), roleHandler.UpdateRolePermissions) // به‌روزرسانی مجوزهای نقش

		// مسیرهای DELETE (نیاز به دسترسی delete)
		roleRoutesNoPrefixDev.DELETE("/:id", middleware.RequirePermission("role.delete"), roleHandler.DeleteRole)
	}
}
