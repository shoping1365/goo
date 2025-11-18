package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"my-go-backend/internal/models"
)

// FarazSMSGateway پیاده‌سازی ساده برای درگاه فراز اس‌ام‌اس
// توجه: برای استفاده در تولید باید با مستندات رسمی فراز یکپارچه شود.

type FarazSMSGateway struct {
	gateway models.SMSGateway
}

func NewFarazSMSGateway(gw models.SMSGateway) *FarazSMSGateway {
	return &FarazSMSGateway{gateway: gw}
}

func (g *FarazSMSGateway) Record() models.SMSGateway { return g.gateway }

// HealthCheck با بررسی وضعیت موجودی (نمونه فرضی)
func (g *FarazSMSGateway) HealthCheck() error {
	url := fmt.Sprintf("https://api.sms.farazpardazan.com/v1/credit?api_key=%s", g.gateway.ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status %d", resp.StatusCode)
	}
	return nil
}

// Send ارسال پیامک عادی؛ فراز برای پترن‌ها نیز endpoint مجزا دارد که اینجا ساده شده است.
func (g *FarazSMSGateway) Send(pattern models.SMSPattern, req models.SMSSendRequest) (*models.SMSSendResponse, error) {
	url := "https://api.sms.farazpardazan.com/v1/send"
	body := map[string]interface{}{
		"api_key": g.gateway.ApiKey,
		"secret":  g.gateway.Password,
		"from":    g.gateway.SenderNumber,
		"to":      req.Mobile,
		"text":    pattern.MessageTemplate, // باید جایگذاری متغیرها انجام شود
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	return &models.SMSSendResponse{Success: true, Message: "ارسال با فراز اس‌ام‌اس انجام شد"}, nil
}
