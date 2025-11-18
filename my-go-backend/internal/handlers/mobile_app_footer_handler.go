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

type MobileAppFooterHandler struct {
	settingService *services.SettingService
}

// NewMobileAppFooterHandler یک نمونه جدید از کنترلر تنظیمات فوتر موبایل و اپلیکیشن ایجاد می‌کند
func NewMobileAppFooterHandler(settingService *services.SettingService) *MobileAppFooterHandler {
	return &MobileAppFooterHandler{
		settingService: settingService,
	}
}

// GetMobileAppFooterSettings دریافت تنظیمات فوتر موبایل و اپلیکیشن برای نمایش عمومی
func (h *MobileAppFooterHandler) GetMobileAppFooterSettings(c *gin.Context) {
	fmt.Printf("🔍 GetMobileAppFooterSettings: شروع دریافت تنظیمات فوتر موبایل و اپلیکیشن\n")

	// دریافت فوترهای فعال از جدول mobile_app_footers
	var footers []models.MobileAppFooter
	if err := h.settingService.DB().Where("is_active = ?", true).Find(&footers).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت فوترهای موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات فوتر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("📊 تعداد فوترهای فعال موبایل و اپلیکیشن یافت شده: %d\n", len(footers))

	// اگر فوتری وجود ندارد، پیام مناسب برگردان
	if len(footers) == 0 {
		fmt.Printf("⚠️ هیچ فوتر فعال موبایل و اپلیکیشنی یافت نشد\n")
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []interface{}{},
			"message": "هیچ فوتر فعال موبایل و اپلیکیشنی یافت نشد",
		})
		return
	}

	// تبدیل به فرمت مناسب برای frontend
	var responseData []map[string]interface{}
	for _, footer := range footers {
		footerData := map[string]interface{}{
			"id":             footer.ID,
			"name":           footer.Name,
			"description":    footer.Description,
			"platform":       footer.Platform,
			"page_selection": footer.PageSelection,
			"specific_pages": footer.SpecificPages,
			"excluded_pages": footer.ExcludedPages,
			"is_active":      footer.IsActive,
			"created_at":     footer.CreatedAt,
			"updated_at":     footer.UpdatedAt,
			"layers":         footer.Layers,
		}
		responseData = append(responseData, footerData)
	}

	fmt.Printf("✅ تنظیمات فوتر موبایل و اپلیکیشن با موفقیت ارسال شد\n")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    responseData,
		"message": "تنظیمات فوتر موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// ListMobileAppFooters لیست کردن تمام فوترهای موبایل و اپلیکیشن (برای ادمین)
