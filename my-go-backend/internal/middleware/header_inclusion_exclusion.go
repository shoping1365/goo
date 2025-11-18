package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HeaderInclusionExclusionConfig تنظیمات شامل کردن و مستثنی کردن هدر
type HeaderInclusionExclusionConfig struct {
	ExcludedPaths []string  `json:"excluded_paths"` // صفحات مستثنی
	SpecificPaths []string  `json:"specific_paths"` // صفحات خاص
	UpdatedAt     time.Time `json:"updated_at"`
}

// HeaderInclusionExclusionMiddleware middleware برای مدیریت شامل کردن و مستثنی کردن هدر
func HeaderInclusionExclusionMiddleware(db *gorm.DB) gin.HandlerFunc {
	// cache برای تنظیمات
	var configCache HeaderInclusionExclusionConfig
	var lastUpdate time.Time
	cacheDuration := 5 * time.Minute // 5 دقیقه cache

	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// بررسی cache و به‌روزرسانی در صورت نیاز
		if time.Since(lastUpdate) > cacheDuration {
			updateInclusionExclusionConfig(db, &configCache)
			lastUpdate = time.Now()
		}

		// بررسی اینکه آیا مسیر باید مستثنی شود
		if isPathExcludedInclusion(path, configCache.ExcludedPaths) {
			// اگر مسیر مستثنی است، هدر را برنگردان
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// بررسی اینکه آیا مسیر باید شامل شود
		if isPathIncluded(path, configCache.SpecificPaths) {
			// اگر مسیر شامل است، هدر را برگردان
			c.Next()
			return
		}

		// پیش‌فرض: هدر در تمام صفحات نمایش داده می‌شود (به جز صفحات مستثنی)
		c.Next()
	}
}

// updateInclusionExclusionConfig به‌روزرسانی تنظیمات شامل کردن و مستثنی کردن از دیتابیس
func updateInclusionExclusionConfig(db *gorm.DB, config *HeaderInclusionExclusionConfig) {
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
		config.SpecificPaths = []string{}
		return
	}

	// پردازش تنظیمات
	var excludedPaths []string
	var specificPaths []string

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

// isPathExcludedInclusion بررسی می‌کند که آیا مسیر در لیست مستثنی‌ها قرار دارد
func isPathExcludedInclusion(path string, excludedPaths []string) bool {
	for _, excludedPath := range excludedPaths {
		// بررسی آدرس کامل (exact match)
		if path == excludedPath {
			return true
		}
	}
	return false
}

// isPathIncluded بررسی می‌کند که آیا مسیر در لیست شامل‌ها قرار دارد
func isPathIncluded(path string, specificPaths []string) bool {
	for _, specificPath := range specificPaths {
		// بررسی آدرس کامل (exact match)
		if path == specificPath {
			return true
		}
	}
	return false
}

// getDefaultExcludedPathsInclusion مسیرهای پیش‌فرض مستثنی (آدرس کامل)
func getDefaultExcludedPathsInclusion() []string {
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

// getDefaultSpecificPaths مسیرهای پیش‌فرض شامل (آدرس کامل)
func getDefaultSpecificPaths() []string {
	return []string{
		"/",                           // صفحه اصلی
		"/product/sku-123456/example", // محصول خاص (مثال)
		"/product/sku-187139/tst-12",  // محصول خاص (مثال)
		"/category/electronics",       // دسته‌بندی خاص
		"/about",                      // صفحه درباره ما
		"/contact",                    // صفحه تماس
		"/blog",                       // صفحه وبلاگ
		"/products",                   // صفحه محصولات
	}
}
