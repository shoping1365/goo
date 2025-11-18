package migrations

import "gorm.io/gorm"

// Up060ProductQuestionsDeviceV2 تایید ستون‌های دستگاه در سوالات محصول
func Up060ProductQuestionsDeviceV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("ALTER TABLE product_questions ADD COLUMN IF NOT EXISTS user_agent TEXT")
		_ = tx.Exec("ALTER TABLE product_questions ADD COLUMN IF NOT EXISTS device_info JSONB")
		return nil
	})
}
