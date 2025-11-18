package migrations

import "gorm.io/gorm"

// Up068CategoryQACore ایجاد جدول category_qas برای مدیریت پرسش/پاسخ دسته‌بندی‌ها (idempotent)
func Up068CategoryQACore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS category_qas (
                id BIGSERIAL PRIMARY KEY,
                category_id BIGINT NOT NULL,
                question TEXT NOT NULL,
                answer TEXT,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                sort_order INT NOT NULL DEFAULT 0,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_category_qas_category ON category_qas(category_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_category_qas_active ON category_qas(is_active)`).Error; err != nil {
			return err
		}
		return nil
	})
}


