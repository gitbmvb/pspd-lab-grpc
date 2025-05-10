package main

import (
	"backend/grpc_services"
	"backend/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	grpc_services.Init()
	
	// Use gorilla/mux instead of http.ServeMux
	router := mux.NewRouter().StrictSlash(true) // StrictSlash handles trailing slashes

	// User routes
	userRouter := router.PathPrefix("/api/users").Subrouter()
	userRouter.HandleFunc("", handlers.ListUsers).Methods("GET")
	userRouter.HandleFunc("/{email}", handlers.GetUser).Methods("GET")
	userRouter.HandleFunc("/create", handlers.CreateUser).Methods("POST")
	userRouter.HandleFunc("/update/{email}", handlers.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/delete/{email}", handlers.DeleteUser).Methods("DELETE")

	// Chat routes
	chatRouter := router.PathPrefix("/api/chats").Subrouter()
	chatRouter.HandleFunc("/{email}", handlers.ListChats).Methods("GET")
	chatRouter.HandleFunc("/{id}", handlers.GetChat).Methods("GET")
	chatRouter.HandleFunc("/create", handlers.CreateChat).Methods("POST")
	chatRouter.HandleFunc("/update/{id}", handlers.UpdateChat).Methods("PUT")
	chatRouter.HandleFunc("/delete/{id}", handlers.DeleteChat).Methods("DELETE")

	// Message routes
	messageRouter := router.PathPrefix("/api/messages").Subrouter()
	messageRouter.HandleFunc("", handlers.ListMessages).Methods("GET")
	messageRouter.HandleFunc("/{idMessage}", handlers.GetMessage).Methods("GET")
	messageRouter.HandleFunc("/create", handlers.CreateMessage).Methods("POST")
	messageRouter.HandleFunc("/update/{id}", handlers.UpdateMessage).Methods("PUT")
	messageRouter.HandleFunc("/delete/{id}", handlers.DeleteMessage).Methods("DELETE")

	// Auth & Misc routes
	router.HandleFunc("/api/login", handlers.LoginUser).Methods("POST")
	router.HandleFunc("/api/ask", handlers.AskHandler).Methods("POST")

	// CORS middleware
	handler := cors.Default().Handler(router)
	
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}