package migrations

import "gorm.io/gorm"

// Up003Catalog ایجاد جداول کاتالوگ پایه: brands, categories, category_brand_pages
// این مایگریشن بدون وابستگی به محصولات قابل اجراست و باید idempotent باشد.
func Up003Catalog(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// brands
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS brands (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL UNIQUE,
                slug VARCHAR(120) NOT NULL UNIQUE,
                official_name VARCHAR(150),
                description TEXT,
                published BOOLEAN NOT NULL DEFAULT FALSE,
                show_on_home BOOLEAN NOT NULL DEFAULT FALSE,
                show_in_menu BOOLEAN NOT NULL DEFAULT FALSE,
                meta_title VARCHAR(150),
                meta_keywords VARCHAR(255),
                meta_description VARCHAR(255),
                image_url VARCHAR(255),
                video_url VARCHAR(255),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_brands_published ON brands(published)`).Error; err != nil {
			return err
		}

		// categories
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS categories (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL UNIQUE,
                name_en VARCHAR(150),
                slug VARCHAR(120) NOT NULL UNIQUE,
                parent_id BIGINT,
                description TEXT,
                image_url VARCHAR(255),
                icon VARCHAR(100),
                "order" INT NOT NULL DEFAULT 1,
                published BOOLEAN NOT NULL DEFAULT TRUE,
                show_on_home BOOLEAN NOT NULL DEFAULT FALSE,
                show_in_menu BOOLEAN NOT NULL DEFAULT FALSE,
                banner_url VARCHAR(255),
                notice_message TEXT,
                meta_title VARCHAR(150),
                meta_keywords VARCHAR(255),
                meta_description VARCHAR(255),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ,
                CONSTRAINT fk_categories_parent FOREIGN KEY(parent_id) REFERENCES categories(id) ON DELETE SET NULL
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_categories_published ON categories(published)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_categories_order ON categories("order")`).Error; err != nil {
			return err
		}

		// category_brand_pages
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS category_brand_pages (
                id BIGSERIAL PRIMARY KEY,
                category_id BIGINT NOT NULL,
                brand_id BIGINT NOT NULL,
                slug VARCHAR(180) NOT NULL UNIQUE,
                title VARCHAR(150),
                description TEXT,
                banner_url VARCHAR(255),
                video_url VARCHAR(255),
                meta_title VARCHAR(150),
                meta_keywords VARCHAR(255),
                meta_description VARCHAR(255),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ,
                CONSTRAINT fk_cbp_category FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE,
                CONSTRAINT fk_cbp_brand FOREIGN KEY(brand_id) REFERENCES brands(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_cbp_category_brand ON category_brand_pages(category_id, brand_id)`).Error; err != nil {
			return err
		}

		return nil
	})
}
