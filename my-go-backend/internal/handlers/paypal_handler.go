package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// PayPalHandler handler مدیریت درگاه پرداخت پی‌پال
type PayPalHandler struct {
	gatewayRepo *repository.PaymentGatewayRepository
	paymentRepo *repository.PaymentTransactionRepository
}

// NewPayPalHandler سازنده handler پی‌پال
func NewPayPalHandler(gatewayRepo *repository.PaymentGatewayRepository, paymentRepo *repository.PaymentTransactionRepository) *PayPalHandler {
	return &PayPalHandler{
		gatewayRepo: gatewayRepo,
		paymentRepo: paymentRepo,
	}
}

// GetPayPalSettings دریافت تنظیمات درگاه پی‌پال
func (h *PayPalHandler) GetPayPalSettings(c *gin.Context) {
	// دریافت درگاه پی‌پال از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("paypal")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پی‌پال یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    gateway,
		"message": "تنظیمات درگاه پی‌پال با موفقیت دریافت شد",
	})
}

// UpdatePayPalSettings به‌روزرسانی تنظیمات درگاه پی‌پال
func (h *PayPalHandler) UpdatePayPalSettings(c *gin.Context) {
	var request struct {
		MerchantId string  `json:"merchant_id" binding:"required"`
		PrivateKey string  `json:"private_key" binding:"required"`
		PublicKey  string  `json:"public_key" binding:"required"`
		TestKey    string  `json:"test_key"`
		IsActive   bool    `json:"is_active"`
		IsTestMode bool    `json:"is_test_mode"`
		MinAmount  int64   `json:"min_amount"`
		MaxAmount  int64   `json:"max_amount"`
		Fee        float64 `json:"fee"`
		Priority   int     `json:"priority"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	// دریافت درگاه پی‌پال از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("paypal")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پی‌پال یافت نشد",
		})
		return
	}

	// به‌روزرسانی تنظیمات
	gateway.MerchantId = request.MerchantId
	gateway.ApiKeys.PrivateKey = request.PrivateKey
	gateway.ApiKeys.PublicKey = request.PublicKey
	gateway.ApiKeys.TestKey = request.TestKey
	gateway.MinAmount = request.MinAmount
	gateway.MaxAmount = request.MaxAmount
	gateway.Fee = request.Fee
	gateway.Priority = request.Priority

	// تنظیم وضعیت
	if request.IsActive {
		gateway.Status = "active"
	} else {
		gateway.Status = "inactive"
	}

	gateway.IsTestMode = request.IsTestMode

	// ذخیره در دیتابیس
	err = h.gatewayRepo.Update(gateway)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در به‌روزرسانی تنظیمات: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    gateway,
		"message": "تنظیمات درگاه پی‌پال با موفقیت به‌روزرسانی شد",
	})
}

// TestPayPalConnection تست اتصال به درگاه پی‌پال
func (h *PayPalHandler) TestPayPalConnection(c *gin.Context) {
	// دریافت درگاه پی‌پال از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("paypal")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پی‌پال یافت نشد",
		})
		return
	}

	// ایجاد سرویس پی‌پال
	paypalService := services.NewPayPalService(gateway)

	// تست اتصال
	err = paypalService.TestConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در تست اتصال: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "اتصال به درگاه پی‌پال موفقیت‌آمیز بود",
	})
}

// CreatePayPalPayment ایجاد پرداخت پی‌پال
func (h *PayPalHandler) CreatePayPalPayment(c *gin.Context) {
	var request struct {
		Amount      int    `json:"amount" binding:"required,min=1"`
		CallbackURL string `json:"callback_url" binding:"required"`
		Description string `json:"description" binding:"required"`
		Email       string `json:"email"`
		Mobile      string `json:"mobile"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	// دریافت درگاه پی‌پال از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("paypal")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پی‌پال یافت نشد",
		})
		return
	}

	// بررسی فعال بودن درگاه
	if gateway.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "درگاه پی‌پال غیرفعال است",
		})
		return
	}

	// ایجاد سرویس پی‌پال
	paypalService := services.NewPayPalService(gateway)

	// ایجاد پرداخت
	transaction, err := paypalService.CreatePayment(request.Amount, request.CallbackURL, request.Description, request.Email, request.Mobile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در ایجاد پرداخت: " + err.Error(),
		})
		return
	}

	// ذخیره تراکنش در دیتابیس
	transaction.GatewayID = gateway.ID
	err = h.paymentRepo.Create(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در ذخیره تراکنش: " + err.Error(),
		})
		return
	}

	// دریافت آدرس پرداخت
	paymentURL := paypalService.GetPaymentURL(transaction, request.CallbackURL)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transaction_id": transaction.TransactionID,
			"order_id":       transaction.OrderID,
			"amount":         transaction.Amount,
			"payment_url":    paymentURL,
			"payment_form":   paypalService.GetPaymentForm(transaction, request.CallbackURL),
		},
		"message": "درخواست پرداخت پی‌پال با موفقیت ایجاد شد",
	})
}

