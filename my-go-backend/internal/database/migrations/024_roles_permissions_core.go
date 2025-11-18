package migrations

import "gorm.io/gorm"

// Up024RolesPermissionsCore جداول نقش‌ها/مجوزها و ارتباط آن‌ها (در صورتی‌که 002 آن‌ها را نساخته باشد)
func Up024RolesPermissionsCore(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS roles (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL UNIQUE,
                display_name VARCHAR(150) NOT NULL,
                description TEXT,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                is_system BOOLEAN NOT NULL DEFAULT FALSE,
                priority INT NOT NULL DEFAULT 0,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS permissions (
                id BIGSERIAL PRIMARY KEY,
                name VARCHAR(100) NOT NULL UNIQUE,
                display_name VARCHAR(150) NOT NULL,
                description TEXT,
                module VARCHAR(50) NOT NULL,
                sub_module VARCHAR(50),
                action VARCHAR(50) NOT NULL,
                resource VARCHAR(50) NOT NULL,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                priority INT NOT NULL DEFAULT 0,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS role_permissions (
                id BIGSERIAL PRIMARY KEY,
                role_id BIGINT NOT NULL,
                permission_id BIGINT NOT NULL,
                granted_by BIGINT,
                granted_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                expires_at TIMESTAMPTZ,
                is_active BOOLEAN NOT NULL DEFAULT TRUE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		_ = tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_role_permissions_unique ON role_permissions(role_id, permission_id)`)
		return nil
	})
}
