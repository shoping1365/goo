package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
)

// SMSPatternProxyHandler فقط فراخوانی ایجاد پترن را به مسیر موجود تفویض می‌کند (عدم تکرار منطق)
// این هندلر نمونه‌ی کوچک برای راه‌اندازی سریع کلاینت‌هاست؛ مسیر اصلی ساخت پترن در sms_pattern_router موجود است.
type SMSPatternProxyHandler struct {
	DB         *gorm.DB
	smsService *services.SMSService
}

func NewSMSPatternProxyHandler(db *gorm.DB) *SMSPatternProxyHandler {
	return &SMSPatternProxyHandler{
		DB:         db,
		smsService: services.NewSMSService(db),
	}
}

// CreateCustomerDiscountNotifyPattern درخواست ساخت پترن «اطلاع‌رسانی تخفیف» برای مشتری را به روت اصلی forward می‌کند.
// توجه: این متد پترن را نمی‌سازد، فقط API متمرکز مدیریت الگو را صدا می‌زند. اگر پترنی وجود نداشت، باید در همان سرویس ساخته شود.
func (h *SMSPatternProxyHandler) CreateCustomerDiscountNotifyPattern(c *gin.Context) {
	c.JSON(http.StatusBadRequest, utils.New("FORWARD_REQUIRED", "از مسیر مدیریت پترن‌ها استفاده کنید: /api/admin/sms-patterns", nil))
}

// CreateCustomerLowStockNotifyPattern مشابه بالا برای «کمبود موجودی»
func (h *SMSPatternProxyHandler) CreateCustomerLowStockNotifyPattern(c *gin.Context) {
	c.JSON(http.StatusBadRequest, utils.New("FORWARD_REQUIRED", "از مسیر مدیریت پترن‌ها استفاده کنید: /api/admin/sms-patterns", nil))
}

// SendByPattern ارسال پیام بر اساس pattern_code موجود
// این مسیر یک نمونه اجرایی است: با pattern_code و mobile و variables از SMSService استفاده می‌کند.
type sendByPatternRequest struct {
	PatternCode string            `json:"pattern_code"`
	Mobile      string            `json:"mobile" binding:"required"`
	GatewayID   uint              `json:"gateway_id"`
	Scope       string            `json:"scope"`   // دریافت‌کننده (manager/customer)
	Feature     string            `json:"feature"` // هدف/آیتم (auth_otp, order_confirmation, ...)
	Variables   map[string]string `json:"variables"`
}

func (h *SMSPatternProxyHandler) SendByPattern(c *gin.Context) {
	var req sendByPatternRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}

	// استفاده از مدل SMSSendRequest (ارسال واقعی در سرویس مرکزی انجام می‌شود)
	smsRequest := models.SMSSendRequest{
		PatternCode: req.PatternCode,
		Mobile:      req.Mobile,
		GatewayID:   req.GatewayID,
		Scope:       req.Scope,
		Feature:     req.Feature,
		Variables:   req.Variables,
	}

	// ارسال واقعی پیامک
	response, err := h.smsService.SendSMS(smsRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("SMS_SEND_ERROR", "خطا در ارسال پیامک", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "پیامک با موفقیت ارسال شد",
		"data":    response,
	})
}

// SMSPatternHandler مدیریت درخواست‌های مربوط به پترن‌های پیامک
type SMSPatternHandler struct {
	db         *gorm.DB
	smsService *services.SMSService
}

// NewSMSPatternHandler ایجاد نمونه جدید از handler
func NewSMSPatternHandler(db *gorm.DB) *SMSPatternHandler {
	return &SMSPatternHandler{
		db:         db,
		smsService: services.NewSMSService(db),
	}
}

// ifEmpty اگر رشته خالی بود مقدار پیش‌فرض برمی‌گرداند
func ifEmpty(s string, def string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return def
	}
	return s
}

// generateNextPatternID تولید ID بعدی برای سیستم
func (h *SMSPatternHandler) generateNextPatternID(gatewayID uint) (uint, error) {
	var lastPattern models.SMSPattern
	if err := h.db.Order("id DESC").First(&lastPattern).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// اگر هیچ پترنی در سیستم وجود ندارد، از ID 1 شروع کن
			return 1, nil
		}
		return 0, err
	}

	// ID بعدی (بدون محدودیت)
	nextID := lastPattern.ID + 1

	return nextID, nil
}

