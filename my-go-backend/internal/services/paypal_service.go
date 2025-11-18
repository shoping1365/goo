package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"my-go-backend/internal/models"
)

// PayPalService سرویس درگاه پرداخت پی‌پال
// شامل تمام عملیات پرداخت، تایید و بازگشت وجه

type PayPalService struct {
	gateway *models.PaymentGateway
}

// NewPayPalService سازنده سرویس پی‌پال
func NewPayPalService(gateway *models.PaymentGateway) *PayPalService {
	return &PayPalService{
		gateway: gateway,
	}
}

// ساختارهای JSON برای API requests
type PayPalAccessTokenRequest struct {
	GrantType string `json:"grant_type"`
}

type PayPalAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type PayPalPaymentRequest struct {
	Intent       string              `json:"intent"`
	Payer        PayPalPayer         `json:"payer"`
	Transactions []PayPalTransaction `json:"transactions"`
	RedirectUrls PayPalRedirectUrls  `json:"redirect_urls"`
}

type PayPalPayer struct {
	PaymentMethod string `json:"payment_method"`
}

type PayPalTransaction struct {
	Amount        PayPalAmount `json:"amount"`
	Description   string       `json:"description"`
	InvoiceNumber string       `json:"invoice_number"`
}

type PayPalAmount struct {
	Total    string `json:"total"`
	Currency string `json:"currency"`
}

type PayPalRedirectUrls struct {
	ReturnURL string `json:"return_url"`
	CancelURL string `json:"cancel_url"`
}

type PayPalPaymentResponse struct {
	ID    string       `json:"id"`
	State string       `json:"state"`
	Links []PayPalLink `json:"links"`
}

type PayPalLink struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type PayPalPaymentExecuteRequest struct {
	PayerID string `json:"payer_id"`
}

type PayPalPaymentExecuteResponse struct {
	ID           string                  `json:"id"`
	State        string                  `json:"state"`
	Transactions []PayPalTransactionInfo `json:"transactions"`
}

type PayPalTransactionInfo struct {
	Amount           PayPalAmount            `json:"amount"`
	RelatedResources []PayPalRelatedResource `json:"related_resources"`
}

type PayPalRelatedResource struct {
	Sale PayPalSale `json:"sale"`
}

type PayPalSale struct {
	ID     string       `json:"id"`
	State  string       `json:"state"`
	Amount PayPalAmount `json:"amount"`
}

type PayPalRefundRequest struct {
	Amount PayPalAmount `json:"amount"`
}

type PayPalRefundResponse struct {
	ID     string       `json:"id"`
	State  string       `json:"state"`
	Amount PayPalAmount `json:"amount"`
}

// CreatePayment ایجاد درخواست پرداخت پی‌پال
func (p *PayPalService) CreatePayment(amount int, callbackURL, description, email, mobile string) (*models.PaymentTransaction, error) {
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

	// دریافت access token
	accessToken, err := p.getAccessToken()
	if err != nil {
		return nil, fmt.Errorf("خطا در دریافت access token: %v", err)
	}

	// ایجاد تراکنش
	transaction := &models.PaymentTransaction{
		TransactionID: generatePayPalTransactionID(),
		OrderID:       generatePayPalOrderID(),
		Amount:        int64(amount),
		Status:        "pending",
		PaymentMethod: "paypal",
		Description:   description,
		GatewayType:   "paypal",
		CreatedAt:     time.Now(),
	}

	// تبدیل مبلغ به دلار (فرض بر این است که مبلغ به تومان است)
	amountUSD := float64(amount) / 50000 // نرخ تقریبی دلار به تومان

	// ساخت درخواست
	request := PayPalPaymentRequest{
		Intent: "sale",
		Payer: PayPalPayer{
			PaymentMethod: "paypal",
		},
		Transactions: []PayPalTransaction{
			{
				Amount: PayPalAmount{
					Total:    fmt.Sprintf("%.2f", amountUSD),
					Currency: "USD",
				},
				Description:   description,
				InvoiceNumber: transaction.OrderID,
			},
		},
		RedirectUrls: PayPalRedirectUrls{
			ReturnURL: callbackURL + "?success=true",
			CancelURL: callbackURL + "?success=false",
		},
	}

	// ارسال درخواست
	apiURL := p.gateway.ApiEndpoints.Payment
	if p.gateway.IsTestMode {
		apiURL = "https://api-m.sandbox.paypal.com/v1/payments/payment"
	} else {
		apiURL = "https://api-m.paypal.com/v1/payments/payment"
	}

	resp, err := p.sendAPIRequest(apiURL, request, accessToken)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست به درگاه پی‌پال: %v", err)
	}

	// پردازش پاسخ
	var response PayPalPaymentResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, fmt.Errorf("خطا در پردازش پاسخ درگاه پی‌پال: %v", err)
	}

	// بررسی نتیجه
	if response.State != "created" {
		return nil, fmt.Errorf("خطای درگاه پی‌پال: وضعیت %s", response.State)
	}

	// ذخیره شناسه پرداخت در تراکنش
	transaction.GatewayToken = response.ID

	// پیدا کردن لینک تایید
	var approvalURL string
	for _, link := range response.Links {
		if link.Rel == "approval_url" {
			approvalURL = link.Href
			break
		}
	}

	// ذخیره URL در توضیحات
	transaction.Description = fmt.Sprintf("PayPal Payment - Approval URL: %s", approvalURL)

	return transaction, nil
}

