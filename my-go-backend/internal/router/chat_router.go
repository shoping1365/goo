package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupChatRoutes تنظیم مسیرهای چت
func RegisterChatRoutes(router *gin.Engine, chatHandler *handlers.ChatHandler, wsManager *services.WebSocketManager) {
	// گروه مسیرهای چت
	chatGroup := router.Group("/api/chat")
	chatGroup.Use(middleware.RateLimit())
	{
		// مسیرهای عمومی (بدون احراز هویت)
		chatGroup.POST("/sessions", chatHandler.CreateChatSession)
		chatGroup.GET("/sessions/:session_id", chatHandler.GetChatSession)
		chatGroup.POST("/sessions/:session_id/messages", chatHandler.SendMessage)
		chatGroup.GET("/sessions/:session_id/messages", chatHandler.GetSessionMessages)
		chatGroup.PUT("/messages/:message_id/read", chatHandler.MarkMessageAsRead)
		chatGroup.PUT("/sessions/:session_id/end", chatHandler.EndChatSession)
		chatGroup.GET("/knowledge-base/search", chatHandler.SearchKnowledgeBase)

		// WebSocket endpoint
		chatGroup.GET("/ws", func(c *gin.Context) {
			wsManager.HandleWebSocket(c.Writer, c.Request)
		})

		// گروه مسیرهای ادمین (با احراز هویت)
		adminChatGroup := chatGroup.Group("/admin")
		adminChatGroup.Use(middleware.AuthMiddleware())
		adminChatGroup.Use(middleware.AdminMiddleware())
		{
			// لیست جلسات بر اساس وضعیت
			adminChatGroup.GET("/sessions", chatHandler.ListSessions)

			// مدیریت جلسه‌های چت
			adminChatGroup.GET("/sessions/waiting", chatHandler.GetWaitingSessions)
			adminChatGroup.PUT("/sessions/:session_id/assign", chatHandler.AssignOperator)
			adminChatGroup.PUT("/sessions/:session_id/read", chatHandler.MarkSessionRead)

			// مدیریت اپراتورها
			adminChatGroup.GET("/operators/available", chatHandler.GetAvailableOperators)
			adminChatGroup.PUT("/operators/:operator_id/status", chatHandler.UpdateOperatorStatus)
			adminChatGroup.POST("/operators/invite", chatHandler.InviteOperator)

			// آمار چت
			adminChatGroup.GET("/statistics", chatHandler.GetChatStatistics)

			// تنظیمات نرخ پیام
			adminChatGroup.GET("/settings/rate-limit", chatHandler.GetChatRateLimitSettings)
			adminChatGroup.PUT("/settings/rate-limit", chatHandler.UpdateChatRateLimitSettings)

			// متریک‌های چت
			adminChatGroup.GET("/metrics", chatHandler.GetChatMetrics)

			// مدیریت ابزارک‌های چت
			adminChatGroup.POST("/widgets", chatHandler.CreateChatWidget)
			adminChatGroup.GET("/widgets", chatHandler.GetActiveChatWidgets)
			adminChatGroup.PUT("/widgets/:id", chatHandler.UpdateChatWidget)
			adminChatGroup.DELETE("/widgets/:id", chatHandler.DeleteChatWidget)

			// مدیریت بات‌های هوش مصنوعی
			adminChatGroup.POST("/ai-bots", chatHandler.CreateAIBot)
			adminChatGroup.GET("/ai-bots", chatHandler.GetActiveAIBots)

			// مدیریت پایگاه دانش
			adminChatGroup.GET("/knowledge-base", chatHandler.ListKnowledgeBase)
			adminChatGroup.POST("/knowledge-base", chatHandler.CreateKnowledgeBase)
			adminChatGroup.GET("/knowledge-base/:id", chatHandler.GetKnowledgeBase)
			adminChatGroup.PUT("/knowledge-base/:id", chatHandler.UpdateKnowledgeBase)
			adminChatGroup.DELETE("/knowledge-base/:id", chatHandler.DeleteKnowledgeBase)
		}
	}

	// گروه مسیرهای اپراتور (با احراز هویت)
	operatorChatGroup := router.Group("/api/chat/operator")
	operatorChatGroup.Use(middleware.AuthMiddleware())
	operatorChatGroup.Use(middleware.OperatorMiddleware())
	{
		// اپراتورها می‌توانند جلسه‌های در انتظار را ببینند
		operatorChatGroup.GET("/sessions/waiting", chatHandler.GetWaitingSessions)

		// اپراتورها می‌توانند خودشان را به جلسه تخصیص دهند
		operatorChatGroup.PUT("/sessions/:session_id/assign-self", chatHandler.AssignOperator)

		// اپراتورها می‌توانند وضعیت خود را تغییر دهند
		operatorChatGroup.PUT("/status", chatHandler.UpdateOperatorStatus)

		// اپراتورها می‌توانند آمار خود را ببینند
		operatorChatGroup.GET("/statistics", chatHandler.GetChatStatistics)
	}
}

