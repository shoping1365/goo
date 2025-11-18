package migrations

import "gorm.io/gorm"

// Up051UserBankAccountsV2 تایید ایندکس‌های حساب بانکی
func Up051UserBankAccountsV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_user_bank_accounts_user ON user_bank_accounts(user_id)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_user_bank_accounts_default ON user_bank_accounts(is_default)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_user_bank_accounts_iban ON user_bank_accounts(iban)")
		return nil
	})
}
