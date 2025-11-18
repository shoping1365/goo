package migrations

import "gorm.io/gorm"

// Up129AddCreatedAtToProductWarehouseStocks اضافه کردن ستون created_at به جدول product_warehouse_stocks
func Up129AddCreatedAtToProductWarehouseStocks(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن ستون created_at
		if err := tx.Exec(`
			ALTER TABLE product_warehouse_stocks 
			ADD COLUMN IF NOT EXISTS created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		`).Error; err != nil {
			return err
		}

		// به‌روزرسانی رکوردهای موجود با زمان ایجاد پیش‌فرض
		if err := tx.Exec(`
			UPDATE product_warehouse_stocks 
			SET created_at = NOW() 
			WHERE created_at IS NULL
		`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down129AddCreatedAtToProductWarehouseStocks حذف ستون created_at (در صورت نیاز)
func Down129AddCreatedAtToProductWarehouseStocks(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
			ALTER TABLE product_warehouse_stocks 
			DROP COLUMN IF EXISTS created_at
		`).Error; err != nil {
			return err
		}
		return nil
	})
}
