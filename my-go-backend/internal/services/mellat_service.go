package services

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"my-go-backend/internal/models"
)

// MellatService سرویس مدیریت درگاه پرداخت ملت
// شامل تمام عملیات پرداخت، تایید و بازگشت وجه

type MellatService struct {
	gateway *models.PaymentGateway
}

// NewMellatService سازنده سرویس ملت
func NewMellatService(gateway *models.PaymentGateway) *MellatService {
	return &MellatService{
		gateway: gateway,
	}
}

// ساختارهای SOAP برای ملت

// درخواست پرداخت
type mellatPayRequest struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XMLNS   string   `xml:"xmlns:soap,attr"`
	Body    struct {
		BpPayRequest struct {
			XMLName        xml.Name `xml:"bpPayRequest"`
			XMLNS          string   `xml:"xmlns,attr"`
			TerminalId     int64    `xml:"terminalId"`
			UserName       string   `xml:"userName"`
			UserPassword   string   `xml:"userPassword"`
			OrderId        int64    `xml:"orderId"`
			Amount         int64    `xml:"amount"`
			LocalDate      string   `xml:"localDate"`
			LocalTime      string   `xml:"localTime"`
			AdditionalData string   `xml:"additionalData"`
			CallBackUrl    string   `xml:"callBackUrl"`
			PayerId        int64    `xml:"payerId"`
		} `xml:"bpPayRequest"`
	} `xml:"soap:Body"`
}

// پاسخ پرداخت
type mellatPayResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		BpPayRequestResponse struct {
			Return string `xml:"return"`
		} `xml:"bpPayRequestResponse"`
	} `xml:"Body"`
}

// درخواست تایید
type mellatVerifyRequest struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XMLNS   string   `xml:"xmlns:soap,attr"`
	Body    struct {
		BpVerifyRequest struct {
			XMLName         xml.Name `xml:"bpVerifyRequest"`
			XMLNS           string   `xml:"xmlns,attr"`
			TerminalId      int64    `xml:"terminalId"`
			UserName        string   `xml:"userName"`
			UserPassword    string   `xml:"userPassword"`
			OrderId         int64    `xml:"orderId"`
			SaleOrderId     int64    `xml:"saleOrderId"`
			SaleReferenceId int64    `xml:"saleReferenceId"`
		} `xml:"bpVerifyRequest"`
	} `xml:"soap:Body"`
}

// پاسخ تایید
type mellatVerifyResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		BpVerifyRequestResponse struct {
			Return string `xml:"return"`
		} `xml:"bpVerifyRequestResponse"`
	} `xml:"Body"`
}

// درخواست بازگشت وجه
type mellatRefundRequest struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XMLNS   string   `xml:"xmlns:soap,attr"`
	Body    struct {
		BpRefundRequest struct {
			XMLName         xml.Name `xml:"bpRefundRequest"`
			XMLNS           string   `xml:"xmlns,attr"`
			TerminalId      int64    `xml:"terminalId"`
			UserName        string   `xml:"userName"`
			UserPassword    string   `xml:"userPassword"`
			OrderId         int64    `xml:"orderId"`
			SaleOrderId     int64    `xml:"saleOrderId"`
			SaleReferenceId int64    `xml:"saleReferenceId"`
			RefundAmount    int64    `xml:"refundAmount"`
		} `xml:"bpRefundRequest"`
	} `xml:"soap:Body"`
}

// پاسخ بازگشت وجه
type mellatRefundResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		BpRefundRequestResponse struct {
			Return string `xml:"return"`
		} `xml:"bpRefundRequestResponse"`
	} `xml:"Body"`
}

// درخواست استعلام
type mellatInquiryRequest struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XMLNS   string   `xml:"xmlns:soap,attr"`
	Body    struct {
		BpInquiryRequest struct {
			XMLName         xml.Name `xml:"bpInquiryRequest"`
			XMLNS           string   `xml:"xmlns,attr"`
			TerminalId      int64    `xml:"terminalId"`
			UserName        string   `xml:"userName"`
			UserPassword    string   `xml:"userPassword"`
			OrderId         int64    `xml:"orderId"`
			SaleOrderId     int64    `xml:"saleOrderId"`
			SaleReferenceId int64    `xml:"saleReferenceId"`
		} `xml:"bpInquiryRequest"`
	} `xml:"soap:Body"`
}

