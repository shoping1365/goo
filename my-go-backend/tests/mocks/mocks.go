package mocks

import (
	"context"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"time"

	"gorm.io/gorm"
)

// MockUnitOfWorkFactory is a mock implementation of UnitOfWorkFactory
type MockUnitOfWorkFactory struct {
	CreateFunc func() unitofwork.UnitOfWork
}

func (m *MockUnitOfWorkFactory) Create() unitofwork.UnitOfWork {
	if m.CreateFunc != nil {
		return m.CreateFunc()
	}
	return &MockUnitOfWork{}
}

// MockUnitOfWork is a mock implementation of UnitOfWork
type MockUnitOfWork struct {
	BeginTxFunc func(ctx context.Context, opts ...*gorm.DB) error
	CommitFunc  func() error
	RollbackFunc func() error
	
	WalletRepoFunc func() repository.WalletRepositoryInterface
	WalletTxRepoFunc func() repository.WalletTransactionRepositoryInterface
	
	InventoryMovementRepoFunc func() repository.InventoryMovementRepositoryInterface
	ProductWarehouseStockRepoFunc func() repository.ProductWarehouseStockRepositoryInterface
	WarehouseRepoFunc func() repository.WarehouseRepositoryInterface

	BlockedIPRepoFunc func() repository.BlockedIPRepositoryInterface
	TrafficLogRepoFunc func() repository.TrafficLogRepositoryInterface
	TrafficSessionRepoFunc func() repository.TrafficSessionRepositoryInterface
	CategoryRepoFunc func() repository.CategoryRepositoryInterface

	FraudRuleRepoFunc func() repository.FraudRuleRepositoryInterface
	FraudCaseRepoFunc func() repository.FraudCaseRepositoryInterface
	FraudEventRepoFunc func() repository.FraudEventRepositoryInterface

	ProductRepoFunc func() repository.ProductRepositoryInterface
	BrandRepoFunc func() repository.BrandRepositoryInterface

	// Add other repo funcs as needed
}

func (m *MockUnitOfWork) BeginTx(ctx context.Context, opts ...*gorm.DB) error {
	if m.BeginTxFunc != nil {
		return m.BeginTxFunc(ctx, opts...)
	}
	return nil
}

func (m *MockUnitOfWork) Commit() error {
	if m.CommitFunc != nil {
		return m.CommitFunc()
	}
	return nil
}

func (m *MockUnitOfWork) Rollback() error {
	if m.RollbackFunc != nil {
		return m.RollbackFunc()
	}
	return nil
}

func (m *MockUnitOfWork) GetTx() *gorm.DB { return nil }
func (m *MockUnitOfWork) DB() *gorm.DB { return nil }

func (m *MockUnitOfWork) WalletRepository() repository.WalletRepositoryInterface {
	if m.WalletRepoFunc != nil {
		return m.WalletRepoFunc()
	}
	return &MockWalletRepository{}
}

func (m *MockUnitOfWork) WalletTransactionRepository() repository.WalletTransactionRepositoryInterface {
	if m.WalletTxRepoFunc != nil {
		return m.WalletTxRepoFunc()
	}
	return &MockWalletTransactionRepository{}
}

// Implement other methods of UnitOfWork interface with panic or empty return to satisfy interface
func (m *MockUnitOfWork) MediaFileRepository() repository.MediaFileRepositoryInterface { return nil }
func (m *MockUnitOfWork) MediaVersionRepository() repository.MediaVersionRepositoryInterface { return nil }
func (m *MockUnitOfWork) CompressionJobRepository() repository.CompressionJobRepositoryInterface { return nil }
func (m *MockUnitOfWork) CompressionSettingRepository() repository.CompressionSettingRepositoryInterface { return nil }
func (m *MockUnitOfWork) CompressionStatRepository() repository.CompressionStatRepositoryInterface { return nil }
func (m *MockUnitOfWork) BrandRepository() repository.BrandRepositoryInterface {
	if m.BrandRepoFunc != nil {
		return m.BrandRepoFunc()
	}
	return &MockBrandRepository{}
}
func (m *MockUnitOfWork) CategoryRepository() repository.CategoryRepositoryInterface {
	if m.CategoryRepoFunc != nil {
		return m.CategoryRepoFunc()
	}
	return &MockCategoryRepository{}
}
func (m *MockUnitOfWork) ProductRepository() repository.ProductRepositoryInterface {
	if m.ProductRepoFunc != nil {
		return m.ProductRepoFunc()
	}
	return &MockProductRepository{}
}
func (m *MockUnitOfWork) AttributeRepository() repository.AttributeRepositoryInterface { return nil }
func (m *MockUnitOfWork) AttributeGroupRepository() repository.AttributeGroupRepositoryInterface { return nil }
func (m *MockUnitOfWork) ReviewRepository() repository.ReviewRepositoryInterface { return nil }
func (m *MockUnitOfWork) WishlistRepository() repository.WishlistRepositoryInterface { return nil }
func (m *MockUnitOfWork) StockNotificationRepository() repository.StockNotificationRepositoryInterface { return nil }
func (m *MockUnitOfWork) ProductQARepository() repository.ProductQARepositoryInterface { return nil }
func (m *MockUnitOfWork) MenuRepository() repository.MenuRepositoryInterface { return nil }
func (m *MockUnitOfWork) WarehouseRepository() repository.WarehouseRepositoryInterface {
	if m.WarehouseRepoFunc != nil {
		return m.WarehouseRepoFunc()
	}
	return &MockWarehouseRepository{}
}
func (m *MockUnitOfWork) ProductWarehouseStockRepository() repository.ProductWarehouseStockRepositoryInterface {
	if m.ProductWarehouseStockRepoFunc != nil {
		return m.ProductWarehouseStockRepoFunc()
	}
	return &MockProductWarehouseStockRepository{}
}
func (m *MockUnitOfWork) InventoryMovementRepository() repository.InventoryMovementRepositoryInterface {
	if m.InventoryMovementRepoFunc != nil {
		return m.InventoryMovementRepoFunc()
	}
	return &MockInventoryMovementRepository{}
}
func (m *MockUnitOfWork) LoginAttemptRepository() repository.LoginAttemptRepositoryInterface { return nil }
func (m *MockUnitOfWork) BankAccountRepository() repository.BankAccountRepositoryInterface { return nil }
func (m *MockUnitOfWork) FraudCaseRepository() repository.FraudCaseRepositoryInterface {
	if m.FraudCaseRepoFunc != nil {
		return m.FraudCaseRepoFunc()
	}
	return &MockFraudCaseRepository{}
}
func (m *MockUnitOfWork) FraudEventRepository() repository.FraudEventRepositoryInterface {
	if m.FraudEventRepoFunc != nil {
		return m.FraudEventRepoFunc()
	}
	return &MockFraudEventRepository{}
}
func (m *MockUnitOfWork) FraudRuleRepository() repository.FraudRuleRepositoryInterface {
	if m.FraudRuleRepoFunc != nil {
		return m.FraudRuleRepoFunc()
	}
	return &MockFraudRuleRepository{}
}
func (m *MockUnitOfWork) FraudListRepository() repository.FraudListRepositoryInterface { return nil }
func (m *MockUnitOfWork) FraudActionLogRepository() repository.FraudActionLogRepositoryInterface { return nil }
func (m *MockUnitOfWork) TrafficLogRepository() repository.TrafficLogRepositoryInterface {
	if m.TrafficLogRepoFunc != nil {
		return m.TrafficLogRepoFunc()
	}
	return &MockTrafficLogRepository{}
}
func (m *MockUnitOfWork) TrafficSessionRepository() repository.TrafficSessionRepositoryInterface {
	if m.TrafficSessionRepoFunc != nil {
		return m.TrafficSessionRepoFunc()
	}
	return &MockTrafficSessionRepository{}
}
func (m *MockUnitOfWork) BlockedIPRepository() repository.BlockedIPRepositoryInterface {
	if m.BlockedIPRepoFunc != nil {
		return m.BlockedIPRepoFunc()
	}
	return &MockBlockedIPRepository{}
}
func (m *MockUnitOfWork) ChatSessionRepository() repository.ChatSessionRepositoryInterface { return nil }
func (m *MockUnitOfWork) ChatMessageRepository() repository.ChatMessageRepositoryInterface { return nil }
func (m *MockUnitOfWork) ChatWidgetRepository() repository.ChatWidgetRepositoryInterface { return nil }
func (m *MockUnitOfWork) ChatAIBotRepository() repository.ChatAIBotRepositoryInterface { return nil }
func (m *MockUnitOfWork) ChatKnowledgeBaseRepository() repository.ChatKnowledgeBaseRepositoryInterface { return nil }
func (m *MockUnitOfWork) ChatOperatorRepository() repository.ChatOperatorRepositoryInterface { return nil }


