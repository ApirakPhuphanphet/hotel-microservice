package main

import (
	"context"
	"fmt"
	"net"

	"user/handler"
	userpb "user/proto"
	"user/repository"
	"user/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	db := client.Database("user")
	userRepository := repository.NewUserRepository(db, "users")
	UserService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(UserService)

	defer client.Disconnect(context.TODO())

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	defer grpcServer.Stop()
	userpb.RegisterUserServiceServer(grpcServer, userHandler)

	fmt.Print("gRPC server is running on port :50051 \n")
	grpcServer.Serve(listener)

}
