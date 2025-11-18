package migrations

import "gorm.io/gorm"

// Up025FraudCoreAndIndexes پوشش Up61/62 (fraud core + indexes)
func Up025FraudCoreAndIndexes(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// core tables
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS fraud_cases (
                id BIGSERIAL PRIMARY KEY,
                order_id BIGINT,
                user_id BIGINT,
                risk_score INT NOT NULL DEFAULT 0,
                risk_level VARCHAR(10),
                status VARCHAR(15),
                summary VARCHAR(500),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                closed_at TIMESTAMPTZ,
                metadata JSONB NOT NULL DEFAULT '{}'::jsonb
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS fraud_events (
                id BIGSERIAL PRIMARY KEY,
                case_id BIGINT NOT NULL,
                event_type VARCHAR(50),
                weight INT NOT NULL DEFAULT 0,
                impact INT NOT NULL DEFAULT 0,
                details JSONB NOT NULL DEFAULT '{}'::jsonb,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS fraud_scores (
                id BIGSERIAL PRIMARY KEY,
                case_id BIGINT NOT NULL,
                score_before INT NOT NULL DEFAULT 0,
                score_after INT NOT NULL DEFAULT 0,
                reason JSONB NOT NULL DEFAULT '{}'::jsonb,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS fraud_rules (
                id BIGSERIAL PRIMARY KEY,
                key VARCHAR(100) UNIQUE,
                description VARCHAR(500),
                weight INT NOT NULL DEFAULT 0,
                enabled BOOLEAN NOT NULL DEFAULT TRUE,
                params JSONB NOT NULL DEFAULT '{}'::jsonb,
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS fraud_list (
                id BIGSERIAL PRIMARY KEY,
                kind VARCHAR(30),
                value_hash VARCHAR(128),
                note VARCHAR(300),
                expires_at TIMESTAMPTZ,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS fraud_actions_log (
                id BIGSERIAL PRIMARY KEY,
                case_id BIGINT NOT NULL,
                action VARCHAR(20),
                actor_id BIGINT,
                note VARCHAR(500),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// indexes
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_fraud_cases_user_created ON fraud_cases(user_id, created_at DESC)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_fraud_cases_level_created ON fraud_cases(risk_level, created_at DESC)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_fraud_cases_status_created ON fraud_cases(status, created_at DESC)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_fraud_events_case_type_created ON fraud_events(case_id, event_type, created_at DESC)")
		_ = tx.Exec("CREATE UNIQUE INDEX IF NOT EXISTS ux_fraud_list_kind_value ON fraud_list(kind, value_hash)")
		return nil
	})
}
