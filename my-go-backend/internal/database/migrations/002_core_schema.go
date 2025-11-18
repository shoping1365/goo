package migrations

import (
	"gorm.io/gorm"
)

// Up002CoreSchema ایجاد اسکیمای هسته (بدون AutoMigrate)
// شامل: users, roles, permissions, role_permissions, user_roles, sessions, settings
// نکته: این مایگریشن باید بدون خطا چندباره قابل اجرا باشد (IF NOT EXISTS / ON CONFLICT DO NOTHING)
func Up002CoreSchema(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// Extensions
		// Note: pgcrypto extension is optional and may not be available on all PostgreSQL installations
		// Ignore error if extension is not available
		if err := tx.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto;`); err != nil {
			// Log warning but continue - pgcrypto is optional
			tx.Logger.Warn(tx.Statement.Context, "Warning: pgcrypto extension not available, continuing without it")
		}

		// settings
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS settings (
                id           BIGSERIAL PRIMARY KEY,
                key          VARCHAR(255) NOT NULL UNIQUE,
                value        TEXT,
                description  TEXT,
                category     VARCHAR(50) NOT NULL,
                type         VARCHAR(20),
                is_public    BOOLEAN NOT NULL DEFAULT FALSE,
                created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                created_by   VARCHAR(100),
                updated_by   VARCHAR(100)
            );
        `).Error; err != nil {
			return err
		}
		// ستون soft-delete برای سازگاری با GORM
		_ = tx.Exec(`ALTER TABLE settings ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMPTZ`)

		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_settings_category ON settings(category);`).Error; err != nil {
			return err
		}

		// roles
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS roles (
                id           BIGSERIAL PRIMARY KEY,
                name         VARCHAR(100) NOT NULL UNIQUE,
                display_name VARCHAR(150) NOT NULL,
                description  TEXT,
                is_active    BOOLEAN NOT NULL DEFAULT TRUE,
                is_system    BOOLEAN NOT NULL DEFAULT FALSE,
                priority     INT NOT NULL DEFAULT 0,
                created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_roles_active ON roles(is_active);`).Error; err != nil {
			return err
		}

		// permissions
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS permissions (
                id           BIGSERIAL PRIMARY KEY,
                name         VARCHAR(100) NOT NULL UNIQUE,
                display_name VARCHAR(150) NOT NULL,
                description  TEXT,
                module       VARCHAR(50) NOT NULL,
                sub_module   VARCHAR(50),
                action       VARCHAR(50) NOT NULL,
                resource     VARCHAR(50) NOT NULL,
                is_active    BOOLEAN NOT NULL DEFAULT TRUE,
                priority     INT NOT NULL DEFAULT 0,
                created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_permissions_active ON permissions(is_active);`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_permissions_module_action ON permissions(module, action);`).Error; err != nil {
			return err
		}

		// role_permissions (junction)
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS role_permissions (
                id            BIGSERIAL PRIMARY KEY,
                role_id       BIGINT NOT NULL,
                permission_id BIGINT NOT NULL,
                granted_by    BIGINT,
                granted_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                expires_at    TIMESTAMPTZ,
                is_active     BOOLEAN NOT NULL DEFAULT TRUE,
                created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_rp_role FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE,
                CONSTRAINT fk_rp_perm FOREIGN KEY(permission_id) REFERENCES permissions(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_role_permissions_unique ON role_permissions(role_id, permission_id);`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_role_permissions_active ON role_permissions(is_active);`).Error; err != nil {
			return err
		}

		// users
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS users (
                id                 BIGSERIAL PRIMARY KEY,
                username           VARCHAR(100) UNIQUE,
                email              VARCHAR(255) UNIQUE,
                mobile             VARCHAR(20) NOT NULL UNIQUE,
                name               VARCHAR(100),
                password_hash      VARCHAR(255),
                role_id            BIGINT NOT NULL DEFAULT 1,
                status             VARCHAR(20) NOT NULL DEFAULT 'active',
                email_verified     BOOLEAN NOT NULL DEFAULT FALSE,
                mobile_verified    BOOLEAN NOT NULL DEFAULT TRUE,
                failed_login_count INT NOT NULL DEFAULT 0,
                locked_until       TIMESTAMPTZ,
                last_login_at      TIMESTAMPTZ,
                last_login_ip      VARCHAR(45),
                registered_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                preferences        JSONB NOT NULL DEFAULT '{}'::jsonb,
                security_settings  JSONB NOT NULL DEFAULT '{}'::jsonb,
                profile_data       JSONB NOT NULL DEFAULT '{}'::jsonb,
                CONSTRAINT fk_users_role FOREIGN KEY(role_id) REFERENCES roles(id)
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_users_status ON users(status);`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_users_locked_until ON users(locked_until);`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_users_failed_login_count ON users(failed_login_count);`).Error; err != nil {
			return err
		}

		// user_roles (many-to-many per user)
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS user_roles (
                id         BIGSERIAL PRIMARY KEY,
                user_id    BIGINT NOT NULL,
                role_id    BIGINT NOT NULL,
                valid_from TIMESTAMPTZ,
                valid_to   TIMESTAMPTZ,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_ur_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
                CONSTRAINT fk_ur_role FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_user_roles_unique ON user_roles(user_id, role_id);`).Error; err != nil {
			return err
		}

		// sessions
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS sessions (
                id           BIGSERIAL PRIMARY KEY,
                user_id      BIGINT,
                token        VARCHAR(255) UNIQUE,
                ip_address   VARCHAR(45),
                user_agent   VARCHAR(255),
                expires_at   TIMESTAMPTZ NOT NULL,
                mobile       VARCHAR(20),
                auth_method  VARCHAR(50) NOT NULL DEFAULT 'password',
                is_verified  BOOLEAN NOT NULL DEFAULT TRUE,
                verified_at  TIMESTAMPTZ,
                is_active    BOOLEAN NOT NULL DEFAULT TRUE,
                last_used_at TIMESTAMPTZ,
                created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_sessions_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_sessions_expires_active ON sessions(expires_at, is_active);`).Error; err != nil {
			return err
		}

		return nil
	})
}
