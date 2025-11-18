package migrations

import "gorm.io/gorm"

// Up081FixZeroStockIssue اصلاح مشکل صفر شدن موجودی‌ها
// این migration مشکل migration قبلی را حل می‌کند
func Up081FixZeroStockIssue(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. حذف رکوردهای انبار که quantity = 0 هستند و محصول stock_quantity > 0 دارد
		if err := tx.Exec(`
			DELETE FROM product_warehouse_stocks 
			WHERE quantity = 0 
			AND product_id IN (
				SELECT id FROM products WHERE COALESCE(stock_quantity, 0) > 0
			)
		`).Error; err != nil {
			return err
		}

		// 2. ایجاد مجدد رکوردهای انبار برای محصولاتی که stock_quantity > 0 دارند
		if err := tx.Exec(`
			INSERT INTO product_warehouse_stocks (product_id, warehouse_id, quantity, reserved, min_qty, max_qty)
			SELECT 
				p.id, 
				(SELECT id FROM warehouses WHERE is_default = TRUE LIMIT 1),
				COALESCE(p.stock_quantity, 0),
				0,
				COALESCE(p.min_stock_quantity, 0),
				COALESCE(p.max_stock_quantity, 0)
			FROM products p
			WHERE COALESCE(p.stock_quantity, 0) > 0
			AND NOT EXISTS (
				SELECT 1 FROM product_warehouse_stocks pws 
				WHERE pws.product_id = p.id 
				AND pws.warehouse_id = (SELECT id FROM warehouses WHERE is_default = TRUE LIMIT 1)
			)
		`).Error; err != nil {
			return err
		}

		// 3. به‌روزرسانی موجودی انبارها بر اساس stock_quantity محصولات
		if err := tx.Exec(`
			UPDATE product_warehouse_stocks 
			SET quantity = p.stock_quantity
			FROM products p
			WHERE product_warehouse_stocks.product_id = p.id
			AND COALESCE(p.stock_quantity, 0) > 0
			AND product_warehouse_stocks.warehouse_id = (SELECT id FROM warehouses WHERE is_default = TRUE LIMIT 1)
		`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down081FixZeroStockIssue برگرداندن تغییرات (در صورت نیاز)
func Down081FixZeroStockIssue(db *gorm.DB) error {
	// این migration قابل برگشت نیست چون مشکل را حل می‌کند
	return nil
}
