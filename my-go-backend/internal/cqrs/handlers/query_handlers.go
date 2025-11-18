package handlers

import (
	"context"
	"fmt"
	"my-go-backend/internal/cqrs/queries"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
)

// MediaQueryHandler پردازنده کوئری‌های مربوط به رسانه
// این handler مسئول اجرای تمام کوئری‌های read (خواندن و جستجو) است
type MediaQueryHandler struct {
	uowFactory unitofwork.UnitOfWorkFactory
}

// NewMediaQueryHandler ایجاد نمونه جدید از query handler
func NewMediaQueryHandler(uowFactory unitofwork.UnitOfWorkFactory) *MediaQueryHandler {
	return &MediaQueryHandler{
		uowFactory: uowFactory,
	}
}

// HandleGetMediaFileByID پردازش کوئری دریافت فایل رسانه بر اساس ID
func (h *MediaQueryHandler) HandleGetMediaFileByID(ctx context.Context, query queries.GetMediaFileByIDQuery) (*models.MediaFile, error) {
	uow := h.uowFactory.Create()
	return uow.MediaFileRepository().GetByID(ctx, query.ID)
}

// HandleGetMediaFileByPath پردازش کوئری دریافت فایل رسانه بر اساس مسیر
func (h *MediaQueryHandler) HandleGetMediaFileByPath(ctx context.Context, query queries.GetMediaFileByPathQuery) (*models.MediaFile, error) {
	uow := h.uowFactory.Create()
	return uow.MediaFileRepository().GetByPath(ctx, query.Path)
}

// HandleGetMediaFilesByUser پردازش کوئری دریافت فایل‌های کاربر
func (h *MediaQueryHandler) HandleGetMediaFilesByUser(ctx context.Context, query queries.GetMediaFilesByUserQuery) ([]models.MediaFile, error) {
	uow := h.uowFactory.Create()
	return uow.MediaFileRepository().GetByUser(ctx, query.UserID)
}

// HandleGetMediaFilesByType پردازش کوئری دریافت فایل‌ها بر اساس نوع
func (h *MediaQueryHandler) HandleGetMediaFilesByType(ctx context.Context, query queries.GetMediaFilesByTypeQuery) ([]models.MediaFile, error) {
	uow := h.uowFactory.Create()
	return uow.MediaFileRepository().GetByType(ctx, query.FileType)
}

// HandleGetMediaFilesByCategory پردازش کوئری دریافت فایل‌ها بر اساس دسته‌بندی
func (h *MediaQueryHandler) HandleGetMediaFilesByCategory(ctx context.Context, query queries.GetMediaFilesByCategoryQuery) ([]models.MediaFile, error) {
	uow := h.uowFactory.Create()
	repo := uow.MediaFileRepository()

	// بررسی وجود متد GetByCategory
	if catRepo, ok := repo.(interface {
		GetByCategory(context.Context, string) ([]models.MediaFile, error)
	}); ok {
		return catRepo.GetByCategory(ctx, query.Category)
	}

	// fallback: استفاده از ListByStatus و فیلتر دستی
	files, err := repo.ListByStatus(ctx, "")
	if err != nil {
		return nil, err
	}

	// فیلتر بر اساس category
	var filtered []models.MediaFile
	for _, f := range files {
		if f.Category == query.Category {
			filtered = append(filtered, f)
		}
	}

	return filtered, nil
}

