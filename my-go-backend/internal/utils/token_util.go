package utils

// این فایل util برای تولید توکن میهمان (guest) است و حاوی تابعی ساده برای صدور JWT
// با تاریخ انقضای یک هفته می‌باشد. این توکن صرفاً برای مسیرهای چت معتبر است.

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateGuestToken تولید توکن JWT مهمان با sessionID، نام و شماره تلفن
func GenerateGuestToken(sessionID, fullName, phone string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-super-secret-jwt-key-change-this-in-production"
	}
	claims := jwt.MapClaims{
		"guest":      true,
		"session_id": sessionID,
		"full_name":  fullName,
		"phone":      phone,
		"exp":        time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":        time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
