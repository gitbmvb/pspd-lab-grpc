package grpc_services

import (
	"context"
    "time"
	"encoding/json"
	"log"
	"net/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var ClientDB DataServiceClient
var ClientLLM LLMServiceClient

func Init() {
	log.Println("Initializing gRPC clients...")
	InitDBClient()
	log.Println("gRPC client for DB initialized")
	InitLLMClient()
	log.Println("gRPC client for LLM initialized")
}

func InitDBClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	conn, err := grpc.DialContext(
		ctx,
		"192.168.100.3:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	if conn.GetState().String() != "READY" {
		log.Fatalf("gRPC connection is not ready, current state: %v", conn.GetState())
	} else {
		log.Printf("gRPC connection is ready, current state: %v", conn.GetState())
	}

	ClientDB = NewDataServiceClient(conn)
	
	log.Println("Successfully connected to gRPC server")
}

func InitLLMClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	conn, err := grpc.DialContext(
		ctx,
		"192.168.100.2:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	if conn.GetState().String() != "READY" {
		log.Fatalf("gRPC connection is not ready, current state: %v", conn.GetState())
	} else {
		log.Printf("gRPC connection is ready, current state: %v", conn.GetState())
	}

	ClientLLM = NewLLMServiceClient(conn)
	log.Println("Successfully connected to gRPC server")
}


func SendJSONResponse(w http.ResponseWriter, statusCode int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
