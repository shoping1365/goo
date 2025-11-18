package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterProductWarehouseStockRoutes مسیرهای موجودی محصول به تفکیک انبار
// Route group: /api/product-warehouse-stocks
func RegisterProductWarehouseStockRoutes(public *gin.RouterGroup, db *gorm.DB) {
	uowf := unitofwork.NewUnitOfWorkFactory(db)
	h := handlers.NewProductWarehouseStockHandler(uowf, db)
	grp := public.Group("/product-warehouse-stocks")
	{
		grp.GET("/:productId", h.GetStocksByProduct)
		grp.PUT("/:productId", middleware.Auth(), middleware.RequirePermission("products_edit"), h.PutStocksByProduct)
		grp.POST("/:productId/adjust", middleware.Auth(), middleware.RequirePermission("products_edit"), h.AdjustStock)
	}
}
