package migrations

import "gorm.io/gorm"

// Up005OrdersCore ایجاد جداول سفارش و اقلام وابسته با ایندکس‌های کلیدی
func Up005OrdersCore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// orders
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS orders (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT,
                shipping_address_id BIGINT NOT NULL,
                payment_method VARCHAR(20) NOT NULL DEFAULT 'online',
                status VARCHAR(30) NOT NULL DEFAULT 'processing_queue',
                payment_status VARCHAR(20) NOT NULL DEFAULT 'awaiting_payment',
                is_paid BOOLEAN NOT NULL DEFAULT FALSE,
                total_amount NUMERIC(12,2),
                final_amount NUMERIC(12,2),
                tracking_code VARCHAR(16) UNIQUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_orders_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_user_status_created ON orders(user_id, status, created_at DESC)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_payment_status_paid ON orders(payment_status, is_paid)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_created_desc ON orders(created_at DESC)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_paid_created ON orders(created_at DESC) WHERE is_paid = true`).Error; err != nil {
			return err
		}

		// order_items
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS order_items (
                id BIGSERIAL PRIMARY KEY,
                order_id BIGINT NOT NULL,
                product_id BIGINT NOT NULL,
                quantity INT NOT NULL,
                unit_price NUMERIC(12,2) NOT NULL,
                final_price NUMERIC(12,2) NOT NULL,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_order_items_order FOREIGN KEY(order_id) REFERENCES orders(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_order_items_order ON order_items(order_id)`).Error; err != nil {
			return err
		}

		return nil
	})
}
