package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DigikalaImportHandler handler برای import از دیجی‌کالا
type DigikalaImportHandler struct {
	DB            *gorm.DB
	ImportService *services.DigikalaImportService
	ImportRepo    repository.DigikalaImportRepositoryInterface
	LogRepo       repository.DigikalaImportLogRepositoryInterface
}

// NewDigikalaImportHandler ایجاد instance جدید
func NewDigikalaImportHandler(db *gorm.DB, importService *services.DigikalaImportService) *DigikalaImportHandler {
	importRepo := repository.NewDigikalaImportRepository(db)
	logRepo := repository.NewDigikalaImportLogRepository(db)

	return &DigikalaImportHandler{
		DB:            db,
		ImportService: importService,
		ImportRepo:    importRepo,
		LogRepo:       logRepo,
	}
}

// GetAvailableCategories دریافت دسته‌بندی‌های موجود در دیجی‌کالا
func (h *DigikalaImportHandler) GetAvailableCategories(c *gin.Context) {
	categories, err := h.ImportService.GetAvailableCategories()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DIGIKALA_API_ERROR", "خطا در دریافت دسته‌بندی‌ها از دیجی‌کالا", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "دسته‌بندی‌های دیجی‌کالا با موفقیت دریافت شدند",
		"data":    categories,
		"total":   len(categories),
	})
}

// SearchCategories جستجو در دسته‌بندی‌های دیجی‌کالا
func (h *DigikalaImportHandler) SearchCategories(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "پارامتر جستجو الزامی است", nil))
		return
	}

	categories, err := h.ImportService.SearchCategories(query)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DIGIKALA_SEARCH_ERROR", "خطا در جستجو دسته‌بندی‌ها", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "جستجو با موفقیت انجام شد",
		"data":    categories,
		"total":   len(categories),
		"query":   query,
	})
}

// ImportCategoriesRequest درخواست import دسته‌بندی‌ها
type ImportCategoriesRequest struct {
	CategoryIDs      []string `json:"category_ids" binding:"required"`
	ParentCategoryID *uint    `json:"parent_category_id"`
}

// ImportCategories import دسته‌بندی‌های انتخابی
func (h *DigikalaImportHandler) ImportCategories(c *gin.Context) {
	var req ImportCategoriesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	if len(req.CategoryIDs) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "حداقل یک دسته‌بندی باید انتخاب شود", nil))
		return
	}

	imported, failed, err := h.ImportService.ImportCategories(c.Request.Context(), req.CategoryIDs, req.ParentCategoryID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("IMPORT_ERROR", "خطا در import دسته‌بندی‌ها", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "عملیات import با موفقیت انجام شد",
		"imported": imported,
		"failed":   failed,
		"total":    imported + failed,
	})
}

// GetCategoryDetails دریافت جزئیات یک دسته‌بندی خاص از دیجی‌کالا
func (h *DigikalaImportHandler) GetCategoryDetails(c *gin.Context) {
	categoryID := c.Param("id")
	if categoryID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه دسته‌بندی الزامی است", nil))
		return
	}

	categories, err := h.ImportService.FetchCategories(categoryID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DIGIKALA_API_ERROR", "خطا در دریافت جزئیات دسته‌بندی", err.Error()))
		return
	}

	if len(categories) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("CATEGORY_NOT_FOUND", "دسته‌بندی یافت نشد", nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "جزئیات دسته‌بندی با موفقیت دریافت شد",
		"data":    categories[0],
	})
}

// ImportSpecificCategories import دسته‌بندی‌های خاص (مثل موبایل، لپ‌تاپ و غیره)
func (h *DigikalaImportHandler) ImportSpecificCategories(c *gin.Context) {
	categoryType := c.Param("type")

	var categoryIDs []string
	var parentCategoryID *uint

	// دسته‌بندی‌های پیش‌تعریف شده دیجی‌کالا
	predefinedCategories := map[string][]string{
		"mobile":  {"1", "2", "3"},    // شناسه‌های واقعی دسته‌بندی موبایل
		"laptop":  {"10", "11", "12"}, // شناسه‌های واقعی دسته‌بندی لپ‌تاپ
		"fashion": {"20", "21", "22"}, // شناسه‌های واقعی دسته‌بندی مد و پوشاک
		"home":    {"30", "31", "32"}, // شناسه‌های واقعی دسته‌بندی خانه و آشپزخانه
		"beauty":  {"40", "41", "42"}, // شناسه‌های واقعی دسته‌بندی زیبایی
		"sport":   {"50", "51", "52"}, // شناسه‌های واقعی دسته‌بندی ورزش
		"book":    {"60", "61", "62"}, // شناسه‌های واقعی دسته‌بندی کتاب
		"toy":     {"70", "71", "72"}, // شناسه‌های واقعی دسته‌بندی اسباب‌بازی
	}

	if ids, exists := predefinedCategories[categoryType]; exists {
		categoryIDs = ids
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_CATEGORY_TYPE", "نوع دسته‌بندی نامعتبر", map[string]interface{}{
			"available_types": []string{"mobile", "laptop", "fashion", "home", "beauty", "sport", "book", "toy"},
		}))
		return
	}

	// دریافت parent_category_id از query params
	if parentIDStr := c.Query("parent_id"); parentIDStr != "" {
		if parentID, err := strconv.ParseUint(parentIDStr, 10, 32); err == nil {
			parentIDUint := uint(parentID)
			parentCategoryID = &parentIDUint
		}
	}

	imported, failed, err := h.ImportService.ImportCategories(c.Request.Context(), categoryIDs, parentCategoryID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("IMPORT_ERROR", "خطا در import دسته‌بندی‌ها", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":       true,
		"message":       fmt.Sprintf("دسته‌بندی‌های %s با موفقیت import شدند", categoryType),
		"category_type": categoryType,
		"imported":      imported,
		"failed":        failed,
		"total":         imported + failed,
	})
}

