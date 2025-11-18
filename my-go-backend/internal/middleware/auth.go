package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"my-go-backend/internal/database"
	"my-go-backend/internal/models"
)

/*
میدلور احراز هویت

این میدلور برای بررسی توکن JWT و احراز هویت کاربر استفاده می‌شود.
*/

// Auth میدلور احراز هویت را ایجاد می‌کند
var (
	sessionTableOnce    sync.Once
	sessionTableEnabled bool
	errSessionInactive  = errors.New("session revoked or inactive")
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("🔐 Auth Debug for path: %s\n", c.Request.URL.Path)

		// حذف مسیر هدرهای سفارشی توسعه (هماهنگ‌سازی با احراز هویت متمرکز)
		// در سیستم فعلی فقط JWT معتبر پذیرفته می‌شود

		// بررسی JWT token
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token := authHeader[7:]
			if len(token) >= 10 {
				fmt.Printf("   - 🔑 Using JWT token: %s...\n", token[:10])
			} else {
				fmt.Printf("   - 🔑 Using JWT token: %s\n", token)
			}

			// ابتدا توکن را parse می‌کنیم تا ببینیم مهمان است یا خیر
			if claims, ok := parseGuestToken(token); ok {
				path := c.Request.URL.Path
				if strings.HasPrefix(path, "/api/chat/admin") || strings.HasPrefix(path, "/api/admin/chat") {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "دسترسی مهمان به مسیرهای ادمین مجاز نیست"})
					c.Abort()
					return
				}
				// توکن مهمان فقط برای مسیرهای چت معتبر است
				if strings.HasPrefix(path, "/api/chat") {
					c.Set("is_guest", true)
					c.Set("guest_session_id", claims["session_id"])
					c.Set("guest_name", claims["full_name"])
					c.Set("guest_phone", claims["phone"])
					c.Next()
					return
				}
				// مهمان به مسیر دیگر دسترسی ندارد
				c.JSON(http.StatusUnauthorized, gin.H{"error": "دسترسی مهمان محدود به چت است"})
				c.Abort()
				return
			}

			// اگر مهمان نبود، مسیر عادی کاربر لاگین شده
			userID, session, err := validateJWTToken(token)
			if err != nil {
				fmt.Printf("   - ❌ JWT validation failed: %v\n", err)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "توکن نامعتبر است"})
				c.Abort()
				return
			}

			// بررسی وجود کاربر در دیتابیس
			var user models.User
			if err := database.GormDB.Preload("UserRole").First(&user, userID).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر یافت نشد"})
				c.Abort()
				return
			}

			// ست کردن اطلاعات کاربر در context
			c.Set("user_id", userID)
			c.Set("role_id", user.RoleID)
			c.Set("role", user.UserRole.Name)
			c.Set("username", user.Username)
			if session != nil {
				c.Set("session_token", session.SessionToken)
				c.Set("session_id", session.ID)
			}
			c.Next()
			return
		}

		// تلاش برای خواندن JWT از کوکی HttpOnly
		if cookieToken, err := c.Cookie("access_token"); err == nil && cookieToken != "" {
			fmt.Printf("   - 🍪 Using JWT cookie: %s...\n", cookieToken[:10])

			userID, session, err := validateJWTToken(cookieToken)
			if err != nil {
				fmt.Printf("   - ❌ Cookie JWT validation failed: %v\n", err)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "توکن نامعتبر است"})
				c.Abort()
				return
			}

			// بررسی وجود کاربر در دیتابیس
			var user models.User
			if err := database.GormDB.Preload("UserRole").First(&user, userID).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر یافت نشد"})
				c.Abort()
				return
			}

			// ست کردن اطلاعات کاربر در context
			c.Set("user_id", userID)
			c.Set("role_id", user.RoleID)
			c.Set("role", user.UserRole.Name)
			c.Set("username", user.Username)
			if session != nil {
				c.Set("session_token", session.SessionToken)
				c.Set("session_id", session.ID)
			}
			c.Next()
			return
		}

		// در نهایت اگر هیچ راهی احراز هویت نشد
		fmt.Printf("   - ❌ No valid authentication source, returning unauthorized\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		c.Abort()
		return
	}
}

// AuthOptional میدلور احراز هویت اختیاری: اگر کاربر لاگین بود، نقش را ست می‌کند، اگر نبود، ادامه می‌دهد
func AuthOptional() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("🔐 AuthOptional Debug for path: %s\n", c.Request.URL.Path)

		// تلاش برای خواندن JWT از کوکی HttpOnly
		if cookieToken, err := c.Cookie("access_token"); err == nil && cookieToken != "" {
			fmt.Printf("   - 🍪 JWT cookie detected: %s...\n", cookieToken[:10])

			if userID, session, err := validateJWTToken(cookieToken); err == nil {
				// خواندن کاربر از دیتابیس و ست کردن نقش
				var user models.User
				if err := database.GormDB.Preload("UserRole").First(&user, userID).Error; err == nil {
					c.Set("user_id", userID)
					c.Set("role_id", user.RoleID)
					c.Set("role", user.UserRole.Name)
					c.Set("username", user.Username)
					if session != nil {
						c.Set("session_token", session.SessionToken)
						c.Set("session_id", session.ID)
					}
					fmt.Printf("   - ✅ User authenticated (optional) - ID: %d, Role: %s\n", userID, user.UserRole.Name)
				} else {
					fmt.Printf("   - ❌ User not found in DB\n")
				}
			} else if errors.Is(err, errSessionInactive) {
				fmt.Printf("   - ⚠️ Session inactive or revoked\n")
			} else {
				fmt.Printf("   - ❌ Cookie JWT validation failed: %v\n", err)
			}
		} else {
			fmt.Printf("   - No JWT cookie, continuing as guest\n")
		}

		c.Next()
	}
}

