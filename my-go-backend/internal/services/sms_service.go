package services

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	ippanel "github.com/ippanel/go-rest-sdk/v2"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
)

// SMSService سرویس ارسال پیامک
type SMSService struct {
	db             *gorm.DB
	failedGateways map[uint]time.Time
	mu             sync.RWMutex
}

// NewSMSService ایجاد نمونه جدید از سرویس پیامک
func NewSMSService(db *gorm.DB) *SMSService {
	service := &SMSService{
		db:             db,
		failedGateways: make(map[uint]time.Time),
	}
	go service.startHealthChecker()
	return service
}

// startHealthChecker بررسی سلامت درگاه‌های از کار افتاده در فواصل زمانی ثابت
func (s *SMSService) startHealthChecker() {
	ticker := time.NewTicker(30 * time.Minute)
	for range ticker.C {
		s.runHealthChecks()
	}
}

// runHealthChecks اجرای بررسی سلامت برای درگاه‌های خراب
func (s *SMSService) runHealthChecks() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for gatewayID, failedAt := range s.failedGateways {
		if time.Since(failedAt) < 30*time.Minute {
			continue
		}

		var gateway models.SMSGateway
		if err := s.db.First(&gateway, gatewayID).Error; err != nil {
			delete(s.failedGateways, gatewayID)
			continue
		}

		if s.healthCheckGateway(gateway) {
			delete(s.failedGateways, gatewayID)
		} else {
			s.failedGateways[gatewayID] = time.Now()
		}
	}
}

// healthCheckGateway بررسی سلامت یک درگاه خاص
func (s *SMSService) healthCheckGateway(gateway models.SMSGateway) bool {
	_, err := s.GetGatewayBalance(gateway.ID)
	return err == nil
}

// markGatewayFailed علامت‌گذاری درگاه به عنوان خراب
func (s *SMSService) markGatewayFailed(id uint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.failedGateways[id] = time.Now()
}

// SendSMS ارسال پیامک با پترن
func (s *SMSService) SendSMS(request models.SMSSendRequest) (*models.SMSSendResponse, error) {
	log.Printf("🔍 درخواست ارسال پیامک: PatternCode='%s', Mobile='%s', GatewayID=%d, Scope='%s', Feature='%s'", request.PatternCode, request.Mobile, request.GatewayID, request.Scope, request.Feature)
	log.Printf("🔍 طول PatternCode: %d", len(request.PatternCode))

	// دریافت لیست درگاه‌های فعال بر اساس اولویت
	var gateways []models.SMSGateway
	if err := s.db.Where("is_active = ?", true).Order("priority ASC").Find(&gateways).Error; err != nil {
		return nil, fmt.Errorf("خطا در دریافت درگاه‌ها: %w", err)
	}

	if len(gateways) == 0 {
		return nil, fmt.Errorf("هیچ درگاه فعالی یافت نشد")
	}

	log.Printf("درگاه‌های فعال یافت شدند: %d درگاه", len(gateways))
	for i, gw := range gateways {
		log.Printf("درگاه %d: ID=%d, Type=%s, Name=%s, PatternBased=%v", i+1, gw.ID, gw.Type, gw.Name, gw.PatternBased)
	}

	// اگر کاربر درگاه مشخصی را تعیین کرده باشد، ابتدا همان را امتحان می‌کنیم
	if request.GatewayID != 0 {
		for i, gw := range gateways {
			if gw.ID == request.GatewayID {
				gateways[0], gateways[i] = gateways[i], gateways[0]
				break
			}
		}
	}

	// حلقه روی درگاه‌ها به ترتیب اولویت
	for _, gateway := range gateways {
		// رد شدن از درگاه‌های خراب
		s.mu.RLock()
		_, failed := s.failedGateways[gateway.ID]
		s.mu.RUnlock()
		if failed {
			continue
		}

		log.Printf("🔍 بررسی درگاه: ID=%d, Type=%s, PatternBased=%v", gateway.ID, gateway.Type, gateway.PatternBased)

		// بررسی اینکه آیا این درگاه از پترن استفاده می‌کند
		if gateway.PatternBased {
			// جستجوی پترن بر اساس scope و feature
			var pattern models.SMSPattern
			query := s.db.Where("gateway_id = ? AND status = ?", gateway.ID, "active")

			// اگر scope و feature مشخص شده باشند، بر اساس آن‌ها جستجو کن
			if request.Scope != "" && request.Feature != "" {
				query = query.Where("scope = ? AND feature = ?", request.Scope, request.Feature)
				log.Printf("🔍 جستجوی پترن: GatewayID=%d, Scope='%s', Feature='%s'", gateway.ID, request.Scope, request.Feature)
			} else if request.PatternCode != "" {
				// اگر pattern_code مشخص شده باشد، بر اساس آن جستجو کن
				query = query.Where("pattern_code = ?", request.PatternCode)
				log.Printf("🔍 جستجوی پترن: PatternCode='%s', GatewayID=%d", request.PatternCode, gateway.ID)
			} else {
				// اگر هیچ‌کدام مشخص نشده باشد، خطا بده
				log.Printf("❌ هیچ معیاری برای جستجوی پترن مشخص نشده است")
				continue
			}

			if err := query.First(&pattern).Error; err != nil {
				log.Printf("❌ پترن مناسب برای درگاه %d یافت نشد: %v", gateway.ID, err)
				continue
			}

			log.Printf("✅ پترن یافت شد: ID=%d, PatternCode=%s, GatewayID=%d, Scope=%s, Feature=%s", pattern.ID, pattern.PatternCode, pattern.GatewayID, pattern.Scope, pattern.Feature)

			var resp *models.SMSSendResponse
			var err error

			switch strings.ToLower(gateway.Type) {
			case "ippanel":
				resp, err = s.sendWithIPPanel(pattern, gateway, request)
			case "kavenegar":
				resp, err = s.sendWithKavenegar(pattern, gateway, request)
			case "meli_payamak":
				resp, err = s.sendWithMeliPayamak(pattern, gateway, request)
			case "farazsms":
				resp, err = s.sendWithFarazSMS(pattern, gateway, request)
			default:
				err = fmt.Errorf("درگاه پیامک %s پشتیبانی نمی‌شود", gateway.Type)
			}

			if err == nil {
				return resp, nil
			}

			// در صورت خطا، درگاه را خراب علامت بزنیم و ادامه دهیم
			log.Printf("خطا در ارسال پیامک با درگاه %s (ID=%d): %v", gateway.Type, gateway.ID, err)
			s.markGatewayFailed(gateway.ID)
		} else {
			// اگر درگاه از پترن استفاده نمی‌کند، ارسال مستقیم
			log.Printf("🔍 درگاه %d از پترن استفاده نمی‌کند، ارسال مستقیم", gateway.ID)

			var resp *models.SMSSendResponse
			var err error

			switch strings.ToLower(gateway.Type) {
			case "ippanel":
				resp, err = s.sendWithIPPanelDirect(gateway, request)
			case "kavenegar":
				// برای درگاه‌های دیگر که هنوز پیاده‌سازی نشده‌اند، خطا بده
				err = fmt.Errorf("ارسال مستقیم با کاوه‌نگار هنوز پیاده‌سازی نشده است")
			case "meli_payamak":
				err = fmt.Errorf("ارسال مستقیم با ملی پیامک هنوز پیاده‌سازی نشده است")
			case "farazsms":
				err = fmt.Errorf("ارسال مستقیم با فراز اس‌ام‌اس هنوز پیاده‌سازی نشده است")
			default:
				err = fmt.Errorf("درگاه پیامک %s پشتیبانی نمی‌شود", gateway.Type)
			}

			if err == nil {
				return resp, nil
			}

			// در صورت خطا، درگاه را خراب علامت بزنیم و ادامه دهیم
			log.Printf("خطا در ارسال پیامک با درگاه %s (ID=%d): %v", gateway.Type, gateway.ID, err)
			s.markGatewayFailed(gateway.ID)
		}
	}

	return nil, fmt.Errorf("ارسال پیامک با تمام درگاه‌ها ناموفق بود")
}

