package repository

import (
	"context"
	"fmt"
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

// MediaVersionRepositoryInterface defines the contract for media version data access
// اینترفیس ریپازیتوری نسخه‌های رسانه برای تست‌پذیری و جداسازی لایه‌ها
// این نسخه بازطراحی شده است تا تراکنش‌ها توسط UnitOfWork مدیریت شوند
type MediaVersionRepositoryInterface interface {
	Create(ctx context.Context, mv *models.MediaVersion) error
	Update(ctx context.Context, mv *models.MediaVersion) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*models.MediaVersion, error)
	GetByMediaID(ctx context.Context, mediaID uint) ([]models.MediaVersion, error)
	GetActiveVersion(ctx context.Context, mediaID uint) (*models.MediaVersion, error)
	ListBackups(ctx context.Context, mediaID uint) ([]models.MediaVersion, error)
	GetActiveVersions(ctx context.Context, mediaID uint) ([]models.MediaVersion, error)
	GetBackups(ctx context.Context, mediaID uint) ([]models.MediaVersion, error)
	ListByFormat(ctx context.Context, mediaID uint, format string) ([]models.MediaVersion, error)

	// متدهای اضافی برای عملیات پیچیده‌تر
	BulkCreate(ctx context.Context, versions []models.MediaVersion) error
	SetActiveVersion(ctx context.Context, mediaID, versionID uint) error
	CreateBackupFromVersion(ctx context.Context, versionID uint) (*models.MediaVersion, error)
	GetVersionHistory(ctx context.Context, mediaID uint) ([]models.MediaVersion, error)
	DeleteOldBackups(ctx context.Context, mediaID uint, keepCount int) error
	GetVersionsByQuality(ctx context.Context, mediaID uint, quality string) ([]models.MediaVersion, error)
	UpdateCompressionRatio(ctx context.Context, versionID uint, ratio float32) error
	GetTotalSizeByMedia(ctx context.Context, mediaID uint) (uint64, error)
	CleanupOrphanedVersions(ctx context.Context) (int64, error)
}

// MediaVersionRepository is the GORM implementation of MediaVersionRepositoryInterface
// پیاده‌سازی ریپازیتوری نسخه‌های رسانه با GORM
// توجه: این ریپازیتوری دیگر تراکنش‌های داخلی ندارد و از DB تزریق شده استفاده می‌کند
type MediaVersionRepository struct {
	DB *gorm.DB // این DB می‌تواند یک connection عادی یا یک transaction از UnitOfWork باشد
}

// NewMediaVersionRepository ایجاد نمونه جدید ریپازیتوری
// db: می‌تواند یک اتصال معمولی یا تراکنش از UnitOfWork باشد
func NewMediaVersionRepository(db *gorm.DB) MediaVersionRepositoryInterface {
	return &MediaVersionRepository{DB: db}
}

// Create درج یک نسخه رسانه جدید
// توجه: مدیریت تراکنش بر عهده UnitOfWork است
func (r *MediaVersionRepository) Create(ctx context.Context, mv *models.MediaVersion) error {
	if mv == nil {
		return fmt.Errorf("نسخه رسانه نمی‌تواند nil باشد")
	}
	if mv.MediaID == 0 {
		return fmt.Errorf("شناسه رسانه اصلی الزامی است")
	}
	return r.DB.WithContext(ctx).Create(mv).Error
}

// Update به‌روزرسانی یک نسخه رسانه موجود
func (r *MediaVersionRepository) Update(ctx context.Context, mv *models.MediaVersion) error {
	if mv == nil || mv.ID == 0 {
		return fmt.Errorf("نسخه رسانه نامعتبر است")
	}
	return r.DB.WithContext(ctx).Save(mv).Error
}

