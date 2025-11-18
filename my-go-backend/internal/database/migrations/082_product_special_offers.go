package migrations

import "gorm.io/gorm"

// Up082ProductSpecialOffers ایجاد جدول طرح‌های فروش ویژه مرحله‌ای برای محصولات
// هر ردیف شامل یک قیمت فروش ویژه و سقف تعداد قابل فروش در آن پله است.
func Up082ProductSpecialOffers(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS product_special_offers (
                id BIGSERIAL PRIMARY KEY,
                product_id BIGINT NOT NULL,
                price NUMERIC(12,2) NOT NULL,
                quantity INT NOT NULL,
                sort_order INT NOT NULL DEFAULT 0,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_pso_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_pso_product_order ON product_special_offers(product_id, sort_order, id)`).Error; err != nil {
			return err
		}
		return nil
	})
}

