package migrations

import (
	"gorm.io/gorm"
)

// Up001SchemaMigrationsLedger ایجاد جدول رهگیری نسخه‌های مایگریشن
// این جدول فقط یکبار ساخته می‌شود و اجرای دوباره تغییری ایجاد نمی‌کند.
func Up001SchemaMigrationsLedger(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// ایجاد جدول در صورت عدم وجود
		ddl := `
        CREATE TABLE IF NOT EXISTS schema_migrations (
            version      INT PRIMARY KEY,
            name         TEXT NOT NULL,
            applied_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
        );`
		if err := tx.Exec(ddl).Error; err != nil {
			return err
		}

		// ایندکس زمانی برای گزارش‌گیری
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_schema_migrations_applied_at ON schema_migrations(applied_at)`).Error; err != nil {
			return err
		}
		return nil
	})
}
