package migrations

import "gorm.io/gorm"

// Up029HeaderLayersFixes اعمال Up30/31/32/33 روی headers/header_layers به‌صورت idempotent
func Up029HeaderLayersFixes(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// ایجاد جدول headers/header_layers اگر نیستند
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS headers (
            id BIGSERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL UNIQUE,
            description TEXT,
            page_selection VARCHAR(50) DEFAULT 'all',
            specific_pages TEXT,
            excluded_pages TEXT,
            is_active BOOLEAN DEFAULT TRUE,
            created_at TIMESTAMPTZ DEFAULT NOW(),
            updated_at TIMESTAMPTZ DEFAULT NOW()
        );`)
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS header_layers (
            id BIGSERIAL PRIMARY KEY,
            header_id BIGINT NOT NULL REFERENCES headers(id) ON DELETE CASCADE,
            name VARCHAR(255) NOT NULL,
            width INT DEFAULT 100,
            height INT DEFAULT 50,
            row_count INT DEFAULT 1,
            color VARCHAR(50) DEFAULT '#ffffff',
            opacity DOUBLE PRECISION DEFAULT 1.0,
            show_separator BOOLEAN DEFAULT FALSE,
            separator_type VARCHAR(50) DEFAULT 'line',
            separator_color VARCHAR(50) DEFAULT '#000000',
            separator_opacity DOUBLE PRECISION DEFAULT 1.0,
            separator_width INT DEFAULT 1,
            page_selection VARCHAR(50) DEFAULT 'all',
            specific_pages TEXT,
            excluded_pages TEXT,
            items TEXT,
            created_at TIMESTAMPTZ DEFAULT NOW(),
            updated_at TIMESTAMPTZ DEFAULT NOW()
        );`)
		// 31: add show_on_mobile/desktop
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS show_on_mobile BOOLEAN DEFAULT TRUE`)
		_ = tx.Exec(`ALTER TABLE header_layers ADD COLUMN IF NOT EXISTS show_on_desktop BOOLEAN DEFAULT TRUE`)
		// 32/33: حذف/بازگردانی deleted_at را لازم نمی‌دانیم؛ اطمینان از نبود آن
		_ = tx.Exec(`DO $$ BEGIN IF EXISTS(SELECT 1 FROM information_schema.columns WHERE table_name='headers' AND column_name='deleted_at') THEN ALTER TABLE headers DROP COLUMN deleted_at; END IF; END $$;`)
		_ = tx.Exec(`DO $$ BEGIN IF EXISTS(SELECT 1 FROM information_schema.columns WHERE table_name='header_layers' AND column_name='deleted_at') THEN ALTER TABLE header_layers DROP COLUMN deleted_at; END IF; END $$;`)
		return nil
	})
}
