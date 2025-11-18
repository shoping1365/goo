package migrations

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Up065EnsureDevUser اطمینان از وجود کاربر dev با role_id=100 و پسورد مشخص (idempotent)
func Up065EnsureDevUser(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 1) اطمینان از وجود نقش با id=100 (برای سازگاری با میدلورهایی که role_id=100 را developer می‌دانند)
		// ابتدا تلاش برای درج با نام developer، اگر نام قبلاً وجود داشت، درج با نام developer_100
		if err := tx.Exec(`
            INSERT INTO roles(id, name, display_name, is_system, priority, is_active)
            SELECT 100, 'developer', 'Developer', true, 0, true
            WHERE NOT EXISTS (SELECT 1 FROM roles WHERE id = 100)
              AND NOT EXISTS (SELECT 1 FROM roles WHERE name = 'developer');
        `).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            INSERT INTO roles(id, name, display_name, is_system, priority, is_active)
            SELECT 100, 'developer_100', 'Developer', true, 0, true
            WHERE NOT EXISTS (SELECT 1 FROM roles WHERE id = 100);
        `).Error; err != nil {
			return err
		}

		// 2) ساخت هش پسورد «1365»
		hashed, _ := bcrypt.GenerateFromPassword([]byte("1365"), bcrypt.DefaultCost)

		// 3) ایجاد کاربر dev اگر وجود ندارد (با role_id=100)
		if err := tx.Exec(`
            INSERT INTO users(username, name, mobile, email, password_hash, role_id, status)
            SELECT 'dev', 'Dev', '09203214155', 'dev@example.com', ?, 100, 'active'
            WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = 'dev');
        `, string(hashed)).Error; err != nil {
			return err
		}

		// 4) اگر dev وجود دارد ولی پسورد ندارد یا نقش 100 نیست، به‌روزرسانی کن
		if err := tx.Exec(`
            UPDATE users
            SET password_hash = COALESCE(NULLIF(password_hash, ''), ?),
                role_id = 100,
                status = 'active'
            WHERE username = 'dev';
        `, string(hashed)).Error; err != nil {
			return err
		}

		return nil
	})
}
