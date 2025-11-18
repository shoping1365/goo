package handlers

import (
	"errors"
	"log"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
	"net/http"

	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{DB: db}
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	// دریافت پارامترهای pagination از query string
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// دریافت پارامترهای جستجو و فیلتر
	search := strings.TrimSpace(c.Query("search"))
	categoryFilter := strings.TrimSpace(c.Query("filter_category"))
	statusFilter := strings.TrimSpace(c.Query("filter_status"))

	// محدودیت حداکثر limit برای جلوگیری از overload
	if limit > 500 {
		limit = 500
	}
	if limit < 1 {
		limit = 10
	}
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	var products []models.Product
	var total int64

	// ساخت query پایه
	query := h.DB.Preload("Category").
		Preload("Brand").
		Preload("Images", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC, id ASC").Limit(1)
		})

	// اعمال فیلتر جستجو
	if search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ? OR sku ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// اعمال فیلتر دسته‌بندی
	if categoryFilter != "" {
		query = query.Where("category_id = ?", categoryFilter)
	}

	// اعمال فیلتر وضعیت
	if statusFilter != "" {
		query = query.Where("status = ?", statusFilter)
	}

	// شمارش تعداد کل محصولات با فیلتر‌های اعمال شده
	if err := query.Model(&models.Product{}).Count(&total).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در شمارش محصولات", err.Error()))
		return
	}

	// دریافت محصولات با pagination
	if err := query.
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		Find(&products).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست محصولات", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       products,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	})
}

