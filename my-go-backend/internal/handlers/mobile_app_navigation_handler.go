package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
)

type MobileAppNavigationHandler struct {
	settingService *services.SettingService
}

// NewMobileAppNavigationHandler یک نمونه جدید از کنترلر تنظیمات ناوبری موبایل و اپلیکیشن ایجاد می‌کند
func NewMobileAppNavigationHandler(settingService *services.SettingService) *MobileAppNavigationHandler {
	return &MobileAppNavigationHandler{
		settingService: settingService,
	}
}

// GetMobileAppNavigationSettings دریافت تنظیمات ناوبری موبایل و اپلیکیشن برای نمایش عمومی
func (h *MobileAppNavigationHandler) GetMobileAppNavigationSettings(c *gin.Context) {
	fmt.Printf("🔍 GetMobileAppNavigationSettings: شروع دریافت تنظیمات ناوبری موبایل و اپلیکیشن\n")

	// بررسی User-Agent برای تشخیص موبایل - فقط در موبایل نمایش داده شود
	userAgent := c.GetHeader("User-Agent")
	fmt.Printf("🔍 User-Agent: %s\n", userAgent)
	if IsDesktopDevice(userAgent) {
		fmt.Printf("⚠️ درخواست از دستگاه دسکتاپ: %s - ناوبری موبایل در دسکتاپ نمایش داده نمی‌شود\n", userAgent)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []interface{}{},
			"message": "ناوبری موبایل فقط در دستگاه‌های موبایل نمایش داده می‌شود",
		})
		return
	}

	fmt.Printf("📱 درخواست از دستگاه موبایل: %s\n", userAgent)

	// دریافت ناوبری‌های فعال از جدول mobile_app_navigations
	var navigations []models.MobileAppNavigation
	if err := h.settingService.DB().Where("is_active = ?", true).Find(&navigations).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت ناوبری‌های موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات ناوبری موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("📊 تعداد ناوبری‌های فعال موبایل و اپلیکیشن یافت شده: %d\n", len(navigations))

	// اگر ناوبری‌ای وجود ندارد، پیام مناسب برگردان
	if len(navigations) == 0 {
		fmt.Printf("⚠️ هیچ ناوبری فعال موبایل و اپلیکیشنی یافت نشد\n")
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []interface{}{},
			"message": "هیچ ناوبری فعال موبایل و اپلیکیشنی یافت نشد",
		})
		return
	}

	// تبدیل به فرمت مناسب برای frontend
	var responseData []map[string]interface{}
	for _, navigation := range navigations {
		navigationData := map[string]interface{}{
			"id":               navigation.ID,
			"name":             navigation.Name,
			"description":      navigation.Description,
			"platform":         navigation.Platform,
			"page_selection":   navigation.PageSelection,
			"specific_pages":   navigation.SpecificPages,
			"excluded_pages":   navigation.ExcludedPages,
			"is_active":        navigation.IsActive,
			"created_at":       navigation.CreatedAt,
			"updated_at":       navigation.UpdatedAt,
			"navigation_items": navigation.NavigationItems,
			"layers":           navigation.Layers,
		}
		responseData = append(responseData, navigationData)
	}

	fmt.Printf("✅ تنظیمات ناوبری موبایل و اپلیکیشن با موفقیت ارسال شد\n")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    responseData,
		"message": "تنظیمات ناوبری موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// ListMobileAppNavigations لیست کردن تمام ناوبری‌های موبایل و اپلیکیشن (برای ادمین)
