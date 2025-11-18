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
	"strconv"
	"time"

	"my-go-backend/internal/models"
)

// ParsianService سرویس درگاه پرداخت بانک پارسیان
// شامل تمام عملیات پرداخت، تایید و بازگشت وجه

type ParsianService struct {
	gateway *models.PaymentGateway
}

// NewParsianService سازنده سرویس پارسیان
func NewParsianService(gateway *models.PaymentGateway) *ParsianService {
	return &ParsianService{
		gateway: gateway,
	}
}

// ساختارهای JSON برای API requests
type ParsianPaymentRequest struct {
	Pin         string `json:"Pin"`
	Amount      string `json:"Amount"`
	OrderId     string `json:"OrderId"`
	RedirectUrl string `json:"RedirectUrl"`
	CellNumber  string `json:"CellNumber,omitempty"`
}

type ParsianPaymentResponse struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
	Token   string `json:"Token"`
	OrderId string `json:"OrderId"`
}

type ParsianVerificationRequest struct {
	Pin     string `json:"Pin"`
	Token   string `json:"Token"`
	OrderId string `json:"OrderId"`
}

type ParsianVerificationResponse struct {
	Status     int    `json:"Status"`
	Message    string `json:"Message"`
	Amount     string `json:"Amount"`
	OrderId    string `json:"OrderId"`
	Token      string `json:"Token"`
	RRN        string `json:"RRN"`
	CardNumber string `json:"CardNumber"`
	TraceNo    string `json:"TraceNo"`
}

type ParsianRefundRequest struct {
	Pin     string `json:"Pin"`
	Token   string `json:"Token"`
	OrderId string `json:"OrderId"`
}

type ParsianRefundResponse struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}

// CreatePayment ایجاد درخواست پرداخت پارسیان
func (p *ParsianService) CreatePayment(amount int, callbackURL, description, email, mobile string) (*models.PaymentTransaction, error) {
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
		TransactionID: generateParsianTransactionID(),
		OrderID:       generateParsianOrderID(),
		Amount:        int64(amount),
		Status:        "pending",
		PaymentMethod: "parsian",
		Description:   description,
		GatewayType:   "parsian",
		CreatedAt:     time.Now(),
	}

	// آماده‌سازی داده‌های درخواست
	pin := p.gateway.MerchantId
	amountStr := strconv.Itoa(amount)
	orderId := transaction.OrderID

	// ساخت درخواست
	request := ParsianPaymentRequest{
		Pin:         pin,
		Amount:      amountStr,
		OrderId:     orderId,
		RedirectUrl: callbackURL,
		CellNumber:  mobile,
	}

	// ارسال درخواست
	apiURL := p.gateway.ApiEndpoints.Payment
	if p.gateway.IsTestMode {
		apiURL = "https://pec.shaparak.ir/NewIPGServices/Sale/SaleService.asmx/SalePaymentRequest"
	} else {
		apiURL = "https://pec.shaparak.ir/NewIPGServices/Sale/SaleService.asmx/SalePaymentRequest"
	}

	resp, err := p.sendAPIRequest(apiURL, request)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست به درگاه پارسیان: %v", err)
	}

	// پردازش پاسخ
	var response ParsianPaymentResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, fmt.Errorf("خطا در پردازش پاسخ درگاه پارسیان: %v", err)
	}

	// بررسی نتیجه
	if response.Status != 0 {
		return nil, fmt.Errorf("خطای درگاه پارسیان: %s", response.Message)
	}

	// ذخیره توکن در تراکنش
	transaction.GatewayToken = response.Token

	return transaction, nil
}

