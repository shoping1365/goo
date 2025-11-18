package models

import (
	"time"

	"gorm.io/gorm"
)

// MediaFile مدل اصلی رسانه مطابق جدول media_files
// این مدل برای مدیریت انواع فایل‌ها (تصویر، ویدیو، صوت، سند و ...) استفاده می‌شود.
type MediaFile struct {
	ID               uint   `gorm:"primaryKey" json:"id"`                             // کلید اصلی (عددی)
	FileName         string `gorm:"type:varchar(255);not null" json:"file_name"`      // نام فایل اصلی
	Title            string `gorm:"type:varchar(150)" json:"title"`                   // عنوان رسانه
	AltText          string `gorm:"type:varchar(255)" json:"alt_text"`                // متن جایگزین
	DisplayTitle     string `gorm:"type:varchar(150)" json:"display_title"`           // عنوان نمایشی
	ShortDescription string `gorm:"type:varchar(255)" json:"short_description"`       // توضیحات مختصر
	Description      string `gorm:"type:text" json:"description"`                     // توضیحات کامل
	FileType         string `gorm:"type:varchar(30);not null;index" json:"file_type"` // نوع فایل
	MimeType         string `gorm:"type:varchar(50)" json:"mime_type"`                // نوع MIME
	Format           string `gorm:"type:varchar(20)" json:"format"`                   // فرمت فعلی فایل
	OriginalFormat   string `gorm:"type:varchar(20)" json:"original_format"`          // فرمت اولیه فایل
	Size             uint   `gorm:"not null" json:"size"`                             // حجم فایل اصلی (بایت)
	Compressed       bool   `gorm:"default:false;index" json:"compressed"`            // آیا فشرده‌سازی شده؟
	CompressedSize   *uint  `json:"compressed_size"`                                  // حجم بعد از فشرده‌سازی (nullable)
	URL              string `gorm:"type:varchar(255);not null" json:"url"`            // آدرس فایل (مسیر نسبی)
	// Meta (Video specific) — nullable برای تصاویر مقدار صفر است.
	DurationSeconds float64        `json:"duration_seconds"`                       // طول ویدیو (ثانیه)
	Width           int            `json:"width"`                                  // عرض ویدیو (پیکسل)
	Height          int            `json:"height"`                                 // ارتفاع ویدیو (پیکسل)
	BitrateKbps     uint           `json:"bitrate_kbps"`                           // نرخ بیت (kbps)
	Category        string         `gorm:"type:varchar(30);index" json:"category"` // دسته‌بندی رسانه
	UploadedBy      *uint          `gorm:"index" json:"uploaded_by"`               // user_id آپلودکننده (nullable)
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	// ارتباط با واریانت‌ها (در صورت وجود جدول واریانت)
	Variants []MediaVariant `gorm:"foreignKey:MediaID" json:"variants"`

	// ارتباط با jobهای فشرده‌سازی
	CompressionJobs []CompressionJob `gorm:"foreignKey:MediaID" json:"compression_jobs,omitempty"`
}

// MediaVariant مدل واریانت‌های رسانه (مثلاً سایزهای مختلف تصویر)
type MediaVariant struct {
	ID        uint           `gorm:"primaryKey" json:"id"`                        // کلید اصلی
	MediaID   uint           `gorm:"not null;index" json:"media_id"`              // کلید خارجی به media_files
	Width     int            `json:"width"`                                       // عرض تصویر
	Height    int            `json:"height"`                                      // ارتفاع تصویر
	FilePath  string         `gorm:"type:varchar(255);not null" json:"file_path"` // مسیر فایل واریانت (نسبی)
	FileSize  uint           `json:"file_size"`                                   // حجم فایل واریانت
	Purpose   string         `gorm:"type:varchar(30)" json:"purpose"`             // نوع واریانت (thumbnail, small, ...)
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