// پاسخ استعلام
type mellatInquiryResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		BpInquiryRequestResponse struct {
			Return string `xml:"return"`
		} `xml:"bpInquiryRequestResponse"`
	} `xml:"Body"`
}

// CreatePayment ایجاد درخواست پرداخت جدید
func (m *MellatService) CreatePayment(amount int, callbackURL, description, email, mobile string) (*models.PaymentTransaction, error) {
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

	// تبدیل مرچنت ID به عدد
	terminalId, err := strconv.ParseInt(m.gateway.MerchantId, 10, 64)
	if err != nil {
		return nil, errors.New("مرچنت ID نامعتبر است")
	}

	// تولید OrderId یکتا
	orderId := time.Now().UnixNano() % 1000000000 // حداکثر 9 رقم

	// دریافت تاریخ و زمان فعلی
	now := time.Now()
	localDate := now.Format("20060102") // YYYYMMDD
	localTime := now.Format("150405")   // HHMMSS

	// ساخت درخواست SOAP
	request := mellatPayRequest{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
	}
	request.Body.BpPayRequest.XMLNS = "http://interfaces.core.sw.bps.com/"
	request.Body.BpPayRequest.TerminalId = terminalId
	request.Body.BpPayRequest.UserName = m.gateway.ApiKeys.PublicKey
	request.Body.BpPayRequest.UserPassword = m.gateway.ApiKeys.PrivateKey
	request.Body.BpPayRequest.OrderId = orderId
	request.Body.BpPayRequest.Amount = int64(amount)
	request.Body.BpPayRequest.LocalDate = localDate
	request.Body.BpPayRequest.LocalTime = localTime
	request.Body.BpPayRequest.AdditionalData = description
	request.Body.BpPayRequest.CallBackUrl = callbackURL
	request.Body.BpPayRequest.PayerId = 0

	// ارسال درخواست به ملت
	response := &mellatPayResponse{}
	err = m.makeSOAPRequest("bpPayRequest", request, response)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست پرداخت: %v", err)
	}

	// پردازش پاسخ
	result := response.Body.BpPayRequestResponse.Return
	if result == "" {
		return nil, errors.New("پاسخ خالی از سرور ملت دریافت شد")
	}

	// بررسی کد پاسخ
	parts := strings.Split(result, ",")
	if len(parts) < 2 {
		return nil, fmt.Errorf("فرمت پاسخ نامعتبر: %s", result)
	}

	resCode := parts[0]
	refId := parts[1]

	if resCode != "0" {
		return nil, fmt.Errorf("خطای ملت: کد %s - %s", resCode, m.getErrorMessage(resCode))
	}

	// ایجاد تراکنش در دیتابیس
	transaction := &models.PaymentTransaction{
		GatewayID:     m.gateway.ID,
		OrderID:       fmt.Sprintf("ML%d", orderId),
		TransactionID: refId,
		Amount:        int64(amount),
		Status:        "pending",
		PaymentMethod: "mellat",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return transaction, nil
}

