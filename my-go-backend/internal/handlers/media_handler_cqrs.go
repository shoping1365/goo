package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"my-go-backend/internal/cqrs/bus"
	"my-go-backend/internal/cqrs/commands"
	"my-go-backend/internal/cqrs/queries"
	"my-go-backend/internal/utils"
)

// MediaCQRSHandler هندلر جدید مدیریت رسانه با استفاده از CQRS
// این handler از Command و Query buses برای جداسازی عملیات خواندن و نوشتن استفاده می‌کند
type MediaCQRSHandler struct {
	commandBus bus.CommandBus
	queryBus   bus.QueryBus
}

// NewMediaCQRSHandler ایجاد نمونه جدید از MediaCQRSHandler
func NewMediaCQRSHandler(cmdBus bus.CommandBus, queryBus bus.QueryBus) *MediaCQRSHandler {
	return &MediaCQRSHandler{
		commandBus: cmdBus,
		queryBus:   queryBus,
	}
}

// GetMediaFile دریافت اطلاعات یک فایل رسانه
// GET /api/media/cqrs/{id}
func (h *MediaCQRSHandler) GetMediaFile(w http.ResponseWriter, r *http.Request) {
	// استخراج ID از URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/media/cqrs/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "INVALID_ID", "شناسه نامعتبر است")
		return
	}

	// استفاده از Query Bus برای دریافت اطلاعات
	file, err := h.queryBus.GetMediaFileByID(r.Context(), uint(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "NOT_FOUND", "فایل یافت نشد")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, file)
}

// GetMediaFileWithVariants دریافت فایل رسانه به همراه نسخه‌ها
// GET /api/media/cqrs/{id}/with-variants
func (h *MediaCQRSHandler) GetMediaFileWithVariants(w http.ResponseWriter, r *http.Request) {
	// استخراج ID از URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		utils.RespondWithError(w, http.StatusBadRequest, "INVALID_URL", "آدرس نامعتبر است")
		return
	}

	idStr := parts[4]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "INVALID_ID", "شناسه نامعتبر است")
		return
	}

	// استفاده از Query Bus
	file, err := h.queryBus.GetMediaFileWithVariants(r.Context(), uint(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "NOT_FOUND", "فایل یافت نشد")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, file)
}

// ListMediaFilesPaged لیست فایل‌های رسانه به صورت صفحه‌بندی شده
// GET /api/media/cqrs?page=1&pageSize=20&category=products&fileType=image/jpeg
func (h *MediaCQRSHandler) ListMediaFilesPaged(w http.ResponseWriter, r *http.Request) {
	// پارس پارامترها
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	// ایجاد فیلتر
	filter := queries.MediaFileFilter{
		Category: r.URL.Query().Get("category"),
		FileType: r.URL.Query().Get("fileType"),
	}

	// compressed filter
	if compressedStr := r.URL.Query().Get("compressed"); compressedStr != "" {
		compressed := compressedStr == "true" || compressedStr == "1"
		filter.Compressed = &compressed
	}

	// user filter
	if userIDStr := r.URL.Query().Get("userID"); userIDStr != "" {
		if userID, err := strconv.ParseUint(userIDStr, 10, 32); err == nil {
			uid := uint(userID)
			filter.UserID = &uid
		}
	}

	// استفاده از Query Bus
	files, total, err := h.queryBus.GetMediaFilesPaged(r.Context(), page, pageSize, filter)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "QUERY_ERROR", "خطا در دریافت لیست")
		return
	}

	// پاسخ با اطلاعات صفحه‌بندی
	response := map[string]interface{}{
		"data":       files,
		"total":      total,
		"page":       page,
		"pageSize":   pageSize,
		"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

// SearchMediaFiles جستجو در فایل‌های رسانه
// GET /api/media/cqrs/search?q=example
func (h *MediaCQRSHandler) SearchMediaFiles(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "QUERY_REQUIRED", "عبارت جستجو الزامی است")
		return
	}

	// استفاده از Query Bus
	files, err := h.queryBus.SearchMediaFiles(r.Context(), query)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "SEARCH_ERROR", "خطا در جستجو")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"data":  files,
		"query": query,
		"count": len(files),
	})
}

// CreateMediaFile ایجاد فایل رسانه جدید
// POST /api/media/cqrs
func (h *MediaCQRSHandler) CreateMediaFile(w http.ResponseWriter, r *http.Request) {
	var cmd commands.CreateMediaFileCommand

	// پارس body
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "INVALID_JSON", "داده‌های نامعتبر")
		return
	}

	// دریافت user ID از context (اگر احراز هویت شده باشد)
	if userID, ok := r.Context().Value("user_id").(uint); ok {
		cmd.UploadedBy = &userID
	}

	// استفاده از Command Bus
	file, err := h.commandBus.CreateMediaFile(r.Context(), cmd)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "CREATE_ERROR", err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, file)
}

