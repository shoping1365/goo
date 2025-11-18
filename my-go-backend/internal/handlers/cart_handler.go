package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartHandler struct {
	db *gorm.DB
}

func NewCartHandler(db *gorm.DB) *CartHandler {
	return &CartHandler{db: db}
}

// GetCart دریافت سبد خرید کاربر
func (h *CartHandler) GetCart(c *gin.Context) {
	var userIDPtr *uint
	if v, ok := c.Get("user_id"); ok {
		id := v.(uint)
		userIDPtr = &id
	}

	sessionID := h.getSessionID(c)

	var cart models.Cart
	var err error

	if userIDPtr != nil {
		// بهینه‌سازی: فقط فیلدهای مورد نیاز را Select کنیم
		err = h.db.Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, cart_id, product_id, quantity, unit_price, discount_amount, final_price, is_next_purchase").
				Preload("Product", func(db *gorm.DB) *gorm.DB {
					return db.Select("id, name, price, image, stock_quantity, track_inventory")
				})
		}).Where("user_id = ?", *userIDPtr).First(&cart).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// اگر برای کاربر سبدی نبود، بر اساس session_id جستجو کن
			err = h.db.Preload("Items", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, cart_id, product_id, quantity, unit_price, discount_amount, final_price, is_next_purchase").
					Preload("Product", func(db *gorm.DB) *gorm.DB {
						return db.Select("id, name, price, image, stock_quantity, track_inventory")
					})
			}).Where("session_id = ?", sessionID).First(&cart).Error

			if err == nil {
				// مالکیت سبد مهمان را به کاربر بده
				h.db.Model(&cart).Update("user_id", userIDPtr)
			}
		}
	} else {
		// کاربر مهمان
		if sessionID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "شناسه نشست معتبر نیست"})
			return
		}
		err = h.db.Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, cart_id, product_id, quantity, unit_price, discount_amount, final_price, is_next_purchase").
				Preload("Product", func(db *gorm.DB) *gorm.DB {
					return db.Select("id, name, price, image, stock_quantity, track_inventory")
				})
		}).Where("session_id = ?", sessionID).First(&cart).Error
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// اگر سبد خرید وجود ندارد، یک سبد خرید جدید ایجاد می‌کنیم
			cart = models.Cart{
				SessionID: sessionID,
				Items:     []models.CartItem{},
			}
			h.db.Create(&cart)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "خطا در دریافت سبد خرید",
			})
			return
		}
	}

	// بهینه‌سازی: دریافت تصاویر محصولات به صورت batch
	if len(cart.Items) > 0 {
		productIDs := make([]uint, len(cart.Items))
		productMap := make(map[uint]*models.Product)

		for i, item := range cart.Items {
			productIDs[i] = item.ProductID
			productMap[item.ProductID] = &cart.Items[i].Product
		}

		// دریافت تصاویر اصلی همه محصولات در یک کوئری
		var productImages []struct {
			ProductID uint   `json:"product_id"`
			ImageURL  string `json:"image_url"`
		}

		h.db.Table("product_images").
			Select("product_images.product_id, media_files.file_path as image_url").
			Joins("JOIN media_files ON media_files.id = product_images.media_file_id").
			Where("product_images.product_id IN ? AND product_images.is_primary = true", productIDs).
			Where("media_files.deleted_at IS NULL").
			Scan(&productImages)

		// تنظیم تصاویر در محصولات
		imageMap := make(map[uint]string)
		for _, img := range productImages {
			imageMap[img.ProductID] = img.ImageURL
		}

		// اعمال تصاویر یا fallback
		for idx := range cart.Items {
			if imageURL, exists := imageMap[cart.Items[idx].ProductID]; exists && imageURL != "" {
				cart.Items[idx].Product.Image = imageURL
			} else {
				// Fallback به تصویر پیش‌فرض محصول
				if cart.Items[idx].Product.Image == "" {
					thumb := h.getProductThumbnail(int(cart.Items[idx].ProductID))
					if thumb != "" {
						cart.Items[idx].Product.Image = thumb
					}
				}
			}
		}
	}

	// محاسبه مجموع
	totalItems := 0
	totalPrice := 0.0
	for _, item := range cart.Items {
		if item.IsNextPurchase {
			continue // در محاسبه مجموع لحاظ نشود
		}
		totalItems += item.Quantity
		// استفاده از قیمت نهایی آیتم (پس از اعمال طرح ویژه)
		totalPrice += float64(item.Quantity) * item.FinalPrice
	}

	// اولین نمایش سبد: حذف آیتم‌های اتمام موجودی در حالت پیگیری موجودی
	// مطابق نیاز: اگر کاربر کالا را وارد سبد کرده باشد و موجودی تمام شود، در اولین نمایش سبد حذف شود
	// توجه: در فاز 1 رزروی انجام نمی‌شود
	var removedCount int
	{
		// بارگذاری وضعیت فعلی محصول هر آیتم
		var toRemoveIDs []uint
		for _, item := range cart.Items {
			// اگر محصول track_inventory فعال و موجودی کافی نیست ⇒ حذف
			if item.Product.TrackInventory && item.Product.StockQuantity <= 0 {
				toRemoveIDs = append(toRemoveIDs, item.ID)
			}
		}
		if len(toRemoveIDs) > 0 {
			h.db.Where("id IN ?", toRemoveIDs).Delete(&models.CartItem{})
			removedCount = len(toRemoveIDs)
			// بازخوانی اقلام پس از حذف برای پاسخ تمیز
			_ = h.db.Preload("Items", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, cart_id, product_id, quantity, unit_price, discount_amount, final_price, is_next_purchase").
					Preload("Product", func(db *gorm.DB) *gorm.DB {
						return db.Select("id, name, price, image, stock_quantity, track_inventory")
					})
			}).First(&cart, cart.ID).Error
		}
	}

	// تبدیل به پاسخ مناسب
	response := models.CartResponse{
		ID:         cart.ID,
		SessionID:  cart.SessionID,
		Items:      cart.Items,
		TotalItems: totalItems,
		TotalPrice: totalPrice,
		CreatedAt:  cart.CreatedAt,
		UpdatedAt:  cart.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"success":                    true,
		"data":                       response,
		"removed_out_of_stock_count": removedCount,
	})
}

