package migrations

import (
	"gorm.io/gorm"
)

// Up121AddURLToProducts اضافه کردن فیلد url به جدول products
func Up121AddURLToProducts(db *gorm.DB) error {
	return db.Exec(`
		ALTER TABLE products 
		ADD COLUMN url VARCHAR(500) DEFAULT '';
	`).Error
}

// Down121AddURLToProducts حذف فیلد url از جدول products
func Down121AddURLToProducts(db *gorm.DB) error {
	return db.Exec(`
		ALTER TABLE products 
		DROP COLUMN url;
	`).Error
}