// MockWalletRepository
type MockWalletRepository struct {
	GetByUserIDFunc func(ctx context.Context, userID uint) (*models.UserWallet, error)
	CreateFunc func(ctx context.Context, wallet *models.UserWallet) error
	UpdateFunc func(ctx context.Context, wallet *models.UserWallet) error
	UpdateBalanceFunc func(ctx context.Context, walletID uint, newBalance float64) error
	UpdateStatusFunc func(ctx context.Context, walletID uint, status string) error
	SumAllBalancesFunc func(ctx context.Context) (float64, error)
}

func (m *MockWalletRepository) GetByUserID(ctx context.Context, userID uint) (*models.UserWallet, error) {
	if m.GetByUserIDFunc != nil {
		return m.GetByUserIDFunc(ctx, userID)
	}
	return nil, nil
}

func (m *MockWalletRepository) Create(ctx context.Context, wallet *models.UserWallet) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, wallet)
	}
	return nil
}

func (m *MockWalletRepository) Update(ctx context.Context, wallet *models.UserWallet) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, wallet)
	}
	return nil
}

func (m *MockWalletRepository) UpdateBalance(ctx context.Context, walletID uint, newBalance float64) error {
	if m.UpdateBalanceFunc != nil {
		return m.UpdateBalanceFunc(ctx, walletID, newBalance)
	}
	return nil
}

func (m *MockWalletRepository) UpdateStatus(ctx context.Context, walletID uint, status string) error {
	if m.UpdateStatusFunc != nil {
		return m.UpdateStatusFunc(ctx, walletID, status)
	}
	return nil
}

func (m *MockWalletRepository) SumAllBalances(ctx context.Context) (float64, error) {
	if m.SumAllBalancesFunc != nil {
		return m.SumAllBalancesFunc(ctx)
	}
	return 0, nil
}

// MockWalletTransactionRepository
type MockWalletTransactionRepository struct {
	CreateFunc func(ctx context.Context, tx *models.WalletTransaction) error
	ListByWalletFunc func(ctx context.Context, walletID uint, limit int, offset int, filter repository.WalletTransactionFilter) ([]models.WalletTransaction, int64, error)
	AdminListTransactionsFunc func(ctx context.Context, limit int, offset int, filter repository.AdminTransactionFilter) ([]repository.AdminTransactionRow, int64, error)
	GetByIDFunc func(ctx context.Context, id uint) (*models.WalletTransaction, error)
	UpdateStatusFunc func(ctx context.Context, id uint, status string) error
}

func (m *MockWalletTransactionRepository) Create(ctx context.Context, tx *models.WalletTransaction) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, tx)
	}
	return nil
}

func (m *MockWalletTransactionRepository) ListByWallet(ctx context.Context, walletID uint, limit int, offset int, filter repository.WalletTransactionFilter) ([]models.WalletTransaction, int64, error) {
	if m.ListByWalletFunc != nil {
		return m.ListByWalletFunc(ctx, walletID, limit, offset, filter)
	}
	return nil, 0, nil
}

func (m *MockWalletTransactionRepository) AdminListTransactions(ctx context.Context, limit int, offset int, filter repository.AdminTransactionFilter) ([]repository.AdminTransactionRow, int64, error) {
	if m.AdminListTransactionsFunc != nil {
		return m.AdminListTransactionsFunc(ctx, limit, offset, filter)
	}
	return nil, 0, nil
}

func (m *MockWalletTransactionRepository) GetByID(ctx context.Context, id uint) (*models.WalletTransaction, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockWalletTransactionRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	if m.UpdateStatusFunc != nil {
		return m.UpdateStatusFunc(ctx, id, status)
	}
	return nil
}

// MockInventoryMovementRepository
type MockInventoryMovementRepository struct {
	CreateFunc func(ctx context.Context, m *models.InventoryMovement) error
	ListByProductFunc func(ctx context.Context, productID uint, limit, offset int) ([]models.InventoryMovement, error)
}

func (m *MockInventoryMovementRepository) Create(ctx context.Context, mm *models.InventoryMovement) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, mm)
	}
	return nil
}

func (m *MockInventoryMovementRepository) ListByProduct(ctx context.Context, productID uint, limit, offset int) ([]models.InventoryMovement, error) {
	if m.ListByProductFunc != nil {
		return m.ListByProductFunc(ctx, productID, limit, offset)
	}
	return nil, nil
}

// MockProductWarehouseStockRepository
type MockProductWarehouseStockRepository struct {
	GetByProductFunc func(ctx context.Context, productID uint) ([]models.ProductWarehouseStock, error)
	GetByProductAndWarehouseForUpdateFunc func(ctx context.Context, productID uint, warehouseID uint) (*models.ProductWarehouseStock, error)
	SumByProductFunc func(ctx context.Context, productID uint) (int, int, error)
	UpsertFunc func(ctx context.Context, s *models.ProductWarehouseStock) error
	AdjustQuantityFunc func(ctx context.Context, productID, warehouseID uint, delta int) error
}

func (m *MockProductWarehouseStockRepository) GetByProduct(ctx context.Context, productID uint) ([]models.ProductWarehouseStock, error) {
	if m.GetByProductFunc != nil {
		return m.GetByProductFunc(ctx, productID)
	}
	return nil, nil
}

func (m *MockProductWarehouseStockRepository) GetByProductAndWarehouseForUpdate(ctx context.Context, productID uint, warehouseID uint) (*models.ProductWarehouseStock, error) {
	if m.GetByProductAndWarehouseForUpdateFunc != nil {
		return m.GetByProductAndWarehouseForUpdateFunc(ctx, productID, warehouseID)
	}
	return nil, nil
}

func (m *MockProductWarehouseStockRepository) SumByProduct(ctx context.Context, productID uint) (int, int, error) {
	if m.SumByProductFunc != nil {
		return m.SumByProductFunc(ctx, productID)
	}
	return 0, 0, nil
}

func (m *MockProductWarehouseStockRepository) Upsert(ctx context.Context, s *models.ProductWarehouseStock) error {
	if m.UpsertFunc != nil {
		return m.UpsertFunc(ctx, s)
	}
	return nil
}

func (m *MockProductWarehouseStockRepository) AdjustQuantity(ctx context.Context, productID, warehouseID uint, delta int) error {
	if m.AdjustQuantityFunc != nil {
		return m.AdjustQuantityFunc(ctx, productID, warehouseID, delta)
	}
	return nil
}

// MockWarehouseRepository
type MockWarehouseRepository struct {
	CreateFunc func(ctx context.Context, entity *models.Warehouse) error
	GetByIDFunc func(ctx context.Context, id uint) (*models.Warehouse, error)
	GetAllFunc func(ctx context.Context) ([]models.Warehouse, error)
	UpdateFunc func(ctx context.Context, entity *models.Warehouse) error
	DeleteFunc func(ctx context.Context, id uint) error
	CountFunc func(ctx context.Context) (int64, error)
	GetDefaultFunc func(ctx context.Context) (*models.Warehouse, error)
	SetDefaultFunc func(ctx context.Context, id uint) error
}

