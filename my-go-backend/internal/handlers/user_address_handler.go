package handlers

import (
	"net/http"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddressHandler مدیریت آدرس‌های کاربر
type AddressHandler struct {
	DB *gorm.DB
}

func NewAddressHandler(db *gorm.DB) *AddressHandler {
	return &AddressHandler{DB: db}
}

// ---------------- GET /user/addresses ----------------
func (h *AddressHandler) List(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}

	var list []models.UserAddress
	if err := h.DB.Where("user_id = ?", userID).Order("is_default DESC, created_at DESC").Find(&list).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت آدرس‌ها", err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

// ---------------- POST /user/addresses ----------------
func (h *AddressHandler) Create(c *gin.Context) {
	session := sessions.Default(c)
	userIDVal := session.Get("user_id")
	if userIDVal == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}
	userID := userIDVal.(uint)

	var payload struct {
		Street          string `json:"street" binding:"required"`
		PostalCode      string `json:"postal_code"`
		RecipientName   string `json:"recipient_name"`
		RecipientMobile string `json:"recipient_mobile"`
		Phone           string `json:"phone"`
		Province        string `json:"province"`
		City            string `json:"city"`
		ProvinceID      *uint  `json:"province_id"`
		CityID          *uint  `json:"city_id"`
		IsDefault       bool   `json:"is_default"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.SendError(c, http.StatusBadRequest, "VALIDATION_ERROR", "داده‌های ورودی نامعتبر است", err.Error())
		return
	}

	// اگر is_default=true باشد، ابتدا سایر آدرس‌ها را غیرفعال کن
	if payload.IsDefault {
		_ = h.DB.Model(&models.UserAddress{}).Where("user_id = ?", userID).Update("is_default", false).Error
	}

	addr := models.UserAddress{
		UserID:          userID,
		Street:          payload.Street,
		PostalCode:      payload.PostalCode,
		RecipientName:   payload.RecipientName,
		RecipientMobile: payload.RecipientMobile,
		Phone:           payload.Phone,
		Province:        payload.Province,
		City:            payload.City,
		ProvinceID:      payload.ProvinceID,
		CityID:          payload.CityID,
		IsDefault:       payload.IsDefault,
	}
	if err := h.DB.Create(&addr).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در ذخیره آدرس", err.Error())
		return
	}
	c.JSON(http.StatusCreated, addr)
}

// ---------------- PUT /user/addresses/:id ----------------
func (h *AddressHandler) Update(c *gin.Context) {
	session := sessions.Default(c)
	userIDVal := session.Get("user_id")
	if userIDVal == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}
	userID := userIDVal.(uint)

	// دریافت ID آدرس از URL
	addressID := c.Param("id")
	if addressID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "شناسه آدرس الزامی است"})
		return
	}

	var payload struct {
		Street          string `json:"street" binding:"required"`
		PostalCode      string `json:"postal_code"`
		RecipientName   string `json:"recipient_name"`
		RecipientMobile string `json:"recipient_mobile"`
		Phone           string `json:"phone"`
		Province        string `json:"province"`
		City            string `json:"city"`
		ProvinceID      *uint  `json:"province_id"`
		CityID          *uint  `json:"city_id"`
		IsDefault       bool   `json:"is_default"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.SendError(c, http.StatusBadRequest, "VALIDATION_ERROR", "داده‌های ورودی نامعتبر است", err.Error())
		return
	}

	// بررسی مالکیت آدرس
	var existingAddr models.UserAddress
	if err := h.DB.Where("id = ? AND user_id = ?", addressID, userID).First(&existingAddr).Error; err != nil {
		utils.SendError(c, http.StatusNotFound, "ADDRESS_NOT_FOUND", "آدرس یافت نشد یا متعلق به کاربر نیست", nil)
		return
	}

	// اگر is_default=true باشد، ابتدا سایر آدرس‌ها را غیرفعال کن
	if payload.IsDefault {
		_ = h.DB.Model(&models.UserAddress{}).Where("user_id = ? AND id != ?", userID, addressID).Update("is_default", false).Error
	}

	// به‌روزرسانی آدرس
	updates := map[string]interface{}{
		"street":           payload.Street,
		"postal_code":      payload.PostalCode,
		"recipient_name":   payload.RecipientName,
		"recipient_mobile": payload.RecipientMobile,
		"phone":            payload.Phone,
		"province":         payload.Province,
		"city":             payload.City,
		"province_id":      payload.ProvinceID,
		"city_id":          payload.CityID,
		"is_default":       payload.IsDefault,
	}

	if err := h.DB.Model(&existingAddr).Updates(updates).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در به‌روزرسانی آدرس", err.Error())
		return
	}

	// دریافت آدرس به‌روزرسانی شده
	var updatedAddr models.UserAddress
	h.DB.First(&updatedAddr, addressID)
	c.JSON(http.StatusOK, updatedAddr)
}
