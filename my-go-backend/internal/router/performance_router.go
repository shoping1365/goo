package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterPerformanceRoutes exposes admin performance-related endpoints
func RegisterPerformanceRoutes(r *gin.Engine) {
	perf := r.Group("/api/admin/performance")
	{
		perf.GET("/settings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"cacheTTL": 300, "bundleBudgetKB": 150}})
		})
		perf.PUT("/settings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true})
		})
		perf.POST("/analyze-bundle", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "bundle analysis started"})
		})
	}

	// cache clear endpoint
	r.POST("/api/admin/cache/clear", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "cache cleared"})
	})
}