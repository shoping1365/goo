package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"
)

// ProductPriceHandler encapsulates all logic related to reading & updating product pricing fields.
// Keeping it separate from the main ProductHandler respects SRP and results in clearer maintainability.
// Only price–related columns are exposed here; any non-pricing changes belong to other handlers.

type ProductPriceHandler struct {
	DB *gorm.DB
}

// NewProductPriceHandler returns a fresh ProductPriceHandler.
func NewProductPriceHandler(db *gorm.DB) *ProductPriceHandler {
	return &ProductPriceHandler{DB: db}
}

// -----------------------------------------------------------------------------
// READ – GET /product-prices/:id
// -----------------------------------------------------------------------------
// GetPrice returns only the price-specific columns of a single product.
func (h *ProductPriceHandler) GetPrice(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "محصول پیدا نشد", err.Error()))
		return
	}

	// دریافت پله‌های فروش ویژه
	var offers []models.ProductSpecialOffer
	_ = h.DB.Where("product_id = ?", product.ID).Order("sort_order ASC, id ASC").Find(&offers).Error

	c.JSON(http.StatusOK, gin.H{
		"id":                 product.ID,
		"price":              product.Price,
		"old_price":          product.OldPrice,
		"cost":               product.Cost,
		"profit":             product.Profit,
		"discount_percent":   product.DiscountPercent,
		"discount_amount":    product.DiscountAmount,
		"sale_price":         product.SalePrice,
		"sale_start_at":      product.SaleStartAt,
		"sale_end_at":        product.SaleEndAt,
		"special_offers":     offers,
		"disable_buy_button": product.DisableBuyButton,
		"call_for_price":     product.CallForPrice,
	})
}

// -----------------------------------------------------------------------------
// UPDATE – PUT /product-prices/:id
// -----------------------------------------------------------------------------
// updatePriceInput defines allowed fields for price updates. Pointers are used so that
// omitted fields stay untouched (partial update semantics).

type updatePriceInput struct {
	Price           *float64 `json:"price"`            // قیمت فروش
	OldPrice        *float64 `json:"old_price"`        // قیمت خط خورده
	Cost            *float64 `json:"cost"`             // هزینه خرید برای مالک فروشگاه
	DiscountPercent *float64 `json:"discount_percent"` // درصد تخفیف
	DiscountAmount  *float64 `json:"discount_amount"`  // مبلغ تخفیف
	SalePrice       *float64 `json:"sale_price"`       // قیمت ویژه (در بازهٔ زمان‌بندی‌شده)
	SaleStartAt     *string  `json:"sale_start_at"`    // ISO8601/RFC3339
	SaleEndAt       *string  `json:"sale_end_at"`      // ISO8601/RFC3339
	SpecialOffers   []struct {
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
		Sort     int     `json:"sort_order"`
	} `json:"special_offers"`
	DisableBuyButton *bool `json:"disable_buy_button"` // غیرفعال کردن دکمه خرید
	CallForPrice     *bool `json:"call_for_price"`     // تماس برای قیمت
}

