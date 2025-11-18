package services

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"my-go-backend/internal/models"
)

// ProductSpecService handles validation & persistence of product specifications.
type ProductSpecService struct {
	DB *gorm.DB
}

func NewProductSpecService(db *gorm.DB) *ProductSpecService {
	return &ProductSpecService{DB: db}
}

// SpecInput represents incoming JSON for a single attribute value.
type SpecInput struct {
	AttributeID uint    `json:"attribute_id"`
	OptionID    *uint   `json:"option_id,omitempty"`  // تک‌انتخاب
	OptionIDs   []uint  `json:"option_ids,omitempty"` // چندانتخاب
	ValueText   *string `json:"value_text,omitempty"`
	Sort        *uint   `json:"sort,omitempty"`
}

// ValidateAndSave validates inputs and persists them inside a transaction.
func (s *ProductSpecService) ValidateAndSave(productID uint, categoryID uint, specs []SpecInput) error {
	// اگر کاربر مشخصات ارسال نکرد، به‌صورت امن رد می‌شویم (درخواست مجاز است)
	if len(specs) == 0 {
		return nil
	}

	// Load category attribute schema (could be cached)
	schemaSvc := NewAttributeSchemaService(s.DB)
	groups, err := schemaSvc.GetCategorySchema(categoryID)
	if err != nil {
		return err
	}

	// map of attribute meta for quick lookup
	meta := make(map[uint]AttributeSchemaElement)
	requiredMissing := make(map[uint]bool)
	for _, g := range groups {
		for _, a := range g.Attributes {
			meta[a.ID] = a
			if a.IsRequired {
				requiredMissing[a.ID] = true
			}
		}
	}

	// validation loop
	for _, sp := range specs {
		m, ok := meta[sp.AttributeID]
		if !ok {
			return errors.New("attribute not allowed for category")
		}
		delete(requiredMissing, sp.AttributeID)

		// Normalise type to handle both Persian labels and English slugs stored in DB.
		t := strings.ToLower(strings.TrimSpace(m.Type))
		ct := strings.ToLower(strings.TrimSpace(m.ControlType))

		// هوشمندتر: ابتدا نوع «متنی» را تشخیص دهیم، سپس انتخابی؛ در صورت ناشناخته بودن، یکی از دو مقدار باید وجود داشته باشد.
		isTextAttr := t == "custom_text" || t == "متن سفارشی" || strings.HasPrefix(ct, "textbox")

		isSingleSelect := t == "select" || t == "انتخاب" || strings.HasPrefix(ct, "dropdown_single")

		isMultiSelect := t == "multi_select" || t == "چندانتخاب" || t == "چند انتخابی" || strings.HasPrefix(ct, "dropdown_multi")

		if isTextAttr {
			// قانون (موقتی): فقط بررسی کنیم که مقدار متنی وجود داشته باشد؛ بقیهٔ بررسی‌ها غیرفعال شد.
			if sp.ValueText == nil || len(strings.TrimSpace(*sp.ValueText)) == 0 {
				return fmt.Errorf("برای ویژگی «%s» مقدار وارد کنید", m.Name)
			}
			if sp.OptionID != nil || len(sp.OptionIDs) > 0 {
				return fmt.Errorf("ویژگی «%s» متنی است، نباید گزینه انتخاب شود", m.Name)
			}
		} else if isSingleSelect {
			if sp.OptionID == nil {
				return fmt.Errorf("برای ویژگی «%s» باید یک گزینه انتخاب شود", m.Name)
			}
			// ⬇️ بررسی مقدار متنی موقتاً غیرفعال شد
		} else if isMultiSelect {
			if len(sp.OptionIDs) == 0 {
				return fmt.Errorf("برای ویژگی «%s» حداقل یک گزینه انتخاب کنید", m.Name)
			}
			// ⬇️ بررسی مقدار متنی و option_id موقتاً غیرفعال شد
		} else {
			// اگر نوع ناشناخته بود، حداقل یکی از مقادیر باید پر باشد (تطابق قبلی)
			if sp.OptionID == nil && len(sp.OptionIDs) == 0 && (sp.ValueText == nil || len(strings.TrimSpace(*sp.ValueText)) == 0) {
				return fmt.Errorf("برای ویژگی «%s» مقدار وارد کنید", m.Name)
			}
		}
	}

	if len(requiredMissing) > 0 {
		names := make([]string, 0, len(requiredMissing))
		for id := range requiredMissing {
			if m, ok := meta[id]; ok {
				names = append(names, "«"+m.Name+"»")
			}
		}
		// اگر هیچ ورودی‌ای برای specs نیامده، اجبار ویژگی‌های ضروری را موقتاً نادیده می‌گیریم
		// چون کاربر می‌تواند بدون مشخصات محصول را بسازد.
		if len(specs) > 0 {
			return fmt.Errorf("ویژگی‌های ضروری تکمیل نشده‌اند: %s", strings.Join(names, "، "))
		}
	}

	// persist in transaction: delete old & insert new
	return s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("product_id = ?", productID).Delete(&models.ProductAttributeValue{}).Error; err != nil {
			return err
		}
		for _, sp := range specs {
			// Helper to insert one record
			insertOne := func(optID *uint, valText *string) error {
				pav := models.ProductAttributeValue{
					ProductID:        productID,
					AttributeID:      sp.AttributeID,
					AttributeValueID: optID,
					SortOrder:        0,
				}
				if sp.Sort != nil {
					pav.SortOrder = *sp.Sort
				}
				if valText != nil {
					pav.ValueText = valText
				}
				return tx.Create(&pav).Error
			}

			if len(sp.OptionIDs) > 0 { // multi-select
				for _, oid := range sp.OptionIDs {
					idCopy := oid
					if err := insertOne(&idCopy, nil); err != nil {
						return err
					}
				}
			} else if sp.OptionID != nil { // single select
				if err := insertOne(sp.OptionID, nil); err != nil {
					return err
				}
			} else { // text value
				if err := insertOne(nil, sp.ValueText); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
