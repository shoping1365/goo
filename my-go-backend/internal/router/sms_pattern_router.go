package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// SetupSMSPatternRoutes تنظیم مسیرهای API پترن‌های پیامک
func SetupSMSPatternRoutes(r *gin.Engine, db *gorm.DB) {
	// ایجاد handler
	patternHandler := handlers.NewSMSPatternHandler(db)
	testPatternHandler := handlers.NewTestPatternHandler(db)

	// گروه مسیرهای پترن‌های پیامک
	smsPatterns := r.Group("/api/admin/sms-patterns")
	{
		// اعمال middleware احراز هویت و دسترسی
		smsPatterns.Use(middleware.AuthMiddleware())
		// smsPatterns.Use(middleware.AdminPermissionMiddleware("sms_patterns")) // فعلاً غیرفعال

		// دریافت لیست پترن‌ها
		smsPatterns.GET("", patternHandler.GetPatterns)

		// دریافت پترن بر اساس feature برای مصرف اپلیکیشن‌ها
		smsPatterns.GET("/by-feature", patternHandler.GetPatternByFeature)

		// دریافت آمار پترن‌ها
		smsPatterns.GET("/stats", patternHandler.GetPatternStats)

		// دریافت پترن خاص
		smsPatterns.GET("/:id", patternHandler.GetPattern)

		// ایجاد پترن جدید
		smsPatterns.POST("", patternHandler.CreatePattern)

		// ویرایش پترن
		smsPatterns.PUT("/:id", patternHandler.UpdatePattern)

		// حذف پترن
		smsPatterns.DELETE("/:id", patternHandler.DeletePattern)

		// تست پترن
		smsPatterns.POST("/:id/test", testPatternHandler.TestPattern)

		// کپی کردن پترن
		smsPatterns.POST("/:id/duplicate", patternHandler.DuplicatePattern)

		// ارسال پیامک بر اساس scope و feature
		smsPatterns.POST("/send-by-scope-feature", patternHandler.SendSMSByScopeAndFeature)

		// بررسی وضعیت درگاه‌های پیامکی
		smsPatterns.GET("/gateways/status", patternHandler.GetGatewayStatus)

		// پاکسازی پترن‌های نامعتبر (یکباره)
		smsPatterns.POST("/cleanup", patternHandler.CleanupInvalidPatterns)
	}

	// مسیرهای بدون پیشوند /api برای محیط dev
	smsPatternsNoPrefix := r.Group("/admin/sms-patterns")
	{
		smsPatternsNoPrefix.GET("", patternHandler.GetPatterns)
		smsPatternsNoPrefix.GET("/stats", patternHandler.GetPatternStats)
		smsPatternsNoPrefix.GET("/:id", patternHandler.GetPattern)
		smsPatternsNoPrefix.POST("", patternHandler.CreatePattern)
		smsPatternsNoPrefix.PUT("/:id", patternHandler.UpdatePattern)
		smsPatternsNoPrefix.DELETE("/:id", patternHandler.DeletePattern)
		smsPatternsNoPrefix.POST("/:id/test", testPatternHandler.TestPattern)
		smsPatternsNoPrefix.POST("/:id/duplicate", patternHandler.DuplicatePattern)
		smsPatternsNoPrefix.POST("/send-by-scope-feature", patternHandler.SendSMSByScopeAndFeature)
		smsPatternsNoPrefix.GET("/gateways/status", patternHandler.GetGatewayStatus)
		smsPatternsNoPrefix.POST("/cleanup", patternHandler.CleanupInvalidPatterns)
	}

}