// UpdatePrice persists new price information for a product while automatically
// recalculating dependent columns such as profit, discount_amount & discount_percent.
func (h *ProductPriceHandler) UpdatePrice(c *gin.Context) {
	// 1) Fetch product or 404
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr) // error ignored – non-numeric id already 404s in query below

	var product models.Product
	if err := h.DB.First(&product, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "محصول پیدا نشد", err.Error()))
		return
	}

	// 2) Bind input JSON
	var input updatePriceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("BIND_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}

	// 3) Selectively update fields (nil means keep current value)
	if input.Price != nil {
		product.Price = *input.Price
	}
	if input.OldPrice != nil {
		product.OldPrice = *input.OldPrice
	}
	if input.Cost != nil {
		product.Cost = *input.Cost
	}

	// Discount logic – precedence: explicit percent > explicit amount.
	if input.DiscountPercent != nil {
		product.DiscountPercent = *input.DiscountPercent
		product.DiscountAmount = (product.Price * product.DiscountPercent) / 100
	} else if input.DiscountAmount != nil {
		product.DiscountAmount = *input.DiscountAmount
		if product.Price != 0 {
			product.DiscountPercent = (product.DiscountAmount / product.Price) * 100
		}
	}

	// قیمت ویژه و زمان‌بندی
	if input.SalePrice != nil {
		product.SalePrice = *input.SalePrice
	}
	// Parse RFC3339 timestamps if provided
	parseTime := func(ptr *string) (*time.Time, error) {
		if ptr == nil || *ptr == "" {
			return nil, nil
		}
		t, err := time.Parse(time.RFC3339, *ptr)
		if err != nil {
			return nil, err
		}
		return &t, nil
	}
	if input.SaleStartAt != nil {
		if t, err := parseTime(input.SaleStartAt); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("BIND_ERROR", "فرمت تاریخ شروع نامعتبر است (RFC3339)", err.Error()))
			return
		} else {
			product.SaleStartAt = t
		}
	}
	if input.SaleEndAt != nil {
		if t, err := parseTime(input.SaleEndAt); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("BIND_ERROR", "فرمت تاریخ پایان نامعتبر است (RFC3339)", err.Error()))
			return
		} else {
			product.SaleEndAt = t
		}
	}

	// اعتبارسنجی: اگر هردو زمان تنظیم شده‌اند، پایان باید بعد از شروع باشد
	if product.SaleStartAt != nil && product.SaleEndAt != nil {
		if product.SaleEndAt.Before(*product.SaleStartAt) || product.SaleEndAt.Equal(*product.SaleStartAt) {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "تاریخ پایان باید بعد از تاریخ شروع باشد", nil))
			return
		}
	}

	// ذخیره طرح‌های فروش ویژه (ساده: حذف و درج مجدد)
	if input.SpecialOffers != nil {
		if err := h.DB.Where("product_id = ?", product.ID).Delete(&models.ProductSpecialOffer{}).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در حذف طرح‌های قبلی", err.Error()))
			return
		}
		for i, o := range input.SpecialOffers {
			if o.Price <= 0 || o.Quantity <= 0 {
				continue
			}
			rec := models.ProductSpecialOffer{ProductID: product.ID, Price: o.Price, Quantity: o.Quantity, SortOrder: o.Sort}
			// اگر sort ارسال نشد، از اندیس استفاده کن
			if rec.SortOrder == 0 {
				rec.SortOrder = i + 1
			}
			if err := h.DB.Create(&rec).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ثبت طرح ویژه", err.Error()))
				return
			}
		}
	}

	// 4) Update pricing settings
	if input.DisableBuyButton != nil {
		product.DisableBuyButton = *input.DisableBuyButton
	}
	if input.CallForPrice != nil {
		product.CallForPrice = *input.CallForPrice
	}

	// 5) Derived field – profit (اگر هیچ قیمتی ثبت نشده بود، صفر بماند)
	if product.Price == 0 && product.Cost == 0 {
		product.Profit = 0
	} else {
		product.Profit = product.Price - product.Cost
	}

	// 6) Persist
	if err := h.DB.Save(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ذخیره قیمت", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "قیمت با موفقیت بروزرسانی شد",
		"product": gin.H{
			"id":               product.ID,
			"price":            product.Price,
			"old_price":        product.OldPrice,
			"cost":             product.Cost,
			"profit":           product.Profit,
			"discount_percent": product.DiscountPercent,
			"discount_amount":  product.DiscountAmount,
		},
	})
}

// -----------------------------------------------------------------------------
// BULK UPDATE – POST /product-prices/bulk-update
// -----------------------------------------------------------------------------

type bulkPriceUpdateInput struct {
	ProductIDs []int   `json:"product_ids" binding:"required"`
	Action     string  `json:"action" binding:"required,oneof=increase decrease"` // increase or decrease
	Percent    float64 `json:"percent" binding:"required,min=0,max=1000"`
}

// BulkUpdatePrices handles bulk price updates for multiple products
func (h *ProductPriceHandler) BulkUpdatePrices(c *gin.Context) {
	var input bulkPriceUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("BIND_ERROR", "ورودی نامعتبر", err.Error()))
		return
	}

	// Validate percent based on action
	if input.Action == "decrease" && input.Percent > 100 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "درصد کاهش نمی‌تواند بیشتر از ۱۰۰ باشد", nil))
		return
	}

	var products []models.Product
	if err := h.DB.Where("id IN ?", input.ProductIDs).Find(&products).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت محصولات", err.Error()))
		return
	}

	if len(products) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", "هیچ محصولی یافت نشد", nil))
		return
	}

	// Calculate multiplier
	var multiplier float64
	if input.Action == "increase" {
		multiplier = 1 + (input.Percent / 100)
	} else {
		multiplier = 1 - (input.Percent / 100)
	}

	updatedCount := 0
	for _, product := range products {
		// Store original price for old_price assignment in decrease case
		originalPrice := product.Price

		// Apply price change
		product.Price = float64(int(product.Price * multiplier))

		// Handle special cases based on action
		if input.Action == "increase" {
			// For increase: remove discounted price if exists
			if product.DiscountAmount > 0 || product.DiscountPercent > 0 {
				product.DiscountAmount = 0
				product.DiscountPercent = 0
			}
		} else if input.Action == "decrease" {
			// For decrease: move current price to old_price and set new price
			product.OldPrice = originalPrice
		}

		// Recalculate profit
		product.Profit = product.Price - product.Cost

		// Save the product
		if err := h.DB.Save(&product).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در ذخیره محصول", err.Error()))
			return
		}
		updatedCount++
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "قیمت‌ها با موفقیت بروزرسانی شدند",
		"updated_count": updatedCount,
		"action":        input.Action,
		"percent":       input.Percent,
	})
}