// VerifyPayment تایید پرداخت
func (m *MellatService) VerifyPayment(amount int, authority string) (bool, string, error) {
	// بررسی پارامترهای ورودی
	if amount <= 0 {
		return false, "", errors.New("مبلغ باید عدد مثبت باشد")
	}
	if authority == "" {
		return false, "", errors.New("Authority الزامی است")
	}

	// تبدیل مرچنت ID به عدد
	terminalId, err := strconv.ParseInt(m.gateway.MerchantId, 10, 64)
	if err != nil {
		return false, "", errors.New("مرچنت ID نامعتبر است")
	}

	// استعلام تراکنش برای دریافت SaleOrderId و SaleReferenceId
	saleOrderId, saleReferenceId, err := m.inquiryTransaction(authority)
	if err != nil {
		return false, "", fmt.Errorf("خطا در استعلام تراکنش: %v", err)
	}

	// ساخت درخواست تایید SOAP
	request := mellatVerifyRequest{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
	}
	request.Body.BpVerifyRequest.XMLNS = "http://interfaces.core.sw.bps.com/"
	request.Body.BpVerifyRequest.TerminalId = terminalId
	request.Body.BpVerifyRequest.UserName = m.gateway.ApiKeys.PublicKey
	request.Body.BpVerifyRequest.UserPassword = m.gateway.ApiKeys.PrivateKey
	request.Body.BpVerifyRequest.OrderId = saleOrderId
	request.Body.BpVerifyRequest.SaleOrderId = saleOrderId
	request.Body.BpVerifyRequest.SaleReferenceId = saleReferenceId

	// ارسال درخواست به ملت
	response := &mellatVerifyResponse{}
	err = m.makeSOAPRequest("bpVerifyRequest", request, response)
	if err != nil {
		return false, "", fmt.Errorf("خطا در تایید پرداخت: %v", err)
	}

	// بررسی پاسخ
	result := response.Body.BpVerifyRequestResponse.Return
	if result == "0" {
		return true, authority, nil
	} else {
		return false, "", fmt.Errorf("خطای ملت در تایید: کد %s - %s", result, m.getErrorMessage(result))
	}
}

// RefundPayment بازگشت وجه
func (m *MellatService) RefundPayment(refID string, amount int) error {
	// بررسی پارامترهای ورودی
	if refID == "" {
		return errors.New("شماره ارجاع الزامی است")
	}
	if amount <= 0 {
		return errors.New("مبلغ باید عدد مثبت باشد")
	}

	// تبدیل مرچنت ID به عدد
	terminalId, err := strconv.ParseInt(m.gateway.MerchantId, 10, 64)
	if err != nil {
		return errors.New("مرچنت ID نامعتبر است")
	}

	// استعلام تراکنش برای دریافت SaleOrderId و SaleReferenceId
	saleOrderId, saleReferenceId, err := m.inquiryTransaction(refID)
	if err != nil {
		return fmt.Errorf("خطا در استعلام تراکنش: %v", err)
	}

	// ساخت درخواست بازگشت وجه SOAP
	request := mellatRefundRequest{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
	}
	request.Body.BpRefundRequest.XMLNS = "http://interfaces.core.sw.bps.com/"
	request.Body.BpRefundRequest.TerminalId = terminalId
	request.Body.BpRefundRequest.UserName = m.gateway.ApiKeys.PublicKey
	request.Body.BpRefundRequest.UserPassword = m.gateway.ApiKeys.PrivateKey
	request.Body.BpRefundRequest.OrderId = saleOrderId
	request.Body.BpRefundRequest.SaleOrderId = saleOrderId
	request.Body.BpRefundRequest.SaleReferenceId = saleReferenceId
	request.Body.BpRefundRequest.RefundAmount = int64(amount)

	// ارسال درخواست به ملت
	response := &mellatRefundResponse{}
	err = m.makeSOAPRequest("bpRefundRequest", request, response)
	if err != nil {
		return fmt.Errorf("خطا در بازگشت وجه: %v", err)
	}

	// بررسی پاسخ
	result := response.Body.BpRefundRequestResponse.Return
	if result == "0" {
		return nil // موفقیت‌آمیز
	} else {
		return fmt.Errorf("خطای ملت در بازگشت وجه: کد %s - %s", result, m.getErrorMessage(result))
	}
}