// Delete حذف یک نسخه رسانه بر اساس ID (soft delete)
func (r *MediaVersionRepository) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return fmt.Errorf("شناسه نسخه نامعتبر است")
	}
	result := r.DB.WithContext(ctx).Delete(&models.MediaVersion{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetByID دریافت نسخه رسانه بر اساس ID
func (r *MediaVersionRepository) GetByID(ctx context.Context, id uint) (*models.MediaVersion, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه نسخه نامعتبر است")
	}
	var mv models.MediaVersion
	err := r.DB.WithContext(ctx).First(&mv, id).Error
	if err != nil {
		return nil, err
	}
	return &mv, nil
}

// GetByMediaID دریافت تمام نسخه‌های یک فایل رسانه
func (r *MediaVersionRepository) GetByMediaID(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	var versions []models.MediaVersion
	err := r.DB.WithContext(ctx).
		Where("media_id = ?", mediaID).
		Order("created_at DESC").
		Find(&versions).Error
	return versions, err
}

// GetActiveVersion دریافت نسخه فعال یک فایل رسانه
func (r *MediaVersionRepository) GetActiveVersion(ctx context.Context, mediaID uint) (*models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	var mv models.MediaVersion
	err := r.DB.WithContext(ctx).
		Where("media_id = ? AND is_active = ?", mediaID, true).
		Order("created_at DESC").
		First(&mv).Error
	if err != nil {
		return nil, err
	}
	return &mv, nil
}

// ListBackups دریافت تمام نسخه‌های بکاپ یک فایل رسانه (Deprecated - use GetBackups)
func (r *MediaVersionRepository) ListBackups(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	return r.GetBackups(ctx, mediaID)
}

// GetActiveVersions دریافت تمام نسخه‌های فعال یک فایل رسانه
func (r *MediaVersionRepository) GetActiveVersions(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	var versions []models.MediaVersion
	err := r.DB.WithContext(ctx).
		Where("media_id = ? AND is_active = ?", mediaID, true).
		Order("created_at DESC").
		Find(&versions).Error
	return versions, err
}

// GetBackups دریافت تمام نسخه‌های بکاپ یک فایل رسانه
func (r *MediaVersionRepository) GetBackups(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	var versions []models.MediaVersion
	err := r.DB.WithContext(ctx).
		Where("media_id = ? AND is_backup = ?", mediaID, true).
		Order("created_at DESC").
		Find(&versions).Error
	return versions, err
}

// ListByFormat دریافت تمام نسخه‌های یک فایل رسانه با فرمت خاص
func (r *MediaVersionRepository) ListByFormat(ctx context.Context, mediaID uint, format string) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	if format == "" {
		return nil, fmt.Errorf("فرمت نمی‌تواند خالی باشد")
	}
	var versions []models.MediaVersion
	err := r.DB.WithContext(ctx).
		Where("media_id = ? AND format = ?", mediaID, format).
		Order("created_at DESC").
		Find(&versions).Error
	return versions, err
}

// BulkCreate درج دسته‌ای نسخه‌های رسانه
func (r *MediaVersionRepository) BulkCreate(ctx context.Context, versions []models.MediaVersion) error {
	if len(versions) == 0 {
		return nil
	}
	// بررسی صحت داده‌ها
	for i, v := range versions {
		if v.MediaID == 0 {
			return fmt.Errorf("نسخه شماره %d: شناسه رسانه اصلی الزامی است", i+1)
		}
	}
	return r.DB.WithContext(ctx).CreateInBatches(versions, 50).Error
}

// SetActiveVersion تنظیم یک نسخه به عنوان نسخه فعال و غیرفعال کردن سایر نسخه‌ها
func (r *MediaVersionRepository) SetActiveVersion(ctx context.Context, mediaID, versionID uint) error {
	if mediaID == 0 || versionID == 0 {
		return fmt.Errorf("شناسه‌ها نامعتبر هستند")
	}

	// بررسی وجود نسخه
	var version models.MediaVersion
	if err := r.DB.WithContext(ctx).First(&version, versionID).Error; err != nil {
		return fmt.Errorf("نسخه مورد نظر یافت نشد: %w", err)
	}

	if version.MediaID != mediaID {
		return fmt.Errorf("نسخه متعلق به رسانه دیگری است")
	}

	// غیرفعال کردن سایر نسخه‌های فعال
	if err := r.DB.WithContext(ctx).
		Model(&models.MediaVersion{}).
		Where("media_id = ? AND is_active = ? AND id != ?", mediaID, true, versionID).
		Update("is_active", false).Error; err != nil {
		return err
	}

	// فعال کردن نسخه انتخابی
	return r.DB.WithContext(ctx).
		Model(&models.MediaVersion{}).
		Where("id = ?", versionID).
		Update("is_active", true).Error
}

// CreateBackupFromVersion ایجاد یک بکاپ از نسخه موجود
func (r *MediaVersionRepository) CreateBackupFromVersion(ctx context.Context, versionID uint) (*models.MediaVersion, error) {
	if versionID == 0 {
		return nil, fmt.Errorf("شناسه نسخه نامعتبر است")
	}

	// دریافت نسخه اصلی
	original, err := r.GetByID(ctx, versionID)
	if err != nil {
		return nil, fmt.Errorf("نسخه اصلی یافت نشد: %w", err)
	}

	// ایجاد کپی به عنوان بکاپ
	backup := &models.MediaVersion{
		MediaID:           original.MediaID,
		VersionCode:       fmt.Sprintf("backup_%s_%d", original.VersionCode, time.Now().Unix()),
		FilePath:          original.FilePath,
		FileSize:          original.FileSize,
		Format:            original.Format,
		Quality:           original.Quality,
		IsActive:          false,
		IsBackup:          true,
		CompressionRatio:  original.CompressionRatio,
		CompressionMethod: original.CompressionMethod,
		ParentVersionID:   &original.ID,
		BackupPath:        original.FilePath, // مسیر بکاپ همان مسیر اصلی است
		Note:              fmt.Sprintf("بکاپ از نسخه %s - %s", original.VersionCode, time.Now().Format("2006-01-02 15:04:05")),
		Meta:              original.Meta,
		CreatedBy:         original.CreatedBy,
	}

	if err := r.Create(ctx, backup); err != nil {
		return nil, fmt.Errorf("خطا در ایجاد بکاپ: %w", err)
	}

	return backup, nil
}

