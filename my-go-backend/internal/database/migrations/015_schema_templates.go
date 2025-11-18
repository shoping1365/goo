package migrations

import "gorm.io/gorm"

// Up015SchemaTemplates جداول schemas و admin_settings
func Up015SchemaTemplates(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS schemas (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(255) NOT NULL,
                type VARCHAR(50) NOT NULL,
                description TEXT,
                is_template BOOLEAN NOT NULL DEFAULT TRUE,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                title VARCHAR(255),
                slug VARCHAR(255),
                excerpt TEXT,
                content TEXT,
                site_name VARCHAR(255),
                featured_image VARCHAR(500),
                article_id VARCHAR(100),
                article_url VARCHAR(500),
                author VARCHAR(255),
                published_at TIMESTAMPTZ,
                price DECIMAL(10,2),
                currency VARCHAR(10) NOT NULL DEFAULT 'IRR',
                address VARCHAR(500),
                telephone VARCHAR(50),
                meta_title VARCHAR(255),
                meta_keywords VARCHAR(255),
                meta_description VARCHAR(500),
                og_title VARCHAR(255),
                og_description VARCHAR(500),
                og_image VARCHAR(500),
                og_type VARCHAR(50),
                og_site_name VARCHAR(100),
                extra_fields JSONB,
                created_by BIGINT,
                updated_by BIGINT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS admin_settings (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT,
                key VARCHAR(100) NOT NULL,
                value TEXT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		return nil
	})
}
