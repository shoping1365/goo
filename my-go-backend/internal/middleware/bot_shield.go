package middleware

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/services"
)

// BotShieldConfig تنظیمات کش و کلیدها
type BotShieldConfig struct {
	// کلیدهای تنظیمات برای خواندن از SettingService
	EnabledKey             string
	GlobalModeKey          string
	AllowlistUserAgentsKey string
	BlocklistUserAgentsKey string
	AllowlistIPsKey        string
	BlocklistIPsKey        string
	MaliciousUserAgentsKey string
	// کلیدهای کنترلی برای کاهش سربار و جداسازی ترافیک داخلی
	ExcludePathsKey   string // CSV از prefix ها برای عبور (بدون مانیتور/مسدود)
	SkipAuthorizedKey string // اگر true و هدر Authorization وجود داشت → عبور
	SkipLocalIPsKey   string // اگر true و IP محلی بود → عبور
	SkipMethodsKey    string // CSV متدهایی که باید عبور کنند (مثلاً OPTIONS,HEAD)
	MonitorOnlyKey    string // اگر true فقط مانیتور و هیچ مسدودی اعمال نشود
}

// defaultConfig تنظیمات پیش‌فرض کلیدها (کامنت شده)
/*
var defaultConfig = BotShieldConfig{
    EnabledKey:             "bots.enabled",
    GlobalModeKey:          "bots.global_mode",           // allow_all | block_all | restricted
    AllowlistUserAgentsKey: "bots.allowlist.user_agents", // CSV
    BlocklistUserAgentsKey: "bots.blocklist.user_agents", // CSV
    AllowlistIPsKey:        "bots.allowlist.ips",         // CSV
    BlocklistIPsKey:        "bots.blocklist.ips",         // CSV
    MaliciousUserAgentsKey: "bots.malicious.user_agents", // CSV یا الگوها
    ExcludePathsKey:        "bots.exclude_paths",
    SkipAuthorizedKey:      "bots.skip_authorized",
    SkipLocalIPsKey:        "bots.skip_local_ips",
    SkipMethodsKey:         "bots.skip_methods",
    MonitorOnlyKey:         "bots.monitor_only",
}
*/

// BotShield میدلور محافظت از ربات‌ها بر اساس تنظیمات سیستم
// این میدلور در تمام درخواست‌ها اجرا می‌شود و با توجه به حالت کلی و لیست‌های مجاز/مسدود عمل می‌کند.
// قوانین:
// - Googlebot همیشه اجازه دسترسی دارد (allowlist ذاتی + بررسی UA)
// - در حالت block_all همه به‌جز allowlist مسدود می‌شوند
// - در حالت restricted فقط UAهای مسدود/بدخواه بلاک می‌شوند
// - در حالت allow_all فقط IPهای صراحتاً مسدود بلاک می‌شوند
func BotShield(settingService *services.SettingService, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			        // 1) شَرط‌های عبور سریع برای حذف اثر روی سیستم داخلی و درخواست‌های کم‌اهمیت
			        path := c.Request.URL.Path
			        method := c.Request.Method

					// عبور فوق‌سریع برای تمام مسیرهای Admin و مسیرهای مدیریت موجودی محصول/انبار
					if strings.HasPrefix(path, "/admin") || strings.HasPrefix(path, "/api/admin") || strings.HasPrefix(path, "/product-warehouse-stocks") || strings.HasPrefix(path, "/api/product-warehouse-stocks") {
						c.Next()
						return
					}

					// متدهای عبوری (پیش‌فرض: OPTIONS, HEAD)
					skipMethodsCSV := readSettingValue(c, settingService, defaultConfig.SkipMethodsKey, "OPTIONS,HEAD")
					if methodMatchesAny(method, splitCSV(skipMethodsCSV)) {
						c.Next()
						return
					}

					// عبور درخواست‌های دارای Authorization (عملیات داخلی/ادمین)
					if strings.EqualFold(strings.TrimSpace(readSettingValue(c, settingService, defaultConfig.SkipAuthorizedKey, "true")), "true") {
						if auth := c.Request.Header.Get("Authorization"); strings.TrimSpace(auth) != "" {
							c.Next()
							return
						}
					}

					// عبور IP های محلی (لوکال هاست، شبکه داخلی)
					if strings.EqualFold(strings.TrimSpace(readSettingValue(c, settingService, defaultConfig.SkipLocalIPsKey, "true")), "true") {
						ip := clientIP(c.Request)
						if isLocalIP(ip) {
							c.Next()
							return
						}
					}

					// عبور مسیرهای استاتیک/ادمین/سیستمی (پیش‌فرض امن)
					// افزودن استثناهای اضافی طبق مستند: heartbeat و مسیرهای warehouses ادمین
					// نکته: /admin و /api/admin به‌صورت کلی عبور می‌دهند، اما برای شفافیت و لاگ بهتر، مسیرهای زیر هم اضافه شده‌اند
					defaultExcludes := "/_nuxt,/.nuxt,/assets,/public,/uploads,/static,/favicon.ico,/robots.txt,/health,/admin,/api/admin,/users/heartbeat,/api/users/heartbeat,/admin/warehouses,/api/admin/warehouses"
					excludesCSV := readSettingValue(c, settingService, defaultConfig.ExcludePathsKey, defaultExcludes)
					if pathHasAnyPrefix(path, splitCSV(excludesCSV)) {
						c.Next()
						return
					}

					// 2) اگر BotShield خاموش است → عبور
					enabled := strings.ToLower(strings.TrimSpace(readSettingValue(c, settingService, defaultConfig.EnabledKey, "true")))
					if enabled == "false" || enabled == "0" || enabled == "no" {
						c.Next()
						return
					}

					ua := c.Request.Header.Get("User-Agent")
					ip := clientIP(c.Request)

					// Googlebot همیشه مجاز
			        if isGoogleBot(ua) {
			            c.Next()
			            return
			        }

					// حالت کلی سیاست
					globalMode := readSettingValue(c, settingService, defaultConfig.GlobalModeKey, "allow_all")

					// خواندن لیست‌ها
					allowUA := splitCSV(readSettingValue(c, settingService, defaultConfig.AllowlistUserAgentsKey, ""))
					blockUA := splitCSV(readSettingValue(c, settingService, defaultConfig.BlocklistUserAgentsKey, ""))
					allowIPs := splitCSV(readSettingValue(c, settingService, defaultConfig.AllowlistIPsKey, ""))
					blockIPs := splitCSV(readSettingValue(c, settingService, defaultConfig.BlocklistIPsKey, ""))
					maliciousUA := splitCSV(readSettingValue(c, settingService, defaultConfig.MaliciousUserAgentsKey, ""))

					// اجازه IP های allowlist
					if ipInList(ip, allowIPs) {
						c.Next()
						return
					}
					// بلاک IP های blocklist
					if ipInList(ip, blockIPs) {
						if isMonitorOnly(c, settingService) {
							c.Next()
							return
						}
						c.AbortWithStatus(http.StatusNotFound)
						return
					}

					// 3) اعمال سیاست‌ها
			        switch strings.ToLower(globalMode) {
					case "block_all":
						if uaMatchesAny(ua, allowUA) {
							c.Next()
							return
						}
						if isMonitorOnly(c, settingService) {
							c.Next()
							return
						}
						c.AbortWithStatus(http.StatusNotFound)
						return
					case "restricted":
						if uaMatchesAny(ua, blockUA) || uaMatchesAny(ua, maliciousUA) {
							if isMonitorOnly(c, settingService) {
								c.Next()
								return
							}
							c.AbortWithStatus(http.StatusNotFound)
							return
						}
					default:
						// allow_all: فقط موارد بلاک‌شدهٔ IP در بالا رسیدگی شدند
			        }
		*/

		// pass-through
		c.Next()
	}
}

