package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadVideoHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "فایل ویدیو یافت نشد",
		})
		return
	}
	defer file.Close()

	// بررسی نوع فایل
	contentType := header.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "video/") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "فایل باید ویدیو باشد",
		})
		return
	}

	// بررسی اندازه فایل (50MB)
	if header.Size > 50*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "حجم فایل بیش از حد مجاز است (حداکثر 50MB)",
		})
		return
	}

	// ایجاد نام یکتا برای فایل
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("video_%d%s", time.Now().Unix(), ext)

	// دریافت category از فرم
	category := c.DefaultPostForm("category", "library")
	category = strings.ToLower(category)
	// جلوگیری از مقادیر غیرمجاز
	allowedCategories := map[string]bool{"products": true, "banners": true, "customer": true, "library": true}
	if !allowedCategories[category] {
		category = "library"
	}

	// ایجاد مسیر ذخیره بر اساس category
	uploadPath := filepath.Join("public", "uploads", "media", category, "videos")
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "خطا در ایجاد پوشه",
		})
		return
	}

	// ذخیره فایل
	filePath := filepath.Join(uploadPath, filename)
	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "خطا در ذخیره فایل",
		})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "خطا در کپی فایل",
		})
		return
	}

	// برگرداندن آدرس ویدیو
	videoURL := fmt.Sprintf("/uploads/media/%s/videos/%s", category, filename)
	c.JSON(http.StatusOK, gin.H{
		"video_url": videoURL,
		"message":   "ویدیو با موفقیت آپلود شد",
	})
}
