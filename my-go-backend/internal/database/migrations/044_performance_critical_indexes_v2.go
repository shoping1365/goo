package migrations

import (
	"gorm.io/gorm"
	"log"
)

// Up044PerformanceCriticalIndexesV2 ایندکس‌های پرترافیک (نسخه منسجم)
func Up044PerformanceCriticalIndexesV2(db *gorm.DB) error {
	// بدون transaction اجرا می‌شود تا اگر یکی fail شد بقیه اجرا شوند

	// Enable pg_trgm extension (ممکن است قبلاً نصب شده باشد)
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm").Error; err != nil {
		log.Printf("⚠️  Warning: Could not create pg_trgm extension: %v (may already exist or need superuser)", err)
	}

	// Orders indexes
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_orders_user_status_created ON orders(user_id, status, created_at DESC)").Error; err != nil {
		log.Printf("⚠️  Warning: idx_orders_user_status_created: %v", err)
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_orders_paid_created ON orders(created_at DESC) WHERE is_paid = true").Error; err != nil {
		log.Printf("⚠️  Warning: idx_orders_paid_created: %v", err)
	}

	// Products indexes (پویا با فیلترهای متداول)
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_products_status_cat_updated ON products(category_id, is_active, status, updated_at DESC) WHERE status = 'published' AND is_active = true").Error; err != nil {
		log.Printf("⚠️  Warning: idx_products_status_cat_updated: %v", err)
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_products_brand_active_updated ON products(brand_id, is_active, updated_at DESC) WHERE is_active = true").Error; err != nil {
		log.Printf("⚠️  Warning: idx_products_brand_active_updated: %v", err)
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_products_status_active_price ON products(category_id, is_active, price) WHERE is_active = true").Error; err != nil {
		log.Printf("⚠️  Warning: idx_products_status_active_price: %v", err)
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_products_stock_track_status ON products(stock_quantity, track_inventory, status)").Error; err != nil {
		log.Printf("⚠️  Warning: idx_products_stock_track_status: %v", err)
	}

	// Search acceleration (ILIKE): pg_trgm + GIN indexes
	// فقط اگر pg_trgm موجود باشد
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_products_name_trgm ON products USING GIN (name gin_trgm_ops)").Error; err != nil {
		log.Printf("⚠️  Warning: idx_products_name_trgm (needs pg_trgm): %v", err)
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_products_desc_trgm ON products USING GIN (full_description gin_trgm_ops)").Error; err != nil {
		log.Printf("⚠️  Warning: idx_products_desc_trgm (needs pg_trgm): %v", err)
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_products_sku_trgm ON products USING GIN (sku gin_trgm_ops)").Error; err != nil {
		log.Printf("⚠️  Warning: idx_products_sku_trgm (needs pg_trgm): %v", err)
	}

	// Sessions index
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_user_sessions_expires_active ON user_sessions(expires_at, is_active) WHERE is_active = true").Error; err != nil {
		log.Printf("⚠️  Warning: idx_user_sessions_expires_active: %v", err)
	}

	log.Println("✅ Migration 044: Performance indexes applied (with warnings if any)")
	return nil
}