// inquiryTransaction استعلام تراکنش
func (m *MellatService) inquiryTransaction(refId string) (int64, int64, error) {
	// تبدیل مرچنت ID به عدد
	terminalId, err := strconv.ParseInt(m.gateway.MerchantId, 10, 64)
	if err != nil {
		return 0, 0, errors.New("مرچنت ID نامعتبر است")
	}

	// تبدیل RefId به OrderId
	orderId, err := strconv.ParseInt(refId, 10, 64)
	if err != nil {
		return 0, 0, errors.New("شماره ارجاع نامعتبر است")
	}

	// ساخت درخواست استعلام SOAP
	request := mellatInquiryRequest{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
	}
	request.Body.BpInquiryRequest.XMLNS = "http://interfaces.core.sw.bps.com/"
	request.Body.BpInquiryRequest.TerminalId = terminalId
	request.Body.BpInquiryRequest.UserName = m.gateway.ApiKeys.PublicKey
	request.Body.BpInquiryRequest.UserPassword = m.gateway.ApiKeys.PrivateKey
	request.Body.BpInquiryRequest.OrderId = orderId
	request.Body.BpInquiryRequest.SaleOrderId = orderId
	request.Body.BpInquiryRequest.SaleReferenceId = orderId

	// ارسال درخواست به ملت
	response := &mellatInquiryResponse{}
	err = m.makeSOAPRequest("bpInquiryRequest", request, response)
	if err != nil {
		return 0, 0, fmt.Errorf("خطا در استعلام تراکنش: %v", err)
	}

	// پردازش پاسخ
	result := response.Body.BpInquiryRequestResponse.Return
	if result == "" {
		return 0, 0, errors.New("پاسخ خالی از سرور ملت دریافت شد")
	}

	// بررسی کد پاسخ
	parts := strings.Split(result, ",")
	if len(parts) < 3 {
		return 0, 0, fmt.Errorf("فرمت پاسخ نامعتبر: %s", result)
	}

	resCode := parts[0]
	saleOrderId, _ := strconv.ParseInt(parts[1], 10, 64)
	saleReferenceId, _ := strconv.ParseInt(parts[2], 10, 64)

	if resCode != "0" {
		return 0, 0, fmt.Errorf("خطای ملت در استعلام: کد %s - %s", resCode, m.getErrorMessage(resCode))
	}

	return saleOrderId, saleReferenceId, nil
}

// GetPaymentURL دریافت آدرس پرداخت
func (m *MellatService) GetPaymentURL(transaction *models.PaymentTransaction, callbackURL string) string {
	// ملت از فرم HTML برای پرداخت استفاده می‌کند
	// این آدرس باید به فرم پرداخت ملت هدایت شود
	var baseURL string
	if m.gateway.IsTestMode {
		baseURL = "https://pgwstest.bpm.bankmellat.ir/pgwchannel/startpay.mellat"
	} else {
		baseURL = "https://bpm.shaparak.ir/pgwchannel/startpay.mellat"
	}
	return baseURL
}

// GetPaymentForm دریافت فرم HTML پرداخت
func (m *MellatService) GetPaymentForm(transaction *models.PaymentTransaction, callbackURL string) string {
	// فرم HTML برای ارسال به ملت
	form := fmt.Sprintf(`
		<form id="mellat-payment-form" method="post" action="%s">
			<input type="hidden" name="RefId" value="%s" />
		</form>
		<script>document.getElementById('mellat-payment-form').submit();</script>
	`,
		m.gateway.ApiEndpoints.Payment,
		transaction.TransactionID,
	)
	return form
}

