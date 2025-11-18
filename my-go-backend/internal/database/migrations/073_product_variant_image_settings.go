package migrations

import "gorm.io/gorm"

// Up073ProductVariantImageSettings افزودن ستون‌های تنظیمات نمایش تصاویر تنوع به جدول products
func Up073ProductVariantImageSettings(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن ستون‌ها در صورت نبود
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS show_variant_image_in_list BOOLEAN DEFAULT FALSE`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS change_main_image_on_variant_select BOOLEAN DEFAULT FALSE`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS show_all_variant_images_in_gallery BOOLEAN DEFAULT FALSE`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE products ADD COLUMN IF NOT EXISTS auto_compress_variant_images BOOLEAN DEFAULT FALSE`).Error; err != nil {
			return err
		}
		return nil
	})
}
