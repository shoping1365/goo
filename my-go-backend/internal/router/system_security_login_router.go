package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterSystemSecurityLoginRoutes ثبت مسیرهای صفحه Admin › System Security › Login
func RegisterSystemSecurityLoginRoutes(r *gin.Engine, db *gorm.DB) {
	h := handlers.NewSystemSecurityLoginHandler(db)

	grp := r.Group("/api/admin/system-security/login")
	{
		// احراز هویت و سطح دسترسی ادمین
		grp.Use(middleware.AuthMiddleware())
		grp.Use(middleware.RequirePermission("security_login.read"))

		grp.GET("/counters", h.GetCounters)
		grp.GET("/history", h.ListHistory)

		grp.POST("/block-ip", middleware.RequirePermission("security_login.create"), h.BlockIP)
		grp.DELETE("/unblock-ip/:ip", middleware.RequirePermission("security_login.update"), h.UnblockIP)
	}
}
