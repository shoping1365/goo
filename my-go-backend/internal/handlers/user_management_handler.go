package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"
)

type UserManagementHandler struct {
	db              *gorm.DB
	authService     *services.AuthService
	settingsService *services.SettingsService
	// UoW و سرویس مانیتورینگ ورود
	uowFactory   unitofwork.UnitOfWorkFactory
	loginService *services.LoginMonitoringService
}

func NewUserManagementHandler(db *gorm.DB) *UserManagementHandler {
	uowf := unitofwork.NewUnitOfWorkFactory(db)
	return &UserManagementHandler{
		db:              db,
		authService:     services.NewAuthService(db),
		settingsService: services.NewSettingsService(db),
		uowFactory:      uowf,
		loginService:    services.NewLoginMonitoringService(uowf),
	}
}

// GetUsers دریافت لیست کاربران
func (h *UserManagementHandler) GetUsers(c *gin.Context) {
	var users []models.User
	var total int64

	// پارامترهای فیلتر
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	role := c.Query("role")
	status := c.Query("status")
	search := c.Query("search")

	query := h.db.Model(&models.User{})

	// اعمال فیلترها
	if role != "" {
		query = query.Where("role = ?", role)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if search != "" {
		query = query.Where("name ILIKE ? OR mobile ILIKE ? OR email ILIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// شمارش کل
	query.Count(&total)

	// دریافت کاربران
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت کاربران"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// GetUser دریافت اطلاعات یک کاربر
func (h *UserManagementHandler) GetUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه کاربر نامعتبر است"})
		return
	}

	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "کاربر یافت نشد"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// BlockUser بلاک کردن کاربر
func (h *UserManagementHandler) BlockUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه کاربر نامعتبر است"})
		return
	}

	var req struct {
		Reason string `json:"reason" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "دلیل بلاک کردن الزامی است"})
		return
	}

	// بررسی وجود کاربر
	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "کاربر یافت نشد"})
		return
	}

	// بررسی بلاک بودن
	if user.Status == "blocked" || user.Status == "suspended" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کاربر قبلاً بلاک شده است"})
		return
	}

	// بلاک کردن کاربر
	adminID := c.GetUint("user_id")
	updates := map[string]interface{}{
		"is_blocked":   true,
		"blocked_at":   time.Now(),
		"blocked_by":   adminID,
		"block_reason": req.Reason,
	}

	if err := h.db.Model(&user).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بلاک کردن کاربر"})
		return
	}

	// ثبت فعالیت
	h.authService.LogUserActivity(adminID, "block_user",
		"کاربر "+user.Mobile+" بلاک شد. دلیل: "+req.Reason,
		c.ClientIP(), c.GetHeader("User-Agent"))

	// ثبت رویداد امنیتی
	h.createSecurityEvent(user.ID, "user_blocked", "high",
		"کاربر توسط ادمین بلاک شد. دلیل: "+req.Reason,
		c.ClientIP(), c.GetHeader("User-Agent"))

	c.JSON(http.StatusOK, gin.H{"message": "کاربر با موفقیت بلاک شد"})
}

// UnblockUser آنبلاک کردن کاربر
func (h *UserManagementHandler) UnblockUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه کاربر نامعتبر است"})
		return
	}

	// بررسی وجود کاربر
	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "کاربر یافت نشد"})
		return
	}

	// بررسی بلاک بودن
	if user.Status != "blocked" && user.Status != "suspended" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کاربر بلاک نیست"})
		return
	}

	// آنبلاک کردن کاربر
	updates := map[string]interface{}{
		"is_blocked":   false,
		"blocked_at":   nil,
		"blocked_by":   nil,
		"block_reason": nil,
	}

	if err := h.db.Model(&user).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در آنبلاک کردن کاربر"})
		return
	}

	// ثبت فعالیت
	adminID := c.GetUint("user_id")
	h.authService.LogUserActivity(adminID, "unblock_user",
		"کاربر "+user.Mobile+" آنبلاک شد",
		c.ClientIP(), c.GetHeader("User-Agent"))

	c.JSON(http.StatusOK, gin.H{"message": "کاربر با موفقیت آنبلاک شد"})
}

// GetUserLoginHistory دریافت تاریخچه ورود کاربر
func (h *UserManagementHandler) GetUserLoginHistory(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه کاربر نامعتبر است"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	var attempts []models.LoginAttempt
	var total int64

	query := h.db.Model(&models.LoginAttempt{}).Where("user_id = ?", userID)
	query.Count(&total)

	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&attempts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت تاریخچه"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attempts": attempts,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

// GetUserAnalytics دریافت آمار کاربر
func (h *UserManagementHandler) GetUserAnalytics(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه کاربر نامعتبر است"})
		return
	}

	// دریافت اطلاعات کاربر
	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "کاربر یافت نشد"})
		return
	}

	// آمار کاربر
	analytics := gin.H{
		"user_id": userID,
	}

	// شمارش تلاش‌های ورود از جدول auth_events
	var totalLogins, successfulLogins, failedLogins int64
	h.db.Model(&models.AuthEvent{}).Where("mobile = ?", user.Mobile).Count(&totalLogins)
	h.db.Model(&models.AuthEvent{}).Where("mobile = ? AND status = 'success'", user.Mobile).Count(&successfulLogins)
	h.db.Model(&models.AuthEvent{}).Where("mobile = ? AND status = 'failed'", user.Mobile).Count(&failedLogins)

	analytics["total_logins"] = totalLogins
	analytics["successful_logins"] = successfulLogins
	analytics["failed_logins"] = failedLogins

	// محاسبه نرخ موفقیت
	if totalLogins > 0 {
		analytics["success_rate"] = float64(successfulLogins) / float64(totalLogins) * 100
	} else {
		analytics["success_rate"] = 0
	}

	// آخرین ورود
	var lastLoginEvent models.AuthEvent
	if err := h.db.Where("mobile = ? AND status = 'success'", user.Mobile).Order("created_at DESC").First(&lastLoginEvent).Error; err == nil {
		analytics["last_login_at"] = lastLoginEvent.CreatedAt
	}

	// نشست‌های فعال
	var activeSessions int64
	h.db.Model(&models.UserSession{}).Where("user_id = ? AND is_active = true", userID).Count(&activeSessions)
	analytics["active_sessions"] = activeSessions

	// رویدادهای امنیتی
	var securityEvents, unresolvedEvents int64
	h.db.Model(&models.AuthEvent{}).Where("user_id = ? AND severity IN ('warning', 'error', 'critical')", userID).Count(&securityEvents)
	h.db.Model(&models.AuthEvent{}).Where("user_id = ? AND severity = 'error'", userID).Count(&unresolvedEvents)
	analytics["security_events"] = securityEvents
	analytics["unresolved_events"] = unresolvedEvents

	c.JSON(http.StatusOK, analytics)
}

// GetLoginHistory دریافت تاریخچه ورود
// این متد صرفاً خواندنی است و از سرویس استفاده می‌کند تا وابستگی‌های دیتابیس از هندلر جدا بماند.
func (h *UserManagementHandler) GetLoginHistory(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	filter := repository.LoginAttemptFilter{
		Search:      c.Query("search"),
		AttemptType: c.Query("method"),
		DateFrom:    c.Query("dateFrom"),
		IP:          c.Query("ip"),
	}
	if success := c.Query("success"); success != "" {
		v := success == "true"
		filter.IsSuccessful = &v
	}

	attempts, total, err := h.loginService.ListHistory(c.Request.Context(), filter, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت تاریخچه"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attempts": attempts,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

// GetLoginStats دریافت آمار ورود
// شامل تعداد ورودهای موفق/ناموفق، نرخ موفقیت و تعداد IP های مشکوک و کاربران آنلاین
func (h *UserManagementHandler) GetLoginStats(c *gin.Context) {
	stats, err := h.loginService.GetStats(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت آمار"})
		return
	}
	c.JSON(http.StatusOK, stats)
}

// createSecurityEvent ایجاد رویداد امنیتی
func (h *UserManagementHandler) createSecurityEvent(userID uint, eventType, severity, details, ipAddress, userAgent string) {
	event := &models.AuthEvent{
		UserID:    &userID,
		EventType: eventType,
		Status:    "info",
		Severity:  severity,
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}

	h.db.Create(event)
}
