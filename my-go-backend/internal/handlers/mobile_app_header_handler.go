package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
	"mime/multipart"
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
	"strconv"

	"gorm.io/gorm"
)

type MobileAppHeaderHandler struct {
	settingService *services.SettingService
}

// NewMobileAppHeaderHandler یک نمونه جدید از کنترلر تنظیمات هدر موبایل و اپلیکیشن ایجاد می‌کند
func NewMobileAppHeaderHandler(settingService *services.SettingService) *MobileAppHeaderHandler {
	return &MobileAppHeaderHandler{
		settingService: settingService,
	}
}

// GetMobileAppHeaderSettings دریافت تنظیمات هدر موبایل و اپلیکیشن برای نمایش عمومی
func (h *MobileAppHeaderHandler) GetMobileAppHeaderSettings(c *gin.Context) {
	fmt.Printf("🔍 GetMobileAppHeaderSettings: شروع دریافت تنظیمات هدر موبایل و اپلیکیشن\n")

	// دریافت هدرهای فعال از جدول mobile_app_headers
	var headers []models.MobileAppHeader
	if err := h.settingService.DB().Where("is_active = ?", true).Find(&headers).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت هدرهای موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات هدر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("📊 Number of active mobile and app headers found: %d\n", len(headers))

	// اگر هدری وجود ندارد، پیام مناسب برگردان
	if len(headers) == 0 {
		fmt.Printf("⚠️ No active mobile and app headers found\n")
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []interface{}{},
			"message": "No active mobile and app headers found",
		})
		return
	}

	// تبدیل به فرمت مناسب برای frontend
	var responseData []map[string]interface{}
	for _, header := range headers {
		headerData := map[string]interface{}{
			"id":                 header.ID,
			"name":               header.Name,
			"description":        header.Description,
			"platform":           header.Platform,
			"page_selection":     header.PageSelection,
			"specific_pages":     header.SpecificPages,
			"excluded_pages":     header.ExcludedPages,
			"header_type":        header.HeaderType,
			"logo_url":           header.LogoURL,
			"logo_alt":           header.LogoAlt,
			"show_search":        header.ShowSearch,
			"show_cart":          header.ShowCart,
			"show_user_menu":     header.ShowUserMenu,
			"show_notifications": header.ShowNotifications,
			"show_menu_button":   header.ShowMenuButton,
			"background_color":   header.BackgroundColor,
			"text_color":         header.TextColor,
			"top_image_url":      header.TopImageURL,
			"top_image_alt":      header.TopImageAlt,
			"bottom_image_url":   header.BottomImageURL,
			"bottom_image_alt":   header.BottomImageAlt,
			"is_active":          header.IsActive,
			"created_at":         header.CreatedAt,
			"updated_at":         header.UpdatedAt,
		}
		responseData = append(responseData, headerData)
	}

	fmt.Printf("✅ تنظیمات هدر موبایل و اپلیکیشن با موفقیت ارسال شد\n")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    responseData,
		"message": "تنظیمات هدر موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// ListMobileAppHeaders لیست کردن تمام هدرهای موبایل و اپلیکیشن (برای ادمین)
