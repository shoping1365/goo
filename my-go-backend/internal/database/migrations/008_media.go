package migrations

import "gorm.io/gorm"

// Up008Media ایجاد جداول رسانه اصلی: media_files, media_variants, media_versions, media_delete_logs, compression_jobs/settings/stats, backups
func Up008Media(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// media_files
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS media_files (
                id BIGSERIAL PRIMARY KEY,
                file_name VARCHAR(255) NOT NULL,
                title VARCHAR(150),
                alt_text VARCHAR(255),
                display_title VARCHAR(150),
                short_description VARCHAR(255),
                description TEXT,
                file_type VARCHAR(30) NOT NULL,
                mime_type VARCHAR(50),
                format VARCHAR(20),
                original_format VARCHAR(20),
                size BIGINT NOT NULL,
                compressed BOOLEAN NOT NULL DEFAULT FALSE,
                compressed_size BIGINT,
                url VARCHAR(255) NOT NULL,
                duration_seconds DOUBLE PRECISION,
                width INT,
                height INT,
                bitrate_kbps INT,
                category VARCHAR(30),
                uploaded_by BIGINT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_media_files_category_created ON media_files(category, created_at DESC) WHERE deleted_at IS NULL`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_media_files_uploaded_by_created ON media_files(uploaded_by, created_at DESC) WHERE deleted_at IS NULL`).Error; err != nil {
			return err
		}

		// media_variants
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS media_variants (
                id BIGSERIAL PRIMARY KEY,
                media_id BIGINT NOT NULL,
                width INT,
                height INT,
                file_path VARCHAR(255) NOT NULL,
                file_size BIGINT,
                purpose VARCHAR(30),
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ,
                CONSTRAINT fk_media_variants_media FOREIGN KEY(media_id) REFERENCES media_files(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}

		// media_versions
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS media_versions (
                id BIGSERIAL PRIMARY KEY,
                media_id BIGINT NOT NULL,
                version_code VARCHAR(50) NOT NULL,
                file_path VARCHAR(255) NOT NULL,
                file_size BIGINT,
                format VARCHAR(20),
                quality VARCHAR(20),
                is_active BOOLEAN NOT NULL DEFAULT FALSE,
                is_backup BOOLEAN NOT NULL DEFAULT FALSE,
                compression_ratio REAL,
                compression_method VARCHAR(50),
                created_by BIGINT,
                meta JSON,
                parent_version_id BIGINT,
                backup_path VARCHAR(255),
                note TEXT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ,
                CONSTRAINT fk_media_versions_media FOREIGN KEY(media_id) REFERENCES media_files(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}

		// media_delete_logs
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS media_delete_logs (
                id BIGSERIAL PRIMARY KEY,
                rel_path VARCHAR(255),
                deleted_by BIGINT,
                note TEXT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// compression_jobs
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS compression_jobs (
                id BIGSERIAL PRIMARY KEY,
                media_id BIGINT NOT NULL,
                version_id BIGINT,
                requested_by BIGINT,
                status VARCHAR(20) NOT NULL,
                started_at TIMESTAMPTZ,
                finished_at TIMESTAMPTZ,
                error_message TEXT,
                target_format VARCHAR(20),
                target_quality VARCHAR(20),
                batch_id BIGINT,
                job_type VARCHAR(30) NOT NULL,
                progress INT NOT NULL DEFAULT 0,
                log TEXT,
                original_size BIGINT,
                compressed_size BIGINT,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                CONSTRAINT fk_comp_jobs_media FOREIGN KEY(media_id) REFERENCES media_files(id) ON DELETE CASCADE
            );
        `).Error; err != nil {
			return err
		}

		// compression_settings
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS compression_settings (
                id BIGSERIAL PRIMARY KEY,
                scope VARCHAR(20) NOT NULL,
                user_id BIGINT,
                media_id BIGINT,
                job_id BIGINT,
                key VARCHAR(50) NOT NULL,
                value JSON NOT NULL,
                is_default BOOLEAN NOT NULL DEFAULT FALSE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		// compression_stats
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS compression_stats (
                id BIGSERIAL PRIMARY KEY,
                stat_type VARCHAR(50) NOT NULL,
                value VARCHAR(100) NOT NULL,
                media_id BIGINT,
                user_id BIGINT,
                format VARCHAR(20),
                period VARCHAR(20),
                meta JSON,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                deleted_at TIMESTAMPTZ
            );
        `).Error; err != nil {
			return err
		}

		// backups
		if err := tx.Exec(`
            CREATE TABLE IF NOT EXISTS backups (
                id BIGSERIAL PRIMARY KEY,
                rel_path VARCHAR(255) UNIQUE,
                period CHAR(7),
                archived BOOLEAN NOT NULL DEFAULT FALSE,
                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
            );
        `).Error; err != nil {
			return err
		}

		return nil
	})
}
