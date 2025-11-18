package unitofwork

import (
	"context"
	"fmt"
	"my-go-backend/internal/repository"

	"gorm.io/gorm"
)

// unitOfWorkImpl پیاده‌سازی عملی اینترفیس UnitOfWork
// این پیاده‌سازی از GORM برای مدیریت تراکنش‌ها استفاده می‌کند
type unitOfWorkImpl struct {
	db *gorm.DB // اتصال اصلی دیتابیس
	tx *gorm.DB // تراکنش فعال

	// ریپازیتوری‌های lazy-loaded
	mediaFileRepo          repository.MediaFileRepositoryInterface
	mediaVersionRepo       repository.MediaVersionRepositoryInterface
	compressionJobRepo     repository.CompressionJobRepositoryInterface
	compressionSettingRepo repository.CompressionSettingRepositoryInterface
	compressionStatRepo    repository.CompressionStatRepositoryInterface

	// ریپازیتوری‌های مرتبط با محصولات و فروشگاه
	brandRepo             repository.BrandRepositoryInterface
	categoryRepo          repository.CategoryRepositoryInterface
	productRepo           repository.ProductRepositoryInterface
	attributeRepo         repository.AttributeRepositoryInterface
	attributeGroupRepo    repository.AttributeGroupRepositoryInterface
	reviewRepo            repository.ReviewRepositoryInterface
	wishlistRepo          repository.WishlistRepositoryInterface
	stockNotificationRepo repository.StockNotificationRepositoryInterface
	productQARepo         repository.ProductQARepositoryInterface

	// سایر ریپازیتوری‌ها
	menuRepo repository.MenuRepositoryInterface

	// انبار
	warehouseRepo     repository.WarehouseRepositoryInterface
	pwsRepo           repository.ProductWarehouseStockRepositoryInterface
	inventoryMoveRepo repository.InventoryMovementRepositoryInterface

	// احراز هویت/تاریخچه لاگین
	loginAttemptRepo repository.LoginAttemptRepositoryInterface

	// کیف پول
	walletRepo      repository.WalletRepositoryInterface
	walletTxRepo    repository.WalletTransactionRepositoryInterface
	bankAccountRepo repository.BankAccountRepositoryInterface

	// Fraud repositories
	fraudCaseRepo      repository.FraudCaseRepositoryInterface
	fraudEventRepo     repository.FraudEventRepositoryInterface
	fraudRuleRepo      repository.FraudRuleRepositoryInterface
	fraudListRepo      repository.FraudListRepositoryInterface
	fraudActionLogRepo repository.FraudActionLogRepositoryInterface

	// Traffic repositories
	trafficLogRepo     repository.TrafficLogRepositoryInterface
	trafficSessionRepo repository.TrafficSessionRepositoryInterface
	blockedIPRepo      repository.BlockedIPRepositoryInterface

	// Chat repositories
	chatSessionRepo  repository.ChatSessionRepositoryInterface
	chatMessageRepo  repository.ChatMessageRepositoryInterface
	chatOperatorRepo repository.ChatOperatorRepositoryInterface
	chatWidgetRepo   repository.ChatWidgetRepositoryInterface
	chatAIBotRepo    repository.ChatAIBotRepositoryInterface
	chatKBRepo       repository.ChatKnowledgeBaseRepositoryInterface
}

// NewGormUnitOfWork ایجاد یک نمونه جدید از UnitOfWork
// db: اتصال GORM به دیتابیس
func NewGormUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWorkImpl{
		db: db,
	}
}

// BeginTx شروع یک تراکنش جدید با گزینه‌های داده شده
func (u *unitOfWorkImpl) BeginTx(ctx context.Context, opts ...*gorm.DB) error {
	if u.tx != nil {
		// اگر تراکنش قبلی وجود دارد، ابتدا آن را rollback کن
		if err := u.tx.Rollback().Error; err != nil {
			return fmt.Errorf("خطا در rollback تراکنش قبلی: %w", err)
		}
		u.tx = nil
		u.resetRepositories()
	}

	// اگر گزینه‌های سفارشی ارائه شده، از آن استفاده کن
	if len(opts) > 0 && opts[0] != nil {
		u.tx = opts[0].WithContext(ctx).Begin()
	} else {
		// در غیر این صورت، تراکنش معمولی آغاز کن
		u.tx = u.db.WithContext(ctx).Begin()
	}

	if u.tx.Error != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", u.tx.Error)
	}

	return nil
}

