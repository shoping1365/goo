package migrations

import "gorm.io/gorm"

// Up011Notifications اعلان‌های موجودی و تخفیف + لاگ اعلان‌ها
func Up011Notifications(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS stock_notifications (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT NOT NULL,
                product_id BIGINT NOT NULL,
                email VARCHAR(255) NOT NULL,
                phone VARCHAR(20),
                status VARCHAR(20) NOT NULL DEFAULT 'pending',
                sent_at TIMESTAMPTZ,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS discount_notifications (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT NOT NULL,
                product_id BIGINT NOT NULL,
                email VARCHAR(255) NOT NULL,
                phone VARCHAR(20),
                threshold_type VARCHAR(20) NOT NULL DEFAULT 'any',
                threshold_value NUMERIC(12,2) NOT NULL DEFAULT 0,
                status VARCHAR(20) NOT NULL DEFAULT 'pending',
                sent_at TIMESTAMPTZ,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
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
