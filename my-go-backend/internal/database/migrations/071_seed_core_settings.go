package migrations

import "gorm.io/gorm"

// Up071SeedCoreSettings درج کلیدهای ضروری settings به صورت ایدمپوتنت
func Up071SeedCoreSettings(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// bots.malicious.user_agents مقدار پیش‌فرض خالی (CSV)
		if err := tx.Exec(`
            INSERT INTO settings (key, value, description, category, type, is_public)
            VALUES ('bots.malicious.user_agents', '', 'لیست یو‌زر ایجنت‌های بدخواه (CSV)', 'bots', 'string', false)
            ON CONFLICT (key) DO NOTHING;
        `).Error; err != nil {
			return err
		}

		return nil
	})
}


