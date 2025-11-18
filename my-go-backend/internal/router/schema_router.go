package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupSchemaRoutes تنظیم مسیرهای API اسکیمای پیش‌فرض
func SetupSchemaRoutes(r *gin.Engine, db *gorm.DB) {
	// ایجاد سرویس و هندلر
	schemaService := services.NewSchemaService(db)
	schemaHandler := handlers.NewSchemaHandler(schemaService)

	// گروه مسیرهای API اسکیما
	schemaGroup := r.Group("/api/schemas")
	{
		// دریافت تمام تمپلیت‌ها
		schemaGroup.GET("/templates", schemaHandler.GetAllTemplates)

		// دریافت انواع اسکیما
		schemaGroup.GET("/types", schemaHandler.GetSchemaTypes)

		// دریافت تمپلیت‌ها بر اساس نوع
		schemaGroup.GET("/templates/type/:type", schemaHandler.GetTemplatesByType)

		// دریافت تمپلیت بر اساس ID
		schemaGroup.GET("/templates/:id", schemaHandler.GetTemplateByID)

		// ایجاد تمپلیت جدید
		schemaGroup.POST("/templates", schemaHandler.CreateTemplate)

		// به‌روزرسانی تمپلیت
		schemaGroup.PUT("/templates/:id", schemaHandler.UpdateTemplate)

		// حذف تمپلیت
		schemaGroup.DELETE("/templates/:id", schemaHandler.DeleteTemplate)

		// تغییر وضعیت تمپلیت
		schemaGroup.PATCH("/templates/:id/toggle", schemaHandler.ToggleTemplateStatus)

		// تولید اسکیما از تمپلیت
		schemaGroup.POST("/templates/:id/generate", schemaHandler.GenerateSchemaFromTemplate)
	}
}
