package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"

	"gorm.io/gorm"
)

// WalletService سرویس دامنهٔ کیف پول با استفاده از UnitOfWork و Repository
// تمام تراکنش‌ها باید از طریق UnitOfWork مدیریت شوند. هیچ منطق دیتابیسی در Handler قرار نمی‌گیرد.
type WalletService struct {
	uowFactory unitofwork.UnitOfWorkFactory
}

func NewWalletService(uowFactory unitofwork.UnitOfWorkFactory) *WalletService {
	return &WalletService{uowFactory: uowFactory}
}

// CreditWallet شارژ کیف پول کاربر
// amount باید > 0 باشد
func (s *WalletService) CreditWallet(ctx context.Context, adminUserID, userID uint, amount float64, method, gateway string, metadata map[string]any) error {
	if userID == 0 || amount <= 0 {
		return fmt.Errorf("ورودی نامعتبر: userID و amount")
	}

	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() { _ = uow.Rollback() }()

	walletRepo := uow.WalletRepository()
	txRepo := uow.WalletTransactionRepository()

	wallet, err := walletRepo.GetByUserID(ctx, userID)
	if err != nil {
		// اگر پیدا نشد، ایجاد می‌کنیم
		if errors.Is(err, gorm.ErrRecordNotFound) {
			wallet = &models.UserWallet{UserID: userID, Balance: 0, Currency: "IRR", Status: "active"}
			if err := walletRepo.Create(ctx, wallet); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	newBalance := wallet.Balance + amount
	if err := walletRepo.UpdateBalance(ctx, wallet.ID, newBalance); err != nil {
		return err
	}

	// تراکنش
	txm := &models.WalletTransaction{
		WalletID:  wallet.ID,
		Type:      "credit",
		Method:    method,
		Amount:    amount,
		Gateway:   gateway,
		Status:    "success",
		CreatedAt: time.Now(),
	}
	if err := txRepo.Create(ctx, txm); err != nil {
		return err
	}

	if err := uow.Commit(); err != nil {
		return err
	}
	return nil
}

// DebitWallet کسر از کیف پول کاربر
func (s *WalletService) DebitWallet(ctx context.Context, adminUserID, userID uint, amount float64, method, reason string, metadata map[string]any) error {
	if userID == 0 || amount <= 0 {
		return fmt.Errorf("ورودی نامعتبر: userID و amount")
	}

	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() { _ = uow.Rollback() }()

	walletRepo := uow.WalletRepository()
	txRepo := uow.WalletTransactionRepository()

	wallet, err := walletRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if wallet.Status == "blocked" {
		return fmt.Errorf("کیف پول مسدود است")
	}
	if wallet.Balance < amount {
		return fmt.Errorf("موجودی کافی نیست")
	}

	newBalance := wallet.Balance - amount
	if err := walletRepo.UpdateBalance(ctx, wallet.ID, newBalance); err != nil {
		return err
	}

	txm := &models.WalletTransaction{
		WalletID:  wallet.ID,
		Type:      "debit",
		Method:    method,
		Amount:    amount,
		Status:    "success",
		CreatedAt: time.Now(),
	}
	if err := txRepo.Create(ctx, txm); err != nil {
		return err
	}

	if err := uow.Commit(); err != nil {
		return err
	}
	return nil
}

// Transfer انتقال بین دو کیف پول کاربر
func (s *WalletService) Transfer(ctx context.Context, adminUserID, fromUserID, toUserID uint, amount float64) error {
	if fromUserID == 0 || toUserID == 0 || amount <= 0 {
		return fmt.Errorf("ورودی نامعتبر برای انتقال")
	}
	if fromUserID == toUserID {
		return fmt.Errorf("انتقال بین دو شناسه یکسان مجاز نیست")
	}

	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() { _ = uow.Rollback() }()

	walletRepo := uow.WalletRepository()
	txRepo := uow.WalletTransactionRepository()

	fromWallet, err := walletRepo.GetByUserID(ctx, fromUserID)
	if err != nil {
		return err
	}
	if fromWallet.Status == "blocked" {
		return fmt.Errorf("کیف پول مبدا مسدود است")
	}
	if fromWallet.Balance < amount {
		return fmt.Errorf("موجودی مبدا کافی نیست")
	}

	toWallet, err := walletRepo.GetByUserID(ctx, toUserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			toWallet = &models.UserWallet{UserID: toUserID, Balance: 0, Currency: "IRR", Status: "active"}
			if err := walletRepo.Create(ctx, toWallet); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// بروزرسانی بالانس‌ها
	if err := walletRepo.UpdateBalance(ctx, fromWallet.ID, fromWallet.Balance-amount); err != nil {
		return err
	}
	if err := walletRepo.UpdateBalance(ctx, toWallet.ID, toWallet.Balance+amount); err != nil {
		return err
	}

	// ثبت تراکنش‌ها
	if err := txRepo.Create(ctx, &models.WalletTransaction{WalletID: fromWallet.ID, Type: "debit", Amount: amount, Status: "success", CreatedAt: time.Now()}); err != nil {
		return err
	}
	if err := txRepo.Create(ctx, &models.WalletTransaction{WalletID: toWallet.ID, Type: "credit", Amount: amount, Status: "success", CreatedAt: time.Now()}); err != nil {
		return err
	}

	if err := uow.Commit(); err != nil {
		return err
	}
	return nil
}

// BlockWallet مسدودسازی کیف پول
func (s *WalletService) BlockWallet(ctx context.Context, adminUserID, userID uint) error {
	if userID == 0 {
		return fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() { _ = uow.Rollback() }()
	if err := uow.WalletRepository().UpdateStatus(ctx, s.mustGetWalletID(ctx, uow, userID), "blocked"); err != nil {
		return err
	}
	return uow.Commit()
}

// UnblockWallet آزادسازی کیف پول
func (s *WalletService) UnblockWallet(ctx context.Context, adminUserID, userID uint) error {
	if userID == 0 {
		return fmt.Errorf("شناسه کاربر نامعتبر است")
	}
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() { _ = uow.Rollback() }()
	if err := uow.WalletRepository().UpdateStatus(ctx, s.mustGetWalletID(ctx, uow, userID), "active"); err != nil {
		return err
	}
	return uow.Commit()
}

// ListTransactionsAdmin فهرست تراکنش‌های سراسری برای ادمین
func (s *WalletService) ListTransactionsAdmin(ctx context.Context, page, pageSize int, filter repository.AdminTransactionFilter) ([]repository.AdminTransactionRow, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	uow := s.uowFactory.Create()
	rows, total, err := uow.WalletTransactionRepository().AdminListTransactions(ctx, pageSize, offset, filter)
	return rows, total, err
}

// Summary آمار خلاصه برای داشبورد کیف پول
func (s *WalletService) Summary(ctx context.Context) (map[string]any, error) {
	uow := s.uowFactory.Create()
	total, err := uow.WalletRepository().SumAllBalances(ctx)
	if err != nil {
		return nil, err
	}
	// تعداد کیف پول‌های مسدود
	var blockedCount int64
	if err := uow.GetTx().WithContext(ctx).Model(&models.UserWallet{}).Where("status = ?", "blocked").Count(&blockedCount).Error; err != nil {
		return nil, err
	}
	// تعداد تراکنش‌های امروز
	var txToday int64
	if err := uow.GetTx().WithContext(ctx).Model(&models.WalletTransaction{}).Where("created_at::date = CURRENT_DATE").Count(&txToday).Error; err != nil {
		return nil, err
	}
	// تعداد تراکنش‌های در انتظار (اگر استفاده شود)
	var pendingCount int64
	if err := uow.GetTx().WithContext(ctx).Model(&models.WalletTransaction{}).Where("status = ?", "pending").Count(&pendingCount).Error; err != nil {
		return nil, err
	}
	return map[string]any{
		"totalBalance": total,
		"blockedCount": blockedCount,
		"todayTxCount": txToday,
		"pendingCount": pendingCount,
	}, nil
}

// TrendLastNDays محاسبه روند خالص تراکنش‌ها در N روز اخیر (credit - debit)
func (s *WalletService) TrendLastNDays(ctx context.Context, days int) ([]struct {
	Day string
	Net int64
}, error) {
	if days <= 0 || days > 60 {
		days = 30
	}
	uow := s.uowFactory.Create()
	type row struct {
		Day string
		Net int64
	}
	var rows []row
	err := uow.GetTx().WithContext(ctx).
		Table("wallet_transactions").
		Select("to_char(created_at::date, 'YYYY-MM-DD') as day, COALESCE(SUM(CASE WHEN type='credit' THEN amount ELSE -amount END),0) as net").
		Where("created_at::date >= CURRENT_DATE - (?-1)", days).
		Group("created_at::date").
		Order("created_at::date").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	dayIndex := map[string]int64{}
	for _, r := range rows {
		dayIndex[r.Day] = r.Net
	}
	result := make([]struct {
		Day string
		Net int64
	}, 0, days)
	for i := days - 1; i >= 0; i-- {
		d := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		result = append(result, struct {
			Day string
			Net int64
		}{Day: d, Net: dayIndex[d]})
	}
	return result, nil
}

// UpdateTransactionStatus تغییر وضعیت تراکنش کیف پول (برای عملیات ادمین مانند تأیید/رد برداشت)
func (s *WalletService) UpdateTransactionStatus(ctx context.Context, txID uint, status string) error {
	if txID == 0 || status == "" {
		return fmt.Errorf("invalid input")
	}
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() { _ = uow.Rollback() }()
	if err := uow.WalletTransactionRepository().UpdateStatus(ctx, txID, status); err != nil {
		return err
	}
	return uow.Commit()
}

// mustGetWalletID گرفتن یا ساختن کیف پول و بازگرداندن ID آن
func (s *WalletService) mustGetWalletID(ctx context.Context, uow unitofwork.UnitOfWork, userID uint) uint {
	w, err := uow.WalletRepository().GetByUserID(ctx, userID)
	if err == nil {
		return w.ID
	}
	// اگر نبود ایجاد کن
	_ = uow.WalletRepository().Create(ctx, &models.UserWallet{UserID: userID, Balance: 0, Currency: "IRR", Status: "active"})
	w2, _ := uow.WalletRepository().GetByUserID(ctx, userID)
	if w2 != nil {
		return w2.ID
	}
	return 0
}