func (m *MockWarehouseRepository) Create(ctx context.Context, entity *models.Warehouse) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, entity)
	}
	return nil
}

func (m *MockWarehouseRepository) GetByID(ctx context.Context, id uint) (*models.Warehouse, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockWarehouseRepository) GetAll(ctx context.Context) ([]models.Warehouse, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc(ctx)
	}
	return nil, nil
}

func (m *MockWarehouseRepository) Update(ctx context.Context, entity *models.Warehouse) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, entity)
	}
	return nil
}

func (m *MockWarehouseRepository) Delete(ctx context.Context, id uint) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return nil
}

func (m *MockWarehouseRepository) Count(ctx context.Context) (int64, error) {
	if m.CountFunc != nil {
		return m.CountFunc(ctx)
	}
	return 0, nil
}

func (m *MockWarehouseRepository) GetDefault(ctx context.Context) (*models.Warehouse, error) {
	if m.GetDefaultFunc != nil {
		return m.GetDefaultFunc(ctx)
	}
	return nil, nil
}

func (m *MockWarehouseRepository) SetDefault(ctx context.Context, id uint) error {
	if m.SetDefaultFunc != nil {
		return m.SetDefaultFunc(ctx, id)
	}
	return nil
}

// MockBlockedIPRepository
type MockBlockedIPRepository struct {
	FindActiveByIPFunc func(ctx context.Context, ip string) (*models.BlockedIP, error)
	CreateOrReactivateFunc func(ctx context.Context, ip, reason, blockedBy string, expiresAt *time.Time) error
	DeactivateFunc func(ctx context.Context, ip string) error
	ListActiveFunc func(ctx context.Context, page, limit int) ([]models.BlockedIP, int64, error)
}

func (m *MockBlockedIPRepository) FindActiveByIP(ctx context.Context, ip string) (*models.BlockedIP, error) {
	if m.FindActiveByIPFunc != nil {
		return m.FindActiveByIPFunc(ctx, ip)
	}
	return nil, nil
}

func (m *MockBlockedIPRepository) CreateOrReactivate(ctx context.Context, ip, reason, blockedBy string, expiresAt *time.Time) error {
	if m.CreateOrReactivateFunc != nil {
		return m.CreateOrReactivateFunc(ctx, ip, reason, blockedBy, expiresAt)
	}
	return nil
}

func (m *MockBlockedIPRepository) Deactivate(ctx context.Context, ip string) error {
	if m.DeactivateFunc != nil {
		return m.DeactivateFunc(ctx, ip)
	}
	return nil
}

func (m *MockBlockedIPRepository) ListActive(ctx context.Context, page, limit int) ([]models.BlockedIP, int64, error) {
	if m.ListActiveFunc != nil {
		return m.ListActiveFunc(ctx, page, limit)
	}
	return nil, 0, nil
}

// MockTrafficLogRepository
type MockTrafficLogRepository struct {
	CreateFunc func(ctx context.Context, log *models.TrafficLog) error
	BulkCreateFunc func(ctx context.Context, logs []models.TrafficLog) error
	ListFunc func(ctx context.Context, filter repository.TrafficLogFilter, page, limit int) ([]models.TrafficLog, int64, error)
	ListByIPFunc func(ctx context.Context, ip string, page, limit int) ([]models.TrafficLog, int64, error)
	CountSinceFunc func(ctx context.Context, from time.Time) (int64, error)
	CountSuspiciousSinceFunc func(ctx context.Context, from time.Time) (int64, error)
	HourlyCountTodayFunc func(ctx context.Context) ([]models.HourlyTrafficData, error)
	ListLatestByUserAgentContainsFunc func(ctx context.Context, patterns []string, page, limit int) ([]models.TrafficLog, int64, error)
}

func (m *MockTrafficLogRepository) Create(ctx context.Context, log *models.TrafficLog) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, log)
	}
	return nil
}

func (m *MockTrafficLogRepository) BulkCreate(ctx context.Context, logs []models.TrafficLog) error {
	if m.BulkCreateFunc != nil {
		return m.BulkCreateFunc(ctx, logs)
	}
	return nil
}

func (m *MockTrafficLogRepository) List(ctx context.Context, filter repository.TrafficLogFilter, page, limit int) ([]models.TrafficLog, int64, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx, filter, page, limit)
	}
	return nil, 0, nil
}

func (m *MockTrafficLogRepository) ListByIP(ctx context.Context, ip string, page, limit int) ([]models.TrafficLog, int64, error) {
	if m.ListByIPFunc != nil {
		return m.ListByIPFunc(ctx, ip, page, limit)
	}
	return nil, 0, nil
}

func (m *MockTrafficLogRepository) CountSince(ctx context.Context, from time.Time) (int64, error) {
	if m.CountSinceFunc != nil {
		return m.CountSinceFunc(ctx, from)
	}
	return 0, nil
}

func (m *MockTrafficLogRepository) CountSuspiciousSince(ctx context.Context, from time.Time) (int64, error) {
	if m.CountSuspiciousSinceFunc != nil {
		return m.CountSuspiciousSinceFunc(ctx, from)
	}
	return 0, nil
}

func (m *MockTrafficLogRepository) HourlyCountToday(ctx context.Context) ([]models.HourlyTrafficData, error) {
	if m.HourlyCountTodayFunc != nil {
		return m.HourlyCountTodayFunc(ctx)
	}
	return nil, nil
}

func (m *MockTrafficLogRepository) ListLatestByUserAgentContains(ctx context.Context, patterns []string, page, limit int) ([]models.TrafficLog, int64, error) {
	if m.ListLatestByUserAgentContainsFunc != nil {
		return m.ListLatestByUserAgentContainsFunc(ctx, patterns, page, limit)
	}
	return nil, 0, nil
}

// MockTrafficSessionRepository
type MockTrafficSessionRepository struct {
	CountOnlineFunc func(ctx context.Context, since time.Time) (int64, error)
	ListOnlineFunc func(ctx context.Context, since time.Time, page, limit int) ([]models.TrafficSession, int64, error)
	DeactivateByIPFunc func(ctx context.Context, ip string) error
	CleanupExpiredFunc func(ctx context.Context, before time.Time) error
	GetBySessionIDFunc func(ctx context.Context, sessionID string) (*models.TrafficSession, error)
	CreateFunc func(ctx context.Context, sess *models.TrafficSession) error
	SaveFunc func(ctx context.Context, sess *models.TrafficSession) error
}

func (m *MockTrafficSessionRepository) CountOnline(ctx context.Context, since time.Time) (int64, error) {
	if m.CountOnlineFunc != nil {
		return m.CountOnlineFunc(ctx, since)
	}
	return 0, nil
}

func (m *MockTrafficSessionRepository) ListOnline(ctx context.Context, since time.Time, page, limit int) ([]models.TrafficSession, int64, error) {
	if m.ListOnlineFunc != nil {
		return m.ListOnlineFunc(ctx, since, page, limit)
	}
	return nil, 0, nil
}

func (m *MockTrafficSessionRepository) DeactivateByIP(ctx context.Context, ip string) error {
	if m.DeactivateByIPFunc != nil {
		return m.DeactivateByIPFunc(ctx, ip)
	}
	return nil
}

func (m *MockTrafficSessionRepository) CleanupExpired(ctx context.Context, before time.Time) error {
	if m.CleanupExpiredFunc != nil {
		return m.CleanupExpiredFunc(ctx, before)
	}
	return nil
}

