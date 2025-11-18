package services

import (
	"context"
	"fmt"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
)

// CompressionJobService provides business logic for CompressionJob entities.
// It is a thin wrapper over the repository layer for now but gives a place
// to put extra validation / orchestration in the future.
type CompressionJobService struct {
	Repo repository.CompressionJobRepositoryInterface
}

func NewCompressionJobService(repo repository.CompressionJobRepositoryInterface) *CompressionJobService {
	return &CompressionJobService{Repo: repo}
}

func (s *CompressionJobService) Create(ctx context.Context, job *models.CompressionJob) error {
	return s.Repo.Create(ctx, job)
}

func (s *CompressionJobService) Update(ctx context.Context, job *models.CompressionJob) error {
	return s.Repo.Update(ctx, job)
}

func (s *CompressionJobService) Delete(ctx context.Context, id uint) error {
	return s.Repo.Delete(ctx, id)
}

func (s *CompressionJobService) GetByID(ctx context.Context, id uint) (*models.CompressionJob, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *CompressionJobService) GetByMedia(ctx context.Context, mediaID uint) ([]models.CompressionJob, error) {
	return s.Repo.GetByMediaID(ctx, mediaID)
}

func (s *CompressionJobService) GetByStatus(ctx context.Context, status string) ([]models.CompressionJob, error) {
	return s.Repo.GetByStatus(ctx, status)
}

func (s *CompressionJobService) ListByBatch(ctx context.Context, batchID uint) ([]models.CompressionJob, error) {
	return s.Repo.ListByBatch(ctx, batchID)
}

// CompressionStatService provides business logic for CompressionStat entities.

type CompressionStatService struct {
	Repo repository.CompressionStatRepositoryInterface
}

func NewCompressionStatService(repo repository.CompressionStatRepositoryInterface) *CompressionStatService {
	return &CompressionStatService{Repo: repo}
}

func (s *CompressionStatService) Create(ctx context.Context, stat *models.CompressionStat) error {
	return s.Repo.Create(ctx, stat)
}

func (s *CompressionStatService) Update(ctx context.Context, stat *models.CompressionStat) error {
	return s.Repo.Update(ctx, stat)
}

func (s *CompressionStatService) Delete(ctx context.Context, id uint) error {
	return s.Repo.Delete(ctx, id)
}

func (s *CompressionStatService) GetByID(ctx context.Context, id uint) (*models.CompressionStat, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *CompressionStatService) GetByType(ctx context.Context, statType string) ([]models.CompressionStat, error) {
	return s.Repo.GetByType(ctx, statType)
}

func (s *CompressionStatService) GetByMedia(ctx context.Context, mediaID uint) ([]models.CompressionStat, error) {
	return s.Repo.GetByMedia(ctx, mediaID)
}

func (s *CompressionStatService) GetByUser(ctx context.Context, userID uint) ([]models.CompressionStat, error) {
	return s.Repo.GetByUser(ctx, userID)
}

func (s *CompressionStatService) GetByFormat(ctx context.Context, format string) ([]models.CompressionStat, error) {
	return s.Repo.GetByFormat(ctx, format)
}

func (s *CompressionStatService) GetByPeriod(ctx context.Context, period string) ([]models.CompressionStat, error) {
	return s.Repo.GetByPeriod(ctx, period)
}

// CompressionSettingService provides business logic for CompressionSetting entities.

type CompressionSettingService struct {
	Repo repository.CompressionSettingRepositoryInterface
}

func NewCompressionSettingService(repo repository.CompressionSettingRepositoryInterface) *CompressionSettingService {
	return &CompressionSettingService{Repo: repo}
}

func (s *CompressionSettingService) Create(ctx context.Context, setting *models.CompressionSetting) error {
	return s.Repo.Create(ctx, setting)
}

func (s *CompressionSettingService) Update(ctx context.Context, setting *models.CompressionSetting) error {
	return s.Repo.Update(ctx, setting)
}

func (s *CompressionSettingService) Delete(ctx context.Context, id uint) error {
	return s.Repo.Delete(ctx, id)
}

func (s *CompressionSettingService) GetByID(ctx context.Context, id uint) (*models.CompressionSetting, error) {
	return s.Repo.GetByID(ctx, id)
}

func (s *CompressionSettingService) GetByScope(ctx context.Context, scope string) ([]models.CompressionSetting, error) {
	return s.Repo.GetByScope(ctx, scope)
}

func (s *CompressionSettingService) GetByUser(ctx context.Context, userID uint) ([]models.CompressionSetting, error) {
	return s.Repo.GetByUser(ctx, userID)
}

func (s *CompressionSettingService) GetByMedia(ctx context.Context, mediaID uint) ([]models.CompressionSetting, error) {
	return s.Repo.GetByMedia(ctx, mediaID)
}

func (s *CompressionSettingService) GetByJob(ctx context.Context, jobID uint) ([]models.CompressionSetting, error) {
	return s.Repo.GetByJob(ctx, jobID)
}

