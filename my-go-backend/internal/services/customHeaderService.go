package services

import (
	"encoding/json"
	"errors"
	"my-go-backend/internal/models"
	"my-go-backend/internal/repository"
	"strconv"
	"strings"
)

type CustomHeaderService struct {
	headerRepo *repository.CustomHeaderRepository
}

func NewCustomHeaderService(headerRepo *repository.CustomHeaderRepository) *CustomHeaderService {
	return &CustomHeaderService{
		headerRepo: headerRepo,
	}
}

// GetAllHeaders دریافت تمام هدرها
func (s *CustomHeaderService) GetAllHeaders() ([]models.Header, error) {
	return s.headerRepo.GetAllHeaders()
}

// GetHeaderByID دریافت هدر بر اساس ID
func (s *CustomHeaderService) GetHeaderByID(id uint) (*models.Header, error) {
	return s.headerRepo.GetHeaderByID(id)
}

// CreateHeader ایجاد هدر جدید
func (s *CustomHeaderService) CreateHeader(headerData map[string]interface{}) (*models.Header, error) {
	// اعتبارسنجی نام هدر
	name, ok := headerData["name"].(string)
	if !ok || strings.TrimSpace(name) == "" {
		return nil, errors.New("نام هدر الزامی است")
	}

	// بررسی تکراری نبودن نام
	exists, err := s.headerRepo.HeaderExists(name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("هدری با این نام قبلاً وجود دارد")
	}

	// تبدیل لایه‌ها
	layersData, ok := headerData["layers"].([]interface{})
	if !ok || len(layersData) == 0 {
		return nil, errors.New("حداقل یک لایه باید ایجاد شود")
	}

	var layers []models.HeaderLayer
	for _, layerData := range layersData {
		layer, err := s.convertLayerData(layerData)
		if err != nil {
			return nil, err
		}
		layers = append(layers, layer)
	}

	// ایجاد هدر جدید
	header := &models.Header{
		Name:          strings.TrimSpace(name),
		Description:   s.getStringValue(headerData, "description"),
		PageSelection: s.getStringValue(headerData, "pageSelection"),
		SpecificPages: s.getStringValue(headerData, "specificPages"),
		ExcludedPages: s.getStringValue(headerData, "excludedPages"),
		IsActive:      true,
		Layers:        layers,
	}

	err = s.headerRepo.CreateHeader(header)
	if err != nil {
		return nil, err
	}

	return header, nil
}

// UpdateHeader به‌روزرسانی هدر
func (s *CustomHeaderService) UpdateHeader(id uint, headerData map[string]interface{}) (*models.Header, error) {
	// دریافت هدر موجود
	existingHeader, err := s.headerRepo.GetHeaderByID(id)
	if err != nil {
		return nil, errors.New("هدر مورد نظر یافت نشد")
	}

	// اعتبارسنجی نام هدر
	name, ok := headerData["name"].(string)
	if !ok || strings.TrimSpace(name) == "" {
		return nil, errors.New("نام هدر الزامی است")
	}

	// بررسی تکراری نبودن نام (به جز خود هدر)
	exists, err := s.headerRepo.HeaderExists(name, id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("هدری با این نام قبلاً وجود دارد")
	}

	// تبدیل لایه‌ها
	layersData, ok := headerData["layers"].([]interface{})
	if !ok || len(layersData) == 0 {
		return nil, errors.New("حداقل یک لایه باید ایجاد شود")
	}

	var layers []models.HeaderLayer
	for _, layerData := range layersData {
		layer, err := s.convertLayerData(layerData)
		if err != nil {
			return nil, err
		}
		layers = append(layers, layer)
	}

	// به‌روزرسانی هدر
	existingHeader.Name = strings.TrimSpace(name)
	existingHeader.Description = s.getStringValue(headerData, "description")
	existingHeader.PageSelection = s.getStringValue(headerData, "pageSelection")
	existingHeader.SpecificPages = s.getStringValue(headerData, "specificPages")
	existingHeader.ExcludedPages = s.getStringValue(headerData, "excludedPages")
	existingHeader.Layers = layers

	err = s.headerRepo.UpdateHeader(existingHeader)
	if err != nil {
		return nil, err
	}

	return existingHeader, nil
}

