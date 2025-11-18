package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"
)

// ProductSpecHandler handles saving specs.
type ProductSpecHandler struct {
	DB      *gorm.DB
	Service *services.ProductSpecService
}

func NewProductSpecHandler(db *gorm.DB) *ProductSpecHandler {
	return &ProductSpecHandler{DB: db, Service: services.NewProductSpecService(db)}
}

// SaveSpecs POST /products/:id/specs
func (h *ProductSpecHandler) SaveSpecs(c *gin.Context) {
	productIDStr := c.Param("id")
	pid, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	// For category we can query product row quickly
	var catID uint
	_ = h.DB.Table("products").Select("category_id").Where("id = ?", pid).Scan(&catID).Error

	var input struct {
		Values []services.SpecInput `json:"values"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	if err := h.Service.ValidateAndSave(uint(pid), catID, input.Values); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	c.Status(http.StatusCreated)
}

// GetSpecs GET /products/:id/specs – returns saved attribute values for the product.
func (h *ProductSpecHandler) GetSpecs(c *gin.Context) {
	productIDStr := c.Param("id")
	pid, err := strconv.Atoi(productIDStr)
	if err != nil || pid <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), nil))
		return
	}

	// Join to get attribute names, units, option values, and group names (if any)
	type SpecDTO struct {
		ID               uint    `json:"id"`
		AttributeID      uint    `json:"attribute_id"`
		AttributeValueID *uint   `json:"attribute_value_id,omitempty"`
		Name             string  `json:"name"`
		Value            string  `json:"value"`
		ValueText        *string `json:"value_text,omitempty"`
		Unit             *string `json:"unit,omitempty"`
		Group            *string `json:"group,omitempty"`
		SortOrder        uint    `json:"sort_order"`
	}

	var specs []SpecDTO
	query := `SELECT pav.id,
					 pav.attribute_id                  AS attribute_id,
					 pav.attribute_value_id            AS attribute_value_id,
					 a.name                            AS name,
					 COALESCE(pav.value_text, av.value) AS value,
					 pav.value_text                      AS value_text,
					 a.unit                            AS unit,
					 g.name                            AS "group",
					 pav.sort_order                    AS sort_order
			  FROM product_attribute_values pav
					   JOIN attributes a ON a.id = pav.attribute_id
					   LEFT JOIN attribute_values av ON av.id = pav.attribute_value_id
					   LEFT JOIN attribute_group_attributes aga ON aga.attribute_id = a.id AND aga.deleted_at IS NULL
					   LEFT JOIN attribute_groups g ON g.id = aga.attribute_group_id AND g.deleted_at IS NULL
			  WHERE pav.product_id = ? AND pav.deleted_at IS NULL
			  ORDER BY COALESCE(aga.sort_order, 0), a.sort_order, pav.sort_order, pav.id`

	if err := h.DB.Raw(query, pid).Scan(&specs).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": specs})
}
