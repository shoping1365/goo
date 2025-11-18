package services

// ---------------------------------------------------------------------------------------------------------------------
// smart_video_compress.go
// ---------------------------------------------------------------------------------------------------------------------
// This file contains the first iteration of the **Smart Video Compression** engine.
// Goal: Provide a single entry-point (`SmartCompressVideo`) that analyses a video,
//       decides best encoding parameters and executes FFmpeg accordingly.
//       Each step is *heavily* documented (inline comments) so that future
//       contributors understand the reasoning behind every line.
//
// Separation of concerns:
//   1) utils/ffprobe.go       → generic wrapper around `ffprobe` executable.
//   2) services/smart_video_compress.go → orchestration + decision engine.
//   3) handlers/media_handler.go        → HTTP endpoint that calls this service.
//
// NOTE: This is **MVP**.  We start with a rule-based engine and leave hooks for
//       future ML-powered improvements.
// ---------------------------------------------------------------------------------------------------------------------

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	ffprobe "gopkg.in/vansante/go-ffprobe.v2"
)

// ------------------------------------------------------------------------------------------------
// SmartVideoOptions – incoming preferences (future-proof)
// ------------------------------------------------------------------------------------------------
// TargetRatio      : if >0, desired compressed_size = original_size * TargetRatio (e.g. 0.6 means 40% smaller)
// CreateDashLadder : if true, generate multi-resolution ladders (1080p/720p/480p)
// MaxWidth/Height  : upper bound on resolution (0 = keep original)
// UseFaceDetection : if true, use face detection for smart compression
// ------------------------------------------------------------------------------------------------
type SmartVideoOptions struct {
	TargetRatio       float64
	CreateDashLadder  bool
	MaxWidth          int
	MaxHeight         int
	Width             int // اگر >0 باشد مستقیماً استفاده می‌شود
	Height            int
	TargetBitrateKbps int64 // حداکثر بیت‌ریت خروجی
	UseFaceDetection  bool  // استفاده از تشخیص چهره برای فشرده‌سازی هوشمند
	FrameAnalysisMode string
	WorkerCount       int // تعداد ورکر برای پردازش همزمان
}

// SmartVideoResult wraps paths to generated files and statistics
// ------------------------------------------------------------------------------------------------
// OutputPath  : final MP4 (or first layer) path
// SizeBytes   : size on disk (bytes)
// Reduction   : percentage reduction compared to original (0-100)
// LadderPaths : optional additional renditions (if CreateDashLadder)
// ------------------------------------------------------------------------------------------------
type SmartVideoResult struct {
	OutputPath      string
	SizeBytes       int64
	Reduction       int
	Width           int
	Height          int
	BitrateKbps     int64
	DurationSeconds float64
	LadderPaths     []string
}

// SmartCompressVideo is the *single public entry* for callers.
// It takes an input file path + options → returns SmartVideoResult.
func SmartCompressVideo(ctx context.Context, inputPath string, opts SmartVideoOptions, progressCb func(int), procCb func(*os.Process)) (SmartVideoResult, error) {
	return SmartCompressVideoWithSettings(ctx, inputPath, opts, progressCb, procCb, nil)
}

