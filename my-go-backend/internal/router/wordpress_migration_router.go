package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterWordPressMigrationRoutes wires admin endpoints for WordPress migration
func RegisterWordPressMigrationRoutes(r *gin.Engine, db *gorm.DB) {
	h := handlers.NewWordPressMigrationHandler(db)

	admin := r.Group("/api/admin/wordpress-migration")
	admin.Use(middleware.Auth()) // اضافه کردن middleware احراز هویت
	{
		admin.POST("/test-connection", middleware.RequirePermission("developer_migration"), h.TestConnection)
		admin.POST("/start", middleware.RequirePermission("developer_migration"), h.StartMigration)
		admin.POST("/save-plan", middleware.RequirePermission("developer_migration"), h.SavePlan)
		admin.POST("/run-scheduled", middleware.RequirePermission("developer_migration"), h.RunScheduled)
		admin.GET("/logs", middleware.RequirePermission("developer_migration"), h.GetMigrationLogs)
		admin.POST("/abort", middleware.RequirePermission("developer_migration"), h.AbortMigration)
		admin.POST("/extract-meta", middleware.RequirePermission("developer_migration"), h.ExtractMetaFromWordPressProduct)
		admin.POST("/total-products-count", middleware.RequirePermission("developer_migration"), h.GetTotalProductsCount)
	}

	// dev plain prefix for local tools
	dev := r.Group("/admin/wordpress-migration")
	{
		dev.POST("/test-connection", h.TestConnection)
		dev.POST("/start", h.StartMigration)
		dev.POST("/save-plan", h.SavePlan)
		dev.POST("/run-scheduled", h.RunScheduled)
		dev.GET("/logs", h.GetMigrationLogs)
		dev.POST("/abort", h.AbortMigration)
		dev.POST("/extract-meta", h.ExtractMetaFromWordPressProduct)
		dev.POST("/total-products-count", h.GetTotalProductsCount)
	}
}
