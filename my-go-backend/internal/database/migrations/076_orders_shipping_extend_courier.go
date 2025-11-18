package migrations

import "gorm.io/gorm"

// Up076OrdersShippingExtendCourier افزودن مقدار 'courier' به مقادیر مجاز روش ارسال
func Up076OrdersShippingExtendCourier(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف محدودیت قبلی (در صورت وجود)
		if err := tx.Exec(`ALTER TABLE orders DROP CONSTRAINT IF EXISTS ck_orders_shipping_method_values`).Error; err != nil {
			return err
		}
		// ایجاد محدودیت جدید با مقدار 'courier' اضافه
		if err := tx.Exec(`
            ALTER TABLE orders
            ADD CONSTRAINT ck_orders_shipping_method_values
            CHECK (shipping_method IN ('post','tipax','chapar','freight','pickup','courier'))
        `).Error; err != nil {
			return err
		}
		return nil
	})
}