// sendWithIPPanel ارسال پیامک با درگاه IPPanel
func (s *SMSService) sendWithIPPanel(pattern models.SMSPattern, gateway models.SMSGateway, request models.SMSSendRequest) (*models.SMSSendResponse, error) {
	log.Printf("🔍 IPPanel: PatternCode='%s', Variables=%v", pattern.PatternCode, request.Variables)

	// ایجاد کلاینت IPPanel
	client := ippanel.New(gateway.ApiKey)

	// آماده‌سازی متغیرهای پترن
	patternValues := make(map[string]string)

	// متغیرهای ارسالی از درخواست را اضافه کن
	for key, value := range request.Variables {
		patternValues[key] = value
	}

	// اگر متغیر date وجود دارد، ساعت فعلی را به آن اضافه کن
	if _, exists := patternValues["date"]; exists {
		now := time.Now()
		timeStr := now.Format("15:04")
		patternValues["date"] = patternValues["date"] + " " + timeStr
	}

	log.Printf("🔍 متغیرهای پترن: %v", patternValues)
	// فرمت‌بندی شماره موبایل برای IPPanel (فرمت بین‌المللی)
	formattedMobile := request.Mobile
	if !strings.HasPrefix(formattedMobile, "+") {
		if strings.HasPrefix(formattedMobile, "0") {
			// تبدیل 09203214155 به +989203214155
			formattedMobile = "+98" + strings.TrimPrefix(formattedMobile, "0")
		} else if strings.HasPrefix(formattedMobile, "98") {
			// تبدیل 989203214155 به +989203214155
			formattedMobile = "+" + formattedMobile
		}
	}

	log.Printf("🔍 شماره موبایل فرمت شده: '%s'", formattedMobile)

	// ارسال با پترن
	log.Printf("📤 آماده ارسال با SendPattern...")
	log.Printf("📤 پارامترها: pattern_code='%s', sender='%s', mobile='%s', variables=%v",
		pattern.PatternCode, gateway.SenderNumber, formattedMobile, patternValues)

	messageID, err := client.SendPattern(
		pattern.PatternCode,  // کد پترن
		gateway.SenderNumber, // شماره فرستنده
		formattedMobile,      // شماره گیرنده (فرمت شده)
		patternValues,        // متغیرهای پترن
	)

	log.Printf("🔍 نتیجه SendPattern: messageID=%v, err=%v", messageID, err)
	log.Printf("🔍 نتیجه SendPattern: messageID=%v, err=%v", messageID, err)

	// بررسی اینکه آیا messageID معتبر است
	if messageID == 0 {
		log.Printf("❌ messageID صفر است - احتمالاً مشکلی در پارامترها وجود دارد")
		log.Printf("❌ بررسی پارامترها:")
		log.Printf("❌   - PatternCode: '%s'", pattern.PatternCode)
		log.Printf("❌   - SenderNumber: '%s'", gateway.SenderNumber)
		log.Printf("❌   - Mobile: '%s'", formattedMobile)
		log.Printf("❌   - Variables: %v", patternValues)
		return nil, fmt.Errorf("خطا در ارسال پترن: messageID صفر است")
	}

	if err != nil {
		log.Printf("❌ خطا در SendPattern: %v", err)
		log.Printf("❌ نوع خطا: %T", err)
		// بررسی نوع خطا
		var ippanelErr ippanel.Error
		if errors.As(err, &ippanelErr) {
			log.Printf("❌ خطای IPPanel: Code=%d, Message=%v", ippanelErr.Code, ippanelErr.Message)
			log.Printf("❌ جزئیات خطا: %+v", ippanelErr)
			// نمایش پیام خطا به صورت string
			if errMsg, ok := ippanelErr.Message.(string); ok {
				log.Printf("❌ پیام خطا: %s", errMsg)
			} else {
				log.Printf("❌ پیام خطا (نوع نامشخص): %v", ippanelErr.Message)
			}
			switch ippanelErr.Code {
			case ippanel.ErrUnprocessableEntity:
				return nil, fmt.Errorf("خطا در اعتبارسنجی داده‌ها: %v", ippanelErr.Message)
			case 401: // ippanel.ErrStatusUnauthorized
				return nil, fmt.Errorf("کلید API نامعتبر است")
			default:
				return nil, fmt.Errorf("خطای IPPanel: %v", ippanelErr.Message)
			}
		}
		return nil, fmt.Errorf("خطا در ارسال پیامک: %w", err)
	}

	// افزایش شمارنده استفاده
	s.db.Model(&pattern).Update("usage_count", pattern.UsageCount+1)

	// ساخت پاسخ
	response := &models.SMSSendResponse{
		Success: true,
		Message: "پیامک با موفقیت ارسال شد",
		Data: struct {
			MessageID string `json:"message_id,omitempty"`
			Cost      int    `json:"cost,omitempty"`
			Balance   int    `json:"balance,omitempty"`
		}{
			MessageID: fmt.Sprintf("%d", messageID),
			Cost:      0, // اطلاعات هزینه در دسترس نیست
			Balance:   0, // باید از API موجودی دریافت شود
		},
	}

	return response, nil
}

