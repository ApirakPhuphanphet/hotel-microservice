package main

import (
	"log"
	"strconv"

	"github.com/ApirakPhuphanphet/hotel-microservice/roomService"
	roompb "github.com/ApirakPhuphanphet/hotel-microservice/roomService/proto"
	"github.com/ApirakPhuphanphet/hotel-microservice/userService"
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

	// Connect to the user service
	roomConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer roomConn.Close()
	roomClient := roompb.NewRoomServiceClient(roomConn)

	app := fiber.New()
	userHandler(app, userClient)
	roomHandler(app, roomClient)

	app.Listen(":3000")
}

func roomHandler(app *fiber.App, grpcRoomClient roompb.RoomServiceClient) {
	app.Post("/room", func(c *fiber.Ctx) error {
		room := new(roompb.Room)
		if err := c.BodyParser(room); err != nil {
			log.Printf("Error while parsing request body : %v", err)
		}

		res, err := roomService.CreateRoom(room, grpcRoomClient)
		if err != nil {
			log.Printf("Error while calling CreateRoom : %v", err)
			return err
		}
		return c.JSON(res)
	})

	app.Get("/room", func(c *fiber.Ctx) error {
		rooms, err := roomService.GetAllRoom(grpcRoomClient)
		if err != nil {
			log.Printf("Error while calling GetAllRoom : %v", err)
			return err
		}
		return c.JSON(rooms)
	})

	app.Get("/room/:number", func(c *fiber.Ctx) error {
		number, err := strconv.ParseInt(c.Params("number"), 10, 64)
		if err != nil {
			log.Printf("Error while parsing room number: %v", err)
			return err
		}

		room, err := roomService.GetRoomById(number, grpcRoomClient)
		if err != nil {
			log.Printf("Error while calling GetRoomById : %v", err)
			return err
		}
		return c.JSON(room)
	})

	app.Put("/room/updatePrice/:number", func(c *fiber.Ctx) error {
		number, err := strconv.ParseInt(c.Params("number"), 10, 64)
		if err != nil {
			log.Printf("Error while parsing room number: %v", err)
			return err
		}

		type PriceRequest struct {
			Price float64 `json:"price"`
		}
		priceReq := new(PriceRequest)
		if err := c.BodyParser(priceReq); err != nil {
			log.Printf("Error while parsing request body : %v", err)
			return err
		}
		price := priceReq.Price

		res, err := roomService.UpdatePrice(number, float32(price), grpcRoomClient)
		if err != nil {
			log.Printf("Error while calling UpdateRoom : %v", err)
			return err
		}
		return c.JSON(res)
	})

	app.Put("room/addBooking/:number", func(c *fiber.Ctx) error {
		number, err := strconv.ParseInt(c.Params("number"), 10, 64)
		if err != nil {
			log.Printf("Error while parsing room number: %v", err)
			return err
		}

		booking := new(roompb.Booking)
		if err := c.BodyParser(booking); err != nil {
			log.Printf("Error while parsing request body : %v", err)
			return err
		}
		log.Printf("Booking: %v", booking)
		res, err := roomService.AddBook(number, booking, grpcRoomClient)
		if err != nil {
			log.Printf("Error while calling UpdateRoom : %v", err)
			return err
		}
		return c.JSON(res)
	})

	app.Put("room/removeBooking/:number", func(c *fiber.Ctx) error {
		number, err := strconv.ParseInt(c.Params("number"), 10, 64)
		if err != nil {
			log.Printf("Error while parsing room number: %v", err)
			return err
		}

		booking := new(roompb.Booking)
		if err := c.BodyParser(booking); err != nil {
			log.Printf("Error while parsing request body : %v", err)
			return err
		}
		log.Printf("Booking: %v", booking)
		res, err := roomService.DeleteBooking(number, booking.Date, grpcRoomClient)
		if err != nil {
			log.Printf("Error while calling UpdateRoom : %v", err)
			return err
		}
		return c.JSON(res)
	})

	app.Delete("/room/:number", func(c *fiber.Ctx) error {
		number, err := strconv.ParseInt(c.Params("number"), 10, 64)
		if err != nil {
			log.Printf("Error while parsing room number: %v", err)
			return err
		}
		log.Printf("Room number: %d", number)
		res, err := roomService.DeleteRoom(number, grpcRoomClient)
		if err != nil {
			log.Printf("Error while calling DeleteRoom : %v", err)
			return err
		}
		return c.JSON(res)
	})
}

func userHandler(app *fiber.App, grpcUserClient userpb.UserServiceClient) {

	app.Post("/user", func(c *fiber.Ctx) error {
		user := new(userpb.User)
		if err := c.BodyParser(user); err != nil {
			log.Printf("Error while parsing request body: %v", err)
			return err
		}
		log.Printf("User: %v", user)
		res, err := userService.CreateUser(user, grpcUserClient)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		res, err := userService.GetAllUser(grpcUserClient)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		log.Print(id)
		res, err := userService.GetUserById(id, grpcUserClient)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})

	app.Put("/user/:id/:role", func(c *fiber.Ctx) error {
		id := c.Params("id")
		role := c.Params("role")
		log.Printf("ID: %s, Role: %s", id, role)
		res, err := userService.ChangeRole(id, role, grpcUserClient)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})

	app.Put("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		user := new(userpb.User)
		if err := c.BodyParser(user); err != nil {
			log.Printf("Error while parsing request body: %v", err)
			return err
		}
		log.Printf("User: %v", user)
		res, err := userService.UpdateUser(id, user, grpcUserClient)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})

	app.Delete("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		log.Printf("ID: %s", id)
		res, err := userService.DeleteUser(id, grpcUserClient)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})
}
