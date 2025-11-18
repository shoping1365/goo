package migrations

import "gorm.io/gorm"

// Up075OrdersShippingTrackingCode افزودن ستون کد رهگیری ارسال به جدول orders
func Up075OrdersShippingTrackingCode(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS shipping_tracking_code VARCHAR(100)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_shipping_tracking_code ON orders(shipping_tracking_code)`).Error; err != nil {
			return err
		}
		return nil
	})
}
