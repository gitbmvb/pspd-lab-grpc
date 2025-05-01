package main

import (
	"log"
	"net/http"
	"backend/db"
	"backend/handlers"
	"github.com/rs/cors"
)

func main() {
	db.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/login", handlers.Login)
	mux.HandleFunc("/api/ask", handlers.AskHandler)

	handler := cors.Default().Handler(mux)
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", handler)
}