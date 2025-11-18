package handlers

import (
	"math"
	"net/http"
	"strconv"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TrafficMonitoringHandler handler برای نظارت بر ترافیک
type TrafficMonitoringHandler struct {
	db  *gorm.DB
	svc *services.TrafficService
}

// NewTrafficMonitoringHandler ایجاد instance جدید از handler
func NewTrafficMonitoringHandler(db *gorm.DB) *TrafficMonitoringHandler {
	// ساخت سرویس با UoW
	uowFactory := unitofwork.NewUnitOfWorkFactory(db)
	svc := services.NewTrafficService(uowFactory)
	return &TrafficMonitoringHandler{db: db, svc: svc}
}

// GetTrafficStats دریافت آمار کلی ترافیک
func (h *TrafficMonitoringHandler) GetTrafficStats(c *gin.Context) {
	stats, err := h.svc.GetStats(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در دریافت آمار"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": stats})
}

// GetOnlineUsers دریافت لیست کاربران آنلاین با صفحه‌بندی
func (h *TrafficMonitoringHandler) GetOnlineUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit > 100 {
		limit = 100
	}
	if limit < 1 {
		limit = 10
	}

	sessions, total, err := h.svc.ListOnlineSessions(c, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در دریافت کاربران آنلاین"})
		return
	}

	var onlineUsers []gin.H
	for _, session := range sessions {
		userName := "کاربر مهمان"
		userEmail := ""
		if session.User != nil {
			userName = session.User.Name
			userEmail = session.User.Email
		}
		onlineUsers = append(onlineUsers, gin.H{
			"id":          session.ID,
			"name":        userName,
			"email":       userEmail,
			"ip":          session.IPAddress,
			"currentPage": session.CurrentPage,
			"loginTime":   session.LoginTime.Format("15:04"),
			"lastSeen":    session.LastSeen.Format("15:04"),
			"userAgent":   session.UserAgent,
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": onlineUsers, "pagination": gin.H{
		"current_page":   page,
		"total_pages":    totalPages,
		"total_items":    total,
		"items_per_page": limit,
		"has_next_page":  page < totalPages,
		"has_prev_page":  page > 1,
	}})
}

