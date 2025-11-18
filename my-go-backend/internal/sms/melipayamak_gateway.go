package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"my-go-backend/internal/models"
)

// MeliPayamakGateway پیاده‌سازی درگاه ملی پیامک (REST)

type MeliPayamakGateway struct {
	gateway models.SMSGateway
}

// NewMeliPayamakGateway سازنده درگاه
func NewMeliPayamakGateway(gw models.SMSGateway) *MeliPayamakGateway {
	return &MeliPayamakGateway{gateway: gw}
}

func (g *MeliPayamakGateway) Record() models.SMSGateway { return g.gateway }

// HealthCheck ساده – درخواست موجودی
func (g *MeliPayamakGateway) HealthCheck() error {
	url := "https://rest.payamak-panel.com/api/credit/"
	// در نسخه نمایشی فقط 200 را چک می‌کنیم
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

// Send ارسال پیامک با ملی پیامک
func (g *MeliPayamakGateway) Send(pattern models.SMSPattern, req models.SMSSendRequest) (*models.SMSSendResponse, error) {
	url := "https://rest.payamak-panel.com/api/SendSMS/SendSMS"

	bodyMap := map[string]interface{}{
		"api_key": g.gateway.ApiKey,
		"to":      req.Mobile,
		"from":    g.gateway.SenderNumber,
		"text":    pattern.MessageTemplate, // در نسخه پترن، باید جایگذاری متغیرها انجام شود
	}
	b, err := json.Marshal(bodyMap)
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
	return &models.SMSSendResponse{Success: true, Message: "ارسال با ملی پیامک انجام شد"}, nil
}
