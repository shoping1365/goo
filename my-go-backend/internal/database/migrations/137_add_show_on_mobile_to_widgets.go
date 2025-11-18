package migrations

import "gorm.io/gorm"

// Migration137AddShowOnMobileToWidgets افزودن فیلد show_on_mobile به جدول widgets
// این فیلد مشخص می‌کند که آیا ویجت باید در موبایل نمایش داده شود یا خیر
func Migration137AddShowOnMobileToWidgets(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن ستون show_on_mobile با مقدار پیش‌فرض true
		if err := tx.Exec(`
			ALTER TABLE widgets 
			ADD COLUMN IF NOT EXISTS show_on_mobile BOOLEAN DEFAULT true NOT NULL
		`).Error; err != nil {
			return err
		}

		// به‌روزرسانی تمام رکوردهای موجود به true
		if err := tx.Exec(`
			UPDATE widgets 
			SET show_on_mobile = true 
			WHERE show_on_mobile IS NULL
		`).Error; err != nil {
			return err
		}

		// ایجاد ایندکس برای بهبود عملکرد query ها
		if err := tx.Exec(`
			CREATE INDEX IF NOT EXISTS idx_widgets_show_on_mobile 
			ON widgets(show_on_mobile)
		`).Error; err != nil {
			return err
		}

		// اضافه کردن کامنت توضیحی
		if err := tx.Exec(`
			COMMENT ON COLUMN widgets.show_on_mobile 
			IS 'آیا این ویجت در موبایل نمایش داده شود؟ (true/false)'
		`).Error; err != nil {
			return err
		}

		return nil
	})
}
