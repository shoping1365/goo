package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
)

// RegisterBrandRoutes defines endpoints for product brands
func RegisterBrandRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewBrandHandler(db)

	brands := public.Group("/brands")
	{
		brands.GET("", h.ListBrands)
		brands.POST("", h.CreateBrand)
		brands.GET("/:id", h.GetBrand)
		brands.PUT("/:id", h.UpdateBrand)
		brands.DELETE("/:id", h.DeleteBrand)
	}
}
