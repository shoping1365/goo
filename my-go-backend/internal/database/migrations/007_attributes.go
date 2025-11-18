package migrations

import "gorm.io/gorm"

// Up007Attributes سیستم ویژگی‌ها و گروه‌بندی آن‌ها
func Up007Attributes(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// attributes
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS attributes (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL UNIQUE,
                display_name VARCHAR(150),
                code VARCHAR(50) UNIQUE,
                data_type VARCHAR(20) NOT NULL,
                unit VARCHAR(30),
                sort_order INT NOT NULL DEFAULT 0,
                is_required BOOLEAN NOT NULL DEFAULT FALSE,
                is_filterable BOOLEAN NOT NULL DEFAULT FALSE,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                has_options BOOLEAN NOT NULL DEFAULT FALSE,
                attribute_group_id BIGINT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_attributes_active ON attributes(is_active)`).Error; err != nil {
			return err
		}

		// attribute_values
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS attribute_values (
                id BIGSERIAL PRIMARY KEY,
                attribute_id BIGINT NOT NULL,
                value VARCHAR(100) NOT NULL,
                text_value VARCHAR(255),
                numeric_value DOUBLE PRECISION,
                bool_value BOOLEAN,
                date_value TIMESTAMPTZ,
                slug VARCHAR(100),
                sort_order INT NOT NULL DEFAULT 0,
                meta JSONB,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_attribute_values_attr FOREIGN KEY(attribute_id) REFERENCES attributes(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_attribute_values_attr ON attribute_values(attribute_id)`).Error; err != nil {
			return err
		}

		// attribute_groups
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS attribute_groups (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL,
                category_id BIGINT NOT NULL,
                description TEXT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ,
                CONSTRAINT fk_attribute_groups_category FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_attribute_groups_category ON attribute_groups(category_id)`).Error; err != nil {
			return err
		}

		// attribute_group_attributes (join)
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS attribute_group_attributes (
                id BIGSERIAL PRIMARY KEY,
                attribute_group_id BIGINT NOT NULL,
                attribute_id BIGINT NOT NULL,
                sort_order INT NOT NULL DEFAULT 1,
                type VARCHAR(30) NOT NULL,
                control_type VARCHAR(30) NOT NULL,
                has_filter BOOLEAN NOT NULL DEFAULT FALSE,
                is_key BOOLEAN NOT NULL DEFAULT FALSE,
                show_on_product BOOLEAN NOT NULL DEFAULT TRUE,
                is_required BOOLEAN NOT NULL DEFAULT FALSE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ,
                CONSTRAINT fk_aga_group FOREIGN KEY(attribute_group_id) REFERENCES attribute_groups(id) ON DELETE CASCADE,
                CONSTRAINT fk_aga_attr FOREIGN KEY(attribute_id) REFERENCES attributes(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_aga_unique ON attribute_group_attributes(attribute_group_id, attribute_id)`).Error; err != nil {
			return err
		}
		return nil
	})
}