// Admin میدلور بررسی دسترسی ادمین را ایجاد می‌کند
func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// بررسی نقش کاربر
		role, exists := c.Get("role")
		if !exists || role == "customer" {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			c.Abort()
			return
		}

		// بررسی نقش‌های مدیریتی
		roleStr := role.(string)
		adminRoles := []string{"admin", "developer"}

		isAdmin := false
		for _, adminRole := range adminRoles {
			if roleStr == adminRole {
				isAdmin = true
				break
			}
		}

		if !isAdmin {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// validateJWTToken بررسی توکن JWT و استخراج user_id
func validateJWTToken(tokenString string) (uint, *models.UserSession, error) {
	/*
		بررسی و اعتبارسنجی توکن JWT

		نکته: ابتدا از متغیر محیطی JWT_SECRET استفاده می‌کنیم (استاندارد صنعت).
		در صورت عدم وجود، از جدول تنظیمات (AuthSettings) برای backward compatibility
		و در نهایت از یک مقدار پیش‌فرض توسعه استفاده می‌کنیم.
	*/
	secretKey := ""

	// ابتدا تلاش برای خواندن از متغیر محیطی (استاندارد صنعت)
	secretKey = os.Getenv("JWT_SECRET")

	// در صورت عدم وجود، تلاش برای خواندن از جدول settings (برای backward compatibility)
	if secretKey == "" {
		var jwtSetting models.Setting
		if err := database.GormDB.Where("category = ? AND key = ?", "auth", "jwt_secret").First(&jwtSetting).Error; err == nil && jwtSetting.Value != "" {
			secretKey = jwtSetting.Value
		}
	}

	// مقدار پیش‌فرض (فقط برای محیط توسعه)
	if secretKey == "" {
		secretKey = "your-super-secret-jwt-key-change-this-in-production"
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// بررسی الگوریتم
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	}, jwt.WithLeeway(30*time.Second))

	if err != nil {
		return 0, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// استخراج user_id از claims
		if userID, exists := claims["user_id"]; exists {
			switch v := userID.(type) {
			case float64:
				uid := uint(v)
				session, err := validateActiveSession(uid, tokenString)
				if err != nil {
					return 0, nil, err
				}
				return uid, session, nil
			case int:
				uid := uint(v)
				session, err := validateActiveSession(uid, tokenString)
				if err != nil {
					return 0, nil, err
				}
				return uid, session, nil
			case int64:
				uid := uint(v)
				session, err := validateActiveSession(uid, tokenString)
				if err != nil {
					return 0, nil, err
				}
				return uid, session, nil
			default:
				return 0, nil, fmt.Errorf("invalid user_id type in token")
			}
		}
		return 0, nil, fmt.Errorf("user_id not found in token")
	}

	return 0, nil, fmt.Errorf("invalid token")
}

// parseGuestToken بررسی می‌کند آیا توکن دارای claim guest=true است و اگر بله Claims را برمی‌گرداند
func parseGuestToken(tokenString string) (jwt.MapClaims, bool) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-super-secret-jwt-key-change-this-in-production"
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}
	if g, _ := claims["guest"].(bool); g {
		return claims, true
	}
	return nil, false
}

func validateActiveSession(userID uint, tokenString string) (*models.UserSession, error) {
	if !isSessionTableEnabled() {
		return nil, nil
	}

	if database.GormDB == nil {
		return nil, nil
	}

	var session models.UserSession
	err := database.GormDB.
		Where("session_token = ? AND is_active = ?", tokenString, true).
		Where("expires_at > ?", time.Now()).
		First(&session).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errSessionInactive
		}
		return nil, err
	}

	if session.UserID != userID {
		return nil, errSessionInactive
	}

	touchSession(session.ID)

	return &session, nil
}

func touchSession(sessionID uint) {
	if sessionID == 0 || database.GormDB == nil {
		return
	}

	_ = database.GormDB.Model(&models.UserSession{}).
		Where("id = ?", sessionID).
		Updates(map[string]interface{}{
			"last_activity": time.Now(),
		}).Error
}

func isSessionTableEnabled() bool {
	sessionTableOnce.Do(func() {
		if database.GormDB == nil {
			sessionTableEnabled = false
			return
		}
		sessionTableEnabled = database.GormDB.Migrator().HasTable(&models.UserSession{})
	})
	return sessionTableEnabled
}
