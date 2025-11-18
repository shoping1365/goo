package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"my-go-backend/internal/models"
)

// ZarinpalService سرویس مدیریت درگاه پرداخت زرین‌پال
// شامل تمام عملیات پرداخت، تایید و مدیریت تراکنش‌ها

type ZarinpalService struct {
	gateway *models.PaymentGateway
}

// NewZarinpalService سازنده سرویس زرین‌پال
func NewZarinpalService(gateway *models.PaymentGateway) *ZarinpalService {
	return &ZarinpalService{
		gateway: gateway,
	}
}

// ساختارهای درخواست و پاسخ زرین‌پال

// درخواست ایجاد پرداخت
type zarinpalPaymentRequest struct {
	MerchantID  string `json:"merchant_id"`
	Amount      int    `json:"amount"`
	CallbackURL string `json:"callback_url"`
	Description string `json:"description"`
	Email       string `json:"email,omitempty"`
	Mobile      string `json:"mobile,omitempty"`
}

// پاسخ ایجاد پرداخت
type zarinpalPaymentResponse struct {
	Status    int    `json:"Status"`
	Authority string `json:"Authority"`
}

// درخواست تایید پرداخت
type zarinpalVerificationRequest struct {
	MerchantID string `json:"merchant_id"`
	Authority  string `json:"authority"`
	Amount     int    `json:"amount"`
}

// پاسخ تایید پرداخت
type zarinpalVerificationResponse struct {
	Status int         `json:"Status"`
	RefID  json.Number `json:"RefID"`
}

// درخواست تراکنش‌های تایید نشده
type zarinpalUnverifiedRequest struct {
	MerchantID string `json:"merchant_id"`
}

// تراکنش تایید نشده
type zarinpalUnverifiedAuthority struct {
	Authority   string `json:"Authority"`
	Amount      int    `json:"Amount"`
	Channel     string `json:"Channel"`
	CallbackURL string `json:"CallbackURL"`
	Referer     string `json:"Referer"`
	Email       string `json:"Email"`
	CellPhone   string `json:"CellPhone"`
	Date        string `json:"Date"`
}

// پاسخ تراکنش‌های تایید نشده
type zarinpalUnverifiedResponse struct {
	Status      int                           `json:"Status"`
	Authorities []zarinpalUnverifiedAuthority `json:"Authorities"`
}

// درخواست تمدید Authority
type zarinpalRefreshRequest struct {
	MerchantID string `json:"merchant_id"`
	Authority  string `json:"authority"`
	ExpireIn   int    `json:"expire_in"`
}

// پاسخ تمدید Authority
type zarinpalRefreshResponse struct {
	Status int `json:"Status"`
}

// CreatePayment ایجاد درخواست پرداخت جدید
// مبلغ به تومان ارسال می‌شود (نه ریال)
func (s *ZarinpalService) CreatePayment(amount int, callbackURL, description, email, mobile string) (*models.PaymentTransaction, error) {
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

	// ساخت درخواست
	request := zarinpalPaymentRequest{
		MerchantID:  s.gateway.MerchantId,
		Amount:      amount,
		CallbackURL: callbackURL,
		Description: description,
		Email:       email,
		Mobile:      mobile,
	}

	// ارسال درخواست به زرین‌پال
	response := &zarinpalPaymentResponse{}
	err := s.makeRequest("PaymentRequest.json", request, response)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست پرداخت: %v", err)
	}

	// بررسی پاسخ
	if response.Status != 100 {
		return nil, fmt.Errorf("خطای زرین‌پال: کد %d - %s", response.Status, s.getErrorMessage(response.Status))
	}

	// ایجاد تراکنش در دیتابیس
	transaction := &models.PaymentTransaction{
		GatewayID:     s.gateway.ID,
		OrderID:       fmt.Sprintf("ZP%d", time.Now().Unix()),
		TransactionID: response.Authority,
		Amount:        int64(amount),
		Status:        "pending",
		PaymentMethod: "zarinpal",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return transaction, nil
}

