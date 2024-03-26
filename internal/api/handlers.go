package handlers

import (
	"fmt"
	"net/http"
)

// Hello Hello world handler
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
