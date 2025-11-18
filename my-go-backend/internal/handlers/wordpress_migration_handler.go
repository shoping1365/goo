package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WordPressMigrationHandler struct {
	DB      *gorm.DB
	Service *services.WordPressMigrationService
}

var (
	migrationLogsMu    sync.Mutex
	migrationLogsBuf   []map[string]interface{}
	currentMigCancelMu sync.Mutex
	currentMigCancel   context.CancelFunc
)

func pushMigrationLog(logType, message string) {
	migrationLogsMu.Lock()
	defer migrationLogsMu.Unlock()
	if len(migrationLogsBuf) >= 200 {
		migrationLogsBuf = migrationLogsBuf[1:]
	}
	migrationLogsBuf = append(migrationLogsBuf, map[string]interface{}{
		"timestamp": time.Now().Format("15:04:05"),
		"type":      strings.ToLower(logType),
		"message":   message,
	})
}

func NewWordPressMigrationHandler(db *gorm.DB) *WordPressMigrationHandler {
	return &WordPressMigrationHandler{DB: db, Service: services.NewWordPressMigrationService()}
}

// ---------------- Scheduling -----------------

type MigrationStage struct {
	Time                     string `json:"time"`
	BatchSize                int    `json:"batchSize"`
	Products                 bool   `json:"products"`
	Orders                   bool   `json:"orders"`
	WordpressUsers           bool   `json:"wordpressUsers"`
	DigitsUsers              bool   `json:"digitsUsers"`
	Categories               bool   `json:"categories"`
	Brands                   bool   `json:"brands"`
	Attributes               bool   `json:"attributes"`
	RedirectProducts         bool   `json:"redirectProducts"`
	RedirectCode             int    `json:"redirectCode"`
	RedirectGroupName        string `json:"redirectGroupName"`
	DraftMode                bool   `json:"draftMode"`
	IncludeSlugInURL         bool   `json:"includeSlugInURL"`
	TransferSEOfromYoast     bool   `json:"transferSEOfromYoast"`
	TransferSEOfromRankMath  bool   `json:"transferSEOfromRankMath"`
	TransferPrices           bool   `json:"transferPrices"`
	TransferStock            bool   `json:"transferStock"`
	TransferVariants         bool   `json:"transferVariants"`
	TransferDescriptions     bool   `json:"transferDescriptions"`
	TransferShortDescription bool   `json:"transferShortDescription"`
	TransferImages           bool   `json:"transferImages"`
	TransferAttributesFields bool   `json:"transferAttributesFields"`
}