// GetImportStatus وضعیت import های انجام شده
func (h *DigikalaImportHandler) GetImportStatus(c *gin.Context) {
	// بررسی تعداد دسته‌بندی‌های import شده از دیجی‌کالا (بر اساس description)
	var count int64
	h.DB.Table("categories").Where("description LIKE ?", "%Import شده از دیجی‌کالا%").Count(&count)

	// دریافت آخرین import ها
	var recentImports []map[string]interface{}
	h.DB.Table("categories").
		Select("name, slug, created_at, description").
		Where("description LIKE ?", "%Import شده از دیجی‌کالا%").
		Order("created_at DESC").
		Limit(10).
		Find(&recentImports)

	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"message":        "وضعیت import با موفقیت دریافت شد",
		"total_imported": count,
		"recent_imports": recentImports,
	})
}

// GetStats دریافت آمار کلی import ها
func (h *DigikalaImportHandler) GetStats(c *gin.Context) {
	stats, err := h.ImportRepo.GetStats()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("STATS_ERROR", "خطا در دریافت آمار", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "آمار با موفقیت دریافت شد",
		"data":    stats,
	})
}

// GetImportHistory دریافت تاریخچه import ها
func (h *DigikalaImportHandler) GetImportHistory(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	imports, err := h.ImportRepo.List(limit, offset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("HISTORY_ERROR", "خطا در دریافت تاریخچه", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "تاریخچه با موفقیت دریافت شد",
		"data":    imports,
		"limit":   limit,
		"offset":  offset,
	})
}

// ValidateURLRequest درخواست اعتبارسنجی URL
type ValidateURLRequest struct {
	URL string `json:"url" binding:"required"`
}

// ValidateURL اعتبارسنجی URL دیجی‌کالا
func (h *DigikalaImportHandler) ValidateURL(c *gin.Context) {
	var req ValidateURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	// بررسی فرمت URL دیجی‌کالا
	isValid, categoryTitle, categoryID := h.ImportService.ValidateDigikalaURL(req.URL)

	if !isValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_URL", "URL نامعتبر است", map[string]interface{}{
			"url":             req.URL,
			"expected_format": "https://www.digikala.com/search/category-*/",
		}))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"message":        "URL معتبر است",
		"url":            req.URL,
		"category_title": categoryTitle,
		"category_id":    categoryID,
		"is_valid":       true,
	})
}

// StartImportRequest درخواست شروع import
type StartImportRequest struct {
	CategoryURL      string      `json:"category_url" binding:"required"`
	ItemsPerMinute   int         `json:"items_per_minute"`
	MaxItems         int         `json:"max_items"`
	MaxProducts      int         `json:"max_products"`
	SkipExisting     bool        `json:"skip_existing"`
	ImportImages     bool        `json:"import_images"`
	TargetCategoryID interface{} `json:"target_category_id"` // Accept both string and int

	// Legacy support
	Settings struct {
		ItemsPerMinute int  `json:"items_per_minute"`
		MaxProducts    int  `json:"max_products"`
		SkipExisting   bool `json:"skip_existing"`
	} `json:"settings"`
}