// Commit ذخیره تمام تغییرات انجام شده در تراکنش
func (u *unitOfWorkImpl) Commit() error {
	if u.tx == nil {
		return fmt.Errorf("هیچ تراکنش فعالی وجود ندارد")
	}

	if err := u.tx.Commit().Error; err != nil {
		// در صورت خطا در commit، تراکنش را rollback کن
		if rollbackErr := u.tx.Rollback().Error; rollbackErr != nil {
			return fmt.Errorf("خطا در تأیید تراکنش: %w و خطا در rollback: %v", err, rollbackErr)
		}
		return fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	// پاکسازی منابع
	u.resetRepositories()
	u.tx = nil

	return nil
}

// Rollback برگرداندن تمام تغییرات انجام شده در تراکنش
func (u *unitOfWorkImpl) Rollback() error {
	if u.tx == nil {
		return fmt.Errorf("هیچ تراکنش فعالی وجود ندارد")
	}

	if err := u.tx.Rollback().Error; err != nil {
		return fmt.Errorf("خطا در بازگردانی تراکنش: %w", err)
	}

	// پاکسازی منابع
	u.resetRepositories()
	u.tx = nil

	return nil
}

// GetTx دریافت نمونه تراکنش فعال
func (u *unitOfWorkImpl) GetTx() *gorm.DB {
	if u.tx != nil {
		return u.tx
	}
	return u.db
}

// DB بازگرداندن نمونه GORM DB (اتصال یا تراکنش)
// این متد برای سازگاری با کدهای قدیمی اضافه شده
func (u *unitOfWorkImpl) DB() *gorm.DB {
	return u.GetTx()
}

// resetRepositories پاکسازی تمام ریپازیتوری‌های ساخته شده
// این متد پس از commit یا rollback فراخوانی می‌شود
func (u *unitOfWorkImpl) resetRepositories() {
	u.mediaFileRepo = nil
	u.mediaVersionRepo = nil
	u.compressionJobRepo = nil
	u.compressionSettingRepo = nil
	u.compressionStatRepo = nil

	// ریپازیتوری‌های محصولات و فروشگاه
	u.brandRepo = nil
	u.categoryRepo = nil
	u.productRepo = nil
	u.attributeRepo = nil
	u.attributeGroupRepo = nil
	u.reviewRepo = nil
	u.wishlistRepo = nil
	u.stockNotificationRepo = nil
	u.productQARepo = nil

	// سایر ریپازیتوری‌ها
	u.menuRepo = nil
	u.warehouseRepo = nil
	u.pwsRepo = nil
	u.inventoryMoveRepo = nil
	u.loginAttemptRepo = nil

	// کیف پول
	u.walletRepo = nil
	u.walletTxRepo = nil
	u.bankAccountRepo = nil

	// Fraud
	u.fraudCaseRepo = nil
	u.fraudEventRepo = nil
	u.fraudRuleRepo = nil
	u.fraudListRepo = nil
	u.fraudActionLogRepo = nil

	// Traffic
	u.trafficLogRepo = nil
	u.trafficSessionRepo = nil
	u.blockedIPRepo = nil

	// Chat
	u.chatSessionRepo = nil
	u.chatMessageRepo = nil
	u.chatOperatorRepo = nil
	u.chatWidgetRepo = nil
	u.chatAIBotRepo = nil
	u.chatKBRepo = nil
}

// MediaFileRepository دریافت یا ایجاد ریپازیتوری فایل‌های رسانه
// از الگوی lazy loading استفاده می‌کند
func (u *unitOfWorkImpl) MediaFileRepository() repository.MediaFileRepositoryInterface {
	if u.mediaFileRepo == nil {
		u.mediaFileRepo = &repository.MediaFileRepository{
			DB: u.GetTx(),
		}
	}
	return u.mediaFileRepo
}

// MediaVersionRepository دریافت یا ایجاد ریپازیتوری نسخه‌های رسانه
func (u *unitOfWorkImpl) MediaVersionRepository() repository.MediaVersionRepositoryInterface {
	if u.mediaVersionRepo == nil {
		u.mediaVersionRepo = &repository.MediaVersionRepository{
			DB: u.GetTx(),
		}
	}
	return u.mediaVersionRepo
}

// CompressionJobRepository دریافت یا ایجاد ریپازیتوری کارهای فشرده‌سازی
func (u *unitOfWorkImpl) CompressionJobRepository() repository.CompressionJobRepositoryInterface {
	if u.compressionJobRepo == nil {
		u.compressionJobRepo = &repository.CompressionJobRepository{
			DB: u.GetTx(),
		}
	}
	return u.compressionJobRepo
}

// CompressionSettingRepository دریافت یا ایجاد ریپازیتوری تنظیمات فشرده‌سازی
func (u *unitOfWorkImpl) CompressionSettingRepository() repository.CompressionSettingRepositoryInterface {
	if u.compressionSettingRepo == nil {
		u.compressionSettingRepo = &repository.CompressionSettingRepository{
			DB: u.GetTx(),
		}
	}
	return u.compressionSettingRepo
}

// CompressionStatRepository دریافت یا ایجاد ریپازیتوری آمار فشرده‌سازی
func (u *unitOfWorkImpl) CompressionStatRepository() repository.CompressionStatRepositoryInterface {
	if u.compressionStatRepo == nil {
		u.compressionStatRepo = &repository.CompressionStatRepository{
			DB: u.GetTx(),
		}
	}
	return u.compressionStatRepo
}

// ==================== Product & E-commerce Repositories ====================

// BrandRepository دریافت یا ایجاد ریپازیتوری برندها
func (u *unitOfWorkImpl) BrandRepository() repository.BrandRepositoryInterface {
	if u.brandRepo == nil {
		u.brandRepo = repository.NewBrandRepository(u.GetTx())
	}
	return u.brandRepo
}

// CategoryRepository دریافت یا ایجاد ریپازیتوری دسته‌بندی‌ها
func (u *unitOfWorkImpl) CategoryRepository() repository.CategoryRepositoryInterface {
	if u.categoryRepo == nil {
		u.categoryRepo = repository.NewCategoryRepository(u.GetTx())
	}
	return u.categoryRepo
}

// ProductRepository دریافت یا ایجاد ریپازیتوری محصولات
func (u *unitOfWorkImpl) ProductRepository() repository.ProductRepositoryInterface {
	if u.productRepo == nil {
		u.productRepo = repository.NewProductRepository(u.GetTx())
	}
	return u.productRepo
}

// AttributeRepository دریافت یا ایجاد ریپازیتوری ویژگی‌ها
func (u *unitOfWorkImpl) AttributeRepository() repository.AttributeRepositoryInterface {
	if u.attributeRepo == nil {
		u.attributeRepo = repository.NewAttributeRepository(u.GetTx())
	}
	return u.attributeRepo
}

// AttributeGroupRepository دریافت یا ایجاد ریپازیتوری گروه‌های ویژگی
func (u *unitOfWorkImpl) AttributeGroupRepository() repository.AttributeGroupRepositoryInterface {
	if u.attributeGroupRepo == nil {
		u.attributeGroupRepo = repository.NewAttributeGroupRepository(u.GetTx())
	}
	return u.attributeGroupRepo
}

// ReviewRepository دریافت یا ایجاد ریپازیتوری نظرات
func (u *unitOfWorkImpl) ReviewRepository() repository.ReviewRepositoryInterface {
	if u.reviewRepo == nil {
		u.reviewRepo = repository.NewReviewRepository(u.GetTx())
	}
	return u.reviewRepo
}

// WishlistRepository دریافت یا ایجاد ریپازیتوری لیست علاقمندی‌ها
func (u *unitOfWorkImpl) WishlistRepository() repository.WishlistRepositoryInterface {
	if u.wishlistRepo == nil {
		u.wishlistRepo = repository.NewWishlistRepository(u.GetTx())
	}
	return u.wishlistRepo
}

// StockNotificationRepository دریافت یا ایجاد ریپازیتوری اعلان‌های موجودی
func (u *unitOfWorkImpl) StockNotificationRepository() repository.StockNotificationRepositoryInterface {
	if u.stockNotificationRepo == nil {
		u.stockNotificationRepo = repository.NewStockNotificationRepository(u.GetTx())
	}
	return u.stockNotificationRepo
}

// ProductQARepository دریافت یا ایجاد ریپازیتوری پرسش و پاسخ محصولات
func (u *unitOfWorkImpl) ProductQARepository() repository.ProductQARepositoryInterface {
	if u.productQARepo == nil {
		u.productQARepo = repository.NewProductQARepository(u.GetTx())
	}
	return u.productQARepo
}

// MenuRepository دریافت یا ایجاد ریپازیتوری منوها
func (u *unitOfWorkImpl) MenuRepository() repository.MenuRepositoryInterface {
	if u.menuRepo == nil {
		u.menuRepo = repository.NewMenuRepository(u.GetTx())
	}
	return u.menuRepo
}

// WarehouseRepository دریافت یا ایجاد ریپازیتوری انبارها
func (u *unitOfWorkImpl) WarehouseRepository() repository.WarehouseRepositoryInterface {
	if u.warehouseRepo == nil {
		u.warehouseRepo = repository.NewWarehouseRepository(u.GetTx())
	}
	return u.warehouseRepo
}

// ProductWarehouseStockRepository دریافت ریپازیتوری موجودی محصول-انبار
func (u *unitOfWorkImpl) ProductWarehouseStockRepository() repository.ProductWarehouseStockRepositoryInterface {
	if u.pwsRepo == nil {
		u.pwsRepo = repository.NewProductWarehouseStockRepository(u.GetTx())
	}
	return u.pwsRepo
}

// InventoryMovementRepository دریافت ریپازیتوری حرکات موجودی
func (u *unitOfWorkImpl) InventoryMovementRepository() repository.InventoryMovementRepositoryInterface {
	if u.inventoryMoveRepo == nil {
		u.inventoryMoveRepo = repository.NewInventoryMovementRepository(u.GetTx())
	}
	return u.inventoryMoveRepo
}

// LoginAttemptRepository دریافت یا ایجاد ریپازیتوری تاریخچه ورود
func (u *unitOfWorkImpl) LoginAttemptRepository() repository.LoginAttemptRepositoryInterface {
	if u.loginAttemptRepo == nil {
		u.loginAttemptRepo = &repository.LoginAttemptRepository{DB: u.GetTx()}
	}
	return u.loginAttemptRepo
}

// WalletRepository دریافت یا ایجاد ریپازیتوری کیف پول
func (u *unitOfWorkImpl) WalletRepository() repository.WalletRepositoryInterface {
	if u.walletRepo == nil {
		u.walletRepo = repository.NewWalletRepository(u.GetTx())
	}
	return u.walletRepo
}

// WalletTransactionRepository دریافت یا ایجاد ریپازیتوری تراکنش‌های کیف پول
func (u *unitOfWorkImpl) WalletTransactionRepository() repository.WalletTransactionRepositoryInterface {
	if u.walletTxRepo == nil {
		u.walletTxRepo = repository.NewWalletTransactionRepository(u.GetTx())
	}
	return u.walletTxRepo
}

// BankAccountRepository دریافت یا ایجاد ریپازیتوری کارت/حساب بانکی
func (u *unitOfWorkImpl) BankAccountRepository() repository.BankAccountRepositoryInterface {
	if u.bankAccountRepo == nil {
		u.bankAccountRepo = repository.NewBankAccountRepository(u.GetTx())
	}
	return u.bankAccountRepo
}

// ==================== Fraud Repositories ====================

func (u *unitOfWorkImpl) FraudCaseRepository() repository.FraudCaseRepositoryInterface {
	if u.fraudCaseRepo == nil {
		u.fraudCaseRepo = &repository.FraudCaseRepository{DB: u.GetTx()}
	}
	return u.fraudCaseRepo
}

func (u *unitOfWorkImpl) FraudEventRepository() repository.FraudEventRepositoryInterface {
	if u.fraudEventRepo == nil {
		u.fraudEventRepo = &repository.FraudEventRepository{DB: u.GetTx()}
	}
	return u.fraudEventRepo
}

func (u *unitOfWorkImpl) FraudRuleRepository() repository.FraudRuleRepositoryInterface {
	if u.fraudRuleRepo == nil {
		u.fraudRuleRepo = &repository.FraudRuleRepository{DB: u.GetTx()}
	}
	return u.fraudRuleRepo
}

func (u *unitOfWorkImpl) FraudListRepository() repository.FraudListRepositoryInterface {
	if u.fraudListRepo == nil {
		u.fraudListRepo = &repository.FraudListRepository{DB: u.GetTx()}
	}
	return u.fraudListRepo
}

func (u *unitOfWorkImpl) FraudActionLogRepository() repository.FraudActionLogRepositoryInterface {
	if u.fraudActionLogRepo == nil {
		u.fraudActionLogRepo = &repository.FraudActionLogRepository{DB: u.GetTx()}
	}
	return u.fraudActionLogRepo
}

// TrafficLogRepository دریافت یا ایجاد ریپازیتوری لاگ ترافیک
func (u *unitOfWorkImpl) TrafficLogRepository() repository.TrafficLogRepositoryInterface {
	if u.trafficLogRepo == nil {
		u.trafficLogRepo = &repository.TrafficLogRepository{DB: u.GetTx()}
	}
	return u.trafficLogRepo
}

func (u *unitOfWorkImpl) TrafficSessionRepository() repository.TrafficSessionRepositoryInterface {
	if u.trafficSessionRepo == nil {
		u.trafficSessionRepo = &repository.TrafficSessionRepository{DB: u.GetTx()}
	}
	return u.trafficSessionRepo
}

func (u *unitOfWorkImpl) BlockedIPRepository() repository.BlockedIPRepositoryInterface {
	if u.blockedIPRepo == nil {
		u.blockedIPRepo = &repository.BlockedIPRepository{DB: u.GetTx()}
	}
	return u.blockedIPRepo
}

// ==================== Chat Repositories ====================

func (u *unitOfWorkImpl) ChatSessionRepository() repository.ChatSessionRepositoryInterface {
	if u.chatSessionRepo == nil {
		u.chatSessionRepo = &repository.ChatSessionRepository{DB: u.GetTx()}
	}
	return u.chatSessionRepo
}

func (u *unitOfWorkImpl) ChatMessageRepository() repository.ChatMessageRepositoryInterface {
	if u.chatMessageRepo == nil {
		u.chatMessageRepo = &repository.ChatMessageRepository{DB: u.GetTx()}
	}
	return u.chatMessageRepo
}

func (u *unitOfWorkImpl) ChatOperatorRepository() repository.ChatOperatorRepositoryInterface {
	if u.chatOperatorRepo == nil {
		u.chatOperatorRepo = &repository.ChatOperatorRepository{DB: u.GetTx()}
	}
	return u.chatOperatorRepo
}

// ChatWidgetRepository دریافت یا ایجاد ریپازیتوری ابزارک چت
func (u *unitOfWorkImpl) ChatWidgetRepository() repository.ChatWidgetRepositoryInterface {
	if u.chatWidgetRepo == nil {
		u.chatWidgetRepo = &repository.ChatWidgetRepository{DB: u.GetTx()}
	}
	return u.chatWidgetRepo
}

// ChatAIBotRepository دریافت ریپازیتوری ربات هوش مصنوعی چت
func (u *unitOfWorkImpl) ChatAIBotRepository() repository.ChatAIBotRepositoryInterface {
	if u.chatAIBotRepo == nil {
		u.chatAIBotRepo = &repository.ChatAIBotRepository{DB: u.GetTx()}
	}
	return u.chatAIBotRepo
}

// ChatKnowledgeBaseRepository دریافت ریپازیتوری پایگاه دانش چت
func (u *unitOfWorkImpl) ChatKnowledgeBaseRepository() repository.ChatKnowledgeBaseRepositoryInterface {
	if u.chatKBRepo == nil {
		u.chatKBRepo = &repository.ChatKnowledgeBaseRepository{DB: u.GetTx()}
	}
	return u.chatKBRepo
}
