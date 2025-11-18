package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/utils"
)

// StockDiscountNotificationHandler
// کنترلر مدیریت درخواست‌های خبرم کن (موجودی و تخفیف)
type StockDiscountNotificationHandler struct {
	DB                       *gorm.DB
	StockNotificationRepo    repository.StockNotificationRepositoryInterface
	DiscountNotificationRepo repository.DiscountNotificationRepositoryInterface
}

func NewStockDiscountNotificationHandler(db *gorm.DB) *StockDiscountNotificationHandler {
	return &StockDiscountNotificationHandler{
		DB:                       db,
		StockNotificationRepo:    repository.NewStockNotificationRepository(db),
		DiscountNotificationRepo: repository.NewDiscountNotificationRepository(db),
	}
}

// ------------------------------ STOCK --------------------------------------

// ListStockNotifications فهرست اعلان‌های موجودی
func (h *StockDiscountNotificationHandler) ListStockNotifications(c *gin.Context) {
	list, err := h.StockNotificationRepo.GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت اعلان‌های موجودی", err.Error()))
		return
	}
	c.JSON(http.StatusOK, list)
}

// CreateStockNotification ایجاد اعلان جدید موجودی
func (h *StockDiscountNotificationHandler) CreateStockNotification(c *gin.Context) {
	var req struct {
		UserID    uint   `json:"user_id" binding:"required"`
		ProductID uint   `json:"product_id" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Phone     string `json:"phone"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}

	n := &models.StockNotification{UserID: req.UserID, ProductID: req.ProductID, Email: req.Email, Phone: req.Phone}
	if err := h.StockNotificationRepo.Create(c.Request.Context(), n); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد اعلان موجودی", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, n)
}

// DeleteStockNotification حذف اعلان موجودی
func (h *StockDiscountNotificationHandler) DeleteStockNotification(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.StockNotificationRepo.Delete(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف اعلان", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// MarkStockAsSent علامت‌گذاری اعلان به عنوان ارسال شده
func (h *StockDiscountNotificationHandler) MarkStockAsSent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.StockNotificationRepo.MarkAsSent(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بروزرسانی وضعیت", err.Error()))
		return
	}
	// ثبت لاگ موفقیت
	_ = h.DB.Create(&models.NotificationLog{Channel: "sms", Type: "stock", Notification: uint(id), Status: "success"}).Error
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// ------------------------------ DISCOUNT -----------------------------------

// ListDiscountNotifications فهرست اعلان‌های تخفیف
func (h *StockDiscountNotificationHandler) ListDiscountNotifications(c *gin.Context) {
	list, err := h.DiscountNotificationRepo.GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت اعلان‌های تخفیف", err.Error()))
		return
	}
	c.JSON(http.StatusOK, list)
}

// CreateDiscountNotification ایجاد اعلان جدید تخفیف
func (h *StockDiscountNotificationHandler) CreateDiscountNotification(c *gin.Context) {
	var req struct {
		UserID         uint    `json:"user_id" binding:"required"`
		ProductID      uint    `json:"product_id" binding:"required"`
		Email          string  `json:"email" binding:"required,email"`
		Phone          string  `json:"phone"`
		ThresholdType  string  `json:"threshold_type" binding:"omitempty,oneof=any percent amount"`
		ThresholdValue float64 `json:"threshold_value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}

	n := &models.DiscountNotification{
		UserID:         req.UserID,
		ProductID:      req.ProductID,
		Email:          req.Email,
		Phone:          req.Phone,
		ThresholdType:  req.ThresholdType,
		ThresholdValue: req.ThresholdValue,
	}
	if err := h.DiscountNotificationRepo.Create(c.Request.Context(), n); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ایجاد اعلان تخفیف", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, n)
}

// DeleteDiscountNotification حذف اعلان تخفیف
func (h *StockDiscountNotificationHandler) DeleteDiscountNotification(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.DiscountNotificationRepo.Delete(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف اعلان", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// MarkDiscountAsSent علامت‌گذاری اعلان تخفیف به عنوان ارسال شده
func (h *StockDiscountNotificationHandler) MarkDiscountAsSent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.DiscountNotificationRepo.MarkAsSent(c.Request.Context(), uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بروزرسانی وضعیت", err.Error()))
		return
	}
	// ثبت لاگ موفقیت
	_ = h.DB.Create(&models.NotificationLog{Channel: "sms", Type: "discount", Notification: uint(id), Status: "success"}).Error
	c.JSON(http.StatusOK, gin.H{"success": true})
}