// AddToCart افزودن محصول به سبد خرید
func (h *CartHandler) AddToCart(c *gin.Context) {
	var req models.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Cart Add - JSON Bind Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "داده‌های ورودی نامعتبر است",
		})
		return
	}

	fmt.Printf("Cart Add Request - ProductID: %d, Quantity: %d\n", req.ProductID, req.Quantity)

	// بررسی وجود محصول
	var product models.Product
	var err error
	if err = h.db.First(&product, req.ProductID).Error; err != nil {
		fmt.Printf("Product not found by ID %d: %v\n", req.ProductID, err)
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "محصول یافت نشد",
		})
		return
	}

	fmt.Printf("Found Product - ID: %d, SKU: %s, Stock: %d, Requested: %d\n",
		product.ID, product.SKU, product.StockQuantity, req.Quantity)

	// بررسی موجودی محصول – فقط زمانی که موجودی پیگیری می‌شود
	if product.TrackInventory && product.StockQuantity < req.Quantity {
		fmt.Printf("Stock check failed - Stock: %d, Requested: %d\n", product.StockQuantity, req.Quantity)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "موجودی محصول کافی نیست",
		})
		return
	}

	var userIDPtr *uint
	if v, ok := c.Get("user_id"); ok {
		id := v.(uint)
		userIDPtr = &id
	}

	sessionID := h.getSessionID(c)
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "شناسه نشست معتبر نیست",
		})
		return
	}

	// دریافت یا ایجاد سبد خرید
	var cart models.Cart
	if userIDPtr != nil {
		err = h.db.Where("user_id = ?", userIDPtr).First(&cart).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// ممکن است سبد مهمان با همین session وجود داشته باشد → مالکیت را منتقل کن
			err = h.db.Where("session_id = ?", sessionID).First(&cart).Error
			if err == nil {
				h.db.Model(&cart).Update("user_id", userIDPtr)
			}
		}
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// هنوز پیدا نشده: سبد جدید بساز
		cart = models.Cart{SessionID: sessionID}
		if userIDPtr != nil {
			cart.UserID = userIDPtr
		}
		h.db.Create(&cart)
		// ست‌کردن کوکی برای مهمان
		http.SetCookie(c.Writer, &http.Cookie{Name: "session_id", Value: sessionID, Path: "/", HttpOnly: true, SameSite: http.SameSiteLaxMode})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا در ایجاد سبد خرید"})
		return
	}

	// بررسی وجود محصول در سبد خرید
	var existingItem models.CartItem
	err = h.db.Where("cart_id = ? AND product_id = ?", cart.ID, uint(req.ProductID)).First(&existingItem).Error
	if err == nil {
		// اگر محصول قبلاً در سبد خرید وجود دارد، تعداد را افزایش می‌دهیم و قیمت نهایی را بر اساس طرح‌های ویژه بازمحاسبه می‌کنیم
		newQuantity := existingItem.Quantity + req.Quantity

		// بررسی مقدار جدید فقط برای محصولات با پیگیری موجودی
		if product.TrackInventory && newQuantity > product.StockQuantity {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "موجودی محصول کافی نیست",
			})
			return
		}

		// محاسبه قیمت نهایی بر اساس طرح‌های ویژه فعال
		finalUnitPrice := h.calculateUnitPriceWithSpecialOffers(&product, newQuantity)
		discountPerUnit := product.Price - finalUnitPrice

		h.db.Model(&existingItem).Updates(map[string]interface{}{
			"quantity":        newQuantity,
			"unit_price":      product.Price,
			"discount_amount": discountPerUnit,
			"final_price":     finalUnitPrice,
		})

	} else {
		// اگر محصول در سبد خرید نیست، آیتم جدید اضافه می‌کنیم با اعمال قیمت ویژه
		finalUnitPrice := h.calculateUnitPriceWithSpecialOffers(&product, req.Quantity)
		discountPerUnit := product.Price - finalUnitPrice
		newItem := models.CartItem{
			CartID:         cart.ID,
			ProductID:      uint(req.ProductID),
			Quantity:       req.Quantity,
			UnitPrice:      product.Price,
			DiscountAmount: discountPerUnit,
			FinalPrice:     finalUnitPrice,
		}
		h.db.Create(&newItem)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "محصول با موفقیت به سبد خرید اضافه شد",
	})
}

