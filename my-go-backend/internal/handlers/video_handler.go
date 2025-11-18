package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"gorm.io/gorm"
)

// VideoHandler is a thin wrapper around existing MediaHandler methods for video operations.
// It allows future separation without breaking current API contracts.
type VideoHandler struct {
	Base *MediaHandler
}

func NewVideoHandler(base *MediaHandler) *VideoHandler {
	return &VideoHandler{Base: base}
}

// ------- passthrough wrappers for common operations --------
func (h *VideoHandler) ListVideosHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.ListVideosHandler(w, r)
}
func (h *VideoHandler) UploadVideoHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.UploadVideoHandler(w, r)
}
func (h *VideoHandler) DownloadVideoHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.DownloadVideoHandler(w, r)
}
func (h *VideoHandler) DeleteVideoHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.DeleteVideoHandler(w, r)
}
func (h *VideoHandler) GetVideoHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.GetVideoHandler(w, r)
}
func (h *VideoHandler) UpdateVideoHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.UpdateVideoHandler(w, r)
}

// -----------------------------------------------------------------------------
// Video-specific operations
// -----------------------------------------------------------------------------

// CompressVideoHandler handles POST /media/compress-video/:id
// فشرده‌سازی ویدیو با قابلیت هوشمند
func (h *VideoHandler) CompressVideoHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "شناسه نامعتبر است", nil))
		return
	}

	// بررسی وجود ویدیو در دیتابیس
	var media models.MediaFile
	if err := h.Base.DB.First(&media, id).Error; err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("VIDEO_NOT_FOUND", "ویدیو یافت نشد", nil))
		return
	}

	// بررسی اینکه فایل واقعاً ویدیو است
	if !strings.HasPrefix(media.FileType, "video/") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("INVALID_FILE_TYPE", "فقط ویدیوها قابل فشرده‌سازی هستند", nil))
		return
	}

	// ساخت مسیر کامل فایل
	projectRoot := ""
	if wd, err := os.Getwd(); err == nil {
		projectRoot = filepath.Dir(wd)
	}
	publicDir := filepath.Join(projectRoot, "public")
	rel := strings.TrimPrefix(media.URL, "/")
	absPath := filepath.Join(publicDir, rel)

	// بررسی وجود فایل
	if _, err := os.Stat(absPath); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("FILE_NOT_FOUND", "فایل یافت نشد", nil))
		return
	}

	// ایجاد job برای فشرده‌سازی
	jobID := generateUniqueName("")
	ctx, cancel := context.WithCancel(context.Background())
	job := &videoJob{Progress: 0, cancel: cancel}
	videoCompressJobs.Store(jobID, job)

	// راه‌اندازی worker ها
	startVideoWorkers()

	// دریافت پارامترهای فشرده‌سازی
	q := r.URL.Query()
	opts := services.SmartVideoOptions{}

	// بررسی پارامتر smart برای فشرده‌سازی هوشمند
	isSmartCompression := q.Get("smart") == "true"
	if isSmartCompression {
		// تنظیمات پیش‌فرض برای فشرده‌سازی هوشمند
		opts.UseFaceDetection = true
		opts.TargetRatio = 0.6 // کاهش حجم تا 60%
		opts.MaxWidth = 1920
		opts.MaxHeight = 1080
		opts.TargetBitrateKbps = 3000
	} else {
		// تنظیمات معمولی
		if v := q.Get("target_ratio"); v != "" {
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				opts.TargetRatio = f
			}
		}
		if v := q.Get("max_width"); v != "" {
			if i, err := strconv.Atoi(v); err == nil {
				opts.MaxWidth = i
			}
		}
		if v := q.Get("max_height"); v != "" {
			if i, err := strconv.Atoi(v); err == nil {
				opts.MaxHeight = i
			}
		}
		if v := q.Get("width"); v != "" {
			if i, err := strconv.Atoi(v); err == nil {
				opts.Width = i
			}
		}
		if v := q.Get("height"); v != "" {
			if i, err := strconv.Atoi(v); err == nil {
				opts.Height = i
			}
		}
		if v := q.Get("bitrate"); v != "" {
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				opts.TargetBitrateKbps = i
			}
		}
	}

	// دریافت روش تحلیل فریم برای ویدیوهای طولانی
	frameAnalysisMode := q.Get("frame_analysis_mode")
	if frameAnalysisMode == "" {
		frameAnalysisMode = "smart" // مقدار پیش‌فرض
	}

	// دریافت تعداد ورکر از query parameter
	workerCount := 4 // مقدار پیش‌فرض
	if v := q.Get("worker_count"); v != "" {
		if i, err := strconv.Atoi(v); err == nil && i >= 1 && i <= 16 {
			workerCount = i
		}
	}

	fmt.Printf("Video compression requested with workerCount: %d, frameAnalysisMode: %s\n", workerCount, frameAnalysisMode)

	// ارسال job به صف
	videoCompressQueue <- &queuedJob{
		jobID:             jobID,
		absPath:           absPath,
		mediaID:           uint(id),
		db:                h.Base.DB,
		opts:              opts,
		frameAnalysisMode: frameAnalysisMode,
		workerCount:       workerCount,
		ctx:               ctx,
		job:               job,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"job_id":  jobID,
		"message": "فشرده‌سازی ویدیو شروع شد",
	})
}

