package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ReviewHandler struct {
	db   *gorm.DB
	uowf unitofwork.UnitOfWorkFactory
}

func NewReviewHandler(db *gorm.DB) *ReviewHandler { // backward compatible
	return &ReviewHandler{db: db}
}

func NewReviewHandlerWithUoW(db *gorm.DB, uowf unitofwork.UnitOfWorkFactory) *ReviewHandler {
	return &ReviewHandler{db: db, uowf: uowf}
}

// getProductThumbnail جستجو برای thumbnail محصول از media variants
func (h *ReviewHandler) getProductThumbnail(productID int) string {
	var mediaFile models.MediaFile

	// جستجو برای media file مربوط به محصول (بر اساس category و نام فایل)
	err := h.db.Where("category = ? AND (file_name LIKE ? OR file_name LIKE ?)",
		"products",
		fmt.Sprintf("%%product_%d%%", productID),
		fmt.Sprintf("%%_%d_%%", productID)).
		First(&mediaFile).Error

	if err != nil {
		return ""
	}

	// جستجو برای thumbnail variant
	var thumbnail models.MediaVariant
	err = h.db.Where("media_id = ? AND purpose = ?", mediaFile.ID, "thumbnail").
		First(&thumbnail).Error

	if err != nil {
		return ""
	}

	// اطمینان از حذف اسلش اضافه برای جلوگیری از //
	cleaned := strings.TrimPrefix(thumbnail.FilePath, "/")
	return "/media/" + cleaned
}

