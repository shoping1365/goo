package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// OrderDTO ساختار داده‌های سفارش برای API
type OrderDTO struct {
	ID                 uint      `json:"id"`
	OrderNumber        string    `json:"orderNumber"`
	CustomerName       string    `json:"customerName"`
	CustomerPhone      string    `json:"customerPhone"`
	CustomerIP         string    `json:"customerIP"`
	UserAgent          string    `json:"userAgent"`
	TotalAmount        float64   `json:"totalAmount"`
	PaymentMethod      string    `json:"paymentMethod"`
	PaymentStatus      string    `json:"paymentStatus"`
	Status             string    `json:"status"`
	OrderIntegrity     string    `json:"orderIntegrity"`
	CreatedAt          time.Time `json:"createdAt"`
	ItemsCount         int       `json:"itemsCount"`
	ShippingAddress    string    `json:"shippingAddress"`
	ShippingCity       string    `json:"shippingCity"`
	ShippingProvince   string    `json:"shippingProvince"`
	ShippingPostalCode string    `json:"shippingPostalCode"`
	RecipientName      string    `json:"recipientName"`
	RecipientPhone     string    `json:"recipientPhone"`
	ShippingMethod     string    `json:"shippingMethod"`
	TrackingCode       string    `json:"trackingCode"`
}

// OrderHandler manages orders
type OrderHandler struct {
	DB           *gorm.DB
	orderService *services.OrderService
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	orderRepo := repository.NewOrderRepository(db)
	inventoryRepo := repository.NewInventoryRepository(db)
	orderService := services.NewOrderService(orderRepo, inventoryRepo)
	return &OrderHandler{DB: db, orderService: orderService}
}

