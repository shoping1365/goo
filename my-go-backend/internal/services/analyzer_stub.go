//go:build !ffmpeg
// +build !ffmpeg

package services

import "fmt"

// FaceAnalyzer ساختار تحلیلگر چهره (stub version)
type FaceAnalyzer struct {
	// در نسخه stub، هیچ فیلدی وجود ندارد
}

// NewFaceAnalyzer ایجاد نمونه جدید از FaceAnalyzer (stub version)
func NewFaceAnalyzer(modelPath string) (*FaceAnalyzer, error) {
	return nil, fmt.Errorf("OpenCV support not compiled in. Use build tags 'ffmpeg' and 'opencv' to enable face detection")
}

// DetectFaces تشخیص چهره در فریم (stub version)
func (a *FaceAnalyzer) DetectFaces(bgrFrame interface{}) bool {
	return false
}

// Close آزادسازی منابع FaceAnalyzer (stub version)
func (a *FaceAnalyzer) Close() {
	// Nothing to clean up in stub version
}