// CreateReview handles the creation of a new review
func (h *ReviewHandler) CreateReview(ctx *gin.Context) {
	rating, _ := strconv.Atoi(ctx.PostForm("rating"))
	productID, _ := strconv.Atoi(ctx.PostForm("product_id"))
	comment := ctx.PostForm("comment")
	title := ctx.PostForm("title")
	pros := ctx.PostForm("pros")
	cons := ctx.PostForm("cons")
	// اطلاعات اختیاری پروفایل برای تکمیل یکباره
	providedFullName := strings.TrimSpace(ctx.PostForm("full_name"))
	providedMobile := strings.TrimSpace(ctx.PostForm("mobile"))
	// anonymous := ctx.PostForm("anonymous") == "true" // برای نمایش در سایت استفاده می‌شود، در پنل ادمین همیشه نام واقعی نمایش داده می‌شود

	// دریافت user_id از context
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "کاربر احراز هویت نشده"})
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user_id نامعتبر است"})
		return
	}

	// چک مقدار rating
	if rating < 1 || rating > 5 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}

	// ابتدا نام کاربر را از جدول users بگیریم
	var user models.User
	if err := h.db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, utils.New("USER_NOT_FOUND", "کاربر یافت نشد", nil))
		return
	}

	// اگر نام کاربر خالی است و full_name ارسال شده، یکبار پروفایل را تکمیل کنیم
	needUpdate := false
	updates := map[string]interface{}{}
	if user.Name == "" && providedFullName != "" {
		updates["name"] = providedFullName
		needUpdate = true
	}
	if user.Mobile == "" && providedMobile != "" {
		updates["mobile"] = providedMobile
		needUpdate = true
	}
	if needUpdate {
		_ = h.db.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error
		// بارگذاری مجدد برای استفاده ادامه مسیر
		_ = h.db.Where("id = ?", userID).First(&user).Error
	}

	// ابتدا بررسی کنیم که آیا مشتری با این userID وجود دارد یا نه
	var existingCustomer models.Customer
	if err := h.db.Where("id = ?", userID).First(&existingCustomer).Error; err != nil {
		// مشتری وجود ندارد، آن را ایجاد کنیم
		// نام واقعی را از جدول users بگیریم
		existingCustomer = models.Customer{
			Model: gorm.Model{ID: userID},
			Name:  user.Name, // نام واقعی از جدول users
			Email: user.Email,
			Phone: user.Mobile,
		}
		if err := h.db.Create(&existingCustomer).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
			return
		}
	} else {
		// مشتری وجود دارد، اطلاعات را به‌روزرسانی کنیم
		existingCustomer.Name = user.Name // نام واقعی از جدول users
		existingCustomer.Email = user.Email
		if strings.TrimSpace(user.Mobile) != "" {
			existingCustomer.Phone = user.Mobile
		}
		if err := h.db.Save(&existingCustomer).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
			return
		}
	}

	// ساخت مدل Review
	review := models.Review{
		CustomerID: userID,
		Rating:     rating,
		ProductID:  uint(productID),
		Title:      strings.TrimSpace(title),
		Comment:    comment,
		Pros:       strings.TrimSpace(pros),
		Cons:       strings.TrimSpace(cons),
		Status:     "pending",
		IPAddress:  ctx.ClientIP(),
		UserAgent:  ctx.Request.UserAgent(),
	}

	// آپلود تصاویر
	form, _ := ctx.MultipartForm()
	files := form.File["images"]
	var imagePaths []string
	// مقادیر پیش‌فرض و قابل‌تغییر از تنظیمات سیستم
	maxFiles := 7
	maxFileSizeMB := 50
	allowedImage := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	allowedVideo := map[string]bool{".mp4": true, ".webm": true, ".mov": true}

	// تلاش برای خواندن از تنظیمات (در صورت موجود بودن در context)
	if v := ctx.GetString("reviews.max_files_per_review"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			maxFiles = n
		}
	}
	if v := ctx.GetString("reviews.max_file_size_mb"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			maxFileSizeMB = n
		}
	}
	for i, file := range files {
		if i >= maxFiles {
			break
		}
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !allowedImage[ext] {
			continue
		}
		if file.Size > int64(maxFileSizeMB*1024*1024) {
			continue
		}
		filename := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
		saveDir := filepath.Join("public", "uploads", "media", "customer", "images")
		_ = os.MkdirAll(saveDir, 0755)
		savePath := filepath.Join(saveDir, filename)
		if err := ctx.SaveUploadedFile(file, savePath); err == nil {
			imagePaths = append(imagePaths, "/uploads/media/customer/images/"+filename)
		}
	}
	// تبدیل به JSON معتبر
	imagesJSON, _ := json.Marshal(imagePaths)
	review.Images = datatypes.JSON(imagesJSON)

	// اگر هیچ عکسی نبود:
	if len(imagePaths) == 0 {
		review.Images = datatypes.JSON([]byte("[]"))
	}

	// آپلود ویدیو (اختیاری)
	videoFiles := form.File["videos"]
	var videoPaths []string
	for i, file := range videoFiles {
		if i >= maxFiles {
			break
		}
		// محدودیت انواع فایل و اندازه طبق تنظیمات
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !allowedVideo[ext] {
			continue
		}
		if file.Size > int64(maxFileSizeMB*1024*1024) {
			continue
		}
		filename := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
		saveDir := filepath.Join("public", "uploads", "media", "customer", "videos")
		_ = os.MkdirAll(saveDir, 0755)
		savePath := filepath.Join(saveDir, filename)
		if err := ctx.SaveUploadedFile(file, savePath); err == nil {
			videoPaths = append(videoPaths, "/uploads/media/customer/videos/"+filename)
		}
	}
	// اگر مدل Review فیلد Videos دارد:
	// review.Videos = videoPaths

	// Merge image and video paths into single slice for attachments
	attachments := append([]string{}, imagePaths...)
	attachments = append(attachments, videoPaths...)

	attachmentsJSON, _ := json.Marshal(attachments)
	review.Images = datatypes.JSON(attachmentsJSON)

	if len(attachments) == 0 {
		review.Images = datatypes.JSON([]byte("[]"))
	}

	// ذخیره با UnitOfWork + Repository
	uow := unitofwork.NewGormUnitOfWork(h.db)
	if err := uow.BeginTx(ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	defer func() { _ = uow.Rollback() }()

	if err := uow.ReviewRepository().Create(ctx, &review); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	if err := uow.Commit(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, review)
}

