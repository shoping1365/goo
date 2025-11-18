package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"

	"gorm.io/gorm"
)

/*
APISettingsService
سرویس مدیریت تنظیمات API

این سرویس شامل عملیات زیر است:
- دریافت تنظیمات API
- بروزرسانی تنظیمات API
- تست اتصال به OpenAI
- مدیریت آمار استفاده
- رمزگذاری و رمزگشایی API Key ها
*/

type APISettingsService struct {
	db *gorm.DB
}

// NewAPISettingsService یک نمونه جدید از سرویس تنظیمات API ایجاد می‌کند
func NewAPISettingsService(db *gorm.DB) *APISettingsService {
	return &APISettingsService{
		db: db,
	}
}

// GetOpenAIConfig تنظیمات OpenAI را از دیتابیس دریافت می‌کند
func (s *APISettingsService) GetOpenAIConfig(ctx context.Context) (*models.OpenAIConfig, error) {
	var setting models.APISettings
	if err := s.db.WithContext(ctx).Where("provider = ?", "openai").First(&setting).Error; err != nil {
		return nil, fmt.Errorf("خطا در دریافت تنظیمات OpenAI از دیتابیس: %w", err)
	}

	var config models.OpenAIConfig
	if err := json.Unmarshal([]byte(setting.Config), &config); err != nil {
		return nil, fmt.Errorf("خطا در پردازش تنظیمات OpenAI: %w", err)
	}

	// رمزگشایی کلید اگر لازم بود
	if config.APIKey != "" && utils.IsEncrypted(config.APIKey) {
		decrypted, err := utils.DecryptString(config.APIKey)
		if err != nil {
			return nil, fmt.Errorf("خطا در رمزگشایی API Key: %w", err)
		}
		config.APIKey = decrypted
	}

	return &config, nil
}

// GetAPISettings تمام تنظیمات API را دریافت می‌کند
func (s *APISettingsService) GetAPISettings(ctx context.Context) (*models.APISettingsResponse, error) {
	fmt.Printf("🔍 شروع دریافت تنظیمات API از دیتابیس...\n")

	var apiSettings []models.APISettings

	if err := s.db.WithContext(ctx).Find(&apiSettings).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت تنظیمات API از دیتابیس: %v\n", err)
		return nil, fmt.Errorf("خطا در دریافت تنظیمات API: %w", err)
	}

	fmt.Printf("📄 تعداد تنظیمات یافت شده: %d\n", len(apiSettings))

	// تبدیل به فرمت مورد نیاز
	data := make(map[string]interface{})

	for _, setting := range apiSettings {
		switch setting.Provider {
		case "openai":
			var config models.OpenAIConfig
			var usageStats models.OpenAIUsageStats

			if err := json.Unmarshal([]byte(setting.Config), &config); err != nil {
				return nil, fmt.Errorf("خطا در پردازش تنظیمات OpenAI: %w", err)
			}

			// رمزگشایی API Key
			if config.APIKey != "" && utils.IsEncrypted(config.APIKey) {
				decryptedKey, err := utils.DecryptString(config.APIKey)
				if err != nil {
					return nil, fmt.Errorf("خطا در رمزگشایی API Key: %w", err)
				}
				config.APIKey = decryptedKey
			}

			if err := json.Unmarshal([]byte(setting.UsageStats), &usageStats); err != nil {
				// اگر آمار استفاده وجود نداشت، مقادیر پیش‌فرض تنظیم کن
				usageStats = models.OpenAIUsageStats{
					AccountBalance:    "0.00",
					MonthlyUsage:      "0.00",
					TotalRequests:     "0",
					TodayRequests:     "0",
					LastBalanceUpdate: "",
					LastUsageUpdate:   "",
				}
			}

			// ترکیب تنظیمات و آمار
			openAIData := map[string]interface{}{
				"enabled":             config.Enabled,
				"api_key":             config.APIKey,
				"api_url":             config.APIUrl,
				"default_model":       config.DefaultModel,
				"temperature":         config.Temperature,
				"rate_limit":          config.RateLimit,
				"timeout":             config.Timeout,
				"max_daily_cost":      config.MaxDailyCost,
				"consuming_pages":     config.ConsumingPages,
				"section_models":      config.SectionModels,
				"available_models":    config.AvailableModels,
				"account_balance":     usageStats.AccountBalance,
				"monthly_usage":       usageStats.MonthlyUsage,
				"total_requests":      usageStats.TotalRequests,
				"today_requests":      usageStats.TodayRequests,
				"last_balance_update": usageStats.LastBalanceUpdate,
			}

			data["openai"] = openAIData
		}
	}

	return &models.APISettingsResponse{
		Status: "success",
		Data:   data,
	}, nil
}

