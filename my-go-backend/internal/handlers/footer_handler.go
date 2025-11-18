package handlers

import (
	"encoding/json"
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

type FooterHandler struct {
	settingService *services.SettingService
}

// NewFooterHandler یک نمونه جدید از کنترلر تنظیمات فوتر ایجاد می‌کند
func NewFooterHandler(settingService *services.SettingService) *FooterHandler {
	return &FooterHandler{
		settingService: settingService,
	}
}

func (h *FooterHandler) GetFooterSettings(c *gin.Context) {
	fmt.Printf("🔍 GetFooterSettings: Starting footer settings retrieval\n")

	// دریافت فوترهای فعال از جدول footers
	var footers []models.Footer
	if err := h.settingService.DB().Preload("Layers").Where("is_active = ?", true).Find(&footers).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت فوترها: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات فوتر", err.Error()))
		return
	}

	fmt.Printf("📊 Number of active footers found: %d\n", len(footers))

	// اگر فوتری وجود ندارد، پیام مناسب برگردان
	if len(footers) == 0 {
		fmt.Printf("⚠️ No active footers found\n")
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []interface{}{},
			"message": "هیچ فوتر فعالی یافت نشد",
		})
		return
	}

	// برگرداندن فوترهای فعال
	fmt.Printf("✅ فوترهای فعال برگردانده شدند\n")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    footers,
		"message": "فوترهای فعال با موفقیت دریافت شد",
	})
}

// لیست همه فوترها
func (h *FooterHandler) ListFooters(c *gin.Context) {
	var footers []models.Footer
	if err := h.settingService.DB().Preload("Layers").Find(&footers).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست فوترها", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": footers})
}

// ایجاد فوتر جدید
func (h *FooterHandler) CreateFooter(c *gin.Context) {
	var req models.Footer
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

	// ایجاد فوتر با لایه‌ها در یک تراکنش
	tx := h.settingService.DB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ایجاد فوتر
	if err := tx.Create(&req).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد فوتر", err.Error()))
		return
	}

	// commit تراکنش
	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در commit تراکنش", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": req})
}

// دریافت یک فوتر خاص
func (h *FooterHandler) GetFooterByID(c *gin.Context) {
	id := c.Param("id")
	var footer models.Footer
	if err := h.settingService.DB().Preload("Layers").First(&footer, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "فوتر پیدا نشد", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": footer})
}

// ویرایش فوتر و لایه‌ها
func (h *FooterHandler) UpdateFooterByID(c *gin.Context) {
	id := c.Param("id")
	var req models.Footer
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
	var footer models.Footer
	if err := h.settingService.DB().Preload("Layers").First(&footer, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "فوتر پیدا نشد", err.Error()))
		return
	}
	// به‌روزرسانی فیلدهای اصلی
	footer.Name = req.Name
	footer.Description = req.Description
	footer.PageSelection = req.PageSelection
	footer.SpecificPages = req.SpecificPages
	footer.ExcludedPages = req.ExcludedPages
	footer.IsActive = req.IsActive
	// --- مدیریت لایه‌ها ---
	// دریافت لیست فعلی لایه‌ها از پایگاه‌داده
	var existingLayers []models.FooterLayer
	if err := h.settingService.DB().Where("footer_id = ?", footer.ID).Find(&existingLayers).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در خواندن لایه‌های فعلی", err.Error()))
		return
	}

	// نگهداری IDهای لایه‌های ارسالی
	incomingIDs := make(map[uint]bool)

	for _, l := range req.Layers {
		if l.ID == 0 {
			// لایه جدید
			l.FooterID = footer.ID
			if err := h.settingService.DB().Create(&l).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد لایه جدید", err.Error()))
				return
			}
		} else {
			// به‌روزرسانی لایه موجود
			l.FooterID = footer.ID
			if err := h.settingService.DB().Model(&l).Where("id = ? AND footer_id = ?", l.ID, footer.ID).Updates(l).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی لایه", err.Error()))
				return
			}
		}
		incomingIDs[l.ID] = true
	}

	// حذف لایه‌هایی که کاربر در فرم حذف کرده است (در incomingIDs نیستند)
	for _, ex := range existingLayers {
		if !incomingIDs[ex.ID] {
			if err := h.settingService.DB().Delete(&models.FooterLayer{}, ex.ID).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف لایه", err.Error()))
				return
			}
		}
	}
	// بعد از افزودن لایه‌های جدید، فقط فیلدهای اصلی فوتر را به‌روزرسانی می‌کنیم (بدون دخالت در Associations)
	if err := h.settingService.DB().Model(&footer).Select("Name", "Description", "PageSelection", "SpecificPages", "ExcludedPages", "IsActive").Updates(footer).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در به‌روزرسانی فوتر", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": footer})
}

// حذف فوتر و لایه‌ها
func (h *FooterHandler) DeleteFooterByID(c *gin.Context) {
	id := c.Param("id")
	if err := h.settingService.DB().Delete(&models.Footer{}, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف فوتر", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "فوتر با موفقیت حذف شد"})
}

func (h *FooterHandler) DeleteFooter(c *gin.Context) {
	id := c.Param("id")
	footerID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه فوتر نامعتبر است", err.Error()))
		return
	}

	// بررسی وجود فوتر
	var footer models.Footer
	if err := h.settingService.DB().First(&footer, footerID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "فوتر مورد نظر یافت نشد", err.Error()))
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بررسی وجود فوتر", err.Error()))
		}
		return
	}

	// حذف فوتر (لایه‌ها به صورت خودکار حذف می‌شوند به دلیل CASCADE)
	if err := h.settingService.DB().Delete(&footer).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف فوتر", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "فوتر با موفقیت حذف شد",
	})
}

