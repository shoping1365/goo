package router

import (
	"github.com/gin-gonic/gin"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterUserRoutes defines public user-related endpoints (registration, list, heartbeat, delete).
func RegisterUserRoutes(public *gin.RouterGroup, userHandler *handlers.UserHandler) {
	users := public.Group("/users")
	{
		// GET /api/users  (authenticated + permission-based, developer always allowed)
		users.GET("", middleware.Auth(), middleware.RequireAnyPermission("user.view", "users_view"), userHandler.GetUsers)

		// POST /api/users/register  - customer registration
		users.POST("/register", userHandler.RegisterUser)

		// PUT /api/users/heartbeat  - update last seen (authenticated)
		users.PUT("/heartbeat", middleware.Auth(), userHandler.UpdateLastSeen)

		// Test endpoint
		users.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "User router is working", "path": c.Request.URL.Path})
		})

		// DELETE /api/users/:id  - delete user (authenticated + permission-based)
		users.DELETE(":id", middleware.Auth(), middleware.RequireAnyPermission("user.delete", "users_delete"), userHandler.DeleteUser)

		// User detail read-only endpoints (authenticated + permission-based)
		users.GET(":id", middleware.Auth(), middleware.RequireAnyPermission("user.view", "users_view"), userHandler.GetUserByID)
		users.GET(":id/addresses", middleware.Auth(), middleware.RequireAnyPermission("user.view", "users_view"), userHandler.GetUserAddresses)
		users.GET(":id/bank-accounts", middleware.Auth(), middleware.RequireAnyPermission("user.view", "users_view"), userHandler.GetUserBankAccounts)
		users.GET(":id/wallet/summary", middleware.Auth(), middleware.RequireAnyPermission("user.view", "users_view"), userHandler.GetUserWalletSummary)
		users.GET(":id/wallet/transactions", middleware.Auth(), middleware.RequireAnyPermission("user.view", "users_view"), userHandler.GetUserWalletTransactions)
		users.GET(":id/notifications", middleware.Auth(), middleware.RequireAnyPermission("user.view", "users_view"), userHandler.GetUserNotifications)

		// Admin actions
		users.POST(":id/role", middleware.Auth(), middleware.RequireAnyPermission("user.update", "users_update"), userHandler.AdminSetUserRole)
		users.POST(":id/block", middleware.Auth(), middleware.RequireAnyPermission("user.update", "users_update"), userHandler.AdminBlockUser)
		users.POST(":id/unblock", middleware.Auth(), middleware.RequireAnyPermission("user.update", "users_update"), userHandler.AdminUnblockUser)
		users.PUT(":id", middleware.Auth(), middleware.RequireAnyPermission("user.update", "users_update"), userHandler.UpdateUserBasicInfo)
		users.POST(":id/force-logout", middleware.Auth(), middleware.RequireAnyPermission("user.update", "users_update"), userHandler.AdminForceLogoutUser)
		users.POST(":id/reset-password", middleware.Auth(), middleware.RequireAnyPermission("user.update", "users_update"), userHandler.AdminResetUserPassword)
		users.POST(":id/send-warning", middleware.Auth(), middleware.RequireAnyPermission("user.update", "users_update"), userHandler.AdminSendUserWarning)
	}

	// Auth endpoints (یکپارچه: username یا mobile + password)
	auth := public.Group("/auth")
	{
		// POST /api/auth/login - ورود یکپارچه با username یا mobile + password
		auth.POST("/login", userHandler.Login)
	}
}
