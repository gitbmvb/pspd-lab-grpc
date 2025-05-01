package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"backend/db"
)

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM \"User\" WHERE email = $1 AND password = $2)", loginRequest.Email, loginRequest.Password).Scan(&exists)
	if err != nil || !exists {
		http.Error(w, "Unathorized", http.StatusUnauthorized)
		return
	}

	log.Printf("User %s logged in successfully", loginRequest.Email)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Login successful"}`))
}