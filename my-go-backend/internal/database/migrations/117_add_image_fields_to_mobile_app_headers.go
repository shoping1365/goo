package migrations

import (
	"gorm.io/gorm"
)

// AddImageFieldsToMobileAppHeaders migration برای اضافه کردن فیلدهای عکس به جدول mobile_app_headers
func AddImageFieldsToMobileAppHeaders(db *gorm.DB) error {
	// اضافه کردن فیلدهای عکس بالا و پایین
	if err := db.Exec(`
		ALTER TABLE mobile_app_headers 
		ADD COLUMN top_image_url VARCHAR(500) DEFAULT '',
		ADD COLUMN top_image_alt VARCHAR(255) DEFAULT '',
		ADD COLUMN bottom_image_url VARCHAR(500) DEFAULT '',
		ADD COLUMN bottom_image_alt VARCHAR(255) DEFAULT ''
	`).Error; err != nil {
		return err
	}

	return nil
}
