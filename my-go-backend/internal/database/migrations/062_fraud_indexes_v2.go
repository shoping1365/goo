package migrations

import "gorm.io/gorm"

// Up062FraudIndexesV2 فراخوانی دوباره ایندکس‌های fraud (کافی‌سازی)
func Up062FraudIndexesV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_fraud_cases_user_created ON fraud_cases(user_id, created_at DESC)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_fraud_cases_level_created ON fraud_cases(risk_level, created_at DESC)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_fraud_cases_status_created ON fraud_cases(status, created_at DESC)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_fraud_events_case_type_created ON fraud_events(case_id, event_type, created_at DESC)")
		_ = tx.Exec("CREATE UNIQUE INDEX IF NOT EXISTS ux_fraud_list_kind_value ON fraud_list(kind, value_hash)")
		return nil
	})
}
