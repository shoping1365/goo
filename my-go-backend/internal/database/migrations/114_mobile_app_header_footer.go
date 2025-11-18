package migrations

import "gorm.io/gorm"

// Up114MobileAppHeaderFooter ایجاد جداول هدر و فوتر موبایل و اپلیکیشن
func Up114MobileAppHeaderFooter(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// ایجاد جدول mobile_app_headers
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS mobile_app_headers (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL UNIQUE,
				description TEXT,
				platform VARCHAR(50) DEFAULT 'mobile',
				page_selection VARCHAR(50) DEFAULT 'all',
				specific_pages TEXT,
				excluded_pages TEXT,
				is_active BOOLEAN DEFAULT true,
				created_at TIMESTAMP DEFAULT NOW(),
				updated_at TIMESTAMP DEFAULT NOW()
			)
		`).Error; err != nil {
			return err
		}

		// ایجاد جدول mobile_app_header_layers
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS mobile_app_header_layers (
				id SERIAL PRIMARY KEY,
				mobile_app_header_id INTEGER NOT NULL REFERENCES mobile_app_headers(id) ON DELETE CASCADE,
				name VARCHAR(255) NOT NULL,
				width INTEGER DEFAULT 100,
				height INTEGER DEFAULT 50,
				row_count INTEGER DEFAULT 1,
				color VARCHAR(50) DEFAULT '#ffffff',
				opacity DECIMAL(3,2) DEFAULT 1.0,
				show_separator BOOLEAN DEFAULT false,
				separator_type VARCHAR(50) DEFAULT 'line',
				separator_color VARCHAR(50) DEFAULT '#000000',
				separator_opacity DECIMAL(3,2) DEFAULT 1.0,
				separator_width INTEGER DEFAULT 1,
				items TEXT,
				padding_left INTEGER DEFAULT 20,
				padding_right INTEGER DEFAULT 20,
				padding_top INTEGER DEFAULT 10,
				padding_bottom INTEGER DEFAULT 10,
				created_at TIMESTAMP DEFAULT NOW(),
				updated_at TIMESTAMP DEFAULT NOW()
			)
		`).Error; err != nil {
			return err
		}

		// ایجاد جدول mobile_app_footers
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS mobile_app_footers (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL UNIQUE,
				description TEXT,
				platform VARCHAR(50) DEFAULT 'mobile',
				page_selection VARCHAR(50) DEFAULT 'all',
				specific_pages TEXT,
				excluded_pages TEXT,
				is_active BOOLEAN DEFAULT true,
				created_at TIMESTAMP DEFAULT NOW(),
				updated_at TIMESTAMP DEFAULT NOW()
			)
		`).Error; err != nil {
			return err
		}

		// ایجاد جدول mobile_app_footer_layers
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS mobile_app_footer_layers (
				id SERIAL PRIMARY KEY,
				mobile_app_footer_id INTEGER NOT NULL REFERENCES mobile_app_footers(id) ON DELETE CASCADE,
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
				created_at TIMESTAMP DEFAULT NOW(),
				updated_at TIMESTAMP DEFAULT NOW()
			)
		`).Error; err != nil {
			return err
		}

		// ایجاد ایندکس‌ها برای بهینه‌سازی عملکرد
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_mobile_app_headers_active ON mobile_app_headers(is_active)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_mobile_app_headers_platform ON mobile_app_headers(platform)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_mobile_app_header_layers_header_id ON mobile_app_header_layers(mobile_app_header_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_mobile_app_footers_active ON mobile_app_footers(is_active)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_mobile_app_footers_platform ON mobile_app_footers(platform)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_mobile_app_footer_layers_footer_id ON mobile_app_footer_layers(mobile_app_footer_id)`).Error; err != nil {
			return err
		}

		// اضافه کردن permissions برای مدیریت هدر و فوتر موبایل و اپلیکیشن
		permissions := []string{
			"INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('mobile_app_header_manage','مدیریت هدر موبایل و اپلیکیشن','content_management','mobile_app_header','manage','mobile_app_header',true) ON CONFLICT (name) DO NOTHING",
			"INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('mobile_app_header.create','ایجاد هدر موبایل و اپلیکیشن','content_management','mobile_app_header','create','mobile_app_header',true) ON CONFLICT (name) DO NOTHING",
			"INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('mobile_app_header.update','ویرایش هدر موبایل و اپلیکیشن','content_management','mobile_app_header','update','mobile_app_header',true) ON CONFLICT (name) DO NOTHING",
			"INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('mobile_app_header.delete','حذف هدر موبایل و اپلیکیشن','content_management','mobile_app_header','delete','mobile_app_header',true) ON CONFLICT (name) DO NOTHING",
			"INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('mobile_app_header.read','مشاهده هدر موبایل و اپلیکیشن','content_management','mobile_app_header','read','mobile_app_header',true) ON CONFLICT (name) DO NOTHING",
			"INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('mobile_app_footer_manage','مدیریت فوتر موبایل و اپلیکیشن','content_management','mobile_app_footer','manage','mobile_app_footer',true) ON CONFLICT (name) DO NOTHING",
			"INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('mobile_app_footer.create','ایجاد فوتر موبایل و اپلیکیشن','content_management','mobile_app_footer','create','mobile_app_footer',true) ON CONFLICT (name) DO NOTHING",
			"INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('mobile_app_footer.update','ویرایش فوتر موبایل و اپلیکیشن','content_management','mobile_app_footer','update','mobile_app_footer',true) ON CONFLICT (name) DO NOTHING",
			"INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('mobile_app_footer.delete','حذف فوتر موبایل و اپلیکیشن','content_management','mobile_app_footer','delete','mobile_app_footer',true) ON CONFLICT (name) DO NOTHING",
			"INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('mobile_app_footer.read','مشاهده فوتر موبایل و اپلیکیشن','content_management','mobile_app_footer','read','mobile_app_footer',true) ON CONFLICT (name) DO NOTHING",
		}

		for _, permission := range permissions {
			if err := tx.Exec(permission).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// Down114MobileAppHeaderFooter حذف جداول هدر و فوتر موبایل و اپلیکیشن
func Down114MobileAppHeaderFooter(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف جداول (به ترتیب معکوس به دلیل foreign key constraints)
		if err := tx.Exec(`DROP TABLE IF EXISTS mobile_app_footer_layers`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`DROP TABLE IF EXISTS mobile_app_footers`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`DROP TABLE IF EXISTS mobile_app_header_layers`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`DROP TABLE IF EXISTS mobile_app_headers`).Error; err != nil {
			return err
		}

		// حذف permissions مربوطه
		permissions := []string{
			"DELETE FROM permissions WHERE name IN ('mobile_app_header_manage', 'mobile_app_header.create', 'mobile_app_header.update', 'mobile_app_header.delete', 'mobile_app_header.read')",
			"DELETE FROM permissions WHERE name IN ('mobile_app_footer_manage', 'mobile_app_footer.create', 'mobile_app_footer.update', 'mobile_app_footer.delete', 'mobile_app_footer.read')",
		}

		for _, permission := range permissions {
			_ = tx.Exec(permission) // خطا را نادیده می‌گیریم چون ممکن است permissions وجود نداشته باشد
		}

		return nil
	})
}
