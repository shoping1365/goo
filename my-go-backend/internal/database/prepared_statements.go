package database

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

// PreparedStatementManager مدیریت Prepared Statements برای بهبود عملکرد
type PreparedStatementManager struct {
	db          *gorm.DB
	statements  map[string]*sql.Stmt
	mutex       sync.RWMutex
	initialized bool
}

// NewPreparedStatementManager ایجاد نمونه جدید از PreparedStatementManager
func NewPreparedStatementManager(db *gorm.DB) *PreparedStatementManager {
	return &PreparedStatementManager{
		db:         db,
		statements: make(map[string]*sql.Stmt),
	}
}

// Initialize آماده‌سازی تمام Prepared Statements
func (pm *PreparedStatementManager) Initialize(ctx context.Context) error {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if pm.initialized {
		return nil
	}

	// دریافت connection از GORM
	sqlDB, err := pm.db.DB()
	if err != nil {
		return fmt.Errorf("خطا در دریافت connection: %w", err)
	}

	// تعریف Prepared Statements
	statements := map[string]string{
		// Product Statements
		"product_get_by_id":       "SELECT * FROM products WHERE id = $1 AND deleted_at IS NULL",
		"product_get_by_sku":      "SELECT * FROM products WHERE sku = $1 AND deleted_at IS NULL",
		"product_get_by_slug":     "SELECT * FROM products WHERE slug = $1 AND deleted_at IS NULL",
		"product_get_published":   "SELECT * FROM products WHERE status = 'published' AND is_active = true AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2",
		"product_get_by_category": "SELECT * FROM products WHERE category_id = $1 AND status = 'published' AND is_active = true AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3",
		"product_get_by_brand":    "SELECT * FROM products WHERE brand_id = $1 AND status = 'published' AND is_active = true AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3",
		"product_get_featured":    "SELECT * FROM products WHERE is_featured = true AND status = 'published' AND is_active = true AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $1",
		"product_get_new":         "SELECT * FROM products WHERE is_new = true AND status = 'published' AND is_active = true AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $1",
		"product_get_discounted":  "SELECT * FROM products WHERE sale_price < price AND status = 'published' AND is_active = true AND deleted_at IS NULL ORDER BY (price - sale_price) DESC LIMIT $1",
		"product_search":          "SELECT * FROM products WHERE (name ILIKE $1 OR description ILIKE $1 OR sku ILIKE $1) AND status = 'published' AND is_active = true AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3",
		"product_get_low_stock":   "SELECT * FROM products WHERE stock_quantity <= min_stock_quantity AND track_inventory = true AND deleted_at IS NULL ORDER BY stock_quantity ASC LIMIT $1 OFFSET $2",
		"product_update_stock":    "UPDATE products SET stock_quantity = $1, stock_status = $2, updated_at = NOW() WHERE id = $3",
		"product_update_status":   "UPDATE products SET status = $1, updated_at = NOW() WHERE id = $2",
		"product_soft_delete":     "UPDATE products SET deleted_at = NOW() WHERE id = $1",

		// Category Statements
		"category_get_by_id":      "SELECT * FROM categories WHERE id = $1 AND deleted_at IS NULL",
		"category_get_by_slug":    "SELECT * FROM categories WHERE slug = $1 AND deleted_at IS NULL",
		"category_get_root":       "SELECT * FROM categories WHERE parent_id IS NULL AND is_active = true AND deleted_at IS NULL ORDER BY sort_order ASC",
		"category_get_children":   "SELECT * FROM categories WHERE parent_id = $1 AND is_active = true AND deleted_at IS NULL ORDER BY sort_order ASC",
		"category_get_featured":   "SELECT * FROM categories WHERE is_featured = true AND is_active = true AND deleted_at IS NULL ORDER BY sort_order ASC LIMIT $1",
		"category_get_with_count": "SELECT c.*, COUNT(p.id) as product_count FROM categories c LEFT JOIN products p ON c.id = p.category_id AND p.status = 'published' AND p.deleted_at IS NULL WHERE c.is_active = true AND c.deleted_at IS NULL GROUP BY c.id ORDER BY c.sort_order ASC",

		// Brand Statements
		"brand_get_by_id":    "SELECT * FROM brands WHERE id = $1 AND deleted_at IS NULL",
		"brand_get_by_slug":  "SELECT * FROM brands WHERE slug = $1 AND deleted_at IS NULL",
		"brand_get_all":      "SELECT * FROM brands WHERE is_active = true AND deleted_at IS NULL ORDER BY name ASC LIMIT $1 OFFSET $2",
		"brand_get_featured": "SELECT * FROM brands WHERE is_featured = true AND is_active = true AND deleted_at IS NULL ORDER BY name ASC LIMIT $1",
		"brand_search":       "SELECT * FROM brands WHERE name ILIKE $1 AND is_active = true AND deleted_at IS NULL ORDER BY name ASC LIMIT $2",

		// Review Statements
		"review_get_by_id":          "SELECT * FROM reviews WHERE id = $1 AND deleted_at IS NULL",
		"review_get_by_product":     "SELECT * FROM reviews WHERE product_id = $1 AND status = $2 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $3 OFFSET $4",
		"review_get_by_user":        "SELECT * FROM reviews WHERE user_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3",
		"review_get_published":      "SELECT * FROM reviews WHERE status = 'approved' AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2",
		"review_get_pending":        "SELECT * FROM reviews WHERE status = 'pending' AND deleted_at IS NULL ORDER BY created_at ASC LIMIT $1 OFFSET $2",
		"review_get_average_rating": "SELECT AVG(rating) as average_rating, COUNT(*) as total_reviews FROM reviews WHERE product_id = $1 AND status = 'approved' AND deleted_at IS NULL",
		"review_update_status":      "UPDATE reviews SET status = $1, updated_at = NOW() WHERE id = $2",

		// Wishlist Statements
		"wishlist_get_by_user":  "SELECT * FROM wishlist_items WHERE user_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC",
		"wishlist_get_products": "SELECT p.* FROM products p INNER JOIN wishlist_items wi ON p.id = wi.product_id WHERE wi.user_id = $1 AND p.deleted_at IS NULL AND wi.deleted_at IS NULL ORDER BY wi.created_at DESC LIMIT $2 OFFSET $3",
		"wishlist_add_item":     "INSERT INTO wishlist_items (user_id, product_id, created_at) VALUES ($1, $2, NOW()) ON CONFLICT (user_id, product_id) DO NOTHING",
		"wishlist_remove_item":  "UPDATE wishlist_items SET deleted_at = NOW() WHERE user_id = $1 AND product_id = $2",
		"wishlist_clear":        "UPDATE wishlist_items SET deleted_at = NOW() WHERE user_id = $1",
		"wishlist_check_item":   "SELECT COUNT(*) FROM wishlist_items WHERE user_id = $1 AND product_id = $2 AND deleted_at IS NULL",

		// Stock Notification Statements
		"stock_notification_get_by_product": "SELECT * FROM stock_notifications WHERE product_id = $1 AND status = $2 AND deleted_at IS NULL",
		"stock_notification_get_by_user":    "SELECT * FROM stock_notifications WHERE user_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC",
		"stock_notification_get_pending":    "SELECT * FROM stock_notifications WHERE status = 'pending' AND deleted_at IS NULL ORDER BY created_at ASC LIMIT $1 OFFSET $2",
		"stock_notification_mark_sent":      "UPDATE stock_notifications SET status = 'sent', sent_at = NOW() WHERE id = $1",

		// Product QA Statements
		"product_qa_get_questions": "SELECT * FROM product_qa WHERE product_id = $1 AND parent_id IS NULL AND status = $2 AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $3 OFFSET $4",
		"product_qa_get_answers":   "SELECT * FROM product_qa WHERE parent_id = $1 AND status = $2 AND deleted_at IS NULL ORDER BY created_at ASC",
		"product_qa_get_published": "SELECT * FROM product_qa WHERE product_id = $1 AND status = 'approved' AND deleted_at IS NULL ORDER BY created_at DESC LIMIT $2 OFFSET $3",
		"product_qa_get_pending":   "SELECT * FROM product_qa WHERE status = 'pending' AND deleted_at IS NULL ORDER BY created_at ASC LIMIT $1 OFFSET $2",
		"product_qa_update_status": "UPDATE product_qa SET status = $1, updated_at = NOW() WHERE id = $2",

		// Menu Statements (reads)
		"menu_get_by_id":               "SELECT * FROM menus WHERE id = $1",
		"menu_get_by_slug":             "SELECT * FROM menus WHERE slug = $1",
		"menu_get_enabled":             "SELECT * FROM menus WHERE enabled = true ORDER BY \"order\" ASC",
		"menu_items_by_menu_root":      "SELECT * FROM menu_items WHERE menu_id = $1 AND parent_id IS NULL ORDER BY \"order\" ASC",
		"menu_children_by_parent":      "SELECT * FROM menu_items WHERE parent_id = $1 ORDER BY \"order\" ASC",
		"menu_columns_by_menu":         "SELECT * FROM menu_columns WHERE menu_id = $1 ORDER BY \"order\" ASC",
		"menu_locations_enabled":       "SELECT * FROM menu_locations ORDER BY \"order\" ASC",
		"menu_location_by_slug":        "SELECT * FROM menu_locations WHERE slug = $1",
		"menu_assignments_by_location": "SELECT * FROM menu_assignments WHERE location_id = $1 AND enabled = true ORDER BY \"order\" ASC",

		// Media quick reads
		"media_get_by_id":   "SELECT * FROM media_files WHERE id = $1",
		"media_get_by_path": "SELECT * FROM media_files WHERE url = $1",
		"media_recent":      "SELECT * FROM media_files ORDER BY created_at DESC LIMIT $1",
	}

	// آماده‌سازی statements
	for name, query := range statements {
		stmt, err := sqlDB.PrepareContext(ctx, query)
		if err != nil {
			return fmt.Errorf("خطا در آماده‌سازی statement %s: %w", name, err)
		}
		pm.statements[name] = stmt
	}

	pm.initialized = true
	return nil
}

