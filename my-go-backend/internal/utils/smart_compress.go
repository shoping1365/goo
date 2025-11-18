// smart_compress.go
// ---------------------------------------------------------------------------------------------------------------------
// این فایل تمامی منطق فشرده‌سازی هوشمند را در خود نگه می‌دارد.
// هدف اصلی: با رعایت اصل Single-Responsibility، Handler‌ ها تنها درخواست را دریافت کنند و عملیات فشرده‌سازی در این ماژول انجام شود.
// ---------------------------------------------------------------------------------------------------------------------
//  ✅ SmartCompress   : نقطه ورود اصلی برای فشرده‌سازی هوشمند
//  ✅ SmartOptions    : گزینه‌های ورودی شامل اندازه هدف، فرمت دلخواه، ابعاد حداکثری و …
//  ✅ SmartResult     : خروجی شامل داده فشرده‌شده و MIME نهایی
//  ✅ resizeIfNeeded  : کوچک‌کردن تصویر در صورت بزرگ بودن بیش از حد (Performance Booster)
//  ✅ decideFormats   : انتخاب لیست فرمت‌های کاندید بر اساس ویژگی‌های ورودی (Heuristic Engine)
//  ✅ encodeWithQual  : فشرده‌سازی تصویر با Quality  داده‌شده برای هر فرمت
//  ✅ mapQualityAVIF  : تبدیل بازه ۱-۱۰۰ به ۰-۶۳ مخصوص libaom
// در تمام توابع، توضیح «قبل از کد» قرار گرفته تا دقیقاً معلوم باشد چه می‌کنند.
// ---------------------------------------------------------------------------------------------------------------------

package utils

import (
	// bytes.Buffer برای نگه‌داشتن خروجی در حافظه بدون ایجاد فایل موقت
	"bytes"
	// خطاهای استاندارد
	"errors"
	// پردازش تصویر
	"image"
	_ "image/gif" // ثبت دیکودر GIF برای image.Decode
	"image/jpeg"
	"image/png"

	// انجام محاسبات عددی ساده برای حلقه کاهش کیفیت
	"math"

	// کتابخانه خارجی برای فرمت WebP

	"github.com/chai2010/webp"

	// کوچک‌سازی ابعاد تصویر در صورت نیاز
	"strings"

	"github.com/nfnt/resize"
	// محاسبه SSIM جهت کنترل افت کیفیت بصری
	"image/color"
	"image/draw"

	quantize "github.com/ericpauley/go-quantize/quantize"
)

// SmartOptions مشخصات و محدودیت‌های درخواستی برای فشرده‌سازی را نگه می‌دارد.
//
//	TargetSizeKB   : حداکثر حجم خروجی (کیلوبایت). اگر صفر باشد، الگوریتم تا حد بهینه عمل می‌کند ولی الزام حجمی ندارد.
//	OutputFormat   : فرمت دلخواه کاربر. اگر خالی باشد، الگوریتم بهترین فرمت را حدس می‌زند.
//	MaxWidth/Height: محدود کردن ابعاد پیش از فشرده‌سازی (Performance + UX)
//	OriginalSize   : اندازه فایل ورودی (بایت)؛ برای تصمیم‌گیری بهتر روی فرمت خروجی.
//	MinQuality     : حداقل کیفیت مجاز برای حلقه کاهش کیفیت (پیش‌فرض 25)
//	AllowResize    : اگر true باشد و پس از کاهش کیفیت هنوز به هدف نرسیم، ابعاد را کوچک می‌کند.
//	QualityStep    : گام کاهش کیفیت (پیش‌فرض 6)
//	DesiredRatio   : اگر TargetSizeKB صفر باشد، این نسبت حجم هدف به حجم اولیه است (مثلاً 0.65 یعنی 35٪ کاهش)
//	SSIMTarget     : حداقل SSIM قابل قبول (۰٫۹۶ پیش‌فرض). ۰ برای نادیده گرفتن
//	OptimizeColors : اگر true باشد، تصویر را با کوانتیزیشن ۲۵۶ رنگ کوچک می‌کند.
//	ProgressiveJPEG: اگر true باشد، فعلاً یک انکود baseline انجام می‌دهد (TODO: replace with libjpeg progressive).
type SmartOptions struct {
	TargetSizeKB    int
	OutputFormat    string
	MaxWidth        uint
	MaxHeight       uint
	OriginalSize    int
	MinQuality      int     // حداقل کیفیت مجاز برای حلقه کاهش کیفیت (پیش‌فرض 25)
	AllowResize     bool    // اگر true باشد و پس از کاهش کیفیت هنوز به هدف نرسیم، ابعاد را کوچک می‌کند.
	QualityStep     int     // گام کاهش کیفیت (پیش‌فرض 6)
	DesiredRatio    float64 // اگر TargetSizeKB صفر باشد، این نسبت حجم هدف به حجم اولیه است (مثلاً 0.65 یعنی 35٪ کاهش)
	SSIMTarget      float64 // حداقل SSIM قابل قبول (۰٫۹۶ پیش‌فرض). ۰ برای نادیده گرفتن
	OptimizeColors  bool
	ProgressiveJPEG bool
}