// UpdateCartItem به‌روزرسانی تعداد محصول در سبد خرید
func (h *CartHandler) UpdateCartItem(c *gin.Context) {
	var req models.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "داده‌های ورودی نامعتبر است",
		})
		return
	}

	// دریافت آیتم سبد خرید
	var cartItem models.CartItem
	if err := h.db.Preload("Product").First(&cartItem, req.CartItemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "آیتم سبد خرید یافت نشد",
		})
		return
	}

	// بررسی موجودی محصول (در حالت پیگیری موجودی)
	if cartItem.Product.TrackInventory && cartItem.Product.StockQuantity < req.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "موجودی محصول کافی نیست",
		})
		return
	}

	// به‌روزرسانی تعداد و قیمت نهایی بر اساس طرح‌های ویژه فعال
	finalUnitPrice := h.calculateUnitPriceWithSpecialOffers(&cartItem.Product, req.Quantity)
	discountPerUnit := cartItem.Product.Price - finalUnitPrice
	h.db.Model(&cartItem).Updates(map[string]interface{}{
		"quantity":        req.Quantity,
		"unit_price":      cartItem.Product.Price,
		"discount_amount": discountPerUnit,
		"final_price":     finalUnitPrice,
	})

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "سبد خرید با موفقیت به‌روزرسانی شد",
	})
}

