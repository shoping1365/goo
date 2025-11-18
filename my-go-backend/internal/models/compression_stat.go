package models

import (
	"time"

	"gorm.io/gorm"
)

// CompressionStat مدل آمار و گزارش‌های فشرده‌سازی
// این مدل برای ذخیره و تحلیل انواع آمار (کاهش حجم، تعداد تصاویر فشرده نشده، بازده فرمت‌ها و ...) استفاده می‌شود.
type CompressionStat struct {
	ID        uint           `gorm:"primaryKey" json:"id"`                             // کلید اصلی
	StatType  string         `gorm:"type:varchar(50);not null;index" json:"stat_type"` // نوع آمار
	Value     string         `gorm:"type:varchar(100);not null" json:"value"`          // مقدار آمار (عدد یا رشته)
	MediaID   *uint          `gorm:"index" json:"media_id"`                            // کلید خارجی به media_files (nullable)
	UserID    *uint          `gorm:"index" json:"user_id"`                             // کلید خارجی به users (nullable)
	Format    string         `gorm:"type:varchar(20);index" json:"format"`             // فرمت مورد نظر (nullable)
	Period    string         `gorm:"type:varchar(20);index" json:"period"`             // بازه زمانی (مثلاً 2024-06)
	Meta      string         `gorm:"type:json" json:"meta"`                            // اطلاعات اضافی (JSON)
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
