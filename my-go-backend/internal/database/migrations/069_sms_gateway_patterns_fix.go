package migrations

import "gorm.io/gorm"

// Up069SMSGatewayPatternsFix ایجاد جدول sms_gateway_patterns در صورت نبود + مهاجرت داده اختیاری از sms_patterns
func Up069SMSGatewayPatternsFix(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS sms_gateway_patterns (
                id BIGSERIAL PRIMARY KEY,
                gateway_id BIGINT NOT NULL,
                pattern_code VARCHAR(100) NOT NULL,
                description VARCHAR(500),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_sgp_gateway FOREIGN KEY(gateway_id) REFERENCES sms_gateways(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}

		// ایندکس‌ها
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_sgp_gateway ON sms_gateway_patterns(gateway_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_sgp_unique ON sms_gateway_patterns(gateway_id, pattern_code)`).Error; err != nil {
			return err
		}

		// مهاجرت داده از جدول قدیمی sms_patterns در صورت وجود
		_ = tx.Exec(`
            INSERT INTO sms_gateway_patterns (gateway_id, pattern_code, description, created_at, updated_at)
            SELECT sp.gateway_id, sp.pattern_code, sp.description,
                   COALESCE(sp.created_at, NOW()), COALESCE(sp.updated_at, NOW())
            FROM sms_patterns sp
            WHERE NOT EXISTS (
                SELECT 1 FROM sms_gateway_patterns sgp
                WHERE sgp.gateway_id = sp.gateway_id AND sgp.pattern_code = sp.pattern_code
            );
        `)
		return nil
	})
}


