package migrations

import "gorm.io/gorm"

// Up014ChatAIMonitoring جداول اکوسیستم چت (اپراتور/جلسه/پیام/ویجت/بات/دانش) و مانیتورینگ
func Up014ChatAIMonitoring(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// chat_operators
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS chat_operators (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT,
                status VARCHAR(20) NOT NULL DEFAULT 'offline',
                max_concurrent_chats INT NOT NULL DEFAULT 5,
                current_chats INT NOT NULL DEFAULT 0,
                work_start_time VARCHAR(20) NOT NULL DEFAULT '09:00:00',
                work_end_time VARCHAR(20) NOT NULL DEFAULT '18:00:00',
                timezone VARCHAR(50) NOT NULL DEFAULT 'Asia/Tehran',
                auto_accept BOOLEAN NOT NULL DEFAULT FALSE,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		// chat_widgets
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS chat_widgets (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(255),
                description TEXT,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                welcome_message TEXT,
                offline_message TEXT,
                theme VARCHAR(50) NOT NULL DEFAULT 'light',
                position VARCHAR(50) NOT NULL DEFAULT 'bottom-right',
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		// chat_sessions
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS chat_sessions (
                id BIGSERIAL PRIMARY KEY,
                session_id VARCHAR(255) UNIQUE,
                customer_name VARCHAR(255),
                customer_phone VARCHAR(50),
                customer_email VARCHAR(255),
                status VARCHAR(20) NOT NULL DEFAULT 'waiting',
                operator_id BIGINT,
                ai_enabled BOOLEAN NOT NULL DEFAULT FALSE,
                widget_id BIGINT,
                ip_address VARCHAR(45),
                user_agent VARCHAR(500),
                started_at TIMESTAMPTZ,
                ended_at TIMESTAMPTZ,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		// chat_messages
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS chat_messages (
                id BIGSERIAL PRIMARY KEY,
                session_id BIGINT NOT NULL,
                sender_type VARCHAR(20) NOT NULL,
                sender_id BIGINT,
                message TEXT NOT NULL,
                message_type VARCHAR(50) NOT NULL DEFAULT 'text',
                file_url VARCHAR(255),
                file_size BIGINT,
                file_type VARCHAR(50),
                is_read BOOLEAN NOT NULL DEFAULT FALSE,
                read_at TIMESTAMPTZ,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		// chat_ai_bots
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS chat_ai_bots (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(255),
                description TEXT,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                openai_key VARCHAR(255),
                model VARCHAR(100) NOT NULL DEFAULT 'gpt-3.5-turbo',
                max_tokens INT NOT NULL DEFAULT 1000,
                temperature DOUBLE PRECISION NOT NULL DEFAULT 0.7,
                system_prompt TEXT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		// chat_knowledge_bases
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS chat_knowledge_bases (
                id BIGSERIAL PRIMARY KEY,
                title VARCHAR(255),
                content TEXT,
                category VARCHAR(100),
                tags VARCHAR(255),
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		// monitoring tables
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS chat_performance_metrics (
                id BIGSERIAL PRIMARY KEY,
                active_connections INT,
                avg_response_time DOUBLE PRECISION,
                messages_per_second DOUBLE PRECISION,
                error_rate DOUBLE PRECISION,
                cpu_usage DOUBLE PRECISION,
                memory_usage DOUBLE PRECISION,
                disk_usage DOUBLE PRECISION,
                network_usage DOUBLE PRECISION,
                performance_score DOUBLE PRECISION,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS chat_alerts (
                id BIGSERIAL PRIMARY KEY,
                level VARCHAR(20),
                message TEXT,
                category VARCHAR(50),
                is_read BOOLEAN NOT NULL DEFAULT FALSE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS chat_system_health (
                id BIGSERIAL PRIMARY KEY,
                status VARCHAR(20),
                uptime VARCHAR(50),
                last_check TIMESTAMPTZ,
                total_connections INT,
                active_sessions INT,
                total_messages INT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		// سازگاری با کوئری‌های GORM که از soft delete استفاده می‌کنند
		_ = tx.Exec(`ALTER TABLE chat_system_health ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ`)
		return nil
	})
}
