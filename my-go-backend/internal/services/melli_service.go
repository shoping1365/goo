package services

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"my-go-backend/internal/models"
)

// MelliService سرویس درگاه پرداخت بانک ملی ایران (BMI)
// شامل تمام عملیات پرداخت، تایید و بازگشت وجه

type MelliService struct {
	gateway *models.PaymentGateway
}

// NewMelliService سازنده سرویس ملی
func NewMelliService(gateway *models.PaymentGateway) *MelliService {
	return &MelliService{
		gateway: gateway,
	}
}

// ساختارهای JSON برای API requests
type MelliPaymentRequest struct {
	TerminalId    string `json:"TerminalId"`
	MerchantId    string `json:"MerchantId"`
	Amount        string `json:"Amount"`
	SignData      string `json:"SignData"`
	ReturnUrl     string `json:"ReturnUrl"`
	LocalDateTime string `json:"LocalDateTime"`
	OrderId       string `json:"OrderId"`
}

type MelliPaymentResponse struct {
	ResCode     int    `json:"ResCode"`
	Token       string `json:"Token"`
	Description string `json:"Description"`
}

type MelliVerificationRequest struct {
	Token    string `json:"Token"`
	SignData string `json:"SignData"`
}

type MelliVerificationResponse struct {
	ResCode       int    `json:"ResCode"`
	RetrivalRefNo string `json:"RetrivalRefNo"`
	SystemTraceNo string `json:"SystemTraceNo"`
	OrderId       string `json:"OrderId"`
	Description   string `json:"Description"`
}

// CreatePayment ایجاد درخواست پرداخت ملی
func (s *MelliService) CreatePayment(amount int, callbackURL, description, email, mobile string) (*models.PaymentTransaction, error) {
	// ایجاد تراکنش
	transaction := &models.PaymentTransaction{
		TransactionID: generateTransactionID(),
		OrderID:       generateOrderID(),
		Amount:        int64(amount),
		Status:        "pending",
		PaymentMethod: "melli",
		Description:   description,
		GatewayType:   "melli",
		CreatedAt:     time.Now(),
	}

	// آماده‌سازی داده‌های درخواست
	terminalId := s.gateway.MerchantId
	merchantId := s.gateway.ApiKeys.PublicKey
	orderId := transaction.OrderID
	amountStr := strconv.Itoa(amount)
	localDateTime := time.Now().Format("01/02/2006 3:04:05 PM")

	// ایجاد SignData
	signData, err := s.encryptPKCS7(fmt.Sprintf("%s;%s;%s", terminalId, orderId, amountStr), s.gateway.ApiKeys.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("خطا در رمزنگاری SignData: %v", err)
	}

	// آماده‌سازی درخواست
	request := MelliPaymentRequest{
		TerminalId:    terminalId,
		MerchantId:    merchantId,
		Amount:        amountStr,
		SignData:      signData,
		ReturnUrl:     callbackURL,
		LocalDateTime: localDateTime,
		OrderId:       orderId,
	}

	// ارسال درخواست
	apiURL := s.gateway.ApiEndpoints.Payment
	if s.gateway.IsTestMode {
		apiURL = "https://test.sadad.shaparak.ir/vpg/api/v0/Request/PaymentRequest"
	} else {
		apiURL = "https://sadad.shaparak.ir/vpg/api/v0/Request/PaymentRequest"
	}

	resp, err := s.sendAPIRequest(apiURL, request)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست به درگاه ملی: %v", err)
	}

	// پردازش پاسخ
	var response MelliPaymentResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, fmt.Errorf("خطا در پردازش پاسخ درگاه ملی: %v", err)
	}

	// بررسی نتیجه
	if response.ResCode != 0 {
		errorMsg := s.getErrorMessage(response.ResCode)
		return nil, fmt.Errorf("خطا از درگاه ملی: %s", errorMsg)
	}

	// ذخیره توکن
	transaction.GatewayToken = response.Token
	transaction.Status = "pending"

	return transaction, nil
}