// VerifyPayment تایید پرداخت
func (s *ZarinpalService) VerifyPayment(amount int, authority string) (bool, string, error) {
	// بررسی پارامترهای ورودی
	if amount <= 0 {
		return false, "", errors.New("مبلغ باید عدد مثبت باشد")
	}
	if authority == "" {
		return false, "", errors.New("Authority الزامی است")
	}

	// ساخت درخواست
	request := zarinpalVerificationRequest{
		MerchantID: s.gateway.MerchantId,
		Amount:     amount,
		Authority:  authority,
	}

	// ارسال درخواست به زرین‌پال
	response := &zarinpalVerificationResponse{}
	err := s.makeRequest("PaymentVerification.json", request, response)
	if err != nil {
		return false, "", fmt.Errorf("خطا در تایید پرداخت: %v", err)
	}

	// بررسی پاسخ
	if response.Status != 100 {
		return false, "", fmt.Errorf("خطای زرین‌پال: کد %d - %s", response.Status, s.getErrorMessage(response.Status))
	}

	// تبدیل RefID به string
	refID := string(response.RefID)
	return true, refID, nil
}

// GetUnverifiedTransactions دریافت تراکنش‌های تایید نشده
func (s *ZarinpalService) GetUnverifiedTransactions() ([]zarinpalUnverifiedAuthority, error) {
	// ساخت درخواست
	request := zarinpalUnverifiedRequest{
		MerchantID: s.gateway.MerchantId,
	}

	// ارسال درخواست به زرین‌پال
	response := &zarinpalUnverifiedResponse{}
	err := s.makeRequest("UnverifiedTransactions.json", request, response)
	if err != nil {
		return nil, fmt.Errorf("خطا در دریافت تراکنش‌های تایید نشده: %v", err)
	}

	// بررسی پاسخ
	if response.Status != 100 {
		return nil, fmt.Errorf("خطای زرین‌پال: کد %d - %s", response.Status, s.getErrorMessage(response.Status))
	}

	return response.Authorities, nil
}

// RefreshAuthority تمدید زمان انقضای Authority
// expire باید بین 1800 تا 3888000 ثانیه باشد
func (s *ZarinpalService) RefreshAuthority(authority string, expire int) error {
	// بررسی پارامترهای ورودی
	if authority == "" {
		return errors.New("Authority الزامی است")
	}
	if expire < 1800 {
		return errors.New("زمان انقضا باید حداقل 1800 ثانیه باشد")
	}
	if expire > 3888000 {
		return errors.New("زمان انقضا نباید بیشتر از 3888000 ثانیه باشد")
	}

	// ساخت درخواست
	request := zarinpalRefreshRequest{
		MerchantID: s.gateway.MerchantId,
		Authority:  authority,
		ExpireIn:   expire,
	}

	// ارسال درخواست به زرین‌پال
	response := &zarinpalRefreshResponse{}
	err := s.makeRequest("RefreshAuthority.json", request, response)
	if err != nil {
		return fmt.Errorf("خطا در تمدید Authority: %v", err)
	}

	// بررسی پاسخ
	if response.Status != 100 {
		return fmt.Errorf("خطای زرین‌پال: کد %d - %s", response.Status, s.getErrorMessage(response.Status))
	}

	return nil
}

// GetPaymentURL دریافت آدرس پرداخت
func (s *ZarinpalService) GetPaymentURL(transaction *models.PaymentTransaction, callbackURL string) string {
	baseURL := s.gateway.ApiEndpoints.Payment
	if s.gateway.IsTestMode {
		baseURL = "https://sandbox.zarinpal.com/pg/StartPay/"
	} else {
		baseURL = "https://www.zarinpal.com/pg/StartPay/"
	}
	return baseURL + transaction.TransactionID
}

// TestConnection تست اتصال به زرین‌پال
func (s *ZarinpalService) TestConnection() error {
	// تست با ایجاد یک درخواست پرداخت تستی
	testAmount := 1000 // 1000 تومان
	testCallback := "https://test.example.com/callback"
	testDescription := "تست اتصال زرین‌پال"

	request := zarinpalPaymentRequest{
		MerchantID:  s.gateway.MerchantId,
		Amount:      testAmount,
		CallbackURL: testCallback,
		Description: testDescription,
	}

	response := &zarinpalPaymentResponse{}
	err := s.makeRequest("PaymentRequest.json", request, response)
	if err != nil {
		return fmt.Errorf("خطا در اتصال: %v", err)
	}

	if response.Status == 100 {
		return nil
	} else if response.Status == -3 {
		return nil // خطای مبلغ قابل قبول است
	} else {
		return fmt.Errorf("خطای زرین‌پال: کد %d", response.Status)
	}
}