// UpdateAPISettings تنظیمات API را بروزرسانی می‌کند
func (s *APISettingsService) UpdateAPISettings(ctx context.Context, request *models.APISettingsRequest) (*models.APISettingsResponse, error) {
	// بروزرسانی تنظیمات OpenAI
	if request.OpenAI != nil {
		if err := s.updateOpenAISettings(ctx, request.OpenAI); err != nil {
			return nil, fmt.Errorf("خطا در بروزرسانی تنظیمات OpenAI: %w", err)
		}
	}

	// دریافت تنظیمات بروزرسانی شده
	return s.GetAPISettings(ctx)
}

// updateOpenAISettings تنظیمات OpenAI را بروزرسانی می‌کند
func (s *APISettingsService) updateOpenAISettings(ctx context.Context, config *models.OpenAIConfig) error {
	// رمزگذاری API Key اگر خالی نباشد و قبلاً رمزگذاری نشده باشد
	if config.APIKey != "" {
		// بررسی اینکه آیا قبلاً رمزگذاری شده یا نه
		if !utils.IsEncrypted(config.APIKey) {
			fmt.Printf("🔐 رمزگذاری API Key جدید...\n")
			encryptedKey, err := utils.EncryptString(config.APIKey)
			if err != nil {
				return fmt.Errorf("خطا در رمزگذاری API Key: %w", err)
			}
			config.APIKey = encryptedKey
			fmt.Printf("✅ API Key رمزگذاری شد\n")
		} else {
			fmt.Printf("🔐 API Key قبلاً رمزگذاری شده است\n")
		}
	} else {
		fmt.Printf("⚠️ API Key خالی است\n")
	}

	// تبدیل تنظیمات به JSON
	configJSON, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("خطا در تبدیل تنظیمات به JSON: %w", err)
	}

	// بررسی وجود تنظیمات OpenAI
	var existingSetting models.APISettings
	result := s.db.WithContext(ctx).Where("provider = ?", "openai").First(&existingSetting)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// ایجاد تنظیمات جدید
			newSetting := models.APISettings{
				Provider: "openai",
				Config:   string(configJSON),
				UsageStats: `{
					"account_balance": "0.00",
					"monthly_usage": "0.00",
					"total_requests": "0",
					"today_requests": "0",
					"last_balance_update": "",
					"last_usage_update": ""
				}`,
			}

			if err := s.db.WithContext(ctx).Create(&newSetting).Error; err != nil {
				return fmt.Errorf("خطا در ایجاد تنظیمات OpenAI: %w", err)
			}
			fmt.Printf("✅ تنظیمات OpenAI جدید ایجاد شد\n")
		} else {
			return fmt.Errorf("خطا در بررسی تنظیمات موجود: %w", result.Error)
		}
	} else {
		// بروزرسانی تنظیمات موجود
		existingSetting.Config = string(configJSON)
		existingSetting.UpdatedAt = time.Now()

		if err := s.db.WithContext(ctx).Save(&existingSetting).Error; err != nil {
			return fmt.Errorf("خطا در بروزرسانی تنظیمات OpenAI: %w", err)
		}
		fmt.Printf("✅ تنظیمات OpenAI بروزرسانی شد\n")
	}

	return nil
}

