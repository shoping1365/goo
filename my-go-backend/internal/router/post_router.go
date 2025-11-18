package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterPostRoutes تعریف endpoint های CRUD برای نوشته‌ها
func RegisterPostRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewPostHandler(db)

	posts := public.Group("/posts")
	{
		posts.GET("", middleware.AuthOptional(), h.ListPosts)
		
		// Endpoint های جدید برای بررسی و تولید slug یکتا (قبل از :id)
		posts.GET("/check-slug", middleware.Auth(), middleware.RequirePermission("posts_manage"), h.CheckSlugUnique)
		posts.GET("/generate-slug", middleware.Auth(), middleware.RequirePermission("posts_manage"), h.GenerateUniqueSlug)
		
		// Endpoint تست برای بررسی routing
		posts.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Posts router is working", "path": c.Request.URL.Path})
		})
		
		// این route باید آخر باشد چون همه چیز را catch می‌کند
		posts.GET("/:id", middleware.AuthOptional(), h.GetPost)
		posts.POST("", middleware.Auth(), middleware.RequirePermission("posts_manage"), h.CreatePost)
		posts.PUT("/:id", middleware.Auth(), middleware.RequirePermission("posts_manage"), h.UpdatePost)
		posts.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("posts_manage"), h.DeletePost)
	}
}