func (h *MobileAppNavigationHandler) ListMobileAppNavigations(c *gin.Context) {
	fmt.Printf("🔍 ListMobileAppNavigations: شروع لیست کردن ناوبری‌های موبایل و اپلیکیشن\n")

	var navigations []models.MobileAppNavigation
	if err := h.settingService.DB().Find(&navigations).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت لیست ناوبری‌های موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست ناوبری‌های موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("📊 تعداد کل ناوبری‌های موبایل و اپلیکیشن: %d\n", len(navigations))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    navigations,
		"message": "لیست ناوبری‌های موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// CreateMobileAppNavigation ایجاد ناوبری جدید موبایل و اپلیکیشن
func (h *MobileAppNavigationHandler) CreateMobileAppNavigation(c *gin.Context) {
	fmt.Printf("🔍 CreateMobileAppNavigation: شروع ایجاد ناوبری جدید موبایل و اپلیکیشن\n")

	var requestData struct {
		Name            string   `json:"name" binding:"required"`
		Description     string   `json:"description"`
		Platform        string   `json:"platform" binding:"required"`
		BackgroundColor string   `json:"background_color"`
		TextColor       string   `json:"text_color"`
		NavigationItems []string `json:"navigation_items"`
		PageSelection   string   `json:"page_selection"`
		SpecificPages   string   `json:"specific_pages"`
		ExcludedPages   string   `json:"excluded_pages"`
		IsActive        bool     `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		fmt.Printf("❌ خطا در پارس کردن JSON: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	// تبدیل navigation_items به JSON
	var navigationItemsJSON string
	if len(requestData.NavigationItems) > 0 {
		jsonBytes, err := json.Marshal(requestData.NavigationItems)
		if err != nil {
			fmt.Printf("❌ خطا در تبدیل navigation_items به JSON: %v\n", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "خطا در تبدیل آیتم‌های ناوبری", err.Error()))
			return
		}
		navigationItemsJSON = string(jsonBytes)
	} else {
		navigationItemsJSON = "[]"
	}

	// ایجاد مدل ناوبری
	navigation := models.MobileAppNavigation{
		Name:            requestData.Name,
		Description:     requestData.Description,
		Platform:        requestData.Platform,
		BackgroundColor: requestData.BackgroundColor,
		TextColor:       requestData.TextColor,
		NavigationItems: string(navigationItemsJSON),
		PageSelection:   requestData.PageSelection,
		SpecificPages:   requestData.SpecificPages,
		ExcludedPages:   requestData.ExcludedPages,
		IsActive:        requestData.IsActive,
	}

	// بررسی وجود نام تکراری
	var existingNavigation models.MobileAppNavigation
	if err := h.settingService.DB().Where("name = ?", navigation.Name).First(&existingNavigation).Error; err == nil {
		fmt.Printf("❌ نام ناوبری تکراری است: %s\n", navigation.Name)
		c.AbortWithStatusJSON(http.StatusConflict, utils.New("DUPLICATE_ERROR", "نام ناوبری موبایل و اپلیکیشن تکراری است", nil))
		return
	}

	// ایجاد ناوبری جدید
	if err := h.settingService.DB().Create(&navigation).Error; err != nil {
		fmt.Printf("❌ خطا در ایجاد ناوبری موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد ناوبری موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ ناوبری موبایل و اپلیکیشن با موفقیت ایجاد شد: %s (ID: %d)\n", navigation.Name, navigation.ID)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    navigation,
		"message": "ناوبری موبایل و اپلیکیشن با موفقیت ایجاد شد",
	})
}

// GetMobileAppNavigationByID دریافت ناوبری موبایل و اپلیکیشن بر اساس ID
func (h *MobileAppNavigationHandler) GetMobileAppNavigationByID(c *gin.Context) {
	fmt.Printf("🔍 GetMobileAppNavigationByID: شروع دریافت ناوبری موبایل و اپلیکیشن بر اساس ID\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var navigation models.MobileAppNavigation
	if err := h.settingService.DB().First(&navigation, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ ناوبری موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "ناوبری موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت ناوبری موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت ناوبری موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ ناوبری موبایل و اپلیکیشن با موفقیت دریافت شد: %s (ID: %d)\n", navigation.Name, navigation.ID)

	// تبدیل به فرمت مناسب برای frontend
	navigationData := map[string]interface{}{
		"id":               navigation.ID,
		"name":             navigation.Name,
		"description":      navigation.Description,
		"platform":         navigation.Platform,
		"background_color": navigation.BackgroundColor,
		"text_color":       navigation.TextColor,
		"navigation_items": navigation.NavigationItems,
		"page_selection":   navigation.PageSelection,
		"specific_pages":   navigation.SpecificPages,
		"excluded_pages":   navigation.ExcludedPages,
		"is_active":        navigation.IsActive,
		"created_at":       navigation.CreatedAt,
		"updated_at":       navigation.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    navigationData,
		"message": "ناوبری موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// UpdateMobileAppNavigationByID به‌روزرسانی ناوبری موبایل و اپلیکیشن بر اساس ID
func (h *MobileAppNavigationHandler) UpdateMobileAppNavigationByID(c *gin.Context) {
	fmt.Printf("🔍 UpdateMobileAppNavigationByID: شروع به‌روزرسانی ناوبری موبایل و اپلیکیشن\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var navigation models.MobileAppNavigation
	if err := h.settingService.DB().First(&navigation, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ ناوبری موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "ناوبری موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت ناوبری موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت ناوبری موبایل و اپلیکیشن", err.Error()))
		return
	}

	var updateData models.MobileAppNavigation
	if err := c.ShouldBindJSON(&updateData); err != nil {
		fmt.Printf("❌ خطا در پارس کردن JSON: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	// بررسی وجود نام تکراری (اگر نام تغییر کرده باشد)
	if updateData.Name != navigation.Name {
		var existingNavigation models.MobileAppNavigation
		if err := h.settingService.DB().Where("name = ? AND id != ?", updateData.Name, id).First(&existingNavigation).Error; err == nil {
			fmt.Printf("❌ نام ناوبری تکراری است: %s\n", updateData.Name)
			c.AbortWithStatusJSON(http.StatusConflict, utils.New("DUPLICATE_ERROR", "نام ناوبری موبایل و اپلیکیشن تکراری است", nil))
			return
		}
	}

	// به‌روزرسانی ناوبری
	if err := h.settingService.DB().Model(&navigation).Updates(updateData).Error; err != nil {
		fmt.Printf("❌ خطا در به‌روزرسانی ناوبری موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی ناوبری موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ ناوبری موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد: %s (ID: %d)\n", navigation.Name, navigation.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    navigation,
		"message": "ناوبری موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد",
	})
}

// DeleteMobileAppNavigationByID حذف ناوبری موبایل و اپلیکیشن بر اساس ID
func (h *MobileAppNavigationHandler) DeleteMobileAppNavigationByID(c *gin.Context) {
	fmt.Printf("🔍 DeleteMobileAppNavigationByID: شروع حذف ناوبری موبایل و اپلیکیشن\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var navigation models.MobileAppNavigation
	if err := h.settingService.DB().First(&navigation, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ ناوبری موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "ناوبری موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت ناوبری موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت ناوبری موبایل و اپلیکیشن", err.Error()))
		return
	}

	// حذف ناوبری (لایه‌ها به صورت خودکار حذف می‌شوند به دلیل CASCADE)
	if err := h.settingService.DB().Delete(&navigation).Error; err != nil {
		fmt.Printf("❌ خطا در حذف ناوبری موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف ناوبری موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("✅ ناوبری موبایل و اپلیکیشن با موفقیت حذف شد: %s (ID: %d)\n", navigation.Name, navigation.ID)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ناوبری موبایل و اپلیکیشن با موفقیت حذف شد",
	})
}

// GetMobileAppNavigationLayers دریافت لایه‌های ناوبری موبایل و اپلیکیشن
func (h *MobileAppNavigationHandler) GetMobileAppNavigationLayers(c *gin.Context) {
	fmt.Printf("🔍 GetMobileAppNavigationLayers: شروع دریافت لایه‌های ناوبری موبایل و اپلیکیشن\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	var layers []models.MobileAppNavigationLayer
	if err := h.settingService.DB().Where("mobile_app_navigation_id = ?", id).Find(&layers).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت لایه‌های ناوبری موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لایه‌های ناوبری موبایل و اپلیکیشن", err.Error()))
		return
	}

	fmt.Printf("📊 تعداد لایه‌های ناوبری موبایل و اپلیکیشن: %d\n", len(layers))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    layers,
		"message": "لایه‌های ناوبری موبایل و اپلیکیشن با موفقیت دریافت شد",
	})
}

// UpdateMobileAppNavigationLayers به‌روزرسانی لایه‌های ناوبری موبایل و اپلیکیشن
func (h *MobileAppNavigationHandler) UpdateMobileAppNavigationLayers(c *gin.Context) {
	fmt.Printf("🔍 UpdateMobileAppNavigationLayers: شروع به‌روزرسانی لایه‌های ناوبری موبایل و اپلیکیشن\n")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Printf("❌ ID نامعتبر: %s\n", idStr)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ID نامعتبر", err.Error()))
		return
	}

	// بررسی وجود ناوبری
	var navigation models.MobileAppNavigation
	if err := h.settingService.DB().First(&navigation, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("❌ ناوبری موبایل و اپلیکیشن یافت نشد: ID %d\n", id)
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "ناوبری موبایل و اپلیکیشن یافت نشد", nil))
			return
		}
		fmt.Printf("❌ خطا در دریافت ناوبری موبایل و اپلیکیشن: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت ناوبری موبایل و اپلیکیشن", err.Error()))
		return
	}

	var layers []models.MobileAppNavigationLayer
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
	if err := tx.Where("mobile_app_navigation_id = ?", id).Delete(&models.MobileAppNavigationLayer{}).Error; err != nil {
		tx.Rollback()
		fmt.Printf("❌ خطا در حذف لایه‌های موجود: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف لایه‌های موجود", err.Error()))
		return
	}

	// اضافه کردن لایه‌های جدید
	for i := range layers {
		layers[i].ID = 0 // اطمینان از ایجاد رکورد جدید
		layers[i].MobileAppNavigationID = uint(id)
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

	fmt.Printf("✅ لایه‌های ناوبری موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد: %d لایه\n", len(layers))

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    layers,
		"message": "لایه‌های ناوبری موبایل و اپلیکیشن با موفقیت به‌روزرسانی شد",
	})
}

// UploadMobileAppNavigationLogo آپلود لوگو ناوبری موبایل و اپلیکیشن
func (h *MobileAppNavigationHandler) UploadMobileAppNavigationLogo(c *gin.Context) {
	fmt.Printf("🔍 UploadMobileAppNavigationLogo: شروع آپلود لوگو ناوبری موبایل و اپلیکیشن\n")

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

	fmt.Printf("✅ لوگو ناوبری موبایل و اپلیکیشن با موفقیت آپلود شد: %s\n", logoURL)

	response := models.MobileAppNavigationLogoUploadResponse{
		Success: true,
		Data: struct {
			LogoURL string `json:"logo_url"`
		}{
			LogoURL: logoURL,
		},
		Message: "لوگو ناوبری موبایل و اپلیکیشن با موفقیت آپلود شد",
	}

	c.JSON(http.StatusOK, response)
}

func (h *MobileAppNavigationHandler) uploadFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	// اینجا باید منطق آپلود واقعی پیاده‌سازی شود
	// برای مثال، آپلود به cloud storage یا ذخیره در پوشه محلی
	// فعلاً یک URL نمونه برمی‌گردانیم
	return fmt.Sprintf("/uploads/mobile-app-navigations/%s", header.Filename), nil
}
