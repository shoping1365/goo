package handlers

import (
	"net/http"
	"strconv"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// PaymentHandler هندلر عملیات پرداخت
// شامل ایجاد، تایید و مدیریت تراکنش‌های پرداخت

type PaymentHandler struct {
	gatewayRepo     *repository.PaymentGatewayRepository
	transactionRepo *repository.PaymentTransactionRepository
	serviceFactory  *services.PaymentServiceFactory
	calculator      *services.PaymentCalculator
}

// NewPaymentHandler سازنده هندلر پرداخت
func NewPaymentHandler(
	gatewayRepo *repository.PaymentGatewayRepository,
	transactionRepo *repository.PaymentTransactionRepository,
	serviceFactory *services.PaymentServiceFactory,
) *PaymentHandler {
	return &PaymentHandler{
		gatewayRepo:     gatewayRepo,
		transactionRepo: transactionRepo,
		serviceFactory:  serviceFactory,
		calculator:      services.NewPaymentCalculator(),
	}
}

// ساختار درخواست ایجاد پرداخت
type CreatePaymentRequest struct {
	GatewayID   uint   `json:"gateway_id" binding:"required"`   // شناسه درگاه پرداخت
	Amount      int64  `json:"amount" binding:"required,min=1"` // مبلغ (تومان)
	OrderID     string `json:"order_id" binding:"required"`     // شماره سفارش
	Description string `json:"description" binding:"required"`  // توضیحات
	Email       string `json:"email"`                           // ایمیل (اختیاری)
	Mobile      string `json:"mobile"`                          // موبایل (اختیاری)
	CallbackURL string `json:"callback_url" binding:"required"` // آدرس callback
}

// ساختار پاسخ ایجاد پرداخت
type CreatePaymentResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	PaymentURL    string `json:"payment_url,omitempty"`
	Authority     string `json:"authority,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
}

// ساختار درخواست تایید پرداخت
type VerifyPaymentRequest struct {
	Authority string `json:"authority" binding:"required"` // Authority از زرین‌پال
	Amount    int64  `json:"amount" binding:"required"`    // مبلغ (تومان)
}

// ساختار پاسخ تایید پرداخت
type VerifyPaymentResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	RefID         string `json:"ref_id,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
}

// CreatePayment ایجاد درخواست پرداخت جدید
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر است",
			"error":   err.Error(),
		})
		return
	}

	// دریافت درگاه پرداخت
	gateway, err := h.gatewayRepo.GetByID(req.GatewayID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پرداخت یافت نشد",
		})
		return
	}

	// بررسی وضعیت درگاه
	if gateway.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "درگاه پرداخت غیرفعال است",
		})
		return
	}

	// بررسی محدودیت‌های مبلغ با استفاده از سرویس محاسبه‌گر
	isValid, errorMessage := h.calculator.ValidateAmount(req.Amount, gateway)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": errorMessage,
		})
		return
	}

	// ایجاد سرویس درگاه پرداخت
	paymentService, err := h.serviceFactory.CreateService(gateway)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در ایجاد سرویس پرداخت",
			"error":   err.Error(),
		})
		return
	}

	// ایجاد تراکنش در درگاه
	transaction, err := paymentService.CreatePayment(int(req.Amount), req.CallbackURL, req.Description, req.Email, req.Mobile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در ایجاد تراکنش",
			"error":   err.Error(),
		})
		return
	}

	// تنظیم OrderID
	transaction.OrderID = req.OrderID

	// ذخیره تراکنش در دیتابیس
	err = h.transactionRepo.Create(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در ذخیره تراکنش",
			"error":   err.Error(),
		})
		return
	}

	// دریافت آدرس پرداخت
	paymentURL := paymentService.GetPaymentURL(transaction, req.CallbackURL)

	// پاسخ موفقیت‌آمیز
	response := CreatePaymentResponse{
		Success:       true,
		Message:       "تراکنش با موفقیت ایجاد شد",
		PaymentURL:    paymentURL,
		Authority:     transaction.TransactionID,
		TransactionID: transaction.TransactionID,
	}

	c.JSON(http.StatusOK, response)
}