// Create ایجاد سفارش جدید
func (h *OrderHandler) Create(c *gin.Context) {
	session := sessions.Default(c)
	userIDVal := session.Get("user_id")
	if userIDVal == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}
	userID := userIDVal.(uint)

	var req struct {
		ShippingAddressID uint   `json:"shipping_address_id" binding:"required"`
		PaymentMethod     string `json:"payment_method" binding:"required,oneof=online cod"`
		Items             []struct {
			ProductID uint `json:"product_id" binding:"required"`
			Quantity  int  `json:"quantity" binding:"required,min=1"`
		} `json:"items" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "VALIDATION_ERROR", "داده‌های ورودی نامعتبر است", err.Error())
		return
	}

	// ایجاد سفارش
	order := &models.Order{
		UserID:            &userID,
		ShippingAddressID: req.ShippingAddressID,
		PaymentMethod:     req.PaymentMethod,
		Status:            models.OrderStatusPending,
		PaymentStatus:     models.PaymentStatusPending,
		IsPaid:            false,
	}

	// محاسبه مبالغ و دریافت اطلاعات محصولات
	var subtotal float64
	for _, item := range req.Items {
		// دریافت اطلاعات محصول از دیتابیس
		var product models.Product
		if err := h.DB.First(&product, item.ProductID).Error; err != nil {
			utils.SendError(c, http.StatusNotFound, "PRODUCT_NOT_FOUND", fmt.Sprintf("محصول با شناسه %d یافت نشد", item.ProductID), nil)
			return
		}

		// لاگ برای دیباگ
		fmt.Printf("🔍 محصول یافت شد: ID=%d, Name=%s, Image=%s, SKU=%s\n",
			product.ID, product.Name, product.Image, product.SKU)

		unitPrice := product.Price
		totalPrice := unitPrice * float64(item.Quantity)
		subtotal += totalPrice

		orderItem := models.OrderItem{
			ProductID:    item.ProductID,
			ProductName:  product.Name,
			ProductImage: product.Image,
			ProductSKU:   product.SKU,
			Quantity:     item.Quantity,
			UnitPrice:    unitPrice,
			TotalPrice:   totalPrice,
			FinalPrice:   totalPrice,
		}

		// لاگ برای دیباگ
		fmt.Printf("📦 OrderItem ایجاد شد: ProductName=%s, ProductImage=%s, TotalPrice=%.2f\n",
			orderItem.ProductName, orderItem.ProductImage, orderItem.TotalPrice)

		order.OrderItems = append(order.OrderItems, orderItem)
	}

	order.SubtotalAmount = subtotal
	order.TotalAmount = subtotal
	order.FinalAmount = subtotal

	err := h.orderService.CreateOrder(order)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "ORDER_CREATE_FAILED", err.Error(), nil)
		return
	}

	// --- هوک ارزیابی تقلب پس از ایجاد سفارش ---
	go func(orderID uint, userID uint, paymentMethod string, finalAmount float64) {
		var u models.User
		_ = h.DB.First(&u, userID).Error
		// دریافت شهر آدرس از روی شناسه آدرس ارسال سفارش
		var addr models.UserAddress
		_ = h.DB.Select("city").Where("id = ?", order.ShippingAddressID).First(&addr).Error
		// IP و DeviceID از هدرها
		ip := c.ClientIP()
		deviceID := c.GetHeader("X-Device-ID")

		// ایجاد fraud service
		uowf := unitofwork.NewUnitOfWorkFactory(h.DB)
		fraud := services.NewFraudService(uowf, "internal_salt")

		_, _ = fraud.EvaluateOrder(c, services.EvaluateOrderInput{
			OrderID:          orderID,
			UserID:           &userID,
			Amount:           finalAmount,
			PaymentMethod:    paymentMethod,
			UserRegisteredAt: &u.RegisteredAt,
			IP:               ip,
			DeviceID:         deviceID,
			ShippingCity:     addr.City,
		})
	}(order.ID, userID, order.PaymentMethod, order.FinalAmount)

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) ListMyOrders(c *gin.Context) {
	uidVal, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}
	uid := uidVal.(uint)
	var orders []models.Order
	if err := h.DB.Preload("OrderItems").Where("user_id = ?", uid).Order("created_at DESC").Find(&orders).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش‌ها", err.Error())
		return
	}
	c.JSON(http.StatusOK, orders)
}

// لیست همه سفارش‌ها برای ادمین با اطلاعات تکمیلی
func (h *OrderHandler) ListAllOrders(c *gin.Context) {
	// --- Pagination Handling ---
	pageParam := c.DefaultQuery("page", "1")
	sizeParam := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(sizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Hard-limit pageSize to prevent abuse and performance issues
	// Maximum reasonable page size for orders list to prevent massive data loading
	// This prevents the system from loading thousands of orders at once which causes performance issues
	if pageSize > 100 {
		pageSize = 100
		// Log when pageSize is limited to help with monitoring
		fmt.Printf("⚠️  Order list pageSize limited to 100 (requested: %d) to prevent performance issues\n", pageSize)
	}

	offset := (page - 1) * pageSize

	// --- Incremental Loading (afterId) ---
	afterIDParam := c.Query("afterId")
	var afterID uint64
	if afterIDParam != "" {
		if parsed, err := strconv.ParseUint(afterIDParam, 10, 64); err == nil {
			afterID = parsed
		}
	}

	// Build base query with optional afterId filter
	dbQuery := h.DB
	if afterID > 0 {
		dbQuery = dbQuery.Where("id > ?", afterID)
	}

	// Total count for frontend (optional)
	var total int64
	if err := dbQuery.Model(&models.Order{}).Count(&total).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌ها", err.Error())
		return
	}

	// بهینه‌سازی: استفاده از JOIN query به جای دو کوئری جداگانه
	type OrderWithUser struct {
		models.Order
		UserName           string `json:"user_name"`
		UserMobile         string `json:"user_mobile"`
		ShippingAddress    string `json:"shipping_address"`
		ShippingCity       string `json:"shipping_city"`
		ShippingProvince   string `json:"shipping_province"`
		ShippingPostalCode string `json:"shipping_postal_code"`
		RecipientName      string `json:"recipient_name"`
		RecipientPhone     string `json:"recipient_phone"`
	}

	var ordersWithUser []OrderWithUser

	// ساخت کوئری اصلی با JOIN
	baseQuery := h.DB.Table("orders").
		Select(`orders.*, 
			COALESCE(users.name, '') as user_name, 
			COALESCE(users.mobile, '') as user_mobile,
			COALESCE(user_addresses.street, '') as shipping_address,
			COALESCE(user_addresses.city, '') as shipping_city,
			COALESCE(user_addresses.province, '') as shipping_province,
			COALESCE(user_addresses.postal_code, '') as shipping_postal_code,
			COALESCE(user_addresses.recipient_name, '') as recipient_name,
			COALESCE(user_addresses.phone, '') as recipient_phone`).
		Joins("LEFT JOIN users ON users.id = orders.user_id").
		Joins("LEFT JOIN user_addresses ON user_addresses.id = orders.shipping_address_id")

	// اعمال فیلتر afterID
	if afterID > 0 {
		baseQuery = baseQuery.Where("orders.id > ?", afterID)
	}

	// اجرای کوئری
	if afterID > 0 {
		if err := baseQuery.Order("orders.created_at DESC").Scan(&ordersWithUser).Error; err != nil {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش‌ها", err.Error())
			return
		}
	} else {
		if err := baseQuery.Order("orders.created_at DESC").Limit(pageSize).Offset(offset).Scan(&ordersWithUser).Error; err != nil {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش‌ها", err.Error())
			return
		}
	}

	// دریافت تعداد OrderItems برای همه سفارشات در یک کوئری (بهینه‌سازی)
	// فقط در صورت وجود سفارشات، کوئری را اجرا کن
	var orderItemCounts []struct {
		OrderID uint `json:"order_id"`
		Count   int  `json:"count"`
	}

	if len(ordersWithUser) > 0 {
		orderIDs := make([]uint, len(ordersWithUser))
		for i, order := range ordersWithUser {
			orderIDs[i] = order.ID
		}

		// استفاده از prepared statement برای بهینه‌سازی عملکرد
		// محدود کردن تعداد سفارشات برای جلوگیری از بارگذاری بیش از حد داده
		if len(orderIDs) > 100 {
			// اگر تعداد سفارشات بیش از 100 باشد، فقط برای 100 سفارش اول آیتم‌ها را شمارش کن
			orderIDs = orderIDs[:100]
			fmt.Printf("⚠️  Order items count query limited to 100 orders (total: %d) to prevent performance issues\n", len(ordersWithUser))
		}

		if err := h.DB.Table("order_items").
			Select("order_id, COUNT(*) as count").
			Where("order_id IN ?", orderIDs).
			Group("order_id").
			Scan(&orderItemCounts).Error; err != nil {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت تعداد آیتم‌های سفارش", err.Error())
			return
		}
	}

	// گروه‌بندی تعداد آیتم‌ها بر اساس OrderID
	orderItemCountMap := make(map[uint]int)
	for _, count := range orderItemCounts {
		orderItemCountMap[count.OrderID] = count.Count
	}

	// ساخت آرایه خروجی

	// ساخت آرایه خروجی بهینه شده
	var dtoList []OrderDTO
	for _, order := range ordersWithUser {
		dto := OrderDTO{
			ID:                 order.ID,
			OrderNumber:        order.OrderNumber, // استفاده از OrderNumber
			CustomerName:       order.UserName,    // از JOIN بدست آمده
			CustomerPhone:      order.UserMobile,  // از JOIN بدست آمده
			CustomerIP:         "",                // IP کاربر - باید از دیتابیس دریافت شود
			UserAgent:          "",                // User-Agent - باید از دیتابیس دریافت شود
			TotalAmount:        order.TotalAmount, // استفاده از TotalAmount به جای FinalAmount
			PaymentMethod:      order.PaymentMethod,
			Status:             order.Status,
			OrderIntegrity:     "verified", // TODO: پیاده‌سازی منطق صحت سفارش
			CreatedAt:          order.CreatedAt,
			ItemsCount:         orderItemCountMap[order.ID], // تعداد آیتم‌ها
			ShippingAddress:    order.ShippingAddress,
			ShippingCity:       order.ShippingCity,
			ShippingProvince:   order.ShippingProvince,
			ShippingPostalCode: order.ShippingPostalCode,
			RecipientName:      order.RecipientName,
			RecipientPhone:     order.RecipientPhone,
		}
		dtoList = append(dtoList, dto)
	}

	// پاسخ مطابق با نوع درخواست (صفحه‌بندی یا لود افزایشی)
	if afterID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":  dtoList,
			"total": total,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  dtoList,
		"page":  page,
		"total": total,
	})
}

// AdminOrderStats محاسبه آمار کارت‌های صفحه سفارشات ادمین
func (h *OrderHandler) AdminOrderStats(c *gin.Context) {
	// شمارنده‌های وضعیت‌ها
	var total int64
	if err := h.DB.Model(&models.Order{}).Count(&total).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش کل سفارش‌ها", err.Error())
		return
	}

	var pending int64
	if err := h.DB.Model(&models.Order{}).Where("payment_status = ?", "awaiting_payment").Count(&pending).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های در انتظار پرداخت", err.Error())
		return
	}

	var processing int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "processing").Count(&processing).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های در حال پردازش", err.Error())
		return
	}

	var shipped int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "shipped").Count(&shipped).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های ارسال شده", err.Error())
		return
	}

	var delivered int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "delivered").Count(&delivered).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های تحویل شده", err.Error())
		return
	}

	var cancelled int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "cancelled").Count(&cancelled).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های لغو شده", err.Error())
		return
	}

	// مبالغ فروش
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	weekStart := now.AddDate(0, 0, -7)
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	var todaySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND created_at >= ?", true, todayStart).
		Scan(&todaySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش امروز", err.Error())
		return
	}

	var weeklySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND created_at >= ?", true, weekStart).
		Scan(&weeklySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش هفتگی", err.Error())
		return
	}

	var monthlySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND created_at >= ?", true, monthStart).
		Scan(&monthlySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش ماه جاری", err.Error())
		return
	}

	// میانگین مبلغ سفارش‌های پرداخت‌شده
	var averageOrder float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(AVG(final_amount),0)").
		Where("is_paid = ?", true).
		Scan(&averageOrder).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه متوسط سفارش", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":        total,
		"pending":      pending,
		"processing":   processing,
		"shipped":      shipped,
		"delivered":    delivered,
		"cancelled":    cancelled,
		"todaySales":   todaySales,
		"weeklySales":  weeklySales,
		"monthlySales": monthlySales,
		"averageOrder": averageOrder,
	})
}

// GetOrderItems دریافت آیتم‌های یک سفارش خاص
func (h *OrderHandler) GetOrderItems(c *gin.Context) {
	orderIDStr := c.Param("id")
	if orderIDStr == "" {
		utils.SendError(c, http.StatusBadRequest, "INVALID_ORDER_ID", "شناسه سفارش الزامی است", nil)
		return
	}

	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "INVALID_ORDER_ID", "شناسه سفارش نامعتبر است", err.Error())
		return
	}

	var orderItems []models.OrderItem
	if err := h.DB.Where("order_id = ?", uint(orderID)).Find(&orderItems).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت آیتم‌های سفارش", err.Error())
		return
	}

	// Debug: Log the query result
	fmt.Printf("Found %d items for order ID %d\n", len(orderItems), orderID)

	// اگر هیچ آیتمی یافت نشد، بررسی کن که آیا سفارش وجود دارد
	if len(orderItems) == 0 {
		var orderExists bool
		h.DB.Model(&models.Order{}).Select("1").Where("id = ?", uint(orderID)).Find(&orderExists)
		if !orderExists {
			utils.SendError(c, http.StatusNotFound, "ORDER_NOT_FOUND", "سفارش یافت نشد", nil)
			return
		}
		fmt.Printf("Order %d exists but has no items\n", orderID)

		// بررسی آیتم‌های dummy (product_id = 0)
		var dummyItems []models.OrderItem
		if err := h.DB.Where("order_id = ? AND product_id = 0", uint(orderID)).Find(&dummyItems).Error; err == nil && len(dummyItems) > 0 {
			fmt.Printf("Found %d dummy items for order %d\n", len(dummyItems), orderID)
			orderItems = dummyItems
		}
	}

	// Debug: Log first few items if any
	if len(orderItems) > 0 {
		fmt.Printf("First item: %+v\n", orderItems[0])
	} else {
		// Check if there are any order_items in the database at all
		var totalItems int64
		h.DB.Model(&models.OrderItem{}).Count(&totalItems)
		fmt.Printf("Total order items in database: %d\n", totalItems)

		// Check if the order exists
		var orderExists bool
		h.DB.Model(&models.Order{}).Select("1").Where("id = ?", uint(orderID)).Find(&orderExists)
		fmt.Printf("Order %d exists: %t\n", orderID, orderExists)
	}

	// Create response with product names
	type OrderItemWithProduct struct {
		ID          uint    `json:"id"`
		OrderID     uint    `json:"order_id"`
		ProductID   uint    `json:"product_id"`
		ProductName string  `json:"product_name"`
		Quantity    int     `json:"quantity"`
		UnitPrice   float64 `json:"unit_price"`
		FinalPrice  float64 `json:"final_price"`
	}

	var itemsWithProducts []OrderItemWithProduct
	for _, item := range orderItems {
		// Get product name
		var productName string

		// اگر product_id = 0 باشد، از اطلاعات موجود در order_item استفاده کن
		if item.ProductID == 0 {
			productName = fmt.Sprintf("محصول سفارش %d", item.ID)
			fmt.Printf("Product ID is 0 for order item %d, using fallback name\n", item.ID)
		} else {
			err := h.DB.Model(&models.Product{}).Select("name").Where("id = ?", item.ProductID).Scan(&productName)
			if err != nil || productName == "" {
				fmt.Printf("Error getting product name for ID %d: %v\n", item.ProductID, err)
				productName = fmt.Sprintf("محصول %d", item.ProductID)
			}
		}

		itemsWithProducts = append(itemsWithProducts, OrderItemWithProduct{
			ID:          item.ID,
			OrderID:     item.OrderID,
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
			FinalPrice:  item.FinalPrice,
		})
	}

	fmt.Printf("Returning %d items with product names\n", len(itemsWithProducts))
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    itemsWithProducts,
	})
}

// ListInProgressOrders لیست سفارشات در صف پردازش برای ادمین
func (h *OrderHandler) ListInProgressOrders(c *gin.Context) {
	// --- Pagination Handling ---
	pageParam := c.DefaultQuery("page", "1")
	sizeParam := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(sizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	// فیلتر سفارشات در صف پردازش
	// شامل سفارشاتی که در حال حاضر در این وضعیت‌ها هستند
	inProgressStatuses := []string{"processing_queue", "awaiting_payment", "pending_review", "processing", "ready_to_ship"}

	// Debug: بررسی تعداد کل سفارشات
	var totalOrders int64
	h.DB.Model(&models.Order{}).Count(&totalOrders)
	fmt.Printf("Total orders in database: %d\n", totalOrders)

	// Debug: بررسی سفارشات با وضعیت‌های مختلف
	for _, status := range inProgressStatuses {
		var count int64
		h.DB.Model(&models.Order{}).Where("status = ?", status).Count(&count)
		fmt.Printf("Orders with status '%s': %d\n", status, count)
	}

	// ساخت کوئری اصلی با JOIN برای سفارشات در صف پردازش
	type OrderWithUser struct {
		models.Order
		UserName           string `json:"user_name"`
		UserMobile         string `json:"user_mobile"`
		ShippingAddress    string `json:"shipping_address"`
		ShippingCity       string `json:"shipping_city"`
		ShippingProvince   string `json:"shipping_province"`
		ShippingPostalCode string `json:"shipping_postal_code"`
		RecipientName      string `json:"recipient_name"`
		RecipientPhone     string `json:"recipient_phone"`
	}

	var ordersWithUser []OrderWithUser

	// ساخت کوئری اصلی با JOIN
	baseQuery := h.DB.Table("orders").
		Select(`orders.*, 
			COALESCE(users.name, '') as user_name, 
			COALESCE(users.mobile, '') as user_mobile,
			COALESCE(user_addresses.street, '') as shipping_address,
			COALESCE(user_addresses.city, '') as shipping_city,
			COALESCE(user_addresses.province, '') as shipping_province,
			COALESCE(user_addresses.postal_code, '') as shipping_postal_code,
			COALESCE(user_addresses.recipient_name, '') as recipient_name,
			COALESCE(user_addresses.phone, '') as recipient_phone`).
		Joins("LEFT JOIN users ON users.id = orders.user_id").
		Joins("LEFT JOIN user_addresses ON user_addresses.id = orders.shipping_address_id").
		Where("orders.status IN ?", inProgressStatuses)

	// شمارش کل سفارشات در صف پردازش
	var total int64
	if err := baseQuery.Model(&models.Order{}).Count(&total).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های در صف پردازش", err.Error())
		return
	}

	// اجرای کوئری با صفحه‌بندی
	if err := baseQuery.Order("orders.created_at DESC").Limit(pageSize).Offset(offset).Scan(&ordersWithUser).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش‌های در صف پردازش", err.Error())
		return
	}

	// دریافت تعداد OrderItems برای همه سفارشات در یک کوئری
	var orderItemCounts []struct {
		OrderID uint `json:"order_id"`
		Count   int  `json:"count"`
	}

	if len(ordersWithUser) > 0 {
		orderIDs := make([]uint, len(ordersWithUser))
		for i, order := range ordersWithUser {
			orderIDs[i] = order.ID
		}

		if err := h.DB.Table("order_items").
			Select("order_id, COUNT(*) as count").
			Where("order_id IN ?", orderIDs).
			Group("order_id").
			Scan(&orderItemCounts).Error; err != nil {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت تعداد آیتم‌های سفارش", err.Error())
			return
		}
	}

	// گروه‌بندی تعداد آیتم‌ها بر اساس OrderID
	orderItemCountMap := make(map[uint]int)
	for _, count := range orderItemCounts {
		orderItemCountMap[count.OrderID] = count.Count
	}

	// ساخت آرایه خروجی بهینه شده
	var dtoList []OrderDTO
	for _, order := range ordersWithUser {
		dto := OrderDTO{
			ID:                 order.ID,
			OrderNumber:        order.OrderNumber, // استفاده از OrderNumber
			CustomerName:       order.UserName,
			CustomerPhone:      order.UserMobile,
			CustomerIP:         "",                // IP کاربر - باید از دیتابیس دریافت شود
			UserAgent:          "",                // User-Agent - باید از دیتابیس دریافت شود
			TotalAmount:        order.TotalAmount, // استفاده از TotalAmount به جای FinalAmount
			PaymentMethod:      order.PaymentMethod,
			Status:             order.Status,
			OrderIntegrity:     "verified",
			CreatedAt:          order.CreatedAt,
			ItemsCount:         orderItemCountMap[order.ID],
			ShippingAddress:    order.ShippingAddress,
			ShippingCity:       order.ShippingCity,
			ShippingProvince:   order.ShippingProvince,
			ShippingPostalCode: order.ShippingPostalCode,
			RecipientName:      order.RecipientName,
			RecipientPhone:     order.RecipientPhone,
			ShippingMethod:     "", // روش ارسال - باید از دیتابیس دریافت شود
			TrackingCode:       order.TrackingCode,
		}
		dtoList = append(dtoList, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  dtoList,
		"page":  page,
		"total": total,
	})
}

// InProgressOrderStats آمار سفارشات در صف پردازش
func (h *OrderHandler) InProgressOrderStats(c *gin.Context) {
	// فیلتر سفارشات در صف پردازش
	inProgressStatuses := []string{"processing_queue", "awaiting_payment", "pending_review", "processing", "ready_to_ship"}

	// شمارنده‌های وضعیت‌ها
	var totalInQueue int64
	if err := h.DB.Model(&models.Order{}).Where("status IN ?", inProgressStatuses).Count(&totalInQueue).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش کل سفارش‌های در صف", err.Error())
		return
	}

	var pendingPayment int64
	if err := h.DB.Model(&models.Order{}).Where("payment_status = ?", "awaiting_payment").Count(&pendingPayment).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های در انتظار پرداخت", err.Error())
		return
	}

	var pendingReview int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "pending_review").Count(&pendingReview).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های در انتظار بررسی", err.Error())
		return
	}

	var processing int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "processing").Count(&processing).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های در حال پردازش", err.Error())
		return
	}

	var readyToShip int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "ready_to_ship").Count(&readyToShip).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های آماده ارسال", err.Error())
		return
	}

	// محاسبه متوسط زمان پردازش (ساعت)
	var avgProcessingTime float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(AVG(EXTRACT(EPOCH FROM (updated_at - created_at))/3600), 0)").
		Where("status IN ? AND updated_at > created_at", inProgressStatuses).
		Scan(&avgProcessingTime).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه متوسط زمان پردازش", err.Error())
		return
	}

	// آمار توزیع وضعیت‌ها
	var statusDistribution []struct {
		Status string `json:"status"`
		Count  int64  `json:"count"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Where("status IN ?", inProgressStatuses).
		Group("status").
		Scan(&statusDistribution).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت توزیع وضعیت‌ها", err.Error())
		return
	}

	// آمار روش‌های پرداخت
	var paymentMethods []struct {
		PaymentMethod string `json:"payment_method"`
		Count         int64  `json:"count"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("payment_method, COUNT(*) as count").
		Where("status IN ?", inProgressStatuses).
		Group("payment_method").
		Scan(&paymentMethods).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت آمار روش‌های پرداخت", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"totalInQueue":       totalInQueue,
		"pendingPayment":     pendingPayment,
		"pendingReview":      pendingReview,
		"processing":         processing,
		"readyToShip":        readyToShip,
		"avgProcessingTime":  avgProcessingTime,
		"statusDistribution": statusDistribution,
		"paymentMethods":     paymentMethods,
	})
}

// ListProcessingOrders لیست سفارشات در حال انجام برای ادمین
func (h *OrderHandler) ListProcessingOrders(c *gin.Context) {
	// --- Pagination Handling ---
	pageParam := c.DefaultQuery("page", "1")
	sizeParam := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(sizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	// فیلتر سفارشات در حال انجام
	// شامل سفارشاتی که وضعیتشان به "processing" تغییر کرده
	processingStatuses := []string{"processing"}

	// ساخت کوئری اصلی با JOIN برای سفارشات در حال انجام
	type OrderWithUser struct {
		models.Order
		UserName           string `json:"user_name"`
		UserMobile         string `json:"user_mobile"`
		ShippingAddress    string `json:"shipping_address"`
		ShippingCity       string `json:"shipping_city"`
		ShippingProvince   string `json:"shipping_province"`
		ShippingPostalCode string `json:"shipping_postal_code"`
		RecipientName      string `json:"recipient_name"`
		RecipientPhone     string `json:"recipient_phone"`
	}

	var ordersWithUser []OrderWithUser

	// ساخت کوئری اصلی با JOIN
	baseQuery := h.DB.Table("orders").
		Select(`orders.*, 
			COALESCE(users.name, '') as user_name, 
			COALESCE(users.mobile, '') as user_mobile,
			COALESCE(user_addresses.street, '') as shipping_address,
			COALESCE(user_addresses.city, '') as shipping_city,
			COALESCE(user_addresses.province, '') as shipping_province,
			COALESCE(user_addresses.postal_code, '') as shipping_postal_code,
			COALESCE(user_addresses.recipient_name, '') as recipient_name,
			COALESCE(user_addresses.phone, '') as recipient_phone`).
		Joins("LEFT JOIN users ON users.id = orders.user_id").
		Joins("LEFT JOIN user_addresses ON user_addresses.id = orders.shipping_address_id").
		Where("orders.status IN ?", processingStatuses)

	// شمارش کل سفارشات در حال انجام
	var total int64
	if err := baseQuery.Model(&models.Order{}).Count(&total).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های در حال انجام", err.Error())
		return
	}

	// اجرای کوئری با صفحه‌بندی
	if err := baseQuery.Order("orders.updated_at DESC").Limit(pageSize).Offset(offset).Scan(&ordersWithUser).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش‌های در حال انجام", err.Error())
		return
	}

	// دریافت تعداد OrderItems برای همه سفارشات در یک کوئری
	var orderItemCounts []struct {
		OrderID uint `json:"order_id"`
		Count   int  `json:"count"`
	}

	if len(ordersWithUser) > 0 {
		orderIDs := make([]uint, len(ordersWithUser))
		for i, order := range ordersWithUser {
			orderIDs[i] = order.ID
		}

		if err := h.DB.Table("order_items").
			Select("order_id, COUNT(*) as count").
			Where("order_id IN ?", orderIDs).
			Group("order_id").
			Scan(&orderItemCounts).Error; err != nil {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت تعداد آیتم‌های سفارش", err.Error())
			return
		}
	}

	// گروه‌بندی تعداد آیتم‌ها بر اساس OrderID
	orderItemCountMap := make(map[uint]int)
	for _, count := range orderItemCounts {
		orderItemCountMap[count.OrderID] = count.Count
	}

	// ساخت آرایه خروجی بهینه شده
	var dtoList []OrderDTO
	for _, order := range ordersWithUser {
		dto := OrderDTO{
			ID:                 order.ID,
			OrderNumber:        order.OrderNumber, // استفاده از OrderNumber
			CustomerName:       order.UserName,
			CustomerPhone:      order.UserMobile,
			CustomerIP:         "",                // IP کاربر - باید از دیتابیس دریافت شود
			UserAgent:          "",                // User-Agent - باید از دیتابیس دریافت شود
			TotalAmount:        order.TotalAmount, // استفاده از TotalAmount به جای FinalAmount
			PaymentMethod:      order.PaymentMethod,
			Status:             order.Status,
			OrderIntegrity:     "verified",
			CreatedAt:          order.CreatedAt,
			ItemsCount:         orderItemCountMap[order.ID],
			ShippingAddress:    order.ShippingAddress,
			ShippingCity:       order.ShippingCity,
			ShippingProvince:   order.ShippingProvince,
			ShippingPostalCode: order.ShippingPostalCode,
			RecipientName:      order.RecipientName,
			RecipientPhone:     order.RecipientPhone,
			ShippingMethod:     "", // روش ارسال - باید از دیتابیس دریافت شود
			TrackingCode:       order.TrackingCode,
		}
		dtoList = append(dtoList, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  dtoList,
		"page":  page,
		"total": total,
	})
}

// ProcessingOrderStats آمار سفارشات در حال انجام
func (h *OrderHandler) ProcessingOrderStats(c *gin.Context) {
	// فیلتر سفارشات در حال انجام
	processingStatuses := []string{"processing"}

	// شمارنده‌های وضعیت‌ها
	var totalProcessing int64
	if err := h.DB.Model(&models.Order{}).Where("status IN ?", processingStatuses).Count(&totalProcessing).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش کل سفارش‌های در حال انجام", err.Error())
		return
	}

	// محاسبه مبلغ کل سفارشات در حال انجام
	var totalAmount float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("status IN ?", processingStatuses).
		Scan(&totalAmount).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه مبلغ کل سفارشات در حال انجام", err.Error())
		return
	}

	// محاسبه متوسط زمان پردازش (ساعت)
	var avgProcessingTime float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(AVG(EXTRACT(EPOCH FROM (updated_at - created_at))/3600), 0)").
		Where("status IN ? AND updated_at > created_at", processingStatuses).
		Scan(&avgProcessingTime).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه متوسط زمان پردازش", err.Error())
		return
	}

	// محاسبه نرخ تکمیل (درصد سفارشات تکمیل شده در 24 ساعت گذشته)
	var completedToday int64
	var startedToday int64

	today := time.Now().AddDate(0, 0, -1) // 24 ساعت گذشته

	if err := h.DB.Model(&models.Order{}).
		Where("status = ? AND updated_at >= ?", "shipped", today).
		Count(&completedToday).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه سفارشات تکمیل شده", err.Error())
		return
	}

	if err := h.DB.Model(&models.Order{}).
		Where("status = ? AND created_at >= ?", "processing", today).
		Count(&startedToday).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه سفارشات شروع شده", err.Error())
		return
	}

	var completionRate float64
	if startedToday > 0 {
		completionRate = float64(completedToday) / float64(startedToday) * 100
	}

	// آمار توزیع روش‌های پرداخت
	var paymentMethods []struct {
		PaymentMethod string `json:"payment_method"`
		Count         int64  `json:"count"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("payment_method, COUNT(*) as count").
		Where("status IN ?", processingStatuses).
		Group("payment_method").
		Scan(&paymentMethods).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت آمار روش‌های پرداخت", err.Error())
		return
	}

	// آمار توزیع وضعیت‌ها (برای مراحل پردازش)
	var statusDistribution []struct {
		Status string `json:"status"`
		Count  int64  `json:"count"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Where("status IN ?", processingStatuses).
		Group("status").
		Scan(&statusDistribution).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت توزیع وضعیت‌ها", err.Error())
		return
	}

	// آمار زمانی تفصیلی
	var fastestProcessing float64
	var longestProcessing float64
	var processingPerDay float64

	// سریع‌ترین پردازش (کمترین زمان)
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(MIN(EXTRACT(EPOCH FROM (updated_at - created_at))/3600), 0)").
		Where("status IN ? AND updated_at > created_at", processingStatuses).
		Scan(&fastestProcessing).Error; err != nil {
		fastestProcessing = 0
	}

	// طولانی‌ترین پردازش
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(MAX(EXTRACT(EPOCH FROM (updated_at - created_at))/3600), 0)").
		Where("status IN ? AND updated_at > created_at", processingStatuses).
		Scan(&longestProcessing).Error; err != nil {
		longestProcessing = 0
	}

	// متوسط پردازش در روز (آخرین 7 روز)
	weekAgo := time.Now().AddDate(0, 0, -7)
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(COUNT(*)/7.0, 0)").
		Where("status IN ? AND created_at >= ?", processingStatuses, weekAgo).
		Scan(&processingPerDay).Error; err != nil {
		processingPerDay = 0
	}

	// داده‌های روند 30 روزه
	var trendData []struct {
		Date  string `json:"date"`
		Count int64  `json:"count"`
	}

	// تولید داده‌های 30 روز گذشته
	for i := 29; i >= 0; i-- {
		date := time.Now().AddDate(0, 0, -i)
		dateStr := date.Format("2006-01-02")

		var count int64
		h.DB.Model(&models.Order{}).
			Where("status IN ? AND DATE(created_at) = ?", processingStatuses, dateStr).
			Count(&count)

		trendData = append(trendData, struct {
			Date  string `json:"date"`
			Count int64  `json:"count"`
		}{
			Date:  dateStr,
			Count: count,
		})
	}

	// سفارشات اخیر در حال انجام (آخرین 10 سفارش)
	type RecentOrder struct {
		ID            uint    `json:"id"`
		OrderNumber   string  `json:"orderNumber"`
		CustomerName  string  `json:"customerName"`
		CustomerPhone string  `json:"customerPhone"`
		TotalAmount   float64 `json:"totalAmount"`
		Status        string  `json:"status"`
		PaymentMethod string  `json:"paymentMethod"`
		CreatedAt     string  `json:"createdAt"`
	}

	var recentOrders []RecentOrder
	if err := h.DB.Table("orders").
		Select(`orders.id, 
			orders.tracking_code as order_number,
			COALESCE(users.name, '') as customer_name,
			COALESCE(users.mobile, '') as customer_phone,
			orders.final_amount as total_amount,
			orders.status,
			orders.payment_method,
			orders.created_at`).
		Joins("LEFT JOIN users ON users.id = orders.user_id").
		Where("orders.status IN ?", processingStatuses).
		Order("orders.created_at DESC").
		Limit(10).
		Scan(&recentOrders).Error; err != nil {
		recentOrders = []RecentOrder{}
	}

	c.JSON(http.StatusOK, gin.H{
		"totalProcessing":    totalProcessing,
		"totalAmount":        totalAmount,
		"completionRate":     completionRate,
		"avgProcessingTime":  avgProcessingTime,
		"paymentMethods":     paymentMethods,
		"statusDistribution": statusDistribution,
		"timeStats": gin.H{
			"avgProcessingTime": avgProcessingTime,
			"fastestProcessing": fastestProcessing,
			"processingPerDay":  processingPerDay,
			"longestProcessing": longestProcessing,
		},
		"trendData":       trendData,
		"recentOrders":    recentOrders,
		"detailedReports": getDetailedReports(h.DB, processingStatuses),
	})
}

// getDetailedReports تولید گزارشات تفصیلی برای دوره‌های مختلف
func getDetailedReports(db *gorm.DB, processingStatuses []string) []gin.H {
	now := time.Now()

	// محاسبه برای دوره‌های مختلف
	periods := []struct {
		name  string
		start time.Time
		end   time.Time
	}{
		{"امروز", time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()), now},
		{"دیروز", time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, now.Location()), time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())},
		{"هفته جاری", now.AddDate(0, 0, -7), now},
		{"ماه جاری", time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()), now},
		{"سال جاری", time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location()), now},
	}

	var detailedReports []gin.H

	for _, p := range periods {
		var count int64
		var revenue float64
		var completedCount int64

		// شمارش سفارشات در حال انجام در این دوره
		db.Model(&models.Order{}).Where("status IN ? AND created_at BETWEEN ? AND ?", processingStatuses, p.start, p.end).Count(&count)

		// محاسبه درآمد
		db.Model(&models.Order{}).
			Select("COALESCE(SUM(final_amount),0)").
			Where("status IN ? AND created_at BETWEEN ? AND ?", processingStatuses, p.start, p.end).
			Scan(&revenue)

		// شمارش سفارشات تکمیل شده در این دوره
		db.Model(&models.Order{}).Where("status = ? AND updated_at BETWEEN ? AND ?", "shipped", p.start, p.end).Count(&completedCount)

		var avgValue float64
		var completionRate float64
		var change float64

		if count > 0 {
			avgValue = revenue / float64(count)
			completionRate = float64(completedCount) / float64(count) * 100
		}

		// محاسبه تغییر نسبت به دوره قبل (ساده‌سازی شده)
		change = 0 // TODO: پیاده‌سازی محاسبه تغییر

		detailedReports = append(detailedReports, gin.H{
			"period":         p.name,
			"orderCount":     count,
			"totalAmount":    revenue,
			"avgOrderValue":  avgValue,
			"completionRate": completionRate,
			"change":         change,
		})
	}

	return detailedReports
}

// GetOrderByID دریافت اطلاعات یک سفارش خاص
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")

	// دریافت ID سفارش
	id, err := strconv.ParseUint(orderID, 10, 32)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "INVALID_ORDER_ID", "شناسه سفارش نامعتبر است", err.Error())
		return
	}

	// ساخت کوئری برای دریافت سفارش با اطلاعات کاربر
	type OrderWithUser struct {
		models.Order
		UserName           string `json:"user_name"`
		UserMobile         string `json:"user_mobile"`
		ShippingAddress    string `json:"shipping_address"`
		ShippingCity       string `json:"shipping_city"`
		ShippingProvince   string `json:"shipping_province"`
		ShippingPostalCode string `json:"shipping_postal_code"`
		RecipientName      string `json:"recipient_name"`
		RecipientPhone     string `json:"recipient_phone"`
	}

	var orderWithUser OrderWithUser

	// دریافت سفارش با JOIN
	if err := h.DB.Table("orders").
		Select(`orders.*, 
			COALESCE(users.name, '') as user_name, 
			COALESCE(users.mobile, '') as user_mobile,
			COALESCE(user_addresses.street, '') as shipping_address,
			COALESCE(user_addresses.city, '') as shipping_city,
			COALESCE(user_addresses.province, '') as shipping_province,
			COALESCE(user_addresses.postal_code, '') as shipping_postal_code,
			COALESCE(user_addresses.recipient_name, '') as recipient_name,
			COALESCE(user_addresses.phone, '') as recipient_phone`).
		Joins("LEFT JOIN users ON users.id = orders.user_id").
		Joins("LEFT JOIN user_addresses ON user_addresses.id = orders.shipping_address_id").
		Where("orders.id = ?", id).
		First(&orderWithUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.SendError(c, http.StatusNotFound, "ORDER_NOT_FOUND", "سفارش یافت نشد", "")
		} else {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش", err.Error())
		}
		return
	}

	// محاسبه تعداد آیتم‌های سفارش
	var itemsCount int64
	if err := h.DB.Model(&models.OrderItem{}).Where("order_id = ?", orderWithUser.ID).Count(&itemsCount).Error; err != nil {
		fmt.Printf("خطا در محاسبه تعداد آیتم‌های سفارش %d: %v\n", orderWithUser.ID, err)
		itemsCount = 0
	}

	// ساخت DTO
	dto := OrderDTO{
		ID:                 orderWithUser.ID,
		OrderNumber:        orderWithUser.OrderNumber, // استفاده از OrderNumber به جای TrackingCode
		CustomerName:       orderWithUser.UserName,
		CustomerPhone:      orderWithUser.UserMobile,
		CustomerIP:         "",                        // IP کاربر - باید از دیتابیس دریافت شود
		UserAgent:          "",                        // User-Agent - باید از دیتابیس دریافت شود
		TotalAmount:        orderWithUser.TotalAmount, // استفاده از TotalAmount به جای FinalAmount
		PaymentMethod:      orderWithUser.PaymentMethod,
		PaymentStatus:      orderWithUser.PaymentStatus,
		Status:             orderWithUser.Status,
		OrderIntegrity:     "verified",
		CreatedAt:          orderWithUser.CreatedAt,
		ItemsCount:         int(itemsCount), // محاسبه صحیح تعداد آیتم‌ها
		ShippingAddress:    orderWithUser.ShippingAddress,
		ShippingCity:       orderWithUser.ShippingCity,
		ShippingProvince:   orderWithUser.ShippingProvince,
		ShippingPostalCode: orderWithUser.ShippingPostalCode,
		RecipientName:      orderWithUser.RecipientName,
		RecipientPhone:     orderWithUser.RecipientPhone,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    dto,
	})
}

// UpdateOrderStatus به‌روزرسانی وضعیت سفارش
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	// دریافت ID سفارش
	id, err := strconv.ParseUint(orderID, 10, 32)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "INVALID_ORDER_ID", "شناسه سفارش نامعتبر است", err.Error())
		return
	}

	// دریافت داده‌های درخواست
	var request struct {
		Status        string `json:"status"`
		PaymentStatus string `json:"paymentStatus"`
		Notes         string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendError(c, http.StatusBadRequest, "INVALID_REQUEST", "درخواست نامعتبر است", err.Error())
		return
	}

	// بررسی وجود سفارش
	var order models.Order
	if err := h.DB.First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.SendError(c, http.StatusNotFound, "ORDER_NOT_FOUND", "سفارش یافت نشد", "")
		} else {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش", err.Error())
		}
		return
	}

	// ذخیره وضعیت‌های قبلی
	oldStatus := order.Status
	oldPaymentStatus := order.PaymentStatus

	// به‌روزرسانی وضعیت‌ها
	updates := make(map[string]interface{})

	if request.Status != "" {
		updates["status"] = request.Status
	}
	if request.PaymentStatus != "" {
		updates["payment_status"] = request.PaymentStatus
	}

	// اگر هیچ تغییری وجود ندارد
	if len(updates) == 0 {
		utils.SendError(c, http.StatusBadRequest, "NO_CHANGES", "هیچ تغییری برای به‌روزرسانی وجود ندارد", "")
		return
	}

	updates["updated_at"] = time.Now()

	if err := h.DB.Model(&order).Updates(updates).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در به‌روزرسانی سفارش", err.Error())
		return
	}

	// لاگ تغییر وضعیت‌ها
	if request.Status != "" {
		fmt.Printf("Order %d status changed from '%s' to '%s'\n", order.ID, oldStatus, request.Status)
	}
	if request.PaymentStatus != "" {
		fmt.Printf("Order %d payment status changed from '%s' to '%s'\n", order.ID, oldPaymentStatus, request.PaymentStatus)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "وضعیت سفارش با موفقیت به‌روزرسانی شد",
		"data": gin.H{
			"id":               order.ID,
			"oldStatus":        oldStatus,
			"newStatus":        order.Status,
			"oldPaymentStatus": oldPaymentStatus,
			"newPaymentStatus": order.PaymentStatus,
			"updatedAt":        time.Now(),
		},
	})
}

// AdminOrderReports گزارشات جامع سفارشات برای ادمین
func (h *OrderHandler) AdminOrderReports(c *gin.Context) {
	// دریافت پارامترهای query
	period := c.DefaultQuery("period", "month") // month, year, week, day
	year := c.DefaultQuery("year", fmt.Sprintf("%d", time.Now().Year()))
	month := c.DefaultQuery("month", fmt.Sprintf("%d", int(time.Now().Month())))

	// محاسبه بازه زمانی
	now := time.Now()
	var startTime, endTime time.Time

	switch period {
	case "day":
		startTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		endTime = startTime.Add(24 * time.Hour)
	case "week":
		startTime = now.AddDate(0, 0, -7)
		endTime = now
	case "month":
		yearInt, _ := strconv.Atoi(year)
		monthInt, _ := strconv.Atoi(month)
		startTime = time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 0, 0, now.Location())
		endTime = startTime.AddDate(0, 1, 0)
	case "year":
		yearInt, _ := strconv.Atoi(year)
		startTime = time.Date(yearInt, 1, 1, 0, 0, 0, 0, now.Location())
		endTime = startTime.AddDate(1, 0, 0)
	default:
		startTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		endTime = startTime.AddDate(0, 1, 0)
	}

	// آمار کلی
	var totalOrders int64
	if err := h.DB.Model(&models.Order{}).Where("created_at BETWEEN ? AND ?", startTime, endTime).Count(&totalOrders).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش کل سفارش‌ها", err.Error())
		return
	}

	var totalRevenue float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND created_at BETWEEN ? AND ?", true, startTime, endTime).
		Scan(&totalRevenue).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه کل درآمد", err.Error())
		return
	}

	var avgOrderValue float64
	if totalOrders > 0 {
		avgOrderValue = totalRevenue / float64(totalOrders)
	}

	// آمار وضعیت‌ها
	var statusStats []struct {
		Status     string  `json:"status"`
		Count      int64   `json:"count"`
		Percentage float64 `json:"percentage"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Where("created_at BETWEEN ? AND ?", startTime, endTime).
		Group("status").
		Scan(&statusStats).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت آمار وضعیت‌ها", err.Error())
		return
	}

	// محاسبه درصدها
	for i := range statusStats {
		if totalOrders > 0 {
			statusStats[i].Percentage = float64(statusStats[i].Count) / float64(totalOrders) * 100
		}
	}

	// آمار روش‌های پرداخت
	var paymentMethodStats []struct {
		PaymentMethod string  `json:"payment_method"`
		Count         int64   `json:"count"`
		Percentage    float64 `json:"percentage"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("payment_method, COUNT(*) as count").
		Where("created_at BETWEEN ? AND ?", startTime, endTime).
		Group("payment_method").
		Scan(&paymentMethodStats).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت آمار روش‌های پرداخت", err.Error())
		return
	}

	// محاسبه درصدها
	for i := range paymentMethodStats {
		if totalOrders > 0 {
			paymentMethodStats[i].Percentage = float64(paymentMethodStats[i].Count) / float64(totalOrders) * 100
		}
	}

	// آمار ماهانه برای نمودار
	var monthlyData []struct {
		Month string  `json:"month"`
		Value float64 `json:"value"`
	}

	if period == "year" {
		// برای سال، داده‌های ماهانه
		for i := 1; i <= 12; i++ {
			monthStart := time.Date(now.Year(), time.Month(i), 1, 0, 0, 0, 0, now.Location())
			monthEnd := monthStart.AddDate(0, 1, 0)

			var monthRevenue float64
			h.DB.Model(&models.Order{}).
				Select("COALESCE(SUM(final_amount),0)").
				Where("is_paid = ? AND created_at BETWEEN ? AND ?", true, monthStart, monthEnd).
				Scan(&monthRevenue)

			monthNames := []string{"فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور",
				"مهر", "آبان", "آذر", "دی", "بهمن", "اسفند"}

			monthlyData = append(monthlyData, struct {
				Month string  `json:"month"`
				Value float64 `json:"value"`
			}{
				Month: monthNames[i-1],
				Value: monthRevenue,
			})
		}
	}

	// گزارشات تفصیلی برای دوره‌های مختلف
	var detailedReports []struct {
		Period         string  `json:"period"`
		OrderCount     int64   `json:"order_count"`
		TotalRevenue   float64 `json:"total_revenue"`
		AvgOrderValue  float64 `json:"avg_order_value"`
		ConversionRate float64 `json:"conversion_rate"`
		Change         float64 `json:"change"`
	}

	// محاسبه برای دوره‌های مختلف
	periods := []struct {
		name  string
		start time.Time
		end   time.Time
	}{
		{"امروز", time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()), now},
		{"دیروز", time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, now.Location()), time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())},
		{"هفته جاری", now.AddDate(0, 0, -7), now},
		{"ماه جاری", time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()), now},
	}

	for _, p := range periods {
		var count int64
		var revenue float64

		h.DB.Model(&models.Order{}).Where("created_at BETWEEN ? AND ?", p.start, p.end).Count(&count)
		h.DB.Model(&models.Order{}).
			Select("COALESCE(SUM(final_amount),0)").
			Where("is_paid = ? AND created_at BETWEEN ? AND ?", true, p.start, p.end).
			Scan(&revenue)

		var avgValue float64
		if count > 0 {
			avgValue = revenue / float64(count)
		}

		detailedReports = append(detailedReports, struct {
			Period         string  `json:"period"`
			OrderCount     int64   `json:"order_count"`
			TotalRevenue   float64 `json:"total_revenue"`
			AvgOrderValue  float64 `json:"avg_order_value"`
			ConversionRate float64 `json:"conversion_rate"`
			Change         float64 `json:"change"`
		}{
			Period:         p.name,
			OrderCount:     count,
			TotalRevenue:   revenue,
			AvgOrderValue:  avgValue,
			ConversionRate: 0, // TODO: محاسبه نرخ تبدیل
			Change:         0, // TODO: محاسبه تغییر نسبت به دوره قبل
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"comprehensiveStats": gin.H{
				"totalOrders":    totalOrders,
				"totalRevenue":   totalRevenue,
				"avgOrderValue":  avgOrderValue,
				"conversionRate": 0, // TODO: محاسبه نرخ تبدیل
			},
			"yearlySalesData":       monthlyData,
			"orderStatusComparison": statusStats,
			"paymentMethodStats":    paymentMethodStats,
			"detailedReports":       detailedReports,
			"period":                period,
			"year":                  year,
			"month":                 month,
		},
	})
}

// ListShippedOrders لیست سفارشات ارسال شده برای ادمین
func (h *OrderHandler) ListShippedOrders(c *gin.Context) {
	// --- Pagination Handling ---
	pageParam := c.DefaultQuery("page", "1")
	sizeParam := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(sizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	// فیلتر سفارشات ارسال شده
	// شامل سفارشاتی که وضعیتشان به "shipped" یا "delivered" یا "in_transit" تغییر کرده
	shippedStatuses := []string{"shipped", "delivered", "in_transit"}

	// ساخت کوئری اصلی با JOIN برای سفارشات ارسال شده
	type OrderWithUser struct {
		models.Order
		UserName           string `json:"user_name"`
		UserMobile         string `json:"user_mobile"`
		ShippingAddress    string `json:"shipping_address"`
		ShippingCity       string `json:"shipping_city"`
		ShippingProvince   string `json:"shipping_province"`
		ShippingPostalCode string `json:"shipping_postal_code"`
		RecipientName      string `json:"recipient_name"`
		RecipientPhone     string `json:"recipient_phone"`
	}

	var ordersWithUser []OrderWithUser

	// ساخت کوئری اصلی با JOIN
	baseQuery := h.DB.Table("orders").
		Select(`orders.*,
			COALESCE(users.name, '') as user_name,
			COALESCE(users.mobile, '') as user_mobile,
			COALESCE(user_addresses.street, '') as shipping_address,
			COALESCE(user_addresses.city, '') as shipping_city,
			COALESCE(user_addresses.province, '') as shipping_province,
			COALESCE(user_addresses.postal_code, '') as shipping_postal_code,
			COALESCE(user_addresses.recipient_name, '') as recipient_name,
			COALESCE(user_addresses.phone, '') as recipient_phone`).
		Joins("LEFT JOIN users ON users.id = orders.user_id").
		Joins("LEFT JOIN user_addresses ON user_addresses.id = orders.shipping_address_id").
		Where("orders.status IN ?", shippedStatuses)

	// شمارش کل سفارشات ارسال شده
	var total int64
	if err := baseQuery.Model(&models.Order{}).Count(&total).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های ارسال شده", err.Error())
		return
	}

	// اجرای کوئری با صفحه‌بندی
	if err := baseQuery.Order("orders.updated_at DESC").Limit(pageSize).Offset(offset).Scan(&ordersWithUser).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش‌های ارسال شده", err.Error())
		return
	}

	// دریافت تعداد OrderItems برای همه سفارشات در یک کوئری
	var orderItemCounts []struct {
		OrderID uint `json:"order_id"`
		Count   int  `json:"count"`
	}

	if len(ordersWithUser) > 0 {
		orderIDs := make([]uint, len(ordersWithUser))
		for i, order := range ordersWithUser {
			orderIDs[i] = order.ID
		}

		if err := h.DB.Table("order_items").
			Select("order_id, COUNT(*) as count").
			Where("order_id IN ?", orderIDs).
			Group("order_id").
			Scan(&orderItemCounts).Error; err != nil {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت تعداد آیتم‌های سفارش", err.Error())
			return
		}
	}

	// گروه‌بندی تعداد آیتم‌ها بر اساس OrderID
	orderItemCountMap := make(map[uint]int)
	for _, count := range orderItemCounts {
		orderItemCountMap[count.OrderID] = count.Count
	}

	// ساخت آرایه خروجی بهینه شده
	var dtoList []OrderDTO
	for _, order := range ordersWithUser {
		dto := OrderDTO{
			ID:                 order.ID,
			OrderNumber:        order.OrderNumber, // استفاده از OrderNumber
			CustomerName:       order.UserName,
			CustomerPhone:      order.UserMobile,
			CustomerIP:         "",                // IP کاربر - باید از دیتابیس دریافت شود
			UserAgent:          "",                // User-Agent - باید از دیتابیس دریافت شود
			TotalAmount:        order.TotalAmount, // استفاده از TotalAmount به جای FinalAmount
			PaymentMethod:      order.PaymentMethod,
			Status:             order.Status,
			OrderIntegrity:     "verified",
			CreatedAt:          order.CreatedAt,
			ItemsCount:         orderItemCountMap[order.ID],
			ShippingAddress:    order.ShippingAddress,
			ShippingCity:       order.ShippingCity,
			ShippingProvince:   order.ShippingProvince,
			ShippingPostalCode: order.ShippingPostalCode,
			RecipientName:      order.RecipientName,
			RecipientPhone:     order.RecipientPhone,
			ShippingMethod:     "", // روش ارسال - باید از دیتابیس دریافت شود
			TrackingCode:       order.TrackingCode,
		}
		dtoList = append(dtoList, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  dtoList,
		"page":  page,
		"total": total,
	})
}

// ShippedOrderStats آمار سفارشات ارسال شده
func (h *OrderHandler) ShippedOrderStats(c *gin.Context) {
	// فیلتر سفارشات ارسال شده
	shippedStatuses := []string{"shipped", "delivered", "in_transit"}

	// شمارنده‌های وضعیت‌ها
	var totalShipped int64
	if err := h.DB.Model(&models.Order{}).Where("status IN ?", shippedStatuses).Count(&totalShipped).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش کل سفارش‌های ارسال شده", err.Error())
		return
	}

	var shipped int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "shipped").Count(&shipped).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های ارسال شده", err.Error())
		return
	}

	var inTransit int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "in_transit").Count(&inTransit).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های در حال ارسال", err.Error())
		return
	}

	var delivered int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "delivered").Count(&delivered).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های تحویل شده", err.Error())
		return
	}

	// مبالغ فروش
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	weekStart := now.AddDate(0, 0, -7)
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	var todaySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, shippedStatuses, todayStart).
		Scan(&todaySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش امروز", err.Error())
		return
	}

	var weeklySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, shippedStatuses, weekStart).
		Scan(&weeklySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش هفتگی", err.Error())
		return
	}

	var monthlySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, shippedStatuses, monthStart).
		Scan(&monthlySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش ماه جاری", err.Error())
		return
	}

	// میانگین مبلغ سفارش‌های پرداخت‌شده
	var averageOrder float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(AVG(final_amount),0)").
		Where("is_paid = ? AND status IN ?", true, shippedStatuses).
		Scan(&averageOrder).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه متوسط سفارش", err.Error())
		return
	}

	// محاسبه نرخ موفقیت تحویل
	var deliverySuccessRate float64
	if totalShipped > 0 {
		deliverySuccessRate = float64(delivered) / float64(totalShipped) * 100
	}

	// محاسبه متوسط زمان تحویل (ساعت)
	var avgDeliveryTime float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(AVG(EXTRACT(EPOCH FROM (updated_at - created_at))/3600), 0)").
		Where("status IN ? AND updated_at > created_at", shippedStatuses).
		Scan(&avgDeliveryTime).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه متوسط زمان تحویل", err.Error())
		return
	}

	// آمار شرکت‌های ارسال
	var shippingMethods []struct {
		ShippingMethod string  `json:"shipping_method"`
		Count          int64   `json:"count"`
		Percentage     float64 `json:"percentage"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("shipping_method, COUNT(*) as count").
		Where("status IN ?", shippedStatuses).
		Group("shipping_method").
		Scan(&shippingMethods).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت آمار شرکت‌های ارسال", err.Error())
		return
	}

	// محاسبه درصدها
	for i := range shippingMethods {
		if totalShipped > 0 {
			shippingMethods[i].Percentage = float64(shippingMethods[i].Count) / float64(totalShipped) * 100
		}
	}

	// آمار توزیع وضعیت‌ها
	var statusDistribution []struct {
		Status string `json:"status"`
		Count  int64  `json:"count"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Where("status IN ?", shippedStatuses).
		Group("status").
		Scan(&statusDistribution).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت توزیع وضعیت‌ها", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"totalShipped":        totalShipped,
		"shipped":             shipped,
		"inTransit":           inTransit,
		"delivered":           delivered,
		"todaySales":          todaySales,
		"weeklySales":         weeklySales,
		"monthlySales":        monthlySales,
		"averageOrder":        averageOrder,
		"deliverySuccessRate": deliverySuccessRate,
		"avgDeliveryTime":     avgDeliveryTime,
		"shippingMethods":     shippingMethods,
		"statusDistribution":  statusDistribution,
	})
}

// ListReturnedOrders لیست سفارشات مرجوع شده برای ادمین
func (h *OrderHandler) ListReturnedOrders(c *gin.Context) {
	// --- Pagination Handling ---
	pageParam := c.DefaultQuery("page", "1")
	sizeParam := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(sizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	// فیلتر سفارشات مرجوع شده
	// شامل سفارشاتی که وضعیتشان به "returned" یا "refunded" تغییر کرده
	returnedStatuses := []string{"returned", "refunded", "return_pending"}

	// ساخت کوئری اصلی با JOIN برای سفارشات مرجوع شده
	type OrderWithUser struct {
		models.Order
		UserName           string `json:"user_name"`
		UserMobile         string `json:"user_mobile"`
		ShippingAddress    string `json:"shipping_address"`
		ShippingCity       string `json:"shipping_city"`
		ShippingProvince   string `json:"shipping_province"`
		ShippingPostalCode string `json:"shipping_postal_code"`
		RecipientName      string `json:"recipient_name"`
		RecipientPhone     string `json:"recipient_phone"`
	}

	var ordersWithUser []OrderWithUser

	// ساخت کوئری اصلی با JOIN
	baseQuery := h.DB.Table("orders").
		Select(`orders.*,
			COALESCE(users.name, '') as user_name,
			COALESCE(users.mobile, '') as user_mobile,
			COALESCE(user_addresses.street, '') as shipping_address,
			COALESCE(user_addresses.city, '') as shipping_city,
			COALESCE(user_addresses.province, '') as shipping_province,
			COALESCE(user_addresses.postal_code, '') as shipping_postal_code,
			COALESCE(user_addresses.recipient_name, '') as recipient_name,
			COALESCE(user_addresses.phone, '') as recipient_phone`).
		Joins("LEFT JOIN users ON users.id = orders.user_id").
		Joins("LEFT JOIN user_addresses ON user_addresses.id = orders.shipping_address_id").
		Where("orders.status IN ?", returnedStatuses)

	// شمارش کل سفارشات مرجوع شده
	var total int64
	if err := baseQuery.Model(&models.Order{}).Count(&total).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های مرجوع شده", err.Error())
		return
	}

	// اجرای کوئری با صفحه‌بندی
	if err := baseQuery.Order("orders.updated_at DESC").Limit(pageSize).Offset(offset).Scan(&ordersWithUser).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش‌های مرجوع شده", err.Error())
		return
	}

	// دریافت تعداد OrderItems برای همه سفارشات در یک کوئری
	var orderItemCounts []struct {
		OrderID uint `json:"order_id"`
		Count   int  `json:"count"`
	}

	if len(ordersWithUser) > 0 {
		orderIDs := make([]uint, len(ordersWithUser))
		for i, order := range ordersWithUser {
			orderIDs[i] = order.ID
		}

		if err := h.DB.Table("order_items").
			Select("order_id, COUNT(*) as count").
			Where("order_id IN ?", orderIDs).
			Group("order_id").
			Scan(&orderItemCounts).Error; err != nil {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت تعداد آیتم‌های سفارش", err.Error())
			return
		}
	}

	// گروه‌بندی تعداد آیتم‌ها بر اساس OrderID
	orderItemCountMap := make(map[uint]int)
	for _, count := range orderItemCounts {
		orderItemCountMap[count.OrderID] = count.Count
	}

	// ساخت آرایه خروجی بهینه شده
	var dtoList []OrderDTO
	for _, order := range ordersWithUser {
		dto := OrderDTO{
			ID:                 order.ID,
			OrderNumber:        order.OrderNumber, // استفاده از OrderNumber
			CustomerName:       order.UserName,
			CustomerPhone:      order.UserMobile,
			CustomerIP:         "",                // IP کاربر - باید از دیتابیس دریافت شود
			UserAgent:          "",                // User-Agent - باید از دیتابیس دریافت شود
			TotalAmount:        order.TotalAmount, // استفاده از TotalAmount به جای FinalAmount
			PaymentMethod:      order.PaymentMethod,
			Status:             order.Status,
			OrderIntegrity:     "verified",
			CreatedAt:          order.CreatedAt,
			ItemsCount:         orderItemCountMap[order.ID],
			ShippingAddress:    order.ShippingAddress,
			ShippingCity:       order.ShippingCity,
			ShippingProvince:   order.ShippingProvince,
			ShippingPostalCode: order.ShippingPostalCode,
			RecipientName:      order.RecipientName,
			RecipientPhone:     order.RecipientPhone,
			ShippingMethod:     "", // روش ارسال - باید از دیتابیس دریافت شود
			TrackingCode:       order.TrackingCode,
		}
		dtoList = append(dtoList, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  dtoList,
		"page":  page,
		"total": total,
	})
}

// ReturnedOrderStats آمار سفارشات مرجوع شده
func (h *OrderHandler) ReturnedOrderStats(c *gin.Context) {
	// فیلتر سفارشات مرجوع شده
	returnedStatuses := []string{"returned", "refunded", "return_pending"}

	// شمارنده‌های وضعیت‌ها
	var totalReturned int64
	if err := h.DB.Model(&models.Order{}).Where("status IN ?", returnedStatuses).Count(&totalReturned).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش کل سفارش‌های مرجوع شده", err.Error())
		return
	}

	var pendingReview int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "return_pending").Count(&pendingReview).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های در انتظار بررسی مرجوعی", err.Error())
		return
	}

	var approved int64
	if err := h.DB.Model(&models.Order{}).Where("status = ?", "returned").Count(&approved).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های مرجوعی تایید شده", err.Error())
		return
	}

	// مبالغ فروش
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	weekStart := now.AddDate(0, 0, -7)
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	var todaySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, returnedStatuses, todayStart).
		Scan(&todaySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش امروز", err.Error())
		return
	}

	var weeklySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, returnedStatuses, weekStart).
		Scan(&weeklySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش هفتگی", err.Error())
		return
	}

	var monthlySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, returnedStatuses, monthStart).
		Scan(&monthlySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش ماه جاری", err.Error())
		return
	}

	// میانگین مبلغ سفارش‌های پرداخت‌شده
	var averageOrder float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(AVG(final_amount),0)").
		Where("is_paid = ? AND status IN ?", true, returnedStatuses).
		Scan(&averageOrder).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه متوسط سفارش", err.Error())
		return
	}

	// آمار شرکت‌های حمل و نقل
	var shippingMethods []struct {
		ShippingMethod string  `json:"shipping_method"`
		Count          int64   `json:"count"`
		Percentage     float64 `json:"percentage"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("shipping_method, COUNT(*) as count").
		Where("status IN ?", returnedStatuses).
		Group("shipping_method").
		Scan(&shippingMethods).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت آمار شرکت‌های حمل و نقل", err.Error())
		return
	}

	// محاسبه درصدها
	for i := range shippingMethods {
		if totalReturned > 0 {
			shippingMethods[i].Percentage = float64(shippingMethods[i].Count) / float64(totalReturned) * 100
		}
	}

	// آمار توزیع وضعیت‌ها
	var statusDistribution []struct {
		Status string `json:"status"`
		Count  int64  `json:"count"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Where("status IN ?", returnedStatuses).
		Group("status").
		Scan(&statusDistribution).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت توزیع وضعیت‌ها", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"totalReturned":      totalReturned,
		"pendingReview":      pendingReview,
		"approved":           approved,
		"todaySales":         todaySales,
		"weeklySales":        weeklySales,
		"monthlySales":       monthlySales,
		"averageOrder":       averageOrder,
		"shippingMethods":    shippingMethods,
		"statusDistribution": statusDistribution,
	})
}

// ListRefundedOrders لیست سفارشات مسترد شده برای ادمین
func (h *OrderHandler) ListRefundedOrders(c *gin.Context) {
	// --- Pagination Handling ---
	pageParam := c.DefaultQuery("page", "1")
	sizeParam := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(sizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	// فیلتر سفارشات مسترد شده
	// شامل سفارشاتی که وضعیتشان به "refunded" تغییر کرده
	refundedStatuses := []string{"refunded"}

	// ساخت کوئری اصلی با JOIN برای سفارشات مسترد شده
	type OrderWithUser struct {
		models.Order
		UserName           string `json:"user_name"`
		UserMobile         string `json:"user_mobile"`
		ShippingAddress    string `json:"shipping_address"`
		ShippingCity       string `json:"shipping_city"`
		ShippingProvince   string `json:"shipping_province"`
		ShippingPostalCode string `json:"shipping_postal_code"`
		RecipientName      string `json:"recipient_name"`
		RecipientPhone     string `json:"recipient_phone"`
	}

	var ordersWithUser []OrderWithUser

	// ساخت کوئری اصلی با JOIN
	baseQuery := h.DB.Table("orders").
		Select(`orders.*,
			COALESCE(users.name, '') as user_name,
			COALESCE(users.mobile, '') as user_mobile,
			COALESCE(user_addresses.street, '') as shipping_address,
			COALESCE(user_addresses.city, '') as shipping_city,
			COALESCE(user_addresses.province, '') as shipping_province,
			COALESCE(user_addresses.postal_code, '') as shipping_postal_code,
			COALESCE(user_addresses.recipient_name, '') as recipient_name,
			COALESCE(user_addresses.phone, '') as recipient_phone`).
		Joins("LEFT JOIN users ON users.id = orders.user_id").
		Joins("LEFT JOIN user_addresses ON user_addresses.id = orders.shipping_address_id").
		Where("orders.status IN ?", refundedStatuses)

	// شمارش کل سفارشات مسترد شده
	var total int64
	if err := baseQuery.Model(&models.Order{}).Count(&total).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های مسترد شده", err.Error())
		return
	}

	// اجرای کوئری با صفحه‌بندی
	if err := baseQuery.Order("orders.updated_at DESC").Limit(pageSize).Offset(offset).Scan(&ordersWithUser).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش‌های مسترد شده", err.Error())
		return
	}

	// دریافت تعداد OrderItems برای همه سفارشات در یک کوئری
	var orderItemCounts []struct {
		OrderID uint `json:"order_id"`
		Count   int  `json:"count"`
	}

	if len(ordersWithUser) > 0 {
		orderIDs := make([]uint, len(ordersWithUser))
		for i, order := range ordersWithUser {
			orderIDs[i] = order.ID
		}

		if err := h.DB.Table("order_items").
			Select("order_id, COUNT(*) as count").
			Where("order_id IN ?", orderIDs).
			Group("order_id").
			Scan(&orderItemCounts).Error; err != nil {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت تعداد آیتم‌های سفارش", err.Error())
			return
		}
	}

	// گروه‌بندی تعداد آیتم‌ها بر اساس OrderID
	orderItemCountMap := make(map[uint]int)
	for _, count := range orderItemCounts {
		orderItemCountMap[count.OrderID] = count.Count
	}

	// ساخت آرایه خروجی بهینه شده
	var dtoList []OrderDTO
	for _, order := range ordersWithUser {
		dto := OrderDTO{
			ID:                 order.ID,
			OrderNumber:        order.OrderNumber, // استفاده از OrderNumber
			CustomerName:       order.UserName,
			CustomerPhone:      order.UserMobile,
			CustomerIP:         "",                // IP کاربر - باید از دیتابیس دریافت شود
			UserAgent:          "",                // User-Agent - باید از دیتابیس دریافت شود
			TotalAmount:        order.TotalAmount, // استفاده از TotalAmount به جای FinalAmount
			PaymentMethod:      order.PaymentMethod,
			Status:             order.Status,
			OrderIntegrity:     "verified",
			CreatedAt:          order.CreatedAt,
			ItemsCount:         orderItemCountMap[order.ID],
			ShippingAddress:    order.ShippingAddress,
			ShippingCity:       order.ShippingCity,
			ShippingProvince:   order.ShippingProvince,
			ShippingPostalCode: order.ShippingPostalCode,
			RecipientName:      order.RecipientName,
			RecipientPhone:     order.RecipientPhone,
			ShippingMethod:     "", // روش ارسال - باید از دیتابیس دریافت شود
			TrackingCode:       order.TrackingCode,
		}
		dtoList = append(dtoList, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  dtoList,
		"page":  page,
		"total": total,
	})
}

// RefundedOrderStats آمار سفارشات مسترد شده
func (h *OrderHandler) RefundedOrderStats(c *gin.Context) {
	// فیلتر سفارشات مسترد شده
	refundedStatuses := []string{"refunded"}

	// شمارنده‌های وضعیت‌ها
	var totalRefunded int64
	if err := h.DB.Model(&models.Order{}).Where("status IN ?", refundedStatuses).Count(&totalRefunded).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش کل سفارش‌های مسترد شده", err.Error())
		return
	}

	// مبالغ فروش
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	weekStart := now.AddDate(0, 0, -7)
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	var todaySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, refundedStatuses, todayStart).
		Scan(&todaySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش امروز", err.Error())
		return
	}

	var weeklySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, refundedStatuses, weekStart).
		Scan(&weeklySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش هفتگی", err.Error())
		return
	}

	var monthlySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, refundedStatuses, monthStart).
		Scan(&monthlySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش ماه جاری", err.Error())
		return
	}

	// میانگین مبلغ سفارش‌های پرداخت‌شده
	var averageOrder float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(AVG(final_amount),0)").
		Where("is_paid = ? AND status IN ?", true, refundedStatuses).
		Scan(&averageOrder).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه متوسط سفارش", err.Error())
		return
	}

	// آمار روش‌های پرداخت
	var paymentMethods []struct {
		PaymentMethod string  `json:"payment_method"`
		Count         int64   `json:"count"`
		Percentage    float64 `json:"percentage"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("payment_method, COUNT(*) as count").
		Where("status IN ?", refundedStatuses).
		Group("payment_method").
		Scan(&paymentMethods).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت آمار روش‌های پرداخت", err.Error())
		return
	}

	// محاسبه درصدها
	for i := range paymentMethods {
		if totalRefunded > 0 {
			paymentMethods[i].Percentage = float64(paymentMethods[i].Count) / float64(totalRefunded) * 100
		}
	}

	// آمار توزیع وضعیت‌ها
	var statusDistribution []struct {
		Status string `json:"status"`
		Count  int64  `json:"count"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Where("status IN ?", refundedStatuses).
		Group("status").
		Scan(&statusDistribution).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت توزیع وضعیت‌ها", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"totalRefunded":      totalRefunded,
		"todaySales":         todaySales,
		"weeklySales":        weeklySales,
		"monthlySales":       monthlySales,
		"averageOrder":       averageOrder,
		"paymentMethods":     paymentMethods,
		"statusDistribution": statusDistribution,
	})
}

// ListCancelledOrders لیست سفارشات لغو شده برای ادمین
func (h *OrderHandler) ListCancelledOrders(c *gin.Context) {
	// --- Pagination Handling ---
	pageParam := c.DefaultQuery("page", "1")
	sizeParam := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(sizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	// فیلتر سفارشات لغو شده
	cancelledStatuses := []string{"cancelled"}

	// ساخت کوئری اصلی با JOIN برای سفارشات لغو شده
	type OrderWithUser struct {
		models.Order
		UserName           string `json:"user_name"`
		UserMobile         string `json:"user_mobile"`
		ShippingAddress    string `json:"shipping_address"`
		ShippingCity       string `json:"shipping_city"`
		ShippingProvince   string `json:"shipping_province"`
		ShippingPostalCode string `json:"shipping_postal_code"`
		RecipientName      string `json:"recipient_name"`
		RecipientPhone     string `json:"recipient_phone"`
	}

	var ordersWithUser []OrderWithUser

	// ساخت کوئری اصلی با JOIN
	baseQuery := h.DB.Table("orders").
		Select(`orders.*,
			COALESCE(users.name, '') as user_name,
			COALESCE(users.mobile, '') as user_mobile,
			COALESCE(user_addresses.street, '') as shipping_address,
			COALESCE(user_addresses.city, '') as shipping_city,
			COALESCE(user_addresses.province, '') as shipping_province,
			COALESCE(user_addresses.postal_code, '') as shipping_postal_code,
			COALESCE(user_addresses.recipient_name, '') as recipient_name,
			COALESCE(user_addresses.phone, '') as recipient_phone`).
		Joins("LEFT JOIN users ON users.id = orders.user_id").
		Joins("LEFT JOIN user_addresses ON user_addresses.id = orders.shipping_address_id").
		Where("orders.status IN ?", cancelledStatuses)

	// شمارش کل سفارشات لغو شده
	var total int64
	if err := baseQuery.Model(&models.Order{}).Count(&total).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش سفارش‌های لغو شده", err.Error())
		return
	}

	// اجرای کوئری با صفحه‌بندی
	if err := baseQuery.Order("orders.updated_at DESC").Limit(pageSize).Offset(offset).Scan(&ordersWithUser).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت سفارش‌های لغو شده", err.Error())
		return
	}

	// دریافت تعداد OrderItems برای همه سفارشات در یک کوئری
	var orderItemCounts []struct {
		OrderID uint `json:"order_id"`
		Count   int  `json:"count"`
	}

	if len(ordersWithUser) > 0 {
		orderIDs := make([]uint, len(ordersWithUser))
		for i, order := range ordersWithUser {
			orderIDs[i] = order.ID
		}

		if err := h.DB.Table("order_items").
			Select("order_id, COUNT(*) as count").
			Where("order_id IN ?", orderIDs).
			Group("order_id").
			Scan(&orderItemCounts).Error; err != nil {
			utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت تعداد آیتم‌های سفارش", err.Error())
			return
		}
	}

	// گروه‌بندی تعداد آیتم‌ها بر اساس OrderID
	orderItemCountMap := make(map[uint]int)
	for _, count := range orderItemCounts {
		orderItemCountMap[count.OrderID] = count.Count
	}

	// ساخت آرایه خروجی بهینه شده
	var dtoList []OrderDTO
	for _, order := range ordersWithUser {
		dto := OrderDTO{
			ID:                 order.ID,
			OrderNumber:        order.OrderNumber, // استفاده از OrderNumber
			CustomerName:       order.UserName,
			CustomerPhone:      order.UserMobile,
			CustomerIP:         "",                // IP کاربر - باید از دیتابیس دریافت شود
			UserAgent:          "",                // User-Agent - باید از دیتابیس دریافت شود
			TotalAmount:        order.TotalAmount, // استفاده از TotalAmount به جای FinalAmount
			PaymentMethod:      order.PaymentMethod,
			Status:             order.Status,
			OrderIntegrity:     "verified",
			CreatedAt:          order.CreatedAt,
			ItemsCount:         orderItemCountMap[order.ID],
			ShippingAddress:    order.ShippingAddress,
			ShippingCity:       order.ShippingCity,
			ShippingProvince:   order.ShippingProvince,
			ShippingPostalCode: order.ShippingPostalCode,
			RecipientName:      order.RecipientName,
			RecipientPhone:     order.RecipientPhone,
			ShippingMethod:     "", // روش ارسال - باید از دیتابیس دریافت شود
			TrackingCode:       order.TrackingCode,
		}
		dtoList = append(dtoList, dto)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  dtoList,
		"page":  page,
		"total": total,
	})
}

// CancelledOrderStats آمار سفارشات لغو شده
func (h *OrderHandler) CancelledOrderStats(c *gin.Context) {
	// فیلتر سفارشات لغو شده
	cancelledStatuses := []string{"cancelled"}

	// شمارنده‌های وضعیت‌ها
	var totalCancelled int64
	if err := h.DB.Model(&models.Order{}).Where("status IN ?", cancelledStatuses).Count(&totalCancelled).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در شمارش کل سفارش‌های لغو شده", err.Error())
		return
	}

	// مبالغ فروش
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	weekStart := now.AddDate(0, 0, -7)
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	var todaySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, cancelledStatuses, todayStart).
		Scan(&todaySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش امروز", err.Error())
		return
	}

	var weeklySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, cancelledStatuses, weekStart).
		Scan(&weeklySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش هفتگی", err.Error())
		return
	}

	var monthlySales float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(SUM(final_amount),0)").
		Where("is_paid = ? AND status IN ? AND created_at >= ?", true, cancelledStatuses, monthStart).
		Scan(&monthlySales).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه فروش ماه جاری", err.Error())
		return
	}

	// میانگین مبلغ سفارش‌های پرداخت‌شده
	var averageOrder float64
	if err := h.DB.Model(&models.Order{}).
		Select("COALESCE(AVG(final_amount),0)").
		Where("is_paid = ? AND status IN ?", true, cancelledStatuses).
		Scan(&averageOrder).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در محاسبه متوسط سفارش", err.Error())
		return
	}

	// آمار روش‌های پرداخت
	var paymentMethods []struct {
		PaymentMethod string  `json:"payment_method"`
		Count         int64   `json:"count"`
		Percentage    float64 `json:"percentage"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("payment_method, COUNT(*) as count").
		Where("status IN ?", cancelledStatuses).
		Group("payment_method").
		Scan(&paymentMethods).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت آمار روش‌های پرداخت", err.Error())
		return
	}

	// محاسبه درصدها
	for i := range paymentMethods {
		if totalCancelled > 0 {
			paymentMethods[i].Percentage = float64(paymentMethods[i].Count) / float64(totalCancelled) * 100
		}
	}

	// آمار توزیع وضعیت‌ها
	var statusDistribution []struct {
		Status string `json:"status"`
		Count  int64  `json:"count"`
	}

	if err := h.DB.Model(&models.Order{}).
		Select("status, COUNT(*) as count").
		Where("status IN ?", cancelledStatuses).
		Group("status").
		Scan(&statusDistribution).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "DB_ERROR", "خطا در دریافت توزیع وضعیت‌ها", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"totalCancelled":     totalCancelled,
		"todaySales":         todaySales,
		"weeklySales":        weeklySales,
		"monthlySales":       monthlySales,
		"averageOrder":       averageOrder,
		"paymentMethods":     paymentMethods,
		"statusDistribution": statusDistribution,
	})
}
