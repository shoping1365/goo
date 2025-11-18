package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

// Up025InventoryManagement ایجاد جداول سیستم مدیریت موجودی
func Up025InventoryManagement(tx *gorm.DB) error {
	// جدول موجودی محصولات
	if err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS product_inventory (
			id BIGSERIAL PRIMARY KEY,
			product_id BIGINT NOT NULL UNIQUE,
			stock_quantity INTEGER NOT NULL DEFAULT 0,
			reserved_quantity INTEGER NOT NULL DEFAULT 0,
			available_quantity INTEGER GENERATED ALWAYS AS (stock_quantity - reserved_quantity) STORED,
			low_stock_threshold INTEGER NOT NULL DEFAULT 10,
			is_tracked BOOLEAN NOT NULL DEFAULT true,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`).Error; err != nil {
		return fmt.Errorf("failed to create product_inventory table: %w", err)
	}

	// ایجاد ایندکس‌ها برای product_inventory
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_inventory_product_id ON product_inventory (product_id)`).Error; err != nil {
		return fmt.Errorf("failed to create index on product_inventory: %w", err)
	}
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_inventory_stock_quantity ON product_inventory (stock_quantity)`).Error; err != nil {
		return fmt.Errorf("failed to create index on product_inventory: %w", err)
	}
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_inventory_available_quantity ON product_inventory (available_quantity)`).Error; err != nil {
		return fmt.Errorf("failed to create index on product_inventory: %w", err)
	}

	// جدول رزرو موجودی
	if err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS inventory_reservations (
			id BIGSERIAL PRIMARY KEY,
			product_id BIGINT NOT NULL,
			order_id BIGINT NOT NULL,
			reserved_quantity INTEGER NOT NULL,
			reserved_until TIMESTAMP NOT NULL,
			is_released BOOLEAN NOT NULL DEFAULT false,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`).Error; err != nil {
		return fmt.Errorf("failed to create inventory_reservations table: %w", err)
	}

	// ایجاد ایندکس‌ها برای inventory_reservations
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_inventory_reservations_product_id ON inventory_reservations (product_id)`).Error; err != nil {
		return fmt.Errorf("failed to create index on inventory_reservations: %w", err)
	}
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_inventory_reservations_order_id ON inventory_reservations (order_id)`).Error; err != nil {
		return fmt.Errorf("failed to create index on inventory_reservations: %w", err)
	}
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_inventory_reservations_reserved_until ON inventory_reservations (reserved_until)`).Error; err != nil {
		return fmt.Errorf("failed to create index on inventory_reservations: %w", err)
	}
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_inventory_reservations_is_released ON inventory_reservations (is_released)`).Error; err != nil {
		return fmt.Errorf("failed to create index on inventory_reservations: %w", err)
	}

	// جدول سفارش‌ها - فقط فیلدهای جدید اضافه می‌کنیم
	if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS order_number VARCHAR(50)`).Error; err != nil {
		return fmt.Errorf("failed to add order_number column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS subtotal_amount DECIMAL(10,2) DEFAULT 0`).Error; err != nil {
		return fmt.Errorf("failed to add subtotal_amount column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS tax_amount DECIMAL(10,2) DEFAULT 0`).Error; err != nil {
		return fmt.Errorf("failed to add tax_amount column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS shipping_amount DECIMAL(10,2) DEFAULT 0`).Error; err != nil {
		return fmt.Errorf("failed to add shipping_amount column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS discount_amount DECIMAL(10,2) DEFAULT 0`).Error; err != nil {
		return fmt.Errorf("failed to add discount_amount column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS shipping_address TEXT`).Error; err != nil {
		return fmt.Errorf("failed to add shipping_address column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS billing_address TEXT`).Error; err != nil {
		return fmt.Errorf("failed to add billing_address column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS notes TEXT`).Error; err != nil {
		return fmt.Errorf("failed to add notes column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE orders ADD COLUMN IF NOT EXISTS reserved_until TIMESTAMP`).Error; err != nil {
		return fmt.Errorf("failed to add reserved_until column: %w", err)
	}

	// ایجاد ایندکس‌ها برای orders (فقط ایندکس‌های جدید)
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_order_number ON orders (order_number)`).Error; err != nil {
		return fmt.Errorf("failed to create index on orders: %w", err)
	}
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_status ON orders (status)`).Error; err != nil {
		return fmt.Errorf("failed to create index on orders: %w", err)
	}
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_payment_status ON orders (payment_status)`).Error; err != nil {
		return fmt.Errorf("failed to create index on orders: %w", err)
	}
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_orders_created_at ON orders (created_at)`).Error; err != nil {
		return fmt.Errorf("failed to create index on orders: %w", err)
	}

	// جدول جزئیات سفارش‌ها - فقط فیلدهای جدید اضافه می‌کنیم
	if err := tx.Exec(`ALTER TABLE order_items ADD COLUMN IF NOT EXISTS product_name VARCHAR(255)`).Error; err != nil {
		return fmt.Errorf("failed to add product_name column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE order_items ADD COLUMN IF NOT EXISTS product_image VARCHAR(500)`).Error; err != nil {
		return fmt.Errorf("failed to add product_image column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE order_items ADD COLUMN IF NOT EXISTS product_sku VARCHAR(100)`).Error; err != nil {
		return fmt.Errorf("failed to add product_sku column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE order_items ADD COLUMN IF NOT EXISTS total_price DECIMAL(10,2) DEFAULT 0`).Error; err != nil {
		return fmt.Errorf("failed to add total_price column: %w", err)
	}
	if err := tx.Exec(`ALTER TABLE order_items ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP`).Error; err != nil {
		return fmt.Errorf("failed to add updated_at column: %w", err)
	}

	// ایجاد ایندکس‌ها برای order_items (فقط ایندکس‌های جدید)
	if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_order_items_product_id ON order_items (product_id)`).Error; err != nil {
		return fmt.Errorf("failed to create index on order_items: %w", err)
	}

	// تنظیمات سیستم موجودی
	if err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS inventory_settings (
			id BIGSERIAL PRIMARY KEY,
			reservation_duration_minutes INTEGER NOT NULL DEFAULT 30, -- مدت رزرو به دقیقه
			auto_release_enabled BOOLEAN NOT NULL DEFAULT true, -- آزادسازی خودکار رزروهای منقضی
			low_stock_notification_enabled BOOLEAN NOT NULL DEFAULT true, -- اعلان موجودی کم
			out_of_stock_notification_enabled BOOLEAN NOT NULL DEFAULT true, -- اعلان اتمام موجودی
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`).Error; err != nil {
		return fmt.Errorf("failed to create inventory_settings table: %w", err)
	}

	// درج تنظیمات پیش‌فرض
	if err := tx.Exec(`
		INSERT INTO inventory_settings (id, reservation_duration_minutes, auto_release_enabled, low_stock_notification_enabled, out_of_stock_notification_enabled)
		VALUES (1, 30, true, true, true)
		ON CONFLICT (id) DO UPDATE SET id = EXCLUDED.id
	`).Error; err != nil {
		return fmt.Errorf("failed to insert default inventory settings: %w", err)
	}

	return nil
}
