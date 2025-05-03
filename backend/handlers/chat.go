package handlers

import (
	"backend/grpc_services"
	"backend/models"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func GetChats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.Client.GetChats(ctx, &grpc_services.Email{Email: email})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to fetch chats",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusOK, grpc_services.Response{
		Status: "ok",
		Data:   resp.Chats,
	})
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatID := vars["chat_id"]

	chatIDInt, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid chat ID format",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.Client.GetChat(ctx, &grpc_services.ChatData{ChatId: chatIDInt})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusNotFound, grpc_services.Response{
			Status:  "error",
			Message: "Chat not found",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusOK, grpc_services.Response{
		Status: "ok",
		Data:   resp,
	})
}

func CreateChat(w http.ResponseWriter, r *http.Request) {
	var chatReq models.Chat
	if err := json.NewDecoder(r.Body).Decode(&chatReq); err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid input data",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.Client.CreateChat(ctx, &grpc_services.ChatRequest{
		ChatId: chatReq.ChatID,
		Topic:  chatReq.Topic,
		Email:  chatReq.Email,
	})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to create chat",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusCreated, grpc_services.Response{
		Status: "ok",
		Data:   resp,
	})
}

func UpdateChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatID := vars["chat_id"]

	chatIDInt, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid chat ID format",
		})
		return
	}

	var chatReq models.Chat
	if err := json.NewDecoder(r.Body).Decode(&chatReq); err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid input data",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.Client.UpdateChat(ctx, &grpc_services.ChatRequest{
		ChatId: chatIDInt,
		Topic:  chatReq.Topic,
		Email:  chatReq.Email,
	})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to update chat",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusOK, grpc_services.Response{
		Status: "ok",
		Data:   resp,
	})
}

func DeleteChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatID := vars["chat_id"]

	chatIDInt, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid chat ID format",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	_, err = grpc_services.Client.DeleteChat(ctx, &grpc_services.ChatId{ChatId: chatIDInt})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to delete chat",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusNoContent, grpc_services.Response{
		Status: "ok",
	})
}
