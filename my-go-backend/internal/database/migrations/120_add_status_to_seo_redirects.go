package migrations

import (
	"gorm.io/gorm"
)

// Up120AddStatusToSEORedirects اضافه کردن فیلد status به جدول seo_redirects
func Up120AddStatusToSEORedirects(db *gorm.DB) error {
	return db.Exec(`
		ALTER TABLE seo_redirects 
		ADD COLUMN status VARCHAR(20) DEFAULT 'active' NOT NULL;
	`).Error
}

// Down120AddStatusToSEORedirects حذف فیلد status از جدول seo_redirects
func Down120AddStatusToSEORedirects(db *gorm.DB) error {
	return db.Exec(`
		ALTER TABLE seo_redirects 
		DROP COLUMN status;
	`).Error
}

