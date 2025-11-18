package migrations

import "gorm.io/gorm"

// Up051OrdersSessionsIndexes ایندکس‌های تکمیلی برای سفارش‌ها و سشن‌ها
func Up051OrdersSessionsIndexes(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// orders – ایندکس‌های partial/ترکیبی
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_user_paid ON orders(user_id) WHERE is_paid = true`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status)`).Error; err != nil {
			return err
		}

		// sessions – تکمیلی
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_sessions_user_current_expires ON sessions(user_id, is_active, expires_at)`).Error; err != nil {
			return err
		}
		return nil
	})
}