// readSettingValue خواندن مقدار تنظیم با مقدار پیش‌فرض
/*
func readSettingValue(c *gin.Context, svc *services.SettingService, key, def string) string {
    st, err := svc.GetSetting(c, key)
    if err != nil || st == nil || strings.TrimSpace(st.Value) == "" {
        return def
    }
    return st.Value
}
*/

// isGoogleBot تشخیص ساده Googlebot بر اساس UA
func isGoogleBot(ua string) bool {
	ua = strings.ToLower(ua)
	return strings.Contains(ua, "googlebot") || strings.Contains(ua, "apis-google")
}

// uaMatchesAny بررسی تطابق UA با هر یک از الگوهای ساده (contains)
func uaMatchesAny(ua string, patterns []string) bool {
	if ua == "" || len(patterns) == 0 {
		return false
	}
	low := strings.ToLower(ua)
	for _, p := range patterns {
		p = strings.TrimSpace(strings.ToLower(p))
		if p == "" {
			continue
		}
		if strings.Contains(low, p) {
			return true
		}
	}
	return false
}

// ipInList بررسی حضور IP در لیست
func ipInList(ip string, list []string) bool {
	if ip == "" || len(list) == 0 {
		return false
	}
	for _, v := range list {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		if ip == v {
			return true
		}
	}
	return false
}

// splitCSV جداکردن CSV به آرایه
func splitCSV(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// pathHasAnyPrefix بررسی می‌کند مسیر با هر prefix داده‌شده شروع می‌شود یا خیر
func pathHasAnyPrefix(path string, prefixes []string) bool {
	if path == "" || len(prefixes) == 0 {
		return false
	}
	for _, p := range prefixes {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}

// methodMatchesAny بررسی عبور متدها
func methodMatchesAny(m string, methods []string) bool {
	if m == "" || len(methods) == 0 {
		return false
	}
	for _, x := range methods {
		if strings.EqualFold(strings.TrimSpace(x), m) {
			return true
		}
	}
	return false
}

// isLocalIP تشخیص IP های لوکال/داخلی
func isLocalIP(ip string) bool {
	if ip == "" {
		return false
	}
	if ip == "127.0.0.1" || ip == "::1" || strings.HasPrefix(ip, "192.168.") || strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "172.16.") {
		return true
	}
	return false
}

// isMonitorOnly خواندن حالت مانیتور-صرف
func isMonitorOnly(c *gin.Context, svc *services.SettingService) bool {
	// وقتی BotShield غیر فعال است، همیشه حالت مانیتور را true در نظر می‌گیریم
	v := "true"
	return v == "true" || v == "1" || v == "yes"
}

// clientIP استخراج IP کلاینت با درنظرگرفتن X-Forwarded-For
func clientIP(r *http.Request) string {
	// ابتدا X-Forwarded-For
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// ممکن است چند IP باشد
		segs := strings.Split(xff, ",")
		if len(segs) > 0 {
			return strings.TrimSpace(segs[0])
		}
	}
	// سپس RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && host != "" {
		return host
	}
	return r.RemoteAddr
}