// StartImport شروع عملیات import جدید
func (h *DigikalaImportHandler) StartImport(c *gin.Context) {
	var req StartImportRequest

	// Log request body for debugging
	bodyBytes, _ := c.GetRawData()
	fmt.Printf("📥 Request Body: %s\n", string(bodyBytes))

	// بازگردانی body برای binding
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("❌ Binding Error: %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	fmt.Printf("✅ Parsed Request: %+v\n", req)

	// پشتیبانی از فرمت‌های مختلف
	maxProducts := req.MaxProducts
	if maxProducts == 0 {
		maxProducts = req.MaxItems
	}
	if maxProducts == 0 {
		maxProducts = req.Settings.MaxProducts
	}
	if maxProducts == 0 {
		maxProducts = 10 // مقدار پیش‌فرض
	}

	itemsPerMinute := req.ItemsPerMinute
	if itemsPerMinute == 0 {
		itemsPerMinute = req.Settings.ItemsPerMinute
	}
	if itemsPerMinute == 0 {
		itemsPerMinute = 10 // مقدار پیش‌فرض
	}

	skipExisting := req.SkipExisting || req.Settings.SkipExisting

	// اعتبارسنجی URL
	isValid, categoryTitle, categoryID := h.ImportService.ValidateDigikalaURL(req.CategoryURL)
	if !isValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_URL", "URL نامعتبر است", nil))
		return
	}

	// ایجاد رکورد import جدید
	importRecord := &models.DigikalaImport{
		CategoryURL:      req.CategoryURL,
		CategoryTitle:    categoryTitle,
		Status:           models.ImportStatuses.Pending,
		TotalProducts:    0,
		ImportedProducts: 0,
		FailedProducts:   0,
		Progress:         0,
	}

	if err := h.ImportRepo.Create(importRecord); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("CREATE_ERROR", "خطا در ایجاد رکورد import", err.Error()))
		return
	}

	// شروع عملیات import در پس‌زمینه (async)
	go func() {
		h.startImportProcess(importRecord.ID, maxProducts, itemsPerMinute, skipExisting)
	}()

	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"message":        "عملیات import شروع شد",
		"import_id":      importRecord.ID,
		"category_title": categoryTitle,
		"category_id":    categoryID,
	})
}

// startImportProcess اجرای عملیات import در پس‌زمینه
func (h *DigikalaImportHandler) startImportProcess(importID uint, maxProducts int, itemsPerMinute int, skipExisting bool) {
	// به‌روزرسانی وضعیت به in_progress
	importRecord, err := h.ImportRepo.GetByID(importID)
	if err != nil {
		return
	}

	importRecord.Status = models.ImportStatuses.InProgress
	now := time.Now()
	importRecord.StartedAt = &now
	h.ImportRepo.Update(importRecord)

	// اضافه کردن لاگ
	h.LogRepo.Create(&models.DigikalaImportLog{
		ImportID: importID,
		Level:    models.LogLevels.Info,
		Message:  "عملیات import شروع شد",
	})

	// ⭐ ایجاد Product Scraper Service
	productRepo := repository.NewProductRepository(h.DB)
	categoryRepo := repository.NewCategoryRepository(h.DB)
	scraperService := services.NewDigikalaProductScraperService(
		h.DB,
		productRepo,
		categoryRepo,
		h.LogRepo,
	)

	// تنظیم import ID برای logging
	scraperService.SetImportID(importID)

	// تنظیم settings
	scraperSettings := services.ImportSettings{
		ItemsPerMinute: itemsPerMinute,
		MaxProducts:    maxProducts,
		SkipExisting:   skipExisting,
		ImportImages:   true,
		ImportSpecs:    true,
		Delay:          time.Duration(60/itemsPerMinute) * time.Second,
	}

	if scraperSettings.ItemsPerMinute == 0 {
		scraperSettings.ItemsPerMinute = 10
		scraperSettings.Delay = 6 * time.Second
	}

	scraperService.SetSettings(scraperSettings)

	// اجرای import واقعی
	ctx := context.Background()

	// Callback برای به‌روزرسانی progress
	progressCallback := func(imported, failed, skipped, total int) {
		importRecord.TotalProducts = total
		importRecord.ImportedProducts = imported
		importRecord.FailedProducts = failed
		if total > 0 {
			importRecord.Progress = (float64(imported+failed+skipped) / float64(total)) * 100
		}
		h.ImportRepo.Update(importRecord)
	}

	imported, failed, skipped, err := scraperService.ImportProductsFromCategory(
		ctx,
		importRecord.CategoryURL,
		scraperSettings.MaxProducts,
		progressCallback,
	)

	// به‌روزرسانی نهایی رکورد import با نتایج
	importRecord.TotalProducts = imported + failed + skipped
	importRecord.ImportedProducts = imported
	importRecord.FailedProducts = failed

	if err != nil {
		importRecord.Status = models.ImportStatuses.Failed
		h.LogRepo.Create(&models.DigikalaImportLog{
			ImportID: importID,
			Level:    models.LogLevels.Error,
			Message:  fmt.Sprintf("خطا در import: %v", err),
		})
	} else {
		importRecord.Status = models.ImportStatuses.Completed
		completedAt := time.Now()
		importRecord.CompletedAt = &completedAt
		importRecord.Progress = 100
		h.LogRepo.Create(&models.DigikalaImportLog{
			ImportID: importID,
			Level:    models.LogLevels.Info,
			Message:  fmt.Sprintf("عملیات import با موفقیت تکمیل شد: %d موفق, %d ناموفق, %d رد شده", imported, failed, skipped),
		})
	}

	h.ImportRepo.Update(importRecord)
}