// TestOpenAIConnection اتصال به OpenAI را تست می‌کند
func (s *APISettingsService) TestOpenAIConnection(ctx context.Context, request *models.TestOpenAIRequest) (*models.TestOpenAIResponse, error) {
	// رمزگشایی API Key اگر رمزگذاری شده باشد
	apiKey := request.APIKey
	if utils.IsEncrypted(apiKey) {
		decryptedKey, err := utils.DecryptString(apiKey)
		if err != nil {
			return &models.TestOpenAIResponse{
				Status:  "error",
				Message: "خطا در رمزگشایی API Key: " + err.Error(),
			}, nil
		}
		apiKey = decryptedKey
	}
	// لاگ کلید نهایی
	maskedKey := ""
	if len(apiKey) > 8 {
		maskedKey = apiKey[:4] + "..." + apiKey[len(apiKey)-4:]
	} else {
		maskedKey = apiKey
	}
	fmt.Printf("🔑 کلید نهایی ارسالی به OpenAI (TestOpenAIConnection): %s\n", maskedKey)

	// ایجاد OpenAI client با SDK رسمی
	client := openai.NewClient(option.WithAPIKey(apiKey))

	fmt.Printf("🔍 تست اتصال به OpenAI...\n")
	fmt.Printf("  - URL: https://api.openai.com/v1\n")
	fmt.Printf("  - API Key: %s\n", maskAPIKey(apiKey))

	// تست اتصال با دریافت لیست مدل‌ها
	modelsResp, err := client.Models.List(ctx)
	if err != nil {
		fmt.Printf("❌ خطا در اتصال: %v\n", err)
		var apiErr *openai.Error
		if errors.As(err, &apiErr) {
			switch apiErr.StatusCode {
			case 401:
				return &models.TestOpenAIResponse{
					Status:  "error",
					Message: "API Key نامعتبر است. لطفاً کلید API صحیح را وارد کنید.",
				}, nil
			case 403:
				return &models.TestOpenAIResponse{
					Status:  "error",
					Message: "دسترسی به OpenAI محدود شده است. لطفاً تنظیمات حساب خود را بررسی کنید.",
				}, nil
			case 429:
				return &models.TestOpenAIResponse{
					Status:  "error",
					Message: "محدودیت تعداد درخواست. لطفاً کمی صبر کنید و دوباره تلاش کنید.",
				}, nil
			default:
				return &models.TestOpenAIResponse{
					Status:  "error",
					Message: fmt.Sprintf("خطا در اتصال به OpenAI: %s", apiErr.Error()),
				}, nil
			}
		}
		return &models.TestOpenAIResponse{
			Status:  "error",
			Message: "خطا در اتصال به OpenAI: " + err.Error(),
		}, nil
	}

	// اتصال موفق
	fmt.Printf("✅ اتصال موفق\n")
	fmt.Printf("📊 تعداد مدل‌های موجود: %d\n", len(modelsResp.Data))

	// نمایش مدل‌های موجود
	availableModels := []string{}
	for _, model := range modelsResp.Data {
		availableModels = append(availableModels, model.ID)
		fmt.Printf("  - %s\n", model.ID)
	}

	data := map[string]interface{}{
		"status":           "connected",
		"message":          "اتصال به OpenAI با موفقیت برقرار شد",
		"models_count":     len(modelsResp.Data),
		"available_models": availableModels,
	}

	return &models.TestOpenAIResponse{
		Status:  "success",
		Data:    data,
		Message: "اتصال به OpenAI با موفقیت برقرار شد",
	}, nil
}

func maskAPIKey(apiKey string) string {
	if len(apiKey) <= 8 {
		return "***"
	}
	return apiKey[:4] + "..." + apiKey[len(apiKey)-4:]
}

