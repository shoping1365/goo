package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterFirewallRoutes ثبت مسیرهای مدیریت فایروال سیستم‌عامل
func RegisterFirewallRoutes(r *gin.Engine, h *handlers.FirewallHandler) {
	apiRead := r.Group("/api/admin/system-security")
	apiWrite := r.Group("/api/admin/system-security")
	plainRead := r.Group("/admin/system-security")
	plainWrite := r.Group("/admin/system-security")
	// Common auth
	apiRead.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware(), middleware.RequireAnyPermission("security_firewall.read", "security_firewall.write"))
	apiWrite.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware(), middleware.RequirePermission("security_firewall.write"))
	plainRead.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware(), middleware.RequireAnyPermission("security_firewall.read", "security_firewall.write"))
	plainWrite.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware(), middleware.RequirePermission("security_firewall.write"))

	// Query (read)
	apiRead.GET("/firewall/status", h.GetStatus)
	apiRead.GET("/firewall/rules", h.ListRules)
	apiRead.GET("/firewall/logs", h.ReadLogs)
	apiRead.GET("/firewall/logging", h.GetLoggingConfig)
	plainRead.GET("/firewall/status", h.GetStatus)
	plainRead.GET("/firewall/rules", h.ListRules)
	plainRead.GET("/firewall/logs", h.ReadLogs)
	plainRead.GET("/firewall/logging", h.GetLoggingConfig)

	// Commands (write)
	apiWrite.POST("/firewall/toggle", h.Toggle)
	apiWrite.POST("/firewall/mode", h.SetMode)
	apiWrite.POST("/firewall/rules", h.CreateRule)
	apiWrite.PUT("/firewall/rules/:name", h.UpdateRule)
	apiWrite.PUT("/firewall/rules/:name/toggle", h.ToggleRule)
	apiWrite.DELETE("/firewall/rules/:name", h.DeleteRule)
	apiWrite.DELETE("/firewall/logs", h.ClearLogs)
	apiWrite.POST("/firewall/backup", h.BackupRules)
	apiWrite.POST("/firewall/restore", h.RestoreRules)
	apiWrite.POST("/firewall/logging", h.SetLoggingConfig)

	plainWrite.POST("/firewall/toggle", h.Toggle)
	plainWrite.POST("/firewall/mode", h.SetMode)
	plainWrite.POST("/firewall/rules", h.CreateRule)
	plainWrite.PUT("/firewall/rules/:name", h.UpdateRule)
	plainWrite.PUT("/firewall/rules/:name/toggle", h.ToggleRule)
	plainWrite.DELETE("/firewall/rules/:name", h.DeleteRule)
	plainWrite.DELETE("/firewall/logs", h.ClearLogs)
	plainWrite.POST("/firewall/backup", h.BackupRules)
	plainWrite.POST("/firewall/restore", h.RestoreRules)
	plainWrite.POST("/firewall/logging", h.SetLoggingConfig)
}
