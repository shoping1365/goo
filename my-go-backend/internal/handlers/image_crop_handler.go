package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

// CropImageHandler handles POST /api/media/image-crop
// برش تصویر بر اساس ابعاد مشخص شده برای موبایل
func (h *MediaHandler) CropImageHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Src        string `json:"src"`
		CropWidth  int    `json:"crop_width"`
		CropHeight int    `json:"crop_height"`
		Mode       string `json:"mode"`    // cover, contain, center, top, bottom
		Device     string `json:"device"`  // mobile, desktop
		Quality    int    `json:"quality"` // 1-100
	}

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	// تنظیمات پیش‌فرض
	if req.Mode == "" {
		req.Mode = "cover"
	}
	if req.Quality == 0 {
		req.Quality = 85
	}
	if req.Device == "" {
		req.Device = "mobile"
	}

	// بررسی وجود فایل منبع
	projectRoot := ""
	if wd, err := os.Getwd(); err == nil {
		projectRoot = filepath.Dir(wd)
	}
	publicDir := filepath.Join(projectRoot, "public")

	// اعتبارسنجی و پاکسازی مسیر برای جلوگیری از path traversal
	absSrcPath, err := validateAndSanitizePath(publicDir, req.Src)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "مسیر نامعتبر", err.Error()))
		return
	}

	if _, err := os.Stat(absSrcPath); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("FILE_NOT_FOUND", "فایل منبع یافت نشد", req.Src))
		return
	}

	// Security: Validate TIFF files to prevent CVE-2023-36808 (crash on crafted TIFF)
	extLower := strings.ToLower(filepath.Ext(absSrcPath))
	if extLower == ".tiff" || extLower == ".tif" {
		if fileInfo, err := os.Stat(absSrcPath); err == nil {
			if fileInfo.Size() > 10*1024*1024 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(utils.New("FILE_TOO_LARGE", "فایل TIFF بیش از حد بزرگ است (حداکثر 10MB)", nil))
				return
			}
		}
	}

	// خواندن تصویر اصلی با استفاده از safeOpenImage برای جلوگیری از panic (CVE-2023-36308)
	img, err := safeOpenImage(absSrcPath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("IMAGE_READ_ERROR", "خطا در خواندن تصویر", err.Error()))
		return
	}

	// برش تصویر بر اساس mode
	// همه حالت‌ها از imaging.Fit استفاده می‌کنند
	croppedImg := imaging.Fit(img, req.CropWidth, req.CropHeight, imaging.Lanczos)

	// تولید نام فایل جدید
	ext := filepath.Ext(absSrcPath)
	baseName := strings.TrimSuffix(filepath.Base(absSrcPath), ext)
	timestamp := time.Now().Unix()
	croppedFileName := fmt.Sprintf("%s_cropped_%dx%d_%d%s", baseName, req.CropWidth, req.CropHeight, timestamp, ext)

	// مسیر ذخیره فایل برش شده - ذخیره در پوشه مشخص شده
	croppedDir := filepath.Join(publicDir, "uploads", "media", "banners", "images", "cropped")
	if err := os.MkdirAll(croppedDir, 0755); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DIRECTORY_ERROR", "خطا در ایجاد پوشه", err.Error()))
		return
	}

	croppedPath := filepath.Join(croppedDir, croppedFileName)
	croppedRelPath := strings.TrimPrefix(croppedPath, publicDir)
	croppedRelPath = strings.ReplaceAll(croppedRelPath, "\\", "/")

	// ذخیره تصویر برش شده
	var buf bytes.Buffer
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		if err := jpeg.Encode(&buf, croppedImg, &jpeg.Options{Quality: req.Quality}); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.New("ENCODE_ERROR", "خطا در کدگذاری JPEG", err.Error()))
			return
		}
	case ".png":
		if err := png.Encode(&buf, croppedImg); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.New("ENCODE_ERROR", "خطا در کدگذاری PNG", err.Error()))
			return
		}
	case ".webp":
		if err := webp.Encode(&buf, croppedImg, &webp.Options{Quality: float32(req.Quality)}); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.New("ENCODE_ERROR", "خطا در کدگذاری WebP", err.Error()))
			return
		}
	default:
		// پیش‌فرض JPEG
		if err := jpeg.Encode(&buf, croppedImg, &jpeg.Options{Quality: req.Quality}); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.New("ENCODE_ERROR", "خطا در کدگذاری تصویر", err.Error()))
			return
		}
	}

	// نوشتن فایل
	if err := os.WriteFile(croppedPath, buf.Bytes(), 0644); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("SAVE_ERROR", "خطا در ذخیره فایل", err.Error()))
		return
	}

	// ثبت در دیتابیس (اختیاری)
	media := models.MediaFile{
		FileName:   croppedFileName,
		Title:      fmt.Sprintf("Cropped %dx%d - %s", req.CropWidth, req.CropHeight, baseName),
		FileType:   "image/" + strings.TrimPrefix(ext, "."),
		MimeType:   "image/" + strings.TrimPrefix(ext, "."),
		Size:       uint(buf.Len()),
		URL:        croppedRelPath,
		Category:   "banners",
		UploadedBy: nil, // می‌توان از context کاربر گرفت
	}

	if err := h.DB.Create(&media).Error; err != nil {
		// اگر ثبت در دیتابیس ناموفق بود، فایل را حذف نکنیم
		fmt.Printf("Warning: Failed to save cropped image to database: %v\n", err)
	}

	// پاسخ موفقیت‌آمیز
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"original_url": req.Src,
			"cropped_url":  croppedRelPath,
			"width":        req.CropWidth,
			"height":       req.CropHeight,
			"mode":         req.Mode,
			"device":       req.Device,
			"quality":      req.Quality,
			"file_size":    buf.Len(),
			"media_id":     media.ID,
		},
	})
}

