package services

import (
	"errors"
	"fmt"
	"math/rand"
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

// MonitoringService سرویس مانیتورینگ عملکرد چت
type MonitoringService struct {
	db                  *gorm.DB
	isMonitoringEnabled bool
	stopChan            chan bool
}

// NewMonitoringService ایجاد نمونه جدید از سرویس مانیتورینگ
func NewMonitoringService(db *gorm.DB) *MonitoringService {
	return &MonitoringService{
		db:                  db,
		isMonitoringEnabled: true,
		stopChan:            make(chan bool),
	}
}

// SavePerformanceMetrics ذخیره متریک‌های عملکرد
func (ms *MonitoringService) SavePerformanceMetrics(metrics *models.ChatPerformanceMetrics) error {
	// برخی دیتابیس‌های قدیمی ستون deleted_at را ندارند؛ از درج آن صرف‌نظر می‌کنیم
	return ms.db.Omit("DeletedAt").Create(metrics).Error
}

// GetPerformanceMetrics دریافت متریک‌های عملکرد
func (ms *MonitoringService) GetPerformanceMetrics(timeRange string, limit int) ([]models.ChatPerformanceMetrics, error) {
	var metrics []models.ChatPerformanceMetrics

	query := ms.db.Model(&models.ChatPerformanceMetrics{})

	// فیلتر بر اساس بازه زمانی
	switch timeRange {
	case "1h":
		query = query.Where("created_at >= ?", time.Now().Add(-1*time.Hour))
	case "24h":
		query = query.Where("created_at >= ?", time.Now().Add(-24*time.Hour))
	case "7d":
		query = query.Where("created_at >= ?", time.Now().Add(-7*24*time.Hour))
	case "30d":
		query = query.Where("created_at >= ?", time.Now().Add(-30*24*time.Hour))
	}

	err := query.Order("created_at DESC").Limit(limit).Find(&metrics).Error
	return metrics, err
}

// GetPerformanceSummary دریافت خلاصه عملکرد
func (ms *MonitoringService) GetPerformanceSummary(timeRange string) (*models.PerformanceSummary, error) {
	var summary models.PerformanceSummary

	// محاسبه میانگین زمان پاسخ
	err := ms.db.Model(&models.ChatPerformanceMetrics{}).
		Select("AVG(avg_response_time) as avg_response_time, MAX(avg_response_time) as max_response_time, MIN(avg_response_time) as min_response_time").
		Where("created_at >= ?", getTimeRangeStart(timeRange)).
		Scan(&summary).Error
	if err != nil {
		return nil, err
	}

	// محاسبه تعداد کل اتصالات
	err = ms.db.Model(&models.ChatPerformanceMetrics{}).
		Select("SUM(active_connections) as total_connections").
		Where("created_at >= ?", getTimeRangeStart(timeRange)).
		Scan(&summary).Error
	if err != nil {
		return nil, err
	}

	// محاسبه میانگین نرخ خطا
	err = ms.db.Model(&models.ChatPerformanceMetrics{}).
		Select("AVG(error_rate) as error_rate").
		Where("created_at >= ?", getTimeRangeStart(timeRange)).
		Scan(&summary).Error
	if err != nil {
		return nil, err
	}

	// محاسبه امتیاز عملکرد
	summary.PerformanceScore = ms.calculatePerformanceScore(&summary)

	return &summary, nil
}

// CreateAlert ایجاد هشدار جدید
func (ms *MonitoringService) CreateAlert(alert *models.ChatAlert) error {
	return ms.db.Create(alert).Error
}

// GetAlerts دریافت هشدارها
func (ms *MonitoringService) GetAlerts(limit int, unreadOnly bool) ([]models.ChatAlert, error) {
	var alerts []models.ChatAlert

	query := ms.db.Model(&models.ChatAlert{})
	if unreadOnly {
		query = query.Where("is_read = ?", false)
	}

	err := query.Order("created_at DESC").Limit(limit).Find(&alerts).Error
	return alerts, err
}

// MarkAlertAsRead علامت‌گذاری هشدار به عنوان خوانده شده
func (ms *MonitoringService) MarkAlertAsRead(alertID uint) error {
	return ms.db.Model(&models.ChatAlert{}).Where("id = ?", alertID).Update("is_read", true).Error
}

// GetSystemHealth دریافت وضعیت سلامت سیستم
func (ms *MonitoringService) GetSystemHealth() (*models.ChatSystemHealth, error) {
	var health models.ChatSystemHealth

	// استفاده از Unscoped برای جلوگیری از فیلتر soft delete در صورت نبود ستون deleted_at
	err := ms.db.Unscoped().First(&health).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ایجاد رکورد اولیه
			health = models.ChatSystemHealth{
				Status:           "healthy",
				Uptime:           "100%",
				LastCheck:        time.Now(),
				TotalConnections: 0,
				ActiveSessions:   0,
				TotalMessages:    0,
			}
			// در ایجاد اولیه از درج ستون DeletedAt صرف‌نظر می‌کنیم
			err = ms.db.Omit("DeletedAt").Create(&health).Error
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &health, nil
}

// UpdateSystemHealth به‌روزرسانی وضعیت سلامت سیستم
func (ms *MonitoringService) UpdateSystemHealth(health *models.ChatSystemHealth) error {
	health.LastCheck = time.Now()
	// Unscoped: جلوگیری از افزودن شرط deleted_at در WHERE
	return ms.db.Unscoped().Model(&models.ChatSystemHealth{}).
		Where("id = ?", health.ID).
		Omit("DeletedAt").
		Updates(map[string]interface{}{
			"status":            health.Status,
			"uptime":            health.Uptime,
			"last_check":        health.LastCheck,
			"total_connections": health.TotalConnections,
			"active_sessions":   health.ActiveSessions,
			"total_messages":    health.TotalMessages,
			"updated_at":        time.Now(),
		}).Error
}

// GetRealTimeMetrics دریافت متریک‌های real-time
func (ms *MonitoringService) GetRealTimeMetrics() (*models.RealTimeMetricsResponse, error) {
	// دریافت آخرین متریک‌های عملکرد
	var latestMetrics models.ChatPerformanceMetrics
	err := ms.db.Order("created_at DESC").First(&latestMetrics).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// اگر متریک وجود ندارد، یک متریک نمونه ایجاد کن
	if errors.Is(err, gorm.ErrRecordNotFound) {
		latestMetrics = ms.generateSampleMetrics()
	}

	// دریافت وضعیت سلامت سیستم
	health, err := ms.GetSystemHealth()
	if err != nil {
		return nil, err
	}

	// دریافت هشدارهای اخیر
	alerts, err := ms.GetAlerts(10, false)
	if err != nil {
		return nil, err
	}

	return &models.RealTimeMetricsResponse{
		Success: true,
		Metrics: latestMetrics,
		Health:  *health,
		Alerts:  alerts,
	}, nil
}

// CollectSystemMetrics جمع‌آوری متریک‌های سیستم
func (ms *MonitoringService) CollectSystemMetrics() error {
	// در یک سیستم واقعی، اینجا متریک‌های واقعی سیستم جمع‌آوری می‌شوند
	// فعلاً از داده‌های نمونه استفاده می‌کنیم

	metrics := ms.generateSampleMetrics()

	// ذخیره متریک‌ها
	err := ms.SavePerformanceMetrics(&metrics)
	if err != nil {
		return fmt.Errorf("خطا در ذخیره متریک‌های عملکرد: %v", err)
	}

	// بررسی و ایجاد هشدارها
	err = ms.checkAndCreateAlerts(&metrics)
	if err != nil {
		return fmt.Errorf("خطا در بررسی هشدارها: %v", err)
	}

	// به‌روزرسانی وضعیت سلامت سیستم
	err = ms.updateSystemHealth(&metrics)
	if err != nil {
		return fmt.Errorf("خطا در به‌روزرسانی وضعیت سلامت: %v", err)
	}

	return nil
}

// generateSampleMetrics تولید متریک‌های نمونه
func (ms *MonitoringService) generateSampleMetrics() models.ChatPerformanceMetrics {
	return models.ChatPerformanceMetrics{
		ActiveConnections: 150 + rand.Intn(50),
		AvgResponseTime:   150 + rand.Float64()*100,
		MessagesPerSecond: 30 + rand.Float64()*40,
		ErrorRate:         rand.Float64() * 2.0,
		CPUUsage:          40 + rand.Float64()*30,
		MemoryUsage:       60 + rand.Float64()*25,
		DiskUsage:         70 + rand.Float64()*20,
		NetworkUsage:      80 + rand.Float64()*15,
		PerformanceScore:  85 + rand.Float64()*15,
	}
}

// checkAndCreateAlerts بررسی و ایجاد هشدارها
func (ms *MonitoringService) checkAndCreateAlerts(metrics *models.ChatPerformanceMetrics) error {
	// بررسی زمان پاسخ
	if metrics.AvgResponseTime > 300 {
		alert := models.ChatAlert{
			Level:    "warning",
			Message:  fmt.Sprintf("زمان پاسخ افزایش یافته است: %.2fms", metrics.AvgResponseTime),
			Category: "performance",
		}
		ms.CreateAlert(&alert)
	}

	// بررسی نرخ خطا
	if metrics.ErrorRate > 1.0 {
		alert := models.ChatAlert{
			Level:    "error",
			Message:  fmt.Sprintf("نرخ خطا بالا است: %.2f%%", metrics.ErrorRate),
			Category: "performance",
		}
		ms.CreateAlert(&alert)
	}

	// بررسی استفاده از CPU
	if metrics.CPUUsage > 80 {
		alert := models.ChatAlert{
			Level:    "warning",
			Message:  fmt.Sprintf("استفاده از CPU بالا است: %.2f%%", metrics.CPUUsage),
			Category: "system",
		}
		ms.CreateAlert(&alert)
	}

	// بررسی استفاده از حافظه
	if metrics.MemoryUsage > 85 {
		alert := models.ChatAlert{
			Level:    "warning",
			Message:  fmt.Sprintf("استفاده از حافظه بالا است: %.2f%%", metrics.MemoryUsage),
			Category: "system",
		}
		ms.CreateAlert(&alert)
	}

	return nil
}

// updateSystemHealth به‌روزرسانی وضعیت سلامت سیستم
func (ms *MonitoringService) updateSystemHealth(metrics *models.ChatPerformanceMetrics) error {
	health, err := ms.GetSystemHealth()
	if err != nil {
		return fmt.Errorf("خطا در دریافت وضعیت سلامت: %v", err)
	}

	if health == nil {
		return fmt.Errorf("وضعیت سلامت null است")
	}

	// تعیین وضعیت بر اساس متریک‌ها
	status := "healthy"
	if metrics.ErrorRate > 2.0 || metrics.AvgResponseTime > 500 {
		status = "error"
	} else if metrics.ErrorRate > 1.0 || metrics.AvgResponseTime > 300 {
		status = "warning"
	}

	health.Status = status
	health.TotalConnections += metrics.ActiveConnections
	health.ActiveSessions = metrics.ActiveConnections
	health.TotalMessages += int(metrics.MessagesPerSecond)

	return ms.UpdateSystemHealth(health)
}

// calculatePerformanceScore محاسبه امتیاز عملکرد
func (ms *MonitoringService) calculatePerformanceScore(summary *models.PerformanceSummary) float64 {
	score := 100.0

	// کاهش امتیاز بر اساس زمان پاسخ
	if summary.AvgResponseTime > 200 {
		score -= (summary.AvgResponseTime - 200) * 0.1
	}

	// کاهش امتیاز بر اساس نرخ خطا
	score -= summary.ErrorRate * 10

	// اطمینان از اینکه امتیاز بین 0 تا 100 باشد
	if score < 0 {
		score = 0
	} else if score > 100 {
		score = 100
	}

	return score
}

// getTimeRangeStart دریافت زمان شروع بازه زمانی
func getTimeRangeStart(timeRange string) time.Time {
	now := time.Now()

	switch timeRange {
	case "1h":
		return now.Add(-1 * time.Hour)
	case "24h":
		return now.Add(-24 * time.Hour)
	case "7d":
		return now.Add(-7 * 24 * time.Hour)
	case "30d":
		return now.Add(-30 * 24 * time.Hour)
	default:
		return now.Add(-24 * time.Hour) // پیش‌فرض 24 ساعت
	}
}

// EnableMonitoring فعال کردن مانیتورینگ
func (ms *MonitoringService) EnableMonitoring() error {
	ms.isMonitoringEnabled = true
	ms.stopChan = make(chan bool)
	go ms.StartMonitoring()
	return nil
}

// DisableMonitoring غیرفعال کردن مانیتورینگ
func (ms *MonitoringService) DisableMonitoring() error {
	ms.isMonitoringEnabled = false
	if ms.stopChan != nil {
		close(ms.stopChan)
	}
	return nil
}

// IsMonitoringEnabled بررسی وضعیت فعال بودن مانیتورینگ
func (ms *MonitoringService) IsMonitoringEnabled() bool {
	return ms.isMonitoringEnabled
}

// StartMonitoring شروع مانیتورینگ خودکار
func (ms *MonitoringService) StartMonitoring() {
	go func() {
		ticker := time.NewTicker(30 * time.Second) // هر 30 ثانیه
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if !ms.isMonitoringEnabled {
					return
				}
				err := ms.CollectSystemMetrics()
				if err != nil {
					fmt.Printf("خطا در جمع‌آوری متریک‌ها: %v\n", err)
				}
			case <-ms.stopChan:
				return
			}
		}
	}()
}
