package migrations

import "gorm.io/gorm"

// Up034WidgetsAndCMSIndexes ایندکس‌های کلیدی برای widgets/posts/menus
func Up034WidgetsAndCMSIndexes(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_widgets_type_status ON widgets(type, status)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_posts_status_published ON posts(status, published_at)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_menu_items_menu_order ON menu_items(menu_id, \"order\")")
		return nil
	})
}
