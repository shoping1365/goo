package migrations

import "gorm.io/gorm"

// Up080WarehouseBasedInventory تبدیل سیستم موجودی به انبارمحور
// این مایگریشن محصولات را کاملاً متکی به انبارها می‌کند
func Up080WarehouseBasedInventory(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. اطمینان از وجود حداقل یک انبار
		if err := tx.Exec(`
			INSERT INTO warehouses (code, name, city, is_default, is_active)
			SELECT 'DEFAULT', 'انبار مرکزی', 'تهران', TRUE, TRUE
			WHERE NOT EXISTS (SELECT 1 FROM warehouses WHERE is_default = TRUE)
		`).Error; err != nil {
			return err
		}

		// 2. برای محصولاتی که هنوز رکورد انبار ندارند، رکورد ایجاد کن
		// فقط برای محصولاتی که stock_quantity > 0 دارند
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

		// 3. برای محصولاتی که stock_quantity > 0 دارند ولی رکورد انبار ندارند، موجودی را به انبار پیش‌فرض منتقل کن
		if err := tx.Exec(`
			UPDATE product_warehouse_stocks 
			SET quantity = p.stock_quantity
			FROM products p
			WHERE product_warehouse_stocks.product_id = p.id
			AND p.stock_quantity > 0
			AND product_warehouse_stocks.warehouse_id = (SELECT id FROM warehouses WHERE is_default = TRUE LIMIT 1)
		`).Error; err != nil {
			return err
		}

		// 4. اطمینان از اینکه تمام محصولات حداقل یک رکورد انبار دارند
		// فقط برای محصولاتی که واقعاً نیاز به رکورد انبار دارند
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
			WHERE NOT EXISTS (
				SELECT 1 FROM product_warehouse_stocks pws WHERE pws.product_id = p.id
			)
			AND (
				COALESCE(p.stock_quantity, 0) > 0 
				OR COALESCE(p.min_stock_quantity, 0) > 0 
				OR COALESCE(p.max_stock_quantity, 0) > 0
			)
		`).Error; err != nil {
			return err
		}

		// 5. تنظیم default_warehouse_id برای محصولاتی که ندارند
		if err := tx.Exec(`
			UPDATE products 
			SET default_warehouse_id = (SELECT id FROM warehouses WHERE is_default = TRUE LIMIT 1)
			WHERE default_warehouse_id IS NULL
		`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down080WarehouseBasedInventory برگرداندن سیستم موجودی به حالت قبلی
func Down080WarehouseBasedInventory(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// برگرداندن موجودی کل محصولات بر اساس مجموع انبارها
		if err := tx.Exec(`
			UPDATE products 
			SET stock_quantity = COALESCE((
				SELECT SUM(quantity) 
				FROM product_warehouse_stocks 
				WHERE product_id = products.id
			), 0)
		`).Error; err != nil {
			return err
		}

		return nil
	})
}