// sendWithIPPanelDirect ارسال پیامک بدون پترن (مستقیم)
func (s *SMSService) sendWithIPPanelDirect(gateway models.SMSGateway, request models.SMSSendRequest) (*models.SMSSendResponse, error) {
	log.Printf("🔍 IPPanel Direct: Mobile='%s', Message='%s'", request.Mobile, request.Message)

	// ایجاد کلاینت IPPanel
	client := ippanel.New(gateway.ApiKey)

	// فرمت‌بندی شماره موبایل برای IPPanel (فرمت بین‌المللی)
	formattedMobile := request.Mobile
	if !strings.HasPrefix(formattedMobile, "+") {
		if strings.HasPrefix(formattedMobile, "0") {
			// تبدیل 09203214155 به +989203214155
			formattedMobile = "+98" + strings.TrimPrefix(formattedMobile, "0")
		} else if strings.HasPrefix(formattedMobile, "98") {
			// تبدیل 989203214155 به +989203214155
			formattedMobile = "+" + formattedMobile
		}
	}

	log.Printf("🔍 شماره موبایل فرمت شده: '%s'", formattedMobile)

	// ارسال مستقیم بدون پترن
	log.Printf("📤 آماده ارسال مستقیم...")
	log.Printf("📤 پارامترها: sender='%s', mobile='%s', message='%s'",
		gateway.SenderNumber, formattedMobile, request.Message)

	messageID, err := client.Send(
		gateway.SenderNumber,      // شماره فرستنده
		[]string{formattedMobile}, // لیست شماره‌های گیرنده
		request.Message,           // متن پیام
		"normal",                  // نوع پیام (normal برای پیامک معمولی)
	)

	log.Printf("🔍 نتیجه Send: messageID=%v, err=%v", messageID, err)
	if err != nil {
		log.Printf("❌ خطا در Send: %v", err)
		log.Printf("❌ نوع خطا: %T", err)
		// بررسی نوع خطا
		var ippanelErr ippanel.Error
		if errors.As(err, &ippanelErr) {
			log.Printf("❌ خطای IPPanel: Code=%d, Message=%v", ippanelErr.Code, ippanelErr.Message)
			log.Printf("❌ جزئیات خطا: %+v", ippanelErr)
			// نمایش پیام خطا به صورت string
			if errMsg, ok := ippanelErr.Message.(string); ok {
				log.Printf("❌ پیام خطا: %s", errMsg)
			} else {
				log.Printf("❌ پیام خطا (نوع نامشخص): %v", ippanelErr.Message)
			}
			switch ippanelErr.Code {
			case ippanel.ErrUnprocessableEntity:
				return nil, fmt.Errorf("خطا در اعتبارسنجی داده‌ها: %v", ippanelErr.Message)
			case 401: // ippanel.ErrStatusUnauthorized
				return nil, fmt.Errorf("کلید API نامعتبر است")
			default:
				return nil, fmt.Errorf("خطای IPPanel: %v", ippanelErr.Message)
			}
		}
		return nil, fmt.Errorf("خطا در ارسال پیامک: %w", err)
	}

	// ساخت پاسخ
	response := &models.SMSSendResponse{
		Success: true,
		Message: "پیامک با موفقیت ارسال شد",
		Data: struct {
			MessageID string `json:"message_id,omitempty"`
			Cost      int    `json:"cost,omitempty"`
			Balance   int    `json:"balance,omitempty"`
		}{
			MessageID: fmt.Sprintf("%d", messageID),
			Cost:      0, // اطلاعات هزینه در دسترس نیست
			Balance:   0, // باید از API موجودی دریافت شود
		},
	}

	return response, nil
}

// sendWithKavenegar ارسال پیامک با درگاه کاوه‌نگار
func (s *SMSService) sendWithKavenegar(_ models.SMSPattern, _ models.SMSGateway, _ models.SMSSendRequest) (*models.SMSSendResponse, error) {
	// TODO: پیاده‌سازی ارسال با کاوه‌نگار
	return nil, fmt.Errorf("ارسال با کاوه‌نگار هنوز پیاده‌سازی نشده است")
}

// sendWithMeliPayamak ارسال پیامک با درگاه ملی پیامک
func (s *SMSService) sendWithMeliPayamak(_ models.SMSPattern, _ models.SMSGateway, _ models.SMSSendRequest) (*models.SMSSendResponse, error) {
	// TODO: پیاده‌سازی ارسال با ملی پیامک
	return nil, fmt.Errorf("ارسال با ملی پیامک هنوز پیاده‌سازی نشده است")
}

// sendWithFarazSMS ارسال پیامک با درگاه فراز اس‌ام‌اس
func (s *SMSService) sendWithFarazSMS(_ models.SMSPattern, _ models.SMSGateway, _ models.SMSSendRequest) (*models.SMSSendResponse, error) {
	// TODO: پیاده‌سازی ارسال با فراز اس‌ام‌اس
	return nil, fmt.Errorf("ارسال با فراز اس‌ام‌اس هنوز پیاده‌سازی نشده است")
}