// UpdateUsageStats آمار استفاده را بروزرسانی می‌کند
func (s *APISettingsService) UpdateUsageStats(ctx context.Context, provider string, stats interface{}) error {
	// تبدیل آمار به JSON
	statsJSON, err := json.Marshal(stats)
	if err != nil {
		return fmt.Errorf("خطا در تبدیل آمار به JSON: %w", err)
	}

	// بروزرسانی آمار استفاده
	result := s.db.WithContext(ctx).Model(&models.APISettings{}).
		Where("provider = ?", provider).
		Update("usage_stats", string(statsJSON))

	if result.Error != nil {
		return fmt.Errorf("خطا در بروزرسانی آمار استفاده: %w", result.Error)
	}

	return nil
}

// FetchOpenAIUsageData آمار استفاده واقعی را از OpenAI دریافت می‌کند
func (s *APISettingsService) FetchOpenAIUsageData(ctx context.Context) (*models.OpenAIUsageStats, error) {
	fmt.Printf("🔍 شروع FetchOpenAIUsageData...\n")

	// دریافت تنظیمات OpenAI از دیتابیس
	var apiSettings models.APISettings
	if err := s.db.WithContext(ctx).Where("provider = ?", "openai").First(&apiSettings).Error; err != nil {
		fmt.Printf("❌ خطا در دریافت تنظیمات از دیتابیس: %v\n", err)
		// بررسی اینکه آیا جدول وجود دارد یا نه
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("تنظیمات OpenAI یافت نشد. لطفاً ابتدا تنظیمات API را ذخیره کنید")
		}
		// بررسی خطاهای دیگر
		if err.Error() == "relation \"api_settings\" does not exist" {
			return nil, fmt.Errorf("جدول تنظیمات API وجود ندارد. لطفاً ابتدا مایگریشن‌ها را اجرا کنید")
		}
		return nil, fmt.Errorf("خطا در دریافت تنظیمات OpenAI: %w", err)
	}

	fmt.Printf("✅ تنظیمات از دیتابیس دریافت شد. ID: %d, Provider: %s\n", apiSettings.ID, apiSettings.Provider)
	fmt.Printf("📄 Config length: %d, UsageStats length: %d\n", len(apiSettings.Config), len(apiSettings.UsageStats))

	// تبدیل تنظیمات
	var config models.OpenAIConfig
	fmt.Printf("🔄 شروع تبدیل JSON تنظیمات...\n")
	if err := json.Unmarshal([]byte(apiSettings.Config), &config); err != nil {
		fmt.Printf("❌ خطا در تبدیل JSON تنظیمات: %v\n", err)
		fmt.Printf("📄 Config content: %s\n", apiSettings.Config)
		return nil, fmt.Errorf("خطا در پردازش تنظیمات OpenAI: %w", err)
	}
	fmt.Printf("✅ تنظیمات JSON تبدیل شد\n")

	// بررسی فعال بودن
	fmt.Printf("🔍 بررسی فعال بودن OpenAI: %v\n", config.Enabled)
	if !config.Enabled {
		return nil, fmt.Errorf("OpenAI فعال نیست. لطفاً ابتدا OpenAI را فعال کنید")
	}

	// رمزگشایی API Key
	fmt.Printf("🔐 شروع رمزگشایی API Key...\n")
	apiKey := config.APIKey
	if apiKey != "" && utils.IsEncrypted(apiKey) {
		fmt.Printf("🔐 API Key رمزگذاری شده است، شروع رمزگشایی...\n")
		decryptedKey, err := utils.DecryptString(apiKey)
		if err != nil {
			fmt.Printf("❌ خطا در رمزگشایی API Key: %v\n", err)
			return nil, fmt.Errorf("خطا در رمزگشایی API Key: %w", err)
		}
		apiKey = decryptedKey
		fmt.Printf("✅ API Key رمزگشایی شد\n")
	} else {
		fmt.Printf("🔐 API Key رمزگذاری نشده است\n")
	}

	if apiKey == "" {
		fmt.Printf("❌ API Key خالی است\n")
		return nil, fmt.Errorf("API Key تنظیم نشده است. لطفاً ابتدا کلید API را وارد کنید")
	}

	fmt.Printf("🔑 API Key موجود است: %s\n", maskAPIKey(apiKey))

	// ایجاد OpenAI client با SDK
	client := openai.NewClient(option.WithAPIKey(apiKey))

	// پردازش داده‌ها
	stats := &models.OpenAIUsageStats{
		AccountBalance:    "0.00",
		MonthlyUsage:      "0.00",
		TotalRequests:     "0",
		TodayRequests:     "0",
		LastBalanceUpdate: time.Now().Format("2006-01-02 15:04:05"),
		LastUsageUpdate:   time.Now().Format("2006-01-02 15:04:05"),
	}

	fmt.Printf("🌐 شروع دریافت اطلاعات از OpenAI با SDK...\n")

	// تلاش برای دریافت اطلاعات حساب
	fmt.Printf("💰 تلاش برای دریافت اطلاعات حساب...\n")
	balanceData, balanceErr := s.fetchAccountBalanceWithSDK(ctx, client, config.APIUrl, apiKey)
	if balanceErr != nil {
		fmt.Printf("⚠️ خطا در دریافت اطلاعات حساب: %v\n", balanceErr)
		// ادامه کار با مقادیر پیش‌فرض
	} else if balanceData != nil {
		fmt.Printf("✅ اطلاعات حساب دریافت شد\n")
		// استخراج اعتبار حساب
		if hardLimit, ok := balanceData["hard_limit_usd"].(float64); ok {
			stats.AccountBalance = fmt.Sprintf("%.2f", hardLimit)
			fmt.Printf("💰 Hard limit: %.2f\n", hardLimit)
		} else if softLimit, ok := balanceData["soft_limit_usd"].(float64); ok {
			stats.AccountBalance = fmt.Sprintf("%.2f", softLimit)
			fmt.Printf("💰 Soft limit: %.2f\n", softLimit)
		}
	}

	// تلاش برای دریافت آمار استفاده
	fmt.Printf("📊 تلاش برای دریافت آمار استفاده...\n")
	usageData, usageErr := s.fetchUsageDataWithSDK(ctx, client, config.APIUrl, apiKey)
	if usageErr != nil {
		fmt.Printf("⚠️ خطا در دریافت آمار استفاده: %v\n", usageErr)
		// ادامه کار با مقادیر پیش‌فرض
	} else if usageData != nil {
		fmt.Printf("✅ آمار استفاده دریافت شد\n")
		// استخراج آمار استفاده
		if totalUsage, ok := usageData["total_usage"].(float64); ok {
			stats.MonthlyUsage = fmt.Sprintf("%.2f", totalUsage/100) // تبدیل از cent به دلار
			fmt.Printf("📊 Total usage: %.2f cents (%.2f USD)\n", totalUsage, totalUsage/100)
		}
	}

	// بروزرسانی آمار در دیتابیس
	fmt.Printf("💾 بروزرسانی آمار در دیتابیس...\n")
	if err := s.UpdateUsageStats(ctx, "openai", stats); err != nil {
		fmt.Printf("❌ خطا در بروزرسانی آمار: %v\n", err)
		return nil, fmt.Errorf("خطا در بروزرسانی آمار: %w", err)
	}

	fmt.Printf("✅ آمار استفاده بروزرسانی شد: %+v\n", stats)
	return stats, nil
}

