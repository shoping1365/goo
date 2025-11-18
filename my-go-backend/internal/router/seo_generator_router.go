package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
)

// RegisterSeoGeneratorRoutes ثبت مسیرهای API تولید پویای SEO
func RegisterSeoGeneratorRoutes(router *gin.Engine, db *gorm.DB) {
	seoHandler := handlers.NewSeoGeneratorHandler(db)

	// گروه مسیرهای SEO Generator
	seo := router.Group("/api/seo-generator")
	{
		// تولید تگ‌های SEO برای یک نوشته موجود
		seo.GET("/post/:id", seoHandler.GenerateSeoTags)

		// تولید تگ‌های SEO از داده‌های ارسالی
		seo.POST("/generate", seoHandler.GenerateSeoTagsFromData)

		// پیش‌نمایش تگ‌های SEO
		seo.POST("/preview", seoHandler.PreviewSeoTags)
	}
}