// GetGatewayBalance دریافت موجودی درگاه
func (s *SMSService) GetGatewayBalance(gatewayID uint) (float64, error) {
	var gateway models.SMSGateway
	if err := s.db.First(&gateway, gatewayID).Error; err != nil {
		return 0, fmt.Errorf("درگاه یافت نشد: %w", err)
	}

	switch strings.ToLower(gateway.Type) {
	case "ippanel":
		return s.getIPPanelBalance(gateway)
	case "kavenegar":
		return s.getKavenegarBalance(gateway)
	case "meli_payamak":
		return s.getMeliPayamakBalance(gateway)
	case "farazsms":
		return s.getFarazSMSBalance(gateway)
	default:
		return 0, fmt.Errorf("درگاه %s پشتیبانی نمی‌شود", gateway.Type)
	}
}

// getIPPanelBalance دریافت موجودی IPPanel
func (s *SMSService) getIPPanelBalance(gateway models.SMSGateway) (float64, error) {
	client := ippanel.New(gateway.ApiKey)
	balance, err := client.GetCredit()
	if err != nil {
		var ippanelErr ippanel.Error
		if errors.As(err, &ippanelErr) {
			switch ippanelErr.Code {
			case 401: // ippanel.ErrStatusUnauthorized
				return 0, fmt.Errorf("کلید API نامعتبر است")
			default:
				return 0, fmt.Errorf("خطای IPPanel: %v", ippanelErr.Message)
			}
		}
		return 0, fmt.Errorf("خطا در دریافت موجودی: %w", err)
	}
	return balance, nil
}

// getKavenegarBalance دریافت موجودی کاوه‌نگار
func (s *SMSService) getKavenegarBalance(_ models.SMSGateway) (float64, error) {
	// TODO: پیاده‌سازی دریافت موجودی کاوه‌نگار
	return 0, fmt.Errorf("دریافت موجودی کاوه‌نگار هنوز پیاده‌سازی نشده است")
}

// getMeliPayamakBalance دریافت تعداد باقی‌مانده SMS ملی پیامک
func (s *SMSService) getMeliPayamakBalance(gateway models.SMSGateway) (float64, error) {
	// ساخت درخواست SOAP برای دریافت تعداد باقی‌مانده SMS
	soapEnvelope := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <GetCredit xmlns="http://tempuri.org/">
      <username>%s</username>
      <password>%s</password>
    </GetCredit>
  </soap:Body>
</soap:Envelope>`, gateway.Username, gateway.Password)

	// ارسال درخواست SOAP
	resp, err := http.Post(gateway.ApiURL, "text/xml; charset=utf-8", strings.NewReader(soapEnvelope))
	if err != nil {
		return 0, fmt.Errorf("خطا در ارسال درخواست: %w", err)
	}
	defer resp.Body.Close()

	// خواندن پاسخ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("خطا در خواندن پاسخ: %w", err)
	}

	// بررسی کد وضعیت HTTP
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("خطای HTTP: %d - %s", resp.StatusCode, string(body))
	}

	// ساختار برای پردازش پاسخ XML
	type GetCreditResponse struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			GetCreditResponse struct {
				GetCreditResult float64 `xml:"GetCreditResult"`
			} `xml:"GetCreditResponse"`
		} `xml:"Body"`
	}

	// پردازش پاسخ XML
	var response GetCreditResponse
	if err := xml.Unmarshal(body, &response); err != nil {
		return 0, fmt.Errorf("خطا در پردازش XML: %w", err)
	}

	balance := response.Body.GetCreditResponse.GetCreditResult

	return balance, nil
}

// getMeliPayamakCredit دریافت موجودی ریالی ملی پیامک
func (s *SMSService) getMeliPayamakCredit(gateway models.SMSGateway) (float64, error) {
	// ساخت درخواست SOAP برای دریافت موجودی ریالی
	soapEnvelope := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <GetUserCredit2 xmlns="http://tempuri.org/">
      <username>%s</username>
      <password>%s</password>
    </GetUserCredit2>
  </soap:Body>
</soap:Envelope>`, gateway.Username, gateway.Password)

	// ارسال درخواست SOAP به آدرس Users.asmx
	usersURL := "https://api.payamak-panel.com/post/Users.asmx"
	resp, err := http.Post(usersURL, "text/xml; charset=utf-8", strings.NewReader(soapEnvelope))
	if err != nil {
		return 0, fmt.Errorf("خطا در ارسال درخواست: %w", err)
	}
	defer resp.Body.Close()

	// خواندن پاسخ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("خطا در خواندن پاسخ: %w", err)
	}

	// بررسی کد وضعیت HTTP
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("خطای HTTP: %d - %s", resp.StatusCode, string(body))
	}

	// ساختار برای پردازش پاسخ XML
	type GetUserCredit2Response struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			GetUserCredit2Response struct {
				GetUserCredit2Result float64 `xml:"GetUserCredit2Result"`
			} `xml:"GetUserCredit2Response"`
		} `xml:"Body"`
	}

	// پردازش پاسخ XML
	var response GetUserCredit2Response
	if err := xml.Unmarshal(body, &response); err != nil {
		return 0, fmt.Errorf("خطا در پردازش XML: %w", err)
	}

	credit := response.Body.GetUserCredit2Response.GetUserCredit2Result

	return credit, nil
}

// getFarazSMSBalance دریافت موجودی فراز اس‌ام‌اس
func (s *SMSService) getFarazSMSBalance(_ models.SMSGateway) (float64, error) {
	// TODO: پیاده‌سازی دریافت موجودی فراز اس‌ام‌اس
	return 0, fmt.Errorf("دریافت موجودی فراز اس‌ام‌اس هنوز پیاده‌سازی نشده است")
}

// GetMeliPayamakInfo دریافت اطلاعات کامل ملی پیامک (تعداد SMS و موجودی ریالی)
func (s *SMSService) GetMeliPayamakInfo(gatewayID uint) (map[string]interface{}, error) {
	var gateway models.SMSGateway
	if err := s.db.First(&gateway, gatewayID).Error; err != nil {
		return nil, fmt.Errorf("درگاه یافت نشد: %w", err)
	}

	// بررسی نوع درگاه
	if strings.ToLower(gateway.Type) != "meli_payamak" {
		return nil, fmt.Errorf("این تابع فقط برای درگاه ملی پیامک قابل استفاده است")
	}

	// دریافت تعداد باقی‌مانده SMS
	remainingSMS, err := s.getMeliPayamakBalance(gateway)
	if err != nil {
		return nil, fmt.Errorf("خطا در دریافت تعداد باقی‌مانده SMS: %w", err)
	}

	// دریافت موجودی ریالی
	credit, err := s.getMeliPayamakCredit(gateway)
	if err != nil {
		return nil, fmt.Errorf("خطا در دریافت موجودی ریالی: %w", err)
	}

	// ساخت پاسخ
	result := map[string]interface{}{
		"remaining_sms": int(remainingSMS), // تعداد باقی‌مانده SMS
		"credit":        credit,            // موجودی ریالی
		"gateway_name":  gateway.Name,
		"gateway_type":  gateway.Type,
	}

	return result, nil
}

