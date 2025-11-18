package migrations

import (
	"gorm.io/gorm"
)

// Migration136UpdateUserBankingInfosTable - به‌روزرسانی جدول اطلاعات بانکی کاربران
// این migration شماره شبا را اختیاری می‌کند و نام صاحب حساب را اجباری می‌کند
func Migration136UpdateUserBankingInfosTable(db *gorm.DB) error {
	return db.Exec(`
		-- تغییر ستون sheba_number به اختیاری
		ALTER TABLE user_banking_infos ALTER COLUMN sheba_number DROP NOT NULL;
		
		-- تغییر ستون account_holder_name به اجباری
		ALTER TABLE user_banking_infos ALTER COLUMN account_holder_name SET NOT NULL;
		
		-- به‌روزرسانی کامنت‌ها
		COMMENT ON COLUMN user_banking_infos.sheba_number IS 'شماره شبا (IBAN) - اختیاری';
		COMMENT ON COLUMN user_banking_infos.account_holder_name IS 'نام کامل صاحب حساب - اجباری';
	`).Error
}