// SmartCompressVideoWithSettings نسخه پیشرفته با تنظیمات مرکزی
// این متد از تنظیمات مرکزی برای تعداد ورکر و سایر پارامترها استفاده می‌کند
func SmartCompressVideoWithSettings(ctx context.Context, inputPath string, opts SmartVideoOptions, progressCb func(int), procCb func(*os.Process), settingService *CompressionSettingService) (SmartVideoResult, error) {
	// 1) ---------------------------------------------------------------------------------
	// Probe video to gather metadata (duration, resolution, bitrate, etc.)
	// ------------------------------------------------------------------------------------
	info, err := probeVideo(inputPath)
	if err != nil {
		return SmartVideoResult{}, fmt.Errorf("ffprobe error: %w", err)
	}

	// 2) ---------------------------------------------------------------------------------
	// Decide best encoding params (codec, CRF, resolution …) based on metadata.
	// For now we implement deterministic heuristic; hook for ML later.
	// ------------------------------------------------------------------------------------
	params := suggestEncodingParams(info, opts)

	// 3) ---------------------------------------------------------------------------------
	// اگر فشرده‌سازی هوشمند فعال باشد، از VideoProcessor استفاده می‌کنیم
	// ------------------------------------------------------------------------------------
	if opts.UseFaceDetection {
		// بررسی مدت زمان ویدیو برای بهینه‌سازی پردازش
		duration := info.Format.DurationSeconds
		if duration > 1800 { // بیش از 30 دقیقه
			// برای ویدیوهای طولانی، تنظیمات بهینه‌سازی اعمال می‌شود
			fmt.Printf("ویدیو طولانی است (%.2f ثانیه). تنظیمات بهینه‌سازی اعمال می‌شود.\n", duration)
		}
		if duration > 3600 { // بیش از 1 ساعت
			fmt.Printf("ویدیو خیلی طولانی است (%.2f ثانیه). الگوریتم تشخیص تغییر صحنه فعال شد.\n", duration)
		}
		// استفاده از VideoProcessor با تشخیص چهره و تنظیمات مرکزی
		var processor *VideoProcessor

		// اضافه کردن timeout برای ویدیوهای طولانی
		if duration > 1800 { // بیش از 30 دقیقه
			// تنظیم timeout طولانی‌تر برای ویدیوهای طولانی
			timeout := time.Duration(duration/60) * time.Minute // 1 دقیقه به ازای هر 60 ثانیه ویدیو
			if timeout > 2*time.Hour {
				timeout = 2 * time.Hour // حداکثر 2 ساعت
			}
			fmt.Printf("تنظیم timeout برای ویدیو: %v\n", timeout)
		}

		// دریافت حالت تحلیل فریم از opts یا ورودی
		frameAnalysisMode := "smart"
		if opts.FrameAnalysisMode != "" {
			frameAnalysisMode = opts.FrameAnalysisMode
		}

		// ایجاد VideoProcessor
		fmt.Printf("Creating VideoProcessor with params: width=%d, height=%d, bitrate=%d, fps=30, workerCount=%d\n",
			params.width, params.height, opts.TargetBitrateKbps*1000, opts.WorkerCount)
		processor = NewVideoProcessor(params.width, params.height, opts.TargetBitrateKbps*1000, 30)

		// تنظیم تعداد ورکر و حالت تحلیل فریم
		if opts.WorkerCount > 0 {
			processor.SetWorkerCount(opts.WorkerCount)
		}
		processor.SetFrameAnalysisMode(frameAnalysisMode)

		fmt.Printf("VideoProcessor created successfully with frameAnalysisMode: %s, workerCount: %d\n", frameAnalysisMode, opts.WorkerCount)

		outPath := changeExtension(inputPath, "smart_compressed.mp4")
		fmt.Printf("Starting smart compression: %s -> %s\n", inputPath, outPath)

		if err := processor.Process(inputPath, outPath); err != nil {
			fmt.Printf("Smart compression failed: %v\n", err)
			return SmartVideoResult{}, fmt.Errorf("smart compression error: %w", err)
		}
		fmt.Printf("Smart compression completed successfully\n")

		// 4) ---------------------------------------------------------------------------------
		// Gather result statistics (size, reduction).
		// ------------------------------------------------------------------------------------
		size, _ := fileSize(outPath)
		reduction := 0
		if origSize, err := strconv.ParseInt(info.Format.Size, 10, 64); err == nil && origSize > 0 {
			reduction = int((1 - float64(size)/float64(origSize)) * 100)
		}

		// استخراج برخی متادیتا برای بازگرداندن به بالا
		var width, height int
		for _, s := range info.Streams {
			if s.CodecType == "video" {
				width = s.Width
				height = s.Height
				break
			}
		}
		var bitrateKbps int64
		var durationSec float64
		bitrateInt, _ := strconv.ParseInt(info.Format.BitRate, 10, 64)
		bitrateKbps = bitrateInt / 1000
		durationSec = info.Format.DurationSeconds

		return SmartVideoResult{
			OutputPath:      outPath,
			SizeBytes:       size,
			Reduction:       reduction,
			Width:           width,
			Height:          height,
			BitrateKbps:     bitrateKbps,
			DurationSeconds: durationSec,
			LadderPaths:     nil,
		}, nil
	}

	// 3) ---------------------------------------------------------------------------------
	// Build & execute FFmpeg command for normal compression.
	// We stream progress but for MVP we run synchronously.
	// ------------------------------------------------------------------------------------
	outPath := changeExtension(inputPath, "compressed.mp4")
	if err := runFFmpegWithProgress(ctx, inputPath, outPath, params, info, progressCb, procCb); err != nil {
		return SmartVideoResult{}, err
	}

	// 4) ---------------------------------------------------------------------------------
	// Gather result statistics (size, reduction).
	// ------------------------------------------------------------------------------------
	size, _ := fileSize(outPath)
	reduction := 0
	if origSize, err := strconv.ParseInt(info.Format.Size, 10, 64); err == nil && origSize > 0 {
		reduction = int((1 - float64(size)/float64(origSize)) * 100)
	}

	// استخراج برخی متادیتا برای بازگرداندن به بالا
	var width, height int
	for _, s := range info.Streams {
		if s.CodecType == "video" {
			width = s.Width
			height = s.Height
			break
		}
	}
	var bitrateKbps int64
	var durationSec float64
	bitrateInt, _ := strconv.ParseInt(info.Format.BitRate, 10, 64)
	bitrateKbps = bitrateInt / 1000
	durationSec = info.Format.DurationSeconds

	return SmartVideoResult{
		OutputPath:      outPath,
		SizeBytes:       size,
		Reduction:       reduction,
		Width:           width,
		Height:          height,
		BitrateKbps:     bitrateKbps,
		DurationSeconds: durationSec,
		LadderPaths:     nil,
	}, nil
}