// GetInboxMessages دریافت پیام‌های دریافتی
func (s *SMSService) GetInboxMessages(gatewayID uint, page, limit int64) ([]models.InboxMessage, *models.InboxData, error) {
	var gateway models.SMSGateway
	if err := s.db.First(&gateway, gatewayID).Error; err != nil {
		return nil, nil, fmt.Errorf("درگاه یافت نشد: %w", err)
	}

	switch strings.ToLower(gateway.Type) {
	case "ippanel":
		return s.getIPPanelInbox(gateway, page, limit)
	default:
		return nil, nil, fmt.Errorf("درگاه %s پشتیبانی نمی‌شود", gateway.Type)
	}
}

// getIPPanelInbox دریافت پیام‌های دریافتی IPPanel
func (s *SMSService) getIPPanelInbox(gateway models.SMSGateway, page, limit int64) ([]models.InboxMessage, *models.InboxData, error) {
	client := ippanel.New(gateway.ApiKey)

	paginationParams := ippanel.ListParams{
		Page:  page,
		Limit: limit,
	}

	messages, paginationInfo, err := client.FetchInbox(paginationParams)
	if err != nil {
		var ippanelErr ippanel.Error
		if errors.As(err, &ippanelErr) {
			switch ippanelErr.Code {
			case 401: // ippanel.ErrStatusUnauthorized
				return nil, nil, fmt.Errorf("کلید API نامعتبر است")
			default:
				return nil, nil, fmt.Errorf("خطای IPPanel: %v", ippanelErr.Message)
			}
		}
		return nil, nil, fmt.Errorf("خطا در دریافت پیام‌های دریافتی: %w", err)
	}

	// تبدیل به مدل داخلی
	var inboxMessages []models.InboxMessage
	for _, msg := range messages {
		inboxMessages = append(inboxMessages, models.InboxMessage{
			To:        msg.To,
			Message:   msg.Message,
			From:      msg.From,
			CreatedAt: msg.CreatedAt,
			Type:      msg.Type,
		})
	}

	inboxData := &models.InboxData{
		Messages:   inboxMessages,
		Total:      int(paginationInfo.Total),
		Page:       int(paginationInfo.Page),
		Limit:      int(paginationInfo.Limit),
		TotalPages: int(paginationInfo.Pages),
	}

	return inboxMessages, inboxData, nil
}

// SendPatternWithGateway ارسال پیامک با پترن از درگاه مشخص شده
func (s *SMSService) SendPatternWithGateway(patternCode string, gatewayID uint, recipient string, patternValues map[string]string) (string, error) {
	log.Printf("🔍 SendPatternWithGateway: PatternCode='%s', GatewayID=%d, Recipient='%s'", patternCode, gatewayID, recipient)
	log.Printf("🔍 متغیرهای پترن: %v", patternValues)

	// دریافت درگاه مشخص شده
	var gateway models.SMSGateway
	if err := s.db.Where("id = ? AND is_active = ? AND pattern_based = ?", gatewayID, true, true).First(&gateway).Error; err != nil {
		return "", fmt.Errorf("درگاه فعالی با قابلیت پترن یافت نشد: %w", err)
	}

	log.Printf("🔍 درگاه یافت شد: ID=%d, Type=%s, PatternBased=%v", gateway.ID, gateway.Type, gateway.PatternBased)

	// بررسی وجود پترن در این درگاه
	var pattern models.SMSPattern
	if err := s.db.Where("pattern_code = ? AND gateway_id = ? AND status = ?", patternCode, gateway.ID, "active").First(&pattern).Error; err != nil {
		return "", fmt.Errorf("پترن با کد '%s' برای درگاه %d یافت نشد: %w", patternCode, gateway.ID, err)
	}

	log.Printf("✅ پترن یافت شد: ID=%d, PatternCode=%s, GatewayID=%d, Scope=%s, Feature=%s", pattern.ID, pattern.PatternCode, pattern.GatewayID, pattern.Scope, pattern.Feature)

	var messageID string
	var err error

	switch strings.ToLower(gateway.Type) {
	case "ippanel":
		messageID, err = s.sendPatternWithIPPanel(pattern, gateway, recipient, patternValues)
	case "kavenegar":
		messageID, err = s.sendPatternWithKavenegar(pattern, gateway, recipient, patternValues)
	case "meli_payamak":
		messageID, err = s.sendPatternWithMeliPayamak(pattern, gateway, recipient, patternValues)
	case "novin":
		messageID, err = s.sendPatternWithFarazSMS(pattern, gateway, recipient, patternValues)
	default:
		err = fmt.Errorf("درگاه پیامک %s پشتیبانی نمی‌شود", gateway.Type)
	}

	if err != nil {
		// در صورت خطا، درگاه را خراب علامت بزنیم
		log.Printf("خطا در ارسال پترن با درگاه %s (ID=%d): %v", gateway.Type, gateway.ID, err)
		s.markGatewayFailed(gateway.ID)
		return "", err
	}

	// افزایش شمارنده استفاده
	s.db.Model(&pattern).Update("usage_count", pattern.UsageCount+1)
	return messageID, nil
}

