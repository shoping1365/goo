package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterMaintenanceRoutes defines temporary developer maintenance endpoints (/dev/*)
func RegisterMaintenanceRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewMaintenanceHandler(db)

	// Developer maintenance endpoints - only for developer role
	dev := public.Group("/dev")
	dev.Use(middleware.Auth()) // احراز هویت فعال شده است
	dev.Use(middleware.RequirePermission("developer_maintenance"))
	{
		dev.POST("/purge-soft-deletes", h.PurgeSoftDeletes)
		dev.GET("/soft-deletes", h.ListSoftDeleted)
		dev.POST("/restore-soft-delete", h.RestoreSoftDeleted)
		dev.POST("/hard-delete", h.HardDelete)
	}
}
