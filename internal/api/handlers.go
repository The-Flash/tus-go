package handlers

import (
	"net/http"
)

// Hello Hello world handler
func Hello(w http.ResponseWriter, r *http.Request) {
	WriteJsonResponse(w, map[string]string{
		"msg": "hello world",
	}, http.StatusOK)
}
