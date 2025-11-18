package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// InstallmentHandler هندلر ماژول اقساط
// این هندلر تمام منطق مربوط به تب‌های صفحهٔ اقساط را پیاده‌سازی می‌کند.
type InstallmentHandler struct {
	DB *gorm.DB
}

func NewInstallmentHandler(db *gorm.DB) *InstallmentHandler { return &InstallmentHandler{DB: db} }

// ============================
//         مدل‌های داخلی
// ============================

type installmentPaymentRow struct {
	ID               uint       `json:"id"`
	CustomerName     string     `json:"customer_name"`
	NationalID       string     `json:"national_id"`
	Mobile           string     `json:"mobile"`
	ProductID        uint       `json:"product_id"`
	ProductName      string     `json:"product_name"`
	ProductSKU       string     `json:"product_sku"`
	Amount           float64    `json:"amount"`
	InstallmentCount int        `json:"installment_count"`
	MonthlyPayment   float64    `json:"monthly_payment"`
	Status           string     `json:"status"`
	CreditScore      int        `json:"credit_score"`
	ApprovedBy       *uint      `json:"approved_by"`
	ApprovedAt       *time.Time `json:"approved_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// ============================
//            API ها
// ============================

// ListInstallments
// دریافت لیست خریدهای اقساطی با فیلتر و صفحه‌بندی
func (h *InstallmentHandler) ListInstallments(c *gin.Context) {
	// ورودی‌ها
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	status := c.Query("status")
	search := c.Query("search")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	// ساخت کوئری پایه
	dbq := h.DB.Table("installment_payments AS ip").
		Select(`ip.id, ip.customer_name, ip.national_id, ip.mobile, ip.product_id, ip.product_name,
                ip.amount, ip.installment_count, ip.monthly_payment, ip.status, ip.credit_score,
                ip.approved_by, ip.approved_at, ip.created_at, ip.updated_at,
                p.sku AS product_sku`).
		Joins("LEFT JOIN products p ON p.id = ip.product_id")

	// فیلترها
	if status != "" {
		dbq = dbq.Where("ip.status = ?", status)
	}
	if search != "" {
		like := "%" + search + "%"
		dbq = dbq.Where("ip.customer_name ILIKE ? OR ip.national_id ILIKE ? OR ip.mobile ILIKE ?", like, like, like)
	}
	if startDate != "" {
		dbq = dbq.Where("ip.created_at >= ?", startDate)
	}
	if endDate != "" {
		dbq = dbq.Where("ip.created_at <= ?", endDate)
	}

	// شمارش
	var total int64
	if err := dbq.Count(&total).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "DB_ERROR"})
		return
	}

	// صفحه‌بندی
	var rows []installmentPaymentRow
	if err := dbq.Order("ip.created_at DESC").Limit(limit).Offset((page - 1) * limit).Scan(&rows).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "DB_ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"data":       rows,
		"pagination": gin.H{"page": page, "limit": limit, "total": total, "totalPages": (total + int64(limit) - 1) / int64(limit)},
	})
}

// GetStatusMetrics
// وضعیت کلی سلامت و متریک‌های سیستم برای داشبورد
func (h *InstallmentHandler) GetStatusMetrics(c *gin.Context) {
	// محاسبات نمونه و سبک — در سیستم واقعی باید از جداول گزارش/متریک یا کوئری‌های بهینه استفاده شود
	type T struct{ Count int64 }
	var tPending, tApproved, tRejected, tCompleted, tOverdue T
	h.DB.Table("installment_payments").Where("status = ?", "pending").Count(&tPending.Count)
	h.DB.Table("installment_payments").Where("status = ?", "approved").Count(&tApproved.Count)
	h.DB.Table("installment_payments").Where("status = ?", "rejected").Count(&tRejected.Count)
	h.DB.Table("installment_payments").Where("status = ?", "completed").Count(&tCompleted.Count)
	h.DB.Table("installment_payments").Where("status = ?", "overdue").Count(&tOverdue.Count)

	resp := gin.H{
		"overallHealth":            "healthy",
		"database":                 "healthy",
		"api":                      "healthy",
		"creditSystem":             "healthy",
		"databaseResponseTime":     10,
		"apiResponseTime":          12,
		"creditSystemAvailability": 99,
		"activeRequests":           0,
		"pendingApprovals":         tPending.Count,
		"processingRequests":       tApproved.Count,
		"completedToday":           0,
		"avgResponseTime":          15,
		"maxResponseTime":          40,
		"minResponseTime":          5,
		"requestSuccessRate":       99,
		"processingSuccessRate":    99,
		"approvalSuccessRate":      95,
		"recentAlerts":             []any{},
		"systemVersion":            "1.0.0",
		"lastSystemUpdate":         time.Now().Format("2006-01-02 15:04"),
		"updateAvailable":          false,
		"cpuUsage":                 12,
		"memoryUsage":              45,
		"diskUsage":                61,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
}

// GetStats آمار کلی برای کارت‌های Summary
func (h *InstallmentHandler) GetStats(c *gin.Context) {
	type S struct {
		TotalInstallments int64   `json:"totalInstallments"`
		TotalAmount       float64 `json:"totalAmount"`
		ApprovalRate      float64 `json:"approvalRate"`
		RepaymentRate     float64 `json:"repaymentRate"`
		InstallmentGrowth float64 `json:"installmentGrowth"`
		AmountGrowth      float64 `json:"amountGrowth"`
	}

	var s S
	h.DB.Table("installment_payments").Count(&s.TotalInstallments)
	h.DB.Table("installment_payments").Select("COALESCE(SUM(amount),0)").Scan(&s.TotalAmount)

	var approved, all int64
	h.DB.Table("installment_payments").Where("status = ?", "approved").Count(&approved)
	h.DB.Table("installment_payments").Count(&all)
	if all > 0 {
		s.ApprovalRate = float64(approved) * 100 / float64(all)
	}
	// در این نسخه ساده، نرخ بازپرداخت و رشد به صورت ثابت/ساده محاسبه می‌شود
	s.RepaymentRate = 90
	s.InstallmentGrowth = 10
	s.AmountGrowth = 12

	c.JSON(http.StatusOK, gin.H{"success": true, "data": s})
}

// GetTrends روند و توزیع‌ها برای نمودارها
func (h *InstallmentHandler) GetTrends(c *gin.Context) {
	// ورودی period اختیاری است
	_ = c.DefaultQuery("period", "30")

	// دادهٔ نمونه بر اساس داده‌های واقعی ساخته می‌شود (اینجا ساختار کامل برمی‌گردانیم)
	resp := gin.H{
		"newRequests":            0,
		"newRequestsChange":      0,
		"approvedRequests":       0,
		"approvedRequestsChange": 0,
		"totalAmount":            0,
		"totalAmountChange":      0,
		"approvalRate":           0,
		"approvalRateChange":     0,
		"dailyData":              []gin.H{},
		"maxRequests":            0,
		"maxApproved":            0,
		"maxRejected":            0,
		"amountRanges":           []gin.H{},
		"installmentRanges":      []gin.H{},
		"avgProcessingTime":      0,
		"avgCreditScore":         0,
		"avgInstallmentAmount":   0,
		"topDays":                []gin.H{},
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
}

// GetTopProducts محصولات برتر از نظر اقساط
func (h *InstallmentHandler) GetTopProducts(c *gin.Context) {
	// ورودی‌ها
	_ = c.DefaultQuery("period", "30")
	_ = c.DefaultQuery("sort", "requests")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// نمونهٔ خروجی — در نسخهٔ بعدی می‌توان با JOIN روی products و installment_payments تکمیل کرد
	resp := gin.H{
		"products": []any{},
		"summary":  gin.H{"totalProducts": 0, "totalRequests": 0, "totalAmount": 0, "avgApprovalRate": 0},
		"hasMore":  false,
	}
	_ = page
	_ = limit
	c.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
}

// ListInstallmentProducts مدیریت محصولات قابل اقساط
func (h *InstallmentHandler) ListInstallmentProducts(c *gin.Context) {
	// محصولات با ستون installment_enabled
	type P struct {
		ID               uint           `json:"id"`
		Name             string         `json:"name"`
		SKU              string         `json:"sku"`
		Image            string         `json:"image"`
		Price            float64        `json:"price"`
		InstallmentPrice float64        `json:"installmentPrice"`
		Installments     datatypes.JSON `json:"installments"`
		Status           string         `json:"status"`
		SalesCount       int64          `json:"salesCount"`
		TotalSales       float64        `json:"totalSales"`
	}

	var items []P
	// توجه: فیلد status در products موجود است؛ InstallmentPrice و Installments در مایگریشن 57 ایجاد شده‌اند
	err := h.DB.Table("products").
		Select("id, name, sku, image, price, installment_price, installment_options as installments, status, 0 as salesCount, 0 as totalSales").
		Where("installment_enabled = ?", true).Order("id DESC").
		Scan(&items).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "DB_ERROR"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": items})
}

// CreditSettingsUpdate ذخیره تنظیمات اعتباری سراسری
func (h *InstallmentHandler) CreditSettingsUpdate(c *gin.Context) {
	// این تنظیمات می‌تواند در جدول settings با Category = installment ذخیره شود
	var payload map[string]any
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "INVALID_PAYLOAD"})
		return
	}
	// Simplified: ذخیره به عنوان key-value
	// در این نسخه نمونه، فقط پاسخ موفق برمی‌گردانیم
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// Logs APIs
func (h *InstallmentHandler) ListLogs(c *gin.Context) {
	type L struct {
		ID        uint      `json:"id"`
		Level     string    `json:"level"`
		Title     string    `json:"title"`
		Message   string    `json:"message"`
		Details   string    `json:"details"`
		Timestamp time.Time `json:"timestamp"`
		Source    string    `json:"source"`
		User      string    `json:"user"`
	}
	// لاگ‌های سیستمی اقساط — اینجا خروجی خالی برمی‌گردانیم، ساختار حاضر است
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"logs": []L{}, "stats": gin.H{"totalLogs": 0, "infoLogs": 0, "warningLogs": 0, "errorLogs": 0}}})
}

func (h *InstallmentHandler) DeleteLogs(c *gin.Context) {
	// پاکسازی لاگ‌ها — در این نسخه فقط پاسخ موفق
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *InstallmentHandler) ExportLogs(c *gin.Context) {
	// آماده‌سازی اکسپورت — پاسخ موفق
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// CreditCheckRequest ورودی اعتبارسنجی
type CreditCheckRequest struct {
	NationalID       string  `json:"nationalId"`
	Mobile           string  `json:"mobile"`
	Amount           float64 `json:"amount"`
	InstallmentCount int     `json:"installmentCount"`
	ProductID        uint    `json:"productId"`
}

// CreditCheckResponse خروجی اعتبارسنجی
type CreditCheckResponse struct {
	Score           int      `json:"score"`
	Status          string   `json:"status"`
	MaxAmount       float64  `json:"maxAmount"`
	History         gin.H    `json:"history"`
	Recommendations []string `json:"recommendations"`
}

// CreditCheck محاسبه و ثبت نتیجه اعتبارسنجی
func (h *InstallmentHandler) CreditCheck(c *gin.Context) {
	var req CreditCheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "INVALID_PAYLOAD"})
		return
	}
	if req.NationalID == "" || req.Mobile == "" || req.Amount <= 0 || req.InstallmentCount <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "REQUIRED_FIELDS"})
		return
	}

	// محاسبهٔ ساده امتیاز اعتباری (نسخه اولیه)
	score := 500
	recommendations := []string{}
	monthly := req.Amount / float64(req.InstallmentCount)
	dti := (monthly / 5000000.0) * 100 // فرض درآمد 5M
	if dti > 50 {
		score -= 100
		recommendations = append(recommendations, "نسبت بدهی به درآمد بالا است")
	} else if dti > 30 {
		score -= 50
	} else {
		score += 50
	}
	if score < 0 {
		score = 0
	}
	if score > 1000 {
		score = 1000
	}

	status := "رد شده"
	maxAmount := 0.0
	if score >= 700 {
		status = "تایید شده"
		maxAmount = req.Amount * 2
	} else if score >= 500 {
		status = "نیاز به بررسی بیشتر"
		maxAmount = req.Amount
	} else {
		recommendations = append(recommendations, "امتیاز اعتباری ناکافی")
	}

	// ذخیره در credit_checks
	type CreditCheck struct {
		ID               uint           `gorm:"primaryKey"`
		NationalID       string         `gorm:"type:varchar(10);index"`
		Mobile           string         `gorm:"type:varchar(20);index"`
		RequestedAmount  float64        `gorm:"type:numeric(18,2);not null"`
		InstallmentCount int            `gorm:"not null"`
		CreditScore      int            `gorm:"index"`
		Status           string         `gorm:"type:varchar(50);index"`
		MaxAmount        float64        `gorm:"type:numeric(18,2);not null"`
		Recommendations  datatypes.JSON `gorm:"type:jsonb;default:'[]'"`
		CheckedBy        *uint          `gorm:"index"`
		CreatedAt        time.Time      `gorm:"autoCreateTime"`
	}
	if err := h.DB.Table("credit_checks").Create(&CreditCheck{
		NationalID:       req.NationalID,
		Mobile:           req.Mobile,
		RequestedAmount:  req.Amount,
		InstallmentCount: req.InstallmentCount,
		CreditScore:      score,
		Status:           status,
		MaxAmount:        maxAmount,
		Recommendations:  datatypes.JSON([]byte("[]")),
		CheckedBy:        nil,
	}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "message": "DB_ERROR"})
		return
	}

	resp := CreditCheckResponse{
		Score:           score,
		Status:          status,
		MaxAmount:       maxAmount,
		History:         gin.H{"totalLoans": 0, "overdueLoans": 0, "membershipDuration": 0},
		Recommendations: recommendations,
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
}

// ApproveInstallment تایید و ایجاد خرید اقساطی بر اساس نتیجه اعتبارسنجی
func (h *InstallmentHandler) ApproveInstallment(c *gin.Context) {
	var body struct {
		NationalID       string               `json:"nationalId"`
		Mobile           string               `json:"mobile"`
		Amount           float64              `json:"amount"`
		InstallmentCount int                  `json:"installmentCount"`
		ProductID        uint                 `json:"productId"`
		CreditResult     *CreditCheckResponse `json:"creditResult"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "INVALID_PAYLOAD"})
		return
	}
	if body.CreditResult == nil || body.CreditResult.Status == "رد شده" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "CREDIT_NOT_APPROVED"})
		return
	}
	if body.Amount > body.CreditResult.MaxAmount {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": "AMOUNT_EXCEEDS_LIMIT"})
		return
	}

	tx := h.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// یافتن یا ساخت کاربر
	var userID uint
	type U struct{ ID uint }
	var u U
	if err := tx.Raw("SELECT id FROM users WHERE national_id = ? OR mobile = ? LIMIT 1", body.NationalID, body.Mobile).Scan(&u).Error; err == nil && u.ID != 0 {
		userID = u.ID
		_ = tx.Exec("UPDATE users SET updated_at = NOW() WHERE id = ?", userID)
	} else {
		// در مدل User ستونی برای national_id نیست؛ در نسخهٔ اولیه فقط با موبایل ایجاد می‌کنیم
		if err := tx.Exec("INSERT INTO users (username, name, mobile, password_hash, role_id, status, registered_at, updated_at) VALUES (?,?,?,?,?,?,NOW(),NOW())",
			body.Mobile, "مشتری "+body.NationalID, body.Mobile, "", 1, "active").Error; err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
			return
		}
		if err := tx.Raw("SELECT id FROM users WHERE mobile = ? ORDER BY id DESC LIMIT 1", body.Mobile).Scan(&u).Error; err != nil || u.ID == 0 {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
			return
		}
		userID = u.ID
	}

	// دریافت محصول
	type P struct {
		ID   uint
		Name string
	}
	var p P
	if err := tx.Raw("SELECT id, name FROM products WHERE id = ?", body.ProductID).Scan(&p).Error; err != nil || p.ID == 0 {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"success": false, "message": "PRODUCT_NOT_FOUND"})
		return
	}

	monthly := float64(int64((body.Amount / float64(body.InstallmentCount)) + 0.5))

	// درج رکورد خرید اقساطی
	var installmentID uint
	if err := tx.Raw(
		`INSERT INTO installment_payments (user_id, customer_name, national_id, mobile, product_id, product_name, amount, installment_count, monthly_payment, credit_score, status, approved_by, approved_at, created_at, updated_at, metadata)
         VALUES (?,?,?,?,?,?,?,?,?,?,?,NULL,NOW(),NOW(),NOW(),'{}') RETURNING id`,
		userID, "مشتری "+body.NationalID, body.NationalID, body.Mobile, p.ID, p.Name, body.Amount, body.InstallmentCount, monthly, body.CreditResult.Score, "approved",
	).Scan(&installmentID).Error; err != nil {
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	// درج اقساط برنامه
	now := time.Now()
	for i := 1; i <= body.InstallmentCount; i++ {
		due := now.AddDate(0, i, 0)
		status := "pending"
		if i == 1 {
			status = "current"
		}
		if err := tx.Exec(
			`INSERT INTO installment_schedules (installment_id, installment_number, amount, due_date, status, created_at)
             VALUES (?,?,?,?,?,NOW())`,
			installmentID, i, monthly, due, status,
		).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
			return
		}
	}

	// ثبت لاگ
	_ = tx.Exec(`INSERT INTO installment_logs (installment_id, action, description, performed_by, created_at) VALUES (?,?,?,NULL,NOW())`,
		installmentID, "approve", "تایید خرید اقساطی").Error

	if err := tx.Commit().Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"installmentId": installmentID,
		"message":       "خرید اقساطی با موفقیت تایید شد",
	}})
}

