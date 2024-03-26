package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Hello Hello world handler
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

// Files creates an upload resource
func Files(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "http://localhost:8000/files/23243424")
	log.Println(r.Header)
	fmt.Fprintln(w, "files")
}
