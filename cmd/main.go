package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/AlexeyErmolenko/microservices-chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50052

type chatServer struct {
	desc.UnimplementedChatV1Server
}

func (c chatServer) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
}

func main() {
	fmt.Println(color.RedString("It's gRPC server"))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &chatServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
