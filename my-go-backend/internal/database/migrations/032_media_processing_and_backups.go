package migrations

import "gorm.io/gorm"

// Up032MediaProcessingAndBackups تکمیل جداول رسانه و بک‌آپ
func Up032MediaProcessingAndBackups(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_media_files_compressed ON media_files(compressed)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_backups_period ON backups(period)")
		return nil
	})
}
