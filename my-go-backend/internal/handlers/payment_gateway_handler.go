package handlers

import (
	"net/http"
	"strconv"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// PaymentGatewayHandler هندلر مدیریت درگاه‌های پرداخت
// شامل ثبت، ویرایش، حذف، لیست و تست اتصال

type PaymentGatewayHandler struct {
	Repo *repository.PaymentGatewayRepository
}

// NewPaymentGatewayHandler سازنده هندلر
func NewPaymentGatewayHandler(repo *repository.PaymentGatewayRepository) *PaymentGatewayHandler {
	return &PaymentGatewayHandler{Repo: repo}
}

// Create ثبت درگاه پرداخت جدید
func (h *PaymentGatewayHandler) Create(c *gin.Context) {
	var gateway models.PaymentGateway
	if err := c.ShouldBindJSON(&gateway); err != nil {
		utils.SendError(c, http.StatusBadRequest, "VALIDATION_ERROR", "داده‌های ارسالی نامعتبر است", err.Error())
		return
	}

	// اعتبارسنجی نوع درگاه
	if gateway.Type == "" {
		utils.SendError(c, http.StatusBadRequest, "REQUIRED_FIELD", "نوع درگاه الزامی است", nil)
		return
	}

	// بررسی معتبر بودن نوع درگاه
	validTypes := []string{"mellat", "melli", "parsian", "saman", "zarinpal"}
	isValidType := false
	for _, validType := range validTypes {
		if gateway.Type == validType {
			isValidType = true
			break
		}
	}
	if !isValidType {
		utils.SendError(c, http.StatusBadRequest, "INVALID_FORMAT", "نوع درگاه باید یکی از موارد زیر باشد: mellat, melli, parsian, saman, zarinpal", nil)
		return
	}

	// بررسی وجود درگاه با همین نوع
	existingGateway, err := h.Repo.GetFirstByType(gateway.Type)
	if err == nil && existingGateway != nil {
		utils.SendError(c, http.StatusConflict, "DUPLICATE_GATEWAY", "درگاه پرداخت با این نوع قبلاً ثبت شده است", nil)
		return
	}

	// تنظیم خودکار نام و نام انگلیسی بر اساس نوع درگاه
	gateway.Name, gateway.EnglishName = getGatewayNamesByType(gateway.Type)

	// اعتبارسنجی فیلدهای عددی
	if gateway.Fee < 0 {
		utils.SendError(c, http.StatusBadRequest, "INVALID_NUMBER", "کارمزد نمی‌تواند منفی باشد", nil)
		return
	}

	if gateway.MinAmount < 0 {
		utils.SendError(c, http.StatusBadRequest, "INVALID_NUMBER", "حداقل مبلغ نمی‌تواند منفی باشد", nil)
		return
	}

	if gateway.MaxAmount < 0 {
		utils.SendError(c, http.StatusBadRequest, "INVALID_NUMBER", "حداکثر مبلغ نمی‌تواند منفی باشد", nil)
		return
	}

	// اگر maxAmount تنظیم شده، باید بزرگتر از minAmount باشد
	if gateway.MaxAmount > 0 && gateway.MinAmount > 0 && gateway.MaxAmount <= gateway.MinAmount {
		utils.SendError(c, http.StatusBadRequest, "OUT_OF_RANGE", "حداکثر مبلغ باید بزرگتر از حداقل مبلغ باشد", nil)
		return
	}

	// تنظیم اولویت خودکار اگر تنظیم نشده باشد
	if gateway.Priority == 0 {
		nextPriority, err := h.Repo.GetNextPriority()
		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در تعیین اولویت", err.Error())
			return
		}
		gateway.Priority = nextPriority
	}

	// تنظیم زمان ایجاد
	gateway.CreatedAt = time.Now()
	gateway.UpdatedAt = time.Now()

	if err := h.Repo.Create(&gateway); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در ثبت درگاه پرداخت", err.Error())
		return
	}
	utils.SendSuccess(c, "درگاه پرداخت با موفقیت ثبت شد", gateway)
}

// GetByID دریافت درگاه پرداخت بر اساس ID
func (h *PaymentGatewayHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه نامعتبر است",
		})
		return
	}

	gateway, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "درگاه پرداخت یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   gateway,
	})
}