// fetchAccountBalanceWithSDK اطلاعات حساب را از OpenAI با استفاده از SDK دریافت می‌کند
func (s *APISettingsService) fetchAccountBalanceWithSDK(ctx context.Context, openaiClient openai.Client, apiUrl, apiKey string) (map[string]interface{}, error) {
	fmt.Printf("🔍 fetchAccountBalanceWithSDK شروع شد...\n")

	// متأسفانه SDK رسمی OpenAI برای Go هنوز endpoint های billing را پشتیبانی نمی‌کند
	// بنابراین از HTTP request مستقیم استفاده می‌کنیم اما با error handling بهتر

	// ایجاد HTTP client
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// تلاش با endpoint اول - استفاده از endpoint صحیح OpenAI
	balanceReq, err := http.NewRequestWithContext(ctx, "GET", "https://api.openai.com/dashboard/billing/subscription", nil)
	if err != nil {
		fmt.Printf("❌ خطا در ایجاد درخواست اطلاعات حساب: %v\n", err)
		return nil, fmt.Errorf("خطا در ایجاد درخواست اطلاعات حساب: %w", err)
	}

	balanceReq.Header.Set("Authorization", "Bearer "+apiKey)
	balanceReq.Header.Set("Content-Type", "application/json")

	fmt.Printf("🔍 دریافت اطلاعات حساب OpenAI...\n")
	fmt.Printf("  - URL: %s\n", balanceReq.URL.String())
	fmt.Printf("  - API Key: %s\n", maskAPIKey(apiKey))

	balanceResp, err := client.Do(balanceReq)
	if err != nil {
		fmt.Printf("❌ خطا در اتصال به OpenAI: %v\n", err)
		return nil, fmt.Errorf("خطا در اتصال به OpenAI: %w", err)
	}
	defer balanceResp.Body.Close()

	fmt.Printf("📡 کد پاسخ اطلاعات حساب: %d\n", balanceResp.StatusCode)

	if balanceResp.StatusCode == 200 {
		var balanceData map[string]interface{}
		if err := json.NewDecoder(balanceResp.Body).Decode(&balanceData); err != nil {
			fmt.Printf("❌ خطا در پردازش پاسخ اطلاعات حساب: %v\n", err)
			return nil, fmt.Errorf("خطا در پردازش پاسخ اطلاعات حساب: %w", err)
		}
		fmt.Printf("✅ اطلاعات حساب دریافت شد\n")
		return balanceData, nil
	} else if balanceResp.StatusCode == 401 {
		fmt.Printf("❌ خطای 401: API Key نامعتبر است\n")
		return nil, fmt.Errorf("API Key نامعتبر است یا دسترسی به اطلاعات حساب ندارد")
	} else if balanceResp.StatusCode == 403 {
		fmt.Printf("❌ خطای 403: دسترسی محدود شده است\n")
		return nil, fmt.Errorf("دسترسی به اطلاعات حساب محدود شده است")
	} else {
		fmt.Printf("❌ خطای %d: تلاش با endpoint جایگزین\n", balanceResp.StatusCode)
		// تلاش با endpoint جایگزین
		return s.fetchAccountBalanceAlternative(ctx, client, apiUrl, apiKey)
	}
}

