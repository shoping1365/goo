package migrations

import "gorm.io/gorm"

// Up037OnlineChatSystem اطمینان از وجود هسته جداول چت آنلاین (در صورت جدا بودن از 014)
func Up037OnlineChatSystem(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// Chat Operators table
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS chat_operators (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			status VARCHAR(20) DEFAULT 'offline',
			max_concurrent_chats INT DEFAULT 5,
			current_chats INT DEFAULT 0,
			work_start_time VARCHAR(20) DEFAULT '09:00:00',
			work_end_time VARCHAR(20) DEFAULT '18:00:00',
			timezone VARCHAR(50) DEFAULT 'Asia/Tehran',
			auto_accept BOOLEAN DEFAULT FALSE,
			is_active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW(),
			deleted_at TIMESTAMPTZ
		)`)

		// Chat Sessions table
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS chat_sessions (
			id BIGSERIAL PRIMARY KEY,
			session_id VARCHAR(255) UNIQUE,
			customer_name TEXT,
			customer_phone TEXT,
			customer_email TEXT,
			status VARCHAR(20) DEFAULT 'waiting',
			operator_id BIGINT,
			ai_enabled BOOLEAN DEFAULT FALSE,
			widget_id BIGINT,
			ip_address VARCHAR(45),
			user_agent TEXT,
			started_at TIMESTAMPTZ DEFAULT NOW(),
			ended_at TIMESTAMPTZ,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW(),
			deleted_at TIMESTAMPTZ
		)`)

		// Chat Messages table
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS chat_messages (
			id BIGSERIAL PRIMARY KEY,
			session_id BIGINT NOT NULL,
			sender_type VARCHAR(20),
			sender_id BIGINT,
			message TEXT,
			message_type VARCHAR(20) DEFAULT 'text',
			file_url TEXT,
			file_size BIGINT,
			file_type VARCHAR(100),
			is_read BOOLEAN DEFAULT FALSE,
			read_at TIMESTAMPTZ,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW(),
			deleted_at TIMESTAMPTZ
		)`)

		// Chat Widgets table
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS chat_widgets (
			id BIGSERIAL PRIMARY KEY,
			name VARCHAR(255),
			description TEXT,
			is_active BOOLEAN DEFAULT TRUE,
			welcome_message TEXT,
			offline_message TEXT,
			theme VARCHAR(20) DEFAULT 'light',
			position VARCHAR(20) DEFAULT 'bottom-right',
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW(),
			deleted_at TIMESTAMPTZ
		)`)

		// Chat AI Bots table
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS chat_ai_bots (
			id BIGSERIAL PRIMARY KEY,
			name VARCHAR(255),
			description TEXT,
			is_active BOOLEAN DEFAULT TRUE,
			openai_key TEXT,
			model VARCHAR(100) DEFAULT 'gpt-3.5-turbo',
			max_tokens INT DEFAULT 1000,
			temperature DECIMAL(3,2) DEFAULT 0.7,
			system_prompt TEXT,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW(),
			deleted_at TIMESTAMPTZ
		)`)

		// Chat Knowledge Base table
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS chat_knowledge_base (
			id BIGSERIAL PRIMARY KEY,
			title VARCHAR(500),
			content TEXT,
			category VARCHAR(100),
			tags TEXT,
			is_active BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW(),
			deleted_at TIMESTAMPTZ
		)`)

		// Chat Statistics table
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS chat_statistics (
			id BIGSERIAL PRIMARY KEY,
			date DATE NOT NULL,
			total_sessions INT DEFAULT 0,
			completed_sessions INT DEFAULT 0,
			total_messages INT DEFAULT 0,
			avg_response_time DECIMAL(10,2) DEFAULT 0,
			operator_id BIGINT,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		)`)

		return nil
	})
}