// calculateUnitPriceWithSpecialOffers محاسبه قیمت نهایی هر واحد بر اساس طرح‌های ویژه فعال و بازه زمانی
// منطق: اگر بازه فعال باشد و پله‌ها تعریف شده باشند، میانگین وزنی قیمت‌ها برای کل تعداد درخواستی محاسبه می‌شود.
// پس از اتمام پله‌ها، قیمت به قیمت اصلی محصول برمی‌گردد.
func (h *CartHandler) calculateUnitPriceWithSpecialOffers(product *models.Product, quantity int) float64 {
	if quantity <= 0 {
		return product.Price
	}
	// بررسی بازه زمانی فروش ویژه
	now := time.Now()
	if product.SaleStartAt != nil {
		if now.Before(product.SaleStartAt.UTC()) { // هنوز شروع نشده
			return product.Price
		}
	}
	if product.SaleEndAt != nil {
		if now.After(product.SaleEndAt.UTC()) { // تمام شده
			return product.Price
		}
	}

	// دریافت پله‌ها
	var offers []models.ProductSpecialOffer
	if err := h.db.Where("product_id = ?", product.ID).Order("sort_order ASC, id ASC").Find(&offers).Error; err != nil || len(offers) == 0 {
		// در صورت نبود پله، اگر قیمت فروش ویژه (تک‌قیمتی) تنظیم شده باشد از آن استفاده کن
		if product.SalePrice > 0 {
			return product.SalePrice
		}
		return product.Price
	}

	remaining := quantity
	totalCost := 0.0
	for _, o := range offers {
		if o.Price <= 0 || o.Quantity <= 0 {
			continue
		}
		take := o.Quantity
		if take > remaining {
			take = remaining
		}
		if take > 0 {
			totalCost += float64(take) * o.Price
			remaining -= take
		}
		if remaining == 0 {
			break
		}
	}
	// اگر هنوز تعداد باقی مانده، با قیمت اصلی محاسبه می‌شود
	if remaining > 0 {
		totalCost += float64(remaining) * product.Price
	}
	avg := totalCost / float64(quantity)
	return avg
}

// RemoveFromCart حذف محصول از سبد خرید
func (h *CartHandler) RemoveFromCart(c *gin.Context) {
	var req models.RemoveFromCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "داده‌های ورودی نامعتبر است",
		})
		return
	}

	// دریافت آیتم از سبد خرید
	var cartItem models.CartItem
	if err := h.db.Preload("Product").First(&cartItem, req.CartItemID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "آیتم سبد خرید یافت نشد",
		})
		return
	}

	// سپس حذف آیتم
	result := h.db.Delete(&models.CartItem{}, req.CartItemID)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "آیتم سبد خرید یافت نشد",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "محصول با موفقیت از سبد خرید حذف شد",
	})
}

// ClearCart پاک کردن کامل سبد خرید
func (h *CartHandler) ClearCart(c *gin.Context) {
	sessionID := h.getSessionID(c)
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "شناسه نشست معتبر نیست",
		})
		return
	}

	var userIDPtr *uint
	if v, ok := c.Get("user_id"); ok {
		id := v.(uint)
		userIDPtr = &id
	}

	if userIDPtr != nil {
		h.db.Where("cart_id IN (SELECT id FROM carts WHERE user_id = ?)", *userIDPtr).Delete(&models.CartItem{})
	} else {
		// دریافت cart id بر اساس session
		var cart models.Cart
		if err := h.db.Where("session_id = ?", sessionID).First(&cart).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "سبد خرید یافت نشد"})
			return
		}
		h.db.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "سبد خرید با موفقیت پاک شد",
	})
}

// GetCartCount دریافت تعداد آیتم‌های سبد خرید
func (h *CartHandler) GetCartCount(c *gin.Context) {
	sessionID := h.getSessionID(c)
	if sessionID == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"count":   0,
		})
		return
	}

	var count int64
	h.db.Model(&models.CartItem{}).
		Joins("JOIN carts ON cart_items.cart_id = carts.id").
		Where("carts.session_id = ?", sessionID).
		Count(&count)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"count":   count,
	})
}

