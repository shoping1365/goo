package middleware

import (
	"fmt"
	"net/http"

	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware میدلور مدیریت خطاهای سراسری
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			// خطای panic به صورت string
			errorResponse := utils.NewErrorResponse(
				utils.ErrorTypeServer,
				"خطای سرور",
				"خطای غیرمنتظره رخ داده است",
				err,
				http.StatusInternalServerError,
			)
			c.JSON(http.StatusInternalServerError, errorResponse)
		} else if err, ok := recovered.(error); ok {
			// خطای panic به صورت error
			errorResponse := utils.NewErrorResponse(
				utils.ErrorTypeServer,
				"خطای سرور",
				"خطای غیرمنتظره رخ داده است",
				err.Error(),
				http.StatusInternalServerError,
			)
			c.JSON(http.StatusInternalServerError, errorResponse)
		} else {
			// سایر انواع panic
			errorResponse := utils.NewErrorResponse(
				utils.ErrorTypeServer,
				"خطای سرور",
				"خطای غیرمنتظره رخ داده است",
				fmt.Sprintf("%v", recovered),
				http.StatusInternalServerError,
			)
			c.JSON(http.StatusInternalServerError, errorResponse)
		}

		// لاگ کردن stack trace برای debugging
		// PANIC logged
	})
}

// ErrorResponseMiddleware میدلور برای تبدیل خطاها به پاسخ‌های استاندارد
func ErrorResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// بررسی خطاهای موجود در context
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// خطای عمومی
			errorResponse := utils.NewErrorResponse(
				utils.ErrorTypeServer,
				"خطای سرور",
				"خطای غیرمنتظره رخ داده است",
				err.Error(),
				http.StatusInternalServerError,
			)

			// ارسال پاسخ خطا
			c.JSON(errorResponse.StatusCode, errorResponse)
			c.Abort()
		}
	}
}

// LoggingMiddleware میدلور برای لاگ کردن درخواست‌ها و خطاها
func LoggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// لاگ کردن درخواست‌های موفق
		if param.StatusCode < 400 {
			return fmt.Sprintf("✅ %s | %d | %s | %s | %s\n",
				param.Method,
				param.StatusCode,
				param.Path,
				param.ClientIP,
				param.Latency,
			)
		}

		// لاگ کردن خطاها با جزئیات بیشتر
		return fmt.Sprintf("❌ %s | %d | %s | %s | %s | %s\n",
			param.Method,
			param.StatusCode,
			param.Path,
			param.ClientIP,
			param.Latency,
			param.ErrorMessage,
		)
	})
}

// CustomLogger میدلور لاگینگ سفارشی برای نمایش لاگ‌ها
func CustomLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// لاگ کردن درخواست‌های موفق
		if param.StatusCode < 400 {
			return fmt.Sprintf("✅ %s | %d | %s | %s | %s | %s\n",
				param.TimeStamp.Format("2006/01/02 15:04:05"),
				param.StatusCode,
				param.Method,
				param.Path,
				param.ClientIP,
				param.Latency,
			)
		}

		// لاگ کردن خطاها با جزئیات بیشتر
		return fmt.Sprintf("❌ %s | %d | %s | %s | %s | %s | %s\n",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			param.StatusCode,
			param.Method,
			param.Path,
			param.ClientIP,
			param.Latency,
			param.ErrorMessage,
		)
	})
}

// ValidationErrorMiddleware میدلور برای مدیریت خطاهای اعتبارسنجی
func ValidationErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// بررسی خطاهای اعتبارسنجی
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if err.Type == gin.ErrorTypeBind {
					// خطای binding (validation)
					errorResponse := utils.NewErrorResponse(
						utils.ErrorTypeValidation,
						"خطای اعتبارسنجی",
						"اطلاعات ارسالی صحیح نیست",
						err.Error(),
						http.StatusBadRequest,
					)
					c.JSON(http.StatusBadRequest, errorResponse)
					c.Abort()
					return
				}
			}
		}
	}
}
