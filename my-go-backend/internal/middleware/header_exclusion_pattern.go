package middleware

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HeaderExclusionPatternConfig تنظیمات مستثنی کردن هدر با pattern matching
type HeaderExclusionPatternConfig struct {
	ExactPaths   []string  `json:"exact_paths"`   // آدرس‌های دقیق
	PatternPaths []string  `json:"pattern_paths"` // آدرس‌های با pattern
	UpdatedAt    time.Time `json:"updated_at"`
}

// HeaderExclusionPatternMiddleware middleware با پشتیبانی از آدرس کامل و pattern
func HeaderExclusionPatternMiddleware(db *gorm.DB) gin.HandlerFunc {
	// cache برای تنظیمات
	var configCache HeaderExclusionPatternConfig
	var lastUpdate time.Time
	cacheDuration := 5 * time.Minute // 5 دقیقه cache

	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// بررسی cache و به‌روزرسانی در صورت نیاز
		if time.Since(lastUpdate) > cacheDuration {
			updatePatternExclusionConfig(db, &configCache)
			lastUpdate = time.Now()
		}

		// بررسی اینکه آیا مسیر مستثنی است
		if isPathExcludedWithPattern(path, configCache.ExactPaths, configCache.PatternPaths) {
			// اگر مسیر مستثنی است، هدر را برنگردان
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// updatePatternExclusionConfig به‌روزرسانی تنظیمات مستثنی کردن و شامل کردن از دیتابیس
func updatePatternExclusionConfig(db *gorm.DB, config *HeaderExclusionPatternConfig) {
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
		config.ExactPaths = []string{}
		config.PatternPaths = []string{}
		return
	}

	// پردازش تنظیمات
	var exactPaths []string
	var patternPaths []string

	for _, setting := range headerSettings {
		if setting.PageSelection == "exclude" && setting.ExcludedPages != "" {
			// تبدیل صفحات مستثنی به آرایه
			pages := strings.Split(setting.ExcludedPages, ",")
			for _, page := range pages {
				page = strings.TrimSpace(page)
				if page != "" {
					// تشخیص نوع آدرس
					if isPatternPath(page) {
						patternPaths = append(patternPaths, page)
					} else {
						exactPaths = append(exactPaths, page)
					}
				}
			}
		}

		if setting.PageSelection == "specific" && setting.SpecificPages != "" {
			// تبدیل صفحات خاص به آرایه
			pages := strings.Split(setting.SpecificPages, ",")
			for _, page := range pages {
				page = strings.TrimSpace(page)
				if page != "" {
					// تشخیص نوع آدرس
					if isPatternPath(page) {
						patternPaths = append(patternPaths, page)
					} else {
						exactPaths = append(exactPaths, page)
					}
				}
			}
		}
	}

	// اگر هیچ تنظیماتی نبود، هیچ محدودیتی اعمال نکن (پیش‌فرض: نمایش در تمام صفحات)
	// exactPaths و patternPaths خالی باقی می‌مانند

	config.ExactPaths = exactPaths
	config.PatternPaths = patternPaths
	config.UpdatedAt = time.Now()
}

// isPatternPath بررسی می‌کند که آیا آدرس شامل pattern است
func isPatternPath(path string) bool {
	// آدرس‌هایی که شامل wildcard یا متغیر هستند
	patternIndicators := []string{
		"*",         // wildcard
		"{",         // متغیر
		"[",         // optional segment
		"sku-",      // SKU pattern
		"product/",  // محصولات
		"category/", // دسته‌بندی‌ها
	}

	for _, indicator := range patternIndicators {
		if strings.Contains(path, indicator) {
			return true
		}
	}

	return false
}

// isPathExcludedWithPattern بررسی می‌کند که آیا مسیر مستثنی است
func isPathExcludedWithPattern(path string, exactPaths []string, patternPaths []string) bool {
	// بررسی آدرس‌های دقیق
	for _, exactPath := range exactPaths {
		if path == exactPath {
			return true
		}
	}

	// بررسی آدرس‌های با pattern
	for _, patternPath := range patternPaths {
		if matchPattern(path, patternPath) {
			return true
		}
	}

	return false
}

// matchPattern بررسی می‌کند که آیا مسیر با pattern مطابقت دارد
func matchPattern(path string, pattern string) bool {
	// تبدیل pattern به regex
	regexPattern := convertPatternToRegex(pattern)

	// بررسی مطابقت
	matched, err := regexp.MatchString(regexPattern, path)
	if err != nil {
		return false
	}

	return matched
}

// convertPatternToRegex تبدیل pattern به regex
func convertPatternToRegex(pattern string) string {
	// جایگزینی wildcard ها
	regex := pattern

	// جایگزینی * با .*
	regex = strings.ReplaceAll(regex, "*", ".*")

	// جایگزینی {variable} با [^/]+
	regex = regexp.MustCompile(`\{[^}]+\}`).ReplaceAllString(regex, `[^/]+`)

	// جایگزینی [optional] با (?:optional)?
	regex = regexp.MustCompile(`\[([^\]]+)\]`).ReplaceAllString(regex, `(?:$1)?`)

	// اضافه کردن ^ و $ برای match کامل
	regex = "^" + regex + "$"

	return regex
}

// getDefaultExactPaths آدرس‌های دقیق پیش‌فرض
func getDefaultExactPaths() []string {
	return []string{
		"/admin",
		"/admin/content/header-management",
		"/admin/content/header-management/create",
		"/auth/login",
		"/auth/register",
		"/checkout",
		"/cart",
		"/api/admin",
	}
}

// getDefaultPatternPaths آدرس‌های با pattern پیش‌فرض
func getDefaultPatternPaths() []string {
	return []string{
		"/product/*",     // تمام محصولات
		"/product/sku-*", // محصولات با SKU
		"/product/*/*",   // محصولات با دو segment
		"/category/*",    // تمام دسته‌بندی‌ها
		"/admin/*",       // تمام صفحات ادمین
		"/auth/*",        // تمام صفحات احراز هویت
	}
}
