package migrations

import "gorm.io/gorm"

// Up067AttributesSoftDeleteFix افزودن ستون deleted_at برای جداول attributes و attribute_values (idempotent)
func Up067AttributesSoftDeleteFix(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`ALTER TABLE attributes ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE attribute_values ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ`).Error; err != nil {
			return err
		}
		return nil
	})
}


