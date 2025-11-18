package models

import (
	"time"

	"gorm.io/gorm"
)

// MediaVersion مدل نسخه‌های مختلف هر رسانه (تصویر/ویدیو/...)
// این مدل برای مدیریت نسخه‌های فشرده، تغییر فرمت، بکاپ و ... استفاده می‌شود.
type MediaVersion struct {
	ID                uint           `gorm:"primaryKey" json:"id"`                                // کلید اصلی
	MediaID           uint           `gorm:"not null;index" json:"media_id"`                      // کلید خارجی به media_files
	VersionCode       string         `gorm:"type:varchar(50);not null;index" json:"version_code"` // کد نسخه (original, v1, webp, ...)
	FilePath          string         `gorm:"type:varchar(255);not null" json:"file_path"`         // مسیر فایل این نسخه
	FileSize          uint           `json:"file_size"`                                           // حجم فایل این نسخه
	Format            string         `gorm:"type:varchar(20)" json:"format"`                      // فرمت این نسخه
	Quality           string         `gorm:"type:varchar(20)" json:"quality"`                     // کیفیت (مثلاً 80% یا high/medium/low)
	IsActive          bool           `gorm:"default:false;index" json:"is_active"`                // آیا نسخه فعال است؟
	IsBackup          bool           `gorm:"default:false;index" json:"is_backup"`                // آیا بکاپ است؟
	CompressionRatio  float32        `json:"compression_ratio"`                                   // درصد کاهش حجم نسبت به نسخه اصلی
	CompressionMethod string         `gorm:"type:varchar(50)" json:"compression_method"`          // الگوریتم یا ابزار فشرده‌سازی
	CreatedBy         *uint          `gorm:"index" json:"created_by"`                             // کاربری که این نسخه را ایجاد کرده
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	Meta              string         `gorm:"type:json" json:"meta"`                // اطلاعات اضافی (JSON: ابعاد، رنگ، EXIF، ...)
	ParentVersionID   *uint          `gorm:"index" json:"parent_version_id"`       // اگر این نسخه بر اساس نسخه دیگری ساخته شده
	BackupPath        string         `gorm:"type:varchar(255)" json:"backup_path"` // مسیر بکاپ (در صورت وجود)
	Note              string         `gorm:"type:text" json:"note"`                // توضیحات یا لاگ اضافی
}
