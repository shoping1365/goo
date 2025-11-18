package migrations

import "gorm.io/gorm"

// Up016ApiSettingsAndDeltas پوشش مایگریشن‌های قدیمی 23، 26، 28–33 به‌صورت idempotent
func Up016ApiSettingsAndDeltas(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 23_api_settings
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS api_settings (
                id BIGSERIAL PRIMARY KEY,
                provider VARCHAR(50) UNIQUE,
                config JSONB NOT NULL DEFAULT '{}'::jsonb,
                usage_stats JSONB NOT NULL DEFAULT '{}'::jsonb,
                key VARCHAR(100) UNIQUE,
                value TEXT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		_ = tx.Exec(`ALTER TABLE api_settings ADD COLUMN IF NOT EXISTS provider VARCHAR(50)`)
		_ = tx.Exec(`ALTER TABLE api_settings ADD COLUMN IF NOT EXISTS config JSONB`)
		_ = tx.Exec(`ALTER TABLE api_settings ADD COLUMN IF NOT EXISTS usage_stats JSONB`)
		_ = tx.Exec(`ALTER TABLE api_settings ADD COLUMN IF NOT EXISTS key VARCHAR(100)`)
		_ = tx.Exec(`ALTER TABLE api_settings ADD COLUMN IF NOT EXISTS value TEXT`)
		_ = tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_api_settings_key_unique ON api_settings(key)`)
		_ = tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_api_settings_provider_unique ON api_settings(provider)`)

		// 26_ai_chat_sessions (ai_sessions + ai_chat_messages)
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS ai_sessions (
                id BIGSERIAL PRIMARY KEY,
                session_id VARCHAR(255) NOT NULL UNIQUE,
                user_id BIGINT,
                context_type VARCHAR(50) NOT NULL,
                context_ref_id BIGINT,
                current_state VARCHAR(50) NOT NULL DEFAULT 'initial',
                selected_title TEXT,
                draft_article JSONB,
                edit_history JSONB,
                metadata JSONB,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_ai_sessions_user_id ON ai_sessions(user_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_ai_sessions_context_type ON ai_sessions(context_type)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_ai_sessions_context_ref_id ON ai_sessions(context_ref_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_ai_sessions_created_at ON ai_sessions(created_at)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS ai_chat_messages (
                id BIGSERIAL PRIMARY KEY,
                session_id VARCHAR(255) NOT NULL,
                role VARCHAR(20) NOT NULL,
                content TEXT NOT NULL,
                type VARCHAR(50),
                metadata JSONB,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_ai_messages_session FOREIGN KEY (session_id) REFERENCES ai_sessions(session_id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_ai_messages_session_id ON ai_chat_messages(session_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_ai_messages_created_at ON ai_chat_messages(created_at)`).Error; err != nil {
			return err
		}

		// 28_header_settings → درج تنظیمات پیش‌فرض در settings (idempotent)
		// کلیدهای header_* فقط اگر وجود نداشتند اضافه می‌شوند
		ensureSetting := func(key, val, typ string, isPublic bool) error {
			q := `INSERT INTO settings(key, value, category, type, is_public) 
                  SELECT ?, ?, 'header', ?, ?
                  WHERE NOT EXISTS (SELECT 1 FROM settings WHERE key = ? AND category = 'header')`
			return tx.Exec(q, key, val, typ, isPublic, key).Error
		}
		if err := ensureSetting("header_logo_url", "/default-logo.png", "string", true); err != nil {
			return err
		}
		if err := ensureSetting("header_logo_alt", "ابزاریست", "string", true); err != nil {
			return err
		}
		if err := ensureSetting("header_show_search", "true", "boolean", true); err != nil {
			return err
		}
		if err := ensureSetting("header_show_cart", "true", "boolean", true); err != nil {
			return err
		}
		if err := ensureSetting("header_show_user_menu", "true", "boolean", true); err != nil {
			return err
		}
		if err := ensureSetting("header_phone_number", "021-12345678", "string", true); err != nil {
			return err
		}
		if err := ensureSetting("header_email", "info@abzariest.com", "string", true); err != nil {
			return err
		}
		if err := ensureSetting("header_show_social_links", "true", "boolean", true); err != nil {
			return err
		}
		if err := ensureSetting("header_facebook_url", "https://facebook.com/abzariest", "string", true); err != nil {
			return err
		}
		if err := ensureSetting("header_instagram_url", "https://instagram.com/abzariest", "string", true); err != nil {
			return err
		}
		if err := ensureSetting("header_telegram_url", "https://t.me/abzariest", "string", true); err != nil {
			return err
		}
		if err := ensureSetting("header_whatsapp_url", "https://wa.me/989123456789", "string", true); err != nil {
			return err
		}
		if err := ensureSetting("header_sticky", "true", "boolean", true); err != nil {
			return err
		}
		if err := ensureSetting("header_background_color", "#ffffff", "string", true); err != nil {
			return err
		}
		if err := ensureSetting("header_text_color", "#333333", "string", true); err != nil {
			return err
		}
		if err := ensureSetting("header_height", "80px", "string", true); err != nil {
			return err
		}

		return nil
	})
}