// CompressVideoProgressHandler handles GET /media/compress-video/progress/:jobID
// بررسی پیشرفت فشرده‌سازی ویدیو
func (h *VideoHandler) CompressVideoProgressHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	jobID := parts[len(parts)-1]

	jobInterface, exists := videoCompressJobs.Load(jobID)
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("JOB_NOT_FOUND", "کار فشرده‌سازی یافت نشد", nil))
		return
	}

	job := jobInterface.(*videoJob)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"progress": job.Progress,
		"done":     job.Done,
		"paused":   job.Paused,
		"error":    job.Error,
		"result":   job.Result,
	})
}

// CancelVideoCompressionHandler handles POST /media/compress-video/cancel/:jobID
// لغو فشرده‌سازی ویدیو
func (h *VideoHandler) CancelVideoCompressionHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	jobID := parts[len(parts)-1]

	jobInterface, exists := videoCompressJobs.Load(jobID)
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("JOB_NOT_FOUND", "کار فشرده‌سازی یافت نشد", nil))
		return
	}

	job := jobInterface.(*videoJob)
	if job.Done {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("JOB_ALREADY_DONE", "کار فشرده‌سازی قبلاً تمام شده است", nil))
		return
	}

	// لغو کار
	job.cancel()
	if job.proc != nil {
		_ = job.proc.Signal(sigStop)
	}

	job.Done = true
	job.Error = "کار توسط کاربر لغو شد"
	videoCompressJobs.Store(jobID, job)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "کار فشرده‌سازی لغو شد",
	})
}

// PauseVideoCompressionHandler handles POST /media/compress-video/pause/:jobID
// توقف موقت فشرده‌سازی ویدیو
func (h *VideoHandler) PauseVideoCompressionHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	jobID := parts[len(parts)-1]

	jobInterface, exists := videoCompressJobs.Load(jobID)
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("JOB_NOT_FOUND", "کار فشرده‌سازی یافت نشد", nil))
		return
	}

	job := jobInterface.(*videoJob)
	if job.Done {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("JOB_ALREADY_DONE", "کار فشرده‌سازی قبلاً تمام شده است", nil))
		return
	}

	if job.Paused {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("JOB_ALREADY_PAUSED", "کار فشرده‌سازی قبلاً متوقف شده است", nil))
		return
	}

	// توقف کار
	job.Paused = true
	if job.proc != nil {
		_ = job.proc.Signal(sigStop)
	}

	videoCompressJobs.Store(jobID, job)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "کار فشرده‌سازی متوقف شد",
	})
}

// ResumeVideoCompressionHandler handles POST /media/compress-video/resume/:jobID
// ادامه فشرده‌سازی ویدیو
func (h *VideoHandler) ResumeVideoCompressionHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	jobID := parts[len(parts)-1]

	jobInterface, exists := videoCompressJobs.Load(jobID)
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("JOB_NOT_FOUND", "کار فشرده‌سازی یافت نشد", nil))
		return
	}

	job := jobInterface.(*videoJob)
	if job.Done {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("JOB_ALREADY_DONE", "کار فشرده‌سازی قبلاً تمام شده است", nil))
		return
	}

	if !job.Paused {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("JOB_NOT_PAUSED", "کار فشرده‌سازی متوقف نشده است", nil))
		return
	}

	// ادامه کار
	job.Paused = false
	if job.proc != nil {
		_ = job.proc.Signal(sigCont)
	}

	videoCompressJobs.Store(jobID, job)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "کار فشرده‌سازی ادامه یافت",
	})
}

