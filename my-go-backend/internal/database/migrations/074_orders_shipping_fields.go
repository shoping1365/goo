package migrations

import "gorm.io/gorm"

// Up074OrdersShippingFields افزودن فیلدهای ساختاریافته ارسال به جدول orders و محدودسازی مقادیر مجاز
func Up074OrdersShippingFields(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// ستون روش ارسال
		if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS shipping_method VARCHAR(30) NOT NULL DEFAULT 'post'`).Error; err != nil {
			return err
		}
		// ستون هزینه ارسال
		if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS shipping_cost NUMERIC(12,2) NOT NULL DEFAULT 0`).Error; err != nil {
			return err
		}
		// ستون IP مشتری
		if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS customer_ip VARCHAR(45)`).Error; err != nil {
			return err
		}

		// محدودسازی مقادیر مجاز روش ارسال (post, tipax, chapar, freight, pickup)
		// توجه: ADD CONSTRAINT IF NOT EXISTS در Postgres وجود ندارد؛ از بلاک DO استفاده می‌کنیم
		if err := tx.Exec(`
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'ck_orders_shipping_method_values'
    ) THEN
        ALTER TABLE orders
        ADD CONSTRAINT ck_orders_shipping_method_values
        CHECK (shipping_method IN ('post','tipax','chapar','freight','pickup'));
    END IF;
END $$;`).Error; err != nil {
			return err
		}

		// ایندکس کم‌هزینه روی روش ارسال برای گزارش‌ها
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_shipping_method ON orders(shipping_method)`).Error; err != nil {
			return err
		}
		return nil
	})
}


