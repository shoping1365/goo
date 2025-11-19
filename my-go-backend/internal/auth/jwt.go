package auth

import (
	"errors"
	"os"
	"time"

	"my-go-backend/internal/database"
	"my-go-backend/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

/*
پکیج احراز هویت

این پکیج برای مدیریت توکن‌های JWT و احراز هویت کاربران استفاده می‌شود.
*/

var (
	// کلید مخفی برای امضای توکن (پیش‌فرض)
	secretKey = []byte("your-secret-key")

	// خطاهای رایج
	ErrInvalidToken = errors.New("توکن نامعتبر است")
	ErrExpiredToken = errors.New("توکن منقضی شده است")
)

// getSecretKey کلید مخفی JWT را از تنظیمات یا متغیر محیطی می‌خواند
func getSecretKey() []byte {
	// اگر اتصال GormDB مقداردهی نشده باشد، مستقیماً از env یا مقدار پیش‌فرض استفاده می‌کنیم
	if database.GormDB == nil {
		if envKey := os.Getenv("JWT_SECRET"); envKey != "" {
			return []byte(envKey)
		}
		return secretKey
	}

	// تلاش برای خواندن از جدول settings (دسته auth)
	var jwtSetting models.Setting
	if err := database.GormDB.Where("category = ? AND key = ?", "auth", "jwt_secret").First(&jwtSetting).Error; err == nil && jwtSetting.Value != "" {
		return []byte(jwtSetting.Value)
	}

	// تلاش برای خواندن از متغیر محیطی
	if envKey := os.Getenv("JWT_SECRET"); envKey != "" {
		return []byte(envKey)
	}

	// استفاده از کلید مخفی ثابت (پیش‌فرض)
	return secretKey
}

// Claims ساختار اطلاعات ذخیره شده در توکن
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken یک توکن JWT جدید ایجاد می‌کند
func GenerateToken(userID uint, username, role string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7 روز
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getSecretKey())
}

// ValidateToken اعتبار توکن را بررسی می‌کند
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getSecretKey(), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// RefreshToken یک توکن جدید با اطلاعات توکن قبلی ایجاد می‌کند
func RefreshToken(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	return GenerateToken(claims.UserID, claims.Username, claims.Role)
}