// fetchAccountBalanceAlternative تلاش جایگزین برای دریافت اطلاعات حساب
func (s *APISettingsService) fetchAccountBalanceAlternative(ctx context.Context, client *http.Client, apiUrl, apiKey string) (map[string]interface{}, error) {
	fmt.Printf("🔍 fetchAccountBalanceAlternative شروع شد...\n")

	// تلاش با endpoint جایگزین - استفاده از endpoint صحیح OpenAI
	balanceReq, err := http.NewRequestWithContext(ctx, "GET", "https://api.openai.com/dashboard/billing/credit_grants", nil)
	if err != nil {
		fmt.Printf("❌ خطا در ایجاد درخواست جایگزین اطلاعات حساب: %v\n", err)
		return nil, fmt.Errorf("خطا در ایجاد درخواست جایگزین اطلاعات حساب: %w", err)
	}

	balanceReq.Header.Set("Authorization", "Bearer "+apiKey)
	balanceReq.Header.Set("Content-Type", "application/json")

	fmt.Printf("🔍 تلاش جایگزین دریافت اطلاعات حساب...\n")
	fmt.Printf("  - URL: %s\n", balanceReq.URL.String())

	balanceResp, err := client.Do(balanceReq)
	if err != nil {
		fmt.Printf("❌ خطا در اتصال جایگزین به OpenAI: %v\n", err)
		return nil, fmt.Errorf("خطا در اتصال جایگزین به OpenAI: %w", err)
	}
	defer balanceResp.Body.Close()

	fmt.Printf("📡 کد پاسخ جایگزین اطلاعات حساب: %d\n", balanceResp.StatusCode)

	if balanceResp.StatusCode == 200 {
		var balanceData map[string]interface{}
		if err := json.NewDecoder(balanceResp.Body).Decode(&balanceData); err != nil {
			fmt.Printf("❌ خطا در پردازش پاسخ جایگزین اطلاعات حساب: %v\n", err)
			return nil, fmt.Errorf("خطا در پردازش پاسخ جایگزین اطلاعات حساب: %w", err)
		}
		fmt.Printf("✅ اطلاعات حساب با روش جایگزین دریافت شد\n")
		return balanceData, nil
	}

	fmt.Printf("❌ هیچ endpoint اطلاعات حساب قابل دسترسی نیست\n")
	return nil, fmt.Errorf("هیچ یک از endpoint های اطلاعات حساب قابل دسترسی نیست")
}