// SendPattern ارسال پیامک با پترن
func (s *SMSService) SendPattern(patternCode, originator, recipient string, patternValues map[string]string) (string, error) {
	log.Printf("🔍 SendPattern: PatternCode='%s', Originator='%s', Recipient='%s'", patternCode, originator, recipient)
	log.Printf("🔍 متغیرهای پترن: %v", patternValues)

	// دریافت درگاه‌های فعال
	var gateways []models.SMSGateway
	if err := s.db.Where("is_active = ? AND pattern_based = ?", true, true).Order("priority ASC").Find(&gateways).Error; err != nil {
		return "", fmt.Errorf("خطا در دریافت درگاه‌های فعال: %w", err)
	}

	if len(gateways) == 0 {
		return "", fmt.Errorf("هیچ درگاه فعالی با قابلیت پترن یافت نشد")
	}

	// حلقه روی درگاه‌ها
	for _, gateway := range gateways {
		log.Printf("🔍 بررسی درگاه: ID=%d, Type=%s, PatternBased=%v", gateway.ID, gateway.Type, gateway.PatternBased)

		// بررسی اینکه آیا این درگاه از پترن پشتیبانی می‌کند
		if !gateway.PatternBased {
			log.Printf("❌ درگاه %d از پترن پشتیبانی نمی‌کند", gateway.ID)
			continue
		}

		// بررسی وجود پترن در این درگاه
		var pattern models.SMSPattern
		if err := s.db.Where("pattern_code = ? AND gateway_id = ? AND status = ?", patternCode, gateway.ID, "active").First(&pattern).Error; err != nil {
			log.Printf("❌ پترن با کد '%s' برای درگاه %d یافت نشد: %v", patternCode, gateway.ID, err)
			continue
		}

		log.Printf("✅ پترن یافت شد: ID=%d, PatternCode=%s, GatewayID=%d, Scope=%s, Feature=%s", pattern.ID, pattern.PatternCode, pattern.GatewayID, pattern.Scope, pattern.Feature)

		var messageID string
		var err error

		switch strings.ToLower(gateway.Type) {
		case "ippanel":
			messageID, err = s.sendPatternWithIPPanel(pattern, gateway, recipient, patternValues)
		case "kavenegar":
			messageID, err = s.sendPatternWithKavenegar(pattern, gateway, recipient, patternValues)
		case "meli_payamak":
			messageID, err = s.sendPatternWithMeliPayamak(pattern, gateway, recipient, patternValues)
		case "novin":
			messageID, err = s.sendPatternWithFarazSMS(pattern, gateway, recipient, patternValues)
		default:
			err = fmt.Errorf("درگاه پیامک %s پشتیبانی نمی‌شود", gateway.Type)
		}

		if err == nil {
			// افزایش شمارنده استفاده
			s.db.Model(&pattern).Update("usage_count", pattern.UsageCount+1)
			return messageID, nil
		}

		// در صورت خطا، درگاه را خراب علامت بزنیم و ادامه دهیم
		log.Printf("خطا در ارسال پترن با درگاه %s (ID=%d): %v", gateway.Type, gateway.ID, err)
		s.markGatewayFailed(gateway.ID)
	}

	return "", fmt.Errorf("ارسال پترن با تمام درگاه‌ها ناموفق بود")
}

// sendPatternWithIPPanel ارسال پترن با IPPanel
func (s *SMSService) sendPatternWithIPPanel(pattern models.SMSPattern, gateway models.SMSGateway, recipient string, patternValues map[string]string) (string, error) {
	log.Printf("🔍 IPPanel Pattern: PatternCode='%s', Variables=%v", pattern.PatternCode, patternValues)

	log.Printf("🔍 Gateway Info: ID=%d, Name='%s', Type='%s', SenderNumber='%s', ApiKey length=%d",
		gateway.ID, gateway.Name, gateway.Type, gateway.SenderNumber, len(gateway.ApiKey))

	// بررسی اعتبار کلید API
	if gateway.ApiKey == "" {
		return "", fmt.Errorf("کلید API درگاه خالی است")
	}

	// ایجاد کلاینت IPPanel
	client := ippanel.New(gateway.ApiKey)

	// فرمت‌بندی شماره موبایل برای IPPanel (فرمت بین‌المللی)
	formattedMobile := recipient
	if !strings.HasPrefix(formattedMobile, "+") {
		if strings.HasPrefix(formattedMobile, "0") {
			// تبدیل 09203214155 به +989203214155
			formattedMobile = "+98" + strings.TrimPrefix(formattedMobile, "0")
		} else if strings.HasPrefix(formattedMobile, "98") {
			// تبدیل 989203214155 به +989203214155
			formattedMobile = "+" + formattedMobile
		}
	}

	log.Printf("🔍 شماره موبایل فرمت شده: '%s'", formattedMobile)

	// بررسی اعتبار شماره فرستنده
	if gateway.SenderNumber == "" {
		return "", fmt.Errorf("شماره فرستنده درگاه خالی است")
	}

	// ارسال با پترن
	log.Printf("📤 آماده ارسال با SendPattern...")
	log.Printf("📤 پارامترها: pattern_code='%s', sender='%s', mobile='%s', variables=%v",
		pattern.PatternCode, gateway.SenderNumber, formattedMobile, patternValues)

	messageID, err := client.SendPattern(
		pattern.PatternCode,  // کد پترن
		gateway.SenderNumber, // شماره فرستنده
		formattedMobile,      // شماره گیرنده (فرمت شده)
		patternValues,        // متغیرهای پترن
	)

	log.Printf("🔍 نتیجه SendPattern: messageID=%v, err=%v", messageID, err)
	if err != nil {
		log.Printf("❌ خطا در SendPattern: %v", err)
		log.Printf("❌ نوع خطا: %T", err)
		// بررسی نوع خطا
		var ippanelErr ippanel.Error
		if errors.As(err, &ippanelErr) {
			log.Printf("❌ خطای IPPanel: Code=%d, Message=%v", ippanelErr.Code, ippanelErr.Message)
			log.Printf("❌ جزئیات خطا: %+v", ippanelErr)

			// بهبود نمایش پیام خطا
			var errorMessage string
			switch msg := ippanelErr.Message.(type) {
			case string:
				errorMessage = msg
			case map[string]interface{}:
				// تبدیل map به string
				if len(msg) > 0 {
					errorMessage = fmt.Sprintf("%v", msg)
				} else {
					errorMessage = "خطای نامشخص"
				}
			case []interface{}:
				errorMessage = fmt.Sprintf("%v", msg)
			default:
				errorMessage = fmt.Sprintf("%v", ippanelErr.Message)
			}

			log.Printf("❌ پیام خطا پردازش شده: %s", errorMessage)

			switch ippanelErr.Code {
			case ippanel.ErrUnprocessableEntity:
				return "", fmt.Errorf("خطا در اعتبارسنجی داده‌ها: %s", errorMessage)
			case 401: // ippanel.ErrStatusUnauthorized
				return "", fmt.Errorf("کلید API نامعتبر است")
			case 400:
				// بررسی دلیل خطای 400
				if len(patternValues) == 0 {
					return "", fmt.Errorf("خطای 400 IPPanel - متغیرهای پترن خالی هستند. لطفاً متغیرهای مورد نیاز را وارد کنید")
				}
				return "", fmt.Errorf("خطای 400 IPPanel - احتمالاً کد پترن یا شماره فرستنده نامعتبر است: %s", errorMessage)
			default:
				return "", fmt.Errorf("خطای IPPanel (کد %d): %s", ippanelErr.Code, errorMessage)
			}
		}
		return "", fmt.Errorf("خطا در ارسال پیامک: %w", err)
	}

	// بررسی اینکه messageID معتبر باشد
	if messageID == 0 {
		return "", fmt.Errorf("خطا در ارسال پیامک: messageID صفر برگردانده شد")
	}

	return fmt.Sprintf("%d", messageID), nil
}

