package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims - محتویات JWT Token
type JWTClaims struct {
	UserID      uint     `json:"user_id"`
	Mobile      string   `json:"mobile"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

// GenerateAccessToken - ایجاد Access Token (15 دقیقه)
func GenerateAccessToken(userID uint, mobile string, role string, permissions []string, secret string) (string, error) {
	if secret == "" {
		return "", errors.New("JWT_SECRET not configured")
	}

	claims := &JWTClaims{
		UserID:      userID,
		Mobile:      mobile,
		Role:        role,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-nuxt",
			Subject:   "auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken - ایجاد Refresh Token (30 روز)
// نکته: توکن واقعی تولید می‌شود، بعد bcrypt hash میکنیم قبل ذخیره‌سازی
func GenerateRefreshToken(userID uint, secret string) (token string, expiresAt time.Time, err error) {
	if secret == "" {
		return "", time.Time{}, errors.New("JWT_SECRET not configured")
	}

	expiresAt = time.Now().Add(30 * 24 * time.Hour)

	claims := &JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-nuxt",
			Subject:   "refresh",
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

// ValidateAccessToken - بررسی اعتبار Access Token
func ValidateAccessToken(tokenString string, secret string) (*JWTClaims, error) {
	if secret == "" {
		return nil, errors.New("JWT_SECRET not configured")
	}

	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// بررسی signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// ValidateRefreshToken - بررسی اعتبار Refresh Token
func ValidateRefreshToken(tokenString string, secret string) (*JWTClaims, error) {
	if secret == "" {
		return nil, errors.New("JWT_SECRET not configured")
	}

	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// بررسی اینکه توکن refresh token باشد
	if claims.Subject != "refresh" {
		return nil, errors.New("not a refresh token")
	}

	return claims, nil
}