// VerifyPayPalPayment تایید پرداخت پی‌پال
func (h *PayPalHandler) VerifyPayPalPayment(c *gin.Context) {
	var request struct {
		Amount int    `json:"amount" binding:"required,min=1"`
		Token  string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	// دریافت درگاه پی‌پال از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("paypal")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پی‌پال یافت نشد",
		})
		return
	}

	// ایجاد سرویس پی‌پال
	paypalService := services.NewPayPalService(gateway)

	// تایید پرداخت
	success, refNum, err := paypalService.VerifyPayment(request.Amount, request.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در تایید پرداخت: " + err.Error(),
		})
		return
	}

	if !success {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پرداخت تایید نشد",
		})
		return
	}

	// به‌روزرسانی وضعیت تراکنش در دیتابیس
	transaction, err := h.paymentRepo.GetByTransactionID(request.Token)
	if err == nil {
		transaction.Status = "success"
		h.paymentRepo.Update(transaction)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"ref_num": refNum,
			"amount":  request.Amount,
		},
		"message": "پرداخت پی‌پال با موفقیت تایید شد",
	})
}

// RefundPayPalPayment بازگشت وجه پی‌پال
func (h *PayPalHandler) RefundPayPalPayment(c *gin.Context) {
	var request struct {
		TransactionID string `json:"transaction_id" binding:"required"`
		Amount        int    `json:"amount" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	// دریافت درگاه پی‌پال از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("paypal")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پی‌پال یافت نشد",
		})
		return
	}

	// ایجاد سرویس پی‌پال
	paypalService := services.NewPayPalService(gateway)

	// بازگشت وجه
	err = paypalService.RefundPayment(request.TransactionID, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در بازگشت وجه: " + err.Error(),
		})
		return
	}

	// به‌روزرسانی وضعیت تراکنش در دیتابیس
	transaction, err := h.paymentRepo.GetByTransactionID(request.TransactionID)
	if err == nil {
		transaction.Status = "refunded"
		h.paymentRepo.Update(transaction)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "بازگشت وجه پی‌پال با موفقیت انجام شد",
	})
}

// ProcessPayPalCallback پردازش callback پی‌پال
func (h *PayPalHandler) ProcessPayPalCallback(c *gin.Context) {
	// دریافت پارامترهای callback
	paymentID := c.Query("paymentId")
	payerID := c.Query("PayerID")
	success := c.Query("success")

	// تبدیل پارامترها به map
	params := map[string]string{
		"paymentId": paymentID,
		"PayerID":   payerID,
		"success":   success,
	}

	// دریافت درگاه پی‌پال از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("paypal")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پی‌پال یافت نشد",
		})
		return
	}

	// ایجاد سرویس پی‌پال
	paypalService := services.NewPayPalService(gateway)

	// پردازش callback
	transaction, err := paypalService.ProcessCallback(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "خطا در پردازش callback: " + err.Error(),
		})
		return
	}

	// ذخیره تراکنش در دیتابیس
	transaction.GatewayID = gateway.ID
	err = h.paymentRepo.Create(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در ذخیره تراکنش: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    transaction,
		"message": "Callback پی‌پال با موفقیت پردازش شد",
	})
}

// GetPayPalTransactions دریافت تراکنش‌های پی‌پال
func (h *PayPalHandler) GetPayPalTransactions(c *gin.Context) {
	// دریافت پارامترهای صفحه‌بندی
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	// دریافت درگاه پی‌پال
	gateway, err := h.gatewayRepo.GetFirstByType("paypal")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پی‌پال یافت نشد",
		})
		return
	}

	// دریافت تراکنش‌ها
	transactions, total, err := h.paymentRepo.GetByGateway(gateway.ID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در دریافت تراکنش‌ها: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transactions": transactions,
			"total":        total,
			"page":         page,
			"limit":        limit,
		},
		"message": "تراکنش‌های پی‌پال با موفقیت دریافت شد",
	})
}
