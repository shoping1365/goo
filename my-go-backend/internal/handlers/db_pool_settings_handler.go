package handlers

import (
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
)

// DBPoolSettingsPayload ساختار ورودی/خروجی تنظیمات کانکشن‌پول
type DBPoolSettingsPayload struct {
	AllocatedCores         int `json:"allocatedCores"`
	MaxIdleConns           int `json:"maxIdleConns"`
	MaxOpenConns           int `json:"maxOpenConns"`
	ConnMaxLifetimeMinutes int `json:"connMaxLifetimeMinutes"`
	ConnMaxIdleTimeMinutes int `json:"connMaxIdleTimeMinutes"`
}

// DBPoolSettingsResponse پاسخ GET شامل وضعیت فعلی، پیشنهادها و اطلاعات سیستم
type DBPoolSettingsResponse struct {
	Success bool   `json:"success"`
	Data    gin.H  `json:"data"`
	Message string `json:"message,omitempty"`
}

// DBPoolSettingsHandler هندلر مدیریت تنظیمات کانکشن‌پول دیتابیس
type DBPoolSettingsHandler struct {
	db             *gorm.DB
	settingService *services.SettingService
}

func NewDBPoolSettingsHandler(db *gorm.DB, settingService *services.SettingService) *DBPoolSettingsHandler {
	return &DBPoolSettingsHandler{db: db, settingService: settingService}
}

// GetSettings دریافت تنظیمات فعلی و پیشنهادها
func (h *DBPoolSettingsHandler) GetSettings(c *gin.Context) {
	ctx := c.Request.Context()

	// خواندن از settings با پیشوند db.pool.
	settings, _ := h.settingService.GetSettingsByKeyPrefix(ctx, "db.pool.")
	kv := map[string]string{}
	for _, s := range settings {
		kv[s.Key] = s.Value
	}

	toInt := func(s string, def int) int {
		if s == "" {
			return def
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			return def
		}
		return v
	}

	payload := DBPoolSettingsPayload{
		AllocatedCores:         toInt(kv["db.pool.allocated_cores"], 0),
		MaxIdleConns:           toInt(kv["db.pool.max_idle_conns"], 0),
		MaxOpenConns:           toInt(kv["db.pool.max_open_conns"], 0),
		ConnMaxLifetimeMinutes: toInt(kv["db.pool.conn_max_lifetime_minutes"], 60),
		ConnMaxIdleTimeMinutes: toInt(kv["db.pool.conn_max_idle_time_minutes"], 0),
	}

	systemCores := runtime.NumCPU()
	baseCores := payload.AllocatedCores
	if baseCores <= 0 {
		baseCores = systemCores
	}

	recommended := gin.H{
		"maxIdleConns":           baseCores,
		"maxOpenConns":           2 * baseCores,
		"connMaxLifetimeMinutes": 60,
		"connMaxIdleTimeMinutes": 30,
	}

	c.JSON(http.StatusOK, DBPoolSettingsResponse{
		Success: true,
		Data: gin.H{
			"current":     payload,
			"recommended": recommended,
			"system": gin.H{
				"cpuCores": systemCores,
			},
		},
	})
}

// UpdateSettings بروزرسانی و اعمال تنظیمات کانکشن‌پول
func (h *DBPoolSettingsHandler) UpdateSettings(c *gin.Context) {
	var p DBPoolSettingsPayload
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "ورودی نامعتبر"})
		return
	}

	// اعمال بر روی sqlDB
	sqlDB, err := h.db.DB()
	if err == nil {
		if p.MaxIdleConns > 0 {
			sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		}
		if p.MaxOpenConns > 0 {
			sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		}
		if p.ConnMaxLifetimeMinutes > 0 {
			sqlDB.SetConnMaxLifetime(time.Duration(p.ConnMaxLifetimeMinutes) * time.Minute)
		}
		if p.ConnMaxIdleTimeMinutes > 0 {
			// از Go 1.15+ در دسترس است
			sqlDB.SetConnMaxIdleTime(time.Duration(p.ConnMaxIdleTimeMinutes) * time.Minute)
		}
	}

	// ذخیره در جدول settings برای پایداری
	ctx := c.Request.Context()
	toSetting := func(key string, val string) models.Setting {
		return models.Setting{
			Key:      key,
			Value:    val,
			Category: "db.pool",
			Type:     "number",
			IsPublic: false,
		}
	}
	newSettings := []models.Setting{
		toSetting("db.pool.allocated_cores", strconv.Itoa(p.AllocatedCores)),
		toSetting("db.pool.max_idle_conns", strconv.Itoa(p.MaxIdleConns)),
		toSetting("db.pool.max_open_conns", strconv.Itoa(p.MaxOpenConns)),
		toSetting("db.pool.conn_max_lifetime_minutes", strconv.Itoa(p.ConnMaxLifetimeMinutes)),
		toSetting("db.pool.conn_max_idle_time_minutes", strconv.Itoa(p.ConnMaxIdleTimeMinutes)),
	}
	_ = h.settingService.UpdateSettings(ctx, newSettings)
	_ = h.settingService.ClearCache(ctx)

	c.JSON(http.StatusOK, gin.H{"success": true})
}
