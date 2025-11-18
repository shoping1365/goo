package repository

import (
	"context"
	"my-go-backend/internal/models"
)

// ChatSessionRepositoryInterface اینترفیس ریپازیتوری جلسات چت
// Rule 0: تمام عملیات فقط در لایه ریپازیتوری انجام شود و هیچ منطق بیزینسی اینجا قرار نگیرد.
type ChatSessionRepositoryInterface interface {
	// Create ایجاد جلسه جدید
	Create(ctx context.Context, session *models.ChatSession) error

	// GetByID دریافت بر اساس شناسه دیتابیس
	GetByID(ctx context.Context, id uint) (*models.ChatSession, error)

	// GetBySessionID دریافت بر اساس شناسه عمومی سشن (uuid/random)
	GetBySessionID(ctx context.Context, sessionID string) (*models.ChatSession, error)

	// GetWaitingSessions دریافت جلسات در وضعیت waiting
	GetWaitingSessions(ctx context.Context) ([]models.ChatSession, error)

	// ListSessionsByStatus دریافت لیست بر اساس وضعیت با صفحه‌بندی
	ListSessionsByStatus(ctx context.Context, status string, limit, offset int) ([]models.ChatSession, error)

	// UpdateFields به روزرسانی فیلدهای دلخواه با map
	UpdateFields(ctx context.Context, id uint, fields map[string]interface{}) error
}

// ChatMessageRepositoryInterface اینترفیس ریپازیتوری پیام‌های چت
type ChatMessageRepositoryInterface interface {
	// Create ایجاد پیام جدید
	Create(ctx context.Context, msg *models.OnlineChatMessage) error

	// GetBySessionID دریافت پیام‌های یک جلسه با صفحه‌بندی (جدیدترین ابتدا)
	GetBySessionID(ctx context.Context, sessionID uint, limit, offset int) ([]models.OnlineChatMessage, error)

	// MarkAsRead خواندن یک پیام
	MarkAsRead(ctx context.Context, messageID uint) error

	// MarkSessionMessagesAsRead خواندن همه پیام‌های یک جلسه
	MarkSessionMessagesAsRead(ctx context.Context, sessionID uint) error
}

// ChatOperatorRepositoryInterface اینترفیس ریپازیتوری اپراتورها
// فقط عملیات مورد نیاز فعلی پیاده‌سازی شده است.
type ChatOperatorRepositoryInterface interface {
	GetByID(ctx context.Context, id uint) (*models.ChatOperator, error)
	UpdateFields(ctx context.Context, id uint, fields map[string]interface{}) error
}

// ChatWidgetRepositoryInterface اینترفیس ریپازیتوری ابزارک‌های چت
type ChatWidgetRepositoryInterface interface {
	Create(ctx context.Context, widget *models.ChatWidget) error
	GetActive(ctx context.Context) ([]models.ChatWidget, error)
	GetByID(ctx context.Context, id uint) (*models.ChatWidget, error)
	Update(ctx context.Context, widget *models.ChatWidget) error
	Delete(ctx context.Context, id uint) error
}

// ChatAIBotRepositoryInterface اینترفیس ریپازیتوری ربات‌های هوش مصنوعی چت
type ChatAIBotRepositoryInterface interface {
	Create(ctx context.Context, bot *models.ChatAIBot) error
	GetActive(ctx context.Context) ([]models.ChatAIBot, error)
	GetByID(ctx context.Context, id uint) (*models.ChatAIBot, error)
	Update(ctx context.Context, bot *models.ChatAIBot) error
	Delete(ctx context.Context, id uint) error
}

// ChatKnowledgeBaseRepositoryInterface اینترفیس ریپازیتوری پایگاه دانش
// شامل متدهای جستجو و فهرست
// برای سادگی متد List ترکیب count را برمی‌گرداند
// (می‌تواند در CQRS به QueryHandler جدا منتقل شود)
type ChatKnowledgeBaseRepositoryInterface interface {
	Create(ctx context.Context, kb *models.ChatKnowledgeBase) error
	Search(ctx context.Context, q string) ([]models.ChatKnowledgeBase, error)
	List(ctx context.Context, q string, limit, offset int) ([]models.ChatKnowledgeBase, int64, error)
	GetByID(ctx context.Context, id uint) (*models.ChatKnowledgeBase, error)
	Update(ctx context.Context, kb *models.ChatKnowledgeBase) error
	Delete(ctx context.Context, id uint) error
}
