package services

import (
	"context"
	"encoding/json"
	"fmt"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

// MediaFileService سرویس مدیریت فایل‌های رسانه‌ای (تصویر، ویدیو و ...)
// این سرویس بازطراحی شده تا از UnitOfWork برای مدیریت تراکنش‌ها استفاده کند
// عملیات CRUD، جستجو، کش و عملیات پیچیده را برای فایل‌های رسانه‌ای فراهم می‌کند
type MediaFileService struct {
	uowFactory unitofwork.UnitOfWorkFactory // factory برای ایجاد unit of work
	redis      *redis.Client                // optional, for caching
}

// NewMediaFileService سازنده سرویس رسانه با پشتیبانی از UnitOfWork
// uowFactory: برای ایجاد unit of work جدید در هر عملیات
// redis: کلاینت Redis برای کش (اختیاری)
func NewMediaFileService(uowFactory unitofwork.UnitOfWorkFactory, redis *redis.Client) *MediaFileService {
	return &MediaFileService{
		uowFactory: uowFactory,
		redis:      redis,
	}
}

// Create ایجاد یک فایل رسانه‌ای جدید
// این متد تراکنش را به صورت داخلی مدیریت می‌کند
func (s *MediaFileService) Create(ctx context.Context, file *models.MediaFile) error {
	if file == nil {
		return fmt.Errorf("فایل رسانه نمی‌تواند nil باشد")
	}

	// شروع unit of work
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// ایجاد فایل رسانه
	mediaRepo := uow.MediaFileRepository()
	if err := mediaRepo.Create(ctx, file); err != nil {
		return fmt.Errorf("خطا در ایجاد فایل رسانه: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true

	// پاکسازی کش در صورت وجود
	s.invalidateCache(ctx, file.ID)

	return nil
}

// CreateWithVersion ایجاد فایل رسانه به همراه نسخه اولیه
// این متد برای زمانی است که می‌خواهیم فایل و نسخه را در یک تراکنش ایجاد کنیم
func (s *MediaFileService) CreateWithVersion(ctx context.Context, file *models.MediaFile, version *models.MediaVersion) error {
	if file == nil || version == nil {
		return fmt.Errorf("فایل رسانه و نسخه نمی‌توانند nil باشند")
	}

	// شروع unit of work
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// ایجاد فایل رسانه
	mediaRepo := uow.MediaFileRepository()
	if err := mediaRepo.Create(ctx, file); err != nil {
		return fmt.Errorf("خطا در ایجاد فایل رسانه: %w", err)
	}

	// تنظیم MediaID برای نسخه
	version.MediaID = file.ID

	// ایجاد نسخه
	versionRepo := uow.MediaVersionRepository()
	if err := versionRepo.Create(ctx, version); err != nil {
		return fmt.Errorf("خطا در ایجاد نسخه: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true

	// پاکسازی کش
	s.invalidateCache(ctx, file.ID)

	return nil
}

// Update به‌روزرسانی اطلاعات یک فایل رسانه‌ای
func (s *MediaFileService) Update(ctx context.Context, file *models.MediaFile) error {
	if file == nil || file.ID == 0 {
		return fmt.Errorf("فایل رسانه نامعتبر است")
	}

	// شروع unit of work
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// به‌روزرسانی فایل
	mediaRepo := uow.MediaFileRepository()
	if err := mediaRepo.Update(ctx, file); err != nil {
		return fmt.Errorf("خطا در به‌روزرسانی فایل رسانه: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true

	// پاکسازی کش
	s.invalidateCache(ctx, file.ID)

	return nil
}

// Delete حذف یک فایل رسانه‌ای و تمام نسخه‌های آن
func (s *MediaFileService) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return fmt.Errorf("شناسه فایل نامعتبر است")
	}

	// شروع unit of work
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// حذف نسخه‌های مرتبط
	versionRepo := uow.MediaVersionRepository()
	versions, err := versionRepo.GetByMediaID(ctx, id)
	if err == nil && len(versions) > 0 {
		for _, v := range versions {
			if err := versionRepo.Delete(ctx, v.ID); err != nil {
				return fmt.Errorf("خطا در حذف نسخه %d: %w", v.ID, err)
			}
		}
	}

	// حذف job های فشرده‌سازی مرتبط
	jobRepo := uow.CompressionJobRepository()
	jobs, err := jobRepo.GetByMediaID(ctx, id)
	if err == nil && len(jobs) > 0 {
		for _, j := range jobs {
			if err := jobRepo.Delete(ctx, j.ID); err != nil {
				return fmt.Errorf("خطا در حذف job %d: %w", j.ID, err)
			}
		}
	}

	// حذف فایل اصلی
	mediaRepo := uow.MediaFileRepository()
	if err := mediaRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("خطا در حذف فایل رسانه: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true

	// پاکسازی کش
	s.invalidateCache(ctx, id)

	return nil
}

// GetByID بازیابی فایل رسانه‌ای بر اساس ID (با کش Redis)
// این متد فقط خواندنی است و نیازی به تراکنش ندارد
func (s *MediaFileService) GetByID(ctx context.Context, id uint) (*models.MediaFile, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه فایل نامعتبر است")
	}

	// بررسی کش
	if s.redis != nil {
		cacheKey := fmt.Sprintf("mediafile:id:%d", id)
		var cached models.MediaFile
		data, err := s.redis.Get(ctx, cacheKey).Bytes()
		if err == nil {
			err = json.Unmarshal(data, &cached)
			if err == nil {
				return &cached, nil
			}
		}
	}

	// دریافت از دیتابیس
	uow := s.uowFactory.Create()
	// برای عملیات خواندنی نیازی به شروع تراکنش نیست
	file, err := uow.MediaFileRepository().GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// ذخیره در کش
	if s.redis != nil {
		cacheKey := fmt.Sprintf("mediafile:id:%d", id)
		if data, err := json.Marshal(file); err == nil {
			_ = s.redis.Set(ctx, cacheKey, data, 12*time.Hour).Err()
		}
	}

	return file, nil
}

// GetByPath بازیابی فایل رسانه‌ای بر اساس مسیر
func (s *MediaFileService) GetByPath(ctx context.Context, path string) (*models.MediaFile, error) {
	if path == "" {
		return nil, fmt.Errorf("مسیر فایل نمی‌تواند خالی باشد")
	}

	uow := s.uowFactory.Create()
	return uow.MediaFileRepository().GetByPath(ctx, path)
}

// GetByUser بازیابی همه فایل‌های یک کاربر
func (s *MediaFileService) GetByUser(ctx context.Context, userID uint) ([]models.MediaFile, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}

	uow := s.uowFactory.Create()
	return uow.MediaFileRepository().GetByUser(ctx, userID)
}

// GetByType بازیابی همه فایل‌های یک نوع خاص
func (s *MediaFileService) GetByType(ctx context.Context, fileType string) ([]models.MediaFile, error) {
	if fileType == "" {
		return nil, fmt.Errorf("نوع فایل نمی‌تواند خالی باشد")
	}

	uow := s.uowFactory.Create()
	return uow.MediaFileRepository().GetByType(ctx, fileType)
}

// ListByStatus لیست فایل‌ها بر اساس وضعیت
func (s *MediaFileService) ListByStatus(ctx context.Context, status string) ([]models.MediaFile, error) {
	uow := s.uowFactory.Create()
	return uow.MediaFileRepository().ListByStatus(ctx, status)
}

// GetWithVariants دریافت فایل رسانه به همراه واریانت‌ها و اطلاعات کامل
func (s *MediaFileService) GetWithVariants(ctx context.Context, id uint) (*models.MediaFile, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه فایل نامعتبر است")
	}

	// استفاده از repository جدید که متد GetWithVariants دارد
	uow := s.uowFactory.Create()

	// چک می‌کنیم آیا repository ما این متد را دارد
	if repo, ok := uow.MediaFileRepository().(*repository.MediaFileRepository); ok {
		// اگر repository قدیمی است، از متد قدیمی استفاده می‌کنیم
		file, err := repo.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}

		// بارگذاری دستی واریانت‌ها
		versionRepo := uow.MediaVersionRepository()
		versions, err := versionRepo.GetByMediaID(ctx, id)
		if err == nil {
			// تبدیل MediaVersion به MediaVariant
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

	// اگر repository جدید است که متد GetWithVariants دارد
	return uow.MediaFileRepository().GetByID(ctx, id)
}

// CompressFile ایجاد یک job فشرده‌سازی برای فایل رسانه
func (s *MediaFileService) CompressFile(ctx context.Context, mediaID uint, targetFormat, targetQuality string) (*models.CompressionJob, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه فایل نامعتبر است")
	}

	// شروع unit of work
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// بررسی وجود فایل
	mediaRepo := uow.MediaFileRepository()
	file, err := mediaRepo.GetByID(ctx, mediaID)
	if err != nil {
		return nil, fmt.Errorf("فایل رسانه یافت نشد: %w", err)
	}

	// ایجاد job فشرده‌سازی
	job := &models.CompressionJob{
		MediaID:       mediaID,
		Status:        "pending",
		JobType:       "compress",
		TargetFormat:  targetFormat,
		TargetQuality: targetQuality,
		OriginalSize:  &file.Size,
		CreatedAt:     time.Now(),
	}

	jobRepo := uow.CompressionJobRepository()
	if err := jobRepo.Create(ctx, job); err != nil {
		return nil, fmt.Errorf("خطا در ایجاد job فشرده‌سازی: %w", err)
	}

	// ایجاد تنظیمات فشرده‌سازی
	settingValue := fmt.Sprintf(`{"format": "%s", "quality": "%s"}`, targetFormat, targetQuality)
	setting := &models.CompressionSetting{
		Scope: "job",
		JobID: &job.ID,
		Key:   "compression_config",
		Value: settingValue,
	}

	settingRepo := uow.CompressionSettingRepository()
	if err := settingRepo.Create(ctx, setting); err != nil {
		return nil, fmt.Errorf("خطا در ذخیره تنظیمات: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return nil, fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true

	return job, nil
}

// GetPaginated دریافت فایل‌های رسانه به صورت صفحه‌بندی شده
func (s *MediaFileService) GetPaginated(ctx context.Context, page, pageSize int, filter repository.MediaFileFilter) ([]models.MediaFile, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	uow := s.uowFactory.Create()

	// چک می‌کنیم آیا repository جدید این متد را دارد
	if newRepo, ok := uow.MediaFileRepository().(*repository.MediaFileRepository); ok && newRepo == nil {
		// اگر repository قدیمی است، از متد ساده استفاده می‌کنیم
		files, err := uow.MediaFileRepository().ListByStatus(ctx, "")
		if err != nil {
			return nil, 0, err
		}

		// صفحه‌بندی دستی
		total := int64(len(files))
		start := (page - 1) * pageSize
		end := start + pageSize

		if start >= len(files) {
			return []models.MediaFile{}, total, nil
		}
		if end > len(files) {
			end = len(files)
		}

		return files[start:end], total, nil
	}

	// repository جدید را type assert می‌کنیم
	repo := uow.MediaFileRepository()

	// به صورت موقت از متد قدیمی استفاده می‌کنیم
	files, err := repo.ListByStatus(ctx, "")
	if err != nil {
		return nil, 0, err
	}

	// فیلتر و صفحه‌بندی دستی
	filtered := make([]models.MediaFile, 0)
	for _, f := range files {
		// اعمال فیلترها
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

		filtered = append(filtered, f)
	}

	// صفحه‌بندی
	total := int64(len(filtered))
	start := (page - 1) * pageSize
	end := start + pageSize

	if start >= len(filtered) {
		return []models.MediaFile{}, total, nil
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], total, nil
}

// GetStatsByUser دریافت آمار فایل‌های رسانه یک کاربر
func (s *MediaFileService) GetStatsByUser(ctx context.Context, userID uint) (*repository.MediaFileStats, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}

	uow := s.uowFactory.Create()

	// محاسبه آمار دستی
	files, err := uow.MediaFileRepository().GetByUser(ctx, userID)
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

// invalidateCache پاکسازی کش برای یک فایل رسانه
func (s *MediaFileService) invalidateCache(ctx context.Context, id uint) {
	if s.redis == nil || id == 0 {
		return
	}

	cacheKey := fmt.Sprintf("mediafile:id:%d", id)
	_ = s.redis.Del(ctx, cacheKey).Err()
}
