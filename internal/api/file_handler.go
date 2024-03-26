package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/The-Flash/tus-go/internal/uid"
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
	fid := generateFileId()
	url := absFileUrl(r, fid)
	log.Println(url)
	uploadLengthStr := r.Header.Get("Upload-Length")
	uploadLength, err := strconv.Atoi(uploadLengthStr)
	if err != nil {
		log.Println("invalid upload length", uploadLengthStr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uploadMetadata := r.Header.Get("Upload-Metadata")
	w.Header().Set("Location", url)
	log.Println(uploadLength, uploadMetadata)
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

func generateFileId() string {
	return uid.Uid()
}

func parseUploadMetadataHeader(header string) map[string]string {
	return map[string]string{}
}
