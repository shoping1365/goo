package migrations

import "gorm.io/gorm"

// Up102FooterCore ایجاد جداول فوتر و لایه‌های فوتر
// این مایگریشن ساختار کامل فوتر را با پشتیبانی از لایه‌بندی پیشرفته ایجاد می‌کند
func Up102FooterCore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// ایجاد جدول footers
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS footers (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL UNIQUE,
				description TEXT,
				page_selection VARCHAR(50) DEFAULT 'all',
				specific_pages TEXT,
				excluded_pages TEXT,
				is_active BOOLEAN DEFAULT true,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
			)
		`).Error; err != nil {
			return err
		}

		// ایجاد جدول footer_layers
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS footer_layers (
				id SERIAL PRIMARY KEY,
				footer_id INTEGER NOT NULL REFERENCES footers(id) ON DELETE CASCADE,
				name VARCHAR(255) NOT NULL,
				width INTEGER DEFAULT 100,
				height INTEGER DEFAULT 200,
				row_count INTEGER DEFAULT 1,
				color VARCHAR(50) DEFAULT '#1f2937',
				opacity DECIMAL(3,2) DEFAULT 1.0,
				show_separator BOOLEAN DEFAULT false,
				separator_type VARCHAR(50) DEFAULT 'line',
				separator_color VARCHAR(50) DEFAULT '#ffffff',
				separator_opacity DECIMAL(3,2) DEFAULT 0.2,
				separator_width INTEGER DEFAULT 1,
				items TEXT,
				padding_left INTEGER DEFAULT 20,
				padding_right INTEGER DEFAULT 20,
				padding_top INTEGER DEFAULT 20,
				padding_bottom INTEGER DEFAULT 20,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
			)
		`).Error; err != nil {
			return err
		}

		// ایجاد indexها برای بهبود عملکرد
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_footers_is_active ON footers(is_active)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_footers_page_selection ON footers(page_selection)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_footer_layers_footer_id ON footer_layers(footer_id)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_footer_layers_padding ON footer_layers(padding_left, padding_right, padding_top, padding_bottom)`).Error; err != nil {
			return err
		}

		// ایجاد function برای updated_at
		if err := tx.Exec(`
			CREATE OR REPLACE FUNCTION update_footer_updated_at()
			RETURNS TRIGGER AS $$
			BEGIN
				NEW.updated_at = NOW();
				RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;
		`).Error; err != nil {
			return err
		}

		// حذف triggerهای موجود اگر وجود دارند
		_ = tx.Exec(`DROP TRIGGER IF EXISTS trigger_footer_updated_at ON footers`)
		_ = tx.Exec(`DROP TRIGGER IF EXISTS trigger_footer_layer_updated_at ON footer_layers`)

		// ایجاد triggerها
		if err := tx.Exec(`
			CREATE TRIGGER trigger_footer_updated_at
				BEFORE UPDATE ON footers
				FOR EACH ROW
				EXECUTE FUNCTION update_footer_updated_at();
		`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
			CREATE TRIGGER trigger_footer_layer_updated_at
				BEFORE UPDATE ON footer_layers
				FOR EACH ROW
				EXECUTE FUNCTION update_footer_updated_at();
		`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down102FooterCore حذف جداول فوتر و لایه‌های فوتر
func Down102FooterCore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف triggerها
		_ = tx.Exec(`DROP TRIGGER IF EXISTS trigger_footer_layer_updated_at ON footer_layers`)
		_ = tx.Exec(`DROP TRIGGER IF EXISTS trigger_footer_updated_at ON footers`)
		_ = tx.Exec(`DROP FUNCTION IF EXISTS update_footer_updated_at()`)

		// حذف indexها
		_ = tx.Exec(`DROP INDEX IF EXISTS idx_footer_layers_padding`)
		_ = tx.Exec(`DROP INDEX IF EXISTS idx_footer_layers_footer_id`)
		_ = tx.Exec(`DROP INDEX IF EXISTS idx_footers_page_selection`)
		_ = tx.Exec(`DROP INDEX IF EXISTS idx_footers_is_active`)

		// حذف جداول (ترتیب معکوس)
		_ = tx.Exec(`DROP TABLE IF EXISTS footer_layers`)
		_ = tx.Exec(`DROP TABLE IF EXISTS footers`)

		return nil
	})
}
