package repository

import (
	"context"
	"my-go-backend/internal/models"
	"strings"
	"time"

	"gorm.io/gorm"
)

// TrafficLogRepositoryInterface قرارداد ریپازیتوری لاگ‌های ترافیک
type TrafficLogRepositoryInterface interface {
	Create(ctx context.Context, log *models.TrafficLog) error
	BulkCreate(ctx context.Context, logs []models.TrafficLog) error
	List(ctx context.Context, filter TrafficLogFilter, page, limit int) ([]models.TrafficLog, int64, error)
	ListByIP(ctx context.Context, ip string, page, limit int) ([]models.TrafficLog, int64, error)
	CountSince(ctx context.Context, from time.Time) (int64, error)
	CountSuspiciousSince(ctx context.Context, from time.Time) (int64, error)
	HourlyCountToday(ctx context.Context) ([]models.HourlyTrafficData, error)
	ListLatestByUserAgentContains(ctx context.Context, patterns []string, page, limit int) ([]models.TrafficLog, int64, error)
}

// TrafficLogFilter فیلترهای جستجو
type TrafficLogFilter struct {
	AdType      string
	IPAddress   string
	City        string
	RequestPath string
	StatusCode  *int
	TimeFrom    *int64 // unix seconds
}

type TrafficLogRepository struct {
	DB *gorm.DB
}

func (r *TrafficLogRepository) Create(ctx context.Context, log *models.TrafficLog) error {
	return r.DB.WithContext(ctx).Create(log).Error
}

func (r *TrafficLogRepository) BulkCreate(ctx context.Context, logs []models.TrafficLog) error {
	if len(logs) == 0 {
		return nil
	}
	return r.DB.WithContext(ctx).CreateInBatches(&logs, 1000).Error
}

func (r *TrafficLogRepository) List(ctx context.Context, filter TrafficLogFilter, page, limit int) ([]models.TrafficLog, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 1000 {
		limit = 50
	}

	q := r.DB.WithContext(ctx).Model(&models.TrafficLog{})
	// فیلترها
	if filter.AdType != "" {
		q = q.Where("ad_type = ?", filter.AdType)
	}
	if filter.IPAddress != "" {
		q = q.Where("ip_address = ?", filter.IPAddress)
	}
	if filter.City != "" {
		q = q.Where("city = ?", filter.City)
	}
	if filter.RequestPath != "" {
		// استفاده از ILIKE برای جستجوی جزئی
		q = q.Where("request_path ILIKE ?", "%"+filter.RequestPath+"%")
	}
	if filter.StatusCode != nil {
		q = q.Where("status_code = ?", *filter.StatusCode)
	}
	if filter.TimeFrom != nil {
		q = q.Where("created_at >= to_timestamp(?)", *filter.TimeFrom)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var logs []models.TrafficLog
	if err := q.Order("created_at DESC").Offset((page - 1) * limit).Limit(limit).Find(&logs).Error; err != nil {
		return nil, 0, err
	}
	return logs, total, nil
}

func (r *TrafficLogRepository) ListByIP(ctx context.Context, ip string, page, limit int) ([]models.TrafficLog, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 1000 {
		limit = 50
	}
	q := r.DB.WithContext(ctx).Model(&models.TrafficLog{}).Where("ip_address = ?", ip)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var logs []models.TrafficLog
	if err := q.Order("created_at DESC").Offset((page - 1) * limit).Limit(limit).Find(&logs).Error; err != nil {
		return nil, 0, err
	}
	return logs, total, nil
}

// ListLatestByUserAgentContains: آخرین لاگ برای هر User-Agent که شامل هر یک از الگوها باشد
func (r *TrafficLogRepository) ListLatestByUserAgentContains(ctx context.Context, patterns []string, page, limit int) ([]models.TrafficLog, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 1000 {
		limit = 50
	}
	db := r.DB.WithContext(ctx)

	likeConds := make([]string, 0, len(patterns))
	args := make([]interface{}, 0, len(patterns))
	for _, p := range patterns {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		likeConds = append(likeConds, "LOWER(user_agent) LIKE ?")
		args = append(args, "%"+strings.ToLower(p)+"%")
	}
	if len(likeConds) == 0 {
		likeConds = []string{"LOWER(user_agent) LIKE ?", "LOWER(user_agent) LIKE ?", "LOWER(user_agent) LIKE ?"}
		args = []interface{}{"%bot%", "%crawler%", "%spider%"}
	}
	where := "(" + strings.Join(likeConds, " OR ") + ")"

	var total int64
	countSQL := "SELECT COUNT(*) FROM (SELECT DISTINCT user_agent FROM traffic_logs WHERE " + where + ") t"
	if err := db.Raw(countSQL, args...).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	sql := "SELECT DISTINCT ON (user_agent) id, created_at, user_id, session_id, ip_address, request_method, request_path, status_code, response_time_ms, user_agent, browser, device_type, os, hostname, ad_type, city, country_code, is_suspicious, meta FROM traffic_logs WHERE " + where + " ORDER BY user_agent, created_at DESC OFFSET ? LIMIT ?"
	var rows []models.TrafficLog
	allArgs := append(args, (page-1)*limit, limit)
	if err := db.Raw(sql, allArgs...).Scan(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (r *TrafficLogRepository) CountSince(ctx context.Context, from time.Time) (int64, error) {
	var cnt int64
	err := r.DB.WithContext(ctx).Model(&models.TrafficLog{}).Where("created_at >= ?", from).Count(&cnt).Error
	return cnt, err
}

func (r *TrafficLogRepository) CountSuspiciousSince(ctx context.Context, from time.Time) (int64, error) {
	var cnt int64
	// معیار اولیه: خطاهای 4xx پرتکرار (401,403,404,429) و همه 5xx
	err := r.DB.WithContext(ctx).Model(&models.TrafficLog{}).
		Where("created_at >= ? AND (status_code IN (401,403,404,429) OR status_code >= 500)", from).
		Count(&cnt).Error
	return cnt, err
}

func (r *TrafficLogRepository) HourlyCountToday(ctx context.Context) ([]models.HourlyTrafficData, error) {
	type row struct {
		Hour  int
		Count int64
	}
	var rows []row
	err := r.DB.WithContext(ctx).
		Raw(`
            SELECT EXTRACT(HOUR FROM created_at)::int AS hour, COUNT(*)::bigint AS count
            FROM traffic_logs
            WHERE created_at::date = CURRENT_DATE
            GROUP BY 1
            ORDER BY 1
        `).Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	var out []models.HourlyTrafficData
	for _, r := range rows {
		out = append(out, models.HourlyTrafficData{Hour: r.Hour, Count: r.Count})
	}
	return out, nil
}
