package migrations

import "gorm.io/gorm"

// Up009CMS جداول پست‌ها، دسته‌بندی پست‌ها، ویجت‌ها، منو و هدر
func Up009CMS(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// post_categories
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS post_categories (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL UNIQUE,
                slug VARCHAR(120) NOT NULL UNIQUE,
                description TEXT,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                "order" INT NOT NULL DEFAULT 1,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		// posts (سازگار با query زمان‌بندی انتشار)
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS posts (
                id BIGSERIAL PRIMARY KEY,
                title VARCHAR(255) NOT NULL,
                slug VARCHAR(255) NOT NULL UNIQUE,
                excerpt TEXT,
                content TEXT,
                status VARCHAR(20) NOT NULL DEFAULT 'draft',
                category_id BIGINT,
                featured_image VARCHAR(255),
                author_id BIGINT NOT NULL,
                published_at TIMESTAMPTZ,
                scheduled_at TIMESTAMPTZ,
                meta_title VARCHAR(255),
                meta_keywords VARCHAR(255),
                meta_description VARCHAR(255),
                robots_meta VARCHAR(50) NOT NULL DEFAULT 'index,follow',
                estimated_time INT NOT NULL DEFAULT 0,
                location VARCHAR(255),
                event_date TIMESTAMPTZ,
                pros TEXT,
                cons TEXT,
                tags JSONB,
                language VARCHAR(10) NOT NULL DEFAULT 'fa',
                is_accessible_for_free BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_posts_author FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE RESTRICT
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_posts_status_scheduled ON posts(status, scheduled_at)").Error; err != nil {
			return err
		}

		// menus
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS menus (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL,
                slug VARCHAR(100) NOT NULL UNIQUE,
                description VARCHAR(500),
                type VARCHAR(50) NOT NULL DEFAULT 'standard',
                theme VARCHAR(50) NOT NULL DEFAULT 'default',
                max_depth INT NOT NULL DEFAULT 3,
                enabled BOOLEAN NOT NULL DEFAULT TRUE,
                show_icons BOOLEAN NOT NULL DEFAULT TRUE,
                show_badges BOOLEAN NOT NULL DEFAULT TRUE,
                "order" INT NOT NULL DEFAULT 0,
                is_mega_menu BOOLEAN NOT NULL DEFAULT FALSE,
                mega_width VARCHAR(20) NOT NULL DEFAULT 'full',
                mega_height VARCHAR(20) NOT NULL DEFAULT 'auto',
                mega_columns INT NOT NULL DEFAULT 2,
                show_on_mobile BOOLEAN NOT NULL DEFAULT TRUE,
                show_on_desktop BOOLEAN NOT NULL DEFAULT TRUE,
                show_on_tablet BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS menu_items (
                id BIGSERIAL PRIMARY KEY,
                menu_id BIGINT NOT NULL,
                parent_id BIGINT,
                column_id BIGINT,
                title VARCHAR(100) NOT NULL,
                path VARCHAR(200),
                icon VARCHAR(100),
                badge VARCHAR(50),
                badge_color VARCHAR(20) NOT NULL DEFAULT 'red',
                enabled BOOLEAN NOT NULL DEFAULT TRUE,
                "order" INT NOT NULL DEFAULT 0,
                open_in_new_tab BOOLEAN NOT NULL DEFAULT FALSE,
                item_type VARCHAR(20) NOT NULL DEFAULT 'link',
                target_id BIGINT,
                target_type VARCHAR(50),
                image_url VARCHAR(500),
                description VARCHAR(300),
                featured BOOLEAN NOT NULL DEFAULT FALSE,
                show_on_mobile BOOLEAN NOT NULL DEFAULT TRUE,
                show_on_desktop BOOLEAN NOT NULL DEFAULT TRUE,
                show_on_tablet BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_menu_items_menu FOREIGN KEY(menu_id) REFERENCES menus(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS menu_columns (
                id BIGSERIAL PRIMARY KEY,
                menu_id BIGINT NOT NULL,
                title VARCHAR(100),
                description VARCHAR(300),
                width VARCHAR(20) NOT NULL DEFAULT 'auto',
                "order" INT NOT NULL DEFAULT 0,
                enabled BOOLEAN NOT NULL DEFAULT TRUE,
                column_type VARCHAR(20) NOT NULL DEFAULT 'items',
                show_on_mobile BOOLEAN NOT NULL DEFAULT TRUE,
                show_on_desktop BOOLEAN NOT NULL DEFAULT TRUE,
                show_on_tablet BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_menu_columns_menu FOREIGN KEY(menu_id) REFERENCES menus(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS menu_locations (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL UNIQUE,
                slug VARCHAR(100) NOT NULL UNIQUE,
                description VARCHAR(500),
                icon VARCHAR(100),
                enabled BOOLEAN NOT NULL DEFAULT TRUE,
                "order" INT NOT NULL DEFAULT 0,
                max_menus INT NOT NULL DEFAULT 1,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS menu_assignments (
                id BIGSERIAL PRIMARY KEY,
                menu_id BIGINT NOT NULL,
                location_id BIGINT NOT NULL,
                "order" INT NOT NULL DEFAULT 0,
                enabled BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_menu_assign_menu FOREIGN KEY(menu_id) REFERENCES menus(id) ON DELETE CASCADE,
                CONSTRAINT fk_menu_assign_location FOREIGN KEY(location_id) REFERENCES menu_locations(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}

		// widgets
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS widgets (
                id BIGSERIAL PRIMARY KEY,
                code VARCHAR(100) NOT NULL UNIQUE,
                title VARCHAR(255) NOT NULL,
                description TEXT,
                type VARCHAR(50) NOT NULL,
                status VARCHAR(20) NOT NULL DEFAULT 'active',
                sort_order INT NOT NULL DEFAULT 1,
                image VARCHAR(500),
                link VARCHAR(500),
                page VARCHAR(50) NOT NULL DEFAULT 'home',
                config JSONB,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// headers & header_layers
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS headers (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(255) NOT NULL UNIQUE,
                description TEXT,
                page_selection VARCHAR(50) NOT NULL DEFAULT 'all',
                specific_pages TEXT,
                excluded_pages TEXT,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS header_layers (
                id BIGSERIAL PRIMARY KEY,
                header_id BIGINT NOT NULL,
                name VARCHAR(255) NOT NULL,
                width INT NOT NULL DEFAULT 100,
                height INT NOT NULL DEFAULT 50,
                row_count INT NOT NULL DEFAULT 1,
                color VARCHAR(50) NOT NULL DEFAULT '#ffffff',
                opacity DOUBLE PRECISION NOT NULL DEFAULT 1.0,
                show_separator BOOLEAN NOT NULL DEFAULT FALSE,
                separator_type VARCHAR(50) NOT NULL DEFAULT 'line',
                separator_color VARCHAR(50) NOT NULL DEFAULT '#000000',
                separator_opacity DOUBLE PRECISION NOT NULL DEFAULT 1.0,
                separator_width INT NOT NULL DEFAULT 1,
                page_selection VARCHAR(50) NOT NULL DEFAULT 'all',
                specific_pages TEXT,
                excluded_pages TEXT,
                items TEXT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_header_layers_header FOREIGN KEY(header_id) REFERENCES headers(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		return nil
	})
}
