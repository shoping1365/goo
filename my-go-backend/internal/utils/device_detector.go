package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// DeviceInfo اطلاعات دستگاه از User-Agent
type DeviceInfo struct {
	DeviceName     string `json:"device_name"`
	Platform       string `json:"platform"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	OSName         string `json:"os_name"`
	OSVersion      string `json:"os_version"`
	IsMobile       bool   `json:"is_mobile"`
	IsTablet       bool   `json:"is_tablet"`
	IsDesktop      bool   `json:"is_desktop"`
}

// DetectDevice تشخیص دستگاه از User-Agent
func DetectDevice(c *gin.Context) *DeviceInfo {
	userAgent := c.GetHeader("User-Agent")
	return ParseUserAgent(userAgent)
}

// ParseUserAgent تجزیه User-Agent string
func ParseUserAgent(userAgent string) *DeviceInfo {
	if userAgent == "" {
		return &DeviceInfo{
			DeviceName: "Unknown Device",
			Platform:   "unknown",
			Browser:    "unknown",
			OSName:     "unknown",
		}
	}

	ua := strings.ToLower(userAgent)
	device := &DeviceInfo{}

	// تشخیص Platform و OS
	device.Platform, device.OSName, device.OSVersion = detectOS(ua, userAgent)

	// تشخیص Browser
	device.Browser, device.BrowserVersion = detectBrowser(ua, userAgent)

	// تشخیص نوع دستگاه
	device.IsMobile = detectMobile(ua)
	device.IsTablet = detectTablet(ua)
	device.IsDesktop = !device.IsMobile && !device.IsTablet

	// تولید نام دستگاه
	device.DeviceName = generateDeviceName(device)

	return device
}

// detectOS تشخیص سیستم عامل
func detectOS(ua, originalUA string) (platform, osName, osVersion string) {
	// iOS
	if strings.Contains(ua, "iphone") {
		platform = "ios"
		osName = "iOS"
		if matches := regexp.MustCompile(`os (\d+)[_.](\d+)`).FindStringSubmatch(ua); len(matches) >= 3 {
			osVersion = matches[1] + "." + matches[2]
		}
		return
	}

	if strings.Contains(ua, "ipad") {
		platform = "ios"
		osName = "iPadOS"
		if matches := regexp.MustCompile(`os (\d+)[_.](\d+)`).FindStringSubmatch(ua); len(matches) >= 3 {
			osVersion = matches[1] + "." + matches[2]
		}
		return
	}

	// Android
	if strings.Contains(ua, "android") {
		platform = "android"
		osName = "Android"
		if matches := regexp.MustCompile(`android (\d+\.?\d*)`).FindStringSubmatch(ua); len(matches) >= 2 {
			osVersion = matches[1]
		}
		return
	}

	// Windows
	if strings.Contains(ua, "windows nt") {
		platform = "windows"
		osName = "Windows"
		if strings.Contains(ua, "windows nt 10.0") {
			osVersion = "10/11"
		} else if strings.Contains(ua, "windows nt 6.3") {
			osVersion = "8.1"
		} else if strings.Contains(ua, "windows nt 6.1") {
			osVersion = "7"
		}
		return
	}

	// macOS
	if strings.Contains(ua, "mac os x") || strings.Contains(ua, "macos") {
		platform = "macos"
		osName = "macOS"
		if matches := regexp.MustCompile(`mac os x (\d+)[_.](\d+)`).FindStringSubmatch(ua); len(matches) >= 3 {
			osVersion = matches[1] + "." + matches[2]
		}
		return
	}

	// Linux
	if strings.Contains(ua, "linux") {
		platform = "linux"
		osName = "Linux"
		if strings.Contains(ua, "ubuntu") {
			osName = "Ubuntu"
		} else if strings.Contains(ua, "fedora") {
			osName = "Fedora"
		}
		return
	}

	// Default
	return "unknown", "Unknown", ""
}

// detectBrowser تشخیص مرورگر
func detectBrowser(ua, originalUA string) (browser, version string) {
	// Chrome (باید قبل از Safari چک شود)
	if strings.Contains(ua, "chrome") && !strings.Contains(ua, "edg") {
		browser = "Chrome"
		if matches := regexp.MustCompile(`chrome\/(\d+\.?\d*)`).FindStringSubmatch(ua); len(matches) >= 2 {
			version = matches[1]
		}
		return
	}

	// Edge
	if strings.Contains(ua, "edg") {
		browser = "Edge"
		if matches := regexp.MustCompile(`edg\/(\d+\.?\d*)`).FindStringSubmatch(ua); len(matches) >= 2 {
			version = matches[1]
		}
		return
	}

	// Firefox
	if strings.Contains(ua, "firefox") {
		browser = "Firefox"
		if matches := regexp.MustCompile(`firefox\/(\d+\.?\d*)`).FindStringSubmatch(ua); len(matches) >= 2 {
			version = matches[1]
		}
		return
	}

	// Safari (باید آخر چک شود)
	if strings.Contains(ua, "safari") && !strings.Contains(ua, "chrome") {
		browser = "Safari"
		if matches := regexp.MustCompile(`version\/(\d+\.?\d*)`).FindStringSubmatch(ua); len(matches) >= 2 {
			version = matches[1]
		}
		return
	}

	// Opera
	if strings.Contains(ua, "opera") || strings.Contains(ua, "opr") {
		browser = "Opera"
		if matches := regexp.MustCompile(`(?:opera|opr)\/(\d+\.?\d*)`).FindStringSubmatch(ua); len(matches) >= 2 {
			version = matches[1]
		}
		return
	}

	return "Unknown", ""
}

// detectMobile تشخیص موبایل
func detectMobile(ua string) bool {
	mobileKeywords := []string{
		"mobile", "android", "iphone", "ipod", "blackberry",
		"windows phone", "palm", "smartphone", "iemobile",
	}

	for _, keyword := range mobileKeywords {
		if strings.Contains(ua, keyword) {
			return true
		}
	}
	return false
}

// detectTablet تشخیص تبلت
func detectTablet(ua string) bool {
	tabletKeywords := []string{"ipad", "tablet", "playbook", "kindle"}

	for _, keyword := range tabletKeywords {
		if strings.Contains(ua, keyword) {
			return true
		}
	}

	// Android tablets (معمولاً mobile نیستند ولی Android هستند)
	return strings.Contains(ua, "android") && !strings.Contains(ua, "mobile")
}

// generateDeviceName تولید نام دستگاه
func generateDeviceName(device *DeviceInfo) string {
	var deviceType string
	if device.IsMobile {
		deviceType = "Mobile"
	} else if device.IsTablet {
		deviceType = "Tablet"
	} else {
		deviceType = "Desktop"
	}

	if device.OSName != "Unknown" && device.Browser != "Unknown" {
		if device.OSVersion != "" {
			return device.OSName + " " + device.OSVersion + " - " + device.Browser + " (" + deviceType + ")"
		}
		return device.OSName + " - " + device.Browser + " (" + deviceType + ")"
	}

	if device.OSName != "Unknown" {
		return device.OSName + " " + deviceType
	}

	if device.Browser != "Unknown" {
		return device.Browser + " " + deviceType
	}

	return "Unknown " + deviceType
}

// GetClientFingerprint تولید fingerprint دستگاه
func GetClientFingerprint(c *gin.Context) string {
	userAgent := c.GetHeader("User-Agent")
	acceptLanguage := c.GetHeader("Accept-Language")
	acceptEncoding := c.GetHeader("Accept-Encoding")

	// ترکیب اطلاعات برای fingerprint
	fingerprint := userAgent + "|" + acceptLanguage + "|" + acceptEncoding

	// Hash کردن برای کوتاه‌تر شدن
	return hashString(fingerprint)[:16] // فقط 16 کاراکتر اول
}

// hashString تولید SHA256 hash
func hashString(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