// GetImportProgress دریافت پیشرفت import
func (h *DigikalaImportHandler) GetImportProgress(c *gin.Context) {
	importIDStr := c.Param("id")
	importID, err := strconv.ParseUint(importIDStr, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه import نامعتبر", nil))
		return
	}

	importRecord, err := h.ImportRepo.GetByID(uint(importID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("IMPORT_NOT_FOUND", "import یافت نشد", nil))
		return
	}

	// دریافت لاگ‌های مربوط به این import
	logs, err := h.LogRepo.GetByImportID(uint(importID))
	if err != nil {
		logs = []models.DigikalaImportLog{} // در صورت خطا، آرایه خالی برگردان
	}

	// محاسبه سرعت و ETA
	var speed int
	var eta *int
	if importRecord.StartedAt != nil && importRecord.ImportedProducts > 0 {
		elapsed := time.Since(*importRecord.StartedAt).Seconds()
		if elapsed > 0 {
			speed = int(float64(importRecord.ImportedProducts) / elapsed)
			if speed > 0 && importRecord.TotalProducts > 0 {
				remaining := importRecord.TotalProducts - importRecord.ImportedProducts
				etaValue := remaining / speed
				eta = &etaValue
			}
		}
	}

	// تبدیل به فرمت frontend
	c.JSON(http.StatusOK, gin.H{
		"processed":  importRecord.ImportedProducts,
		"total":      importRecord.TotalProducts,
		"progress":   importRecord.Progress,
		"successful": importRecord.ImportedProducts,
		"failed":     importRecord.FailedProducts,
		"skipped":    0, // فعلاً ندارم
		"speed":      speed,
		"eta":        eta,
		"status":     importRecord.Status,
		"logs":       logs,
	})
}

// CancelImport لغو import
func (h *DigikalaImportHandler) CancelImport(c *gin.Context) {
	importIDStr := c.Param("id")
	importID, err := strconv.ParseUint(importIDStr, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه import نامعتبر", nil))
		return
	}

	importRecord, err := h.ImportRepo.GetByID(uint(importID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("IMPORT_NOT_FOUND", "import یافت نشد", nil))
		return
	}

	if importRecord.Status != models.ImportStatuses.InProgress {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_STATUS", "فقط import های در حال انجام قابل لغو هستند", nil))
		return
	}

	importRecord.Status = models.ImportStatuses.Cancelled
	h.ImportRepo.Update(importRecord)

	// اضافه کردن لاگ لغو
	h.LogRepo.Create(&models.DigikalaImportLog{
		ImportID: uint(importID),
		Level:    models.LogLevels.Warning,
		Message:  "عملیات import لغو شد",
	})

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"message":   "import با موفقیت لغو شد",
		"import_id": importID,
	})
}

// RetryImportRequest درخواست تکرار import
type RetryImportRequest struct {
	Settings struct {
		ItemsPerMinute int  `json:"items_per_minute"`
		MaxProducts    int  `json:"max_products"`
		SkipExisting   bool `json:"skip_existing"`
	} `json:"settings"`
}

// RetryImport تکرار import ناموفق
func (h *DigikalaImportHandler) RetryImport(c *gin.Context) {
	importIDStr := c.Param("id")
	importID, err := strconv.ParseUint(importIDStr, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_ID", "شناسه import نامعتبر", nil))
		return
	}

	var req RetryImportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های ورودی نامعتبر", err.Error()))
		return
	}

	importRecord, err := h.ImportRepo.GetByID(uint(importID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("IMPORT_NOT_FOUND", "import یافت نشد", nil))
		return
	}

	if importRecord.Status != models.ImportStatuses.Failed {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("INVALID_STATUS", "فقط import های ناموفق قابل تکرار هستند", nil))
		return
	}

	// ریست کردن وضعیت برای تکرار
	importRecord.Status = models.ImportStatuses.Pending
	importRecord.ErrorMessage = nil
	importRecord.Progress = 0
	h.ImportRepo.Update(importRecord)

	// اضافه کردن لاگ تکرار
	h.LogRepo.Create(&models.DigikalaImportLog{
		ImportID: uint(importID),
		Level:    models.LogLevels.Info,
		Message:  "عملیات import مجدداً شروع شد",
	})

	// شروع عملیات import در پس‌زمینه
	go func() {
		// استفاده از مقادیر پیش‌فرض برای retry
		h.startImportProcess(uint(importID), 10, 10, true)
	}()

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"message":   "عملیات import مجدداً شروع شد",
		"import_id": importID,
	})
}
