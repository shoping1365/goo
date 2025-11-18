package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
)

// CompressionStatHandler هندلر مدیریت آمار و گزارش‌های فشرده‌سازی
// این هندلر فقط با CompressionStatService کار می‌کند و عملیات CRUD و جستجو را فراهم می‌کند

type CompressionStatHandler struct {
	Service *services.CompressionStatService
}

// NewCompressionStatHandler سازنده هندلر آمار فشرده‌سازی
func NewCompressionStatHandler(service *services.CompressionStatService) *CompressionStatHandler {
	return &CompressionStatHandler{Service: service}
}

// ListCompressionStats لیست همه آمار (با فیلتر)
func (h *CompressionStatHandler) ListCompressionStats(w http.ResponseWriter, r *http.Request) {
	statType := r.URL.Query().Get("stat_type")
	mediaIDStr := r.URL.Query().Get("media_id")
	userIDStr := r.URL.Query().Get("user_id")
	format := r.URL.Query().Get("format")
	period := r.URL.Query().Get("period")
	var stats []models.CompressionStat
	var err error
	switch {
	case statType != "":
		stats, err = h.Service.GetByType(r.Context(), statType)
	case mediaIDStr != "":
		mediaID, _ := strconv.ParseUint(mediaIDStr, 10, 64)
		stats, err = h.Service.GetByMedia(r.Context(), uint(mediaID))
	case userIDStr != "":
		userID, _ := strconv.ParseUint(userIDStr, 10, 64)
		stats, err = h.Service.GetByUser(r.Context(), uint(userID))
	case format != "":
		stats, err = h.Service.GetByFormat(r.Context(), format)
	case period != "":
		stats, err = h.Service.GetByPeriod(r.Context(), period)
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": stats})
}

// GetCompressionStat بازیابی آمار بر اساس ID
func (h *CompressionStatHandler) GetCompressionStat(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/compression-stat/")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	stat, err := h.Service.GetByID(r.Context(), uint(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": stat})
}

// CreateCompressionStat ایجاد آمار جدید
func (h *CompressionStatHandler) CreateCompressionStat(w http.ResponseWriter, r *http.Request) {
	var req models.CompressionStat
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	if err := h.Service.Create(context.Background(), &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": req})
}

// UpdateCompressionStat ویرایش آمار
func (h *CompressionStatHandler) UpdateCompressionStat(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/compression-stat/")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	var req models.CompressionStat
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	req.ID = uint(id)
	if err := h.Service.Update(context.Background(), &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": req})
}

// DeleteCompressionStat حذف آمار
func (h *CompressionStatHandler) DeleteCompressionStat(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	if err := h.Service.Delete(context.Background(), uint(id)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "Stat deleted"})
}
