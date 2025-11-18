package repository

import (
	"context"
	"fmt"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// MediaFileRepositoryInterface defines the contract for media file data access
// اینترفیس ریپازیتوری فایل‌های رسانه‌ای برای تست‌پذیری و جداسازی لایه‌ها
// این نسخه بازطراحی شده است تا تراکنش‌ها توسط UnitOfWork مدیریت شوند
type MediaFileRepositoryInterface interface {
	Create(ctx context.Context, file *models.MediaFile) error
	Update(ctx context.Context, file *models.MediaFile) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*models.MediaFile, error)
	GetByPath(ctx context.Context, path string) (*models.MediaFile, error)
	GetByUser(ctx context.Context, userID uint) ([]models.MediaFile, error)
	GetByType(ctx context.Context, fileType string) ([]models.MediaFile, error)
	ListByStatus(ctx context.Context, status string) ([]models.MediaFile, error)

	// متدهای اضافی برای عملیات پیچیده‌تر
	BulkCreate(ctx context.Context, files []models.MediaFile) error
	BulkUpdate(ctx context.Context, files []models.MediaFile) error
	GetWithVariants(ctx context.Context, id uint) (*models.MediaFile, error)
	GetPaginated(ctx context.Context, page, pageSize int, filter MediaFileFilter) ([]models.MediaFile, int64, error)
	SearchByTitle(ctx context.Context, query string) ([]models.MediaFile, error)
	UpdateCategory(ctx context.Context, id uint, category string) error
	GetByCategory(ctx context.Context, category string) ([]models.MediaFile, error)
	GetRecentUploads(ctx context.Context, limit int) ([]models.MediaFile, error)
	GetStatsByUser(ctx context.Context, userID uint) (*MediaFileStats, error)
}

// MediaFileFilter ساختار برای فیلتر کردن فایل‌های رسانه
type MediaFileFilter struct {
	FileType   string
	Category   string
	UserID     *uint
	Compressed *bool
	FromDate   *string
	ToDate     *string
	MinSize    *uint
	MaxSize    *uint
}

// MediaFileStats آمار فایل‌های رسانه یک کاربر
type MediaFileStats struct {
	TotalFiles      int64
	TotalSize       uint64
	CompressedFiles int64
	CompressedSize  uint64
	FileTypes       map[string]int64
}

// MediaFileRepository is the GORM implementation of MediaFileRepositoryInterface
// پیاده‌سازی ریپازیتوری فایل‌های رسانه‌ای با GORM
// توجه: این ریپازیتوری دیگر تراکنش‌های داخلی ندارد و از DB تزریق شده استفاده می‌کند
type MediaFileRepository struct {
	DB *gorm.DB // این DB می‌تواند یک connection عادی یا یک transaction از UnitOfWork باشد
}

// NewMediaFileRepository ایجاد نمونه جدید ریپازیتوری
// db: می‌تواند یک اتصال معمولی یا تراکنش از UnitOfWork باشد
func NewMediaFileRepository(db *gorm.DB) MediaFileRepositoryInterface {
	return &MediaFileRepository{DB: db.Session(&gorm.Session{PrepareStmt: true})}
}

// Create درج یک فایل رسانه جدید
// توجه: مدیریت تراکنش بر عهده UnitOfWork است
func (r *MediaFileRepository) Create(ctx context.Context, file *models.MediaFile) error {
	if file == nil {
		return fmt.Errorf("فایل رسانه نمی‌تواند nil باشد")
	}
	return r.DB.WithContext(ctx).Create(file).Error
}

// Update به‌روزرسانی یک فایل رسانه موجود
func (r *MediaFileRepository) Update(ctx context.Context, file *models.MediaFile) error {
	if file == nil || file.ID == 0 {
		return fmt.Errorf("فایل رسانه نامعتبر است")
	}
	return r.DB.WithContext(ctx).Save(file).Error
}

