package migrations

import "gorm.io/gorm"

// Up084ProductVideosLazyLoad اضافه کردن ستون lazy_load به جدول product_videos
// این ستون مشخص می‌کند آیا ویدیو به‌صورت تنبل (پس از کلیک) بارگذاری شود یا خیر.
func Up084ProductVideosLazyLoad(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            ALTER TABLE product_videos
            ADD COLUMN IF NOT EXISTS lazy_load BOOLEAN NOT NULL DEFAULT TRUE;
        `).Error; err != nil {
			return err
		}
		// تعیین مقدار TRUE برای رکوردهای موجود (بجز مواردی که صریحاً FALSE ثبت شده باشند)
		if err := tx.Exec(`UPDATE product_videos SET lazy_load = TRUE WHERE lazy_load IS NULL;`).Error; err != nil {
			return err
		}
		// ایندکس برای فیلد جستجوی سریع بر اساس product_id و lazy_load
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_product_videos_lazy ON product_videos(product_id, lazy_load);`).Error; err != nil {
			return err
		}
		return nil
	})
}
