package migrations

import "gorm.io/gorm"

// Up006UserAddresses ایجاد جدول آدرس‌های کاربر (ساده‌شده و بهینه)
func Up006UserAddresses(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS user_addresses (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT NOT NULL,
                label VARCHAR(100),
                street TEXT,
                postal_code VARCHAR(20),
                recipient_name VARCHAR(100),
                recipient_mobile VARCHAR(20),
                phone VARCHAR(20),
                province VARCHAR(100),
                city VARCHAR(100),
                province_id BIGINT,
                city_id BIGINT,
                is_default BOOLEAN NOT NULL DEFAULT FALSE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ,
                CONSTRAINT fk_user_addresses_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_user_addresses_user_default ON user_addresses(user_id, is_default)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_user_addresses_city_postal ON user_addresses(city, postal_code)`).Error; err != nil {
			return err
		}
		return nil
	})
}
