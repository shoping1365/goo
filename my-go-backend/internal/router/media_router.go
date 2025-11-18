package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
)

// RegisterMediaRoutes defines media upload/manipulation endpoints.
// Requires SettingService for compression and other configs.
func RegisterMediaRoutes(public *gin.RouterGroup, db *gorm.DB, settingService *services.SettingService) {
	// Handlers
	mediaHandler := handlers.NewMediaHandler(db, settingService)
	// ساخت سرویس‌ها برای worker AI
	apiSvc := services.NewAPISettingsService(db)
	seoSvc := services.NewImageSEOService(db, settingService, apiSvc)
	worker := services.NewImageSEOWorker(db, settingService, apiSvc, seoSvc)
	worker.Start()
	mediaHandler.ImageSEOWorker = worker
	imageHandler := handlers.NewImageHandler(mediaHandler)
	videoHandler := handlers.NewVideoHandler(mediaHandler)
	seoHandler := handlers.NewMediaSEOHandler(db)

	// Adapter for standard http.HandlerFunc to gin.HandlerFunc (copies useful context values)
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

			// اضافه کردن gin context به request context
			ctx = context.WithValue(ctx, "gin_context", c)

			handler(c.Writer, c.Request.WithContext(ctx))
		}
	}

	// Endpoint to list available categories (unauthenticated)
	public.GET("/media/categories", func(c *gin.Context) {
		var cats []string
		if err := db.Model(&models.MediaFile{}).Distinct("category").Pluck("category", &cats).Error; err != nil {
			c.JSON(500, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"success": true, "data": cats})
	})

	// Preview compress (no auth)
	public.POST("/media/preview-compress", httpToGin(mediaHandler.PreviewCompressHandler))
	public.POST("/media/preview-compress-video", httpToGin(videoHandler.PreviewCompressVideoHandler))

	// Image cropping endpoints (no auth for now)
	public.POST("/media/image-crop", httpToGin(mediaHandler.CropImageHandler))
	public.GET("/media/image-crop", httpToGin(mediaHandler.GetCroppedImageHandler))

	// Authenticated media group
	media := public.Group("/media")
	media.Use(middleware.Auth()) // احراز هویت فعال شد
	{
		// List endpoint با احراز هویت
		media.GET("/list", middleware.RequirePermission("media_library.read"), httpToGin(imageHandler.ListMediaHandler))
		media.GET("/videos", middleware.RequirePermission("media_library.read"), httpToGin(videoHandler.ListVideosHandler))

		// Download
		media.GET("/download/:id", middleware.RequirePermission("media_download"), httpToGin(imageHandler.DownloadMediaHandler))

		// Image operations
		media.POST("/rotate/:id", middleware.RequirePermission("media_edit"), httpToGin(imageHandler.RotateMediaHandler))
		media.POST("/compress/:id", middleware.RequirePermission("media_edit"), httpToGin(imageHandler.CompressMediaHandler))
		media.POST("/revert/:id", middleware.RequirePermission("media_edit"), httpToGin(imageHandler.RevertCompressionHandler))

		// Video compression operations
		media.POST("/compress-video/:id", middleware.RequirePermission("media_edit"), httpToGin(videoHandler.CompressVideoHandler))
		media.GET("/compress-video/progress/:jobID", middleware.RequirePermission("media_library.read"), httpToGin(videoHandler.CompressVideoProgressHandler))
		media.POST("/compress-video/cancel/:jobID", middleware.RequirePermission("media_edit"), httpToGin(videoHandler.CancelVideoCompressionHandler))
		media.POST("/compress-video/pause/:jobID", middleware.RequirePermission("media_edit"), httpToGin(videoHandler.PauseVideoCompressionHandler))
		media.POST("/compress-video/resume/:jobID", middleware.RequirePermission("media_edit"), httpToGin(videoHandler.ResumeVideoCompressionHandler))

		// Upload & Delete
		media.POST("/upload", middleware.RequirePermission("media_upload"), httpToGin(imageHandler.UploadMediaHandler))
		media.DELETE("/delete", middleware.RequirePermission("media_library.delete"), httpToGin(imageHandler.DeleteMediaHandler))

		// CRUD for single media item
		media.GET("/:id", middleware.RequirePermission("media_library.read"), httpToGin(imageHandler.GetMediaHandler))
		media.PUT("/:id", middleware.RequirePermission("media_edit"), httpToGin(imageHandler.UpdateMediaHandler))

		// Compression jobs monitoring
		media.GET("/compression-jobs", middleware.RequirePermission("media_library.read"), httpToGin(mediaHandler.GetCompressionJobsHandler))
		media.POST("/compression-jobs/clear-errors", middleware.RequirePermission("media_edit"), httpToGin(mediaHandler.ClearCompressionJobErrorsHandler))
		media.POST("/compression-jobs/:id/reset", middleware.RequirePermission("media_edit"), httpToGin(mediaHandler.ResetCompressionJobHandler))
		media.POST("/test-scheduler", middleware.RequirePermission("media_edit"), httpToGin(mediaHandler.TestSchedulerHandler))
		media.POST("/process-pending-video-compression", middleware.RequirePermission("media_edit"), httpToGin(mediaHandler.ProcessPendingVideoCompressionJobsHandler))

		// System health monitoring
		media.GET("/system-health", middleware.RequirePermission("media_library.read"), httpToGin(mediaHandler.GetSystemHealthHandler))

		// Admin image SEO jobs monitoring (under /media, but admin UI will call via proxy /api/admin/image-seo/*)
		media.GET("/admin/image-seo/jobs", middleware.RequirePermission("media_library.read"), httpToGin(seoHandler.ListJobs))
		media.POST("/admin/image-seo/retry", middleware.RequirePermission("media_edit"), httpToGin(seoHandler.Retry))
	}
}
