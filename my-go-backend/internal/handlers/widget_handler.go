package handlers

import (
	"fmt"
	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// WidgetHandler ساختار هندلر برای مدیریت ویجت‌ها
type WidgetHandler struct {
	DB *gorm.DB
}

// NewWidgetHandler ایجاد نمونه جدید از هندلر ویجت
func NewWidgetHandler(db *gorm.DB) *WidgetHandler {
	return &WidgetHandler{DB: db}
}

// CreateWidget ایجاد ابزارک جدید
func (h *WidgetHandler) CreateWidget(c *gin.Context) {
	var widget models.Widget

	if err := c.ShouldBindJSON(&widget); err != nil {
		utils.BadRequest(c, "داده‌های ورودی نامعتبر", err.Error())
		return
	}

	// اعتبارسنجی
	if err := widget.Validate(); err != nil {
		utils.BadRequest(c, "خطا در اعتبارسنجی", err.Error())
		return
	}

	// تولید کد یونیک
	widget.Code = h.generateUniqueCode(widget.Type)

	// تنظیم زمان
	widget.CreatedAt = time.Now()
	widget.UpdatedAt = time.Now()

	// ذخیره در دیتابیس
	if err := h.DB.Create(&widget).Error; err != nil {
		utils.InternalServerError(c, "خطا در ذخیره‌سازی", err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ابزارک با موفقیت ایجاد شد",
		"data":    widget,
	})
}

// GetWidgets دریافت تمام ابزارک‌ها
func (h *WidgetHandler) GetWidgets(c *gin.Context) {
	var widgets []models.Widget

	// فیلتر بر اساس صفحه
	page := c.DefaultQuery("page", "")
	status := c.DefaultQuery("status", "")

	query := h.DB.Model(&models.Widget{})

	if page != "" {
		query = query.Where("page = ?", page)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	// مرتب‌سازی بر اساس sort_order
	if err := query.Order("sort_order ASC, created_at DESC").Find(&widgets).Error; err != nil {
		utils.InternalServerError(c, "خطا در دریافت ابزارک‌ها", err.Error())
		return
	}

	// حذف تبدیل config به map و انتساب آن به مدل Widget
	// فقط هنگام ارسال خروجی به کلاینت، تبدیل انجام شود

	utils.SendSuccess(c, "لیست ابزارک‌ها", widgets)
}

// GetWidget دریافت ابزارک با ID
func (h *WidgetHandler) GetWidget(c *gin.Context) {
	id := c.Param("id")

	var widget models.Widget
	if err := h.DB.First(&widget, id).Error; err != nil {
		utils.NotFound(c, "ابزارک پیدا نشد", err.Error())
		return
	}

	// حذف تبدیل config به map و انتساب آن به مدل Widget
	// فقط هنگام ارسال خروجی به کلاینت، تبدیل انجام شود

	utils.SendSuccess(c, "اطلاعات ابزارک", widget)
}

// GetWidgetByCode دریافت ابزارک با کد
func (h *WidgetHandler) GetWidgetByCode(c *gin.Context) {
	code := c.Param("code")

	var widget models.Widget
	if err := h.DB.Where("code = ? AND status = ?", code, "active").First(&widget).Error; err != nil {
		utils.NotFound(c, "ابزارک پیدا نشد", err.Error())
		return
	}

	utils.SendSuccess(c, "اطلاعات ابزارک", widget)
}

// UpdateWidget به‌روزرسانی ابزارک
func (h *WidgetHandler) UpdateWidget(c *gin.Context) {
	id := c.Param("id")

	var widget models.Widget
	if err := h.DB.First(&widget, id).Error; err != nil {
		utils.NotFound(c, "ابزارک پیدا نشد", err.Error())
		return
	}

	var updateData models.Widget
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.BadRequest(c, "داده‌های ورودی نامعتبر", err.Error())
		return
	}

	// اعتبارسنجی
	if updateData.Title != "" {
		widget.Title = updateData.Title
	}
	if updateData.Description != "" {
		widget.Description = updateData.Description
	}
	if updateData.Type != "" {
		widget.Type = updateData.Type
	}
	if updateData.Status != "" {
		widget.Status = updateData.Status
	}
	if updateData.SortOrder != 0 {
		widget.SortOrder = updateData.SortOrder
	}
	if updateData.Image != "" {
		widget.Image = updateData.Image
	}
	if updateData.Link != "" {
		widget.Link = updateData.Link
	}
	if updateData.Page != "" {
		widget.Page = updateData.Page
	}
	// Update show_on_mobile - this is a boolean so we need to handle it differently
	// We update it regardless of its value (true or false)
	widget.ShowOnMobile = updateData.ShowOnMobile
	if updateData.Config != nil {
		widget.Config = updateData.Config
	}

	// اعتبارسنجی مجدد
	if err := widget.Validate(); err != nil {
		utils.BadRequest(c, "خطا در اعتبارسنجی", err.Error())
		return
	}

	widget.UpdatedAt = time.Now()

	if err := h.DB.Save(&widget).Error; err != nil {
		utils.InternalServerError(c, "خطا در به‌روزرسانی", err.Error())
		return
	}

	utils.SendSuccess(c, "ابزارک با موفقیت به‌روزرسانی شد", widget)
}

// DeleteWidget حذف ابزارک
func (h *WidgetHandler) DeleteWidget(c *gin.Context) {
	id := c.Param("id")

	var widget models.Widget
	if err := h.DB.First(&widget, id).Error; err != nil {
		utils.NotFound(c, "ابزارک پیدا نشد", err.Error())
		return
	}

	if err := h.DB.Delete(&widget).Error; err != nil {
		utils.InternalServerError(c, "خطا در حذف", err.Error())
		return
	}

	utils.SendSuccess(c, "ابزارک با موفقیت حذف شد", nil)
}

// UpdateWidgetOrder به‌روزرسانی ترتیب ابزارک‌ها
func (h *WidgetHandler) UpdateWidgetOrder(c *gin.Context) {
	var orderData struct {
		Items []struct {
			ID    uint `json:"id"`
			Order int  `json:"order"`
		} `json:"items"`
	}

	if err := c.ShouldBindJSON(&orderData); err != nil {
		utils.BadRequest(c, "داده‌های ورودی نامعتبر", err.Error())
		return
	}

	// شروع transaction
	tx := h.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, item := range orderData.Items {
		if err := tx.Model(&models.Widget{}).Where("id = ?", item.ID).
			Update("sort_order", item.Order).Error; err != nil {
			tx.Rollback()
			utils.InternalServerError(c, "خطا در به‌روزرسانی ترتیب", err.Error())
			return
		}
	}

	tx.Commit()
	utils.SendSuccess(c, "ترتیب ابزارک‌ها با موفقیت به‌روزرسانی شد", nil)
}

// ToggleWidgetStatus تغییر وضعیت ابزارک
func (h *WidgetHandler) ToggleWidgetStatus(c *gin.Context) {
	id := c.Param("id")

	var widget models.Widget
	if err := h.DB.First(&widget, id).Error; err != nil {
		utils.NotFound(c, "ابزارک پیدا نشد", err.Error())
		return
	}

	// تغییر وضعیت
	if widget.Status == "active" {
		widget.Status = "inactive"
	} else {
		widget.Status = "active"
	}

	widget.UpdatedAt = time.Now()

	if err := h.DB.Save(&widget).Error; err != nil {
		utils.InternalServerError(c, "خطا در تغییر وضعیت", err.Error())
		return
	}

	utils.SendSuccess(c, "وضعیت ابزارک با موفقیت تغییر کرد", widget)
}

// GetWidgetsByPage دریافت ابزارک‌های یک صفحه خاص
func (h *WidgetHandler) GetWidgetsByPage(c *gin.Context) {
	page := c.Param("page")

	// ابتدا بررسی کنیم که آیا جدول widgets وجود دارد
	var tableExists bool
	h.DB.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'widgets')").Scan(&tableExists)

	if !tableExists {
		// اگر جدول وجود ندارد، آرایه خالی برگردانیم
		c.JSON(200, gin.H{
			"success": true,
			"message": "جدول ویجت‌ها موجود نیست",
			"data":    []interface{}{},
		})
		return
	}

	// تشخیص دستگاه از User-Agent
	userAgent := c.GetHeader("User-Agent")
	isMobile := false

	// بررسی User-Agent برای تشخیص موبایل
	if userAgent != "" {
		// لیست کلمات کلیدی موبایل
		mobileKeywords := []string{
			"Mobile", "Android", "iPhone", "iPad", "iPod",
			"BlackBerry", "Windows Phone", "Opera Mini", "IEMobile",
		}

		for _, keyword := range mobileKeywords {
			if len(userAgent) > 0 && strings.Contains(userAgent, keyword) {
				isMobile = true
				break
			}
		}
	}

	var widgets []models.Widget
	query := h.DB.Where("page = ? AND status = ?", page, "active")

	// اگر درخواست از موبایل است، فقط ویجت‌هایی که show_on_mobile = true هستند را برگردان
	if isMobile {
		query = query.Where("show_on_mobile = ?", true)
	}

	if err := query.Order("sort_order ASC").Find(&widgets).Error; err != nil {
		// اگر خطا در کوئری بود، لاگ کنیم و آرایه خالی برگردانیم
		fmt.Printf("Error fetching widgets for page %s: %v\n", page, err)
		c.JSON(200, gin.H{
			"success": true,
			"message": "خطا در دریافت ویجت‌ها",
			"data":    []interface{}{},
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "لیست ویجت‌های صفحه",
		"data":    widgets,
	})
}

// DuplicateWidget کپی کردن ابزارک
func (h *WidgetHandler) DuplicateWidget(c *gin.Context) {
	id := c.Param("id")

	var originalWidget models.Widget
	if err := h.DB.First(&originalWidget, id).Error; err != nil {
		utils.NotFound(c, "ابزارک پیدا نشد", err.Error())
		return
	}

	// ایجاد کپی
	newWidget := models.Widget{
		Title:       originalWidget.Title + " (کپی)",
		Description: originalWidget.Description,
		Type:        originalWidget.Type,
		Status:      "inactive", // کپی به صورت غیرفعال
		SortOrder:   originalWidget.SortOrder + 1,
		Image:       originalWidget.Image,
		Link:        originalWidget.Link,
		Page:        originalWidget.Page,
		Config:      originalWidget.Config,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// تولید کد یونیک جدید
	newWidget.Code = h.generateUniqueCode(newWidget.Type)

	if err := h.DB.Create(&newWidget).Error; err != nil {
		utils.InternalServerError(c, "خطا در کپی کردن", err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "ابزارک با موفقیت کپی شد",
		"data":    newWidget,
	})
}

// تولید کد یونیک برای ابزارک
func (h *WidgetHandler) generateUniqueCode(widgetType string) string {
	typeMap := map[string]string{
		"main-slider-side-banner": "SLIDER_MAIN",
		"single-slider-side":      "SLIDER_SINGLE",
		"full-slider":             "SLIDER_FULL",

		"products-box-bg": "PRODUCT_BG",
		"full-banner":     "BANNER_FULL",

		"double-banner":    "BANNER_DOUBLE",
		"triple-banner":    "BANNER_TRIPLE",
		"quad-banner":      "BANNER_QUAD",
		"penta-banner":     "BANNER_PENTA",
		"brands-slider":    "BRANDS_SLIDER",
		"services-slider":  "SERVICES_SLIDER",
		"category-box":     "CATEGORY_BOX",
		"blog-posts":       "BLOG_POSTS",
		"product-carousel": "PRODUCT_CAROUSEL",
	}

	prefix := typeMap[widgetType]
	if prefix == "" {
		prefix = "WIDGET"
	}

	// شمارش تعداد ویجت‌های موجود از این نوع
	var count int64
	h.DB.Model(&models.Widget{}).Where("type = ?", widgetType).Count(&count)

	// تولید کد یونیک با شماره بعدی
	code := fmt.Sprintf("%s_%03d", prefix, count+1)

	// بررسی اینکه آیا این کد قبلاً وجود دارد یا نه
	var exists int64
	h.DB.Model(&models.Widget{}).Where("code = ?", code).Count(&exists)

	// اگر کد تکراری بود، شماره را افزایش دهیم
	if exists > 0 {
		code = fmt.Sprintf("%s_%03d", prefix, count+2)
	}

	return code
}