func (h *WordPressMigrationHandler) SavePlan(c *gin.Context) {
	var req struct {
		Config struct {
			SiteURL        string `json:"siteUrl"`
			APIKey         string `json:"apiKey"`
			ConsumerKey    string `json:"consumerKey"`
			ConsumerSecret string `json:"consumerSecret"`
		} `json:"config"`
		Stages []MigrationStage `json:"stages"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Config.SiteURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "درخواست نامعتبر"})
		return
	}
	planKey := "wp_migration.plan"
	stateKey := "wp_migration.plan_state"
	b, _ := json.Marshal(req)
	if err := h.DB.Save(&models.Setting{Key: planKey, Value: string(b), Category: "wp_migration", Type: "json"}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	_ = h.DB.Where("key = ?", stateKey).Delete(&models.Setting{}).Error
	_ = h.DB.Create(&models.Setting{Key: stateKey, Value: `{"nextIndex":0}`, Category: "wp_migration", Type: "json"}).Error
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "پلن زمان‌بندی ذخیره شد"})
}

func (h *WordPressMigrationHandler) RunScheduled(c *gin.Context) {
	var planSetting models.Setting
	if err := h.DB.Where("key = ?", "wp_migration.plan").First(&planSetting).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "پلن یافت نشد"})
		return
	}
	var payload struct {
		Config struct {
			SiteURL        string `json:"siteUrl"`
			APIKey         string `json:"apiKey"`
			ConsumerKey    string `json:"consumerKey"`
			ConsumerSecret string `json:"consumerSecret"`
		} `json:"config"`
		Stages []MigrationStage `json:"stages"`
	}
	if err := json.Unmarshal([]byte(planSetting.Value), &payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "پلن نامعتبر است"})
		return
	}
	var stateSetting models.Setting
	if err := h.DB.Where("key = ?", "wp_migration.plan_state").First(&stateSetting).Error; err != nil {
		stateSetting = models.Setting{Key: "wp_migration.plan_state", Value: `{"nextIndex":0}`, Category: "wp_migration", Type: "json"}
		_ = h.DB.Create(&stateSetting).Error
	}
	var st struct {
		NextIndex int `json:"nextIndex"`
	}
	_ = json.Unmarshal([]byte(stateSetting.Value), &st)
	if st.NextIndex >= len(payload.Stages) {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "همه مراحل اجرا شده‌اند"})
		return
	}
	stage := payload.Stages[st.NextIndex]
	if stage.Time != "" && time.Now().Format("15:04") < stage.Time {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "زمان اجرای این مرحله هنوز نرسیده"})
		return
	}
	uow := unitofwork.NewUnitOfWorkFactory(h.DB).Create()
	if err := uow.BeginTx(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	defer uow.Rollback()
	results := map[string]int{"imported": 0, "skipped": 0, "failed": 0}
	if stage.Products && payload.Config.ConsumerKey != "" && payload.Config.ConsumerSecret != "" {
		items, _ := h.Service.FetchWCProductsPage(payload.Config.SiteURL, payload.Config.ConsumerKey, payload.Config.ConsumerSecret, 1, max(1, stage.BatchSize))
		for _, it := range items {
			created, err := h.Service.UpsertProductFromWC(
				uow,
				it,
				stage.Categories,
				stage.Brands,
				stage.Attributes,
				stage.RedirectProducts,
				stage.RedirectCode,
				stage.DraftMode,
				stage.IncludeSlugInURL,
				stage.TransferSEOfromYoast,
				stage.TransferSEOfromRankMath,
				stage.TransferPrices,
				stage.TransferStock,
				stage.TransferVariants,
				stage.TransferDescriptions,
				stage.TransferShortDescription,
				stage.TransferImages,
				stage.TransferAttributesFields,
				stage.RedirectGroupName,
				payload.Config.SiteURL,
				payload.Config.ConsumerKey,
				payload.Config.ConsumerSecret,
			)
			if err != nil {
				results["failed"]++
				continue
			}
			if created {
				results["imported"]++
			} else {
				results["skipped"]++
			}
		}
	}
	if err := uow.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	st.NextIndex++
	b2, _ := json.Marshal(st)
	_ = h.DB.Model(&models.Setting{}).Where("key = ?", "wp_migration.plan_state").Update("value", string(b2)).Error
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "مرحله زمان‌بندی اجرا شد", "results": results, "nextIndex": st.NextIndex})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// --------------- Connection Test -----------------

func (h *WordPressMigrationHandler) TestConnection(c *gin.Context) {
	var body struct {
		SiteURL        string `json:"siteUrl"`
		APIKey         string `json:"apiKey"`
		ConsumerKey    string `json:"consumerKey"`
		ConsumerSecret string `json:"consumerSecret"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || body.SiteURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "آدرس سایت الزامی است"})
		return
	}
	if body.ConsumerKey != "" && body.ConsumerSecret != "" {
		count, err := h.Service.TestWooCustomers(body.SiteURL, body.ConsumerKey, body.ConsumerSecret)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "اتصال ووکامرس برقرار شد", "data": gin.H{"customersCount": count}})
		return
	}
	if body.APIKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "برای تست وردپرس، کلید API لازم است یا از CK/CS استفاده کنید"})
		return
	}
	var out []map[string]any
	if err := h.Service.FetchTest(body.SiteURL, body.APIKey, &out); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "اتصال وردپرس برقرار شد", "data": gin.H{"postsCount": len(out)}})
}

// ----------------- Start / Logs / Abort ------------------

