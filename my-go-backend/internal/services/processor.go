//go:build ffmpeg
// +build ffmpeg

// این فایل فقط زمانی کامپایل می‌شود که FFmpeg در دسترس باشد
// برای توسعه و تست، می‌توانید از build tag ffmpeg استفاده کنید

package services

/*
#cgo windows,ffmpeg CFLAGS: -IC:/msys64/mingw64/include -IC:/ffmpeg-7.1.1-full_build/include
#cgo windows,ffmpeg LDFLAGS: -LC:/msys64/mingw64/lib -LC:/ffmpeg-7.1.1-full_build/lib -lavformat -lavcodec -lavutil -lswscale -lswresample
#cgo !windows,ffmpeg CFLAGS: -I/mingw64/include
#cgo !windows,ffmpeg LDFLAGS: -L/mingw64/lib -lavformat -lavcodec -lavutil -lswscale -lswresample
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <libswscale/swscale.h>
#include <libswresample/swresample.h>
#include <stdlib.h>
#include <libavutil/opt.h>
#include <libavutil/error.h>
#include <errno.h>

// wrapper function برای دسترسی به stream بر اساس اندیس
static inline AVStream* av_fetch_stream(AVFormatContext *ctx, int idx) {
    return ctx->streams[idx];
}
// ثابت‌های خطا برای دسترسی در Go
const int GO_AVERROR_EOF = AVERROR_EOF;
const int GO_AVERROR_EAGAIN = AVERROR(EAGAIN);
*/
import "C"
import (
	"context"
	"fmt"
	"unsafe"
)

// VideoProcessor ساختار اصلی پردازشگر ویدیو
// این ساختار تمام متغیرهای مورد نیاز در طول فرآیند را نگهداری می‌کند
type VideoProcessor struct {
	// Context های ورودی و خروجی
	inputFormatCtx  *C.AVFormatContext
	outputFormatCtx *C.AVFormatContext

	// Stream index ها
	videoStreamIndex C.int
	audioStreamIndex C.int

	// Context های کدک
	videoDecoderCtx *C.AVCodecContext
	videoEncoderCtx *C.AVCodecContext
	audioDecoderCtx *C.AVCodecContext
	audioEncoderCtx *C.AVCodecContext

	// Context های پردازش
	scalerCtx    *C.struct_SwsContext
	resamplerCtx *C.struct_SwrContext

	// Frame ها
	inputFrame  *C.AVFrame
	outputFrame *C.AVFrame

	// Packet ها
	inputPacket  *C.AVPacket
	outputPacket *C.AVPacket

	// تنظیمات پردازش
	width             int
	height            int
	bitrate           int64
	fps               int
	workerCount       int // تعداد ورکر برای پردازش همزمان (از تنظیمات مرکزی)
	frameAnalysisMode string
}

// Process متد اصلی پردازش ویدیو
// این متد ارکستر کل عملیات فشرده‌سازی ویدیو است
func (p *VideoProcessor) Process(inputPath, outputPath string) error {
	fmt.Printf("VideoProcessor.Process started: %s -> %s\n", inputPath, outputPath)
	fmt.Printf("Processor settings: width=%d, height=%d, bitrate=%d, fps=%d, workerCount=%d, frameAnalysisMode=%s\n",
		p.width, p.height, p.bitrate, p.fps, p.workerCount, p.frameAnalysisMode)

	// فوراً defer cleanup را برای آزادسازی منابع در انتها فراخوانی می‌کنیم
	defer p.cleanup()

	// تبدیل مسیرها به C string
	cInputPath := C.CString(inputPath)
	defer C.free(unsafe.Pointer(cInputPath))
	cOutputPath := C.CString(outputPath)
	defer C.free(unsafe.Pointer(cOutputPath))

	// مرحله ۱: باز کردن فایل ورودی
	if err := p.openInputFile(cInputPath); err != nil {
		return fmt.Errorf("خطا در باز کردن فایل ورودی: %w", err)
	}

	// مرحله ۲: راه‌اندازی output format
	if err := p.setupOutputFormat(outputPath); err != nil {
		return fmt.Errorf("خطا در راه‌اندازی output format: %w", err)
	}

	// مرحله ۳: راه‌اندازی کدک ها
	if err := p.setupCodecs(); err != nil {
		return fmt.Errorf("خطا در راه‌اندازی کدک ها: %w", err)
	}

	// مرحله ۴: باز کردن فایل خروجی
	if err := p.openOutputFile(cOutputPath); err != nil {
		return fmt.Errorf("خطا در باز کردن فایل خروجی: %w", err)
	}

	// مرحله ۵: راه‌اندازی پردازش‌گرها
	if err := p.setupProcessors(); err != nil {
		return fmt.Errorf("خطا در راه‌اندازی پردازش‌گرها: %w", err)
	}

	// مرحله ۶: ساخت فریم‌ها و پکت‌های قابل استفاده مجدد
	if err := p.setupFramesAndPackets(); err != nil {
		return fmt.Errorf("خطا در ساخت فریم‌ها و پکت‌ها: %w", err)
	}

	// مرحله ۷: پردازش فریم‌ها
	if err := p.processFrames(); err != nil {
		return fmt.Errorf("خطا در پردازش فریم‌ها: %w", err)
	}

	// مرحله ۸: نهایی‌سازی
	if err := p.finalize(); err != nil {
		return fmt.Errorf("خطا در نهایی‌سازی: %w", err)
	}

	return nil
}

