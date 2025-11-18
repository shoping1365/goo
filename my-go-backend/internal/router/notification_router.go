package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// SetupNotificationRoutes مسیرهای مدیریت اعلان‌های «خبرم کن» (موجودی و تخفیف)
// - عمومی: ایجاد درخواست توسط کاربر
// - ادمین: مشاهده/حذف/علامت‌گذاری ارسال
func SetupNotificationRoutes(r *gin.Engine, public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewStockDiscountNotificationHandler(db)
	patternProxy := handlers.NewSMSPatternProxyHandler(db)

	// عمومی
	pub := public.Group("/notifications")
	{
		// ایجاد درخواست
		pub.POST("/stock", h.CreateStockNotification)
		pub.POST("/discount", h.CreateDiscountNotification)
	}

	// ادمین
	admin := r.Group("/api/admin/notifications")
	admin.Use(middleware.AuthMiddleware())

	// گروه خواندن
	adminRead := admin.Group("")
	adminRead.Use(middleware.RequirePermission("product_notify_requests.read"))
	{
		adminRead.GET("/stock", h.ListStockNotifications)
		adminRead.GET("/discount", h.ListDiscountNotifications)
	}

	// گروه نوشتن
	adminWrite := admin.Group("")
	adminWrite.Use(middleware.RequirePermission("product_notify_requests.write"))
	{
		adminWrite.DELETE("/stock/:id", h.DeleteStockNotification)
		adminWrite.POST("/stock/:id/mark-sent", h.MarkStockAsSent)
		adminWrite.DELETE("/discount/:id", h.DeleteDiscountNotification)
		adminWrite.POST("/discount/:id/mark-sent", h.MarkDiscountAsSent)
	}

	// ابزارکمکی: ارسال با pattern_code (فقط برای اپراتوری)
	adminAux := r.Group("/api/admin/notify")
	adminAux.Use(middleware.AuthMiddleware())
	adminAux.Use(middleware.RequirePermission("product_notify_requests.write"))
	{
		adminAux.POST("/send-by-pattern", patternProxy.SendByPattern)
	}
}
