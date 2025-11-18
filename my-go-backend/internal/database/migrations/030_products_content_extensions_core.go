package migrations

import "gorm.io/gorm"

// Up030ProductsContentExtensionsCore پوشش 05 (سؤالات/ویدیو/مشخصات/نقدها) در صورت نیاز
func Up030ProductsContentExtensionsCore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// review_helpful_votes ایجاد شده در 020؛ در اینجا صرفاً اطمینان از جداول اصلی
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS product_questions (
            id BIGSERIAL PRIMARY KEY,
            user_id BIGINT,
            product_id BIGINT NOT NULL,
            category_id BIGINT,
            question TEXT NOT NULL,
            is_anonymous BOOLEAN DEFAULT FALSE,
            status VARCHAR(20),
            reject_reason TEXT,
            admin_reply TEXT,
            admin_reply_at TIMESTAMPTZ,
            guest_name VARCHAR(100),
            guest_phone VARCHAR(20),
            ip_address VARCHAR(45),
            created_at TIMESTAMPTZ DEFAULT NOW(),
            updated_at TIMESTAMPTZ DEFAULT NOW()
        );`)
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS product_specifications (
            id BIGSERIAL PRIMARY KEY,
            product_id BIGINT NOT NULL,
            name VARCHAR(255),
            value TEXT,
            created_at TIMESTAMPTZ DEFAULT NOW()
        );`)
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS product_videos (
            id BIGSERIAL PRIMARY KEY,
            product_id BIGINT NOT NULL,
            title VARCHAR(255),
            description TEXT,
            video_url VARCHAR(255),
            thumbnail_url VARCHAR(255),
            show_in_gallery BOOLEAN DEFAULT FALSE,
            autoplay BOOLEAN DEFAULT FALSE,
            show_controls BOOLEAN DEFAULT TRUE,
            sort_order INT DEFAULT 0,
            created_at TIMESTAMPTZ DEFAULT NOW(),
            updated_at TIMESTAMPTZ DEFAULT NOW()
        );`)
		return nil
	})
}
