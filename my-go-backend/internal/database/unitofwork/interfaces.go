package unitofwork

import (
	"context"
	"my-go-backend/internal/repository"

	"gorm.io/gorm"
)

// UnitOfWork اینترفیس اصلی برای مدیریت تراکنش‌ها
// این اینترفیس امکان مدیریت چندین عملیات دیتابیس را در یک تراکنش فراهم می‌کند
// و تضمین می‌کند که یا همه عملیات‌ها موفق می‌شوند یا هیچکدام (اتمیسیتی)
type UnitOfWork interface {
	// BeginTx شروع یک تراکنش جدید
	// ctx: کانتکست برای کنترل timeout و cancellation
	// opts: گزینه‌های تراکنش (مثل isolation level)
	BeginTx(ctx context.Context, opts ...*gorm.DB) error

	// Commit تأیید و ذخیره تمام تغییرات انجام شده در تراکنش
	Commit() error

	// Rollback برگرداندن تمام تغییرات انجام شده در تراکنش
	Rollback() error

	// GetTx دریافت نمونه تراکنش فعال برای استفاده مستقیم (در موارد خاص)
	GetTx() *gorm.DB

	// DB دریافت نمونه GORM (اتصال یا تراکنش) برای استفاده‌های موردی
	DB() *gorm.DB

	// ریپازیتوری‌های مرتبط با رسانه
	MediaFileRepository() repository.MediaFileRepositoryInterface
	MediaVersionRepository() repository.MediaVersionRepositoryInterface
	CompressionJobRepository() repository.CompressionJobRepositoryInterface
	CompressionSettingRepository() repository.CompressionSettingRepositoryInterface
	CompressionStatRepository() repository.CompressionStatRepositoryInterface

	// ریپازیتوری‌های مرتبط با محصولات و فروشگاه
	BrandRepository() repository.BrandRepositoryInterface
	CategoryRepository() repository.CategoryRepositoryInterface
	ProductRepository() repository.ProductRepositoryInterface
	AttributeRepository() repository.AttributeRepositoryInterface
	AttributeGroupRepository() repository.AttributeGroupRepositoryInterface
	ReviewRepository() repository.ReviewRepositoryInterface
	WishlistRepository() repository.WishlistRepositoryInterface
	StockNotificationRepository() repository.StockNotificationRepositoryInterface
	ProductQARepository() repository.ProductQARepositoryInterface

	// سایر ریپازیتوری‌ها
	MenuRepository() repository.MenuRepositoryInterface

	// انبار و موجودی چندانباره
	WarehouseRepository() repository.WarehouseRepositoryInterface
	ProductWarehouseStockRepository() repository.ProductWarehouseStockRepositoryInterface
	InventoryMovementRepository() repository.InventoryMovementRepositoryInterface

	// ریپازیتوری‌های لاگین/احراز هویت
	LoginAttemptRepository() repository.LoginAttemptRepositoryInterface

	// ریپازیتوری‌های کیف پول
	WalletRepository() repository.WalletRepositoryInterface
	WalletTransactionRepository() repository.WalletTransactionRepositoryInterface

	// ریپازیتوری حساب/کارت بانکی
	BankAccountRepository() repository.BankAccountRepositoryInterface

	// ریپازیتوری‌های Fraud
	FraudCaseRepository() repository.FraudCaseRepositoryInterface
	FraudEventRepository() repository.FraudEventRepositoryInterface
	FraudRuleRepository() repository.FraudRuleRepositoryInterface
	FraudListRepository() repository.FraudListRepositoryInterface
	FraudActionLogRepository() repository.FraudActionLogRepositoryInterface

	// ریپازیتوری‌های ترافیک
	TrafficLogRepository() repository.TrafficLogRepositoryInterface
	TrafficSessionRepository() repository.TrafficSessionRepositoryInterface
	BlockedIPRepository() repository.BlockedIPRepositoryInterface

	// ریپازیتوری چت
	ChatSessionRepository() repository.ChatSessionRepositoryInterface
	ChatMessageRepository() repository.ChatMessageRepositoryInterface
	// ریپازیتوری ابزارک، بات و پایگاه دانش چت
	ChatWidgetRepository() repository.ChatWidgetRepositoryInterface
	ChatAIBotRepository() repository.ChatAIBotRepositoryInterface
	ChatKnowledgeBaseRepository() repository.ChatKnowledgeBaseRepositoryInterface
	ChatOperatorRepository() repository.ChatOperatorRepositoryInterface
}

// UnitOfWorkFactory اینترفیس برای ایجاد نمونه‌های جدید UnitOfWork
// این الگو امکان تست و جایگزینی پیاده‌سازی‌های مختلف را فراهم می‌کند
type UnitOfWorkFactory interface {
	// Create ایجاد یک نمونه جدید از UnitOfWork
	Create() UnitOfWork
}

// TransactionOptions گزینه‌های پیکربندی تراکنش
type TransactionOptions struct {
	// IsolationLevel سطح ایزولیشن تراکنش
	// مقادیر ممکن: ReadUncommitted, ReadCommitted, RepeatableRead, Serializable
	IsolationLevel string

	// ReadOnly آیا تراکنش فقط خواندنی است
	ReadOnly bool

	// TimeoutSeconds حداکثر زمان اجرای تراکنش
	TimeoutSeconds int
}
