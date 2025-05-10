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

func ListMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idChat := vars["idChat"]

	idChatInt, err := strconv.ParseInt(idChat, 10, 64)
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid chat ID format",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.ClientDB.ListMessages(ctx, &grpc_services.ChatDeleteRequest{IdChat: idChatInt})
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
	idMessage := vars["idMessage"]

	idMessageInt, err := strconv.ParseInt(idMessage, 10, 64)
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid message ID format",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.ClientDB.GetMessage(ctx, &grpc_services.MessageReadRequest{IdMessage: idMessageInt})
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

	resp, err := grpc_services.ClientDB.CreateMessage(ctx, &grpc_services.MessageCreateRequest{
		Content:   msgReq.Content,
		IdChat:    msgReq.ChatID,
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
	idMessage := vars["idMessage"]

	idMessageInt, err := strconv.ParseInt(idMessage, 10, 64)
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

	resp, err := grpc_services.ClientDB.UpdateMessage(ctx, &grpc_services.MessageUpdateRequest{
		IdMessage: idMessageInt,
		Content:   msgReq.Content,
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
	idMessage := vars["idMessage"]

	idMessageInt, err := strconv.ParseInt(idMessage, 10, 64)
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid message ID format",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	_, err = grpc_services.ClientDB.DeleteMessage(ctx, &grpc_services.MessageDeleteRequest{IdMessage: idMessageInt})
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
