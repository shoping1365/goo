package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
)

// RegisterPostCategoryRoutes تعریف endpoint های CRUD برای دسته‌بندی‌های نوشته‌ها
func RegisterPostCategoryRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewPostCategoryHandler(db)

	categories := public.Group("/post-categories")
	{
		categories.GET("", h.ListCategories)
		categories.POST("", h.CreateCategory)
		categories.GET("/:id", h.GetCategory)
		categories.PUT("/:id", h.UpdateCategory)
		categories.DELETE("/:id", h.DeleteCategory)
		categories.PATCH("/:id/toggle-status", h.ToggleCategoryStatus)
	}
}