// TestConnection تست اتصال به ملت
func (m *MellatService) TestConnection() error {
	// تست با ایجاد یک درخواست پرداخت تستی
	testAmount := 1000 // 1000 تومان
	testCallback := "https://test.example.com/callback"
	testDescription := "تست اتصال ملت"

	// تبدیل مرچنت ID به عدد
	terminalId, err := strconv.ParseInt(m.gateway.MerchantId, 10, 64)
	if err != nil {
		return fmt.Errorf("مرچنت ID نامعتبر است")
	}

	// تولید OrderId تستی
	orderId := time.Now().UnixNano() % 1000000000

	// دریافت تاریخ و زمان فعلی
	now := time.Now()
	localDate := now.Format("20060102")
	localTime := now.Format("150405")

	// ساخت درخواست تست SOAP
	request := mellatPayRequest{
		XMLNS: "http://schemas.xmlsoap.org/soap/envelope/",
	}
	request.Body.BpPayRequest.XMLNS = "http://interfaces.core.sw.bps.com/"
	request.Body.BpPayRequest.TerminalId = terminalId
	request.Body.BpPayRequest.UserName = m.gateway.ApiKeys.PublicKey
	request.Body.BpPayRequest.UserPassword = m.gateway.ApiKeys.PrivateKey
	request.Body.BpPayRequest.OrderId = orderId
	request.Body.BpPayRequest.Amount = int64(testAmount)
	request.Body.BpPayRequest.LocalDate = localDate
	request.Body.BpPayRequest.LocalTime = localTime
	request.Body.BpPayRequest.AdditionalData = testDescription
	request.Body.BpPayRequest.CallBackUrl = testCallback
	request.Body.BpPayRequest.PayerId = 0

	response := &mellatPayResponse{}
	err = m.makeSOAPRequest("bpPayRequest", request, response)
	if err != nil {
		return fmt.Errorf("خطا در اتصال: %v", err)
	}

	// بررسی پاسخ
	result := response.Body.BpPayRequestResponse.Return
	if result == "" {
		return fmt.Errorf("پاسخ خالی از سرور ملت دریافت شد")
	}

	parts := strings.Split(result, ",")
	if len(parts) < 2 {
		return fmt.Errorf("فرمت پاسخ نامعتبر: %s", result)
	}

	resCode := parts[0]
	if resCode == "0" {
		return nil
	} else if resCode == "1" {
		return nil // خطای مبلغ قابل قبول است
	} else {
		return fmt.Errorf("خطای ملت: کد %s - %s", resCode, m.getErrorMessage(resCode))
	}
}

// makeSOAPRequest ارسال درخواست SOAP به ملت
func (m *MellatService) makeSOAPRequest(action string, request interface{}, response interface{}) error {
	// تعیین آدرس Web Service
	wsURL := "https://bpm.shaparak.ir/pgwchannel/services/pgw"
	if m.gateway.IsTestMode {
		wsURL = "https://pgwstest.bpm.bankmellat.ir/pgwchannel/services/pgw"
	}

	// تبدیل درخواست به XML
	xmlData, err := xml.Marshal(request)
	if err != nil {
		return fmt.Errorf("خطا در تبدیل درخواست به XML: %v", err)
	}

	// ایجاد درخواست HTTP
	soapAction := fmt.Sprintf("http://interfaces.core.sw.bps.com/%s", action)
	req, err := http.NewRequest("POST", wsURL, bytes.NewBuffer(xmlData))
	if err != nil {
		return fmt.Errorf("خطا در ایجاد درخواست HTTP: %v", err)
	}

	// تنظیم هدرهای SOAP
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", soapAction)
	req.Header.Set("Accept", "text/xml")

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

	// تبدیل پاسخ XML
	err = xml.Unmarshal(body, response)
	if err != nil {
		return fmt.Errorf("خطا در تبدیل پاسخ XML: %v", err)
	}

	return nil
}

