package models

import (
	"time"

	"gorm.io/gorm"
)

// DigikalaImport نمایانگر یک عملیات import از دیجی‌کالا
type DigikalaImport struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	CategoryURL      string         `json:"category_url" gorm:"not null;size:500"`
	CategoryTitle    string         `json:"category_title" gorm:"size:255"`
	Status           string         `json:"status" gorm:"not null;size:50;default:'pending'"` // pending, in_progress, completed, failed, cancelled
	TotalProducts    int            `json:"total_products" gorm:"default:0"`
	ImportedProducts int            `json:"imported_products" gorm:"default:0"`
	FailedProducts   int            `json:"failed_products" gorm:"default:0"`
	Progress         float64        `json:"progress" gorm:"default:0"` // درصد پیشرفت (0-100)
	ErrorMessage     *string        `json:"error_message" gorm:"type:text"`
	StartedAt        *time.Time     `json:"started_at"`
	CompletedAt      *time.Time     `json:"completed_at"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName نام جدول در دیتابیس
func (DigikalaImport) TableName() string {
	return "digikala_imports"
}

// DigikalaImportLog نمایانگر لاگ‌های عملیات import
type DigikalaImportLog struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	ImportID   uint           `json:"import_id" gorm:"not null"`
	Import     DigikalaImport `json:"import" gorm:"foreignKey:ImportID"`
	Level      string         `json:"level" gorm:"not null;size:20"` // info, warning, error
	Message    string         `json:"message" gorm:"not null;type:text"`
	ProductURL *string        `json:"product_url" gorm:"size:500"`
	CreatedAt  time.Time      `json:"created_at"`
}

// TableName نام جدول در دیتابیس
func (DigikalaImportLog) TableName() string {
	return "digikala_import_logs"
}

// DigikalaImportStats آمار کلی import ها
type DigikalaImportStats struct {
	TotalImports      int     `json:"total_imports"`
	CompletedImports  int     `json:"completed_imports"`
	FailedImports     int     `json:"failed_imports"`
	InProgressImports int     `json:"in_progress_imports"`
	TotalProducts     int     `json:"total_products"`
	ImportedProducts  int     `json:"imported_products"`
	FailedProducts    int     `json:"failed_products"`
	SuccessRate       float64 `json:"success_rate"`
	AverageSpeed      float64 `json:"average_speed"` // محصولات در دقیقه
}

// ImportStatuses وضعیت‌های مختلف import
var ImportStatuses = struct {
	Pending    string
	InProgress string
	Completed  string
	Failed     string
	Cancelled  string
}{
	Pending:    "pending",
	InProgress: "in_progress",
	Completed:  "completed",
	Failed:     "failed",
	Cancelled:  "cancelled",
}

// LogLevels سطوح مختلف لاگ
var LogLevels = struct {
	Info    string
	Warning string
	Error   string
}{
	Info:    "info",
	Warning: "warning",
	Error:   "error",
}
