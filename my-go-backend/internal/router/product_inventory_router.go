package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/handlers"
)

// RegisterProductInventoryRoutes defines endpoints for product inventory management.
// Route group:  /api/product-inventories
func RegisterProductInventoryRoutes(public *gin.RouterGroup, db *gorm.DB) {
	uowf := unitofwork.NewUnitOfWorkFactory(db)
	handler := handlers.NewProductInventoryHandler(uowf, db)

	inventories := public.Group("/product-inventories")
	{
		inventories.GET("/:id", handler.GetInventory) // Retrieve inventory details for a single product
		// هماهنگ با روت قیمت: استفاده از «products_edit» به‌عنوان دسترسی واحد
		inventories.PUT("/:id", handler.UpdateInventory)
	}
}