// CleanupInvalidPatterns پاکسازی یکباره پترن‌های نامعتبر
func (h *SMSPatternHandler) CleanupInvalidPatterns(c *gin.Context) {
	// غیرفعال شده - پترن‌ها حذف نمی‌شوند
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "پاکسازی غیرفعال شده است",
	})
}

// GetPatterns دریافت لیست پترن‌های پیامک
func (h *SMSPatternHandler) GetPatterns(c *gin.Context) {
	var patterns []models.SMSPattern

	// دریافت پارامترهای فیلتر
	typeFilter := c.Query("type")
	categoryFilter := c.Query("category")
	statusFilter := c.Query("status")
	gatewayFilter := c.Query("gateway")
	searchFilter := c.Query("search")
	featureFilter := c.Query("feature")

	query := h.db.Preload("Gateway")

	// اعمال فیلترها (منطبق با ستون‌های موجود جدول)
	if typeFilter != "" {
		// ستون type در جدول فعلی وجود ندارد؛ نادیده گرفته می‌شود
	}
	if categoryFilter != "" {
		// ستون category در جدول فعلی وجود ندارد؛ نادیده گرفته می‌شود
	}
	if statusFilter != "" {
		query = query.Where("status = ?", statusFilter)
	}
	if gatewayFilter != "" {
		// پشتیبانی از نام درگاه (مثلاً IPPanel) → تبدیل به gateway_id
		if strings.TrimSpace(gatewayFilter) != "" && strings.TrimSpace(gatewayFilter) == gatewayFilter {
			// اگر مقدار عددی باشد همان را استفاده کن
			if _, err := strconv.Atoi(gatewayFilter); err == nil {
				query = query.Where("gateway_id = ?", gatewayFilter)
			} else {
				// تبدیل نام به ID
				var gw models.SMSGateway
				if err := h.db.Where("LOWER(type) = LOWER(?) OR LOWER(name) = LOWER(?)", gatewayFilter, gatewayFilter).First(&gw).Error; err == nil {
					query = query.Where("gateway_id = ?", gw.ID)
				}
			}
		}
	}
	if searchFilter != "" {
		searchTerm := "%" + searchFilter + "%"
		query = query.Where("name LIKE ? OR pattern_code LIKE ? OR feature = ?", searchTerm, searchTerm, searchFilter)
	}
	if featureFilter != "" {
		query = query.Where("feature = ?", featureFilter)
	}

	// دریافت نتایج
	if err := query.Find(&patterns).Error; err != nil {
		utils.InternalServerError(c, "خطا در دریافت پترن‌ها", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    patterns,
	})
}

// GetPattern دریافت پترن خاص
func (h *SMSPatternHandler) GetPattern(c *gin.Context) {
	id := c.Param("id")
	patternID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.BadRequest(c, "شناسه پترن نامعتبر است", err)
		return
	}

	var pattern models.SMSPattern
	if err := h.db.Preload("Gateway").First(&pattern, patternID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "پترن یافت نشد", err)
			return
		}
		utils.InternalServerError(c, "خطا در دریافت پترن", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pattern,
	})
}

