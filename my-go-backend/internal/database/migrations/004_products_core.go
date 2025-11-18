package migrations

import "gorm.io/gorm"

// Up004ProductsCore ایجاد جداول محصول، تصاویر، و تنوع‌ها با ایندکس‌های ضروری
func Up004ProductsCore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// products
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS products (
                id BIGSERIAL PRIMARY KEY,
                uuid UUID NOT NULL UNIQUE,
                slug VARCHAR(255) NOT NULL UNIQUE,
                sku VARCHAR(100) UNIQUE,
                name VARCHAR(255),
                name_en VARCHAR(255),
                image VARCHAR(255),
                price NUMERIC(12,2),
                old_price NUMERIC(12,2),
                cost NUMERIC(12,2),
                profit NUMERIC(12,2),
                discount_percent NUMERIC(5,2),
                discount_amount NUMERIC(12,2),
                sale_price NUMERIC(12,2),
                wholesale_price NUMERIC(12,2),
                wholesale_sale_price NUMERIC(12,2),
                description TEXT,
                full_description TEXT,
                status VARCHAR(50),
                is_new BOOLEAN DEFAULT FALSE,
                weight NUMERIC(12,3),
                length NUMERIC(12,3),
                width NUMERIC(12,3),
                height NUMERIC(12,3),
                shipping_cost NUMERIC(12,2),
                shipping_time INT,
                stock_quantity INT,
                min_stock_quantity INT,
                max_stock_quantity INT,
                stock_status VARCHAR(50),
                track_inventory BOOLEAN DEFAULT TRUE,
                show_stock_to_customer BOOLEAN DEFAULT FALSE,
                disable_buy_button BOOLEAN DEFAULT FALSE,
                call_for_price BOOLEAN DEFAULT FALSE,
                seo_title VARCHAR(255),
                meta_description VARCHAR(255),
                meta_keywords VARCHAR(255),
                category_id BIGINT,
                brand_id BIGINT,
                is_active BOOLEAN DEFAULT TRUE,
                is_featured BOOLEAN DEFAULT FALSE,
                is_on_sale BOOLEAN DEFAULT FALSE,
                meta_title VARCHAR(255),
                tags TEXT,
                updated_by BIGINT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                video_url VARCHAR(255),
                CONSTRAINT fk_products_category FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE SET NULL,
                CONSTRAINT fk_products_brand FOREIGN KEY(brand_id) REFERENCES brands(id) ON DELETE SET NULL
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_products_status_category_price ON products(status, category_id, price) WHERE status = 'published'`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_products_brand_status_price ON products(brand_id, status, price) WHERE status = 'published'`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_products_stock_track_status ON products(stock_quantity, track_inventory, status)`).Error; err != nil {
			return err
		}

		// product_images
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS product_images (
                id BIGSERIAL PRIMARY KEY,
                product_id BIGINT NOT NULL,
                image_url VARCHAR(255),
                sort_order INT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_product_images_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_images_product_order ON product_images(product_id, sort_order)`).Error; err != nil {
			return err
		}

		// product_variants
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS product_variants (
                id BIGSERIAL PRIMARY KEY,
                product_id BIGINT NOT NULL,
                name VARCHAR(255),
                value VARCHAR(255),
                price_adjustment NUMERIC(12,2),
                stock_quantity INT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_product_variants_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_variants_product ON product_variants(product_id)`).Error; err != nil {
			return err
		}

		// product_specifications
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS product_specifications (
                id BIGSERIAL PRIMARY KEY,
                product_id BIGINT NOT NULL,
                name VARCHAR(255),
                value TEXT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_product_spec_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_spec_product ON product_specifications(product_id)`).Error; err != nil {
			return err
		}
		return nil
	})
}
