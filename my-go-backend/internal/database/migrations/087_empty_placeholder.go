package migrations

import "gorm.io/gorm"

// Up087EmptyPlaceholder مایگریشن خالی برای حفظ ترتیب شمارهگذاری
// این مایگریشن هیچ تغییری در دیتابیس ایجاد نمی‌کند
func Up087EmptyPlaceholder(db *gorm.DB) error {
	// هیچ عملیاتی انجام نمی‌شود
	// فقط برای حفظ ترتیب شمارهگذاری مایگریشن‌ها
	return nil
}

// Down087EmptyPlaceholder مایگریشن خالی برای حفظ ترتیب شمارهگذاری
func Down087EmptyPlaceholder(db *gorm.DB) error {
	// هیچ عملیاتی انجام نمی‌شود
	// فقط برای حفظ ترتیب شمارهگذاری مایگریشن‌ها
	return nil
}
