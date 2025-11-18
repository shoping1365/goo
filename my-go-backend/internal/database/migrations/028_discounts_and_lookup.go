package migrations

import "gorm.io/gorm"

// Up028DiscountsAndLookup ایجاد جداول Lookup و تخفیف‌ها (معادل 07)
func Up028DiscountsAndLookup(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS lookup_tables (
                id BIGSERIAL PRIMARY KEY,
                type VARCHAR(50),
                key VARCHAR(100),
                value VARCHAR(255)
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_lookup_type_key ON lookup_tables(type, key)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS discount_codes (
                id BIGSERIAL PRIMARY KEY,
                code VARCHAR(50) UNIQUE,
                discount_type VARCHAR(20),
                value DOUBLE PRECISION,
                max_usage BIGINT,
                usage_count BIGINT,
                is_active BOOLEAN NOT NULL DEFAULT TRUE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_discount_active ON discount_codes(is_active)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS discount_usages (
                id BIGSERIAL PRIMARY KEY,
                discount_code_id BIGINT NOT NULL,
                order_id BIGINT,
                user_id BIGINT
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_discount_usage_code ON discount_usages(discount_code_id)`).Error; err != nil {
			return err
		}
		return nil
	})
}