// DeleteHeader حذف هدر
func (s *CustomHeaderService) DeleteHeader(id uint) error {
	// بررسی وجود هدر
	_, err := s.headerRepo.GetHeaderByID(id)
	if err != nil {
		return errors.New("هدر مورد نظر یافت نشد")
	}

	return s.headerRepo.DeleteHeader(id)
}

// GetActiveHeaders دریافت هدرهای فعال
func (s *CustomHeaderService) GetActiveHeaders() ([]models.Header, error) {
	return s.headerRepo.GetActiveHeaders()
}

// convertLayerData تبدیل داده‌های لایه
func (s *CustomHeaderService) convertLayerData(layerData interface{}) (models.HeaderLayer, error) {
	layerMap, ok := layerData.(map[string]interface{})
	if !ok {
		return models.HeaderLayer{}, errors.New("داده‌های لایه نامعتبر است")
	}

	itemsJSON := s.resolveItemsJSON(layerMap["items"])

	layer := models.HeaderLayer{
		Name:             s.getStringFromKeys(layerMap, "", "name"),
		Width:            s.getIntFromKeys(layerMap, 100, "width"),
		Height:           s.getIntFromKeys(layerMap, 50, "height"),
		RowCount:         s.getIntFromKeys(layerMap, 1, "rowCount", "row_count"),
		Color:            s.getStringFromKeys(layerMap, "", "color"),
		Opacity:          s.getFloatFromKeys(layerMap, 1.0, "opacity", "opacity_percentage"),
		ShowSeparator:    s.getBoolFromKeys(layerMap, false, "showSeparator", "show_separator", "separatorEnabled"),
		SeparatorType:    s.getStringFromKeys(layerMap, "", "separatorType", "separator_type"),
		SeparatorColor:   s.getStringFromKeys(layerMap, "", "separatorColor", "separator_color"),
		SeparatorOpacity: s.getFloatFromKeys(layerMap, 1.0, "separatorOpacity", "separator_opacity"),
		SeparatorWidth:   s.getIntFromKeys(layerMap, 1, "separatorWidth", "separator_width"),
		EnableBorder:     s.getBoolFromKeys(layerMap, false, "enable_border", "enableBorder", "borderEnabled"),
		BorderPosition:   s.getStringFromKeys(layerMap, "all", "border_position", "borderPosition"),
		BorderColor:      s.getStringFromKeys(layerMap, "", "border_color", "borderColor"),
		BorderWidth:      s.getIntFromKeys(layerMap, 1, "border_width", "borderWidth"),
		BorderStyle:      s.getStringFromKeys(layerMap, "solid", "border_style", "borderStyle"),
		EnableShadow:     s.getBoolFromKeys(layerMap, false, "enable_shadow", "enableShadow", "shadowEnabled", "elevationEnabled"),
		ShadowIntensity:  s.getStringFromKeys(layerMap, "md", "shadow_intensity", "shadowIntensity"),
		ShadowDirection:  s.getStringFromKeys(layerMap, "top", "shadow_direction", "shadowDirection"),
		Direction:        s.getStringFromKeys(layerMap, "rtl", "direction"),
		MobileResponsive: s.getBoolFromKeys(layerMap, true, "mobile_responsive", "mobileResponsive"),
		TabletResponsive: s.getBoolFromKeys(layerMap, true, "tablet_responsive", "tabletResponsive"),
		PaddingLeft:      s.getIntFromKeys(layerMap, 0, "padding_left", "paddingLeft"),
		PaddingRight:     s.getIntFromKeys(layerMap, 0, "padding_right", "paddingRight"),
		PaddingTop:       s.getIntFromKeys(layerMap, 0, "padding_top", "paddingTop"),
		PaddingBottom:    s.getIntFromKeys(layerMap, 0, "padding_bottom", "paddingBottom"),
		Items:            itemsJSON,
		StyleSettings:    s.resolveStyleSettings(layerMap["style_settings"], layerMap["styleSettings"]),
	}

	return layer, nil
}

