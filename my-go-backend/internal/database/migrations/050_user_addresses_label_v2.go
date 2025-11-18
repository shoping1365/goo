package migrations

import "gorm.io/gorm"

// Up050UserAddressesLabelV2 اطمینان مجدد از ستون label (اگر در 021 نبود)
func Up050UserAddressesLabelV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("ALTER TABLE user_addresses ADD COLUMN IF NOT EXISTS label VARCHAR(100)")
		return nil
	})
}
