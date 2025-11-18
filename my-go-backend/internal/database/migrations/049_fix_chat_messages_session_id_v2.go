package migrations

import "gorm.io/gorm"

// Up049FixChatMessagesSessionIDV2 تبدیل session_id متنی به BIGINT و FK به chat_sessions(id) به‌صورت ایمن
func Up049FixChatMessagesSessionIDV2(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// تبدیل ایمن: اگر ستون session_id از نوع متنی بود، به id تبدیل می‌کنیم؛ اگر از قبل BIGINT است، فقط قیود و ایندکس‌ها را تضمین می‌کنیم
		_ = tx.Exec(`DO $$
        DECLARE v_type text;
        BEGIN
            SELECT data_type INTO v_type FROM information_schema.columns 
            WHERE table_name = 'chat_messages' AND column_name = 'session_id';

            -- ستون موقت برای نگهداری id جلسه
            IF NOT EXISTS (
                SELECT 1 FROM information_schema.columns 
                WHERE table_name = 'chat_messages' AND column_name = 'session_id_new'
            ) THEN
                EXECUTE 'ALTER TABLE chat_messages ADD COLUMN session_id_new BIGINT';
            END IF;

            IF v_type IN ('character varying','text') THEN
                -- backfill بر اساس session_id متنی
                EXECUTE 'UPDATE chat_messages cm 
                         SET session_id_new = cs.id 
                         FROM chat_sessions cs 
                         WHERE cm.session_id = cs.session_id 
                           AND cm.session_id_new IS NULL';

                -- حذف قیود قدیمی در صورت وجود
                IF EXISTS (
                    SELECT 1 FROM information_schema.table_constraints 
                    WHERE table_name = 'chat_messages' AND constraint_name = 'fk_chat_messages_session'
                ) THEN
                    EXECUTE 'ALTER TABLE chat_messages DROP CONSTRAINT fk_chat_messages_session';
                END IF;

                -- جایگزینی ستون‌ها
                IF EXISTS (
                    SELECT 1 FROM information_schema.columns 
                    WHERE table_name='chat_messages' AND column_name='session_id'
                ) THEN
                    EXECUTE 'ALTER TABLE chat_messages DROP COLUMN session_id';
                END IF;
                EXECUTE 'ALTER TABLE chat_messages RENAME COLUMN session_id_new TO session_id';
            ELSE
                -- اگر از قبل BIGINT است، مقداردهی لازم نیست؛ ستون موقت را در صورت وجود حذف کن
                IF EXISTS (
                    SELECT 1 FROM information_schema.columns 
                    WHERE table_name='chat_messages' AND column_name='session_id_new'
                ) THEN
                    EXECUTE 'ALTER TABLE chat_messages DROP COLUMN session_id_new';
                END IF;
            END IF;

            -- افزودن FK در صورت عدم وجود
            IF NOT EXISTS (
                SELECT 1 FROM information_schema.table_constraints 
                WHERE table_name = 'chat_messages' AND constraint_name = 'fk_chat_messages_session'
            ) THEN
                EXECUTE 'ALTER TABLE chat_messages ADD CONSTRAINT fk_chat_messages_session 
                         FOREIGN KEY (session_id) REFERENCES chat_sessions(id) ON DELETE CASCADE';
            END IF;
        END$$;`)

		// ایندکس‌ها
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_chat_messages_session_id ON chat_messages(session_id)")
		_ = tx.Exec("CREATE INDEX IF NOT EXISTS idx_chat_messages_session_created_at ON chat_messages(session_id, created_at DESC)")
		return nil
	})
}
