package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// ------------------------------------------------------------
//  RATE LIMIT MIDDLEWARE (IP-Based)
//  --------------------------------
//  • Mitigates brute-force login & DoS on sensitive endpoints.
//  • Simple token-bucket per IP kept in memory (goroutine-safe).
//  • Default limit: 60 requests / minute / IP.
//  • Returns HTTP 429 on exceeding the quota.
// ------------------------------------------------------------

const defaultLimit = 60
const refillInterval = time.Minute

// bucket represents a token bucket for one IP
type bucket struct {
	tokens int
	last   time.Time
}

var (
	mu       sync.Mutex
	visitors = make(map[string]*bucket)
)

// RateLimit returns a gin.HandlerFunc enforcing the default limit.
func RateLimit() gin.HandlerFunc { return RateLimitWith(defaultLimit, refillInterval, "default") }

// RateLimitWith اجازه تعریف محدوده نرخ سفارشی و scope مجزا را می‌دهد
func RateLimitWith(limit int, window time.Duration, scope string) gin.HandlerFunc {
	if limit <= 0 {
		limit = defaultLimit
	}
	if window <= 0 {
		window = refillInterval
	}
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := scope + ":" + ip

		mu.Lock()
		b, ok := visitors[key]
		if !ok {
			b = &bucket{tokens: limit, last: time.Now()}
			visitors[key] = b
		}
		// refill tokens
		elapsed := time.Since(b.last)
		if elapsed >= window {
			b.tokens = limit
			b.last = time.Now()
		}
		if b.tokens <= 0 {
			mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		}
		b.tokens--
		mu.Unlock()

		c.Next()
	}
}
