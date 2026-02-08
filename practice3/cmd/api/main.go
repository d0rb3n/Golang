package main

import (
	"log"
	"net/http"
	"practice2/internal/handlers"
	"practice2/internal/middleware"
	"practice2/internal/storage"
)

func main() {

	store := storage.NewStore()
	handler := handlers.NewHandler(store)

	mux := http.NewServeMux()
	mux.Handle("/tasks", http.HandlerFunc(handler.Tasks))

	var h http.Handler = mux

	h = middleware.Logging(h)
	h = middleware.APIKey(h)

	log.Println("Server started at :8080")

	log.Fatal(http.ListenAndServe(":8080", h))

}
