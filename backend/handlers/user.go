package handlers

import (
	"backend/grpc_services"
	"backend/models"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"time"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.ClientDB.ListUsers(ctx, &emptypb.Empty{})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusOK, grpc_services.Response{
		Status: "ok",
		Data:   resp.Users,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.ClientDB.GetUser(ctx, &grpc_services.UserReadDeleteRequest{Email: email})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusNotFound, grpc_services.Response{
			Status:  "error",
			Message: "User not found",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusOK, grpc_services.Response{
		Status: "ok",
		Data:   resp,
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userReq models.User
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid input data",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.ClientDB.CreateUser(ctx, &grpc_services.UserCreateUpdateRequest{
		Email:    userReq.Email,
		Name:     userReq.Name,
		Password: userReq.Password,
	})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to create user",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusCreated, grpc_services.Response{
		Status: "ok",
		Data:   resp,
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	var userReq models.User
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid input data",
		})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := grpc_services.ClientDB.UpdateUser(ctx, &grpc_services.UserCreateUpdateRequest{
		Email:    email,
		Name:     userReq.Name,
		Password: userReq.Password,
	})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to update user",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusOK, grpc_services.Response{
		Status: "ok",
		Data:   resp,
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	_, err := grpc_services.ClientDB.DeleteUser(ctx, &grpc_services.UserReadDeleteRequest{Email: email})
	if err != nil {
		grpc_services.SendJSONResponse(w, http.StatusInternalServerError, grpc_services.Response{
			Status:  "error",
			Message: "Failed to delete user",
		})
		return
	}

	grpc_services.SendJSONResponse(w, http.StatusNoContent, grpc_services.Response{
		Status: "ok",
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// 1. Parse request body
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Invalid request format",
		})
		return
	}

	// 2. Validate required fields
	if loginReq.Email == "" || loginReq.Password == "" {
		grpc_services.SendJSONResponse(w, http.StatusBadRequest, grpc_services.Response{
			Status:  "error",
			Message: "Email and password are required",
		})
		return
	}

	// 3. Create context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// 4. Call gRPC service
	_, err := grpc_services.ClientDB.LoginUser(ctx, &grpc_services.UserLoginRequest{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	})

	// 5. Handle response
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMsg := "Login failed"

		// Convert gRPC error to status
		st, _ := status.FromError(err)
		switch st.Code() {
		case codes.NotFound:
			statusCode = http.StatusUnauthorized
			errorMsg = "Invalid email or password" // Generic message for security
		case codes.Internal:
			errorMsg = "Internal server error"
		}

		grpc_services.SendJSONResponse(w, statusCode, grpc_services.Response{
			Status:  "error",
			Message: errorMsg,
		})
		return
	}

	// 6. Successful login response
	grpc_services.SendJSONResponse(w, http.StatusOK, grpc_services.Response{
		Status: "ok",
		Data: map[string]string{
			"email": loginReq.Email,
		},
	})
}
