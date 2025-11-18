package config

import (
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SessionConfig تنظیمات session
type SessionConfig struct {
	Store   sessions.Store
	Options sessions.Options
}

// NewSessionConfig ایجاد تنظیمات session جدید
func NewSessionConfig() (*SessionConfig, error) {
	var store sessions.Store

	// فعلاً فقط از Cookie استفاده می‌کنیم (Redis نیاز به dependency اضافی دارد)
	store = cookie.NewStore([]byte(getSessionSecret()))

	// تنظیمات session
	options := sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(30 * 24 * time.Hour.Seconds()), // 30 روز
		SameSite: 1,                                  // Lax
	}

	// در production، secure را فعال کن
	if os.Getenv("GIN_MODE") == "release" {
		options.Secure = true
	}

	store.Options(options)

	return &SessionConfig{
		Store:   store,
		Options: options,
	}, nil
}

// getSessionSecret دریافت secret key برای session
func getSessionSecret() string {
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		secret = "your-secret-key-change-in-production"
	}
	return secret
}

// GetSessionMiddleware ایجاد middleware session
func GetSessionMiddleware() gin.HandlerFunc {
	config, err := NewSessionConfig()
	if err != nil {
		panic(err)
	}

	return sessions.Sessions("mysession", config.Store)
}