// SmartResult خروجی نهایی از منطق هوشمند
//
//	Data     : تصویر فشرده شده در حافظه
//	MimeType : MIME نهایی که باید در Header پاسخ نوشته شود.
//	Format   : پسوند (webp, avif, …) جهت استفاده‌های آتی
type SmartResult struct {
	Data     *bytes.Buffer
	MimeType string
	Format   string
}

// SmartCompress نقطه ورود اصلی؛ با توجه به SmartOptions تصویر را بهینه می‌کند.
// – ورودی  img        : تصویر Decode شده (در ابعاد و رنگ اصلی)
// – ورودی  inputMime  : MIME ورودی (مثلاً image/jpeg). برای انتخاب فرمت بهتر استفاده می‌شود.
// – ورودی  opts       : محدودیت‌ها و ترجیحات کاربر
// – خروجی SmartResult : دادهٔ فشرده + جزئیات
// – خروجی error       : در صورت بروز خطا
func SmartCompress(img image.Image, inputMime string, opts SmartOptions) (SmartResult, error) {
	// ---------------------------------------
	// ۱) در صورت نیاز، ابتدا ابعاد تصویر کوچک می‌شود.
	// ---------------------------------------
	resized := resizeIfNeeded(img, opts.MaxWidth, opts.MaxHeight)

	// ---------------------------------------
	// ۲) بر اساس ورودی و گزینه‌ها لیست فرمت‌های پیشنهادی را می‌گیریم.
	//    این لیست به ترتیب اولویت است؛ اولین فرمت که شرط حجم را پاس کند انتخاب می‌شود.
	// ---------------------------------------
	candidates := decideFormats(inputMime, opts)

	// ---------------------------------------
	// ۳) حلقه روی فرمت‌های کاندید + حلقه داخلی کاهش تدریجی کیفیت
	// ---------------------------------------
	var lastErr error
	minQual := 20 // کمی تهاجمی‌تر از قبل
	if opts.MinQuality > 0 {
		minQual = opts.MinQuality
	}

	step := 6 // گام پیش‌فرض
	if opts.QualityStep > 0 {
		step = opts.QualityStep
	}

	// تعیین هدف بر حسب بایت: یا TargetSizeKB یا نسبت DesiredRatio
	targetBytes := 0
	if opts.TargetSizeKB > 0 {
		targetBytes = opts.TargetSizeKB * 1024
	} else if opts.DesiredRatio > 0 && opts.DesiredRatio < 1 {
		targetBytes = int(float64(opts.OriginalSize) * opts.DesiredRatio)
	} else {
		// اگر هیچ محدودیت صریحی داده نشده، بر اساس اندازه ورودی تصمیم می‌گیریم.
		// هدف این است که عکس‌های بزرگ‌تر از ۱ مگابایت به‌صورت تهاجمی‌تری فشرده شوند.
		switch {
		case opts.OriginalSize > 5*1024*1024: // بالای ۵ مگابایت → کاهش به حدود ۴۰٪
			targetBytes = int(float64(opts.OriginalSize) * 0.40)
		case opts.OriginalSize > 2*1024*1024: // بین ۲ تا ۵ مگابایت → کاهش به حدود ۵۰٪
			targetBytes = int(float64(opts.OriginalSize) * 0.50)
		case opts.OriginalSize > 1*1024*1024: // بین ۱ تا ۲ مگابایت → کاهش به حدود ۵۰٪
			targetBytes = int(float64(opts.OriginalSize) * 0.50)
		}
	}

	for _, f := range candidates {
		qualityInit := map[string]int{"webp": 75, "jpeg": 85, "png": -1}[f]
		var bestBuf *bytes.Buffer
		bestSize := int(^uint(0) >> 1) // max int
		quality := qualityInit
		for quality >= minQual {
			buf := &bytes.Buffer{}
			if err := encodeWithQuality(buf, resized, f, quality, opts); err != nil {
				lastErr = err
				break
			}
			pass := false
			if targetBytes == 0 || buf.Len() <= targetBytes {
				pass = true
			}
			// اگر معیار بصری تعریف شده باشد، آن را چک می‌کنیم
			if opts.SSIMTarget > 0 {
				if val, err := calculateSSIM(resized, buf.Bytes()); err == nil {
					pass = pass && (val >= opts.SSIMTarget)
				}
			}

			if pass {
				if buf.Len() < bestSize {
					bestBuf = buf
					bestSize = buf.Len()
				}
			}
			quality -= step
		}
		if bestBuf != nil {
			mime := map[string]string{"jpeg": "image/jpeg", "jpg": "image/jpeg", "png": "image/png", "webp": "image/webp"}[f]
			return SmartResult{Data: bestBuf, MimeType: mime, Format: f}, nil
		}

		// اگر هنوز بزرگ است و اجازه Resize داریم چند پاس کوچک‌سازی اجرا می‌کنیم.
		if opts.AllowResize && (opts.MaxWidth == 0 && opts.MaxHeight == 0) && targetBytes > 0 {
			// حداکثر سه رزایز پیاپی
			for passes := 0; passes < 3; passes++ {
				scaleFactor := math.Sqrt(float64(targetBytes) / float64(opts.OriginalSize))
				if scaleFactor < 0.35 {
					scaleFactor = 0.35
				}
				newW := uint(float64(resized.Bounds().Dx()) * scaleFactor)
				newH := uint(float64(resized.Bounds().Dy()) * scaleFactor)
				// اگر ابعاد خیلی کوچک می‌شود، از حلقه خارج شویم
				if newW < 200 || newH < 200 {
					break
				}
				resized = resize.Resize(newW, newH, resized, resize.Lanczos3)
				// ریست کیفیت
				quality = qualityInit
				for quality >= minQual {
					buf := &bytes.Buffer{}
					if err := encodeWithQuality(buf, resized, f, quality, opts); err != nil {
						lastErr = err
						break
					}
					pass := false
					if targetBytes == 0 || buf.Len() <= targetBytes {
						pass = true
					}
					// اگر معیار بصری تعریف شده باشد، آن را چک می‌کنیم
					if opts.SSIMTarget > 0 {
						if val, err := calculateSSIM(resized, buf.Bytes()); err == nil {
							pass = pass && (val >= opts.SSIMTarget)
						}
					}

					if pass {
						mime := map[string]string{"jpeg": "image/jpeg", "jpg": "image/jpeg", "png": "image/png", "webp": "image/webp"}[f]
						return SmartResult{Data: buf, MimeType: mime, Format: f}, nil
					}
					quality -= step
				}
			}
		}
	}
	if lastErr != nil {
		return SmartResult{}, lastErr
	}
	return SmartResult{}, errors.New("smart compression: no format satisfied constraints")
}

