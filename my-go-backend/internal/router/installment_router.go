package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
)

// RegisterInstallmentRoutes ثبت مسیرهای ماژول خرید اقساطی
func RegisterInstallmentRoutes(r *gin.Engine, db *gorm.DB) {
	h := handlers.NewInstallmentHandler(db)

	// هم‌راستا با مسیرهای Nuxt: /api/admin/installment-payments/*
	admin := r.Group("/api/admin/installment-payments")
	{
		// لیست اقساط + فیلتر/جستجو/صفحه‌بندی
		admin.GET("", h.ListInstallments)

		// وضعیت سیستم داشبورد
		admin.GET("/status", h.GetStatusMetrics)

		// آمار خلاصه
		admin.GET("/stats", h.GetStats)

		// روند و تحلیل‌ها
		admin.GET("/trends", h.GetTrends)

		// محصولات برتر
		admin.GET("/top-products", h.GetTopProducts)

		// مدیریت محصولات اقساطی - تغییر مسیر از /products به /installment-products
		admin.GET("/installment-products", h.ListInstallmentProducts)
		admin.POST("/installment-products", h.CreateInstallmentProduct)
		admin.PUT("/installment-products/:id", h.UpdateInstallmentProduct)
		admin.POST("/installment-products/:id/toggle-status", h.ToggleProductStatus)

		// تنظیمات اعتبارسنجی سراسری
		admin.PUT("/credit-settings", h.CreditSettingsUpdate)

		// تب اعتبارسنجی (front expects these)
		admin.POST("/credit-check", h.CreditCheck)
		admin.POST("/approve", h.ApproveInstallment)
		admin.GET("/pending", h.ListPendingRequests)
		admin.POST(":id/approve", h.ApprovePendingRequest)

		// لاگ‌ها
		admin.GET("/logs", h.ListLogs)
		admin.DELETE("/logs", h.DeleteLogs)
		admin.POST("/logs/export", h.ExportLogs)

		// گزارش‌گیری و خروجی
		admin.POST("/export/excel", h.ExportExcel)
		admin.POST("/export/pdf", h.ExportPDF)
		admin.POST("/export/csv", h.ExportCSV)
		admin.PUT("/export/schedule", h.SaveExportSchedule)

		// جغرافیا
		admin.GET("/geography", h.GetGeography)
	}
}
