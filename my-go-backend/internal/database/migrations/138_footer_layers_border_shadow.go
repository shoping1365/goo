package migrations

import "gorm.io/gorm"

// Up138FooterLayersBorderShadow اضافه کردن فیلد JSONB برای تنظیمات استایل
func Up138FooterLayersBorderShadow(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن یک فیلد JSONB برای تمام تنظیمات استایل
		if err := tx.Exec(`
			ALTER TABLE footer_layers 
			ADD COLUMN IF NOT EXISTS style_settings JSONB DEFAULT '{}'::jsonb
		`).Error; err != nil {
			return err
		}

		// ایجاد GIN index برای جستجوی سریع‌تر در JSONB
		if err := tx.Exec(`
			CREATE INDEX IF NOT EXISTS idx_footer_layers_style_settings 
			ON footer_layers USING GIN (style_settings)
		`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down138FooterLayersBorderShadow حذف فیلد style_settings از footer_layers
func Down138FooterLayersBorderShadow(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف index
		if err := tx.Exec(`
			DROP INDEX IF EXISTS idx_footer_layers_style_settings
		`).Error; err != nil {
			return err
		}

		// حذف ستون
		if err := tx.Exec(`
			ALTER TABLE footer_layers 
			DROP COLUMN IF EXISTS style_settings
		`).Error; err != nil {
			return err
		}

		return nil
	})
}