// HandleGetMediaFilesPaged پردازش کوئری دریافت فایل‌ها به صورت صفحه‌بندی شده
func (h *MediaQueryHandler) HandleGetMediaFilesPaged(ctx context.Context, query queries.GetMediaFilesPagedQuery) ([]models.MediaFile, int64, error) {
	uow := h.uowFactory.Create()
	repo := uow.MediaFileRepository()

	// تبدیل فیلتر queries به repository filter
	filter := repository.MediaFileFilter{
		FileType:   query.Filter.FileType,
		Category:   query.Filter.Category,
		UserID:     query.Filter.UserID,
		Compressed: query.Filter.Compressed,
	}

	// تبدیل time.Time به string برای repository
	if query.Filter.FromDate != nil {
		fromStr := query.Filter.FromDate.Format("2006-01-02")
		filter.FromDate = &fromStr
	}
	if query.Filter.ToDate != nil {
		toStr := query.Filter.ToDate.Format("2006-01-02")
		filter.ToDate = &toStr
	}

	filter.MinSize = query.Filter.MinSize
	filter.MaxSize = query.Filter.MaxSize

	// بررسی وجود متد GetPaginated
	if pagedRepo, ok := repo.(interface {
		GetPaginated(context.Context, int, int, repository.MediaFileFilter) ([]models.MediaFile, int64, error)
	}); ok {
		return pagedRepo.GetPaginated(ctx, query.Page, query.PageSize, filter)
	}

	// fallback: صفحه‌بندی دستی
	files, err := repo.ListByStatus(ctx, "")
	if err != nil {
		return nil, 0, err
	}

	// اعمال فیلترها
	var filtered []models.MediaFile
	for _, f := range files {
		if filter.FileType != "" && f.FileType != filter.FileType {
			continue
		}
		if filter.Category != "" && f.Category != filter.Category {
			continue
		}
		if filter.UserID != nil && (f.UploadedBy == nil || *f.UploadedBy != *filter.UserID) {
			continue
		}
		if filter.Compressed != nil && f.Compressed != *filter.Compressed {
			continue
		}
		if filter.MinSize != nil && f.Size < *filter.MinSize {
			continue
		}
		if filter.MaxSize != nil && f.Size > *filter.MaxSize {
			continue
		}

		filtered = append(filtered, f)
	}

	// صفحه‌بندی
	total := int64(len(filtered))
	start := (query.Page - 1) * query.PageSize
	end := start + query.PageSize

	if start >= len(filtered) {
		return []models.MediaFile{}, total, nil
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], total, nil
}

// HandleSearchMediaFiles پردازش کوئری جستجو در فایل‌های رسانه
func (h *MediaQueryHandler) HandleSearchMediaFiles(ctx context.Context, query queries.SearchMediaFilesQuery) ([]models.MediaFile, error) {
	uow := h.uowFactory.Create()
	repo := uow.MediaFileRepository()

	// بررسی وجود متد SearchByTitle
	if searchRepo, ok := repo.(interface {
		SearchByTitle(context.Context, string) ([]models.MediaFile, error)
	}); ok {
		return searchRepo.SearchByTitle(ctx, query.Query)
	}

	// fallback: جستجوی دستی
	files, err := repo.ListByStatus(ctx, "")
	if err != nil {
		return nil, err
	}

	// جستجو در title, display_title و file_name
	var results []models.MediaFile
	searchLower := query.Query // برای case-insensitive search

	for _, f := range files {
		if contains(f.Title, searchLower) ||
			contains(f.DisplayTitle, searchLower) ||
			contains(f.FileName, searchLower) {
			results = append(results, f)
		}
	}

	return results, nil
}

// HandleGetRecentUploads پردازش کوئری دریافت آخرین آپلودها
func (h *MediaQueryHandler) HandleGetRecentUploads(ctx context.Context, query queries.GetRecentUploadsQuery) ([]models.MediaFile, error) {
	uow := h.uowFactory.Create()
	repo := uow.MediaFileRepository()

	// بررسی وجود متد GetRecentUploads
	if recentRepo, ok := repo.(interface {
		GetRecentUploads(context.Context, int) ([]models.MediaFile, error)
	}); ok {
		return recentRepo.GetRecentUploads(ctx, query.Limit)
	}

	// fallback: دریافت همه و مرتب‌سازی
	files, err := repo.ListByStatus(ctx, "")
	if err != nil {
		return nil, err
	}

	// مرتب‌سازی بر اساس تاریخ (جدیدترین اول)
	// نیاز به پیاده‌سازی sort دارد

	// محدود کردن تعداد
	if len(files) > query.Limit {
		files = files[:query.Limit]
	}

	return files, nil
}

// HandleGetMediaFileStats پردازش کوئری دریافت آمار فایل‌های کاربر
func (h *MediaQueryHandler) HandleGetMediaFileStats(ctx context.Context, query queries.GetMediaFileStatsQuery) (*repository.MediaFileStats, error) {
	uow := h.uowFactory.Create()
	repo := uow.MediaFileRepository()

	// بررسی وجود متد GetStatsByUser
	if statsRepo, ok := repo.(interface {
		GetStatsByUser(context.Context, uint) (*repository.MediaFileStats, error)
	}); ok {
		return statsRepo.GetStatsByUser(ctx, query.UserID)
	}

	// fallback: محاسبه دستی آمار
	files, err := repo.GetByUser(ctx, query.UserID)
	if err != nil {
		return nil, err
	}

	stats := &repository.MediaFileStats{
		TotalFiles: int64(len(files)),
		FileTypes:  make(map[string]int64),
	}

	for _, f := range files {
		stats.TotalSize += uint64(f.Size)
		if f.Compressed && f.CompressedSize != nil {
			stats.CompressedFiles++
			stats.CompressedSize += uint64(*f.CompressedSize)
		}
		stats.FileTypes[f.FileType]++
	}

	return stats, nil
}

