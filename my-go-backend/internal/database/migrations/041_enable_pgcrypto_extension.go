package migrations

import "gorm.io/gorm"

// Up041EnablePgcrypto اطمینان مجدد از فعال بودن اکستنشن pgcrypto (در صورت نیاز خارج از 017)
func Up041EnablePgcrypto(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// pgcrypto extension is optional
		_ = tx.Exec("CREATE EXTENSION IF NOT EXISTS pgcrypto")
		return nil
	})
}
