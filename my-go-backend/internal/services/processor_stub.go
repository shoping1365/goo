//go:build !ffmpeg
// +build !ffmpeg

package services

import (
	"context"
	"fmt"
)

// VideoProcessor ساختار اصلی پردازشگر ویدیو (stub version)
// این نسخه برای زمانی است که FFmpeg در دسترس نیست
type VideoProcessor struct {
	width             int
	height            int
	bitrate           int64
	fps               int
	workerCount       int
	frameAnalysisMode string
}

// NewVideoProcessor ایجاد نمونه جدید از VideoProcessor
func NewVideoProcessor(width, height int, bitrate int64, fps int) *VideoProcessor {
	return &VideoProcessor{
		width:             width,
		height:            height,
		bitrate:           bitrate,
		fps:               fps,
		workerCount:       4,
		frameAnalysisMode: "smart",
	}
}

// SetWorkerCount تنظیم تعداد ورکر برای پردازش همزمان
func (p *VideoProcessor) SetWorkerCount(count int) {
	if count > 0 && count <= 16 {
		p.workerCount = count
	}
}

// SetFrameAnalysisMode تنظیم حالت تحلیل فریم
func (p *VideoProcessor) SetFrameAnalysisMode(mode string) {
	p.frameAnalysisMode = mode
}

// Process متد اصلی پردازش ویدیو (stub)
func (p *VideoProcessor) Process(inputPath, outputPath string) error {
	return fmt.Errorf("FFmpeg support not available - VideoProcessor is in stub mode. Use build tag 'ffmpeg' to enable full functionality")
}

// NewVideoProcessorWithSettings ایجاد نمونه جدید از VideoProcessor با تنظیمات مرکزی
func NewVideoProcessorWithSettings(ctx context.Context, settingService *CompressionSettingService, width, height int, bitrate int64, fps int) (*VideoProcessor, error) {
	workerCount := 4
	if settingService != nil {
		if count, err := settingService.GetVideoCompressionWorkerCount(ctx); err == nil {
			workerCount = count
		}
	}

	return &VideoProcessor{
		width:             width,
		height:            height,
		bitrate:           bitrate,
		fps:               fps,
		workerCount:       workerCount,
		frameAnalysisMode: "smart",
	}, nil
}

// cleanup متد آزادسازی منابع (stub version)
func (p *VideoProcessor) cleanup() {
	// Nothing to clean up in stub version
}

// متدهای placeholder (stub versions)
func (p *VideoProcessor) openInputFile(inputPath interface{}) error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) setupDecoder() error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) openOutputFile(outputPath interface{}) error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) setupCodecs() error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) setupProcessors() error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) processFrames() error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) finalize() error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

// متدهای جدید فاز ۳ (stub versions)
func (p *VideoProcessor) setupOutputFormat(path string) error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) setupEncoder() error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) setupFrameConverter() error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) setupFramesAndPackets() error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) transcodeLoop(analyzer interface{}) error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) processVideoPacketWithAnalysis(analyzer interface{}) error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) processFrameWithAnalysis(analyzer interface{}) error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) setEncoderQP(qp int) error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) flushVideoEncoder() error {
	return fmt.Errorf("FFmpeg support not compiled in")
}

func (p *VideoProcessor) flushAudioEncoder() error {
	return fmt.Errorf("FFmpeg support not compiled in")
}
