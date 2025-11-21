package security

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Mock Security Middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "Bearer valid-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Assume user role is set by AuthMiddleware
		// For test, we check a header or context
		role := c.GetHeader("X-Role")
		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.Next()
	}
}

func TestSecurityMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Protected Route
	protected := r.Group("/admin")
	protected.Use(AuthMiddleware())
	protected.Use(AdminMiddleware())
	protected.GET("/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin"})
	})

	t.Run("Unauthorized Access", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/admin/dashboard", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected 401, got %d", w.Code)
		}
	})

	t.Run("Forbidden Access (Not Admin)", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/admin/dashboard", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		req.Header.Set("X-Role", "user")
		
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusForbidden {
			t.Errorf("Expected 403, got %d", w.Code)
		}
	})

	t.Run("Authorized Access", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/admin/dashboard", nil)
		req.Header.Set("Authorization", "Bearer valid-token")
		req.Header.Set("X-Role", "admin")
		
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected 200, got %d", w.Code)
		}
	})
}
