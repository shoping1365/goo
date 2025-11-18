package main

import (
	"my-go-backend/internal/config"
	"my-go-backend/internal/database"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

func main() {
	// بارگذاری تنظیمات
	_ = config.Load()

	// اتصال به دیتابیس
	db, err := database.ConnectGorm()
	if err != nil {

		return
	}

	var tableExists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'api_settings')").Scan(&tableExists).Error
	if err != nil {
		// خطا در بررسی وجود جدول
		return
	}

	if !tableExists {
		// ❌ جدول api_settings وجود ندارد
		// 🔄 ایجاد جدول api_settings...

		if err := db.Exec(`CREATE TABLE IF NOT EXISTS api_settings (
            id BIGSERIAL PRIMARY KEY,
            provider VARCHAR(50) NOT NULL UNIQUE,
            config JSONB,
            usage_stats JSONB,
            created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
            updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
        );`).Error; err != nil {
			// خطا در ایجاد جدول api_settings
			return
		}
		// ✅ جدول api_settings ایجاد شد
	} else {
		// ✅ جدول api_settings وجود دارد
	}

	// بررسی وجود رکورد OpenAI
	// 🔍 بررسی وجود رکورد OpenAI...

	var apiSettings models.APISettings
	err = db.Where("provider = ?", "openai").First(&apiSettings).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// ❌ رکورد OpenAI یافت نشد
			// 🔄 ایجاد رکورد پیش‌فرض OpenAI...

			// ایجاد رکورد پیش‌فرض
			defaultConfig := `{
				"enabled": false,
				"api_key": "",
				"api_url": "https://api.openai.com/v1",
				"default_model": "gpt-3.5-turbo",
				"temperature": 0.7,
				"rate_limit": 60,
				"timeout": 30,
				"max_daily_cost": 10.00,
				"consuming_pages": [],
				"section_models": {},
				"available_models": []
			}`

			defaultUsageStats := `{
				"account_balance": "0.00",
				"monthly_usage": "0.00",
				"total_requests": "0",
				"today_requests": "0",
				"last_balance_update": "",
				"last_usage_update": ""
			}`

			newSetting := models.APISettings{
				Provider:   "openai",
				Config:     defaultConfig,
				UsageStats: defaultUsageStats,
			}

			if err := db.Create(&newSetting).Error; err != nil {
				// خطا در ایجاد رکورد OpenAI
				return
			}
			// ✅ رکورد پیش‌فرض OpenAI ایجاد شد
		} else {
			// خطا در بررسی رکورد OpenAI
			return
		}
	} else {

	}

}