func (m *MockTrafficSessionRepository) GetBySessionID(ctx context.Context, sessionID string) (*models.TrafficSession, error) {
	if m.GetBySessionIDFunc != nil {
		return m.GetBySessionIDFunc(ctx, sessionID)
	}
	return nil, nil
}

func (m *MockTrafficSessionRepository) Create(ctx context.Context, sess *models.TrafficSession) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, sess)
	}
	return nil
}

func (m *MockTrafficSessionRepository) Save(ctx context.Context, sess *models.TrafficSession) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(ctx, sess)
	}
	return nil
}

// MockCategoryRepository
type MockCategoryRepository struct {
	CreateFunc func(ctx context.Context, entity *models.Category) error
	GetByIDFunc func(ctx context.Context, id uint) (*models.Category, error)
	GetAllFunc func(ctx context.Context) ([]models.Category, error)
	UpdateFunc func(ctx context.Context, entity *models.Category) error
	DeleteFunc func(ctx context.Context, id uint) error
	CountFunc func(ctx context.Context) (int64, error)
	GetBySlugFunc func(ctx context.Context, slug string) (*models.Category, error)
	GetByParentIDFunc func(ctx context.Context, parentID *uint) ([]models.Category, error)
	GetRootCategoriesFunc func(ctx context.Context) ([]models.Category, error)
	GetPublishedFunc func(ctx context.Context) ([]models.Category, error)
	GetByHomePageFunc func(ctx context.Context) ([]models.Category, error)
	GetByMenuFunc func(ctx context.Context) ([]models.Category, error)
	GetHierarchyFunc func(ctx context.Context) ([]models.Category, error)
	GetWithProductCountFunc func(ctx context.Context) ([]models.Category, error)
}

func (m *MockCategoryRepository) Create(ctx context.Context, entity *models.Category) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, entity)
	}
	return nil
}

func (m *MockCategoryRepository) GetByID(ctx context.Context, id uint) (*models.Category, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockCategoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc(ctx)
	}
	return nil, nil
}

func (m *MockCategoryRepository) Update(ctx context.Context, entity *models.Category) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, entity)
	}
	return nil
}

func (m *MockCategoryRepository) Delete(ctx context.Context, id uint) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return nil
}

func (m *MockCategoryRepository) Count(ctx context.Context) (int64, error) {
	if m.CountFunc != nil {
		return m.CountFunc(ctx)
	}
	return 0, nil
}

func (m *MockCategoryRepository) GetBySlug(ctx context.Context, slug string) (*models.Category, error) {
	if m.GetBySlugFunc != nil {
		return m.GetBySlugFunc(ctx, slug)
	}
	return nil, nil
}

func (m *MockCategoryRepository) GetByParentID(ctx context.Context, parentID *uint) ([]models.Category, error) {
	if m.GetByParentIDFunc != nil {
		return m.GetByParentIDFunc(ctx, parentID)
	}
	return nil, nil
}

func (m *MockCategoryRepository) GetRootCategories(ctx context.Context) ([]models.Category, error) {
	if m.GetRootCategoriesFunc != nil {
		return m.GetRootCategoriesFunc(ctx)
	}
	return nil, nil
}

func (m *MockCategoryRepository) GetPublished(ctx context.Context) ([]models.Category, error) {
	if m.GetPublishedFunc != nil {
		return m.GetPublishedFunc(ctx)
	}
	return nil, nil
}

func (m *MockCategoryRepository) GetByHomePage(ctx context.Context) ([]models.Category, error) {
	if m.GetByHomePageFunc != nil {
		return m.GetByHomePageFunc(ctx)
	}
	return nil, nil
}

func (m *MockCategoryRepository) GetByMenu(ctx context.Context) ([]models.Category, error) {
	if m.GetByMenuFunc != nil {
		return m.GetByMenuFunc(ctx)
	}
	return nil, nil
}

func (m *MockCategoryRepository) GetHierarchy(ctx context.Context) ([]models.Category, error) {
	if m.GetHierarchyFunc != nil {
		return m.GetHierarchyFunc(ctx)
	}
	return nil, nil
}

func (m *MockCategoryRepository) GetWithProductCount(ctx context.Context) ([]models.Category, error) {
	if m.GetWithProductCountFunc != nil {
		return m.GetWithProductCountFunc(ctx)
	}
	return nil, nil
}

// MockFraudRuleRepository
type MockFraudRuleRepository struct {
	ListActiveFunc func(ctx context.Context) ([]repository.FraudRule, error)
	GetByKeyFunc func(ctx context.Context, key string) (*repository.FraudRule, error)
	GetByIDFunc func(ctx context.Context, id uint) (*repository.FraudRule, error)
	ListAllFunc func(ctx context.Context) ([]repository.FraudRule, error)
	UpdateFunc func(ctx context.Context, r *repository.FraudRule) error
}

func (m *MockFraudRuleRepository) ListActive(ctx context.Context) ([]repository.FraudRule, error) {
	if m.ListActiveFunc != nil {
		return m.ListActiveFunc(ctx)
	}
	return nil, nil
}

func (m *MockFraudRuleRepository) GetByKey(ctx context.Context, key string) (*repository.FraudRule, error) {
	if m.GetByKeyFunc != nil {
		return m.GetByKeyFunc(ctx, key)
	}
	return nil, nil
}

func (m *MockFraudRuleRepository) GetByID(ctx context.Context, id uint) (*repository.FraudRule, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockFraudRuleRepository) ListAll(ctx context.Context) ([]repository.FraudRule, error) {
	if m.ListAllFunc != nil {
		return m.ListAllFunc(ctx)
	}
	return nil, nil
}

func (m *MockFraudRuleRepository) Update(ctx context.Context, r *repository.FraudRule) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, r)
	}
	return nil
}

// MockFraudCaseRepository
type MockFraudCaseRepository struct {
	CreateFunc func(ctx context.Context, c *repository.FraudCase) error
	UpdateFunc func(ctx context.Context, c *repository.FraudCase) error
	GetByIDFunc func(ctx context.Context, id uint) (*repository.FraudCase, error)
	GetByOrderIDFunc func(ctx context.Context, orderID uint) (*repository.FraudCase, error)
	ListFunc func(ctx context.Context, f repository.FraudCaseFilter) ([]repository.FraudCase, int64, error)
	ListEnrichedFunc func(ctx context.Context, f repository.FraudCaseFilter) ([]repository.AdminFraudCaseRow, int64, error)
	GetDetailFunc func(ctx context.Context, id uint) (*repository.FraudCaseDetail, error)
}

func (m *MockFraudCaseRepository) Create(ctx context.Context, c *repository.FraudCase) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, c)
	}
	return nil
}

func (m *MockFraudCaseRepository) Update(ctx context.Context, c *repository.FraudCase) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, c)
	}
	return nil
}

func (m *MockFraudCaseRepository) GetByID(ctx context.Context, id uint) (*repository.FraudCase, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockFraudCaseRepository) GetByOrderID(ctx context.Context, orderID uint) (*repository.FraudCase, error) {
	if m.GetByOrderIDFunc != nil {
		return m.GetByOrderIDFunc(ctx, orderID)
	}
	return nil, nil
}

func (m *MockFraudCaseRepository) List(ctx context.Context, f repository.FraudCaseFilter) ([]repository.FraudCase, int64, error) {
	if m.ListFunc != nil {
		return m.ListFunc(ctx, f)
	}
	return nil, 0, nil
}

func (m *MockFraudCaseRepository) ListEnriched(ctx context.Context, f repository.FraudCaseFilter) ([]repository.AdminFraudCaseRow, int64, error) {
	if m.ListEnrichedFunc != nil {
		return m.ListEnrichedFunc(ctx, f)
	}
	return nil, 0, nil
}

