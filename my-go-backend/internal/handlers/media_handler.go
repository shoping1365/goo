package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"image"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"

	"my-go-backend/internal/models"
	"my-go-backend/internal/services"
	"my-go-backend/internal/utils"

	"gorm.io/gorm"
)

// Allowed image MIME types
var allowedImageTypes = map[string]bool{
	"image/jpeg":    true,
	"image/png":     true,
	"image/gif":     true,
	"image/webp":    true,
	"image/svg+xml": true,
}

// Allowed video MIME types (primary detection)
var allowedVideoTypes = map[string]bool{
	"video/mp4":                true,
	"video/webm":               true,
	"video/quicktime":          true,
	"video/x-msvideo":          true, // AVI
	"video/x-matroska":         true, // MKV
	"video/x-flv":              true, // FLV
	"video/3gpp":               true, // 3GP
	"video/x-ms-wmv":           true, // WMV
	"application/octet-stream": true, // برخی مرورگرها برای ویدیو mp4 ارسال می‌کنند
}

// Fallback allowed video extensions (برای مواقعی که DetectContentType تشخیص اشتباه دهد)
var allowedVideoExt = map[string]bool{
	".mp4":  true,
	".webm": true,
	".mov":  true,
	".mkv":  true,
	".avi":  true,
	".flv":  true,
	".3gp":  true,
	".wmv":  true,
	".m4v":  true,
	".ogv":  true,
}

// Allowed audio MIME types
var allowedAudioTypes = map[string]bool{
	"audio/mpeg": true,
	"audio/wav":  true,
	"audio/ogg":  true,
}

const maxFileSize = 1024 * 1024 * 1024 // 1GB

// ------------------- Smart Video Queue (limited concurrency) -------------------

// پیش‌فرض‌های سایز تصاویر؛ مقادیر می‌توانند در تنظیمات (کلید image_sizes) override شوند.
var imageSizesDefault = map[string]struct{ width, height uint }{
	"thumbnail": {150, 150},
	"small":     {400, 400},
	"medium":    {800, 800},
	"large":     {1600, 1600},
}

// MediaHandler ساختار هندلر برای مدیریت رسانه‌ها
// باید در main مقداردهی شود و *gorm.DB را دریافت کند

type MediaHandler struct {
	DB             *gorm.DB
	SettingService *services.SettingService
	// سرویس اختیاری برای زمان‌بندی AI تصاویر
	ImageSEOWorker *services.ImageSEOWorker
}

func NewMediaHandler(db *gorm.DB, ss *services.SettingService) *MediaHandler {
	return &MediaHandler{DB: db, SettingService: ss}
}

// MediaWithUploader ساختار خروجی رسانه به همراه اطلاعات آپلودکننده
// این struct برای نمایش اطلاعات رسانه و اطلاعات کاربر آپلودکننده در خروجی API استفاده می‌شود.
type MediaWithUploader struct {
	models.MediaFile
	UploaderName     string `json:"uploader_name"`
	UploaderUsername string `json:"uploader_username"`
	UploaderRole     string `json:"uploader_role"`
	Caption          string `json:"caption"`
}

// ListMediaHandler لیست رسانه‌ها را به همراه اطلاعات آپلودکننده برمی‌گرداند
func (h *MediaHandler) ListMediaHandler(w http.ResponseWriter, r *http.Request) {
	var mediaWithUploader []MediaWithUploader

	// دریافت پارامتر category از query string
	category := r.URL.Query().Get("category")

	// ساخت کوئری پایه
	query := h.DB.Table("media_files").
		Select(`media_files.*, 
			media_files.short_description as caption, 
			users.name as uploader_name, 
			users.username as uploader_username, 
			roles.name as uploader_role`).
		Joins("LEFT JOIN users ON users.id = media_files.uploaded_by").
		Joins("LEFT JOIN roles ON roles.id = users.role_id").
		Where("media_files.deleted_at IS NULL")

	// اعمال فیلتر category در صورت وجود
	if category != "" && category != "all" {
		query = query.Where("media_files.category = ?", category)
	}

	// اجرای کوئری
	err := query.Order("media_files.created_at DESC").
		Scan(&mediaWithUploader).Error

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطا در دریافت لیست رسانه‌ها", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": mediaWithUploader})
}

