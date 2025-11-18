package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// FooterLayerInclusionExclusionConfig تنظیمات شامل کردن و مستثنی کردن فوتر و لایه‌ها
type FooterLayerInclusionExclusionConfig struct {
	FooterExcludedPaths []string  `json:"footer_excluded_paths"` // صفحات مستثنی برای فوتر
	FooterSpecificPaths []string  `json:"footer_specific_paths"` // صفحات خاص برای فوتر
	LayerExcludedPaths  []string  `json:"layer_excluded_paths"`  // صفحات مستثنی برای لایه‌ها
	LayerSpecificPaths  []string  `json:"layer_specific_paths"`  // صفحات خاص برای لایه‌ها
	UpdatedAt           time.Time `json:"updated_at"`
}

// FooterLayerInclusionExclusionMiddleware middleware برای مدیریت شامل کردن و مستثنی کردن فوتر و لایه‌ها
func FooterLayerInclusionExclusionMiddleware(db *gorm.DB) gin.HandlerFunc {
	// cache برای تنظیمات
	var configCache FooterLayerInclusionExclusionConfig
	var lastUpdate time.Time
	cacheDuration := 5 * time.Minute // 5 دقیقه cache

	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// بررسی cache و به‌روزرسانی در صورت نیاز
		if time.Since(lastUpdate) > cacheDuration {
			updateFooterLayerInclusionExclusionConfig(db, &configCache)
			lastUpdate = time.Now()
		}

		// بررسی اینکه آیا مسیر باید مستثنی شود (برای فوتر)
		if isFooterLayerPathExcluded(path, configCache.FooterExcludedPaths) {
			// اگر مسیر مستثنی است، فوتر را برنگردان
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// بررسی اینکه آیا مسیر باید شامل شود (برای فوتر)
		if isFooterLayerPathIncluded(path, configCache.FooterSpecificPaths) {
			// اگر مسیر شامل است، فوتر را برگردان
			c.Next()
			return
		}

		// پیش‌فرض: فوتر در تمام صفحات نمایش داده می‌شود (به جز صفحات مستثنی)
		c.Next()
	}
}

// updateFooterLayerInclusionExclusionConfig به‌روزرسانی تنظیمات شامل کردن و مستثنی کردن فوتر و لایه‌ها از دیتابیس
func updateFooterLayerInclusionExclusionConfig(db *gorm.DB, config *FooterLayerInclusionExclusionConfig) {
	// خواندن تنظیمات فوتر از جدول footers
	var footerSettings []struct {
		PageSelection string `json:"page_selection"`
		ExcludedPages string `json:"excluded_pages"`
		SpecificPages string `json:"specific_pages"`
	}

	// خواندن فوترهای فعال
	err := db.Table("footers").
		Select("page_selection, excluded_pages, specific_pages").
		Where("is_active = ?", true).
		Find(&footerSettings).Error

	if err != nil {
		// در صورت خطا، هیچ تنظیماتی اعمال نکن (پیش‌فرض: نمایش در تمام صفحات)
		config.FooterExcludedPaths = []string{}
		config.FooterSpecificPaths = []string{}
		config.LayerExcludedPaths = []string{}
		config.LayerSpecificPaths = []string{}
		return
	}

	// پردازش تنظیمات فوتر
	var footerExcludedPaths []string
	var footerSpecificPaths []string

	for _, setting := range footerSettings {
		if setting.PageSelection == "exclude" && setting.ExcludedPages != "" {
			// تبدیل صفحات مستثنی به آرایه
			pages := strings.Split(setting.ExcludedPages, ",")
			for _, page := range pages {
				page = strings.TrimSpace(page)
				if page != "" {
					footerExcludedPaths = append(footerExcludedPaths, page)
				}
			}
		}

		if setting.PageSelection == "specific" && setting.SpecificPages != "" {
			// تبدیل صفحات خاص به آرایه
			pages := strings.Split(setting.SpecificPages, ",")
			for _, page := range pages {
				page = strings.TrimSpace(page)
				if page != "" {
					footerSpecificPaths = append(footerSpecificPaths, page)
				}
			}
		}
	}

	// تنظیمات لایه را خالی بگذاریم چون در footer_layers این ستون‌ها وجود ندارند
	// فقط از تنظیمات اصلی footers استفاده می‌کنیم
	config.LayerExcludedPaths = []string{}
	config.LayerSpecificPaths = []string{}

	// اگر هیچ تنظیماتی نبود، هیچ محدودیتی اعمال نکن (پیش‌فرض: نمایش در تمام صفحات)
	// تمام آرایه‌ها خالی باقی می‌مانند

	config.FooterExcludedPaths = footerExcludedPaths
	config.FooterSpecificPaths = footerSpecificPaths
	config.UpdatedAt = time.Now()
}

// isFooterLayerPathExcluded بررسی می‌کند که آیا مسیر در لیست مستثنی‌ها قرار دارد
func isFooterLayerPathExcluded(path string, excludedPaths []string) bool {
	for _, excludedPath := range excludedPaths {
		// بررسی آدرس کامل (exact match)
		if path == excludedPath {
			return true
		}
	}
	return false
}

// isFooterLayerPathIncluded بررسی می‌کند که آیا مسیر در لیست شامل‌ها قرار دارد
func isFooterLayerPathIncluded(path string, specificPaths []string) bool {
	for _, specificPath := range specificPaths {
		// بررسی آدرس کامل (exact match)
		if path == specificPath {
			return true
		}
	}
	return false
}
