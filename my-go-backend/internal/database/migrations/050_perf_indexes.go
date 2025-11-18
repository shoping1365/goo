package migrations

import (
	"database/sql"

	"gorm.io/gorm"
)

// Up050PerfIndexes ایندکس‌های تکمیلی پرفورمنسی برای جداول هسته
// این فایل فقط ایندکس‌هایی را ایجاد می‌کند که جدولشان وجود دارد.
func Up050PerfIndexes(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		tableExists := func(name string) bool {
			var cnt int64
			if err := tx.Raw(`SELECT COUNT(1) FROM information_schema.tables WHERE table_name = ?`, name).Scan(&cnt).Error; err != nil {
				return false
			}
			return cnt > 0
		}

		// users
		if tableExists("users") {
			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_users_verified_registered ON users(registered_at DESC) WHERE email_verified = true OR mobile_verified = true`).Error; err != nil {
				return err
			}
		}

		// sessions
		if tableExists("sessions") {
			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_sessions_user_active ON sessions(user_id, is_active)`).Error; err != nil {
				return err
			}
			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_sessions_last_used_active ON sessions(last_used_at DESC) WHERE is_active = true`).Error; err != nil {
				return err
			}
		}

		// role_permissions
		if tableExists("role_permissions") {
			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_role_permissions_role_active ON role_permissions(role_id, is_active)`).Error; err != nil {
				return err
			}
		}

		// products – ایندکس‌های پوششی برای لیست‌ها و فیلترها
		if tableExists("products") {
			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_products_category_active_status_updated ON products(category_id, is_active, status, updated_at DESC)`).Error; err != nil {
				return err
			}
			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_products_category_active_price ON products(category_id, is_active, price)`).Error; err != nil {
				return err
			}
			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_products_brand_active_updated ON products(brand_id, is_active, updated_at DESC)`).Error; err != nil {
				return err
			}
			for _, idx := range []string{
				`CREATE INDEX IF NOT EXISTS idx_products_is_active ON products(is_active)`,
				`CREATE INDEX IF NOT EXISTS idx_products_status ON products(status)`,
				`CREATE INDEX IF NOT EXISTS idx_products_created_at ON products(created_at)`,
				`CREATE INDEX IF NOT EXISTS idx_products_updated_at ON products(updated_at)`,
				`CREATE INDEX IF NOT EXISTS idx_products_stock_status ON products(stock_status)`,
			} {
				if err := tx.Exec(idx).Error; err != nil {
					return err
				}
			}
		}

		// جداول فرعی محصولات
		if tableExists("product_images") {
			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_images_product_id ON product_images(product_id)`).Error; err != nil {
				return err
			}
		}
		if tableExists("product_variants") {
			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_variants_product_id ON product_variants(product_id)`).Error; err != nil {
				return err
			}
		}
		if tableExists("product_specifications") {
			if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_specifications_product_id ON product_specifications(product_id)`).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// helper to avoid unused import removal in some editors
var _ = sql.ErrNoRows
