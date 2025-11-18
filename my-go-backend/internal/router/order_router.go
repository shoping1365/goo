package router

import (
	"my-go-backend/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterOrderRoutes sets up order creation/payment endpoints
func RegisterOrderRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewOrderHandler(db)
	orders := public.Group("/orders")
	// orders.Use(middleware.Auth()) // احراز هویت غیرفعال شده است
	{
		orders.POST("/create", h.Create)
		orders.GET("", h.ListMyOrders)
		orders.GET("/:id/items", h.GetOrderItems)
	}
}
