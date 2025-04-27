package main

import (
	"auth/handler"
	authpb "auth/proto"
	userServicepb "auth/userService/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the user service
	userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer userConn.Close()
	userClient := userServicepb.NewUserServiceClient(userConn)

	authHandler := handler.NewAuthHandler(userClient)

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	defer grpcServer.Stop()
	authpb.RegisterAuthServiceServer(grpcServer, authHandler)

	log.Print("gRPC server is running on port :50053 \n")
	grpcServer.Serve(listener)
}
