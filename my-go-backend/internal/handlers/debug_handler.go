package handlers

import (
	"net/http"

	"my-go-backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DebugHandler struct {
	db *gorm.DB
}

func NewDebugHandler(db *gorm.DB) *DebugHandler {
	return &DebugHandler{db: db}
}

// GetDevUser بررسی کاربر dev
func (h *DebugHandler) GetDevUser(c *gin.Context) {
	var devUser models.User
	err := h.db.Where("username = ?", "dev").First(&devUser).Error
	
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"exists": false,
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"exists": true,
		"user": gin.H{
			"id": devUser.ID,
			"username": devUser.Username,
			"mobile": devUser.Mobile,
			"email": devUser.Email,
			"role_id": devUser.RoleID,
			"status": devUser.Status,
		},
	})
}
