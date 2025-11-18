package services

import (
	"context"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
)

// MediaVersionService سرویس مدیریت نسخه‌های رسانه
// این سرویس عملیات CRUD و جستجو را برای نسخه‌های مختلف هر رسانه فراهم می‌کند

type MediaVersionService struct {
	repo repository.MediaVersionRepositoryInterface
}

// NewMediaVersionService سازنده سرویس نسخه رسانه
func NewMediaVersionService(repo repository.MediaVersionRepositoryInterface) *MediaVersionService {
	return &MediaVersionService{repo: repo}
}

// Create یک نسخه جدید برای رسانه ایجاد می‌کند
func (s *MediaVersionService) Create(ctx context.Context, version *models.MediaVersion) error {
	return s.repo.Create(ctx, version)
}

// Update اطلاعات یک نسخه را بروزرسانی می‌کند
func (s *MediaVersionService) Update(ctx context.Context, version *models.MediaVersion) error {
	return s.repo.Update(ctx, version)
}

// Delete یک نسخه را حذف می‌کند
func (s *MediaVersionService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

// GetByID بازیابی نسخه بر اساس ID
func (s *MediaVersionService) GetByID(ctx context.Context, id uint) (*models.MediaVersion, error) {
	return s.repo.GetByID(ctx, id)
}

// GetByMedia بازیابی همه نسخه‌های یک رسانه
func (s *MediaVersionService) GetByMedia(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	return s.repo.GetByMediaID(ctx, mediaID)
}

// GetActiveVersions بازیابی نسخه‌های فعال یک رسانه
func (s *MediaVersionService) GetActiveVersions(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	return s.repo.GetActiveVersions(ctx, mediaID)
}

// GetBackups بازیابی نسخه‌های بکاپ یک رسانه
func (s *MediaVersionService) GetBackups(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	return s.repo.GetBackups(ctx, mediaID)
}

// ListByFormat لیست نسخه‌ها بر اساس فرمت
func (s *MediaVersionService) ListByFormat(ctx context.Context, mediaID uint, format string) ([]models.MediaVersion, error) {
	return s.repo.ListByFormat(ctx, mediaID, format)
}
