package migrations

import "gorm.io/gorm"

// Up042UserRolesM2MAndSecurityIndexes اطمینان از user_roles و ایندکس‌های امنیتی users
func Up042UserRolesM2MAndSecurityIndexes(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		_ = tx.Exec(`CREATE TABLE IF NOT EXISTS user_roles (
            id BIGSERIAL PRIMARY KEY,
            user_id BIGINT NOT NULL,
            role_id BIGINT NOT NULL,
            valid_from TIMESTAMPTZ,
            valid_to TIMESTAMPTZ,
            created_at TIMESTAMPTZ DEFAULT NOW(),
            updated_at TIMESTAMPTZ DEFAULT NOW()
        );`)
		_ = tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_user_roles_unique ON user_roles(user_id, role_id)`)
		_ = tx.Exec(`CREATE INDEX IF NOT EXISTS idx_users_locked_until ON users(locked_until)`)
		_ = tx.Exec(`CREATE INDEX IF NOT EXISTS idx_users_failed_login_count ON users(failed_login_count)`)
		_ = tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_role_permissions_unique ON role_permissions(role_id, permission_id)`)
		return nil
	})
}
