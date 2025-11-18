package router

import (
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupUserBankingInfoRoutes تنظیم مسیرهای اطلاعات بانکی کاربران
func SetupUserBankingInfoRoutes(r *gin.RouterGroup, bankingInfoHandler *handlers.UserBankingInfoHandler) {
	// مسیرهای عمومی (نیاز به احراز هویت)
	userRoutes := r.Group("/user")
	userRoutes.Use(middleware.Auth())
	{
		// مدیریت اطلاعات بانکی کاربر
		userRoutes.POST("/banking-info", bankingInfoHandler.Create)
		userRoutes.GET("/banking-info", bankingInfoHandler.Get)
		userRoutes.PUT("/banking-info/:id", bankingInfoHandler.Update)
		userRoutes.DELETE("/banking-info/:id", bankingInfoHandler.Delete)
		userRoutes.POST("/banking-info/:id/set-default", bankingInfoHandler.SetDefault)

		// API های تکمیل خودکار
		userRoutes.POST("/banking-info/auto-complete/card", bankingInfoHandler.AutoCompleteFromCard)
		userRoutes.POST("/banking-info/auto-complete/sheba", bankingInfoHandler.AutoCompleteFromSheba)
	}

	// مسیرهای ادمین (نیاز به احراز هویت و مجوز ادمین)
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.Auth(), middleware.RequirePermission("user.manage"))
	{
		// مشاهده اطلاعات بانکی یک کاربر
		adminRoutes.GET("/user/:userID/banking-info", bankingInfoHandler.GetByUserID)

		// تایید/لغو تایید اطلاعات بانکی
		adminRoutes.POST("/banking-info/:id/verify", bankingInfoHandler.Verify)
		adminRoutes.POST("/banking-info/:id/unverify", bankingInfoHandler.Unverify)
	}
}