func (m *MockFraudCaseRepository) GetDetail(ctx context.Context, id uint) (*repository.FraudCaseDetail, error) {
	if m.GetDetailFunc != nil {
		return m.GetDetailFunc(ctx, id)
	}
	return nil, nil
}

// MockFraudEventRepository
type MockFraudEventRepository struct {
	CreateManyFunc func(ctx context.Context, events []repository.FraudEvent) error
	ListByCaseFunc func(ctx context.Context, caseID uint) ([]repository.FraudEvent, error)
}

func (m *MockFraudEventRepository) CreateMany(ctx context.Context, events []repository.FraudEvent) error {
	if m.CreateManyFunc != nil {
		return m.CreateManyFunc(ctx, events)
	}
	return nil
}

func (m *MockFraudEventRepository) ListByCase(ctx context.Context, caseID uint) ([]repository.FraudEvent, error) {
	if m.ListByCaseFunc != nil {
		return m.ListByCaseFunc(ctx, caseID)
	}
	return nil, nil
}

// MockProductRepository
type MockProductRepository struct {
	CreateFunc func(ctx context.Context, p *models.Product) error
	GetByIDFunc func(ctx context.Context, id uint) (*models.Product, error)
	GetBySlugFunc func(ctx context.Context, slug string) (*models.Product, error)
	GetBySKUFunc func(ctx context.Context, sku string) (*models.Product, error)
	GetAllFunc func(ctx context.Context) ([]models.Product, error)
	GetByCategoryFunc func(ctx context.Context, categoryID uint, limit, offset int) ([]models.Product, error)
	GetByBrandFunc func(ctx context.Context, brandID uint, limit, offset int) ([]models.Product, error)
	GetPublishedFunc func(ctx context.Context, limit, offset int) ([]models.Product, error)
	GetPublishedAfterFunc func(ctx context.Context, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error)
	GetByCategoryAfterFunc func(ctx context.Context, categoryID uint, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error)
	GetByBrandAfterFunc func(ctx context.Context, brandID uint, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error)
	GetNewProductsFunc func(ctx context.Context, limit int) ([]models.Product, error)
	GetDiscountedProductsFunc func(ctx context.Context, limit int) ([]models.Product, error)
	SearchByNameFunc func(ctx context.Context, name string, limit, offset int) ([]models.Product, error)
	GetByPriceRangeFunc func(ctx context.Context, minPrice, maxPrice float64, limit, offset int) ([]models.Product, error)
	GetRelatedProductsFunc func(ctx context.Context, productID uint, limit int) ([]models.Product, error)
	UpdateStockFunc func(ctx context.Context, productID uint, quantity int) error
	GetLowStockProductsFunc func(ctx context.Context) ([]models.Product, error)
	UpdateFunc func(ctx context.Context, p *models.Product) error
	DeleteFunc func(ctx context.Context, id uint) error
	CountFunc func(ctx context.Context) (int64, error)
}

func (m *MockProductRepository) Create(ctx context.Context, p *models.Product) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, p)
	}
	return nil
}

func (m *MockProductRepository) GetByID(ctx context.Context, id uint) (*models.Product, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockProductRepository) GetBySlug(ctx context.Context, slug string) (*models.Product, error) {
	if m.GetBySlugFunc != nil {
		return m.GetBySlugFunc(ctx, slug)
	}
	return nil, nil
}

func (m *MockProductRepository) GetBySKU(ctx context.Context, sku string) (*models.Product, error) {
	if m.GetBySKUFunc != nil {
		return m.GetBySKUFunc(ctx, sku)
	}
	return nil, nil
}

func (m *MockProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc(ctx)
	}
	return nil, nil
}

func (m *MockProductRepository) GetByCategory(ctx context.Context, categoryID uint, limit, offset int) ([]models.Product, error) {
	if m.GetByCategoryFunc != nil {
		return m.GetByCategoryFunc(ctx, categoryID, limit, offset)
	}
	return nil, nil
}

func (m *MockProductRepository) GetByBrand(ctx context.Context, brandID uint, limit, offset int) ([]models.Product, error) {
	if m.GetByBrandFunc != nil {
		return m.GetByBrandFunc(ctx, brandID, limit, offset)
	}
	return nil, nil
}

func (m *MockProductRepository) GetPublished(ctx context.Context, limit, offset int) ([]models.Product, error) {
	if m.GetPublishedFunc != nil {
		return m.GetPublishedFunc(ctx, limit, offset)
	}
	return nil, nil
}

func (m *MockProductRepository) GetPublishedAfter(ctx context.Context, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error) {
	if m.GetPublishedAfterFunc != nil {
		return m.GetPublishedAfterFunc(ctx, limit, cursorCreatedAt, cursorID)
	}
	return nil, nil
}

func (m *MockProductRepository) GetByCategoryAfter(ctx context.Context, categoryID uint, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error) {
	if m.GetByCategoryAfterFunc != nil {
		return m.GetByCategoryAfterFunc(ctx, categoryID, limit, cursorCreatedAt, cursorID)
	}
	return nil, nil
}

func (m *MockProductRepository) GetByBrandAfter(ctx context.Context, brandID uint, limit int, cursorCreatedAt time.Time, cursorID uint) ([]models.Product, error) {
	if m.GetByBrandAfterFunc != nil {
		return m.GetByBrandAfterFunc(ctx, brandID, limit, cursorCreatedAt, cursorID)
	}
	return nil, nil
}

func (m *MockProductRepository) GetNewProducts(ctx context.Context, limit int) ([]models.Product, error) {
	if m.GetNewProductsFunc != nil {
		return m.GetNewProductsFunc(ctx, limit)
	}
	return nil, nil
}

func (m *MockProductRepository) GetDiscountedProducts(ctx context.Context, limit int) ([]models.Product, error) {
	if m.GetDiscountedProductsFunc != nil {
		return m.GetDiscountedProductsFunc(ctx, limit)
	}
	return nil, nil
}

func (m *MockProductRepository) SearchByName(ctx context.Context, name string, limit, offset int) ([]models.Product, error) {
	if m.SearchByNameFunc != nil {
		return m.SearchByNameFunc(ctx, name, limit, offset)
	}
	return nil, nil
}

func (m *MockProductRepository) GetByPriceRange(ctx context.Context, minPrice, maxPrice float64, limit, offset int) ([]models.Product, error) {
	if m.GetByPriceRangeFunc != nil {
		return m.GetByPriceRangeFunc(ctx, minPrice, maxPrice, limit, offset)
	}
	return nil, nil
}

func (m *MockProductRepository) GetRelatedProducts(ctx context.Context, productID uint, limit int) ([]models.Product, error) {
	if m.GetRelatedProductsFunc != nil {
		return m.GetRelatedProductsFunc(ctx, productID, limit)
	}
	return nil, nil
}

func (m *MockProductRepository) UpdateStock(ctx context.Context, productID uint, quantity int) error {
	if m.UpdateStockFunc != nil {
		return m.UpdateStockFunc(ctx, productID, quantity)
	}
	return nil
}

func (m *MockProductRepository) Update(ctx context.Context, p *models.Product) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, p)
	}
	return nil
}

func (m *MockProductRepository) Delete(ctx context.Context, id uint) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return nil
}

func (m *MockProductRepository) Count(ctx context.Context) (int64, error) {
	if m.CountFunc != nil {
		return m.CountFunc(ctx)
	}
	return 0, nil
}

func (m *MockProductRepository) GetLowStockProducts(ctx context.Context) ([]models.Product, error) {
	if m.GetLowStockProductsFunc != nil {
		return m.GetLowStockProductsFunc(ctx)
	}
	return nil, nil
}