// GetGeneralStats فهرست لاگ‌های عمومی با فیلترهای فرانت (ادغام در یک خروجی سازگار)
func (h *TrafficMonitoringHandler) GetGeneralStats(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if limit > 1000 {
		limit = 1000
	}
	if limit < 1 {
		limit = 50
	}

	adType := c.Query("adType")
	ipAddress := c.Query("ipAddress")
	city := c.Query("location")
	requestPath := c.Query("pagePath")
	status := c.Query("status")
	timeRange := c.Query("timeRange")

	var statusCodePtr *int
	if status != "" {
		if sc, err := strconv.Atoi(status); err == nil {
			statusCodePtr = &sc
		}
	}

	var fromUnix *int64
	if timeRange != "" {
		fromUnix = h.svc.ConvertTimeRangeToUnix(timeRange)
	}

	logs, pagination, err := h.svc.ListGeneralTraffic(c, repository.TrafficLogFilter{
		AdType:      adType,
		IPAddress:   ipAddress,
		City:        city,
		RequestPath: requestPath,
		StatusCode:  statusCodePtr,
		TimeFrom:    fromUnix,
	}, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در دریافت آمار ترافیک"})
		return
	}

	// تبدیل به ساختار مورد انتظار فرانت (کلیدهای PascalCase)
	var formatted []gin.H
	if logsSlice, ok := logs.([]models.TrafficLog); ok {
		for _, lg := range logsSlice {
			formatted = append(formatted, gin.H{
				"ID":           lg.ID,
				"AdType":       lg.AdType,
				"Location":     lg.City,
				"City":         lg.City,
				"RequestPath":  lg.RequestPath,
				"CreatedAt":    lg.CreatedAt,
				"IPAddress":    lg.IPAddress,
				"Hostname":     lg.Hostname,
				"StatusCode":   lg.StatusCode,
				"ResponseTime": lg.ResponseTimeMs,
				"Browser":      lg.Browser,
				"DeviceType":   lg.DeviceType,
				"OS":           lg.OS,
				"Display":      "",
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"traffic_logs": formatted,
			"pagination":   pagination,
		},
	})
}

// GetTrafficDetailsByIP لاگ‌های مربوط به IP مشخص
func (h *TrafficMonitoringHandler) GetTrafficDetailsByIP(c *gin.Context) {
	ip := c.Param("ip")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if limit > 1000 {
		limit = 1000
	}
	if limit < 1 {
		limit = 50
	}

	logs, pagination, err := h.svc.ListTrafficDetailsByIP(c, ip, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در دریافت جزئیات ترافیک"})
		return
	}
	// تبدیل به ساختار مورد انتظار جزئیات (کلیدهای lowerCamelCase)
	var formatted []gin.H
	if logsSlice, ok := logs.([]models.TrafficLog); ok {
		for _, lg := range logsSlice {
			formatted = append(formatted, gin.H{
				"id":         lg.ID,
				"ipAddress":  lg.IPAddress,
				"adType":     lg.AdType,
				"location":   lg.City,
				"page":       lg.RequestPath,
				"timestamp":  lg.CreatedAt.Format("2006-01-02 15:04:05"),
				"browser":    lg.Browser,
				"deviceType": lg.DeviceType,
				"os":         lg.OS,
				"hostname":   lg.Hostname,
				"status":     lg.StatusCode,
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": formatted, "pagination": pagination})
}

// GetUserDetailsByIP فعالیت‌های کاربر از روی IP (فعلاً همان لاگ‌ها)
func (h *TrafficMonitoringHandler) GetUserDetailsByIP(c *gin.Context) {
	ip := c.Param("ip")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if limit > 1000 {
		limit = 1000
	}
	if limit < 1 {
		limit = 50
	}

	logs, pagination, err := h.svc.ListTrafficDetailsByIP(c, ip, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در دریافت جزئیات کاربر"})
		return
	}

	// تبدیل به ساختار فعالیت کاربر
	var formatted []gin.H
	if logsSlice, ok := logs.([]models.TrafficLog); ok {
		for _, lg := range logsSlice {
			formatted = append(formatted, gin.H{
				"id":        lg.ID,
				"page":      lg.RequestPath,
				"action":    "بازدید صفحه",
				"timestamp": lg.CreatedAt.Format("2006-01-02 15:04:05"),
				"duration":  "-",
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": formatted, "pagination": pagination})
}

// UpdateUserSession به‌روزرسانی سشن ترافیک کاربر
func (h *TrafficMonitoringHandler) UpdateUserSession(c *gin.Context) {
	var req struct {
		SessionID   string `json:"session_id" binding:"required"`
		CurrentPage string `json:"current_page"`
		IPAddress   string `json:"ip_address"`
		UserAgent   string `json:"user_agent"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "داده‌های ورودی نامعتبر است"})
		return
	}
	var uid *uint
	if userID, exists := c.Get("user_id"); exists {
		if v, ok := userID.(uint); ok {
			uid = &v
		}
	}
	if err := h.svc.UpdateSession(c, req.SessionID, req.IPAddress, req.UserAgent, req.CurrentPage, uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در به‌روزرسانی سشن"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "سشن به‌روزرسانی شد"})
}

// LogTrafficRequest ثبت درخواست ترافیک

// BlockIP مسدود کردن آدرس IP
func (h *TrafficMonitoringHandler) BlockIP(c *gin.Context) {
	var req struct {
		IPAddress string `json:"ip_address" binding:"required"`
		Reason    string `json:"reason"`
		Duration  int    `json:"duration"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "داده‌های ورودی نامعتبر است"})
		return
	}
	blockedBy := "سیستم"
	if userName, exists := c.Get("user_name"); exists {
		if name, ok := userName.(string); ok {
			blockedBy = name
		}
	}
	if err := h.svc.BlockIP(c, req.IPAddress, req.Reason, blockedBy, req.Duration); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در مسدود کردن IP"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "آدرس IP با موفقیت مسدود شد"})
}

// UnblockIP آزاد کردن آدرس IP
func (h *TrafficMonitoringHandler) UnblockIP(c *gin.Context) {
	ipAddress := c.Param("ip")
	if err := h.svc.UnblockIP(c, ipAddress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در آزاد کردن IP"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "آدرس IP با موفقیت آزاد شد"})
}

// GetBlockedIPs دریافت لیست آدرس‌های IP مسدود شده با صفحه‌بندی
func (h *TrafficMonitoringHandler) GetBlockedIPs(c *gin.Context) {
	// دریافت پارامترهای صفحه‌بندی
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// محدودیت تعداد آیتم در هر صفحه
	if limit > 100 {
		limit = 100
	}
	if limit < 1 {
		limit = 10
	}

	// محاسبه offset
	offset := (page - 1) * limit

	var blockedIPs []models.BlockedIP
	var total int64

	// شمارش کل رکوردها
	if err := h.db.Model(&models.BlockedIP{}).
		Where("is_active = ?", true).
		Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در شمارش IP های مسدود شده"})
		return
	}

	// دریافت IP های مسدود شده با صفحه‌بندی
	if err := h.db.Where("is_active = ?", true).
		Order("blocked_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&blockedIPs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت لیست IP های مسدود شده"})
		return
	}

	// محاسبه اطلاعات صفحه‌بندی
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   blockedIPs,
		"pagination": gin.H{
			"current_page":   page,
			"total_pages":    totalPages,
			"total_items":    total,
			"items_per_page": limit,
			"has_next_page":  page < totalPages,
			"has_prev_page":  page > 1,
		},
	})
}

// CleanupExpiredSessions پاکسازی سشن‌های منقضی شده
func (h *TrafficMonitoringHandler) CleanupExpiredSessions(c *gin.Context) {
	if err := h.svc.CleanupExpiredSessions(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در پاکسازی سشن‌ها"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "پاکسازی با موفقیت انجام شد"})
}