// آپلود رسانه (media_files + media_variants)
func (h *MediaHandler) UploadMediaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	responseSent := false
	sendJSON := func(status int, payload interface{}) {
		if !responseSent {
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(payload)
			responseSent = true
		}
	}
	defer func() {
		if rec := recover(); rec != nil && !responseSent {
			sendJSON(http.StatusInternalServerError, utils.New("INTERNAL_ERROR", "خطای داخلی سرور", nil))
		}
	}()

	var err error
	// محدودیت حجم فایل
	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)
	if err = r.ParseMultipartForm(maxFileSize); err != nil {
		sendJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "داده‌های فرم نامعتبر است", err.Error()))
		return
	}

	category := strings.ToLower(r.FormValue("category"))
	if category == "" {
		category = strings.ToLower(r.URL.Query().Get("category"))
	}
	if category == "" {
		category = "library"
	}
	if category != "products" && category != "library" && category != "customer" && category != "product-categories" && category != "brands" && category != "banners" {
		sendJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "دسته‌بندی نامعتبر. مجاز: products, library, customer, product-categories, brands, banners", nil))
		return
	}

	purpose := r.URL.Query().Get("purpose")
	if purpose == "" {
		purpose = "general"
	}
	if purpose != "general" && purpose != "banner-desktop" && purpose != "banner-mobile" {
		sendJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "هدف نامعتبر. مجاز: general, banner-desktop, banner-mobile", nil))
		return
	}

	author := r.FormValue("author")
	if author == "" {
		author = r.URL.Query().Get("author")
	}

	// مقداردهی فیلد UploadedBy با آیدی کاربر لاگین‌شده (در صورت وجود)
	var uploadedBy *uint = nil
	// تلاش برای گرفتن user_id از context (مثلاً ست شده توسط middleware احراز هویت)
	if userIDVal, ok := r.Context().Value("user_id").(uint); ok {
		uploadedBy = &userIDVal
	} else if userIDVal, ok := r.Context().Value("user_id").(int64); ok {
		uid := uint(userIDVal)
		uploadedBy = &uid
	} else {
		// اگر در context نبود، از فرم یا query بگیر
		userIDStr := r.FormValue("uploaded_by")
		if userIDStr == "" {
			userIDStr = r.URL.Query().Get("uploaded_by")
		}
		if userIDStr != "" {
			if id, err := parseUintFromString(userIDStr); err == nil {
				uploadedBy = &id
			}
		}
	}

	var file multipart.File
	var header *multipart.FileHeader
	file, header, err = r.FormFile("file")
	if err != nil {
		sendJSON(http.StatusBadRequest, utils.New("FILE_REQUIRED", "هیچ فایلی آپلود نشده است", nil))
		return
	}
	defer file.Close()

	buff := make([]byte, 512)
	if _, err := file.Read(buff); err != nil {
		sendJSON(http.StatusBadRequest, utils.New("FILE_READ_ERROR", "خطا در خواندن فایل", err.Error()))
		return
	}
	filetype := http.DetectContentType(buff)
	ext := strings.ToLower(filepath.Ext(header.Filename))

	// Security: Block dangerous extensions explicitly
	if ext == ".exe" || ext == ".bat" || ext == ".sh" || ext == ".php" || ext == ".go" || ext == ".js" || ext == ".html" {
		sendJSON(http.StatusBadRequest, utils.New("SECURITY_ERROR", "آپلود این نوع فایل مجاز نیست", nil))
		return
	}

	// Security: If content type is generic octet-stream, strictly validate extension against allowed video types
	if filetype == "application/octet-stream" {
		if !allowedVideoExt[ext] {
			sendJSON(http.StatusBadRequest, utils.New("UNSUPPORTED_FILE", "نوع فایل ناشناخته و غیرمجاز است", nil))
			return
		}
	} else if !allowedImageTypes[filetype] && !allowedVideoTypes[filetype] && !allowedAudioTypes[filetype] {
		sendJSON(http.StatusBadRequest, utils.New("UNSUPPORTED_FILE", "نوع فایل پشتیبانی نمی‌شود", nil))
		return
	}

	// Security: Validate TIFF files to prevent CVE-2023-36808 (crash on crafted TIFF)
	if ext == ".tiff" || ext == ".tif" {
		if header.Size > 10*1024*1024 {
			sendJSON(http.StatusBadRequest, utils.New("FILE_TOO_LARGE", "فایل TIFF بیش از حد بزرگ است (حداکثر 10MB)", nil))
			return
		}
	}

	if header.Size > maxFileSize {
		sendJSON(http.StatusBadRequest, utils.New("FILE_TOO_LARGE", "حجم فایل بیش از حد مجاز است", nil))
		return
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		sendJSON(http.StatusBadRequest, utils.New("FILE_READ_ERROR", "خطا در خواندن فایل", err.Error()))
		return
	}

	// قبل از پردازش عکس، بررسی مغایرت MIME و پسوند
	imageExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	videoExts := map[string]bool{".mp4": true, ".webm": true, ".mov": true}

	// اگر MIME عکس و پسوند ویدیو یا بالعکس، آپلود رد شود
	if allowedImageTypes[filetype] && videoExts[ext] || allowedVideoTypes[filetype] && imageExts[ext] {
		sendJSON(http.StatusBadRequest, utils.New("MIME_MISMATCH", "نوع فایل و پسوند همخوانی ندارند. آپلود مجاز نیست.", nil))
		return
	}
	// اگر MIME و پسوند هر دو ویدیویی یا هر دو عکسی باشند، مسیر صحیح اجرا شود (کد پایین به صورت پیش‌فرض همین کار را می‌کند)

	uniqueName := generateUniqueName(ext)

	baseCategory := category // always lowercase

	var subfolder string
	switch {
	case allowedImageTypes[filetype]:
		subfolder = "images"
	case allowedVideoTypes[filetype]:
		subfolder = "videos"
	case allowedAudioTypes[filetype]:
		subfolder = "audio"
	case strings.HasPrefix(filetype, "application/pdf") ||
		strings.HasPrefix(filetype, "application/msword") ||
		strings.HasPrefix(filetype, "application/vnd.openxmlformats-officedocument") ||
		strings.HasPrefix(filetype, "text/"):
		subfolder = "documents"
	default:
		subfolder = "others"
	}

	// اگر category=products است، از ساختار جدید استفاده کن
	var relativePath string
	var absolutePath string

	// اول از environment variable بخون، اگر نبود از working directory محاسبه کن
	projectRoot := os.Getenv("PROJECT_ROOT")
	if projectRoot == "" {
		if wd, err := os.Getwd(); err == nil {
			projectRoot = filepath.Dir(wd)
		}
	}
	publicDir := filepath.Join(projectRoot, "public")

	if baseCategory == "products" && allowedImageTypes[filetype] {
		// استفاده از ساختار جدید برای عکس‌های محصولات
		prodDir, err := getProductImageDir()
		if err != nil {
			// fallback به ساختار قدیمی
			relativePath = filepath.ToSlash(filepath.Join("/uploads", "media", baseCategory, subfolder, uniqueName))
			absolutePath = filepath.Join(publicDir, "uploads", "media", baseCategory, subfolder, uniqueName)
		} else {
			// استفاده از ساختار جدید
			// محاسبه مسیر نسبی از پوشه public
			relPath := strings.TrimPrefix(prodDir, publicDir)
			relPath = filepath.ToSlash(relPath)
			if !strings.HasPrefix(relPath, "/") {
				relPath = "/" + relPath
			}
			relativePath = filepath.ToSlash(filepath.Join(relPath, uniqueName))
			absolutePath = filepath.Join(prodDir, uniqueName)
		}
	} else {
		// ساختار قدیمی برای سایر دسته‌ها - استفاده از تابع helper برای ساخت ایمن مسیر
		var err error
		absolutePath, err = buildSafeMediaPath(publicDir, baseCategory, subfolder, uniqueName)
		if err != nil {
			sendJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "خطا در ساخت مسیر", err.Error()))
			return
		}
		relativePath = filepath.ToSlash(strings.TrimPrefix(absolutePath, publicDir))
		if !strings.HasPrefix(relativePath, "/") {
			relativePath = "/" + relativePath
		}
	}

	if err := os.MkdirAll(filepath.Dir(absolutePath), 0755); err != nil {
		sendJSON(http.StatusInternalServerError, utils.New("DIRECTORY_ERROR", "خطا در ایجاد پوشه", err.Error()))
		return
	}

	// -----------------------------------------------------------
	//  بررسی فعال بودن فشرده‌سازی خودکار از تنظیمات
	// -----------------------------------------------------------

	// کیفیت پیش‌فرض واریانت‌ها (در JPEG/WebP)؛ در ادامه اگر فشرده‌سازی فعال باشد بروزرسانی می‌شود
	variantQuality := 90

	origSize := uint(header.Size)
	finalSize := origSize

	compressEnabled := true
	if h.SettingService != nil {
		if settings, err := h.SettingService.GetSettingsByCategory(r.Context(), "compression"); err == nil {
			for _, s := range settings {
				if s.Key == "compression.enabled" {
					compressEnabled = strings.ToLower(s.Value) != "false" && s.Value != "0"
				}
			}
		}
	}

	// آپلود با checkbox کلاینت: اگر پارامتر compress یا auto_compress=false ارسال شود، غیرفعال می‌شود
	if v := r.FormValue("compress"); v != "" {
		if strings.ToLower(v) == "false" || v == "0" {
			compressEnabled = false
		}
	} else if v := r.FormValue("autoCompress"); v != "" {
		if strings.ToLower(v) == "false" || v == "0" {
			compressEnabled = false
		}
	}

	if allowedImageTypes[filetype] && compressEnabled && filetype != "image/gif" {
		// ابتدا کل فایل را در حافظه می‌خوانیم تا هم بک‌آپ سالم باشد هم برای فشرده‌سازی استفاده شود
		imgBytes, err := io.ReadAll(file)
		if err != nil {
			sendJSON(http.StatusBadRequest, utils.New("FILE_READ_ERROR", "خطا در خواندن تصویر", err.Error()))
			return
		}
		file.Seek(0, io.SeekStart)

		var origImg image.Image
		extLower := strings.ToLower(filepath.Ext(header.Filename))
		switch extLower {
		case ".webp":
			origImg, err = webp.Decode(bytes.NewReader(imgBytes))
		default:
			origImg, _, err = image.Decode(bytes.NewReader(imgBytes))
		}
		if err != nil {
			log.Println("compress decode error:", err)
			http.Error(w, "unsupported image: "+err.Error(), http.StatusBadRequest)
			return
		}

		// تنظیمات فعلی را از SettingService می‌گیریم
		settingMap := map[string]string{}
		if h.SettingService != nil {
			if settings, err := h.SettingService.GetSettingsByCategory(r.Context(), "compression"); err == nil {
				for _, s := range settings {
					settingMap[s.Key] = s.Value
				}
			}
		}

		qualityKey := settingMap["compression.quality"]
		if qualityKey == "" {
			qualityKey = "medium"
		}
		formatKey := settingMap["compression.format"]
		if formatKey == "" {
			formatKey = "original"
		}

		// Allow explicit quality in query param (?quality=80 or ?smart=1)
		queryQ := r.URL.Query().Get("quality")
		querySmart := r.URL.Query().Get("smart")

		// quality percent & smart handling
		qualityPercent := 75
		isSmart := false

		// ------------------------------------------------------------
		// مرحله ۱) ابتدا ورودی صریح کاربر (QueryString) بررسی می‌شود
		// ------------------------------------------------------------
		if querySmart == "1" || strings.ToLower(querySmart) == "true" {
			isSmart = true
		}

		if queryQ != "" {
			if q, err := strconv.Atoi(queryQ); err == nil && q >= 1 && q <= 100 {
				qualityPercent = q
				// وقتی کاربر کیفیت را مستقیم تعیین می‌کند، حالت هوشمند کنار گذاشته می‌شود
				isSmart = false
			}
		}

		// ------------------------------------------------------------
		// مرحله ۲) اگر هیچ پارامتر صریحی نبود، از تنظیمات سراسری بهره ببریم
		// ------------------------------------------------------------
		if querySmart == "" && queryQ == "" {
			switch strings.ToLower(qualityKey) {
			case "smart":
				isSmart = true
			case "high":
				qualityPercent = 90
			case "low":
				qualityPercent = 60
			case "medium":
				qualityPercent = 75
			default:
				// پشتیبانی از مقدار عددی ۱-۱۰۰ ذخیره‌شده در DB
				if q, err := strconv.Atoi(qualityKey); err == nil && q >= 1 && q <= 100 {
					qualityPercent = q
				}
			}
		}

		outFormat := strings.ToLower(formatKey)
		if outFormat == "original" {
			ext := strings.ToLower(filepath.Ext(header.Filename))
			switch ext {
			case ".jpg", ".jpeg":
				outFormat = "jpeg"
			case ".png":
				outFormat = "png"
			case ".webp":
				outFormat = "webp"
			default:
				outFormat = "jpeg"
			}
		}

		var buf bytes.Buffer

		// --------------------------------------------------
		//  تغییر اندازه برای برندها و دسته‌بندی‌ها بر اساس تنظیمات image_sizes
		// --------------------------------------------------
		if baseCategory == "brands" || baseCategory == "product-categories" && h.SettingService != nil {
			if cfg, err := h.SettingService.GetSetting(r.Context(), "image_sizes"); err == nil && cfg != nil {
				type sizeCfg struct {
					Width  uint `json:"width"`
					Height uint `json:"height"`
				}
				var mp map[string]sizeCfg
				if json.Unmarshal([]byte(cfg.Value), &mp) == nil {
					key := ""
					if baseCategory == "brands" {
						key = "brand"
					} else {
						key = "category"
					}
					if val, ok := mp[key]; ok {
						if val.Width > 0 || val.Height > 0 {
							origImg = resize.Resize(val.Width, val.Height, origImg, resize.Lanczos3)
						}
					} else if baseCategory == "product-categories" {
						// پیش‌فرض 500x500 اگر در تنظیمات نبود
						origImg = resize.Resize(500, 500, origImg, resize.Lanczos3)
					}
				}
			}
		}

		if isSmart {
			// Use SmartCompress util
			filetype := http.DetectContentType(imgBytes[:512])
			smartOutFormat := outFormat
			if smartOutFormat == "original" {
				smartOutFormat = "" // Let SmartCompress decide
			}
			opts := utils.SmartOptions{
				OutputFormat: smartOutFormat,
				OriginalSize: len(imgBytes),
			}
			res, err := utils.SmartCompress(origImg, filetype, opts)
			if err != nil {
				log.Println("smart compress error:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			buf = *res.Data
			outFormat = res.Format // may differ
		} else {
			switch outFormat {
			case "jpeg", "jpg":
				if err := jpeg.Encode(&buf, origImg, &jpeg.Options{Quality: qualityPercent}); err != nil {
					log.Println("jpeg encode error:", err)
					http.Error(w, "encode error: "+err.Error(), http.StatusInternalServerError)
					return
				}
			case "png":
				if err := png.Encode(&buf, origImg); err != nil {
					log.Println("png encode error:", err)
					http.Error(w, "encode error: "+err.Error(), http.StatusInternalServerError)
					return
				}
			case "webp":
				if err := webp.Encode(&buf, origImg, &webp.Options{Quality: float32(qualityPercent)}); err != nil {
					log.Println("webp encode error:", err)
					http.Error(w, "encode error: "+err.Error(), http.StatusInternalServerError)
					return
				}
			default:
				log.Println("jpeg encode error default:", err)
				http.Error(w, "encode error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// در صورت تغییر فرمت، پسوند را بروزرسانی می‌کنیم
		extMap := map[string]string{"jpeg": ".jpg", "jpg": ".jpg", "png": ".png", "webp": ".webp"}
		ext = extMap[outFormat]
		if ext == "" {
			ext = filepath.Ext(header.Filename)
		}
		uniqueName = generateUniqueName(ext)
		// استفاده از تابع helper برای ساخت ایمن مسیر
		absolutePath, err = buildSafeMediaPath(publicDir, baseCategory, subfolder, uniqueName)
		if err != nil {
			sendJSON(http.StatusBadRequest, utils.New("VALIDATION_ERROR", "خطا در ساخت مسیر", err.Error()))
			return
		}
		relativePath = filepath.ToSlash(strings.TrimPrefix(absolutePath, publicDir))
		if !strings.HasPrefix(relativePath, "/") {
			relativePath = "/" + relativePath
		}

		// ذخیره فایل فشرده شده
		if err := os.MkdirAll(filepath.Dir(absolutePath), 0755); err != nil {
			sendJSON(http.StatusInternalServerError, map[string]interface{}{"success": false, "error": "Cannot create directory"})
			return
		}
		_ = os.WriteFile(absolutePath, buf.Bytes(), 0644)

		finalSize = uint(buf.Len())
		filetype = "image/" + outFormat
		if outFormat == "jpg" || outFormat == "jpeg" {
			filetype = "image/jpeg"
		}

		// بک‌آپ نسخه اصلی در BackupService (بدون تغییر فرمت)
		{
			bs := services.NewBackupService("")
			// مسیر و نام فایلِ اصلی قبل از فشرده‌سازی
			origExt := filepath.Ext(header.Filename) // همان پسوند اولیه
			// برای اینکه متد Revert بتواند بک‌آپ را پیدا کند، باید basePath یکسان باشد
			basePart := strings.TrimSuffix(uniqueName, ext)
			origName := basePart + origExt
			origRelPath := filepath.ToSlash(filepath.Join("/uploads", "media", baseCategory, subfolder, origName))
			// برای جلوگیری از دوباره‌نویسی روی مسیر uploads، یک فایل موقتی می‌سازیم
			tmpFile, err := os.CreateTemp("", "orig-*"+origExt)
			if err == nil {
				_ = os.WriteFile(tmpFile.Name(), imgBytes, 0644)
				_ = tmpFile.Close()
				_, errBkp := bs.CopyToDailyBackup(tmpFile.Name(), strings.TrimPrefix(origRelPath, "/"))
				if errBkp != nil {
					log.Println("backup original error:", errBkp)
				}
				_ = os.Remove(tmpFile.Name()) // پاک‌سازی فایل موقت
			}
		}

		// اگر فایل گیف است، هیچ فشرده‌سازی یا تبدیل فرمتی انجام نشود
		if filetype == "image/gif" || strings.ToLower(filepath.Ext(header.Filename)) == ".gif" {
			// فایل را به صورت اصلی ذخیره کن و هیچ تغییری نده
			file.Seek(0, io.SeekStart)
			outBytes, err := io.ReadAll(file)
			if err != nil {
				sendJSON(http.StatusBadRequest, utils.New("FILE_READ_ERROR", "خطا در خواندن فایل گیف", err.Error()))
				return
			}
			if err := os.WriteFile(absolutePath, outBytes, 0644); err != nil {
				sendJSON(http.StatusInternalServerError, utils.New("SAVE_ERROR", "خطا در ذخیره فایل گیف", err.Error()))
				return
			}
			// برای GIF سایز همان سایز اصلی است
			finalSize = origSize
			// ادامه می‌دهیم تا رکورد دیتابیس ثبت شود (return حذف شد)
		}
	} else if filetype == "image/gif" {
		// ذخیره مستقیم فایل GIF بدون فشرده‌سازی (وقتی compression disabled است)
		file.Seek(0, io.SeekStart)
		gifBytes, err := io.ReadAll(file)
		if err != nil {
			sendJSON(http.StatusBadRequest, utils.New("FILE_READ_ERROR", "خطا در خواندن فایل گیف", err.Error()))
			return
		}
		if err := os.WriteFile(absolutePath, gifBytes, 0644); err != nil {
			sendJSON(http.StatusInternalServerError, utils.New("SAVE_ERROR", "خطا در ذخیره فایل گیف", err.Error()))
			return
		}
		finalSize = origSize
	} else {
		// برای فایل‌های غیر تصویری (ویدیو، صدا، اسناد)، کل فایل را بخوانیم
		var fileData []byte
		if allowedVideoTypes[filetype] || allowedAudioTypes[filetype] {
			// کل فایل را در حافظه بخوانیم
			log.Printf("Reading full file for %s (size: %d bytes)", filetype, header.Size)
			fileData, err = io.ReadAll(file)
			if err != nil {
				sendJSON(http.StatusBadRequest, utils.New("FILE_READ_ERROR", "خطا در خواندن فایل", err.Error()))
				return
			}
			log.Printf("Successfully read %d bytes for %s", len(fileData), filetype)
		} else {
			// برای سایر فایل‌ها، از همان buff استفاده کنیم
			fileData = buff
			log.Printf("Using buffer for %s (size: %d bytes)", filetype, len(fileData))
		}

		// بک‌آپ فایل در حالت بدون فشرده‌سازی
		bs := services.NewBackupService("")
		err = bs.SaveOriginalAndBackup(absolutePath, strings.TrimPrefix(relativePath, "/"), fileData)
		if err != nil {
			log.Printf("Error saving file %s: %v", absolutePath, err)
			sendJSON(http.StatusInternalServerError, utils.New("SAVE_ERROR", "خطا در ذخیره فایل", err.Error()))
			return
		}
		log.Printf("Successfully saved file: %s (size: %d bytes)", absolutePath, len(fileData))
	}

	var compressedSizePtr *uint = nil
	if finalSize != origSize {
		compressedSizePtr = &finalSize
	}
	media := models.MediaFile{
		FileName:       uniqueName,
		Title:          header.Filename,
		FileType:       filetype,
		MimeType:       filetype,
		Size:           origSize,
		Compressed:     compressedSizePtr != nil,
		CompressedSize: compressedSizePtr,
		URL:            relativePath,
		Category:       baseCategory,
		UploadedBy:     uploadedBy,
	}

	if err := h.DB.Create(&media).Error; err != nil {
		sendJSON(http.StatusInternalServerError, map[string]interface{}{"success": false, "error": "Database error"})
		return
	}

	// Log the newly created media record for debugging purposes
	log.Printf("Media file created with ID: %d, URL: %s", media.ID, media.URL)

	// -----------------------------------------------------------
	//  سیستم زمان‌بندی فشرده‌سازی ویدیو
	// -----------------------------------------------------------
	if allowedVideoTypes[filetype] || allowedVideoExt[ext] {
		// اگر فایل ویدیویی بود، job فشرده‌سازی ویدیو ایجاد کن
		if media.ID != 0 {
			// دریافت تنظیمات فشرده‌سازی از فرم یا query
			// تنظیمات مستقیماً از فرم گرفته می‌شود و در فیلدهای مربوطه قرار می‌گیرد

			// تنظیمات پیش‌فرض از UI
			targetFormat := r.FormValue("format")
			if targetFormat == "" {
				targetFormat = "mp4" // پیش‌فرض
			}

			targetQuality := r.FormValue("quality")
			if targetQuality == "" {
				targetQuality = "720p" // پیش‌فرض
			}

			compressionJob := models.CompressionJob{
				MediaID:       media.ID,
				Status:        "pending",
				JobType:       "video_compression",
				TargetFormat:  targetFormat,
				TargetQuality: targetQuality,
				Progress:      0,
				OriginalSize:  &media.Size,
				StartedAt:     nil,
				FinishedAt:    nil,
				Log:           "",
				ErrorMessage:  "",
			}
			if err := h.DB.Create(&compressionJob).Error; err != nil {
				log.Printf("Error creating compression job for video %d: %v", media.ID, err)
			} else {
				log.Printf("Created pending compression job %d for video %d", compressionJob.ID, media.ID)
				go h.scheduleVideoCompression()
			}
		}
	}

	// اگر فایل تصویری بود و دسته‌بندی products یا product-categories بود، واریانت‌ها را بساز و ذخیره کن
	if allowedImageTypes[filetype] && baseCategory == "products" {
		imgFile, err := os.Open(absolutePath)
		if err == nil {
			defer imgFile.Close()
			img, _, err := image.Decode(imgFile)
			if err == nil {
				// ساخت و ذخیره واریانت‌ها
				// Start with empty map, will be filled from settings or defaults
				sizes := map[string]struct{ width, height uint }{}

				// Load sizes from settings first
				if h.SettingService != nil {
					// برای محصولات و دسته‌ها از کلید مشترک image_sizes استفاده می‌کنیم
					if cfg, err := h.SettingService.GetSetting(r.Context(), "image_sizes"); err == nil && cfg != nil {
						type sizeCfg struct {
							Width  uint `json:"width"`
							Height uint `json:"height"`
						}
						var mp map[string]sizeCfg
						if json.Unmarshal([]byte(cfg.Value), &mp) == nil {
							for k, v := range mp {
								// Add size based on category and key
								shouldAdd := false
								switch baseCategory {
								case "product-categories":
									shouldAdd = k == "category"
								case "brands":
									shouldAdd = k == "brand"
								case "products":
									shouldAdd = k == "thumbnail" || k == "small" || k == "medium" || k == "large"
								default:
									shouldAdd = true // For other categories, add all sizes
								}

								if shouldAdd {
									sizes[k] = struct{ width, height uint }{v.Width, v.Height}
								}
							}
						}
					}
				}

				// Fill missing sizes with defaults only if not in settings
				if baseCategory == "products" {
					// For products, only add the 4 standard sizes if not in settings
					for k, v := range imageSizesDefault {
						if _, exists := sizes[k]; !exists {
							sizes[k] = v
						}
					}
				} else if baseCategory == "brands" {
					// For brands, add only default brand size if not in settings
					if _, exists := sizes["brand"]; !exists {
						sizes["brand"] = struct{ width, height uint }{200, 100}
					}
				} else if baseCategory == "product-categories" {
					// For categories, add only default category size if not in settings
					if _, exists := sizes["category"]; !exists {
						sizes["category"] = struct{ width, height uint }{500, 500}
					}
				} else {
					// For other categories, add all default sizes
					for k, v := range imageSizesDefault {
						if _, exists := sizes[k]; !exists {
							sizes[k] = v
						}
					}
				}
				for sizeName, dimensions := range sizes {
					// Skip resizing if both width and height are 0 (original size)
					if dimensions.width == 0 && dimensions.height == 0 {
						continue
					}

					resized := resize.Resize(dimensions.width, dimensions.height, img, resize.Lanczos3)
					variantName := fmt.Sprintf("%s_%s%s", uniqueName[:len(uniqueName)-len(ext)], sizeName, ext)
					variantRelPath := filepath.ToSlash(filepath.Join("/uploads", "media", baseCategory, subfolder, variantName))
					variantAbsPath := filepath.Join(publicDir, "uploads", "media", baseCategory, subfolder, variantName)
					variantFile, err := os.Create(variantAbsPath)
					if err != nil {
						continue
					}
					switch filetype {
					case "image/jpeg":
						_ = jpeg.Encode(variantFile, resized, &jpeg.Options{Quality: variantQuality})
					case "image/png":
						png.Encode(variantFile, resized)
					case "image/webp":
						_ = webp.Encode(variantFile, resized, &webp.Options{Quality: float32(variantQuality)})
					}
					variantFile.Close()
					variantInfo, _ := os.Stat(variantAbsPath)
					variant := models.MediaVariant{
						MediaID:  media.ID,
						Width:    int(dimensions.width),
						Height:   int(dimensions.height),
						FilePath: variantRelPath,
						FileSize: uint(variantInfo.Size()),
						Purpose:  sizeName,
					}
					_ = h.DB.Create(&variant).Error // خطاهای واریانت را لاگ نمی‌کنیم تا آپلود اصلی fail نشود
				}
			}
		}
	}

	sendJSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"files":   []models.MediaFile{media},
	})

	// زمان‌بندی تولید متادیتای تصویر توسط AI (با تاخیر تنظیم‌شده)
	go func() {
		// فقط برای تصاویر
		if !strings.HasPrefix(strings.ToLower(filetype), "image/") {
			return
		}
		if h.ImageSEOWorker == nil {
			// تلاش برای ساخت سریع worker بر اساس سرویس‌های در دسترس
			apiSvc := services.NewAPISettingsService(h.DB)
			seoSvc := services.NewImageSEOService(h.DB, h.SettingService, apiSvc)
			h.ImageSEOWorker = services.NewImageSEOWorker(h.DB, h.SettingService, apiSvc, seoSvc)
			h.ImageSEOWorker.Start()
		}
		lang := r.FormValue("context_lang")
		pageTitle := r.FormValue("context_title")
		h.ImageSEOWorker.ScheduleOnUpload(r.Context(), media.ID, pageTitle, lang)
	}()
}

// دریافت رسانه بر اساس ID
func (h *MediaHandler) GetMediaHandler(w http.ResponseWriter, r *http.Request) {
	// تلاش برای استخراج شناسه از Gin context (پارامتر مسیر :id)
	var idStr string
	if gctxVal := r.Context().Value("gin_context"); gctxVal != nil {
		if gc, ok := gctxVal.(*gin.Context); ok && gc != nil {
			idStr = strings.TrimSpace(gc.Param("id"))
		}
	}
	// در صورت نبود Gin یا پارامتر، آخرین سگمنت مسیر به‌عنوان ID در نظر گرفته می‌شود
	if idStr == "" {
		path := strings.TrimSuffix(r.URL.Path, "/")
		parts := strings.Split(path, "/")
		if len(parts) > 0 {
			idStr = parts[len(parts)-1]
		}
	}

	id, err := parseUintFromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Invalid media ID"})
		return
	}
	var media models.MediaFile
	if err := h.DB.Preload("Variants").First(&media, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Media not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": media})
}

// حذف رسانه بر اساس ID
func (h *MediaHandler) DeleteMediaHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Invalid media ID"})
		return
	}

	// واکشی رسانه برای بررسی دسته‌بندی
	var media models.MediaFile
	if err := h.DB.First(&media, id).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Media not found"})
		return
	}

	// --- کنترل پرمیژن حذف بر اساس دسته‌بندی ---
	// برای developer، دسترسی کامل فرض می‌شود
	// در آینده می‌توان پرمیژن‌های دسته‌ای مثل media_library.delete:products اضافه کرد

	// حذف فیزیکی واریانت‌ها
	var variants []models.MediaVariant
	h.DB.Where("media_id = ?", id).Find(&variants)
	projectRoot := ""
	if wd, err := os.Getwd(); err == nil {
		projectRoot = filepath.Dir(wd)
	}
	publicDir := filepath.Join(projectRoot, "public")
	for _, v := range variants {
		absPath := filepath.Join(publicDir, v.FilePath)
		_ = os.Remove(absPath) // اگر نبود، خطا مهم نیست
	}
	// حذف رکورد واریانت‌ها
	h.DB.Where("media_id = ?", id).Delete(&models.MediaVariant{})
	// حذف فایل اصلی + پاکسازی بک‌آپ
	absPath := filepath.Join(publicDir, media.URL)
	_ = os.Remove(absPath)
	// remove backups of this file across periods
	bs := services.NewBackupService("")
	relPath := strings.TrimPrefix(media.URL, "/")
	if n, err := bs.RemoveBackups(relPath); err == nil {
		log.Printf("[delete] removed %d backup copies of %s", n, relPath)
	}
	// حذف رسانه
	if err := h.DB.Delete(&models.MediaFile{}, id).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Failed to delete media"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "Media and its variants deleted successfully"})
}

