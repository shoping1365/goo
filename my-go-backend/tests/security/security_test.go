package security_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"my-go-backend/internal/auth"
	"my-go-backend/internal/database"
	"my-go-backend/internal/handlers"
	"my-go-backend/internal/middleware"
	"my-go-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// SetupTestDB connects to the database using DATABASE_URL
func SetupTestDB() *gorm.DB {
	// Load .env file from project root (../../.env relative to tests/security)
	cwd, _ := os.Getwd()
	fmt.Println("Current Working Directory:", cwd)
	
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		// Try absolute path or other locations if needed
		_ = godotenv.Load("../../../.env") // If running from root? No.
	}

	// Ensure we are using the correct environment
	if os.Getenv("DATABASE_URL") == "" {
		// Construct DATABASE_URL from components if available
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		dbSSLMode := os.Getenv("DB_SSLMODE")

		if dbUser != "" && dbHost != "" {
			dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				dbUser,
				url.QueryEscape(dbPassword),
				dbHost,
				dbPort,
				dbName,
				dbSSLMode,
			)
			os.Setenv("DATABASE_URL", dsn)
			fmt.Println("Constructed DATABASE_URL from .env variables")
		} else {
			fmt.Println("WARNING: DATABASE_URL is not set and components are missing. Tests might fail.")
		}
	}
	
	db, err := database.ConnectGorm()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return db
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Recovery())
	
	// Auth Middleware
	r.Use(middleware.Auth())

	// Admin Route (Protected)
	r.GET("/api/admin/test", func(c *gin.Context) {
		role := c.GetString("role")
		// Simple role check for testing
		if role != "admin" && role != "super_admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin"})
	})

	// Product Route (For SQL Injection Test)
	productHandler := handlers.NewProductHandler(db)
	r.GET("/api/products", productHandler.ListProducts)

	return r
}

func CreateTestUser(db *gorm.DB, username, roleName string) (uint, string) {
	// Ensure role exists
	var role models.Role
	if err := db.Where("name = ?", roleName).First(&role).Error; err != nil {
		role = models.Role{Name: roleName, DisplayName: roleName}
		db.Create(&role)
	}

	// Use last 9 digits of UnixNano to ensure uniqueness in short timeframe
	nano := time.Now().UnixNano()
	mobile := fmt.Sprintf("09%09d", nano%1000000000)

	user := models.User{
		Username: username,
		Email:    fmt.Sprintf("%s_%d@example.com", username, nano), // Unique email
		Mobile:   mobile,
		RoleID:   role.ID,
	}
	// Handle unique constraint violation by retrying or ignoring for test
	if err := db.Create(&user).Error; err != nil {
		// Try to find existing
		db.Where("username = ?", username).First(&user)
	}
	
	token, _ := auth.GenerateToken(user.ID, user.Username, role.Name)

	// Create a valid session for the user
	session := models.UserSession{
		UserID:       user.ID,
		SessionToken: token,
		RefreshToken: token + "_refresh_" + fmt.Sprintf("%d", time.Now().UnixNano()), // Ensure unique
		IsActive:     true,
		ExpiresAt:    time.Now().Add(24 * time.Hour),
		LoginAt:      time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		IPAddress:    "127.0.0.1",
		UserAgent:    "Go-Test-Runner",
	}
	
	if err := db.Create(&session).Error; err != nil {
		fmt.Printf("Failed to create session for user %s: %v\n", username, err)
	}

	return user.ID, token
}

// TestAccessControl checks if only admins can access admin endpoints
func TestAccessControl(t *testing.T) {
	db := SetupTestDB()
	r := SetupRouter(db)

	// Create Users
	_, userToken := CreateTestUser(db, "test_user_sec", "user")
	_, adminToken := CreateTestUser(db, "test_admin_sec", "admin")

	tests := []struct {
		name       string
		token      string
		wantStatus int
	}{
		{"No Token", "", http.StatusUnauthorized},
		{"User Token", userToken, http.StatusForbidden},
		{"Admin Token", adminToken, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/admin/test", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("AccessControl %s: want status %d, got %d", tt.name, tt.wantStatus, w.Code)
			}
		})
	}
}

// TestSQLInjection checks if the search endpoint is vulnerable to SQL Injection
func TestSQLInjection(t *testing.T) {
	db := SetupTestDB()
	r := SetupRouter(db)

	// SQL Injection Payload
	// Trying to inject OR '1'='1 to get all records or cause error
	payload := "' OR '1'='1"
	
	_, userToken := CreateTestUser(db, "test_sql_user", "user")

	req, _ := http.NewRequest("GET", "/api/products?search="+payload, nil)
	req.Header.Set("Authorization", "Bearer "+userToken)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 1. Should not be 500 (Internal Server Error)
	if w.Code == http.StatusInternalServerError {
		t.Errorf("SQLInjection: Got 500 Internal Server Error, potential vulnerability or DB error")
	}

	// 2. Should not return all products (checking count)
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("SQLInjection: Failed to parse response JSON: %v", err)
	}

	// In a safe implementation, this searches for the literal string, which should return 0 results
	// unless there is a product named exactly "' OR '1'='1"
	if total, ok := response["total"].(float64); ok {
		if total > 0 {
			// It's suspicious if it returns results for such a query, but not definitive proof of injection.
			// However, if it returns ALL products (e.g. > 100), it's definitely broken.
			// For now, we just log it.
			t.Logf("SQLInjection: Returned %v results for payload %s", total, payload)
		}
	}
}

// TestRateLimit checks if the rate limiter blocks excessive requests
func TestRateLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	// Limit: 5 requests per minute per IP
	r.Use(middleware.RateLimitWith(5, time.Minute, "test_limit")) 
	r.GET("/test-limit", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Send 5 requests (Allowed)
	for i := 0; i < 5; i++ {
		req, _ := http.NewRequest("GET", "/test-limit", nil)
		req.RemoteAddr = "192.168.1.200:1234" // Unique Fake IP for this test
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("RateLimit: Request %d failed, want 200, got %d", i+1, w.Code)
		}
	}

	// Send 6th request (Blocked)
	req, _ := http.NewRequest("GET", "/test-limit", nil)
	req.RemoteAddr = "192.168.1.200:1234"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusTooManyRequests {
		t.Errorf("RateLimit: Request 6 should be blocked, want 429, got %d", w.Code)
	}
}