// sendPatternWithKavenegar ارسال پترن با کاوه‌نگار
func (s *SMSService) sendPatternWithKavenegar(pattern models.SMSPattern, gateway models.SMSGateway, recipient string, patternValues map[string]string) (string, error) {
	log.Printf("🔍 Kavenegar Pattern: PatternCode='%s', Variables=%v", pattern.PatternCode, patternValues)

	// اگر متغیر date وجود دارد، ساعت فعلی را به آن اضافه کن
	if _, exists := patternValues["date"]; exists {
		now := time.Now()
		timeStr := now.Format("15:04")
		patternValues["date"] = patternValues["date"] + " " + timeStr
	}

	log.Printf("🔍 Gateway Info: ID=%d, Name='%s', Type='%s'", gateway.ID, gateway.Name, gateway.Type)

	// TODO: پیاده‌سازی ارسال پترن با کاوه‌نگار
	return "", fmt.Errorf("ارسال پترن با کاوه‌نگار هنوز پیاده‌سازی نشده است")
}

// sendPatternWithMeliPayamak ارسال پترن با ملی پیامک
func (s *SMSService) sendPatternWithMeliPayamak(pattern models.SMSPattern, gateway models.SMSGateway, recipient string, patternValues map[string]string) (string, error) {
	log.Printf("🔍 MeliPayamak Pattern: PatternCode='%s', Variables=%v", pattern.PatternCode, patternValues)

	// اگر متغیر date وجود دارد، ساعت فعلی را به آن اضافه کن
	if _, exists := patternValues["date"]; exists {
		now := time.Now()
		timeStr := now.Format("15:04")
		patternValues["date"] = patternValues["date"] + " " + timeStr
	}

	log.Printf("🔍 Gateway Info: ID=%d, Name='%s', Type='%s'", gateway.ID, gateway.Name, gateway.Type)

	// TODO: پیاده‌سازی ارسال پترن با ملی پیامک
	return "", fmt.Errorf("ارسال پترن با ملی پیامک هنوز پیاده‌سازی نشده است")
}

// SendPatternByScopeAndFeature ارسال پیامک بر اساس scope و feature
func (s *SMSService) SendPatternByScopeAndFeature(scope, feature, recipient string, patternValues map[string]string) (string, error) {
	log.Printf("🔍 SendPatternByScopeAndFeature: Scope='%s', Feature='%s', Recipient='%s'", scope, feature, recipient)
	log.Printf("🔍 متغیرهای پترن: %v", patternValues)

	// دریافت درگاه‌های فعال
	var gateways []models.SMSGateway
	if err := s.db.Where("is_active = ? AND pattern_based = ?", true, true).Order("priority ASC").Find(&gateways).Error; err != nil {
		return "", fmt.Errorf("خطا در دریافت درگاه‌های فعال: %w", err)
	}

	if len(gateways) == 0 {
		return "", fmt.Errorf("هیچ درگاه فعالی با قابلیت پترن یافت نشد")
	}

	// حلقه روی درگاه‌ها
	for _, gateway := range gateways {
		log.Printf("🔍 بررسی درگاه: ID=%d, Type=%s, PatternBased=%v", gateway.ID, gateway.Type, gateway.PatternBased)

		// بررسی اینکه آیا این درگاه از پترن پشتیبانی می‌کند
		if !gateway.PatternBased {
			log.Printf("❌ درگاه %d از پترن پشتیبانی نمی‌کند", gateway.ID)
			continue
		}

		// بررسی وجود پترن در این درگاه بر اساس scope و feature
		var pattern models.SMSPattern
		if err := s.db.Where("gateway_id = ? AND scope = ? AND feature = ? AND status = ?", gateway.ID, scope, feature, "active").First(&pattern).Error; err != nil {
			log.Printf("❌ پترن با Scope='%s' و Feature='%s' برای درگاه %d یافت نشد: %v", scope, feature, gateway.ID, err)
			continue
		}

		log.Printf("✅ پترن یافت شد: ID=%d, PatternCode=%s, GatewayID=%d, Scope=%s, Feature=%s", pattern.ID, pattern.PatternCode, pattern.GatewayID, pattern.Scope, pattern.Feature)

		var messageID string
		var err error

		switch strings.ToLower(gateway.Type) {
		case "ippanel":
			messageID, err = s.sendPatternWithIPPanel(pattern, gateway, recipient, patternValues)
		case "kavenegar":
			messageID, err = s.sendPatternWithKavenegar(pattern, gateway, recipient, patternValues)
		case "meli_payamak":
			messageID, err = s.sendPatternWithMeliPayamak(pattern, gateway, recipient, patternValues)
		case "novin":
			messageID, err = s.sendPatternWithFarazSMS(pattern, gateway, recipient, patternValues)
		default:
			err = fmt.Errorf("درگاه پیامک %s پشتیبانی نمی‌شود", gateway.Type)
		}

		if err == nil {
			// افزایش شمارنده استفاده
			s.db.Model(&pattern).Update("usage_count", pattern.UsageCount+1)
			return messageID, nil
		}

		// در صورت خطا، درگاه را خراب علامت بزنیم و ادامه دهیم
		log.Printf("خطا در ارسال پترن با درگاه %s (ID=%d): %v", gateway.Type, gateway.ID, err)
		s.markGatewayFailed(gateway.ID)
	}

	return "", fmt.Errorf("ارسال پترن با تمام درگاه‌ها ناموفق بود")
}

