package handlers

import "net/http"

// ImageHandler is a thin wrapper around existing MediaHandler methods.
// It allows future separation without breaking current API contracts.
type ImageHandler struct {
	Base *MediaHandler
}

func NewImageHandler(base *MediaHandler) *ImageHandler {
	return &ImageHandler{Base: base}
}

// ------- passthrough wrappers --------
func (h *ImageHandler) ListMediaHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.ListMediaHandler(w, r)
}
func (h *ImageHandler) UploadMediaHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.UploadMediaHandler(w, r)
}
func (h *ImageHandler) CompressMediaHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.CompressMediaHandler(w, r)
}
func (h *ImageHandler) DownloadMediaHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.DownloadMediaHandler(w, r)
}
func (h *ImageHandler) RotateMediaHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.RotateMediaHandler(w, r)
}
func (h *ImageHandler) RevertCompressionHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.RevertCompressionHandler(w, r)
}
func (h *ImageHandler) DeleteMediaHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.DeleteMediaHandler(w, r)
}
func (h *ImageHandler) GetMediaHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.GetMediaHandler(w, r)
}
func (h *ImageHandler) UpdateMediaHandler(w http.ResponseWriter, r *http.Request) {
	h.Base.UpdateMediaHandler(w, r)
}
