package migrations

import "gorm.io/gorm"

// Up070ProductAttributeValuesCore ایجاد جدول product_attribute_values برای آمار ویژگی‌ها
func Up070ProductAttributeValuesCore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS product_attribute_values (
                id BIGSERIAL PRIMARY KEY,
                product_id BIGINT NOT NULL,
                attribute_id BIGINT NOT NULL,
                attribute_value_id BIGINT NULL,
                value_text TEXT NULL,
                value_numeric NUMERIC(18,4) NULL,
                value_bool BOOLEAN NULL,
                value_date TIMESTAMPTZ NULL,
                sort_order INT DEFAULT 0,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ NULL,
                CONSTRAINT fk_pav_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE,
                CONSTRAINT fk_pav_attribute FOREIGN KEY(attribute_id) REFERENCES attributes(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_pav_product ON product_attribute_values(product_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_pav_attribute ON product_attribute_values(attribute_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_pav_deleted_at ON product_attribute_values(deleted_at)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_pav_product_attribute_option ON product_attribute_values(product_id, attribute_id, attribute_value_id)`).Error; err != nil {
			return err
		}
		return nil
	})
}


