package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
)

/*
کنترلر تنظیمات

این کنترلر برای مدیریت APIهای مربوط به تنظیمات سیستم استفاده می‌شود و شامل موارد زیر است:
- دریافت تنظیمات
- بروزرسانی تنظیمات
- مدیریت دسته‌بندی‌های مختلف تنظیمات
*/

type SettingHandler struct {
	settingService  *services.SettingService
	settingsService *services.SettingsService
}

// NewSettingHandler یک نمونه جدید از کنترلر تنظیمات ایجاد می‌کند
func NewSettingHandler(settingService *services.SettingService, settingsService *services.SettingsService) *SettingHandler {
	return &SettingHandler{
		settingService:  settingService,
		settingsService: settingsService,
	}
}

// GetSettings تمام تنظیمات را دریافت می‌کند
// @Summary دریافت تمام تنظیمات
// @Description دریافت لیست تمام تنظیمات سیستم یا تنظیمات یک دسته‌بندی خاص
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Param category query string false "دسته‌بندی تنظیمات (اختیاری)"
// @Success 200 {array} models.Setting
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/settings [get]
func (h *SettingHandler) GetSettings(c *gin.Context) {
	// بررسی وجود پارامتر category در query string
	category := c.Query("category")

	if category != "" {
		// اگر category مشخص شده، تنظیمات آن دسته‌بندی را برگردان
		settings, err := h.settingService.GetSettingsByCategory(c.Request.Context(), category)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات", err.Error()))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    settings,
		})
		return
	}

	// اگر category مشخص نشده، تمام تنظیمات را برگردان
	settings, err := h.settingService.GetAllSettings(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    settings,
	})
}

// GetSettingsByCategory تنظیمات یک دسته‌بندی خاص را دریافت می‌کند
// @Summary دریافت تنظیمات بر اساس دسته‌بندی
// @Description دریافت لیست تنظیمات یک دسته‌بندی خاص
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Param category path string true "دسته‌بندی تنظیمات"
// @Success 200 {array} models.Setting
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/settings/category/{category} [get]
func (h *SettingHandler) GetSettingsByCategory(c *gin.Context) {
	category := c.Param("category")
	if category == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "دسته‌بندی الزامی است", nil))
		return
	}

	settings, err := h.settingService.GetSettingsByCategory(c.Request.Context(), category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات", err.Error()))
		return
	}

	c.JSON(http.StatusOK, settings)
}

// GetHeaderSettings تنظیمات هدر را دریافت می‌کند
// @Summary دریافت تنظیمات هدر
// @Description دریافت لیست تنظیمات مربوط به هدر سایت
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Success 200 {array} models.Setting
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/header-settings [get]
func (h *SettingHandler) GetHeaderSettings(c *gin.Context) {
	settings, err := h.settingService.GetSettingsByCategory(c.Request.Context(), "header")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات هدر", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    settings,
	})
}

// GetSetting تنظیم مورد نظر را بر اساس کلید دریافت می‌کند
// @Summary دریافت تنظیم بر اساس کلید
// @Description دریافت تنظیم مورد نظر بر اساس کلید
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Param key path string true "کلید تنظیم"
// @Success 200 {object} models.Setting
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/settings/{key} [get]
func (h *SettingHandler) GetSetting(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "کلید تنظیم الزامی است", nil))
		return
	}

	setting, err := h.settingService.GetSetting(c.Request.Context(), key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیم", err.Error()))
		return
	}

	if setting == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("SETTING_NOT_FOUND", "تنظیم مورد نظر یافت نشد", nil))
		return
	}

	c.JSON(http.StatusOK, setting)
}

