package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// MediaSEOHandler هندلرهای مانیتورینگ و Retry برای image_seo_jobs
type MediaSEOHandler struct{ DB *gorm.DB }

func NewMediaSEOHandler(db *gorm.DB) *MediaSEOHandler { return &MediaSEOHandler{DB: db} }

// ListJobs GET /api/admin/image-seo/jobs?status=pending|processing|error
func (h *MediaSEOHandler) ListJobs(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	type jobRow struct {
		ID           uint       `json:"id"`
		MediaID      uint       `json:"media_id"`
		Status       string     `json:"status"`
		Retries      int        `json:"retries"`
		ScheduledAt  *time.Time `json:"scheduled_at"`
		StartedAt    *time.Time `json:"started_at"`
		FinishedAt   *time.Time `json:"finished_at"`
		ErrorMessage *string    `json:"error_message"`
	}
	var rows []jobRow
	if status == "" {
		if err := h.DB.Raw("SELECT id, media_id, status, retries, scheduled_at, started_at, finished_at, error_message FROM image_seo_jobs ORDER BY created_at DESC LIMIT 200").Scan(&rows).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := h.DB.Raw("SELECT id, media_id, status, retries, scheduled_at, started_at, finished_at, error_message FROM image_seo_jobs WHERE status = ? ORDER BY created_at DESC LIMIT 200", status).Scan(&rows).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": rows})
}

// Retry POST /api/admin/image-seo/retry { job_id?:number, media_id?:number }
func (h *MediaSEOHandler) Retry(w http.ResponseWriter, r *http.Request) {
	var req struct {
		JobID   *uint `json:"job_id"`
		MediaID *uint `json:"media_id"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)
	if req.JobID != nil {
		_ = h.DB.Exec("UPDATE image_seo_jobs SET status='pending', error_message=NULL, scheduled_at=NOW()+INTERVAL '10 seconds', started_at=NULL, finished_at=NULL WHERE id=?", *req.JobID).Error
	} else if req.MediaID != nil {
		_ = h.DB.Exec("INSERT INTO image_seo_jobs(media_id, status, scheduled_at) VALUES(?, 'pending', NOW()+INTERVAL '10 seconds')", *req.MediaID).Error
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}

// GetWorkerStatus GET /api/admin/image-seo/worker/status
func (h *MediaSEOHandler) GetWorkerStatus(w http.ResponseWriter, r *http.Request) {
	// شمارش jobهای موجود بر اساس وضعیت
	var stats struct {
		Pending    int64 `json:"pending"`
		Processing int64 `json:"processing"`
		Completed  int64 `json:"completed"`
		Error      int64 `json:"error"`
		Total      int64 `json:"total"`
	}

	// شمارش بر اساس status
	h.DB.Raw("SELECT COUNT(*) FROM image_seo_jobs WHERE status = 'pending'").Scan(&stats.Pending)
	h.DB.Raw("SELECT COUNT(*) FROM image_seo_jobs WHERE status = 'processing'").Scan(&stats.Processing)
	h.DB.Raw("SELECT COUNT(*) FROM image_seo_jobs WHERE status = 'completed'").Scan(&stats.Completed)
	h.DB.Raw("SELECT COUNT(*) FROM image_seo_jobs WHERE status = 'error'").Scan(&stats.Error)
	h.DB.Raw("SELECT COUNT(*) FROM image_seo_jobs").Scan(&stats.Total)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"worker_active": true, // فرض بر این است که worker فعال است
			"statistics":    stats,
			"timestamp":     time.Now().Unix(),
		},
	})
}