// Delete حذف یک فایل رسانه بر اساس ID (soft delete)
func (r *MediaFileRepository) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return fmt.Errorf("شناسه فایل نامعتبر است")
	}
	result := r.DB.WithContext(ctx).Delete(&models.MediaFile{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetByID دریافت فایل رسانه بر اساس ID
func (r *MediaFileRepository) GetByID(ctx context.Context, id uint) (*models.MediaFile, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه فایل نامعتبر است")
	}
	var file models.MediaFile
	err := r.DB.WithContext(ctx).First(&file, id).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

// GetByPath دریافت فایل رسانه بر اساس مسیر
func (r *MediaFileRepository) GetByPath(ctx context.Context, path string) (*models.MediaFile, error) {
	if path == "" {
		return nil, fmt.Errorf("مسیر فایل نمی‌تواند خالی باشد")
	}
	var file models.MediaFile
	err := r.DB.WithContext(ctx).Where("url = ?", path).First(&file).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

// GetByUser دریافت تمام فایل‌های رسانه یک کاربر
func (r *MediaFileRepository) GetByUser(ctx context.Context, userID uint) ([]models.MediaFile, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	var files []models.MediaFile
	err := r.DB.WithContext(ctx).
		Where("uploaded_by = ?", userID).
		Order("created_at DESC").
		Find(&files).Error
	return files, err
}

// GetByType دریافت تمام فایل‌های رسانه از یک نوع خاص
func (r *MediaFileRepository) GetByType(ctx context.Context, fileType string) ([]models.MediaFile, error) {
	if fileType == "" {
		return nil, fmt.Errorf("نوع فایل نمی‌تواند خالی باشد")
	}
	var files []models.MediaFile
	err := r.DB.WithContext(ctx).
		Where("file_type = ?", fileType).
		Order("created_at DESC").
		Find(&files).Error
	return files, err
}

// ListByStatus دریافت فایل‌های رسانه بر اساس وضعیت (در حال حاضر مدل MediaFile فیلد status ندارد)
// این متد برای سازگاری با اینترفیس حفظ شده است
func (r *MediaFileRepository) ListByStatus(ctx context.Context, status string) ([]models.MediaFile, error) {
	// چون مدل فعلی فیلد status ندارد، همه فایل‌ها را برمی‌گردانیم
	// در آینده می‌توان این فیلد را اضافه کرد
	var files []models.MediaFile
	err := r.DB.WithContext(ctx).Find(&files).Error
	return files, err
}

// BulkCreate درج دسته‌ای فایل‌های رسانه
func (r *MediaFileRepository) BulkCreate(ctx context.Context, files []models.MediaFile) error {
	if len(files) == 0 {
		return nil
	}
	return r.DB.WithContext(ctx).CreateInBatches(files, 100).Error
}

// BulkUpdate به‌روزرسانی دسته‌ای فایل‌های رسانه
func (r *MediaFileRepository) BulkUpdate(ctx context.Context, files []models.MediaFile) error {
	if len(files) == 0 {
		return nil
	}
	// استفاده از تراکنش داخلی فقط برای این عملیات
	// توجه: اگر DB از قبل یک تراکنش باشد، GORM آن را تشخیص می‌دهد
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, file := range files {
			if err := tx.Save(&file).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// GetWithVariants دریافت فایل رسانه به همراه واریانت‌هایش
func (r *MediaFileRepository) GetWithVariants(ctx context.Context, id uint) (*models.MediaFile, error) {
	if id == 0 {
		return nil, fmt.Errorf("شناسه فایل نامعتبر است")
	}
	var file models.MediaFile
	err := r.DB.WithContext(ctx).
		Preload("Variants").
		Preload("CompressionJobs").
		First(&file, id).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

// GetPaginated دریافت فایل‌های رسانه به صورت صفحه‌بندی شده با فیلتر
func (r *MediaFileRepository) GetPaginated(ctx context.Context, page, pageSize int, filter MediaFileFilter) ([]models.MediaFile, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	var files []models.MediaFile
	var total int64

	query := r.DB.WithContext(ctx).Model(&models.MediaFile{})

	// اعمال فیلترها
	if filter.FileType != "" {
		query = query.Where("file_type = ?", filter.FileType)
	}
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	if filter.UserID != nil {
		query = query.Where("uploaded_by = ?", *filter.UserID)
	}
	if filter.Compressed != nil {
		query = query.Where("compressed = ?", *filter.Compressed)
	}
	if filter.MinSize != nil {
		query = query.Where("size >= ?", *filter.MinSize)
	}
	if filter.MaxSize != nil {
		query = query.Where("size <= ?", *filter.MaxSize)
	}
	if filter.FromDate != nil {
		query = query.Where("created_at >= ?", *filter.FromDate)
	}
	if filter.ToDate != nil {
		query = query.Where("created_at <= ?", *filter.ToDate)
	}

	// شمارش کل
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// دریافت رکوردها
	offset := (page - 1) * pageSize
	err := query.
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&files).Error

	return files, total, err
}

// SearchByTitle جستجو در عنوان فایل‌های رسانه
func (r *MediaFileRepository) SearchByTitle(ctx context.Context, query string) ([]models.MediaFile, error) {
	if query == "" {
		return nil, fmt.Errorf("عبارت جستجو نمی‌تواند خالی باشد")
	}

	var files []models.MediaFile
	searchPattern := "%" + query + "%"
	err := r.DB.WithContext(ctx).
		Where("title ILIKE ? OR display_title ILIKE ? OR file_name ILIKE ?",
			searchPattern, searchPattern, searchPattern).
		Order("created_at DESC").
		Find(&files).Error
	return files, err
}

// UpdateCategory به‌روزرسانی دسته‌بندی یک فایل رسانه
func (r *MediaFileRepository) UpdateCategory(ctx context.Context, id uint, category string) error {
	if id == 0 {
		return fmt.Errorf("شناسه فایل نامعتبر است")
	}

	result := r.DB.WithContext(ctx).
		Model(&models.MediaFile{}).
		Where("id = ?", id).
		Update("category", category)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// GetByCategory دریافت فایل‌های رسانه بر اساس دسته‌بندی
func (r *MediaFileRepository) GetByCategory(ctx context.Context, category string) ([]models.MediaFile, error) {
	var files []models.MediaFile
	query := r.DB.WithContext(ctx)

	if category == "" {
		// دریافت فایل‌های بدون دسته‌بندی
		query = query.Where("category = '' OR category IS NULL")
	} else {
		query = query.Where("category = ?", category)
	}

	err := query.Order("created_at DESC").Find(&files).Error
	return files, err
}

// GetRecentUploads دریافت آخرین فایل‌های آپلود شده
func (r *MediaFileRepository) GetRecentUploads(ctx context.Context, limit int) ([]models.MediaFile, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	var files []models.MediaFile
	err := r.DB.WithContext(ctx).
		Order("created_at DESC").
		Limit(limit).
		Find(&files).Error
	return files, err
}

// GetStatsByUser دریافت آمار فایل‌های رسانه یک کاربر
func (r *MediaFileRepository) GetStatsByUser(ctx context.Context, userID uint) (*MediaFileStats, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}

	stats := &MediaFileStats{
		FileTypes: make(map[string]int64),
	}

	// شمارش کل فایل‌ها
	r.DB.WithContext(ctx).
		Model(&models.MediaFile{}).
		Where("uploaded_by = ?", userID).
		Count(&stats.TotalFiles)

	// محاسبه حجم کل
	var totalSize struct {
		Total uint64
	}
	r.DB.WithContext(ctx).
		Model(&models.MediaFile{}).
		Where("uploaded_by = ?", userID).
		Select("COALESCE(SUM(size), 0) as total").
		Scan(&totalSize)
	stats.TotalSize = totalSize.Total

	// شمارش فایل‌های فشرده شده
	r.DB.WithContext(ctx).
		Model(&models.MediaFile{}).
		Where("uploaded_by = ? AND compressed = ?", userID, true).
		Count(&stats.CompressedFiles)

	// محاسبه حجم فایل‌های فشرده شده
	var compressedSize struct {
		Total uint64
	}
	r.DB.WithContext(ctx).
		Model(&models.MediaFile{}).
		Where("uploaded_by = ? AND compressed = ? AND compressed_size IS NOT NULL", userID, true).
		Select("COALESCE(SUM(compressed_size), 0) as total").
		Scan(&compressedSize)
	stats.CompressedSize = compressedSize.Total

	// آمار بر اساس نوع فایل
	rows, err := r.DB.WithContext(ctx).
		Model(&models.MediaFile{}).
		Where("uploaded_by = ?", userID).
		Select("file_type, COUNT(*) as count").
		Group("file_type").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fileType string
		var count int64
		if err := rows.Scan(&fileType, &count); err != nil {
			return nil, err
		}
		stats.FileTypes[fileType] = count
	}

	return stats, nil
}
