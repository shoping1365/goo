package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateSessionID ایجاد شناسه نشست تصادفی
func GenerateSessionID() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		// در صورت خطا، یک شناسه پیش‌فرض برمی‌گردانیم
		return "default-session-id"
	}
	return hex.EncodeToString(bytes)
}
