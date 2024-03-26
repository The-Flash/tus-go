package main

import (
	"log"
	"net/http"

	handlers "github.com/The-Flash/tus-go/internal/api"
	"github.com/The-Flash/tus-go/internal/middleware"
)

func main() {
	port := ":8000"
	middlewareStack := middleware.MiddlewareStack(
		middleware.EnableCors,
	)
	router := http.NewServeMux()
	fileHandler := handlers.NewFileHandler()
	router.HandleFunc(fileHandler.Pattern(), fileHandler.Handler)
	log.Println("server started on port " + port)
	err := http.ListenAndServe(port, middlewareStack(router))
	if err != nil {
		log.Fatal(err)
	}
}