// UpdateSetting تنظیم مورد نظر را بروزرسانی می‌کند
// @Summary بروزرسانی تنظیم
// @Description بروزرسانی تنظیم مورد نظر
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Param key path string true "کلید تنظیم"
// @Param setting body models.Setting true "اطلاعات تنظیم"
// @Success 200 {object} models.Setting
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/settings/{key} [put]
func (h *SettingHandler) UpdateSetting(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "کلید تنظیم الزامی است", nil))
		return
	}

	var setting models.Setting
	if err := c.ShouldBindJSON(&setting); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "اطلاعات نامعتبر", err.Error()))
		return
	}

	setting.Key = key
	if err := h.settingService.UpdateSetting(c.Request.Context(), &setting); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بروزرسانی تنظیم", err.Error()))
		return
	}

	c.JSON(http.StatusOK, setting)
}

// UpdateSettings چندین تنظیم را همزمان بروزرسانی می‌کند
// @Summary بروزرسانی چندین تنظیم
// @Description بروزرسانی همزمان چندین تنظیم
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Param settings body []models.Setting true "لیست تنظیمات"
// @Success 200 {array} models.Setting
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/settings [put]
func (h *SettingHandler) UpdateSettings(c *gin.Context) {
	var settings []models.Setting
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "اطلاعات نامعتبر", err.Error()))
		return
	}

	if err := h.settingService.UpdateSettings(c.Request.Context(), settings); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بروزرسانی تنظیمات", err.Error()))
		return
	}

	c.JSON(http.StatusOK, settings)
}

