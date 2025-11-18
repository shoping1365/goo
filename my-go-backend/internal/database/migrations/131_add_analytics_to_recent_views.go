package migrations

import (
	"gorm.io/gorm"
)

// Migration131AddAnalyticsToRecentViews adds analytics fields to recent_views table
func Migration131AddAnalyticsToRecentViews(db *gorm.DB) error {
	// افزودن فیلدهای تحلیلی به جدول recent_views (PostgreSQL syntax)

	// اضافه کردن ستون‌ها
	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS duration_seconds INT DEFAULT 0;
	`).Error; err != nil {
		return err
	}

	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS view_count INT DEFAULT 1;
	`).Error; err != nil {
		return err
	}

	if err := db.Exec(`
		ALTER TABLE recent_views 
		ADD COLUMN IF NOT EXISTS last_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
	`).Error; err != nil {
		return err
	}

	// اضافه کردن COMMENT ها (PostgreSQL syntax)
	if err := db.Exec(`
		COMMENT ON COLUMN recent_views.duration_seconds IS 'مدت زمان ماندن در صفحه (به ثانیه)';
	`).Error; err != nil {
		// اگر خطا داد مهم نیست، comment اختیاری است
	}

	if err := db.Exec(`
		COMMENT ON COLUMN recent_views.view_count IS 'تعداد دفعات بازدید';
	`).Error; err != nil {
		// اگر خطا داد مهم نیست
	}

	if err := db.Exec(`
		COMMENT ON COLUMN recent_views.last_updated_at IS 'آخرین بروزرسانی';
	`).Error; err != nil {
		// اگر خطا داد مهم نیست
	}

	// ایجاد trigger برای auto-update last_updated_at (PostgreSQL)
	// ابتدا function را بسازیم
	if err := db.Exec(`
		CREATE OR REPLACE FUNCTION update_recent_views_timestamp()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.last_updated_at = CURRENT_TIMESTAMP;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
	`).Error; err != nil {
		return err
	}

	// حالا trigger را بسازیم
	if err := db.Exec(`
		DROP TRIGGER IF EXISTS update_recent_views_timestamp_trigger ON recent_views;
	`).Error; err != nil {
		// اگر وجود نداشت مهم نیست
	}

	if err := db.Exec(`
		CREATE TRIGGER update_recent_views_timestamp_trigger
		BEFORE UPDATE ON recent_views
		FOR EACH ROW
		EXECUTE FUNCTION update_recent_views_timestamp();
	`).Error; err != nil {
		return err
	}

	return nil
}
