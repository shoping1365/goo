package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterSEORedirectRoutes مسیرهای مدیریت ریدایرکت‌های سئو
func RegisterSEORedirectRoutes(r *gin.Engine, db *gorm.DB) {
	h := handlers.NewSEORedirectHandler(db)
	admin := r.Group("/api/admin/seo", middleware.Auth(), middleware.Admin())
	{
		admin.GET("/redirects", h.ListRedirects)
		admin.POST("/redirects", h.CreateRedirect)
		admin.DELETE("/redirects/:id", h.DeleteRedirect)
		admin.DELETE("/redirects/bulk-delete", h.BulkDeleteRedirects)
	}
}