// MockBrandRepository
type MockBrandRepository struct {
	CreateFunc func(ctx context.Context, b *models.Brand) error
	GetByIDFunc func(ctx context.Context, id uint) (*models.Brand, error)
	GetBySlugFunc func(ctx context.Context, slug string) (*models.Brand, error)
	GetAllFunc func(ctx context.Context) ([]models.Brand, error)
	UpdateFunc func(ctx context.Context, b *models.Brand) error
	DeleteFunc func(ctx context.Context, id uint) error
	CountFunc func(ctx context.Context) (int64, error)
	GetPublishedFunc func(ctx context.Context) ([]models.Brand, error)
	GetByHomePageFunc func(ctx context.Context) ([]models.Brand, error)
	GetByMenuFunc func(ctx context.Context) ([]models.Brand, error)
	SearchByNameFunc func(ctx context.Context, name string) ([]models.Brand, error)
}

func (m *MockBrandRepository) Create(ctx context.Context, b *models.Brand) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, b)
	}
	return nil
}

func (m *MockBrandRepository) GetByID(ctx context.Context, id uint) (*models.Brand, error) {
	if m.GetByIDFunc != nil {
		return m.GetByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockBrandRepository) GetBySlug(ctx context.Context, slug string) (*models.Brand, error) {
	if m.GetBySlugFunc != nil {
		return m.GetBySlugFunc(ctx, slug)
	}
	return nil, nil
}

func (m *MockBrandRepository) GetAll(ctx context.Context) ([]models.Brand, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc(ctx)
	}
	return nil, nil
}

func (m *MockBrandRepository) Update(ctx context.Context, b *models.Brand) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, b)
	}
	return nil
}

func (m *MockBrandRepository) Delete(ctx context.Context, id uint) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}
	return nil
}

func (m *MockBrandRepository) Count(ctx context.Context) (int64, error) {
	if m.CountFunc != nil {
		return m.CountFunc(ctx)
	}
	return 0, nil
}

func (m *MockBrandRepository) GetPublished(ctx context.Context) ([]models.Brand, error) {
	if m.GetPublishedFunc != nil {
		return m.GetPublishedFunc(ctx)
	}
	return nil, nil
}

func (m *MockBrandRepository) GetByHomePage(ctx context.Context) ([]models.Brand, error) {
	if m.GetByHomePageFunc != nil {
		return m.GetByHomePageFunc(ctx)
	}
	return nil, nil
}

func (m *MockBrandRepository) GetByMenu(ctx context.Context) ([]models.Brand, error) {
	if m.GetByMenuFunc != nil {
		return m.GetByMenuFunc(ctx)
	}
	return nil, nil
}

func (m *MockBrandRepository) SearchByName(ctx context.Context, name string) ([]models.Brand, error) {
	if m.SearchByNameFunc != nil {
		return m.SearchByNameFunc(ctx, name)
	}
	return nil, nil
}

// MockOrderRepository
type MockOrderRepository struct {
	CreateOrderFunc func(order *models.Order) error
	GetOrderByIDFunc func(orderID uint) (*models.Order, error)
	GetOrderByOrderNumberFunc func(orderNumber string) (*models.Order, error)
	GetUserOrdersFunc func(userID uint) ([]models.Order, error)
	UpdateOrderFunc func(order *models.Order) error
	UpdateOrderStatusFunc func(orderID uint, status string) error
	UpdatePaymentStatusFunc func(orderID uint, status string) error
	CancelOrderFunc func(orderID uint) error
	ReleaseExpiredReservationsFunc func() error
	GetPendingOrdersFunc func() ([]models.Order, error)
	GetExpiredReservationsFunc func() ([]models.InventoryReservation, error)
}

func (m *MockOrderRepository) CreateOrder(order *models.Order) error {
	if m.CreateOrderFunc != nil {
		return m.CreateOrderFunc(order)
	}
	return nil
}

func (m *MockOrderRepository) GetOrderByID(orderID uint) (*models.Order, error) {
	if m.GetOrderByIDFunc != nil {
		return m.GetOrderByIDFunc(orderID)
	}
	return nil, nil
}

func (m *MockOrderRepository) GetOrderByOrderNumber(orderNumber string) (*models.Order, error) {
	if m.GetOrderByOrderNumberFunc != nil {
		return m.GetOrderByOrderNumberFunc(orderNumber)
	}
	return nil, nil
}

func (m *MockOrderRepository) GetUserOrders(userID uint) ([]models.Order, error) {
	if m.GetUserOrdersFunc != nil {
		return m.GetUserOrdersFunc(userID)
	}
	return nil, nil
}

func (m *MockOrderRepository) UpdateOrder(order *models.Order) error {
	if m.UpdateOrderFunc != nil {
		return m.UpdateOrderFunc(order)
	}
	return nil
}

func (m *MockOrderRepository) UpdateOrderStatus(orderID uint, status string) error {
	if m.UpdateOrderStatusFunc != nil {
		return m.UpdateOrderStatusFunc(orderID, status)
	}
	return nil
}

func (m *MockOrderRepository) UpdatePaymentStatus(orderID uint, status string) error {
	if m.UpdatePaymentStatusFunc != nil {
		return m.UpdatePaymentStatusFunc(orderID, status)
	}
	return nil
}

func (m *MockOrderRepository) CancelOrder(orderID uint) error {
	if m.CancelOrderFunc != nil {
		return m.CancelOrderFunc(orderID)
	}
	return nil
}

func (m *MockOrderRepository) ReleaseExpiredReservations() error {
	if m.ReleaseExpiredReservationsFunc != nil {
		return m.ReleaseExpiredReservationsFunc()
	}
	return nil
}

func (m *MockOrderRepository) GetPendingOrders() ([]models.Order, error) {
	if m.GetPendingOrdersFunc != nil {
		return m.GetPendingOrdersFunc()
	}
	return nil, nil
}

func (m *MockOrderRepository) GetExpiredReservations() ([]models.InventoryReservation, error) {
	if m.GetExpiredReservationsFunc != nil {
		return m.GetExpiredReservationsFunc()
	}
	return nil, nil
}

// MockInventoryRepository
type MockInventoryRepository struct {
	GetInventoryByProductIDFunc func(productID uint) (*models.ProductInventory, error)
	CreateOrUpdateInventoryFunc func(inventory *models.ProductInventory) error
	CheckAvailabilityFunc func(productID uint, quantity int) (bool, error)
	ReserveInventoryFunc func(productID, orderID, quantity uint, reservedUntil time.Time) (*models.InventoryReservation, error)
	ReleaseReservationFunc func(reservationID uint) error
	ReleaseExpiredReservationsFunc func() error
	GetLowStockProductsFunc func() ([]models.ProductInventory, error)
	GetOutOfStockProductsFunc func() ([]models.ProductInventory, error)
	GetInventorySettingsFunc func() (*models.InventorySettings, error)
	UpdateInventorySettingsFunc func(settings *models.InventorySettings) error
	CheckAndReserveInventoryFunc func(order *models.Order) error
	ProcessOrderReservationFunc func(orderID uint) error
	CancelOrderReservationFunc func(orderID uint) error
}

func (m *MockInventoryRepository) GetInventoryByProductID(productID uint) (*models.ProductInventory, error) {
	if m.GetInventoryByProductIDFunc != nil {
		return m.GetInventoryByProductIDFunc(productID)
	}
	return nil, nil
}

func (m *MockInventoryRepository) CreateOrUpdateInventory(inventory *models.ProductInventory) error {
	if m.CreateOrUpdateInventoryFunc != nil {
		return m.CreateOrUpdateInventoryFunc(inventory)
	}
	return nil
}

func (m *MockInventoryRepository) CheckAvailability(productID uint, quantity int) (bool, error) {
	if m.CheckAvailabilityFunc != nil {
		return m.CheckAvailabilityFunc(productID, quantity)
	}
	return false, nil
}

