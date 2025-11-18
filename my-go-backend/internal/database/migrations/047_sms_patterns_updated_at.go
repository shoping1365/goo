package migrations

import "gorm.io/gorm"

// Up047SMSPatternsUpdatedAt افزودن ستون updated_at به sms_patterns
func Up047SMSPatternsUpdatedAt(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("ALTER TABLE sms_patterns ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ DEFAULT NOW()")
		_ = tx.Exec("UPDATE sms_patterns SET updated_at = NOW() WHERE updated_at IS NULL")
		_ = tx.Exec("ALTER TABLE sms_patterns ALTER COLUMN updated_at SET NOT NULL")
		return nil
	})
}