// VerifyPayment تایید پرداخت ملی
func (s *MelliService) VerifyPayment(amount int, token string) (bool, string, error) {
	// ایجاد SignData برای تایید
	signData, err := s.encryptPKCS7(token, s.gateway.ApiKeys.PrivateKey)
	if err != nil {
		return false, "", fmt.Errorf("خطا در رمزنگاری SignData: %v", err)
	}

	// آماده‌سازی درخواست تایید
	request := MelliVerificationRequest{
		Token:    token,
		SignData: signData,
	}

	// ارسال درخواست تایید
	apiURL := s.gateway.ApiEndpoints.Verification
	if s.gateway.IsTestMode {
		apiURL = "https://test.sadad.shaparak.ir/vpg/api/v0/Advice/Verify"
	} else {
		apiURL = "https://sadad.shaparak.ir/vpg/api/v0/Advice/Verify"
	}

	resp, err := s.sendAPIRequest(apiURL, request)
	if err != nil {
		return false, "", fmt.Errorf("خطا در ارسال درخواست تایید به درگاه ملی: %v", err)
	}

	// پردازش پاسخ
	var response MelliVerificationResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return false, "", fmt.Errorf("خطا در پردازش پاسخ تایید: %v", err)
	}

	// بررسی نتیجه
	if response.ResCode != 0 {
		errorMsg := s.getErrorMessage(response.ResCode)
		return false, "", fmt.Errorf("خطا در تایید پرداخت: %s", errorMsg)
	}

	return true, response.RetrivalRefNo, nil
}

// RefundPayment بازگشت وجه ملی
func (s *MelliService) RefundPayment(token string, amount int) error {
	// آماده‌سازی درخواست بازگشت وجه
	amountStr := strconv.Itoa(amount)
	signData, err := s.encryptPKCS7(fmt.Sprintf("%s;%s", token, amountStr), s.gateway.ApiKeys.PrivateKey)
	if err != nil {
		return fmt.Errorf("خطا در رمزنگاری SignData: %v", err)
	}

	request := map[string]string{
		"Token":    token,
		"Amount":   amountStr,
		"SignData": signData,
	}

	// ارسال درخواست بازگشت وجه
	apiURL := s.gateway.ApiEndpoints.Refund
	if s.gateway.IsTestMode {
		apiURL = "https://test.sadad.shaparak.ir/vpg/api/v0/Advice/Refund"
	} else {
		apiURL = "https://sadad.shaparak.ir/vpg/api/v0/Advice/Refund"
	}

	resp, err := s.sendAPIRequest(apiURL, request)
	if err != nil {
		return fmt.Errorf("خطا در ارسال درخواست بازگشت وجه به درگاه ملی: %v", err)
	}

	// پردازش پاسخ
	var response map[string]interface{}
	if err := json.Unmarshal(resp, &response); err != nil {
		return fmt.Errorf("خطا در پردازش پاسخ بازگشت وجه: %v", err)
	}

	// بررسی نتیجه
	if resCode, exists := response["ResCode"]; exists {
		if resCode.(float64) != 0 {
			errorMsg := s.getErrorMessage(int(resCode.(float64)))
			return fmt.Errorf("خطا در بازگشت وجه: %s", errorMsg)
		}
	}

	return nil
}