// ویرایش رسانه (فقط فیلدهای قابل ویرایش)
func (h *MediaHandler) UpdateMediaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/media/")
	id, err := parseUintFromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Invalid media ID"})
		return
	}
	var req struct {
		Title       *string `json:"title"`
		AltText     *string `json:"alt_text"`
		Caption     *string `json:"caption"`
		Description *string `json:"description"`
		MimeType    *string `json:"mime_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Invalid request body"})
		return
	}
	updates := map[string]interface{}{}
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.AltText != nil {
		updates["alt_text"] = *req.AltText
	}
	if req.Caption != nil {
		updates["short_description"] = *req.Caption
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.MimeType != nil {
		updates["mime_type"] = *req.MimeType
	}
	if err := h.DB.Model(&models.MediaFile{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// Escape error message to prevent XSS
		errorMsg := html.EscapeString(err.Error())
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Failed to update media: " + errorMsg})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "Media updated"})
}

// RegenerateMediaHandler handles POST /media/regenerate?type=xxx
func (h *MediaHandler) RegenerateMediaHandler(w http.ResponseWriter, r *http.Request) {
	typeKey := r.URL.Query().Get("type")
	if typeKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "type is required"}`))
		return
	}

	// TODO: Find all media of this type, delete old resized files, and regenerate new ones
	log.Printf("Regenerating images for type: %s", typeKey)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Regeneration started for type: ` + typeKey + `"}`))
}