// GetVideoCompressionSettings تنظیمات مرکزی فشرده‌سازی ویدیو را برمی‌گرداند
// این متد تمام تنظیمات مربوط به فشرده‌سازی ویدیو را از دیتابیس می‌خواند
func (s *CompressionSettingService) GetVideoCompressionSettings(ctx context.Context) (map[string]string, error) {
	// خواندن تنظیمات از جدول admin_settings با کلیدهای video_compression.*
	settings, err := s.Repo.GetByScope(ctx, "global")
	if err != nil {
		return nil, err
	}

	// تبدیل به map برای استفاده آسان‌تر
	result := make(map[string]string)
	for _, setting := range settings {
		if setting.Key != "" {
			result[setting.Key] = setting.Value
		}
	}

	return result, nil
}

// GetVideoCompressionWorkerCount تعداد ورکر برای فشرده‌سازی ویدیو را برمی‌گرداند
// اگر تنظیمات موجود نباشد، مقدار پیش‌فرض 4 برمی‌گرداند
func (s *CompressionSettingService) GetVideoCompressionWorkerCount(ctx context.Context) (int, error) {
	settings, err := s.GetVideoCompressionSettings(ctx)
	if err != nil {
		return 4, err // مقدار پیش‌فرض در صورت خطا
	}

	// بررسی وجود کلید worker_count
	if workerCountStr, exists := settings["video_compression.worker_count"]; exists {
		// تبدیل رشته به عدد
		var workerCount int
		_, err := fmt.Sscanf(workerCountStr, "%d", &workerCount)
		if err == nil && workerCount >= 1 && workerCount <= 16 {
			return workerCount, nil
		}
	}

	return 4, nil // مقدار پیش‌فرض
}

// GetVideoCompressionQuality کیفیت فشرده‌سازی ویدیو را برمی‌گرداند
func (s *CompressionSettingService) GetVideoCompressionQuality(ctx context.Context) (string, error) {
	settings, err := s.GetVideoCompressionSettings(ctx)
	if err != nil {
		return "720p", err // مقدار پیش‌فرض در صورت خطا
	}

	if quality, exists := settings["video_compression.quality"]; exists {
		return quality, nil
	}

	return "720p", nil // مقدار پیش‌فرض
}

// GetVideoCompressionFormat فرمت فشرده‌سازی ویدیو را برمی‌گرداند
func (s *CompressionSettingService) GetVideoCompressionFormat(ctx context.Context) (string, error) {
	settings, err := s.GetVideoCompressionSettings(ctx)
	if err != nil {
		return "mp4", err // مقدار پیش‌فرض در صورت خطا
	}

	if format, exists := settings["video_compression.format"]; exists {
		return format, nil
	}

	return "mp4", nil // مقدار پیش‌فرض
}

// IsVideoCompressionEnabled بررسی می‌کند که آیا فشرده‌سازی خودکار ویدیو فعال است یا خیر
func (s *CompressionSettingService) IsVideoCompressionEnabled(ctx context.Context) (bool, error) {
	settings, err := s.GetVideoCompressionSettings(ctx)
	if err != nil {
		return true, err // مقدار پیش‌فرض در صورت خطا
	}

	if enabled, exists := settings["video_compression.enabled"]; exists {
		return enabled == "true" || enabled == "1", nil
	}

	return true, nil // مقدار پیش‌فرض
}

// در متد GetVideoCompressionSettings و سایر متدهای مرتبط، نیازی به تغییر خاص نیست چون همه کلیدها به صورت داینامیک map می‌شوند.
// اما برای راحتی، متدهای جدید زیر را اضافه کن:

// GetVideoCompressionStartHour ساعت شروع فشرده‌سازی ویدیو را برمی‌گرداند
func (s *CompressionSettingService) GetVideoCompressionStartHour(ctx context.Context) (int, error) {
	settings, err := s.GetVideoCompressionSettings(ctx)
	if err != nil {
		return 1, err // مقدار پیش‌فرض
	}
	if start, exists := settings["video_compression.start_hour"]; exists {
		var h int
		_, err := fmt.Sscanf(start, "%d", &h)
		if err == nil && h >= 0 && h <= 23 {
			return h, nil
		}
	}
	return 1, nil
}

// GetVideoCompressionEndHour ساعت پایان فشرده‌سازی ویدیو را برمی‌گرداند
func (s *CompressionSettingService) GetVideoCompressionEndHour(ctx context.Context) (int, error) {
	settings, err := s.GetVideoCompressionSettings(ctx)
	if err != nil {
		return 13, err // مقدار پیش‌فرض
	}
	if end, exists := settings["video_compression.end_hour"]; exists {
		var h int
		_, err := fmt.Sscanf(end, "%d", &h)
		if err == nil && h >= 0 && h <= 23 {
			return h, nil
		}
	}
	return 13, nil
}
