package router

import (
	"time"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

/*
Router احراز هویت یکپارچه
-----------------------------------------------------------------
این router مسیرهای مربوط به سیستم احراز هویت یکپارچه را تعریف می‌کند.
تمام مسیرها با سیستم JWT موجود سازگار هستند.
*/

// SetupUnifiedAuthRoutes تنظیم مسیرهای احراز هویت یکپارچه
func SetupUnifiedAuthRoutes(r *gin.RouterGroup, handler *handlers.UnifiedAuthHandler) {
	auth := r.Group("/auth")
	{
		// مسیرهای عمومی (بدون نیاز به احراز هویت)
		auth.POST("/check-user", handler.CheckUser) // بررسی وجود کاربر

		// مسیرهای حساس با Rate Limiting
		// محدودیت: 10 درخواست در دقیقه برای هر IP
		auth.POST("/send-otp", middleware.RateLimitWith(10, time.Minute, "send-otp"), handler.SendOTP)
		auth.POST("/verify-otp", middleware.RateLimitWith(10, time.Minute, "verify-otp"), handler.VerifyOTPAndLogin)
		auth.POST("/login-password", middleware.RateLimitWith(10, time.Minute, "login"), handler.LoginWithPassword)

		// مسیرهای عمومی دیگر
		auth.GET("/methods", handler.GetAuthMethods)  // دریافت روش‌های احراز هویت
		auth.GET("/csrf-token", handler.GetCSRFToken) // دریافت CSRF Token

		// مسیرهای نیازمند احراز هویت
		authRequired := auth.Group("")
		authRequired.Use(middleware.Auth())
		{
			authRequired.GET("/me", handler.GetCurrentUser)              // دریافت اطلاعات کاربر جاری
			authRequired.GET("/permissions", handler.GetUserPermissions) // دریافت دسترسی‌های کاربر
			authRequired.POST("/set-password", handler.SetPassword)      // تنظیم پسورد
			authRequired.POST("/logout", handler.Logout)                 // خروج از حساب
			authRequired.POST("/refresh", handler.RefreshToken)          // تجدید توکن
		}
	}

	// مسیرهای مدیریت (فقط ادمین)
	admin := r.Group("/admin/auth")
	admin.Use(middleware.Auth(), middleware.AdminMiddleware())
	{
		admin.GET("/login-attempts", handler.AdminGetLoginAttempts) // فهرست تلاش‌های ورود
		admin.POST("/unlock/:mobile", handler.AdminUnlockAccount)   // رفع قفل حساب
	}
}

/*
نقشه مسیرها:

احراز هویت یکپارچه:
├── POST /api/auth/check-mobile        → بررسی موبایل + ارسال OTP
├── POST /api/auth/verify-otp          → تایید OTP + ورود/ثبت‌نام
├── POST /api/auth/login-password      → ورود با پسورد
├── POST /api/auth/unified-login       → ورود یکپارچه (انتخاب روش)
├── GET  /api/auth/methods             → روش‌های احراز هویت موجود
├── POST /api/auth/set-password        → تنظیم پسورد (نیاز به احراز)
└── GET  /api/auth/status              → وضعیت احراز هویت

مدیریت (ادمین):
├── GET  /api/admin/auth/login-attempts → فهرست تلاش‌های ورود
└── POST /api/admin/auth/unlock/:mobile → رفع قفل حساب

سازگاری:
├── POST /api/send-otp                 → سازگاری با سیستم قدیمی
└── POST /api/verify-otp               → سازگاری با سیستم قدیمی

نکات مهم:
- تمام مسیرها از کوکی‌های HttpOnly JWT استفاده می‌کنند
- سازگاری کامل با سیستم احراز هویت موجود
- پشتیبانی از middleware های امنیتی موجود
- امکان استفاده همزمان با سیستم JWT فعلی
*/
