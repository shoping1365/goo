package services

import (
	"context"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/repository"
)

// LoginMonitoringService
// سرویس مدیریت تاریخچه و آمار ورودها
// این سرویس از UnitOfWork استفاده می‌کند و منطق دامنه را در خود نگه می‌دارد.
type LoginMonitoringService struct {
	uowFactory unitofwork.UnitOfWorkFactory
}

// NewLoginMonitoringService سازنده سرویس
func NewLoginMonitoringService(uowFactory unitofwork.UnitOfWorkFactory) *LoginMonitoringService {
	return &LoginMonitoringService{uowFactory: uowFactory}
}

// ListHistory
// دریافت تاریخچه ورود با فیلتر و صفحه‌بندی
func (s *LoginMonitoringService) ListHistory(ctx context.Context, filter repository.LoginAttemptFilter, page, limit int) (attempts interface{}, total int64, err error) {
	uow := s.uowFactory.Create()
	// فقط خواندن؛ تراکنش الزامی نیست
	list, totalCount, e := uow.LoginAttemptRepository().List(ctx, filter, page, limit)
	if e != nil {
		return nil, 0, e
	}
	return list, totalCount, nil
}

// GetStats
// دریافت آمار کلی ورودها
type LoginStats struct {
	SuccessfulLogins   int64   `json:"successfulLogins"`
	FailedLogins       int64   `json:"failedLogins"`
	SuccessRate        float64 `json:"successRate"`
	SuspiciousActivity int64   `json:"suspiciousActivity"`
	OnlineUsers        int64   `json:"onlineUsers"`
}

func (s *LoginMonitoringService) GetStats(ctx context.Context) (LoginStats, error) {
	uow := s.uowFactory.Create()

	succ, err := uow.LoginAttemptRepository().CountSuccessful(ctx)
	if err != nil {
		return LoginStats{}, err
	}
	fail, err := uow.LoginAttemptRepository().CountFailed(ctx)
	if err != nil {
		return LoginStats{}, err
	}
	total := succ + fail
	var rate float64
	if total > 0 {
		rate = float64(succ) / float64(total) * 100
	}

	since := time.Now().AddDate(0, 0, -7)
	sus, err := uow.LoginAttemptRepository().CountDistinctFailedIPsSince(ctx, since)
	if err != nil {
		return LoginStats{}, err
	}

	// کاربران آنلاین (جلسات فعال)
	// معیار: سشن فعال (is_active=true) با تاریخ انقضا در آینده
	var online int64
	if err := uow.GetTx().WithContext(ctx).Table("user_sessions").
		Where("is_active = ? AND expires_at > NOW()", true).
		Distinct("user_id").Count(&online).Error; err != nil {
		return LoginStats{}, err
	}

	return LoginStats{
		SuccessfulLogins:   succ,
		FailedLogins:       fail,
		SuccessRate:        rate,
		SuspiciousActivity: sus,
		OnlineUsers:        online,
	}, nil
}