// TestAddToCart تست ساده برای debugging
func (h *CartHandler) TestAddToCart(c *gin.Context) {
	var req models.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Validation Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validation failed: " + err.Error(),
		})
		return
	}

	fmt.Printf("Test OK - ProductID: %d, Quantity: %d\n", req.ProductID, req.Quantity)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Test successful",
		"data":    req,
	})
}

// CreateSession ایجاد یک session ID جدید و ست‌کردن کوکی
func (h *CartHandler) CreateSession(c *gin.Context) {
	// اگر قبلاً کوکی session_id وجود دارد همان را برگردان
	if sid, err := c.Cookie("session_id"); err == nil && sid != "" {
		c.JSON(http.StatusOK, gin.H{
			"success":    true,
			"session_id": sid,
			"message":    "session موجود است",
		})
		return
	}

	newSessionID := utils.GenerateSessionID()
	// تنظیم SameSite mode برای امنیت
	c.SetSameSite(http.SameSiteLaxMode)
	// ست‌کردن کوکی با مدت اعتبار 30 روز
	c.SetCookie("session_id", newSessionID, 30*24*60*60, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"session_id": newSessionID,
		"message":    "session جدید ایجاد شد",
	})
}

// getSessionID دریافت شناسه نشست از کوکی یا header
func (h *CartHandler) getSessionID(c *gin.Context) string {
	// ابتدا از کوکی بررسی می‌کنیم
	sessionID, err := c.Cookie("session_id")
	if err == nil && sessionID != "" {
		return sessionID
	}

	// سپس از header بررسی می‌کنیم
	sessionID = c.GetHeader("X-Session-ID")
	if sessionID != "" {
		return sessionID
	}

	// اگر session ID وجود ندارد، یک session ID جدید ایجاد می‌کنیم
	newSessionID := utils.GenerateSessionID()

	// تنظیم SameSite mode برای امنیت
	c.SetSameSite(http.SameSiteLaxMode)
	// تنظیم کوکی
	c.SetCookie("session_id", newSessionID, 30*24*60*60, "/", "", false, true)

	return newSessionID
}

// getProductThumbnail returns /media/... path for 150x150 thumbnail if exists
func (h *CartHandler) getProductThumbnail(productID int) string {
	// ابتدا سعی می‌کنیم از جدول product_images تصویر را بیابیم
	var pImg models.ProductImage
	if err := h.db.Where("product_id = ?", productID).Order("sort_order ASC").First(&pImg).Error; err == nil {
		if pImg.ImageURL != "" {
			return pImg.ImageURL
		}
	}

	// در صورت نیافتن، به رسانه‌ها fallback می‌کنیم (الگوی قدیمی)
	var mediaFile models.MediaFile
	err := h.db.Where("category = ? AND file_name LIKE ?", "products", fmt.Sprintf("product_%d%%", productID)).
		Order("created_at DESC").First(&mediaFile).Error
	if err != nil {
		return ""
	}

	var thumbnail models.MediaVariant
	err = h.db.Where("media_id = ? AND purpose = ?", mediaFile.ID, "thumbnail").
		First(&thumbnail).Error
	if err != nil {
		return ""
	}

	cleaned := strings.TrimPrefix(thumbnail.FilePath, "/")
	return "/media/" + cleaned
}

// MoveToNext علامت‌گذاری آیتم سبد برای خرید بعدی
func (h *CartHandler) MoveToNext(c *gin.Context) {
	var req struct {
		CartItemID uint `json:"cart_item_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.CartItemID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "cart_item_id لازم است"})
		return
	}
	if err := h.db.Model(&models.CartItem{}).Where("id = ?", req.CartItemID).Update("is_next_purchase", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// MoveToCart بازگرداندن آیتم از خرید بعدی به سبد
func (h *CartHandler) MoveToCart(c *gin.Context) {
	var req struct {
		CartItemID uint `json:"cart_item_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.CartItemID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "cart_item_id لازم است"})
		return
	}
	if err := h.db.Model(&models.CartItem{}).Where("id = ?", req.CartItemID).Update("is_next_purchase", false).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "خطا"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