func (h *WordPressMigrationHandler) StartMigration(c *gin.Context) {
	var req struct {
		Config struct {
			SiteURL        string `json:"siteUrl"`
			APIKey         string `json:"apiKey"`
			ConsumerKey    string `json:"consumerKey"`
			ConsumerSecret string `json:"consumerSecret"`
		} `json:"config"`
		Options struct {
			Products                     bool   `json:"products"`
			Orders                       bool   `json:"orders"`
			WordpressUsers               bool   `json:"wordpressUsers"`
			DigitsUsers                  bool   `json:"digitsUsers"`
			Categories                   bool   `json:"categories"`
			Brands                       bool   `json:"brands"`
			Attributes                   bool   `json:"attributes"`
			ReplaceUsernameWithMobile    bool   `json:"replaceUsernameWithMobile"`
			SkipEmails                   bool   `json:"skipEmails"`
			IncludeAddresses             bool   `json:"includeAddresses"`
			RedirectProducts             bool   `json:"redirectProducts"`
			RedirectCode                 int    `json:"redirectCode"`
			RedirectGroupName            string `json:"redirectGroupName"`
			DraftMode                    bool   `json:"draftMode"`
			IncludeSlugInURL             bool   `json:"includeSlugInURL"`
			TransferSEOfromYoast         bool   `json:"transferSEOfromYoast"`
			TransferSEOfromRankMath      bool   `json:"transferSEOfromRankMath"`
			TransferPrices               bool   `json:"transferPrices"`
			TransferStock                bool   `json:"transferStock"`
			TransferVariants             bool   `json:"transferVariants"`
			TransferDescriptions         bool   `json:"transferDescriptions"`
			TransferShortDescription     bool   `json:"transferShortDescription"`
			TransferImages               bool   `json:"transferImages"`
			TransferAttributesFields     bool   `json:"transferAttributesFields"`
			CreateCategoriesFromProducts bool   `json:"createCategoriesFromProducts"`
		} `json:"options"`
		Advanced struct {
			Mode          string `json:"mode"`
			BatchSize     int    `json:"batchSize"`
			IncludeImages bool   `json:"includeImages"`
		} `json:"advanced"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "درخواست نامعتبر"})
		return
	}
	hasAPI := req.Config.APIKey != ""
	hasWoo := req.Config.ConsumerKey != "" && req.Config.ConsumerSecret != ""
	if req.Config.SiteURL == "" || (!hasAPI && !hasWoo) {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "تنظیمات اتصال ناقص است (API Key یا Consumer Key/Secret را وارد کنید)"})
		return
	}
	if !(req.Options.Products || req.Options.Orders || req.Options.WordpressUsers || req.Options.DigitsUsers || req.Options.Categories || req.Options.Brands || req.Options.Attributes) {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "حداقل یک گزینه برای انتقال انتخاب کنید"})
		return
	}
	// reset logs
	migrationLogsMu.Lock()
	migrationLogsBuf = nil
	migrationLogsMu.Unlock()

	addLog := func(message, level string) { log.Printf("[%s] %s", level, message); pushMigrationLog(level, message) }
	addLog("شروع فرآیند انتقال...", "INFO")
	addLog(fmt.Sprintf("Mode=%s, BatchSize=%d", req.Advanced.Mode, req.Advanced.BatchSize), "INFO")

	go func() {
		defer func() {
			if r := recover(); r != nil {
				addLog(fmt.Sprintf("panic: %v", r), "ERROR")
			}
			currentMigCancelMu.Lock()
			currentMigCancel = nil
			currentMigCancelMu.Unlock()
			addLog("=== انتقال تکمیل شد ===", "INFO")
		}()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Hour)
		currentMigCancelMu.Lock()
		currentMigCancel = cancel
		currentMigCancelMu.Unlock()
		defer cancel()

		factory := unitofwork.NewUnitOfWorkFactory(h.DB)

		// Attributes
		if req.Options.Attributes && req.Advanced.Mode != "dry-run" {
			uow := factory.Create()
			if err := uow.BeginTx(ctx); err == nil {
				var logs []string
				impA, impV, skp, fld, err2 := h.Service.ImportWCAttributes(ctx, uow, req.Config.SiteURL, req.Config.ConsumerKey, req.Config.ConsumerSecret, &logs, func(t, m string) { pushMigrationLog(t, m) })
				if err2 == nil {
					_ = uow.Commit()
				}
				addLog(fmt.Sprintf("ویژگی‌ها: %d ویژگی و %d مقدار انتقال یافت ( %d رد، %d ناموفق )", impA, impV, skp, fld), "SUCCESS")
			}
		}

		// Digits Users
		if req.Options.DigitsUsers && req.Advanced.Mode != "dry-run" && hasWoo {
			uow := factory.Create()
			if err := uow.BeginTx(ctx); err == nil {
				var logs []string
				imp, skp, fld, _ := h.Service.ImportDigitsUsers(ctx, uow, req.Config.SiteURL, req.Config.ConsumerKey, req.Config.ConsumerSecret, req.Advanced.BatchSize, req.Options.ReplaceUsernameWithMobile, req.Options.SkipEmails, req.Options.IncludeAddresses, &logs, func(t, m string) { pushMigrationLog(t, m) })
				_ = uow.Commit()
				addLog(fmt.Sprintf("کاربران دیجیتس: %d انتقال یافت، %d رد شد، %d ناموفق", imp, skp, fld), "SUCCESS")
			}
		} else if req.Options.DigitsUsers && !hasWoo {
			addLog("مرحله کاربران دیجیتس رد شد: کلیدهای WooCommerce تنظیم نشده‌اند", "WARNING")
		}

		// Orders
		if req.Options.Orders && req.Advanced.Mode != "dry-run" && hasWoo {
			uow := factory.Create()
			if err := uow.BeginTx(ctx); err == nil {
				var logs []string
				imp, skp, fld, err := h.Service.ImportWooOrders(ctx, uow, req.Config.SiteURL, req.Config.ConsumerKey, req.Config.ConsumerSecret, req.Advanced.BatchSize, req.Options.IncludeAddresses, false, &logs, func(t, m string) { pushMigrationLog(t, m) })
				if err != nil {
					_ = uow.Rollback()
					addLog("خطا در انتقال سفارشات: "+err.Error(), "ERROR")
				} else {
					_ = uow.Commit()
					addLog(fmt.Sprintf("سفارشات: %d انتقال یافت، %d رد شد، %d ناموفق", imp, skp, fld), "SUCCESS")
				}
			}
		} else if req.Options.Orders && !hasWoo {
			addLog("مرحله سفارشات رد شد: CK/CS تنظیم نشده‌اند", "WARNING")
		}

		// Products
		if req.Options.Products && req.Advanced.Mode != "dry-run" {
			if !hasWoo {
				addLog("مرحله محصولات رد شد: CK/CS تنظیم نشده‌اند", "WARNING")
				return
			}

			// دریافت تعداد کل محصولات برای محاسبه دقیق پیشرفت
			_, totalProducts, err := h.Service.FetchWCProductsPageWithMeta(
				req.Config.SiteURL,
				req.Config.ConsumerKey,
				req.Config.ConsumerSecret,
				1, // صفحه اول
				1, // فقط یک محصول برای گرفتن تعداد کل
			)
			if err != nil {
				addLog(fmt.Sprintf("خطا در دریافت تعداد کل محصولات: %v", err), "ERROR")
				return
			}

			if totalProducts == 0 {
				addLog("هیچ محصولی یافت نشد.", "WARNING")
				return
			}

			addLog(fmt.Sprintf("تعداد کل محصولات برای انتقال: %d", totalProducts), "INFO")

			page := 1
			perPage := req.Advanced.BatchSize
			if perPage <= 0 {
				perPage = 20
			}
			processedCount := 0

			for {
				items, err := h.Service.FetchWCProductsPage(req.Config.SiteURL, req.Config.ConsumerKey, req.Config.ConsumerSecret, page, perPage)
				if err != nil {
					addLog(fmt.Sprintf("خطا در دریافت محصولات صفحه %d: %v", page, err), "ERROR")
					break
				}
				if len(items) == 0 {
					break
				}
				addLog(fmt.Sprintf("در حال پردازش %d محصول از صفحه %d...", len(items), page), "INFO")
				for _, it := range items {
					processedCount++
					progressPercent := (processedCount * 100) / totalProducts
					addLog(fmt.Sprintf("در حال انتقال محصول %d/%d (%d%%): %s", processedCount, totalProducts, progressPercent, it.Name), "INFO")
					uow := factory.Create()
					if err := uow.BeginTx(ctx); err != nil {
						addLog("خطا در شروع تراکنش محصول: "+err.Error(), "ERROR")
						continue
					}
					created, err := h.Service.UpsertProductFromWC(
						uow, it,
						req.Options.Categories || req.Options.CreateCategoriesFromProducts,
						req.Options.Brands,
						req.Options.Attributes,
						req.Options.RedirectProducts,
						req.Options.RedirectCode,
						req.Options.DraftMode,
						req.Options.IncludeSlugInURL,
						req.Options.TransferSEOfromYoast,
						req.Options.TransferSEOfromRankMath,
						req.Options.TransferPrices,
						req.Options.TransferStock,
						req.Options.TransferVariants,
						req.Options.TransferDescriptions,
						req.Options.TransferShortDescription,
						req.Options.TransferImages,
						req.Options.TransferAttributesFields,
						req.Options.RedirectGroupName,
						req.Config.SiteURL,
						req.Config.ConsumerKey,
						req.Config.ConsumerSecret,
					)
					if err != nil {
						addLog(fmt.Sprintf("خطا در انتقال محصول '%s': %v", it.Name, err), "ERROR")
						_ = uow.Rollback()
						continue
					}
					if err := uow.Commit(); err != nil {
						addLog(fmt.Sprintf("خطا در commit محصول %s: %v", it.Name, err), "ERROR")
						// اگر commit مشکل داشت، rollback کن
						_ = uow.Rollback()
						continue
					}
					if created {
						addLog(fmt.Sprintf("محصول %s با موفقیت انتقال یافت", it.Name), "SUCCESS")
					} else {
						addLog(fmt.Sprintf("محصول %s قبلاً وجود دارد (رد شد)", it.Name), "INFO")
					}
				}
				page++
			}
			addLog(fmt.Sprintf("انتقال محصولات تکمیل شد. تعداد کل پردازش شده: %d", processedCount), "SUCCESS")
		}
	}()

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "فرآیند انتقال شروع شد. برای مشاهده پیشرفت، از بخش لاگ‌ها استفاده کنید.", "status": "started"})
}

// GetTotalProductsCount دریافت تعداد کل محصولات از سایت وردپرس
func (h *WordPressMigrationHandler) GetTotalProductsCount(c *gin.Context) {
	var req struct {
		Config struct {
			SiteURL        string `json:"siteUrl"`
			ConsumerKey    string `json:"consumerKey"`
			ConsumerSecret string `json:"consumerSecret"`
		} `json:"config"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "درخواست نامعتبر است"})
		return
	}

	if req.Config.SiteURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "آدرس سایت الزامی است"})
		return
	}

	// دریافت اولین صفحه برای گرفتن تعداد کل
	_, totalCount, err := h.Service.FetchWCProductsPageWithMeta(
		req.Config.SiteURL,
		req.Config.ConsumerKey,
		req.Config.ConsumerSecret,
		1, // صفحه اول
		1, // فقط یک محصول برای گرفتن تعداد کل
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا در دریافت تعداد محصولات: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"totalCount": totalCount,
		"message":    fmt.Sprintf("تعداد کل محصولات: %d", totalCount),
	})
}