// InquiryPayment استعلام تراکنش ملی
func (s *MelliService) InquiryPayment(token string) (int, error) {
	// آماده‌سازی درخواست استعلام
	signData, err := s.encryptPKCS7(token, s.gateway.ApiKeys.PrivateKey)
	if err != nil {
		return 0, fmt.Errorf("خطا در رمزنگاری SignData: %v", err)
	}

	request := map[string]string{
		"Token":    token,
		"SignData": signData,
	}

	// ارسال درخواست استعلام
	apiURL := s.gateway.ApiEndpoints.Balance
	if s.gateway.IsTestMode {
		apiURL = "https://test.sadad.shaparak.ir/vpg/api/v0/Advice/Inquiry"
	} else {
		apiURL = "https://sadad.shaparak.ir/vpg/api/v0/Advice/Inquiry"
	}

	resp, err := s.sendAPIRequest(apiURL, request)
	if err != nil {
		return 0, fmt.Errorf("خطا در ارسال درخواست استعلام به درگاه ملی: %v", err)
	}

	// پردازش پاسخ
	var response map[string]interface{}
	if err := json.Unmarshal(resp, &response); err != nil {
		return 0, fmt.Errorf("خطا در پردازش پاسخ استعلام: %v", err)
	}

	// بررسی نتیجه
	if resCode, exists := response["ResCode"]; exists {
		if resCode.(float64) != 0 {
			errorMsg := s.getErrorMessage(int(resCode.(float64)))
			return 0, fmt.Errorf("خطا در استعلام تراکنش: %s", errorMsg)
		}
	}

	// دریافت مبلغ
	if amount, exists := response["Amount"]; exists {
		amountStr := amount.(string)
		amountInt, err := strconv.Atoi(amountStr)
		if err != nil {
			return 0, fmt.Errorf("خطا در تبدیل مبلغ: %v", err)
		}
		return amountInt, nil
	}

	return 0, fmt.Errorf("مبلغ در پاسخ یافت نشد")
}

// GetBalance دریافت موجودی حساب ملی
func (s *MelliService) GetBalance() (float64, error) {
	// درگاه ملی معمولاً قابلیت دریافت موجودی ندارد
	// این متد برای سازگاری با interface اضافه شده
	return 0, fmt.Errorf("درگاه ملی قابلیت دریافت موجودی ندارد")
}

// TestConnection تست اتصال به درگاه ملی
func (s *MelliService) TestConnection() error {
	// تست با ایجاد یک تراکنش تستی
	testTransaction := &models.PaymentTransaction{
		TransactionID: "TEST_" + generateTransactionID(),
		OrderID:       "TEST_" + generateOrderID(),
		Amount:        1000,
		Description:   "تست اتصال درگاه ملی",
		Status:        "pending",
		PaymentMethod: "melli",
		GatewayType:   "melli",
		CreatedAt:     time.Now(),
	}

	// آماده‌سازی داده‌های تست
	terminalId := s.gateway.MerchantId
	merchantId := s.gateway.ApiKeys.PublicKey
	orderId := testTransaction.OrderID
	amountStr := "1000"
	localDateTime := time.Now().Format("01/02/2006 3:04:05 PM")

	// ایجاد SignData
	signData, err := s.encryptPKCS7(fmt.Sprintf("%s;%s;%s", terminalId, orderId, amountStr), s.gateway.ApiKeys.PrivateKey)
	if err != nil {
		return fmt.Errorf("خطا در آماده‌سازی درخواست تست")
	}

	// آماده‌سازی درخواست تست
	request := MelliPaymentRequest{
		TerminalId:    terminalId,
		MerchantId:    merchantId,
		Amount:        amountStr,
		SignData:      signData,
		ReturnUrl:     "https://test.com/callback",
		LocalDateTime: localDateTime,
		OrderId:       orderId,
	}

	// ارسال درخواست تست
	apiURL := "https://test.sadad.shaparak.ir/vpg/api/v0/Request/PaymentRequest"
	resp, err := s.sendAPIRequest(apiURL, request)
	if err != nil {
		return fmt.Errorf("خطا در اتصال به درگاه ملی: %v", err)
	}

	// پردازش پاسخ
	var response MelliPaymentResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return fmt.Errorf("خطا در پردازش پاسخ تست")
	}

	// بررسی نتیجه
	if response.ResCode != 0 {
		errorMsg := s.getErrorMessage(response.ResCode)
		return fmt.Errorf("خطا در تست اتصال: %s", errorMsg)
	}

	return nil
}

// ProcessCallback پردازش callback ملی
func (s *MelliService) ProcessCallback(formData map[string]string) (bool, string, string, error) {
	// دریافت پارامترهای callback
	resCode := formData["ResCode"]
	orderId := formData["OrderId"]
	token := formData["token"]

	// بررسی کد نتیجه
	if resCode != "0" {
		errorMsg := s.getErrorMessage(0) // کد خطا را به عدد تبدیل می‌کنیم
		return false, "", "", fmt.Errorf("خطا در callback: %s", errorMsg)
	}

	// تایید پرداخت
	success, refNum, err := s.VerifyPayment(0, token)
	if err != nil {
		return false, "", "", err
	}

	return success, refNum, orderId, nil
}