// PreviewCompressHandler پیش‌نمایش زنده فشرده‌سازی و تبدیل فرمت تصویر
// ورودی: فایل تصویر (multipart)، فرمت خروجی (jpeg/webp/avif/png)، کیفیت (1-100)
// خروجی: تصویر فشرده‌شده با Content-Type مناسب
func (h *MediaHandler) PreviewCompressHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	// محدودیت حجم فایل (مثلاً 10MB)
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)
	if err := r.ParseMultipartForm(10 * 1024 * 1024); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Invalid form data"})
		return
	}

	// ------------------------------------------------------------------------------------
	// ۱) دریافت فایل و متادیتا
	// ------------------------------------------------------------------------------------
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "No file uploaded"})
		return
	}
	defer file.Close()

	// کل فایل را در حافظه می‌خوانیم (تا سقف ۱۰MB)؛ هم برای image.Decode و هم برای محاسبه اندازه.
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Failed to read file"})
		return
	}

	// MIME ورودی را با نگاه به ۵۱۲ بایت اول تشخیص می‌دهیم.
	inputMime := http.DetectContentType(fileBytes[:512])

	// Decode تصویر از روی بایت‌های خام.
	var origImg image.Image
	extLower := strings.ToLower(filepath.Ext(header.Filename))
	switch extLower {
	case ".webp":
		origImg, err = webp.Decode(bytes.NewReader(fileBytes))
	default:
		origImg, _, err = image.Decode(bytes.NewReader(fileBytes))
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Invalid image file"})
		return
	}

	// ------------------------------------------------------------------------------------
	// ۲) بررسی فعال بودن حالت Smart – اگر فیلد smart=true باشد، از منطق هوشمند استفاده می‌شود.
	// ------------------------------------------------------------------------------------
	if sm := strings.ToLower(r.FormValue("smart")); sm == "true" || sm == "1" || sm == "advanced" {
		// تنظیم گزینه‌های SmartOptions بر اساس ورودی کاربر
		targetKB := 0
		if ts := r.FormValue("targetSizeKB"); ts != "" {
			if v, _ := strconv.Atoi(ts); v > 0 {
				targetKB = v
			}
		}
		// پارس پارامترهای اختیاری پیشرفته
		qStep := 0
		if qs := r.FormValue("qualityStep"); qs != "" {
			if v, _ := strconv.Atoi(qs); v > 0 {
				qStep = v
			}
		}
		ssimTarget := 0.97 // حداقل SSIM پیش‌فرض برای حفظ بافت و رنگ
		if st := r.FormValue("ssimTarget"); st != "" {
			if v, _ := strconv.ParseFloat(st, 64); v > 0 {
				ssimTarget = v
			}
		}

		opts := utils.SmartOptions{
			TargetSizeKB: targetKB,
			OutputFormat: r.FormValue("format"),
			MaxWidth:     0,
			MaxHeight:    0,
			OriginalSize: int(header.Size),
			AllowResize:  r.FormValue("allowResize") == "1" || strings.ToLower(r.FormValue("allowResize")) == "true",
			QualityStep:  qStep,
			SSIMTarget:   ssimTarget,
		}

		res, err := utils.SmartCompress(origImg, inputMime, opts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": err.Error()})
			return
		}
		w.Header().Set("Content-Type", res.MimeType)
		w.Write(res.Data.Bytes())
		return
	}

	// ------------------------------------------------------------------------------------
	// ۳) حالت معمول (غیر هوشمند) – از پارامتر format + quality استفاده می‌کنیم.
	// ------------------------------------------------------------------------------------
	format := strings.ToLower(r.FormValue("format"))
	if format == "" {
		format = "jpeg"
	}
	quality := 80
	if qs := r.FormValue("quality"); qs != "" {
		if q, err := strconv.Atoi(qs); err == nil && q >= 1 && q <= 100 {
			quality = q
		}
	}

	switch format {
	case "jpeg", "jpg":
		w.Header().Set("Content-Type", "image/jpeg")
		jpeg.Encode(w, origImg, &jpeg.Options{Quality: quality})
	case "png":
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, origImg)
	case "webp":
		w.Header().Set("Content-Type", "image/webp")
		if err := webp.Encode(w, origImg, &webp.Options{Quality: float32(quality)}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Failed to encode image"})
			return
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Unsupported output format"})
	}
}

