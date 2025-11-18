package handlers

import (
	"math"
	"net/http"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SystemSecurityLoginHandler هندلر صفحهٔ مانیتورینگ ورودها
type SystemSecurityLoginHandler struct {
	db  *gorm.DB
	svc *services.SecurityLoginService
}

func NewSystemSecurityLoginHandler(db *gorm.DB) *SystemSecurityLoginHandler {
	uowFactory := unitofwork.NewUnitOfWorkFactory(db)
	svc := services.NewSecurityLoginService(uowFactory)
	return &SystemSecurityLoginHandler{db: db, svc: svc}
}

// GetCounters آمار کارت‌های بالای صفحه
func (h *SystemSecurityLoginHandler) GetCounters(c *gin.Context) {
	succ, fail, susp, online, err := h.svc.GetCounters(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("INTERNAL_ERROR", utils.GetErrorMessage("INTERNAL_ERROR"), err.Error()))
		return
	}
	// آنلاین را از ترافیک (اختیاری) صفر می‌گذاریم؛ صفحه می‌تواند جداگانه از /api/admin/traffic/stats بخواند
	c.JSON(http.StatusOK, gin.H{
		"successful": succ,
		"failed":     fail,
		"suspicious": susp,
		"online":     online,
	})
}

// ListHistory فهرست تاریخچه ورود با فیلتر/صفحه‌بندی
func (h *SystemSecurityLoginHandler) ListHistory(c *gin.Context) {
	page := utils.ParseInt(c.DefaultQuery("page", "1"))
	limit := utils.ParseInt(c.DefaultQuery("limit", "10"))
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	filter := repository.LoginAttemptFilter{
		Search:      c.Query("user"),
		AttemptType: c.Query("type"),
		DateFrom:    c.Query("date"),
	}
	// نگاشت وضعیت
	switch c.Query("status") {
	case "successful":
		b := true
		filter.IsSuccessful = &b
	case "failed", "suspicious":
		b := false
		filter.IsSuccessful = &b
	}

	items, total, err := h.svc.ListLoginAttempts(c, filter, page, limit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("INTERNAL_ERROR", utils.GetErrorMessage("INTERNAL_ERROR"), err.Error()))
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	c.JSON(http.StatusOK, gin.H{
		"data":  items,
		"total": total,
		"page":  page,
		"limit": limit,
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

// BlockIP مسدود کردن IP از صفحه
func (h *SystemSecurityLoginHandler) BlockIP(c *gin.Context) {
	var req struct {
		IPAddress string `json:"ip_address" binding:"required"`
		Reason    string `json:"reason"`
		Duration  int    `json:"duration"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	blockedBy := "system"
	if userName, ok := c.Get("user_name"); ok {
		if s, ok2 := userName.(string); ok2 && s != "" {
			blockedBy = s
		}
	}

	if err := h.svc.BlockIP(c, req.IPAddress, req.Reason, blockedBy, req.Duration); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// UnblockIP آزادسازی IP
func (h *SystemSecurityLoginHandler) UnblockIP(c *gin.Context) {
	ip := c.Param("ip")
	if ip == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), "ip is required"))
		return
	}
	if err := h.svc.UnblockIP(c, ip); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
