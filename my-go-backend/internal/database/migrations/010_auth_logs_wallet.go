package migrations

import "gorm.io/gorm"

// Up010AuthLogsWallet جداول OTP، لاگ لاگین، رویدادهای احراز هویت، نشست‌های کاربر، کیف پول و تراکنش‌ها
func Up010AuthLogsWallet(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// otp_codes
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS otp_codes (
                id BIGSERIAL PRIMARY KEY,
                mobile VARCHAR(20) NOT NULL,
                code VARCHAR(10) NOT NULL,
                purpose VARCHAR(50) NOT NULL,
                expires_at TIMESTAMPTZ NOT NULL,
                is_used BOOLEAN NOT NULL DEFAULT FALSE,
                used_at TIMESTAMPTZ,
                ip_address VARCHAR(45),
                user_agent VARCHAR(500),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// login_attempts
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS login_attempts (
                id BIGSERIAL PRIMARY KEY,
                mobile VARCHAR(20) NOT NULL,
                attempt_type VARCHAR(50) NOT NULL,
                is_successful BOOLEAN NOT NULL DEFAULT FALSE,
                failure_reason VARCHAR(255),
                attempt_ip VARCHAR(45),
                user_agent VARCHAR(500),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// user_sessions
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS user_sessions (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT NOT NULL,
                session_token VARCHAR(255) NOT NULL UNIQUE,
                refresh_token VARCHAR(255) NOT NULL UNIQUE,
                device_fingerprint VARCHAR(255),
                device_name VARCHAR(255),
                platform VARCHAR(100),
                browser VARCHAR(100),
                browser_version VARCHAR(50),
                os_name VARCHAR(100),
                os_version VARCHAR(50),
                ip_address VARCHAR(45),
                city VARCHAR(100),
                country VARCHAR(100),
                timezone VARCHAR(100),
                user_agent VARCHAR(500),
                device_info JSON,
                auth_method VARCHAR(50),
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                is_trusted BOOLEAN NOT NULL DEFAULT FALSE,
                is_current BOOLEAN NOT NULL DEFAULT FALSE,
                login_at TIMESTAMPTZ,
                last_activity TIMESTAMPTZ,
                expires_at TIMESTAMPTZ,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// auth_events
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS auth_events (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT,
                mobile VARCHAR(20),
                event_type VARCHAR(100) NOT NULL,
                status VARCHAR(20) NOT NULL,
                severity VARCHAR(20) NOT NULL,
                ip_address VARCHAR(45),
                user_agent VARCHAR(500),
                metadata JSON,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// user_wallets
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS user_wallets (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT NOT NULL UNIQUE,
                balance NUMERIC(18,2) NOT NULL DEFAULT 0,
                currency VARCHAR(10) NOT NULL DEFAULT 'IRR',
                status VARCHAR(20) NOT NULL DEFAULT 'active',
                last_transaction_at TIMESTAMPTZ,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// wallet_transactions
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS wallet_transactions (
                id BIGSERIAL PRIMARY KEY,
                wallet_id BIGINT NOT NULL,
                type VARCHAR(10) NOT NULL,
                method VARCHAR(20),
                amount NUMERIC(18,2) NOT NULL,
                related_type VARCHAR(50),
                related_id BIGINT,
                card_number_last4 VARCHAR(4),
                gateway VARCHAR(50),
                ip VARCHAR(45),
                user_agent VARCHAR(500),
                platform VARCHAR(50),
                browser VARCHAR(50),
                os_name VARCHAR(50),
                session_id VARCHAR(255),
                metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
                status VARCHAR(20) NOT NULL DEFAULT 'success',
                error_code VARCHAR(50),
                error_message VARCHAR(500),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_wallet_tx_wallet FOREIGN KEY(wallet_id) REFERENCES user_wallets(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		return nil
	})
}
