package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRolePermissionRoutes تنظیم مسیرهای مربوط به مدیریت دسترسی‌های نقش‌ها
// این روتر برای مدیریت تخصیص دسترسی‌ها به نقش‌ها استفاده می‌شود
func SetupRolePermissionRoutes(r *gin.Engine, db *gorm.DB) {
	// ایجاد هندلر
	rolePermissionHandler := handlers.NewRolePermissionHandler(db)

	// گروه مسیرهای مربوط به دسترسی‌های نقش‌ها
	rolePermissionGroup := r.Group("/api/admin/role-permissions")
	{
		// اعمال middleware احراز هویت و دسترسی
		rolePermissionGroup.Use(middleware.Auth())
		rolePermissionGroup.Use(middleware.RequirePermission("role_permissions_view"))

		// دریافت تمام دسترسی‌های موجود (برای نمایش در UI)
		rolePermissionGroup.GET("/permissions", rolePermissionHandler.GetAllPermissions)

		// دریافت دسترسی‌های یک نقش خاص
		rolePermissionGroup.GET("/roles/:id", rolePermissionHandler.GetRolePermissions)

		// تخصیص یک دسترسی به نقش
		rolePermissionGroup.POST("/assign", middleware.RequirePermission("role_permissions_write"), rolePermissionHandler.AssignPermission)

		// تخصیص گروهی دسترسی‌ها به نقش
		rolePermissionGroup.POST("/bulk-assign", middleware.RequirePermission("role_permissions_write"), rolePermissionHandler.BulkAssignPermissions)

		// لغو دسترسی از نقش
		rolePermissionGroup.DELETE("/roles/:roleId/permissions/:permissionId", middleware.RequirePermission("role_permissions_write"), rolePermissionHandler.RevokePermission)

		// بررسی دسترسی کاربر
		rolePermissionGroup.GET("/users/:userId/check", rolePermissionHandler.CheckUserPermission)
	}
}
