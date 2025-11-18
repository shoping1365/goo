package queries

import (
	"time"
)

// GetMediaFileByIDQuery کوئری دریافت فایل رسانه بر اساس ID
type GetMediaFileByIDQuery struct {
	ID uint
}

// GetMediaFileByPathQuery کوئری دریافت فایل رسانه بر اساس مسیر
type GetMediaFileByPathQuery struct {
	Path string
}

// GetMediaFilesByUserQuery کوئری دریافت فایل‌های رسانه یک کاربر
type GetMediaFilesByUserQuery struct {
	UserID uint
}

// GetMediaFilesByTypeQuery کوئری دریافت فایل‌های رسانه بر اساس نوع
type GetMediaFilesByTypeQuery struct {
	FileType string
}

// GetMediaFilesByCategoryQuery کوئری دریافت فایل‌های رسانه بر اساس دسته‌بندی
type GetMediaFilesByCategoryQuery struct {
	Category string
}

// GetMediaFilesPagedQuery کوئری دریافت فایل‌های رسانه به صورت صفحه‌بندی شده
type GetMediaFilesPagedQuery struct {
	Page     int
	PageSize int
	Filter   MediaFileFilter
}

// MediaFileFilter فیلترهای جستجوی فایل‌های رسانه
type MediaFileFilter struct {
	FileType   string
	Category   string
	UserID     *uint
	Compressed *bool
	FromDate   *time.Time
	ToDate     *time.Time
	MinSize    *uint
	MaxSize    *uint
}

// SearchMediaFilesQuery کوئری جستجو در فایل‌های رسانه
type SearchMediaFilesQuery struct {
	Query string // عبارت جستجو در title، display_title و file_name
}

// GetRecentUploadsQuery کوئری دریافت آخرین فایل‌های آپلود شده
type GetRecentUploadsQuery struct {
	Limit int
}

// GetMediaFileStatsQuery کوئری دریافت آمار فایل‌های رسانه یک کاربر
type GetMediaFileStatsQuery struct {
	UserID uint
}

// GetMediaFileWithVariantsQuery کوئری دریافت فایل رسانه به همراه نسخه‌ها
type GetMediaFileWithVariantsQuery struct {
	ID uint
}

// GetMediaVersionByIDQuery کوئری دریافت نسخه رسانه بر اساس ID
type GetMediaVersionByIDQuery struct {
	ID uint
}

// GetMediaVersionsByMediaIDQuery کوئری دریافت تمام نسخه‌های یک رسانه
type GetMediaVersionsByMediaIDQuery struct {
	MediaID uint
}

// GetActiveMediaVersionQuery کوئری دریافت نسخه فعال یک رسانه
type GetActiveMediaVersionQuery struct {
	MediaID uint
}

// GetMediaBackupsQuery کوئری دریافت بکاپ‌های یک رسانه
type GetMediaBackupsQuery struct {
	MediaID uint
}

// GetMediaVersionsByFormatQuery کوئری دریافت نسخه‌های یک رسانه با فرمت خاص
type GetMediaVersionsByFormatQuery struct {
	MediaID uint
	Format  string
}

// GetMediaVersionsByQualityQuery کوئری دریافت نسخه‌های یک رسانه با کیفیت خاص
type GetMediaVersionsByQualityQuery struct {
	MediaID uint
	Quality string
}

// GetMediaVersionHistoryQuery کوئری دریافت تاریخچه نسخه‌های یک رسانه
type GetMediaVersionHistoryQuery struct {
	MediaID uint
}

// GetTotalSizeByMediaQuery کوئری محاسبه مجموع حجم نسخه‌های یک رسانه
type GetTotalSizeByMediaQuery struct {
	MediaID uint
}

// GetCompressionJobByIDQuery کوئری دریافت job فشرده‌سازی بر اساس ID
type GetCompressionJobByIDQuery struct {
	ID uint
}

// GetCompressionJobsByMediaQuery کوئری دریافت job های فشرده‌سازی یک رسانه
type GetCompressionJobsByMediaQuery struct {
	MediaID uint
}

// GetCompressionJobsByStatusQuery کوئری دریافت job ها بر اساس وضعیت
type GetCompressionJobsByStatusQuery struct {
	Status string
}

// GetPendingCompressionJobsQuery کوئری دریافت job های در انتظار
type GetPendingCompressionJobsQuery struct {
	Limit int
}

// GetCompressionSettingByIDQuery کوئری دریافت تنظیمات فشرده‌سازی بر اساس ID
type GetCompressionSettingByIDQuery struct {
	ID uint
}

// GetCompressionSettingsByScopeQuery کوئری دریافت تنظیمات بر اساس scope
type GetCompressionSettingsByScopeQuery struct {
	Scope string
}

// GetCompressionSettingsByUserQuery کوئری دریافت تنظیمات کاربر
type GetCompressionSettingsByUserQuery struct {
	UserID uint
}

// GetCompressionSettingsByMediaQuery کوئری دریافت تنظیمات رسانه
type GetCompressionSettingsByMediaQuery struct {
	MediaID uint
}

// GetCompressionStatsByTypeQuery کوئری دریافت آمار فشرده‌سازی بر اساس نوع
type GetCompressionStatsByTypeQuery struct {
	StatType string
}

// GetCompressionStatsByMediaQuery کوئری دریافت آمار فشرده‌سازی یک رسانه
type GetCompressionStatsByMediaQuery struct {
	MediaID uint
}

// GetCompressionStatsByUserQuery کوئری دریافت آمار فشرده‌سازی کاربر
type GetCompressionStatsByUserQuery struct {
	UserID uint
}

// GetMediaStorageReportQuery کوئری گزارش فضای ذخیره‌سازی
type GetMediaStorageReportQuery struct {
	GroupBy string // user, category, type, month
	Limit   int
}

// GetMediaUsageReportQuery کوئری گزارش استفاده از رسانه‌ها
type GetMediaUsageReportQuery struct {
	From    time.Time
	To      time.Time
	GroupBy string // daily, weekly, monthly
}

// GetCompressionEfficiencyReportQuery کوئری گزارش کارایی فشرده‌سازی
type GetCompressionEfficiencyReportQuery struct {
	From time.Time
	To   time.Time
}
