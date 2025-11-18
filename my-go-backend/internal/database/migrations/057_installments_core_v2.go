package migrations

import "gorm.io/gorm"

// Up057InstallmentsCoreV2 تایید ایندکس‌های اقساط و ستون‌های محصولات (در صورت کمبود)
func Up057InstallmentsCoreV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اطمینان از وجود جداول اقساط
		_ = tx.Exec(`
            CREATE TABLE IF NOT EXISTS installment_plans (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100),
                description TEXT,
                total_amount NUMERIC(18,2),
                num_installments INT,
                interest_rate NUMERIC(5,2) DEFAULT 0,
                created_at TIMESTAMPTZ DEFAULT NOW(),
                updated_at TIMESTAMPTZ DEFAULT NOW()
            );
        `)
		_ = tx.Exec(`
            CREATE TABLE IF NOT EXISTS installment_schedules (
                id BIGSERIAL PRIMARY KEY,
                installment_id BIGINT,
                due_date TIMESTAMPTZ,
                amount NUMERIC(18,2),
                status VARCHAR(20) DEFAULT 'pending',
                created_at TIMESTAMPTZ DEFAULT NOW(),
                updated_at TIMESTAMPTZ DEFAULT NOW()
            );
        `)
		_ = tx.Exec(`
            CREATE TABLE IF NOT EXISTS installment_payments (
                id BIGSERIAL PRIMARY KEY,
                schedule_id BIGINT,
                amount NUMERIC(18,2),
                status VARCHAR(20) DEFAULT 'success',
                created_at TIMESTAMPTZ DEFAULT NOW()
            );
        `)

		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_installments_status_created ON installment_payments(status, created_at DESC)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_installment_schedules_due ON installment_schedules(installment_id, due_date)")
		_ = tx.Exec("ALTER TABLE products ADD COLUMN IF NOT EXISTS installment_enabled BOOLEAN DEFAULT FALSE")
		_ = tx.Exec("ALTER TABLE products ADD COLUMN IF NOT EXISTS installment_price NUMERIC(18,2) DEFAULT 0")
		_ = tx.Exec("ALTER TABLE products ADD COLUMN IF NOT EXISTS installment_options JSONB DEFAULT '[]'")
		return nil
	})
}
