package utils

import (
	"fmt"
	"net/http"
	"strings"
)

// ErrorType نوع خطا را تعریف می‌کند
type ErrorType string

const (
	ErrorTypeValidation  ErrorType = "VALIDATION_ERROR"
	ErrorTypeAuth        ErrorType = "AUTH_ERROR"
	ErrorTypePermission  ErrorType = "PERMISSION_ERROR"
	ErrorTypeNotFound    ErrorType = "NOT_FOUND_ERROR"
	ErrorTypeDatabase    ErrorType = "DB_ERROR"
	ErrorTypeExternalAPI ErrorType = "EXTERNAL_API_ERROR"
	ErrorTypeServer      ErrorType = "SERVER_ERROR"
	ErrorTypeNetwork     ErrorType = "NETWORK_ERROR"
	ErrorTypeTimeout     ErrorType = "TIMEOUT_ERROR"
	ErrorTypeRateLimit   ErrorType = "RATE_LIMIT_ERROR"
	ErrorTypeBusiness    ErrorType = "BUSINESS_ERROR"
)

// ErrorResponse ساختار پاسخ خطا
type ErrorResponse struct {
	Type        ErrorType `json:"type"`
	Title       string    `json:"title"`
	Message     string    `json:"message"`
	Details     string    `json:"details,omitempty"`
	StatusCode  int       `json:"status_code"`
	UserMessage string    `json:"user_message"` // پیام قابل نمایش برای کاربر
}

// NewErrorResponse ایجاد یک پاسخ خطای جدید
func NewErrorResponse(errorType ErrorType, title, message, details string, statusCode int) *ErrorResponse {
	return &ErrorResponse{
		Type:        errorType,
		Title:       title,
		Message:     message,
		Details:     details,
		StatusCode:  statusCode,
		UserMessage: generateUserMessage(errorType, message),
	}
}

// generateUserMessage تولید پیام کاربرپسند بر اساس نوع خطا
func generateUserMessage(errorType ErrorType, message string) string {
	switch errorType {
	case ErrorTypeValidation:
		return "اطلاعات وارد شده صحیح نیست. لطفاً دوباره بررسی کنید."
	case ErrorTypeAuth:
		return "احراز هویت ناموفق بود. لطفاً دوباره وارد شوید."
	case ErrorTypePermission:
		return "شما دسترسی لازم برای انجام این عملیات را ندارید."
	case ErrorTypeNotFound:
		return "مورد درخواستی یافت نشد."
	case ErrorTypeDatabase:
		return "خطا در دسترسی به اطلاعات. لطفاً بعداً تلاش کنید."
	case ErrorTypeExternalAPI:
		return "خطا در ارتباط با سرویس خارجی. لطفاً بعداً تلاش کنید."
	case ErrorTypeServer:
		return "خطای سرور. لطفاً بعداً تلاش کنید."
	case ErrorTypeNetwork:
		return "خطا در اتصال شبکه. لطفاً اتصال اینترنت خود را بررسی کنید."
	case ErrorTypeTimeout:
		return "عملیات زمان زیادی طول کشید. لطفاً دوباره تلاش کنید."
	case ErrorTypeRateLimit:
		return "تعداد درخواست‌های شما بیش از حد مجاز است. لطفاً کمی صبر کنید."
	case ErrorTypeBusiness:
		return message // برای خطاهای تجاری، پیام اصلی را نمایش بده
	default:
		return "خطای غیرمنتظره رخ داده است. لطفاً بعداً تلاش کنید."
	}
}

// HTTP Status Code به Error Type تبدیل
func StatusCodeToErrorType(statusCode int) ErrorType {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return ErrorTypeBusiness // موفق
	case statusCode == 400:
		return ErrorTypeValidation
	case statusCode == 401:
		return ErrorTypeAuth
	case statusCode == 403:
		return ErrorTypePermission
	case statusCode == 404:
		return ErrorTypeNotFound
	case statusCode == 408:
		return ErrorTypeTimeout
	case statusCode == 429:
		return ErrorTypeRateLimit
	case statusCode >= 500 && statusCode < 600:
		return ErrorTypeServer
	default:
		return ErrorTypeBusiness
	}
}

