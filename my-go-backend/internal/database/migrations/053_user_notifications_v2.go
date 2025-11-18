package migrations

import "gorm.io/gorm"

// Up053UserNotificationsV2 تایید ساختار و ایندکس‌ها
func Up053UserNotificationsV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_user_notifications_user ON user_notifications(user_id)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_user_notifications_status ON user_notifications(status)")
		return nil
	})
}