// GetCroppedImageHandler handles GET /api/media/image-crop
// دریافت تصویر برش شده بر اساس پارامترهای query
func (h *MediaHandler) GetCroppedImageHandler(w http.ResponseWriter, r *http.Request) {
	src := r.URL.Query().Get("src")
	crop := r.URL.Query().Get("crop") // format: "375x150"
	mode := r.URL.Query().Get("mode")
	device := r.URL.Query().Get("device")
	mobile := r.URL.Query().Get("mobile") == "true"

	if src == "" || crop == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("MISSING_PARAMS", "پارامترهای src و crop الزامی هستند", nil))
		return
	}

	// پارس کردن ابعاد برش
	cropParts := strings.Split(crop, "x")
	if len(cropParts) != 2 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("INVALID_CROP", "فرمت crop نامعتبر. باید به صورت widthxheight باشد", crop))
		return
	}

	cropWidth, err := strconv.Atoi(cropParts[0])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("INVALID_WIDTH", "عرض برش نامعتبر", cropParts[0]))
		return
	}

	cropHeight, err := strconv.Atoi(cropParts[1])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("INVALID_HEIGHT", "ارتفاع برش نامعتبر", cropParts[1]))
		return
	}

	// تنظیمات پیش‌فرض
	if mode == "" {
		mode = "cover"
	}
	if device == "" {
		if mobile {
			device = "mobile"
		} else {
			device = "desktop"
		}
	}

	// بررسی وجود فایل منبع
	projectRoot := ""
	if wd, err := os.Getwd(); err == nil {
		projectRoot = filepath.Dir(wd)
	}
	publicDir := filepath.Join(projectRoot, "public")

	// اعتبارسنجی و پاکسازی مسیر برای جلوگیری از path traversal
	absSrcPath, err := validateAndSanitizePath(publicDir, src)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "مسیر نامعتبر", err.Error()))
		return
	}

	if _, err := os.Stat(absSrcPath); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("FILE_NOT_FOUND", "فایل منبع یافت نشد", src))
		return
	}

	// Security: Validate TIFF files to prevent CVE-2023-36808 (crash on crafted TIFF)
	extLower := strings.ToLower(filepath.Ext(absSrcPath))
	if extLower == ".tiff" || extLower == ".tif" {
		if fileInfo, err := os.Stat(absSrcPath); err == nil {
			if fileInfo.Size() > 10*1024*1024 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(utils.New("FILE_TOO_LARGE", "فایل TIFF بیش از حد بزرگ است (حداکثر 10MB)", nil))
				return
			}
		}
	}

	// خواندن تصویر اصلی با استفاده از safeOpenImage برای جلوگیری از panic (CVE-2023-36308)
	img, err := safeOpenImage(absSrcPath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("IMAGE_READ_ERROR", "خطا در خواندن تصویر", err.Error()))
		return
	}

	// برش تصویر
	// همه حالت‌ها از imaging.Fit استفاده می‌کنند
	croppedImg := imaging.Fit(img, cropWidth, cropHeight, imaging.Lanczos)

	// تنظیم Content-Type مناسب
	ext := strings.ToLower(filepath.Ext(absSrcPath))
	var contentType string
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".webp":
		contentType = "image/webp"
	case ".gif":
		contentType = "image/gif"
	default:
		contentType = "image/jpeg"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "public, max-age=31536000") // کش برای 1 سال

	// ارسال تصویر برش شده
	switch contentType {
	case "image/jpeg":
		jpeg.Encode(w, croppedImg, &jpeg.Options{Quality: 85})
	case "image/png":
		png.Encode(w, croppedImg)
	case "image/webp":
		webp.Encode(w, croppedImg, &webp.Options{Quality: 85})
	case "image/gif":
		// برای GIF، تصویر اصلی را برگردان (برش پیچیده‌تر است)
		http.ServeFile(w, r, absSrcPath)
	default:
		jpeg.Encode(w, croppedImg, &jpeg.Options{Quality: 85})
	}
}
