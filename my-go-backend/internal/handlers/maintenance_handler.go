package handlers

import (
	"net/http"
	"reflect"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MaintenanceHandler contains housekeeping endpoints (developer tools).
type MaintenanceHandler struct {
	DB *gorm.DB
}

func NewMaintenanceHandler(db *gorm.DB) *MaintenanceHandler {
	return &MaintenanceHandler{DB: db}
}

// PurgeSoftDeletesRequest represents JSON body for purge request.
type PurgeSoftDeletesRequest struct {
	OlderThanDays int      `json:"older_than_days"`
	Tables        []string `json:"tables"`
}

// PurgeSoftDeletes permanently removes soft-deleted rows older than given days.
func (h *MaintenanceHandler) PurgeSoftDeletes(c *gin.Context) {
	var req PurgeSoftDeletesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}
	if req.OlderThanDays <= 0 {
		req.OlderThanDays = 90
	}
	cutoff := time.Now().AddDate(0, 0, -req.OlderThanDays)

	allowed := modelFactories()
	if len(req.Tables) == 0 {
		// default to categories only
		req.Tables = []string{"categories"}
	}

	tx := h.DB.Unscoped()
	for _, t := range req.Tables {
		entry, ok := allowed[t]
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "جدول ناشناخته: "+t, nil))
			return
		}
		if err := tx.Where("deleted_at IS NOT NULL AND deleted_at < ?", cutoff).Delete(entry).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در پاک‌سازی جدول "+t, err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// ListSoftDeleted returns soft-deleted rows of a given table
func (h *MaintenanceHandler) ListSoftDeleted(c *gin.Context) {
	table := c.Query("table")
	allowed := modelFactories()

	if table == "" || table == "all" {
		// aggregate
		result := make(map[string]interface{})
		for t, mdl := range allowed {
			slicePtr := newSlice(mdl)
			if err := h.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(slicePtr).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت داده", err.Error()))
				return
			}
			result[t] = reflect.ValueOf(slicePtr).Elem().Interface()
		}
		c.JSON(http.StatusOK, result)
		return
	}

	mdl, ok := allowed[table]
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "جدول پشتیبانی نمی‌شود", nil))
		return
	}
	slicePtr := newSlice(mdl)
	if err := h.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(slicePtr).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت داده", err.Error()))
		return
	}
	c.JSON(http.StatusOK, reflect.ValueOf(slicePtr).Elem().Interface())
}

func newSlice(model interface{}) interface{} {
	t := reflect.TypeOf(model).Elem()
	sliceType := reflect.SliceOf(t)
	return reflect.New(sliceType).Interface()
}

type RestoreRequest struct {
	Table string `json:"table"`
	ID    uint   `json:"id"`
}

// RestoreSoftDeleted un-deletes a record
func (h *MaintenanceHandler) RestoreSoftDeleted(c *gin.Context) {
	var req RestoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}
	allowed := modelFactories()
	factory, ok := allowed[req.Table]
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "جدول پشتیبانی نمی‌شود", nil))
		return
	}
	if err := h.DB.Unscoped().Model(factory).Where("id = ?", req.ID).Update("deleted_at", nil).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در بازیابی رکورد", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

type HardDeleteRequest struct {
	Table string `json:"table"`
	ID    uint   `json:"id"`
}

// HardDelete permanently removes a soft-deleted record
func (h *MaintenanceHandler) HardDelete(c *gin.Context) {
	var req HardDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}
	allowed := modelFactories()
	mdl, ok := allowed[req.Table]
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "جدول پشتیبانی نمی‌شود", nil))
		return
	}
	if err := h.DB.Unscoped().Where("id = ?", req.ID).Delete(mdl).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف رکورد", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func modelFactories() map[string]interface{} {
	return map[string]interface{}{
		"categories":             &models.Category{},
		"products":               &models.Product{},
		"product_images":         &models.ProductImage{},
		"customers":              &models.Customer{},
		"media_files":            &models.MediaFile{},
		"reviews":                &models.Review{},
		"review_images":          &models.ReviewImage{},
		"category_brand_pages":   &models.CategoryBrandPage{},
		"product_variants":       &models.ProductVariant{},
		"product_specifications": &models.ProductSpecification{},
		"product_videos":         &models.ProductVideo{},
		"brands":                 &models.Brand{},
	}
}

// TestSimpleEndpoint تست ساده برای بررسی عملکرد API
func TestSimpleEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "API در حال کار است",
		"time":    time.Now().Format("2006-01-02 15:04:05"),
		"version": "1.0.0",
	})
}

// GetSMSGatewaysTest تست دریافت لیست درگاه‌های پیامک
func GetSMSGatewaysTest(c *gin.Context) {
	// دریافت دیتابیس از context
	dbInterface, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "دیتابیس در دسترس نیست"})
		return
	}

	db, ok := dbInterface.(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در دسترسی به دیتابیس"})
		return
	}

	// دریافت تمام درگاه‌های پیامک
	var gateways []models.SMSGateway
	if err := db.Find(&gateways).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "خطا در دریافت درگاه‌ها"})
		return
	}

	// ساخت پاسخ تستی
	response := gin.H{
		"status": "success",
		"data": gin.H{
			"gateways":  gateways,
			"count":     len(gateways),
			"test_time": time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	c.JSON(http.StatusOK, response)
}
