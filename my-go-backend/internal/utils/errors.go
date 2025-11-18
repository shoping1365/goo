package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
پکیج ابزارها

این پکیج شامل توابع کمکی و ابزارهای مورد نیاز در سراسر برنامه است.
*/

// SuccessResponse ساختار پاسخ موفقیت
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SendError خطا را در قالب ApiError واحد ارسال می‌کند.
// تمام هندلرهای جدید بهتر است مستقیماً از utils.New + c.AbortWithStatusJSON استفاده کنند،
// اما این هلسپر برای سازگاری قدیمی نگه داشته شده است.
func SendError(c *gin.Context, status int, code string, message string, details interface{}) {
	c.AbortWithStatusJSON(status, New(code, message, details))
}

// SendSuccess پاسخ موفقیت را ارسال می‌کند
func SendSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

// BadRequest پاسخ خطای درخواست نامعتبر را ارسال می‌کند
func BadRequest(c *gin.Context, message string, details interface{}) {
	SendError(c, http.StatusBadRequest, "BAD_REQUEST", message, details)
}

// Unauthorized پاسخ خطای عدم احراز هویت را ارسال می‌کند
func Unauthorized(c *gin.Context, message string, details interface{}) {
	SendError(c, http.StatusUnauthorized, "UNAUTHORIZED", message, details)
}

// Forbidden پاسخ خطای عدم دسترسی را ارسال می‌کند
func Forbidden(c *gin.Context, message string, details interface{}) {
	SendError(c, http.StatusNotFound, "NOT_FOUND", message, details)
}

// NotFound پاسخ خطای مورد یافت نشد را ارسال می‌کند
func NotFound(c *gin.Context, message string, details interface{}) {
	SendError(c, http.StatusNotFound, "NOT_FOUND", message, details)
}

// InternalServerError پاسخ خطای سرور را ارسال می‌کند
func InternalServerError(c *gin.Context, message string, details interface{}) {
	SendError(c, http.StatusInternalServerError, "INTERNAL_ERROR", message, details)
}