// VerifyPayment تایید پرداخت پی‌پال
func (p *PayPalService) VerifyPayment(amount int, token string) (bool, string, error) {
	// بررسی پارامترهای ورودی
	if amount < 1 {
		return false, "", errors.New("مبلغ باید عدد مثبت باشد")
	}
	if token == "" {
		return false, "", errors.New("توکن الزامی است")
	}

	// دریافت access token
	accessToken, err := p.getAccessToken()
	if err != nil {
		return false, "", fmt.Errorf("خطا در دریافت access token: %v", err)
	}

	// دریافت اطلاعات پرداخت
	apiURL := p.gateway.ApiEndpoints.Verification
	if p.gateway.IsTestMode {
		apiURL = fmt.Sprintf("https://api-m.sandbox.paypal.com/v1/payments/payment/%s", token)
	} else {
		apiURL = fmt.Sprintf("https://api-m.paypal.com/v1/payments/payment/%s", token)
	}

	resp, err := p.sendAPIRequest(apiURL, nil, accessToken)
	if err != nil {
		return false, "", fmt.Errorf("خطا در دریافت اطلاعات پرداخت: %v", err)
	}

	// پردازش پاسخ
	var response PayPalPaymentExecuteResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return false, "", fmt.Errorf("خطا در پردازش پاسخ تایید: %v", err)
	}

	// بررسی وضعیت
	if response.State != "approved" {
		return false, "", fmt.Errorf("پرداخت تایید نشده. وضعیت: %s", response.State)
	}

	// بررسی مبلغ
	if len(response.Transactions) > 0 {
		responseAmount, err := strconv.ParseFloat(response.Transactions[0].Amount.Total, 64)
		if err != nil {
			return false, "", fmt.Errorf("خطا در پردازش مبلغ پاسخ: %v", err)
		}

		expectedAmount := float64(amount) / 50000 // تبدیل تومان به دلار
		if responseAmount != expectedAmount {
			return false, "", fmt.Errorf("مبلغ تایید شده (%.2f USD) با مبلغ اصلی (%.2f USD) مطابقت ندارد", responseAmount, expectedAmount)
		}
	}

	// دریافت شناسه فروش
	var saleID string
	if len(response.Transactions) > 0 && len(response.Transactions[0].RelatedResources) > 0 {
		saleID = response.Transactions[0].RelatedResources[0].Sale.ID
	}

	return true, saleID, nil
}

// RefundPayment بازگشت وجه پی‌پال
func (p *PayPalService) RefundPayment(transactionID string, amount int) error {
	// بررسی پارامترهای ورودی
	if transactionID == "" {
		return errors.New("شناسه تراکنش الزامی است")
	}
	if amount < 1 {
		return errors.New("مبلغ باید عدد مثبت باشد")
	}

	// دریافت access token
	accessToken, err := p.getAccessToken()
	if err != nil {
		return fmt.Errorf("خطا در دریافت access token: %v", err)
	}

	// تبدیل مبلغ به دلار
	amountUSD := float64(amount) / 50000

	// ساخت درخواست
	request := PayPalRefundRequest{
		Amount: PayPalAmount{
			Total:    fmt.Sprintf("%.2f", amountUSD),
			Currency: "USD",
		},
	}

	// ارسال درخواست
	apiURL := p.gateway.ApiEndpoints.Refund
	if p.gateway.IsTestMode {
		apiURL = fmt.Sprintf("https://api-m.sandbox.paypal.com/v1/payments/sale/%s/refund", transactionID)
	} else {
		apiURL = fmt.Sprintf("https://api-m.paypal.com/v1/payments/sale/%s/refund", transactionID)
	}

	resp, err := p.sendAPIRequest(apiURL, request, accessToken)
	if err != nil {
		return fmt.Errorf("خطا در ارسال درخواست بازگشت وجه به درگاه پی‌پال: %v", err)
	}

	// پردازش پاسخ
	var response PayPalRefundResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return fmt.Errorf("خطا در پردازش پاسخ بازگشت وجه درگاه پی‌پال: %v", err)
	}

	// بررسی نتیجه
	if response.State != "completed" {
		return fmt.Errorf("خطای بازگشت وجه درگاه پی‌پال: وضعیت %s", response.State)
	}

	return nil
}

