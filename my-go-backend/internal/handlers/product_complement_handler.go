package handlers

import (
	"net/http"
	"strconv"

	"my-go-backend/internal/models"
	"my-go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProductImageHandler handles product image operations
type ProductImageHandler struct {
	DB *gorm.DB
}

func NewProductImageHandler(db *gorm.DB) *ProductImageHandler {
	return &ProductImageHandler{DB: db}
}

func (h *ProductImageHandler) AddImage(c *gin.Context) {
	productIDStr := c.Param("productId")
	productID, _ := strconv.Atoi(productIDStr)
	var img models.ProductImage
	if err := c.ShouldBindJSON(&img); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	img.ProductID = productID
	if err := h.DB.Create(&img).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusCreated, img)
}

func (h *ProductImageHandler) DeleteImage(c *gin.Context) {
	imageIDStr := c.Param("imageId")
	imageID, _ := strconv.Atoi(imageIDStr)
	if err := h.DB.Delete(&models.ProductImage{}, imageID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Image deleted"})
}

// NOTE: Variant handlers moved to product_variant_handler.go

// ProductSpecificationHandler handles product specification operations
type ProductSpecificationHandler struct {
	DB *gorm.DB
}

func (h *ProductSpecificationHandler) AddSpecification(c *gin.Context) {
	productIDStr := c.Param("productId")
	productID, _ := strconv.Atoi(productIDStr)
	var s models.ProductSpecification
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	s.ProductID = productID
	if err := h.DB.Create(&s).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusCreated, s)
}

func (h *ProductSpecificationHandler) UpdateSpecification(c *gin.Context) {
	specIDStr := c.Param("specId")
	specID, _ := strconv.Atoi(specIDStr)
	var s models.ProductSpecification
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	if err := h.DB.Model(&models.ProductSpecification{}).Where("id = ?", specID).Updates(s).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Specification updated"})
}

func (h *ProductSpecificationHandler) DeleteSpecification(c *gin.Context) {
	specIDStr := c.Param("specId")
	specID, _ := strconv.Atoi(specIDStr)
	if err := h.DB.Delete(&models.ProductSpecification{}, specID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Specification deleted"})
}

// ProductVideoHandler handles product video operations
type ProductVideoHandler struct {
	DB *gorm.DB
}

func (h *ProductVideoHandler) AddVideo(c *gin.Context) {
	productIDStr := c.Param("productId")
	productID, _ := strconv.Atoi(productIDStr)
	var v models.ProductVideo
	if err := c.ShouldBindJSON(&v); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	v.ProductID = productID
	if err := h.DB.Create(&v).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusCreated, v)
}

func (h *ProductVideoHandler) GetVideos(c *gin.Context) {
	productIDStr := c.Param("productId")
	productID, _ := strconv.Atoi(productIDStr)

	var videos []models.ProductVideo
	if err := h.DB.Where("product_id = ?", productID).Find(&videos).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": videos})
}

func (h *ProductVideoHandler) UpdateVideo(c *gin.Context) {
	videoIDStr := c.Param("videoId")
	videoID, _ := strconv.Atoi(videoIDStr)
	var v models.ProductVideo
	if err := c.ShouldBindJSON(&v); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}
	if err := h.DB.Model(&models.ProductVideo{}).Where("id = ?", videoID).Updates(v).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Video updated"})
}

func (h *ProductVideoHandler) DeleteVideo(c *gin.Context) {
	videoIDStr := c.Param("videoId")
	videoID, _ := strconv.Atoi(videoIDStr)
	if err := h.DB.Delete(&models.ProductVideo{}, videoID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Video deleted"})
}

// Legacy note: RelatedProductHandler removed. Related products are derived from categories and not stored.

// ProductComplementHandler مسئول عملیات مربوط به "محصولات مکمل" است
type ProductComplementHandler struct {
	DB *gorm.DB
}

// GetComplements دریافت لیست محصولات مکمل یک محصول
func (h *ProductComplementHandler) GetComplements(c *gin.Context) {
	productIDStr := c.Param("productId")
	productID, _ := strconv.Atoi(productIDStr)

	// SELECT complement products via join on product_complements
	// شامل تصویر اصلی از جدول product_images به عنوان image_url
	type Product struct {
		ID       uint    `json:"id"`
		Name     string  `json:"name"`
		Slug     string  `json:"slug"`
		SKU      string  `json:"sku"`
		Image    string  `json:"main_image" gorm:"column:image"`
		ImageURL string  `json:"image_url"`
		Price    float64 `json:"price"`
		Category string  `json:"category"`
		Priority *int    `json:"priority"`
		Quantity *int    `json:"quantity"`
		Required *bool   `json:"required"`
	}

	var list []Product
	if err := h.DB.Raw(`
        SELECT p.id,
               p.name,
               p.slug,
               p.sku,
               p.image,
               (
                 SELECT pi.image_url
                 FROM product_images pi
                 WHERE pi.product_id = p.id
                 ORDER BY pi.sort_order ASC, pi.id ASC
                 LIMIT 1
               ) AS image_url,
               p.price,
               COALESCE(c.name, '') AS category,
               pc.priority,
               pc.quantity,
               pc.required
        FROM product_complements pc
        JOIN products p ON p.id = pc.complement_product_id
        LEFT JOIN categories c ON c.id = p.category_id
        WHERE pc.product_id = ?
        ORDER BY COALESCE(pc.priority, 9999), p.id
    `, productID).Scan(&list).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

// AddComplement افزودن یک محصول به‌عنوان مکمل
func (h *ProductComplementHandler) AddComplement(c *gin.Context) {
	productIDStr := c.Param("productId")
	productID, _ := strconv.Atoi(productIDStr)
	var payload struct {
		ComplementProductID int   `json:"complement_product_id"`
		Priority            *int  `json:"priority"`
		Required            *bool `json:"required"`
		Quantity            *int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", utils.GetErrorMessage("VALIDATION_ERROR"), err.Error()))
		return
	}

	pri := 0
	if payload.Priority != nil {
		pri = *payload.Priority
	}
	req := false
	if payload.Required != nil {
		req = *payload.Required
	}
	qty := 1
	if payload.Quantity != nil && *payload.Quantity > 0 {
		qty = *payload.Quantity
	}

	if err := h.DB.Exec(`
        INSERT INTO product_complements (product_id, complement_product_id, priority, required, quantity, created_at)
        VALUES (?, ?, ?, ?, ?, NOW())
        ON CONFLICT (product_id, complement_product_id) DO UPDATE
        SET priority = EXCLUDED.priority,
            required = EXCLUDED.required,
            quantity = EXCLUDED.quantity
    `, productID, payload.ComplementProductID, pri, req, qty).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Complement added"})
}

// RemoveComplement حذف محصول مکمل از یک محصول
func (h *ProductComplementHandler) RemoveComplement(c *gin.Context) {
	productIDStr := c.Param("productId")
	productID, _ := strconv.Atoi(productIDStr)
	complementIDStr := c.Param("complementId")
	complementID, _ := strconv.Atoi(complementIDStr)
	if err := h.DB.Exec("DELETE FROM product_complements WHERE product_id = ? AND complement_product_id = ?", productID, complementID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.New("DB_ERROR", utils.GetErrorMessage("DB_ERROR"), err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Complement removed"})
}