// GetPublicReviewsByProduct
// دریافت فهرست نظرات تأییدشدهٔ یک محصول برای نمایش عمومی (بدون نیاز به احراز هویت)
// این خروجی حاوی اطلاعات حداقلی و ایمن برای فرانت‌اند است و ایمیل/اطلاعات حساس کاربر را بازنمی‌گرداند.
func (h *ReviewHandler) GetPublicReviewsByProduct(ctx *gin.Context) {
	// شناسه محصول
	productIDStr := ctx.Param("id")
	if productIDStr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), "product id is required"))
		return
	}

	// فیلترهای اختیاری
	ratingFilter := ctx.Query("rating")

	var reviews []models.Review
	query := h.db.Model(&models.Review{}).
		Preload("Customer").
		Where("product_id = ? AND status = ?", productIDStr, "approved").
		Order("created_at DESC")

	if ratingFilter != "" {
		query = query.Where("rating = ?", ratingFilter)
	}

	if err := query.Find(&reviews).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	// نگاشت به ساختار ایمن و مورد نیاز UI
	safe := make([]map[string]interface{}, 0, len(reviews))
	for _, r := range reviews {
		// استخراج تصاویر از JSON
		var imgs []string
		if len(r.Images) > 0 {
			_ = json.Unmarshal(r.Images, &imgs)
		}

		item := map[string]interface{}{
			"id":            r.ID,
			"user_name":     r.Customer.Name,
			"rating":        r.Rating,
			"title":         nil, // در مدل فعلی فیلد عنوان ذخیره نمی‌شود
			"comment":       r.Comment,
			"images":        imgs,
			"is_verified":   r.HasPurchased,
			"helpful_count": r.HelpfulCount,
			"created_at":    r.CreatedAt,
		}
		if strings.TrimSpace(r.AdminReply) != "" {
			item["reply"] = map[string]interface{}{
				"content":    r.AdminReply,
				"created_at": r.UpdatedAt,
			}
		}

		safe = append(safe, item)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"reviews": safe,
	})
}

// GetReviews retrieves all reviews with pagination and filters
func (h *ReviewHandler) GetReviews(ctx *gin.Context) {
	var reviews []models.Review
	query := h.db.Model(&models.Review{}).Preload("Customer").Preload("Product").Preload("Product.Images")

	// Apply filters
	if status := ctx.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}
	if rating := ctx.Query("rating"); rating != "" {
		query = query.Where("rating = ?", rating)
	}
	if search := ctx.Query("search"); search != "" {
		query = query.Where("comment LIKE ?", "%"+search+"%")
	}
	if productID := ctx.Query("product_id"); productID != "" {
		query = query.Where("product_id = ?", productID)
	}

	// Pagination
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))
	offset := (page - 1) * perPage

	var total int64
	query.Count(&total)

	query = query.Offset(offset).Limit(perPage)
	if err := query.Find(&reviews).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	// Map reviews to include only needed customer fields
	mapped := make([]map[string]interface{}, len(reviews))
	for i, r := range reviews {
		mapped[i] = map[string]interface{}{
			"id":         r.ID,
			"product_id": r.ProductID,
			"product": map[string]interface{}{
				"id":    r.Product.ID,
				"name":  r.Product.Name,
				"price": r.Product.Price,
				"image": func() string {
					// اگر مسیر نسبی است همان را برمی‌گردانیم، در فرانت‌اند با آدرس سرور ترکیب می‌شود
					return r.Product.Image
				}(),
				"thumbnail": func() string {
					// اولویت اول: thumbnail از media_variants
					if thumbnailURL := h.getProductThumbnail(r.Product.ID); thumbnailURL != "" {
						return thumbnailURL
					}

					// اولویت دوم: اولین تصویر محصول (در صورت موجود بودن)
					if len(r.Product.Images) > 0 {
						return r.Product.Images[0].ImageURL
					}

					// اولویت سوم: عکس اصلی محصول
					if r.Product.Image != "" {
						return r.Product.Image
					}

					// پیش‌فرض
					return "/default-product.svg"
				}(),
			},
			"customer": map[string]interface{}{
				"id":    r.CustomerID,
				"name":  r.Customer.Name,
				"email": r.Customer.Email,
				"phone": r.Customer.Phone,
			},
			"rating":        r.Rating,
			"comment":       r.Comment,
			"status":        r.Status,
			"images":        r.Images,
			"admin_reply":   r.AdminReply,
			"ip_address":    r.IPAddress,
			"user_agent":    r.UserAgent,
			"device_info":   r.DeviceInfo,
			"has_purchased": r.HasPurchased,
			"helpful_count": r.HelpfulCount,
			"created_at":    r.CreatedAt,
			"updated_at":    r.UpdatedAt,
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"reviews": mapped,
		"total":   total,
		"page":    page,
		"perPage": perPage,
	})
}