func (h *MobileAppFooterHandler) ListMobileAppFooters(c *gin.Context) {
	fmt.Printf("🔍 ListMobileAppFooters: شروع لیست کردن فوترهای موبایل و اپلیکیشن\n")

	var footers []models.MobileAppFooter
	if err := h.settingService.DB().Find(&footers).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت لیست فوترهای موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست فوترهای موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("📊 تعداد کل فوترهای موبایل و اپلیکیشن: %d\n", len(footers))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    footers,
		"message": "لیست فوترهای موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// CreateMobileAppFooter ایجاد فوتر جدید موبایل و اپلیکیشن
func (h *MobileAppFooterHandler) CreateMobileAppFooter(c *gin.Context) {
	fmt.Printf("🔍 CreateMobileAppFooter: شروع ایجاد فوتر جدید موبایل و اپلیکیشن\n")

	var footer models.MobileAppFooter
	if err := c.ShouldBindJSON(&footer); err != nil {
		fmt.Printf("❌ خطا در پارس کردن JSON: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	// بررسی وجود نام تکراری
	var existingFooter models.MobileAppFooter
	if err := h.settingService.DB().Where("name = ?", footer.Name).First(&existingFooter).Error; err == nil {
		fmt.Printf("❌ نام فوتر تکراری است: %s\n", footer.Name)
		c.AbortWithStatusJSON(http.StatusConflict, utils.New("DUPLICATE_ERROR", "نام فوتر موبایل و اپلیکیشن تکراری است", nil))
		return
	}

	// ایجاد فوتر جدید
	if err := h.settingService.DB().Create(&footer).Error; err != nil {
		fmt.Printf("❌ خطا در ایجاد فوتر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد فوتر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ فوتر موبایل و اپلیکیشن با موفقیت ایجاد شد: %s (ID: %d)\n", footer.Name, footer.ID)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    footer,
		"message": "فوتر موبایل و اپلیکیشن با موفقیت ایجاد شد",
	})
}

// GetMobileAppFooterByID دریافت فوتر موبایل و اپلیکیشن بر اساس ID
func (h *MobileAppFooterHandler) GetMobileAppFooterByID(c *gin.Context) {
	fmt.Printf("🔍 GetMobileAppFooterByID: شروع دریافت فوتر موبایل و اپلیکیشن بر اساس ID\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var footer models.MobileAppFooter
	if err := h.settingService.DB().First(&footer, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ فوتر موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "فوتر موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت فوتر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت فوتر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ فوتر موبایل و اپلیکیشن با موفقیت دریافت شد: %s (ID: %d)\n", footer.Name, footer.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    footer,
		"message": "فوتر موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// UpdateMobileAppFooterByID به‌روزرسانی فوتر موبایل و اپلیکیشن بر اساس ID
func (h *MobileAppFooterHandler) UpdateMobileAppFooterByID(c *gin.Context) {
	fmt.Printf("🔍 UpdateMobileAppFooterByID: شروع به‌روزرسانی فوتر موبایل و اپلیکیشن\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var footer models.MobileAppFooter
	if err := h.settingService.DB().First(&footer, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ فوتر موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "فوتر موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت فوتر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت فوتر موبایل و اپلیکیشن", err.Error()))
		return
	}

	var updateData models.MobileAppFooter
	if err := c.ShouldBindJSON(&updateData); err != nil {
		fmt.Printf("❌ خطا در پارس کردن JSON: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	// بررسی وجود نام تکراری (اگر نام تغییر کرده باشد)
	if updateData.Name != footer.Name {
		var existingFooter models.MobileAppFooter
		if err := h.settingService.DB().Where("name = ? AND id != ?", updateData.Name, id).First(&existingFooter).Error; err == nil {
			fmt.Printf("❌ نام فوتر تکراری است: %s\n", updateData.Name)
			c.AbortWithStatusJSON(http.StatusConflict, utils.New("DUPLICATE_ERROR", "نام فوتر موبایل و اپلیکیشن تکراری است", nil))
			return
		}
	}

	// به‌روزرسانی فوتر
	if err := h.settingService.DB().Model(&footer).Updates(updateData).Error; err != nil {
		fmt.Printf("❌ خطا در به‌روزرسانی فوتر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی فوتر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ فوتر موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد: %s (ID: %d)\n", footer.Name, footer.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    footer,
		"message": "فوتر موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد",
	})
}

// DeleteMobileAppFooterByID حذف فوتر موبایل و اپلیکیشن بر اساس ID
func (h *MobileAppFooterHandler) DeleteMobileAppFooterByID(c *gin.Context) {
	fmt.Printf("🔍 DeleteMobileAppFooterByID: شروع حذف فوتر موبایل و اپلیکیشن\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var footer models.MobileAppFooter
	if err := h.settingService.DB().First(&footer, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ فوتر موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "فوتر موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت فوتر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت فوتر موبایل و اپلیکیشن", err.Error()))
		return
	}

	// حذف فوتر (لایه‌ها به صورت خودکار حذف می‌شوند به دلیل CASCADE)
	if err := h.settingService.DB().Delete(&footer).Error; err != nil {
		fmt.Printf("❌ خطا در حذف فوتر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف فوتر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ فوتر موبایل و اپلیکیشن با موفقیت حذف شد: %s (ID: %d)\n", footer.Name, footer.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "فوتر موبایل و اپلیکیشن با موفقیت حذف شد",
	})
}

// GetMobileAppFooterLayers دریافت لایه‌های فوتر موبایل و اپلیکیشن
func (h *MobileAppFooterHandler) GetMobileAppFooterLayers(c *gin.Context) {
	fmt.Printf("🔍 GetMobileAppFooterLayers: شروع دریافت لایه‌های فوتر موبایل و اپلیکیشن\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var layers []models.MobileAppFooterLayer
	if err := h.settingService.DB().Where("mobile_app_footer_id = ?", id).Find(&layers).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت لایه‌های فوتر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لایه‌های فوتر موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("📊 تعداد لایه‌های فوتر موبایل و اپلیکیشن: %d\n", len(layers))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    layers,
		"message": "لایه‌های فوتر موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// UpdateMobileAppFooterLayers به‌روزرسانی لایه‌های فوتر موبایل و اپلیکیشن
func (h *MobileAppFooterHandler) UpdateMobileAppFooterLayers(c *gin.Context) {
	fmt.Printf("🔍 UpdateMobileAppFooterLayers: شروع به‌روزرسانی لایه‌های فوتر موبایل و اپلیکیشن\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	// بررسی وجود فوتر
	var footer models.MobileAppFooter
	if err := h.settingService.DB().First(&footer, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ فوتر موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "فوتر موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت فوتر موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت فوتر موبایل و اپلیکیشن", err.Error()))
		return
	}

	var layers []models.MobileAppFooterLayer
	if err := c.ShouldBindJSON(&layers); err != nil {
		fmt.Printf("❌ خطا در پارس کردن JSON: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	// شروع تراکنش
	tx := h.settingService.DB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// حذف لایه‌های موجود
	if err := tx.Where("mobile_app_footer_id = ?", id).Delete(&models.MobileAppFooterLayer{}).Error; err != nil {
		tx.Rollback()
		fmt.Printf("❌ خطا در حذف لایه‌های موجود: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف لایه‌های موجود", err.Error()))
		return
	}

	// اضافه کردن لایه‌های جدید
	for i := range layers {
		layers[i].ID = 0 // اطمینان از ایجاد رکورد جدید
		layers[i].MobileAppFooterID = uint(id)
	}

	if err := tx.Create(&layers).Error; err != nil {
		tx.Rollback()
		fmt.Printf("❌ خطا در ایجاد لایه‌های جدید: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد لایه‌های جدید", err.Error()))
		return
	}

	// تایید تراکنش
	if err := tx.Commit().Error; err != nil {
		fmt.Printf("❌ خطا در تایید تراکنش: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در تایید تراکنش", err.Error()))
		return
	}

	fmt.Printf("✅ لایه‌های فوتر موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد: %d لایه\n", len(layers))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    layers,
		"message": "لایه‌های فوتر موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد",
	})
}

// UploadMobileAppFooterLogo آپلود لوگو فوتر موبایل و اپلیکیشن
func (h *MobileAppFooterHandler) UploadMobileAppFooterLogo(c *gin.Context) {
	fmt.Printf("🔍 UploadMobileAppFooterLogo: شروع آپلود لوگو فوتر موبایل و اپلیکیشن\n")

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

	fmt.Printf("✅ لوگو فوتر موبایل و اپلیکیشن با موفقیت آپلود شد: %s\n", logoURL)

	response := models.MobileAppFooterLogoUploadResponse{
		Success: true,
		Data: struct {
			LogoURL string `json:"logo_url"`
		}{
			LogoURL: logoURL,
		},
		Message: "لوگو فوتر موبایل و اپلیکیشن با موفقیت آپلود شد",
	}

	c.JSON(http.StatusOK, response)
}

// Helper functions
func (h *MobileAppFooterHandler) uploadFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	// اینجا باید منطق آپلود واقعی پیاده‌سازی شود
	// برای مثال، آپلود به cloud storage یا ذخیره در پوشه محلی
	// فعلاً یک URL نمونه برمی‌گردانیم
	return fmt.Sprintf("/uploads/mobile-app-footers/%s", header.Filename), nil
}
