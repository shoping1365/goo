package migrations

import "gorm.io/gorm"

// Up141SearchIndexQueue ایجاد جدول صف ایندکس جستجو برای همگام‌سازی بلادرنگ
func Up141SearchIndexQueue(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS search_index_queue (
				id BIGSERIAL PRIMARY KEY,
				resource_type VARCHAR(50) NOT NULL,
				resource_id BIGINT NOT NULL,
				operation VARCHAR(20) NOT NULL,
				payload JSONB,
				status VARCHAR(20) NOT NULL DEFAULT 'pending',
				attempts INTEGER NOT NULL DEFAULT 0,
				error_message TEXT,
				available_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
				created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
				updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
			);
		`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_search_index_queue_status ON search_index_queue(status, available_at)`).Error; err != nil {
			return err
		}

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_search_index_queue_resource ON search_index_queue(resource_type, resource_id)`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down141SearchIndexQueue حذف جدول صف ایندکس جستجو
func Down141SearchIndexQueue(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`DROP TABLE IF EXISTS search_index_queue`).Error; err != nil {
			return err
		}
		return nil
	})
}
