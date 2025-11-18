package migrations

import "gorm.io/gorm"

// Up126OrdersUserAgent افزودن ستون user_agent به جدول orders
func Up126OrdersUserAgent(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن ستون user_agent
		if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS user_agent TEXT DEFAULT NULL`).Error; err != nil {
			return err
		}

		// ایجاد ایندکس برای user_agent (اختیاری - برای جستجو)
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_user_agent ON orders(user_agent)`).Error; err != nil {
			return err
		}

		return nil
	})
}
