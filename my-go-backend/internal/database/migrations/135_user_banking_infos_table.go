package migrations

import (
	"gorm.io/gorm"
)

// Migration135AddUserBankingInfosTable - ایجاد جدول اطلاعات بانکی کاربران
// این جدول برای ذخیره اطلاعات بانکی کاربران شامل شماره کارت، نام بانک، شماره حساب و شماره شبا استفاده می‌شود
func Migration135AddUserBankingInfosTable(db *gorm.DB) error {
	return db.Exec(`
		-- ایجاد جدول user_banking_infos
		CREATE TABLE IF NOT EXISTS user_banking_infos (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			
			-- اطلاعات بانکی
			bank_name VARCHAR(100) NOT NULL,           -- نام بانک
			card_number VARCHAR(20) NOT NULL,           -- شماره کارت
			account_number VARCHAR(50) NOT NULL,        -- شماره حساب
			sheba_number VARCHAR(26),                     -- شماره شبا (اختیاری)
			account_holder_name VARCHAR(200),           -- نام صاحب حساب
			account_type VARCHAR(50),                   -- نوع حساب (جاری، قرض الحسنه، پس انداز و...)
			
			-- وضعیت پیشفرض
			is_default BOOLEAN DEFAULT FALSE,           -- آیا این حساب پیشفرض است
			
			-- وضعیت تایید
			is_verified BOOLEAN DEFAULT FALSE,          -- آیا اطلاعات تایید شده است
			
			-- اطلاعات تایید
			verified_at TIMESTAMPTZ,                   -- زمان تایید
			verified_by BIGINT,                        -- شناسه ادمین تایید کننده
			verification_note TEXT,                    -- یادداشت تایید
			
			-- زمان‌ها
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			
			-- کلیدهای خارجی
			CONSTRAINT fk_user_banking_infos_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT fk_user_banking_infos_verified_by FOREIGN KEY(verified_by) REFERENCES users(id) ON DELETE SET NULL
		);
		
		-- ایجاد ایندکس‌ها برای بهبود عملکرد
		CREATE INDEX IF NOT EXISTS idx_user_banking_infos_user_id ON user_banking_infos(user_id);
		CREATE INDEX IF NOT EXISTS idx_user_banking_infos_is_verified ON user_banking_infos(is_verified);
		CREATE INDEX IF NOT EXISTS idx_user_banking_infos_is_default ON user_banking_infos(is_default);
		CREATE INDEX IF NOT EXISTS idx_user_banking_infos_verified_by ON user_banking_infos(verified_by);
		CREATE INDEX IF NOT EXISTS idx_user_banking_infos_created_at ON user_banking_infos(created_at DESC);
		
		-- ایندکس برای جستجوی سریع بر اساس نام بانک
		CREATE INDEX IF NOT EXISTS idx_user_banking_infos_bank_name ON user_banking_infos(bank_name);
		
		-- کامنت‌ها برای مستندسازی
		COMMENT ON TABLE user_banking_infos IS 'جدول اطلاعات بانکی کاربران شامل شماره کارت، نام بانک، شماره حساب و شماره شبا';
		COMMENT ON COLUMN user_banking_infos.user_id IS 'شناسه کاربر';
		COMMENT ON COLUMN user_banking_infos.bank_name IS 'نام بانک (مثل بانک ملی، بانک سپه و...)';
		COMMENT ON COLUMN user_banking_infos.card_number IS 'شماره کارت بانکی';
		COMMENT ON COLUMN user_banking_infos.account_number IS 'شماره حساب بانکی';
		COMMENT ON COLUMN user_banking_infos.sheba_number IS 'شماره شبا (IBAN)';
		COMMENT ON COLUMN user_banking_infos.account_holder_name IS 'نام کامل صاحب حساب';
		COMMENT ON COLUMN user_banking_infos.account_type IS 'نوع حساب: جاری، قرض الحسنه، پس انداز، سپرده کوتاه مدت، سپرده بلند مدت';
		COMMENT ON COLUMN user_banking_infos.is_verified IS 'آیا اطلاعات بانکی توسط ادمین تایید شده است';
		COMMENT ON COLUMN user_banking_infos.verified_by IS 'شناسه ادمینی که اطلاعات را تایید کرده';
		COMMENT ON COLUMN user_banking_infos.verification_note IS 'یادداشت یا توضیحات مربوط به تایید اطلاعات';
	`).Error
}
