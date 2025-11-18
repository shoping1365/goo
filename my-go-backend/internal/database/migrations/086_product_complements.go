package migrations

import "gorm.io/gorm"

// Up086ProductComplements ایجاد جدول روابط محصولات مکمل
func Up086ProductComplements(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS product_complements (
                product_id BIGINT NOT NULL,
                complement_product_id BIGINT NOT NULL,
                priority INT,
                required BOOLEAN NOT NULL DEFAULT FALSE,
                quantity INT NOT NULL DEFAULT 1,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_pc_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE,
                CONSTRAINT fk_pc_complement FOREIGN KEY(complement_product_id) REFERENCES products(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_product_complements_unique ON product_complements(product_id, complement_product_id);`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_complements_product ON product_complements(product_id);`).Error; err != nil {
			return err
		}
		return nil
	})
}
