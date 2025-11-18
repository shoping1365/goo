//go:build ffmpeg
// +build ffmpeg

package services

/*
#cgo windows,ffmpeg CFLAGS: -IC:/msys64/mingw64/include -IC:/ffmpeg-7.1.1-full_build/include
#cgo windows,ffmpeg LDFLAGS: -LC:/msys64/mingw64/lib -LC:/ffmpeg-7.1.1-full_build/lib -lavutil
#cgo !windows,ffmpeg CFLAGS: -I/mingw64/include
#cgo !windows,ffmpeg LDFLAGS: -L/mingw64/lib -lavutil
#include <libavutil/frame.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

// FaceAnalyzer ساختار برای تحلیل ساده چهره و حرکت
// در حال حاضر، تشخیص چهره انجام نمی‌شود و فقط حرکت ساده اندازه‌گیری می‌شود
// برای سادگی و عدم وابستگی به OpenCV، از مقایسه فریم قبلی و فعلی استفاده می‌کنیم

type FaceAnalyzer struct {
	lastFrameData []byte
}

// NewFaceAnalyzer ایجاد نمونه جدید FaceAnalyzer
// در نسخه فعلی، cascadePath نادیده گرفته می‌شود
func NewFaceAnalyzer(cascadePath string) (*FaceAnalyzer, error) {
	// می‌توانید در آینده چک کنید که فایل cascade وجود دارد
	return &FaceAnalyzer{}, nil
}

// Close آزادسازی منابع (در این نسخه چیزی برای آزادسازی وجود ندارد)
func (fa *FaceAnalyzer) Close() { fa.lastFrameData = nil }

// DetectFaces تشخیص چهره (فعلاً پیاده‌سازی نشده، همیشه false برمی‌گرداند)
func (fa *FaceAnalyzer) DetectFaces(frame *C.AVFrame) bool { return false }

// DetectMotion محاسبه شدت حرکت بین فریم قبلی و فعلی (0..100)
func (fa *FaceAnalyzer) DetectMotion(frame *C.AVFrame) float64 {
	if frame == nil {
		return 0
	}
	width := int(frame.width)
	height := int(frame.height)
	if width == 0 || height == 0 {
		return 0
	}
	// فقط صفحه Y (در data[0]) را بررسی می‌کنیم
	lineSize := int(frame.linesize[0])
	if lineSize == 0 {
		return 0
	}
	size := height * lineSize
	dataPtr := unsafe.Pointer(frame.data[0])
	if dataPtr == nil {
		return 0
	}
	curFrame := unsafe.Slice((*byte)(dataPtr), size)

	// اگر فریم قبلی نداریم، ذخیره کرده و خروجی 0 می‌دهیم
	if fa.lastFrameData == nil || len(fa.lastFrameData) != len(curFrame) {
		fa.lastFrameData = append([]byte(nil), curFrame...)
		return 0
	}

	var diffSum int64
	for i := 0; i < len(curFrame); i++ {
		d := int64(int(curFrame[i]) - int(fa.lastFrameData[i]))
		if d < 0 {
			d = -d
		}
		diffSum += d
	}
	// به‌روز رسانی فریم قبلی
	copy(fa.lastFrameData, curFrame)

	// میانگین اختلاف پیکسل‌ها را به درصد تبدیل می‌کنیم
	avgDiff := float64(diffSum) / float64(len(curFrame))
	motion := (avgDiff / 255.0) * 100.0
	return motion
}
