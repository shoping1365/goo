package migrations

import "gorm.io/gorm"

// Up124PaymentGatewaysMissingColumns اضافه کردن ستون‌های گمشده به جدول payment_gateways
// این مایگریشن ستون‌های مربوط به embedded structs را اضافه می‌کند
func Up124PaymentGatewaysMissingColumns(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن ستون‌های PaymentGatewayApiKeys
		if err := tx.Exec(`
            ALTER TABLE payment_gateways 
            ADD COLUMN IF NOT EXISTS public_key VARCHAR(255),
            ADD COLUMN IF NOT EXISTS private_key VARCHAR(255),
            ADD COLUMN IF NOT EXISTS test_key VARCHAR(255),
            ADD COLUMN IF NOT EXISTS secret_key VARCHAR(255);
        `).Error; err != nil {
			return err
		}

		// اضافه کردن ستون‌های PaymentGatewayEndpoints
		if err := tx.Exec(`
            ALTER TABLE payment_gateways 
            ADD COLUMN IF NOT EXISTS payment VARCHAR(255),
            ADD COLUMN IF NOT EXISTS verification VARCHAR(255),
            ADD COLUMN IF NOT EXISTS refund VARCHAR(255),
            ADD COLUMN IF NOT EXISTS balance VARCHAR(255),
            ADD COLUMN IF NOT EXISTS callback VARCHAR(255);
        `).Error; err != nil {
			return err
		}

		// اضافه کردن ستون‌های PaymentGatewaySettings
		if err := tx.Exec(`
            ALTER TABLE payment_gateways 
            ADD COLUMN IF NOT EXISTS auto_redirect BOOLEAN NOT NULL DEFAULT TRUE,
            ADD COLUMN IF NOT EXISTS show_gateway_name BOOLEAN NOT NULL DEFAULT TRUE,
            ADD COLUMN IF NOT EXISTS require_card_info BOOLEAN NOT NULL DEFAULT FALSE,
            ADD COLUMN IF NOT EXISTS support_installment BOOLEAN NOT NULL DEFAULT FALSE,
            ADD COLUMN IF NOT EXISTS max_installments INT NOT NULL DEFAULT 1,
            ADD COLUMN IF NOT EXISTS currency VARCHAR(10) NOT NULL DEFAULT 'IRR',
            ADD COLUMN IF NOT EXISTS language VARCHAR(10) NOT NULL DEFAULT 'fa';
        `).Error; err != nil {
			return err
		}

		return nil
	})
}

// Down124PaymentGatewaysMissingColumns حذف ستون‌های اضافه شده
func Down124PaymentGatewaysMissingColumns(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف ستون‌های PaymentGatewayApiKeys
		if err := tx.Exec(`
            ALTER TABLE payment_gateways 
            DROP COLUMN IF EXISTS public_key,
            DROP COLUMN IF EXISTS private_key,
            DROP COLUMN IF EXISTS test_key,
            DROP COLUMN IF EXISTS secret_key;
        `).Error; err != nil {
			return err
		}

		// حذف ستون‌های PaymentGatewayEndpoints
		if err := tx.Exec(`
            ALTER TABLE payment_gateways 
            DROP COLUMN IF EXISTS payment,
            DROP COLUMN IF EXISTS verification,
            DROP COLUMN IF EXISTS refund,
            DROP COLUMN IF EXISTS balance,
            DROP COLUMN IF EXISTS callback;
        `).Error; err != nil {
			return err
		}

		// حذف ستون‌های PaymentGatewaySettings
		if err := tx.Exec(`
            ALTER TABLE payment_gateways 
            DROP COLUMN IF EXISTS auto_redirect,
            DROP COLUMN IF EXISTS show_gateway_name,
            DROP COLUMN IF EXISTS require_card_info,
            DROP COLUMN IF EXISTS support_installment,
            DROP COLUMN IF EXISTS max_installments,
            DROP COLUMN IF EXISTS currency,
            DROP COLUMN IF EXISTS language;
        `).Error; err != nil {
			return err
		}

		return nil
	})
}
