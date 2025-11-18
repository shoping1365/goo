package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"my-go-backend/internal/search/engine"
	"my-go-backend/internal/search/models"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Service سرویس اصلی جستجو است که بین لایه HTTP و موتور جستجوی داخلی قرار می‌گیرد.
type Service struct {
	client       engine.Client
	metrics      *metrics
	defaultIndex string
}

type metrics struct {
	searchLatency *prometheus.HistogramVec
	searchTotal   *prometheus.CounterVec
	suggestTotal  *prometheus.CounterVec
}

// New سازنده‌ی سرویس.
func New(client engine.Client, allowedIndexes []string, enableMetrics bool) *Service {
	var m *metrics
	if enableMetrics {
		m = &metrics{
			searchLatency: promauto.NewHistogramVec(prometheus.HistogramOpts{
				Name:    "search_service_latency_seconds",
				Help:    "مدت زمان اجرای درخواست جستجو",
				Buckets: prometheus.DefBuckets,
			}, []string{"index", "status"}),
			searchTotal: promauto.NewCounterVec(prometheus.CounterOpts{
				Name: "search_service_requests_total",
				Help: "تعداد کل درخواست‌های جستجو",
			}, []string{"index", "status"}),
			suggestTotal: promauto.NewCounterVec(prometheus.CounterOpts{
				Name: "search_service_suggest_total",
				Help: "تعداد کل درخواست‌های پیشنهاد",
			}, []string{"index", "status"}),
		}
	}

	var defaultIndex string
	for _, idx := range allowedIndexes {
		if strings.TrimSpace(idx) != "" {
			defaultIndex = strings.TrimSpace(idx)
			break
		}
	}

	return &Service{client: client, metrics: m, defaultIndex: defaultIndex}
}

var (
	errEmptyQuery   = errors.New("query is required")
	errInvalidLimit = errors.New("invalid page size")
)

// ErrEmptyQuery خطای مربوط به خالی بودن جستجو را برمی‌گرداند.
func ErrEmptyQuery() error {
	return errEmptyQuery
}

// ErrInvalidLimit خطای مربوط به اندازه صفحه‌ی نامعتبر را برمی‌گرداند.
func ErrInvalidLimit() error {
	return errInvalidLimit
}

// Search اجرای درخواست جستجو با اعتبارسنجی اولیه.
func (s *Service) Search(ctx context.Context, req models.SearchRequest) (models.SearchResponse, error) {
	req.Query = strings.TrimSpace(req.Query)
	if req.Query == "" {
		s.recordSearchMetrics(req.Index, "invalid", 0)
		return models.SearchResponse{}, errEmptyQuery
	}

	if req.Page.Size == 0 {
		req.Page.Size = 20
	}
	if req.Page.Size < 0 || req.Page.Size > 500 {
		s.recordSearchMetrics(req.Index, "invalid", 0)
		return models.SearchResponse{}, errInvalidLimit
	}

	if req.Index == "" {
		req.Index = s.defaultIndex
	}

	start := time.Now()
	resp, err := s.client.Search(ctx, req)
	status := statusFromError(err)
	s.recordSearchMetrics(req.Index, status, time.Since(start))
	return resp, err
}

// Suggest اجرای پیشنهاد.
func (s *Service) Suggest(ctx context.Context, req models.SuggestRequest) (models.SuggestResponse, error) {
	if req.Limit <= 0 {
		req.Limit = 5
	}
	if req.Index == "" {
		req.Index = s.defaultIndex
	}

	resp, err := s.client.Suggest(ctx, req)
	status := statusFromError(err)
	if s.metrics != nil {
		s.metrics.suggestTotal.WithLabelValues(req.Index, status).Inc()
	}
	return resp, err
}

// Health وضع سلامت موتور را برمی‌گرداند.
func (s *Service) Health(ctx context.Context) (models.HealthResponse, error) {
	return s.client.Health(ctx)
}

// IndexDocument درج یا به‌روزرسانی سند در ایندکس مشخص
func (s *Service) IndexDocument(ctx context.Context, index string, id string, doc map[string]any) error {
	index = strings.TrimSpace(index)
	if index == "" {
		if s.defaultIndex == "" {
			return errors.New("no default index configured")
		}
		index = s.defaultIndex
	}

	if id = strings.TrimSpace(id); id == "" {
		return errors.New("document id is required")
	}

	return s.client.IndexDocument(ctx, index, id, doc)
}

// DeleteDocument حذف سند از ایندکس مشخص
func (s *Service) DeleteDocument(ctx context.Context, index string, id string) error {
	index = strings.TrimSpace(index)
	if index == "" {
		if s.defaultIndex == "" {
			return errors.New("no default index configured")
		}
		index = s.defaultIndex
	}

	if id = strings.TrimSpace(id); id == "" {
		return errors.New("document id is required")
	}

	return s.client.DeleteDocument(ctx, index, id)
}

func (s *Service) recordSearchMetrics(index, status string, dur time.Duration) {
	if s.metrics == nil {
		return
	}
	if index == "" {
		index = "_unknown"
	}
	s.metrics.searchTotal.WithLabelValues(index, status).Inc()
	if dur > 0 {
		s.metrics.searchLatency.WithLabelValues(index, status).Observe(dur.Seconds())
	}
}

func statusFromError(err error) string {
	if err == nil {
		return "success"
	}
	if errors.Is(err, errEmptyQuery) || errors.Is(err, errInvalidLimit) {
		return "invalid"
	}
	return "error"
}
