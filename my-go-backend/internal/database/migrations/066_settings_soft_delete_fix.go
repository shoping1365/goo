package migrations

import "gorm.io/gorm"

// Up066SettingsSoftDeleteFix افزودن ستون deleted_at به جدول settings (idempotent)
func Up066SettingsSoftDeleteFix(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`ALTER TABLE settings ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ`).Error; err != nil {
			return err
		}
		return nil
	})
}


