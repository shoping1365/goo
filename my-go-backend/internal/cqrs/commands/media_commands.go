package commands

import (
	"time"
)

// CreateMediaFileCommand دستور ایجاد فایل رسانه جدید
// این دستور تمام اطلاعات لازم برای ایجاد یک فایل رسانه را در خود دارد
type CreateMediaFileCommand struct {
	FileName         string
	Title            string
	AltText          string
	DisplayTitle     string
	ShortDescription string
	Description      string
	FileType         string
	MimeType         string
	Format           string
	OriginalFormat   string
	Size             uint
	URL              string
	Category         string
	UploadedBy       *uint

	// فیلدهای مربوط به ویدیو
	DurationSeconds float64
	Width           int
	Height          int
	BitrateKbps     uint
}

// UpdateMediaFileCommand دستور به‌روزرسانی فایل رسانه
// فقط فیلدهایی که مقدار دارند به‌روزرسانی می‌شوند
type UpdateMediaFileCommand struct {
	ID               uint
	Title            *string
	AltText          *string
	DisplayTitle     *string
	ShortDescription *string
	Description      *string
	Category         *string
}

// DeleteMediaFileCommand دستور حذف فایل رسانه
// این دستور علاوه بر حذف فایل اصلی، تمام نسخه‌ها و job های مرتبط را نیز حذف می‌کند
type DeleteMediaFileCommand struct {
	ID uint
}

// CreateMediaVersionCommand دستور ایجاد نسخه جدید برای یک رسانه
type CreateMediaVersionCommand struct {
	MediaID           uint
	VersionCode       string
	FilePath          string
	FileSize          uint
	Format            string
	Quality           string
	IsActive          bool
	IsBackup          bool
	CompressionRatio  float32
	CompressionMethod string
	CreatedBy         *uint
	Meta              string
	ParentVersionID   *uint
	BackupPath        string
	Note              string
}

// CompressMediaFileCommand دستور فشرده‌سازی یک فایل رسانه
// این دستور یک job فشرده‌سازی ایجاد می‌کند و آن را در صف قرار می‌دهد
type CompressMediaFileCommand struct {
	MediaID       uint
	TargetFormat  string
	TargetQuality string
	RequestedBy   *uint
}

// CreateMediaWithVersionCommand دستور ایجاد فایل رسانه به همراه نسخه اولیه
// این دستور در یک تراکنش اتمیک اجرا می‌شود
type CreateMediaWithVersionCommand struct {
	MediaFile CreateMediaFileCommand
	Version   CreateMediaVersionCommand
}

// SetActiveVersionCommand دستور تنظیم یک نسخه به عنوان نسخه فعال
type SetActiveVersionCommand struct {
	MediaID   uint
	VersionID uint
}

// CreateBackupCommand دستور ایجاد بکاپ از یک نسخه
type CreateBackupCommand struct {
	VersionID uint
	CreatedBy *uint
}

// CleanupOldBackupsCommand دستور پاکسازی بکاپ‌های قدیمی
type CleanupOldBackupsCommand struct {
	MediaID   uint
	KeepCount int // تعداد بکاپ‌هایی که باید نگه داشته شوند
}

// BulkCreateMediaFilesCommand دستور ایجاد دسته‌ای فایل‌های رسانه
type BulkCreateMediaFilesCommand struct {
	Files []CreateMediaFileCommand
}

// UpdateMediaCategoryCommand دستور به‌روزرسانی دسته‌بندی یک فایل رسانه
type UpdateMediaCategoryCommand struct {
	ID       uint
	Category string
}

// CreateCompressionJobCommand دستور ایجاد job فشرده‌سازی
type CreateCompressionJobCommand struct {
	MediaID       uint
	VersionID     *uint
	RequestedBy   *uint
	TargetFormat  string
	TargetQuality string
	BatchID       *uint
	JobType       string
	OriginalSize  *uint
}

// UpdateCompressionJobCommand دستور به‌روزرسانی وضعیت job فشرده‌سازی
type UpdateCompressionJobCommand struct {
	ID             uint
	Status         *string
	StartedAt      *time.Time
	FinishedAt     *time.Time
	ErrorMessage   *string
	Progress       *int
	Log            *string
	CompressedSize *uint
}

// CreateCompressionSettingCommand دستور ایجاد تنظیمات فشرده‌سازی
type CreateCompressionSettingCommand struct {
	Scope     string
	UserID    *uint
	MediaID   *uint
	JobID     *uint
	Key       string
	Value     string
	IsDefault bool
}

// CommandResult نتیجه اجرای یک دستور
// این ساختار برای بازگرداندن نتیجه و خطاهای احتمالی استفاده می‌شود
type CommandResult struct {
	Success bool
	Error   error
	Data    interface{} // داده‌های اضافی در صورت نیاز
}

// MediaCommandResult نتیجه دستورات مربوط به رسانه
type MediaCommandResult struct {
	CommandResult
	MediaID   uint
	VersionID uint
}
