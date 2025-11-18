package handlers

import (
	"fmt"
	"image"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
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

// safeOpenImage باز کردن ایمن تصویر با panic recovery برای جلوگیری از CVE-2023-36308
// این تابع imaging.Open را wrap می‌کند و در صورت panic، خطا را به صورت کنترل شده برمی‌گرداند
func safeOpenImage(filename string) (image.Image, error) {
	var img image.Image
	var err error
	var panicErr interface{}

	// استفاده از defer/recover برای catch کردن panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicErr = r
			}
		}()
		img, err = imaging.Open(filename)
	}()

	// اگر panic رخ داده باشد، خطا برگردان
	if panicErr != nil {
		return nil, fmt.Errorf("panic در پردازش تصویر (احتمالاً فایل TIFF آسیب‌پذیر): %v", panicErr)
	}

	// اگر خطای عادی رخ داده باشد، همان را برگردان
	if err != nil {
		return nil, err
	}

	return img, nil
}
