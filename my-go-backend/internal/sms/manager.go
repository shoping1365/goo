package sms

import (
	"fmt"
	"log"
	"sync"
	"time"

	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

var (
	globalManager *GatewayManager
	managerOnce   sync.Once
)

// GatewayManager وظیفه مدیریت، سوییچ و Failover بین درگاه‌ها را بر عهده دارد

type GatewayManager struct {
	db             *gorm.DB
	gateways       []SmsGateway
	failedGateways map[uint]time.Time // کلید: ID درگاه
	mu             sync.RWMutex
}

// NewGatewayManager لود اولیه درگاه‌های فعال از دیتابیس
func NewGatewayManager(db *gorm.DB) (*GatewayManager, error) {
	managerOnce.Do(func() {
		manager := &GatewayManager{
			db:             db,
			failedGateways: make(map[uint]time.Time),
		}
		if err := manager.loadGateways(); err != nil {
			log.Printf("خطا در بارگذاری درگاه‌ها: %v", err)
			return
		}

		go manager.startHealthChecker()
		globalManager = manager
	})
	return globalManager, nil
}

// cleanupDuplicatePatterns پاک کردن پترن‌های تکراری
func (m *GatewayManager) cleanupDuplicatePatterns() {
	systemPatterns := []string{"admin_failover", "admin_test", "otp_login", "admin_order", "low_stock", "security_issue"}

	// دریافت همه درگاه‌های فعال از دیتابیس
	var gateways []models.SMSGateway
	if err := m.db.Where("is_active = ?", true).Find(&gateways).Error; err != nil {
		log.Printf("خطا در دریافت درگاه‌ها برای پاک‌سازی: %v", err)
		return
	}

	for _, patternCode := range systemPatterns {
		// یافتن همه پترن‌های با این کد برای هر درگاه
		for _, gateway := range gateways {
			var patterns []models.SMSPattern
			m.db.Where("pattern_code = ? AND gateway_id = ?", patternCode, gateway.ID).Find(&patterns)

			if len(patterns) > 1 {
				// نگه داشتن اولین پترن و حذف بقیه
				for i := 1; i < len(patterns); i++ {
					m.db.Delete(&patterns[i])
					log.Printf("پترن تکراری حذف شد: %s (ID: %d) برای درگاه %d", patternCode, patterns[i].ID, gateway.ID)
				}
			}
		}
	}

	// همچنین پاک کردن پترن‌هایی که نام آن‌ها تکراری است
	systemNames := []string{"اعلان خطای سیستم", "تست پیامک", "کد تأیید ورود", "اعلان سفارشات به مدیر", "کمبود موجودی محصول", "مشکلات امنیتی"}

	for _, name := range systemNames {
		for _, gateway := range gateways {
			var patterns []models.SMSPattern
			m.db.Where("name = ? AND gateway_id = ?", name, gateway.ID).Find(&patterns)

			if len(patterns) > 1 {
				// نگه داشتن اولین پترن و حذف بقیه
				for i := 1; i < len(patterns); i++ {
					m.db.Delete(&patterns[i])
					log.Printf("پترن تکراری حذف شد (بر اساس نام): %s (ID: %d) برای درگاه %d", name, patterns[i].ID, gateway.ID)
				}
			}
		}
	}
}

// loadGateways خواندن درگاه‌های فعال مرتب شده بر اساس اولویت
func (m *GatewayManager) loadGateways() error {
	var records []models.SMSGateway
	if err := m.db.Where("is_active = ?", true).Order("priority ASC").Find(&records).Error; err != nil {
		return err
	}

	var list []SmsGateway
	for _, rec := range records {
		switch rec.Type {
		case "ippanel":
			list = append(list, NewIPPanelGateway(rec))
		case "kavenegar":
			// list = append(list, NewKavenegarGateway(rec)) // TODO: پیاده‌سازی Kavenegar
		case "meli_payamak":
			list = append(list, NewMeliPayamakGateway(rec))
		case "farazsms":
			list = append(list, NewFarazSMSGateway(rec))
			// سایر درگاه‌ها
		default:
			continue
		}
	}
	m.gateways = list
	return nil
}

// CleanupDuplicatesManually پاک‌سازی دستی پترن‌های تکراری (برای API endpoint)
func CleanupDuplicatesManually(db *gorm.DB) error {
	systemPatterns := []string{"admin_failover", "admin_test", "otp_login", "admin_order", "low_stock", "security_issue"}
	systemNames := []string{"اعلان خطای سیستم", "تست پیامک", "کد تأیید ورود", "اعلان سفارشات به مدیر", "کمبود موجودی محصول", "مشکلات امنیتی"}

	// دریافت همه درگاه‌های فعال از دیتابیس
	var gateways []models.SMSGateway
	if err := db.Where("is_active = ?", true).Find(&gateways).Error; err != nil {
		return fmt.Errorf("خطا در دریافت درگاه‌ها: %w", err)
	}

	deletedCount := 0

	// پاک کردن بر اساس pattern_code
	for _, patternCode := range systemPatterns {
		for _, gateway := range gateways {
			var patterns []models.SMSPattern
			db.Where("pattern_code = ? AND gateway_id = ?", patternCode, gateway.ID).Find(&patterns)

			if len(patterns) > 1 {
				for i := 1; i < len(patterns); i++ {
					db.Delete(&patterns[i])
					deletedCount++
					log.Printf("پترن تکراری حذف شد: %s (ID: %d) برای درگاه %d", patternCode, patterns[i].ID, gateway.ID)
				}
			}
		}
	}

	// پاک کردن بر اساس نام
	for _, name := range systemNames {
		for _, gateway := range gateways {
			var patterns []models.SMSPattern
			db.Where("name = ? AND gateway_id = ?", name, gateway.ID).Find(&patterns)

			if len(patterns) > 1 {
				for i := 1; i < len(patterns); i++ {
					db.Delete(&patterns[i])
					deletedCount++
					log.Printf("پترن تکراری حذف شد (بر اساس نام): %s (ID: %d) برای درگاه %d", name, patterns[i].ID, gateway.ID)
				}
			}
		}
	}

	log.Printf("تعداد %d پترن تکراری حذف شد", deletedCount)
	return nil
}

// SendSMS تلاش برای ارسال پیامک به ترتیب اولویت
func (m *GatewayManager) SendSMS(req models.SMSSendRequest) (*models.SMSSendResponse, error) {
	for _, gw := range m.gateways {
		// بررسی خرابی
		rec := gw.Record()
		m.mu.RLock()
		_, failed := m.failedGateways[rec.ID]
		m.mu.RUnlock()
		if failed {
			continue
		}

		// دریافت پترن منطبق با این درگاه
		var pattern models.SMSPattern
		if err := m.db.Where("pattern_code = ? AND gateway_id = ? AND status = ?", req.PatternCode, rec.ID, "active").First(&pattern).Error; err != nil {
			continue
		}

		resp, err := gw.Send(pattern, req)
		if err == nil {
			return resp, nil
		}

		log.Printf("خطا در درگاه %s: %v", rec.Name, err)
		m.markFailed(rec.ID)
	}
	return nil, fmt.Errorf("ارسال پیامک با همه درگاه‌ها ناموفق بود")
}

// markFailed علامت‌گذاری درگاه خراب
func (m *GatewayManager) markFailed(id uint) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.failedGateways[id] = time.Now()
}

// startHealthChecker مانیتور خرابی درگاه‌ها
func (m *GatewayManager) startHealthChecker() {
	ticker := time.NewTicker(30 * time.Minute)
	for range ticker.C {
		m.checkFailedGateways()
	}
}

// checkFailedGateways تلاش برای احیای درگاه‌های خراب
func (m *GatewayManager) checkFailedGateways() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for id, t := range m.failedGateways {
		if time.Since(t) < 30*time.Minute {
			continue
		}

		var rec models.SMSGateway
		if err := m.db.First(&rec, id).Error; err != nil {
			delete(m.failedGateways, id)
			continue
		}

		var gw SmsGateway
		switch rec.Type {
		case "ippanel":
			gw = NewIPPanelGateway(rec)
		default:
			continue
		}

		if gw.HealthCheck() == nil {
			delete(m.failedGateways, id)
			log.Printf("درگاه %s مجدداً فعال شد", rec.Name)
		} else {
			m.failedGateways[id] = time.Now()
		}
	}
}