// GetBalance دریافت موجودی پی‌پال
func (p *PayPalService) GetBalance() (float64, error) {
	// پی‌پال API مستقیم برای دریافت موجودی ندارد
	// این متد برای سازگاری با interface پیاده‌سازی شده
	return 0, errors.New("درگاه پی‌پال قابلیت دریافت موجودی ندارد")
}

// TestConnection تست اتصال به پی‌پال
func (p *PayPalService) TestConnection() error {
	// تست اتصال با دریافت access token
	_, err := p.getAccessToken()
	if err != nil {
		return fmt.Errorf("خطا در تست اتصال به درگاه پی‌پال: %v", err)
	}

	return nil
}

// ProcessCallback پردازش callback پی‌پال
func (p *PayPalService) ProcessCallback(params map[string]string) (*models.PaymentTransaction, error) {
	// بررسی پارامترهای callback
	paymentID := params["paymentId"]
	payerID := params["PayerID"]
	success := params["success"]

	if paymentID == "" || payerID == "" {
		return nil, errors.New("پارامترهای callback ناقص است")
	}

	// بررسی موفقیت
	if success != "true" {
		return nil, errors.New("پرداخت توسط کاربر لغو شد")
	}

	// ایجاد تراکنش
	transaction := &models.PaymentTransaction{
		TransactionID: paymentID,
		OrderID:       generatePayPalOrderID(),
		Status:        "pending", // نیاز به تایید دارد
		PaymentMethod: "paypal",
		GatewayType:   "paypal",
		GatewayToken:  paymentID,
		CreatedAt:     time.Now(),
	}

	return transaction, nil
}

// GetPaymentURL دریافت آدرس پرداخت
func (p *PayPalService) GetPaymentURL(transaction *models.PaymentTransaction, callbackURL string) string {
	// استخراج URL از توضیحات
	if strings.Contains(transaction.Description, "Approval URL: ") {
		parts := strings.Split(transaction.Description, "Approval URL: ")
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return ""
}

// GetPaymentForm دریافت فرم HTML پرداخت
func (p *PayPalService) GetPaymentForm(transaction *models.PaymentTransaction, callbackURL string) string {
	// پی‌پال از redirect استفاده می‌کند
	paymentURL := p.GetPaymentURL(transaction, callbackURL)
	return fmt.Sprintf(`
		<script>
			window.location.href = '%s';
		</script>
	`, paymentURL)
}

// getAccessToken دریافت access token از پی‌پال
func (p *PayPalService) getAccessToken() (string, error) {
	// ساخت درخواست
	request := PayPalAccessTokenRequest{
		GrantType: "client_credentials",
	}

	// تعیین آدرس API
	apiURL := "https://api-m.paypal.com/v1/oauth2/token"
	if p.gateway.IsTestMode {
		apiURL = "https://api-m.sandbox.paypal.com/v1/oauth2/token"
	}

	// ایجاد درخواست HTTP
	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("خطا در تبدیل داده به JSON: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("خطا در ایجاد درخواست HTTP: %v", err)
	}

	// تنظیم headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	// تنظیم Basic Auth
	clientID := p.gateway.ApiKeys.PublicKey
	clientSecret := p.gateway.ApiKeys.PrivateKey
	auth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	req.Header.Set("Authorization", "Basic "+auth)

	// ارسال درخواست
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("خطا در ارسال درخواست: %v", err)
	}
	defer resp.Body.Close()

	// خواندن پاسخ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("خطا در خواندن پاسخ: %v", err)
	}

	// بررسی کد وضعیت
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("خطای HTTP: %d - %s", resp.StatusCode, string(body))
	}

	// پردازش پاسخ
	var response PayPalAccessTokenResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("خطا در پردازش پاسخ access token: %v", err)
	}

	return response.AccessToken, nil
}

// sendAPIRequest ارسال درخواست HTTP به API پی‌پال
func (p *PayPalService) sendAPIRequest(url string, data interface{}, accessToken string) ([]byte, error) {
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
	req.Header.Set("Authorization", "Bearer "+accessToken)

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

// generatePayPalTransactionID تولید شناسه یکتا برای تراکنش پی‌پال
func generatePayPalTransactionID() string {
	return fmt.Sprintf("PP%d", time.Now().UnixNano())
}

// generatePayPalOrderID تولید شناسه یکتا برای سفارش پی‌پال
func generatePayPalOrderID() string {
	return fmt.Sprintf("ORD%d", time.Now().UnixNano())
}

// createSignature ایجاد امضای دیجیتال (در صورت نیاز)
func (p *PayPalService) createSignature(data string) string {
	// پی‌پال از OAuth2 استفاده می‌کند
	// این متد برای سازگاری با سایر درگاه‌ها اضافه شده
	h := hmac.New(sha256.New, []byte(p.gateway.ApiKeys.PrivateKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
