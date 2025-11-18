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
APISettingsHandler
کنترلر تنظیمات API

این کنترلر برای مدیریت APIهای مربوط به تنظیمات API های خارجی استفاده می‌شود و شامل موارد زیر است:
- دریافت تنظیمات API
- بروزرسانی تنظیمات API
- تست اتصال به OpenAI
*/

type APISettingsHandler struct {
	apiSettingsService *services.APISettingsService
}

// NewAPISettingsHandler یک نمونه جدید از کنترلر تنظیمات API ایجاد می‌کند
func NewAPISettingsHandler(apiSettingsService *services.APISettingsService) *APISettingsHandler {
	return &APISettingsHandler{
		apiSettingsService: apiSettingsService,
	}
}

// GetAPISettings تمام تنظیمات API را دریافت می‌کند
// @Summary دریافت تنظیمات API
// @Description دریافت تمام تنظیمات API های خارجی
// @Tags تنظیمات API
// @Accept json
// @Produce json
// @Success 200 {object} models.APISettingsResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/api-settings [get]
func (h *APISettingsHandler) GetAPISettings(c *gin.Context) {
	settings, err := h.apiSettingsService.GetAPISettings(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات API", err.Error()))
		return
	}

	c.JSON(http.StatusOK, settings)
}

// UpdateAPISettings تنظیمات API را بروزرسانی می‌کند
// @Summary بروزرسانی تنظیمات API
// @Description بروزرسانی تنظیمات API های خارجی
// @Tags تنظیمات API
// @Accept json
// @Produce json
// @Param settings body models.APISettingsRequest true "تنظیمات API"
// @Success 200 {object} models.APISettingsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/api-settings [put]
func (h *APISettingsHandler) UpdateAPISettings(c *gin.Context) {
	var request models.APISettingsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Printf("❌ خطا در تبدیل JSON: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "اطلاعات نامعتبر", err.Error()))
		return
	}

	// اعتبارسنجی تنظیمات OpenAI
	if request.OpenAI != nil {
		if err := h.validateOpenAIConfig(request.OpenAI); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "تنظیمات OpenAI نامعتبر", err.Error()))
			return
		}
	}

	settings, err := h.apiSettingsService.UpdateAPISettings(c.Request.Context(), &request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بروزرسانی تنظیمات API", err.Error()))
		return
	}
	c.JSON(http.StatusOK, settings)
}

// TestOpenAIConnection اتصال به OpenAI را تست می‌کند
// @Summary تست اتصال OpenAI
// @Description تست اتصال به OpenAI با استفاده از کلید API
// @Tags تنظیمات API
// @Accept json
// @Produce json
// @Param request body models.TestOpenAIRequest true "اطلاعات تست"
// @Success 200 {object} models.TestOpenAIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/api-settings/test-openai [post]
func (h *APISettingsHandler) TestOpenAIConnection(c *gin.Context) {
	fmt.Printf("🧪 شروع تست اتصال OpenAI...\n")

	// دریافت تنظیمات OpenAI از دیتابیس
	config, err := h.apiSettingsService.GetOpenAIConfig(c.Request.Context())
	if err != nil {
		fmt.Printf("❌ خطا در دریافت تنظیمات OpenAI: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت تنظیمات OpenAI", err.Error()))
		return
	}

	if !config.Enabled || config.APIKey == "" {
		fmt.Printf("❌ OpenAI فعال نیست یا API Key تنظیم نشده است\n")
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "OpenAI فعال نیست یا API Key تنظیم نشده است", nil))
		return
	}

	// ایجاد درخواست تست با تنظیمات موجود
	request := &models.TestOpenAIRequest{
		APIKey: config.APIKey,
		APIUrl: config.APIUrl,
	}

	if request.APIUrl == "" {
		request.APIUrl = "https://api.openai.com/v1"
	}

	fmt.Printf("🔑 تست اتصال با کلید: %s\n", maskAPIKey(request.APIKey))
	fmt.Printf("🌐 تست اتصال با URL: %s\n", request.APIUrl)

	response, err := h.apiSettingsService.TestOpenAIConnection(c.Request.Context(), request)
	if err != nil {
		fmt.Printf("❌ خطا در تست اتصال: %v\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("API_ERROR", "خطا در تست اتصال", err.Error()))
		return
	}

	fmt.Printf("✅ تست اتصال موفق\n")
	c.JSON(http.StatusOK, response)
}