func (h *WordPressMigrationHandler) GetMigrationLogs(c *gin.Context) {
	migrationLogsMu.Lock()
	logsCopy := make([]map[string]interface{}, len(migrationLogsBuf))
	copy(logsCopy, migrationLogsBuf)
	migrationLogsMu.Unlock()
	c.JSON(http.StatusOK, gin.H{"success": true, "logs": logsCopy})
}

func (h *WordPressMigrationHandler) AbortMigration(c *gin.Context) {
	currentMigCancelMu.Lock()
	if currentMigCancel != nil {
		currentMigCancel()
		pushMigrationLog("warning", "درخواست توقف توسط کاربر دریافت شد؛ عملیات متوقف می‌شود")
	}
	currentMigCancelMu.Unlock()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "درخواست توقف ارسال شد"})
}

// ExtractMetaFromWordPressProduct استخراج متا تگ‌ها از صفحه محصول وردپرس
func (h *WordPressMigrationHandler) ExtractMetaFromWordPressProduct(c *gin.Context) {
	var req struct {
		ProductURL string `json:"productUrl"`
		Options    struct {
			ExtractFromYoast     bool `json:"extractFromYoast"`
			ExtractFromRankMath  bool `json:"extractFromRankMath"`
			ExtractFromOpenGraph bool `json:"extractFromOpenGraph"`
			ExtractFromMetaTags  bool `json:"extractFromMetaTags"`
		} `json:"options"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.ProductURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "URL محصول الزامی است"})
		return
	}

	// بررسی معتبر بودن URL
	if _, err := url.Parse(req.ProductURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "URL نامعتبر است"})
		return
	}

	// استخراج متا تگ‌ها از صفحه محصول
	options := services.MetaExtractionOptions{
		ExtractFromYoast:     req.Options.ExtractFromYoast,
		ExtractFromRankMath:  req.Options.ExtractFromRankMath,
		ExtractFromOpenGraph: req.Options.ExtractFromOpenGraph,
		ExtractFromMetaTags:  req.Options.ExtractFromMetaTags,
	}
	metaData, err := h.Service.ExtractMetaFromProductPage(req.ProductURL, options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "خطا در استخراج متا تگ‌ها: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "متا تگ‌ها با موفقیت استخراج شدند",
		"data":    metaData,
	})
}
