package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/pemarsao/fc2-grpc/pb"
	"github.com/pemarsao/fc2-grpc/services"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grcpServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grcpServer, services.NewUserService())
	reflection.Register(grcpServer)

	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}