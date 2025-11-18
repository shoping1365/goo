package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HeaderExclusionConfig تنظیمات مستثنی کردن هدر
type HeaderExclusionConfig struct {
	ExcludedPaths []string  `json:"excluded_paths"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// HeaderExclusionAdvancedMiddleware middleware پیشرفته با تنظیمات از دیتابیس
func HeaderExclusionAdvancedMiddleware(db *gorm.DB) gin.HandlerFunc {
	// cache برای تنظیمات
	var configCache HeaderExclusionConfig
	var lastUpdate time.Time
	cacheDuration := 5 * time.Minute // 5 دقیقه cache

	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// بررسی cache و به‌روزرسانی در صورت نیاز
		if time.Since(lastUpdate) > cacheDuration {
			updateExclusionConfig(db, &configCache)
			lastUpdate = time.Now()
		}

		// بررسی اینکه آیا مسیر مستثنی است
		if isPathExcluded(path, configCache.ExcludedPaths) {
			// اگر مسیر مستثنی است، هدر را برنگردان
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// updateExclusionConfig به‌روزرسانی تنظیمات مستثنی کردن و شامل کردن از دیتابیس
func updateExclusionConfig(db *gorm.DB, config *HeaderExclusionConfig) {
	// خواندن تنظیمات از جدول headers
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
		config.ExcludedPaths = []string{}
		return
	}

	// پردازش تنظیمات
	var excludedPaths []string

	for _, setting := range headerSettings {
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
					excludedPaths = append(excludedPaths, page)
				}
			}
		}
	}

	// اگر هیچ تنظیماتی نبود، هیچ محدودیتی اعمال نکن (پیش‌فرض: نمایش در تمام صفحات)
	// excludedPaths خالی باقی می‌ماند

	config.ExcludedPaths = excludedPaths
	config.UpdatedAt = time.Now()
}

// isPathExcluded بررسی می‌کند که آیا مسیر در لیست مستثنی‌ها قرار دارد
func isPathExcluded(path string, excludedPaths []string) bool {
	for _, excludedPath := range excludedPaths {
		// بررسی آدرس کامل (exact match)
		if path == excludedPath {
			return true
		}
	}
	return false
}

// getDefaultExcludedPaths مسیرهای پیش‌فرض مستثنی (آدرس کامل)
func getDefaultExcludedPaths() []string {
	return []string{
		"/admin",                                  // صفحه اصلی ادمین
		"/admin/content/header-management",        // صفحه مدیریت هدر
		"/admin/content/header-management/create", // صفحه ایجاد هدر
		"/auth/login",                             // صفحه ورود
		"/auth/register",                          // صفحه ثبت نام
		"/checkout",                               // صفحه پرداخت
		"/cart",                                   // سبد خرید
		"/api/admin",                              // API های ادمین
		"/_nuxt",                                  // فایل‌های Nuxt
		"/__nuxt",                                 // فایل‌های Nuxt (نسخه قدیمی)
	}
}
