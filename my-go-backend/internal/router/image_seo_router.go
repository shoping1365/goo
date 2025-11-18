package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterImageSEORoutes ثبت مسیرهای ادمین برای مانیتورینگ و Retry صف AI تصاویر
// این مسیرها معادل alias برای مسیرهای موجود زیر /media هستند تا دسترسی مستقیم /admin/* هم فراهم باشد.
func RegisterImageSEORoutes(r *gin.Engine, db *gorm.DB) {
	seoHandler := handlers.NewMediaSEOHandler(db)

	// آداپتر برای تبدیل http.HandlerFunc به gin.HandlerFunc
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
			ctx = context.WithValue(ctx, "gin_context", c)
			handler(c.Writer, c.Request.WithContext(ctx))
		}
	}

	// با پیشوند /api/admin
	adminAPI := r.Group("/api/admin")
	adminAPI.Use(middleware.Auth())
	{
		adminAPI.GET("/image-seo/jobs", middleware.RequirePermission("media_library.read"), httpToGin(seoHandler.ListJobs))
		adminAPI.POST("/image-seo/retry", middleware.RequirePermission("media_edit"), httpToGin(seoHandler.Retry))
		adminAPI.GET("/image-seo/worker/status", middleware.RequirePermission("media_library.read"), httpToGin(seoHandler.GetWorkerStatus))
	}

	// بدون پیشوند /api برای محیط dev: /admin/*
	adminDev := r.Group("/admin")
	adminDev.Use(middleware.Auth())
	{
		adminDev.GET("/image-seo/jobs", middleware.RequirePermission("media_library.read"), httpToGin(seoHandler.ListJobs))
		adminDev.POST("/image-seo/retry", middleware.RequirePermission("media_edit"), httpToGin(seoHandler.Retry))
		adminDev.GET("/image-seo/worker/status", middleware.RequirePermission("media_library.read"), httpToGin(seoHandler.GetWorkerStatus))
	}
}
