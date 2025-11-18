package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"errors"
	"mime/multipart"
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
	"strconv"

	"gorm.io/gorm"
)

type HeaderHandler struct {
	settingService *services.SettingService
}

// NewHeaderHandler یک نمونه جدید از کنترلر تنظیمات هدر ایجاد می‌کند
func NewHeaderHandler(settingService *services.SettingService) *HeaderHandler {
	return &HeaderHandler{
		settingService: settingService,
	}
}

func (h *HeaderHandler) GetHeaderSettings(c *gin.Context) {
	fmt.Printf("🔍 GetHeaderSettings: شروع دریافت تنظیمات هدر\n")

	// بررسی User-Agent برای تشخیص دسکتاپ
	userAgent := c.GetHeader("User-Agent")
	if !IsDesktopDevice(userAgent) {
		fmt.Printf("⚠️ درخواست از دستگاه غیر دسکتاپ: %s\n", userAgent)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []interface{}{},
			"message": "هدر دسکتاپ فقط در دسکتاپ نمایش داده می‌شود",
		})
		return
	}

	// دریافت هدرهای فعال از جدول headers
	var headers []models.Header
	if err := h.settingService.DB().Preload("Layers").Where("is_active = ?", true).Find(&headers).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت هدرها: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات هدر", err.Error()))
		return
	}

	fmt.Printf("📊 Number of active headers found: %d\n", len(headers))

	// اگر هدری وجود ندارد، پیام مناسب برگردان
	if len(headers) == 0 {
		fmt.Printf("⚠️ No active headers found\n")
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []interface{}{},
			"message": "No active headers found",
		})
		return
	}

	// برگرداندن هدرهای فعال
	fmt.Printf("✅ Active headers returned successfully\n")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    headers,
		"message": "Active headers retrieved successfully",
	})
}

// لیست همه هدرها
func (h *HeaderHandler) ListHeaders(c *gin.Context) {
	var headers []models.Header
	if err := h.settingService.DB().Preload("Layers").Find(&headers).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست هدرها", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": headers})
}

// ایجاد هدر جدید
func (h *HeaderHandler) CreateHeader(c *gin.Context) {
	var req models.Header
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	// تبدیل items هر لایه به رشته JSON معتبر
	for i := range req.Layers {
		if !json.Valid([]byte(req.Layers[i].Items)) {
			b, _ := json.Marshal(req.Layers[i].Items)
			req.Layers[i].Items = string(b)
		}
	}

	// ایجاد هدر با لایه‌ها در یک تراکنش
	tx := h.settingService.DB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ایجاد هدر
	if err := tx.Create(&req).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد هدر", err.Error()))
		return
	}

	// commit تراکنش
	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در commit تراکنش", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": req})
}

