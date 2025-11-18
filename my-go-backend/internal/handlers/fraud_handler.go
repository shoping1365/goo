package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

// FraudHandler هندلر روت‌های ادمین تقلب
type FraudHandler struct {
	service *services.FraudService
	uowf    unitofwork.UnitOfWorkFactory
}

func NewFraudHandler(service *services.FraudService, uowf unitofwork.UnitOfWorkFactory) *FraudHandler {
	return &FraudHandler{service: service, uowf: uowf}
}

// GET /api/admin/fraud/cases
// دریافت لیست پرونده‌های تقلب به صورت صفحه‌بندی شده با اطلاعات سفارش/کاربر
func (h *FraudHandler) ListCases(c *gin.Context) {
	uow := h.uowf.Create()
	ctx := c.Request.Context()
	q := c.Request.URL.Query()
	risk := q.Get("risk")
	status := q.Get("status")
	page, _ := strconv.Atoi(q.Get("page"))
	size, _ := strconv.Atoi(q.Get("pageSize"))

	filter := struct {
		RiskLevel *string
		Status    *string
		Page      int
		PageSize  int
	}{Page: page, PageSize: size}
	if risk != "" {
		filter.RiskLevel = &risk
	}
	if status != "" {
		filter.Status = &status
	}

	repo := uow.FraudCaseRepository()
	cases, total, err := repo.ListEnriched(ctx, repository.FraudCaseFilter{
		RiskLevel: filter.RiskLevel,
		Status:    filter.Status,
		Page:      filter.Page,
		PageSize:  filter.PageSize,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا در دریافت لیست"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": cases, "total": total})
}

