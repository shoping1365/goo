package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// SamanHandler handler مدیریت درگاه پرداخت سامان
type SamanHandler struct {
	gatewayRepo *repository.PaymentGatewayRepository
	paymentRepo *repository.PaymentTransactionRepository
}

// NewSamanHandler سازنده handler سامان
func NewSamanHandler(gatewayRepo *repository.PaymentGatewayRepository, paymentRepo *repository.PaymentTransactionRepository) *SamanHandler {
	return &SamanHandler{
		gatewayRepo: gatewayRepo,
		paymentRepo: paymentRepo,
	}
}

// GetSamanSettings دریافت تنظیمات درگاه سامان
func (h *SamanHandler) GetSamanSettings(c *gin.Context) {
	// دریافت درگاه سامان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("saman")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه سامان یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    gateway,
		"message": "تنظیمات درگاه سامان با موفقیت دریافت شد",
	})
}

// UpdateSamanSettings به‌روزرسانی تنظیمات درگاه سامان
func (h *SamanHandler) UpdateSamanSettings(c *gin.Context) {
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

	// دریافت درگاه سامان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("saman")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه سامان یافت نشد",
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
		"message": "تنظیمات درگاه سامان با موفقیت به‌روزرسانی شد",
	})
}

// TestSamanConnection تست اتصال به درگاه سامان
func (h *SamanHandler) TestSamanConnection(c *gin.Context) {
	// دریافت درگاه سامان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("saman")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه سامان یافت نشد",
		})
		return
	}

	// ایجاد سرویس سامان
	samanService := services.NewSamanService(gateway)

	// تست اتصال
	err = samanService.TestConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در تست اتصال: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "اتصال به درگاه سامان موفقیت‌آمیز بود",
	})
}

// CreateSamanPayment ایجاد پرداخت سامان
func (h *SamanHandler) CreateSamanPayment(c *gin.Context) {
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

	// دریافت درگاه سامان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("saman")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه سامان یافت نشد",
		})
		return
	}

	// بررسی فعال بودن درگاه
	if gateway.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "درگاه سامان غیرفعال است",
		})
		return
	}

	// ایجاد سرویس سامان
	samanService := services.NewSamanService(gateway)

	// ایجاد پرداخت
	transaction, err := samanService.CreatePayment(request.Amount, request.CallbackURL, request.Description, request.Email, request.Mobile)
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
	paymentURL := samanService.GetPaymentURL(transaction, request.CallbackURL)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transaction_id": transaction.TransactionID,
			"order_id":       transaction.OrderID,
			"amount":         transaction.Amount,
			"payment_url":    paymentURL,
			"payment_form":   samanService.GetPaymentForm(transaction, request.CallbackURL),
		},
		"message": "درخواست پرداخت سامان با موفقیت ایجاد شد",
	})
}

// VerifySamanPayment تایید پرداخت سامان
func (h *SamanHandler) VerifySamanPayment(c *gin.Context) {
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

	// دریافت درگاه سامان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("saman")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه سامان یافت نشد",
		})
		return
	}

	// ایجاد سرویس سامان
	samanService := services.NewSamanService(gateway)

	// تایید پرداخت
	success, refNum, err := samanService.VerifyPayment(request.Amount, request.Token)
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
		"message": "پرداخت سامان با موفقیت تایید شد",
	})
}

// RefundSamanPayment بازگشت وجه سامان
func (h *SamanHandler) RefundSamanPayment(c *gin.Context) {
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

	// دریافت درگاه سامان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("saman")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه سامان یافت نشد",
		})
		return
	}

	// ایجاد سرویس سامان
	samanService := services.NewSamanService(gateway)

	// بازگشت وجه
	err = samanService.RefundPayment(request.TransactionID, request.Amount)
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
		"message": "بازگشت وجه سامان با موفقیت انجام شد",
	})
}

// ProcessSamanCallback پردازش callback سامان
func (h *SamanHandler) ProcessSamanCallback(c *gin.Context) {
	// دریافت پارامترهای callback
	state := c.PostForm("State")
	refNum := c.PostForm("RefNum")
	resNum := c.PostForm("ResNum")
	terminalId := c.PostForm("TerminalId")
	amount := c.PostForm("Amount")

	// تبدیل پارامترها به map
	params := map[string]string{
		"State":      state,
		"RefNum":     refNum,
		"ResNum":     resNum,
		"TerminalId": terminalId,
		"Amount":     amount,
	}

	// دریافت درگاه سامان از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("saman")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه سامان یافت نشد",
		})
		return
	}

	// ایجاد سرویس سامان
	samanService := services.NewSamanService(gateway)

	// پردازش callback
	transaction, err := samanService.ProcessCallback(params)
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
		"message": "Callback سامان با موفقیت پردازش شد",
	})
}

// GetSamanTransactions دریافت تراکنش‌های سامان
func (h *SamanHandler) GetSamanTransactions(c *gin.Context) {
	// دریافت پارامترهای صفحه‌بندی
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	// دریافت درگاه سامان
	gateway, err := h.gatewayRepo.GetFirstByType("saman")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه سامان یافت نشد",
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
		"message": "تراکنش‌های سامان با موفقیت دریافت شد",
	})
}
