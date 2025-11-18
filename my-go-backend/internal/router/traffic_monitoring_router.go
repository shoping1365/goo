package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterTrafficMonitoringRoutes ثبت مسیرهای مربوط به نظارت بر ترافیک
func RegisterTrafficMonitoringRoutes(r *gin.Engine, db *gorm.DB) {
	// ایجاد handler
	trafficHandler := handlers.NewTrafficMonitoringHandler(db)

	// گروه مسیرهای نظارت بر ترافیک
	traffic := r.Group("/api/admin/traffic")
	{
		// اعمال middleware احراز هویت و دسترسی
		traffic.Use(middleware.AuthMiddleware())
		traffic.Use(middleware.RequirePermission("security_traffic.read"))

		// دریافت آمار کلی ترافیک
		traffic.GET("/stats", trafficHandler.GetTrafficStats)

		// دریافت لیست کاربران آنلاین
		traffic.GET("/online-users", trafficHandler.GetOnlineUsers)

		// آمار کلی/لیست لاگ‌ها با فیلتر
		traffic.GET("/general-stats", trafficHandler.GetGeneralStats)

		// به‌روزرسانی سشن کاربر
		traffic.POST("/update-session", trafficHandler.UpdateUserSession)

		// مسدود کردن آدرس IP
		traffic.POST("/block-ip", middleware.RequirePermission("security_traffic.create"), trafficHandler.BlockIP)

		// آزاد کردن آدرس IP
		traffic.DELETE("/unblock-ip/:ip", middleware.RequirePermission("security_traffic.update"), trafficHandler.UnblockIP)

		// دریافت لیست آدرس‌های IP مسدود شده
		traffic.GET("/blocked-ips", trafficHandler.GetBlockedIPs)

		// جزئیات بر اساس IP
		traffic.GET("/details/:ip", trafficHandler.GetTrafficDetailsByIP)
		traffic.GET("/user-details/:ip", trafficHandler.GetUserDetailsByIP)

		// پاکسازی سشن‌های منقضی شده
		traffic.POST("/cleanup", middleware.RequirePermission("security_traffic.delete"), trafficHandler.CleanupExpiredSessions)
	}
}
