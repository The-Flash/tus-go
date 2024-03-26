package handlers

import (
	"net/http"
)

type RouteHandler interface {
	Path() string
	Handler(w http.ResponseWriter, r *http.Request)
	Method() string
	Pattern() string
}
