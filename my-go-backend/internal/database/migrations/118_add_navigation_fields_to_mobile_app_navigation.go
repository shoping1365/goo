package migrations

import (
	"gorm.io/gorm"
)

// AddNavigationFieldsToMobileAppNavigation اضافه کردن فیلدهای جدید به جدول mobile_app_navigations
func AddNavigationFieldsToMobileAppNavigation(db *gorm.DB) error {
	// اضافه کردن فیلدهای جدید
	if err := db.Exec(`
		ALTER TABLE mobile_app_navigations 
		ADD COLUMN background_color VARCHAR(50) DEFAULT '#f8fafc',
		ADD COLUMN text_color VARCHAR(50) DEFAULT '#000000',
		ADD COLUMN navigation_items TEXT
	`).Error; err != nil {
		return err
	}

	return nil
}
