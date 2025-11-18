package migrations

import "gorm.io/gorm"

// Up112MediaLibraryPermissions اضافه کردن دسترسی‌های کتابخانه رسانه
func Up112MediaLibraryPermissions(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن permission برای کتابخانه رسانه
		_ = tx.Exec("INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('media_library.read','مشاهده کتابخانه رسانه','media_management','library','read','media_library',true) ON CONFLICT (name) DO NOTHING")

		// اختصاص permission به نقش‌های admin و developer
		_ = tx.Exec("INSERT INTO role_permissions(role_id, permission_id) SELECT r.id, p.id FROM roles r, permissions p WHERE r.name IN ('admin','developer') AND p.name = 'media_library.read' ON CONFLICT DO NOTHING")

		return nil
	})
}
