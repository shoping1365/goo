package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUserAddressRoutes ثبت مسیرهای آدرس‌های کاربر
func RegisterUserAddressRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewAddressHandler(db)

	addr := public.Group("/user/addresses")
	addr.Use(middleware.Auth()) // کاربر باید لاگین باشد
	{
		addr.GET("", h.List)
		addr.POST("", h.Create)
		addr.PUT("/:id", h.Update)
	}
}
