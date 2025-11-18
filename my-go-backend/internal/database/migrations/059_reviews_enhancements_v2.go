package migrations

import (
	"gorm.io/gorm"
)

// Up059ReviewsEnhancementsV2 تکمیل جدول review_helpful_votes و ستون‌های reviews
func Up059ReviewsEnhancementsV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// ایجاد ایمن جدول reviews در صورت نبود
		_ = tx.Exec(`
            CREATE TABLE IF NOT EXISTS reviews (
                id BIGSERIAL PRIMARY KEY,
                customer_id BIGINT,
                product_id BIGINT,
                rating INT CHECK (rating >= 1 AND rating <= 5),
                title VARCHAR(255),
                comment TEXT,
                pros TEXT,
                cons TEXT,
                status VARCHAR(20) DEFAULT 'pending',
                images JSONB,
                admin_reply TEXT,
                ip_address VARCHAR(45),
                user_agent TEXT,
                device_info JSONB,
                has_purchased BOOLEAN DEFAULT FALSE,
                helpful_count INT DEFAULT 0,
                created_at TIMESTAMPTZ DEFAULT NOW(),
                updated_at TIMESTAMPTZ DEFAULT NOW()
            );
        `)

		// helpful votes
		_ = tx.Exec("CREATE TABLE IF NOT EXISTS review_helpful_votes (id BIGSERIAL PRIMARY KEY, review_id BIGINT NOT NULL, user_id BIGINT NOT NULL)")
		_ = tx.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_review_vote ON review_helpful_votes (review_id, user_id)")

		// تضمین ستون‌های اضافی (idempotent)
		_ = tx.Exec("ALTER TABLE reviews ADD COLUMN IF NOT EXISTS title VARCHAR(255)")
		_ = tx.Exec("ALTER TABLE reviews ADD COLUMN IF NOT EXISTS pros TEXT")
		_ = tx.Exec("ALTER TABLE reviews ADD COLUMN IF NOT EXISTS cons TEXT")
		_ = tx.Exec("ALTER TABLE reviews ADD COLUMN IF NOT EXISTS user_agent TEXT")
		_ = tx.Exec("ALTER TABLE reviews ADD COLUMN IF NOT EXISTS device_info JSONB")
		return nil
	})
}
