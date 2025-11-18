package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/datatypes"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/repository"
)

// FraudService سرویس مرکزی ارزیابی تقلب
// تمام منطق امتیازدهی، اعمال قوانین و ثبت وقایع در اینجا انجام می‌شود.
type FraudService struct {
	uowFactory unitofwork.UnitOfWorkFactory
	salt       string
	ipintel    IPIntelService
}

func NewFraudService(uowFactory unitofwork.UnitOfWorkFactory, salt string) *FraudService {
	return &FraudService{uowFactory: uowFactory, salt: salt, ipintel: NewIPIntelServiceFromEnv()}
}

// EvaluateOrder ورودی‌های ارزیابی: شناسه سفارش و ویژگی‌های زمینه‌ای (ip، device، geo، ageOfAccount، ...)
type EvaluateOrderInput struct {
	OrderID                    uint
	UserID                     *uint
	Amount                     float64
	PaymentMethod              string
	UserRegisteredAt           *time.Time
	IP                         string
	DeviceID                   string
	ShippingCity               string
	ShippingProvince           string
	GeoIPCity                  string
	GeoIPCountry               string
	IsDatacenterIP             bool
	RecentFailedPaymentsFromIP int
}

type EvaluateOrderResult struct {
	CaseID    uint
	RiskScore int
	RiskLevel string
	Status    string
}

// EvaluateOrder
// این متد ورودی‌های سفارش را دریافت می‌کند و با اعمال قوانین تشخیص تقلب،
// امتیاز ریسک (۰..۱۰۰)، سطح ریسک (low|medium|high) و وضعیت پیشنهادی
// (approved|review|blocked) را محاسبه و در جداول مربوطه ثبت می‌کند.
// تمامی عملیات دیتابیس در یک UnitOfWork انجام می‌شود تا اتمیسیتی تضمین گردد.
func (s *FraudService) EvaluateOrder(ctx context.Context, in EvaluateOrderInput) (*EvaluateOrderResult, error) {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, err
	}
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	rules, err := uow.FraudRuleRepository().ListActive(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot load rules: %w", err)
	}

	// base score
	score := 0
	var events []repository.FraudEvent

	// geoip_mismatch_city
	if in.GeoIPCity != "" && in.ShippingCity != "" && !strings.EqualFold(in.GeoIPCity, in.ShippingCity) {
		if ru := findRule(rules, "geoip_mismatch_city"); ru != nil {
			impact := ru.Weight
			score += impact
			events = append(events, repository.FraudEvent{EventType: ru.Key, Weight: ru.Weight, Impact: impact, Details: toJSON(map[string]any{"geo": in.GeoIPCity, "ship": in.ShippingCity})})
		}
	}

	// new_account_high_amount
	if in.UserRegisteredAt != nil {
		if days := int(time.Since(*in.UserRegisteredAt).Hours() / 24); days < 7 && in.Amount > 0 {
			if ru := findRule(rules, "new_account_high_amount"); ru != nil {
				impact := ru.Weight
				score += impact
				events = append(events, repository.FraudEvent{EventType: ru.Key, Weight: ru.Weight, Impact: impact, Details: toJSON(map[string]any{"days": days, "amount": in.Amount})})
			}
		}
	}

	// rapid_failed_payments
	if in.RecentFailedPaymentsFromIP >= 3 {
		if ru := findRule(rules, "rapid_failed_payments"); ru != nil {
			impact := ru.Weight
			score += impact
			events = append(events, repository.FraudEvent{EventType: ru.Key, Weight: ru.Weight, Impact: impact, Details: toJSON(map[string]any{"failed": in.RecentFailedPaymentsFromIP})})
		}
	}

	// datacenter_or_proxy_ip (با سرویس IPIntel یا ورودی مستقیم)
	dc := in.IsDatacenterIP
	if !dc && in.IP != "" && s.ipintel != nil {
		if out, _ := s.ipintel.Lookup(in.IP); out != nil && out.IsDatacenter {
			dc = true
		}
	}
	if dc {
		if ru := findRule(rules, "datacenter_or_proxy_ip"); ru != nil {
			impact := ru.Weight
			score += impact
			events = append(events, repository.FraudEvent{EventType: ru.Key, Weight: ru.Weight, Impact: impact})
		}
	}

	// whitelist_hit
	if in.IP != "" {
		if ok, _ := uow.FraudListRepository().IsWhitelisted(ctx, "ip_whitelist", s.hash(in.IP)); ok {
			if ru := findRule(rules, "whitelist_hit"); ru != nil {
				impact := ru.Weight
				score += impact
				events = append(events, repository.FraudEvent{EventType: ru.Key, Weight: ru.Weight, Impact: impact})
			}
		}
	}

	// clamp and level
	if score < -40 {
		score = -40
	}
	if score > 100 {
		score = 100
	}
	level := "low"
	status := "approved"
	if score >= 70 {
		level = "high"
		status = "blocked"
	}
	if score >= 40 && score < 70 {
		level = "medium"
		status = "review"
	}

	// upsert case بر اساس order_id: اگر قبلاً برای این سفارش پرونده‌ای ثبت شده است، به‌روزرسانی و رویدادهای جدید اضافه شوند
	var userID *uint = in.UserID
	existing, _ := uow.FraudCaseRepository().GetByOrderID(ctx, in.OrderID)
	if existing != nil && existing.ID > 0 {
		// ثبت تغییر امتیاز
		_ = uow.GetTx().Create(&repository.FraudScoreLog{CaseID: existing.ID, ScoreBefore: existing.RiskScore, ScoreAfter: score, Reason: toJSON(map[string]any{"re_eval": true})}).Error
		existing.RiskScore = score
		existing.RiskLevel = level
		existing.Status = status
		existing.Summary = "auto-re-evaluated"
		existing.Metadata = toJSON(map[string]any{"amount": in.Amount, "payment_method": in.PaymentMethod})
		if err := uow.FraudCaseRepository().Update(ctx, existing); err != nil {
			return nil, err
		}
		// پیوست رویدادها
		for i := range events {
			events[i].CaseID = existing.ID
		}
		if err := uow.FraudEventRepository().CreateMany(ctx, events); err != nil {
			return nil, err
		}
		if err := uow.Commit(); err != nil {
			return nil, err
		}
		committed = true
		return &EvaluateOrderResult{CaseID: existing.ID, RiskScore: score, RiskLevel: level, Status: status}, nil
	}
	fc := &repository.FraudCase{OrderID: &in.OrderID, UserID: userID, RiskScore: score, RiskLevel: level, Status: status, Summary: "auto-evaluated", Metadata: toJSON(map[string]any{"amount": in.Amount, "payment_method": in.PaymentMethod})}
	if err := uow.FraudCaseRepository().Create(ctx, fc); err != nil {
		return nil, err
	}

	// attach foreign keys and persist events
	for i := range events {
		events[i].CaseID = fc.ID
	}
	if err := uow.FraudEventRepository().CreateMany(ctx, events); err != nil {
		return nil, err
	}

	if err := uow.Commit(); err != nil {
		return nil, err
	}
	committed = true
	return &EvaluateOrderResult{CaseID: fc.ID, RiskScore: score, RiskLevel: level, Status: status}, nil
}

