package migrations

import "gorm.io/gorm"

// Up012SMSPayment جداول پیامک و درگاه پرداخت/تراکنش‌ها/لاگ‌ها
func Up012SMSPayment(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// sms_gateways
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS sms_gateways (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100),
                type VARCHAR(50) NOT NULL,
                sender_number VARCHAR(30),
                api_url VARCHAR(255),
                api_key VARCHAR(255),
                username VARCHAR(100),
                password VARCHAR(100),
                pattern_based BOOLEAN NOT NULL DEFAULT FALSE,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                priority INT NOT NULL DEFAULT 1,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// sms_gateway_patterns (سازگار با ریپازیتوری)
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
		// جدول قبلی sms_patterns را نیز در صورت وجود نگه می‌داریم (سازگاری)
		_ = tx.Exec(`
            CREATE TABLE IF NOT EXISTS sms_patterns (
                id BIGSERIAL PRIMARY KEY,
                type VARCHAR(20) NOT NULL DEFAULT 'sms',
                scope VARCHAR(20) NOT NULL DEFAULT 'customer',
                feature VARCHAR(50),
                name VARCHAR(255) NOT NULL,
                pattern_code VARCHAR(100) NOT NULL UNIQUE,
                description VARCHAR(500),
                variables TEXT,
                template TEXT NOT NULL,
                gateway_id BIGINT NOT NULL,
                status VARCHAR(20) NOT NULL DEFAULT 'active',
                usage_count INT NOT NULL DEFAULT 0,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `)

		// payment_gateways
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS payment_gateways (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL,
                english_name VARCHAR(100) NOT NULL,
                type VARCHAR(50) NOT NULL,
                status VARCHAR(20) NOT NULL DEFAULT 'inactive',
                icon VARCHAR(10),
                color VARCHAR(20) NOT NULL DEFAULT 'bg-blue-500',
                fee DOUBLE PRECISION NOT NULL DEFAULT 0,
                min_amount BIGINT NOT NULL DEFAULT 0,
                max_amount BIGINT NOT NULL DEFAULT 0,
                priority INT NOT NULL DEFAULT 1,
                is_test_mode BOOLEAN NOT NULL DEFAULT TRUE,
                description VARCHAR(500),
                instructions TEXT,
                terminal_id VARCHAR(100),
                merchant_id VARCHAR(100),
                today_transactions BIGINT NOT NULL DEFAULT 0,
                today_revenue BIGINT NOT NULL DEFAULT 0,
                total_transactions BIGINT NOT NULL DEFAULT 0,
                total_revenue BIGINT NOT NULL DEFAULT 0,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// payment_transactions
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS payment_transactions (
                id BIGSERIAL PRIMARY KEY,
                gateway_id BIGINT NOT NULL,
                order_id VARCHAR(100) NOT NULL,
                transaction_id VARCHAR(100) UNIQUE,
                amount BIGINT NOT NULL,
                status VARCHAR(20) NOT NULL,
                payment_method VARCHAR(50),
                description VARCHAR(500),
                gateway_token VARCHAR(255),
                gateway_type VARCHAR(50),
                card_number VARCHAR(20),
                card_holder VARCHAR(100),
                gateway_response TEXT,
                error_message VARCHAR(500),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_payment_tx_gateway FOREIGN KEY(gateway_id) REFERENCES payment_gateways(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}

		// payment_gateway_logs
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS payment_gateway_logs (
                id BIGSERIAL PRIMARY KEY,
                gateway_id BIGINT NOT NULL,
                level VARCHAR(20) NOT NULL,
                message VARCHAR(500) NOT NULL,
                details TEXT,
                ip_address VARCHAR(45),
                user_agent VARCHAR(500),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_payment_logs_gateway FOREIGN KEY(gateway_id) REFERENCES payment_gateways(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		return nil
	})
}
