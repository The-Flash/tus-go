package handlers

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	return "/files/"
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
	uploadMetadata := r.Header.Get("Upload-Metadata")
	uploadLengthStr := r.Header.Get("Upload-Length")
	uploadLength, err := strconv.Atoi(uploadLengthStr)
	if err != nil {
		log.Println("invalid upload length", uploadLengthStr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ = uploadLength
	w.Header().Set("Location", url)
	metadata := parseUploadMetadataHeader(uploadMetadata)
	log.Println(metadata)
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
	metadata := make(map[string]string)
	for _, element := range strings.Split(header, ",") {
		element := strings.TrimSpace(element)
		parts := strings.Split(element, " ")
		if len(parts) > 2 {
			continue
		}
		key := parts[0]
		if key == "" {
			continue
		}
		value := ""
		if len(parts) == 2 {
			dec, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				continue
			}
			value = string(dec)
		}
		metadata[key] = value
	}
	return metadata
}
