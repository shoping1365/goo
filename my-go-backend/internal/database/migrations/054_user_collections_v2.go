package migrations

import "gorm.io/gorm"

// Up054UserCollectionsV2 تایید ایندکس یکتا
func Up054UserCollectionsV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE UNIQUE INDEX IF NOT EXISTS ux_collection_product ON user_collection_items(collection_id, product_id)")
		return nil
	})
}
