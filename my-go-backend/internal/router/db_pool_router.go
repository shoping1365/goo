package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterDBPoolSettingsRoutes ثبت مسیرهای تنظیمات کانکشن‌پول دیتابیس
func RegisterDBPoolSettingsRoutes(r *gin.Engine, db *gorm.DB, settingService *services.SettingService) {
	h := handlers.NewDBPoolSettingsHandler(db, settingService)
	sys := r.Group("/api/admin/system")
	{
		sys.GET("/db-pool", h.GetSettings)
		sys.POST("/db-pool", h.UpdateSettings)
	}
}
