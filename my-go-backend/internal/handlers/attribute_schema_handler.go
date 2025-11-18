package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// AttributeSchemaHandler exposes category attribute schema & filters.
type AttributeSchemaHandler struct {
	Service *services.AttributeSchemaService
}

func NewAttributeSchemaHandler(svc *services.AttributeSchemaService) *AttributeSchemaHandler {
	return &AttributeSchemaHandler{Service: svc}
}

// GetSchema returns groups + attributes for a category.
func (h *AttributeSchemaHandler) GetSchema(c *gin.Context) {
	idStr := c.Param("id")
	catID, err := strconv.Atoi(idStr)
	if err != nil || catID <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه دسته‌بندی نامعتبر است", nil))
		return
	}
	schema, err := h.Service.GetCategorySchema(uint(catID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت طرح ویژگی", err.Error()))
		return
	}
	c.JSON(http.StatusOK, schema)
}

// GetFilters returns sidebar filter data for category.
func (h *AttributeSchemaHandler) GetFilters(c *gin.Context) {
	idStr := c.Param("id")
	catID, err := strconv.Atoi(idStr)
	if err != nil || catID <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "شناسه دسته‌بندی نامعتبر است", nil))
		return
	}
	filters, err := h.Service.GetCategoryFilters(uint(catID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", "خطا در دریافت فیلترها", err.Error()))
		return
	}
	c.JSON(http.StatusOK, filters)
}
