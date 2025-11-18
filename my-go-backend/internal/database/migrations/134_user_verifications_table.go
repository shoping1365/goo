package migrations

import (
	"gorm.io/gorm"
)

// Migration134AddUserVerificationsTable - ایجاد جدول احراز هویت کاربران (KYC)
// این جدول برای ذخیره درخواست‌های احراز هویت و تایید اطلاعات شخصی کاربران استفاده می‌شود
func Migration134AddUserVerificationsTable(db *gorm.DB) error {
	return db.Exec(`
		-- ایجاد جدول user_verifications
		CREATE TABLE IF NOT EXISTS user_verifications (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			
			-- نوع فیلد مورد نظر برای تایید
			field_name VARCHAR(50) NOT NULL,  -- 'national_code', 'sheba_number', 'birth_date', 'gender', 'job', 'economic_code', 'disability_type', 'legal_info'
			field_value TEXT NOT NULL,        -- مقدار فیلد (به صورت متنی یا JSON)
			
			-- وضعیت تایید
			status VARCHAR(20) NOT NULL DEFAULT 'pending',  -- 'pending', 'verified', 'rejected'
			
			-- اطلاعات تایید/رد
			verified_at TIMESTAMPTZ,
			verified_by BIGINT,  -- ID ادمینی که تایید کرده
			rejection_reason TEXT,
			
			-- مدارک پیوست (لیست URLها به صورت JSON)
			documents JSONB DEFAULT '[]'::jsonb,
			
			-- متاداده اضافی (برای ذخیره اطلاعات تکمیلی)
			metadata JSONB DEFAULT '{}'::jsonb,
			
			-- روش احراز هویت
			verification_method VARCHAR(20) DEFAULT 'manual',  -- 'manual', 'api', 'ekyc'
			
			-- اطلاعات API (اگر از API استفاده شده باشد)
			api_provider VARCHAR(50),     -- 'raypa', 'fanap', 'rahkareshna', ...
			api_response JSONB,
			api_verified_at TIMESTAMPTZ,
			
			-- زمان‌ها
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			
			-- کلیدهای خارجی
			CONSTRAINT fk_user_verifications_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT fk_user_verifications_verified_by FOREIGN KEY(verified_by) REFERENCES users(id) ON DELETE SET NULL
		);
		
		-- ایجاد ایندکس‌ها برای بهبود عملکرد
		CREATE INDEX IF NOT EXISTS idx_user_verifications_user_id ON user_verifications(user_id);
		CREATE INDEX IF NOT EXISTS idx_user_verifications_status ON user_verifications(status);
		CREATE INDEX IF NOT EXISTS idx_user_verifications_field_name ON user_verifications(field_name);
		CREATE INDEX IF NOT EXISTS idx_user_verifications_user_field ON user_verifications(user_id, field_name);
		CREATE INDEX IF NOT EXISTS idx_user_verifications_verified_by ON user_verifications(verified_by);
		CREATE INDEX IF NOT EXISTS idx_user_verifications_created_at ON user_verifications(created_at DESC);
		
		-- کامنت‌ها برای مستندسازی
		COMMENT ON TABLE user_verifications IS 'جدول احراز هویت و تایید اطلاعات شخصی کاربران (KYC)';
		COMMENT ON COLUMN user_verifications.user_id IS 'شناسه کاربر';
		COMMENT ON COLUMN user_verifications.field_name IS 'نام فیلد مورد نظر برای تایید (کد ملی، شبا، تاریخ تولد و...)';
		COMMENT ON COLUMN user_verifications.field_value IS 'مقدار فیلد (به صورت متنی یا JSON)';
		COMMENT ON COLUMN user_verifications.status IS 'وضعیت: pending (در انتظار), verified (تایید شده), rejected (رد شده)';
		COMMENT ON COLUMN user_verifications.verified_by IS 'شناسه ادمینی که اطلاعات را تایید یا رد کرده';
		COMMENT ON COLUMN user_verifications.rejection_reason IS 'دلیل رد شدن درخواست';
		COMMENT ON COLUMN user_verifications.documents IS 'لیست URLهای مدارک پیوست شده (آرایه JSON)';
		COMMENT ON COLUMN user_verifications.verification_method IS 'روش احراز هویت: manual (دستی توسط ادمین), api (از طریق سرویس خارجی), ekyc (احراز هویت چهره‌ای)';
		COMMENT ON COLUMN user_verifications.api_provider IS 'نام ارائه‌دهنده سرویس احراز هویت (در صورت استفاده از API)';
		COMMENT ON COLUMN user_verifications.api_response IS 'پاسخ کامل API (برای لاگ و بررسی)';
	`).Error
}

