package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// HeaderExclusionMiddleware middleware برای مستثنی کردن هدر در صفحات خاص
func HeaderExclusionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// دریافت مسیر درخواست
		path := c.Request.URL.Path

		// بررسی اینکه آیا این مسیر باید مستثنی شود
		if isExcludedPath(path) {
			// اگر مسیر مستثنی است، هدر را برنگردان
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// isExcludedPath بررسی می‌کند که آیا مسیر باید مستثنی شود
func isExcludedPath(path string) bool {
	// لیست مسیرهای مستثنی (آدرس کامل) - خالی برای پیش‌فرض
	excludedPaths := []string{}

	// بررسی آدرس کامل (exact match)
	for _, excludedPath := range excludedPaths {
		if path == excludedPath {
			return true
		}
	}

	return false
}

// isSpecificPath بررسی می‌کند که آیا مسیر باید شامل شود
func isSpecificPath(path string) bool {
	// لیست مسیرهای خاص (آدرس کامل) - خالی برای پیش‌فرض
	specificPaths := []string{}

	// بررسی آدرس کامل (exact match)
	for _, specificPath := range specificPaths {
		if path == specificPath {
			return true
		}
	}

	return false
}

// HeaderExclusionMiddlewareWithConfig middleware با تنظیمات قابل تغییر
func HeaderExclusionMiddlewareWithConfig(excludedPaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// بررسی با لیست سفارشی
		for _, excludedPath := range excludedPaths {
			if strings.HasPrefix(path, excludedPath) {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		}

		c.Next()
	}
}
