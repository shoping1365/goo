package migrations

import "gorm.io/gorm"

// Up104HeaderFooterPermissions اضافه کردن مجوزهای مربوط به هدر و فوتر
func Up104HeaderFooterPermissions(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// اضافه کردن مجوزهای هدر و فوتر
		permissions := []struct {
			name        string
			displayName string
			module      string
			subModule   string
			action      string
			resource    string
		}{
			// مجوزهای هدر
			{
				name:        "header_manage",
				displayName: "مدیریت کامل هدر",
				module:      "cms",
				subModule:   "header",
				action:      "manage",
				resource:    "header",
			},
			{
				name:        "header_view",
				displayName: "مشاهده هدر",
				module:      "cms",
				subModule:   "header",
				action:      "view",
				resource:    "header",
			},
			{
				name:        "header_create",
				displayName: "ایجاد هدر",
				module:      "cms",
				subModule:   "header",
				action:      "create",
				resource:    "header",
			},
			{
				name:        "header_edit",
				displayName: "ویرایش هدر",
				module:      "cms",
				subModule:   "header",
				action:      "edit",
				resource:    "header",
			},
			{
				name:        "header_delete",
				displayName: "حذف هدر",
				module:      "cms",
				subModule:   "header",
				action:      "delete",
				resource:    "header",
			},
			// مجوزهای فوتر
			{
				name:        "footer_manage",
				displayName: "مدیریت کامل فوتر",
				module:      "cms",
				subModule:   "footer",
				action:      "manage",
				resource:    "footer",
			},
			{
				name:        "footer_view",
				displayName: "مشاهده فوتر",
				module:      "cms",
				subModule:   "footer",
				action:      "view",
				resource:    "footer",
			},

			// مجوزهای چت
			{
				name:        "support_live_chat.read",
				displayName: "مشاهده چت زنده",
				module:      "support",
				subModule:   "live_chat",
				action:      "read",
				resource:    "support_live_chat",
			},
			{
				name:        "chat_settings.read",
				displayName: "مشاهده تنظیمات چت",
				module:      "support",
				subModule:   "live_chat",
				action:      "read",
				resource:    "chat_settings",
			},
			{
				name:        "chat_settings.write",
				displayName: "مدیریت تنظیمات چت",
				module:      "support",
				subModule:   "live_chat",
				action:      "write",
				resource:    "chat_settings",
			},
			{
				name:        "chat_metrics.read",
				displayName: "مشاهده متریک‌های چت",
				module:      "support",
				subModule:   "live_chat",
				action:      "read",
				resource:    "chat_metrics",
			},
			{
				name:        "footer_create",
				displayName: "ایجاد فوتر",
				module:      "cms",
				subModule:   "footer",
				action:      "create",
				resource:    "footer",
			},
			{
				name:        "footer_edit",
				displayName: "ویرایش فوتر",
				module:      "cms",
				subModule:   "footer",
				action:      "edit",
				resource:    "footer",
			},
			{
				name:        "footer_delete",
				displayName: "حذف فوتر",
				module:      "cms",
				subModule:   "footer",
				action:      "delete",
				resource:    "footer",
			},
		}

		for _, perm := range permissions {
			// بررسی وجود مجوز
			var count int64
			if err := tx.Raw("SELECT COUNT(1) FROM permissions WHERE name = ?", perm.name).Scan(&count).Error; err != nil {
				return err
			}

			// اگر مجوز وجود نداشت، اضافه کن
			if count == 0 {
				if err := tx.Exec(`
					INSERT INTO permissions (name, display_name, module, sub_module, action, resource, is_active, created_at, updated_at)
					VALUES (?, ?, ?, ?, ?, ?, true, NOW(), NOW())
				`, perm.name, perm.displayName, perm.module, perm.subModule, perm.action, perm.resource).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

// Down104HeaderFooterPermissions حذف مجوزهای مربوط به هدر و فوتر
func Down104HeaderFooterPermissions(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// حذف مجوزهای هدر و فوتر و چت
		headerFooterPermissions := []string{
			"header_manage", "header_view", "header_create", "header_edit", "header_delete",
			"footer_manage", "footer_view", "footer_create", "footer_edit", "footer_delete",
			"support_live_chat.read", "chat_settings.read", "chat_settings.write", "chat_metrics.read",
		}

		for _, permName := range headerFooterPermissions {
			_ = tx.Exec("DELETE FROM permissions WHERE name = ?", permName)
		}

		return nil
	})
}
