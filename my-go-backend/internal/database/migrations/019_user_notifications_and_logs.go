package migrations

import "gorm.io/gorm"

// Up019UserNotificationsAndLogs ایجاد جداول user_notifications و notification_logs
func Up019UserNotificationsAndLogs(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS user_notifications (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT NOT NULL,
                channel VARCHAR(20) NOT NULL,
                source VARCHAR(20) NOT NULL,
                template_id BIGINT,
                title VARCHAR(200),
                body TEXT,
                payload JSONB NOT NULL DEFAULT '{}'::jsonb,
                status VARCHAR(20) NOT NULL DEFAULT 'sent',
                error_code VARCHAR(50),
                error_msg VARCHAR(500),
                sent_at TIMESTAMPTZ,
                read_at TIMESTAMPTZ,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_user_notifications_user ON user_notifications(user_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_user_notifications_status ON user_notifications(status)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS notification_logs (
                id BIGSERIAL PRIMARY KEY,
                channel VARCHAR(20),
                type VARCHAR(20),
                notification BIGINT,
                product_id BIGINT,
                user_id BIGINT,
                status VARCHAR(20),
                error_code VARCHAR(50),
                error_message VARCHAR(500),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		return nil
	})
}