// makeRequest ارسال درخواست HTTP به زرین‌پال
func (s *ZarinpalService) makeRequest(method string, data interface{}, response interface{}) error {
	// تعیین آدرس API بر اساس حالت تست
	apiEndpoint := s.gateway.ApiEndpoints.Payment
	if s.gateway.IsTestMode {
		apiEndpoint = "https://sandbox.zarinpal.com/pg/rest/WebGate/"
	} else {
		apiEndpoint = "https://www.zarinpal.com/pg/rest/WebGate/"
	}

	// تبدیل داده به JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("خطا در تبدیل داده به JSON: %v", err)
	}

	// ایجاد درخواست HTTP
	req, err := http.NewRequest("POST", apiEndpoint+method, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("خطا در ایجاد درخواست HTTP: %v", err)
	}

	// تنظیم هدرها
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// ارسال درخواست
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("خطا در ارسال درخواست: %v", err)
	}
	defer resp.Body.Close()

	// خواندن پاسخ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("خطا در خواندن پاسخ: %v", err)
	}

	// بررسی کد وضعیت HTTP
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("خطای HTTP: %d - %s", resp.StatusCode, string(body))
	}

	// تبدیل پاسخ JSON
	err = json.Unmarshal(body, response)
	if err != nil {
		return fmt.Errorf("خطا در تبدیل پاسخ JSON: %v", err)
	}

	return nil
}

// getErrorMessage تبدیل کد خطای زرین‌پال به پیام فارسی
func (s *ZarinpalService) getErrorMessage(status int) string {
	switch status {
	case -1:
		return "اطلاعات ارسال شده ناقص است"
	case -2:
		return "IP یا مرچنت کد پذیرنده صحیح نیست"
	case -3:
		return "با توجه به محدودیت‌های شاپینگ امتیاز شما کافی نیست"
	case -4:
		return "سطح تایید پذیرنده پایین تر از سطح نقره‌ای است"
	case -11:
		return "درخواست مورد نظر یافت نشد"
	case -12:
		return "امکان ویرایش درخواست وجود ندارد"
	case -21:
		return "هیچ نوع عملیات مالی برای این تراکنش یافت نشد"
	case -22:
		return "تراکنش ناموفق می‌باشد"
	case -33:
		return "مبلغ تراکنش با مبلغ پرداخت شده مطابقت ندارد"
	case -34:
		return "سقف تقسیم تراکنش از لحاظ تعداد یا رقم عبور نموده"
	case -40:
		return "اجازه دسترسی به متد مورد نظر وجود ندارد"
	case -41:
		return "اطلاعات ارسال شده مربوط به AdditionalData غیرمعتبر می‌باشد"
	case -42:
		return "مدت زمان معتبر طول عمر شناسه پرداخت باید بین ۳۰ دقیقه تا ۴۵ روز می‌باشد"
	case -54:
		return "درخواست مورد نظر آرشیو شده"
	case 100:
		return "عملیات با موفقیت انجام شد"
	case 101:
		return "عملیات پرداخت قبلاً انجام شده"
	default:
		return fmt.Sprintf("خطای نامشخص با کد %d", status)
	}
}

// GetBalance دریافت موجودی حساب (در صورت پشتیبانی)
func (s *ZarinpalService) GetBalance() (float64, error) {
	// زرین‌پال API مستقیم برای دریافت موجودی ندارد
	// این متد برای سازگاری با سایر درگاه‌ها اضافه شده
	return 0, errors.New("زرین‌پال API دریافت موجودی ندارد")
}

// RefundPayment بازگشت وجه (در صورت پشتیبانی)
func (s *ZarinpalService) RefundPayment(refID string, amount int) error {
	// زرین‌پال API مستقیم برای بازگشت وجه ندارد
	// این متد برای سازگاری با سایر درگاه‌ها اضافه شده
	return errors.New("زرین‌پال API بازگشت وجه ندارد")
}
