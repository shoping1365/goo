package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
)

// RegisterCompressionSettingRoutes ثبت مسیر خواندن تنظیمات فشرده‌سازی.
// مسیر:
//
//	GET /api/settings/compression
//
// این درخواست تمام رکوردهای جدول compression_settings را بدون فیلتر برمی‌گرداند.
func RegisterCompressionSettingRoutes(public *gin.RouterGroup, db *gorm.DB) {
	if public == nil {
		return
	}

	public.GET("/settings/compression", func(c *gin.Context) {
		var settings []models.CompressionSetting
		if err := db.Find(&settings).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "data": settings})
	})
}
