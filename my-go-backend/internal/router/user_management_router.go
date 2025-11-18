package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserManagementRoutes(r *gin.Engine, userManagementHandler *handlers.UserManagementHandler) {
	admin := r.Group("/api/admin")
	// احراز هویت مبتنی بر JWT/Sidebase و بررسی نقش ادمین با سیاست 404
	admin.Use(middleware.Auth(), middleware.Admin())
	{
		// مدیریت کاربران
		users := admin.Group("/users")
		{
			users.GET("", userManagementHandler.GetUsers)
			users.GET("/:id", userManagementHandler.GetUser)
			users.POST("/:id/block", userManagementHandler.BlockUser)
			users.POST("/:id/unblock", userManagementHandler.UnblockUser)
			users.GET("/:id/login-history", userManagementHandler.GetUserLoginHistory)
			users.GET("/:id/analytics", userManagementHandler.GetUserAnalytics)
		}

		// تاریخچه ورود و آمار (CQRS - مسیرهای خواندنی)
		admin.GET("/login-history", middleware.RequirePermission("security_login.read"), userManagementHandler.GetLoginHistory)
		admin.GET("/login-stats", middleware.RequirePermission("security_login.read"), userManagementHandler.GetLoginStats)
	}
}
