package migrations

import "gorm.io/gorm"

// Up083DropProductsVideoURL حذف ستون قدیمی video_url از جدول products پس از انتقال داده‌ها
// این مایگریشن برای یکپارچه‌سازی کامل منطق ویدیوها با جدول product_videos است.
func Up083DropProductsVideoURL(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 1) انتقال داده‌های قدیمی products.video_url → product_videos (در صورت وجود)
		if err := tx.Exec(`
			INSERT INTO product_videos (product_id, title, description, video_url, thumbnail_url, show_in_gallery, autoplay, show_controls, sort_order, created_at, updated_at)
			SELECT p.id,
			       COALESCE(p.name, ''),
			       '',
			       p.video_url,
			       '',
			       TRUE,
			       FALSE,
			       TRUE,
			       0,
			       NOW(),
			       NOW()
			FROM products p
			WHERE p.video_url IS NOT NULL AND p.video_url <> ''
			  AND NOT EXISTS (
			    SELECT 1 FROM product_videos v
			    WHERE v.product_id = p.id AND v.video_url = p.video_url
			  );
		`).Error; err != nil {
			return err
		}

		// 2) پاک‌سازی مقدار ستون قدیمی (اختیاری)
		if err := tx.Exec(`UPDATE products SET video_url = NULL WHERE video_url IS NOT NULL AND video_url <> ''`).Error; err != nil {
			return err
		}

		// 3) حذف ستون قدیمی
		if err := tx.Exec(`ALTER TABLE products DROP COLUMN IF EXISTS video_url`).Error; err != nil {
			return err
		}

		return nil
	})
}


