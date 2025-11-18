package router

import (
	"my-go-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupCartRoutes تنظیم مسیرهای سبد خرید
func SetupCartRoutes(r *gin.RouterGroup, cartHandler *handlers.CartHandler) {
	cartGroup := r.Group("/cart")
	// احراز هویت اختیاری: اگر کاربر لاگین باشد user_id در context ست می‌شود
	// cartGroup.Use(middleware.AuthOptional()) // احراز هویت غیرفعال شده است
	{
		// دریافت سبد خرید
		cartGroup.GET("", cartHandler.GetCart)

		// افزودن محصول به سبد خرید
		cartGroup.POST("/add", cartHandler.AddToCart)

		// ایجاد session سبد خرید برای مهمان‌ها
		cartGroup.POST("/create-session", cartHandler.CreateSession)

		// انتقال به خرید بعدی / بازگرداندن
		cartGroup.PUT("/move-next", cartHandler.MoveToNext)
		cartGroup.PUT("/move-cart", cartHandler.MoveToCart)

		// به‌روزرسانی آیتم سبد خرید
		cartGroup.PUT("/update", cartHandler.UpdateCartItem)

		// حذف از سبد خرید
		cartGroup.DELETE("/remove", cartHandler.RemoveFromCart)

		// پاک کردن سبد خرید
		cartGroup.DELETE("/clear", cartHandler.ClearCart)

		// دریافت تعداد آیتم‌های سبد خرید
		cartGroup.GET("/count", cartHandler.GetCartCount)
	}
}