// UpdateMediaFile به‌روزرسانی فایل رسانه
// PUT /api/media/cqrs/{id}
func (h *MediaCQRSHandler) UpdateMediaFile(w http.ResponseWriter, r *http.Request) {
	// استخراج ID از URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/media/cqrs/")
	idStr = strings.TrimSuffix(idStr, "/update")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "INVALID_ID", "شناسه نامعتبر است")
		return
	}

	// پارس body
	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "INVALID_JSON", "داده‌های نامعتبر")
		return
	}

	// ایجاد command
	cmd := commands.UpdateMediaFileCommand{
		ID: uint(id),
	}

	// تنظیم فیلدها (فقط اگر در body موجود باشند)
	if title, ok := updateData["title"].(string); ok {
		cmd.Title = &title
	}
	if altText, ok := updateData["altText"].(string); ok {
		cmd.AltText = &altText
	}
	if displayTitle, ok := updateData["displayTitle"].(string); ok {
		cmd.DisplayTitle = &displayTitle
	}
	if shortDescription, ok := updateData["shortDescription"].(string); ok {
		cmd.ShortDescription = &shortDescription
	}
	if description, ok := updateData["description"].(string); ok {
		cmd.Description = &description
	}
	if category, ok := updateData["category"].(string); ok {
		cmd.Category = &category
	}

	// استفاده از Command Bus
	if err := h.commandBus.UpdateMediaFile(r.Context(), cmd); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "UPDATE_ERROR", err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "فایل با موفقیت به‌روزرسانی شد",
	})
}

// DeleteMediaFile حذف فایل رسانه
// DELETE /api/media/cqrs/{id}
func (h *MediaCQRSHandler) DeleteMediaFile(w http.ResponseWriter, r *http.Request) {
	// استخراج ID از URL
	idStr := strings.TrimPrefix(r.URL.Path, "/api/media/cqrs/")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "INVALID_ID", "شناسه نامعتبر است")
		return
	}

	// ایجاد command
	cmd := commands.DeleteMediaFileCommand{
		ID: uint(id),
	}

	// استفاده از Command Bus
	if err := h.commandBus.DeleteMediaFile(r.Context(), cmd); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "DELETE_ERROR", err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "فایل با موفقیت حذف شد",
	})
}

// CompressMediaFile ایجاد job فشرده‌سازی برای فایل رسانه
// POST /api/media/cqrs/{id}/compress
func (h *MediaCQRSHandler) CompressMediaFile(w http.ResponseWriter, r *http.Request) {
	// استخراج ID از URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		utils.RespondWithError(w, http.StatusBadRequest, "INVALID_URL", "آدرس نامعتبر است")
		return
	}

	idStr := parts[4]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "INVALID_ID", "شناسه نامعتبر است")
		return
	}

	// پارس body
	var compressData struct {
		TargetFormat  string `json:"targetFormat"`
		TargetQuality string `json:"targetQuality"`
	}

	if err := json.NewDecoder(r.Body).Decode(&compressData); err != nil {
		// اگر body خالی باشد، مقادیر پیش‌فرض
		compressData.TargetFormat = "webp"
		compressData.TargetQuality = "medium"
	}

	// دریافت user ID از context
	var requestedBy *uint
	if userID, ok := r.Context().Value("user_id").(uint); ok {
		requestedBy = &userID
	}

	// ایجاد command
	cmd := commands.CompressMediaFileCommand{
		MediaID:       uint(id),
		TargetFormat:  compressData.TargetFormat,
		TargetQuality: compressData.TargetQuality,
		RequestedBy:   requestedBy,
	}

	// استفاده از Command Bus
	job, err := h.commandBus.CompressMediaFile(r.Context(), cmd)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "COMPRESS_ERROR", err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusAccepted, map[string]interface{}{
		"success": true,
		"message": "درخواست فشرده‌سازی ثبت شد",
		"jobID":   job.ID,
		"status":  job.Status,
	})
}

// GetMediaStats دریافت آمار فایل‌های کاربر
// GET /api/media/cqrs/stats?userID=123
func (h *MediaCQRSHandler) GetMediaStats(w http.ResponseWriter, r *http.Request) {
	// دریافت user ID
	var userID uint

	// اول از query string
	if userIDStr := r.URL.Query().Get("userID"); userIDStr != "" {
		if uid, err := strconv.ParseUint(userIDStr, 10, 32); err == nil {
			userID = uint(uid)
		}
	}

	// اگر در query نبود، از context
	if userID == 0 {
		if uid, ok := r.Context().Value("user_id").(uint); ok {
			userID = uid
		}
	}

	if userID == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "USER_REQUIRED", "شناسه کاربر الزامی است")
		return
	}

	// استفاده از Query Bus
	stats, err := h.queryBus.GetMediaFileStats(r.Context(), userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "STATS_ERROR", "خطا در دریافت آمار")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, stats)
}

// GetRecentUploads دریافت آخرین آپلودها
// GET /api/media/cqrs/recent?limit=10
func (h *MediaCQRSHandler) GetRecentUploads(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	// استفاده از Query Bus
	files, err := h.queryBus.GetRecentUploads(r.Context(), limit)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "QUERY_ERROR", "خطا در دریافت لیست")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"data":  files,
		"count": len(files),
		"limit": limit,
	})
}

// GetMediaStorageReport گزارش فضای ذخیره‌سازی
// GET /api/media/cqrs/reports/storage?groupBy=category&limit=10
func (h *MediaCQRSHandler) GetMediaStorageReport(w http.ResponseWriter, r *http.Request) {
	groupBy := r.URL.Query().Get("groupBy")
	if groupBy == "" {
		groupBy = "category"
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 {
		limit = 10
	}

	// استفاده از Query Bus
	report, err := h.queryBus.GetMediaStorageReport(r.Context(), groupBy, limit)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "REPORT_ERROR", "خطا در تولید گزارش")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"report":  report,
		"groupBy": groupBy,
	})
}
