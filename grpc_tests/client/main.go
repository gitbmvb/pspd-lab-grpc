package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "client/service"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Example: CreateUsuario
	usuario := &pb.UsuarioRequest{
		Email: "alice@example.com",
		Nome:  "Alice",
		Senha: "123456",
	}

	res, err := client.CreateUsuario(ctx, usuario)
	if err != nil {
		log.Fatalf("CreateUsuario failed: %v", err)
	}

	log.Printf("Usuario criado: %v", res)
}