// Decide
// این متد تصمیم ادمین/سیستم روی پرونده تقلب را ثبت می‌کند و وضعیت کیس را
// به یکی از حالت‌های approved|blocked|review تغییر می‌دهد. اکشن در لاگ اکشن ثبت می‌شود.
func (s *FraudService) Decide(ctx context.Context, caseID uint, action string, actorID *uint, note string) error {
	uow := s.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return err
	}
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	fc, err := uow.FraudCaseRepository().GetByID(ctx, caseID)
	if err != nil {
		return err
	}
	switch action {
	case "approve":
		fc.Status = "approved"
	case "block":
		fc.Status = "blocked"
	case "review":
		fc.Status = "review"
	default:
		return fmt.Errorf("invalid action")
	}
	if err := uow.FraudCaseRepository().Update(ctx, fc); err != nil {
		return err
	}
	if err := uow.FraudActionLogRepository().Create(ctx, &repository.FraudActionLog{CaseID: caseID, Action: action, ActorID: actorID, Note: note}); err != nil {
		return err
	}
	if err := uow.Commit(); err != nil {
		return err
	}
	committed = true
	return nil
}

func (s *FraudService) hash(v string) string {
	h := sha256.Sum256([]byte(s.salt + "::" + v))
	return hex.EncodeToString(h[:])
}

// Hash بیرون‌داد نسخه export برای استفاده هندلر (عدم تکرار salt)
func (s *FraudService) Hash(v string) string { return s.hash(v) }

func findRule(rules []repository.FraudRule, key string) *repository.FraudRule {
	for i := range rules {
		if rules[i].Key == key && rules[i].Enabled {
			return &rules[i]
		}
	}
	return nil
}

func toJSON(m map[string]any) datatypes.JSON {
	if m == nil {
		return datatypes.JSON([]byte("{}"))
	}
	b, _ := json.Marshal(m)
	return datatypes.JSON(b)
}