// PreviewCompressVideoHandler پیش‌نمایش زنده فشرده‌سازی ویدیو
// ورودی: فایل ویدیو (multipart)، پارامترهای فشرده‌سازی
// خروجی: ویدیوی فشرده‌شده با Content-Type مناسب
func (h *VideoHandler) PreviewCompressVideoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PATH:", os.Getenv("PATH"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// محدودیت حجم فایل (مثلاً 100MB)
	r.Body = http.MaxBytesReader(w, r.Body, 100*1024*1024)
	if err := r.ParseMultipartForm(100 * 1024 * 1024); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Invalid form data"})
		return
	}

	// دریافت فایل ویدیو
	file, header, err := r.FormFile("video")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "No video file uploaded"})
		return
	}
	defer file.Close()

	// بررسی نوع فایل
	if !strings.HasPrefix(header.Header.Get("Content-Type"), "video/") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "File is not a video"})
		return
	}

	// ایجاد فایل موقت
	tempDir := os.TempDir()
	tempInput := filepath.Join(tempDir, "preview_input_"+generateUniqueName("")+filepath.Ext(header.Filename))
	tempOutput := filepath.Join(tempDir, "preview_output_"+generateUniqueName("")+".mp4")

	// کپی فایل ورودی به فایل موقت
	inputFile, err := os.Create(tempInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Failed to create temp file"})
		return
	}
	defer func() {
		inputFile.Close()
		os.Remove(tempInput)
		os.Remove(tempOutput)
	}()

	_, err = io.Copy(inputFile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Failed to save uploaded file"})
		return
	}
	inputFile.Close()

	// دریافت پارامترهای فشرده‌سازی
	q := r.URL.Query()
	opts := services.SmartVideoOptions{}

	// دریافت تعداد ورکر از تنظیمات مرکزی (برای استفاده در نسخه‌های بعدی)
	_ = q.Get("worker_count") // فعلاً استفاده نمی‌شود

	// بررسی پارامتر smart برای فشرده‌سازی هوشمند
	isSmartCompression := q.Get("smart") == "true"
	if isSmartCompression {
		// تنظیمات پیش‌فرض برای فشرده‌سازی هوشمند
		// فعلاً OpenCV را غیرفعال می‌کنیم تا مشکل برطرف شود
		opts.UseFaceDetection = false // موقتاً غیرفعال
		opts.TargetRatio = 0.6
		opts.MaxWidth = 1920
		opts.MaxHeight = 1080
		opts.TargetBitrateKbps = 3000
	} else {
		// تنظیمات معمولی
		if v := q.Get("quality"); v != "" {
			switch v {
			case "1080p":
				opts.MaxHeight = 1080
				opts.TargetBitrateKbps = 5000
			case "720p":
				opts.MaxHeight = 720
				opts.TargetBitrateKbps = 3000
			case "480p":
				opts.MaxHeight = 480
				opts.TargetBitrateKbps = 1500
			case "360p":
				opts.MaxHeight = 360
				opts.TargetBitrateKbps = 800
			}
		}
		if v := q.Get("bitrate"); v != "" {
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				opts.TargetBitrateKbps = i
			}
		}
		if v := q.Get("format"); v != "" {
			// برای پیش‌نمایش همیشه MP4 استفاده می‌کنیم
			// opts.OutputFormat = "mp4" // این فیلد در SmartVideoOptions وجود ندارد
		}
	}

	// فشرده‌سازی ویدیو
	// برای حال حاضر از SmartCompressVideo معمولی استفاده می‌کنیم
	// در نسخه‌های بعدی، SmartCompressVideoWithSettings را اضافه خواهیم کرد
	result, err := services.SmartCompressVideo(context.Background(), tempInput, opts, func(p int) {
		// progress callback - برای پیش‌نمایش نیازی نیست
	}, func(proc *os.Process) {
		// process callback - برای پیش‌نمایش نیازی نیست
	})
	if err != nil {
		fmt.Println("خطا در فشرده‌سازی پیش‌نمایش ویدیو:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Compression failed: " + err.Error()})
		return
	}

	// استفاده از مسیر خروجی از نتیجه
	tempOutput = result.OutputPath

	// خواندن فایل خروجی و ارسال
	outputFile, err := os.Open(tempOutput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Failed to read output file"})
		return
	}
	defer outputFile.Close()

	// تنظیم headers
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Disposition", "inline; filename=compressed_video.mp4")
	w.Header().Set("Cache-Control", "no-cache")

	// ارسال فایل
	io.Copy(w, outputFile)
}