// Helper functions
func (s *CustomHeaderService) getStringFromKeys(data map[string]interface{}, defaultValue string, keys ...string) string {
	for _, key := range keys {
		if value, ok := data[key]; ok {
			switch v := value.(type) {
			case string:
				if strings.TrimSpace(v) != "" {
					return v
				}
			}
		}
	}
	return defaultValue
}

func (s *CustomHeaderService) getIntFromKeys(data map[string]interface{}, defaultValue int, keys ...string) int {
	for _, key := range keys {
		if value, ok := data[key]; ok {
			switch v := value.(type) {
			case float64:
				return int(v)
			case int:
				return v
			case int64:
				return int(v)
			case string:
				if parsed, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
					return parsed
				}
			}
		}
	}
	return defaultValue
}

func (s *CustomHeaderService) getFloatFromKeys(data map[string]interface{}, defaultValue float64, keys ...string) float64 {
	for _, key := range keys {
		if value, ok := data[key]; ok {
			switch v := value.(type) {
			case float64:
				return v
			case int:
				return float64(v)
			case string:
				if parsed, err := strconv.ParseFloat(strings.TrimSpace(v), 64); err == nil {
					return parsed
				}
			}
		}
	}
	return defaultValue
}

func (s *CustomHeaderService) getBoolFromKeys(data map[string]interface{}, defaultValue bool, keys ...string) bool {
	for _, key := range keys {
		if value, ok := data[key]; ok {
			switch v := value.(type) {
			case bool:
				return v
			case float64:
				return v != 0
			case int:
				return v != 0
			case string:
				n := strings.TrimSpace(strings.ToLower(v))
				if n == "true" || n == "1" || n == "yes" || n == "on" {
					return true
				}
				if n == "false" || n == "0" || n == "no" || n == "off" {
					return false
				}
			}
		}
	}
	return defaultValue
}

func (s *CustomHeaderService) resolveItemsJSON(items interface{}) string {
	if items == nil {
		return ""
	}

	switch v := items.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case []interface{}:
		if len(v) == 0 {
			return ""
		}
		if data, err := json.Marshal(v); err == nil {
			return string(data)
		}
	case map[string]interface{}:
		if data, err := json.Marshal(v); err == nil {
			return string(data)
		}
	}

	return ""
}

func (s *CustomHeaderService) resolveStyleSettings(primary interface{}, fallback interface{}) string {
	sources := []interface{}{primary, fallback}
	for _, src := range sources {
		switch val := src.(type) {
		case string:
			if strings.TrimSpace(val) != "" {
				return val
			}
		case []byte:
			return string(val)
		case map[string]interface{}:
			if data, err := json.Marshal(val); err == nil {
				return string(data)
			}
		}
	}
	return "{}"
}

// Backwards compatible helper wrappers
func (s *CustomHeaderService) getStringValue(data map[string]interface{}, key string) string {
	return s.getStringFromKeys(data, "", key)
}

func (s *CustomHeaderService) getIntValue(data map[string]interface{}, key string, defaultValue int) int {
	return s.getIntFromKeys(data, defaultValue, key)
}

func (s *CustomHeaderService) getFloatValue(data map[string]interface{}, key string, defaultValue float64) float64 {
	return s.getFloatFromKeys(data, defaultValue, key)
}

func (s *CustomHeaderService) getBoolValue(data map[string]interface{}, key string, defaultValue bool) bool {
	return s.getBoolFromKeys(data, defaultValue, key)
}
