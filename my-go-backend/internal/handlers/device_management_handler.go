package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
)

type DeviceManagementHandler struct {
	DB *gorm.DB
}

func NewDeviceManagementHandler(db *gorm.DB) *DeviceManagementHandler {
	return &DeviceManagementHandler{DB: db}
}

// GetUserDevices - دریافت لیست دستگاه‌های لاگین شده کاربر
// @route GET /api/user/devices
// Response: لیست دستگاه‌های فعال و غیرفعال کاربر با جزئیات کامل
func (h *DeviceManagementHandler) GetUserDevices(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", "لطفاً وارد شوید", nil))
		return
	}

	var sessions []models.UserSession
	err := h.DB.Where("user_id = ?", userID).
		Order("is_current DESC, last_activity DESC, created_at DESC").
		Find(&sessions).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DATABASE_ERROR", "خطا در دریافت اطلاعات", err.Error()))
		return
	}

	// آماده‌سازی response با اطلاعات کامل
	devices := make([]gin.H, 0, len(sessions))
	currentSessionToken, _ := c.Get("session_token")

	for _, session := range sessions {
		// تشخیص session فعلی
		isCurrent := session.SessionToken == currentSessionToken

		device := gin.H{
			"id":              session.ID,
			"device_name":     session.DeviceName,
			"platform":        session.Platform,
			"browser":         session.Browser,
			"browser_version": session.BrowserVersion,
			"os_name":         session.OSName,
			"os_version":      session.OSVersion,
			"ip_address":      session.IPAddress,
			"city":            session.City,
			"country":         session.Country,
			"timezone":        session.Timezone,
			"auth_method":     session.AuthMethod,
			"is_active":       session.IsActive,
			"is_trusted":      session.IsTrusted,
			"is_current":      isCurrent,
			"expires_at":      session.ExpiresAt,
			"last_activity":   session.LastActivity,
			"created_at":      session.CreatedAt,
			"device_info":     session.DeviceInfo,
		}

		// اضافه کردن وضعیت انقضا
		if session.ExpiresAt.Before(time.Now()) {
			device["status"] = "expired"
		} else if session.IsActive {
			device["status"] = "active"
		} else {
			device["status"] = "inactive"
		}

		devices = append(devices, device)
	}

	// آمار کلی
	stats := gin.H{
		"total_devices":   len(sessions),
		"active_devices":  0,
		"trusted_devices": 0,
		"current_device":  nil,
	}

	for _, session := range sessions {
		if session.IsActive && session.ExpiresAt.After(time.Now()) {
			stats["active_devices"] = stats["active_devices"].(int) + 1
		}
		if session.IsTrusted {
			stats["trusted_devices"] = stats["trusted_devices"].(int) + 1
		}
		if session.SessionToken == currentSessionToken {
			stats["current_device"] = session.ID
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"devices": devices,
		"stats":   stats,
	})
}

// LogoutDevice - خروج از دستگاه خاص
// @route DELETE /api/user/devices/:id
// Body: {} (اختیاری)
func (h *DeviceManagementHandler) LogoutDevice(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", "لطفاً وارد شوید", nil))
		return
	}

	deviceID := c.Param("id")
	if deviceID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه دستگاه الزامی است", nil))
		return
	}

	currentSessionToken, _ := c.Get("session_token")

	// پیدا کردن session مورد نظر
	var session models.UserSession
	err := h.DB.Where("id = ? AND user_id = ?", deviceID, userID).First(&session).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("DEVICE_NOT_FOUND", "دستگاه یافت نشد", nil))
		return
	}

	// جلوگیری از خروج از دستگاه فعلی
	if session.SessionToken == currentSessionToken {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "نمی‌توانید از دستگاه فعلی خارج شوید", nil))
		return
	}

	// غیرفعال کردن session
	err = h.DB.Model(&session).Updates(map[string]interface{}{
		"is_active":     false,
		"last_activity": time.Now(),
	}).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DATABASE_ERROR", "خطا در خروج از دستگاه", err.Error()))
		return
	}

	// لاگ امنیتی
	authService := services.NewAuthService(h.DB)
	authService.LogAuthEvent(userID.(uint), "device_logout", gin.H{
		"device_id":   session.ID,
		"device_name": session.DeviceName,
		"platform":    session.Platform,
		"ip_address":  session.IPAddress,
	}, c.ClientIP(), c.GetHeader("User-Agent"))

	c.JSON(http.StatusOK, gin.H{
		"message": "با موفقیت از دستگاه خارج شدید",
	})
}