// cleanup متد آزادسازی منابع
// این متد تمام اشاره‌گرهای C را بررسی کرده و در صورت وجود، آنها را آزاد می‌کند
func (p *VideoProcessor) cleanup() {
	// آزادسازی Format Context ها
	if p.inputFormatCtx != nil {
		C.avformat_close_input(&p.inputFormatCtx)
		p.inputFormatCtx = nil
	}
	if p.outputFormatCtx != nil {
		C.avformat_free_context(p.outputFormatCtx)
		p.outputFormatCtx = nil
	}

	// آزادسازی Codec Context ها
	if p.videoDecoderCtx != nil {
		C.avcodec_free_context(&p.videoDecoderCtx)
		p.videoDecoderCtx = nil
	}
	if p.videoEncoderCtx != nil {
		C.avcodec_free_context(&p.videoEncoderCtx)
		p.videoEncoderCtx = nil
	}
	if p.audioDecoderCtx != nil {
		C.avcodec_free_context(&p.audioDecoderCtx)
		p.audioDecoderCtx = nil
	}
	if p.audioEncoderCtx != nil {
		C.avcodec_free_context(&p.audioEncoderCtx)
		p.audioEncoderCtx = nil
	}

	// آزادسازی پردازش‌گرها
	if p.scalerCtx != nil {
		C.sws_freeContext(p.scalerCtx)
		p.scalerCtx = nil
	}
	if p.resamplerCtx != nil {
		C.swr_free(&p.resamplerCtx)
		p.resamplerCtx = nil
	}

	// آزادسازی Frame ها
	if p.inputFrame != nil {
		C.av_frame_free(&p.inputFrame)
		p.inputFrame = nil
	}
	if p.outputFrame != nil {
		C.av_frame_free(&p.outputFrame)
		p.outputFrame = nil
	}

	// آزادسازی Packet ها
	if p.inputPacket != nil {
		C.av_packet_free(&p.inputPacket)
		p.inputPacket = nil
	}
	if p.outputPacket != nil {
		C.av_packet_free(&p.outputPacket)
		p.outputPacket = nil
	}
}

