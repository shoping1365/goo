package services

import (
	"context"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"time"
)

// TrafficService سرویس خواندنی/نوشتنی برای ترافیک (CQRS با UoW)
type TrafficService struct {
	uowFactory unitofwork.UnitOfWorkFactory
}

func NewTrafficService(uowFactory unitofwork.UnitOfWorkFactory) *TrafficService {
	return &TrafficService{uowFactory: uowFactory}
}

// ListGeneralTraffic دریافت لاگ‌های عمومی با فیلتر و صفحه‌بندی
func (s *TrafficService) ListGeneralTraffic(ctx context.Context, filter repository.TrafficLogFilter, page, limit int) (logs interface{}, pagination map[string]interface{}, err error) {
	uow := s.uowFactory.Create()
	logsList, total, err := uow.TrafficLogRepository().List(ctx, filter, page, limit)
	if err != nil {
		return nil, nil, err
	}
	totalPages := (total + int64(limit) - 1) / int64(limit)
	return logsList, map[string]interface{}{
		"current_page":   page,
		"total_pages":    totalPages,
		"total_items":    total,
		"items_per_page": limit,
		"has_next":       page < int(totalPages),
		"has_prev":       page > 1,
	}, nil
}

// ListTrafficDetailsByIP لاگ‌های یک IP خاص
func (s *TrafficService) ListTrafficDetailsByIP(ctx context.Context, ip string, page, limit int) (logs interface{}, pagination map[string]interface{}, err error) {
	uow := s.uowFactory.Create()
	logsList, total, err := uow.TrafficLogRepository().ListByIP(ctx, ip, page, limit)
	if err != nil {
		return nil, nil, err
	}
	totalPages := (total + int64(limit) - 1) / int64(limit)
	return logsList, map[string]interface{}{
		"current_page":   page,
		"total_pages":    totalPages,
		"total_items":    total,
		"items_per_page": limit,
		"has_next":       page < int(totalPages),
		"has_prev":       page > 1,
	}, nil
}

// ConvertTimeRangeToUnix تبدیل کلید بازه زمانی به timestamp شروع
func (s *TrafficService) ConvertTimeRangeToUnix(key string) *int64 {
	now := time.Now()
	var from time.Time
	switch key {
	case "1h":
		from = now.Add(-1 * time.Hour)
	case "6h":
		from = now.Add(-6 * time.Hour)
	case "24h":
		from = now.Add(-24 * time.Hour)
	case "7d":
		from = now.AddDate(0, 0, -7)
	case "30d":
		from = now.AddDate(0, 0, -30)
	default:
		return nil
	}
	ts := from.Unix()
	return &ts
}

// GetStats محاسبه آمار کلی صفحه کارت‌های بالای داشبورد ترافیک
func (s *TrafficService) GetStats(ctx context.Context) (models.TrafficStats, error) {
	uow := s.uowFactory.Create()
	since5m := time.Now().Add(-5 * time.Minute)
	online, err := uow.TrafficSessionRepository().CountOnline(ctx, since5m)
	if err != nil {
		return models.TrafficStats{}, err
	}
	todayStart := time.Now().Truncate(24 * time.Hour)
	todayCnt, err := uow.TrafficLogRepository().CountSince(ctx, todayStart)
	if err != nil {
		return models.TrafficStats{}, err
	}
	suspCnt, err := uow.TrafficLogRepository().CountSuspiciousSince(ctx, todayStart)
	if err != nil {
		return models.TrafficStats{}, err
	}
	hourly, err := uow.TrafficLogRepository().HourlyCountToday(ctx)
	if err != nil {
		return models.TrafficStats{}, err
	}
	// تعداد حملات مسدود شده: به صورت ساده از تعداد آیتم‌های فعال استفاده می‌کنیم
	blockedList, totalActive, err := uow.BlockedIPRepository().ListActive(ctx, 1, 1)
	_ = blockedList
	if err != nil {
		totalActive = 0
	}
	return models.TrafficStats{
		OnlineUsers:        online,
		TodayRequests:      todayCnt,
		SuspiciousRequests: suspCnt,
		BlockedAttacks:     totalActive,
		HourlyTraffic:      hourly,
	}, nil
}

// ListOnlineSessions دریافت سشن‌های آنلاین در 5 دقیقه گذشته با صفحه‌بندی
func (s *TrafficService) ListOnlineSessions(ctx context.Context, page, limit int) ([]models.TrafficSession, int64, error) {
	uow := s.uowFactory.Create()
	since := time.Now().Add(-5 * time.Minute)
	return uow.TrafficSessionRepository().ListOnline(ctx, since, page, limit)
}

// BlockIP مسدودسازی IP با ثبت دلیل و مدت و غیر فعال کردن سشن‌های مرتبط
func (s *TrafficService) BlockIP(ctx context.Context, ip, reason, blockedBy string, durationHours int) error {
	uow := s.uowFactory.Create()
	var expiresAt *time.Time
	if durationHours > 0 {
		t := time.Now().Add(time.Duration(durationHours) * time.Hour)
		expiresAt = &t
	}
	if err := uow.BlockedIPRepository().CreateOrReactivate(ctx, ip, reason, blockedBy, expiresAt); err != nil {
		return err
	}
	if err := uow.TrafficSessionRepository().DeactivateByIP(ctx, ip); err != nil {
		return err
	}
	return nil
}

// UnblockIP آزادسازی IP
func (s *TrafficService) UnblockIP(ctx context.Context, ip string) error {
	uow := s.uowFactory.Create()
	return uow.BlockedIPRepository().Deactivate(ctx, ip)
}

// CleanupExpiredSessions غیر فعال کردن سشن‌های منقضی شده (بیش از 30 دقیقه غیر فعال)
func (s *TrafficService) CleanupExpiredSessions(ctx context.Context) error {
	uow := s.uowFactory.Create()
	before := time.Now().Add(-30 * time.Minute)
	return uow.TrafficSessionRepository().CleanupExpired(ctx, before)
}

// UpdateSession به‌روزرسانی یا ایجاد سشن ترافیکی کاربر بر اساس SessionID
func (s *TrafficService) UpdateSession(ctx context.Context, sessionID, ip, ua, currentPage string, userID *uint) error {
	uow := s.uowFactory.Create()
	sess, err := uow.TrafficSessionRepository().GetBySessionID(ctx, sessionID)
	if err != nil {
		return err
	}
	now := time.Now()
	if sess == nil {
		ns := &models.TrafficSession{
			SessionID:   sessionID,
			IPAddress:   ip,
			UserAgent:   ua,
			CurrentPage: currentPage,
			LoginTime:   now,
			LastSeen:    now,
			IsActive:    true,
		}
		if userID != nil {
			ns.UserID = userID
		}
		return uow.TrafficSessionRepository().Create(ctx, ns)
	}
	sess.CurrentPage = currentPage
	sess.LastSeen = now
	sess.IsActive = true
	return uow.TrafficSessionRepository().Save(ctx, sess)
}