// CreatePattern ایجاد پترن جدید
func (h *SMSPatternHandler) CreatePattern(c *gin.Context) {
	var form models.SMSPatternForm
	if err := c.ShouldBindJSON(&form); err != nil {
		utils.BadRequest(c, "داده‌های ورودی نامعتبر است", err)
		return
	}

	// پاکسازی ورودی‌ها
	form.Name = strings.TrimSpace(form.Name)
	form.PatternCode = strings.TrimSpace(form.PatternCode)
	form.Status = strings.TrimSpace(form.Status)

	// تولید ID خودکار برای درگاه
	if form.FixedID == 0 {
		nextID, err := h.generateNextPatternID(form.GatewayID)
		if err != nil {
			utils.BadRequest(c, err.Error(), err)
			return
		}
		form.FixedID = nextID
	}

	// بررسی یکتا بودن ID
	var existingPattern models.SMSPattern
	if err := h.db.Where("id = ?", form.FixedID).First(&existingPattern).Error; err == nil {
		utils.SendError(c, http.StatusConflict, "CONFLICT", "این ID قبلاً استفاده شده است", nil)
		return
	}

	// بررسی وجود پترن فعال برای همین درگاه و feature
	var activePattern models.SMSPattern
	if err := h.db.Where("gateway_id = ? AND feature = ? AND status = 'active'", form.GatewayID, form.Feature).First(&activePattern).Error; err == nil {
		utils.SendError(c, http.StatusConflict, "DUPLICATE_FEATURE", fmt.Sprintf("برای درگاه %d و هدف '%s' قبلاً پترن فعالی وجود دارد", form.GatewayID, form.Feature), nil)
		return
	}

	// ایجاد پترن جدید با ID مشخص
	createData := map[string]interface{}{
		"id":           form.FixedID,
		"type":         ifEmpty(form.Type, "sms"),
		"scope":        ifEmpty(form.Scope, "customer"),
		"feature":      strings.TrimSpace(form.Feature),
		"name":         form.Name,
		"pattern_code": form.PatternCode,
		"description":  strings.TrimSpace(form.Description),
		"variables":    strings.TrimSpace(form.Variables),
		"template":     strings.TrimSpace(form.MessageTemplate),
		"gateway_id":   form.GatewayID,
		"status":       form.Status,
	}

	if err := h.db.Table("sms_patterns").Create(&createData).Error; err != nil {
		utils.InternalServerError(c, "خطا در ایجاد پترن", err)
		return
	}

	// واکشی رکورد ایجادشده جهت بازگرداندن شناسه
	var created models.SMSPattern
	if err := h.db.Preload("Gateway").Where("id = ?", form.FixedID).First(&created).Error; err != nil {
		// اگر نتوانست بخواند، حداقل همان داده‌های ورودی را برگردانیم
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"data":    createData,
			"message": "پترن ایجاد شد (شناسه قابل بازیابی نبود)",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    created,
		"message": "پترن با موفقیت ایجاد شد",
	})
}

// UpdatePattern ویرایش پترن
func (h *SMSPatternHandler) UpdatePattern(c *gin.Context) {
	id := c.Param("id")
	patternID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.BadRequest(c, "شناسه پترن نامعتبر است", err)
		return
	}

	var form models.SMSPatternForm
	if err := c.ShouldBindJSON(&form); err != nil {
		utils.BadRequest(c, "داده‌های ورودی نامعتبر است", err)
		return
	}

	// پاکسازی ورودی‌ها
	form.Name = strings.TrimSpace(form.Name)
	form.PatternCode = strings.TrimSpace(form.PatternCode)
	form.Status = strings.TrimSpace(form.Status)

	// بررسی وجود پترن
	var pattern models.SMSPattern
	if err := h.db.First(&pattern, patternID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "پترن یافت نشد", err)
			return
		}
		utils.InternalServerError(c, "خطا در دریافت پترن", err)
		return
	}

	// بررسی یکتا بودن ID (اگر تغییر کرده باشد)
	if pattern.ID != form.FixedID {
		var existingPattern models.SMSPattern
		if err := h.db.Where("id = ? AND id != ?", form.FixedID, patternID).First(&existingPattern).Error; err == nil {
			utils.SendError(c, http.StatusConflict, "CONFLICT", "این ID قبلاً استفاده شده است", nil)
			return
		}
	}

	// بررسی وجود پترن فعال برای همین درگاه و feature (به جز خود پترن فعلی)
	var activePattern models.SMSPattern
	if err := h.db.Where("gateway_id = ? AND feature = ? AND status = 'active' AND id != ?", form.GatewayID, form.Feature, patternID).First(&activePattern).Error; err == nil {
		utils.SendError(c, http.StatusConflict, "DUPLICATE_FEATURE", fmt.Sprintf("برای درگاه %d و هدف '%s' قبلاً پترن فعالی وجود دارد", form.GatewayID, form.Feature), nil)
		return
	}

	// به‌روزرسانی پترن (فقط ستون‌های موجود)
	updates := map[string]interface{}{
		"id":           form.FixedID,
		"type":         ifEmpty(form.Type, "sms"),
		"scope":        ifEmpty(form.Scope, "customer"),
		"feature":      strings.TrimSpace(form.Feature),
		"name":         form.Name,
		"pattern_code": form.PatternCode,
		"description":  strings.TrimSpace(form.Description),
		"variables":    strings.TrimSpace(form.Variables),
		"template":     strings.TrimSpace(form.MessageTemplate),
		"gateway_id":   form.GatewayID,
		"status":       form.Status,
	}

	if err := h.db.Model(&pattern).Select("id", "name", "pattern_code", "gateway_id", "status", "scope", "feature", "description", "variables", "template").Updates(updates).Error; err != nil {
		utils.InternalServerError(c, "خطا در ویرایش پترن", err)
		return
	}

	// بارگذاری مجدد پترن به‌روزرسانی شده
	if err := h.db.Preload("Gateway").First(&pattern, patternID).Error; err != nil {
		utils.InternalServerError(c, "خطا در بارگذاری پترن به‌روزرسانی شده", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pattern,
		"message": "پترن با موفقیت ویرایش شد",
	})
}

