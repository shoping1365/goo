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

// SamanService سرویس درگاه پرداخت بانک سامان
// شامل تمام عملیات پرداخت، تایید و بازگشت وجه

type SamanService struct {
	gateway *models.PaymentGateway
}

// NewSamanService سازنده سرویس سامان
func NewSamanService(gateway *models.PaymentGateway) *SamanService {
	return &SamanService{
		gateway: gateway,
	}
}

// ساختارهای JSON برای API requests
type SamanPaymentRequest struct {
	TerminalId  string `json:"TerminalId"`
	Amount      string `json:"Amount"`
	ResNum      string `json:"ResNum"`
	RedirectUrl string `json:"RedirectUrl"`
	CellNumber  string `json:"CellNumber,omitempty"`
	Action      string `json:"Action"`
}

type SamanPaymentResponse struct {
	Token       string `json:"Token"`
	ResCode     int    `json:"ResCode"`
	Description string `json:"Description"`
}

type SamanVerificationRequest struct {
	RefNum string `json:"RefNum"`
	TermId string `json:"TermId"`
}

type SamanVerificationResponse struct {
	ResCode     int    `json:"ResCode"`
	Amount      string `json:"Amount"`
	Description string `json:"Description"`
	RefNum      string `json:"RefNum"`
	TermId      string `json:"TermId"`
	TraceNo     string `json:"TraceNo"`
}

type SamanRefundRequest struct {
	RefNum string `json:"RefNum"`
	TermId string `json:"TermId"`
}

type SamanRefundResponse struct {
	ResCode     int    `json:"ResCode"`
	Description string `json:"Description"`
}

// CreatePayment ایجاد درخواست پرداخت سامان
func (s *SamanService) CreatePayment(amount int, callbackURL, description, email, mobile string) (*models.PaymentTransaction, error) {
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
		TransactionID: generateSamanTransactionID(),
		OrderID:       generateSamanOrderID(),
		Amount:        int64(amount),
		Status:        "pending",
		PaymentMethod: "saman",
		Description:   description,
		GatewayType:   "saman",
		CreatedAt:     time.Now(),
	}

	// آماده‌سازی داده‌های درخواست
	terminalId := s.gateway.MerchantId
	amountStr := strconv.Itoa(amount)
	resNum := transaction.OrderID

	// ساخت درخواست
	request := SamanPaymentRequest{
		TerminalId:  terminalId,
		Amount:      amountStr,
		ResNum:      resNum,
		RedirectUrl: callbackURL,
		CellNumber:  mobile,
		Action:      "Token",
	}

	// ارسال درخواست
	apiURL := s.gateway.ApiEndpoints.Payment
	if s.gateway.IsTestMode {
		apiURL = "https://sep.shaparak.ir/api/v1/Payment/GetToken"
	} else {
		apiURL = "https://sep.shaparak.ir/api/v1/Payment/GetToken"
	}

	resp, err := s.sendAPIRequest(apiURL, request)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست به درگاه سامان: %v", err)
	}

	// پردازش پاسخ
	var response SamanPaymentResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, fmt.Errorf("خطا در پردازش پاسخ درگاه سامان: %v", err)
	}

	// بررسی نتیجه
	if response.ResCode != 0 {
		return nil, fmt.Errorf("خطای درگاه سامان: %s", response.Description)
	}

	// ذخیره توکن در تراکنش
	transaction.GatewayToken = response.Token

	return transaction, nil
}

// VerifyPayment تایید پرداخت سامان
func (s *SamanService) VerifyPayment(amount int, token string) (bool, string, error) {
	// بررسی پارامترهای ورودی
	if amount < 1 {
		return false, "", errors.New("مبلغ باید عدد مثبت باشد")
	}
	if token == "" {
		return false, "", errors.New("توکن الزامی است")
	}

	// آماده‌سازی داده‌های درخواست
	terminalId := s.gateway.MerchantId

	// ساخت درخواست
	request := SamanVerificationRequest{
		RefNum: token,
		TermId: terminalId,
	}

	// ارسال درخواست
	apiURL := s.gateway.ApiEndpoints.Verification
	if s.gateway.IsTestMode {
		apiURL = "https://sep.shaparak.ir/api/v1/Payment/VerifyPayment"
	} else {
		apiURL = "https://sep.shaparak.ir/api/v1/Payment/VerifyPayment"
	}

	resp, err := s.sendAPIRequest(apiURL, request)
	if err != nil {
		return false, "", fmt.Errorf("خطا در ارسال درخواست تایید به درگاه سامان: %v", err)
	}

	// پردازش پاسخ
	var response SamanVerificationResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return false, "", fmt.Errorf("خطا در پردازش پاسخ تایید درگاه سامان: %v", err)
	}

	// بررسی نتیجه
	if response.ResCode != 0 {
		return false, "", fmt.Errorf("خطای تایید درگاه سامان: %s", response.Description)
	}

	// بررسی مبلغ
	responseAmount, err := strconv.Atoi(response.Amount)
	if err != nil {
		return false, "", fmt.Errorf("خطا در پردازش مبلغ پاسخ: %v", err)
	}

	if responseAmount != amount {
		return false, "", fmt.Errorf("مبلغ تایید شده (%d) با مبلغ اصلی (%d) مطابقت ندارد", responseAmount, amount)
	}

	return true, response.RefNum, nil
}

