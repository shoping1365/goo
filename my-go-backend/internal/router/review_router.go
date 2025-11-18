package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterReviewRoutes exposes product review management endpoints
// این تابع مسیرهای مربوط به مدیریت نظرات محصولات را با محافظت امنیتی تعریف می‌کند
func RegisterReviewRoutes(public *gin.RouterGroup, db *gorm.DB) {
	uowf := unitofwork.NewUnitOfWorkFactory(db)
	h := handlers.NewReviewHandlerWithUoW(db, uowf)

	// Public, no-auth: لیست نظرات تاییدشدهٔ یک محصول
	public.GET("/products/:id/reviews", h.GetPublicReviewsByProduct)

	// Customer endpoints (auth required, بدون نیاز به مجوز ادمین)
	customer := public.Group("/reviews")
	customer.Use(middleware.Auth())
	{
		customer.POST("", h.CreateReview)
		customer.POST("/:id/helpful", h.MarkHelpful)
	}

	// Admin endpoints (auth + permission)
	admin := public.Group("/admin/reviews")
	admin.Use(middleware.Auth())
	{
		admin.GET("", middleware.RequirePermission("reviews_view"), h.GetReviews)
		admin.GET("/:id", middleware.RequirePermission("reviews_view"), h.GetReview)
		admin.PUT("/:id", middleware.RequirePermission("reviews_manage"), h.UpdateReview)
		admin.DELETE("/:id", middleware.RequirePermission("reviews_manage"), h.DeleteReview)
		admin.POST("/:id/approve", middleware.RequirePermission("reviews_manage"), h.ApproveReview)
		admin.POST("/:id/reject", middleware.RequirePermission("reviews_manage"), h.RejectReview)
		admin.POST("/:id/reply", middleware.RequirePermission("reviews_manage"), h.AddReply)
	}
}