// GetReview retrieves a single review by ID
func (h *ReviewHandler) GetReview(ctx *gin.Context) {
	id := ctx.Param("id")
	var review models.Review

	if err := h.db.First(&review, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}

	ctx.JSON(http.StatusOK, review)
}

// UpdateReview updates an existing review
func (h *ReviewHandler) UpdateReview(ctx *gin.Context) {
	id := ctx.Param("id")
	var review models.Review

	if err := h.db.First(&review, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}

	if err := ctx.ShouldBindJSON(&review); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	if err := h.db.Save(&review).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, review)
}

// DeleteReview deletes a review
func (h *ReviewHandler) DeleteReview(ctx *gin.Context) {
	id := ctx.Param("id")
	var review models.Review

	if err := h.db.First(&review, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}

	if err := h.db.Delete(&review).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}

// ApproveReview approves a review
func (h *ReviewHandler) ApproveReview(ctx *gin.Context) {
	id := ctx.Param("id")
	var review models.Review

	if err := h.db.First(&review, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}

	review.Status = "approved"
	if err := h.db.Save(&review).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, review)
}

// RejectReview rejects a review
func (h *ReviewHandler) RejectReview(ctx *gin.Context) {
	id := ctx.Param("id")
	var review models.Review

	if err := h.db.First(&review, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}

	review.Status = "rejected"
	if err := h.db.Save(&review).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, review)
}

// AddReply adds an admin reply to a review
func (h *ReviewHandler) AddReply(ctx *gin.Context) {
	id := ctx.Param("id")
	var review models.Review
	var reply struct {
		Reply string `json:"reply"`
	}

	if err := ctx.ShouldBindJSON(&reply); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	if err := h.db.First(&review, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}

	review.AdminReply = reply.Reply
	if err := h.db.Save(&review).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, review)
}

// MarkHelpful marks a review as helpful
func (h *ReviewHandler) MarkHelpful(ctx *gin.Context) {
	id := ctx.Param("id")
	var review models.Review
	if err := h.db.First(&review, id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}

	// نیاز به کاربر احراز هویت شده برای ثبت رای مفید
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", utils.GetErrorMessage("UNAUTHORIZED"), nil))
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.New("UNAUTHORIZED", utils.GetErrorMessage("UNAUTHORIZED"), nil))
		return
	}

	// ثبت رأی یکتا با Repository + UoW
	uow := unitofwork.NewGormUnitOfWork(h.db)
	if err := uow.BeginTx(ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	defer func() { _ = uow.Rollback() }()

	has, err := uow.ReviewRepository().HasHelpfulVote(ctx, review.ID, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	if has {
		ctx.JSON(http.StatusOK, review)
		return
	}
	if err := uow.ReviewRepository().CreateHelpfulVote(ctx, review.ID, userID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	if err := uow.ReviewRepository().IncrementHelpfulCount(ctx, review.ID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	if err := uow.Commit(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	_ = h.db.First(&review, review.ID).Error
	ctx.JSON(http.StatusOK, review)
}
