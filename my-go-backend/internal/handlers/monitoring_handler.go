package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// MonitoringHandler هندلر مانیتورینگ عملکرد چت
type MonitoringHandler struct {
	monitoringService *services.MonitoringService
}

// NewMonitoringHandler ایجاد نمونه جدید از هندلر مانیتورینگ
func NewMonitoringHandler(monitoringService *services.MonitoringService) *MonitoringHandler {
	return &MonitoringHandler{
		monitoringService: monitoringService,
	}
}

// GetPerformanceMetrics دریافت متریک‌های عملکرد
// @Summary دریافت متریک‌های عملکرد چت
// @Tags Monitoring
// @Accept json
// @Produce json
// @Param time_range query string false "بازه زمانی (1h, 24h, 7d, 30d)" default(24h)
// @Param limit query int false "تعداد رکوردها" default(100)
// @Success 200 {object} models.PerformanceMetricsResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/performance [get]
func (h *MonitoringHandler) GetPerformanceMetrics(c *gin.Context) {
	timeRange := c.DefaultQuery("time_range", "24h")
	limitStr := c.DefaultQuery("limit", "100")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "مقدار limit نامعتبر است"})
		return
	}

	// دریافت متریک‌ها
	metrics, err := h.monitoringService.GetPerformanceMetrics(timeRange, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت متریک‌های عملکرد"})
		return
	}

	// دریافت خلاصه عملکرد
	summary, err := h.monitoringService.GetPerformanceSummary(timeRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت خلاصه عملکرد"})
		return
	}

	c.JSON(http.StatusOK, models.PerformanceMetricsResponse{
		Success: true,
		Data:    metrics,
		Summary: *summary,
	})
}

// GetRealTimeMetrics دریافت متریک‌های real-time
// @Summary دریافت متریک‌های real-time چت
// @Tags Monitoring
// @Accept json
// @Produce json
// @Success 200 {object} models.RealTimeMetricsResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/realtime [get]
func (h *MonitoringHandler) GetRealTimeMetrics(c *gin.Context) {
	metrics, err := h.monitoringService.GetRealTimeMetrics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت متریک‌های real-time"})
		return
	}

	c.JSON(http.StatusOK, metrics)
}

// GetSystemHealth دریافت وضعیت سلامت سیستم
// @Summary دریافت وضعیت سلامت سیستم چت
// @Tags Monitoring
// @Accept json
// @Produce json
// @Success 200 {object} models.SystemHealthResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/health [get]
func (h *MonitoringHandler) GetSystemHealth(c *gin.Context) {
	health, err := h.monitoringService.GetSystemHealth()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت وضعیت سلامت سیستم"})
		return
	}

	alerts, err := h.monitoringService.GetAlerts(10, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت هشدارها"})
		return
	}

	c.JSON(http.StatusOK, models.SystemHealthResponse{
		Success: true,
		Health:  *health,
		Alerts:  alerts,
	})
}

// GetAlerts دریافت هشدارها
// @Summary دریافت هشدارهای سیستم چت
// @Tags Monitoring
// @Accept json
// @Produce json
// @Param limit query int false "تعداد هشدارها" default(50)
// @Param unread_only query bool false "فقط هشدارهای نخوانده" default(false)
// @Success 200 {array} models.ChatAlert
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/alerts [get]
func (h *MonitoringHandler) GetAlerts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "50")
	unreadOnlyStr := c.DefaultQuery("unread_only", "false")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "مقدار limit نامعتبر است"})
		return
	}

	unreadOnly := unreadOnlyStr == "true"

	alerts, err := h.monitoringService.GetAlerts(limit, unreadOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت هشدارها"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"alerts":  alerts,
		"count":   len(alerts),
	})
}

// CreateAlert ایجاد هشدار جدید
// @Summary ایجاد هشدار جدید
// @Tags Monitoring
// @Accept json
// @Produce json
// @Param alert body models.AlertRequest true "اطلاعات هشدار"
// @Success 200 {object} models.AlertResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/alerts [post]
func (h *MonitoringHandler) CreateAlert(c *gin.Context) {
	var req models.AlertRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "داده‌های ورودی نامعتبر است"})
		return
	}

	// اعتبارسنجی level
	if req.Level != "error" && req.Level != "warning" && req.Level != "info" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "سطح هشدار نامعتبر است"})
		return
	}

	// اعتبارسنجی category
	if req.Category != "performance" && req.Category != "security" && req.Category != "system" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "دسته‌بندی هشدار نامعتبر است"})
		return
	}

	alert := models.ChatAlert{
		Level:    req.Level,
		Message:  req.Message,
		Category: req.Category,
		IsRead:   false,
	}

	err := h.monitoringService.CreateAlert(&alert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ایجاد هشدار"})
		return
	}

	c.JSON(http.StatusOK, models.AlertResponse{
		Success: true,
		Alert:   alert,
	})
}

// MarkAlertAsRead علامت‌گذاری هشدار به عنوان خوانده شده
// @Summary علامت‌گذاری هشدار به عنوان خوانده شده
// @Tags Monitoring
// @Accept json
// @Produce json
// @Param id path int true "شناسه هشدار"
// @Success 200 {object} gin.H
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/alerts/{id}/read [put]
func (h *MonitoringHandler) MarkAlertAsRead(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه هشدار نامعتبر است"})
		return
	}

	err = h.monitoringService.MarkAlertAsRead(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در به‌روزرسانی هشدار"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "هشدار به عنوان خوانده شده علامت‌گذاری شد",
	})
}