// ListPendingRequests لیست درخواست‌های در انتظار
func (h *InstallmentHandler) ListPendingRequests(c *gin.Context) {
	type R struct {
		ID           uint      `json:"id"`
		CustomerName string    `json:"customerName"`
		NationalID   string    `json:"nationalId"`
		ProductName  string    `json:"productName"`
		Amount       float64   `json:"amount"`
		CreatedAt    time.Time `json:"createdAt"`
	}
	var rows []R
	if err := h.DB.Table("installment_payments").Select("id, customer_name, national_id, product_name, amount, created_at").Where("status = ?", "pending").Order("created_at DESC").Scan(&rows).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": rows})
}

// ApprovePendingRequest تایید یک درخواست در انتظار با شناسه
func (h *InstallmentHandler) ApprovePendingRequest(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	if err := h.DB.Exec("UPDATE installment_payments SET status = 'approved', approved_at = NOW(), updated_at = NOW() WHERE id = ?", id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// SettingsUpdate ذخیره تنظیمات عمومی صفحهٔ تنظیمات اقساط
func (h *InstallmentHandler) SettingsUpdate(c *gin.Context) {
	var payload map[string]any
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	// در این نسخه نمونه، فقط موفقیت برمی‌گردد
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// BackupCreate ایجاد پشتیبان (نمونه)
func (h *InstallmentHandler) BackupCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// Validation methods APIs (نمونه)
func (h *InstallmentHandler) GetValidationMethods(c *gin.Context) {
	resp := gin.H{
		"overview": gin.H{"totalValidations": 0, "successRate": 0, "avgTime": 0, "activeMethods": 4},
		"methods": gin.H{
			"nationalId":   gin.H{"status": "active", "totalChecks": 0, "successRate": 0, "avgResponseTime": 0, "errorRate": 0},
			"mobile":       gin.H{"status": "active", "totalChecks": 0, "successRate": 0, "avgResponseTime": 0, "errorRate": 0},
			"bankAccount":  gin.H{"status": "active", "totalChecks": 0, "successRate": 0, "avgResponseTime": 0, "errorRate": 0},
			"creditBureau": gin.H{"status": "active", "totalChecks": 0, "successRate": 0, "avgResponseTime": 0, "errorRate": 0},
			"employment":   gin.H{"status": "active", "totalChecks": 0, "successRate": 0, "avgResponseTime": 0, "errorRate": 0},
		},
		"globalSettings": gin.H{"timeout": 30, "retryCount": 3},
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
}

func (h *InstallmentHandler) TestValidationMethod(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *InstallmentHandler) TestAllValidationMethods(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *InstallmentHandler) UpdateValidationSettings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// Geography داده‌های جغرافیایی (نمونه)
func (h *InstallmentHandler) GetGeography(c *gin.Context) {
	resp := gin.H{
		"summary":        gin.H{"totalProvinces": 0, "totalCities": 0, "topProvince": "", "coverage": 0},
		"topProvinces":   []any{},
		"topCities":      []any{},
		"regionalData":   []any{},
		"growthAnalysis": gin.H{"fastestGrowing": "", "slowestGrowing": "", "mostPotential": ""},
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
}

// Export endpoints (نمونه)
func (h *InstallmentHandler) ExportExcel(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}
func (h *InstallmentHandler) ExportPDF(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) }
func (h *InstallmentHandler) ExportCSV(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true}) }
func (h *InstallmentHandler) SaveExportSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// Product management for installments
func (h *InstallmentHandler) CreateInstallmentProduct(c *gin.Context) {
	var payload struct {
		Name             string           `json:"name"`
		SKU              string           `json:"sku"`
		Price            float64          `json:"price"`
		InstallmentPrice float64          `json:"installmentPrice"`
		Installments     []map[string]any `json:"installments"`
		CreditCheck      map[string]any   `json:"creditCheck"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	// درج محصول حداقلی
	if err := h.DB.Exec(
		`INSERT INTO products (uuid, slug, sku, name, price, status, installment_enabled, installment_price, installment_options, created_at, updated_at)
         VALUES (gen_random_uuid(), ?, ?, ?, ?, 'active', TRUE, ?, ?, NOW(), NOW())`,
		payload.Name, payload.SKU, payload.Name, payload.Price, payload.InstallmentPrice, datatypes.JSON([]byte("[]")),
	).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *InstallmentHandler) UpdateInstallmentProduct(c *gin.Context) {
	id := c.Param("id")
	var payload struct {
		Name             *string          `json:"name"`
		SKU              *string          `json:"sku"`
		Price            *float64         `json:"price"`
		InstallmentPrice *float64         `json:"installmentPrice"`
		Installments     []map[string]any `json:"installments"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}
	// به‌روزرسانی ساده
	updates := map[string]any{"updated_at": time.Now()}
	if payload.Name != nil {
		updates["name"] = *payload.Name
	}
	if payload.SKU != nil {
		updates["sku"] = *payload.SKU
	}
	if payload.Price != nil {
		updates["price"] = *payload.Price
	}
	if payload.InstallmentPrice != nil {
		updates["installment_price"] = *payload.InstallmentPrice
	}
	if len(payload.Installments) > 0 {
		b, _ := json.Marshal(payload.Installments)
		updates["installment_options"] = datatypes.JSON(b)
		updates["installment_enabled"] = true
	}
	if err := h.DB.Table("products").Where("id = ?", id).Updates(updates).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *InstallmentHandler) ToggleProductStatus(c *gin.Context) {
	id := c.Param("id")
	// toggle between active/inactive
	if err := h.DB.Exec(`UPDATE products SET status = CASE WHEN status='active' THEN 'inactive' ELSE 'active' END, updated_at = NOW() WHERE id = ?`, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