// SetupChatAdminRoutes تنظیم مسیرهای ادمین چت (جداسازی بهتر)
func SetupChatAdminRoutes(router *gin.Engine, chatHandler *handlers.ChatHandler) {
	// گروه مسیرهای ادمین چت
	adminGroup := router.Group("/api/admin/chat")
	adminGroup.Use(middleware.AuthMiddleware())
	adminGroup.Use(middleware.AdminMiddleware())
	{
		// لیست جلسات بر اساس وضعیت
		adminGroup.GET("/sessions", chatHandler.ListSessions)
		// مدیریت اپراتورها
		operatorsGroup := adminGroup.Group("/operators")
		{
			operatorsGroup.GET("/", chatHandler.GetAvailableOperators)
			operatorsGroup.PUT("/:operator_id/status", chatHandler.UpdateOperatorStatus)
		}

		// مدیریت جلسه‌ها
		sessionsGroup := adminGroup.Group("/sessions")
		{
			sessionsGroup.GET("/waiting", chatHandler.GetWaitingSessions)
			sessionsGroup.PUT("/:session_id/assign", chatHandler.AssignOperator)
			sessionsGroup.PUT("/:session_id/end", chatHandler.EndChatSession)
			sessionsGroup.PUT("/:session_id/read", chatHandler.MarkSessionRead)
		}

		// مدیریت ابزارک‌ها
		widgetsGroup := adminGroup.Group("/widgets")
		{
			widgetsGroup.POST("/", chatHandler.CreateChatWidget)
			widgetsGroup.GET("/", chatHandler.GetActiveChatWidgets)
		}

		// مدیریت بات‌های AI
		aiBotsGroup := adminGroup.Group("/ai-bots")
		{
			aiBotsGroup.POST("/", chatHandler.CreateAIBot)
			aiBotsGroup.GET("/", chatHandler.GetActiveAIBots)
		}

		// مدیریت پایگاه دانش
		knowledgeBaseGroup := adminGroup.Group("/knowledge-base")
		{
			knowledgeBaseGroup.POST("/", chatHandler.CreateKnowledgeBase)
			knowledgeBaseGroup.GET("/search", chatHandler.SearchKnowledgeBase)
		}

		// آمار و گزارشات
		adminGroup.GET("/statistics", chatHandler.GetChatStatistics)
	}
}

// SetupChatCustomerRoutes تنظیم مسیرهای مشتری چت
func SetupChatCustomerRoutes(router *gin.Engine, chatHandler *handlers.ChatHandler) {
	// گروه مسیرهای مشتری چت
	customerGroup := router.Group("/api/customer/chat")
	{
		// ایجاد جلسه چت جدید
		customerGroup.POST("/sessions", chatHandler.CreateChatSession)

		// دریافت اطلاعات جلسه
		customerGroup.GET("/sessions/:session_id", chatHandler.GetChatSession)

		// ارسال و دریافت پیام‌ها
		customerGroup.POST("/sessions/:session_id/messages", chatHandler.SendMessage)
		customerGroup.GET("/sessions/:session_id/messages", chatHandler.GetSessionMessages)

		// علامت‌گذاری پیام‌ها به عنوان خوانده شده
		customerGroup.PUT("/messages/:message_id/read", chatHandler.MarkMessageAsRead)

		// پایان دادن به جلسه
		customerGroup.PUT("/sessions/:session_id/end", chatHandler.EndChatSession)

		// جستجو در پایگاه دانش
		customerGroup.GET("/knowledge-base/search", chatHandler.SearchKnowledgeBase)
	}
}

// SetupChatWebSocketRoutes تنظیم مسیرهای WebSocket چت
func SetupChatWebSocketRoutes(router *gin.Engine, wsManager *services.WebSocketManager) {
	// مسیر WebSocket برای مشتریان
	router.GET("/api/chat/ws", func(c *gin.Context) {
		wsManager.HandleWebSocket(c.Writer, c.Request)
	})

	// مسیر WebSocket برای اپراتورها (با احراز هویت)
	operatorWSGroup := router.Group("/api/chat/operator/ws")
	operatorWSGroup.Use(middleware.AuthMiddleware())
	operatorWSGroup.Use(middleware.OperatorMiddleware())
	{
		operatorWSGroup.GET("/", func(c *gin.Context) {
			wsManager.HandleWebSocket(c.Writer, c.Request)
		})
	}

	// مسیر WebSocket برای ادمین‌ها (با احراز هویت)
	adminWSGroup := router.Group("/api/chat/admin/ws")
	adminWSGroup.Use(middleware.AuthMiddleware())
	adminWSGroup.Use(middleware.AdminMiddleware())
	{
		adminWSGroup.GET("/", func(c *gin.Context) {
			wsManager.HandleWebSocket(c.Writer, c.Request)
		})
	}
}
