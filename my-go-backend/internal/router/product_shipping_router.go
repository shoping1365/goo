package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
)

// RegisterProductShippingRoutes مسیردهی برای مدیریت حمل‌ونقل/ابعاد محصول را ثبت می‌کند.
// Route group: /api/product-shipping
func RegisterProductShippingRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewProductShippingHandler(db)
	g := public.Group("/product-shipping")
	{
		g.GET("/:id", h.GetShipping)
		g.PUT("/:id", h.UpdateShipping)
	}
}