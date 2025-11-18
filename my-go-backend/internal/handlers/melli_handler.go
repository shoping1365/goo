package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// MelliHandler handler مدیریت درگاه پرداخت ملی
type MelliHandler struct {
	gatewayRepo *repository.PaymentGatewayRepository
	paymentRepo *repository.PaymentTransactionRepository
}

// NewMelliHandler سازنده handler ملی
func NewMelliHandler(gatewayRepo *repository.PaymentGatewayRepository, paymentRepo *repository.PaymentTransactionRepository) *MelliHandler {
	return &MelliHandler{
		gatewayRepo: gatewayRepo,
		paymentRepo: paymentRepo,
	}
}

// GetMelliSettings دریافت تنظیمات درگاه ملی
func (h *MelliHandler) GetMelliSettings(c *gin.Context) {
	// دریافت درگاه ملی از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("melli")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملی یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    gateway,
		"message": "تنظیمات درگاه ملی با موفقیت دریافت شد",
	})
}

// UpdateMelliSettings به‌روزرسانی تنظیمات درگاه ملی
func (h *MelliHandler) UpdateMelliSettings(c *gin.Context) {
	var request struct {
		MerchantId  string `json:"merchant_id" binding:"required"`
		PublicKey   string `json:"public_key" binding:"required"`
		PrivateKey  string `json:"private_key" binding:"required"`
		CallbackUrl string `json:"callback_url" binding:"required"`
		IsTestMode  bool   `json:"is_test_mode"`
		IsActive    bool   `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	// دریافت درگاه ملی از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("melli")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملی یافت نشد",
		})
		return
	}

	// به‌روزرسانی تنظیمات
	gateway.MerchantId = request.MerchantId
	gateway.ApiKeys.PublicKey = request.PublicKey
	gateway.ApiKeys.PrivateKey = request.PrivateKey
	gateway.ApiEndpoints.Callback = request.CallbackUrl
	gateway.IsTestMode = request.IsTestMode

	// تغییر وضعیت فعال/غیرفعال
	if request.IsActive {
		gateway.Status = "active"
	} else {
		gateway.Status = "inactive"
	}

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
		"message": "تنظیمات درگاه ملی با موفقیت به‌روزرسانی شد",
	})
}

// TestMelliConnection تست اتصال درگاه ملی
func (h *MelliHandler) TestMelliConnection(c *gin.Context) {
	var request struct {
		MerchantId string `json:"merchant_id" binding:"required"`
		PublicKey  string `json:"public_key" binding:"required"`
		PrivateKey string `json:"private_key" binding:"required"`
		IsTestMode bool   `json:"is_test_mode"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	// ایجاد درگاه موقت برای تست
	testGateway := &models.PaymentGateway{
		MerchantId: request.MerchantId,
		ApiKeys: models.PaymentGatewayApiKeys{
			PublicKey:  request.PublicKey,
			PrivateKey: request.PrivateKey,
		},
		IsTestMode: request.IsTestMode,
	}

	// ایجاد سرویس ملی
	melliService := services.NewMelliService(testGateway)

	// تست اتصال
	testErr := melliService.TestConnection()

	var success bool
	var message string
	if testErr != nil {
		success = false
		message = testErr.Error()
	} else {
		success = true
		message = "اتصال به درگاه ملی با موفقیت برقرار شد"
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": message,
	})
}

// CreateMelliPayment ایجاد پرداخت ملی
func (h *MelliHandler) CreateMelliPayment(c *gin.Context) {
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

	// دریافت درگاه ملی از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("melli")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملی یافت نشد",
		})
		return
	}

	// بررسی فعال بودن درگاه
	if gateway.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "درگاه ملی غیرفعال است",
		})
		return
	}

	// ایجاد سرویس ملی
	melliService := services.NewMelliService(gateway)

	// ایجاد پرداخت
	transaction, err := melliService.CreatePayment(request.Amount, request.CallbackURL, request.Description, request.Email, request.Mobile)
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
	paymentURL := melliService.GetPaymentURL(transaction, request.CallbackURL)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transaction_id": transaction.TransactionID,
			"order_id":       transaction.OrderID,
			"amount":         transaction.Amount,
			"payment_url":    paymentURL,
			"payment_form":   melliService.GetPaymentForm(transaction, request.CallbackURL),
		},
		"message": "درخواست پرداخت ملی با موفقیت ایجاد شد",
	})
}

// VerifyMelliPayment تایید پرداخت ملی
func (h *MelliHandler) VerifyMelliPayment(c *gin.Context) {
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

	// دریافت درگاه ملی از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("melli")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملی یافت نشد",
		})
		return
	}

	// ایجاد سرویس ملی
	melliService := services.NewMelliService(gateway)

	// تایید پرداخت
	success, refNum, err := melliService.VerifyPayment(request.Amount, request.Token)
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
		"message": "پرداخت ملی با موفقیت تایید شد",
	})
}

// RefundMelliPayment بازگشت وجه ملی
func (h *MelliHandler) RefundMelliPayment(c *gin.Context) {
	var request struct {
		Token  string `json:"token" binding:"required"`
		Amount int    `json:"amount" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	// دریافت درگاه ملی از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("melli")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملی یافت نشد",
		})
		return
	}

	// ایجاد سرویس ملی
	melliService := services.NewMelliService(gateway)

	// بازگشت وجه
	err = melliService.RefundPayment(request.Token, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در بازگشت وجه: " + err.Error(),
		})
		return
	}

	// به‌روزرسانی وضعیت تراکنش در دیتابیس
	transaction, err := h.paymentRepo.GetByTransactionID(request.Token)
	if err == nil {
		transaction.Status = "refunded"
		h.paymentRepo.Update(transaction)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"token":  request.Token,
			"amount": request.Amount,
		},
		"message": "بازگشت وجه ملی با موفقیت انجام شد",
	})
}

// ProcessMelliCallback پردازش callback ملی
func (h *MelliHandler) ProcessMelliCallback(c *gin.Context) {
	// دریافت پارامترهای callback
	resCode := c.PostForm("ResCode")
	orderId := c.PostForm("OrderId")
	token := c.PostForm("token")

	if resCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "کد نتیجه مشخص نشده",
		})
		return
	}

	// دریافت درگاه ملی از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("melli")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملی یافت نشد",
		})
		return
	}

	// ایجاد سرویس ملی
	melliService := services.NewMelliService(gateway)

	// پردازش callback
	formData := map[string]string{
		"ResCode": resCode,
		"OrderId": orderId,
		"token":   token,
	}

	success, refNum, orderId, err := melliService.ProcessCallback(formData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در پردازش callback: " + err.Error(),
		})
		return
	}

	// به‌روزرسانی وضعیت تراکنش در دیتابیس
	if success && refNum != "" {
		transaction, err := h.paymentRepo.GetByTransactionID(token)
		if err == nil {
			transaction.Status = "success"
			h.paymentRepo.Update(transaction)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"data": gin.H{
			"ref_num":  refNum,
			"order_id": orderId,
			"token":    token,
			"res_code": resCode,
		},
		"message": "Callback ملی با موفقیت پردازش شد",
	})
}

// GetMelliTransactions دریافت تراکنش‌های ملی
func (h *MelliHandler) GetMelliTransactions(c *gin.Context) {
	// دریافت پارامترهای صفحه‌بندی
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	// دریافت درگاه ملی
	gateway, err := h.gatewayRepo.GetFirstByType("melli")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملی یافت نشد",
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
		"message": "تراکنش‌های ملی با موفقیت دریافت شد",
	})
}
