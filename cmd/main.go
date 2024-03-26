package main

import (
	"log"
	"net/http"

	handlers "github.com/The-Flash/tus-go/internal/api"
)

func main() {
	port := ":8000"
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.Hello)
	router.HandleFunc("POST /files/", handlers.Files)
	log.Println("server started on port " + port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
