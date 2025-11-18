package migrations

import "gorm.io/gorm"

// Up013Traffic جداول ترافیک و IPهای مسدود
func Up013Traffic(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS traffic_logs (
                id BIGSERIAL PRIMARY KEY,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                user_id BIGINT,
                session_id VARCHAR(255),
                ip_address VARCHAR(45),
                request_method VARCHAR(10) NOT NULL,
                request_path VARCHAR(1000) NOT NULL,
                status_code SMALLINT NOT NULL,
                response_time_ms INT NOT NULL DEFAULT 0,
                user_agent TEXT,
                browser VARCHAR(50),
                device_type VARCHAR(20),
                os VARCHAR(50),
                hostname VARCHAR(255),
                ad_type VARCHAR(20),
                city VARCHAR(100),
                country_code CHAR(2),
                is_suspicious BOOLEAN NOT NULL DEFAULT FALSE,
                meta JSONB NOT NULL DEFAULT '{}'::jsonb
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_traffic_logs_created ON traffic_logs(created_at)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_traffic_logs_path ON traffic_logs(request_path)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS traffic_sessions (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT,
                session_id VARCHAR(255) NOT NULL UNIQUE,
                ip_address VARCHAR(45) NOT NULL,
                user_agent TEXT,
                current_page VARCHAR(500),
                login_time TIMESTAMPTZ NOT NULL,
                last_seen TIMESTAMPTZ NOT NULL,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS blocked_ips (
                id BIGSERIAL PRIMARY KEY,
                ip_address VARCHAR(45) NOT NULL UNIQUE,
                reason TEXT,
                blocked_by VARCHAR(100),
                blocked_at TIMESTAMPTZ NOT NULL,
                expires_at TIMESTAMPTZ,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		return nil
	})
}
