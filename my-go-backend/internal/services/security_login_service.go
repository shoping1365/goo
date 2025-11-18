package services

import (
	"context"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/repository"
)

// SecurityLoginService سرویس خواندن/نوشتن برای صفحه مانیتورینگ ورودها
// این سرویس از Unit of Work استفاده می‌کند و منطق تجاری را از لایه دیتابیس جدا نگه می‌دارد
type SecurityLoginService struct {
	uowFactory unitofwork.UnitOfWorkFactory
}

func NewSecurityLoginService(uowFactory unitofwork.UnitOfWorkFactory) *SecurityLoginService {
	return &SecurityLoginService{uowFactory: uowFactory}
}

// GetCounters شمارنده‌های بالای صفحه: موفق/ناموفق/مشکوک/آنلاین
func (s *SecurityLoginService) GetCounters(ctx context.Context) (successful, failed, suspicious, online int64, err error) {
	uow := s.uowFactory.Create()
	// کوئری‌های شمارنده (read-only)
	succ, err := uow.LoginAttemptRepository().CountSuccessful(ctx)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	fail, err := uow.LoginAttemptRepository().CountFailed(ctx)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	// مشکوک: IP های یکتای ناموفق در 24 ساعت اخیر
	since := time.Now().Add(-24 * time.Hour)
	susp, err := uow.LoginAttemptRepository().CountDistinctFailedIPsSince(ctx, since)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	// آنلاین: بر مبنای user_sessions فعال در 2 دقیقه اخیر به‌صورت ساده (از لایه مانیتورینگ اگر داشتیم)
	// اینجا عدد آنلاین را صفر می‌گذاریم تا از هندلر/دیگر سرویس پر شود در صورت نیاز.
	return succ, fail, susp, 0, nil
}

// ListLoginAttempts فهرست تاریخچه ورود با فیلتر/صفحه‌بندی
func (s *SecurityLoginService) ListLoginAttempts(ctx context.Context, filter repository.LoginAttemptFilter, page, limit int) (items interface{}, total int64, err error) {
	uow := s.uowFactory.Create()
	list, cnt, err := uow.LoginAttemptRepository().List(ctx, filter, page, limit)
	if err != nil {
		return nil, 0, err
	}
	return list, cnt, nil
}

// BlockIP مسدودکردن IP با علت و مدت زمان
func (s *SecurityLoginService) BlockIP(ctx context.Context, ip, reason, blockedBy string, durationHours int) error {
	uow := s.uowFactory.Create()
	// نوشتن در یک تراکنش
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() {
		_ = uow.Rollback()
	}()

	var expires *time.Time
	if durationHours > 0 {
		t := time.Now().Add(time.Duration(durationHours) * time.Hour)
		expires = &t
	}
	// از ریپازیتوری BlockedIP استفاده می‌کنیم (مستقیماً از uow در دسترس نیست، تزریق روی DB انجام می‌شود)
	blockedRepo := &repository.BlockedIPRepository{DB: uow.DB()}
	if err := blockedRepo.CreateOrReactivate(ctx, ip, reason, blockedBy, expires); err != nil {
		return err
	}

	return uow.Commit()
}

// UnblockIP آزادسازی IP
func (s *SecurityLoginService) UnblockIP(ctx context.Context, ip string) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	defer func() {
		_ = uow.Rollback()
	}()
	blockedRepo := &repository.BlockedIPRepository{DB: uow.DB()}
	if err := blockedRepo.Deactivate(ctx, ip); err != nil {
		return err
	}
	return uow.Commit()
}
