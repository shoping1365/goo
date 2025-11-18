package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my-go-backend/internal/handlers"
)

// RegisterAttributeGroupRoutes registers endpoints related to attribute groups.
// Route prefix: /attribute-groups
func RegisterAttributeGroupRoutes(public *gin.RouterGroup, db *gorm.DB) {
	handler := handlers.NewAttributeGroupHandler(db)

	groups := public.Group("/attribute-groups")
	{
		groups.GET("", handler.ListAttributeGroups)
		groups.GET("/by-attribute/:attrId", handler.ListGroupsByAttribute)
		groups.GET("/:id", handler.GetAttributeGroup)
		groups.POST("", handler.CreateAttributeGroup)
		groups.PUT("/:id", handler.UpdateAttributeGroup)
		groups.PUT("/:id/attributes", handler.UpsertGroupAttributes)
		groups.DELETE("/:id", handler.DeleteAttributeGroup)
		groups.DELETE("/:id/attributes/:attrId", handler.RemoveAttributeFromGroup)
	}
}
