package migrations

import (
	"gorm.io/gorm"
)

// Up072SEORedirects ایجاد جدول ریدایرکت‌های سئو در صورت عدم وجود
func Up072SEORedirects(db *gorm.DB) error {
	// ایجاد جدول اگر وجود ندارد
	if err := db.Exec(`
    CREATE TABLE IF NOT EXISTS seo_redirects (
        id SERIAL PRIMARY KEY,
        source_path VARCHAR(500) NOT NULL UNIQUE,
        target_path VARCHAR(500) NOT NULL,
        code INT NOT NULL DEFAULT 301,
        group_name VARCHAR(100) NOT NULL DEFAULT 'Imported',
        page_title VARCHAR(200) NULL,
        redirect_type VARCHAR(50) NOT NULL DEFAULT 'permanent',
        visit_count BIGINT NOT NULL DEFAULT 0,
        last_visited_at TIMESTAMPTZ NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );
    `).Error; err != nil {
		return err
	}
	// در صورت وجود جدول، ستون‌ها را در صورت نبود اضافه کن
	if err := db.Exec(`
        ALTER TABLE seo_redirects
        ADD COLUMN IF NOT EXISTS page_title VARCHAR(200) NULL,
        ADD COLUMN IF NOT EXISTS redirect_type VARCHAR(50) NOT NULL DEFAULT 'permanent',
        ADD COLUMN IF NOT EXISTS visit_count BIGINT NOT NULL DEFAULT 0,
        ADD COLUMN IF NOT EXISTS last_visited_at TIMESTAMPTZ NULL;
    `).Error; err != nil {
		return err
	}
	if err := db.Exec(`
        CREATE INDEX IF NOT EXISTS idx_seo_redirects_group_name ON seo_redirects(group_name);
        CREATE INDEX IF NOT EXISTS idx_seo_redirects_last_visited ON seo_redirects(last_visited_at);
    `).Error; err != nil {
		return err
	}
	return nil
}