// CompressMediaHandler handles POST /media/compress/:id
// It compresses the image in-place (using JPEG 75) after copying the original to the central backup folder.
func (h *MediaHandler) CompressMediaHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := parseUintFromString(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var media models.MediaFile
	if err := h.DB.First(&media, id).Error; err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if !strings.HasPrefix(media.FileType, "image/") {
		http.Error(w, "only images supported", http.StatusBadRequest)
		return
	}

	// Build abs path
	projectRoot := ""
	if wd, err := os.Getwd(); err == nil {
		projectRoot = filepath.Dir(wd)
	}
	publicDir := filepath.Join(projectRoot, "public")
	rel := strings.TrimPrefix(media.URL, "/")
	absPath := filepath.Join(publicDir, rel)

	// (removed) debug placeholders; actual renaming handled later

	// -----------------------------------------------------------
	//  Centralised backup copy (monthly folder)
	// -----------------------------------------------------------
	bs := services.NewBackupService("")
	if _, err := bs.CopyToDailyBackup(absPath, strings.TrimPrefix(rel, "/")); err != nil {
		log.Println("backup copy error:", err)
		http.Error(w, "cannot create backup: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// حالا فایل اصلی را می‌خوانیم
	imgBytes, err := os.ReadFile(absPath)
	if err != nil {
		log.Println("compress read image error:", err)
		http.Error(w, "cannot read image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var origImg image.Image
	extLower := strings.ToLower(filepath.Ext(absPath))
	switch extLower {
	case ".webp":
		origImg, err = webp.Decode(bytes.NewReader(imgBytes))
	default:
		origImg, _, err = image.Decode(bytes.NewReader(imgBytes))
	}
	if err != nil {
		log.Println("compress decode error:", err)
		http.Error(w, "unsupported image: "+err.Error(), http.StatusBadRequest)
		return
	}

	// --- Load compression settings (same logic as upload) ---
	settingMap := map[string]string{}
	if h.SettingService != nil {
		if settings, err := h.SettingService.GetSettingsByCategory(r.Context(), "compression"); err == nil {
			for _, s := range settings {
				settingMap[s.Key] = s.Value
			}
		}
	}

	qualityKey := settingMap["compression.quality"]
	if qualityKey == "" {
		qualityKey = "medium"
	}
	formatKey := settingMap["compression.format"]
	if formatKey == "" {
		formatKey = "original"
	}

	// quality percent & smart handling
	qualityPercent := 75
	isSmart := false
	switch strings.ToLower(qualityKey) {
	case "high":
		qualityPercent = 90
	case "medium":
		qualityPercent = 75
	case "low":
		qualityPercent = 60
	case "custom":
		if v, _ := strconv.Atoi(settingMap["compression.custom_quality"]); v >= 1 && v <= 100 {
			qualityPercent = v
		}
	case "smart":
		isSmart = true
	}

	// بررسی پارامترهای query برای override
	qStr := r.URL.Query().Get("quality")
	smartStr := r.URL.Query().Get("smart")
	formatParam := strings.ToLower(r.URL.Query().Get("format"))
	if smartStr == "1" || strings.ToLower(smartStr) == "true" {
		isSmart = true
	}
	outFormat := strings.ToLower(formatKey)
	// Override via query param, if provided
	if formatParam != "" {
		outFormat = formatParam
	}
	if qStr != "" {
		if q, err := strconv.Atoi(qStr); err == nil && q >= 1 && q <= 100 {
			qualityPercent = q
			isSmart = false // اگر کیفیت سفارشی ارسال شود، حالت هوشمند خاموش می‌شود
		}
	}

	if outFormat == "original" {
		ext := strings.ToLower(filepath.Ext(absPath))
		switch ext {
		case ".jpg", ".jpeg":
			outFormat = "jpeg"
		case ".png":
			outFormat = "png"
		case ".webp":
			outFormat = "webp"
		default:
			outFormat = "jpeg"
		}
	}

	var buf bytes.Buffer

	if isSmart {
		// Use SmartCompress util
		filetype := http.DetectContentType(imgBytes[:512])
		smartOutFormat := outFormat
		if smartOutFormat == "original" {
			smartOutFormat = "" // Let SmartCompress decide
		}
		opts := utils.SmartOptions{
			OutputFormat: smartOutFormat,
			OriginalSize: len(imgBytes),
		}
		res, err := utils.SmartCompress(origImg, filetype, opts)
		if err != nil {
			log.Println("smart compress error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		buf = *res.Data
		outFormat = res.Format // may differ
	} else {
		switch outFormat {
		case "jpeg", "jpg":
			if err := jpeg.Encode(&buf, origImg, &jpeg.Options{Quality: qualityPercent}); err != nil {
				log.Println("jpeg encode error:", err)
				http.Error(w, "encode error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		case "png":
			if err := png.Encode(&buf, origImg); err != nil {
				log.Println("png encode error:", err)
				http.Error(w, "encode error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		case "webp":
			if err := webp.Encode(&buf, origImg, &webp.Options{Quality: float32(qualityPercent)}); err != nil {
				log.Println("webp encode error:", err)
				http.Error(w, "encode error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			log.Println("jpeg encode error default:", err)
			http.Error(w, "encode error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// ---------------------------------------------
	// Rename file if output format requires new ext
	// ---------------------------------------------
	originalExt := strings.ToLower(filepath.Ext(absPath))
	var outExt string
	switch outFormat {
	case "jpeg", "jpg":
		outExt = ".jpg"
	case "png":
		outExt = ".png"
	case "webp":
		outExt = ".webp"
	default:
		outExt = originalExt
	}

	// Determine final absolute/relative path for writing
	finalAbsPath := absPath
	finalRel := rel // keep track for DB
	needRename := outExt != originalExt
	if needRename {
		baseNoExt := strings.TrimSuffix(filepath.Base(absPath), originalExt)
		newBase := baseNoExt + outExt
		finalRel = filepath.ToSlash(filepath.Join(filepath.Dir(rel), newBase))
		finalAbsPath = filepath.Join(publicDir, finalRel)
		// Ensure directory exists
		if err := os.MkdirAll(filepath.Dir(finalAbsPath), 0o750); err != nil {
			log.Println("mkdir error:", err)
			http.Error(w, "cannot save image: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// write compressed bytes to file (new or same path)
	if err := os.WriteFile(finalAbsPath, buf.Bytes(), 0644); err != nil {
		log.Println("save image error:", err)
		http.Error(w, "cannot save image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If renamed, delete original file
	if needRename {
		_ = os.Remove(absPath)
	}

	// ---------------------------------------------
	// Update database metadata
	// ---------------------------------------------
	newURL := "/" + filepath.ToSlash(finalRel)
	newFileName := filepath.Base(finalRel)
	mime := "image/" + outFormat
	if outFormat == "jpeg" {
		mime = "image/jpeg"
	}
	newSize := uint(len(buf.Bytes()))
	updates := map[string]interface{}{
		"compressed":      true,
		"compressed_size": newSize,
		"size":            newSize,
	}
	if needRename {
		updates["url"] = newURL
		updates["file_name"] = newFileName
		updates["file_type"] = mime
		updates["mime_type"] = mime
		updates["format"] = outFormat
		if media.OriginalFormat == "" {
			updates["original_format"] = strings.TrimPrefix(originalExt, ".")
		}
		// Keep OriginalFormat as is
		media.URL = newURL
	}
	h.DB.Model(&models.MediaFile{}).Where("id = ?", id).Updates(updates)

	// Build backup URL (original image before compression)
	backupRel := strings.TrimPrefix(rel, "/")
	if strings.HasPrefix(backupRel, "media/uploads/") {
		backupRel = strings.TrimPrefix(backupRel, "media/uploads/")
	} else if strings.HasPrefix(backupRel, "uploads/") {
		backupRel = strings.TrimPrefix(backupRel, "uploads/")
	}
	period := time.Now().Format("2006-01")
	backupURL := filepath.ToSlash(filepath.Join("/uploads", "media", "backup", period, backupRel))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":         true,
		"size":            newSize, // سایز فعلی فایل بعد از فشرده‌سازی
		"compressed_size": newSize, // برای سازگاری کلاینت
		"url":             media.URL,
		"original_url":    backupURL,
	})
}

// -----------------------------------------------------------------------------
// Helper functions for centralised backups
// -----------------------------------------------------------------------------

// DownloadMediaHandler handles GET /media/download/:id and streams the file as attachment
func (h *MediaHandler) DownloadMediaHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) == 0 {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}
	idStr := parts[len(parts)-1]
	id, err := parseUintFromString(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var media models.MediaFile
	if err := h.DB.First(&media, id).Error; err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	projectRoot := ""
	if wd, err := os.Getwd(); err == nil {
		projectRoot = filepath.Dir(wd)
	}
	publicDir := filepath.Join(projectRoot, "public")
	// اعتبارسنجی و پاکسازی مسیر برای جلوگیری از path traversal
	absPath, err := validateAndSanitizePath(publicDir, media.URL)
	if err != nil {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	// اصلاح هدر Content-Disposition بر اساس نوع فایل
	if strings.HasPrefix(media.FileType, "video/") || strings.HasPrefix(media.FileType, "image/") {
		// برای ویدیو و عکس، به صورت inline باشد تا مرورگر پخش کند
		w.Header().Set("Content-Disposition", "inline; filename="+media.FileName)
	} else {
		// برای سایر فایل‌ها (مثلاً اسناد)، attachment باشد
		w.Header().Set("Content-Disposition", "attachment; filename="+media.FileName)
	}
	http.ServeFile(w, r, absPath)
}

// RotateMediaHandler handles POST /media/rotate/:id?angle=90 (supports 90,180,270)
func (h *MediaHandler) RotateMediaHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) == 0 {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}
	idStr := parts[len(parts)-1]
	id, err := parseUintFromString(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var media models.MediaFile
	if err := h.DB.First(&media, id).Error; err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if !strings.HasPrefix(media.FileType, "image/") {
		http.Error(w, "only images supported", http.StatusBadRequest)
		return
	}

	angle := 90
	if aStr := r.URL.Query().Get("angle"); aStr != "" {
		if a, err := strconv.Atoi(aStr); err == nil && a == 90 || a == 180 || a == 270 {
			angle = a
		}
	}

	projectRoot := ""
	if wd, err := os.Getwd(); err == nil {
		projectRoot = filepath.Dir(wd)
	}
	publicDir := filepath.Join(projectRoot, "public")
	rel := strings.TrimPrefix(media.URL, "/")
	absPath := filepath.Join(publicDir, rel)

	// Security: Validate TIFF files to prevent CVE-2023-36808 (crash on crafted TIFF)
	ext := strings.ToLower(filepath.Ext(absPath))
	if ext == ".tiff" || ext == ".tif" {
		if fileInfo, err := os.Stat(absPath); err == nil {
			// Limit TIFF file size to 10MB to mitigate risk
			if fileInfo.Size() > 10*1024*1024 {
				http.Error(w, "TIFF file too large (max 10MB)", http.StatusBadRequest)
				return
			}
		}
	}

	// استفاده از safeOpenImage برای جلوگیری از panic (CVE-2023-36308)
	img, err := safeOpenImage(absPath)
	if err != nil {
		http.Error(w, "cannot open image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var rotated *image.NRGBA
	switch angle {
	case 90:
		rotated = imaging.Rotate90(img)
	case 180:
		rotated = imaging.Rotate180(img)
	case 270:
		rotated = imaging.Rotate270(img)
	default:
		rotated = imaging.Rotate90(img)
	}

	switch strings.ToLower(filepath.Ext(absPath)) {
	case ".webp":
		f, err := os.Create(absPath)
		if err != nil {
			http.Error(w, "cannot save image", http.StatusInternalServerError)
			return
		}
		defer f.Close()
		if err := webp.Encode(f, rotated, &webp.Options{Lossless: true}); err != nil {
			http.Error(w, "cannot encode webp", http.StatusInternalServerError)
			return
		}
	default:
		if err := imaging.Save(rotated, absPath); err != nil {
			http.Error(w, "cannot save image", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "width": rotated.Bounds().Dx(), "height": rotated.Bounds().Dy()})
}

// RevertCompressionHandler handles POST /media/revert/:id
// It restores the original (uncompressed) version of a media file using BackupService utilities.
func (h *MediaHandler) RevertCompressionHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, err := parseUintFromString(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var media models.MediaFile
	if err := h.DB.First(&media, id).Error; err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	// Initialise backup service (auto-detect project root)
	bsvc := services.NewBackupService("")
	projectRoot := bsvc.ProjectRoot
	publicDir := filepath.Join(projectRoot, "public")

	rel := strings.TrimPrefix(media.URL, "/")
	absPath := filepath.Join(publicDir, rel)

	// Variables for potential rename when the original extension differs
	needRename := false
	newRel := rel
	newAbs := absPath

	// Utility to try a list of candidate relative paths and return the first existing backup.
	tryCandidates := func(candidates []string) string {
		for _, cand := range candidates {
			// 1) original relative path
			if p, err := bsvc.FindLatestBackup(cand); err == nil {
				return p
			}
			// 2) trimmed variants matching CopyToDailyBackup logic
			slash := filepath.ToSlash(cand)
			trimmed := strings.TrimPrefix(slash, "uploads/")
			trimmed = strings.TrimPrefix(trimmed, "media/uploads/")
			if trimmed != slash {
				if p, err := bsvc.FindLatestBackup(trimmed); err == nil {
					return p
				}
			}
		}
		return ""
	}

	basePath := strings.TrimSuffix(rel, filepath.Ext(rel))

	// Priority 1: original format if we have it stored
	var backupPath string
	if media.OriginalFormat != "" {
		origExt := strings.ToLower(media.OriginalFormat)
		if origExt != "" && origExt[0] != '.' {
			origExt = "." + origExt
		}
		candList := []string{basePath + origExt}
		if origExt == ".jpeg" {
			candList = append(candList, basePath+".jpg")
		} else if origExt == ".jpg" {
			candList = append(candList, basePath+".jpeg")
		}
		backupPath = tryCandidates(candList)
		if backupPath != "" {
			needRename = filepath.Ext(rel) != filepath.Ext(candList[0])
			if needRename {
				newRel = basePath + filepath.Ext(backupPath)
				newAbs = filepath.Join(publicDir, newRel)
			}
		}
	}

	// Fallback: look for same relative path in backups or any other extension
	if backupPath == "" {
		// first direct match
		if p, err := bsvc.FindLatestBackup(rel); err == nil {
			backupPath = p
		} else {
			// extensions list to probe
			extCandidates := []string{".jpg", ".jpeg", ".png", ".webp", ".gif", ".bmp", ".tiff"}
			probe := make([]string, 0, len(extCandidates))
			for _, ext := range extCandidates {
				probe = append(probe, basePath+ext)
			}
			backupPath = tryCandidates(probe)
			if backupPath != "" && filepath.Ext(rel) != filepath.Ext(backupPath) {
				needRename = true
				newRel = basePath + filepath.Ext(backupPath)
				newAbs = filepath.Join(publicDir, newRel)
			}
		}
	}

	if backupPath == "" {
		http.Error(w, "backup not found", http.StatusBadRequest)
		return
	}

	// Ensure target directory exists
	if err := os.MkdirAll(filepath.Dir(newAbs), 0o750); err != nil {
		http.Error(w, "cannot restore", http.StatusInternalServerError)
		return
	}

	if err := services.CopyFile(backupPath, newAbs); err != nil {
		http.Error(w, "cannot restore", http.StatusInternalServerError)
		return
	}

	// compute restored file size (after copying)
	fi, _ := os.Stat(newAbs)
	size := uint(fi.Size())

	if needRename {
		_ = os.Remove(absPath)
		mime := "image/jpeg" // default MIME
		switch filepath.Ext(newRel) {
		case ".jpg", ".jpeg":
			mime = "image/jpeg"
		case ".png":
			mime = "image/png"
		case ".webp":
			mime = "image/webp"
		case ".gif":
			mime = "image/gif"
		}
		updates := map[string]interface{}{
			"url":             "/" + filepath.ToSlash(newRel),
			"file_name":       filepath.Base(newRel),
			"file_type":       mime,
			"mime_type":       mime,
			"format":          strings.TrimPrefix(filepath.Ext(newRel), "."),
			"compressed":      false,
			"compressed_size": nil,
			"size":            size,
		}
		h.DB.Model(&models.MediaFile{}).Where("id = ?", id).Updates(updates)
	} else {
		// only compressed flags changed but size must be refreshed
		h.DB.Model(&models.MediaFile{}).Where("id = ?", id).Updates(map[string]interface{}{
			"compressed":      false,
			"compressed_size": nil,
			"size":            size,
		})
	}

	// Fetch updated record for response
	var updated models.MediaFile
	h.DB.First(&updated, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "url": updated.URL, "size": updated.Size, "format": updated.Format, "file_name": updated.FileName})
}

// -----------------------------------------------------------------------------
// Video-specific handlers (for VideoHandler wrapper)
// -----------------------------------------------------------------------------

// ListVideosHandler لیست ویدیوها
func (h *MediaHandler) ListVideosHandler(w http.ResponseWriter, r *http.Request) {
	// فعلاً از همان ListMediaHandler استفاده می‌کنیم
	// در آینده می‌توانیم فیلتر ویدیو اضافه کنیم
	h.ListMediaHandler(w, r)
}

// UploadVideoHandler آپلود ویدیو
func (h *MediaHandler) UploadVideoHandler(w http.ResponseWriter, r *http.Request) {
	// فعلاً از همان UploadMediaHandler استفاده می‌کنیم
	// در آینده می‌توانیم منطق خاص ویدیو اضافه کنیم
	h.UploadMediaHandler(w, r)
}

// DownloadVideoHandler دانلود ویدیو
func (h *MediaHandler) DownloadVideoHandler(w http.ResponseWriter, r *http.Request) {
	// فعلاً از همان DownloadMediaHandler استفاده می‌کنیم
	h.DownloadMediaHandler(w, r)
}

// DeleteVideoHandler حذف ویدیو
func (h *MediaHandler) DeleteVideoHandler(w http.ResponseWriter, r *http.Request) {
	// فعلاً از همان DeleteMediaHandler استفاده می‌کنیم
	h.DeleteMediaHandler(w, r)
}

// GetVideoHandler دریافت اطلاعات ویدیو
func (h *MediaHandler) GetVideoHandler(w http.ResponseWriter, r *http.Request) {
	// فعلاً از همان GetMediaHandler استفاده می‌کنیم
	h.GetMediaHandler(w, r)
}

// UpdateVideoHandler بروزرسانی اطلاعات ویدیو
func (h *MediaHandler) UpdateVideoHandler(w http.ResponseWriter, r *http.Request) {
	// پیاده‌سازی به‌روزرسانی ویدیو
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "Video update endpoint"})
}

// scheduleVideoCompression سیستم زمان‌بندی فشرده‌سازی ویدیو
// این متد در background اجرا می‌شود و jobهای pending را در بازه زمانی مشخص شده پردازش می‌کند
func (h *MediaHandler) scheduleVideoCompression() {
	// جلوگیری از اجرای همزمان چندین scheduler
	staticSchedulerOnce.Do(func() {
		go func() {
			ticker := time.NewTicker(5 * time.Minute) // بررسی هر 5 دقیقه
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					h.processPendingVideoCompressionJobs()
				}
			}
		}()
	})
}

// processPendingVideoCompressionJobs پردازش jobهای pending فشرده‌سازی ویدیو
func (h *MediaHandler) processPendingVideoCompressionJobs() {
	// بررسی وضعیت فعال بودن فشرده‌سازی - بدون مقادیر پیش‌فرض
	if h.SettingService == nil {
		log.Printf("Video compression scheduler: SettingService is not available.")
		return
	}

	// دریافت وضعیت فعال بودن از تنظیمات
	compressionEnabled, err := h.SettingService.GetVideoCompressionEnabled(context.Background())
	if err != nil {
		log.Printf("Video compression scheduler: Error getting compression enabled status: %v", err)
		return
	}

	if !compressionEnabled {
		log.Printf("Video compression scheduler: Compression is disabled by frontend settings.")
		return
	}

	// دریافت تنظیمات زمان‌بندی - بدون مقادیر پیش‌فرض
	startHour, err := h.SettingService.GetVideoCompressionStartHour(context.Background())
	if err != nil {
		log.Printf("Video compression scheduler: Error getting start hour: %v", err)
		return
	}

	endHour, err := h.SettingService.GetVideoCompressionEndHour(context.Background())
	if err != nil {
		log.Printf("Video compression scheduler: Error getting end hour: %v", err)
		return
	}

	// بررسی اینکه آیا زمان فعلی در بازه مجاز است
	now := time.Now()
	currentHour := now.Hour()

	// بررسی بازه زمانی (شامل ساعت شروع و پایان)
	var inTimeWindow bool
	if startHour <= endHour {
		// بازه عادی (مثل 0 تا 23)
		inTimeWindow = currentHour >= startHour && currentHour <= endHour
	} else {
		// بازه شبانه‌روزی (مثل 22 تا 6)
		inTimeWindow = currentHour >= startHour || currentHour <= endHour
	}

	if !inTimeWindow {
		log.Printf("Video compression scheduler: Outside time window (%d-%d), current hour: %d", startHour, endHour, currentHour)
		return
	}

	// دریافت jobهای pending
	var pendingJobs []models.CompressionJob
	if err := h.DB.Where("status = ? AND job_type = ?", "pending", "video_compression").
		Preload("Media").Limit(5).Find(&pendingJobs).Error; err != nil {
		log.Printf("Error fetching pending video compression jobs: %v", err)
		return
	}

	if len(pendingJobs) == 0 {
		log.Printf("Video compression scheduler: No pending jobs found")
		return
	}

	log.Printf("Video compression scheduler: Processing %d pending jobs", len(pendingJobs))

	// پردازش jobها
	for _, job := range pendingJobs {
		go h.processVideoCompressionJob(&job)
	}
}

// processVideoCompressionJob پردازش یک job فشرده‌سازی ویدیو
func (h *MediaHandler) processVideoCompressionJob(job *models.CompressionJob) {
	// بروزرسانی وضعیت به processing
	h.DB.Model(job).Update("status", "processing")
	h.DB.Model(job).Update("started_at", time.Now())

	log.Printf("Starting video compression for job %d, media %d", job.ID, job.MediaID)

	// دریافت اطلاعات media
	var media models.MediaFile
	if err := h.DB.First(&media, job.MediaID).Error; err != nil {
		log.Printf("Error fetching media %d for job %d: %v", job.MediaID, job.ID, err)
		h.DB.Model(job).Updates(map[string]interface{}{
			"status":        "error",
			"error_message": "Media not found",
			"finished_at":   time.Now(),
		})
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
		log.Printf("Error: video file not found at %s", absPath)
		h.DB.Model(job).Updates(map[string]interface{}{
			"status":        "error",
			"error_message": "Video file not found",
			"finished_at":   time.Now(),
		})
		return
	}

	// تنظیمات فشرده‌سازی بر اساس job
	opts := services.SmartVideoOptions{
		UseFaceDetection:  true,
		TargetRatio:       0.6,
		MaxWidth:          1920,
		MaxHeight:         1080,
		TargetBitrateKbps: 3000,
		FrameAnalysisMode: "smart",
		WorkerCount:       4,
	}

	// تنظیم کیفیت بر اساس job.TargetQuality
	switch job.TargetQuality {
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

	// تنظیم فرمت بر اساس job.TargetFormat
	if job.TargetFormat != "" && job.TargetFormat != "mp4" {
		// فعلاً فقط MP4 پشتیبانی می‌شود
		log.Printf("Warning: Target format %s not supported, using MP4", job.TargetFormat)
	}

	// اجرای فشرده‌سازی
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	result, err := services.SmartCompressVideo(ctx, absPath, opts, func(progress int) {
		// بروزرسانی پیشرفت
		h.DB.Model(job).Update("progress", progress)
	}, func(proc *os.Process) {
		// ذخیره process برای امکان لغو
	})

	if err != nil {
		log.Printf("Video compression failed for job %d: %v", job.ID, err)
		h.DB.Model(job).Updates(map[string]interface{}{
			"status":        "error",
			"error_message": err.Error(),
			"finished_at":   time.Now(),
		})
		return
	}

	// بروزرسانی media file با اطلاعات فشرده‌سازی
	compressedSize := uint(result.SizeBytes)
	h.DB.Model(&media).Updates(map[string]interface{}{
		"compressed":       true,
		"compressed_size":  compressedSize,
		"format":           strings.TrimPrefix(filepath.Ext(result.OutputPath), "."),
		"width":            result.Width,
		"height":           result.Height,
		"bitrate_kbps":     result.BitrateKbps,
		"duration_seconds": result.DurationSeconds,
	})

	// بروزرسانی job به completed
	h.DB.Model(job).Updates(map[string]interface{}{
		"status":          "completed",
		"progress":        100,
		"finished_at":     time.Now(),
		"compressed_size": compressedSize, // اضافه کردن اندازه فایل فشرده شده
	})

	log.Printf("Video compression completed for job %d, media %d", job.ID, job.MediaID)
}

// متغیر برای جلوگیری از اجرای همزمان چندین scheduler
var staticSchedulerOnce sync.Once

// GetCompressionJobsHandler دریافت وضعیت jobهای فشرده‌سازی
func (h *MediaHandler) GetCompressionJobsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// دریافت پارامترهای فیلتر
	status := r.URL.Query().Get("status")
	jobType := r.URL.Query().Get("job_type")
	mediaID := r.URL.Query().Get("media_id")

	// ساخت کوئری
	query := h.DB.Model(&models.CompressionJob{})

	// اعمال فیلترها
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if jobType != "" {
		query = query.Where("job_type = ?", jobType)
	}
	if mediaID != "" {
		if id, err := parseUintFromString(mediaID); err == nil {
			query = query.Where("media_id = ?", id)
		}
	}

	// دریافت نتایج
	var jobs []models.CompressionJob
	if err := query.Preload("Media").Order("created_at DESC").Limit(100).Find(&jobs).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.New("DB_ERROR", "خطا در دریافت jobها", err.Error()))
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    jobs,
		"count":   len(jobs),
	})
}

// TestSchedulerHandler تست دستی سیستم زمان‌بندی فشرده‌سازی
func (h *MediaHandler) TestSchedulerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// اجرای مستقیم پردازش jobهای pending بدون بررسی زمان
	log.Printf("Manual test: Processing pending video compression jobs")
	h.processPendingVideoCompressionJobs()

	// دریافت تنظیمات زمان‌بندی
	startHour, endHour := 1, 13
	if h.SettingService != nil {
		if start, err := h.SettingService.GetVideoCompressionStartHour(context.Background()); err == nil {
			startHour = start
		}
		if end, err := h.SettingService.GetVideoCompressionEndHour(context.Background()); err == nil {
			endHour = end
		}
	}

	// بررسی زمان فعلی
	now := time.Now()
	currentHour := now.Hour()

	var inTimeWindow bool
	if startHour <= endHour {
		inTimeWindow = currentHour >= startHour && currentHour <= endHour
	} else {
		inTimeWindow = currentHour >= startHour || currentHour <= endHour
	}

	// شمارش jobهای pending
	var pendingCount int64
	h.DB.Model(&models.CompressionJob{}).Where("status = ? AND job_type = ?", "pending", "video_compression").Count(&pendingCount)

	// اجرای دستی scheduler
	if inTimeWindow {
		go h.processPendingVideoCompressionJobs()
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"current_hour":   currentHour,
			"start_hour":     startHour,
			"end_hour":       endHour,
			"in_time_window": inTimeWindow,
			"pending_jobs":   pendingCount,
			"current_time":   now.Format("15:04:05"),
			"scheduler_ran":  inTimeWindow,
		},
	})
}

// GetSystemHealthHandler دریافت اطلاعات سلامت سیستم
func (h *MediaHandler) GetSystemHealthHandler(w http.ResponseWriter, r *http.Request) {
	// دریافت اطلاعات CPU
	cpuUsage, err := getCPUUsage()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"success":false,"message":"خطا در دریافت اطلاعات CPU","error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// دریافت اطلاعات Memory
	memoryUsage, err := getMemoryUsage()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"success":false,"message":"خطا در دریافت اطلاعات Memory","error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// دریافت اطلاعات Disk
	diskUsage, err := getDiskUsage()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"success":false,"message":"خطا در دریافت اطلاعات Disk","error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"cpu":    cpuUsage,
			"memory": memoryUsage,
			"disk":   diskUsage,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getCPUUsage دریافت درصد استفاده CPU
func getCPUUsage() (float64, error) {
	// در ویندوز از WMI استفاده می‌کنیم
	cmd := exec.Command("wmic", "cpu", "get", "loadpercentage")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// پردازش خروجی
	lines := strings.Split(string(output), "\n")
	if len(lines) < 2 {
		return 0, fmt.Errorf("خروجی نامعتبر")
	}

	// خط اول عنوان است، خط دوم مقدار است
	usageStr := strings.TrimSpace(lines[1])
	usage, err := strconv.ParseFloat(usageStr, 64)
	if err != nil {
		return 0, err
	}

	return usage, nil
}

// getMemoryUsage دریافت درصد استفاده Memory
func getMemoryUsage() (float64, error) {
	// در ویندوز از WMI استفاده می‌کنیم
	cmd := exec.Command("wmic", "OS", "get", "TotalVisibleMemorySize,FreePhysicalMemory", "/Value")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// پردازش خروجی
	lines := strings.Split(string(output), "\n")
	var totalMemory, freeMemory float64

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "FreePhysicalMemory=") {
			freeStr := strings.TrimPrefix(line, "FreePhysicalMemory=")
			freeMemory, _ = strconv.ParseFloat(freeStr, 64)
		} else if strings.HasPrefix(line, "TotalVisibleMemorySize=") {
			totalStr := strings.TrimPrefix(line, "TotalVisibleMemorySize=")
			totalMemory, _ = strconv.ParseFloat(totalStr, 64)
		}
	}

	if totalMemory == 0 {
		return 0, fmt.Errorf("خطا در محاسبه حافظه")
	}

	usedMemory := totalMemory - freeMemory
	usagePercent := (usedMemory / totalMemory) * 100

	return usagePercent, nil
}

// getDiskUsage دریافت درصد استفاده Disk
func getDiskUsage() (float64, error) {
	// در ویندوز از WMI استفاده می‌کنیم
	cmd := exec.Command("wmic", "logicaldisk", "where", "DeviceID='C:'", "get", "Size,FreeSpace", "/Value")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// پردازش خروجی
	lines := strings.Split(string(output), "\n")
	var totalSpace, freeSpace float64

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "FreeSpace=") {
			freeStr := strings.TrimPrefix(line, "FreeSpace=")
			freeSpace, _ = strconv.ParseFloat(freeStr, 64)
		} else if strings.HasPrefix(line, "Size=") {
			totalStr := strings.TrimPrefix(line, "Size=")
			totalSpace, _ = strconv.ParseFloat(totalStr, 64)
		}
	}

	if totalSpace == 0 {
		return 0, fmt.Errorf("خطا در محاسبه فضای دیسک")
	}

	usedSpace := totalSpace - freeSpace
	usagePercent := (usedSpace / totalSpace) * 100

	return usagePercent, nil
}

// ProcessPendingVideoCompressionJobsHandler پردازش دستی jobهای pending فشرده‌سازی ویدیو
func (h *MediaHandler) ProcessPendingVideoCompressionJobsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// تبدیل job های error به pending
	log.Printf("Manual processing: Converting error jobs to pending")
	if err := h.DB.Model(&models.CompressionJob{}).
		Where("status = ? AND job_type = ?", "error", "video_compression").
		Updates(map[string]interface{}{
			"status":        "pending",
			"error_message": nil,
			"started_at":    nil,
			"finished_at":   nil,
		}).Error; err != nil {
		log.Printf("Error converting error jobs to pending: %v", err)
	}

	// اجرای مستقیم پردازش jobهای pending بدون بررسی زمان
	log.Printf("Manual processing: Starting pending video compression jobs")
	h.processPendingVideoCompressionJobs()

	// شمارش jobهای pending
	var pendingCount int64
	h.DB.Model(&models.CompressionJob{}).Where("status = ? AND job_type = ?", "pending", "video_compression").Count(&pendingCount)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":       true,
		"message":       "پردازش jobهای pending شروع شد",
		"pending_count": pendingCount,
	})
}

// پاک‌سازی پیام خطاهای همه jobهای فشرده‌سازی ویدیو
// این هندلر مقدار error_message را برای همه رکوردهای compression_jobs خالی می‌کند
// فقط توسط ادمین یا کاربر مجاز باید فراخوانی شود
func (h *MediaHandler) ClearCompressionJobErrorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// مقدار error_message را برای همه jobها null می‌کند
	if err := h.DB.Model(&models.CompressionJob{}).Where("error_message IS NOT NULL").Update("error_message", nil).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "خطا در پاک‌سازی پیام خطاها"})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "message": "تمام پیام‌های خطا با موفقیت پاک شدند"})
}

