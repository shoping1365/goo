package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// BotHandler هندلر تنظیمات و مدیریت ربات‌ها (بدون نیاز به مایگریشن جدید)
// داده‌ها در جدول settings با کلیدهای bots.* ذخیره می‌شوند.
type BotHandler struct {
	settings *services.SettingService
	botSvc   *services.BotService
}

func NewBotHandler(settings *services.SettingService, botSvc *services.BotService) *BotHandler {
	return &BotHandler{settings: settings, botSvc: botSvc}
}

// GetConfig دریافت پیکربندی فعلی ربات‌ها
func (h *BotHandler) GetConfig(c *gin.Context) {
	// خواندن چند کلید مرتبط
	keys := []string{
		"bots.enabled",
		"bots.global_mode",
		"bots.allowlist.user_agents",
		"bots.blocklist.user_agents",
		"bots.allowlist.ips",
		"bots.blocklist.ips",
		"bots.malicious.user_agents",
		// کلیدهای کنترلی جدید برای کاهش سربار
		"bots.exclude_paths",
		"bots.skip_authorized",
		"bots.skip_local_ips",
		"bots.skip_methods",
		"bots.monitor_only",
	}
	resp := gin.H{}
	for _, k := range keys {
		st, _ := h.settings.GetSetting(c, k)
		if st != nil {
			resp[k] = st.Value
		} else {
			resp[k] = ""
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": resp})
}

// UpdateConfig بروزرسانی پیکربندی ربات‌ها
func (h *BotHandler) UpdateConfig(c *gin.Context) {
	var payload map[string]string
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "داده نامعتبر"})
		return
	}
	var toUpdate []models.Setting
	for k, v := range payload {
		if !strings.HasPrefix(k, "bots.") {
			continue
		}
		toUpdate = append(toUpdate, models.Setting{Key: k, Value: v, Category: "bots", Type: "string"})
	}
	if len(toUpdate) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "هیچ کلید معتبری ارسال نشد"})
		return
	}
	if err := h.settings.UpdateSettings(c, toUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در بروزرسانی تنظیمات"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Helpers برای ساده‌سازی CSVها
func csvJoin(list []string) string { return strings.Join(list, ",") }

// ListBots فهرست ربات‌های شناسایی‌شده بر اساس UA pattern ها (از traffic_logs)
func (h *BotHandler) ListBots(c *gin.Context) {
	// اگر سپر ربات‌ها خاموش باشد، هیچ کوئری روی ترافیک اجرا نکن
	if st, _ := h.settings.GetSetting(c, "bots.enabled"); st != nil {
		v := strings.ToLower(strings.TrimSpace(st.Value))
		if v == "false" || v == "0" || v == "no" {
			page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
			limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
			c.JSON(http.StatusOK, gin.H{"status": "success", "data": []any{}, "pagination": gin.H{
				"current_page":   page,
				"items_per_page": limit,
				"total_items":    0,
			}})
			return
		}
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	patterns := c.Query("patterns")
	var pats []string
	if strings.TrimSpace(patterns) != "" {
		pats = strings.Split(patterns, ",")
	}
	logs, total, err := h.botSvc.ListKnownBots(c, pats, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در دریافت ربات‌ها"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": logs, "pagination": gin.H{
		"current_page":   page,
		"items_per_page": limit,
		"total_items":    total,
	}})
}

// ImportMaliciousPatterns ادغام لیست الگوهای مخرب دریافتی با تنظیمات موجود
func (h *BotHandler) ImportMaliciousPatterns(c *gin.Context) {
	var payload struct {
		Patterns []string `json:"patterns"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil || len(payload.Patterns) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "الگوهای نامعتبر"})
		return
	}
	key := "bots.malicious.user_agents"
	st, _ := h.settings.GetSetting(c, key)
	cur := ""
	if st != nil {
		cur = st.Value
	}
	merged := services.NormalizeCSV(cur, strings.Join(payload.Patterns, ","))
	if err := h.settings.UpdateSettings(c, []models.Setting{{Key: key, Value: merged, Category: "bots", Type: "string"}}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در بروزرسانی"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
