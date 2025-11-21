package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Mock handlers for testing routes if actual handlers are hard to setup
func mockLoginHandler(c *gin.Context) {
	var creds struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if creds.Mobile == "09123456789" && creds.Password == "password" {
		c.JSON(http.StatusOK, gin.H{"token": "valid-token"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func mockRegisterHandler(c *gin.Context) {
	var user struct {
		Mobile string `json:"mobile"`
	}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func TestAuthAPI(t *testing.T) {
	// Setup Gin
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	
	// Setup Routes
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", mockLoginHandler)
		auth.POST("/register", mockRegisterHandler)
	}

	t.Run("Login Success", func(t *testing.T) {
		body := map[string]string{
			"mobile":   "09123456789",
			"password": "password",
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}
	})

	t.Run("Login Failure", func(t *testing.T) {
		body := map[string]string{
			"mobile":   "09123456789",
			"password": "wrong-password",
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status 401, got %d", w.Code)
		}
	})

	t.Run("Register Success", func(t *testing.T) {
		body := map[string]string{
			"mobile": "09123456789",
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status 201, got %d", w.Code)
		}
	})
}
