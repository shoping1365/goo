package migrations

import "gorm.io/gorm"

// Up085ImageSEOJobs ایجاد جدول مجزای صف پردازش متادیتای تصاویر با هوش مصنوعی
// این جدول برای زمان‌بندی، مانیتورینگ و گزارش خطاهای تولید alt/caption/description استفاده می‌شود.
// ایندکس‌های حیاتی: (status, scheduled_at), (media_id), (created_at DESC)
func Up085ImageSEOJobs(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// جدول اصلی صف
		if err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS image_seo_jobs (
				id BIGSERIAL PRIMARY KEY,
				media_id BIGINT NOT NULL,
				status VARCHAR(20) NOT NULL DEFAULT 'pending', -- pending|processing|completed|error
				retries INT NOT NULL DEFAULT 0,
				scheduled_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
				started_at TIMESTAMPTZ,
				finished_at TIMESTAMPTZ,
				error_message TEXT,
				model VARCHAR(50),
				lang VARCHAR(10),
				context_title TEXT,
				prompt_template TEXT,
				output_json JSONB,
				created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
				updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
				CONSTRAINT fk_image_seo_media FOREIGN KEY(media_id) REFERENCES media_files(id) ON DELETE CASCADE
			);
		`).Error; err != nil {
			return err
		}

		// ایندکس‌ها
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_image_seo_jobs_status_sched ON image_seo_jobs(status, scheduled_at)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_image_seo_jobs_media ON image_seo_jobs(media_id)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_image_seo_jobs_created ON image_seo_jobs(created_at DESC)`).Error; err != nil {
			return err
		}

		// کلیدهای تنظیمات پیش‌فرض در جدول settings (idempotent)
		if err := tx.Exec(`
			INSERT INTO settings(key, value, description, category, type)
			VALUES
			('image_seo.enabled','true','فعال‌سازی تولید خودکار متادیتای تصاویر','image_seo','boolean'),
			('image_seo.delay_seconds','60','تاخیر اجرای job پس از آپلود (ثانیه)','image_seo','number'),
			('image_seo.batch_size','10','تعداد job در هر اجرای worker','image_seo','number'),
			('image_seo.rate_limit','30','حداکثر تعداد درخواست در دقیقه','image_seo','number'),
			('image_seo.model','gpt-4o-mini','مدل پیش‌فرض OpenAI برای توصیف تصویر','image_seo','string'),
			('image_seo.lang','fa','زبان تولید متن','image_seo','string'),
			('image_seo.generate_title','false','آیا عنوان تصویر هم تولید شود؟','image_seo','boolean'),
			('image_seo.overwrite_existing','false','بازنویسی مقادیر موجود در صورت وجود','image_seo','boolean')
			ON CONFLICT (key) DO NOTHING;
		`).Error; err != nil {
			return err
		}

		return nil
	})
}

