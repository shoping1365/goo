package migrations

import "gorm.io/gorm"

// Up031CartOrderEnhancements تکمیل ستون‌ها و ایندکس‌های سفارش/سبد خرید (ایمن و idempotent)
func Up031CartOrderEnhancements(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اطمینان از وجود جداول سبد خرید قبل از ایجاد ایندکس‌ها
		// carts
		_ = tx.Exec(`
            CREATE TABLE IF NOT EXISTS carts (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT,
                session_id VARCHAR(255),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `)
		_ = tx.Exec(`CREATE INDEX IF NOT EXISTS idx_carts_user_id ON carts(user_id)`)

		// cart_items
		_ = tx.Exec(`
            CREATE TABLE IF NOT EXISTS cart_items (
                id BIGSERIAL PRIMARY KEY,
                cart_id BIGINT,
                product_id BIGINT,
                quantity INT NOT NULL DEFAULT 1,
                is_next_purchase BOOLEAN NOT NULL DEFAULT FALSE,
                unit_price DOUBLE PRECISION NOT NULL DEFAULT 0,
                discount_amount DOUBLE PRECISION NOT NULL DEFAULT 0,
                final_price DOUBLE PRECISION NOT NULL DEFAULT 0,
                added_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `)
		// افزودن ستون‌های موردنیاز در صورت نبود
		_ = tx.Exec(`ALTER TABLE cart_items ADD COLUMN IF NOT EXISTS cart_id BIGINT`)
		_ = tx.Exec(`ALTER TABLE cart_items ADD COLUMN IF NOT EXISTS product_id BIGINT`)
		_ = tx.Exec(`ALTER TABLE cart_items ADD COLUMN IF NOT EXISTS quantity INT NOT NULL DEFAULT 1`)
		_ = tx.Exec(`ALTER TABLE cart_items ADD COLUMN IF NOT EXISTS is_next_purchase BOOLEAN NOT NULL DEFAULT FALSE`)
		_ = tx.Exec(`ALTER TABLE cart_items ADD COLUMN IF NOT EXISTS unit_price DOUBLE PRECISION NOT NULL DEFAULT 0`)
		_ = tx.Exec(`ALTER TABLE cart_items ADD COLUMN IF NOT EXISTS discount_amount DOUBLE PRECISION NOT NULL DEFAULT 0`)
		_ = tx.Exec(`ALTER TABLE cart_items ADD COLUMN IF NOT EXISTS final_price DOUBLE PRECISION NOT NULL DEFAULT 0`)
		_ = tx.Exec(`ALTER TABLE cart_items ADD COLUMN IF NOT EXISTS added_at TIMESTAMPTZ NOT NULL DEFAULT NOW()`)
		_ = tx.Exec(`ALTER TABLE cart_items ADD COLUMN IF NOT EXISTS updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()`)
		_ = tx.Exec(`ALTER TABLE cart_items ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ`)

		// orders: اطمینان از وجود برخی ستون‌ها
		_ = tx.Exec("ALTER TABLE orders ADD COLUMN IF NOT EXISTS is_paid BOOLEAN NOT NULL DEFAULT FALSE")
		_ = tx.Exec("ALTER TABLE orders ADD COLUMN IF NOT EXISTS payment_status VARCHAR(20) NOT NULL DEFAULT 'awaiting_payment'")
		_ = tx.Exec("ALTER TABLE orders ADD COLUMN IF NOT EXISTS tracking_code VARCHAR(100)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_orders_payment_status_paid ON orders(payment_status, is_paid)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_orders_user_status_created ON orders(user_id, status, created_at DESC)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_orders_created_desc ON orders(created_at DESC)")
		// cart_items: ایندکس پرکاربرد
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_cart_items_cart_product ON cart_items(cart_id, product_id)")
		return nil
	})
}
