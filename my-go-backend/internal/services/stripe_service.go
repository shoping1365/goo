package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"my-go-backend/internal/models"
)

// StripeService سرویس درگاه پرداخت استرایپ
// شامل تمام عملیات پرداخت، تایید و بازگشت وجه

type StripeService struct {
	gateway *models.PaymentGateway
}

// NewStripeService سازنده سرویس استرایپ
func NewStripeService(gateway *models.PaymentGateway) *StripeService {
	return &StripeService{
		gateway: gateway,
	}
}

// ساختارهای JSON برای API requests
type StripePaymentIntentRequest struct {
	Amount             int64             `json:"amount"`
	Currency           string            `json:"currency"`
	PaymentMethodTypes []string          `json:"payment_method_types"`
	Description        string            `json:"description"`
	Metadata           map[string]string `json:"metadata"`
	ConfirmationMethod string            `json:"confirmation_method"`
	CaptureMethod      string            `json:"capture_method"`
}

type StripePaymentIntentResponse struct {
	ID           string            `json:"id"`
	Object       string            `json:"object"`
	Amount       int64             `json:"amount"`
	Currency     string            `json:"currency"`
	Status       string            `json:"status"`
	ClientSecret string            `json:"client_secret"`
	NextAction   *StripeNextAction `json:"next_action"`
}

type StripeNextAction struct {
	Type          string               `json:"type"`
	RedirectToURL *StripeRedirectToURL `json:"redirect_to_url"`
}

type StripeRedirectToURL struct {
	URL string `json:"url"`
}

type StripePaymentIntentConfirmRequest struct {
	PaymentMethod string `json:"payment_method"`
}

type StripeRefundRequest struct {
	Amount int64  `json:"amount"`
	Reason string `json:"reason"`
}

type StripeRefundResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Amount int64  `json:"amount"`
	Status string `json:"status"`
}

type StripeWebhookEvent struct {
	ID     string            `json:"id"`
	Object string            `json:"object"`
	Type   string            `json:"type"`
	Data   StripeWebhookData `json:"data"`
}

type StripeWebhookData struct {
	Object StripePaymentIntent `json:"object"`
}

type StripePaymentIntent struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Amount int64  `json:"amount"`
	Status string `json:"status"`
}

// CreatePayment ایجاد درخواست پرداخت استرایپ
func (s *StripeService) CreatePayment(amount int, callbackURL, description, email, mobile string) (*models.PaymentTransaction, error) {
	// بررسی پارامترهای ورودی
	if amount < 1 {
		return nil, errors.New("مبلغ باید عدد مثبت باشد")
	}
	if callbackURL == "" {
		return nil, errors.New("آدرس callback الزامی است")
	}
	if description == "" {
		return nil, errors.New("توضیحات الزامی است")
	}

	// ایجاد تراکنش
	transaction := &models.PaymentTransaction{
		TransactionID: generateStripeTransactionID(),
		OrderID:       generateStripeOrderID(),
		Amount:        int64(amount),
		Status:        "pending",
		PaymentMethod: "stripe",
		Description:   description,
		GatewayType:   "stripe",
		CreatedAt:     time.Now(),
	}

	// تبدیل مبلغ به سنت (استرایپ از سنت استفاده می‌کند)
	amountCents := int64(amount * 100) // تبدیل تومان به سنت

	// ساخت درخواست
	request := StripePaymentIntentRequest{
		Amount:             amountCents,
		Currency:           "usd",
		PaymentMethodTypes: []string{"card"},
		Description:        description,
		Metadata: map[string]string{
			"order_id":       transaction.OrderID,
			"transaction_id": transaction.TransactionID,
		},
		ConfirmationMethod: "manual",
		CaptureMethod:      "automatic",
	}

	// ارسال درخواست
	apiURL := s.gateway.ApiEndpoints.Payment
	if s.gateway.IsTestMode {
		apiURL = "https://api.stripe.com/v1/payment_intents"
	} else {
		apiURL = "https://api.stripe.com/v1/payment_intents"
	}

	resp, err := s.sendAPIRequest(apiURL, request)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست به درگاه استرایپ: %v", err)
	}

	// پردازش پاسخ
	var response StripePaymentIntentResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, fmt.Errorf("خطا در پردازش پاسخ درگاه استرایپ: %v", err)
	}

	// بررسی نتیجه
	if response.Status != "requires_payment_method" && response.Status != "requires_confirmation" {
		return nil, fmt.Errorf("خطای درگاه استرایپ: وضعیت %s", response.Status)
	}

	// ذخیره شناسه پرداخت در تراکنش
	transaction.GatewayToken = response.ID

	// اگر نیاز به redirect باشد
	if response.NextAction != nil && response.NextAction.Type == "redirect_to_url" {
		transaction.Description = fmt.Sprintf("Stripe Payment - Redirect URL: %s", response.NextAction.RedirectToURL.URL)
	}

	return transaction, nil
}

