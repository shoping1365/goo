package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// MellatHandler handler مدیریت درگاه پرداخت ملت
type MellatHandler struct {
	gatewayRepo *repository.PaymentGatewayRepository
	paymentRepo *repository.PaymentTransactionRepository
}

// NewMellatHandler سازنده handler ملت
func NewMellatHandler(gatewayRepo *repository.PaymentGatewayRepository, paymentRepo *repository.PaymentTransactionRepository) *MellatHandler {
	return &MellatHandler{
		gatewayRepo: gatewayRepo,
		paymentRepo: paymentRepo,
	}
}

// GetMellatSettings دریافت تنظیمات درگاه ملت
func (h *MellatHandler) GetMellatSettings(c *gin.Context) {
	// دریافت درگاه ملت از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("mellat")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملت یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    gateway,
		"message": "تنظیمات درگاه ملت با موفقیت دریافت شد",
	})
}

// UpdateMellatSettings به‌روزرسانی تنظیمات درگاه ملت
func (h *MellatHandler) UpdateMellatSettings(c *gin.Context) {
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

	// دریافت درگاه ملت از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("mellat")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملت یافت نشد",
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
		"message": "تنظیمات درگاه ملت با موفقیت به‌روزرسانی شد",
	})
}

// TestMellatConnection تست اتصال درگاه ملت
func (h *MellatHandler) TestMellatConnection(c *gin.Context) {
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

	// ایجاد سرویس ملت
	mellatService := services.NewMellatService(testGateway)

	// تست اتصال
	testErr := mellatService.TestConnection()

	var success bool
	var message string
	if testErr != nil {
		success = false
		message = testErr.Error()
	} else {
		success = true
		message = "اتصال به درگاه ملت با موفقیت برقرار شد"
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": message,
	})
}

// CreateMellatPayment ایجاد پرداخت ملت
func (h *MellatHandler) CreateMellatPayment(c *gin.Context) {
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

	// دریافت درگاه ملت از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("mellat")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملت یافت نشد",
		})
		return
	}

	// بررسی فعال بودن درگاه
	if gateway.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "درگاه ملت غیرفعال است",
		})
		return
	}

	// ایجاد سرویس ملت
	mellatService := services.NewMellatService(gateway)

	// ایجاد پرداخت
	transaction, err := mellatService.CreatePayment(request.Amount, request.CallbackURL, request.Description, request.Email, request.Mobile)
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
	paymentURL := mellatService.GetPaymentURL(transaction, request.CallbackURL)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"transaction_id": transaction.TransactionID,
			"order_id":       transaction.OrderID,
			"amount":         transaction.Amount,
			"payment_url":    paymentURL,
			"payment_form":   mellatService.GetPaymentForm(transaction, request.CallbackURL),
		},
		"message": "درخواست پرداخت ملت با موفقیت ایجاد شد",
	})
}

// VerifyMellatPayment تایید پرداخت ملت
func (h *MellatHandler) VerifyMellatPayment(c *gin.Context) {
	var request struct {
		Amount int    `json:"amount" binding:"required,min=1"`
		RefId  string `json:"ref_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	// دریافت درگاه ملت از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("mellat")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملت یافت نشد",
		})
		return
	}

	// ایجاد سرویس ملت
	mellatService := services.NewMellatService(gateway)

	// تایید پرداخت
	success, authority, err := mellatService.VerifyPayment(request.Amount, request.RefId)
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
	transaction, err := h.paymentRepo.GetByTransactionID(authority)
	if err == nil {
		transaction.Status = "success"
		h.paymentRepo.Update(transaction)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"authority": authority,
			"amount":    request.Amount,
		},
		"message": "پرداخت ملت با موفقیت تایید شد",
	})
}

// RefundMellatPayment بازگشت وجه ملت
func (h *MellatHandler) RefundMellatPayment(c *gin.Context) {
	var request struct {
		RefId  string `json:"ref_id" binding:"required"`
		Amount int    `json:"amount" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "پارامترهای ورودی نامعتبر: " + err.Error(),
		})
		return
	}

	// دریافت درگاه ملت از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("mellat")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملت یافت نشد",
		})
		return
	}

	// ایجاد سرویس ملت
	mellatService := services.NewMellatService(gateway)

	// بازگشت وجه
	err = mellatService.RefundPayment(request.RefId, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در بازگشت وجه: " + err.Error(),
		})
		return
	}

	// به‌روزرسانی وضعیت تراکنش در دیتابیس
	transaction, err := h.paymentRepo.GetByTransactionID(request.RefId)
	if err == nil {
		transaction.Status = "refunded"
		h.paymentRepo.Update(transaction)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"ref_id": request.RefId,
			"amount": request.Amount,
		},
		"message": "بازگشت وجه ملت با موفقیت انجام شد",
	})
}

// ProcessMellatCallback پردازش callback ملت
func (h *MellatHandler) ProcessMellatCallback(c *gin.Context) {
	// دریافت پارامترهای callback
	refId := c.PostForm("RefId")
	resCode := c.PostForm("ResCode")
	saleOrderId := c.PostForm("SaleOrderId")
	saleReferenceId := c.PostForm("SaleReferenceId")

	if resCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "کد نتیجه مشخص نشده",
		})
		return
	}

	// دریافت درگاه ملت از دیتابیس
	gateway, err := h.gatewayRepo.GetFirstByType("mellat")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملت یافت نشد",
		})
		return
	}

	// ایجاد سرویس ملت
	mellatService := services.NewMellatService(gateway)

	// پردازش callback
	formData := map[string]string{
		"RefId":           refId,
		"ResCode":         resCode,
		"SaleOrderId":     saleOrderId,
		"SaleReferenceId": saleReferenceId,
	}

	success, authority, orderId, err := mellatService.ProcessCallback(formData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در پردازش callback: " + err.Error(),
		})
		return
	}

	// به‌روزرسانی وضعیت تراکنش در دیتابیس
	if success && authority != "" {
		transaction, err := h.paymentRepo.GetByTransactionID(authority)
		if err == nil {
			transaction.Status = "success"
			h.paymentRepo.Update(transaction)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"data": gin.H{
			"authority": authority,
			"order_id":  orderId,
			"ref_id":    refId,
			"res_code":  resCode,
		},
		"message": "Callback ملت با موفقیت پردازش شد",
	})
}

// GetMellatTransactions دریافت تراکنش‌های ملت
func (h *MellatHandler) GetMellatTransactions(c *gin.Context) {
	// دریافت پارامترهای صفحه‌بندی
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	// دریافت درگاه ملت
	gateway, err := h.gatewayRepo.GetFirstByType("mellat")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "درگاه ملت یافت نشد",
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
		"message": "تراکنش‌های ملت با موفقیت دریافت شد",
	})
}
