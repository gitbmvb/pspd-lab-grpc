package main

import (
	"backend/grpc_services"
	"backend/handlers"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	grpc_services.Init()
	mux := http.NewServeMux()

	mux.HandleFunc("/api/users", handlers.ListUsers)
	mux.HandleFunc("/api/users/", handlers.GetUser)
	mux.HandleFunc("/api/users/create", handlers.CreateUser)
	mux.HandleFunc("/api/users/update/", handlers.UpdateUser)
	mux.HandleFunc("/api/users/delete/", handlers.DeleteUser)

	mux.HandleFunc("/api/chats", handlers.ListChats)
	mux.HandleFunc("/api/chats/", handlers.GetChat)
	mux.HandleFunc("/api/chats/create", handlers.CreateChat)
	mux.HandleFunc("/api/chats/update/", handlers.UpdateChat)
	mux.HandleFunc("/api/chats/delete/", handlers.DeleteChat)

	mux.HandleFunc("/api/messages", handlers.ListMessages)
	mux.HandleFunc("/api/messages/", handlers.GetMessage)
	mux.HandleFunc("/api/messages/create", handlers.CreateMessage)
	mux.HandleFunc("/api/messages/update/", handlers.UpdateMessage)
	mux.HandleFunc("/api/messages/delete/", handlers.DeleteMessage)

	mux.HandleFunc("/api/login", handlers.Login)
	mux.HandleFunc("/api/ask", handlers.AskHandler)

	handler := cors.Default().Handler(mux)
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", handler)
}
