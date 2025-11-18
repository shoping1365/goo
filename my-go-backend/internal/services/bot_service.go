package services

import (
	"context"
	"strings"

	"my-go-backend/internal/database/unitofwork"
)

// BotService سرویس خواندن/نوشتن تنظیمات ربات و استخراج فهرست ربات‌ها از لاگ‌ها
type BotService struct {
	uowFactory unitofwork.UnitOfWorkFactory
}

func NewBotService(uowFactory unitofwork.UnitOfWorkFactory) *BotService {
	return &BotService{uowFactory: uowFactory}
}

// ListKnownBots استخراج آخرین بازدید هر ربات بر اساس الگوی UA
func (s *BotService) ListKnownBots(ctx context.Context, patterns []string, page, limit int) (interface{}, int64, error) {
	uow := s.uowFactory.Create()
	logs, total, err := uow.TrafficLogRepository().ListLatestByUserAgentContains(ctx, patterns, page, limit)
	if err != nil {
		return nil, 0, err
	}
	return logs, total, nil
}

// NormalizeCSV اضافه‌کردن بدون تکرار به CSV
func NormalizeCSV(current, add string) string {
	cur := strings.Split(strings.TrimSpace(current), ",")
	set := make(map[string]struct{})
	for _, v := range cur {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		set[v] = struct{}{}
	}
	for _, v := range strings.Split(add, ",") {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		set[v] = struct{}{}
	}
	out := make([]string, 0, len(set))
	for k := range set {
		out = append(out, k)
	}
	return strings.Join(out, ",")
}
