package migrations

import "gorm.io/gorm"

// Up018UserBankAccounts ایجاد جدول user_bank_accounts
func Up018UserBankAccounts(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		ddl := `
        CREATE TABLE IF NOT EXISTS user_bank_accounts (
            id BIGSERIAL PRIMARY KEY,
            user_id BIGINT NOT NULL,
            bank_name VARCHAR(100),
            card_number_last4 VARCHAR(4),
            card_token VARCHAR(255),
            iban VARCHAR(26),
            account_number VARCHAR(50),
            is_default BOOLEAN NOT NULL DEFAULT FALSE,
            verified_at TIMESTAMPTZ,
            status VARCHAR(20) NOT NULL DEFAULT 'active',
            metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
            created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
            updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
        );`
		if err := tx.Exec(ddl).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_user_bank_accounts_user ON user_bank_accounts(user_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_user_bank_accounts_default ON user_bank_accounts(is_default)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_user_bank_accounts_iban ON user_bank_accounts(iban)`).Error; err != nil {
			return err
		}
		return nil
	})
}
