package migrations

import "gorm.io/gorm"

// Up020ReviewsAndQuestionsDevice بهبود reviews و افزودن ستون‌های device به product_questions
func Up020ReviewsAndQuestionsDevice(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// review_helpful_votes
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS review_helpful_votes (
                id BIGSERIAL PRIMARY KEY,
                review_id BIGINT NOT NULL,
                user_id BIGINT NOT NULL
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_review_vote ON review_helpful_votes (review_id, user_id)`).Error; err != nil {
			return err
		}

		// اطمینان از وجود product_questions قبل از ALTER
		_ = tx.Exec(`
            CREATE TABLE IF NOT EXISTS product_questions (
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
                user_agent TEXT,
                device_info JSONB,
                created_at TIMESTAMPTZ DEFAULT NOW(),
                updated_at TIMESTAMPTZ DEFAULT NOW()
            );
        `)

		// product_questions: ستون‌های user_agent و device_info
		if err := tx.Exec(`ALTER TABLE product_questions ADD COLUMN IF NOT EXISTS user_agent TEXT`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`ALTER TABLE product_questions ADD COLUMN IF NOT EXISTS device_info JSONB`).Error; err != nil {
			return err
		}
		return nil
	})
}
