package handlers

import (
	"encoding/json"
	"net/http"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"gorm.io/gorm"
)

type MediaSettingsHandler struct {
	DB *gorm.DB
}

func NewMediaSettingsHandler(db *gorm.DB) *MediaSettingsHandler {
	return &MediaSettingsHandler{DB: db}
}

// GetMediaSettings returns all active media settings
func (h *MediaSettingsHandler) GetMediaSettings(w http.ResponseWriter, r *http.Request) {
	var settings []models.MediaSettings
	if err := h.DB.Where("is_active = ?", true).Order("id").Find(&settings).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطای پایگاه‌داده", err.Error()))
		return
	}
	response := models.MediaSettingsResponse{
		Settings: settings,
		Total:    len(settings),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateMediaSetting creates a new media setting
func (h *MediaSettingsHandler) CreateMediaSetting(w http.ResponseWriter, r *http.Request) {
	var setting models.MediaSettings
	if err := json.NewDecoder(r.Body).Decode(&setting); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "بدنه درخواست نامعتبر است", err.Error()))
		return
	}
	if setting.Name == "" || setting.Width <= 0 || setting.Height <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "ورودی نامعتبر است", nil))
		return
	}
	if setting.Quality < 0 || setting.Quality > 100 {
		setting.Quality = 85
	}
	if err := h.DB.Create(&setting).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطای پایگاه‌داده", err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(setting)
}

// UpdateMediaSetting updates an existing media setting
func (h *MediaSettingsHandler) UpdateMediaSetting(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "شناسه ارسال نشده است", nil))
		return
	}
	var setting models.MediaSettings
	if err := json.NewDecoder(r.Body).Decode(&setting); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "بدنه درخواست نامعتبر است", err.Error()))
		return
	}
	if setting.Name == "" || setting.Width <= 0 || setting.Height <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "ورودی نامعتبر است", nil))
		return
	}
	if setting.Quality < 0 || setting.Quality > 100 {
		setting.Quality = 85
	}
	if err := h.DB.Model(&models.MediaSettings{}).Where("id = ?", id).Updates(setting).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطای پایگاه‌داده", err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(setting)
}

// DeleteMediaSetting deletes a media setting
func (h *MediaSettingsHandler) DeleteMediaSetting(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "شناسه ارسال نشده است", nil))
		return
	}
	if err := h.DB.Delete(&models.MediaSettings{}, id).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطای پایگاه‌داده", err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
