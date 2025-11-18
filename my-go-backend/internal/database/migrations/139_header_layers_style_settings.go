package migrations

import "gorm.io/gorm"

// Up139HeaderLayersStyleSettings اضافه کردن فیلد JSONB برای تنظیمات استایل هدر
func Up139HeaderLayersStyleSettings(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن فیلد style_settings از نوع JSONB برای ذخیره تنظیمات مرز و سایه
		if err := tx.Exec(`
			ALTER TABLE header_layers 
			ADD COLUMN IF NOT EXISTS style_settings JSONB DEFAULT '{}'::jsonb
		`).Error; err != nil {
			return err
		}

		// ایجاد ایندکس GIN روی style_settings برای جستجوی سریع‌تر
		if err := tx.Exec(`
			CREATE INDEX IF NOT EXISTS idx_header_layers_style_settings 
			ON header_layers USING GIN (style_settings)
		`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down139HeaderLayersStyleSettings حذف فیلد style_settings از header_layers
func Down139HeaderLayersStyleSettings(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف ایندکس مربوط به style_settings
		if err := tx.Exec(`
			DROP INDEX IF EXISTS idx_header_layers_style_settings
		`).Error; err != nil {
			return err
		}

		// حذف ستون style_settings
		if err := tx.Exec(`
			ALTER TABLE header_layers 
			DROP COLUMN IF EXISTS style_settings
		`).Error; err != nil {
			return err
		}

		return nil
	})
}
