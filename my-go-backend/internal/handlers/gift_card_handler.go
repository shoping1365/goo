package handlers

import (
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GiftCardHandler struct {
	service *services.GiftCardService
}

func NewGiftCardHandler(service *services.GiftCardService) *GiftCardHandler {
	return &GiftCardHandler{service: service}
}

// GetDashboardStats دریافت آمار داشبورد
func (h *GiftCardHandler) GetDashboardStats(c *gin.Context) {
	stats, err := h.service.GetDashboardStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در دریافت آمار داشبورد",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    stats,
	})
}

// GetChartData دریافت داده‌های نمودار
func (h *GiftCardHandler) GetChartData(c *gin.Context) {
	data, err := h.service.GetChartData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در دریافت داده‌های نمودار",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// GetRecentActivities دریافت فعالیت‌های اخیر
func (h *GiftCardHandler) GetRecentActivities(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	activities, err := h.service.GetRecentActivities(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در دریافت فعالیت‌های اخیر",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    activities,
	})
}

// GetAllGiftCards دریافت تمام گیفت کارت‌ها
func (h *GiftCardHandler) GetAllGiftCards(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "20")
	status := c.Query("status")
	cardType := c.Query("type")
	category := c.Query("category")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 20
	}

	giftCards, total, err := h.service.GetAllGiftCards(page, pageSize, status, cardType, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در دریافت لیست گیفت کارت‌ها",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"gift_cards": giftCards,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		},
	})
}

// GetGiftCard دریافت گیفت کارت بر اساس ID
func (h *GiftCardHandler) GetGiftCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر",
		})
		return
	}

	giftCard, err := h.service.GetGiftCard(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "گیفت کارت یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    giftCard,
	})
}

// GetGiftCardByCode دریافت گیفت کارت بر اساس کد
func (h *GiftCardHandler) GetGiftCardByCode(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "کد گیفت کارت الزامی است",
		})
		return
	}

	giftCard, err := h.service.GetGiftCardByCode(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "گیفت کارت یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    giftCard,
	})
}

// CreateGiftCard ایجاد گیفت کارت جدید
func (h *GiftCardHandler) CreateGiftCard(c *gin.Context) {
	var giftCard models.GiftCard
	if err := c.ShouldBindJSON(&giftCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "داده‌های ورودی نامعتبر",
			"details": err.Error(),
		})
		return
	}

	// دریافت کاربر از context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "کاربر احراز هویت نشده",
		})
		return
	}
	giftCard.CreatedBy = userID.(uint)

	// اعتبارسنجی
	if err := h.service.ValidateGiftCard(&giftCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// ایجاد گیفت کارت
	if err := h.service.CreateGiftCard(&giftCard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در ایجاد گیفت کارت",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    giftCard,
		"message": "گیفت کارت با موفقیت ایجاد شد",
	})
}

// UpdateGiftCard به‌روزرسانی گیفت کارت
func (h *GiftCardHandler) UpdateGiftCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر",
		})
		return
	}

	var giftCard models.GiftCard
	if err := c.ShouldBindJSON(&giftCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "داده‌های ورودی نامعتبر",
			"details": err.Error(),
		})
		return
	}

	giftCard.ID = uint(id)

	// دریافت کاربر از context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "کاربر احراز هویت نشده",
		})
		return
	}
	giftCard.CreatedBy = userID.(uint)

	// اعتبارسنجی
	if err := h.service.ValidateGiftCard(&giftCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// به‌روزرسانی
	if err := h.service.UpdateGiftCard(&giftCard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در به‌روزرسانی گیفت کارت",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    giftCard,
		"message": "گیفت کارت با موفقیت به‌روزرسانی شد",
	})
}

// DeleteGiftCard حذف گیفت کارت
func (h *GiftCardHandler) DeleteGiftCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر",
		})
		return
	}

	// دریافت کاربر از context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "کاربر احراز هویت نشده",
		})
		return
	}

	if err := h.service.DeleteGiftCard(uint(id), userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در حذف گیفت کارت",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "گیفت کارت با موفقیت حذف شد",
	})
}

// UseGiftCard استفاده از گیفت کارت
func (h *GiftCardHandler) UseGiftCard(c *gin.Context) {
	var request struct {
		Code   string `json:"code" binding:"required"`
		Amount int64  `json:"amount" binding:"required,gt=0"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "داده‌های ورودی نامعتبر",
			"details": err.Error(),
		})
		return
	}

	// دریافت کاربر از context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "کاربر احراز هویت نشده",
		})
		return
	}

	if err := h.service.UseGiftCard(request.Code, request.Amount, userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "گیفت کارت با موفقیت استفاده شد",
	})
}

// ActivateGiftCard فعال کردن گیفت کارت
func (h *GiftCardHandler) ActivateGiftCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر",
		})
		return
	}

	// دریافت کاربر از context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "کاربر احراز هویت نشده",
		})
		return
	}

	if err := h.service.ActivateGiftCard(uint(id), userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "گیفت کارت با موفقیت فعال شد",
	})
}

// DeactivateGiftCard غیرفعال کردن گیفت کارت
func (h *GiftCardHandler) DeactivateGiftCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر",
		})
		return
	}

	// دریافت کاربر از context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "کاربر احراز هویت نشده",
		})
		return
	}

	if err := h.service.DeactivateGiftCard(uint(id), userID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "گیفت کارت با موفقیت غیرفعال شد",
	})
}

// SearchGiftCards جستجو در گیفت کارت‌ها
func (h *GiftCardHandler) SearchGiftCards(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "عبارت جستجو الزامی است",
		})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "20")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 20
	}

	giftCards, total, err := h.service.SearchGiftCards(query, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در جستجو",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"gift_cards": giftCards,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
			"query":      query,
		},
	})
}

// GetTransactions دریافت تراکنش‌های گیفت کارت
func (h *GiftCardHandler) GetTransactions(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "شناسه نامعتبر",
		})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "20")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 20
	}

	transactions, total, err := h.service.GetTransactions(uint(id), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "خطا در دریافت تراکنش‌ها",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transactions": transactions,
			"total":        total,
			"page":         page,
			"page_size":    pageSize,
		},
	})
}
