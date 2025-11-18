package handlers

import (
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// FirewallHandler هندلر مدیریت فایروال سیستم‌عامل
// تمامی منطق بیزینسی در سرویس قرار دارد
type FirewallHandler struct {
	svc *services.FirewallService
}

func NewFirewallHandler(svc *services.FirewallService) *FirewallHandler {
	return &FirewallHandler{svc: svc}
}

// GetStatus دریافت وضعیت فعلی فایروال + حالت امنیتی
// @Summary وضعیت فایروال
// @Tags security_firewall
// @Produce json
func (h *FirewallHandler) GetStatus(c *gin.Context) {
	status, err := h.svc.GetStatus(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در دریافت وضعیت فایروال"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": status})
}

// Toggle فعال/غیرفعال کردن فایروال
func (h *FirewallHandler) Toggle(c *gin.Context) {
	var body struct {
		Enabled bool `json:"enabled"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "درخواست نامعتبر"})
		return
	}
	if err := h.svc.Toggle(c.Request.Context(), body.Enabled); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در به‌روزرسانی وضعیت فایروال"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "وضعیت فایروال بروزرسانی شد"})
}

// SetMode تغییر حالت امنیتی
func (h *FirewallHandler) SetMode(c *gin.Context) {
	var body struct {
		Mode string `json:"mode"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "درخواست نامعتبر"})
		return
	}
	mode := repository.FirewallMode(strings.ToLower(body.Mode))
	if err := h.svc.SetMode(c.Request.Context(), mode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "حالت نامعتبر یا خطا در تنظیم"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "حالت امنیتی بروزرسانی شد"})
}

// ListRules فهرست قوانین فعلی سیستم عامل
func (h *FirewallHandler) ListRules(c *gin.Context) {
	rules, err := h.svc.ListRules(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در دریافت قوانین"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rules})
}

// CreateRule افزودن قانون جدید
func (h *FirewallHandler) CreateRule(c *gin.Context) {
	var body repository.CreateFirewallRule
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "درخواست نامعتبر"})
		return
	}
	if err := validateRuleInput(body.Name, body.Direction, body.Action, body.Source, body.Destination, body.Port); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.svc.CreateRule(c.Request.Context(), body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در ایجاد قانون"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "قانون ایجاد شد"})
}

// UpdateRule بروزرسانی قانون موجود (با نام)
func (h *FirewallHandler) UpdateRule(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "نام قانون الزامی است"})
		return
	}
	var body repository.UpdateFirewallRule
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "درخواست نامعتبر"})
		return
	}
	// validate fields if present
	dir := ""
	if body.Direction != nil {
		dir = *body.Direction
	}
	act := ""
	if body.Action != nil {
		act = *body.Action
	}
	src := ""
	if body.Source != nil {
		src = *body.Source
	}
	dst := ""
	if body.Destination != nil {
		dst = *body.Destination
	}
	port := ""
	if body.Port != nil {
		port = *body.Port
	}
	if err := validateRuleInput(name, dir, act, src, dst, port); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.svc.UpdateRule(c.Request.Context(), name, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در بروزرسانی قانون"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "قانون بروزرسانی شد"})
}

// ToggleRule فعال/غیرفعال‌سازی قانون موجود
func (h *FirewallHandler) ToggleRule(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "نام قانون الزامی است"})
		return
	}
	enabledStr := c.Query("enabled")
	enabled := enabledStr == "1" || strings.EqualFold(enabledStr, "true")
	if err := h.svc.ToggleRule(c.Request.Context(), name, enabled); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در تغییر وضعیت قانون"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "وضعیت قانون تغییر کرد"})
}

// DeleteRule حذف قانون
func (h *FirewallHandler) DeleteRule(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "نام قانون الزامی است"})
		return
	}
	if err := h.svc.DeleteRule(c.Request.Context(), name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در حذف قانون"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "قانون حذف شد"})
}

// ReadLogs خواندن لاگ‌های فایروال
func (h *FirewallHandler) ReadLogs(c *gin.Context) {
	limit := 200
	if l := c.Query("limit"); l != "" {
		if n, err := strconv.Atoi(l); err == nil {
			limit = n
		}
	}
	logs, err := h.svc.ReadLogs(c.Request.Context(), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در دریافت لاگ‌ها"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": logs})
}

// ClearLogs پاکسازی لاگ‌ها
func (h *FirewallHandler) ClearLogs(c *gin.Context) {
	if err := h.svc.ClearLogs(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در پاک‌سازی لاگ‌ها"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "لاگ‌ها پاک شدند"})
}

// BackupRules پشتیبان قوانین
func (h *FirewallHandler) BackupRules(c *gin.Context) {
	dest := c.Query("dest")
	path, err := h.svc.Backup(c.Request.Context(), dest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در پشتیبان‌گیری"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "پشتیبان‌گیری انجام شد", "path": path})
}

// RestoreRules بازگردانی قوانین از فایل wfw
func (h *FirewallHandler) RestoreRules(c *gin.Context) {
	var body struct {
		Path string `json:"path"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || strings.TrimSpace(body.Path) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "مسیر فایل نامعتبر است"})
		return
	}
	if err := h.svc.Restore(c.Request.Context(), body.Path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در بازگردانی"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "بازگردانی انجام شد"})
}

// GetLoggingConfig دریافت تنظیمات لاگ فایروال
func (h *FirewallHandler) GetLoggingConfig(c *gin.Context) {
	cfg, err := h.svc.GetLogging(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در دریافت تنظیمات لاگ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cfg})
}

// SetLoggingConfig تنظیم لاگ فایروال
func (h *FirewallHandler) SetLoggingConfig(c *gin.Context) {
	var cfg repository.FirewallLoggingConfig
	if err := c.ShouldBindJSON(&cfg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "درخواست نامعتبر"})
		return
	}
	if err := h.svc.SetLogging(c.Request.Context(), cfg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "خطا در ذخیره تنظیمات لاگ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "تنظیمات لاگ ذخیره شد"})
}

// validateRuleInput: اعتبارسنجی ورودی قانون فایروال
func validateRuleInput(name, direction, action, source, destination, port string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("نام قانون الزامی است")
	}
	if direction != "" && direction != "inbound" && direction != "outbound" {
		return fmt.Errorf("جهت قانون نامعتبر است")
	}
	if action != "" && action != "allow" && action != "deny" && action != "drop" {
		return fmt.Errorf("عملیات قانون نامعتبر است")
	}
	if source != "" && !isValidIPorCIDR(source) && source != "any" {
		return fmt.Errorf("IP منبع نامعتبر است")
	}
	if destination != "" && !isValidIPorCIDR(destination) && destination != "any" {
		return fmt.Errorf("IP مقصد نامعتبر است")
	}
	if port != "" && port != "any" {
		portRe := regexp.MustCompile(`^(\d{1,5})(-(\d{1,5}))?(,(\d{1,5})(-(\d{1,5}))?)*$`)
		if !portRe.MatchString(port) {
			return fmt.Errorf("فرمت پورت نامعتبر است")
		}
	}
	return nil
}

func isValidIPorCIDR(s string) bool {
	if ip := net.ParseIP(s); ip != nil {
		return true
	}
	if strings.Contains(s, "/") {
		_, _, err := net.ParseCIDR(s)
		return err == nil
	}
	return false
}
