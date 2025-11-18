package handlers

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// generateUniqueName تولید نام یکتا برای فایل‌ها
func generateUniqueName(ext string) string {
	ts := time.Now().UnixNano() / 1e6
	randPart := rand.Intn(0xffff)
	return fmt.Sprintf("%s-%04x%s", strconv.FormatInt(ts, 36), randPart, ext)
}

// validateAndSanitizePath اعتبارسنجی و پاکسازی مسیر برای جلوگیری از path traversal
// این تابع بررسی می‌کند که مسیر در محدوده baseDir باشد و از path traversal جلوگیری می‌کند
func validateAndSanitizePath(baseDir, userPath string) (string, error) {
	if baseDir == "" {
		return "", fmt.Errorf("baseDir نمی‌تواند خالی باشد")
	}

	// پاک کردن / از ابتدای userPath
	userPath = strings.TrimPrefix(userPath, "/")
	if userPath == "" {
		return "", fmt.Errorf("مسیر نمی‌تواند خالی باشد")
	}

	// ساخت مسیر کامل
	fullPath := filepath.Join(baseDir, userPath)

	// Clean کردن مسیر برای حذف .. و .
	cleanPath := filepath.Clean(fullPath)

	// بررسی اینکه مسیر clean شده هنوز در baseDir باشد (جلوگیری از path traversal)
	baseDirAbs, err := filepath.Abs(baseDir)
	if err != nil {
		return "", fmt.Errorf("خطا در محاسبه مسیر مطلق baseDir: %v", err)
	}

	cleanPathAbs, err := filepath.Abs(cleanPath)
	if err != nil {
		return "", fmt.Errorf("خطا در محاسبه مسیر مطلق: %v", err)
	}

	// بررسی اینکه cleanPathAbs با baseDirAbs شروع می‌شود
	if !strings.HasPrefix(cleanPathAbs, baseDirAbs) {
		return "", fmt.Errorf("مسیر خارج از محدوده مجاز است")
	}

	return cleanPath, nil
}

// buildSafeMediaPath ساخت ایمن مسیر برای فایل‌های رسانه
// این تابع category و subfolder را validate می‌کند و مسیر ایمن می‌سازد
func buildSafeMediaPath(baseDir, category, subfolder, fileName string) (string, error) {
	// Validate category
	allowedCategories := map[string]bool{
		"products": true, "library": true, "customer": true,
		"product-categories": true, "brands": true, "banners": true,
	}
	if !allowedCategories[category] {
		return "", fmt.Errorf("دسته‌بندی نامعتبر: %s", category)
	}

	// Validate subfolder
	allowedSubfolders := map[string]bool{
		"images": true, "videos": true, "audio": true,
		"documents": true, "others": true,
	}
	if !allowedSubfolders[subfolder] {
		return "", fmt.Errorf("زیرپوشه نامعتبر: %s", subfolder)
	}

	// Validate fileName - باید فقط شامل کاراکترهای مجاز باشد
	if fileName == "" {
		return "", fmt.Errorf("نام فایل نمی‌تواند خالی باشد")
	}
	if strings.Contains(fileName, "..") || strings.Contains(fileName, "/") || strings.Contains(fileName, "\\") {
		return "", fmt.Errorf("نام فایل نامعتبر: %s", fileName)
	}

	// ساخت مسیر
	path := filepath.Join(baseDir, "uploads", "media", category, subfolder, fileName)
	return filepath.Clean(path), nil
}
