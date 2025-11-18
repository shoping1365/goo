package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
)

// TestPatternHandler هندلر جداگانه برای تست پترن‌های پیامک
type TestPatternHandler struct {
	db         *gorm.DB
	smsService *services.SMSService
}

// NewTestPatternHandler ایجاد نمونه جدید از هندلر تست پترن
func NewTestPatternHandler(db *gorm.DB) *TestPatternHandler {
	return &TestPatternHandler{
		db:         db,
		smsService: services.NewSMSService(db),
	}
}

// TestPattern تست پترن
func (h *TestPatternHandler) TestPattern(c *gin.Context) {
	id := c.Param("id")
	patternID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		utils.BadRequest(c, "شناسه پترن نامعتبر است", err)
		return
	}

	// دریافت پترن
	var pattern models.SMSPattern
	if err := h.db.Preload("Gateway").First(&pattern, patternID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "پترن یافت نشد", err)
			return
		}
		utils.InternalServerError(c, "خطا در دریافت پترن", err)
		return
	}

	// بررسی فعال بودن پترن
	if pattern.Status != "active" {
		utils.BadRequest(c, "پترن غیرفعال است و قابل تست نیست", nil)
		return
	}

	// دریافت شماره موبایل از درخواست
	var request struct {
		Phone     string            `json:"phone" binding:"required"`
		Variables map[string]string `json:"variables"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.BadRequest(c, "داده‌های ارسالی نامعتبر است", err)
		return
	}

	log.Printf("🔍 درخواست دریافت شده: Phone='%s', Variables=%v", request.Phone, request.Variables)

	// آماده‌سازی متغیرهای پترن - فقط متغیرهای تعریف شده در پترن
	patternValues := make(map[string]string)

	// اگر متغیرهای پترن تعریف شده باشند، فقط آن‌ها را پردازش کن
	if pattern.Variables != "" {
		log.Printf("🔍 متغیرهای پترن خام: '%s'", pattern.Variables)

		// متغیرهای پترن به صورت JSON ذخیره شده‌اند
		var patternVars []string
		if err := json.Unmarshal([]byte(pattern.Variables), &patternVars); err != nil {
			log.Printf("🔍 JSON نبود، پردازش به صورت رشته: %v", err)
			// اگر JSON نبود، به صورت رشته معمولی پردازش کن
			variables := strings.Split(pattern.Variables, ",")
			for _, variable := range variables {
				variable = strings.TrimSpace(variable)
				if variable != "" {
					// حذف % از ابتدا و انتها اگر وجود داشته باشد
					if strings.HasPrefix(variable, "%") && strings.HasSuffix(variable, "%") {
						variable = variable[1 : len(variable)-1]
					}
					patternVars = append(patternVars, variable)
				}
			}
		} else {
			log.Printf("🔍 متغیرهای پترن JSON: %v", patternVars)
		}

		// برای هر متغیر پترن، تمام متغیرها را ارسال کن (حتی اگر خالی باشند)
		for _, variable := range patternVars {
			log.Printf("🔍 بررسی متغیر: '%s'", variable)
			if value, exists := request.Variables[variable]; exists {
				patternValues[variable] = value
				log.Printf("✅ متغیر اضافه شد: '%s' = '%s'", variable, value)
			} else {
				// اگر متغیر در درخواست وجود ندارد، مقدار خالی ارسال کن
				patternValues[variable] = ""
				log.Printf("✅ متغیر با مقدار خالی اضافه شد: '%s' = ''", variable)
			}
		}
	}

	log.Printf("🔍 متغیرهای نهایی: %v", patternValues)

	// بررسی اینکه آیا درگاه با پترن کار می‌کند یا نه
	if pattern.Gateway.PatternBased {
		log.Printf("🔍 ارسال با پترن - Scope='%s', Feature='%s'", pattern.Scope, pattern.Feature)
		// ارسال با پترن
		messageID, err := h.smsService.SendPatternByScopeAndFeature(pattern.Scope, pattern.Feature, request.Phone, patternValues)
		if err != nil {
			utils.InternalServerError(c, "خطا در ارسال پیامک تست", err)
			return
		}

		// افزایش شمارنده استفاده
		h.db.Model(&pattern).Update("usage_count", pattern.UsageCount+1)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "پیامک تست با موفقیت ارسال شد",
			"data": gin.H{
				"message_id": messageID,
				"pattern":    pattern,
				"variables":  patternValues,
				"method":     "pattern",
			},
		})
	} else {
		log.Printf("🔍 ارسال مستقیم - GatewayID=%d", pattern.GatewayID)
		// ارسال بدون پترن (مستقیم)
		messageID, err := h.smsService.SendDirectSMS(pattern.GatewayID, request.Phone, pattern.MessageTemplate, patternValues)
		if err != nil {
			utils.InternalServerError(c, "خطا در ارسال پیامک تست", err)
			return
		}

		// افزایش شمارنده استفاده
		h.db.Model(&pattern).Update("usage_count", pattern.UsageCount+1)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "پیامک تست با موفقیت ارسال شد",
			"data": gin.H{
				"message_id": messageID,
				"pattern":    pattern,
				"variables":  patternValues,
				"method":     "direct",
			},
		})
	}
}