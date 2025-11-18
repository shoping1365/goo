package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// StripeHandler handler مدیریت درگاه پرداخت استرایپ
type StripeHandler struct {
	gatewayRepo *repository.PaymentGatewayRepository
	paymentRepo *repository.PaymentTransactionRepository
}

// NewStripeHandler سازنده handler استرایپ
func NewStripeHandler(gatewayRepo *repository.PaymentGatewayRepository, paymentRepo *repository.PaymentTransactionRepository) *StripeHandler {
	return &StripeHandler{
		gatewayRepo: gatewayRepo,
		paymentRepo: paymentRepo,
	}
}

// GetStripeSettings دریافت تنظیمات درگاه استرایپ
func (h *StripeHandler) GetStripeSettings(c *gin.Context) {
	// دریافت درگاه استرایپ از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("stripe")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه استرایپ یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    gateway,
		"message": "تنظیمات درگاه استرایپ با موفقیت دریافت شد",
	})
}

// UpdateStripeSettings به‌روزرسانی تنظیمات درگاه استرایپ
func (h *StripeHandler) UpdateStripeSettings(c *gin.Context) {
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

	// دریافت درگاه استرایپ از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("stripe")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه استرایپ یافت نشد",
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
		"message": "تنظیمات درگاه استرایپ با موفقیت به‌روزرسانی شد",
	})
}

// TestStripeConnection تست اتصال به درگاه استرایپ
func (h *StripeHandler) TestStripeConnection(c *gin.Context) {
	// دریافت درگاه استرایپ از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("stripe")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه استرایپ یافت نشد",
		})
		return
	}

	// ایجاد سرویس استرایپ
	stripeService := services.NewStripeService(gateway)

	// تست اتصال
	err = stripeService.TestConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در تست اتصال: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "اتصال به درگاه استرایپ موفقیت‌آمیز بود",
	})
}

// CreateStripePayment ایجاد پرداخت استرایپ
func (h *StripeHandler) CreateStripePayment(c *gin.Context) {
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

	// دریافت درگاه استرایپ از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("stripe")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه استرایپ یافت نشد",
		})
		return
	}

	// بررسی فعال بودن درگاه
	if gateway.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "درگاه استرایپ غیرفعال است",
		})
		return
	}

	// ایجاد سرویس استرایپ
	stripeService := services.NewStripeService(gateway)

	// ایجاد پرداخت
	transaction, err := stripeService.CreatePayment(request.Amount, request.CallbackURL, request.Description, request.Email, request.Mobile)
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
	paymentURL := stripeService.GetPaymentURL(transaction, request.CallbackURL)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transaction_id": transaction.TransactionID,
			"order_id":       transaction.OrderID,
			"amount":         transaction.Amount,
			"payment_url":    paymentURL,
			"payment_form":   stripeService.GetPaymentForm(transaction, request.CallbackURL),
		},
		"message": "درخواست پرداخت استرایپ با موفقیت ایجاد شد",
	})
}

// VerifyStripePayment تایید پرداخت استرایپ
func (h *StripeHandler) VerifyStripePayment(c *gin.Context) {
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

	// دریافت درگاه استرایپ از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("stripe")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه استرایپ یافت نشد",
		})
		return
	}

	// ایجاد سرویس استرایپ
	stripeService := services.NewStripeService(gateway)

	// تایید پرداخت
	success, refNum, err := stripeService.VerifyPayment(request.Amount, request.Token)
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
		"message": "پرداخت استرایپ با موفقیت تایید شد",
	})
}

// RefundStripePayment بازگشت وجه استرایپ
func (h *StripeHandler) RefundStripePayment(c *gin.Context) {
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

	// دریافت درگاه استرایپ از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("stripe")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه استرایپ یافت نشد",
		})
		return
	}

	// ایجاد سرویس استرایپ
	stripeService := services.NewStripeService(gateway)

	// بازگشت وجه
	err = stripeService.RefundPayment(request.TransactionID, request.Amount)
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
		"message": "بازگشت وجه استرایپ با موفقیت انجام شد",
	})
}

// ProcessStripeCallback پردازش callback استرایپ
func (h *StripeHandler) ProcessStripeCallback(c *gin.Context) {
	// دریافت پارامترهای callback
	paymentIntentID := c.Query("payment_intent")
	paymentIntentClientSecret := c.Query("payment_intent_client_secret")

	// تبدیل پارامترها به map
	params := map[string]string{
		"payment_intent":               paymentIntentID,
		"payment_intent_client_secret": paymentIntentClientSecret,
	}

	// دریافت درگاه استرایپ از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("stripe")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه استرایپ یافت نشد",
		})
		return
	}

	// ایجاد سرویس استرایپ
	stripeService := services.NewStripeService(gateway)

	// پردازش callback
	transaction, err := stripeService.ProcessCallback(params)
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
		"message": "Callback استرایپ با موفقیت پردازش شد",
	})
}

// ProcessStripeWebhook پردازش webhook استرایپ
func (h *StripeHandler) ProcessStripeWebhook(c *gin.Context) {
	// خواندن payload
	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "خطا در خواندن payload: " + err.Error(),
		})
		return
	}

	// دریافت signature header
	signature := c.GetHeader("Stripe-Signature")
	if signature == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "امضای webhook یافت نشد",
		})
		return
	}

	// دریافت درگاه استرایپ از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("stripe")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه استرایپ یافت نشد",
		})
		return
	}

	// ایجاد سرویس استرایپ
	stripeService := services.NewStripeService(gateway)

	// پردازش webhook
	transaction, err := stripeService.ProcessWebhook(payload, signature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "خطا در پردازش webhook: " + err.Error(),
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
		"message": "Webhook استرایپ با موفقیت پردازش شد",
	})
}

// GetStripeTransactions دریافت تراکنش‌های استرایپ
func (h *StripeHandler) GetStripeTransactions(c *gin.Context) {
	// دریافت پارامترهای صفحه‌بندی
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	// دریافت درگاه استرایپ
	gateway, err := h.gatewayRepo.GetFirstByType("stripe")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه استرایپ یافت نشد",
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
		"message": "تراکنش‌های استرایپ با موفقیت دریافت شد",
	})
}