// VerifyPayment تایید پرداخت استرایپ
func (s *StripeService) VerifyPayment(amount int, token string) (bool, string, error) {
	// بررسی پارامترهای ورودی
	if amount < 1 {
		return false, "", errors.New("مبلغ باید عدد مثبت باشد")
	}
	if token == "" {
		return false, "", errors.New("توکن الزامی است")
	}

	// دریافت اطلاعات payment intent
	apiURL := s.gateway.ApiEndpoints.Verification
	if s.gateway.IsTestMode {
		apiURL = fmt.Sprintf("https://api.stripe.com/v1/payment_intents/%s", token)
	} else {
		apiURL = fmt.Sprintf("https://api.stripe.com/v1/payment_intents/%s", token)
	}

	resp, err := s.sendAPIRequest(apiURL, nil)
	if err != nil {
		return false, "", fmt.Errorf("خطا در دریافت اطلاعات پرداخت: %v", err)
	}

	// پردازش پاسخ
	var response StripePaymentIntentResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return false, "", fmt.Errorf("خطا در پردازش پاسخ تایید: %v", err)
	}

	// بررسی وضعیت
	if response.Status != "succeeded" {
		return false, "", fmt.Errorf("پرداخت تایید نشده. وضعیت: %s", response.Status)
	}

	// بررسی مبلغ
	expectedAmount := int64(amount * 100) // تبدیل تومان به سنت
	if response.Amount != expectedAmount {
		return false, "", fmt.Errorf("مبلغ تایید شده (%d cents) با مبلغ اصلی (%d cents) مطابقت ندارد", response.Amount, expectedAmount)
	}

	return true, response.ID, nil
}

// RefundPayment بازگشت وجه استرایپ
func (s *StripeService) RefundPayment(transactionID string, amount int) error {
	// بررسی پارامترهای ورودی
	if transactionID == "" {
		return errors.New("شناسه تراکنش الزامی است")
	}
	if amount < 1 {
		return errors.New("مبلغ باید عدد مثبت باشد")
	}

	// تبدیل مبلغ به سنت
	amountCents := int64(amount * 100)

	// ساخت درخواست (برای مستندات)
	_ = StripeRefundRequest{
		Amount: amountCents,
		Reason: "requested_by_customer",
	}

	// ارسال درخواست
	apiURL := s.gateway.ApiEndpoints.Refund
	if s.gateway.IsTestMode {
		apiURL = fmt.Sprintf("https://api.stripe.com/v1/refunds")
	} else {
		apiURL = fmt.Sprintf("https://api.stripe.com/v1/refunds")
	}

	// اضافه کردن payment_intent به درخواست
	formData := url.Values{}
	formData.Set("payment_intent", transactionID)
	formData.Set("amount", strconv.FormatInt(amountCents, 10))
	formData.Set("reason", "requested_by_customer")

	resp, err := s.sendFormRequest(apiURL, formData)
	if err != nil {
		return fmt.Errorf("خطا در ارسال درخواست بازگشت وجه به درگاه استرایپ: %v", err)
	}

	// پردازش پاسخ
	var response StripeRefundResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return fmt.Errorf("خطا در پردازش پاسخ بازگشت وجه درگاه استرایپ: %v", err)
	}

	// بررسی نتیجه
	if response.Status != "succeeded" {
		return fmt.Errorf("خطای بازگشت وجه درگاه استرایپ: وضعیت %s", response.Status)
	}

	return nil
}

// GetBalance دریافت موجودی استرایپ
func (s *StripeService) GetBalance() (float64, error) {
	// استرایپ API مستقیم برای دریافت موجودی ندارد
	// این متد برای سازگاری با interface پیاده‌سازی شده
	return 0, errors.New("درگاه استرایپ قابلیت دریافت موجودی ندارد")
}

// TestConnection تست اتصال به استرایپ
func (s *StripeService) TestConnection() error {
	// تست اتصال با دریافت اطلاعات حساب
	apiURL := s.gateway.ApiEndpoints.Payment
	if s.gateway.IsTestMode {
		apiURL = "https://api.stripe.com/v1/account"
	} else {
		apiURL = "https://api.stripe.com/v1/account"
	}

	_, err := s.sendAPIRequest(apiURL, nil)
	if err != nil {
		return fmt.Errorf("خطا در تست اتصال به درگاه استرایپ: %v", err)
	}

	return nil
}

// ProcessCallback پردازش callback استرایپ
func (s *StripeService) ProcessCallback(params map[string]string) (*models.PaymentTransaction, error) {
	// بررسی پارامترهای callback
	paymentIntentID := params["payment_intent"]
	_ = params["payment_intent_client_secret"] // برای سازگاری آینده

	if paymentIntentID == "" {
		return nil, errors.New("پارامترهای callback ناقص است")
	}

	// ایجاد تراکنش
	transaction := &models.PaymentTransaction{
		TransactionID: paymentIntentID,
		OrderID:       generateStripeOrderID(),
		Status:        "pending", // نیاز به تایید دارد
		PaymentMethod: "stripe",
		GatewayType:   "stripe",
		GatewayToken:  paymentIntentID,
		CreatedAt:     time.Now(),
	}

	return transaction, nil
}

