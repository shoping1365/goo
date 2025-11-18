package migrations

import "gorm.io/gorm"

// Up116MobileAppHeaderSimplify حذف لایه‌ها و ساده‌سازی ساختار هدر موبایل
func Up116MobileAppHeaderSimplify(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف جداول لایه‌ها
		if err := tx.Exec(`DROP TABLE IF EXISTS mobile_app_header_layers`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`DROP TABLE IF EXISTS mobile_app_footer_layers`).Error; err != nil {
			return err
		}

		// اضافه کردن فیلدهای جدید به جدول هدرها
		if err := tx.Exec(`ALTER TABLE mobile_app_headers ADD COLUMN IF NOT EXISTS header_type VARCHAR(50) DEFAULT 'default'`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_headers ADD COLUMN IF NOT EXISTS logo_url VARCHAR(500)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_headers ADD COLUMN IF NOT EXISTS logo_alt VARCHAR(255)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_headers ADD COLUMN IF NOT EXISTS show_search BOOLEAN DEFAULT true`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_headers ADD COLUMN IF NOT EXISTS show_cart BOOLEAN DEFAULT true`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_headers ADD COLUMN IF NOT EXISTS show_user_menu BOOLEAN DEFAULT true`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_headers ADD COLUMN IF NOT EXISTS show_notifications BOOLEAN DEFAULT true`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_headers ADD COLUMN IF NOT EXISTS show_menu_button BOOLEAN DEFAULT true`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_headers ADD COLUMN IF NOT EXISTS background_color VARCHAR(50) DEFAULT '#ffffff'`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_headers ADD COLUMN IF NOT EXISTS text_color VARCHAR(50) DEFAULT '#000000'`).Error; err != nil {
			return err
		}

		// اضافه کردن فیلدهای جدید به جدول فوترها
		if err := tx.Exec(`ALTER TABLE mobile_app_footers ADD COLUMN IF NOT EXISTS footer_type VARCHAR(50) DEFAULT 'default'`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_footers ADD COLUMN IF NOT EXISTS background_color VARCHAR(50) DEFAULT '#1f2937'`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_footers ADD COLUMN IF NOT EXISTS text_color VARCHAR(50) DEFAULT '#ffffff'`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_footers ADD COLUMN IF NOT EXISTS show_social_links BOOLEAN DEFAULT true`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE mobile_app_footers ADD COLUMN IF NOT EXISTS show_contact_info BOOLEAN DEFAULT true`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down116MobileAppHeaderSimplify بازگردانی تغییرات
func Down116MobileAppHeaderSimplify(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف فیلدهای اضافه شده
		fields := []string{
			"header_type", "logo_url", "logo_alt", "show_search", "show_cart",
			"show_user_menu", "show_notifications", "show_menu_button", "background_color", "text_color",
		}

		for _, field := range fields {
			_ = tx.Exec(`ALTER TABLE mobile_app_headers DROP COLUMN IF EXISTS ` + field)
		}

		footerFields := []string{
			"footer_type", "background_color", "text_color", "show_social_links", "show_contact_info",
		}

		for _, field := range footerFields {
			_ = tx.Exec(`ALTER TABLE mobile_app_footers DROP COLUMN IF EXISTS ` + field)
		}

		// بازسازی جداول لایه‌ها (اختیاری)
		// در صورت نیاز می‌توان جداول لایه‌ها را دوباره ایجاد کرد

		return nil
	})
}
