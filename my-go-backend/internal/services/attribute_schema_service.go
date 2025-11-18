package services

import (
	"gorm.io/gorm"

	"my-go-backend/internal/models"
)

// AttributeSchemaService aggregates attribute groups + attributes for a category.
// It is used by both admin panel and storefront.
type AttributeSchemaService struct {
	DB *gorm.DB
}

func NewAttributeSchemaService(db *gorm.DB) *AttributeSchemaService {
	return &AttributeSchemaService{DB: db}
}

// GroupWithAttributes is a DTO returned to API callers.
type GroupWithAttributes struct {
	ID         uint                     `json:"id"`
	Name       string                   `json:"name"`
	Attributes []AttributeSchemaElement `json:"attributes"`
}

// AttributeSchemaElement contains attribute meta and its options if any.
type AttributeSchemaElement struct {
	ID            uint                    `json:"id"`
	Name          string                  `json:"name"`
	Type          string                  `json:"type"`
	ControlType   string                  `json:"control_type"`
	HasFilter     bool                    `json:"has_filter"`
	IsKey         bool                    `json:"is_key"`
	ShowOnProduct bool                    `json:"show_on_product"`
	IsRequired    bool                    `json:"is_required"`
	Options       []models.AttributeValue `json:"options,omitempty"`
}

// GetCategorySchema returns groups + attributes sorted by display orders.
func (s *AttributeSchemaService) GetCategorySchema(categoryID uint) ([]GroupWithAttributes, error) {
	var groups []models.AttributeGroup
	if err := s.DB.Where("category_id = ?", categoryID).
		Preload("Attributes", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC")
		}).
		Preload("Attributes.Attribute").
		Order("id ASC").
		Find(&groups).Error; err != nil {
		return nil, err
	}
	if len(groups) == 0 {
		return []GroupWithAttributes{}, nil
	}

	// build DTO
	result := make([]GroupWithAttributes, 0, len(groups))
	for _, g := range groups {
		dto := GroupWithAttributes{ID: g.ID, Name: g.Name}
		for _, aga := range g.Attributes {
			attr := AttributeSchemaElement{
				ID:            aga.Attribute.ID,
				Name:          aga.Attribute.Name,
				Type:          aga.Type,
				ControlType:   aga.ControlType,
				HasFilter:     aga.HasFilter,
				IsKey:         aga.IsKey,
				ShowOnProduct: aga.ShowOnProduct,
				IsRequired:    aga.IsRequired,
			}
			if aga.Attribute.HasOptions {
				var opts []models.AttributeValue
				_ = s.DB.Where("attribute_id = ?", aga.Attribute.ID).Order("sort_order ASC").Find(&opts).Error
				attr.Options = opts
			}
			dto.Attributes = append(dto.Attributes, attr)
		}
		result = append(result, dto)
	}
	return result, nil
}

// FilterDTO is minimal info for sidebar filters.
type FilterDTO struct {
	AttributeID uint                    `json:"attribute_id"`
	Name        string                  `json:"name"`
	Options     []models.AttributeValue `json:"options"`
}

func (s *AttributeSchemaService) GetCategoryFilters(categoryID uint) ([]FilterDTO, error) {
	// join tables to get attributes with has_filter = true
	type row struct {
		AttributeID uint
		Name        string
	}
	var attrs []row
	if err := s.DB.Table("attribute_group_attributes as aga").
		Select("DISTINCT aga.attribute_id, a.name").
		Joins("JOIN attributes a ON a.id = aga.attribute_id").
		Joins("JOIN attribute_groups g ON g.id = aga.attribute_group_id").
		Where("g.category_id = ? AND aga.has_filter = ?", categoryID, true).
		Scan(&attrs).Error; err != nil {
		return nil, err
	}
	filters := make([]FilterDTO, 0, len(attrs))
	for _, r := range attrs {
		var opts []models.AttributeValue
		_ = s.DB.Where("attribute_id = ?", r.AttributeID).Order("sort_order ASC").Find(&opts).Error
		filters = append(filters, FilterDTO{
			AttributeID: r.AttributeID,
			Name:        r.Name,
			Options:     opts,
		})
	}
	return filters, nil
}
