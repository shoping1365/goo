package middleware

import "github.com/gin-gonic/gin"

// SecurityHeaders sets various HTTP security headers on every response.
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer-when-downgrade")
		c.Writer.Header().Set("Permissions-Policy", "geolocation=(), camera=(), microphone=()")
		// Basic CSP restricting to self and data URIs for images/media
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; img-src 'self' data:; media-src 'self' data:; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'; connect-src 'self' http://localhost:9090 http://localhost:3000")
		c.Next()
	}
}

// CORS middleware برای حل مشکل Cross-Origin Request
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// تنظیم هدرهای CORS
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-User-ID, X-User-Role")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// اگر درخواست OPTIONS است، پاسخ بده و متوقف شو
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