// GET /api/admin/fraud/cases/:id
// دریافت جزئیات یک پرونده: پرونده + رویدادها + اکشن‌های ادمین
func (h *FraudHandler) GetCaseDetail(c *gin.Context) {
	idStr := c.Param("id")
	cid, _ := strconv.Atoi(idStr)
	uow := h.uowf.Create()
	detail, err := uow.FraudCaseRepository().GetDetail(c, uint(cid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "پرونده یافت نشد"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": detail})
}

// POST /api/fraud/evaluate/order/:orderId
// ارزیابی یک سفارش بر اساس سیگنال‌های ورودی و ثبت پرونده تقلب
func (h *FraudHandler) EvaluateOrder(c *gin.Context) {
	orderIDStr := c.Param("orderId")
	oid, _ := strconv.Atoi(orderIDStr)

	var body struct {
		UserID                     *uint      `json:"user_id"`
		Amount                     float64    `json:"amount"`
		PaymentMethod              string     `json:"payment_method"`
		UserRegisteredAt           *time.Time `json:"user_registered_at"`
		IP                         string     `json:"ip"`
		DeviceID                   string     `json:"device_id"`
		ShippingCity               string     `json:"shipping_city"`
		ShippingProvince           string     `json:"shipping_province"`
		GeoIPCity                  string     `json:"geoip_city"`
		GeoIPCountry               string     `json:"geoip_country"`
		IsDatacenterIP             bool       `json:"is_datacenter_ip"`
		RecentFailedPaymentsFromIP int        `json:"recent_failed_payments_from_ip"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "ورودی نامعتبر"})
		return
	}

	res, err := h.service.EvaluateOrder(c, services.EvaluateOrderInput{
		OrderID: uint(oid), UserID: body.UserID, Amount: body.Amount, PaymentMethod: body.PaymentMethod,
		UserRegisteredAt: body.UserRegisteredAt, IP: body.IP, DeviceID: body.DeviceID, ShippingCity: body.ShippingCity,
		ShippingProvince: body.ShippingProvince, GeoIPCity: body.GeoIPCity, GeoIPCountry: body.GeoIPCountry,
		IsDatacenterIP: body.IsDatacenterIP, RecentFailedPaymentsFromIP: body.RecentFailedPaymentsFromIP,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا در ارزیابی"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": res})
}

// POST /api/admin/fraud/cases/:id/decision
// ثبت تصمیم ادمین روی یک پرونده تقلب
func (h *FraudHandler) Decide(c *gin.Context) {
	idStr := c.Param("id")
	cid, _ := strconv.Atoi(idStr)
	var body struct {
		Action string `json:"action"`
		Note   string `json:"note"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	var actorID *uint
	if v, ok := c.Get("user_id"); ok {
		if vv, ok2 := v.(uint); ok2 {
			actorID = &vv
		}
	}
	if err := h.service.Decide(c, uint(cid), body.Action, actorID, body.Note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// POST /api/admin/fraud/re-evaluate-open
// بازارزیابی همه پرونده‌های باز (review) بعد از تغییر قوانین
func (h *FraudHandler) ReEvaluateOpenCases(c *gin.Context) {
	uow := h.uowf.Create()
	ctx := c.Request.Context()
	// دریافت لیست کیس‌های باز برای بازارزیابی
	status := "review"
	cases, _, err := uow.FraudCaseRepository().List(ctx, repository.FraudCaseFilter{Status: &status, Page: 1, PageSize: 1000})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا در دریافت پرونده‌ها"})
		return
	}
	for _, fc := range cases {
		orderID := uint(0)
		if fc.OrderID != nil {
			orderID = *fc.OrderID
		}
		// اطلاعات ورودی حداقلی؛ UI یا کران‌جاب می‌تواند جزئیات بیشتری تزریق کند
		_, _ = h.service.EvaluateOrder(ctx, services.EvaluateOrderInput{OrderID: orderID, UserID: fc.UserID, Amount: 0, PaymentMethod: "", IP: ""})
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// GET /api/admin/fraud/rules
// دریافت لیست قوانین فعال/غیرفعال برای مدیریت وزن‌ها و فعال‌سازی
func (h *FraudHandler) ListRules(c *gin.Context) {
	uow := h.uowf.Create()
	rules, err := uow.FraudRuleRepository().ListAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا در دریافت قوانین"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": rules})
}

// PUT /api/admin/fraud/rules/:id
// بروزرسانی وزن/فعال بودن/پارامترهای قانون
func (h *FraudHandler) UpdateRule(c *gin.Context) {
	uow := h.uowf.Create()
	idStr := c.Param("id")
	rid, _ := strconv.Atoi(idStr)
	var body struct {
		Weight      *int           `json:"weight"`
		Enabled     *bool          `json:"enabled"`
		Params      map[string]any `json:"params"`
		Description *string        `json:"description"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "ورودی نامعتبر"})
		return
	}
	rule, err := uow.FraudRuleRepository().GetByID(c, uint(rid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "قانون یافت نشد"})
		return
	}
	if body.Weight != nil {
		rule.Weight = *body.Weight
	}
	if body.Enabled != nil {
		rule.Enabled = *body.Enabled
	}
	if body.Description != nil {
		rule.Description = *body.Description
	}
	if body.Params != nil {
		b, _ := json.Marshal(body.Params)
		rule.Params = datatypes.JSON(b)
	}
	if err := uow.FraudRuleRepository().Update(c, rule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا در بروزرسانی قانون"})
		return
	}
	// بازارزیابی پرونده‌های درحال‌بررسی به صورت غیرمسدودکننده
	go func() {
		uow2 := h.uowf.Create()
		ctx := context.Background()
		status := "review"
		cases, _, err := uow2.FraudCaseRepository().List(ctx, repository.FraudCaseFilter{Status: &status, Page: 1, PageSize: 1000})
		if err != nil {
			return
		}
		for _, fc := range cases {
			orderID := uint(0)
			if fc.OrderID != nil {
				orderID = *fc.OrderID
			}
			_, _ = h.service.EvaluateOrder(ctx, services.EvaluateOrderInput{OrderID: orderID, UserID: fc.UserID})
		}
	}()
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// GET /api/admin/fraud/stats
// آمار سطح ریسک و وضعیت پرونده‌ها برای داشبورد
func (h *FraudHandler) Stats(c *gin.Context) {
	uow := h.uowf.Create()
	db := uow.GetTx()
	type KV struct {
		K string
		V int64
	}
	var byLevel []KV
	var byStatus []KV
	db.Raw("SELECT risk_level as k, COUNT(*) as v FROM fraud_cases GROUP BY risk_level").Scan(&byLevel)
	db.Raw("SELECT status as k, COUNT(*) as v FROM fraud_cases GROUP BY status").Scan(&byStatus)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"byLevel": byLevel, "byStatus": byStatus}})
}

// POST /api/admin/fraud/list/whitelist
// افزودن آیتم به لیست سفید (hash شده)
func (h *FraudHandler) UpsertWhitelist(c *gin.Context) {
	h.upsertList(c, true)
}

// POST /api/admin/fraud/list/blacklist
// افزودن آیتم به لیست سیاه (hash شده)
func (h *FraudHandler) UpsertBlacklist(c *gin.Context) {
	h.upsertList(c, false)
}

func (h *FraudHandler) upsertList(c *gin.Context, whitelist bool) {
	uow := h.uowf.Create()
	var body struct {
		Kind      string     `json:"kind"`
		Value     string     `json:"value"`
		Note      string     `json:"note"`
		ExpiresAt *time.Time `json:"expires_at"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "ورودی نامعتبر"})
		return
	}
	if body.Kind == "" || body.Value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "kind و value الزامی است"})
		return
	}
	kindKey := body.Kind
	if whitelist {
		kindKey += "_whitelist"
	} else {
		kindKey += "_blacklist"
	}
	item := &repository.FraudListItem{Kind: kindKey, ValueHash: h.service.Hash(body.Value), Note: body.Note, ExpiresAt: body.ExpiresAt}
	if err := uow.FraudListRepository().Upsert(c, item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا در ثبت"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// GET /api/admin/fraud/cases (پیشرفته)
// افزودن فیلترهای بیشتر: جستجو، بازه مبلغ و تاریخ و روش پرداخت
// توجه: فقط اگر پارامترها ارسال شده باشند فیلتر می‌شوند
func (h *FraudHandler) ListCasesAdvanced(c *gin.Context) {
	uow := h.uowf.Create()
	ctx := c.Request.Context()
	q := c.Request.URL.Query()
	risk := q.Get("risk")
	status := q.Get("status")
	search := q.Get("search")
	pay := q.Get("paymentMethod")
	minAmount := q.Get("minAmount")
	maxAmount := q.Get("maxAmount")
	page, _ := strconv.Atoi(q.Get("page"))
	size, _ := strconv.Atoi(q.Get("pageSize"))
	// پارامترهای اضافی
	from := q.Get("from")
	to := q.Get("to")
	orderID := q.Get("orderId")
	userID := q.Get("userId")

	var fromT, toT *time.Time
	if from != "" {
		if tt, err := time.Parse(time.RFC3339, from); err == nil {
			fromT = &tt
		}
	}
	if to != "" {
		if tt, err := time.Parse(time.RFC3339, to); err == nil {
			toT = &tt
		}
	}
	var oidPtr *uint
	if orderID != "" {
		if v, err := strconv.Atoi(orderID); err == nil {
			vv := uint(v)
			oidPtr = &vv
		}
	}
	var uidPtr *uint
	if userID != "" {
		if v, err := strconv.Atoi(userID); err == nil {
			vv := uint(v)
			uidPtr = &vv
		}
	}

	filter := repository.FraudCaseFilter{Page: page, PageSize: size, From: fromT, To: toT, OrderID: oidPtr, UserID: uidPtr}
	if search != "" {
		filter.Search = &search
	}
	if pay != "" {
		filter.PaymentMethod = &pay
	}
	if minAmount != "" {
		if v, err := strconv.ParseFloat(minAmount, 64); err == nil {
			filter.AmountMin = &v
		}
	}
	if maxAmount != "" {
		if v, err := strconv.ParseFloat(maxAmount, 64); err == nil {
			filter.AmountMax = &v
		}
	}
	if risk != "" {
		filter.RiskLevel = &risk
	}
	if status != "" {
		filter.Status = &status
	}

	cases, total, err := uow.FraudCaseRepository().ListEnriched(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا در دریافت لیست"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": cases, "total": total})
}
