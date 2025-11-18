package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterProductCategoryRoutes defines CRUD endpoints for product categories
func RegisterProductCategoryRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewProductCategoryHandler(db)

	// Auth + Permission برای عملیات نوشتاری
	categories := public.Group("/product-categories")
	{
		categories.GET("", h.ListCategories)
		categories.POST("", middleware.Auth(), middleware.RequirePermission("categories_manage"), h.CreateCategory)
		categories.GET("/:id", h.GetCategory)
		categories.PUT("/:id", middleware.Auth(), middleware.RequirePermission("categories_manage"), h.UpdateCategory)
		categories.DELETE("/:id", middleware.Auth(), middleware.RequirePermission("categories_manage"), h.DeleteCategory)
		categories.POST("/reorder", middleware.Auth(), middleware.RequirePermission("categories_manage"), h.ReorderCategories)
		// public endpoint: get category by slug with optional preview=1
		categories.GET("/slug/:slug", h.GetCategoryBySlug)
	}
}
