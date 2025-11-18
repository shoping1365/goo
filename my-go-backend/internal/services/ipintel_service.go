package services

import (
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

// IPIntelResult نتیجه بررسی IP
// IsDatacenter: آیا IP متعلق به دیتاسنتر/پروکسی شناخته‌شده است
// City: شهر تخمینی بر اساس سرویس GeoIP (در این نسخه می‌تواند خالی باشد)
// Country: کشور تخمینی
type IPIntelResult struct {
	IsDatacenter bool
	City         string
	Country      string
}

// IPIntelService اینترفیس سرویس تشخیص اطلاعات IP
type IPIntelService interface {
	Lookup(ip string) (*IPIntelResult, error)
}

// ipIntelDefaultService پیاده‌سازی پیش‌فرض (سبک) برای IPIntelService
// در صورت تنظیم ENV برای ارائه‌دهنده خارجی (مثلاً URL یا TOKEN)، می‌توان بعداً توسعه داد.
type ipIntelDefaultService struct {
	providerURL   string
	providerToken string
	cache         *ipintelCache
}

// NewIPIntelServiceFromEnv ایجاد سرویس از متغیرهای محیطی
// ENV ها (اختیاری): IPINTEL_URL, IPINTEL_TOKEN
func NewIPIntelServiceFromEnv() IPIntelService {
	return &ipIntelDefaultService{
		providerURL:   os.Getenv("IPINTEL_URL"),
		providerToken: os.Getenv("IPINTEL_TOKEN"),
		cache:         newIPIntelCache(10 * time.Minute),
	}
}

// Lookup
// در این نسخه: تشخیص private/reserved و برخی الگوهای دیتاسنتر ساده.
// اگر provider تنظیم شده باشد، می‌توان اینجا call بیرونی را اضافه کرد (فعلاً بدون call خارجی).
func (s *ipIntelDefaultService) Lookup(ip string) (*IPIntelResult, error) {
	if s.cache != nil {
		if v, ok := s.cache.get(ip); ok {
			return v, nil
		}
	}
	res := &IPIntelResult{IsDatacenter: false, City: "", Country: ""}
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return res, nil
	}
	// Private/Reserved: نه دیتاسنتر
	if isPrivateIP(parsed) {
		return res, nil
	}
	// Heuristic ساده: اگر IP شامل بلوک‌های شناخته‌شده برخی دیتاسنترها باشد (نمونه اولیه)
	// توجه: برای دقت بالا باید از دیتابیس ASN یا سرویس معتبر استفاده شود.
	if strings.HasPrefix(ip, "172.") || strings.HasPrefix(ip, "13.") || strings.HasPrefix(ip, "52.") || strings.HasPrefix(ip, "54.") {
		res.IsDatacenter = true
	}
	if s.cache != nil {
		s.cache.set(ip, res)
	}
	return res, nil
}

func isPrivateIP(ip net.IP) bool {
	privateBlocks := []string{
		"10.", "192.168.", "172.16.", "172.17.", "172.18.", "172.19.", "172.2",
	}
	s := ip.String()
	for _, p := range privateBlocks {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}

// ساده: کش درون‌حافظه‌ای با TTL برای نتایج IPIntel
type ipintelCache struct {
	mu    sync.RWMutex
	ttl   time.Duration
	store map[string]ipintelCacheEntry
}

type ipintelCacheEntry struct {
	val   *IPIntelResult
	expAt time.Time
}

func newIPIntelCache(ttl time.Duration) *ipintelCache {
	if ttl <= 0 {
		ttl = 10 * time.Minute
	}
	return &ipintelCache{ttl: ttl, store: make(map[string]ipintelCacheEntry)}
}

func (c *ipintelCache) get(key string) (*IPIntelResult, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.store[key]
	if !ok {
		return nil, false
	}
	if time.Now().After(e.expAt) {
		return nil, false
	}
	return e.val, true
}

func (c *ipintelCache) set(key string, val *IPIntelResult) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = ipintelCacheEntry{val: val, expAt: time.Now().Add(c.ttl)}
}
