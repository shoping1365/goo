package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterWarehouseRoutes ثبت مسیرهای مدیریت انبار (ادمین)
func RegisterWarehouseRoutes(r *gin.Engine, db *gorm.DB) {
	uowf := unitofwork.NewUnitOfWorkFactory(db)
	h := handlers.NewWarehouseHandler(uowf)
	// مسیرهای استاندارد ادمین با پیشوند /api
	admin := r.Group("/api/admin")
	{
		admin.GET("/warehouses", middleware.Auth(), middleware.RequirePermission("warehouses_read"), h.List)
		admin.POST("/warehouses", middleware.Auth(), middleware.RequirePermission("warehouses_write"), h.Create)
		admin.PUT("/warehouses/:id", middleware.Auth(), middleware.RequirePermission("warehouses_write"), h.Update)
		admin.DELETE("/warehouses/:id", middleware.Auth(), middleware.RequirePermission("warehouses_write"), h.Delete)
	}

	// مسیرهای توسعه بدون پیشوند /api برای سازگاری با پروکسی Dev
	// نکته: احراز هویت و مجوزها دقیقاً مانند مسیرهای /api اعمال می‌شوند
	devAdmin := r.Group("/admin")
	{
		devAdmin.GET("/warehouses", middleware.Auth(), middleware.RequirePermission("warehouses_read"), h.List)
		devAdmin.POST("/warehouses", middleware.Auth(), middleware.RequirePermission("warehouses_write"), h.Create)
		devAdmin.PUT("/warehouses/:id", middleware.Auth(), middleware.RequirePermission("warehouses_write"), h.Update)
		devAdmin.DELETE("/warehouses/:id", middleware.Auth(), middleware.RequirePermission("warehouses_write"), h.Delete)
	}
}
