package middleware

import (
	"github.com/gin-gonic/gin"
)

// clientIP استانداردسازی IP واقعی از هدرهای متداول
// از c.ClientIP استفاده می‌کنیم تا با میدلورهای موجود سازگار باشد

// TrafficLogger: اندازه‌گیری و ارسال لاگ به صف درج انبوه
func TrafficLogger() gin.HandlerFunc {
	// [TEMP DISABLED]
	return func(c *gin.Context) { c.Next() }
}
