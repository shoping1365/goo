package migrations

import "gorm.io/gorm"

// Up088SeedAuthSettings اضافه کردن تنظیمات احراز هویت ضروری
func Up088SeedAuthSettings(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// JWT Secret - کلید مخفی برای امضای توکن‌ها
		if err := tx.Exec(`
			INSERT INTO settings (key, value, description, category, type, is_public, created_at, updated_at)
			VALUES ('jwt_secret', 'c0b9d2f18a4e4b6d9f3a1c7e5b2d8f4a', 'کلید مخفی JWT برای امضای توکن‌ها', 'auth', 'string', false, NOW(), NOW())
			ON CONFLICT (key) DO UPDATE SET 
				value = EXCLUDED.value,
				updated_at = NOW();
		`).Error; err != nil {
			return err
		}

		// تنظیمات OTP
		otpSettings := []struct {
			key, value, description string
		}{
			{"mobile_auth_enabled", "true", "فعال‌سازی احراز هویت با موبایل"},
			{"otp_length", "5", "طول کد تایید (5 رقم)"},
			{"otp_expiry_minutes", "5", "مدت اعتبار کد تایید (5 دقیقه)"},
			{"max_otp_attempts", "3", "حداکثر تلاش برای کد تایید"},
			{"otp_resend_delay_seconds", "60", "تاخیر ارسال مجدد کد (60 ثانیه)"},
		}

		for _, setting := range otpSettings {
			if err := tx.Exec(`
				INSERT INTO settings (key, value, description, category, type, is_public, created_at, updated_at)
				VALUES (?, ?, ?, 'auth', 'string', false, NOW(), NOW())
				ON CONFLICT (key) DO NOTHING;
			`, setting.key, setting.value, setting.description).Error; err != nil {
				return err
			}
		}

		// تنظیمات امنیتی
		securitySettings := []struct {
			key, value, description string
		}{
			{"session_timeout_hours", "24", "مدت اعتبار نشست (24 ساعت)"},
			{"max_login_attempts", "5", "حداکثر تلاش برای ورود"},
			{"lockout_duration_minutes", "15", "مدت قفل شدن حساب (15 دقیقه)"},
			{"password_min_length", "8", "حداقل طول رمز عبور"},
			{"require_strong_password", "true", "نیاز به رمز عبور قوی"},
		}

		for _, setting := range securitySettings {
			if err := tx.Exec(`
				INSERT INTO settings (key, value, description, category, type, is_public, created_at, updated_at)
				VALUES (?, ?, ?, 'security', 'string', false, NOW(), NOW())
				ON CONFLICT (key) DO NOTHING;
			`, setting.key, setting.value, setting.description).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
