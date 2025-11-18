package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HeaderLayerInclusionExclusionConfig تنظیمات شامل کردن و مستثنی کردن هدر و لایه‌ها
type HeaderLayerInclusionExclusionConfig struct {
	HeaderExcludedPaths []string  `json:"header_excluded_paths"` // صفحات مستثنی برای هدر
	HeaderSpecificPaths []string  `json:"header_specific_paths"` // صفحات خاص برای هدر
	LayerExcludedPaths  []string  `json:"layer_excluded_paths"`  // صفحات مستثنی برای لایه‌ها
	LayerSpecificPaths  []string  `json:"layer_specific_paths"`  // صفحات خاص برای لایه‌ها
	UpdatedAt           time.Time `json:"updated_at"`
}

// HeaderLayerInclusionExclusionMiddleware middleware برای مدیریت شامل کردن و مستثنی کردن هدر و لایه‌ها
func HeaderLayerInclusionExclusionMiddleware(db *gorm.DB) gin.HandlerFunc {
	// cache برای تنظیمات
	var configCache HeaderLayerInclusionExclusionConfig
	var lastUpdate time.Time
	cacheDuration := 5 * time.Minute // 5 دقیقه cache

	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// بررسی cache و به‌روزرسانی در صورت نیاز
		if time.Since(lastUpdate) > cacheDuration {
			updateHeaderLayerInclusionExclusionConfig(db, &configCache)
			lastUpdate = time.Now()
		}

		// بررسی اینکه آیا مسیر باید مستثنی شود (برای هدر)
		if isHeaderLayerPathExcluded(path, configCache.HeaderExcludedPaths) {
			// اگر مسیر مستثنی است، هدر را برنگردان
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// بررسی اینکه آیا مسیر باید شامل شود (برای هدر)
		if isHeaderLayerPathIncluded(path, configCache.HeaderSpecificPaths) {
			// اگر مسیر شامل است، هدر را برگردان
			c.Next()
			return
		}

		// پیش‌فرض: هدر در تمام صفحات نمایش داده می‌شود (به جز صفحات مستثنی)
		c.Next()
	}
}

// updateHeaderLayerInclusionExclusionConfig به‌روزرسانی تنظیمات شامل کردن و مستثنی کردن هدر و لایه‌ها از دیتابیس
func updateHeaderLayerInclusionExclusionConfig(db *gorm.DB, config *HeaderLayerInclusionExclusionConfig) {
	// خواندن تنظیمات هدر از جدول headers
	var headerSettings []struct {
		PageSelection string `json:"page_selection"`
		ExcludedPages string `json:"excluded_pages"`
		SpecificPages string `json:"specific_pages"`
	}

	// خواندن هدرهای فعال
	err := db.Table("headers").
		Select("page_selection, excluded_pages, specific_pages").
		Where("is_active = ?", true).
		Find(&headerSettings).Error

	if err != nil {
		// در صورت خطا، هیچ تنظیماتی اعمال نکن (پیش‌فرض: نمایش در تمام صفحات)
		config.HeaderExcludedPaths = []string{}
		config.HeaderSpecificPaths = []string{}
		config.LayerExcludedPaths = []string{}
		config.LayerSpecificPaths = []string{}
		return
	}

	// پردازش تنظیمات هدر
	var headerExcludedPaths []string
	var headerSpecificPaths []string

	for _, setting := range headerSettings {
		if setting.PageSelection == "exclude" && setting.ExcludedPages != "" {
			// تبدیل صفحات مستثنی به آرایه
			pages := strings.Split(setting.ExcludedPages, ",")
			for _, page := range pages {
				page = strings.TrimSpace(page)
				if page != "" {
					headerExcludedPaths = append(headerExcludedPaths, page)
				}
			}
		}

		if setting.PageSelection == "specific" && setting.SpecificPages != "" {
			// تبدیل صفحات خاص به آرایه
			pages := strings.Split(setting.SpecificPages, ",")
			for _, page := range pages {
				page = strings.TrimSpace(page)
				if page != "" {
					headerSpecificPaths = append(headerSpecificPaths, page)
				}
			}
		}
	}

	// خواندن تنظیمات لایه از جدول header_layers
	var layerSettings []struct {
		PageSelection string `json:"page_selection"`
		ExcludedPages string `json:"excluded_pages"`
		SpecificPages string `json:"specific_pages"`
	}

	// خواندن تمام لایه‌ها (جدول header_layers ستون is_active ندارد)
	err = db.Table("header_layers").
		Select("page_selection, excluded_pages, specific_pages").
		Find(&layerSettings).Error

	if err != nil {
		// در صورت خطا، تنظیمات لایه را خالی کن
		config.LayerExcludedPaths = []string{}
		config.LayerSpecificPaths = []string{}
	} else {
		// پردازش تنظیمات لایه
		var layerExcludedPaths []string
		var layerSpecificPaths []string

		for _, setting := range layerSettings {
			if setting.PageSelection == "exclude" && setting.ExcludedPages != "" {
				// تبدیل صفحات مستثنی به آرایه
				pages := strings.Split(setting.ExcludedPages, ",")
				for _, page := range pages {
					page = strings.TrimSpace(page)
					if page != "" {
						layerExcludedPaths = append(layerExcludedPaths, page)
					}
				}
			}

			if setting.PageSelection == "specific" && setting.SpecificPages != "" {
				// تبدیل صفحات خاص به آرایه
				pages := strings.Split(setting.SpecificPages, ",")
				for _, page := range pages {
					page = strings.TrimSpace(page)
					if page != "" {
						layerSpecificPaths = append(layerSpecificPaths, page)
					}
				}
			}
		}

		config.LayerExcludedPaths = layerExcludedPaths
		config.LayerSpecificPaths = layerSpecificPaths
	}

	// اگر هیچ تنظیماتی نبود، هیچ محدودیتی اعمال نکن (پیش‌فرض: نمایش در تمام صفحات)
	// تمام آرایه‌ها خالی باقی می‌مانند

	config.HeaderExcludedPaths = headerExcludedPaths
	config.HeaderSpecificPaths = headerSpecificPaths
	config.UpdatedAt = time.Now()
}

// isHeaderLayerPathExcluded بررسی می‌کند که آیا مسیر در لیست مستثنی‌ها قرار دارد
func isHeaderLayerPathExcluded(path string, excludedPaths []string) bool {
	for _, excludedPath := range excludedPaths {
		// بررسی آدرس کامل (exact match)
		if path == excludedPath {
			return true
		}
	}
	return false
}

// isHeaderLayerPathIncluded بررسی می‌کند که آیا مسیر در لیست شامل‌ها قرار دارد
func isHeaderLayerPathIncluded(path string, specificPaths []string) bool {
	for _, specificPath := range specificPaths {
		// بررسی آدرس کامل (exact match)
		if path == specificPath {
			return true
		}
	}
	return false
}
