package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupWidgetRoutes تنظیم مسیرهای ابزارک
// این تابع مسیرهای مربوط به مدیریت ابزارک‌ها را با محافظت امنیتی تعریف می‌کند
func SetupWidgetRoutes(api *gin.RouterGroup, widgetHandler *handlers.WidgetHandler) {
	// گروه ابزارک‌ها
	widgets := api.Group("/widgets")
	{
		// Debug endpoint for testing auth - بدون نیاز به permission
		widgets.GET("/debug", middleware.Auth(), func(c *gin.Context) {
			userID, exists := c.Get("user_id")
			role, roleExists := c.Get("role")
			username, usernameExists := c.Get("username")

			c.JSON(200, gin.H{
				"message":         "Debug info",
				"user_id":         userID,
				"user_id_exists":  exists,
				"role":            role,
				"role_exists":     roleExists,
				"username":        username,
				"username_exists": usernameExists,
				"headers": gin.H{
					"x_user_id":     c.GetHeader("X-User-ID"),
					"x_user_role":   c.GetHeader("X-User-Role"),
					"authorization": c.GetHeader("Authorization"),
				},
			})
		})

		// Test endpoint بدون نیاز به permission
		widgets.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message":   "Widget endpoint accessible",
				"timestamp": time.Now().Format(time.RFC3339),
			})
		})

		// CRUD عملیات اصلی
		// موقتاً بدون نیاز به permission برای تست
		widgets.POST("/", middleware.Auth(), widgetHandler.CreateWidget)      // ایجاد ابزارک جدید
		widgets.GET("/", middleware.Auth(), widgetHandler.GetWidgets)         // دریافت تمام ابزارک‌ها (با فیلتر)
		widgets.GET("/:id", middleware.Auth(), widgetHandler.GetWidget)       // دریافت ابزارک با ID (ادمین)
		widgets.GET("/public/:id", widgetHandler.GetWidget)                   // دریافت ابزارک با ID (عمومی)
		widgets.PUT("/:id", middleware.Auth(), widgetHandler.UpdateWidget)    // به‌روزرسانی ابزارک
		widgets.DELETE("/:id", middleware.Auth(), widgetHandler.DeleteWidget) // حذف ابزارک

		// عملیات خاص
		widgets.GET("/code/:code", middleware.Auth(), widgetHandler.GetWidgetByCode)     // دریافت ابزارک با کد
		widgets.POST("/order", middleware.Auth(), widgetHandler.UpdateWidgetOrder)       // به‌روزرسانی ترتیب ابزارک‌ها
		widgets.POST("/:id/toggle", middleware.Auth(), widgetHandler.ToggleWidgetStatus) // تغییر وضعیت ابزارک
		widgets.POST("/:id/duplicate", middleware.Auth(), widgetHandler.DuplicateWidget) // کپی کردن ابزارک
		widgets.GET("/page/:page", widgetHandler.GetWidgetsByPage)                       // دریافت ابزارک‌های یک صفحه (عمومی)
	}
}