// GetVersionHistory دریافت تاریخچه کامل نسخه‌های یک رسانه
func (r *MediaVersionRepository) GetVersionHistory(ctx context.Context, mediaID uint) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}

	var versions []models.MediaVersion
	err := r.DB.WithContext(ctx).
		Where("media_id = ?", mediaID).
		Order("created_at DESC").
		Find(&versions).Error
	return versions, err
}

// DeleteOldBackups حذف بکاپ‌های قدیمی و نگه داشتن تعداد مشخصی از جدیدترین‌ها
func (r *MediaVersionRepository) DeleteOldBackups(ctx context.Context, mediaID uint, keepCount int) error {
	if mediaID == 0 {
		return fmt.Errorf("شناسه رسانه نامعتبر است")
	}
	if keepCount < 0 {
		keepCount = 0
	}

	// دریافت ID بکاپ‌هایی که باید حذف شوند
	var deleteIDs []uint
	subQuery := r.DB.WithContext(ctx).
		Model(&models.MediaVersion{}).
		Select("id").
		Where("media_id = ? AND is_backup = ?", mediaID, true).
		Order("created_at DESC").
		Offset(keepCount)

	if err := subQuery.Pluck("id", &deleteIDs).Error; err != nil {
		return err
	}

	if len(deleteIDs) == 0 {
		return nil
	}

	// حذف بکاپ‌های قدیمی
	return r.DB.WithContext(ctx).
		Where("id IN ?", deleteIDs).
		Delete(&models.MediaVersion{}).Error
}

// GetVersionsByQuality دریافت نسخه‌های یک رسانه بر اساس کیفیت
func (r *MediaVersionRepository) GetVersionsByQuality(ctx context.Context, mediaID uint, quality string) ([]models.MediaVersion, error) {
	if mediaID == 0 {
		return nil, fmt.Errorf("شناسه رسانه نامعتبر است")
	}

	var versions []models.MediaVersion
	query := r.DB.WithContext(ctx).Where("media_id = ?", mediaID)

	if quality != "" {
		query = query.Where("quality = ?", quality)
	}

	err := query.Order("created_at DESC").Find(&versions).Error
	return versions, err
}

// UpdateCompressionRatio به‌روزرسانی نسبت فشرده‌سازی یک نسخه
func (r *MediaVersionRepository) UpdateCompressionRatio(ctx context.Context, versionID uint, ratio float32) error {
	if versionID == 0 {
		return fmt.Errorf("شناسه نسخه نامعتبر است")
	}
	if ratio < 0 || ratio > 100 {
		return fmt.Errorf("نسبت فشرده‌سازی باید بین 0 و 100 باشد")
	}

	result := r.DB.WithContext(ctx).
		Model(&models.MediaVersion{}).
		Where("id = ?", versionID).
		Update("compression_ratio", ratio)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetTotalSizeByMedia محاسبه مجموع حجم تمام نسخه‌های یک رسانه
func (r *MediaVersionRepository) GetTotalSizeByMedia(ctx context.Context, mediaID uint) (uint64, error) {
	if mediaID == 0 {
		return 0, fmt.Errorf("شناسه رسانه نامعتبر است")
	}

	var totalSize struct {
		Total uint64
	}

	err := r.DB.WithContext(ctx).
		Model(&models.MediaVersion{}).
		Where("media_id = ?", mediaID).
		Select("COALESCE(SUM(file_size), 0) as total").
		Scan(&totalSize).Error

	return totalSize.Total, err
}

// CleanupOrphanedVersions حذف نسخه‌هایی که رسانه اصلی آنها حذف شده است
func (r *MediaVersionRepository) CleanupOrphanedVersions(ctx context.Context) (int64, error) {
	// این query نسخه‌هایی را حذف می‌کند که media_id آنها در جدول media_files وجود ندارد
	result := r.DB.WithContext(ctx).
		Where("media_id NOT IN (SELECT id FROM media_files)").
		Delete(&models.MediaVersion{})

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
