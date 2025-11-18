package migrations

import "gorm.io/gorm"

// Up110FixOtpCodesUpdatedAt اضافه کردن فیلد updated_at به جدول otp_codes
func Up110FixOtpCodesUpdatedAt(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن فیلد updated_at به جدول otp_codes
		if err := tx.Exec(`
			ALTER TABLE otp_codes 
			ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
		`).Error; err != nil {
			return err
		}

		// اضافه کردن فیلد updated_at به جدول login_attempts
		if err := tx.Exec(`
			ALTER TABLE login_attempts 
			ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
		`).Error; err != nil {
			return err
		}

		// اضافه کردن فیلد updated_at به جدول auth_events
		if err := tx.Exec(`
			ALTER TABLE auth_events 
			ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
		`).Error; err != nil {
			return err
		}

		return nil
	})
}
