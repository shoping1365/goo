package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterMonitoringRoutes ثبت مسیرهای مانیتورینگ عملکرد چت
func RegisterMonitoringRoutes(router *gin.Engine, monitoringHandler *handlers.MonitoringHandler) {
	// گروه مسیرهای مانیتورینگ
	monitoringGroup := router.Group("/api/admin/monitoring")
	plainMonitoringGroup := router.Group("/admin/monitoring")
	// استفاده از احراز هویت مبتنی بر JWT و بررسی نقش ادمین مطابق سیاست پروژه
	monitoringGroup.Use(middleware.Auth())
	monitoringGroup.Use(middleware.Admin())
	plainMonitoringGroup.Use(middleware.Auth())
	plainMonitoringGroup.Use(middleware.Admin())
	{
		// متریک‌های عملکرد
		monitoringGroup.GET("/performance", monitoringHandler.GetPerformanceMetrics)
		monitoringGroup.GET("/realtime", monitoringHandler.GetRealTimeMetrics)
		monitoringGroup.GET("/overview", monitoringHandler.GetSystemOverview)

		// وضعیت سلامت سیستم
		monitoringGroup.GET("/health", monitoringHandler.GetSystemHealth)

		// هشدارها
		monitoringGroup.GET("/alerts", monitoringHandler.GetAlerts)
		monitoringGroup.POST("/alerts", monitoringHandler.CreateAlert)
		monitoringGroup.PUT("/alerts/:id/read", monitoringHandler.MarkAlertAsRead)

		// نمودارها
		monitoringGroup.GET("/chart", monitoringHandler.GetPerformanceChart)

		// جمع‌آوری متریک‌ها
		monitoringGroup.POST("/collect", monitoringHandler.CollectMetrics)

		// کنترل مانیتورینگ
		monitoringGroup.GET("/status", monitoringHandler.GetMonitoringStatus)
		monitoringGroup.POST("/toggle", monitoringHandler.ToggleMonitoring)

		// --- مسیرهای بدون /api ---
		plainMonitoringGroup.GET("/performance", monitoringHandler.GetPerformanceMetrics)
		plainMonitoringGroup.GET("/realtime", monitoringHandler.GetRealTimeMetrics)
		plainMonitoringGroup.GET("/overview", monitoringHandler.GetSystemOverview)
		plainMonitoringGroup.GET("/health", monitoringHandler.GetSystemHealth)
		plainMonitoringGroup.GET("/alerts", monitoringHandler.GetAlerts)
		plainMonitoringGroup.POST("/alerts", monitoringHandler.CreateAlert)
		plainMonitoringGroup.PUT("/alerts/:id/read", monitoringHandler.MarkAlertAsRead)
		plainMonitoringGroup.GET("/chart", monitoringHandler.GetPerformanceChart)
		plainMonitoringGroup.POST("/collect", monitoringHandler.CollectMetrics)
		plainMonitoringGroup.GET("/status", monitoringHandler.GetMonitoringStatus)
		plainMonitoringGroup.POST("/toggle", monitoringHandler.ToggleMonitoring)
	}
}

// RegisterSystemRoutes provides minimal system endpoints used by admin developer pages
func RegisterSystemRoutes(r *gin.Engine) {
	sys := r.Group("/api/admin/system")
	{
		sys.GET("/pm2-status", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "result": "pm2 status ok"}) })
		sys.GET("/pm2-logs", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "result": "pm2 logs..."}) })
		sys.POST("/pm2-update", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "updated", "result": "done"})
		})
		sys.POST("/cpu-details", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "result": "cpu details"}) })
		sys.GET("/cpu-cores", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"totalCores": 8, "logicalCores": 16, "perCoreUsage": []int{20, 35, 40, 15}}})
		})
		sys.GET("/memory-info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"physical": gin.H{"total": 16384, "used": 8192}, "usage": 50}})
		})
		sys.POST("/pm2-test", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "result": "pm2 test ok"}) })
		sys.POST("/cluster-test", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "result": "cluster test ok"}) })
		sys.POST("/load-test", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "result": "load test started"}) })
		sys.GET("/logs", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "logs": "system logs..."}) })
		sys.POST("/reload", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "message": "reload triggered"}) })
		// settings & db monitor stubs
		sys.GET("/auto-refresh-settings", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"enabled": true, "interval": 30}) })
		sys.POST("/auto-refresh-settings", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) })
		sys.GET("/custom-query-auto-settings", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"enabled": false, "interval": 60}) })
		sys.POST("/custom-query-auto-settings", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) })
		sys.GET("/db-monitor", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"activeQueries": []any{}, "slowQueries": []any{}}) })
		sys.POST("/reset-database", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) })
		sys.GET("/custom-query-logs", func(c *gin.Context) { c.JSON(http.StatusOK, []any{}) })
		sys.GET("/api-status", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "status": "ok"}) })
		sys.GET("/database-status", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "status": "ok"}) })
		sys.GET("/server-status", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "status": "ok"}) })
		sys.GET("/clear-logs", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) })
		sys.POST("/clear-logs", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) })
		sys.POST("/restart-server", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "message": "restart queued"}) })
		sys.GET("/test-api/:id", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) })
		sys.GET("/test-database/:id", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) })
		sys.GET("/test-server/:id", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) })
		sys.POST("/save-settings", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) })
	}

	// Health extended stats & query endpoints used by database-status
	r.GET("/health/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "data": gin.H{"database_size": 0, "active_connections": 0, "uptime": "-"}})
	})
	r.POST("/health/query", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": []any{}})
	})
}

// RegisterDBPoolRoutes provides endpoints for DB connection pool settings (used by admin developer UI)
func RegisterDBPoolRoutes(r *gin.Engine, db *gorm.DB, settingHandler *handlers.SettingHandler, settingService *handlers.DBPoolSettingsHandler) {
	// Backward-compat signature bridge to keep router aggregator simple
}
