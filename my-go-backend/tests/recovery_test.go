package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPanicRecovery(t *testing.T) {
	// Setup Gin with Recovery middleware
	// gin.Default() includes Logger and Recovery middleware automatically
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// Define a route that intentionally panics
	r.GET("/trigger-panic", func(c *gin.Context) {
		// Simulate a critical error like index out of range or nil pointer dereference
		panic("Critical system failure!")
	})

	// Perform request
	req, _ := http.NewRequest("GET", "/trigger-panic", nil)
	w := httptest.NewRecorder()
	
	// Because gin.Recovery() prints to stderr, we expect this to log the panic but NOT crash the test
	// The recovery middleware should catch the panic and return 500
	r.ServeHTTP(w, req)

	// Assert that server handled it gracefully
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected 500 Internal Server Error (Panic Recovered), got %d", w.Code)
	}
}