// fetchUsageDataWithSDK آمار استفاده را از OpenAI با استفاده از SDK دریافت می‌کند
func (s *APISettingsService) fetchUsageDataWithSDK(ctx context.Context, openaiClient openai.Client, apiUrl, apiKey string) (map[string]interface{}, error) {
	fmt.Printf("🔍 fetchUsageDataWithSDK شروع شد...\n")

	// متأسفانه SDK رسمی OpenAI برای Go هنوز endpoint های billing را پشتیبانی نمی‌کند
	// بنابراین از HTTP request مستقیم استفاده می‌کنیم اما با error handling بهتر

	// ایجاد HTTP client
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// استفاده از endpoint صحیح OpenAI
	usageReq, err := http.NewRequestWithContext(ctx, "GET", "https://api.openai.com/dashboard/billing/usage", nil)
	if err != nil {
		fmt.Printf("❌ خطا در ایجاد درخواست آمار استفاده: %v\n", err)
		return nil, fmt.Errorf("خطا در ایجاد درخواست آمار استفاده: %w", err)
	}

	usageReq.Header.Set("Authorization", "Bearer "+apiKey)
	usageReq.Header.Set("Content-Type", "application/json")

	fmt.Printf("🔍 دریافت آمار استفاده OpenAI...\n")
	fmt.Printf("  - URL: %s\n", usageReq.URL.String())

	usageResp, err := client.Do(usageReq)
	if err != nil {
		fmt.Printf("❌ خطا در اتصال به OpenAI: %v\n", err)
		return nil, fmt.Errorf("خطا در اتصال به OpenAI: %w", err)
	}
	defer usageResp.Body.Close()

	fmt.Printf("📡 کد پاسخ آمار استفاده: %d\n", usageResp.StatusCode)

	if usageResp.StatusCode == 200 {
		var usageData map[string]interface{}
		if err := json.NewDecoder(usageResp.Body).Decode(&usageData); err != nil {
			fmt.Printf("❌ خطا در پردازش پاسخ آمار استفاده: %v\n", err)
			return nil, fmt.Errorf("خطا در پردازش پاسخ آمار استفاده: %w", err)
		}
		fmt.Printf("✅ آمار استفاده دریافت شد\n")
		return usageData, nil
	} else if usageResp.StatusCode == 401 {
		fmt.Printf("❌ خطای 401: API Key نامعتبر است\n")
		return nil, fmt.Errorf("API Key نامعتبر است یا دسترسی به آمار استفاده ندارد")
	} else if usageResp.StatusCode == 403 {
		fmt.Printf("❌ خطای 403: دسترسی محدود شده است\n")
		return nil, fmt.Errorf("دسترسی به آمار استفاده محدود شده است")
	}

	fmt.Printf("❌ خطا در دریافت آمار استفاده: کد وضعیت %d\n", usageResp.StatusCode)
	return nil, fmt.Errorf("خطا در دریافت آمار استفاده: کد وضعیت %d", usageResp.StatusCode)
}