func (h *MobileAppHeaderHandler) ListMobileAppHeaders(c *gin.Context) {
	fmt.Printf("🔍 ListMobileAppHeaders: شروع لیست کردن هدرهای موبایل و اپلیکیشن\n")

	var headers []models.MobileAppHeader
	if err := h.settingService.DB().Find(&headers).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت لیست هدرهای موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست هدرهای موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("📊 تعداد کل هدرهای موبایل و اپلیکیشن: %d\n", len(headers))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    headers,
		"message": "لیست هدرهای موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// CreateMobileAppHeader ایجاد هدر جدید موبایل و اپلیکیشن
func (h *MobileAppHeaderHandler) CreateMobileAppHeader(c *gin.Context) {
	fmt.Printf("🔍 CreateMobileAppHeader: شروع ایجاد هدر جدید موبایل و اپلیکیشن\n")

	var header models.MobileAppHeader
	if err := c.ShouldBindJSON(&header); err != nil {
		fmt.Printf("❌ خطا در پارس کردن JSON: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	// بررسی وجود نام تکراری
	var existingHeader models.MobileAppHeader
	if err := h.settingService.DB().Where("name = ?", header.Name).First(&existingHeader).Error; err == nil {
		fmt.Printf("❌ نام هدر تکراری است: %s\n", header.Name)
		c.AbortWithStatusJSON(http.StatusConflict, utils.New("DUPLICATE_ERROR", "نام هدر موبایل و اپلیکیشن تکراری است", nil))
		return
	}

	// ایجاد هدر جدید
	if err := h.settingService.DB().Create(&header).Error; err != nil {
		fmt.Printf("❌ خطا در ایجاد هدر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد هدر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ هدر موبایل و اپلیکیشن با موفقیت ایجاد شد: %s (ID: %d)\n", header.Name, header.ID)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    header,
		"message": "هدر موبایل و اپلیکیشن با موفقیت ایجاد شد",
	})
}

// GetMobileAppHeaderByID دریافت هدر موبایل و اپلیکیشن بر اساس ID
func (h *MobileAppHeaderHandler) GetMobileAppHeaderByID(c *gin.Context) {
	fmt.Printf("🔍 GetMobileAppHeaderByID: شروع دریافت هدر موبایل و اپلیکیشن بر اساس ID\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var header models.MobileAppHeader
	if err := h.settingService.DB().First(&header, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ هدر موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "هدر موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت هدر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت هدر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ هدر موبایل و اپلیکیشن با موفقیت دریافت شد: %s (ID: %d)\n", header.Name, header.ID)

	// تبدیل به فرمت مناسب برای frontend
	headerData := map[string]interface{}{
		"id":                 header.ID,
		"name":               header.Name,
		"description":        header.Description,
		"platform":           header.Platform,
		"page_selection":     header.PageSelection,
		"specific_pages":     header.SpecificPages,
		"excluded_pages":     header.ExcludedPages,
		"header_type":        header.HeaderType,
		"logo_url":           header.LogoURL,
		"logo_alt":           header.LogoAlt,
		"show_search":        header.ShowSearch,
		"show_cart":          header.ShowCart,
		"show_user_menu":     header.ShowUserMenu,
		"show_notifications": header.ShowNotifications,
		"show_menu_button":   header.ShowMenuButton,
		"background_color":   header.BackgroundColor,
		"text_color":         header.TextColor,
		"top_image_url":      header.TopImageURL,
		"top_image_alt":      header.TopImageAlt,
		"bottom_image_url":   header.BottomImageURL,
		"bottom_image_alt":   header.BottomImageAlt,
		"is_active":          header.IsActive,
		"created_at":         header.CreatedAt,
		"updated_at":         header.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    headerData,
		"message": "هدر موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// UpdateMobileAppHeaderByID به‌روزرسانی هدر موبایل و اپلیکیشن بر اساس ID
func (h *MobileAppHeaderHandler) UpdateMobileAppHeaderByID(c *gin.Context) {
	fmt.Printf("🔍 UpdateMobileAppHeaderByID: شروع به‌روزرسانی هدر موبایل و اپلیکیشن\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var header models.MobileAppHeader
	if err := h.settingService.DB().First(&header, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ هدر موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "هدر موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت هدر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت هدر موبایل و اپلیکیشن", err.Error()))
		return
	}

	var updateData models.MobileAppHeader
	if err := c.ShouldBindJSON(&updateData); err != nil {
		fmt.Printf("❌ خطا در پارس کردن JSON: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	// بررسی وجود نام تکراری (اگر نام تغییر کرده باشد)
	if updateData.Name != header.Name {
		var existingHeader models.MobileAppHeader
		if err := h.settingService.DB().Where("name = ? AND id != ?", updateData.Name, id).First(&existingHeader).Error; err == nil {
			fmt.Printf("❌ نام هدر تکراری است: %s\n", updateData.Name)
			c.AbortWithStatusJSON(http.StatusConflict, utils.New("DUPLICATE_ERROR", "نام هدر موبایل و اپلیکیشن تکراری است", nil))
			return
		}
	}

	// به‌روزرسانی هدر
	if err := h.settingService.DB().Model(&header).Updates(updateData).Error; err != nil {
		fmt.Printf("❌ خطا در به‌روزرسانی هدر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی هدر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ هدر موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد: %s (ID: %d)\n", header.Name, header.ID)

	// دریافت هدر به‌روزرسانی شده
	var updatedHeader models.MobileAppHeader
	if err := h.settingService.DB().First(&updatedHeader, uint(id)).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت هدر به‌روزرسانی شده: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت هدر به‌روزرسانی شده", err.Error()))
		return
	}

	// تبدیل به فرمت مناسب برای frontend
	headerData := map[string]interface{}{
		"id":                 updatedHeader.ID,
		"name":               updatedHeader.Name,
		"description":        updatedHeader.Description,
		"platform":           updatedHeader.Platform,
		"page_selection":     updatedHeader.PageSelection,
		"specific_pages":     updatedHeader.SpecificPages,
		"excluded_pages":     updatedHeader.ExcludedPages,
		"header_type":        updatedHeader.HeaderType,
		"logo_url":           updatedHeader.LogoURL,
		"logo_alt":           updatedHeader.LogoAlt,
		"show_search":        updatedHeader.ShowSearch,
		"show_cart":          updatedHeader.ShowCart,
		"show_user_menu":     updatedHeader.ShowUserMenu,
		"show_notifications": updatedHeader.ShowNotifications,
		"show_menu_button":   updatedHeader.ShowMenuButton,
		"background_color":   updatedHeader.BackgroundColor,
		"text_color":         updatedHeader.TextColor,
		"top_image_url":      updatedHeader.TopImageURL,
		"top_image_alt":      updatedHeader.TopImageAlt,
		"bottom_image_url":   updatedHeader.BottomImageURL,
		"bottom_image_alt":   updatedHeader.BottomImageAlt,
		"is_active":          updatedHeader.IsActive,
		"created_at":         updatedHeader.CreatedAt,
		"updated_at":         updatedHeader.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    headerData,
		"message": "هدر موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد",
	})
}

// DeleteMobileAppHeaderByID حذف هدر موبایل و اپلیکیشن بر اساس ID
func (h *MobileAppHeaderHandler) DeleteMobileAppHeaderByID(c *gin.Context) {
	fmt.Printf("🔍 DeleteMobileAppHeaderByID: شروع حذف هدر موبایل و اپلیکیشن\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var header models.MobileAppHeader
	if err := h.settingService.DB().First(&header, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ هدر موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "هدر موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت هدر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت هدر موبایل و اپلیکیشن", err.Error()))
		return
	}

	// حذف هدر (لایه‌ها به صورت خودکار حذف می‌شوند به دلیل CASCADE)
	if err := h.settingService.DB().Delete(&header).Error; err != nil {
		fmt.Printf("❌ خطا در حذف هدر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف هدر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ هدر موبایل و اپلیکیشن با موفقیت حذف شد: %s (ID: %d)\n", header.Name, header.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "هدر موبایل و اپلیکیشن با موفقیت حذف شد",
	})
}

// UploadMobileAppHeaderLogo آپلود لوگو هدر موبایل و اپلیکیشن
func (h *MobileAppHeaderHandler) UploadMobileAppHeaderLogo(c *gin.Context) {
	fmt.Printf("🔍 UploadMobileAppHeaderLogo: شروع آپلود لوگو هدر موبایل و اپلیکیشن\n")

	// دریافت فایل از فرم
	file, header, err := c.Request.FormFile("logo")
	if err != nil {
		fmt.Printf("❌ خطا در دریافت فایل: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("FILE_ERROR", "خطا در دریافت فایل", err.Error()))
		return
	}
	defer file.Close()

	// بررسی نوع فایل
	if !utils.IsValidImageType(header) {
		fmt.Printf("❌ نوع فایل نامعتبر: %s\n", header.Header.Get("Content-Type"))
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("FILE_TYPE_ERROR", "نوع فایل نامعتبر. فقط تصاویر مجاز هستند", nil))
		return
	}

	// بررسی اندازه فایل (حداکثر 5MB)
	if header.Size > 5*1024*1024 {
		fmt.Printf("❌ اندازه فایل بیش از حد مجاز: %d bytes\n", header.Size)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("FILE_SIZE_ERROR", "اندازه فایل بیش از حد مجاز (حداکثر 5MB)", nil))
		return
	}

	// آپلود فایل (اینجا باید منطق آپلود واقعی پیاده‌سازی شود)
	logoURL, err := h.uploadFile(file, header)
	if err != nil {
		fmt.Printf("❌ خطا در آپلود فایل: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("UPLOAD_ERROR", "خطا در آپلود فایل", err.Error()))
		return
	}

	fmt.Printf("✅ لوگو هدر موبایل و اپلیکیشن با موفقیت آپلود شد: %s\n", logoURL)

	response := models.MobileAppHeaderLogoUploadResponse{
		Success: true,
		Data: struct {
			LogoURL string `json:"logo_url"`
		}{
			LogoURL: logoURL,
		},
		Message: "لوگو هدر موبایل و اپلیکیشن با موفقیت آپلود شد",
	}

	c.JSON(http.StatusOK, response)
}

func (h *MobileAppHeaderHandler) uploadFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	// اینجا باید منطق آپلود واقعی پیاده‌سازی شود
	// برای مثال، آپلود به cloud storage یا ذخیره در پوشه محلی
	// فعلاً یک URL نمونه برمی‌گردانیم
	return fmt.Sprintf("/uploads/mobile-app-headers/%s", header.Filename), nil
}
