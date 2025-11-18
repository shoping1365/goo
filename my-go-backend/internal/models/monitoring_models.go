package models

import (
	"time"

	"gorm.io/gorm"
)

// ChatPerformanceMetrics متریک‌های عملکرد چت
type ChatPerformanceMetrics struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	ActiveConnections int            `json:"active_connections"`
	AvgResponseTime   float64        `json:"avg_response_time"`
	MessagesPerSecond float64        `json:"messages_per_second"`
	ErrorRate         float64        `json:"error_rate"`
	CPUUsage          float64        `json:"cpu_usage"`
	MemoryUsage       float64        `json:"memory_usage"`
	DiskUsage         float64        `json:"disk_usage"`
	NetworkUsage      float64        `json:"network_usage"`
	PerformanceScore  float64        `json:"performance_score"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// ChatAlert هشدارهای سیستم چت
type ChatAlert struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Level     string         `json:"level"` // error, warning, info
	Message   string         `json:"message"`
	Category  string         `json:"category"` // performance, security, system
	IsRead    bool           `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// ChatSystemHealth وضعیت سلامت سیستم چت
type ChatSystemHealth struct {
	ID               uint           `json:"id" gorm:"primaryKey"`
	Status           string         `json:"status"` // healthy, warning, error
	Uptime           string         `json:"uptime"`
	LastCheck        time.Time      `json:"last_check"`
	TotalConnections int            `json:"total_connections"`
	ActiveSessions   int            `json:"active_sessions"`
	TotalMessages    int            `json:"total_messages"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName تعیین نام جدول برای مدل وضعیت سلامت سیستم چت
// توجه: برای هم‌خوانی با مایگریشن 38 که جدول را با نام مفرد ایجاد کرده است
func (ChatSystemHealth) TableName() string { // فارسی: نام جدول سفارشی برای جلوگیری از جمع بستن خودکار GORM
	return "chat_system_health"
}

// PerformanceMetricsRequest درخواست متریک‌های عملکرد
type PerformanceMetricsRequest struct {
	TimeRange string `json:"time_range"` // 1h, 24h, 7d, 30d
	Interval  string `json:"interval"`   // 1m, 5m, 1h, 1d
}

// PerformanceMetricsResponse پاسخ متریک‌های عملکرد
type PerformanceMetricsResponse struct {
	Success bool                     `json:"success"`
	Data    []ChatPerformanceMetrics `json:"data"`
	Summary PerformanceSummary       `json:"summary"`
}

// PerformanceSummary خلاصه عملکرد
type PerformanceSummary struct {
	AvgResponseTime  float64 `json:"avg_response_time"`
	MaxResponseTime  float64 `json:"max_response_time"`
	MinResponseTime  float64 `json:"min_response_time"`
	TotalConnections int     `json:"total_connections"`
	TotalMessages    int     `json:"total_messages"`
	ErrorRate        float64 `json:"error_rate"`
	PerformanceScore float64 `json:"performance_score"`
}

// SystemHealthResponse پاسخ وضعیت سیستم
type SystemHealthResponse struct {
	Success bool             `json:"success"`
	Health  ChatSystemHealth `json:"health"`
	Alerts  []ChatAlert      `json:"alerts"`
}

// AlertRequest درخواست هشدار
type AlertRequest struct {
	Level    string `json:"level" binding:"required"`
	Message  string `json:"message" binding:"required"`
	Category string `json:"category" binding:"required"`
}

// AlertResponse پاسخ هشدار
type AlertResponse struct {
	Success bool      `json:"success"`
	Alert   ChatAlert `json:"alert"`
}

// RealTimeMetricsResponse پاسخ متریک‌های real-time
type RealTimeMetricsResponse struct {
	Success bool                   `json:"success"`
	Metrics ChatPerformanceMetrics `json:"metrics"`
	Health  ChatSystemHealth       `json:"health"`
	Alerts  []ChatAlert            `json:"alerts"`
}
