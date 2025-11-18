package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// SearchIndexQueue جدول صف همگام‌سازی ایندکس‌های جستجو
// رکوردها توسط گوروتین Indexer پردازش و پس از موفقیت به وضعیت completed تغییر می‌کنند.
type SearchIndexQueue struct {
	ID           uint64            `json:"id" gorm:"primaryKey"`
	ResourceType string            `json:"resource_type" gorm:"size:50;not null"`
	ResourceID   uint64            `json:"resource_id" gorm:"not null"`
	Operation    string            `json:"operation" gorm:"size:20;not null"`
	Payload      datatypes.JSONMap `json:"payload,omitempty"`
	Status       string            `json:"status" gorm:"size:20;not null;default:pending"`
	Attempts     int               `json:"attempts" gorm:"not null;default:0"`
	ErrorMessage string            `json:"error_message"`
	AvailableAt  time.Time         `json:"available_at" gorm:"not null;default:now()"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

const (
	// SearchOpUpsert ثبت یا به‌روزرسانی سند در ایندکس
	SearchOpUpsert = "upsert"
	// SearchOpDelete حذف سند از ایندکس
	SearchOpDelete = "delete"

	// SearchResourceProduct شناسه نوع منبع برای محصولات
	SearchResourceProduct = "product"
	// SearchResourcePost شناسه نوع منبع برای نوشته‌ها/مقالات
	SearchResourcePost = "post"

	// SearchStatusPending رکورد در انتظار پردازش
	SearchStatusPending = "pending"
	// SearchStatusCompleted پردازش موفق
	SearchStatusCompleted = "completed"
	// SearchStatusFailed پردازش ناموفق (پس از اتمام تلاش‌ها)
	SearchStatusFailed = "failed"
)

// TableName نام جدول پایگاه داده
func (SearchIndexQueue) TableName() string {
	return "search_index_queue"
}

// enqueueSearchEvent درج رکورد جدید در صف ایندکس با اطمینان از صحت ورودی‌ها
func enqueueSearchEvent(tx *gorm.DB, resourceType string, resourceID uint64, operation string) error {
	if tx == nil {
		return errors.New("nil transaction for search indexing")
	}

	resourceType = strings.TrimSpace(strings.ToLower(resourceType))
	if resourceType == "" {
		return errors.New("empty resource type")
	}

	op := strings.TrimSpace(strings.ToLower(operation))
	switch op {
	case SearchOpUpsert, SearchOpDelete:
	default:
		return errors.New("invalid operation for search indexing")
	}

	event := &SearchIndexQueue{
		ResourceType: resourceType,
		ResourceID:   resourceID,
		Operation:    op,
		Status:       SearchStatusPending,
	}

	session := tx.Session(&gorm.Session{NewDB: true})
	if ctx := tx.Statement.Context; ctx != nil {
		session = session.WithContext(ctx)
	}

	return session.Model(&SearchIndexQueue{}).Create(event).Error
}
