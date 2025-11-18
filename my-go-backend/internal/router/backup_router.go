package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterBackupRoutes adds /api/admin/media/backup/* endpoints
func RegisterBackupRoutes(r *gin.Engine, db *gorm.DB) {
	h := handlers.NewBackupHandler(db)

	// adapter to preserve context keys similar to main router
	httpToGin := func(handler func(http.ResponseWriter, *http.Request)) gin.HandlerFunc {
		return func(c *gin.Context) {
			ctx := c.Request.Context()
			if v, ok := c.Get("user_id"); ok {
				ctx = context.WithValue(ctx, "user_id", v)
			}
			if v, ok := c.Get("username"); ok {
				ctx = context.WithValue(ctx, "username", v)
			}
			if v, ok := c.Get("role"); ok {
				ctx = context.WithValue(ctx, "role", v)
			}
			handler(c.Writer, c.Request.WithContext(ctx))
		}
	}

	r.GET("/api/admin/media/backup/periods", middleware.Auth(), middleware.RequirePermission("media_backup_access"), httpToGin(h.ListPeriods))
	r.POST("/api/admin/media/backup/restore/:period", middleware.Auth(), middleware.RequirePermission("media_backup_restore"), httpToGin(h.RestorePeriod))
}
