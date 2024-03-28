package handlers

import (
	"fmt"
	"net/http"
)

type UploadHandler struct {
}

// NewFileHandler handle creation of file upload
func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

// Path url to access file handler
func (f *UploadHandler) Path() string {
	return "/files/{fid}"
}

func (f *UploadHandler) Method() string {
	return http.MethodPatch
}

func (f *UploadHandler) Pattern() string {
	return fmt.Sprintf("%s %s", f.Method(), f.Path())
}

func (f *UploadHandler) Handler(w http.ResponseWriter, r *http.Request) {
	fid := r.PathValue("fid")
}
