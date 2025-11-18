package router

import (
	"github.com/gin-gonic/gin"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
)

// RegisterUserVerificationRoutes - ثبت روت‌های احراز هویت کاربران
func RegisterUserVerificationRoutes(public *gin.RouterGroup, verificationHandler *handlers.UserVerificationHandler) {
	// مسیرهای عمومی (نیاز به احراز هویت دارند)
	verifications := public.Group("/verifications")
	verifications.Use(middleware.Auth()) // همه مسیرها نیاز به احراز هویت دارند
	{
		// ایجاد درخواست احراز هویت جدید
		// POST /api/verifications
		verifications.POST("", verificationHandler.CreateVerificationRequest)

		// دریافت اطلاعات یک درخواست
		// GET /api/verifications/:id
		verifications.GET("/:id", verificationHandler.GetVerificationByID)
	}

	// مسیرهای مربوط به کاربر
	users := public.Group("/users")
	users.Use(middleware.Auth())
	{
		// دریافت لیست درخواست‌های احراز هویت یک کاربر
		// GET /api/users/:id/verifications
		users.GET(":id/verifications", verificationHandler.GetUserVerifications)

		// دریافت فیلدهای تایید شده یک کاربر
		// GET /api/users/:id/verified-fields
		users.GET(":id/verified-fields", verificationHandler.GetVerifiedFields)
	}

	// مسیرهای ادمین (نیاز به مجوز مدیریت کاربران)
	admin := public.Group("/admin/verifications")
	admin.Use(middleware.Auth())
	admin.Use(middleware.RequireAnyPermission("user.view", "users_view", "verification.manage"))
	{
		// دریافت تمام درخواست‌ها
		// GET /api/admin/verifications?status=pending&page=1&page_size=20
		admin.GET("", verificationHandler.GetAllVerifications)

		// دریافت درخواست‌های در انتظار
		// GET /api/admin/verifications/pending
		admin.GET("/pending", verificationHandler.GetPendingVerifications)

		// دریافت تعداد درخواست‌های در انتظار
		// GET /api/admin/verifications/pending/count
		admin.GET("/pending/count", verificationHandler.GetPendingCount)

		// تایید درخواست (فاز 1)
		// POST /api/admin/verifications/:id/verify
		admin.POST("/:id/verify", middleware.RequireAnyPermission("user.update", "users_update", "verification.manage"), verificationHandler.VerifyRequest)

		// رد درخواست (فاز 1)
		// POST /api/admin/verifications/:id/reject
		admin.POST("/:id/reject", middleware.RequireAnyPermission("user.update", "users_update", "verification.manage"), verificationHandler.RejectRequest)

		// حذف درخواست
		// DELETE /api/admin/verifications/:id
		admin.DELETE("/:id", middleware.RequireAnyPermission("user.delete", "users_delete", "verification.manage"), verificationHandler.DeleteVerification)
	}
}
