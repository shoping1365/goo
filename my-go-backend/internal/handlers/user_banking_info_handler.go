package handlers

import (
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserBankingInfoHandler هندلر مدیریت اطلاعات بانکی کاربران
type UserBankingInfoHandler struct {
	bankingInfoService *services.UserBankingInfoService
}

// NewUserBankingInfoHandler ایجاد نمونه جدید از UserBankingInfoHandler
func NewUserBankingInfoHandler(bankingInfoService *services.UserBankingInfoService) *UserBankingInfoHandler {
	return &UserBankingInfoHandler{
		bankingInfoService: bankingInfoService,
	}
}

// Create ایجاد اطلاعات بانکی جدید
func (h *UserBankingInfoHandler) Create(c *gin.Context) {
	// دریافت شناسه کاربر از توکن
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده است"})
		return
	}

	var req models.BankingInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "داده‌های ورودی نامعتبر است"})
		return
	}

	// ایجاد اطلاعات بانکی جدید
	response, err := h.bankingInfoService.Create(userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "حساب جدید با موفقیت اضافه شد",
		"data":    response,
	})
}

// Update به‌روزرسانی اطلاعات بانکی موجود
func (h *UserBankingInfoHandler) Update(c *gin.Context) {
	// دریافت شناسه کاربر از توکن
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده است"})
		return
	}

	// دریافت شناسه حساب از پارامتر
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه حساب نامعتبر است"})
		return
	}

	var req models.BankingInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "داده‌های ورودی نامعتبر است"})
		return
	}

	// به‌روزرسانی اطلاعات بانکی
	response, err := h.bankingInfoService.Update(userID.(uint), uint(id), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "حساب با موفقیت به‌روزرسانی شد",
		"data":    response,
	})
}

// Get دریافت تمام اطلاعات بانکی کاربر
func (h *UserBankingInfoHandler) Get(c *gin.Context) {
	// دریافت شناسه کاربر از توکن
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده است"})
		return
	}

	// دریافت تمام اطلاعات بانکی
	responses, err := h.bankingInfoService.GetAllByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت اطلاعات بانکی"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "اطلاعات بانکی با موفقیت دریافت شد",
		"data":    responses,
	})
}

// Delete حذف اطلاعات بانکی
func (h *UserBankingInfoHandler) Delete(c *gin.Context) {
	// دریافت شناسه کاربر از توکن
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده است"})
		return
	}

	// دریافت شناسه حساب از پارامتر
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه حساب نامعتبر است"})
		return
	}

	// حذف اطلاعات بانکی
	err = h.bankingInfoService.Delete(userID.(uint), uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "حساب با موفقیت حذف شد"})
}

// SetDefault تنظیم حساب پیشفرض
func (h *UserBankingInfoHandler) SetDefault(c *gin.Context) {
	// دریافت شناسه کاربر از توکن
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده است"})
		return
	}

	// دریافت شناسه حساب از پارامتر
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه حساب نامعتبر است"})
		return
	}

	// تنظیم حساب پیشفرض
	err = h.bankingInfoService.SetDefault(userID.(uint), uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "حساب پیشفرض با موفقیت تنظیم شد"})
}

// AutoCompleteFromCard تکمیل خودکار از شماره کارت
func (h *UserBankingInfoHandler) AutoCompleteFromCard(c *gin.Context) {
	var req struct {
		CardNumber string `json:"card_number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "داده‌های ورودی نامعتبر است",
			"error":   err.Error(),
		})
		return
	}

	result, err := h.bankingInfoService.AutoCompleteFromCard(req.CardNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "اطلاعات با موفقیت تکمیل شد",
		"data":    result,
	})
}

// AutoCompleteFromSheba تکمیل خودکار از شماره شبا
func (h *UserBankingInfoHandler) AutoCompleteFromSheba(c *gin.Context) {
	var req struct {
		ShebaNumber string `json:"sheba_number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "داده‌های ورودی نامعتبر است",
			"error":   err.Error(),
		})
		return
	}

	result, err := h.bankingInfoService.AutoCompleteFromSheba(req.ShebaNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "اطلاعات با موفقیت تکمیل شد",
		"data":    result,
	})
}

// Verify تایید اطلاعات بانکی توسط ادمین
func (h *UserBankingInfoHandler) Verify(c *gin.Context) {
	// دریافت شناسه ادمین از توکن
	adminID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ادمین احراز هویت نشده است"})
		return
	}

	// دریافت شناسه اطلاعات بانکی از پارامتر
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه اطلاعات بانکی نامعتبر است"})
		return
	}

	var req struct {
		Note string `json:"note"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "داده‌های ورودی نامعتبر است"})
		return
	}

	// تایید اطلاعات بانکی
	err = h.bankingInfoService.Verify(uint(id), adminID.(uint), req.Note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "اطلاعات بانکی با موفقیت تایید شد"})
}

// Unverify لغو تایید اطلاعات بانکی توسط ادمین
func (h *UserBankingInfoHandler) Unverify(c *gin.Context) {
	// دریافت شناسه ادمین از توکن
	adminID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ادمین احراز هویت نشده است"})
		return
	}

	// دریافت شناسه اطلاعات بانکی از پارامتر
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه اطلاعات بانکی نامعتبر است"})
		return
	}

	var req struct {
		Note string `json:"note"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "داده‌های ورودی نامعتبر است"})
		return
	}

	// لغو تایید اطلاعات بانکی
	err = h.bankingInfoService.Unverify(uint(id), adminID.(uint), req.Note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "تایید اطلاعات بانکی لغو شد"})
}

// GetByUserID دریافت اطلاعات بانکی یک کاربر توسط ادمین
func (h *UserBankingInfoHandler) GetByUserID(c *gin.Context) {
	// دریافت شناسه کاربر از پارامتر
	userIDStr := c.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شناسه کاربر نامعتبر است"})
		return
	}

	// دریافت تمام اطلاعات بانکی کاربر
	responses, err := h.bankingInfoService.GetAllByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت اطلاعات بانکی"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "اطلاعات بانکی با موفقیت دریافت شد",
		"data":    responses,
	})
}
