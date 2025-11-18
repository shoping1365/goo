package migrations

import "gorm.io/gorm"

// Up021UserCollectionsAndAddressLabel ایجاد user_collections و user_collection_items + ستون label برای user_addresses
func Up021UserCollectionsAndAddressLabel(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS user_collections (
                id BIGSERIAL PRIMARY KEY,
                user_id BIGINT NOT NULL,
                name VARCHAR(100) NOT NULL,
                is_default BOOLEAN NOT NULL DEFAULT FALSE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_user_collections_user ON user_collections(user_id)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS user_collection_items (
                id BIGSERIAL PRIMARY KEY,
                collection_id BIGINT NOT NULL,
                product_id BIGINT NOT NULL,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS ux_collection_product ON user_collection_items(collection_id, product_id)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`ALTER TABLE user_addresses ADD COLUMN IF NOT EXISTS label VARCHAR(100)`).Error; err != nil {
			return err
		}
		return nil
	})
}
