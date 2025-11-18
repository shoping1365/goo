package migrations

import "gorm.io/gorm"

// Up078ProductReservationFlag افزودن فلگ allow_reservation به products
func Up078ProductReservationFlag(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// ستون allow_reservation برای کنترل فعال‌سازی رزرو در سطح محصول
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS allow_reservation BOOLEAN NOT NULL DEFAULT FALSE`).Error; err != nil {
			return err
		}
		return nil
	})
}

// Down078ProductReservationFlag حذف ستون اضافه شده
func Down078ProductReservationFlag(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec(`ALTER TABLE products DROP COLUMN IF EXISTS allow_reservation`).Error
		return nil
	})
}
