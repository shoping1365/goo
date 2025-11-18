package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

// RecentViewHandler handles recent view operations
type RecentViewHandler struct {
	db *gorm.DB
}

// NewRecentViewHandler creates a new recent view handler
func NewRecentViewHandler(db *gorm.DB) *RecentViewHandler {
	return &RecentViewHandler{db: db}
}

// AddRecentView adds a product to user's recent views
func (h *RecentViewHandler) AddRecentView(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}

	productIDStr := c.Param("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه محصول نامعتبر است"})
		return
	}

	// Check if product exists
	var product models.Product
	if err := h.db.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "محصول یافت نشد"})
		return
	}

	// Parse User-Agent and get device info
	userAgent := c.GetHeader("User-Agent")
	deviceInfo := utils.ParseUserAgent(userAgent)

	// Get client IP
	clientIP := c.ClientIP()

	// Determine device type
	var deviceType string
	if deviceInfo.IsMobile {
		deviceType = "Mobile"
	} else if deviceInfo.IsTablet {
		deviceType = "Tablet"
	} else {
		deviceType = "Desktop"
	}

	// DEBUG: Log the parsed device info
	fmt.Printf("🔍 DEBUG - User-Agent: %s\n", userAgent)
	fmt.Printf("🔍 DEBUG - Browser: %s %s\n", deviceInfo.Browser, deviceInfo.BrowserVersion)
	fmt.Printf("🔍 DEBUG - OS: %s %s\n", deviceInfo.OSName, deviceInfo.OSVersion)
	fmt.Printf("🔍 DEBUG - Device: %s - %s\n", deviceType, deviceInfo.DeviceName)
	fmt.Printf("🔍 DEBUG - IP: %s\n", clientIP)

	// Check if already viewed recently
	var existingView models.RecentView
	err = h.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&existingView).Error

	if err == nil {
		// Update existing view: increment view_count and update timestamp + device info
		h.db.Model(&existingView).Updates(map[string]interface{}{
			"viewed_at":       time.Now(),
			"view_count":      existingView.ViewCount + 1,
			"last_updated_at": time.Now(),
			"browser":         deviceInfo.Browser,
			"browser_version": deviceInfo.BrowserVersion,
			"os":              deviceInfo.OSName,
			"os_version":      deviceInfo.OSVersion,
			"device_type":     deviceType,
			"device_model":    deviceInfo.DeviceName,
			"ip_address":      clientIP,
			"user_agent":      userAgent,
		})
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new recent view with device info
		recentView := models.RecentView{
			UserID:          userID.(uint),
			ProductID:       productID,
			ViewedAt:        time.Now(),
			DurationSeconds: 0,
			ViewCount:       1,
			LastUpdatedAt:   time.Now(),
			Browser:         deviceInfo.Browser,
			BrowserVersion:  deviceInfo.BrowserVersion,
			OS:              deviceInfo.OSName,
			OSVersion:       deviceInfo.OSVersion,
			DeviceType:      deviceType,
			DeviceModel:     deviceInfo.DeviceName,
			IPAddress:       clientIP,
			UserAgent:       userAgent,
		}
		if err := h.db.Create(&recentView).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ثبت بازدید"})
			return
		}

		// Keep only last 30 views per user
		h.cleanupOldViews(userID.(uint))
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بررسی بازدید"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "بازدید ثبت شد"})
}

// GetRecentViews returns user's recent views (last 30)
func (h *RecentViewHandler) GetRecentViews(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}

	var recentViews []models.RecentView
	err := h.db.Preload("Product").
		Where("user_id = ?", userID).
		Order("viewed_at DESC").
		Limit(30).
		Find(&recentViews).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت بازدیدهای اخیر"})
		return
	}

	// Convert to response format
	var responses []models.RecentViewResponse
	for _, rv := range recentViews {
		response := models.RecentViewResponse{
			ID:              rv.ID,
			ProductID:       rv.ProductID,
			ViewedAt:        rv.ViewedAt,
			DurationSeconds: rv.DurationSeconds,
			ViewCount:       rv.ViewCount,
			LastUpdatedAt:   rv.LastUpdatedAt,
			Browser:         rv.Browser,
			BrowserVersion:  rv.BrowserVersion,
			OS:              rv.OS,
			OSVersion:       rv.OSVersion,
			DeviceType:      rv.DeviceType,
			DeviceModel:     rv.DeviceModel,
			IPAddress:       rv.IPAddress,
			UserAgent:       rv.UserAgent,
		}
		response.Product.ID = rv.Product.ID
		response.Product.UUID = rv.Product.UUID
		response.Product.Slug = rv.Product.Slug
		response.Product.Name = rv.Product.Name
		response.Product.Image = rv.Product.Image
		response.Product.ImageURL = rv.Product.ImageURL
		response.Product.Price = rv.Product.Price
		response.Product.OldPrice = rv.Product.OldPrice

		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  responses,
		"count": len(responses),
	})
}