// GetStatement دریافت Prepared Statement
func (pm *PreparedStatementManager) GetStatement(name string) (*sql.Stmt, error) {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()

	if !pm.initialized {
		return nil, fmt.Errorf("PreparedStatementManager هنوز آماده نشده است")
	}

	stmt, exists := pm.statements[name]
	if !exists {
		return nil, fmt.Errorf("statement %s یافت نشد", name)
	}

	return stmt, nil
}

// Close بستن تمام Prepared Statements
func (pm *PreparedStatementManager) Close() error {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	for name, stmt := range pm.statements {
		if err := stmt.Close(); err != nil {
			return fmt.Errorf("خطا در بستن statement %s: %w", name, err)
		}
	}

	pm.statements = make(map[string]*sql.Stmt)
	pm.initialized = false
	return nil
}

// IsInitialized بررسی آماده بودن PreparedStatementManager
func (pm *PreparedStatementManager) IsInitialized() bool {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	return pm.initialized
}

// ==================== Helper Functions ====================

// ExecutePreparedQuery اجرای کوئری با Prepared Statement
func ExecutePreparedQuery(ctx context.Context, stmt *sql.Stmt, args ...interface{}) (*sql.Rows, error) {
	return stmt.QueryContext(ctx, args...)
}

// ExecutePreparedQueryRow اجرای کوئری تک ردیف با Prepared Statement
func ExecutePreparedQueryRow(ctx context.Context, stmt *sql.Stmt, args ...interface{}) *sql.Row {
	return stmt.QueryRowContext(ctx, args...)
}

// ExecutePreparedExec اجرای دستور INSERT/UPDATE/DELETE با Prepared Statement
func ExecutePreparedExec(ctx context.Context, stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {
	return stmt.ExecContext(ctx, args...)
}
