package handlers

import (
	"fmt"
	"net/http"
)

type FileHandler struct {
}

// NewFileHandler handle creation of file upload
func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

// Path url to access file handler
func (f *FileHandler) Path() string {
	return "/files"
}

func (f *FileHandler) Method() string {
	return http.MethodPost
}

func (f *FileHandler) Pattern() string {
	return fmt.Sprintf("%s %s", f.Method(), f.Path())
}

// Handler handler for file upload creation
func (f *FileHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := absFileUrl(r, "generated-id")
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusCreated)
}

func absFileUrl(r *http.Request, id string) (url string) {
	host, proto := getHostAndProtocol(r)
	url = proto + "://" + host + "/files/" + id
	return
}

func getHostAndProtocol(r *http.Request) (host string, proto string) {
	if r.TLS != nil {
		proto = "https"
	} else {
		proto = "http"
	}
	host = r.Host
	if h := r.Header.Get("X-Forwarded-Host"); h != "" {
		host = h
	}
	if p := r.Header.Get("X-Forwarded-Proto"); p != "" {
		proto = p
	}
	return
}