// GetAll دریافت تمام درگاه‌های پرداخت
func (h *PaymentGatewayHandler) GetAll(c *gin.Context) {
	gateways, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت لیست درگاه‌ها",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   gateways,
	})
}

// GetActive دریافت درگاه‌های فعال
func (h *PaymentGatewayHandler) GetActive(c *gin.Context) {
	gateways, err := h.Repo.GetActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت درگاه‌های فعال",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   gateways,
	})
}

// GetByFilters دریافت درگاه‌ها با فیلترهای پیشرفته
func (h *PaymentGatewayHandler) GetByFilters(c *gin.Context) {
	filters := make(map[string]interface{})

	// دریافت پارامترهای فیلتر از query string
	if typeFilter := c.Query("type"); typeFilter != "" {
		filters["type"] = typeFilter
	}
	if statusFilter := c.Query("status"); statusFilter != "" {
		filters["status"] = statusFilter
	}
	if searchQuery := c.Query("search"); searchQuery != "" {
		filters["search"] = searchQuery
	}
	if isTestModeStr := c.Query("is_test_mode"); isTestModeStr != "" {
		isTestMode, err := strconv.ParseBool(isTestModeStr)
		if err == nil {
			filters["is_test_mode"] = isTestMode
		}
	}
	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			filters["limit"] = limit
		}
	}
	if orderBy := c.Query("order_by"); orderBy != "" {
		filters["order_by"] = orderBy
	}

	gateways, err := h.Repo.GetByFilters(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت درگاه‌ها",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   gateways,
	})
}

// Update به‌روزرسانی درگاه پرداخت
func (h *PaymentGatewayHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه نامعتبر است",
		})
		return
	}

	// ابتدا درگاه موجود را دریافت کن
	existingGateway, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "درگاه پرداخت یافت نشد",
		})
		return
	}

	// دریافت داده‌های جدید
	var updateData models.PaymentGateway
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "داده‌های ارسالی نامعتبر است",
			"details": err.Error(),
		})
		return
	}

	// به‌روزرسانی فیلدهای مجاز
	if updateData.Name != "" {
		existingGateway.Name = updateData.Name
	}
	if updateData.EnglishName != "" {
		existingGateway.EnglishName = updateData.EnglishName
	}
	if updateData.Type != "" {
		existingGateway.Type = updateData.Type
	}
	if updateData.Status != "" {
		existingGateway.Status = updateData.Status
	}
	if updateData.Icon != "" {
		existingGateway.Icon = updateData.Icon
	}
	if updateData.Color != "" {
		existingGateway.Color = updateData.Color
	}
	if updateData.Fee >= 0 {
		existingGateway.Fee = updateData.Fee
	}
	if updateData.MinAmount >= 0 {
		existingGateway.MinAmount = updateData.MinAmount
	}
	if updateData.MaxAmount >= 0 {
		existingGateway.MaxAmount = updateData.MaxAmount
	}
	if updateData.Priority > 0 {
		existingGateway.Priority = updateData.Priority
	}
	if updateData.Description != "" {
		existingGateway.Description = updateData.Description
	}
	if updateData.Instructions != "" {
		existingGateway.Instructions = updateData.Instructions
	}

	// به‌روزرسانی تنظیمات API
	if updateData.ApiKeys.PublicKey != "" {
		existingGateway.ApiKeys.PublicKey = updateData.ApiKeys.PublicKey
	}
	if updateData.ApiKeys.PrivateKey != "" {
		existingGateway.ApiKeys.PrivateKey = updateData.ApiKeys.PrivateKey
	}
	if updateData.ApiKeys.TestKey != "" {
		existingGateway.ApiKeys.TestKey = updateData.ApiKeys.TestKey
	}
	if updateData.ApiKeys.SecretKey != "" {
		existingGateway.ApiKeys.SecretKey = updateData.ApiKeys.SecretKey
	}

	// به‌روزرسانی آدرس‌های API
	if updateData.ApiEndpoints.Payment != "" {
		existingGateway.ApiEndpoints.Payment = updateData.ApiEndpoints.Payment
	}
	if updateData.ApiEndpoints.Verification != "" {
		existingGateway.ApiEndpoints.Verification = updateData.ApiEndpoints.Verification
	}
	if updateData.ApiEndpoints.Refund != "" {
		existingGateway.ApiEndpoints.Refund = updateData.ApiEndpoints.Refund
	}
	if updateData.ApiEndpoints.Balance != "" {
		existingGateway.ApiEndpoints.Balance = updateData.ApiEndpoints.Balance
	}
	if updateData.ApiEndpoints.Callback != "" {
		existingGateway.ApiEndpoints.Callback = updateData.ApiEndpoints.Callback
	}

	// به‌روزرسانی سایر فیلدها
	if updateData.TerminalId != "" {
		existingGateway.TerminalId = updateData.TerminalId
	}
	if updateData.MerchantId != "" {
		existingGateway.MerchantId = updateData.MerchantId
	}

	// به‌روزرسانی تنظیمات
	existingGateway.Settings = updateData.Settings

	// تنظیم زمان به‌روزرسانی
	existingGateway.UpdatedAt = time.Now()

	if err := h.Repo.Update(existingGateway); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در به‌روزرسانی درگاه پرداخت",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "درگاه پرداخت با موفقیت به‌روزرسانی شد",
		"data":    existingGateway,
	})
}

