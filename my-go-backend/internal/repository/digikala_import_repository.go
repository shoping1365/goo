package repository

import (
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// DigikalaImportRepositoryInterface رابط repository برای مدیریت import های دیجی‌کالا
type DigikalaImportRepositoryInterface interface {
	Create(importRecord *models.DigikalaImport) error
	GetByID(id uint) (*models.DigikalaImport, error)
	Update(importRecord *models.DigikalaImport) error
	Delete(id uint) error
	List(limit, offset int) ([]models.DigikalaImport, error)
	GetStats() (*models.DigikalaImportStats, error)
	GetByStatus(status string) ([]models.DigikalaImport, error)
	GetInProgress() ([]models.DigikalaImport, error)
}

// DigikalaImportRepository پیاده‌سازی repository برای import های دیجی‌کالا
type DigikalaImportRepository struct {
	db *gorm.DB
}

// NewDigikalaImportRepository ایجاد instance جدید از repository
func NewDigikalaImportRepository(db *gorm.DB) DigikalaImportRepositoryInterface {
	return &DigikalaImportRepository{db: db}
}

// Create ایجاد یک import جدید
func (r *DigikalaImportRepository) Create(importRecord *models.DigikalaImport) error {
	return r.db.Create(importRecord).Error
}

// GetByID دریافت import بر اساس ID
func (r *DigikalaImportRepository) GetByID(id uint) (*models.DigikalaImport, error) {
	var importRecord models.DigikalaImport
	err := r.db.First(&importRecord, id).Error
	if err != nil {
		return nil, err
	}
	return &importRecord, nil
}

// Update به‌روزرسانی import
func (r *DigikalaImportRepository) Update(importRecord *models.DigikalaImport) error {
	return r.db.Save(importRecord).Error
}

// Delete حذف import
func (r *DigikalaImportRepository) Delete(id uint) error {
	return r.db.Delete(&models.DigikalaImport{}, id).Error
}

// List دریافت لیست import ها
func (r *DigikalaImportRepository) List(limit, offset int) ([]models.DigikalaImport, error) {
	var imports []models.DigikalaImport
	err := r.db.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&imports).Error
	return imports, err
}

// GetStats دریافت آمار کلی import ها
func (r *DigikalaImportRepository) GetStats() (*models.DigikalaImportStats, error) {
	var stats models.DigikalaImportStats

	// شمارش کل import ها
	var totalImports int64
	if err := r.db.Model(&models.DigikalaImport{}).Count(&totalImports).Error; err != nil {
		return nil, err
	}
	stats.TotalImports = int(totalImports)

	// شمارش import های تکمیل شده
	var completedImports int64
	if err := r.db.Model(&models.DigikalaImport{}).
		Where("status = ?", models.ImportStatuses.Completed).
		Count(&completedImports).Error; err != nil {
		return nil, err
	}
	stats.CompletedImports = int(completedImports)

	// شمارش import های ناموفق
	var failedImports int64
	if err := r.db.Model(&models.DigikalaImport{}).
		Where("status = ?", models.ImportStatuses.Failed).
		Count(&failedImports).Error; err != nil {
		return nil, err
	}
	stats.FailedImports = int(failedImports)

	// شمارش import های در حال انجام
	var inProgressImports int64
	if err := r.db.Model(&models.DigikalaImport{}).
		Where("status = ?", models.ImportStatuses.InProgress).
		Count(&inProgressImports).Error; err != nil {
		return nil, err
	}
	stats.InProgressImports = int(inProgressImports)

	// مجموع محصولات
	if err := r.db.Model(&models.DigikalaImport{}).
		Select("COALESCE(SUM(total_products), 0)").
		Scan(&stats.TotalProducts).Error; err != nil {
		return nil, err
	}

	// مجموع محصولات import شده
	if err := r.db.Model(&models.DigikalaImport{}).
		Select("COALESCE(SUM(imported_products), 0)").
		Scan(&stats.ImportedProducts).Error; err != nil {
		return nil, err
	}

	// مجموع محصولات ناموفق
	if err := r.db.Model(&models.DigikalaImport{}).
		Select("COALESCE(SUM(failed_products), 0)").
		Scan(&stats.FailedProducts).Error; err != nil {
		return nil, err
	}

	// محاسبه نرخ موفقیت
	if stats.TotalProducts > 0 {
		stats.SuccessRate = float64(stats.ImportedProducts) / float64(stats.TotalProducts) * 100
	}

	// محاسبه سرعت متوسط (محصولات در دقیقه)
	var totalDuration float64
	var completedImportRecords []models.DigikalaImport
	if err := r.db.Where("status = ? AND started_at IS NOT NULL AND completed_at IS NOT NULL",
		models.ImportStatuses.Completed).Find(&completedImportRecords).Error; err == nil {
		for _, imp := range completedImportRecords {
			if imp.StartedAt != nil && imp.CompletedAt != nil {
				duration := imp.CompletedAt.Sub(*imp.StartedAt).Minutes()
				if duration > 0 {
					totalDuration += duration
				}
			}
		}
		if totalDuration > 0 {
			stats.AverageSpeed = float64(stats.ImportedProducts) / totalDuration
		}
	}

	return &stats, nil
}

// GetByStatus دریافت import ها بر اساس وضعیت
func (r *DigikalaImportRepository) GetByStatus(status string) ([]models.DigikalaImport, error) {
	var imports []models.DigikalaImport
	err := r.db.Where("status = ?", status).
		Order("created_at DESC").
		Find(&imports).Error
	return imports, err
}

// GetInProgress دریافت import های در حال انجام
func (r *DigikalaImportRepository) GetInProgress() ([]models.DigikalaImport, error) {
	return r.GetByStatus(models.ImportStatuses.InProgress)
}

// DigikalaImportLogRepositoryInterface رابط repository برای لاگ‌های import
type DigikalaImportLogRepositoryInterface interface {
	Create(log *models.DigikalaImportLog) error
	GetByImportID(importID uint) ([]models.DigikalaImportLog, error)
	GetByLevel(level string) ([]models.DigikalaImportLog, error)
}

// DigikalaImportLogRepository پیاده‌سازی repository برای لاگ‌های import
type DigikalaImportLogRepository struct {
	db *gorm.DB
}

// NewDigikalaImportLogRepository ایجاد instance جدید از repository
func NewDigikalaImportLogRepository(db *gorm.DB) DigikalaImportLogRepositoryInterface {
	return &DigikalaImportLogRepository{db: db}
}

// Create ایجاد یک لاگ جدید
func (r *DigikalaImportLogRepository) Create(log *models.DigikalaImportLog) error {
	return r.db.Create(log).Error
}

// GetByImportID دریافت لاگ‌ها بر اساس ID import
func (r *DigikalaImportLogRepository) GetByImportID(importID uint) ([]models.DigikalaImportLog, error) {
	var logs []models.DigikalaImportLog
	err := r.db.Where("import_id = ?", importID).
		Order("created_at DESC").
		Find(&logs).Error
	return logs, err
}

// GetByLevel دریافت لاگ‌ها بر اساس سطح
func (r *DigikalaImportLogRepository) GetByLevel(level string) ([]models.DigikalaImportLog, error) {
	var logs []models.DigikalaImportLog
	err := r.db.Where("level = ?", level).
		Order("created_at DESC").
		Find(&logs).Error
	return logs, err
}
