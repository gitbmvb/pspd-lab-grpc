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

func GetMessages(w http.ResponseWriter, r *http.Request) {
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

	resp, err := grpc_services.Client.GetMessages(ctx, &grpc_services.ChatId{ChatId: chatIDInt})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to fetch messages",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusOK, grpc_services.Response{
		Status: "ok",
		Data:   resp.Messages,
	})
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID := vars["message_id"]

	messageIDInt, err := strconv.ParseInt(messageID, 10, 64)
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid message ID format",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.Client.GetMessage(ctx, &grpc_services.MessageData{MessageId: messageIDInt})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusNotFound, grpc_services.Response{
			Status:  "error",
			Message: "Message not found",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusOK, grpc_services.Response{
		Status: "ok",
		Data:   resp,
	})
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msgReq models.Message
	if err := json.NewDecoder(r.Body).Decode(&msgReq); err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid input data",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.Client.CreateMessage(ctx, &grpc_services.MessageRequest{
		MessageId: msgReq.MessageID,
		Question:  msgReq.Question,
		Answer:    msgReq.Answer,
		ChatId:    msgReq.ChatID,
	})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to create message",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusCreated, grpc_services.Response{
		Status: "ok",
		Data:   resp,
	})
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID := vars["message_id"]

	messageIDInt, err := strconv.ParseInt(messageID, 10, 64)
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid message ID format",
		})
		return
	}

	var msgReq models.Message
	if err := json.NewDecoder(r.Body).Decode(&msgReq); err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid input data",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.Client.UpdateMessage(ctx, &grpc_services.MessageRequest{
		MessageId: messageIDInt,
		Question:  msgReq.Question,
		Answer:    msgReq.Answer,
		ChatId:    msgReq.ChatID,
	})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to update message",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusOK, grpc_services.Response{
		Status: "ok",
		Data:   resp,
	})
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID := vars["message_id"]

	messageIDInt, err := strconv.ParseInt(messageID, 10, 64)
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid message ID format",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	_, err = grpc_services.Client.DeleteMessage(ctx, &grpc_services.MessageId{MessageId: messageIDInt})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to delete message",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusNoContent, grpc_services.Response{
		Status: "ok",
	})
}
