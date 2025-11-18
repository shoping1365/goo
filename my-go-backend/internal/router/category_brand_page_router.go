package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterCategoryBrandPageRoutes registers endpoints for category-brand pages under /admin/category-brand-pages
func RegisterCategoryBrandPageRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewCategoryBrandPageHandler(db)

	admin := public.Group("/admin")
	admin.Use(middleware.Auth())
	admin.Use(middleware.RequirePermission("category_brand_pages_manage"))
	{
		cbp := admin.Group("/category-brand-pages")
		cbp.GET("/check", h.CheckUnique)
		cbp.POST("", h.CreatePage)
	}
}
