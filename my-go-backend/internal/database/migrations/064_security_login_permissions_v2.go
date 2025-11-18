package migrations

import "gorm.io/gorm"

// Up064SecurityLoginPermissionsV2 ایجاد دسترسی‌های امنیت ورود (ادغام 64)
func Up064SecurityLoginPermissionsV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// permissions
		_ = tx.Exec("INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('security_login.read','مشاهده آمار/تاریخچه ورود','system_security','login','read','security_login',true) ON CONFLICT (name) DO NOTHING")
		_ = tx.Exec("INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('security_login.create','مسدود کردن IP در ورود','system_security','login','create','security_login',true) ON CONFLICT (name) DO NOTHING")
		_ = tx.Exec("INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active) VALUES ('security_login.update','آزاد کردن IP در ورود','system_security','login','update','security_login',true) ON CONFLICT (name) DO NOTHING")
		// assign to admin/developer if roles exist
		_ = tx.Exec("INSERT INTO role_permissions(role_id, permission_id) SELECT r.id, p.id FROM roles r, permissions p WHERE r.name IN ('admin','developer') AND p.name IN ('security_login.read','security_login.create','security_login.update') ON CONFLICT DO NOTHING")
		return nil
	})
}
