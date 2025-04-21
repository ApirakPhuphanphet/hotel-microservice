package main

import (
	"context"
	"fmt"
	"net"

	"room/handler"
	roompb "room/proto"
	"room/repository"
	"room/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	db := client.Database("room")
	roomRepository := repository.NewRoomRepository(db, "rooms")
	roomService := service.NewRoomService(roomRepository)
	roomHandler := handler.NewRoomHandler(roomService)

	defer client.Disconnect(context.TODO())

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	defer grpcServer.Stop()
	roompb.RegisterRoomServiceServer(grpcServer, roomHandler)

	fmt.Print("gRPC server is running on port :50052 \n")
	grpcServer.Serve(listener)

}