// دریافت یک هدر خاص
func (h *HeaderHandler) GetHeaderByID(c *gin.Context) {
	id := c.Param("id")
	var header models.Header
	if err := h.settingService.DB().Preload("Layers").First(&header, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "هدر پیدا نشد", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": header})
}

// ویرایش هدر و لایه‌ها
func (h *HeaderHandler) UpdateHeaderByID(c *gin.Context) {
	id := c.Param("id")
	var req models.Header
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}
	// تبدیل items هر لایه به رشته JSON معتبر
	for i := range req.Layers {
		if !json.Valid([]byte(req.Layers[i].Items)) {
			b, _ := json.Marshal(req.Layers[i].Items)
			req.Layers[i].Items = string(b)
		}
	}
	var header models.Header
	if err := h.settingService.DB().Preload("Layers").First(&header, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "هدر پیدا نشد", err.Error()))
		return
	}
	// به‌روزرسانی فیلدهای اصلی
	header.Name = req.Name
	header.Description = req.Description
	header.PageSelection = req.PageSelection
	header.SpecificPages = req.SpecificPages
	header.ExcludedPages = req.ExcludedPages
	header.IsActive = req.IsActive
	// --- مدیریت لایه‌ها ---
	// دریافت لیست فعلی لایه‌ها از پایگاه‌داده
	var existingLayers []models.HeaderLayer
	if err := h.settingService.DB().Where("header_id = ?", header.ID).Find(&existingLayers).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در خواندن لایه‌های فعلی", err.Error()))
		return
	}

	// نگهداری IDهای لایه‌های ارسالی
	incomingIDs := make(map[uint]bool)

	for _, l := range req.Layers {
		if l.ID == 0 {
			// لایه جدید
			l.HeaderID = header.ID
			if err := h.settingService.DB().Create(&l).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد لایه جدید", err.Error()))
				return
			}
		} else {
			// به‌روزرسانی لایه موجود
			l.HeaderID = header.ID
			if err := h.settingService.DB().Model(&l).Where("id = ? AND header_id = ?", l.ID, header.ID).Updates(l).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی لایه", err.Error()))
				return
			}
		}
		incomingIDs[l.ID] = true
	}

	// حذف لایه‌هایی که کاربر در فرم حذف کرده است (در incomingIDs نیستند)
	for _, ex := range existingLayers {
		if !incomingIDs[ex.ID] {
			if err := h.settingService.DB().Delete(&models.HeaderLayer{}, ex.ID).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف لایه", err.Error()))
				return
			}
		}
	}
	// بعد از افزودن لایه‌های جدید، فقط فیلدهای اصلی هدر را به‌روزرسانی می‌کنیم (بدون دخالت در Associations)
	if err := h.settingService.DB().Model(&header).Select("Name", "Description", "PageSelection", "SpecificPages", "ExcludedPages", "IsActive").Updates(header).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی هدر", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": header})
}

// حذف هدر و لایه‌ها
func (h *HeaderHandler) DeleteHeaderByID(c *gin.Context) {
	id := c.Param("id")
	if err := h.settingService.DB().Delete(&models.Header{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف هدر", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "هدر با موفقیت حذف شد"})
}

func (h *HeaderHandler) DeleteHeader(c *gin.Context) {
	id := c.Param("id")
	headerID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه هدر نامعتبر است", err.Error()))
		return
	}

	// بررسی وجود هدر
	var header models.Header
	if err := h.settingService.DB().First(&header, headerID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "هدر مورد نظر یافت نشد", err.Error()))
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بررسی وجود هدر", err.Error()))
		}
		return
	}

	// حذف هدر (لایه‌ها به صورت خودکار حذف می‌شوند به دلیل CASCADE)
	if err := h.settingService.DB().Delete(&header).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف هدر", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "هدر با موفقیت حذف شد",
	})
}