// sendPatternWithFarazSMS ارسال پترن با فراز اس‌ام‌اس
func (s *SMSService) sendPatternWithFarazSMS(_ models.SMSPattern, _ models.SMSGateway, _ string, _ map[string]string) (string, error) {
	// TODO: پیاده‌سازی ارسال پترن با فراز اس‌ام‌اس
	return "", fmt.Errorf("ارسال پترن با فراز اس‌ام‌اس هنوز پیاده‌سازی نشده است")
}

// SendDirectSMS ارسال مستقیم پیامک بدون پترن
func (s *SMSService) SendDirectSMS(gatewayID uint, recipient, messageTemplate string, variables map[string]string) (string, error) {
	log.Printf("🔍 SendDirectSMS: GatewayID=%d, Recipient='%s'", gatewayID, recipient)
	log.Printf("🔍 متغیرهای پیام: %v", variables)

	// دریافت درگاه
	var gateway models.SMSGateway
	if err := s.db.First(&gateway, gatewayID).Error; err != nil {
		return "", fmt.Errorf("خطا در دریافت درگاه: %w", err)
	}

	// بررسی فعال بودن درگاه
	if !gateway.IsActive {
		return "", fmt.Errorf("درگاه غیرفعال است")
	}

	// جایگزینی متغیرها در متن پیام
	finalMessage := messageTemplate
	log.Printf("🔍 متن پیام اولیه: '%s'", finalMessage)

	for key, value := range variables {
		placeholder := "%" + key + "%"
		oldMessage := finalMessage
		finalMessage = strings.ReplaceAll(finalMessage, placeholder, value)
		if oldMessage != finalMessage {
			log.Printf("🔍 جایگزینی: '%s' -> '%s'", placeholder, value)
		}
	}

	log.Printf("🔍 پیام نهایی: '%s'", finalMessage)

	var messageID string
	var err error

	switch strings.ToLower(gateway.Type) {
	case "ippanel":
		messageID, err = s.sendDirectWithIPPanel(gateway, recipient, finalMessage)
	case "kavenegar":
		messageID, err = s.sendDirectWithKavenegar(gateway, recipient, finalMessage)
	case "meli_payamak":
		messageID, err = s.sendDirectWithMeliPayamak(gateway, recipient, finalMessage)
	case "novin":
		messageID, err = s.sendDirectWithFarazSMS(gateway, recipient, finalMessage)
	default:
		err = fmt.Errorf("درگاه پیامک %s پشتیبانی نمی‌شود", gateway.Type)
	}

	if err != nil {
		log.Printf("خطا در ارسال مستقیم با درگاه %s (ID=%d): %v", gateway.Type, gateway.ID, err)
		s.markGatewayFailed(gateway.ID)
		return "", err
	}

	return messageID, nil
}

// sendDirectWithIPPanel ارسال مستقیم با IPPanel
func (s *SMSService) sendDirectWithIPPanel(gateway models.SMSGateway, recipient, message string) (string, error) {
	log.Printf("📤 آماده ارسال مستقیم با IPPanel...")
	log.Printf("📤 پارامترها: sender='%s', mobile='%s', message='%s'", gateway.SenderNumber, recipient, message)

	// ایجاد کلاینت IPPanel
	client := ippanel.New(gateway.ApiKey)

	// فرمت‌بندی شماره موبایل برای IPPanel (فرمت بین‌المللی)
	formattedMobile := recipient
	if !strings.HasPrefix(formattedMobile, "+") {
		if strings.HasPrefix(formattedMobile, "0") {
			// تبدیل 09203214155 به +989203214155
			formattedMobile = "+98" + strings.TrimPrefix(formattedMobile, "0")
		} else if strings.HasPrefix(formattedMobile, "98") {
			// تبدیل 989203214155 به +989203214155
			formattedMobile = "+" + formattedMobile
		}
	}

	log.Printf("🔍 شماره موبایل فرمت شده: '%s'", formattedMobile)

	// ارسال پیامک مستقیم
	messageID, err := client.Send(
		gateway.SenderNumber,      // شماره فرستنده
		[]string{formattedMobile}, // لیست شماره‌های گیرنده
		message,                   // متن پیام
		"normal",                  // نوع پیام (normal برای پیامک معمولی)
	)

	if err != nil {
		log.Printf("❌ خطا در Send: %v", err)
		return "", fmt.Errorf("خطای IPPanel: %w", err)
	}

	log.Printf("✅ پیامک مستقیم با موفقیت ارسال شد: messageID=%d", messageID)
	return fmt.Sprintf("%d", messageID), nil
}

// sendDirectWithKavenegar ارسال مستقیم با کاوه‌نگار
func (s *SMSService) sendDirectWithKavenegar(_ models.SMSGateway, _, _ string) (string, error) {
	// TODO: پیاده‌سازی ارسال مستقیم با کاوه‌نگار
	return "", fmt.Errorf("ارسال مستقیم با کاوه‌نگار هنوز پیاده‌سازی نشده است")
}

// sendDirectWithMeliPayamak ارسال مستقیم با ملی پیامک
func (s *SMSService) sendDirectWithMeliPayamak(_ models.SMSGateway, _, _ string) (string, error) {
	// TODO: پیاده‌سازی ارسال مستقیم با ملی پیامک
	return "", fmt.Errorf("ارسال مستقیم با ملی پیامک هنوز پیاده‌سازی نشده است")
}

// sendDirectWithFarazSMS ارسال مستقیم با فراز اس‌ام‌اس
func (s *SMSService) sendDirectWithFarazSMS(_ models.SMSGateway, _, _ string) (string, error) {
	// TODO: پیاده‌سازی ارسال مستقیم با فراز اس‌ام‌اس
	return "", fmt.Errorf("ارسال مستقیم با فراز اس‌ام‌اس هنوز پیاده‌سازی نشده است")
}
