package handlers

import (
	"context"
	"encoding/json"
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
	"net/http"
	"strings"
)

// CompressionSettingHandler هندلر مدیریت تنظیمات فشرده‌سازی
// این هندلر فقط با CompressionSettingService کار می‌کند و عملیات CRUD و جستجو را فراهم می‌کند

type CompressionSettingHandler struct {
	Service *services.CompressionSettingService
}

// NewCompressionSettingHandler سازنده هندلر تنظیمات فشرده‌سازی
func NewCompressionSettingHandler(service *services.CompressionSettingService) *CompressionSettingHandler {
	return &CompressionSettingHandler{Service: service}
}

// ListCompressionSettings لیست همه تنظیمات (با فیلتر)
func (h *CompressionSettingHandler) ListCompressionSettings(w http.ResponseWriter, r *http.Request) {
	scope := r.URL.Query().Get("scope")
	userIDStr := r.URL.Query().Get("user_id")
	mediaIDStr := r.URL.Query().Get("media_id")
	jobIDStr := r.URL.Query().Get("job_id")
	var settings []models.CompressionSetting
	var err error
	switch {
	case scope != "":
		settings, err = h.Service.GetByScope(r.Context(), scope)
	case userIDStr != "":
		userID, err := parseUintFromString(userIDStr)
		if err == nil {
			settings, err = h.Service.GetByUser(r.Context(), userID)
		}
	case mediaIDStr != "":
		mediaID, err := parseUintFromString(mediaIDStr)
		if err == nil {
			settings, err = h.Service.GetByMedia(r.Context(), mediaID)
		}
	case jobIDStr != "":
		jobID, err := parseUintFromString(jobIDStr)
		if err == nil {
			settings, err = h.Service.GetByJob(r.Context(), jobID)
		}
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": settings})
}

// GetCompressionSetting بازیابی تنظیم بر اساس ID
func (h *CompressionSettingHandler) GetCompressionSetting(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/compression-setting/")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	setting, err := h.Service.GetByID(r.Context(), id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("SETTING_NOT_FOUND", utils.GetErrorMessage("SETTING_NOT_FOUND"), nil))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": setting})
}

// CreateCompressionSetting ایجاد تنظیم جدید
func (h *CompressionSettingHandler) CreateCompressionSetting(w http.ResponseWriter, r *http.Request) {
	var req models.CompressionSetting
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	if err := h.Service.Create(context.Background(), &req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": req})
}

// UpdateCompressionSetting ویرایش تنظیم
func (h *CompressionSettingHandler) UpdateCompressionSetting(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/compression-setting/")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	var req models.CompressionSetting
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	req.ID = id
	if err := h.Service.Update(context.Background(), &req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": req})
}

// DeleteCompressionSetting حذف تنظیم
func (h *CompressionSettingHandler) DeleteCompressionSetting(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	if err := h.Service.Delete(context.Background(), id); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "Setting deleted"})
}
