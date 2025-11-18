package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
)

// RegisterAttributeRoutes defines all endpoints related to product attributes and their values.
// This keeps SetupRouter clean and makes attribute functionality modular.
func RegisterAttributeRoutes(public *gin.RouterGroup, db *gorm.DB) {
	// Handlers
	attributeHandler := handlers.NewAttributeHandler(db)
	avHandler := handlers.NewAttributeValueHandler(db)

	// ---------------- Attribute routes ----------------
	attributes := public.Group("/attributes")
	{
		attributes.GET("", attributeHandler.ListAttributes)
		attributes.GET("/stats", attributeHandler.GetStats)
		attributes.GET("/:id", attributeHandler.GetAttribute)
		attributes.POST("", attributeHandler.CreateAttribute)
		attributes.PUT("/:id", attributeHandler.UpdateAttribute)
		attributes.DELETE("/:id", attributeHandler.DeleteAttribute)
		attributes.POST("/bulk-delete", attributeHandler.BulkDelete)
	}

	// ---------------- AttributeValue stand-alone routes ----------------
	public.GET("/attribute-values/by-attribute/:attrId", avHandler.ListAttributeValues)
	public.POST("/attribute-values/by-attribute/:attrId", avHandler.CreateAttributeValue)

	public.PUT("/attribute-values/:id", avHandler.UpdateAttributeValue)
	public.DELETE("/attribute-values/:id", avHandler.DeleteAttributeValue)
}
