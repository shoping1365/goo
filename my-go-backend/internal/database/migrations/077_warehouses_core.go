package migrations

import "gorm.io/gorm"

// Up077WarehousesCore ایجاد جداول انبار و موجودی چندانباره
func Up077WarehousesCore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// warehouses
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS warehouses (
                id BIGSERIAL PRIMARY KEY,
                code VARCHAR(64) NOT NULL UNIQUE,
                name VARCHAR(255) NOT NULL,
                city VARCHAR(100),
                address VARCHAR(500),
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                is_default BOOLEAN NOT NULL DEFAULT FALSE,
                priority INT NOT NULL DEFAULT 100,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// product_warehouse_stocks
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS product_warehouse_stocks (
                id BIGSERIAL PRIMARY KEY,
                product_id BIGINT NOT NULL,
                warehouse_id BIGINT NOT NULL,
                quantity INT NOT NULL DEFAULT 0,
                reserved INT NOT NULL DEFAULT 0,
                min_qty INT NOT NULL DEFAULT 0,
                max_qty INT NOT NULL DEFAULT 0,
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_pws_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE,
                CONSTRAINT fk_pws_warehouse FOREIGN KEY(warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE,
                CONSTRAINT uq_pws UNIQUE (product_id, warehouse_id)
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_pws_product ON product_warehouse_stocks(product_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_pws_warehouse ON product_warehouse_stocks(warehouse_id)`).Error; err != nil {
			return err
		}

		// inventory_movements
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS inventory_movements (
                id BIGSERIAL PRIMARY KEY,
                product_id BIGINT NOT NULL,
                warehouse_id BIGINT NOT NULL,
                delta INT NOT NULL,
                type VARCHAR(32) NOT NULL,
                ref_order_id BIGINT,
                notes VARCHAR(500),
                created_by BIGINT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_im_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE,
                CONSTRAINT fk_im_warehouse FOREIGN KEY(warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_im_product_created ON inventory_movements(product_id, created_at DESC)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_im_warehouse_created ON inventory_movements(warehouse_id, created_at DESC)`).Error; err != nil {
			return err
		}

		// Seed a default warehouse if none exists
		if err := tx.Exec(`
            INSERT INTO warehouses (code, name, city, is_default)
            SELECT 'DEFAULT', 'Default Warehouse', '', TRUE
            WHERE NOT EXISTS (SELECT 1 FROM warehouses)
        `).Error; err != nil {
			return err
		}

		// Backfill existing product stocks into default warehouse
		if err := tx.Exec(`
            INSERT INTO product_warehouse_stocks (product_id, warehouse_id, quantity, reserved)
            SELECT p.id, (SELECT id FROM warehouses WHERE is_default = TRUE LIMIT 1), COALESCE(p.stock_quantity,0), 0
            FROM products p
            WHERE COALESCE(p.stock_quantity,0) > 0
              AND NOT EXISTS (
                SELECT 1 FROM product_warehouse_stocks x WHERE x.product_id = p.id AND x.warehouse_id = (SELECT id FROM warehouses WHERE is_default = TRUE LIMIT 1)
              )
        `).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down077WarehousesCore حذف جداول ایجاد شده (در صورت نیاز به rollback)
func Down077WarehousesCore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`DROP TABLE IF EXISTS inventory_movements`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`DROP TABLE IF EXISTS product_warehouse_stocks`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`DROP TABLE IF EXISTS warehouses`).Error; err != nil {
			return err
		}
		return nil
	})
}