// convertHeaderToSettings تبدیل تنظیمات هدر به فرمت دیتابیس
func (h *HeaderHandler) convertHeaderToSettings(request *models.HeaderSettingsRequest) []models.Setting {
	var settings []models.Setting

	// تبدیل فیلدهای request به تنظیمات
	if request.LogoURL != "" {
		settings = append(settings, models.Setting{Key: "header_logo_url", Value: request.LogoURL, Category: "header", Type: "string"})
	}
	if request.LogoAlt != "" {
		settings = append(settings, models.Setting{Key: "header_logo_alt", Value: request.LogoAlt, Category: "header", Type: "string"})
	}
	if request.ShowSearch != nil {
		settings = append(settings, models.Setting{Key: "header_show_search", Value: utils.BoolToString(*request.ShowSearch), Category: "header", Type: "boolean"})
	}
	if request.ShowCart != nil {
		settings = append(settings, models.Setting{Key: "header_show_cart", Value: utils.BoolToString(*request.ShowCart), Category: "header", Type: "boolean"})
	}
	if request.ShowUserMenu != nil {
		settings = append(settings, models.Setting{Key: "header_show_user_menu", Value: utils.BoolToString(*request.ShowUserMenu), Category: "header", Type: "boolean"})
	}
	if request.PhoneNumber != "" {
		settings = append(settings, models.Setting{Key: "header_phone_number", Value: request.PhoneNumber, Category: "header", Type: "string"})
	}
	if request.Email != "" {
		settings = append(settings, models.Setting{Key: "header_email", Value: request.Email, Category: "header", Type: "string"})
	}
	if request.ShowSocialLinks != nil {
		settings = append(settings, models.Setting{Key: "header_show_social_links", Value: utils.BoolToString(*request.ShowSocialLinks), Category: "header", Type: "boolean"})
	}
	if request.FacebookURL != "" {
		settings = append(settings, models.Setting{Key: "header_facebook_url", Value: request.FacebookURL, Category: "header", Type: "string"})
	}
	if request.InstagramURL != "" {
		settings = append(settings, models.Setting{Key: "header_instagram_url", Value: request.InstagramURL, Category: "header", Type: "string"})
	}
	if request.TelegramURL != "" {
		settings = append(settings, models.Setting{Key: "header_telegram_url", Value: request.TelegramURL, Category: "header", Type: "string"})
	}
	if request.WhatsappURL != "" {
		settings = append(settings, models.Setting{Key: "header_whatsapp_url", Value: request.WhatsappURL, Category: "header", Type: "string"})
	}
	if request.Sticky != nil {
		settings = append(settings, models.Setting{Key: "header_sticky", Value: utils.BoolToString(*request.Sticky), Category: "header", Type: "boolean"})
	}
	if request.BackgroundColor != "" {
		settings = append(settings, models.Setting{Key: "header_background_color", Value: request.BackgroundColor, Category: "header", Type: "string"})
	}
	if request.TextColor != "" {
		settings = append(settings, models.Setting{Key: "header_text_color", Value: request.TextColor, Category: "header", Type: "string"})
	}
	if request.Height != "" {
		settings = append(settings, models.Setting{Key: "header_height", Value: request.Height, Category: "header", Type: "string"})
	}

	return settings
}

// validateHeaderSettings اعتبارسنجی تنظیمات هدر
func (h *HeaderHandler) validateHeaderSettings(request *models.HeaderSettingsRequest) error {
	// اعتبارسنجی ایمیل
	if request.Email != "" && !utils.IsValidEmail(request.Email) {
		return utils.NewError("ایمیل نامعتبر است")
	}

	// اعتبارسنجی رنگ‌ها
	if request.BackgroundColor != "" && !utils.IsValidColor(request.BackgroundColor) {
		return utils.NewError("رنگ پس‌زمینه نامعتبر است")
	}

	if request.TextColor != "" && !utils.IsValidColor(request.TextColor) {
		return utils.NewError("رنگ متن نامعتبر است")
	}

	// اعتبارسنجی URL ها
	if request.FacebookURL != "" && !utils.IsValidURL(request.FacebookURL) {
		return utils.NewError("لینک فیسبوک نامعتبر است")
	}

	if request.InstagramURL != "" && !utils.IsValidURL(request.InstagramURL) {
		return utils.NewError("لینک اینستاگرام نامعتبر است")
	}

	if request.TelegramURL != "" && !utils.IsValidURL(request.TelegramURL) {
		return utils.NewError("لینک تلگرام نامعتبر است")
	}

	if request.WhatsappURL != "" && !utils.IsValidURL(request.WhatsappURL) {
		return utils.NewError("لینک واتساپ نامعتبر است")
	}

	return nil
}

