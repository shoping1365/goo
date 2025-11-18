package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

func RegisterProductSpecRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewProductSpecHandler(db)
	// هماهنگ با سایر ماژول‌های محصول: استفاده از نام دسترسی واحد «products_edit»
	public.POST("/products/:id/specs", middleware.Auth(), middleware.RequirePermission("products_edit"), h.SaveSpecs)
	public.GET("/products/:id/specs", h.GetSpecs)
}
