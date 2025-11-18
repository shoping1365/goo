package migrations

import "gorm.io/gorm"

// Up119SEORedirectsPermissions ایجاد مجوزهای مربوط به SEO redirects
func Up119SEORedirectsPermissions(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// ایجاد مجوزهای SEO redirects
		permissions := []struct {
			name        string
			displayName string
			module      string
			subModule   string
			action      string
			resource    string
		}{
			{"developer_migration", "مدیریت مهاجرت و توسعه", "developer", "migration", "manage", "seo_redirects"},
			{"seo_redirects.read", "مشاهده ریدایرکت‌های سئو", "seo", "redirects", "read", "seo_redirects"},
			{"seo_redirects.create", "ایجاد ریدایرکت سئو", "seo", "redirects", "create", "seo_redirects"},
			{"seo_redirects.update", "ویرایش ریدایرکت سئو", "seo", "redirects", "update", "seo_redirects"},
			{"seo_redirects.delete", "حذف ریدایرکت سئو", "seo", "redirects", "delete", "seo_redirects"},
		}

		for _, perm := range permissions {
			if err := tx.Exec(`
				INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active, created_at, updated_at)
				VALUES (?, ?, ?, ?, ?, ?, true, NOW(), NOW())
				ON CONFLICT (name) DO NOTHING
			`, perm.name, perm.displayName, perm.module, perm.subModule, perm.action, perm.resource).Error; err != nil {
				return err
			}
		}

		// اختصاص مجوزها به نقش developer
		if err := tx.Exec(`
			INSERT INTO role_permissions(role_id, permission_id) 
			SELECT r.id, p.id 
			FROM roles r, permissions p 
			WHERE r.name = 'developer' AND p.name IN ('developer_migration', 'seo_redirects.read', 'seo_redirects.create', 'seo_redirects.update', 'seo_redirects.delete')
			ON CONFLICT DO NOTHING
		`).Error; err != nil {
			return err
		}

		// اختصاص مجوزها به نقش admin
		if err := tx.Exec(`
			INSERT INTO role_permissions(role_id, permission_id) 
			SELECT r.id, p.id 
			FROM roles r, permissions p 
			WHERE r.name = 'admin' AND p.name IN ('seo_redirects.read', 'seo_redirects.create', 'seo_redirects.update', 'seo_redirects.delete')
			ON CONFLICT DO NOTHING
		`).Error; err != nil {
			return err
		}

		return nil
	})
}