// validateLogoFile اعتبارسنجی فایل لوگو
func (h *HeaderHandler) validateLogoFile(file *multipart.FileHeader) error {
	// بررسی نوع فایل
	allowedTypes := []string{"image/jpeg", "image/jpg", "image/png", "image/gif", "image/webp"}
	if !utils.Contains(allowedTypes, file.Header.Get("Content-Type")) {
		return utils.NewError("نوع فایل پشتیبانی نمی‌شود. فقط JPG، PNG، GIF و WebP مجاز است")
	}

	// بررسی اندازه فایل (حداکثر 2MB)
	if file.Size > 2*1024*1024 {
		return utils.NewError("حجم فایل نباید بیشتر از 2 مگابایت باشد")
	}

	return nil
}

// uploadFile آپلود فایل (این تابع باید با سیستم آپلود شما سازگار باشد)
func (h *HeaderHandler) uploadFile(file *multipart.FileHeader) (string, error) {
	// این قسمت باید با سیستم آپلود شما سازگار باشد
	// فعلاً یک مسیر نمونه برمی‌گردانیم
	return "/uploads/headers/" + file.Filename, nil
}

// GetHeaderLayers دریافت لایه‌های یک هدر خاص
func (h *HeaderHandler) GetHeaderLayers(c *gin.Context) {
	headerID := c.Param("id")
	if headerID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه هدر الزامی است", nil))
		return
	}

	headerIDUint, err := strconv.ParseUint(headerID, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه هدر نامعتبر است", err.Error()))
		return
	}

	var layers []models.HeaderLayer
	if err := h.settingService.DB().Where("header_id = ?", headerIDUint).Find(&layers).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لایه‌های هدر", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    layers,
		"message": "لایه‌های هدر با موفقیت دریافت شدند",
	})
}

// UpdateHeaderLayers به‌روزرسانی لایه‌های یک هدر خاص
func (h *HeaderHandler) UpdateHeaderLayers(c *gin.Context) {
	headerID := c.Param("id")
	if headerID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه هدر الزامی است", nil))
		return
	}

	headerIDUint, err := strconv.ParseUint(headerID, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه هدر نامعتبر است", err.Error()))
		return
	}

	var request struct {
		Layers []models.HeaderLayer `json:"layers" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
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
	if err := tx.Where("header_id = ?", headerIDUint).Delete(&models.HeaderLayer{}).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف لایه‌های موجود", err.Error()))
		return
	}

	// اضافه کردن لایه‌های جدید
	for i := range request.Layers {
		request.Layers[i].HeaderID = uint(headerIDUint)
		request.Layers[i].ID = 0 // اطمینان از ایجاد لایه جدید

		// تبدیل items به رشته JSON اگر آرایه باشد
		if !json.Valid([]byte(request.Layers[i].Items)) {
			b, _ := json.Marshal(request.Layers[i].Items)
			request.Layers[i].Items = string(b)
		}

		if err := tx.Create(&request.Layers[i]).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد لایه جدید", err.Error()))
			return
		}
	}

	// تایید تراکنش
	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در تایید تراکنش", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "لایه‌های هدر با موفقیت به‌روزرسانی شدند",
		"data":    request.Layers,
	})
}

// IsDesktopDevice بررسی می‌کند که آیا درخواست از دستگاه دسکتاپ است یا نه
func IsDesktopDevice(userAgent string) bool {
	// تبدیل به حروف کوچک برای مقایسه بهتر
	userAgent = strings.ToLower(userAgent)

	// کلمات کلیدی که نشان‌دهنده دستگاه‌های غیر دسکتاپ هستند
	mobileKeywords := []string{
		"mobile", "android", "iphone", "ipad", "ipod",
		"blackberry", "windows phone", "opera mini",
		"mobile safari", "webos", "palm", "symbian",
		"kindle", "silk", "fennec", "maemo", "tablet",
	}

	// اگر هر یک از کلمات کلیدی موبایل در User-Agent وجود داشته باشد، دستگاه موبایل است
	for _, keyword := range mobileKeywords {
		if strings.Contains(userAgent, keyword) {
			return false
		}
	}

	// اگر هیچ کلمه کلیدی موبایل یافت نشد، احتمالاً دسکتاپ است
	return true
}
