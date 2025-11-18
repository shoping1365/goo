package sms

import (
	"fmt"
	"my-go-backend/internal/models"

	ippanel "github.com/ippanel/go-rest-sdk/v2"
)

// IPPanelGateway پیاده‌سازی درگاه IPPanel
// این struct رکورد دیتابیس را نگه می‌دارد تا بتوان به تنظیمات آن دسترسی داشت

type IPPanelGateway struct {
	gateway models.SMSGateway
}

// NewIPPanelGateway سازنده درگاه
func NewIPPanelGateway(gw models.SMSGateway) *IPPanelGateway {
	return &IPPanelGateway{gateway: gw}
}

// Record برگرداندن رکورد درگاه
func (g *IPPanelGateway) Record() models.SMSGateway {
	return g.gateway
}

// HealthCheck بررسی سلامت درگاه با دریافت موجودی
func (g *IPPanelGateway) HealthCheck() error {
	client := ippanel.New(g.gateway.ApiKey)
	_, err := client.GetCredit()
	return err
}

// Send ارسال پیامک (پترنی یا عادی بسته به تنظیمات)
func (g *IPPanelGateway) Send(pattern models.SMSPattern, req models.SMSSendRequest) (*models.SMSSendResponse, error) {
	client := ippanel.New(g.gateway.ApiKey)

	// آماده‌سازی متغیرهای پترن
	patternValues := make(map[string]string)
	for k, v := range req.Variables {
		patternValues[k] = fmt.Sprintf("%v", v)
	}

	messageID, err := client.SendPattern(
		pattern.PatternCode,
		g.gateway.SenderNumber,
		req.Mobile,
		patternValues,
	)
	if err != nil {
		return nil, err
	}

	response := &models.SMSSendResponse{
		Success: true,
		Message: "پیامک با موفقیت ارسال شد",
	}
	response.Data.MessageID = fmt.Sprintf("%d", messageID)
	return response, nil
}
