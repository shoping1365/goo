package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
)

// MediaVersionHandler هندلر نسخه‌های رسانه
// این هندلر فقط با MediaVersionService کار می‌کند و عملیات CRUD و جستجو را فراهم می‌کند

type MediaVersionHandler struct {
	Service *services.MediaVersionService
}

// NewMediaVersionHandler سازنده هندلر نسخه رسانه
func NewMediaVersionHandler(service *services.MediaVersionService) *MediaVersionHandler {
	return &MediaVersionHandler{Service: service}
}

// ListMediaVersions لیست همه نسخه‌ها
func (h *MediaVersionHandler) ListMediaVersions(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := r.URL.Query().Get("media_id")
	if mediaIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "media_id الزامی است", nil))
		return
	}

	mediaID, err := parseUintFromString(mediaIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "شناسه رسانه نامعتبر است", nil))
		return
	}

	versions, err := h.Service.GetByMedia(r.Context(), mediaID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطا در دریافت نسخه‌ها", err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": versions})
}

// GetMediaVersion بازیابی نسخه بر اساس ID
func (h *MediaVersionHandler) GetMediaVersion(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/media-version/")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "شناسه نسخه نامعتبر است", nil))
		return
	}
	version, err := h.Service.GetByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("NOT_FOUND", "نسخه یافت نشد", nil))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": version})
}

// CreateMediaVersion ایجاد نسخه جدید
func (h *MediaVersionHandler) CreateMediaVersion(w http.ResponseWriter, r *http.Request) {
	var req models.MediaVersion
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "بدنه درخواست نامعتبر است", err.Error()))
		return
	}
	if err := h.Service.Create(context.Background(), &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطا در ایجاد نسخه", err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": req})
}

// UpdateMediaVersion ویرایش نسخه
func (h *MediaVersionHandler) UpdateMediaVersion(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/media-version/")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "شناسه نسخه نامعتبر است", nil))
		return
	}
	var req models.MediaVersion
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "بدنه درخواست نامعتبر است", err.Error()))
		return
	}
	req.ID = id
	if err := h.Service.Update(context.Background(), &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطا در بروزرسانی نسخه", err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": req})
}

// DeleteMediaVersion حذف نسخه
func (h *MediaVersionHandler) DeleteMediaVersion(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "شناسه نسخه نامعتبر است", nil))
		return
	}
	if err := h.Service.Delete(context.Background(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطا در حذف نسخه", err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "Version deleted"})
}