func (m *MockInventoryRepository) ReserveInventory(productID, orderID, quantity uint, reservedUntil time.Time) (*models.InventoryReservation, error) {
	if m.ReserveInventoryFunc != nil {
		return m.ReserveInventoryFunc(productID, orderID, quantity, reservedUntil)
	}
	return nil, nil
}

func (m *MockInventoryRepository) ReleaseReservation(reservationID uint) error {
	if m.ReleaseReservationFunc != nil {
		return m.ReleaseReservationFunc(reservationID)
	}
	return nil
}

func (m *MockInventoryRepository) ReleaseExpiredReservations() error {
	if m.ReleaseExpiredReservationsFunc != nil {
		return m.ReleaseExpiredReservationsFunc()
	}
	return nil
}

func (m *MockInventoryRepository) GetLowStockProducts() ([]models.ProductInventory, error) {
	if m.GetLowStockProductsFunc != nil {
		return m.GetLowStockProductsFunc()
	}
	return nil, nil
}

func (m *MockInventoryRepository) GetOutOfStockProducts() ([]models.ProductInventory, error) {
	if m.GetOutOfStockProductsFunc != nil {
		return m.GetOutOfStockProductsFunc()
	}
	return nil, nil
}

func (m *MockInventoryRepository) GetInventorySettings() (*models.InventorySettings, error) {
	if m.GetInventorySettingsFunc != nil {
		return m.GetInventorySettingsFunc()
	}
	return nil, nil
}

func (m *MockInventoryRepository) UpdateInventorySettings(settings *models.InventorySettings) error {
	if m.UpdateInventorySettingsFunc != nil {
		return m.UpdateInventorySettingsFunc(settings)
	}
	return nil
}

func (m *MockInventoryRepository) CheckAndReserveInventory(order *models.Order) error {
	if m.CheckAndReserveInventoryFunc != nil {
		return m.CheckAndReserveInventoryFunc(order)
	}
	return nil
}

func (m *MockInventoryRepository) ProcessOrderReservation(orderID uint) error {
	if m.ProcessOrderReservationFunc != nil {
		return m.ProcessOrderReservationFunc(orderID)
	}
	return nil
}

func (m *MockInventoryRepository) CancelOrderReservation(orderID uint) error {
	if m.CancelOrderReservationFunc != nil {
		return m.CancelOrderReservationFunc(orderID)
	}
	return nil
}

// MockMenuRepository
type MockMenuRepository struct {
	CreateMenuFunc func(ctx context.Context, menu *models.Menu) error
	GetMenuByIDFunc func(ctx context.Context, id uint) (*models.Menu, error)
	GetMenuBySlugFunc func(ctx context.Context, slug string) (*models.Menu, error)
	GetAllMenusFunc func(ctx context.Context) ([]models.Menu, error)
	GetEnabledMenusFunc func(ctx context.Context) ([]models.Menu, error)
	UpdateMenuFunc func(ctx context.Context, menu *models.Menu) error
	DeleteMenuFunc func(ctx context.Context, id uint) error
	CreateMenuItemFunc func(ctx context.Context, item *models.MenuItem) error
	GetMenuItemByIDFunc func(ctx context.Context, id uint) (*models.MenuItem, error)
	GetMenuItemsByMenuIDFunc func(ctx context.Context, menuID uint) ([]models.MenuItem, error)
	UpdateMenuItemFunc func(ctx context.Context, item *models.MenuItem) error
	DeleteMenuItemFunc func(ctx context.Context, id uint) error
	ReorderMenuItemsFunc func(ctx context.Context, menuID uint, itemOrders []struct {
		ID    uint `json:"id"`
		Order int  `json:"order"`
	}) error
	CreateMenuColumnFunc func(ctx context.Context, column *models.MenuColumn) error
	GetMenuColumnByIDFunc func(ctx context.Context, id uint) (*models.MenuColumn, error)
	GetMenuColumnsByMenuIDFunc func(ctx context.Context, menuID uint) ([]models.MenuColumn, error)
	UpdateMenuColumnFunc func(ctx context.Context, column *models.MenuColumn) error
	DeleteMenuColumnFunc func(ctx context.Context, id uint) error
	ReorderMenuColumnsFunc func(ctx context.Context, menuID uint, columnOrders []struct {
		ID    uint `json:"id"`
		Order int  `json:"order"`
	}) error
	CreateMenuLocationFunc func(ctx context.Context, location *models.MenuLocation) error
	GetMenuLocationByIDFunc func(ctx context.Context, id uint) (*models.MenuLocation, error)
	GetMenuLocationBySlugFunc func(ctx context.Context, slug string) (*models.MenuLocation, error)
	GetAllMenuLocationsFunc func(ctx context.Context) ([]models.MenuLocation, error)
	UpdateMenuLocationFunc func(ctx context.Context, location *models.MenuLocation) error
	DeleteMenuLocationFunc func(ctx context.Context, id uint) error
	CreateMenuAssignmentFunc func(ctx context.Context, assignment *models.MenuAssignment) error
	GetMenuAssignmentByIDFunc func(ctx context.Context, id uint) (*models.MenuAssignment, error)
	GetMenuAssignmentsByLocationIDFunc func(ctx context.Context, locationID uint) ([]models.MenuAssignment, error)
	GetMenuAssignmentsByMenuIDFunc func(ctx context.Context, menuID uint) ([]models.MenuAssignment, error)
	UpdateMenuAssignmentFunc func(ctx context.Context, assignment *models.MenuAssignment) error
	DeleteMenuAssignmentFunc func(ctx context.Context, id uint) error
	AssignMenuToLocationFunc func(ctx context.Context, menuID, locationID uint, order int) error
	UnassignMenuFromLocationFunc func(ctx context.Context, menuID, locationID uint) error
	GetMenusByLocationFunc func(ctx context.Context, locationSlug string) ([]models.Menu, error)
	GetMenuContentPagesFunc func(ctx context.Context) ([]struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Slug  string `json:"slug"`
	}, error)
	GetMenuContentPostsFunc func(ctx context.Context) ([]struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Slug  string `json:"slug"`
	}, error)
	GetMenuContentCategoriesFunc func(ctx context.Context) ([]struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}, error)
	GetMenuContentProductCategoriesFunc func(ctx context.Context) ([]struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}, error)
}

func (m *MockMenuRepository) CreateMenu(ctx context.Context, menu *models.Menu) error {
	if m.CreateMenuFunc != nil {
		return m.CreateMenuFunc(ctx, menu)
	}
	return nil
}