// DeleteRecentView removes a specific recent view
func (h *RecentViewHandler) DeleteRecentView(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}

	viewIDStr := c.Param("view_id")
	viewID, err := strconv.ParseUint(viewIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه بازدید نامعتبر است"})
		return
	}

	result := h.db.Where("id = ? AND user_id = ?", viewID, userID).Delete(&models.RecentView{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در حذف بازدید"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "بازدید یافت نشد"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "بازدید حذف شد"})
}

// ClearAllRecentViews removes all recent views for a user
func (h *RecentViewHandler) ClearAllRecentViews(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}

	if err := h.db.Where("user_id = ?", userID).Delete(&models.RecentView{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در پاک کردن بازدیدها"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "تمام بازدیدهای اخیر پاک شدند"})
}

// UpdateDuration updates the duration of a product view
func (h *RecentViewHandler) UpdateDuration(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}

	productIDStr := c.Param("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه محصول نامعتبر است"})
		return
	}

	var req models.UpdateDurationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "مدت زمان نامعتبر است (باید بین 1 تا 3600 ثانیه باشد)"})
		return
	}

	// Find the recent view record
	var recentView models.RecentView
	err = h.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&recentView).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "بازدید یافت نشد"})
		return
	}

	// Update duration
	if err := h.db.Model(&recentView).Update("duration_seconds", req.Duration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بروزرسانی مدت زمان"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "مدت زمان بروزرسانی شد",
		"duration": req.Duration,
	})
}

// GetUserRecentViewsAdmin returns recent views for a specific user (admin only)
func (h *RecentViewHandler) GetUserRecentViewsAdmin(c *gin.Context) {
	targetUserIDStr := c.Param("user_id")
	targetUserID, err := strconv.ParseUint(targetUserIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه کاربر نامعتبر است"})
		return
	}

	// Get limit from query params (default 50, max 200)
	limit := 50
	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 200 {
			limit = l
		}
	}

	var recentViews []models.RecentView
	err = h.db.Preload("Product").
		Where("user_id = ?", targetUserID).
		Order("viewed_at DESC").
		Limit(limit).
		Find(&recentViews).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت بازدیدهای کاربر"})
		return
	}

	// Convert to response format
	var responses []models.RecentViewResponse
	for _, rv := range recentViews {
		response := models.RecentViewResponse{
			ID:              rv.ID,
			ProductID:       rv.ProductID,
			ViewedAt:        rv.ViewedAt,
			DurationSeconds: rv.DurationSeconds,
			ViewCount:       rv.ViewCount,
			LastUpdatedAt:   rv.LastUpdatedAt,
			Browser:         rv.Browser,
			BrowserVersion:  rv.BrowserVersion,
			OS:              rv.OS,
			OSVersion:       rv.OSVersion,
			DeviceType:      rv.DeviceType,
			DeviceModel:     rv.DeviceModel,
			IPAddress:       rv.IPAddress,
			UserAgent:       rv.UserAgent,
		}
		response.Product.ID = rv.Product.ID
		response.Product.UUID = rv.Product.UUID
		response.Product.Slug = rv.Product.Slug
		response.Product.Name = rv.Product.Name
		response.Product.Image = rv.Product.Image
		response.Product.ImageURL = rv.Product.ImageURL
		response.Product.Price = rv.Product.Price
		response.Product.OldPrice = rv.Product.OldPrice

		responses = append(responses, response)
	}

	// Calculate some analytics
	totalDuration := 0
	for _, rv := range recentViews {
		totalDuration += rv.DurationSeconds
	}

	c.JSON(http.StatusOK, gin.H{
		"data":           responses,
		"count":          len(responses),
		"total_duration": totalDuration,
		"avg_duration":   float64(totalDuration) / float64(len(responses)),
	})
}

// cleanupOldViews keeps only the last 30 views for a user
func (h *RecentViewHandler) cleanupOldViews(userID uint) {
	var count int64
	h.db.Model(&models.RecentView{}).Where("user_id = ?", userID).Count(&count)

	if count > 30 {
		// Get IDs of views to delete (keep only last 30)
		var oldViews []models.RecentView
		h.db.Where("user_id = ?", userID).
			Order("viewed_at DESC").
			Offset(30).
			Find(&oldViews)

		var idsToDelete []uint
		for _, view := range oldViews {
			idsToDelete = append(idsToDelete, view.ID)
		}

		if len(idsToDelete) > 0 {
			h.db.Delete(&models.RecentView{}, idsToDelete)
		}
	}
}

// CleanupOldViewsForAllUsers removes views older than 90 days (for data retention policy)
func (h *RecentViewHandler) CleanupOldViewsForAllUsers() {
	// Delete views older than 90 days
	h.db.Where("viewed_at < ?", time.Now().AddDate(0, 0, -90)).Delete(&models.RecentView{})
}

// CleanupUnknownViews removes views that don't have device information (old records)
func (h *RecentViewHandler) CleanupUnknownViews(c *gin.Context) {
	result := h.db.Where("browser = ? OR device_type = ? OR os = ?", "Unknown", "Unknown", "Unknown").Delete(&models.RecentView{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در پاک کردن بازدیدهای قدیمی"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "بازدیدهای قدیمی با موفقیت پاک شدند",
		"deleted_count": result.RowsAffected,
	})
}
