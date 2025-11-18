package repository

import (
	"context"
	"fmt"

	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// WalletRepositoryInterface اینترفیس عملیات دیتابیسی کیف پول کاربران
// توجه: هیچ منطق دامنه‌ای اینجا قرار نمی‌گیرد؛ فقط کوئری‌ها و CRUD. تمام تراکنش‌ها باید توسط UnitOfWork مدیریت شوند.
type WalletRepositoryInterface interface {
	// GetByUserID دریافت کیف پول بر اساس شناسه کاربر
	GetByUserID(ctx context.Context, userID uint) (*models.UserWallet, error)

	// Create ایجاد رکورد کیف پول
	Create(ctx context.Context, wallet *models.UserWallet) error

	// Update به‌روزرسانی کامل رکورد (برای وضعیت و زمان‌ها)
	Update(ctx context.Context, wallet *models.UserWallet) error

	// UpdateBalance به‌روزرسانی بالانس به صورت امن
	UpdateBalance(ctx context.Context, walletID uint, newBalance float64) error

	// UpdateStatus تغییر وضعیت کیف پول
	UpdateStatus(ctx context.Context, walletID uint, status string) error

	// SumAllBalances جمع کل موجودی‌ها (برای داشبورد)
	SumAllBalances(ctx context.Context) (float64, error)
}

// WalletTransactionRepositoryInterface اینترفیس تراکنش‌های کیف پول
type WalletTransactionRepositoryInterface interface {
	// Create ایجاد تراکنش جدید
	Create(ctx context.Context, tx *models.WalletTransaction) error

	// ListByWallet فهرست تراکنش‌ها بر اساس کیف پول با فیلتر و صفحه‌بندی
	ListByWallet(ctx context.Context, walletID uint, limit int, offset int, filter WalletTransactionFilter) ([]models.WalletTransaction, int64, error)

	// AdminListTransactions فهرست سراسری تراکنش‌ها (برای صفحه ادمین) با فیلتر و صفحه‌بندی
	AdminListTransactions(ctx context.Context, limit int, offset int, filter AdminTransactionFilter) ([]AdminTransactionRow, int64, error)

	// GetByID دریافت تراکنش بر اساس ID
	GetByID(ctx context.Context, id uint) (*models.WalletTransaction, error)

	// UpdateStatus تغییر وضعیت تراکنش
	UpdateStatus(ctx context.Context, id uint, status string) error
}

// WalletTransactionFilter فیلتر تراکنش‌ها برای نمایش کیف پول کاربر
type WalletTransactionFilter struct {
	Type   string // credit|debit|""
	Status string // pending|success|failed|""
}

// AdminTransactionFilter فیلتر تراکنش‌ها برای لیست ادمین
type AdminTransactionFilter struct {
	Type     string // credit|debit|""
	Status   string // pending|success|failed|""
	Query    string // شماره تراکنش یا ایمیل/نام کاربر (در صورت موجود بودن)
	FromDate string // اختیاری: >= created_at
	ToDate   string // اختیاری: <= created_at
	UserID   *uint  // اختیاری: فیلتر بر اساس کاربر
	Method   string // online|refund|manual|withdraw|""
}

// AdminTransactionRow نمای سطری برای خروجی لیست ادمین
type AdminTransactionRow struct {
	ID        uint    `json:"id"`
	WalletID  uint    `json:"wallet_id"`
	UserID    uint    `json:"user_id"`
	Type      string  `json:"type"`
	Method    string  `json:"method"`
	Amount    float64 `json:"amount"`
	Gateway   string  `json:"gateway"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	// فیلدهای کاربر برای نمایش
	Username string `json:"username"`
	Email    string `json:"email"`
}

// WalletRepository پیاده‌سازی GORM ریپازیتوری کیف پول
type WalletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepositoryInterface {
	return &WalletRepository{DB: db}
}

func (r *WalletRepository) GetByUserID(ctx context.Context, userID uint) (*models.UserWallet, error) {
	if userID == 0 {
		return nil, fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	var wallet models.UserWallet
	err := r.DB.WithContext(ctx).Where("user_id = ?", userID).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) Create(ctx context.Context, wallet *models.UserWallet) error {
	if wallet == nil || wallet.UserID == 0 {
		return fmt.Errorf("رکورد کیف پول نامعتبر است")
	}
	return r.DB.WithContext(ctx).Create(wallet).Error
}

func (r *WalletRepository) Update(ctx context.Context, wallet *models.UserWallet) error {
	if wallet == nil || wallet.ID == 0 {
		return fmt.Errorf("رکورد کیف پول نامعتبر است")
	}
	return r.DB.WithContext(ctx).Save(wallet).Error
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, walletID uint, newBalance float64) error {
	if walletID == 0 {
		return fmt.Errorf("شناسه کیف پول نامعتبر است")
	}
	return r.DB.WithContext(ctx).Model(&models.UserWallet{}).Where("id = ?", walletID).Update("balance", newBalance).Error
}

func (r *WalletRepository) UpdateStatus(ctx context.Context, walletID uint, status string) error {
	if walletID == 0 {
		return fmt.Errorf("شناسه کیف پول نامعتبر است")
	}
	return r.DB.WithContext(ctx).Model(&models.UserWallet{}).Where("id = ?", walletID).Update("status", status).Error
}

func (r *WalletRepository) SumAllBalances(ctx context.Context) (float64, error) {
	type row struct{ Total float64 }
	var out row
	err := r.DB.WithContext(ctx).Table("user_wallets").Select("COALESCE(SUM(balance),0) as total").Scan(&out).Error
	if err != nil {
		return 0, err
	}
	return out.Total, nil
}

// WalletTransactionRepository پیاده‌سازی GORM تراکنش‌ها
type WalletTransactionRepository struct {
	DB *gorm.DB
}

func NewWalletTransactionRepository(db *gorm.DB) WalletTransactionRepositoryInterface {
	return &WalletTransactionRepository{DB: db}
}

func (r *WalletTransactionRepository) Create(ctx context.Context, txm *models.WalletTransaction) error {
	if txm == nil || txm.WalletID == 0 || txm.Amount == 0 || txm.Type == "" {
		return fmt.Errorf("تراکنش نامعتبر است")
	}
	return r.DB.WithContext(ctx).Create(txm).Error
}

func (r *WalletTransactionRepository) ListByWallet(ctx context.Context, walletID uint, limit int, offset int, filter WalletTransactionFilter) ([]models.WalletTransaction, int64, error) {
	var items []models.WalletTransaction
	var total int64

	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	q := r.DB.WithContext(ctx).Model(&models.WalletTransaction{}).Where("wallet_id = ?", walletID)
	if filter.Type != "" {
		q = q.Where("type = ?", filter.Type)
	}
	if filter.Status != "" {
		q = q.Where("status = ?", filter.Status)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order("created_at DESC").Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *WalletTransactionRepository) AdminListTransactions(ctx context.Context, limit int, offset int, filter AdminTransactionFilter) ([]AdminTransactionRow, int64, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	base := r.DB.WithContext(ctx).Table("wallet_transactions wt").
		Select("wt.id, wt.wallet_id, uw.user_id, wt.type, wt.method, wt.amount, wt.gateway, wt.status, to_char(wt.created_at, 'YYYY-MM-DD HH24:MI:SS') as created_at, u.username, u.email").
		Joins("JOIN user_wallets uw ON uw.id = wt.wallet_id").
		Joins("JOIN users u ON u.id = uw.user_id")

	if filter.Type != "" {
		base = base.Where("wt.type = ?", filter.Type)
	}
	if filter.Status != "" {
		base = base.Where("wt.status = ?", filter.Status)
	}
	if filter.Method != "" {
		base = base.Where("wt.method = ?", filter.Method)
	}
	if filter.UserID != nil {
		base = base.Where("uw.user_id = ?", *filter.UserID)
	}
	if filter.FromDate != "" {
		base = base.Where("wt.created_at >= ?", filter.FromDate)
	}
	if filter.ToDate != "" {
		base = base.Where("wt.created_at <= ?", filter.ToDate)
	}
	if filter.Query != "" {
		like := "%" + filter.Query + "%"
		base = base.Where("u.username ILIKE ? OR u.email ILIKE ? OR CAST(wt.id AS TEXT) ILIKE ?", like, like, like)
	}

	var total int64
	if err := base.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var rows []AdminTransactionRow
	if err := base.Order("wt.created_at DESC").Offset(offset).Limit(limit).Scan(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (r *WalletTransactionRepository) GetByID(ctx context.Context, id uint) (*models.WalletTransaction, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid id")
	}
	var txm models.WalletTransaction
	if err := r.DB.WithContext(ctx).First(&txm, id).Error; err != nil {
		return nil, err
	}
	return &txm, nil
}

func (r *WalletTransactionRepository) UpdateStatus(ctx context.Context, id uint, status string) error {
	if id == 0 {
		return fmt.Errorf("invalid id")
	}
	return r.DB.WithContext(ctx).Model(&models.WalletTransaction{}).Where("id = ?", id).Update("status", status).Error
}