// DeletePattern حذف پترن
func (h *SMSPatternHandler) DeletePattern(c *gin.Context) {
	// کدهای پترن رزرو شده که نباید حذف شوند
	reserved := map[string]bool{
		"admin_failover": true,
		"admin_test":     true,
		"otp_login":      true,
		"admin_order":    true,
		"low_stock":      true,
		"security_issue": true,
	}

	id := c.Param("id")
	patternID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.BadRequest(c, "شناسه پترن نامعتبر است", err)
		return
	}

	// بررسی وجود پترن
	var pattern models.SMSPattern
	if err := h.db.First(&pattern, patternID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "پترن یافت نشد", err)
			return
		}
		utils.InternalServerError(c, "خطا در دریافت پترن", err)
		return
	}

	// جلوگیری از حذف پترن‌های سیستمی
	if reserved[pattern.PatternCode] {
		utils.BadRequest(c, "این پترن سیستمی است و قابل حذف نیست", nil)
		return
	}

	// حذف پترن
	if err := h.db.Delete(&pattern).Error; err != nil {
		utils.InternalServerError(c, "خطا در حذف پترن", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "پترن با موفقیت حذف شد",
	})
}

// GetPatternStats دریافت آمار پترن‌ها
func (h *SMSPatternHandler) GetPatternStats(c *gin.Context) {
	var stats struct {
		TotalPatterns    int64 `json:"total_patterns"`
		ActivePatterns   int64 `json:"active_patterns"`
		InactivePatterns int64 `json:"inactive_patterns"`
		DraftPatterns    int64 `json:"draft_patterns"`
		TotalUsage       int64 `json:"total_usage"`
	}

	// محاسبه آمار
	h.db.Model(&models.SMSPattern{}).Count(&stats.TotalPatterns)
	h.db.Model(&models.SMSPattern{}).Where("status = ?", "active").Count(&stats.ActivePatterns)
	h.db.Model(&models.SMSPattern{}).Where("status = ?", "inactive").Count(&stats.InactivePatterns)
	h.db.Model(&models.SMSPattern{}).Where("status = ?", "draft").Count(&stats.DraftPatterns)
	h.db.Model(&models.SMSPattern{}).Select("COALESCE(SUM(usage_count), 0)").Scan(&stats.TotalUsage)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    stats,
	})
}

// GetPatternByFeature دریافت آخرین پترن فعال بر اساس نوع/اسکوپ/فیچر برای مصرف در سایر ماژول‌ها
// نمونه مصرف: /api/admin/sms-patterns/by-feature?type=sms&scope=customer&feature=discount_notify
func (h *SMSPatternHandler) GetPatternByFeature(c *gin.Context) {
	typ := strings.TrimSpace(c.Query("type"))
	if typ == "" {
		typ = "sms"
	}
	scope := strings.TrimSpace(c.Query("scope"))
	feature := strings.TrimSpace(c.Query("feature"))
	if scope == "" || feature == "" {
		utils.BadRequest(c, "پارامترهای scope و feature الزامی هستند", nil)
		return
	}

	var pattern models.SMSPattern
	if err := h.db.Where("type = ? AND scope = ? AND feature = ? AND status = ?", typ, scope, feature, "active").
		Order("id DESC").First(&pattern).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "پترن فعال یافت نشد", err)
			return
		}
		utils.InternalServerError(c, "خطا در جستجوی پترن", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pattern,
	})
}

