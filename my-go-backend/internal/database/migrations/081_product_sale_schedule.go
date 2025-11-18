package migrations

import "gorm.io/gorm"

// Up081ProductSaleSchedule افزودن ستون‌های زمان‌بندی قیمت ویژه به جدول محصولات
// این مایگریشن دو ستون sale_start_at و sale_end_at را به جدول products اضافه می‌کند
// و یک کانسترینت برای اطمینان از معتبر بودن بازه زمانی ایجاد می‌کند.
func Up081ProductSaleSchedule(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. افزودن ستون‌های جدید در صورت نبود
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS sale_start_at TIMESTAMPTZ`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS sale_end_at TIMESTAMPTZ`).Error; err != nil {
			return err
		}

		// 2. ایجاد کانسترینت برای اطمینان از اینکه sale_end_at بعد از sale_start_at است (در صورت وجود هر دو)
		if err := tx.Exec(`
			DO $$
			BEGIN
				IF NOT EXISTS (
					SELECT 1 FROM pg_constraint WHERE conname = 'chk_products_sale_window'
				) THEN
					ALTER TABLE products
						ADD CONSTRAINT chk_products_sale_window CHECK (
							sale_start_at IS NULL
							OR sale_end_at IS NULL
							OR sale_end_at > sale_start_at
						);
				END IF;
			END$$;
		`).Error; err != nil {
			return err
		}

		// 3. ایجاد ایندکس برای بازه فروش ویژه جهت بهبود کوئری‌ها
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_products_sale_window ON products (sale_start_at, sale_end_at)`).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down081ProductSaleSchedule حذف ستون‌ها و محدودیت بازه زمانی فروش ویژه
func Down081ProductSaleSchedule(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`ALTER TABLE products DROP CONSTRAINT IF EXISTS chk_products_sale_window`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE products DROP COLUMN IF EXISTS sale_start_at`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE products DROP COLUMN IF EXISTS sale_end_at`).Error; err != nil {
			return err
		}
		return nil
	})
}
