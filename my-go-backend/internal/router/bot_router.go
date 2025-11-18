package router

import (
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterBotRoutes مسیرهای مدیریت ربات‌ها را ثبت می‌کند
func RegisterBotRoutes(r *gin.Engine, db *gorm.DB, settingService *services.SettingService) {
	uowFactory := unitofwork.NewUnitOfWorkFactory(db)
	botSvc := services.NewBotService(uowFactory)
	h := handlers.NewBotHandler(settingService, botSvc)
	g := r.Group("/api/admin/bots")
	dev := r.Group("/admin/bots") // mirror بدون /api برای dev
	for _, grp := range []*gin.RouterGroup{g, dev} {
		grp.Use(middleware.Auth())
		grp.Use(middleware.RequirePermission("security_bots.read"))
		grp.GET("/config", h.GetConfig)
		grp.PUT("/config", middleware.RequirePermission("security_bots.update"), h.UpdateConfig)
		grp.GET("", h.ListBots)
		grp.POST("/malicious/import", middleware.RequirePermission("security_bots.update"), h.ImportMaliciousPatterns)
	}
}
