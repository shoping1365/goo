package migrations

import (
	"gorm.io/gorm"
)

// Migration132AddDeviceInfoToRecentViews adds device and browser information to recent_views
func Migration132AddDeviceInfoToRecentViews(db *gorm.DB) error {
	// اضافه کردن فیلدهای مربوط به دستگاه و مرورگر (PostgreSQL)

	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS browser VARCHAR(100) DEFAULT 'Unknown';
	`).Error; err != nil {
		return err
	}

	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS browser_version VARCHAR(50) DEFAULT '';
	`).Error; err != nil {
		return err
	}

	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS os VARCHAR(100) DEFAULT 'Unknown';
	`).Error; err != nil {
		return err
	}

	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS os_version VARCHAR(50) DEFAULT '';
	`).Error; err != nil {
		return err
	}

	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS device_type VARCHAR(50) DEFAULT 'Unknown';
	`).Error; err != nil {
		return err
	}

	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS device_model VARCHAR(100) DEFAULT 'Unknown Device';
	`).Error; err != nil {
		return err
	}

	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS ip_address VARCHAR(45) DEFAULT '';
	`).Error; err != nil {
		return err
	}

	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS user_agent TEXT DEFAULT '';
	`).Error; err != nil {
		return err
	}

	// حذف UNIQUE constraint قدیمی (چون الان می‌خواهیم چند بازدید داشته باشیم)
	db.Exec(`ALTER TABLE recent_views DROP CONSTRAINT IF EXISTS unique_user_product_recent_view;`)

	// آپدیت کردن رکوردهای قدیمی که NULL هستند
	db.Exec(`
		UPDATE recent_views 
		SET 
			browser = COALESCE(browser, 'Unknown'),
			browser_version = COALESCE(browser_version, ''),
			os = COALESCE(os, 'Unknown'),
			os_version = COALESCE(os_version, ''),
			device_type = COALESCE(device_type, 'Unknown'),
			device_model = COALESCE(device_model, 'Unknown Device'),
			ip_address = COALESCE(ip_address, ''),
			user_agent = COALESCE(user_agent, '')
		WHERE browser IS NULL OR browser = '';
	`)

	// اضافه کردن COMMENT ها
	db.Exec(`COMMENT ON COLUMN recent_views.browser IS 'نام مرورگر';`)
	db.Exec(`COMMENT ON COLUMN recent_views.browser_version IS 'نسخه مرورگر';`)
	db.Exec(`COMMENT ON COLUMN recent_views.os IS 'سیستم عامل';`)
	db.Exec(`COMMENT ON COLUMN recent_views.os_version IS 'نسخه سیستم عامل';`)
	db.Exec(`COMMENT ON COLUMN recent_views.device_type IS 'نوع دستگاه (Mobile/Desktop/Tablet)';`)
	db.Exec(`COMMENT ON COLUMN recent_views.device_model IS 'مدل دستگاه';`)
	db.Exec(`COMMENT ON COLUMN recent_views.ip_address IS 'آدرس IP کاربر';`)
	db.Exec(`COMMENT ON COLUMN recent_views.user_agent IS 'User-Agent کامل';`)

	// ایجاد index برای جستجوهای سریع‌تر
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_recent_views_device_type ON recent_views(device_type);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_recent_views_os ON recent_views(os);`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_recent_views_browser ON recent_views(browser);`)

	return nil
}