func (m *MockMenuRepository) GetMenuByID(ctx context.Context, id uint) (*models.Menu, error) {
	if m.GetMenuByIDFunc != nil {
		return m.GetMenuByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetMenuBySlug(ctx context.Context, slug string) (*models.Menu, error) {
	if m.GetMenuBySlugFunc != nil {
		return m.GetMenuBySlugFunc(ctx, slug)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetAllMenus(ctx context.Context) ([]models.Menu, error) {
	if m.GetAllMenusFunc != nil {
		return m.GetAllMenusFunc(ctx)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetEnabledMenus(ctx context.Context) ([]models.Menu, error) {
	if m.GetEnabledMenusFunc != nil {
		return m.GetEnabledMenusFunc(ctx)
	}
	return nil, nil
}

func (m *MockMenuRepository) UpdateMenu(ctx context.Context, menu *models.Menu) error {
	if m.UpdateMenuFunc != nil {
		return m.UpdateMenuFunc(ctx, menu)
	}
	return nil
}

func (m *MockMenuRepository) DeleteMenu(ctx context.Context, id uint) error {
	if m.DeleteMenuFunc != nil {
		return m.DeleteMenuFunc(ctx, id)
	}
	return nil
}

func (m *MockMenuRepository) CreateMenuItem(ctx context.Context, item *models.MenuItem) error {
	if m.CreateMenuItemFunc != nil {
		return m.CreateMenuItemFunc(ctx, item)
	}
	return nil
}

func (m *MockMenuRepository) GetMenuItemByID(ctx context.Context, id uint) (*models.MenuItem, error) {
	if m.GetMenuItemByIDFunc != nil {
		return m.GetMenuItemByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetMenuItemsByMenuID(ctx context.Context, menuID uint) ([]models.MenuItem, error) {
	if m.GetMenuItemsByMenuIDFunc != nil {
		return m.GetMenuItemsByMenuIDFunc(ctx, menuID)
	}
	return nil, nil
}

func (m *MockMenuRepository) UpdateMenuItem(ctx context.Context, item *models.MenuItem) error {
	if m.UpdateMenuItemFunc != nil {
		return m.UpdateMenuItemFunc(ctx, item)
	}
	return nil
}

func (m *MockMenuRepository) DeleteMenuItem(ctx context.Context, id uint) error {
	if m.DeleteMenuItemFunc != nil {
		return m.DeleteMenuItemFunc(ctx, id)
	}
	return nil
}

func (m *MockMenuRepository) ReorderMenuItems(ctx context.Context, menuID uint, itemOrders []struct {
	ID    uint `json:"id"`
	Order int  `json:"order"`
}) error {
	if m.ReorderMenuItemsFunc != nil {
		return m.ReorderMenuItemsFunc(ctx, menuID, itemOrders)
	}
	return nil
}

func (m *MockMenuRepository) CreateMenuColumn(ctx context.Context, column *models.MenuColumn) error {
	if m.CreateMenuColumnFunc != nil {
		return m.CreateMenuColumnFunc(ctx, column)
	}
	return nil
}

func (m *MockMenuRepository) GetMenuColumnByID(ctx context.Context, id uint) (*models.MenuColumn, error) {
	if m.GetMenuColumnByIDFunc != nil {
		return m.GetMenuColumnByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetMenuColumnsByMenuID(ctx context.Context, menuID uint) ([]models.MenuColumn, error) {
	if m.GetMenuColumnsByMenuIDFunc != nil {
		return m.GetMenuColumnsByMenuIDFunc(ctx, menuID)
	}
	return nil, nil
}

func (m *MockMenuRepository) UpdateMenuColumn(ctx context.Context, column *models.MenuColumn) error {
	if m.UpdateMenuColumnFunc != nil {
		return m.UpdateMenuColumnFunc(ctx, column)
	}
	return nil
}

func (m *MockMenuRepository) DeleteMenuColumn(ctx context.Context, id uint) error {
	if m.DeleteMenuColumnFunc != nil {
		return m.DeleteMenuColumnFunc(ctx, id)
	}
	return nil
}

func (m *MockMenuRepository) ReorderMenuColumns(ctx context.Context, menuID uint, columnOrders []struct {
	ID    uint `json:"id"`
	Order int  `json:"order"`
}) error {
	if m.ReorderMenuColumnsFunc != nil {
		return m.ReorderMenuColumnsFunc(ctx, menuID, columnOrders)
	}
	return nil
}

func (m *MockMenuRepository) CreateMenuLocation(ctx context.Context, location *models.MenuLocation) error {
	if m.CreateMenuLocationFunc != nil {
		return m.CreateMenuLocationFunc(ctx, location)
	}
	return nil
}

func (m *MockMenuRepository) GetMenuLocationByID(ctx context.Context, id uint) (*models.MenuLocation, error) {
	if m.GetMenuLocationByIDFunc != nil {
		return m.GetMenuLocationByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetMenuLocationBySlug(ctx context.Context, slug string) (*models.MenuLocation, error) {
	if m.GetMenuLocationBySlugFunc != nil {
		return m.GetMenuLocationBySlugFunc(ctx, slug)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetAllMenuLocations(ctx context.Context) ([]models.MenuLocation, error) {
	if m.GetAllMenuLocationsFunc != nil {
		return m.GetAllMenuLocationsFunc(ctx)
	}
	return nil, nil
}

func (m *MockMenuRepository) UpdateMenuLocation(ctx context.Context, location *models.MenuLocation) error {
	if m.UpdateMenuLocationFunc != nil {
		return m.UpdateMenuLocationFunc(ctx, location)
	}
	return nil
}

func (m *MockMenuRepository) DeleteMenuLocation(ctx context.Context, id uint) error {
	if m.DeleteMenuLocationFunc != nil {
		return m.DeleteMenuLocationFunc(ctx, id)
	}
	return nil
}

func (m *MockMenuRepository) CreateMenuAssignment(ctx context.Context, assignment *models.MenuAssignment) error {
	if m.CreateMenuAssignmentFunc != nil {
		return m.CreateMenuAssignmentFunc(ctx, assignment)
	}
	return nil
}

func (m *MockMenuRepository) GetMenuAssignmentByID(ctx context.Context, id uint) (*models.MenuAssignment, error) {
	if m.GetMenuAssignmentByIDFunc != nil {
		return m.GetMenuAssignmentByIDFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetMenuAssignmentsByLocationID(ctx context.Context, locationID uint) ([]models.MenuAssignment, error) {
	if m.GetMenuAssignmentsByLocationIDFunc != nil {
		return m.GetMenuAssignmentsByLocationIDFunc(ctx, locationID)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetMenuAssignmentsByMenuID(ctx context.Context, menuID uint) ([]models.MenuAssignment, error) {
	if m.GetMenuAssignmentsByMenuIDFunc != nil {
		return m.GetMenuAssignmentsByMenuIDFunc(ctx, menuID)
	}
	return nil, nil
}

func (m *MockMenuRepository) UpdateMenuAssignment(ctx context.Context, assignment *models.MenuAssignment) error {
	if m.UpdateMenuAssignmentFunc != nil {
		return m.UpdateMenuAssignmentFunc(ctx, assignment)
	}
	return nil
}

func (m *MockMenuRepository) DeleteMenuAssignment(ctx context.Context, id uint) error {
	if m.DeleteMenuAssignmentFunc != nil {
		return m.DeleteMenuAssignmentFunc(ctx, id)
	}
	return nil
}

func (m *MockMenuRepository) AssignMenuToLocation(ctx context.Context, menuID, locationID uint, order int) error {
	if m.AssignMenuToLocationFunc != nil {
		return m.AssignMenuToLocationFunc(ctx, menuID, locationID, order)
	}
	return nil
}

func (m *MockMenuRepository) UnassignMenuFromLocation(ctx context.Context, menuID, locationID uint) error {
	if m.UnassignMenuFromLocationFunc != nil {
		return m.UnassignMenuFromLocationFunc(ctx, menuID, locationID)
	}
	return nil
}

func (m *MockMenuRepository) GetMenusByLocation(ctx context.Context, locationSlug string) ([]models.Menu, error) {
	if m.GetMenusByLocationFunc != nil {
		return m.GetMenusByLocationFunc(ctx, locationSlug)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetMenuContentPages(ctx context.Context) ([]struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}, error) {
	if m.GetMenuContentPagesFunc != nil {
		return m.GetMenuContentPagesFunc(ctx)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetMenuContentPosts(ctx context.Context) ([]struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}, error) {
	if m.GetMenuContentPostsFunc != nil {
		return m.GetMenuContentPostsFunc(ctx)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetMenuContentCategories(ctx context.Context) ([]struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}, error) {
	if m.GetMenuContentCategoriesFunc != nil {
		return m.GetMenuContentCategoriesFunc(ctx)
	}
	return nil, nil
}

func (m *MockMenuRepository) GetMenuContentProductCategories(ctx context.Context) ([]struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}, error) {
	if m.GetMenuContentProductCategoriesFunc != nil {
		return m.GetMenuContentProductCategoriesFunc(ctx)
	}
	return nil, nil
}
