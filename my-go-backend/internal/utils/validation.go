package utils

import (
	"encoding/json"
	"mime/multipart"
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// IsValidEmail بررسی معتبر بودن ایمیل
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// IsValidMobile بررسی معتبر بودن شماره موبایل ایرانی
func IsValidMobile(mobile string) bool {
	// حذف فاصله و کاراکترهای غیرضروری
	mobile = strings.ReplaceAll(mobile, " ", "")
	mobile = strings.ReplaceAll(mobile, "-", "")
	mobile = strings.ReplaceAll(mobile, "(", "")
	mobile = strings.ReplaceAll(mobile, ")", "")
	
	// بررسی الگوی شماره موبایل ایرانی
	mobileRegex := regexp.MustCompile(`^(\+98|0098|98|0)?9[0-9]{9}$`)
	return mobileRegex.MatchString(mobile)
}

// IsValidURL بررسی معتبر بودن URL
func IsValidURL(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	return err == nil
}

// IsValidColor بررسی معتبر بودن رنگ (hex یا نام رنگ)
func IsValidColor(color string) bool {
	// بررسی رنگ hex
	if strings.HasPrefix(color, "#") {
		hexRegex := regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)
		return hexRegex.MatchString(color)
	}

	// بررسی نام رنگ‌های CSS
	validColors := []string{
		"red", "green", "blue", "yellow", "black", "white", "gray", "grey",
		"orange", "purple", "pink", "brown", "cyan", "magenta", "lime",
		"navy", "olive", "teal", "silver", "maroon", "fuchsia", "aqua",
		"transparent", "currentColor", "inherit", "initial", "unset",
	}

	color = strings.ToLower(color)
	for _, validColor := range validColors {
		if color == validColor {
			return true
		}
	}

	return false
}

// StringToInt تبدیل رشته به عدد
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// ParseInt تبدیل امن رشته به عدد با مقدار پیش‌فرض
func ParseInt(s string) int {
	if val, err := strconv.Atoi(s); err == nil {
		return val
	}
	return 0
}

// BoolToString تبدیل boolean به string
func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// Contains بررسی وجود یک عنصر در slice
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// NewError ایجاد خطای جدید
func NewError(message string) error {
	return &ValidationError{Message: message}
}

// ValidationError نوع خطای اعتبارسنجی
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// IsValidImageType بررسی معتبر بودن نوع فایل تصویر
func IsValidImageType(header *multipart.FileHeader) bool {
	contentType := header.Header.Get("Content-Type")
	validTypes := []string{"image/jpeg", "image/jpg", "image/png", "image/gif", "image/webp"}
	for _, validType := range validTypes {
		if contentType == validType {
			return true
		}
	}
	return false
}

// ToJSON تبدیل ساختار داده به JSON
func ToJSON(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
