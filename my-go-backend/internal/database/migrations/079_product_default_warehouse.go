package migrations

import "gorm.io/gorm"

// Up079ProductDefaultWarehouse افزودن ستون پیش‌فرض انبار برای هر محصول
func Up079ProductDefaultWarehouse(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS default_warehouse_id BIGINT`).Error; err != nil {
			return err
		}
		// ایندکس جهت جستجوهای احتمالی
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_products_default_warehouse ON products(default_warehouse_id)`).Error; err != nil {
			return err
		}
		return nil
	})
}

func Down079ProductDefaultWarehouse(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec(`DROP INDEX IF EXISTS idx_products_default_warehouse`).Error
		_ = tx.Exec(`ALTER TABLE products DROP COLUMN IF EXISTS default_warehouse_id`).Error
		return nil
	})
}
