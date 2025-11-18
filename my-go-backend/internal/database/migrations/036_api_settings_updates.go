package migrations

import "gorm.io/gorm"

// Up036APISettingsUpdates بروزرسانی api_settings و ایندکس‌ها (اگر نیازی باشد)
func Up036APISettingsUpdates(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_api_settings_provider_unique ON api_settings(provider)")
		return nil
	})
}
