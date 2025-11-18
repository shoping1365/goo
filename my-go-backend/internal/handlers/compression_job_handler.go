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

// CompressionJobHandler هندلر مدیریت jobهای فشرده‌سازی
// این هندلر فقط با CompressionJobService کار می‌کند و عملیات CRUD و جستجو را فراهم می‌کند

type CompressionJobHandler struct {
	Service *services.CompressionJobService
}

// NewCompressionJobHandler سازنده هندلر job فشرده‌سازی
func NewCompressionJobHandler(service *services.CompressionJobService) *CompressionJobHandler {
	return &CompressionJobHandler{Service: service}
}

// ListCompressionJobs لیست همه jobها (با فیلتر)
func (h *CompressionJobHandler) ListCompressionJobs(w http.ResponseWriter, r *http.Request) {
	mediaIDStr := r.URL.Query().Get("media_id")
	status := r.URL.Query().Get("status")
	batchIDStr := r.URL.Query().Get("batch_id")
	var jobs []models.CompressionJob
	var err error
	switch {
	case mediaIDStr != "":
		mediaID, err := parseUintFromString(mediaIDStr)
		if err == nil {
			jobs, err = h.Service.GetByMedia(r.Context(), mediaID)
		}
	case status != "":
		jobs, err = h.Service.GetByStatus(r.Context(), status)
	case batchIDStr != "":
		batchID, err := parseUintFromString(batchIDStr)
		if err == nil {
			jobs, err = h.Service.ListByBatch(r.Context(), batchID)
		}
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
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": jobs})
}

// GetCompressionJob بازیابی job بر اساس ID
func (h *CompressionJobHandler) GetCompressionJob(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/compression-job/")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	job, err := h.Service.GetByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("NOT_FOUND", utils.GetErrorMessage("NOT_FOUND"), nil))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": job})
}

// CreateCompressionJob ایجاد job جدید
func (h *CompressionJobHandler) CreateCompressionJob(w http.ResponseWriter, r *http.Request) {
	var req models.CompressionJob
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

// UpdateCompressionJob ویرایش job
func (h *CompressionJobHandler) UpdateCompressionJob(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/compression-job/")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	var req models.CompressionJob
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	req.ID = id
	if err := h.Service.Update(context.Background(), &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": req})
}

// DeleteCompressionJob حذف job
func (h *CompressionJobHandler) DeleteCompressionJob(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}
	if err := h.Service.Delete(context.Background(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "Job deleted"})
}