// CollectMetrics جمع‌آوری متریک‌های سیستم
// @Summary جمع‌آوری متریک‌های سیستم
// @Tags Monitoring
// @Accept json
// @Produce json
// @Success 200 {object} gin.H
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/collect [post]
func (h *MonitoringHandler) CollectMetrics(c *gin.Context) {
	err := h.monitoringService.CollectSystemMetrics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در جمع‌آوری متریک‌ها"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"message":   "متریک‌های سیستم با موفقیت جمع‌آوری شدند",
		"timestamp": time.Now(),
	})
}

// GetPerformanceChart دریافت داده‌های نمودار عملکرد
// @Summary دریافت داده‌های نمودار عملکرد
// @Tags Monitoring
// @Accept json
// @Produce json
// @Param time_range query string false "بازه زمانی (1h, 24h, 7d, 30d)" default(24h)
// @Param metric query string false "نوع متریک (response_time, messages, connections)" default(response_time)
// @Success 200 {object} gin.H
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/chart [get]
func (h *MonitoringHandler) GetPerformanceChart(c *gin.Context) {
	timeRange := c.DefaultQuery("time_range", "24h")
	metric := c.DefaultQuery("metric", "response_time")

	// دریافت متریک‌ها برای نمودار
	metrics, err := h.monitoringService.GetPerformanceMetrics(timeRange, 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت داده‌های نمودار"})
		return
	}

	// تبدیل به فرمت مناسب برای نمودار
	var labels []string
	var data []float64

	for i := len(metrics) - 1; i >= 0; i-- {
		labels = append(labels, metrics[i].CreatedAt.Format("15:04"))

		switch metric {
		case "response_time":
			data = append(data, metrics[i].AvgResponseTime)
		case "messages":
			data = append(data, metrics[i].MessagesPerSecond)
		case "connections":
			data = append(data, float64(metrics[i].ActiveConnections))
		default:
			data = append(data, metrics[i].AvgResponseTime)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"labels":    labels,
		"data":      data,
		"metric":    metric,
		"timeRange": timeRange,
	})
}

// GetSystemOverview دریافت نمای کلی سیستم
// @Summary دریافت نمای کلی سیستم چت
// @Tags Monitoring
// @Accept json
// @Produce json
// @Success 200 {object} gin.H
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/overview [get]
func (h *MonitoringHandler) GetSystemOverview(c *gin.Context) {
	// دریافت متریک‌های real-time
	realTimeMetrics, err := h.monitoringService.GetRealTimeMetrics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت نمای کلی سیستم"})
		return
	}

	// دریافت خلاصه عملکرد 24 ساعته
	summary, err := h.monitoringService.GetPerformanceSummary("24h")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت خلاصه عملکرد"})
		return
	}

	// دریافت هشدارهای نخوانده
	unreadAlerts, err := h.monitoringService.GetAlerts(10, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت هشدارها"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"overview": gin.H{
			"currentMetrics":     realTimeMetrics.Metrics,
			"systemHealth":       realTimeMetrics.Health,
			"performanceSummary": summary,
			"unreadAlerts":       unreadAlerts,
			"lastUpdated":        time.Now(),
		},
	})
}

// ToggleMonitoring تغییر وضعیت مانیتورینگ
// @Summary تغییر وضعیت مانیتورینگ (فعال/غیرفعال)
// @Tags Monitoring
// @Accept json
// @Produce json
// @Param enabled body bool true "وضعیت فعال بودن"
// @Success 200 {object} gin.H
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/toggle [post]
func (h *MonitoringHandler) ToggleMonitoring(c *gin.Context) {
	// نکته: برای bool از binding:"required" استفاده نکنید، false به عنوان مقدار صفر باعث 400 می‌شود
	var req struct {
		Enabled *bool `json:"enabled"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Enabled == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "داده‌های ورودی نامعتبر است"})
		return
	}

	var err error
	if *req.Enabled {
		err = h.monitoringService.EnableMonitoring()
	} else {
		err = h.monitoringService.DisableMonitoring()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در تغییر وضعیت مانیتورینگ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("مانیتورینگ %s شد", map[bool]string{true: "فعال", false: "غیرفعال"}[*req.Enabled]),
		"enabled": *req.Enabled,
	})
}

// GetMonitoringStatus دریافت وضعیت مانیتورینگ
// @Summary دریافت وضعیت فعلی مانیتورینگ
// @Tags Monitoring
// @Accept json
// @Produce json
// @Success 200 {object} gin.H
// @Failure 500 {object} models.ErrorResponse
// @Router /api/admin/monitoring/status [get]
func (h *MonitoringHandler) GetMonitoringStatus(c *gin.Context) {
	enabled := h.monitoringService.IsMonitoringEnabled()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"enabled": enabled,
		"message": fmt.Sprintf("مانیتورینگ %s است", map[bool]string{true: "فعال", false: "غیرفعال"}[enabled]),
	})
}
