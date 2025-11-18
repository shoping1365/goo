package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
	"my-go-backend/internal/services"
)

// RegisterAttributeSchemaRoutes defines endpoints for category attribute schema and filters.
func RegisterAttributeSchemaRoutes(public *gin.RouterGroup, db *gorm.DB) {
	svc := services.NewAttributeSchemaService(db)
	h := handlers.NewAttributeSchemaHandler(svc)

	cats := public.Group("/categories")
	{
		cats.GET("/:id/attribute-schema", h.GetSchema)
		cats.GET("/:id/filters", h.GetFilters)
	}
}