// VerifyPayment تایید پرداخت پارسیان
func (p *ParsianService) VerifyPayment(amount int, token string) (bool, string, error) {
	// بررسی پارامترهای ورودی
	if amount < 1 {
		return false, "", errors.New("مبلغ باید عدد مثبت باشد")
	}
	if token == "" {
		return false, "", errors.New("توکن الزامی است")
	}

	// آماده‌سازی داده‌های درخواست
	pin := p.gateway.MerchantId

	// ساخت درخواست
	request := ParsianVerificationRequest{
		Pin:     pin,
		Token:   token,
		OrderId: "", // باید از تراکنش دریافت شود
	}

	// ارسال درخواست
	apiURL := p.gateway.ApiEndpoints.Verification
	if p.gateway.IsTestMode {
		apiURL = "https://pec.shaparak.ir/NewIPGServices/Confirm/ConfirmService.asmx/ConfirmPayment"
	} else {
		apiURL = "https://pec.shaparak.ir/NewIPGServices/Confirm/ConfirmService.asmx/ConfirmPayment"
	}

	resp, err := p.sendAPIRequest(apiURL, request)
	if err != nil {
		return false, "", fmt.Errorf("خطا در ارسال درخواست تایید به درگاه پارسیان: %v", err)
	}

	// پردازش پاسخ
	var response ParsianVerificationResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return false, "", fmt.Errorf("خطا در پردازش پاسخ تایید درگاه پارسیان: %v", err)
	}

	// بررسی نتیجه
	if response.Status != 0 {
		return false, "", fmt.Errorf("خطای تایید درگاه پارسیان: %s", response.Message)
	}

	// بررسی مبلغ
	responseAmount, err := strconv.Atoi(response.Amount)
	if err != nil {
		return false, "", fmt.Errorf("خطا در پردازش مبلغ پاسخ: %v", err)
	}

	if responseAmount != amount {
		return false, "", fmt.Errorf("مبلغ تایید شده (%d) با مبلغ اصلی (%d) مطابقت ندارد", responseAmount, amount)
	}

	return true, response.Token, nil
}

// RefundPayment بازگشت وجه پارسیان
func (p *ParsianService) RefundPayment(transactionID string, amount int) error {
	// بررسی پارامترهای ورودی
	if transactionID == "" {
		return errors.New("شناسه تراکنش الزامی است")
	}
	if amount < 1 {
		return errors.New("مبلغ باید عدد مثبت باشد")
	}

	// آماده‌سازی داده‌های درخواست
	pin := p.gateway.MerchantId

	// ساخت درخواست
	request := ParsianRefundRequest{
		Pin:     pin,
		Token:   transactionID,
		OrderId: "", // باید از تراکنش دریافت شود
	}

	// ارسال درخواست
	apiURL := p.gateway.ApiEndpoints.Refund
	if p.gateway.IsTestMode {
		apiURL = "https://pec.shaparak.ir/NewIPGServices/Reverse/ReversalService.asmx/ReversalRequest"
	} else {
		apiURL = "https://pec.shaparak.ir/NewIPGServices/Reverse/ReversalService.asmx/ReversalRequest"
	}

	resp, err := p.sendAPIRequest(apiURL, request)
	if err != nil {
		return fmt.Errorf("خطا در ارسال درخواست بازگشت وجه به درگاه پارسیان: %v", err)
	}

	// پردازش پاسخ
	var response ParsianRefundResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return fmt.Errorf("خطا در پردازش پاسخ بازگشت وجه درگاه پارسیان: %v", err)
	}

	// بررسی نتیجه
	if response.Status != 0 {
		return fmt.Errorf("خطای بازگشت وجه درگاه پارسیان: %s", response.Message)
	}

	return nil
}

// GetBalance دریافت موجودی پارسیان
func (p *ParsianService) GetBalance() (float64, error) {
	// پارسیان API مستقیم برای دریافت موجودی ندارد
	// این متد برای سازگاری با interface پیاده‌سازی شده
	return 0, errors.New("درگاه پارسیان قابلیت دریافت موجودی ندارد")
}

