package main

import (
	"log"

	authpb "github.com/ApirakPhuphanphet/hotel-microservice/authService/proto"
	handlers "github.com/ApirakPhuphanphet/hotel-microservice/handler"
	roompb "github.com/ApirakPhuphanphet/hotel-microservice/roomService/proto"
	userpb "github.com/ApirakPhuphanphet/hotel-microservice/userService/proto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the user service
	userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer userConn.Close()
	userClient := userpb.NewUserServiceClient(userConn)

	// Connect to the room service
	roomConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer roomConn.Close()
	roomClient := roompb.NewRoomServiceClient(roomConn)

	// Connect to the auth service
	authConn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer authConn.Close()
	authClient := authpb.NewAuthServiceClient(authConn)

	app := fiber.New()
	// Handler orther services
	handlers.AuthHandler(app, authClient)
	handlers.RoomHandler(app, roomClient, authClient)
	handlers.UserHandler(app, userClient, authClient)

	app.Listen(":3000")
}
