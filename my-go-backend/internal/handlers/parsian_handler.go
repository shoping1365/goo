package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// ParsianHandler handler مدیریت درگاه پرداخت پارسیان
type ParsianHandler struct {
	gatewayRepo *repository.PaymentGatewayRepository
	paymentRepo *repository.PaymentTransactionRepository
}

// NewParsianHandler سازنده handler پارسیان
func NewParsianHandler(gatewayRepo *repository.PaymentGatewayRepository, paymentRepo *repository.PaymentTransactionRepository) *ParsianHandler {
	return &ParsianHandler{
		gatewayRepo: gatewayRepo,
		paymentRepo: paymentRepo,
	}
}

// GetParsianSettings دریافت تنظیمات درگاه پارسیان
func (h *ParsianHandler) GetParsianSettings(c *gin.Context) {
	// دریافت درگاه پارسیان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("parsian")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پارسیان یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    gateway,
		"message": "تنظیمات درگاه پارسیان با موفقیت دریافت شد",
	})
}

// UpdateParsianSettings به‌روزرسانی تنظیمات درگاه پارسیان
func (h *ParsianHandler) UpdateParsianSettings(c *gin.Context) {
	var request struct {
		MerchantId string  `json:"merchant_id" binding:"required"`
		PrivateKey string  `json:"private_key" binding:"required"`
		PublicKey  string  `json:"public_key"`
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

	// دریافت درگاه پارسیان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("parsian")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پارسیان یافت نشد",
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
		"message": "تنظیمات درگاه پارسیان با موفقیت به‌روزرسانی شد",
	})
}

// TestParsianConnection تست اتصال به درگاه پارسیان
func (h *ParsianHandler) TestParsianConnection(c *gin.Context) {
	// دریافت درگاه پارسیان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("parsian")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پارسیان یافت نشد",
		})
		return
	}

	// ایجاد سرویس پارسیان
	parsianService := services.NewParsianService(gateway)

	// تست اتصال
	err = parsianService.TestConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در تست اتصال: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "اتصال به درگاه پارسیان موفقیت‌آمیز بود",
	})
}

// CreateParsianPayment ایجاد پرداخت پارسیان
func (h *ParsianHandler) CreateParsianPayment(c *gin.Context) {
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

	// دریافت درگاه پارسیان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("parsian")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پارسیان یافت نشد",
		})
		return
	}

	// بررسی فعال بودن درگاه
	if gateway.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "درگاه پارسیان غیرفعال است",
		})
		return
	}

	// ایجاد سرویس پارسیان
	parsianService := services.NewParsianService(gateway)

	// ایجاد پرداخت
	transaction, err := parsianService.CreatePayment(request.Amount, request.CallbackURL, request.Description, request.Email, request.Mobile)
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
	paymentURL := parsianService.GetPaymentURL(transaction, request.CallbackURL)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transaction_id": transaction.TransactionID,
			"order_id":       transaction.OrderID,
			"amount":         transaction.Amount,
			"payment_url":    paymentURL,
			"payment_form":   parsianService.GetPaymentForm(transaction, request.CallbackURL),
		},
		"message": "درخواست پرداخت پارسیان با موفقیت ایجاد شد",
	})
}

// VerifyParsianPayment تایید پرداخت پارسیان
func (h *ParsianHandler) VerifyParsianPayment(c *gin.Context) {
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

	// دریافت درگاه پارسیان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("parsian")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پارسیان یافت نشد",
		})
		return
	}

	// ایجاد سرویس پارسیان
	parsianService := services.NewParsianService(gateway)

	// تایید پرداخت
	success, refNum, err := parsianService.VerifyPayment(request.Amount, request.Token)
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
		"message": "پرداخت پارسیان با موفقیت تایید شد",
	})
}

// RefundParsianPayment بازگشت وجه پارسیان
func (h *ParsianHandler) RefundParsianPayment(c *gin.Context) {
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

	// دریافت درگاه پارسیان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("parsian")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پارسیان یافت نشد",
		})
		return
	}

	// ایجاد سرویس پارسیان
	parsianService := services.NewParsianService(gateway)

	// بازگشت وجه
	err = parsianService.RefundPayment(request.TransactionID, request.Amount)
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
		"message": "بازگشت وجه پارسیان با موفقیت انجام شد",
	})
}

// ProcessParsianCallback پردازش callback پارسیان
func (h *ParsianHandler) ProcessParsianCallback(c *gin.Context) {
	// دریافت پارامترهای callback
	token := c.PostForm("Token")
	orderId := c.PostForm("OrderId")
	status := c.PostForm("Status")
	amount := c.PostForm("Amount")

	// تبدیل پارامترها به map
	params := map[string]string{
		"Token":   token,
		"OrderId": orderId,
		"Status":  status,
		"Amount":  amount,
	}

	// دریافت درگاه پارسیان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("parsian")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پارسیان یافت نشد",
		})
		return
	}

	// ایجاد سرویس پارسیان
	parsianService := services.NewParsianService(gateway)

	// پردازش callback
	transaction, err := parsianService.ProcessCallback(params)
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
		"message": "Callback پارسیان با موفقیت پردازش شد",
	})
}

// GetParsianTransactions دریافت تراکنش‌های پارسیان
func (h *ParsianHandler) GetParsianTransactions(c *gin.Context) {
	// دریافت پارامترهای صفحه‌بندی
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	// دریافت درگاه پارسیان
	gateway, err := h.gatewayRepo.GetFirstByType("parsian")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه پارسیان یافت نشد",
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
		"message": "تراکنش‌های پارسیان با موفقیت دریافت شد",
	})
}
