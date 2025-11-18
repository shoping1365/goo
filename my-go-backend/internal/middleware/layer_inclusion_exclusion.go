package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LayerInclusionExclusionConfig تنظیمات شامل کردن و مستثنی کردن لایه‌ها
type LayerInclusionExclusionConfig struct {
	ExcludedPaths []string  `json:"excluded_paths"` // صفحات مستثنی برای لایه‌ها
	SpecificPaths []string  `json:"specific_paths"` // صفحات خاص برای لایه‌ها
	UpdatedAt     time.Time `json:"updated_at"`
}

// LayerInclusionExclusionMiddleware middleware برای مدیریت شامل کردن و مستثنی کردن لایه‌ها
func LayerInclusionExclusionMiddleware(db *gorm.DB) gin.HandlerFunc {
	// cache برای تنظیمات
	var configCache LayerInclusionExclusionConfig
	var lastUpdate time.Time
	cacheDuration := 5 * time.Minute // 5 دقیقه cache

	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// بررسی cache و به‌روزرسانی در صورت نیاز
		if time.Since(lastUpdate) > cacheDuration {
			updateLayerInclusionExclusionConfig(db, &configCache)
			lastUpdate = time.Now()
		}

		// بررسی اینکه آیا مسیر باید مستثنی شود
		if isLayerPathExcluded(path, configCache.ExcludedPaths) {
			// اگر مسیر مستثنی است، لایه‌ها را برنگردان
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// بررسی اینکه آیا مسیر باید شامل شود
		if isLayerPathIncluded(path, configCache.SpecificPaths) {
			// اگر مسیر شامل است، لایه‌ها را برگردان
			c.Next()
			return
		}

		// پیش‌فرض: لایه‌ها در تمام صفحات نمایش داده می‌شوند (به جز صفحات مستثنی)
		c.Next()
	}
}

// updateLayerInclusionExclusionConfig به‌روزرسانی تنظیمات شامل کردن و مستثنی کردن لایه‌ها از دیتابیس
func updateLayerInclusionExclusionConfig(db *gorm.DB, config *LayerInclusionExclusionConfig) {
	// خواندن تنظیمات از جدول header_layers
	var layerSettings []struct {
		PageSelection string `json:"page_selection"`
		ExcludedPages string `json:"excluded_pages"`
		SpecificPages string `json:"specific_pages"`
	}

	// خواندن لایه‌های فعال
	err := db.Table("header_layers").
		Select("page_selection, excluded_pages, specific_pages").
		Where("is_active = ?", true).
		Find(&layerSettings).Error

	if err != nil {
		// در صورت خطا، هیچ تنظیماتی اعمال نکن (پیش‌فرض: نمایش در تمام صفحات)
		config.ExcludedPaths = []string{}
		config.SpecificPaths = []string{}
		return
	}

	// پردازش تنظیمات
	var excludedPaths []string
	var specificPaths []string

	for _, setting := range layerSettings {
		if setting.PageSelection == "exclude" && setting.ExcludedPages != "" {
			// تبدیل صفحات مستثنی به آرایه
			pages := strings.Split(setting.ExcludedPages, ",")
			for _, page := range pages {
				page = strings.TrimSpace(page)
				if page != "" {
					excludedPaths = append(excludedPaths, page)
				}
			}
		}

		if setting.PageSelection == "specific" && setting.SpecificPages != "" {
			// تبدیل صفحات خاص به آرایه
			pages := strings.Split(setting.SpecificPages, ",")
			for _, page := range pages {
				page = strings.TrimSpace(page)
				if page != "" {
					specificPaths = append(specificPaths, page)
				}
			}
		}
	}

	// اگر هیچ تنظیماتی نبود، هیچ محدودیتی اعمال نکن (پیش‌فرض: نمایش در تمام صفحات)
	// excludedPaths و specificPaths خالی باقی می‌مانند

	config.ExcludedPaths = excludedPaths
	config.SpecificPaths = specificPaths
	config.UpdatedAt = time.Now()
}

// isLayerPathExcluded بررسی می‌کند که آیا مسیر در لیست مستثنی‌های لایه‌ها قرار دارد
func isLayerPathExcluded(path string, excludedPaths []string) bool {
	for _, excludedPath := range excludedPaths {
		// بررسی آدرس کامل (exact match)
		if path == excludedPath {
			return true
		}
	}
	return false
}

// isLayerPathIncluded بررسی می‌کند که آیا مسیر در لیست شامل‌های لایه‌ها قرار دارد
func isLayerPathIncluded(path string, specificPaths []string) bool {
	for _, specificPath := range specificPaths {
		// بررسی آدرس کامل (exact match)
		if path == specificPath {
			return true
		}
	}
	return false
}
