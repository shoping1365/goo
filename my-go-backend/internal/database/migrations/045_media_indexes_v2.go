package migrations

import "gorm.io/gorm"

// Up045MediaIndexesV2 تکمیل ایندکس‌های رسانه (اگر 022 کافی نباشد)
func Up045MediaIndexesV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_media_files_category_created ON media_files(category, created_at DESC) WHERE deleted_at IS NULL")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_media_files_uploaded_by_created ON media_files(uploaded_by, created_at DESC) WHERE deleted_at IS NULL")
		return nil
	})
}