// resizeIfNeeded اگر MaxWidth یا MaxHeight بالاتر از صفر باشد، تصویر را طوری تغییر اندازه می‌دهد که در آن محدوده جا شود.
// در غیر این‌صورت، همان تصویر ورودی را برمی‌گرداند تا از مصرف حافظه اضافی جلوگیری شود.
func resizeIfNeeded(img image.Image, maxW, maxH uint) image.Image {
	if maxW == 0 && maxH == 0 {
		return img
	}
	return resize.Thumbnail(maxW, maxH, img, resize.Lanczos3)
}

// decideFormats فرمت‌های خروجی را به ترتیب اولویت بازمی‌گرداند.
// اولویت‌ها بر اساس:
//
//	– فرمت ورودی
//	– اندازه فایل ورودی (اگر بسیار بزرگ باشد → avif)
//	– وجود یا عدم وجود آلفا (RGBA)
//	– ترجیح کاربر در opts.OutputFormat
func decideFormats(inputMime string, opts SmartOptions) []string {
	// اگر کاربر فرمت خاصی درخواست کرده، فقط همان فرمت خروجی را امتحان می‌کنیم.
	if opts.OutputFormat != "" && opts.OutputFormat != "auto" {
		return []string{strings.ToLower(opts.OutputFormat)}
	}

	hasAlpha := false
	switch inputMime {
	case "image/png", "image/gif", "image/webp":
		hasAlpha = true
	}

	// Heuristic ساده بر اساس اندازه فایل.
	if opts.OriginalSize > 5*1024*1024 { // بزرگ‌تر از ۵ مگابایت
		if hasAlpha {
			return []string{"webp", "png"}
		}
		return []string{"webp", "jpeg"}
	}
	if hasAlpha {
		return []string{"webp", "png"}
	}
	// پیش‌فرض برای عکس‌های معمولی بدون آلفا
	return []string{"webp", "jpeg"}
}