// ListPublicProducts دریافت محصولات منتشر شده برای صفحات عمومی (بدون احراز هویت)
func (h *ProductHandler) ListPublicProducts(c *gin.Context) {
	var products []models.Product
	// ابتدا محصولات با status "published" را جستجو کن
	if err := h.DB.Preload("Images").Preload("Category").Preload("Brand").Where("status = ?", "published").Find(&products).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست محصولات", err.Error()))
		return
	}

	// اگر محصولی با status "published" پیدا نشد، همه محصولات را برگردان
	if len(products) == 0 {
		if err := h.DB.Preload("Images").Preload("Category").Preload("Brand").Find(&products).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست محصولات", err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	identifier := c.Param("id")
	lookup := c.Query("lookup") // sku | slug | (empty => auto)

	var product models.Product

	// اگر صراحتاً lookup=sku ارسال شده باشد، ابتدا با SKU جستجو کن
	if lookup == "sku" {
		if err := h.DB.Preload("Images").Preload("Variants").Preload("Specifications").Preload("Videos").Preload("Category").Preload("Brand").Preload("SpecialOffers", func(db *gorm.DB) *gorm.DB { return db.Order("sort_order ASC, id ASC") }).Where("sku = ?", identifier).First(&product).Error; err == nil {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	// اگر صراحتاً lookup=slug ارسال شده باشد، ابتدا با Slug جستجو کن
	if lookup == "slug" {
		if err := h.DB.Preload("Images").Preload("Variants").Preload("Specifications").Preload("Videos").Preload("Category").Preload("Brand").Preload("SpecialOffers", func(db *gorm.DB) *gorm.DB { return db.Order("sort_order ASC, id ASC") }).Where("slug = ?", identifier).First(&product).Error; err == nil {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	// در غیر این صورت: ابتدا اگر عددی بود با ID جستجو کن
	if id, err := strconv.Atoi(identifier); err == nil {
		if err := h.DB.Preload("Images").Preload("Variants").Preload("Specifications").Preload("Videos").Preload("Category").Preload("Brand").Preload("SpecialOffers", func(db *gorm.DB) *gorm.DB { return db.Order("sort_order ASC, id ASC") }).First(&product, id).Error; err == nil {
			// Attach complements
			var complements []models.Product
			_ = h.DB.Raw(`
                SELECT p.*, (
                  SELECT pi.image_url FROM product_images pi
                  WHERE pi.product_id = p.id
                  ORDER BY pi.sort_order ASC, pi.id ASC
                  LIMIT 1
                ) AS image_url
                FROM product_complements pc
                JOIN products p ON p.id = pc.complement_product_id
                WHERE pc.product_id = ?
                ORDER BY COALESCE(pc.priority, 9999), p.id
            `, product.ID).Scan(&complements).Error
			product.ComplementProducts = complements
			c.JSON(http.StatusOK, product)
			return
		}
	}

	// اگر عددی نبود یا پیدا نشد، با SKU یا Slug جستجو کن
	if err := h.DB.Preload("Images").Preload("Variants").Preload("Specifications").Preload("Videos").Preload("Category").Preload("Brand").Preload("SpecialOffers", func(db *gorm.DB) *gorm.DB { return db.Order("sort_order ASC, id ASC") }).Where("sku = ? OR slug = ?", identifier, identifier).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	// Attach complements
	{
		var complements []models.Product
		_ = h.DB.Raw(`
            SELECT p.*, (
              SELECT pi.image_url FROM product_images pi
              WHERE pi.product_id = p.id
              ORDER BY pi.sort_order ASC, pi.id ASC
              LIMIT 1
            ) AS image_url
            FROM product_complements pc
            JOIN products p ON p.id = pc.complement_product_id
            WHERE pc.product_id = ?
            ORDER BY COALESCE(pc.priority, 9999), p.id
        `, product.ID).Scan(&complements).Error
		product.ComplementProducts = complements
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	// Expected payload from frontend for initial implementation
	type MediaImage struct {
		ID        int    `json:"id"`
		URL       string `json:"url"`
		Thumbnail string `json:"thumbnail"`
	}

	type createProductInput struct {
		Name            string       `json:"name"`
		Slug            string       `json:"slug"`
		SKU             string       `json:"sku"`
		NameEn          string       `json:"name_en"` // English name
		Description     string       `json:"description"`
		DescriptionEn   string       `json:"description_en"` // English description
		MetaDescription string       `json:"meta_description"`
		Price           float64      `json:"price"`
		SalePrice       *float64     `json:"sale_price"`
		StockQuantity   int          `json:"stock_quantity"`
		Weight          *float64     `json:"weight"`
		Dimensions      string       `json:"dimensions"`
		Status          string       `json:"status"`
		BrandID         *uint        `json:"brand_id"`
		CategoryID      *uint        `json:"category_id"`
		SeoTitle        string       `json:"seo_title"`
		Images          []MediaImage `json:"images"`
		SubCategoryID   *uint        `json:"sub_category_id"`
	}

	var input createProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// خطای اعتبارسنجی ورودی‌ها
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.NewErrorResponse(
			utils.ErrorTypeValidation,
			"خطای اعتبارسنجی",
			"اطلاعات ارسالی نامعتبر است",
			err.Error(),
			http.StatusBadRequest,
		))
		return
	}

	// Generate uuid
	newUUID := uuid.NewString()

	// --- Generate slug if not provided ---
	trimmedSlug := strings.TrimSpace(input.Slug)
	var newSlug string
	if trimmedSlug == "" {
		// اگر کاربر slug نداد، از نام (یا در نهایت UUID) بسازیم
		base := strings.TrimSpace(input.Name)
		if base == "" {
			base = newUUID
		}
		newSlug = slug.Make(base)
	} else {
		newSlug = slug.Make(trimmedSlug)
	}

	// اگر slug پس از sanitize خالی شد (مثلاً کاراکتر فارسی فقط)
	if newSlug == "" {
		newSlug = uuid.NewString()
	}

	// اطمینان از یکتا بودن slug با یک کوئری بهینه
	// به جای حلقهٔ چندباره: تعداد اسلاگ‌های هم‌نام/هم‌خانواده را می‌شماریم و بر اساس آن suffix می‌زنیم
	{
		baseSlug := newSlug
		var cnt int64
		if err := h.DB.Model(&models.Product{}).
			Where("slug = ? OR slug LIKE ?", baseSlug, baseSlug+"-%").
			Count(&cnt).Error; err == nil {
			if cnt > 0 {
				// مثال: base, base-2, base-3 ... → cnt شامل خود base هم می‌شود
				newSlug = slug.Make(baseSlug + "-" + strconv.Itoa(int(cnt+1)))
			}
		}
	}

	// Provide default name if empty
	finalName := strings.TrimSpace(input.Name)
	if finalName == "" {
		finalName = "محصول جدید"
	}

	// اعتبارسنجی تجاری حداقلی: نام و دسته‌بندی الزامی
	var fieldErrors []map[string]string
	if strings.TrimSpace(finalName) == "" {
		fieldErrors = append(fieldErrors, map[string]string{"field": "name", "message": "نام محصول الزامی است"})
	}
	if input.CategoryID == nil && input.SubCategoryID == nil {
		fieldErrors = append(fieldErrors, map[string]string{"field": "category_id", "message": "انتخاب دسته‌بندی الزامی است"})
	}
	if len(fieldErrors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":         "VALIDATION_ERROR",
			"type":         utils.ErrorTypeValidation,
			"title":        "خطای اعتبارسنجی",
			"message":      "اطلاعات وارد شده معتبر نیست",
			"user_message": "لطفاً فیلدهای الزامی را تکمیل کنید",
			"status_code":  http.StatusBadRequest,
			"errors":       fieldErrors,
		})
		return
	}

	// Create product
	product := models.Product{
		UUID: newUUID,
		// استفاده از slug محاسبه‌شده برای جلوگیری از تکراری/نامعتبر بودن
		Slug: newSlug,
		// استفاده از نام نهایی (در صورت خالی بودن ورودی)
		Name:            finalName,
		NameEn:          input.NameEn,
		Description:     input.Description,
		FullDescription: input.DescriptionEn,
		MetaDescription: input.MetaDescription,
		Price:           input.Price,
		StockQuantity:   input.StockQuantity,
		Status:          input.Status,
		BrandID:         input.BrandID,
		CategoryID:      input.CategoryID,
		SeoTitle:        input.SeoTitle,
	}

	// Handle optional fields safely
	if input.SalePrice != nil {
		product.SalePrice = *input.SalePrice
	}
	if input.Weight != nil {
		product.Weight = *input.Weight
	}

	if product.Status == "" {
		product.Status = "active"
	}

	// مقداردهی پیش‌فرض موجودی/وضعیت انبار
	if product.StockQuantity <= 0 {
		product.StockQuantity = 0
		product.StockStatus = "out_of_stock"
	} else if product.StockStatus == "" {
		product.StockStatus = "in_stock"
	}

	// اگر کاربر SKU صراحتاً داده، قبل از Create روی مدل ست کنیم تا از BeforeCreate عبور کند
	// و constraint یکتایی رعایت شود
	// توجه: اگر خالی باشد، BeforeCreate خودش تولید می‌کند
	// فیلد ورودی از بالا اضافه شده: input.SKU
	if strings.TrimSpace(input.SKU) != "" {
		product.SKU = strings.TrimSpace(input.SKU)
	}

	if err := h.DB.Create(&product).Error; err != nil {
		// مدیریت بهتر خطای یکتایی روی فیلد slug بدون تکیه بر نام کانسترینت
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && strings.Contains(strings.ToLower(pgErr.Error()), "slug") {
				c.AbortWithStatusJSON(http.StatusBadRequest, utils.NewErrorResponse(
					utils.ErrorTypeValidation,
					"URL تکراری",
					"آدرس یکتا (slug) تکراری است. لطفاً در تب سئو مقدار دیگری انتخاب کنید.",
					pgErr.Error(),
					http.StatusBadRequest,
				))
				return
			}
			if pgErr.Code == "23505" && strings.Contains(strings.ToLower(pgErr.Error()), "sku") {
				c.AbortWithStatusJSON(http.StatusBadRequest, utils.NewErrorResponse(
					utils.ErrorTypeValidation,
					"SKU تکراری",
					"کد محصول (SKU) تکراری است. لطفاً مقدار دیگری انتخاب کنید.",
					pgErr.Error(),
					http.StatusBadRequest,
				))
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.NewErrorResponse(
			utils.ErrorTypeDatabase,
			"خطای پایگاه داده",
			"خطا در ایجاد محصول",
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}

	// Save images if provided
	if len(input.Images) > 0 {
		for i, img := range input.Images {
			productImage := models.ProductImage{
				ProductID: product.ID,
				ImageURL:  img.URL,
				SortOrder: i,
			}
			h.DB.Create(&productImage)
		}
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.NewErrorResponse(
			utils.ErrorTypeNotFound,
			"یافت نشد",
			"محصول مورد نظر یافت نشد",
			err.Error(),
			http.StatusNotFound,
		))
		return
	}

	type MediaImage struct {
		ID        int    `json:"id"`
		URL       string `json:"url"`
		Thumbnail string `json:"thumbnail"`
	}

	type updateProductInput struct {
		Name             *string      `json:"name"`
		NameEn           *string      `json:"name_en"`
		Slug             *string      `json:"slug"`
		Description      *string      `json:"description"`
		FullDescription  *string      `json:"full_description"`
		MetaDescription  *string      `json:"meta_description"`
		Status           *string      `json:"status"`
		BrandID          *uint        `json:"brand_id"`
		CategoryID       *uint        `json:"category_id"`
		SeoTitle         *string      `json:"seo_title"`
		Images           []MediaImage `json:"images"`
		DisableBuyButton *bool        `json:"disable_buy_button"`
		CallForPrice     *bool        `json:"call_for_price"`
		SKU              *string      `json:"sku"`
		// VideoURL         *string      `json:"video_url"` // حذف: ویدیوها در جدول جدا ذخیره می‌شوند
	}

	var input updateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.NewErrorResponse(
			utils.ErrorTypeValidation,
			"خطای اعتبارسنجی",
			"اطلاعات ارسالی نامعتبر است",
			err.Error(),
			http.StatusBadRequest,
		))
		return
	}

	// Update product fields only if provided in payload
	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.NameEn != nil {
		product.NameEn = *input.NameEn
	}
	if input.Description != nil {
		product.Description = *input.Description
	}
	if input.FullDescription != nil {
		product.FullDescription = *input.FullDescription
	}
	if input.MetaDescription != nil {
		trimmed := strings.TrimSpace(*input.MetaDescription)
		if len(trimmed) > 255 {
			trimmed = trimmed[:255]
		}
		product.MetaDescription = trimmed
	}
	if input.Status != nil {
		product.Status = *input.Status
	}
	if input.BrandID != nil {
		product.BrandID = input.BrandID
	}
	if input.CategoryID != nil {
		product.CategoryID = input.CategoryID
	}
	if input.SeoTitle != nil {
		product.SeoTitle = *input.SeoTitle
	}

	// به‌روزرسانی slug در صورت ارسال
	if input.Slug != nil {
		candidate := strings.TrimSpace(*input.Slug)
		if candidate != "" {
			newSlug := slug.Make(candidate)
			if newSlug == "" {
				newSlug = uuid.NewString()
			}
			// اگر تغییر کرده، یکتا سازی کن
			if newSlug != product.Slug {
				baseSlug := newSlug
				var cnt int64
				if err := h.DB.Model(&models.Product{}).
					Where("(slug = ? OR slug LIKE ?) AND id <> ?", baseSlug, baseSlug+"-%", product.ID).
					Count(&cnt).Error; err == nil {
					if cnt > 0 {
						newSlug = slug.Make(baseSlug + "-" + strconv.Itoa(int(cnt+1)))
					}
				}
				product.Slug = newSlug
			}
		}
	}

	// Update pricing settings
	if input.DisableBuyButton != nil {
		product.DisableBuyButton = *input.DisableBuyButton
	}
	if input.CallForPrice != nil {
		product.CallForPrice = *input.CallForPrice
	}

	// فیلد video_url حذف شده است؛ ویدیوها از مسیر /product-videos مدیریت می‌شوند

	// Update SKU if provided and non-empty
	if input.SKU != nil {
		trimmed := strings.TrimSpace(*input.SKU)
		if trimmed != "" {
			product.SKU = trimmed
		}
	}

	if err := h.DB.Save(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.NewErrorResponse(
			utils.ErrorTypeDatabase,
			"خطای پایگاه داده",
			"خطا در بروزرسانی محصول",
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}

	// Update images: delete existing and create new ones
	if len(input.Images) > 0 {
		// Delete existing images
		h.DB.Where("product_id = ?", product.ID).Delete(&models.ProductImage{})

		// Create new images
		for i, img := range input.Images {
			productImage := models.ProductImage{
				ProductID: product.ID,
				ImageURL:  img.URL,
				SortOrder: i,
			}
			h.DB.Create(&productImage)
		}
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (h *ProductHandler) CheckSlug(c *gin.Context) {
	productSlug := c.Query("slug")
	if strings.TrimSpace(productSlug) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slug الزامی است"})
		return
	}
	var count int64
	if err := h.DB.Model(&models.Product{}).Where("slug = ?", productSlug).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"exists": count > 0, "slug": productSlug})
}