// LogoutAllDevices - خروج از همه دستگاه‌ها (به جز فعلی)
// @route POST /api/user/devices/logout-all
func (h *DeviceManagementHandler) LogoutAllDevices(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", "لطفاً وارد شوید", nil))
		return
	}

	currentSessionToken, _ := c.Get("session_token")

	// غیرفعال کردن همه session ها به جز فعلی
	result := h.DB.Model(&models.UserSession{}).
		Where("user_id = ? AND session_token != ? AND is_active = ?", userID, currentSessionToken, true).
		Updates(map[string]interface{}{
			"is_active":     false,
			"last_activity": time.Now(),
		})

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DATABASE_ERROR", "خطا در خروج از دستگاه‌ها", result.Error.Error()))
		return
	}

	// لاگ امنیتی
	authService := services.NewAuthService(h.DB)
	authService.LogAuthEvent(userID.(uint), "logout_all_devices", gin.H{
		"affected_count": result.RowsAffected,
	}, c.ClientIP(), c.GetHeader("User-Agent"))

	c.JSON(http.StatusOK, gin.H{
		"message":        "با موفقیت از همه دستگاه‌ها خارج شدید",
		"affected_count": result.RowsAffected,
	})
}

// TrustDevice - اعتماد به دستگاه
// @route POST /api/user/devices/:id/trust
func (h *DeviceManagementHandler) TrustDevice(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", "لطفاً وارد شوید", nil))
		return
	}

	deviceID := c.Param("id")
	if deviceID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه دستگاه الزامی است", nil))
		return
	}

	// پیدا کردن session مورد نظر
	var session models.UserSession
	err := h.DB.Where("id = ? AND user_id = ?", deviceID, userID).First(&session).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("DEVICE_NOT_FOUND", "دستگاه یافت نشد", nil))
		return
	}

	// تغییر وضعیت اعتماد
	err = h.DB.Model(&session).Update("is_trusted", true).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DATABASE_ERROR", "خطا در تنظیم اعتماد", err.Error()))
		return
	}

	// لاگ امنیتی
	authService := services.NewAuthService(h.DB)
	authService.LogAuthEvent(userID.(uint), "device_trusted", gin.H{
		"device_id":   session.ID,
		"device_name": session.DeviceName,
	}, c.ClientIP(), c.GetHeader("User-Agent"))

	c.JSON(http.StatusOK, gin.H{
		"message": "دستگاه به عنوان مورد اعتماد علامت‌گذاری شد",
	})
}

// UntrustDevice - حذف اعتماد از دستگاه
// @route DELETE /api/user/devices/:id/trust
func (h *DeviceManagementHandler) UntrustDevice(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", "لطفاً وارد شوید", nil))
		return
	}

	deviceID := c.Param("id")
	deviceIDInt, err := strconv.Atoi(deviceID)
	if err != nil || deviceIDInt <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه دستگاه نامعتبر است", nil))
		return
	}

	// پیدا کردن session مورد نظر
	var session models.UserSession
	err = h.DB.Where("id = ? AND user_id = ?", deviceIDInt, userID).First(&session).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("DEVICE_NOT_FOUND", "دستگاه یافت نشد", nil))
		return
	}

	// حذف اعتماد
	err = h.DB.Model(&session).Update("is_trusted", false).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DATABASE_ERROR", "خطا در حذف اعتماد", err.Error()))
		return
	}

	// لاگ امنیتی
	authService := services.NewAuthService(h.DB)
	authService.LogAuthEvent(userID.(uint), "device_untrusted", gin.H{
		"device_id":   session.ID,
		"device_name": session.DeviceName,
	}, c.ClientIP(), c.GetHeader("User-Agent"))

	c.JSON(http.StatusOK, gin.H{
		"message": "اعتماد از دستگاه حذف شد",
	})
}