// GetPaymentURL دریافت آدرس پرداخت ملی
func (s *MelliService) GetPaymentURL(transaction *models.PaymentTransaction, callbackURL string) string {
	baseURL := "https://sadad.shaparak.ir/VPG/Purchase"
	if s.gateway.IsTestMode {
		baseURL = "https://test.sadad.shaparak.ir/VPG/Purchase"
	}

	return fmt.Sprintf("%s?Token=%s", baseURL, transaction.GatewayToken)
}

// GetPaymentForm دریافت فرم HTML پرداخت ملی
func (s *MelliService) GetPaymentForm(transaction *models.PaymentTransaction, callbackURL string) string {
	paymentURL := s.GetPaymentURL(transaction, callbackURL)

	form := fmt.Sprintf(`
		<form id="melli-payment-form" method="POST" action="%s">
			<input type="hidden" name="Token" value="%s">
		</form>
		<script>
			document.getElementById('melli-payment-form').submit();
		</script>
	`, paymentURL, transaction.GatewayToken)

	return form
}

// sendAPIRequest ارسال درخواست API
func (s *MelliService) sendAPIRequest(url string, data interface{}) ([]byte, error) {
	// تبدیل داده به JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// ایجاد درخواست HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// تنظیم headers
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "MelliPaymentGateway/1.0")

	// ارسال درخواست
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// خواندن پاسخ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// بررسی کد وضعیت
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("خطا در پاسخ سرور: %d - %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// encryptPKCS7 رمزنگاری PKCS7 (شبیه‌سازی DES-EDE3)
func (s *MelliService) encryptPKCS7(data, key string) (string, error) {
	// رمزگشایی کلید از base64
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", fmt.Errorf("خطا در رمزگشایی کلید: %v", err)
	}

	// ایجاد block cipher
	block, err := des.NewTripleDESCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("خطا در ایجاد cipher: %v", err)
	}

	// padding داده
	paddedData := s.pkcs7Pad([]byte(data), block.BlockSize())

	// رمزنگاری
	ciphertext := make([]byte, len(paddedData))
	block.Encrypt(ciphertext, paddedData)

	// تبدیل به base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// pkcs7Pad اضافه کردن padding PKCS7
func (s *MelliService) pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// getErrorMessage تبدیل کد خطا به پیام فارسی
func (s *MelliService) getErrorMessage(resCode int) string {
	errorMessages := map[int]string{
		-1:  "خطا در اتصال به درگاه ملی",
		-2:  "خطا در احراز هویت",
		-3:  "خطا در پارامترهای ورودی",
		-4:  "خطا در مبلغ تراکنش",
		-5:  "خطا در شماره سفارش",
		-6:  "خطا در آدرس callback",
		-7:  "خطا در شماره ترمینال",
		-8:  "خطا در نام کاربری",
		-9:  "خطا در رمز عبور",
		-10: "خطا در توکن",
		-11: "خطا در تایید تراکنش",
		-12: "خطا در بازگشت وجه",
		-13: "خطا در استعلام تراکنش",
		-14: "تراکنش قبلاً تایید شده",
		-15: "تراکنش ناموفق",
		-16: "تراکنش لغو شده",
		-17: "تراکنش منقضی شده",
		-18: "خطا در سیستم",
		-19: "خطا در شبکه",
		-20: "خطا در پایگاه داده",
	}

	if msg, exists := errorMessages[resCode]; exists {
		return msg
	}

	return fmt.Sprintf("خطای نامشخص با کد: %d", resCode)
}

// generateTransactionID تولید شناسه تراکنش
func generateTransactionID() string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("MEL%013d", timestamp%10000000000000)
}

// generateOrderID تولید شناسه سفارش
func generateOrderID() string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("MEL%013d", timestamp%10000000000000)
}