// getErrorMessage تبدیل کد خطای ملت به پیام فارسی
func (m *MellatService) getErrorMessage(code string) string {
	switch code {
	case "0":
		return "موفقیت‌آمیز"
	case "1":
		return "خطای مبلغ"
	case "2":
		return "خطای شماره سفارش"
	case "3":
		return "خطای شماره ترمینال"
	case "4":
		return "خطای نام کاربری یا رمز عبور"
	case "5":
		return "خطای تاریخ"
	case "6":
		return "خطای زمان"
	case "7":
		return "خطای شماره ارجاع"
	case "8":
		return "خطای شماره فروش"
	case "9":
		return "خطای شماره ارجاع فروش"
	case "10":
		return "خطای مبلغ بازگشتی"
	case "11":
		return "خطای شماره سفارش بازگشتی"
	case "12":
		return "خطای شماره ارجاع بازگشتی"
	case "13":
		return "خطای شماره فروش بازگشتی"
	case "14":
		return "خطای شماره ارجاع فروش بازگشتی"
	case "15":
		return "خطای مبلغ تسویه"
	case "16":
		return "خطای شماره سفارش تسویه"
	case "17":
		return "خطای شماره ارجاع تسویه"
	case "18":
		return "خطای شماره فروش تسویه"
	case "19":
		return "خطای شماره ارجاع فروش تسویه"
	case "20":
		return "خطای مبلغ استعلام"
	case "21":
		return "خطای شماره سفارش استعلام"
	case "22":
		return "خطای شماره ارجاع استعلام"
	case "23":
		return "خطای شماره فروش استعلام"
	case "24":
		return "خطای شماره ارجاع فروش استعلام"
	case "25":
		return "خطای مبلغ برگشت"
	case "26":
		return "خطای شماره سفارش برگشت"
	case "27":
		return "خطای شماره ارجاع برگشت"
	case "28":
		return "خطای شماره فروش برگشت"
	case "29":
		return "خطای شماره ارجاع فروش برگشت"
	case "30":
		return "خطای مبلغ تسویه برگشت"
	case "31":
		return "خطای شماره سفارش تسویه برگشت"
	case "32":
		return "خطای شماره ارجاع تسویه برگشت"
	case "33":
		return "خطای شماره فروش تسویه برگشت"
	case "34":
		return "خطای شماره ارجاع فروش تسویه برگشت"
	case "35":
		return "خطای مبلغ استعلام برگشت"
	case "36":
		return "خطای شماره سفارش استعلام برگشت"
	case "37":
		return "خطای شماره ارجاع استعلام برگشت"
	case "38":
		return "خطای شماره فروش استعلام برگشت"
	case "39":
		return "خطای شماره ارجاع فروش استعلام برگشت"
	case "40":
		return "خطای مبلغ تسویه برگشت"
	case "41":
		return "خطای شماره سفارش تسویه برگشت"
	case "42":
		return "خطای شماره ارجاع تسویه برگشت"
	case "43":
		return "خطای شماره فروش تسویه برگشت"
	case "44":
		return "خطای شماره ارجاع فروش تسویه برگشت"
	case "45":
		return "خطای مبلغ استعلام برگشت"
	case "46":
		return "خطای شماره سفارش استعلام برگشت"
	case "47":
		return "خطای شماره ارجاع استعلام برگشت"
	case "48":
		return "خطای شماره فروش استعلام برگشت"
	case "49":
		return "خطای شماره ارجاع فروش استعلام برگشت"
	case "50":
		return "خطای مبلغ تسویه برگشت"
	case "51":
		return "خطای شماره سفارش تسویه برگشت"
	case "52":
		return "خطای شماره ارجاع تسویه برگشت"
	case "53":
		return "خطای شماره فروش تسویه برگشت"
	case "54":
		return "خطای شماره ارجاع فروش تسویه برگشت"
	case "55":
		return "خطای مبلغ استعلام برگشت"
	case "56":
		return "خطای شماره سفارش استعلام برگشت"
	case "57":
		return "خطای شماره ارجاع استعلام برگشت"
	case "58":
		return "خطای شماره فروش استعلام برگشت"
	case "59":
		return "خطای شماره ارجاع فروش استعلام برگشت"
	case "60":
		return "خطای مبلغ تسویه برگشت"
	case "61":
		return "خطای شماره سفارش تسویه برگشت"
	case "62":
		return "خطای شماره ارجاع تسویه برگشت"
	case "63":
		return "خطای شماره فروش تسویه برگشت"
	case "64":
		return "خطای شماره ارجاع فروش تسویه برگشت"
	case "65":
		return "خطای مبلغ استعلام برگشت"
	case "66":
		return "خطای شماره سفارش استعلام برگشت"
	case "67":
		return "خطای شماره ارجاع استعلام برگشت"
	case "68":
		return "خطای شماره فروش استعلام برگشت"
	case "69":
		return "خطای شماره ارجاع فروش استعلام برگشت"
	case "70":
		return "خطای مبلغ تسویه برگشت"
	case "71":
		return "خطای شماره سفارش تسویه برگشت"
	case "72":
		return "خطای شماره ارجاع تسویه برگشت"
	case "73":
		return "خطای شماره فروش تسویه برگشت"
	case "74":
		return "خطای شماره ارجاع فروش تسویه برگشت"
	case "75":
		return "خطای مبلغ استعلام برگشت"
	case "76":
		return "خطای شماره سفارش استعلام برگشت"
	case "77":
		return "خطای شماره ارجاع استعلام برگشت"
	case "78":
		return "خطای شماره فروش استعلام برگشت"
	case "79":
		return "خطای شماره ارجاع فروش استعلام برگشت"
	case "80":
		return "خطای مبلغ تسویه برگشت"
	case "81":
		return "خطای شماره سفارش تسویه برگشت"
	case "82":
		return "خطای شماره ارجاع تسویه برگشت"
	case "83":
		return "خطای شماره فروش تسویه برگشت"
	case "84":
		return "خطای شماره ارجاع فروش تسویه برگشت"
	case "85":
		return "خطای مبلغ استعلام برگشت"
	case "86":
		return "خطای شماره سفارش استعلام برگشت"
	case "87":
		return "خطای شماره ارجاع استعلام برگشت"
	case "88":
		return "خطای شماره فروش استعلام برگشت"
	case "89":
		return "خطای شماره ارجاع فروش استعلام برگشت"
	case "90":
		return "خطای مبلغ تسویه برگشت"
	case "91":
		return "خطای شماره سفارش تسویه برگشت"
	case "92":
		return "خطای شماره ارجاع تسویه برگشت"
	case "93":
		return "خطای شماره فروش تسویه برگشت"
	case "94":
		return "خطای شماره ارجاع فروش تسویه برگشت"
	case "95":
		return "خطای مبلغ استعلام برگشت"
	case "96":
		return "خطای شماره سفارش استعلام برگشت"
	case "97":
		return "خطای شماره ارجاع استعلام برگشت"
	case "98":
		return "خطای شماره فروش استعلام برگشت"
	case "99":
		return "خطای شماره ارجاع فروش استعلام برگشت"
	default:
		return fmt.Sprintf("خطای نامشخص با کد %s", code)
	}
}

