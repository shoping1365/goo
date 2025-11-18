package migrations

import "gorm.io/gorm"

// Up033AdminOperationalTables ایجاد جدول admin_settings (اگر قبلا 015 ایجاد نکرده باشد)
func Up033AdminOperationalTables(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS admin_settings (
            id BIGSERIAL PRIMARY KEY,
            user_id BIGINT,
            key VARCHAR(100) NOT NULL,
            value TEXT,
            created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
            updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
        );`)
		return nil
	})
}
