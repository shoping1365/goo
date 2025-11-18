package handlers

import (
	"encoding/json"
	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gorm.io/gorm"
)

// TODO: Refactor backup handler functions names to avoid duplicates

// BackupHandler handles backup listing and restoration
// All routes should be behind Auth/Admin middleware.
type BackupHandler struct{ DB *gorm.DB }

func NewBackupHandler(db *gorm.DB) *BackupHandler { return &BackupHandler{DB: db} }

// backupRootBH helper (reuse logic)
func backupRootBH() string {
	wd, _ := os.Getwd()
	projectRoot := filepath.Dir(wd)
	return filepath.Join(projectRoot, "public", "uploads", "media", "backup")
}

// ListPeriods returns array of YYYY-MM folders
func (h *BackupHandler) ListPeriods(w http.ResponseWriter, r *http.Request) {
	root := backupRootBH()
	entries, err := os.ReadDir(root)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطا در خواندن پوشه پشتیبان", err.Error()))
		return
	}
	periods := make([]string, 0)
	for _, e := range entries {
		if e.IsDir() && strings.Count(e.Name(), "-") == 1 {
			periods = append(periods, e.Name())
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(periods)))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "periods": periods})
}

// RestorePeriod copies all files from backup/<period> back to uploads directory, overwriting existing.
func (h *BackupHandler) RestorePeriod(w http.ResponseWriter, r *http.Request) {
	// URL pattern: /api/admin/media/backup/restore/:period
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "مسیر نامعتبر است", "مسیر نامعتبر است"))
		return
	}
	period := parts[len(parts)-1]
	if len(period) != 7 || period[4] != '-' {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "فرمت دوره زمانی نامعتبر است", "فرمت دوره زمانی نامعتبر است"))
		return
	}
	root := backupRootBH()
	srcDir := filepath.Join(root, period)
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("FILE_NOT_FOUND", "پشتیبان مورد نظر یافت نشد", nil))
		return
	}

	// uploadsRoot should be the parent directory two levels up (public/uploads)
	uploadsRoot := filepath.Dir(filepath.Dir(root))

	var totalFiles int
	var totalBytes int64
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(srcDir, path)
		relUnix := filepath.ToSlash(rel) // maybe starts with uploads/ or media/uploads/
		if strings.HasPrefix(relUnix, "uploads/") {
			relUnix = strings.TrimPrefix(relUnix, "uploads/")
		}
		if strings.HasPrefix(relUnix, "media/uploads/") {
			relUnix = strings.TrimPrefix(relUnix, "media/uploads/")
		}
		publicURL := "/media/uploads/" + relUnix
		// ensure dir exists
		if err := os.MkdirAll(filepath.Dir(filepath.Join(uploadsRoot, relUnix)), 0o750); err != nil {
			return err
		}
		// copy file to uploads
		if err := services.CopyFile(path, filepath.Join(uploadsRoot, relUnix)); err != nil {
			return err
		}
		// Update DB record: set compressed=false, compressed_size=NULL
		if h.DB != nil {
			h.DB.Model(&models.MediaFile{}).
				Where("url = ?", publicURL).
				Updates(map[string]interface{}{"compressed": false, "compressed_size": nil})
		}
		totalFiles++
		totalBytes += info.Size()
		return nil
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("INTERNAL_ERROR", "خطا در بازگردانی پشتیبان", err.Error()))
		return
	}
	// Simple stats
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":        true,
		"period":         period,
		"restored_files": totalFiles,
		"total_size":     totalBytes,
	})
}
