//go:build ffmpeg && opencv
// +build ffmpeg,opencv

package services

/*
#cgo pkg-config: opencv4
#include <opencv2/opencv.hpp>
#include <opencv2/objdetect.hpp>
#include <opencv2/imgproc.hpp>
#include <opencv2/highgui.hpp>
#include <stdlib.h>
#include <string.h>

// تعریف توابع C برای استفاده در Go
cv_CascadeClassifier* cv_CascadeClassifier_new() {
    return new cv::CascadeClassifier();
}

int cv_CascadeClassifier_load(cv_CascadeClassifier* classifier, const char* filename) {
    return classifier->load(filename) ? 1 : 0;
}

void cv_CascadeClassifier_delete(cv_CascadeClassifier* classifier) {
    delete classifier;
}

// تابع تشخیص چهره
int cv_detect_faces(cv_CascadeClassifier* classifier,
                   unsigned char* data, int width, int height, int channels, int step) {
    cv::Mat frame(height, width, CV_8UC3, data, step);
    cv::Mat gray;
    cv::cvtColor(frame, gray, cv::COLOR_BGR2GRAY);
    cv::equalizeHist(gray, gray);

    std::vector<cv::Rect> faces;
    classifier->detectMultiScale(gray, faces, 1.1, 2, 0|CV_HAAR_SCALE_IMAGE, cv::Size(30, 30));

    return faces.size() > 0 ? 1 : 0;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// FaceAnalyzer ساختار تحلیلگر چهره و حرکت
// این ساختار علاوه بر مدل تشخیص چهره، فریم قبلی را برای تشخیص حرکت نگه می‌دارد
// اگر فقط تشخیص چهره نیاز باشد، prevFrame مقدار nil خواهد بود
// اگر تشخیص حرکت فعال باشد، prevFrame برای اختلاف‌گیری استفاده می‌شود
//
type FaceAnalyzer struct {
	classifier *C.cv_CascadeClassifier // مدل طبقه‌بند OpenCV برای تشخیص چهره
	prevFrame  unsafe.Pointer          // اشاره‌گر به فریم قبلی (cv::Mat*) برای تشخیص حرکت
}

// NewFaceAnalyzer ایجاد نمونه جدید از FaceAnalyzer
// این تابع فایل مدل .xml تشخیص چهره را بارگذاری کرده و یک نمونه آماده برمی‌گرداند
func NewFaceAnalyzer(modelPath string) (*FaceAnalyzer, error) {
	// تبدیل مسیر مدل به C string
	cModelPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cModelPath))

	// ایجاد classifier جدید
	classifier := C.cv_CascadeClassifier_new()
	if classifier == nil {
		return nil, fmt.Errorf("خطا در ایجاد cascade classifier")
	}

	// بارگذاری مدل از فایل
	ret := C.cv_CascadeClassifier_load(classifier, cModelPath)
	if ret == 0 {
		C.cv_CascadeClassifier_delete(classifier)
		return nil, fmt.Errorf("خطا در بارگذاری مدل از مسیر: %s", modelPath)
	}

	return &FaceAnalyzer{
		classifier: classifier,
	}, nil
}

// DetectFaces تشخیص چهره در فریم
// این متد داده‌های bgrFrame را در یک cv::Mat قرار داده و تابع detectMultiScale را اجرا می‌کند
func (a *FaceAnalyzer) DetectFaces(bgrFrame *C.AVFrame) bool {
	if a.classifier == nil {
		return false
	}

	// تبدیل داده‌های فریم به فرمت مناسب برای OpenCV
	// bgrFrame باید در فرمت BGR24 باشد
	data := (*C.uchar)(unsafe.Pointer(bgrFrame.data[0]))
	width := int(bgrFrame.width)
	height := int(bgrFrame.height)
	channels := 3 // BGR24
	step := int(bgrFrame.linesize[0])

	// اجرای تشخیص چهره
	ret := C.cv_detect_faces(a.classifier, data, C.int(width), C.int(height), C.int(channels), C.int(step))

	// اگر حداقل یک چهره پیدا شد، true و در غیر این صورت false برمی‌گرداند
	return ret == 1
}

// DetectMotion تشخیص حرکت بین دو فریم متوالی
// این تابع دو فریم BGR24 را دریافت می‌کند و درصد حرکت (0 تا 100) را برمی‌گرداند
// اگر prevFrame مقدار nil باشد، مقدار اولیه را ست می‌کند و 0 برمی‌گرداند
// اگر اختلاف قابل توجه باشد، مقدار بالاتر (مثلاً >10) نشان‌دهنده حرکت زیاد است
func (a *FaceAnalyzer) DetectMotion(bgrFrame *C.AVFrame) float64 {
	// اگر فریم قبلی وجود ندارد، مقدار اولیه را ست کن
	if a.prevFrame == nil {
		// تبدیل فریم فعلی به cv::Mat و ذخیره در prevFrame
		C.free(a.prevFrame)
		// تابع C برای تبدیل AVFrame به cv::Mat و کپی آن (در ادامه اضافه می‌شود)
		a.prevFrame = C.avframe_to_mat_copy(bgrFrame)
		return 0
	}
	// اختلاف فریم فعلی و قبلی را محاسبه کن
	motion := C.cv_frame_difference(a.prevFrame, bgrFrame)
	// فریم فعلی را جایگزین قبلی کن
	C.free(a.prevFrame)
	a.prevFrame = C.avframe_to_mat_copy(bgrFrame)
	return float64(motion)
}

// ResetMotionState آزادسازی و ریست وضعیت فریم قبلی
func (a *FaceAnalyzer) ResetMotionState() {
	if a.prevFrame != nil {
		C.free(a.prevFrame)
		a.prevFrame = nil
	}
}

// Close آزادسازی منابع FaceAnalyzer
func (a *FaceAnalyzer) Close() {
	if a.classifier != nil {
		C.cv_CascadeClassifier_delete(a.classifier)
		a.classifier = nil
	}
}

// --- کد C برای تبدیل AVFrame به cv::Mat و اختلاف فریم ---
C.cv::Mat* avframe_to_mat_copy(C.AVFrame* frame) {
    // فرض: ورودی BGR24 است
    int width = frame->width;
    int height = frame->height;
    int step = frame->linesize[0];
    unsigned char* data = frame->data[0];
    C.cv::Mat* mat = new C.cv::Mat(height, width, C.CV_8UC3);
    C.memcpy(mat->data, data, height * step);
    return mat;
}

// محاسبه درصد اختلاف بین دو فریم (cv::Mat*) و فریم فعلی
// خروجی: درصد حرکت (0 تا 100)
double cv_frame_difference(C.cv::Mat* prev, C.AVFrame* curr) {
    int width = curr->width;
    int height = curr->height;
    int step = curr->linesize[0];
    unsigned char* data = curr->data[0];
    C.cv::Mat currMat(height, width, C.CV_8UC3, data, step);
    C.cv::Mat diff;
    C.cv::absdiff(*prev, currMat, diff);
    C.cv::Mat gray;
    C.cv::cvtColor(diff, gray, C.cv::COLOR_BGR2GRAY);
    double motion = C.cv::sum(gray)[0] / (width * height * 255.0) * 100.0;
    return motion;
}
