package migrations

import (
	"gorm.io/gorm"
)

// Migration133FixRecentViewsDefaults fixes default values and constraints for recent_views
func Migration133FixRecentViewsDefaults(db *gorm.DB) error {
	// حذف UNIQUE constraint (اگر وجود دارد)
	// چون می‌خواهیم هر بازدید را جداگانه ثبت کنیم
	db.Exec(`ALTER TABLE recent_views DROP CONSTRAINT IF EXISTS unique_user_product_recent_view;`)

	// اضافه کردن DEFAULT برای ستون‌های موجود (اگر ندارند)
	db.Exec(`ALTER TABLE recent_views ALTER COLUMN browser SET DEFAULT 'Unknown';`)
	db.Exec(`ALTER TABLE recent_views ALTER COLUMN browser_version SET DEFAULT '';`)
	db.Exec(`ALTER TABLE recent_views ALTER COLUMN os SET DEFAULT 'Unknown';`)
	db.Exec(`ALTER TABLE recent_views ALTER COLUMN os_version SET DEFAULT '';`)
	db.Exec(`ALTER TABLE recent_views ALTER COLUMN device_type SET DEFAULT 'Unknown';`)
	db.Exec(`ALTER TABLE recent_views ALTER COLUMN device_model SET DEFAULT 'Unknown Device';`)
	db.Exec(`ALTER TABLE recent_views ALTER COLUMN ip_address SET DEFAULT '';`)
	db.Exec(`ALTER TABLE recent_views ALTER COLUMN user_agent SET DEFAULT '';`)

	// آپدیت کردن رکوردهای موجود که NULL یا خالی هستند
	if err := db.Exec(`
		UPDATE recent_views 
		SET 
			browser = COALESCE(NULLIF(browser, ''), 'Unknown'),
			browser_version = COALESCE(browser_version, ''),
			os = COALESCE(NULLIF(os, ''), 'Unknown'),
			os_version = COALESCE(os_version, ''),
			device_type = COALESCE(NULLIF(device_type, ''), 'Unknown'),
			device_model = COALESCE(NULLIF(device_model, ''), 'Unknown Device'),
			ip_address = COALESCE(ip_address, ''),
			user_agent = COALESCE(user_agent, '')
		WHERE browser IS NULL OR browser = '' 
		   OR os IS NULL OR os = ''
		   OR device_type IS NULL OR device_type = '';
	`).Error; err != nil {
		return err
	}

	return nil
}
