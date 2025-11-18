package services

import (
	"errors"
	"my-go-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

// SchemaService سرویس مدیریت اسکیمای پیش‌فرض
type SchemaService struct {
	db *gorm.DB
}

// NewSchemaService ایجاد نمونه جدید از SchemaService
func NewSchemaService(db *gorm.DB) *SchemaService {
	return &SchemaService{db: db}
}

// GetAllTemplates دریافت تمام تمپلیت‌های فعال
func (s *SchemaService) GetAllTemplates() ([]models.Schema, error) {
	var schemas []models.Schema
	err := s.db.Where("is_template = ? AND is_active = ?", true, true).
		Preload("Creator").
		Order("created_at DESC").
		Find(&schemas).Error
	return schemas, err
}

// GetTemplatesByType دریافت تمپلیت‌ها بر اساس نوع
func (s *SchemaService) GetTemplatesByType(schemaType string) ([]models.Schema, error) {
	var schemas []models.Schema
	err := s.db.Where("type = ? AND is_template = ? AND is_active = ?", schemaType, true, true).
		Preload("Creator").
		Order("created_at DESC").
		Find(&schemas).Error
	return schemas, err
}

// GetTemplateByID دریافت تمپلیت بر اساس ID
func (s *SchemaService) GetTemplateByID(id uint) (*models.Schema, error) {
	var schema models.Schema
	err := s.db.Where("id = ? AND is_template = ?", id, true).
		Preload("Creator").
		Preload("Updater").
		First(&schema).Error
	if err != nil {
		return nil, err
	}
	return &schema, nil
}

// CreateTemplate ایجاد تمپلیت جدید
func (s *SchemaService) CreateTemplate(schema *models.Schema, userID uint) error {
	schema.IsTemplate = true
	schema.IsActive = true
	schema.CreatedBy = userID
	schema.UpdatedBy = userID
	schema.CreatedAt = time.Now()
	schema.UpdatedAt = time.Now()

	return s.db.Create(schema).Error
}

// UpdateTemplate به‌روزرسانی تمپلیت
func (s *SchemaService) UpdateTemplate(id uint, updates map[string]interface{}, userID uint) error {
	updates["updated_by"] = userID
	updates["updated_at"] = time.Now()

	return s.db.Model(&models.Schema{}).
		Where("id = ? AND is_template = ?", id, true).
		Updates(updates).Error
}

// DeleteTemplate حذف تمپلیت (soft delete)
func (s *SchemaService) DeleteTemplate(id uint) error {
	return s.db.Where("id = ? AND is_template = ?", id, true).
		Delete(&models.Schema{}).Error
}

// ToggleTemplateStatus تغییر وضعیت فعال/غیرفعال تمپلیت
func (s *SchemaService) ToggleTemplateStatus(id uint, userID uint) error {
	var schema models.Schema
	if err := s.db.Where("id = ? AND is_template = ?", id, true).First(&schema).Error; err != nil {
		return err
	}

	schema.IsActive = !schema.IsActive
	schema.UpdatedBy = userID
	schema.UpdatedAt = time.Now()

	return s.db.Save(&schema).Error
}

// GenerateSchemaFromTemplate تولید اسکیما از تمپلیت
func (s *SchemaService) GenerateSchemaFromTemplate(templateID uint, data map[string]interface{}) (*models.Schema, error) {
	// دریافت تمپلیت
	template, err := s.GetTemplateByID(templateID)
	if err != nil {
		return nil, err
	}

	// ایجاد اسکیما جدید از تمپلیت
	schema := &models.Schema{
		Name:            template.Name,
		Type:            template.Type,
		Description:     template.Description,
		IsTemplate:      false, // این یک اسکیما واقعی است، نه تمپلیت
		IsActive:        true,
		Title:           template.Title,
		Slug:            template.Slug,
		Excerpt:         template.Excerpt,
		Content:         template.Content,
		Author:          template.Author,
		PublishedAt:     template.PublishedAt,
		Price:           template.Price,
		Currency:        template.Currency,
		Address:         template.Address,
		Telephone:       template.Telephone,
		MetaTitle:       template.MetaTitle,
		MetaKeywords:    template.MetaKeywords,
		MetaDescription: template.MetaDescription,
		OgTitle:         template.OgTitle,
		OgDescription:   template.OgDescription,
		OgImage:         template.OgImage,
		OgType:          template.OgType,
		OgSiteName:      template.OgSiteName,
		ExtraFields:     template.ExtraFields,
	}

	// جایگزینی متغیرها در فیلدها
	schema = s.replaceTemplateVariables(schema, data)

	return schema, nil
}

// replaceTemplateVariables جایگزینی متغیرهای تمپلیت با مقادیر واقعی
func (s *SchemaService) replaceTemplateVariables(schema *models.Schema, data map[string]interface{}) *models.Schema {
	// این تابع متغیرهای {{variable}} را در فیلدهای اسکیما جایگزین می‌کند
	// مثال: {{title}} -> "عنوان مقاله"

	// جایگزینی در فیلدهای متنی
	if title, ok := data["title"].(string); ok {
		schema.Title = title
	}
	if slug, ok := data["slug"].(string); ok {
		schema.Slug = slug
	}
	if excerpt, ok := data["excerpt"].(string); ok {
		schema.Excerpt = excerpt
	}
	if content, ok := data["content"].(string); ok {
		schema.Content = content
	}
	if author, ok := data["author"].(string); ok {
		schema.Author = author
	}

	// جایگزینی در فیلدهای متا
	if metaTitle, ok := data["meta_title"].(string); ok {
		schema.MetaTitle = metaTitle
	}
	if metaDescription, ok := data["meta_description"].(string); ok {
		schema.MetaDescription = metaDescription
	}
	if metaKeywords, ok := data["meta_keywords"].(string); ok {
		schema.MetaKeywords = metaKeywords
	}

	// جایگزینی در فیلدهای Open Graph
	if ogTitle, ok := data["og_title"].(string); ok {
		schema.OgTitle = ogTitle
	}
	if ogDescription, ok := data["og_description"].(string); ok {
		schema.OgDescription = ogDescription
	}
	if ogImage, ok := data["og_image"].(string); ok {
		schema.OgImage = ogImage
	}

	// جایگزینی در فیلدهای اضافی
	if _, ok := data["extra_fields"].(map[string]interface{}); ok {
		// اینجا می‌توانید منطق جایگزینی در ExtraFields را پیاده‌سازی کنید
		// schema.ExtraFields = extraFields
	}

	return schema
}

// GetSchemaTypes دریافت انواع مختلف اسکیما
func (s *SchemaService) GetSchemaTypes() []string {
	return []string{
		"Article",
		"Product",
		"Organization",
		"WebPage",
		"BreadcrumbList",
		"FAQPage",
		"LocalBusiness",
		"Review",
	}
}

// ValidateTemplate اعتبارسنجی تمپلیت
func (s *SchemaService) ValidateTemplate(schema *models.Schema) error {
	if schema.Name == "" {
		return errors.New("نام تمپلیت الزامی است")
	}
	if schema.Type == "" {
		return errors.New("نوع اسکیما الزامی است")
	}

	// بررسی تکراری نبودن نام
	var count int64
	s.db.Model(&models.Schema{}).
		Where("name = ? AND is_template = ? AND id != ?", schema.Name, true, schema.ID).
		Count(&count)

	if count > 0 {
		return errors.New("نام تمپلیت تکراری است")
	}

	return nil
}