// UpdateReviewSettings بروزرسانی تنظیمات نظرات
// @Summary بروزرسانی تنظیمات نظرات
// @Description مجاز کردن فرمت‌ها و محدودیت‌های آپلود نظرات
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Param settings body map[string]interface{} true "تنظیمات نظرات"
// @Success 200 {object} map[string]bool
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/settings/reviews [put]
func (h *SettingHandler) UpdateReviewSettings(c *gin.Context) {
	var payload struct {
		AllowedImageTypes string `json:"allowedImageTypes"`
		AllowedVideoTypes string `json:"allowedVideoTypes"`
		MaxFilesPerReview int    `json:"maxFilesPerReview"`
		MaxFileSizeMb     int    `json:"maxFileSizeMb"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "اطلاعات نامعتبر", err.Error()))
		return
	}

	settings := []models.Setting{
		{Key: "reviews.allowed_image_types", Value: payload.AllowedImageTypes, Category: "reviews", Type: "string"},
		{Key: "reviews.allowed_video_types", Value: payload.AllowedVideoTypes, Category: "reviews", Type: "string"},
		{Key: "reviews.max_files_per_review", Value: fmt.Sprintf("%d", payload.MaxFilesPerReview), Category: "reviews", Type: "number"},
		{Key: "reviews.max_file_size_mb", Value: fmt.Sprintf("%d", payload.MaxFileSizeMb), Category: "reviews", Type: "number"},
	}

	if err := h.settingService.UpdateSettings(c.Request.Context(), settings); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بروزرسانی تنظیمات نظرات", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// GetPublicSettings تنظیمات عمومی را دریافت می‌کند
// @Summary دریافت تنظیمات عمومی
// @Description دریافت لیست تنظیمات عمومی سیستم
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Success 200 {array} models.Setting
// @Failure 500 {object} models.ErrorResponse
// @Router /api/settings [get]
func (h *SettingHandler) GetPublicSettings(c *gin.Context) {
	settings, err := h.settingService.GetPublicSettings(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات", err.Error()))
		return
	}

	c.JSON(http.StatusOK, settings)
}

// ClearCache کش تنظیمات را پاک می‌کند
// @Summary پاک کردن کش تنظیمات
// @Description پاک کردن کش تمام تنظیمات
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/settings/cache [delete]
func (h *SettingHandler) ClearCache(c *gin.Context) {
	if err := h.settingService.ClearCache(c.Request.Context()); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("INTERNAL_ERROR", "خطا در پاک کردن کش", err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "کش با موفقیت پاک شد",
	})
}

// GetAuthSettings دریافت تنظیمات احراز هویت برای پنل ادمین
func (h *SettingHandler) GetAuthSettings(c *gin.Context) {
	settings, err := h.settingsService.GetAuthSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در دریافت تنظیمات احراز هویت",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "تنظیمات احراز هویت با موفقیت دریافت شد",
		"data":    settings,
	})
}

// UpdateAuthSettings بروزرسانی تنظیمات احراز هویت
// @Summary بروزرسانی تنظیمات احراز هویت
// @Description بروزرسانی تنظیمات احراز هویت شامل OTP، JWT و سایر تنظیمات امنیتی
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Param settings body map[string]interface{} true "تنظیمات احراز هویت"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/settings/auth [put]
func (h *SettingHandler) UpdateAuthSettings(c *gin.Context) {
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	// تبدیل داده‌های دریافتی به تنظیمات
	var settings []models.Setting

	// تنظیمات سیستم یکپارچه
	if val, exists := requestData["unified_auth_enabled"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.unified_auth_enabled",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "boolean",
			Description: "فعال‌سازی سیستم احراز هویت یکپارچه",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["auto_register_enabled"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.auto_register_enabled",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "boolean",
			Description: "ثبت‌نام خودکار کاربران",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["default_auth_method"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.default_auth_method",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "string",
			Description: "روش پیش‌فرض احراز هویت",
			IsPublic:    false,
		})
	}

	// تنظیمات OTP
	if val, exists := requestData["mobile_auth_enabled"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.mobile_auth_enabled",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "boolean",
			Description: "فعال‌سازی احراز هویت با موبایل",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["otp_length"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.otp.length",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "طول کد OTP",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["otp_expiry_minutes"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.otp.expiry_minutes",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "زمان انقضای کد OTP (دقیقه)",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["max_otp_attempts"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.max_otp_attempts",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "حداکثر تعداد تلاش برای OTP",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["otp_resend_cooldown"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.otp.resend_cooldown",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "زمان انتظار برای ارسال مجدد OTP (ثانیه)",
			IsPublic:    false,
		})
	}

	// کد پترن OTP - مهم‌ترین بخش
	if val, exists := requestData["otp_pattern_code"]; exists {
		settings = append(settings, models.Setting{
			Key:         "otp_pattern_code",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "string",
			Description: "کد پترن SMS برای ارسال کد تایید OTP",
			IsPublic:    false,
		})
	}

	// تنظیمات JWT
	if val, exists := requestData["jwt_expiry_hours"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.jwt_expiry_hours",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "زمان انقضای JWT (ساعت)",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["refresh_token_expiry_days"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.refresh_token_expiry_days",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "زمان انقضای توکن تازه‌سازی (روز)",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["jwt_secret"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.jwt_secret",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "string",
			Description: "کلید مخفی JWT",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["max_concurrent_sessions"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.max_concurrent_sessions",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "حداکثر تعداد نشست‌های همزمان",
			IsPublic:    false,
		})
	}

	// تنظیمات ورود با یوزرنیم
	if val, exists := requestData["username_auth_enabled"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.username_auth_enabled",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "boolean",
			Description: "فعال‌سازی ورود با نام کاربری",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["min_password_length"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.min_password_length",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "حداقل طول رمز عبور",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["max_login_attempts"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.max_login_attempts",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "حداکثر تعداد تلاش برای ورود",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["account_lockout_minutes"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.account_lockout_minutes",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "زمان قفل شدن حساب (دقیقه)",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["password_expiry_days"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.password_expiry_days",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "زمان انقضای رمز عبور (روز)",
			IsPublic:    false,
		})
	}

	// تنظیمات امنیتی
	if val, exists := requestData["two_factor_enabled"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.two_factor_enabled",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "boolean",
			Description: "فعال‌سازی احراز هویت دو مرحله‌ای",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["suspicious_activity_detection"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.suspicious_activity_detection",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "boolean",
			Description: "تشخیص فعالیت مشکوک",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["session_timeout_minutes"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.session_timeout_minutes",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "زمان انقضای نشست (دقیقه)",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["login_history_retention_days"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.login_history_retention_days",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "مدت نگهداری تاریخچه ورود (روز)",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["auto_block_failed_logins"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.auto_block_failed_logins",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "تعداد تلاش‌های ناموفق قبل از مسدودسازی خودکار",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["auto_block_duration_hours"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.auto_block_duration_hours",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "number",
			Description: "مدت زمان مسدودسازی خودکار (ساعت)",
			IsPublic:    false,
		})
	}

	// تنظیمات نقش‌ها
	if val, exists := requestData["default_user_role"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.default_user_role",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "string",
			Description: "نقش پیش‌فرض کاربران جدید",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["email_verification_enabled"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.email_verification_enabled",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "boolean",
			Description: "فعال‌سازی تایید ایمیل",
			IsPublic:    false,
		})
	}

	if val, exists := requestData["phone_verification_enabled"]; exists {
		settings = append(settings, models.Setting{
			Key:         "auth.phone_verification_enabled",
			Value:       fmt.Sprintf("%v", val),
			Category:    "auth",
			Type:        "boolean",
			Description: "فعال‌سازی تایید شماره تلفن",
			IsPublic:    false,
		})
	}

	// بروزرسانی تنظیمات در دیتابیس
	if err := h.settingService.UpdateSettings(c.Request.Context(), settings); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بروزرسانی تنظیمات احراز هویت", err.Error()))
		return
	}

	// پاک کردن کش تنظیمات
	if err := h.settingService.ClearCache(c.Request.Context()); err != nil {
		// خطا در پاک کردن کش نباید مانع پاسخ موفقیت‌آمیز شود
		fmt.Printf("هشدار: خطا در پاک کردن کش تنظیمات: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "تنظیمات احراز هویت با موفقیت بروزرسانی شد",
	})
}

// GetCompressionSettings دریافت تنظیمات فشرده‌سازی تصاویر از دیتابیس (دسته‌بندی compression)
func (h *SettingHandler) GetCompressionSettings(c *gin.Context) {
	settings, err := h.settingService.GetSettingsByCategory(c.Request.Context(), "compression")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات فشرده‌سازی", err.Error()))
		return
	}
	c.JSON(http.StatusOK, settings)
}

// GetVideoCompressionSettings دریافت تنظیمات فشرده‌سازی ویدیو
func (h *SettingHandler) GetVideoCompressionSettings(c *gin.Context) {
	// تنظیمات Mock برای فشرده‌سازی ویدیو
	settings := []map[string]interface{}{
		{"key": "video_compression.quality", "value": "85", "category": "video_compression", "type": "string"},
		{"key": "video_compression.format", "value": "mp4", "category": "video_compression", "type": "string"},
		{"key": "video_compression.codec", "value": "h264", "category": "video_compression", "type": "string"},
		{"key": "video_compression.two_pass", "value": "false", "category": "video_compression", "type": "boolean"},
		{"key": "video_compression.remove_metadata", "value": "true", "category": "video_compression", "type": "boolean"},
		{"key": "video_compression.create_thumbnail", "value": "true", "category": "video_compression", "type": "boolean"},
		{"key": "video_compression.frame_rate", "value": "30", "category": "video_compression", "type": "string"},
		{"key": "video_compression.custom_width", "value": "1920", "category": "video_compression", "type": "string"},
		{"key": "video_compression.custom_height", "value": "1080", "category": "video_compression", "type": "string"},
		{"key": "video_compression.custom_bitrate", "value": "2000", "category": "video_compression", "type": "string"},
		{"key": "video_compression.worker_count", "value": "4", "category": "video_compression", "type": "string"},
		{"key": "video_compression.frame_analysis_mode", "value": "auto", "category": "video_compression", "type": "string"},
		{"key": "video_compression.enabled", "value": "true", "category": "video_compression", "type": "boolean"},
		{"key": "video_compression.start_hour", "value": "0", "category": "video_compression", "type": "string"},
		{"key": "video_compression.end_hour", "value": "6", "category": "video_compression", "type": "string"},
	}

	c.JSON(http.StatusOK, settings)
}

// GetInvoicePrintSettings دریافت تنظیمات فاکتور و چاپ
// @Summary دریافت تنظیمات فاکتور و چاپ
// @Description دریافت تنظیمات مربوط به فاکتورها، قالب‌های چاپ و خروجی‌های سیستم
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/invoice-print-settings [get]
func (h *SettingHandler) GetInvoicePrintSettings(c *gin.Context) {
	// تنظیمات پیش‌فرض فاکتور و چاپ
	settings := map[string]interface{}{
		"invoice": map[string]interface{}{
			"companyName":    "شرکت من",
			"taxNumber":      "",
			"companyAddress": "",
			"companyPhone":   "",
		},
		"print": map[string]interface{}{
			"defaultPaperSize": "a4",
			"orientation":      "portrait",
			"defaultFont":      "iransans",
			"fontSize":         "12",
			"showLogo":         true,
			"showQRCode":       true,
			"showBarcode":      true,
		},
		"template": map[string]interface{}{
			"primaryColor":    "#3B82F6",
			"secondaryColor":  "#1E40AF",
			"showHeader":      true,
			"showFooter":      true,
			"showPageNumbers": true,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   settings,
	})
}

// UpdateInvoicePrintSettings بروزرسانی تنظیمات فاکتور و چاپ
// @Summary بروزرسانی تنظیمات فاکتور و چاپ
// @Description بروزرسانی تنظیمات مربوط به فاکتورها، قالب‌های چاپ و خروجی‌های سیستم
// @Tags تنظیمات
// @Accept json
// @Produce json
// @Param settings body map[string]interface{} true "تنظیمات فاکتور و چاپ"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/invoice-print-settings [put]
func (h *SettingHandler) UpdateInvoicePrintSettings(c *gin.Context) {
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error()))
		return
	}

	// تبدیل تنظیمات به فرمت مناسب برای ذخیره
	var settings []models.Setting

	// تنظیمات فاکتور
	if invoiceData, exists := requestData["invoice"].(map[string]interface{}); exists {
		for key, value := range invoiceData {
			settings = append(settings, models.Setting{
				Key:         fmt.Sprintf("invoice.%s", key),
				Value:       fmt.Sprintf("%v", value),
				Category:    "invoice",
				Type:        "string",
				Description: fmt.Sprintf("تنظیمات فاکتور - %s", key),
				IsPublic:    false,
			})
		}
	}

	// تنظیمات چاپ
	if printData, exists := requestData["print"].(map[string]interface{}); exists {
		for key, value := range printData {
			settings = append(settings, models.Setting{
				Key:         fmt.Sprintf("print.%s", key),
				Value:       fmt.Sprintf("%v", value),
				Category:    "print",
				Type:        "string",
				Description: fmt.Sprintf("تنظیمات چاپ - %s", key),
				IsPublic:    false,
			})
		}
	}

	// تنظیمات قالب
	if templateData, exists := requestData["template"].(map[string]interface{}); exists {
		for key, value := range templateData {
			settings = append(settings, models.Setting{
				Key:         fmt.Sprintf("template.%s", key),
				Value:       fmt.Sprintf("%v", value),
				Category:    "template",
				Type:        "string",
				Description: fmt.Sprintf("تنظیمات قالب - %s", key),
				IsPublic:    false,
			})
		}
	}

	// بروزرسانی تنظیمات در دیتابیس
	if err := h.settingService.UpdateSettings(c.Request.Context(), settings); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بروزرسانی تنظیمات فاکتور و چاپ", err.Error()))
		return
	}

	// پاک کردن کش تنظیمات
	if err := h.settingService.ClearCache(c.Request.Context()); err != nil {
		// خطا در پاک کردن کش نباید مانع پاسخ موفقیت‌آمیز شود
		fmt.Printf("هشدار: خطا در پاک کردن کش تنظیمات: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "تنظیمات فاکتور و چاپ با موفقیت بروزرسانی شد",
		"data":    requestData,
	})
}
