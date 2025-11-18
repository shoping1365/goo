package handlers

import (
	"context"
	"fmt"
	"my-go-backend/internal/cqrs/commands"
	"my-go-backend/internal/database/unitofwork"
	"my-go-backend/internal/models"
	"time"
)

// MediaCommandHandler پردازنده دستورات مربوط به رسانه
// این handler مسئول اجرای تمام دستورات write (ایجاد، به‌روزرسانی، حذف) است
type MediaCommandHandler struct {
	uowFactory unitofwork.UnitOfWorkFactory
}

// NewMediaCommandHandler ایجاد نمونه جدید از command handler
func NewMediaCommandHandler(uowFactory unitofwork.UnitOfWorkFactory) *MediaCommandHandler {
	return &MediaCommandHandler{
		uowFactory: uowFactory,
	}
}

// HandleCreateMediaFile پردازش دستور ایجاد فایل رسانه
func (h *MediaCommandHandler) HandleCreateMediaFile(ctx context.Context, cmd commands.CreateMediaFileCommand) (*models.MediaFile, error) {
	// شروع unit of work
	uow := h.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// ایجاد مدل MediaFile از command
	mediaFile := &models.MediaFile{
		FileName:         cmd.FileName,
		Title:            cmd.Title,
		AltText:          cmd.AltText,
		DisplayTitle:     cmd.DisplayTitle,
		ShortDescription: cmd.ShortDescription,
		Description:      cmd.Description,
		FileType:         cmd.FileType,
		MimeType:         cmd.MimeType,
		Format:           cmd.Format,
		OriginalFormat:   cmd.OriginalFormat,
		Size:             cmd.Size,
		URL:              cmd.URL,
		Category:         cmd.Category,
		UploadedBy:       cmd.UploadedBy,
		DurationSeconds:  cmd.DurationSeconds,
		Width:            cmd.Width,
		Height:           cmd.Height,
		BitrateKbps:      cmd.BitrateKbps,
	}

	// ذخیره در دیتابیس
	mediaRepo := uow.MediaFileRepository()
	if err := mediaRepo.Create(ctx, mediaFile); err != nil {
		return nil, fmt.Errorf("خطا در ایجاد فایل رسانه: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return nil, fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true
	return mediaFile, nil
}

// HandleUpdateMediaFile پردازش دستور به‌روزرسانی فایل رسانه
func (h *MediaCommandHandler) HandleUpdateMediaFile(ctx context.Context, cmd commands.UpdateMediaFileCommand) error {
	// شروع unit of work
	uow := h.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// دریافت فایل موجود
	mediaRepo := uow.MediaFileRepository()
	mediaFile, err := mediaRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("فایل رسانه یافت نشد: %w", err)
	}

	// به‌روزرسانی فیلدها (فقط اگر مقدار داشته باشند)
	if cmd.Title != nil {
		mediaFile.Title = *cmd.Title
	}
	if cmd.AltText != nil {
		mediaFile.AltText = *cmd.AltText
	}
	if cmd.DisplayTitle != nil {
		mediaFile.DisplayTitle = *cmd.DisplayTitle
	}
	if cmd.ShortDescription != nil {
		mediaFile.ShortDescription = *cmd.ShortDescription
	}
	if cmd.Description != nil {
		mediaFile.Description = *cmd.Description
	}
	if cmd.Category != nil {
		mediaFile.Category = *cmd.Category
	}

	// ذخیره تغییرات
	if err := mediaRepo.Update(ctx, mediaFile); err != nil {
		return fmt.Errorf("خطا در به‌روزرسانی فایل رسانه: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true
	return nil
}

// HandleDeleteMediaFile پردازش دستور حذف فایل رسانه
func (h *MediaCommandHandler) HandleDeleteMediaFile(ctx context.Context, cmd commands.DeleteMediaFileCommand) error {
	// شروع unit of work
	uow := h.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// حذف نسخه‌های مرتبط
	versionRepo := uow.MediaVersionRepository()
	versions, err := versionRepo.GetByMediaID(ctx, cmd.ID)
	if err == nil && len(versions) > 0 {
		for _, v := range versions {
			if err := versionRepo.Delete(ctx, v.ID); err != nil {
				return fmt.Errorf("خطا در حذف نسخه %d: %w", v.ID, err)
			}
		}
	}

	// حذف job های فشرده‌سازی مرتبط
	jobRepo := uow.CompressionJobRepository()
	jobs, err := jobRepo.GetByMediaID(ctx, cmd.ID)
	if err == nil && len(jobs) > 0 {
		for _, j := range jobs {
			if err := jobRepo.Delete(ctx, j.ID); err != nil {
				return fmt.Errorf("خطا در حذف job %d: %w", j.ID, err)
			}
		}
	}

	// حذف فایل اصلی
	mediaRepo := uow.MediaFileRepository()
	if err := mediaRepo.Delete(ctx, cmd.ID); err != nil {
		return fmt.Errorf("خطا در حذف فایل رسانه: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true
	return nil
}

// HandleCreateMediaWithVersion پردازش دستور ایجاد فایل رسانه به همراه نسخه
func (h *MediaCommandHandler) HandleCreateMediaWithVersion(ctx context.Context, cmd commands.CreateMediaWithVersionCommand) (*models.MediaFile, *models.MediaVersion, error) {
	// شروع unit of work
	uow := h.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, nil, fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// ایجاد فایل رسانه
	mediaFile := &models.MediaFile{
		FileName:         cmd.MediaFile.FileName,
		Title:            cmd.MediaFile.Title,
		AltText:          cmd.MediaFile.AltText,
		DisplayTitle:     cmd.MediaFile.DisplayTitle,
		ShortDescription: cmd.MediaFile.ShortDescription,
		Description:      cmd.MediaFile.Description,
		FileType:         cmd.MediaFile.FileType,
		MimeType:         cmd.MediaFile.MimeType,
		Format:           cmd.MediaFile.Format,
		OriginalFormat:   cmd.MediaFile.OriginalFormat,
		Size:             cmd.MediaFile.Size,
		URL:              cmd.MediaFile.URL,
		Category:         cmd.MediaFile.Category,
		UploadedBy:       cmd.MediaFile.UploadedBy,
		DurationSeconds:  cmd.MediaFile.DurationSeconds,
		Width:            cmd.MediaFile.Width,
		Height:           cmd.MediaFile.Height,
		BitrateKbps:      cmd.MediaFile.BitrateKbps,
	}

	mediaRepo := uow.MediaFileRepository()
	if err := mediaRepo.Create(ctx, mediaFile); err != nil {
		return nil, nil, fmt.Errorf("خطا در ایجاد فایل رسانه: %w", err)
	}

	// ایجاد نسخه
	version := &models.MediaVersion{
		MediaID:           mediaFile.ID,
		VersionCode:       cmd.Version.VersionCode,
		FilePath:          cmd.Version.FilePath,
		FileSize:          cmd.Version.FileSize,
		Format:            cmd.Version.Format,
		Quality:           cmd.Version.Quality,
		IsActive:          cmd.Version.IsActive,
		IsBackup:          cmd.Version.IsBackup,
		CompressionRatio:  cmd.Version.CompressionRatio,
		CompressionMethod: cmd.Version.CompressionMethod,
		CreatedBy:         cmd.Version.CreatedBy,
		Meta:              cmd.Version.Meta,
		ParentVersionID:   cmd.Version.ParentVersionID,
		BackupPath:        cmd.Version.BackupPath,
		Note:              cmd.Version.Note,
	}

	versionRepo := uow.MediaVersionRepository()
	if err := versionRepo.Create(ctx, version); err != nil {
		return nil, nil, fmt.Errorf("خطا در ایجاد نسخه: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return nil, nil, fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true
	return mediaFile, version, nil
}

// HandleCompressMediaFile پردازش دستور فشرده‌سازی فایل رسانه
func (h *MediaCommandHandler) HandleCompressMediaFile(ctx context.Context, cmd commands.CompressMediaFileCommand) (*models.CompressionJob, error) {
	// شروع unit of work
	uow := h.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// بررسی وجود فایل
	mediaRepo := uow.MediaFileRepository()
	file, err := mediaRepo.GetByID(ctx, cmd.MediaID)
	if err != nil {
		return nil, fmt.Errorf("فایل رسانه یافت نشد: %w", err)
	}

	// ایجاد job فشرده‌سازی
	job := &models.CompressionJob{
		MediaID:       cmd.MediaID,
		RequestedBy:   cmd.RequestedBy,
		Status:        "pending",
		JobType:       "compress",
		TargetFormat:  cmd.TargetFormat,
		TargetQuality: cmd.TargetQuality,
		OriginalSize:  &file.Size,
		CreatedAt:     time.Now(),
	}

	jobRepo := uow.CompressionJobRepository()
	if err := jobRepo.Create(ctx, job); err != nil {
		return nil, fmt.Errorf("خطا در ایجاد job فشرده‌سازی: %w", err)
	}

	// ایجاد تنظیمات فشرده‌سازی
	settingValue := fmt.Sprintf(`{"format": "%s", "quality": "%s"}`, cmd.TargetFormat, cmd.TargetQuality)
	setting := &models.CompressionSetting{
		Scope: "job",
		JobID: &job.ID,
		Key:   "compression_config",
		Value: settingValue,
	}

	settingRepo := uow.CompressionSettingRepository()
	if err := settingRepo.Create(ctx, setting); err != nil {
		return nil, fmt.Errorf("خطا در ذخیره تنظیمات: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return nil, fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true
	return job, nil
}

// HandleSetActiveVersion پردازش دستور تنظیم نسخه فعال
func (h *MediaCommandHandler) HandleSetActiveVersion(ctx context.Context, cmd commands.SetActiveVersionCommand) error {
	// شروع unit of work
	uow := h.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// تنظیم نسخه فعال
	versionRepo := uow.MediaVersionRepository()
	if err := versionRepo.SetActiveVersion(ctx, cmd.MediaID, cmd.VersionID); err != nil {
		return fmt.Errorf("خطا در تنظیم نسخه فعال: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true
	return nil
}

// HandleCreateBackup پردازش دستور ایجاد بکاپ
func (h *MediaCommandHandler) HandleCreateBackup(ctx context.Context, cmd commands.CreateBackupCommand) (*models.MediaVersion, error) {
	// شروع unit of work
	uow := h.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return nil, fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// ایجاد بکاپ
	versionRepo := uow.MediaVersionRepository()
	backup, err := versionRepo.CreateBackupFromVersion(ctx, cmd.VersionID)
	if err != nil {
		return nil, fmt.Errorf("خطا در ایجاد بکاپ: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return nil, fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true
	return backup, nil
}

// HandleCleanupOldBackups پردازش دستور پاکسازی بکاپ‌های قدیمی
func (h *MediaCommandHandler) HandleCleanupOldBackups(ctx context.Context, cmd commands.CleanupOldBackupsCommand) error {
	// شروع unit of work
	uow := h.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// پاکسازی بکاپ‌های قدیمی
	versionRepo := uow.MediaVersionRepository()
	if err := versionRepo.DeleteOldBackups(ctx, cmd.MediaID, cmd.KeepCount); err != nil {
		return fmt.Errorf("خطا در پاکسازی بکاپ‌ها: %w", err)
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true
	return nil
}

// HandleBulkCreateMediaFiles پردازش دستور ایجاد دسته‌ای فایل‌های رسانه
func (h *MediaCommandHandler) HandleBulkCreateMediaFiles(ctx context.Context, cmd commands.BulkCreateMediaFilesCommand) error {
	// شروع unit of work
	uow := h.uowFactory.Create()
	if err := uow.BeginTx(ctx); err != nil {
		return fmt.Errorf("خطا در شروع تراکنش: %w", err)
	}

	// تضمین rollback در صورت خطا
	committed := false
	defer func() {
		if !committed {
			_ = uow.Rollback()
		}
	}()

	// تبدیل commands به models
	files := make([]models.MediaFile, len(cmd.Files))
	for i, f := range cmd.Files {
		files[i] = models.MediaFile{
			FileName:         f.FileName,
			Title:            f.Title,
			AltText:          f.AltText,
			DisplayTitle:     f.DisplayTitle,
			ShortDescription: f.ShortDescription,
			Description:      f.Description,
			FileType:         f.FileType,
			MimeType:         f.MimeType,
			Format:           f.Format,
			OriginalFormat:   f.OriginalFormat,
			Size:             f.Size,
			URL:              f.URL,
			Category:         f.Category,
			UploadedBy:       f.UploadedBy,
			DurationSeconds:  f.DurationSeconds,
			Width:            f.Width,
			Height:           f.Height,
			BitrateKbps:      f.BitrateKbps,
		}
	}

	// ایجاد دسته‌ای
	mediaRepo := uow.MediaFileRepository()
	if bulkRepo, ok := mediaRepo.(interface {
		BulkCreate(context.Context, []models.MediaFile) error
	}); ok {
		if err := bulkRepo.BulkCreate(ctx, files); err != nil {
			return fmt.Errorf("خطا در ایجاد دسته‌ای: %w", err)
		}
	} else {
		// fallback به ایجاد تکی
		for _, file := range files {
			if err := mediaRepo.Create(ctx, &file); err != nil {
				return fmt.Errorf("خطا در ایجاد فایل: %w", err)
			}
		}
	}

	// تأیید تراکنش
	if err := uow.Commit(); err != nil {
		return fmt.Errorf("خطا در تأیید تراکنش: %w", err)
	}

	committed = true
	return nil
}
