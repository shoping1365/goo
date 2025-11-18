package migrations

import "gorm.io/gorm"

// Up140HeaderLayersCompleteFields اضافه کردن فیلدهای مستقیم border, shadow, direction و responsive به header_layers
func Up140HeaderLayersCompleteFields(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن فیلدهای Border
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS enable_border BOOLEAN DEFAULT FALSE`)
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS border_position VARCHAR(50) DEFAULT 'all'`)
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS border_color VARCHAR(50) DEFAULT '#e5e7eb'`)
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS border_width INTEGER DEFAULT 1`)
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS border_style VARCHAR(50) DEFAULT 'solid'`)

		// اضافه کردن فیلدهای Shadow
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS enable_shadow BOOLEAN DEFAULT FALSE`)
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS shadow_intensity VARCHAR(50) DEFAULT 'md'`)
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS shadow_direction VARCHAR(50) DEFAULT 'top'`)

		// اضافه کردن فیلد Direction
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS direction VARCHAR(10) DEFAULT 'rtl'`)

		// تغییر نام show_on_mobile به mobile_responsive (اگر وجود داشت)
		_ = tx.Exec(`
			DO $$ 
			BEGIN 
				IF EXISTS(
					SELECT 1 FROM information_schema.columns 
					WHERE table_name='header_layers' AND column_name='show_on_mobile'
				) THEN
					ALTER TABLE header_layers RENAME COLUMN show_on_mobile TO mobile_responsive;
				ELSE
					ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS mobile_responsive BOOLEAN DEFAULT TRUE;
				END IF;
			END $$;
		`)

		// اضافه کردن فیلد tablet_responsive
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS tablet_responsive BOOLEAN DEFAULT TRUE`)

		// حذف show_on_desktop اگر وجود دارد (جایگزین با mobile/tablet responsive شد)
		_ = tx.Exec(`
			DO $$ 
			BEGIN 
				IF EXISTS(
					SELECT 1 FROM information_schema.columns 
					WHERE table_name='header_layers' AND column_name='show_on_desktop'
				) THEN
					ALTER TABLE header_layers DROP COLUMN show_on_desktop;
				END IF;
			END $$;
		`)

		// ایجاد ایندکس برای فیلدهای جدید
		_ = tx.Exec(`CREATE INDEX IF NOT EXISTS idx_header_layers_border ON header_layers(enable_border)`)
		_ = tx.Exec(`CREATE INDEX IF NOT EXISTS idx_header_layers_shadow ON header_layers(enable_shadow)`)
		_ = tx.Exec(`CREATE INDEX IF NOT EXISTS idx_header_layers_responsive ON header_layers(mobile_responsive, tablet_responsive)`)

		return nil
	})
}

// Down140HeaderLayersCompleteFields حذف فیلدهای اضافه شده
func Down140HeaderLayersCompleteFields(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف ایندکس‌ها
		_ = tx.Exec(`DROP INDEX IF EXISTS idx_header_layers_responsive`)
		_ = tx.Exec(`DROP INDEX IF EXISTS idx_header_layers_shadow`)
		_ = tx.Exec(`DROP INDEX IF EXISTS idx_header_layers_border`)

		// بازگرداندن show_on_desktop
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS show_on_desktop BOOLEAN DEFAULT TRUE`)

		// تغییر نام mobile_responsive به show_on_mobile
		_ = tx.Exec(`
			DO $$ 
			BEGIN 
				IF EXISTS(
					SELECT 1 FROM information_schema.columns 
					WHERE table_name='header_layers' AND column_name='mobile_responsive'
				) THEN
					ALTER TABLE header_layers RENAME COLUMN mobile_responsive TO show_on_mobile;
				END IF;
			END $$;
		`)

		// حذف فیلد tablet_responsive
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS tablet_responsive`)

		// حذف فیلد direction
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS direction`)

		// حذف فیلدهای Shadow
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS shadow_direction`)
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS shadow_intensity`)
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS enable_shadow`)

		// حذف فیلدهای Border
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS border_style`)
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS border_width`)
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS border_color`)
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS border_position`)
		_ = tx.Exec(`ALTER TABLE header_layers DROP COLUMN IF EXISTS enable_border`)

		return nil
	})
}
