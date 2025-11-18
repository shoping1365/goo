package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/services"
)

// RegisterVideoRoutes defines video-specific endpoints.
// This router handles all video operations separately from general media operations.
func RegisterVideoRoutes(public *gin.RouterGroup, db *gorm.DB, settingService *services.SettingService) {
	// Handlers
	mediaHandler := handlers.NewMediaHandler(db, settingService)
	videoHandler := handlers.NewVideoHandler(mediaHandler)

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
			handler(c.Writer, c.Request.WithContext(ctx))
		}
	}

	// Video-specific group
	videos := public.Group("/videos")
	videos.Use(middleware.Auth())
	videos.Use(middleware.RequirePermission("video_access"))
	{
		// List videos
		videos.GET("/list", httpToGin(videoHandler.ListVideosHandler))
		videos.GET("/", httpToGin(videoHandler.ListVideosHandler))

		// Upload video
		videos.POST("/upload", httpToGin(videoHandler.UploadVideoHandler))

		// Download video
		videos.GET("/download/:id", httpToGin(videoHandler.DownloadVideoHandler))

		// Video compression operations
		videos.POST("/compress/:id", httpToGin(videoHandler.CompressVideoHandler))
		videos.GET("/compress/progress/:jobID", httpToGin(videoHandler.CompressVideoProgressHandler))
		videos.POST("/compress/cancel/:jobID", httpToGin(videoHandler.CancelVideoCompressionHandler))
		videos.POST("/compress/pause/:jobID", httpToGin(videoHandler.PauseVideoCompressionHandler))
		videos.POST("/compress/resume/:jobID", httpToGin(videoHandler.ResumeVideoCompressionHandler))

		// Video revert operation
		videos.POST("/revert/:id", httpToGin(videoHandler.RevertVideoCompressionHandler))

		// Delete video
		videos.DELETE("/delete", httpToGin(videoHandler.DeleteVideoHandler))

		// CRUD for single video item
		videos.GET("/:id", httpToGin(videoHandler.GetVideoHandler))
		videos.PUT("/:id", httpToGin(videoHandler.UpdateVideoHandler))
	}
}
