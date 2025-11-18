package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

// ImageSEOWorker پردازنده دوره‌ای jobهای image_seo_jobs
type ImageSEOWorker struct {
	DB             *gorm.DB
	SettingService *SettingService
	APIService     *APISettingsService
	SEOService     *ImageSEOService
	ticker         *time.Ticker
	stopCh         chan struct{}
}

func NewImageSEOWorker(db *gorm.DB, ss *SettingService, api *APISettingsService, seo *ImageSEOService) *ImageSEOWorker {
	return &ImageSEOWorker{DB: db, SettingService: ss, APIService: api, SEOService: seo, stopCh: make(chan struct{})}
}

// Start اجرای لوپ با بازه 30 ثانیه
func (w *ImageSEOWorker) Start() {
	if w.ticker != nil {
		return
	}
	w.ticker = time.NewTicker(30 * time.Second)
	go func() {
		for {
			select {
			case <-w.ticker.C:
				w.process()
			case <-w.stopCh:
				w.ticker.Stop()
				return
			}
		}
	}()
}

func (w *ImageSEOWorker) Stop() {
	if w.ticker != nil {
		close(w.stopCh)
	}
}

// ScheduleOnUpload درج job با تاخیر تنظیم‌شده
func (w *ImageSEOWorker) ScheduleOnUpload(ctx context.Context, mediaID uint, pageTitle, lang string) {
	delay := 60
	if st, err := w.SettingService.GetSetting(ctx, "image_seo.delay_seconds"); err == nil && st != nil {
		var d int
		fmt.Sscanf(st.Value, "%d", &d)
		if d > 0 && d < 24*3600 {
			delay = d
		}
	}
	sched := time.Now().Add(time.Duration(delay) * time.Second)
	_ = w.DB.Exec(`INSERT INTO image_seo_jobs(media_id, status, scheduled_at, lang, context_title) VALUES(?, 'pending', ?, ?, ?)`, mediaID, sched, lang, pageTitle).Error
}

func (w *ImageSEOWorker) process() {
	ctx := context.Background()
	// enabled?
	if st, err := w.SettingService.GetSetting(ctx, "image_seo.enabled"); err == nil {
		v := strings.ToLower(strings.TrimSpace(st.Value))
		if v == "false" || v == "0" {
			return
		}
	}

	batch := 10
	if st, err := w.SettingService.GetSetting(ctx, "image_seo.batch_size"); err == nil && st != nil {
		var b int
		fmt.Sscanf(st.Value, "%d", &b)
		if b >= 1 && b <= 100 {
			batch = b
		}
	}
	// generate title?
	genTitle := false
	if st, err := w.SettingService.GetSetting(ctx, "image_seo.generate_title"); err == nil && st != nil {
		v := strings.ToLower(strings.TrimSpace(st.Value))
		genTitle = v == "true" || v == "1"
	}
	// overwrite?
	overwrite := false
	if st, err := w.SettingService.GetSetting(ctx, "image_seo.overwrite_existing"); err == nil && st != nil {
		v := strings.ToLower(strings.TrimSpace(st.Value))
		overwrite = v == "true" || v == "1"
	}

	// انتخاب job های آماده
	rows, err := w.DB.Raw(`SELECT id, media_id, lang, COALESCE(context_title,'') FROM image_seo_jobs WHERE status='pending' AND scheduled_at <= NOW() ORDER BY scheduled_at ASC, id ASC LIMIT ?`, batch).Rows()
	if err != nil {
		log.Println("image seo jobs query error:", err)
		return
	}
	defer rows.Close()
	type item struct {
		id      uint
		mediaID uint
		lang    string
		title   string
	}
	var items []item
	for rows.Next() {
		var it item
		var lang sql.NullString
		var title sql.NullString
		if err := rows.Scan(&it.id, &it.mediaID, &lang, &title); err == nil {
			it.lang = lang.String
			it.title = title.String
			items = append(items, it)
		}
	}
	for _, it := range items {
		_ = w.DB.Exec("UPDATE image_seo_jobs SET status='processing', started_at=NOW() WHERE id=?", it.id).Error
		go w.handleOne(it, genTitle, overwrite)
	}
}

func (w *ImageSEOWorker) handleOne(it struct {
	id      uint
	mediaID uint
	lang    string
	title   string
}, genTitle, overwrite bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	// fetch media
	var rec struct {
		ID          uint
		URL         string
		Alt         string
		Title       string
		Caption     string
		Description string
		Mime        string
	}
	if err := w.DB.Raw(`SELECT id, url, COALESCE(alt_text,''), COALESCE(title,''), COALESCE(short_description,''), COALESCE(description,''), mime_type FROM media_files WHERE id=?`, it.mediaID).Scan(&rec).Error; err != nil {
		w.fail(it.id, fmt.Errorf("media not found: %w", err))
		return
	}
	// اگر تصویر نیست، skip
	if !strings.HasPrefix(strings.ToLower(rec.Mime), "image/") {
		_ = w.DB.Exec("UPDATE image_seo_jobs SET status='completed', finished_at=NOW(), output_json='{}'::jsonb WHERE id=?", it.id).Error
		return
	}
	// overwrite یا فقط فیلدهای خالی
	if !overwrite && rec.Alt != "" && rec.Caption != "" && rec.Description != "" {
		_ = w.DB.Exec("UPDATE image_seo_jobs SET status='completed', finished_at=NOW(), output_json='{}'::jsonb WHERE id=?", it.id).Error
		return
	}

	// read image as base64
	img64, err := w.SEOService.ReadFileAsBase64(rec.URL)
	if err != nil {
		w.fail(it.id, err)
		return
	}
	lang := it.lang
	if lang == "" {
		lang = "fa"
	}
	res, err := w.SEOService.CallOpenAI(ctx, img64, it.title, lang, genTitle)
	if err != nil {
		w.fail(it.id, err)
		return
	}

	// ساخت updates
	updates := map[string]interface{}{}
	if overwrite || rec.Alt == "" {
		updates["alt_text"] = strings.TrimSpace(res.AltText)
	}
	if genTitle && (overwrite || rec.Title == "") && strings.TrimSpace(res.Title) != "" {
		updates["title"] = strings.TrimSpace(res.Title)
	}
	if overwrite || rec.Caption == "" {
		updates["short_description"] = strings.TrimSpace(res.Caption)
	}
	if overwrite || rec.Description == "" {
		updates["description"] = strings.TrimSpace(res.Description)
	}
	if len(updates) > 0 {
		if err := w.DB.Model(&struct{}{}).Table("media_files").Where("id = ?", rec.ID).Updates(updates).Error; err != nil {
			w.fail(it.id, err)
			return
		}
	}
	// ثبت خروجی در job
	b, _ := json.Marshal(res)
	_ = w.DB.Exec("UPDATE image_seo_jobs SET status='completed', finished_at=NOW(), output_json=? WHERE id=?", string(b), it.id).Error
}

func (w *ImageSEOWorker) fail(jobID uint, err error) {
	log.Println("image seo job failed:", err)
	_ = w.DB.Exec("UPDATE image_seo_jobs SET status='error', finished_at=NOW(), error_message=? WHERE id=?", err.Error(), jobID).Error
}
