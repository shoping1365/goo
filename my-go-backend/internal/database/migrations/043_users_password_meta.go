package migrations

import "gorm.io/gorm"

// Up043UsersPasswordMeta افزودن ستون‌های الگوریتم و هزینه رمز عبور (idempotent)
func Up043UsersPasswordMeta(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("ALTER TABLE users ADD COLUMN IF NOT EXISTS password_algo VARCHAR(20) NOT NULL DEFAULT 'bcrypt'")
		_ = tx.Exec("ALTER TABLE users ADD COLUMN IF NOT EXISTS password_cost INTEGER NOT NULL DEFAULT 14")
		return nil
	})
}
