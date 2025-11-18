package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

const smsSettingsFile = "sms_settings.json"

// GetSmsSettings returns the current SMS settings
func GetSmsSettings(c *gin.Context) {
	file, err := os.Open(smsSettingsFile)
	if err != nil {
		c.JSON(http.StatusOK, models.SmsSettings{Provider: models.ProviderFarazSMS})
		return
	}
	defer file.Close()
	var settings models.SmsSettings
	if err := json.NewDecoder(file).Decode(&settings); err != nil {
		c.JSON(http.StatusOK, models.SmsSettings{Provider: models.ProviderFarazSMS})
		return
	}
	c.JSON(http.StatusOK, settings)
}

// UpdateSmsSettings updates the SMS settings
func UpdateSmsSettings(c *gin.Context) {
	var settings models.SmsSettings
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	file, err := os.Create(smsSettingsFile)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("FILE_ERROR", utils.GetErrorMessage("FILE_ERROR"), err.Error()))
		return
	}
	defer file.Close()
	if err := json.NewEncoder(file).Encode(settings); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("FILE_ERROR", utils.GetErrorMessage("FILE_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Settings updated"})
}
