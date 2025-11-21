package tests

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

// TestAPIStress simulates concurrent load on the API
// Run with: go test -v -run TestAPIStress
func TestAPIStress(t *testing.T) {
	// Setup a simple handler
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/heavy-process", func(c *gin.Context) {
		// Simulate some processing time (e.g., DB query)
		time.Sleep(5 * time.Millisecond)
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Configuration for stress test
	concurrentUsers := 50
	requestsPerUser := 20
	totalRequests := concurrentUsers * requestsPerUser

	var wg sync.WaitGroup
	var failedRequests int
	var mu sync.Mutex

	start := time.Now()

	// Launch concurrent users
	for i := 0; i < concurrentUsers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerUser; j++ {
				req, _ := http.NewRequest("GET", "/heavy-process", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)

				if w.Code != 200 {
					mu.Lock()
					failedRequests++
					mu.Unlock()
				}
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	rps := float64(totalRequests) / duration.Seconds()

	t.Logf("Stress Test Results:")
	t.Logf("Total Requests: %d", totalRequests)
	t.Logf("Concurrent Users: %d", concurrentUsers)
	t.Logf("Total Time: %v", duration)
	t.Logf("Requests Per Second (RPS): %.2f", rps)
	t.Logf("Failed Requests: %d", failedRequests)

	if failedRequests > 0 {
		t.Errorf("Stress test failed with %d errors", failedRequests)
	}
}
