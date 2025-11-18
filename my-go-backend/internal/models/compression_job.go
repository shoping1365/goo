package models

import (
	"time"
)

// CompressionJob مدل صف و لاگ عملیات فشرده‌سازی و تبدیل فرمت
// این مدل برای مدیریت وضعیت، خطا، تنظیمات و گروه‌بندی عملیات استفاده می‌شود.
type CompressionJob struct {
	ID             uint       `gorm:"primaryKey" json:"id"`                            // کلید اصلی
	MediaID        uint       `gorm:"not null;index" json:"media_id"`                  // کلید خارجی به media_files
	VersionID      *uint      `gorm:"index" json:"version_id"`                         // کلید خارجی به media_versions (nullable)
	RequestedBy    *uint      `gorm:"index" json:"requested_by"`                       // کاربری که درخواست داده
	Status         string     `gorm:"type:varchar(20);not null;index" json:"status"`   // وضعیت عملیات
	StartedAt      *time.Time `json:"started_at"`                                      // زمان شروع عملیات
	FinishedAt     *time.Time `json:"finished_at"`                                     // زمان پایان عملیات
	ErrorMessage   string     `gorm:"type:text" json:"error_message"`                  // پیام خطا (در صورت وجود)
	TargetFormat   string     `gorm:"type:varchar(20)" json:"target_format"`           // فرمت خروجی
	TargetQuality  string     `gorm:"type:varchar(20)" json:"target_quality"`          // کیفیت خروجی
	BatchID        *uint      `gorm:"index" json:"batch_id"`                           // شناسه گروه (برای عملیات دسته‌ای)
	JobType        string     `gorm:"type:varchar(30);not null;index" json:"job_type"` // نوع عملیات
	Progress       int        `gorm:"default:0" json:"progress"`                       // درصد پیشرفت (۰ تا ۱۰۰)
	Log            string     `gorm:"type:text" json:"log"`                            // لاگ کامل عملیات
	OriginalSize   *uint      `gorm:"column:original_size" json:"original_size"`       // اندازه فایل اصلی
	CompressedSize *uint      `gorm:"column:compressed_size" json:"compressed_size"`   // اندازه فایل فشرده شده
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`

	// روابط
	Media MediaFile `gorm:"foreignKey:MediaID" json:"media,omitempty"` // رابطه با MediaFile
}
