package migrations

import "gorm.io/gorm"

// Up115MobileAppNavigation ایجاد جداول ناوبری موبایل و اپلیکیشن
func Up115MobileAppNavigation(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// ایجاد جدول mobile_app_navigations
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS mobile_app_navigations (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL UNIQUE,
				description TEXT,
				platform VARCHAR(50) NOT NULL DEFAULT 'mobile',
				page_selection VARCHAR(50) NOT NULL DEFAULT 'all',
				specific_pages TEXT,
				excluded_pages TEXT,
				is_active BOOLEAN DEFAULT true,
				created_at TIMESTAMP DEFAULT NOW(),
				updated_at TIMESTAMP DEFAULT NOW(),
				deleted_at TIMESTAMP NULL
			)
		`).Error; err != nil {
			return err
		}

		// ایجاد جدول mobile_app_navigation_layers
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS mobile_app_navigation_layers (
				id SERIAL PRIMARY KEY,
				mobile_app_navigation_id INTEGER NOT NULL,
				name VARCHAR(255) NOT NULL,
				type VARCHAR(50) NOT NULL,
				position VARCHAR(50),
				items TEXT,
				style TEXT,
				is_active BOOLEAN DEFAULT true,
				sort_order INTEGER DEFAULT 0,
				created_at TIMESTAMP DEFAULT NOW(),
				updated_at TIMESTAMP DEFAULT NOW(),
				CONSTRAINT fk_mobile_app_navigation_layers_navigation_id 
					FOREIGN KEY (mobile_app_navigation_id) 
					REFERENCES mobile_app_navigations(id) 
					ON DELETE CASCADE
			)
		`).Error; err != nil {
			return err
		}

		// ایجاد ایندکس‌ها
		if err := tx.Exec(`
			CREATE INDEX IF NOT EXISTS idx_mobile_app_navigations_active 
				ON mobile_app_navigations(is_active)
		`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
			CREATE INDEX IF NOT EXISTS idx_mobile_app_navigations_platform 
				ON mobile_app_navigations(platform)
		`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
			CREATE INDEX IF NOT EXISTS idx_mobile_app_navigation_layers_navigation_id 
				ON mobile_app_navigation_layers(mobile_app_navigation_id)
		`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
			CREATE INDEX IF NOT EXISTS idx_mobile_app_navigation_layers_sort_order 
				ON mobile_app_navigation_layers(sort_order)
		`).Error; err != nil {
			return err
		}

		// ایجاد ایندکس برای soft delete
		if err := tx.Exec(`
			CREATE INDEX IF NOT EXISTS idx_mobile_app_navigations_deleted_at 
				ON mobile_app_navigations(deleted_at)
		`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down115MobileAppNavigation حذف جداول ناوبری موبایل و اپلیکیشن
func Down115MobileAppNavigation(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف جداول (لایه‌ها به صورت خودکار حذف می‌شوند به دلیل CASCADE)
		if err := tx.Exec(`DROP TABLE IF EXISTS mobile_app_navigation_layers`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`DROP TABLE IF EXISTS mobile_app_navigations`).Error; err != nil {
			return err
		}

		return nil
	})
}
