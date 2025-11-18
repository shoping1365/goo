package migrations

import "gorm.io/gorm"

// Up048SMSPatternsExtend توسعه ستون‌ها و ایندکس sms_patterns
func Up048SMSPatternsExtend(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("ALTER TABLE sms_patterns ADD COLUMN IF NOT EXISTS type VARCHAR(20) DEFAULT 'sms'")
		_ = tx.Exec("ALTER TABLE sms_patterns ADD COLUMN IF NOT EXISTS scope VARCHAR(20) DEFAULT 'customer'")
		_ = tx.Exec("ALTER TABLE sms_patterns ADD COLUMN IF NOT EXISTS feature VARCHAR(50)")
		_ = tx.Exec("ALTER TABLE sms_patterns ADD COLUMN IF NOT EXISTS description TEXT")
		_ = tx.Exec("ALTER TABLE sms_patterns ADD COLUMN IF NOT EXISTS variables TEXT")
		_ = tx.Exec("ALTER TABLE sms_patterns ADD COLUMN IF NOT EXISTS template TEXT")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_sms_patterns_feature_scope_gateway_status ON sms_patterns (feature, scope, gateway_id, status)")
		return nil
	})
}