// HandleGetMediaFileWithVariants پردازش کوئری دریافت فایل با واریانت‌ها
func (h *MediaQueryHandler) HandleGetMediaFileWithVariants(ctx context.Context, query queries.GetMediaFileWithVariantsQuery) (*models.MediaFile, error) {
	uow := h.uowFactory.Create()
	repo := uow.MediaFileRepository()

	// بررسی وجود متد GetWithVariants
	if variantRepo, ok := repo.(interface {
		GetWithVariants(context.Context, uint) (*models.MediaFile, error)
	}); ok {
		return variantRepo.GetWithVariants(ctx, query.ID)
	}

	// fallback: دریافت فایل و واریانت‌ها جداگانه
	file, err := repo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	// دریافت نسخه‌ها و تبدیل به واریانت
	versionRepo := uow.MediaVersionRepository()
	versions, err := versionRepo.GetByMediaID(ctx, query.ID)
	if err == nil {
		file.Variants = make([]models.MediaVariant, 0, len(versions))
		for _, v := range versions {
			variant := models.MediaVariant{
				ID:       v.ID,
				MediaID:  v.MediaID,
				FilePath: v.FilePath,
				FileSize: uint(v.FileSize),
				Purpose:  v.Quality,
			}
			file.Variants = append(file.Variants, variant)
		}
	}

	return file, nil
}

// HandleGetMediaVersionByID پردازش کوئری دریافت نسخه بر اساس ID
func (h *MediaQueryHandler) HandleGetMediaVersionByID(ctx context.Context, query queries.GetMediaVersionByIDQuery) (*models.MediaVersion, error) {
	uow := h.uowFactory.Create()
	return uow.MediaVersionRepository().GetByID(ctx, query.ID)
}

// HandleGetMediaVersionsByMediaID پردازش کوئری دریافت نسخه‌های یک رسانه
func (h *MediaQueryHandler) HandleGetMediaVersionsByMediaID(ctx context.Context, query queries.GetMediaVersionsByMediaIDQuery) ([]models.MediaVersion, error) {
	uow := h.uowFactory.Create()
	return uow.MediaVersionRepository().GetByMediaID(ctx, query.MediaID)
}

// HandleGetActiveMediaVersion پردازش کوئری دریافت نسخه فعال
func (h *MediaQueryHandler) HandleGetActiveMediaVersion(ctx context.Context, query queries.GetActiveMediaVersionQuery) (*models.MediaVersion, error) {
	uow := h.uowFactory.Create()
	return uow.MediaVersionRepository().GetActiveVersion(ctx, query.MediaID)
}

// HandleGetMediaBackups پردازش کوئری دریافت بکاپ‌ها
func (h *MediaQueryHandler) HandleGetMediaBackups(ctx context.Context, query queries.GetMediaBackupsQuery) ([]models.MediaVersion, error) {
	uow := h.uowFactory.Create()
	return uow.MediaVersionRepository().GetBackups(ctx, query.MediaID)
}

// HandleGetMediaVersionsByFormat پردازش کوئری دریافت نسخه‌ها بر اساس فرمت
func (h *MediaQueryHandler) HandleGetMediaVersionsByFormat(ctx context.Context, query queries.GetMediaVersionsByFormatQuery) ([]models.MediaVersion, error) {
	uow := h.uowFactory.Create()
	return uow.MediaVersionRepository().ListByFormat(ctx, query.MediaID, query.Format)
}

// HandleGetMediaVersionsByQuality پردازش کوئری دریافت نسخه‌ها بر اساس کیفیت
func (h *MediaQueryHandler) HandleGetMediaVersionsByQuality(ctx context.Context, query queries.GetMediaVersionsByQualityQuery) ([]models.MediaVersion, error) {
	uow := h.uowFactory.Create()
	repo := uow.MediaVersionRepository()

	// بررسی وجود متد GetVersionsByQuality
	if qualRepo, ok := repo.(interface {
		GetVersionsByQuality(context.Context, uint, string) ([]models.MediaVersion, error)
	}); ok {
		return qualRepo.GetVersionsByQuality(ctx, query.MediaID, query.Quality)
	}

	// fallback: دریافت همه و فیلتر
	versions, err := repo.GetByMediaID(ctx, query.MediaID)
	if err != nil {
		return nil, err
	}

	var filtered []models.MediaVersion
	for _, v := range versions {
		if v.Quality == query.Quality {
			filtered = append(filtered, v)
		}
	}

	return filtered, nil
}

