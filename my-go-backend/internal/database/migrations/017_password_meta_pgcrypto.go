package migrations

import "gorm.io/gorm"

// Up017PasswordMetaAndPgcrypto برگردان 41 و 43 به طراحی جدید
func Up017PasswordMetaAndPgcrypto(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// Enable pgcrypto (optional)
		_ = tx.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto`)
		// users.password_algo/password_cost
		if err := tx.Exec(`ALTER TABLE users ADD COLUMN IF NOT EXISTS password_algo VARCHAR(20) NOT NULL DEFAULT 'bcrypt'`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE users ADD COLUMN IF NOT EXISTS password_cost INTEGER NOT NULL DEFAULT 14`).Error; err != nil {
			return err
		}
		return nil
	})
}