// External API Error Handler برای خطاهای API های خارجی
func HandleExternalAPIError(serviceName, endpoint string, statusCode int, responseBody string) *ErrorResponse {
	errorType := StatusCodeToErrorType(statusCode)

	var title, message string

	switch statusCode {
	case 401:
		title = fmt.Sprintf("خطای احراز هویت %s", serviceName)
		message = fmt.Sprintf("کلید API %s نامعتبر است. لطفاً کلید API صحیح را وارد کنید.", serviceName)
	case 403:
		title = fmt.Sprintf("خطای دسترسی %s", serviceName)
		message = fmt.Sprintf("دسترسی به %s محدود شده است. لطفاً تنظیمات حساب خود را بررسی کنید.", serviceName)
	case 404:
		title = fmt.Sprintf("خطای آدرس %s", serviceName)
		message = fmt.Sprintf("آدرس API %s نامعتبر است. لطفاً URL صحیح را وارد کنید.", serviceName)
	case 429:
		title = fmt.Sprintf("محدودیت نرخ درخواست %s", serviceName)
		message = fmt.Sprintf("تعداد درخواست‌های شما به %s بیش از حد مجاز است. لطفاً کمی صبر کنید.", serviceName)
	case 500, 502, 503, 504:
		title = fmt.Sprintf("خطای سرور %s", serviceName)
		message = fmt.Sprintf("سرور %s در دسترس نیست. لطفاً بعداً تلاش کنید.", serviceName)
	default:
		title = fmt.Sprintf("خطا در ارتباط با %s", serviceName)
		message = fmt.Sprintf("خطا در اتصال به %s. کد وضعیت: %d", serviceName, statusCode)
	}

	return NewErrorResponse(errorType, title, message, responseBody, statusCode)
}

// Database Error Handler
func HandleDatabaseError(operation string, err error) *ErrorResponse {
	errorMsg := err.Error()

	var title, message string

	if strings.Contains(errorMsg, "duplicate key") {
		title = "خطای تکراری"
		message = "این مورد قبلاً در سیستم ثبت شده است."
	} else if strings.Contains(errorMsg, "foreign key") {
		title = "خطای وابستگی"
		message = "این مورد به موارد دیگری وابسته است و قابل حذف نیست."
	} else if strings.Contains(errorMsg, "not found") {
		title = "مورد یافت نشد"
		message = "مورد درخواستی در پایگاه داده یافت نشد."
	} else if strings.Contains(errorMsg, "connection") {
		title = "خطای اتصال پایگاه داده"
		message = "خطا در اتصال به پایگاه داده. لطفاً بعداً تلاش کنید."
	} else {
		title = "خطای پایگاه داده"
		message = fmt.Sprintf("خطا در عملیات %s: %s", operation, errorMsg)
	}

	return NewErrorResponse(ErrorTypeDatabase, title, message, errorMsg, http.StatusInternalServerError)
}

// Validation Error Handler
func HandleValidationError(field, message string) *ErrorResponse {
	title := "خطای اعتبارسنجی"
	details := fmt.Sprintf("فیلد '%s': %s", field, message)

	return NewErrorResponse(ErrorTypeValidation, title, message, details, http.StatusBadRequest)
}

// Business Logic Error Handler
func HandleBusinessError(title, message string) *ErrorResponse {
	return NewErrorResponse(ErrorTypeBusiness, title, message, "", http.StatusBadRequest)
}

// Network Error Handler
func HandleNetworkError(serviceName string, err error) *ErrorResponse {
	errorMsg := err.Error()

	var title, message string

	if strings.Contains(errorMsg, "timeout") {
		title = fmt.Sprintf("خطای تایم‌اوت %s", serviceName)
		message = fmt.Sprintf("اتصال به %s زمان زیادی طول کشید. لطفاً دوباره تلاش کنید.", serviceName)
	} else if strings.Contains(errorMsg, "connection refused") {
		title = fmt.Sprintf("خطای اتصال %s", serviceName)
		message = fmt.Sprintf("نمی‌توان به %s متصل شد. لطفاً اتصال شبکه خود را بررسی کنید.", serviceName)
	} else {
		title = fmt.Sprintf("خطای شبکه %s", serviceName)
		message = fmt.Sprintf("خطا در اتصال به %s: %s", serviceName, errorMsg)
	}

	return NewErrorResponse(ErrorTypeNetwork, title, message, errorMsg, http.StatusServiceUnavailable)
}