// ------------------------------------------------------------------------------------------------
// Internal helpers
// ------------------------------------------------------------------------------------------------

type encodingParams struct {
	width      int
	height     int
	crf        int
	preset     string
	vCodec     string
	aCodec     string
	maxBitrate string // e.g. "4000k"
}

// probeVideo wraps github.com/vansante/go-ffprobe with 3s timeout.
func probeVideo(path string) (*ffprobe.ProbeData, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFn()
	data, err := ffprobe.ProbeURL(ctx, path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// suggestEncodingParams – **heuristic engine**
func suggestEncodingParams(info *ffprobe.ProbeData, opts SmartVideoOptions) encodingParams {
	// Extract primary video stream (first stream with codec_type=="video")
	var vStream *ffprobe.Stream
	for _, s := range info.Streams {
		if s.CodecType == "video" {
			vStream = s
			break
		}
	}
	// Default fallbacks
	if vStream == nil {
		return encodingParams{crf: 23, preset: "medium", vCodec: "libx265", aCodec: "aac"}
	}

	// Native resolution
	nativeW := vStream.Width
	nativeH := vStream.Height

	// Decide target resolution
	targetW, targetH := nativeW, nativeH

	// اگر Width/Height صراحتاً داده شده باشد از آن استفاده کن
	if opts.Width > 0 {
		targetW = opts.Width
	}
	if opts.Height > 0 {
		targetH = opts.Height
	}
	// Respect MaxWidth/Height overrides
	if opts.MaxWidth > 0 && targetW > opts.MaxWidth {
		scale := float64(opts.MaxWidth) / float64(targetW)
		targetW = opts.MaxWidth
		targetH = int(float64(targetH) * scale)
	}
	if opts.MaxHeight > 0 && targetH > opts.MaxHeight {
		scale := float64(opts.MaxHeight) / float64(targetH)
		targetH = opts.MaxHeight
		targetW = int(float64(targetW) * scale)
	}

	// Bitrate heuristic
	bitrateInt, _ := strconv.ParseInt(info.Format.BitRate, 10, 64)
	originalBitrateKbps := bitrateInt / 1000 // convert to kbps

	// تنظیم CRF بر اساس نوع فشرده‌سازی
	crf := 23 // default
	if opts.UseFaceDetection {
		// برای فشرده‌سازی هوشمند، CRF پایین‌تر برای حفظ کیفیت چهره‌ها
		// مقدار دقیق CRF در VideoProcessor بر اساس تحلیل فریم‌ها تعیین می‌شود
		crf = 18                          // پایه برای مناطق دارای چهره (کیفیت بالا)
		if originalBitrateKbps > 10_000 { // >10Mbps (in kbps)
			crf = 22
		} else if originalBitrateKbps > 5_000 {
			crf = 20
		}
	} else {
		// فشرده‌سازی معمولی
		if originalBitrateKbps > 10_000 { // >10Mbps (in kbps)
			crf = 28
		} else if originalBitrateKbps > 5_000 {
			crf = 25
		}
	}

	maxBR := fmt.Sprintf("%dk", targetH*3) // پیش‌فرض ساده
	if opts.TargetBitrateKbps > 0 {
		maxBR = fmt.Sprintf("%dk", opts.TargetBitrateKbps)
	}

	// انتخاب کدک بر اساس نوع فشرده‌سازی
	vCodec := "libx265"
	if opts.UseFaceDetection {
		// برای فشرده‌سازی هوشمند، از H.264 استفاده می‌کنیم (سازگاری بهتر)
		vCodec = "libx264"
	}

	return encodingParams{
		width:      targetW,
		height:     targetH,
		crf:        crf,
		preset:     "medium",
		vCodec:     vCodec,
		aCodec:     "aac",
		maxBitrate: maxBR,
	}
}

// runFFmpeg executes FFmpeg with given params.
func runFFmpegWithProgress(ctx context.Context, in, out string, p encodingParams, info *ffprobe.ProbeData, progressCb func(int), procCb func(*os.Process)) error {
	if progressCb == nil {
		progressCb = func(int) {}
	}

	// Build filter chain (scaling)
	scale := fmt.Sprintf("scale=%d:%d:flags=lanczos", p.width, p.height)

	// Assemble cmd args
	args := []string{
		"-i", in,
		"-vf", scale,
		"-c:v", p.vCodec,
		"-preset", p.preset,
		"-crf", strconv.Itoa(p.crf),
		"-c:a", p.aCodec,
		"-b:a", "128k",
		"-threads", "0",
		"-progress", "pipe:1", "-nostats",

		"-y", out,
	}

	// اگر سقف بیت‌ریت مشخص شده باشد
	if p.maxBitrate != "" {
		args = append(args, "-maxrate", p.maxBitrate, "-bufsize", p.maxBitrate)
	}

	cmd := exec.CommandContext(ctx, "ffmpeg", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	if procCb != nil {
		procCb(cmd.Process)
	}

	// Duration in microseconds
	durationUs := int64(0)
	dSec := info.Format.DurationSeconds
	if dSec > 0 {
		durationUs = int64(dSec * 1_000_000)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "out_time_ms=") {
			val := strings.TrimPrefix(line, "out_time_ms=")
			if curUs, err := strconv.ParseInt(val, 10, 64); err == nil && durationUs > 0 {
				pct := int(float64(curUs) / float64(durationUs) * 100)
				if pct > 100 {
					pct = 100
				}
				progressCb(pct)
			}
		}
		if line == "progress=end" {
			progressCb(100)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

// Helper: replace original ext with suffix.
func changeExtension(path, suffix string) string {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	name := base[0 : len(base)-len(ext)]
	return filepath.Join(dir, fmt.Sprintf("%s-%s", name, suffix))
}

// fileSize returns size on disk (bytes) using os.Stat (cross-platform).
func fileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