// ResetCompressionJobHandler ریست کردن job فشرده‌سازی خاص
func (h *MediaHandler) ResetCompressionJobHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// استخراج job ID از URL به روش ساده
	pathParts := strings.Split(r.URL.Path, "/")
	log.Printf("Full URL: %s", r.URL.Path)
	log.Printf("Path parts: %v", pathParts)

	// پیدا کردن job ID در path parts
	var jobIDStr string
	for i, part := range pathParts {
		if part == "compression-jobs" && i+1 < len(pathParts) {
			jobIDStr = pathParts[i+1]
			break
		}
	}

	log.Printf("Extracted job ID: %s", jobIDStr)

	if jobIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Job ID یافت نشد"})
		return
	}

	jobID, err := strconv.ParseUint(jobIDStr, 10, 32)
	if err != nil {
		log.Printf("Error parsing job ID: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Job ID نامعتبر"})
		return
	}

	// پیدا کردن job
	var job models.CompressionJob
	if err := h.DB.First(&job, jobID).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "Job یافت نشد"})
		return
	}

	// بررسی اینکه job قابل ریست باشد
	if job.Status != "error" && job.Status != "pending" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "این job قابل ریست نیست"})
		return
	}

	// ریست کردن job
	if err := h.DB.Model(&job).Updates(map[string]interface{}{
		"status":        "pending",
		"error_message": nil,
		"started_at":    nil,
		"finished_at":   nil,
		"progress":      0,
	}).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "error": "خطا در ریست کردن job"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Job با موفقیت ریست شد",
		"job_id":  jobID,
	})
}