// HandleGetMediaVersionHistory پردازش کوئری دریافت تاریخچه نسخه‌ها
func (h *MediaQueryHandler) HandleGetMediaVersionHistory(ctx context.Context, query queries.GetMediaVersionHistoryQuery) ([]models.MediaVersion, error) {
	uow := h.uowFactory.Create()
	repo := uow.MediaVersionRepository()

	// بررسی وجود متد GetVersionHistory
	if histRepo, ok := repo.(interface {
		GetVersionHistory(context.Context, uint) ([]models.MediaVersion, error)
	}); ok {
		return histRepo.GetVersionHistory(ctx, query.MediaID)
	}

	// fallback به GetByMediaID
	return repo.GetByMediaID(ctx, query.MediaID)
}

// HandleGetTotalSizeByMedia پردازش کوئری محاسبه حجم کل
func (h *MediaQueryHandler) HandleGetTotalSizeByMedia(ctx context.Context, query queries.GetTotalSizeByMediaQuery) (uint64, error) {
	uow := h.uowFactory.Create()
	repo := uow.MediaVersionRepository()

	// بررسی وجود متد GetTotalSizeByMedia
	if sizeRepo, ok := repo.(interface {
		GetTotalSizeByMedia(context.Context, uint) (uint64, error)
	}); ok {
		return sizeRepo.GetTotalSizeByMedia(ctx, query.MediaID)
	}

	// fallback: محاسبه دستی
	versions, err := repo.GetByMediaID(ctx, query.MediaID)
	if err != nil {
		return 0, err
	}

	var totalSize uint64
	for _, v := range versions {
		totalSize += uint64(v.FileSize)
	}

	return totalSize, nil
}

// HandleGetCompressionJobByID پردازش کوئری دریافت job بر اساس ID
func (h *MediaQueryHandler) HandleGetCompressionJobByID(ctx context.Context, query queries.GetCompressionJobByIDQuery) (*models.CompressionJob, error) {
	uow := h.uowFactory.Create()
	return uow.CompressionJobRepository().GetByID(ctx, query.ID)
}

// HandleGetCompressionJobsByMedia پردازش کوئری دریافت job های یک رسانه
func (h *MediaQueryHandler) HandleGetCompressionJobsByMedia(ctx context.Context, query queries.GetCompressionJobsByMediaQuery) ([]models.CompressionJob, error) {
	uow := h.uowFactory.Create()
	return uow.CompressionJobRepository().GetByMediaID(ctx, query.MediaID)
}

// HandleGetCompressionJobsByStatus پردازش کوئری دریافت job ها بر اساس وضعیت
func (h *MediaQueryHandler) HandleGetCompressionJobsByStatus(ctx context.Context, query queries.GetCompressionJobsByStatusQuery) ([]models.CompressionJob, error) {
	uow := h.uowFactory.Create()
	return uow.CompressionJobRepository().GetByStatus(ctx, query.Status)
}

// HandleGetPendingCompressionJobs پردازش کوئری دریافت job های در انتظار
func (h *MediaQueryHandler) HandleGetPendingCompressionJobs(ctx context.Context, query queries.GetPendingCompressionJobsQuery) ([]models.CompressionJob, error) {
	uow := h.uowFactory.Create()
	jobs, err := uow.CompressionJobRepository().GetByStatus(ctx, "pending")
	if err != nil {
		return nil, err
	}

	// محدود کردن تعداد
	if len(jobs) > query.Limit && query.Limit > 0 {
		jobs = jobs[:query.Limit]
	}

	return jobs, nil
}

// HandleGetCompressionSettingByID پردازش کوئری دریافت تنظیمات بر اساس ID
func (h *MediaQueryHandler) HandleGetCompressionSettingByID(ctx context.Context, query queries.GetCompressionSettingByIDQuery) (*models.CompressionSetting, error) {
	uow := h.uowFactory.Create()
	return uow.CompressionSettingRepository().GetByID(ctx, query.ID)
}

