package migrations

import "gorm.io/gorm"

// Up089HeaderLayersPaddingFields مایگریشن خالی برای حفظ ترتیب شماره‌گذاری
// این مایگریشن هیچ تغییری در دیتابیس ایجاد نمی‌کند
func Up089HeaderLayersPaddingFields(db *gorm.DB) error {
	// هیچ عملیاتی انجام نمی‌شود
	// فقط برای حفظ ترتیب شماره‌گذاری مایگریشن‌ها
	return nil
}

// Down089HeaderLayersPaddingFields مایگریشن خالی برای حفظ ترتیب شماره‌گذاری
func Down089HeaderLayersPaddingFields(db *gorm.DB) error {
	// هیچ عملیاتی انجام نمی‌شود
	// فقط برای حفظ ترتیب شماره‌گذاری مایگریشن‌ها
	return nil
}