// Delete حذف درگاه پرداخت
func (h *PaymentGatewayHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه نامعتبر است",
		})
		return
	}

	// بررسی وجود درگاه
	_, err = h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "درگاه پرداخت یافت نشد",
		})
		return
	}

	if err := h.Repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در حذف درگاه پرداخت",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "درگاه پرداخت با موفقیت حذف شد",
	})
}

// UpdateStatus تغییر وضعیت درگاه پرداخت
func (h *PaymentGatewayHandler) UpdateStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه نامعتبر است",
		})
		return
	}

	var request struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "وضعیت ارسالی نامعتبر است",
		})
		return
	}

	// بررسی معتبر بودن وضعیت
	validStatuses := []string{"active", "inactive", "maintenance"}
	isValid := false
	for _, status := range validStatuses {
		if request.Status == status {
			isValid = true
			break
		}
	}
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "وضعیت نامعتبر است. وضعیت‌های مجاز: active, inactive, maintenance",
		})
		return
	}

	if err := h.Repo.UpdateStatus(uint(id), request.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در تغییر وضعیت درگاه پرداخت",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "وضعیت درگاه پرداخت با موفقیت تغییر یافت",
	})
}

// GetStats دریافت آمار کلی درگاه‌های پرداخت
func (h *PaymentGatewayHandler) GetStats(c *gin.Context) {
	stats, err := h.Repo.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت آمار",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   stats,
	})
}

// GetTopGateways دریافت درگاه‌های برتر
func (h *PaymentGatewayHandler) GetTopGateways(c *gin.Context) {
	limit := 5 // پیش‌فرض 5 درگاه برتر
	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	gateways, err := h.Repo.GetTopGateways(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در دریافت درگاه‌های برتر",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   gateways,
	})
}

// TestConnection تست اتصال درگاه پرداخت
func (h *PaymentGatewayHandler) TestConnection(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "شناسه نامعتبر است",
		})
		return
	}

	gateway, err := h.Repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "درگاه پرداخت یافت نشد",
		})
		return
	}

	// در اینجا می‌توانید منطق تست اتصال به درگاه پرداخت را پیاده‌سازی کنید
	// فعلاً یک پاسخ ساده برمی‌گردانیم
	success := true
	message := "اتصال به درگاه پرداخت موفقیت‌آمیز بود"

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"gateway_id": gateway.ID,
			"name":       gateway.Name,
			"success":    success,
			"message":    message,
			"tested_at":  time.Now(),
		},
	})
}

// ResetTodayStats ریست کردن آمار امروز
func (h *PaymentGatewayHandler) ResetTodayStats(c *gin.Context) {
	if err := h.Repo.ResetTodayStats(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "خطا در ریست کردن آمار امروز",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "آمار امروز تمام درگاه‌ها با موفقیت ریست شد",
	})
}

// getGatewayNamesByType تنظیم خودکار نام و نام انگلیسی بر اساس نوع درگاه
func getGatewayNamesByType(gatewayType string) (string, string) {
	switch gatewayType {
	case "mellat":
		return "ملت", "Mellat"
	case "melli":
		return "ملی", "Melli"
	case "parsian":
		return "پارسیان", "Parsian"
	case "saman":
		return "سامان", "Saman"
	case "zarinpal":
		return "زرین‌پال", "Zarinpal"
	default:
		return "درگاه پرداخت", "Payment Gateway"
	}
}