// getProductImageDir مسیر مناسب برای ذخیره عکس محصول بر اساس سال/ماه و تعداد فایل‌ها را برمی‌گرداند
func getProductImageDir() (string, error) {
	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")

	baseDir := productImagesBaseDir()
	yearMonthDir := filepath.Join(baseDir, year, month)

	// اگر پوشه وجود ندارد، آن را ایجاد کن
	if err := ensureDir(yearMonthDir); err != nil {
		return "", err
	}

	// شمارش فایل‌های موجود در پوشه سال/ماه
	files, err := os.ReadDir(yearMonthDir)
	if err != nil {
		return "", err
	}

	// شمارش کل تصاویر موجود در پوشه سال/ماه (شامل پوشه‌های batch)
	totalImageCount := 0

	// ابتدا تصاویر مستقیم در پوشه سال/ماه را بشمار
	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp" || ext == ".svg" || ext == ".avif" {
				totalImageCount++
			}
		}
	}

	// سپس تصاویر موجود در پوشه‌های batch را بشمار
	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), "batch_") {
			batchDir := filepath.Join(yearMonthDir, file.Name())
			batchFiles, err := os.ReadDir(batchDir)
			if err != nil {
				log.Printf("خطا در خواندن پوشه %s: %v", file.Name(), err)
				continue
			}

			for _, batchFile := range batchFiles {
				if !batchFile.IsDir() {
					ext := strings.ToLower(filepath.Ext(batchFile.Name()))
					if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp" || ext == ".svg" || ext == ".avif" {
						totalImageCount++
					}
				}
			}
		}
	}

	// محاسبه شماره پوشه batch مناسب
	// تصاویر 1-500 در batch_1، 501-1000 در batch_2، و الی آخر
	batchNumber := (totalImageCount / 500) + 1
	folderName := fmt.Sprintf("batch_%d", batchNumber)
	batchDir := filepath.Join(yearMonthDir, folderName)

	// ایجاد پوشه batch اگر وجود ندارد
	if err := ensureDir(batchDir); err != nil {
		log.Printf("خطا در ایجاد پوشه batch_%d: %v", batchNumber, err)
		return yearMonthDir, nil
	}

	return batchDir, nil
}

// productImagesBaseDir مسیر پایه عکس‌های محصولات
func productImagesBaseDir() string {
	projectRoot := ""
	if wd, err := os.Getwd(); err == nil {
		projectRoot = filepath.Dir(wd)
	}
	return filepath.Join(projectRoot, "public", "uploads", "media", "products", "images")
}

// ensureDir ایجاد پوشه با مجوزهای مناسب
func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0o755)
}