// DuplicatePattern کپی کردن پترن
func (h *SMSPatternHandler) DuplicatePattern(c *gin.Context) {
	id := c.Param("id")
	patternID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.BadRequest(c, "شناسه پترن نامعتبر است", err)
		return
	}

	// دریافت پترن اصلی
	var originalPattern models.SMSPattern
	if err := h.db.Preload("Gateway").First(&originalPattern, patternID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "پترن یافت نشد", err)
			return
		}
		utils.InternalServerError(c, "خطا در دریافت پترن", err)
		return
	}

	// تولید ID خودکار برای کپی
	nextID, err := h.generateNextPatternID(originalPattern.GatewayID)
	if err != nil {
		utils.BadRequest(c, err.Error(), err)
		return
	}

	// ایجاد کپی جدید با ID خالی
	newPattern := models.SMSPattern{
		ID:              nextID,
		Name:            originalPattern.Name + " (کپی)",
		PatternCode:     fmt.Sprintf("%s_copy_%d", originalPattern.PatternCode, time.Now().Unix()),
		Description:     originalPattern.Description,
		Variables:       originalPattern.Variables,
		MessageTemplate: originalPattern.MessageTemplate,
		GatewayID:       originalPattern.GatewayID,
		Status:          "draft", // کپی به صورت پیش‌نویس
		UsageCount:      0,       // تعداد استفاده صفر
	}

	if err := h.db.Create(&newPattern).Error; err != nil {
		utils.InternalServerError(c, "خطا در کپی کردن پترن", err)
		return
	}

	// بارگذاری اطلاعات درگاه برای پاسخ
	if err := h.db.Preload("Gateway").First(&newPattern, newPattern.ID).Error; err != nil {
		utils.InternalServerError(c, "خطا در بارگذاری اطلاعات کپی شده", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    newPattern,
		"message": "پترن با موفقیت کپی شد",
	})
}

// SendSMSByScopeAndFeature ارسال پیامک بر اساس scope و feature
func (h *SMSPatternHandler) SendSMSByScopeAndFeature(c *gin.Context) {
	var req struct {
		Mobile    string            `json:"mobile" binding:"required"`
		Scope     string            `json:"scope" binding:"required"`   // دریافت‌کننده (manager/customer)
		Feature   string            `json:"feature" binding:"required"` // هدف/آیتم (auth_otp, order_confirmation, ...)
		GatewayID uint              `json:"gateway_id"`
		Variables map[string]string `json:"variables"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "داده‌های ورودی نامعتبر است", err)
		return
	}

	// ارسال پیامک بر اساس scope و feature
	messageID, err := h.smsService.SendPatternByScopeAndFeature(req.Scope, req.Feature, req.Mobile, req.Variables)
	if err != nil {
		utils.InternalServerError(c, "خطا در ارسال پیامک", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "پیامک با موفقیت ارسال شد",
		"data": map[string]interface{}{
			"message_id": messageID,
			"scope":      req.Scope,
			"feature":    req.Feature,
		},
	})
}

// GetGatewayStatus بررسی وضعیت درگاه‌های پیامکی
func (h *SMSPatternHandler) GetGatewayStatus(c *gin.Context) {
	var gateways []models.SMSGateway
	if err := h.db.Find(&gateways).Error; err != nil {
		utils.InternalServerError(c, "خطا در دریافت درگاه‌ها", err)
		return
	}

	var status []map[string]interface{}
	for _, gateway := range gateways {
		status = append(status, map[string]interface{}{
			"id":            gateway.ID,
			"name":          gateway.Name,
			"type":          gateway.Type,
			"is_active":     gateway.IsActive,
			"priority":      gateway.Priority,
			"sender_number": gateway.SenderNumber,
		})
	}

	// اضافه کردن اطلاعات debug
	var activeCount int
	for _, gateway := range gateways {
		if gateway.IsActive {
			activeCount++
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    status,
		"message": fmt.Sprintf("وضعیت درگاه‌های پیامکی - تعداد کل: %d, تعداد فعال: %d", len(gateways), activeCount),
	})
}