// VerifyPayment تایید پرداخت
func (h *PaymentHandler) VerifyPayment(c *gin.Context) {
	var req VerifyPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر است",
			"error":   err.Error(),
		})
		return
	}

	// دریافت تراکنش از دیتابیس
	transaction, err := h.transactionRepo.GetByTransactionID(req.Authority)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "تراکنش یافت نشد",
		})
		return
	}

	// بررسی وضعیت تراکنش
	if transaction.Status == "success" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "این تراکنش قبلاً تایید شده است",
		})
		return
	}

	if transaction.Status == "failed" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "این تراکنش ناموفق بوده است",
		})
		return
	}

	// دریافت درگاه پرداخت
	gateway, err := h.gatewayRepo.GetByID(transaction.GatewayID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پرداخت یافت نشد",
		})
		return
	}

	// ایجاد سرویس درگاه پرداخت
	paymentService, err := h.serviceFactory.CreateService(gateway)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در ایجاد سرویس پرداخت",
			"error":   err.Error(),
		})
		return
	}

	// تایید پرداخت در درگاه
	verified, refID, err := paymentService.VerifyPayment(int(req.Amount), req.Authority)
	if err != nil {
		// بروزرسانی وضعیت تراکنش به ناموفق
		transaction.Status = "failed"
		transaction.ErrorMessage = err.Error()
		transaction.UpdatedAt = time.Now()
		h.transactionRepo.Update(transaction)

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در تایید پرداخت",
			"error":   err.Error(),
		})
		return
	}

	if !verified {
		// بروزرسانی وضعیت تراکنش به ناموفق
		transaction.Status = "failed"
		transaction.ErrorMessage = "پرداخت تایید نشد"
		transaction.UpdatedAt = time.Now()
		h.transactionRepo.Update(transaction)

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پرداخت تایید نشد",
		})
		return
	}

	// بروزرسانی وضعیت تراکنش به موفق
	transaction.Status = "success"
	transaction.GatewayResponse = refID
	transaction.UpdatedAt = time.Now()
	err = h.transactionRepo.Update(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در بروزرسانی تراکنش",
			"error":   err.Error(),
		})
		return
	}

	// در صورت اتصال تراکنش به سفارش، به‌روزرسانی سفارش و کسر موجودی در زمان پرداخت
	if transaction.OrderID != "" {
		// یافتن سفارش بر اساس tracking_code به همراه آیتم‌ها
		var ord models.Order
		if err := h.transactionRepo.DB().Preload("OrderItems").Where("tracking_code = ?", transaction.OrderID).First(&ord).Error; err == nil {
			// علامت‌گذاری سفارش به paid
			_ = h.transactionRepo.DB().Model(&models.Order{}).Where("id = ?", ord.ID).Updates(map[string]interface{}{
				"is_paid":        true,
				"payment_status": "paid",
				"updated_at":     time.Now(),
			}).Error
			// کسر موجودی اقلام از انبار پیش‌فرض
			inventoryRepo := repository.NewInventoryRepository(h.transactionRepo.DB())
			orderRepo := repository.NewOrderRepository(h.transactionRepo.DB())
			inv := services.NewInventoryService(inventoryRepo, orderRepo)
			for _, it := range ord.OrderItems {
				_ = inv.DeductOnPayment(c.Request.Context(), it.ProductID, it.Quantity)
			}
		}
	}

	// پاسخ موفقیت‌آمیز
	response := VerifyPaymentResponse{
		Success:       true,
		Message:       "پرداخت با موفقیت تایید شد",
		RefID:         refID,
		TransactionID: transaction.TransactionID,
	}

	// --- هوک ارزیابی تقلب پس از تایید پرداخت ---
	go func(trx *repository.PaymentTransactionRepository, tID string) {
		// دریافت تراکنش مجدد برای OrderID
		tx, _ := trx.GetByTransactionID(tID)
		if tx != nil && tx.OrderID != "" {
			// تلاش برای یافتن سفارش
			var ord models.Order
			if err := h.transactionRepo.DB().Where("tracking_code = ?", tx.OrderID).First(&ord).Error; err == nil {
				uowf := unitofwork.NewUnitOfWorkFactory(h.transactionRepo.DB())
				fraud := services.NewFraudService(uowf, "internal_salt")
				// دریافت آدرس
				var addr models.UserAddress
				_ = h.transactionRepo.DB().Where("id = ?", ord.ShippingAddressID).First(&addr).Error
				// دریافت کاربر
				var u models.User
				if ord.UserID != nil {
					_ = h.transactionRepo.DB().First(&u, *ord.UserID).Error
				}
				ip := "" // در callback ممکن است IP متفاوت باشد
				deviceID := ""
				_, _ = fraud.EvaluateOrder(c, services.EvaluateOrderInput{
					OrderID:          ord.ID,
					UserID:           ord.UserID,
					Amount:           ord.FinalAmount,
					PaymentMethod:    ord.PaymentMethod,
					UserRegisteredAt: &u.RegisteredAt,
					IP:               ip,
					DeviceID:         deviceID,
					ShippingCity:     addr.City,
				})
			}
		}
	}(h.transactionRepo, transaction.TransactionID)

	c.JSON(http.StatusOK, response)
}