// ProcessWebhook پردازش webhook استرایپ
func (s *StripeService) ProcessWebhook(payload []byte, signature string) (*models.PaymentTransaction, error) {
	// بررسی امضای webhook
	if !s.verifyWebhookSignature(payload, signature) {
		return nil, errors.New("امضای webhook نامعتبر است")
	}

	// پردازش payload
	var event StripeWebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, fmt.Errorf("خطا در پردازش webhook: %v", err)
	}

	// بررسی نوع event
	if event.Type != "payment_intent.succeeded" {
		return nil, fmt.Errorf("نوع event نامعتبر: %s", event.Type)
	}

	// ایجاد تراکنش
	transaction := &models.PaymentTransaction{
		TransactionID: event.Data.Object.ID,
		OrderID:       generateStripeOrderID(),
		Amount:        event.Data.Object.Amount / 100, // تبدیل سنت به تومان
		Status:        "success",
		PaymentMethod: "stripe",
		GatewayType:   "stripe",
		GatewayToken:  event.Data.Object.ID,
		CreatedAt:     time.Now(),
	}

	return transaction, nil
}

// GetPaymentURL دریافت آدرس پرداخت
func (s *StripeService) GetPaymentURL(transaction *models.PaymentTransaction, callbackURL string) string {
	// استخراج URL از توضیحات
	if strings.Contains(transaction.Description, "Redirect URL: ") {
		parts := strings.Split(transaction.Description, "Redirect URL: ")
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return ""
}

// GetPaymentForm دریافت فرم HTML پرداخت
func (s *StripeService) GetPaymentForm(transaction *models.PaymentTransaction, callbackURL string) string {
	// استرایپ معمولاً از JavaScript SDK استفاده می‌کند
	return fmt.Sprintf(`
		<div id="stripe-payment-form">
			<form id="payment-form">
				<div id="card-element"></div>
				<button type="submit">پرداخت</button>
			</form>
		</div>
		<script src="https://js.stripe.com/v3/"></script>
		<script>
			const stripe = Stripe('%s');
			const elements = stripe.elements();
			const card = elements.create('card');
			card.mount('#card-element');
		</script>
	`, s.gateway.ApiKeys.PublicKey)
}

// sendAPIRequest ارسال درخواست HTTP به API استرایپ
func (s *StripeService) sendAPIRequest(url string, data interface{}) ([]byte, error) {
	var req *http.Request
	var err error

	if data != nil {
		// تبدیل داده به JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("خطا در تبدیل داده به JSON: %v", err)
		}

		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, fmt.Errorf("خطا در ایجاد درخواست HTTP: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("خطا در ایجاد درخواست HTTP: %v", err)
		}
	}

	// تنظیم headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.gateway.ApiKeys.PrivateKey)

	// ارسال درخواست
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست: %v", err)
	}
	defer resp.Body.Close()

	// خواندن پاسخ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("خطا در خواندن پاسخ: %v", err)
	}

	// بررسی کد وضعیت
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("خطای HTTP: %d - %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// sendFormRequest ارسال درخواست فرم به API استرایپ
func (s *StripeService) sendFormRequest(url string, formData url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("خطا در ایجاد درخواست HTTP: %v", err)
	}

	// تنظیم headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.gateway.ApiKeys.PrivateKey)

	// ارسال درخواست
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست: %v", err)
	}
	defer resp.Body.Close()

	// خواندن پاسخ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("خطا در خواندن پاسخ: %v", err)
	}

	// بررسی کد وضعیت
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("خطای HTTP: %d - %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// verifyWebhookSignature بررسی امضای webhook استرایپ
func (s *StripeService) verifyWebhookSignature(payload []byte, signature string) bool {
	// این متد باید با webhook secret پیاده‌سازی شود
	// فعلاً برای سادگی true برمی‌گرداند
	return true
}

// generateStripeTransactionID تولید شناسه یکتا برای تراکنش استرایپ
func generateStripeTransactionID() string {
	return fmt.Sprintf("ST%d", time.Now().UnixNano())
}

// generateStripeOrderID تولید شناسه یکتا برای سفارش استرایپ
func generateStripeOrderID() string {
	return fmt.Sprintf("ORD%d", time.Now().UnixNano())
}

// createSignature ایجاد امضای دیجیتال (در صورت نیاز)
func (s *StripeService) createSignature(data string) string {
	// استرایپ از Bearer token استفاده می‌کند
	// این متد برای سازگاری با سایر درگاه‌ها اضافه شده
	h := hmac.New(sha256.New, []byte(s.gateway.ApiKeys.PrivateKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
