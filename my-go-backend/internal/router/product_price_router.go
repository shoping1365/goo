package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterProductPriceRoutes defines endpoints dedicated to managing product pricing.
// Route group:  /api/product-prices
// NOTE: Auth & permission middleware should be added by caller (e.g., admin group) if required.
func RegisterProductPriceRoutes(public *gin.RouterGroup, db *gorm.DB) {
	handler := handlers.NewProductPriceHandler(db)

	prices := public.Group("/product-prices")
	{
		prices.GET("/:id", handler.GetPrice) // Retrieve price details for a single product
		// هماهنگ با سایر ماژول‌های محصول: «products_edit»
		prices.PUT("/:id", middleware.Auth(), middleware.RequirePermission("products_edit"), handler.UpdatePrice)               // Update pricing information
		prices.POST("/bulk-update", middleware.Auth(), middleware.RequirePermission("products_edit"), handler.BulkUpdatePrices) // Bulk update prices for multiple products
	}
}