// TestConnection تست اتصال به پارسیان
func (p *ParsianService) TestConnection() error {
	// تست اتصال با ارسال درخواست تست
	testRequest := ParsianPaymentRequest{
		Pin:         p.gateway.MerchantId,
		Amount:      "1000",
		OrderId:     "test_" + strconv.FormatInt(time.Now().Unix(), 10),
		RedirectUrl: "https://example.com/test",
	}

	apiURL := p.gateway.ApiEndpoints.Payment
	if p.gateway.IsTestMode {
		apiURL = "https://pec.shaparak.ir/NewIPGServices/Sale/SaleService.asmx/SalePaymentRequest"
	}

	_, err := p.sendAPIRequest(apiURL, testRequest)
	if err != nil {
		return fmt.Errorf("خطا در تست اتصال به درگاه پارسیان: %v", err)
	}

	return nil
}

// ProcessCallback پردازش callback پارسیان
func (p *ParsianService) ProcessCallback(params map[string]string) (*models.PaymentTransaction, error) {
	// بررسی پارامترهای callback
	token := params["Token"]
	orderId := params["OrderId"]
	status := params["Status"]
	amount := params["Amount"]

	if token == "" || orderId == "" || status == "" || amount == "" {
		return nil, errors.New("پارامترهای callback ناقص است")
	}

	// بررسی وضعیت
	if status != "0" {
		return nil, fmt.Errorf("پرداخت ناموفق بود. وضعیت: %s", status)
	}

	// ایجاد تراکنش
	transaction := &models.PaymentTransaction{
		TransactionID: token,
		OrderID:       orderId,
		Status:        "pending", // نیاز به تایید دارد
		PaymentMethod: "parsian",
		GatewayType:   "parsian",
		GatewayToken:  token,
		CreatedAt:     time.Now(),
	}

	// تبدیل مبلغ
	amountInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("خطا در پردازش مبلغ: %v", err)
	}
	transaction.Amount = amountInt

	return transaction, nil
}

// GetPaymentURL دریافت آدرس پرداخت
func (p *ParsianService) GetPaymentURL(transaction *models.PaymentTransaction, callbackURL string) string {
	// پارسیان از فرم HTML برای پرداخت استفاده می‌کند
	baseURL := p.gateway.ApiEndpoints.Payment
	if p.gateway.IsTestMode {
		baseURL = "https://pec.shaparak.ir/NewIPGServices/Sale/SaleService.asmx/SalePaymentRequest"
	}
	return baseURL
}

// GetPaymentForm دریافت فرم HTML پرداخت
func (p *ParsianService) GetPaymentForm(transaction *models.PaymentTransaction, callbackURL string) string {
	// فرم HTML برای ارسال به پارسیان
	form := fmt.Sprintf(`
		<form id="parsian-payment-form" method="post" action="%s">
			<input type="hidden" name="Token" value="%s" />
			<input type="hidden" name="RedirectURL" value="%s" />
		</form>
		<script>document.getElementById('parsian-payment-form').submit();</script>
	`,
		"https://pec.shaparak.ir/NewIPGServices/Sale/SaleService.asmx/GoToPayment",
		transaction.GatewayToken,
		callbackURL,
	)
	return form
}

// sendAPIRequest ارسال درخواست HTTP به API پارسیان
func (p *ParsianService) sendAPIRequest(url string, data interface{}) ([]byte, error) {
	// تبدیل داده به JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("خطا در تبدیل داده به JSON: %v", err)
	}

	// ایجاد درخواست HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("خطا در ایجاد درخواست HTTP: %v", err)
	}

	// تنظیم headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("خطای HTTP: %d - %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// generateParsianTransactionID تولید شناسه یکتا برای تراکنش پارسیان
func generateParsianTransactionID() string {
	return fmt.Sprintf("PAR%d", time.Now().UnixNano())
}

// generateParsianOrderID تولید شناسه یکتا برای سفارش پارسیان
func generateParsianOrderID() string {
	return fmt.Sprintf("ORD%d", time.Now().UnixNano())
}

// createSignature ایجاد امضای دیجیتال (در صورت نیاز)
func (p *ParsianService) createSignature(data string) string {
	// پارسیان معمولاً از امضای دیجیتال استفاده نمی‌کند
	// این متد برای سازگاری با سایر درگاه‌ها اضافه شده
	h := hmac.New(sha256.New, []byte(p.gateway.ApiKeys.PrivateKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