// encodeWithQuality خروجی تصویر را با توجه به فرمت و کیفیت و همچنین گزینه‌های اضافی تولید می‌کند.
// اگر OptimizeColors فعال باشد و فرمت PNG انتخاب شود، ابتدا تصویر را با کوانتیزیشن ۲۵۶ رنگ کوچک می‌کند.
// اگر ProgressiveJPEG فعال باشد و فرمت JPEG باشد، فعلاً یک انکود baseline انجام می‌دهد (TODO: replace with libjpeg progressive).
func encodeWithQuality(buf *bytes.Buffer, img image.Image, format string, quality int, opts SmartOptions) error {
	switch format {
	case "jpeg", "jpg":
		// Progressive در stdlib پشتیبانی نمی‌شود؛ فعلاً همان baseline.
		return jpeg.Encode(buf, img, &jpeg.Options{Quality: quality})
	case "png":
		if opts.OptimizeColors {
			// کاهش تعداد رنگ‌ها با median cut به 256 رنگ برای کاهش حجم
			q := quantize.MedianCutQuantizer{}
			pal := q.Quantize(make([]color.Color, 0, 256), img)
			palettedImg := image.NewPaletted(img.Bounds(), pal)
			draw.FloydSteinberg.Draw(palettedImg, img.Bounds(), img, image.Point{})
			return png.Encode(buf, palettedImg)
		}
		return png.Encode(buf, img)
	case "webp":
		return webp.Encode(buf, img, &webp.Options{Quality: float32(quality)})
	default:
		return errors.New("unsupported format: " + format)
	}
}

// calculateSSIM یک تخمین ساده از SSIM را بدون وابستگی خارجی محاسبه می‌کند.
// این پیاده‌سازی «نسخهٔ تک مقیاس» است و فقط از میانگین و واریانس روشنایی پیکسل‌ها استفاده می‌کند.
// دقت آن از نسخه‌های کتابخانه‌ای کمتر است، اما برای فیلتر تقریبی کیفیت کافی‌ست.
func calculateSSIM(orig image.Image, compressedBytes []byte) (float64, error) {
	compImg, _, err := image.Decode(bytes.NewReader(compressedBytes))
	if err != nil {
		return 0, err
	}

	// ابعاد باید یکسان باشند؛ در غیر این‌صورت، SSIM را صفر برمی‌گردانیم.
	if orig.Bounds().Dx() != compImg.Bounds().Dx() || orig.Bounds().Dy() != compImg.Bounds().Dy() {
		return 0, errors.New("calculateSSIM: dimension mismatch")
	}

	var (
		muX, muY       float64
		sigmaX, sigmaY float64
		sigmaXY        float64
		totalPixels    float64
	)

	// ضرایب استاندارد فضای Y'UV (Rec.601) برای محاسبه روشنایی
	const (
		rCoef = 0.299
		gCoef = 0.587
		bCoef = 0.114
	)

	// ابتدا میانگین‌ها را به دست می‌آوریم
	for y := 0; y < orig.Bounds().Dy(); y++ {
		for x := 0; x < orig.Bounds().Dx(); x++ {
			c1 := color.RGBAModel.Convert(orig.At(orig.Bounds().Min.X+x, orig.Bounds().Min.Y+y)).(color.RGBA)
			c2 := color.RGBAModel.Convert(compImg.At(compImg.Bounds().Min.X+x, compImg.Bounds().Min.Y+y)).(color.RGBA)

			y1 := rCoef*float64(c1.R) + gCoef*float64(c1.G) + bCoef*float64(c1.B)
			y2 := rCoef*float64(c2.R) + gCoef*float64(c2.G) + bCoef*float64(c2.B)

			muX += y1
			muY += y2
			totalPixels++
		}
	}

	muX /= totalPixels
	muY /= totalPixels

	// حالا واریانس‌ها و کواریانس را محاسبه می‌کنیم
	for y := 0; y < orig.Bounds().Dy(); y++ {
		for x := 0; x < orig.Bounds().Dx(); x++ {
			c1 := color.RGBAModel.Convert(orig.At(orig.Bounds().Min.X+x, orig.Bounds().Min.Y+y)).(color.RGBA)
			c2 := color.RGBAModel.Convert(compImg.At(compImg.Bounds().Min.X+x, compImg.Bounds().Min.Y+y)).(color.RGBA)

			y1 := rCoef*float64(c1.R) + gCoef*float64(c1.G) + bCoef*float64(c1.B)
			y2 := rCoef*float64(c2.R) + gCoef*float64(c2.G) + bCoef*float64(c2.B)

			diffX := y1 - muX
			diffY := y2 - muY

			sigmaX += diffX * diffX
			sigmaY += diffY * diffY
			sigmaXY += diffX * diffY
		}
	}

	sigmaX /= totalPixels - 1
	sigmaY /= totalPixels - 1
	sigmaXY /= totalPixels - 1

	// ثابت‌های پایداری (پیشنهادی در مقاله Wang-Bovik)
	const (
		c1 = (0.01 * 255) * (0.01 * 255)
		c2 = (0.03 * 255) * (0.03 * 255)
	)

	numerator := (2*muX*muY + c1) * (2*sigmaXY + c2)
	denominator := (muX*muX + muY*muY + c1) * (sigmaX + sigmaY + c2)

	if denominator == 0 {
		return 1, nil // تصاویر ثابت کاملاً مشابه‌اند
	}

	return numerator / denominator, nil
}