// FetchOpenAIUsageData آمار استفاده OpenAI را دریافت می‌کند
// @Summary دریافت آمار استفاده OpenAI
// @Description دریافت آمار استفاده واقعی از OpenAI API
// @Tags تنظیمات API
// @Accept json
// @Produce json
// @Success 200 {object} models.OpenAIUsageStats
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/api-settings/fetch-usage [post]
func (h *APISettingsHandler) FetchOpenAIUsageData(c *gin.Context) {
	stats, err := h.apiSettingsService.FetchOpenAIUsageData(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("API_ERROR", "خطا در دریافت آمار استفاده", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    stats,
		"message": "آمار استفاده با موفقیت دریافت شد",
	})
}

// validateOpenAIConfig تنظیمات OpenAI را اعتبارسنجی می‌کند
func (h *APISettingsHandler) validateOpenAIConfig(config *models.OpenAIConfig) error {
	fmt.Printf("🔍 اعتبارسنجی با مقادیر:\n")
	fmt.Printf("  - API Key: %s\n", maskAPIKey(config.APIKey))
	fmt.Printf("  - API URL: %s\n", config.APIUrl)
	fmt.Printf("  - Default Model: %s\n", config.DefaultModel)
	fmt.Printf("  - Temperature: %f\n", config.Temperature)
	fmt.Printf("  - Rate Limit: %d\n", config.RateLimit)
	fmt.Printf("  - Timeout: %d\n", config.Timeout)
	fmt.Printf("  - Max Daily Cost: %f\n", config.MaxDailyCost)
	fmt.Printf("  - Consuming Pages: %v\n", config.ConsumingPages)
	fmt.Printf("  - Section Models: %v\n", config.SectionModels)
	fmt.Printf("  - Available Models Count: %d\n", len(config.AvailableModels))

	// بررسی کلید API
	if config.APIKey == "" {
		return fmt.Errorf("کلید API الزامی است")
	}

	// تنظیم URL پیش‌فرض
	if config.APIUrl == "" {
		config.APIUrl = "https://api.openai.com/v1"
	}

	// تنظیم مدل پیش‌فرض
	if config.DefaultModel == "" {
		config.DefaultModel = "gpt-4.1-nano-2025-04-14"
	}

	// تنظیم مقادیر پیش‌فرض برای فیلدهای عددی
	if config.Temperature == 0 {
		config.Temperature = 0.7
	}
	if config.RateLimit == 0 {
		config.RateLimit = 60
	}
	if config.Timeout == 0 {
		config.Timeout = 30
	}
	if config.MaxDailyCost == 0 {
		config.MaxDailyCost = 10.0
	}

	// اعتبارسنجی temperature
	if config.Temperature < 0 || config.Temperature > 2 {
		return fmt.Errorf("دمای پاسخ‌دهی باید بین 0 تا 2 باشد (مقدار فعلی: %f)", config.Temperature)
	}

	// اعتبارسنجی rate limit
	if config.RateLimit <= 0 {
		return fmt.Errorf("محدودیت درخواست باید بزرگتر از صفر باشد (مقدار فعلی: %d)", config.RateLimit)
	}

	// اعتبارسنجی timeout
	if config.Timeout < 5 || config.Timeout > 120 {
		return fmt.Errorf("تایم‌اوت باید بین 5 تا 120 ثانیه باشد (مقدار فعلی: %d)", config.Timeout)
	}

	// اعتبارسنجی max daily cost
	if config.MaxDailyCost < 0 {
		return fmt.Errorf("حداکثر هزینه روزانه نمی‌تواند منفی باشد (مقدار فعلی: %f)", config.MaxDailyCost)
	}

	// اعتبارسنجی consuming pages
	if config.ConsumingPages == nil {
		config.ConsumingPages = []string{}
	}

	// اعتبارسنجی section models
	if config.SectionModels == nil {
		config.SectionModels = make(map[string]string)
	}

	// اعتبارسنجی available models
	if config.AvailableModels == nil {
		config.AvailableModels = []models.OpenAIModel{}
	}

	return nil
}

func maskAPIKey(apiKey string) string {
	if len(apiKey) <= 8 {
		return "***"
	}
	return apiKey[:4] + "..." + apiKey[len(apiKey)-4:]
}