// convertFooterToSettings تبدیل تنظیمات فوتر به فرمت دیتابیس
func (h *FooterHandler) convertFooterToSettings(request *models.FooterSettingsRequest) []models.Setting {
	var settings []models.Setting

	// تبدیل فیلدهای request به تنظیمات
	if request.LogoURL != "" {
		settings = append(settings, models.Setting{Key: "footer_logo_url", Value: request.LogoURL, Category: "footer", Type: "string"})
	}
	if request.LogoAlt != "" {
		settings = append(settings, models.Setting{Key: "footer_logo_alt", Value: request.LogoAlt, Category: "footer", Type: "string"})
	}
	if request.ShowCopyright != nil {
		settings = append(settings, models.Setting{Key: "footer_show_copyright", Value: utils.BoolToString(*request.ShowCopyright), Category: "footer", Type: "boolean"})
	}
	if request.ShowSocialLinks != nil {
		settings = append(settings, models.Setting{Key: "footer_show_social_links", Value: utils.BoolToString(*request.ShowSocialLinks), Category: "footer", Type: "boolean"})
	}
	if request.PhoneNumber != "" {
		settings = append(settings, models.Setting{Key: "footer_phone_number", Value: request.PhoneNumber, Category: "footer", Type: "string"})
	}
	if request.Email != "" {
		settings = append(settings, models.Setting{Key: "footer_email", Value: request.Email, Category: "footer", Type: "string"})
	}
	if request.Address != "" {
		settings = append(settings, models.Setting{Key: "footer_address", Value: request.Address, Category: "footer", Type: "string"})
	}
	if request.FacebookURL != "" {
		settings = append(settings, models.Setting{Key: "footer_facebook_url", Value: request.FacebookURL, Category: "footer", Type: "string"})
	}
	if request.InstagramURL != "" {
		settings = append(settings, models.Setting{Key: "footer_instagram_url", Value: request.InstagramURL, Category: "footer", Type: "string"})
	}
	if request.TelegramURL != "" {
		settings = append(settings, models.Setting{Key: "footer_telegram_url", Value: request.TelegramURL, Category: "footer", Type: "string"})
	}
	if request.WhatsappURL != "" {
		settings = append(settings, models.Setting{Key: "footer_whatsapp_url", Value: request.WhatsappURL, Category: "footer", Type: "string"})
	}
	if request.BackgroundColor != "" {
		settings = append(settings, models.Setting{Key: "footer_background_color", Value: request.BackgroundColor, Category: "footer", Type: "string"})
	}
	if request.TextColor != "" {
		settings = append(settings, models.Setting{Key: "footer_text_color", Value: request.TextColor, Category: "footer", Type: "string"})
	}
	if request.Height != "" {
		settings = append(settings, models.Setting{Key: "footer_height", Value: request.Height, Category: "footer", Type: "string"})
	}

	return settings
}

// validateFooterSettings اعتبارسنجی تنظیمات فوتر
func (h *FooterHandler) validateFooterSettings(request *models.FooterSettingsRequest) error {
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
func (h *FooterHandler) validateLogoFile(file *multipart.FileHeader) error {
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
func (h *FooterHandler) uploadFile(file *multipart.FileHeader) (string, error) {
	// این قسمت باید با سیستم آپلود شما سازگار باشد
	// فعلاً یک مسیر نمونه برمی‌گردانیم
	return "/uploads/footers/" + file.Filename, nil
}

// GetFooterLayers دریافت لایه‌های یک فوتر خاص
func (h *FooterHandler) GetFooterLayers(c *gin.Context) {
	footerID := c.Param("id")
	if footerID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه فوتر الزامی است", nil))
		return
	}

	footerIDUint, err := strconv.ParseUint(footerID, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه فوتر نامعتبر است", err.Error()))
		return
	}

	var layers []models.FooterLayer
	if err := h.settingService.DB().Where("footer_id = ?", footerIDUint).Find(&layers).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لایه‌های فوتر", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    layers,
		"message": "لایه‌های فوتر با موفقیت دریافت شدند",
	})
}

// UpdateFooterLayers به‌روزرسانی لایه‌های یک فوتر خاص
func (h *FooterHandler) UpdateFooterLayers(c *gin.Context) {
	footerID := c.Param("id")
	if footerID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه فوتر الزامی است", nil))
		return
	}

	footerIDUint, err := strconv.ParseUint(footerID, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه فوتر نامعتبر است", err.Error()))
		return
	}

	var request struct {
		Layers []models.FooterLayer `json:"layers" binding:"required"`
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
	if err := tx.Where("footer_id = ?", footerIDUint).Delete(&models.FooterLayer{}).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف لایه‌های موجود", err.Error()))
		return
	}

	// اضافه کردن لایه‌های جدید
	for i := range request.Layers {
		request.Layers[i].FooterID = uint(footerIDUint)
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
		"message": "لایه‌های فوتر با موفقیت به‌روزرسانی شدند",
		"data":    request.Layers,
	})
}
