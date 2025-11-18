package migrations

import "gorm.io/gorm"

// Up111SeedBaseRoles اضافه کردن نقش‌های پایه سیستم
func Up111SeedBaseRoles(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// نقش‌های پایه سیستم
		baseRoles := []struct {
			id          int
			name        string
			displayName string
			description string
			isSystem    bool
			priority    int
		}{
			{1, "customer", "مشتری", "کاربران عادی سایت", true, 1},
			{2, "admin", "مدیر", "مدیر سیستم", true, 10},
			{100, "developer", "توسعه‌دهنده", "توسعه‌دهنده سیستم", true, 100},
		}

		for _, role := range baseRoles {
			if err := tx.Exec(`
				INSERT INTO roles (id, name, display_name, description, is_system, priority, is_active, created_at, updated_at)
				VALUES (?, ?, ?, ?, ?, ?, true, NOW(), NOW())
				ON CONFLICT (id) DO UPDATE SET 
					name = EXCLUDED.name,
					display_name = EXCLUDED.display_name,
					description = EXCLUDED.description,
					is_system = EXCLUDED.is_system,
					priority = EXCLUDED.priority,
					is_active = true,
					updated_at = NOW();
			`, role.id, role.name, role.displayName, role.description, role.isSystem, role.priority).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
