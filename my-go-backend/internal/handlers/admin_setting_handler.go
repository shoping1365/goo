package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
)

type AdminSettingHandler struct {
	db *gorm.DB
}

func NewAdminSettingHandler(db *gorm.DB) *AdminSettingHandler {
	return &AdminSettingHandler{db: db}
}

// GetAdminSetting تنظیمات ادمین را برای کاربر و کلید مشخص شده برمی‌گرداند
func (h *AdminSettingHandler) GetAdminSetting(c *gin.Context) {
	// دریافت user_id از context (از middleware auth)
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}
	userID := userIDInterface.(uint)

	// دریافت key از URL parameter
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کلید تنظیمات الزامی است"})
		return
	}

	var setting models.AdminSetting
	err := h.db.Where("user_id = ? AND key = ?", userID, key).First(&setting).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// اگر تنظیمات وجود نداشت، مقدار پیش‌فرض برمی‌گردانیم
			c.JSON(http.StatusOK, gin.H{
				"user_id": userID,
				"key":     key,
				"value":   "",
				"exists":  false,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت تنظیمات"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         setting.ID,
		"user_id":    setting.UserID,
		"key":        setting.Key,
		"value":      setting.Value,
		"exists":     true,
		"created_at": setting.CreatedAt,
		"updated_at": setting.UpdatedAt,
	})
}

// UpdateAdminSetting تنظیمات ادمین را برای کاربر و کلید مشخص شده به‌روزرسانی می‌کند
func (h *AdminSettingHandler) UpdateAdminSetting(c *gin.Context) {
	// دریافت user_id از context (از middleware auth)
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}
	userID := userIDInterface.(uint)

	// دریافت key از URL parameter
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کلید تنظیمات الزامی است"})
		return
	}

	// دریافت value از request body
	var requestBody struct {
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "مقدار تنظیمات الزامی است"})
		return
	}

	// بررسی وجود تنظیمات
	var existingSetting models.AdminSetting
	err := h.db.Where("user_id = ? AND key = ?", userID, key).First(&existingSetting).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بررسی تنظیمات موجود"})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// ایجاد تنظیمات جدید
		userIDPtr := &userID
		newSetting := models.AdminSetting{
			UserID: userIDPtr,
			Key:    key,
			Value:  requestBody.Value,
		}

		if err := h.db.Create(&newSetting).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ایجاد تنظیمات"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "تنظیمات با موفقیت ایجاد شد",
			"id":      newSetting.ID,
			"user_id": newSetting.UserID,
			"key":     newSetting.Key,
			"value":   newSetting.Value,
		})
		return
	}

	// به‌روزرسانی تنظیمات موجود
	existingSetting.Value = requestBody.Value
	if err := h.db.Save(&existingSetting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در به‌روزرسانی تنظیمات"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "تنظیمات با موفقیت به‌روزرسانی شد",
		"id":      existingSetting.ID,
		"user_id": existingSetting.UserID,
		"key":     existingSetting.Key,
		"value":   existingSetting.Value,
	})
}

// GetAllAdminSettings تمام تنظیمات ادمین کاربر را برمی‌گرداند
func (h *AdminSettingHandler) GetAllAdminSettings(c *gin.Context) {
	// دریافت user_id از context (از middleware auth)
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}
	userID := userIDInterface.(uint)

	var settings []models.AdminSetting
	err := h.db.Where("user_id = ?", userID).Find(&settings).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت تنظیمات"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":  userID,
		"settings": settings,
	})
}

// GetBulkEditColumns تنظیمات ستون‌های ویرایش انبوه محصول را برمی‌گرداند
func (h *AdminSettingHandler) GetBulkEditColumns(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}

	// رفع خطای type assertion: پشتیبانی از int، uint، string
	var userID uint
	switch v := userIDInterface.(type) {
	case uint:
		userID = v
	case int:
		userID = uint(v)
	case int64:
		userID = uint(v)
	case float64:
		userID = uint(v)
	case string:
		// تلاش برای تبدیل رشته به uint
		var parsed int
		if _, err := fmt.Sscanf(v, "%d", &parsed); err == nil {
			userID = uint(parsed)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id نامعتبر است"})
			return
		}
	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id نامعتبر است"})
		return
	}

	key := "bulk_edit_product_columns"

	var setting models.AdminSetting
	err := h.db.Where("user_id = ? AND key = ?", userID, key).First(&setting).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// اگر تنظیمات وجود نداشت، مقدار خالی برگردان تا فرانت پیش‌فرض‌ها را ست کند
			c.JSON(http.StatusOK, gin.H{"value": ""})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت تنظیمات ستون‌ها"})
		return
	}

	// اگر مقدار "[]" یا "null" ذخیره شده بود، آن را به خالی تبدیل کن تا UI پیش‌فرض‌ها را اعمال کند
	clean := strings.TrimSpace(setting.Value)
	if clean == "[]" || clean == "null" || clean == "" {
		c.JSON(http.StatusOK, gin.H{"value": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": setting.Value})
}

// UpdateBulkEditColumns تنظیمات ستون‌های ویرایش انبوه محصول را به‌روزرسانی می‌کند
func (h *AdminSettingHandler) UpdateBulkEditColumns(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}

	// رفع خطای type assertion: پشتیبانی از int، uint، string
	var userID uint
	switch v := userIDInterface.(type) {
	case uint:
		userID = v
	case int:
		userID = uint(v)
	case int64:
		userID = uint(v)
	case float64:
		userID = uint(v)
	case string:
		var parsed int
		if _, err := fmt.Sscanf(v, "%d", &parsed); err == nil {
			userID = uint(parsed)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id نامعتبر است"})
			return
		}
	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id نامعتبر است"})
		return
	}

	key := "bulk_edit_product_columns"

	var requestBody struct {
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "مقدار ستون‌ها الزامی است"})
		return
	}

	var existingSetting models.AdminSetting
	err := h.db.Where("user_id = ? AND key = ?", userID, key).First(&existingSetting).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در بررسی تنظیمات موجود"})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		userIDPtr := &userID
		newSetting := models.AdminSetting{
			UserID: userIDPtr,
			Key:    key,
			Value:  requestBody.Value,
		}

		if err := h.db.Create(&newSetting).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ایجاد تنظیمات"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "تنظیمات ستون‌ها با موفقیت ایجاد شد"})
		return
	}

	existingSetting.Value = requestBody.Value
	if err := h.db.Save(&existingSetting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در به‌روزرسانی تنظیمات"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "تنظیمات ستون‌ها با موفقیت به‌روزرسانی شد"})
}
