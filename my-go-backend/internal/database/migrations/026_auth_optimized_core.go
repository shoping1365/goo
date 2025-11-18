package migrations

import "gorm.io/gorm"

// Up026AuthOptimizedCore پوشش بخش‌هایی از 39 (اگر 010 و 017 کافی نبودند)
func Up026AuthOptimizedCore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اطمینان از ایندکس‌های مهم OTP و auth_events
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_otp_codes_mobile_purpose ON otp_codes(mobile, purpose)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_otp_codes_expires_used ON otp_codes(expires_at, is_used)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_auth_events_type_status ON auth_events(event_type, status)")
		// اگر ستون metadata از نوع JSON باشد، ابتدا به JSONB تبدیل می‌کنیم تا GIN قابل استفاده باشد
		_ = tx.Exec(`DO $$
        BEGIN
            IF EXISTS (
                SELECT 1 FROM information_schema.columns 
                WHERE table_name = 'auth_events' AND column_name = 'metadata' AND udt_name = 'json'
            ) THEN
                ALTER TABLE auth_events ALTER COLUMN metadata TYPE JSONB USING metadata::jsonb;
            END IF;
        END$$;`)
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_auth_events_metadata ON auth_events USING GIN (metadata)")
		return nil
	})
}
