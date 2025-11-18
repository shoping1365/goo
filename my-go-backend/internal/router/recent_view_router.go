package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRecentViewRoutes registers recent view endpoints
func RegisterRecentViewRoutes(public *gin.RouterGroup, admin *gin.RouterGroup, recentViewHandler *handlers.RecentViewHandler) {
	recentViews := public.Group("/recent-views")
	recentViews.Use(middleware.Auth()) // Require authentication for all recent view endpoints

	// Get user's recent views (last 30)
	recentViews.GET("", recentViewHandler.GetRecentViews)

	// Add a product to recent views (when user visits product page)
	recentViews.POST("/product/:product_id", recentViewHandler.AddRecentView)

	// Update duration of a product view (for analytics)
	recentViews.PATCH("/product/:product_id/duration", recentViewHandler.UpdateDuration)

	// Delete a specific recent view
	recentViews.DELETE("/:view_id", recentViewHandler.DeleteRecentView)

	// Clear all recent views for user
	recentViews.DELETE("", recentViewHandler.ClearAllRecentViews)

	// Admin endpoints
	if admin != nil {
		adminRecentViews := admin.Group("/recent-views")
		// Get recent views for a specific user (admin only)
		adminRecentViews.GET("/user/:user_id", recentViewHandler.GetUserRecentViewsAdmin)
		// Cleanup old views without device information
		adminRecentViews.DELETE("/cleanup-unknown", recentViewHandler.CleanupUnknownViews)
	}
}
