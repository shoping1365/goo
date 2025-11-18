package migrations

import "gorm.io/gorm"

// Up091HeaderLayersPaddingFields اضافه کردن فیلدهای padding به جدول header_layers
// این مایگریشن فیلدهای padding را برای کنترل فاصله داخلی لایه‌های هدر اضافه می‌کند
func Up091HeaderLayersPaddingFields(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن فیلدهای padding به جدول header_layers
		// این فیلدها برای کنترل فاصله داخلی لایه‌های هدر استفاده می‌شوند

		// padding_left - فاصله داخلی از چپ
		if err := tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS padding_left INTEGER DEFAULT 20`).Error; err != nil {
			return err
		}

		// padding_right - فاصله داخلی از راست
		if err := tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS padding_right INTEGER DEFAULT 20`).Error; err != nil {
			return err
		}

		// padding_top - فاصله داخلی از بالا
		if err := tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS padding_top INTEGER DEFAULT 10`).Error; err != nil {
			return err
		}

		// padding_bottom - فاصله داخلی از پایین
		if err := tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS padding_bottom INTEGER DEFAULT 10`).Error; err != nil {
			return err
		}

		// اضافه کردن index برای بهبود عملکرد queries مربوط به padding
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_header_layers_padding ON header_layers(padding_left, padding_right, padding_top, padding_bottom)`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down091HeaderLayersPaddingFields حذف فیلدهای padding از جدول header_layers
func Down091HeaderLayersPaddingFields(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف index
		_ = tx.Exec(`DROP INDEX IF EXISTS idx_header_layers_padding`)

		// حذف فیلدهای padding
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS padding_bottom`)
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS padding_top`)
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS padding_right`)
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS padding_left`)

		return nil
	})
}