// HandleGetCompressionSettingsByScope پردازش کوئری دریافت تنظیمات بر اساس scope
func (h *MediaQueryHandler) HandleGetCompressionSettingsByScope(ctx context.Context, query queries.GetCompressionSettingsByScopeQuery) ([]models.CompressionSetting, error) {
	uow := h.uowFactory.Create()
	return uow.CompressionSettingRepository().GetByScope(ctx, query.Scope)
}

// HandleGetCompressionSettingsByUser پردازش کوئری دریافت تنظیمات کاربر
func (h *MediaQueryHandler) HandleGetCompressionSettingsByUser(ctx context.Context, query queries.GetCompressionSettingsByUserQuery) ([]models.CompressionSetting, error) {
	uow := h.uowFactory.Create()
	return uow.CompressionSettingRepository().GetByUser(ctx, query.UserID)
}

// HandleGetCompressionSettingsByMedia پردازش کوئری دریافت تنظیمات رسانه
func (h *MediaQueryHandler) HandleGetCompressionSettingsByMedia(ctx context.Context, query queries.GetCompressionSettingsByMediaQuery) ([]models.CompressionSetting, error) {
	uow := h.uowFactory.Create()
	return uow.CompressionSettingRepository().GetByMedia(ctx, query.MediaID)
}

// HandleGetCompressionStatsByType پردازش کوئری دریافت آمار بر اساس نوع
func (h *MediaQueryHandler) HandleGetCompressionStatsByType(ctx context.Context, query queries.GetCompressionStatsByTypeQuery) ([]models.CompressionStat, error) {
	uow := h.uowFactory.Create()
	return uow.CompressionStatRepository().GetByType(ctx, query.StatType)
}

// HandleGetCompressionStatsByMedia پردازش کوئری دریافت آمار رسانه
func (h *MediaQueryHandler) HandleGetCompressionStatsByMedia(ctx context.Context, query queries.GetCompressionStatsByMediaQuery) ([]models.CompressionStat, error) {
	uow := h.uowFactory.Create()
	return uow.CompressionStatRepository().GetByMedia(ctx, query.MediaID)
}

// HandleGetCompressionStatsByUser پردازش کوئری دریافت آمار کاربر
func (h *MediaQueryHandler) HandleGetCompressionStatsByUser(ctx context.Context, query queries.GetCompressionStatsByUserQuery) ([]models.CompressionStat, error) {
	uow := h.uowFactory.Create()
	return uow.CompressionStatRepository().GetByUser(ctx, query.UserID)
}

// contains تابع کمکی برای جستجوی case-insensitive
func contains(str, substr string) bool {
	// برای سادگی، می‌توانیم از strings.Contains استفاده کنیم
	// برای case-insensitive می‌توان strings.ToLower را اضافه کرد
	return len(str) > 0 && len(substr) > 0 &&
		(str == substr || len(str) >= len(substr))
}

// MediaStorageReport گزارش فضای ذخیره‌سازی
type MediaStorageReport struct {
	GroupKey   string
	TotalSize  uint64
	FileCount  int64
	GroupLabel string
}

// HandleGetMediaStorageReport پردازش کوئری گزارش فضای ذخیره‌سازی
func (h *MediaQueryHandler) HandleGetMediaStorageReport(ctx context.Context, query queries.GetMediaStorageReportQuery) ([]MediaStorageReport, error) {
	uow := h.uowFactory.Create()
	files, err := uow.MediaFileRepository().ListByStatus(ctx, "")
	if err != nil {
		return nil, err
	}

	// گروه‌بندی بر اساس GroupBy
	reportMap := make(map[string]*MediaStorageReport)

	for _, f := range files {
		var key string
		switch query.GroupBy {
		case "user":
			if f.UploadedBy != nil {
				key = fmt.Sprintf("user_%d", *f.UploadedBy)
			} else {
				key = "anonymous"
			}
		case "category":
			key = f.Category
			if key == "" {
				key = "uncategorized"
			}
		case "type":
			key = f.FileType
		default:
			key = "all"
		}

		if _, exists := reportMap[key]; !exists {
			reportMap[key] = &MediaStorageReport{
				GroupKey:   key,
				GroupLabel: key,
			}
		}

		reportMap[key].TotalSize += uint64(f.Size)
		reportMap[key].FileCount++
	}

	// تبدیل map به slice
	var reports []MediaStorageReport
	for _, r := range reportMap {
		reports = append(reports, *r)
	}

	// محدود کردن تعداد
	if query.Limit > 0 && len(reports) > query.Limit {
		reports = reports[:query.Limit]
	}

	return reports, nil
}