// RefundPayment بازگشت وجه سامان
func (s *SamanService) RefundPayment(transactionID string, amount int) error {
	// بررسی پارامترهای ورودی
	if transactionID == "" {
		return errors.New("شناسه تراکنش الزامی است")
	}
	if amount < 1 {
		return errors.New("مبلغ باید عدد مثبت باشد")
	}

	// آماده‌سازی داده‌های درخواست
	terminalId := s.gateway.MerchantId

	// ساخت درخواست
	request := SamanRefundRequest{
		RefNum: transactionID,
		TermId: terminalId,
	}

	// ارسال درخواست
	apiURL := s.gateway.ApiEndpoints.Refund
	if s.gateway.IsTestMode {
		apiURL = "https://sep.shaparak.ir/api/v1/Payment/ReversePayment"
	} else {
		apiURL = "https://sep.shaparak.ir/api/v1/Payment/ReversePayment"
	}

	resp, err := s.sendAPIRequest(apiURL, request)
	if err != nil {
		return fmt.Errorf("خطا در ارسال درخواست بازگشت وجه به درگاه سامان: %v", err)
	}

	// پردازش پاسخ
	var response SamanRefundResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return fmt.Errorf("خطا در پردازش پاسخ بازگشت وجه درگاه سامان: %v", err)
	}

	// بررسی نتیجه
	if response.ResCode != 0 {
		return fmt.Errorf("خطای بازگشت وجه درگاه سامان: %s", response.Description)
	}

	return nil
}

// GetBalance دریافت موجودی سامان
func (s *SamanService) GetBalance() (float64, error) {
	// سامان API مستقیم برای دریافت موجودی ندارد
	// این متد برای سازگاری با interface پیاده‌سازی شده
	return 0, errors.New("درگاه سامان قابلیت دریافت موجودی ندارد")
}

// TestConnection تست اتصال به سامان
func (s *SamanService) TestConnection() error {
	// تست اتصال با ارسال درخواست تست
	testRequest := SamanPaymentRequest{
		TerminalId:  s.gateway.MerchantId,
		Amount:      "1000",
		ResNum:      "test_" + strconv.FormatInt(time.Now().Unix(), 10),
		RedirectUrl: "https://example.com/test",
		Action:      "Token",
	}

	apiURL := s.gateway.ApiEndpoints.Payment
	if s.gateway.IsTestMode {
		apiURL = "https://sep.shaparak.ir/api/v1/Payment/GetToken"
	}

	_, err := s.sendAPIRequest(apiURL, testRequest)
	if err != nil {
		return fmt.Errorf("خطا در تست اتصال به درگاه سامان: %v", err)
	}

	return nil
}

// ProcessCallback پردازش callback سامان
func (s *SamanService) ProcessCallback(params map[string]string) (*models.PaymentTransaction, error) {
	// بررسی پارامترهای callback
	state := params["State"]
	refNum := params["RefNum"]
	resNum := params["ResNum"]
	terminalId := params["TerminalId"]
	amount := params["Amount"]

	if state == "" || refNum == "" || resNum == "" || terminalId == "" || amount == "" {
		return nil, errors.New("پارامترهای callback ناقص است")
	}

	// بررسی وضعیت
	if state != "OK" {
		return nil, fmt.Errorf("پرداخت ناموفق بود. وضعیت: %s", state)
	}

	// بررسی ترمینال
	if terminalId != s.gateway.MerchantId {
		return nil, errors.New("شماره ترمینال نامعتبر است")
	}

	// ایجاد تراکنش
	transaction := &models.PaymentTransaction{
		TransactionID: refNum,
		OrderID:       resNum,
		Status:        "pending", // نیاز به تایید دارد
		PaymentMethod: "saman",
		GatewayType:   "saman",
		GatewayToken:  refNum,
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
func (s *SamanService) GetPaymentURL(transaction *models.PaymentTransaction, callbackURL string) string {
	// سامان از فرم HTML برای پرداخت استفاده می‌کند
	baseURL := s.gateway.ApiEndpoints.Payment
	if s.gateway.IsTestMode {
		baseURL = "https://sep.shaparak.ir/api/v1/Payment/GetToken"
	}
	return baseURL
}

// GetPaymentForm دریافت فرم HTML پرداخت
func (s *SamanService) GetPaymentForm(transaction *models.PaymentTransaction, callbackURL string) string {
	// فرم HTML برای ارسال به سامان
	form := fmt.Sprintf(`
		<form id="saman-payment-form" method="post" action="%s">
			<input type="hidden" name="Token" value="%s" />
			<input type="hidden" name="RedirectURL" value="%s" />
		</form>
		<script>document.getElementById('saman-payment-form').submit();</script>
	`,
		"https://sep.shaparak.ir/api/v1/Payment/GoToPayment",
		transaction.GatewayToken,
		callbackURL,
	)
	return form
}

// sendAPIRequest ارسال درخواست HTTP به API سامان
func (s *SamanService) sendAPIRequest(url string, data interface{}) ([]byte, error) {
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

// generateSamanTransactionID تولید شناسه یکتا برای تراکنش سامان
func generateSamanTransactionID() string {
	return fmt.Sprintf("SAM%d", time.Now().UnixNano())
}

// generateSamanOrderID تولید شناسه یکتا برای سفارش سامان
func generateSamanOrderID() string {
	return fmt.Sprintf("ORD%d", time.Now().UnixNano())
}

// createSignature ایجاد امضای دیجیتال (در صورت نیاز)
func (s *SamanService) createSignature(data string) string {
	// سامان معمولاً از امضای دیجیتال استفاده نمی‌کند
	// این متد برای سازگاری با سایر درگاه‌ها اضافه شده
	h := hmac.New(sha256.New, []byte(s.gateway.ApiKeys.PrivateKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
