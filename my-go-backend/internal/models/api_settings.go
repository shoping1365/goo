package models

import (
	"time"
)

/*
مدل APISettings برای مدیریت تنظیمات API های خارجی

این مدل شامل فیلدهای زیر است:
- ID: شناسه یکتای تنظیم
- Provider: نام ارائه دهنده API (مثل openai, google, etc.)
- Config: تنظیمات JSON مربوط به هر ارائه دهنده
- UsageStats: آمار استفاده از API
- CreatedAt: زمان ایجاد
- UpdatedAt: زمان آخرین بروزرسانی
*/

type APISettings struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Provider   string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"provider"`
	Config     string    `gorm:"type:jsonb" json:"config"`
	UsageStats string    `gorm:"type:jsonb" json:"usage_stats"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// OpenAIConfig تنظیمات OpenAI
type OpenAIConfig struct {
	Enabled         bool              `json:"enabled"`
	APIKey          string            `json:"api_key"`
	APIUrl          string            `json:"api_url"`
	DefaultModel    string            `json:"default_model"`
	Temperature     float64           `json:"temperature"`
	RateLimit       int               `json:"rate_limit"`
	Timeout         int               `json:"timeout"`
	MaxDailyCost    float64           `json:"max_daily_cost"`
	ConsumingPages  []string          `json:"consuming_pages"`
	SectionModels   map[string]string `json:"section_models"`   // مدل‌های مخصوص هر بخش
	AvailableModels []OpenAIModel     `json:"available_models"` // لیست تمام مدل‌های موجود
}

// OpenAIModel مدل OpenAI
type OpenAIModel struct {
	ID          string `json:"id"`          // شناسه مدل (مثل gpt-4)
	Name        string `json:"name"`        // نام نمایشی (مثل GPT-4)
	Description string `json:"description"` // توضیحات (مثل پیشرفته)
	Category    string `json:"category"`    // دسته‌بندی (مثل chat, text, image)
	MaxTokens   int    `json:"max_tokens"`  // حداکثر توکن
	IsActive    bool   `json:"is_active"`   // آیا فعال است
	CostPer1K   string `json:"cost_per_1k"` // هزینه هر 1000 توکن
}

// OpenAISectionConfig تنظیمات مدل برای هر بخش
type OpenAISectionConfig struct {
	Section     string  `json:"section"`     // نام بخش
	Model       string  `json:"model"`       // مدل انتخاب شده
	Temperature float64 `json:"temperature"` // دمای پاسخ‌دهی
	MaxTokens   int     `json:"max_tokens"`  // حداکثر توکن
	IsEnabled   bool    `json:"is_enabled"`  // آیا فعال است
}

// OpenAIUsageStats آمار استفاده از OpenAI
type OpenAIUsageStats struct {
	AccountBalance    string `json:"account_balance"`
	MonthlyUsage      string `json:"monthly_usage"`
	TotalRequests     string `json:"total_requests"`
	TodayRequests     string `json:"today_requests"`
	LastBalanceUpdate string `json:"last_balance_update"`
	LastUsageUpdate   string `json:"last_usage_update"`
}

// APISettingsRequest درخواست تنظیمات API
type APISettingsRequest struct {
	OpenAI *OpenAIConfig `json:"openai"`
}

// APISettingsResponse پاسخ تنظیمات API
type APISettingsResponse struct {
	Status  string                 `json:"status"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message,omitempty"`
}

// TestOpenAIRequest درخواست تست اتصال OpenAI
type TestOpenAIRequest struct {
	APIKey string `json:"api_key"`
	APIUrl string `json:"api_url"`
}

// TestOpenAIResponse پاسخ تست اتصال OpenAI
type TestOpenAIResponse struct {
	Status  string                 `json:"status"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Message string                 `json:"message,omitempty"`
}
