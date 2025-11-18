package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

/*
ShopSettingsHandler
هندلر مدیریت تنظیمات فروشگاه

این هندلر شامل متدهای زیر است:
- GetShopSettings: دریافت تنظیمات فروشگاه
- UpdateShopSettings: بروزرسانی تنظیمات فروشگاه
- GetPublicShopSettings: دریافت تنظیمات عمومی فروشگاه
*/

type ShopSettingsHandler struct {
	settingService *services.SettingService
}

// NewShopSettingsHandler یک نمونه جدید از هندلر تنظیمات فروشگاه ایجاد می‌کند
func NewShopSettingsHandler(settingService *services.SettingService) *ShopSettingsHandler {
	return &ShopSettingsHandler{
		settingService: settingService,
	}
}

// Location ساختار هر location با آدرس و شماره‌های تلفن
type Location struct {
	ID      interface{} `json:"id"`
	Title   string      `json:"title"`
	Address string      `json:"address"`
	Phones  []string    `json:"phones"`
}

// ShopSettingsRequest ساختار درخواست بروزرسانی تنظیمات فروشگاه
type ShopSettingsRequest struct {
	// اطلاعات پایه
	ShopNameFa string `json:"shopNameFa"`
	ShopNameEn string `json:"shopNameEn"`
	Logo       string `json:"logo"`
	LogoRetina string `json:"logoRetina"`
	Favicon    string `json:"favicon"`

	// تنظیمات منطقه‌ای
	DefaultLanguage string `json:"defaultLanguage"`
	Timezone        string `json:"timezone"`
	DefaultCurrency string `json:"defaultCurrency"`

	// وضعیت فروشگاه
	MaintenanceMode    bool   `json:"maintenanceMode"`
	MaintenanceMessage string `json:"maintenanceMessage"`

	// اطلاعات تماس (ساختار جدید)
	Locations []Location `json:"locations"`

	// اطلاعات تماس (ساختار قدیمی - Backward Compatibility)
	Phones      []string `json:"phones"`
	Email       string   `json:"email"`
	AdminPhones []string `json:"adminPhones"`
	Address     string   `json:"address"`
	Latitude    string   `json:"latitude"`
	Longitude   string   `json:"longitude"`

	// اطلاعات اضافی
	WorkingHours     []string `json:"workingHours"` // تغییر از string به []string
	ShortDescription string   `json:"shortDescription"`
}

// ShopSettingsResponse ساختار پاسخ تنظیمات فروشگاه
type ShopSettingsResponse struct {
	// اطلاعات پایه
	ShopNameFa string `json:"shopNameFa"`
	ShopNameEn string `json:"shopNameEn"`
	Logo       string `json:"logo"`
	LogoRetina string `json:"logoRetina"`
	Favicon    string `json:"favicon"`

	// تنظیمات منطقه‌ای
	DefaultLanguage string `json:"defaultLanguage"`
	Timezone        string `json:"timezone"`
	DefaultCurrency string `json:"defaultCurrency"`

	// وضعیت فروشگاه
	MaintenanceMode    bool   `json:"maintenanceMode"`
	MaintenanceMessage string `json:"maintenanceMessage"`

	// اطلاعات تماس (ساختار جدید)
	Locations []Location `json:"locations"`

	// اطلاعات تماس (ساختار قدیمی - Backward Compatibility)
	Phones      []string `json:"phones"`
	Email       string   `json:"email"`
	AdminPhones []string `json:"adminPhones"`
	Address     string   `json:"address"`
	Latitude    string   `json:"latitude"`
	Longitude   string   `json:"longitude"`

	// اطلاعات اضافی
	WorkingHours     []string `json:"workingHours"` // تغییر از string به []string
	ShortDescription string   `json:"shortDescription"`
}

// GetShopSettings دریافت تمام تنظیمات فروشگاه
func (h *ShopSettingsHandler) GetShopSettings(c *gin.Context) {
	ctx := c.Request.Context()

	// دریافت تنظیمات فروشگاه از سرویس
	settings, err := h.settingService.GetSettingsByCategory(ctx, "shop")
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("SETTINGS_ERROR", "خطا در دریافت تنظیمات فروشگاه", err.Error()))
		return
	}

	// تبدیل تنظیمات به ساختار پاسخ
	response := h.convertSettingsToResponse(settings)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   response,
	})
}

// UpdateShopSettings بروزرسانی تنظیمات فروشگاه
func (h *ShopSettingsHandler) UpdateShopSettings(c *gin.Context) {
	ctx := c.Request.Context()

	var request ShopSettingsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("خطا در اعتبارسنجی JSON: %v", err)

		// اصلاح خطای GetRawData
		// rawData, _ := c.GetRawData()
		// log.Printf("بدنه درخواست: %s", string(rawData))

		c.JSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "خطا در اعتبارسنجی داده‌ها", err.Error()))
		return
	}

	// log.Printf("درخواست دریافتی: %+v", request)

	// تبدیل درخواست به تنظیمات
	settings := h.convertRequestToSettings(request)

	// استفاده از UpdateSettings که خودش UPSERT را مدیریت می‌کند
	if err := h.settingService.UpdateSettings(ctx, settings); err != nil {
		log.Printf("خطا در بروزرسانی تنظیمات: %v", err)
		c.JSON(http.StatusInternalServerError, utils.New("UPDATE_ERROR", "خطا در بروزرسانی تنظیمات", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "تنظیمات فروشگاه با موفقیت بروزرسانی شد",
	})
}