// NewVideoProcessor ایجاد نمونه جدید از VideoProcessor
func NewVideoProcessor(width, height int, bitrate int64, fps int) *VideoProcessor {
	return &VideoProcessor{
		width:             width,
		height:            height,
		bitrate:           bitrate,
		fps:               fps,
		workerCount:       4,       // مقدار پیش‌فرض
		frameAnalysisMode: "smart", // مقدار پیش‌فرض
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

// NewVideoProcessorWithSettings ایجاد نمونه جدید از VideoProcessor با تنظیمات مرکزی
// این متد تنظیمات را از سرویس تنظیمات فشرده‌سازی می‌خواند
func NewVideoProcessorWithSettings(ctx context.Context, settingService *CompressionSettingService, width, height int, bitrate int64, fps int) (*VideoProcessor, error) {
	// خواندن تعداد ورکر از تنظیمات مرکزی
	workerCount, err := settingService.GetVideoCompressionWorkerCount(ctx)
	if err != nil {
		// در صورت خطا، از مقدار پیش‌فرض استفاده می‌کنیم
		workerCount = 4
	}

	processor := &VideoProcessor{
		width:             width,
		height:            height,
		bitrate:           bitrate,
		fps:               fps,
		workerCount:       workerCount, // تعداد ورکر از تنظیمات مرکزی
		frameAnalysisMode: "smart",     // مقدار پیش‌فرض
	}

	return processor, nil
}

// متدهای placeholder برای کامپایل شدن کد
// این متدها در مراحل بعدی پیاده‌سازی خواهند شد

func (p *VideoProcessor) openInputFile(inputPath *C.char) error {
	fmt.Printf("Opening input file: %s\n", C.GoString(inputPath))

	// باز کردن فایل ورودی با avformat_open_input
	var inputCtx *C.AVFormatContext
	ret := C.avformat_open_input(&inputCtx, inputPath, nil, nil)
	if ret < 0 {
		fmt.Printf("Failed to open input file: %s, error: %d\n", C.GoString(inputPath), ret)
		return fmt.Errorf("خطا در باز کردن فایل ورودی: %d", ret)
	}
	p.inputFormatCtx = inputCtx
	fmt.Printf("Input file opened successfully\n")

	// خواندن اطلاعات stream ها
	ret = C.avformat_find_stream_info(p.inputFormatCtx, nil)
	if ret < 0 {
		return fmt.Errorf("خطا در خواندن اطلاعات stream ها: %d", ret)
	}

	// حلقه روی stream ها برای پیدا کردن اولین stream ویدیو
	p.videoStreamIndex = -1
	p.audioStreamIndex = -1

	for i := C.uint(0); i < p.inputFormatCtx.nb_streams; i++ {
		stream := C.av_fetch_stream(p.inputFormatCtx, C.int(i))
		if stream == nil {
			continue
		}

		// بررسی نوع stream
		if stream.codecpar.codec_type == C.AVMEDIA_TYPE_VIDEO && p.videoStreamIndex == -1 {
			p.videoStreamIndex = C.int(i)
		} else if stream.codecpar.codec_type == C.AVMEDIA_TYPE_AUDIO && p.audioStreamIndex == -1 {
			p.audioStreamIndex = C.int(i)
		}
	}

	// بررسی اینکه حداقل یک stream ویدیو پیدا شده است
	if p.videoStreamIndex == -1 {
		return fmt.Errorf("هیچ stream ویدیویی در فایل یافت نشد")
	}

	return nil
}

func (p *VideoProcessor) openOutputFile(outputPath *C.char) error {
	fmt.Printf("Opening output file: %s\n", C.GoString(outputPath))

	// باز کردن فایل خروجی با avio_open
	ret := C.avio_open(&p.outputFormatCtx.pb, outputPath, C.AVIO_FLAG_WRITE)
	if ret < 0 {
		fmt.Printf("Failed to open output file: %s, error: %d\n", C.GoString(outputPath), ret)
		return fmt.Errorf("خطا در باز کردن فایل خروجی: %d", ret)
	}
	fmt.Printf("Output file opened successfully\n")

	// نوشتن هدر فایل با avformat_write_header
	ret = C.avformat_write_header(p.outputFormatCtx, nil)
	if ret < 0 {
		return fmt.Errorf("خطا در نوشتن header فایل: %d", ret)
	}

	return nil
}

func (p *VideoProcessor) setupCodecs() error {
	fmt.Printf("Setting up codecs...\n")

	// راه‌اندازی decoder ویدیو
	if err := p.setupDecoder(); err != nil {
		fmt.Printf("Failed to setup video decoder: %v\n", err)
		return fmt.Errorf("خطا در راه‌اندازی decoder ویدیو: %w", err)
	}
	fmt.Printf("Video decoder setup completed\n")

	// راه‌اندازی encoder ویدیو
	if err := p.setupEncoder(); err != nil {
		fmt.Printf("Failed to setup video encoder: %v\n", err)
		return fmt.Errorf("خطا در راه‌اندازی encoder ویدیو: %w", err)
	}
	fmt.Printf("Video encoder setup completed\n")

	fmt.Printf("All codecs setup completed successfully\n")
	return nil
}

// setupDecoder راه‌اندازی decoder ویدیو
// این متد دیکودر مناسب را پیدا کرده و راه‌اندازی می‌کند
func (p *VideoProcessor) setupDecoder() error {
	fmt.Printf("Setting up video decoder...\n")

	// پیدا کردن stream ویدیو ورودی
	inputStream := C.av_fetch_stream(p.inputFormatCtx, p.videoStreamIndex)
	if inputStream == nil {
		fmt.Printf("Video input stream not found at index: %d\n", p.videoStreamIndex)
		return fmt.Errorf("stream ویدیو ورودی یافت نشد")
	}
	fmt.Printf("Video input stream found at index: %d\n", p.videoStreamIndex)

	// پیدا کردن دیکودر مناسب با avcodec_find_decoder
	decoder := C.avcodec_find_decoder(inputStream.codecpar.codec_id)
	if decoder == nil {
		return fmt.Errorf("دیکودر مناسب برای کدک ID %d یافت نشد", inputStream.codecpar.codec_id)
	}

	// ایجاد AVCodecContext جدید با avcodec_alloc_context3
	p.videoDecoderCtx = C.avcodec_alloc_context3(decoder)
	if p.videoDecoderCtx == nil {
		return fmt.Errorf("خطا در ایجاد decoder context")
	}

	// انتقال پارامترهای استریم ویدیو با avcodec_parameters_to_context
	ret := C.avcodec_parameters_to_context(p.videoDecoderCtx, inputStream.codecpar)
	if ret < 0 {
		return fmt.Errorf("خطا در انتقال پارامترهای کدک: %d", ret)
	}

	// تنظیم تعداد thread برای دیکودر
	workerCount := p.workerCount
	if workerCount <= 0 {
		workerCount = 4
	}
	p.videoDecoderCtx.thread_count = C.int(workerCount)
	fmt.Printf("FFmpeg decoder thread count set to: %d\n", workerCount)

	// باز کردن دیکودر با avcodec_open2
	ret = C.avcodec_open2(p.videoDecoderCtx, decoder, nil)
	if ret < 0 {
		return fmt.Errorf("خطا در باز کردن دیکودر: %d", ret)
	}

	return nil
}

// setupEncoder راه‌اندازی انکودر libx264
// این متد انکودر libx264 را پیدا کرده و راه‌اندازی می‌کند
func (p *VideoProcessor) setupEncoder() error {
	fmt.Printf("Setting up video encoder (libx264)...\n")

	// پیدا کردن انکودر libx264 با avcodec_find_encoder_by_name
	encoder := C.avcodec_find_encoder_by_name(C.CString("libx264"))
	if encoder == nil {
		fmt.Printf("libx264 encoder not found\n")
		return fmt.Errorf("انکودر libx264 یافت نشد")
	}
	fmt.Printf("libx264 encoder found successfully\n")

	// ایجاد یک استریم جدید با avformat_new_stream در outputFormatCtx
	videoStream := C.avformat_new_stream(p.outputFormatCtx, encoder)
	if videoStream == nil {
		return fmt.Errorf("خطا در ایجاد video stream")
	}

	// ایجاد یک AVCodecContext جدید برای انکودر
	p.videoEncoderCtx = C.avcodec_alloc_context3(encoder)
	if p.videoEncoderCtx == nil {
		return fmt.Errorf("خطا در ایجاد encoder context")
	}

	// کپی پارامترهای ضروری از decoderCtx
	p.videoEncoderCtx.width = C.int(p.width)
	p.videoEncoderCtx.height = C.int(p.height)
	p.videoEncoderCtx.bit_rate = C.int64_t(p.bitrate)
	p.videoEncoderCtx.framerate = C.AVRational{num: C.int(p.fps), den: 1}
	p.videoEncoderCtx.time_base = C.AVRational{num: 1, den: C.int(p.fps)}
	p.videoEncoderCtx.pix_fmt = C.AV_PIX_FMT_YUV420P
	p.videoEncoderCtx.gop_size = 12    // GOP size
	p.videoEncoderCtx.max_b_frames = 2 // تعداد B-frames

	// تنظیم تعداد thread برای FFmpeg
	workerCount := p.workerCount
	if workerCount <= 0 {
		workerCount = 4
	}
	p.videoEncoderCtx.thread_count = C.int(workerCount)
	fmt.Printf("FFmpeg encoder thread count set to: %d\n", workerCount)

	// باز کردن انکودر با avcodec_open2
	ret := C.avcodec_open2(p.videoEncoderCtx, encoder, nil)
	if ret < 0 {
		return fmt.Errorf("خطا در باز کردن encoder: %d", ret)
	}

	// کپی پارامترهای کدک به stream
	ret = C.avcodec_parameters_from_context(videoStream.codecpar, p.videoEncoderCtx)
	if ret < 0 {
		return fmt.Errorf("خطا در کپی پارامترهای کدک: %d", ret)
	}

	// تنظیم time_base برای stream
	videoStream.time_base = p.videoEncoderCtx.time_base

	return nil
}

func (p *VideoProcessor) setupProcessors() error {
	fmt.Printf("Setting up processors...\n")

	// راه‌اندازی مبدل فریم برای تحلیل هوشمند
	if err := p.setupFrameConverter(); err != nil {
		fmt.Printf("Failed to setup frame converter: %v\n", err)
		return fmt.Errorf("خطا در راه‌اندازی frame converter: %w", err)
	}
	fmt.Printf("Frame converter setup completed\n")

	// راه‌اندازی پردازش‌گر تصویر (scaler) برای تغییر اندازه
	if err := p.setupScaler(); err != nil {
		fmt.Printf("Failed to setup scaler: %v\n", err)
		return fmt.Errorf("خطا در راه‌اندازی scaler: %w", err)
	}
	fmt.Printf("Scaler setup completed\n")

	fmt.Printf("All processors setup completed successfully\n")
	return nil
}

// setupScaler راه‌اندازی پردازش‌گر تصویر
func (p *VideoProcessor) setupScaler() error {
	// پیدا کردن stream ویدیو ورودی
	inputStream := C.av_fetch_stream(p.inputFormatCtx, p.videoStreamIndex)
	if inputStream == nil {
		return fmt.Errorf("stream ویدیو ورودی یافت نشد")
	}

	// ایجاد scaler context
	p.scalerCtx = C.sws_getContext(
		inputStream.codecpar.width,                        // عرض ورودی
		inputStream.codecpar.height,                       // ارتفاع ورودی
		C.enum_AVPixelFormat(inputStream.codecpar.format), // فرمت پیکسل ورودی
		C.int(p.width),                                    // عرض خروجی
		C.int(p.height),                                   // ارتفاع خروجی
		C.AV_PIX_FMT_YUV420P,                              // فرمت پیکسل خروجی
		C.SWS_BILINEAR,                                    // الگوریتم scaling
		nil, nil, nil,
	)

	if p.scalerCtx == nil {
		return fmt.Errorf("خطا در ایجاد scaler context")
	}

	return nil
}

// processFrames پردازش تمام فریم‌های ویدیو
func (p *VideoProcessor) processFrames() error {
	fmt.Printf("Starting frame processing...\n")

	// ایجاد تحلیلگر چهره (اگر OpenCV نصب باشد)
	analyzer, err := NewFaceAnalyzer("haarcascade_frontalface_default.xml")
	if err != nil {
		// اگر تحلیلگر چهره در دسترس نباشد، بدون تحلیل ادامه می‌دهیم
		fmt.Printf("Face analyzer not available: %v, continuing without face detection\n", err)
		analyzer = nil
	} else {
		fmt.Printf("Face analyzer initialized successfully\n")
		defer analyzer.Close()
	}

	// اجرای حلقه اصلی پردازش
	fmt.Printf("Starting transcode loop...\n")
	if err := p.transcodeLoop(analyzer); err != nil {
		fmt.Printf("Transcode loop failed: %v\n", err)
		return fmt.Errorf("خطا در حلقه پردازش: %w", err)
	}
	fmt.Printf("Frame processing completed successfully\n")

	return nil
}

// processVideoPacket پردازش packet ویدیو
func (p *VideoProcessor) processVideoPacket() error {
	// ارسال packet به decoder
	ret := C.avcodec_send_packet(p.videoDecoderCtx, p.inputPacket)
	if ret < 0 {
		return fmt.Errorf("خطا در ارسال packet به decoder: %d", ret)
	}

	// دریافت فریم‌ها از decoder
	for ret >= 0 {
		ret = C.avcodec_receive_frame(p.videoDecoderCtx, p.inputFrame)
		if ret == C.GO_AVERROR_EOF || ret == C.GO_AVERROR_EAGAIN {
			break
		} else if ret < 0 {
			return fmt.Errorf("خطا در دریافت frame از decoder: %d", ret)
		}

		// پردازش فریم
		if err := p.processVideoFrame(); err != nil {
			return fmt.Errorf("خطا در پردازش video frame: %w", err)
		}
	}

	return nil
}

// processVideoFrame پردازش فریم ویدیو
func (p *VideoProcessor) processVideoFrame() error {
	// تنظیم پارامترهای output frame
	p.outputFrame.width = C.int(p.width)
	p.outputFrame.height = C.int(p.height)
	p.outputFrame.format = C.int(C.AV_PIX_FMT_YUV420P)

	// تخصیص buffer برای output frame
	ret := C.av_frame_get_buffer(p.outputFrame, 0)
	if ret < 0 {
		return fmt.Errorf("خطا در تخصیص buffer برای output frame: %d", ret)
	}

	// اطمینان از writable بودن frame
	ret = C.av_frame_make_writable(p.outputFrame)
	if ret < 0 {
		return fmt.Errorf("خطا در writable کردن output frame: %d", ret)
	}

	// تبدیل اندازه و فرمت پیکسل
	C.sws_scale(
		p.scalerCtx,
		(**C.uint8_t)(unsafe.Pointer(&p.inputFrame.data[0])),
		(*C.int)(unsafe.Pointer(&p.inputFrame.linesize[0])),
		0,
		p.inputFrame.height,
		(**C.uint8_t)(unsafe.Pointer(&p.outputFrame.data[0])),
		(*C.int)(unsafe.Pointer(&p.outputFrame.linesize[0])),
	)

	// تنظیم timestamp
	p.outputFrame.pts = p.inputFrame.pts

	// ارسال frame به encoder
	ret = C.avcodec_send_frame(p.videoEncoderCtx, p.outputFrame)
	if ret < 0 {
		return fmt.Errorf("خطا در ارسال frame به encoder: %d", ret)
	}

	// دریافت packet از encoder
	for ret >= 0 {
		ret = C.avcodec_receive_packet(p.videoEncoderCtx, p.outputPacket)
		if ret == C.GO_AVERROR_EOF || ret == C.GO_AVERROR_EAGAIN {
			break
		} else if ret < 0 {
			return fmt.Errorf("خطا در دریافت packet از encoder: %d", ret)
		}

		// تنظیم stream index
		p.outputPacket.stream_index = 0

		// تنظیم timestamp packet (encoder خودش timestamp را تنظیم می‌کند)
		// p.outputPacket.pts و p.outputPacket.dts توسط encoder تنظیم می‌شوند

		// نوشتن packet به فایل خروجی
		ret = C.av_interleaved_write_frame(p.outputFormatCtx, p.outputPacket)
		if ret < 0 {
			return fmt.Errorf("خطا در نوشتن video packet: %d", ret)
		}
	}

	return nil
}

// flushVideoEncoder ارسال فریم‌های باقی‌مانده در video encoder
func (p *VideoProcessor) flushVideoEncoder() error {
	// ارسال NULL frame برای flush کردن encoder
	ret := C.avcodec_send_frame(p.videoEncoderCtx, nil)
	if ret < 0 {
		return fmt.Errorf("خطا در flush کردن video encoder: %d", ret)
	}

	// دریافت packet های باقی‌مانده
	for ret >= 0 {
		ret = C.avcodec_receive_packet(p.videoEncoderCtx, p.outputPacket)
		if ret == C.GO_AVERROR_EOF {
			break
		} else if ret < 0 {
			return fmt.Errorf("خطا در دریافت packet از video encoder: %d", ret)
		}

		// تنظیم stream index
		p.outputPacket.stream_index = 0

		// تنظیم timestamp packet (encoder خودش timestamp را تنظیم می‌کند)
		// p.outputPacket.pts و p.outputPacket.dts توسط encoder تنظیم می‌شوند

		// نوشتن packet به فایل خروجی
		ret = C.av_interleaved_write_frame(p.outputFormatCtx, p.outputPacket)
		if ret < 0 {
			return fmt.Errorf("خطا در نوشتن video packet: %d", ret)
		}
	}

	return nil
}

// finalize نهایی‌سازی فایل خروجی
func (p *VideoProcessor) finalize() error {
	fmt.Printf("Finalizing output file...\n")

	// خالی کردن انکودر ویدیو
	if err := p.flushVideoEncoder(); err != nil {
		fmt.Printf("Failed to flush video encoder: %v\n", err)
		return fmt.Errorf("خطا در خالی کردن انکودر ویدیو: %w", err)
	}
	fmt.Printf("Video encoder flushed successfully\n")

	// نوشتن trailer فایل خروجی
	ret := C.av_write_trailer(p.outputFormatCtx)
	if ret < 0 {
		fmt.Printf("Failed to write trailer: %d\n", ret)
		return fmt.Errorf("خطا در نوشتن trailer: %d", ret)
	}
	fmt.Printf("Trailer written successfully\n")

	fmt.Printf("Finalization completed successfully\n")
	return nil
}

// setupOutputStreams تنظیم stream های خروجی
func (p *VideoProcessor) setupOutputStreams() error {
	// تنظیم video stream
	videoStream := C.av_fetch_stream(p.outputFormatCtx, 0)
	if videoStream == nil {
		return fmt.Errorf("video stream خروجی یافت نشد")
	}

	// کپی پارامترهای کدک ویدیو به stream
	ret := C.avcodec_parameters_from_context(videoStream.codecpar, p.videoEncoderCtx)
	if ret < 0 {
		return fmt.Errorf("خطا در کپی پارامترهای video codec: %d", ret)
	}

	// تنظیم time base برای video stream
	videoStream.time_base = p.videoEncoderCtx.time_base

	// کپی مستقیم stream صدا (اگر وجود داشته باشد)
	if p.audioStreamIndex >= 0 {
		fmt.Printf("Audio stream found, copying without processing...\n")

		// پیدا کردن stream صدا ورودی
		inputAudioStream := C.av_fetch_stream(p.inputFormatCtx, p.audioStreamIndex)
		if inputAudioStream == nil {
			return fmt.Errorf("stream صدا ورودی یافت نشد")
		}

		// ایجاد stream صدا جدید در فایل خروجی
		outputAudioStream := C.avformat_new_stream(p.outputFormatCtx, nil)
		if outputAudioStream == nil {
			return fmt.Errorf("خطا در ایجاد audio stream خروجی")
		}

		// کپی مستقیم پارامترهای کدک صدا
		ret = C.avcodec_parameters_copy(outputAudioStream.codecpar, inputAudioStream.codecpar)
		if ret < 0 {
			return fmt.Errorf("خطا در کپی پارامترهای audio codec: %d", ret)
		}

		// کپی time base از stream ورودی
		outputAudioStream.time_base = inputAudioStream.time_base

		fmt.Printf("Audio stream copied successfully\n")
	} else {
		fmt.Printf("No audio stream found in input file\n")
	}

	return nil
}

// setupOutputFormat ایجاد output format context
// این متد با avformat_alloc_output_context2 یک AVFormatContext برای فایل خروجی می‌سازد
func (p *VideoProcessor) setupOutputFormat(path string) error {
	fmt.Printf("Setting up output format for: %s\n", path)

	// تبدیل مسیر به C string
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	// ایجاد output format context با avformat_alloc_output_context2
	var outputCtx *C.AVFormatContext
	ret := C.avformat_alloc_output_context2(&outputCtx, nil, nil, cPath)
	if ret < 0 {
		fmt.Printf("Failed to create output format context: %d\n", ret)
		return fmt.Errorf("خطا در ایجاد output format context: %d", ret)
	}

	// ذخیره AVFormatContext در p.outputFormatCtx
	p.outputFormatCtx = outputCtx
	fmt.Printf("Output format context created successfully\n")

	return nil
}

// setupFrameConverter راه‌اندازی مبدل فریم برای تحلیل هوشمند
// این متد با sws_getContext یک SwsContext می‌سازد که فریم‌ها را به فرمت AV_PIX_FMT_BGR24 تبدیل می‌کند
func (p *VideoProcessor) setupFrameConverter() error {
	// ایجاد SwsContext برای تبدیل به فرمت BGR24
	scalerCtx := C.sws_getContext(
		C.int(p.width),            // عرض منبع
		C.int(p.height),           // ارتفاع منبع
		p.videoDecoderCtx.pix_fmt, // فرمت پیکسل منبع
		C.int(p.width),            // عرض مقصد
		C.int(p.height),           // ارتفاع مقصد
		C.AV_PIX_FMT_BGR24,        // فرمت پیکسل مقصد (BGR24 برای OpenCV)
		C.SWS_BILINEAR,            // الگوریتم تبدیل
		nil,                       // srcFilter
		nil,                       // dstFilter
		nil,                       // param
	)
	if scalerCtx == nil {
		return fmt.Errorf("خطا در ایجاد frame converter context")
	}

	// ذخیره context در p.scalerCtx
	p.scalerCtx = scalerCtx

	return nil
}

// setupFramesAndPackets ساخت فریم‌ها و پکت‌های قابل استفاده مجدد
func (p *VideoProcessor) setupFramesAndPackets() error {
	fmt.Printf("Setting up frames and packets...\n")

	// تنظیم فریم‌ها و پکت‌های قابل استفاده مجدد
	p.inputFrame = C.av_frame_alloc()
	if p.inputFrame == nil {
		fmt.Printf("Failed to allocate input frame\n")
		return fmt.Errorf("خطا در ایجاد input frame")
	}
	fmt.Printf("Input frame allocated successfully\n")

	p.outputFrame = C.av_frame_alloc()
	if p.outputFrame == nil {
		fmt.Printf("Failed to allocate output frame\n")
		return fmt.Errorf("خطا در ایجاد output frame")
	}
	fmt.Printf("Output frame allocated successfully\n")

	p.inputPacket = C.av_packet_alloc()
	if p.inputPacket == nil {
		fmt.Printf("Failed to allocate input packet\n")
		return fmt.Errorf("خطا در ایجاد input packet")
	}
	fmt.Printf("Input packet allocated successfully\n")

	p.outputPacket = C.av_packet_alloc()
	if p.outputPacket == nil {
		fmt.Printf("Failed to allocate output packet\n")
		return fmt.Errorf("خطا در ایجاد output packet")
	}
	fmt.Printf("Output packet allocated successfully\n")

	fmt.Printf("All frames and packets setup completed successfully\n")
	return nil
}

// transcodeLoop حلقه اصلی پردازش و انکود
// این متد حلقه اصلی برنامه خواهد بود که فریم‌ها را پردازش و فشرده می‌کند
func (p *VideoProcessor) transcodeLoop(analyzer *FaceAnalyzer) error {
	// خواندن پکت ها از فایل ورودی
	for {
		// خواندن پکت بعدی از ورودی
		ret := C.av_read_frame(p.inputFormatCtx, p.inputPacket)
		if ret == C.GO_AVERROR_EOF {
			break
		} else if ret < 0 {
			return fmt.Errorf("خطا در خواندن پکت: %d", ret)
		}

		// پردازش پکت‌های ویدیو
		if p.inputPacket.stream_index == p.videoStreamIndex {
			if err := p.processVideoPacketWithAnalysis(analyzer); err != nil {
				return fmt.Errorf("خطا در پردازش پکت ویدیو: %w", err)
			}
		} else if p.inputPacket.stream_index == p.audioStreamIndex {
			// کپی مستقیم پکت‌های صدا بدون پردازش
			if err := p.copyAudioPacket(); err != nil {
				return fmt.Errorf("خطا در کپی پکت صدا: %w", err)
			}
		}

		// آزادسازی پکت
		C.av_packet_unref(p.inputPacket)
	}

	return nil
}

// processVideoPacketWithAnalysis پردازش هوشمند پکت ویدیو با تحلیل چهره (موازی)
func (p *VideoProcessor) processVideoPacketWithAnalysis(analyzer *FaceAnalyzer) error {
	fmt.Printf("Processing video packet with analysis, workerCount: %d\n", p.workerCount)

	// تعیین حالت تحلیل فریم بر اساس تنظیمات
	frameAnalysisMode := p.frameAnalysisMode // مقدار از ساختار خوانده می‌شود
	workerCount := p.workerCount
	if workerCount <= 0 {
		workerCount = 4 // مقدار پیش‌فرض
	}

	// مقداردهی اولیه
	frameAnalysisInterval := 1
	useSceneDetection := false
	if frameAnalysisMode == "all" {
		// تحلیل همه فریم‌ها
		frameAnalysisInterval = 1
		useSceneDetection = false
	} else if frameAnalysisMode == "scene" {
		// فقط با تشخیص تغییر صحنه
		frameAnalysisInterval = 5
		useSceneDetection = true
	} else if frameAnalysisMode == "interval" {
		// هر N فریم تحلیل شود (N=5)
		frameAnalysisInterval = 5
		useSceneDetection = false
	} else {
		// حالت پیش‌فرض (smart): هر 5 فریم + تشخیص صحنه
		frameAnalysisInterval = 5
		useSceneDetection = true
	}

	fmt.Printf("Frame analysis mode: %s, interval: %d, scene detection: %v\n",
		frameAnalysisMode, frameAnalysisInterval, useSceneDetection)

	// شمارنده و متغیرهای کمکی
	frameCounter := 0
	lastSceneChange := 0
	sceneChangeThreshold := 30
	lastFaceResult := false

	// ارسال packet به دیکودر
	ret := C.avcodec_send_packet(p.videoDecoderCtx, p.inputPacket)
	if ret < 0 {
		return fmt.Errorf("خطا در ارسال پکت به دیکودر: %d", ret)
	}

	// دریافت فریم‌ها از دیکودر و پردازش
	for ret >= 0 {
		ret = C.avcodec_receive_frame(p.videoDecoderCtx, p.inputFrame)
		if ret == C.GO_AVERROR_EOF || ret == C.GO_AVERROR_EAGAIN {
			break
		} else if ret < 0 {
			return fmt.Errorf("خطا در دریافت فریم از دیکودر: %d", ret)
		}

		// منطق تحلیل فریم بر اساس حالت انتخابی
		var hasFace bool
		var motionValue float64
		shouldAnalyze := frameCounter%frameAnalysisInterval == 0
		if useSceneDetection && frameCounter-lastSceneChange > sceneChangeThreshold {
			shouldAnalyze = true
			lastSceneChange = frameCounter
		}

		if shouldAnalyze && analyzer != nil {
			// تحلیل چهره در thread جداگانه
			hasFace = analyzer.DetectFaces(p.inputFrame)
			// تحلیل حرکت (Motion Detection)
			motionValue = analyzer.DetectMotion(p.inputFrame)
			lastFaceResult = hasFace
			if frameCounter%30 == 0 { // لاگ هر 30 فریم
				fmt.Printf("Frame %d: Face detected: %v, Motion: %.2f\n", frameCounter, hasFace, motionValue)
			}
		} else {
			hasFace = lastFaceResult
			motionValue = 0
		}
		frameCounter++

		// ارسال نتایج تحلیل چهره و حرکت به منطق فشرده‌سازی
		if err := p.processFrameWithAnalysisResult(p.inputFrame, hasFace, motionValue); err != nil {
			return fmt.Errorf("خطا در پردازش فریم: %w", err)
		}
	}

	return nil
}

// copyAudioPacket کپی مستقیم پکت صدا بدون پردازش
func (p *VideoProcessor) copyAudioPacket() error {
	// تنظیم stream index برای پکت خروجی (صدا معمولاً stream index 1 است)
	p.outputPacket.stream_index = 1

	// کپی مستقیم داده‌های پکت
	ret := C.av_packet_ref(p.outputPacket, p.inputPacket)
	if ret < 0 {
		return fmt.Errorf("خطا در کپی پکت صدا: %d", ret)
	}

	// نوشتن پکت در فایل خروجی
	ret = C.av_interleaved_write_frame(p.outputFormatCtx, p.outputPacket)
	if ret < 0 {
		return fmt.Errorf("خطا در نوشتن پکت صدا: %d", ret)
	}

	return nil
}

// processFrameWithAnalysisResult پردازش فریم با نتیجه تحلیل چهره و حرکت
// این تابع بر اساس وجود چهره و شدت حرکت، QP یا CRF را به صورت هوشمند انتخاب می‌کند
func (p *VideoProcessor) processFrameWithAnalysisResult(inputFrame *C.AVFrame, hasFace bool, motionValue float64) error {
	// انتخاب پارامترهای فشرده‌سازی بر اساس تحلیل چهره و حرکت
	bitrate, outWidth, outHeight, codec, preset, crf := selectCompressionParams(hasFace, motionValue, p.width, p.height, p.bitrate)

	// تنظیم پارامترهای انکودر بر اساس انتخاب‌ها
	p.videoEncoderCtx.width = C.int(outWidth)
	p.videoEncoderCtx.height = C.int(outHeight)
	p.videoEncoderCtx.bit_rate = C.int64_t(bitrate)
	p.videoEncoderCtx.framerate = C.AVRational{num: C.int(p.fps), den: 1}
	p.videoEncoderCtx.time_base = C.AVRational{num: 1, den: C.int(p.fps)}
	p.videoEncoderCtx.pix_fmt = C.AV_PIX_FMT_YUV420P
	p.videoEncoderCtx.gop_size = 12    // GOP size
	p.videoEncoderCtx.max_b_frames = 2 // تعداد B-frames

	encName := C.CString(codec)
	defer C.free(unsafe.Pointer(encName))
	encPtr := C.avcodec_find_encoder_by_name(encName)
	if encPtr == nil {
		return fmt.Errorf("انکودر %s پیدا نشد", codec)
	}
	p.videoEncoderCtx.codec_id = encPtr.id

	// تنظیم پارامترهای فشرده‌سازی برای فریم فعلی
	presetC := C.CString(preset)
	defer C.free(unsafe.Pointer(presetC))
	ret := C.av_opt_set(p.videoEncoderCtx.priv_data, C.CString("preset"), presetC, 0)
	if ret < 0 {
		fmt.Printf("خطا در تنظیم preset، ادامه با تنظیم QP/CRF: %d\n", ret)
	}

	ret = C.av_opt_set(p.videoEncoderCtx.priv_data, C.CString("qp"), C.CString(fmt.Sprintf("%d", crf)), 0)
	if ret < 0 {
		// اگر تنظیم QP مستقیم کار نکرد، از CRF استفاده می‌کنیم
		// تبدیل QP به CRF (تقریبی)
		crf := 23 - (crf - 23) // تبدیل تقریبی QP به CRF
		if crf < 0 {
			crf = 0
		} else if crf > 51 {
			crf = 51
		}

		crfStr := C.CString(fmt.Sprintf("%d", crf))
		defer C.free(unsafe.Pointer(crfStr))

		ret = C.av_opt_set(p.videoEncoderCtx.priv_data, C.CString("crf"), crfStr, 0)
		if ret < 0 {
			return fmt.Errorf("خطا در تنظیم QP/CRF: %d", ret)
		}
	}

	// ارسال فریم به انکودر
	ret = C.avcodec_send_frame(p.videoEncoderCtx, inputFrame)
	if ret < 0 {
		return fmt.Errorf("خطا در ارسال فریم به انکودر: %d", ret)
	}

	// دریافت پکت‌های فشرده شده از انکودر
	for ret >= 0 {
		ret = C.avcodec_receive_packet(p.videoEncoderCtx, p.outputPacket)
		if ret == C.GO_AVERROR_EOF || ret == C.GO_AVERROR_EAGAIN {
			break
		} else if ret < 0 {
			return fmt.Errorf("خطا در دریافت پکت از انکودر: %d", ret)
		}

		// تنظیم stream index
		p.outputPacket.stream_index = 0

		// تنظیم timestamp پکت
		p.outputPacket.pts = C.av_rescale_q(inputFrame.pts, p.videoDecoderCtx.time_base, p.videoEncoderCtx.time_base)
		p.outputPacket.dts = p.outputPacket.pts

		// نوشتن پکت در فایل خروجی
		ret = C.av_interleaved_write_frame(p.outputFormatCtx, p.outputPacket)
		if ret < 0 {
			return fmt.Errorf("خطا در نوشتن پکت: %d", ret)
		}
	}

	return nil
}

// selectCompressionParams انتخاب هوشمند پارامترهای فشرده‌سازی بر اساس تحلیل چهره و حرکت و رزولوشن ورودی
// ورودی: وجود چهره، شدت حرکت، رزولوشن و bitrate ورودی
// خروجی: bitrate، رزولوشن، codec، preset، crf
func selectCompressionParams(hasFace bool, motionValue float64, inputWidth, inputHeight int, inputBitrate int64) (bitrate int64, outWidth, outHeight int, codec, preset string, crf int) {
	// مقدار پایه
	bitrate = inputBitrate
	outWidth, outHeight = inputWidth, inputHeight
	codec = "libx264"
	preset = "medium"
	crf = 23

	if hasFace && motionValue > 10 {
		bitrate = int64(float64(inputBitrate) * 1.3)
		crf = 18
		preset = "medium"
	} else if hasFace {
		bitrate = int64(float64(inputBitrate) * 1.15)
		crf = 20
		preset = "medium"
	} else if motionValue > 10 {
		bitrate = int64(float64(inputBitrate) * 1.15)
		crf = 22
		preset = "veryfast"
	} else {
		bitrate = int64(float64(inputBitrate) * 0.7)
		crf = 28
		preset = "fast"
	}

	// منطق کاهش رزولوشن اگر ویدیو ساده باشد
	if !hasFace && motionValue < 5 && (inputWidth >= 1280 && inputHeight >= 720) {
		// اگر ویدیو ساده و رزولوشن بالا است، کاهش به 720p
		outWidth = 1280
		outHeight = 720
	}

	// اگر چهره یا متن وجود دارد، رزولوشن حفظ می‌شود
	// (در آینده: اگر OCR فعال شد، شرط hasText هم اضافه شود)

	// در صورت نیاز به codec دیگر (مثلاً h265)، می‌توان منطق را توسعه داد
	return
}
