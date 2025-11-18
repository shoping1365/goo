package migrations

import "gorm.io/gorm"

// Up061FraudCoreV2 فراخوانی دوباره core fraud (در صورت نیاز محیط جدید)
func Up061FraudCoreV2(db *gorm.DB) error {
	return Up025FraudCoreAndIndexes(db)
}
