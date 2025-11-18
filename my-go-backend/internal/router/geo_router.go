package router

import (
	"my-go-backend/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterGeoRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := handlers.NewGeoHandler(db)
	geo := public.Group("/geo")
	{
		geo.GET("/provinces", h.ListProvinces)
		geo.GET("/provinces/:id/cities", h.ListCities)
	}
}