// GetTransaction دریافت اطلاعات تراکنش
func (h *PaymentHandler) GetTransaction(c *gin.Context) {
	transactionID := c.Param("id")
	if transactionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "شناسه تراکنش الزامی است",
		})
		return
	}

	// دریافت تراکنش از دیتابیس
	transaction, err := h.transactionRepo.GetByTransactionID(transactionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "تراکنش یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    transaction,
	})
}

// GetTransactions دریافت لیست تراکنش‌ها
func (h *PaymentHandler) GetTransactions(c *gin.Context) {
	// پارامترهای فیلتر
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	status := c.Query("status")
	gatewayID, _ := strconv.ParseUint(c.Query("gateway_id"), 10, 32)

	// تنظیم محدودیت‌ها
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// دریافت تراکنش‌ها
	transactions, total, err := h.transactionRepo.GetPaginated(page, limit, status, uint(gatewayID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در دریافت تراکنش‌ها",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transactions": transactions,
			"pagination": gin.H{
				"page":       page,
				"limit":      limit,
				"total":      total,
				"totalPages": (total + int64(limit) - 1) / int64(limit),
			},
		},
	})
}

// TestGatewayConnection تست اتصال درگاه پرداخت
func (h *PaymentHandler) TestGatewayConnection(c *gin.Context) {
	gatewayID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "شناسه درگاه نامعتبر است",
		})
		return
	}

	// دریافت درگاه پرداخت
	gateway, err := h.gatewayRepo.GetByID(uint(gatewayID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پرداخت یافت نشد",
		})
		return
	}

	// تست اتصال
	testErr := h.serviceFactory.TestGatewayConnection(gateway)

	var success bool
	var message string
	if testErr != nil {
		success = false
		message = testErr.Error()
	} else {
		success = true
		message = "اتصال به درگاه با موفقیت برقرار شد"
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": message,
		"gateway": gin.H{
			"id":   gateway.ID,
			"name": gateway.Name,
			"type": gateway.Type,
		},
	})
}

// GetGatewayCapabilities دریافت قابلیت‌های درگاه پرداخت
func (h *PaymentHandler) GetGatewayCapabilities(c *gin.Context) {
	gatewayType := c.Param("type")
	if gatewayType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "نوع درگاه الزامی است",
		})
		return
	}

	capabilities := h.serviceFactory.GetGatewayCapabilities(gatewayType)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"gateway_type": gatewayType,
			"capabilities": capabilities,
		},
	})
}

// CalculatePaymentBreakdown محاسبه تفصیلی مبلغ و کارمزد
func (h *PaymentHandler) CalculatePaymentBreakdown(c *gin.Context) {
	var req struct {
		GatewayID uint  `json:"gateway_id" binding:"required"`
		Amount    int64 `json:"amount" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر است",
			"error":   err.Error(),
		})
		return
	}

	// دریافت درگاه پرداخت
	gateway, err := h.gatewayRepo.GetByID(req.GatewayID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پرداخت یافت نشد",
		})
		return
	}

	// بررسی وضعیت درگاه
	if gateway.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "درگاه پرداخت غیرفعال است",
		})
		return
	}

	// محاسبه تفصیلی
	breakdown := h.calculator.CalculateBreakdown(req.Amount, gateway)

	// بررسی اعتبار مبلغ
	isValid, errorMessage := h.calculator.ValidateAmount(req.Amount, gateway)

	// دریافت محدودیت‌ها
	minAmount, maxAmount, hasMinLimit, hasMaxLimit := h.calculator.GetAmountLimits(gateway)
	fee, hasFee := h.calculator.GetFeeInfo(gateway)

	response := gin.H{
		"success":   isValid,
		"message":   errorMessage,
		"breakdown": breakdown,
		"limits": gin.H{
			"min_amount":    minAmount,
			"max_amount":    maxAmount,
			"has_min_limit": hasMinLimit,
			"has_max_limit": hasMaxLimit,
		},
		"fee": gin.H{
			"percentage": fee,
			"has_fee":    hasFee,
		},
	}

	c.JSON(http.StatusOK, response)
}