func (h *ProductHandler) GenerateSlug(c *gin.Context) {
	base := c.Query("base_slug")
	base = strings.TrimSpace(base)
	if base == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "base_slug الزامی است"})
		return
	}
	candidate := slug.Make(base)
	if candidate == "" {
		candidate = uuid.NewString()
	}
	// ensure uniqueness by suffixing -n
	i := 0
	final := candidate
	for {
		var count int64
		if err := h.DB.Model(&models.Product{}).Where("slug = ?", final).Count(&count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count == 0 {
			break
		}
		i++
		final = candidate + "-" + strconv.Itoa(i)
	}
	c.JSON(http.StatusOK, gin.H{"slug": final})
}

// BulkTransferProducts انتقال گروهی محصولات به دسته‌بندی دیگر
func (h *ProductHandler) BulkTransferProducts(c *gin.Context) {
	var req struct {
		ProductIDs   []uint `json:"product_ids" binding:"required,min=1"`
		ToCategoryID uint   `json:"to_category_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر: " + err.Error()})
		return
	}

	// ایجاد Unit of Work
	uowFactory := unitofwork.NewUnitOfWorkFactory(h.DB)
	uow := uowFactory.Create()

	// شروع تراکنش
	if err := uow.BeginTx(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در شروع تراکنش: " + err.Error()})
		return
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			if rollbackErr := uow.Rollback(); rollbackErr != nil {
				log.Printf("خطا در rollback تراکنش: %v", rollbackErr)
			}
		}
	}()

	// بررسی وجود دسته‌بندی مقصد
	targetCategory, err := uow.CategoryRepository().GetByID(c.Request.Context(), req.ToCategoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "دسته‌بندی مقصد یافت نشد"})
		return
	}

	// به‌روزرسانی محصولات
	updatedCount := 0
	for _, productID := range req.ProductIDs {
		product, err := uow.ProductRepository().GetByID(c.Request.Context(), productID)
		if err != nil {
			continue // رد کردن محصولات ناموجود
		}

		// به‌روزرسانی دسته‌بندی
		product.CategoryID = &req.ToCategoryID
		if err := uow.ProductRepository().Update(c.Request.Context(), product); err != nil {
			continue
		}
		updatedCount++
	}

	// Commit تراکنش
	if err := uow.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ذخیره تغییرات: " + err.Error()})
		return
	}

	committed = true

	c.JSON(http.StatusOK, gin.H{
		"message":           "انتقال محصولات با موفقیت انجام شد",
		"transferred_count": updatedCount,
		"requested_count":   len(req.ProductIDs),
		"to_category_id":    req.ToCategoryID,
		"to_category_name":  targetCategory.Name,
	})
}

// ListProductImages لیست تمام تصاویر محصولات را برمی‌گرداند
func (h *ProductHandler) ListProductImages(c *gin.Context) {
	var results []map[string]interface{}

	// دریافت تمام تصاویر محصولات به همراه اطلاعات محصول
	err := h.DB.Table("product_images").
		Select(`product_images.*, 
			products.name as product_name,
			products.slug as product_slug`).
		Joins("LEFT JOIN products ON products.id = product_images.product_id").
		Order("product_images.created_at DESC").
		Scan(&results).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت لیست تصاویر محصولات", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": results})
}
