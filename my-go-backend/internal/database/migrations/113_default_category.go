package migrations

import "gorm.io/gorm"

// Up113DefaultCategory ایجاد دسته‌بندی پیش‌فرض سیستم
func Up113DefaultCategory(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// بررسی وجود دسته‌بندی پیش‌فرض
		var count int64
		if err := tx.Model(&struct {
			ID uint `gorm:"primaryKey"`
		}{}).Table("categories").Where("slug = ?", "default").Count(&count).Error; err != nil {
			return err
		}

		// اگر وجود ندارد، ایجاد کن
		if count == 0 {
			if err := tx.Exec(`
				INSERT INTO categories (name, slug, published, "order", created_at, updated_at)
				VALUES ('دسته‌بندی پیش‌فرض', 'default', false, 0, NOW(), NOW())
			`).Error; err != nil {
				return err
			}
		}

		// اضافه کردن permissions برای مدیریت دسته‌بندی محصولات
		_ = tx.Exec("INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('categories_manage','مدیریت دسته‌بندی محصولات','product_management','categories','manage','categories',true) ON CONFLICT (name) DO NOTHING")
		_ = tx.Exec("INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('category.create','ایجاد دسته‌بندی محصولات','product_management','categories','create','category',true) ON CONFLICT (name) DO NOTHING")
		_ = tx.Exec("INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('category.update','ویرایش دسته‌بندی محصولات','product_management','categories','update','category',true) ON CONFLICT (name) DO NOTHING")
		_ = tx.Exec("INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('category.delete','حذف دسته‌بندی محصولات','product_management','categories','delete','category',true) ON CONFLICT (name) DO NOTHING")
		_ = tx.Exec("INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('developer_maintenance','دسترسی به ابزارهای توسعه','system','maintenance','manage','developer_tools',true) ON CONFLICT (name) DO NOTHING")

		// اختصاص permissions به نقش‌های admin و developer
		_ = tx.Exec("INSERT INTO role_permissions(role_id, permission_id) SELECT r.id, p.id FROM roles r, permissions p WHERE r.name IN ('admin','developer') AND p.name IN ('categories_manage','category.create','category.update','category.delete','developer_maintenance') ON CONFLICT DO NOTHING")

		return nil
	})
}
