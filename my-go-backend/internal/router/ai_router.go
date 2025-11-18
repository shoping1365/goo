package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupAIRoutes ثبت مسیرهای AI Content Generation
func SetupAIRoutes(router *gin.Engine, aiHandler *handlers.AIHandler) {
	adminGroup := router.Group("/api/admin")
	adminGroup.Use(middleware.AuthMiddleware())
	adminGroup.Use(middleware.RequirePermission("admin_panel_access"))

	{
		adminGroup.POST("/ai/generate-content", aiHandler.GenerateContent)
		adminGroup.POST("/ai/chat", aiHandler.HandleChat) // چت تعاملی جدید

		// مسیرهای مدیریت جلسات چت
		adminGroup.GET("/ai/sessions", aiHandler.GetUserSessions)                    // دریافت همه جلسات کاربر
		adminGroup.GET("/ai/sessions/:session_id/history", aiHandler.GetChatHistory) // دریافت تاریخچه یک جلسه
	}
}