// GetPublicShopSettings دریافت تنظیمات عمومی فروشگاه
func (h *ShopSettingsHandler) GetPublicShopSettings(c *gin.Context) {
	ctx := c.Request.Context()

	// دریافت تنظیمات عمومی فروشگاه
	settings, err := h.settingService.GetSettingsByCategory(ctx, "shop")
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.New("SETTINGS_ERROR", "خطا در دریافت تنظیمات فروشگاه", err.Error()))
		return
	}

	// فیلتر کردن تنظیمات عمومی
	var publicSettings []models.Setting
	for _, setting := range settings {
		if setting.IsPublic {
			publicSettings = append(publicSettings, setting)
		}
	}

	// تبدیل تنظیمات به ساختار پاسخ
	response := h.convertSettingsToResponse(publicSettings)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   response,
	})
}

// convertSettingsToResponse تبدیل تنظیمات به ساختار پاسخ
func (h *ShopSettingsHandler) convertSettingsToResponse(settings []models.Setting) ShopSettingsResponse {
	response := ShopSettingsResponse{
		// مقادیر پیش‌فرض
		ShopNameFa:         "فروشگاه من",
		ShopNameEn:         "My Shop",
		DefaultLanguage:    "fa",
		Timezone:           "Asia/Tehran",
		DefaultCurrency:    "IRR",
		MaintenanceMode:    false,
		MaintenanceMessage: "فروشگاه در حال تعمیر و نگهداری است. لطفاً بعداً مراجعه کنید.",
		Phones:             []string{},
		AdminPhones:        []string{},
		WorkingHours:       []string{},
		Locations:          []Location{},
	}

	// تبدیل تنظیمات به ساختار پاسخ
	for _, setting := range settings {
		switch setting.Key {
		case "shop_name_fa":
			response.ShopNameFa = setting.Value
		case "shop_name_en":
			response.ShopNameEn = setting.Value
		case "shop_logo":
			response.Logo = setting.Value
		case "shop_logo_retina":
			response.LogoRetina = setting.Value
		case "shop_favicon":
			response.Favicon = setting.Value
		case "default_language":
			response.DefaultLanguage = setting.Value
		case "timezone":
			response.Timezone = setting.Value
		case "default_currency":
			response.DefaultCurrency = setting.Value
		case "maintenance_mode":
			response.MaintenanceMode = setting.Value == "true"
		case "maintenance_message":
			response.MaintenanceMessage = setting.Value
		case "shop_locations":
			if setting.Value != "" {
				json.Unmarshal([]byte(setting.Value), &response.Locations)
			}
		case "shop_phones":
			if setting.Value != "" {
				json.Unmarshal([]byte(setting.Value), &response.Phones)
			}
		case "shop_email":
			response.Email = setting.Value
		case "admin_phones":
			if setting.Value != "" {
				json.Unmarshal([]byte(setting.Value), &response.AdminPhones)
			}
		case "shop_address":
			response.Address = setting.Value
		case "shop_latitude":
			response.Latitude = setting.Value
		case "shop_longitude":
			response.Longitude = setting.Value
		case "working_hours":
			if setting.Value != "" {
				json.Unmarshal([]byte(setting.Value), &response.WorkingHours) // تغییر
			}
		case "shop_description":
			response.ShortDescription = setting.Value
		}
	}

	return response
}

// convertRequestToSettings تبدیل درخواست به تنظیمات
func (h *ShopSettingsHandler) convertRequestToSettings(request ShopSettingsRequest) []models.Setting {
	var settings []models.Setting

	// تبدیل آرایه‌ها به JSON
	locationsJSON, _ := json.Marshal(request.Locations)
	phonesJSON, _ := json.Marshal(request.Phones)
	adminPhonesJSON, _ := json.Marshal(request.AdminPhones)
	workingHoursJSON, _ := json.Marshal(request.WorkingHours)

	// ایجاد تنظیمات
	settingsMap := map[string]string{
		"shop_name_fa":     request.ShopNameFa,
		"shop_name_en":     request.ShopNameEn,
		"shop_logo":        request.Logo,
		"shop_logo_retina": request.LogoRetina,
		"shop_favicon":     request.Favicon,
		"default_language": request.DefaultLanguage,
		"timezone":         request.Timezone,
		"default_currency": request.DefaultCurrency,
		"maintenance_mode": func() string {
			if request.MaintenanceMode {
				return "true"
			} else {
				return "false"
			}
		}(),
		"maintenance_message": request.MaintenanceMessage,
		"shop_locations":      string(locationsJSON),
		"shop_phones":         string(phonesJSON),
		"shop_email":          request.Email,
		"admin_phones":        string(adminPhonesJSON),
		"shop_address":        request.Address,
		"shop_latitude":       request.Latitude,
		"shop_longitude":      request.Longitude,
		"working_hours":       string(workingHoursJSON),
		"shop_description":    request.ShortDescription,
	}

	// ایجاد آبجکت‌های Setting
	for key, value := range settingsMap {
		setting := models.Setting{
			Key:       key,
			Value:     value,
			Category:  "shop",
			Type:      "string",
			IsPublic:  true,
			CreatedBy: "system",
			UpdatedBy: "admin",
		}

		// تنظیم نوع برای فیلدهای خاص
		if key == "maintenance_mode" {
			setting.Type = "boolean"
		} else if key == "shop_locations" || key == "shop_phones" || key == "admin_phones" || key == "working_hours" {
			setting.Type = "json"
		}

		// تنظیم is_public برای فیلدهای خصوصی
		if key == "admin_phones" {
			setting.IsPublic = false
		}

		settings = append(settings, setting)
	}

	return settings
}
