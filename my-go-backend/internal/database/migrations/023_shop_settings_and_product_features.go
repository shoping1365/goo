package migrations

import "gorm.io/gorm"

// Up023ShopSettingsAndProductFeatures جدول shop_settings و ستون‌های محصول مرتبط
func Up023ShopSettingsAndProductFeatures(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// shop_settings (key/value)
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS shop_settings (
                key VARCHAR(100) PRIMARY KEY,
                value TEXT
            );
        `).Error; err != nil {
			return err
		}

		// محصولات: اطمینان از وجود/هماهنگی ستون‌ها
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS show_stock_to_customer BOOLEAN DEFAULT FALSE`).Error; err != nil {
			return err
		}
		// rename track_stock -> track_inventory (اگر قدیمی وجود داشت)
		if err := tx.Exec(`DO $$
        BEGIN
            IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='products' AND column_name='track_stock')
               AND NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='products' AND column_name='track_inventory') THEN
                ALTER TABLE products RENAME COLUMN track_stock TO track_inventory;
            END IF;
        END $$;`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS track_inventory BOOLEAN DEFAULT TRUE`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS disable_buy_button BOOLEAN DEFAULT FALSE`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS call_for_price BOOLEAN DEFAULT FALSE`).Error; err != nil {
			return err
		}
		return nil
	})
}
