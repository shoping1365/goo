package handlers

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type GeoHandler struct {
	DB *gorm.DB
}

func NewGeoHandler(db *gorm.DB) *GeoHandler {
	return &GeoHandler{DB: db}
}

type ProvinceDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CityDTO struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Slug      string  `json:"slug"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ListProvinces GET /api/geo/provinces
func (h *GeoHandler) ListProvinces(c *gin.Context) {
	var provinces []ProvinceDTO
	h.DB.Table("provinces").Select("id, name").Order("name").Scan(&provinces)
	c.JSON(http.StatusOK, provinces)
}

// ListCities GET /api/geo/provinces/:id/cities
func (h *GeoHandler) ListCities(c *gin.Context) {
	pid := c.Param("id")
	var cities []CityDTO
	h.DB.Table("cities").
		Select("id, title as name, slug, latitude, longitude").
		Where("province_id = ?", pid).
		Order("title").
		Scan(&cities)
	c.JSON(http.StatusOK, cities)
}