// GetBalance دریافت موجودی حساب
func (m *MellatService) GetBalance() (float64, error) {
	// ملت API مستقیم برای دریافت موجودی ندارد
	// این متد برای سازگاری با سایر درگاه‌ها اضافه شده
	return 0, errors.New("ملت API دریافت موجودی ندارد")
}

// ProcessCallback پردازش callback از ملت
func (m *MellatService) ProcessCallback(formData map[string]string) (bool, string, string, error) {
	// بررسی پارامترهای callback
	refId := formData["RefId"]
	resCode := formData["ResCode"]
	saleOrderId := formData["SaleOrderId"]

	if resCode == "" {
		return false, "", "", errors.New("کد نتیجه مشخص نشده")
	}

	if resCode == "0" {
		if refId == "" {
			return false, "", "", errors.New("شماره ارجاع دریافت نشد")
		}
		return true, refId, saleOrderId, nil
	} else {
		// بررسی انواع خطا
		switch resCode {
		case "1":
			return false, "", "", errors.New("خطای مبلغ")
		case "2":
			return false, "", "", errors.New("خطای شماره سفارش")
		case "3":
			return false, "", "", errors.New("خطای شماره ترمینال")
		case "4":
			return false, "", "", errors.New("خطای نام کاربری یا رمز عبور")
		case "5":
			return false, "", "", errors.New("خطای تاریخ")
		case "6":
			return false, "", "", errors.New("خطای زمان")
		case "7":
			return false, "", "", errors.New("خطای شماره ارجاع")
		case "8":
			return false, "", "", errors.New("خطای شماره فروش")
		case "9":
			return false, "", "", errors.New("خطای شماره ارجاع فروش")
		default:
			return false, "", "", fmt.Errorf("خطای نامشخص: کد %s", resCode)
		}
	}
}