// RevertVideoCompressionHandler handles POST /media/revert-video/:id
// بازگردانی ویدیو به نسخه اصلی
func (h *VideoHandler) RevertVideoCompressionHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("VALIDATION_ERROR", "شناسه نامعتبر است", nil))
		return
	}

	// بررسی وجود ویدیو در دیتابیس
	var media models.MediaFile
	if err := h.Base.DB.First(&media, id).Error; err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("VIDEO_NOT_FOUND", "ویدیو یافت نشد", nil))
		return
	}

	// بررسی اینکه فایل واقعاً ویدیو است
	if !strings.HasPrefix(media.FileType, "video/") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.New("INVALID_FILE_TYPE", "فقط ویدیوها قابل بازگردانی هستند", nil))
		return
	}

	// استفاده از BackupService برای بازگردانی
	bs := services.NewBackupService("")
	projectRoot := bs.ProjectRoot
	publicDir := filepath.Join(projectRoot, "public")

	rel := strings.TrimPrefix(media.URL, "/")
	absPath := filepath.Join(publicDir, rel)

	// جستجوی بک‌آپ
	backupPath, err := bs.FindLatestBackup(rel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(utils.New("BACKUP_NOT_FOUND", "نسخه پشتیبان یافت نشد", nil))
		return
	}

	// کپی بک‌آپ به محل اصلی
	if err := os.Remove(absPath); err != nil && !os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("FILE_DELETE_ERROR", "خطا در حذف فایل فعلی", err.Error()))
		return
	}

	if err := copyFile(backupPath, absPath); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("BACKUP_RESTORE_ERROR", "خطا در بازگردانی نسخه پشتیبان", err.Error()))
		return
	}

	// بروزرسانی دیتابیس
	fileInfo, _ := os.Stat(absPath)
	size := uint(fileInfo.Size())

	h.Base.DB.Model(&models.MediaFile{}).Where("id = ?", id).Updates(map[string]interface{}{
		"compressed":      false,
		"compressed_size": nil,
		"size":            size,
	})

	// دریافت رکورد بروزرسانی شده
	var updated models.MediaFile
	h.Base.DB.First(&updated, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "ویدیو با موفقیت به نسخه اصلی بازگردانده شد",
		"url":     updated.URL,
		"size":    updated.Size,
		"format":  updated.Format,
	})
}

// -----------------------------------------------------------------------------
// Queue & Worker infrastructure (limited concurrency)
// -----------------------------------------------------------------------------
var (
	maxConcurrentVideo = 2
	videoQueueSize     = 100
	videoCompressQueue = make(chan *queuedJob, videoQueueSize)
	startWorkersOnce   sync.Once

	sigStop syscall.Signal
	sigCont syscall.Signal
)

func init() {
	// Numerical values of SIGSTOP (19) & SIGCONT (18) on Unix-like OS.
	if runtime.GOOS != "windows" {
		sigStop = syscall.Signal(19)
		sigCont = syscall.Signal(18)
	}
}

type videoJob struct {
	Progress int                        `json:"progress"`
	Result   *services.SmartVideoResult `json:"result,omitempty"`
	Error    string                     `json:"error,omitempty"`
	Done     bool                       `json:"done"`
	Paused   bool                       `json:"paused"`
	cancel   context.CancelFunc         `json:"-"`
	proc     *os.Process                `json:"-"`
}

type queuedJob struct {
	jobID             string
	absPath           string
	mediaID           uint
	db                *gorm.DB
	opts              services.SmartVideoOptions
	frameAnalysisMode string // روش تحلیل فریم برای ویدیوهای طولانی
	workerCount       int    // تعداد ورکر برای پردازش همزمان
	ctx               context.Context
	job               *videoJob
}

func startVideoWorkers() {
	startWorkersOnce.Do(func() {
		for i := 0; i < maxConcurrentVideo; i++ {
			go videoWorker()
		}
	})
}

func videoWorker() {
	for qj := range videoCompressQueue {
		// تنظیم تعداد ورکر در تنظیمات
		qj.opts.WorkerCount = qj.workerCount
		qj.opts.FrameAnalysisMode = qj.frameAnalysisMode

		fmt.Printf("Processing video with workerCount: %d, frameAnalysisMode: %s\n", qj.workerCount, qj.frameAnalysisMode)

		result, err := services.SmartCompressVideo(qj.ctx, qj.absPath, qj.opts, func(p int) {
			qj.job.Progress = p
		}, func(proc *os.Process) {
			qj.job.proc = proc
		})
		if err != nil {
			qj.job.Error = err.Error()
			qj.job.Done = true
			videoCompressJobs.Store(qj.jobID, qj.job)
			continue
		}
		qj.job.Progress = 100
		qj.job.Result = &result
		qj.job.Done = true

		// بروزرسانی دیتابیس
		if qj.mediaID > 0 && qj.db != nil {
			compressedSize := uint(result.SizeBytes)
			ext := strings.TrimPrefix(filepath.Ext(result.OutputPath), ".")
			dbUpdates := map[string]interface{}{
				"compressed":       true,
				"compressed_size":  compressedSize,
				"format":           ext,
				"width":            result.Width,
				"height":           result.Height,
				"bitrate_kbps":     result.BitrateKbps,
				"duration_seconds": result.DurationSeconds,
			}
			if err := qj.db.Table("media_files").Where("id = ?", qj.mediaID).Updates(dbUpdates).Error; err != nil {
				fmt.Println("update media file error:", err)
			}
		}

		videoCompressJobs.Store(qj.jobID, qj.job)
	}
}

// Mapping of jobID -> *videoJob for progress polling
var videoCompressJobs sync.Map

// -----------------------------------------------------------------------------
// Utility functions
// -----------------------------------------------------------------------------

// copyFile کپی فایل از مبدا به مقصد
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
