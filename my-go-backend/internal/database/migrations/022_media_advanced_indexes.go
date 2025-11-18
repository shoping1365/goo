package migrations

import "gorm.io/gorm"

// Up022MediaAdvancedIndexes ایندکس‌های تکمیلی رسانه (معادل 45)
func Up022MediaAdvancedIndexes(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_media_versions_active_created ON media_versions(media_id, is_active, created_at DESC) WHERE deleted_at IS NULL")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_media_versions_format_quality ON media_versions(media_id, format, quality) WHERE deleted_at IS NULL")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_compression_jobs_media_status_created ON compression_jobs(media_id, status, created_at DESC)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_compression_jobs_status_created ON compression_jobs(status, created_at DESC)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_compression_settings_scope_key ON compression_settings(scope, key)")
		return nil
	})
}
