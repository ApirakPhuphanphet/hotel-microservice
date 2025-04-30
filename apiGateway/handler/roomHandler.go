// handlers/room_handler.go
package handlers

import (
	"log"
	"strconv"

	authpb "github.com/ApirakPhuphanphet/hotel-microservice/authService/proto"
	"github.com/ApirakPhuphanphet/hotel-microservice/roomService"
	roompb "github.com/ApirakPhuphanphet/hotel-microservice/roomService/proto"
	"github.com/gofiber/fiber/v2"
)

func RoomHandler(app *fiber.App, grpcRoomClient roompb.RoomServiceClient, authClient authpb.AuthServiceClient) {
	app.Post("/room", func(c *fiber.Ctx) error {
		// Check if user is admin?
		userValidated, err := ValidateToken(c, authClient)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if userValidated.Role != "admin" {
			return c.JSON(fiber.Map{
				"error": "You don't have permission",
			})
		}
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
		// Check if user is valid?
		_, err := ValidateToken(c, authClient)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		rooms, err := roomService.GetAllRoom(grpcRoomClient)
		if err != nil {
			log.Printf("Error while calling GetAllRoom : %v", err)
			return err
		}
		return c.JSON(rooms)
	})

	app.Get("/room/:number", func(c *fiber.Ctx) error {
		// Check if user is Valid?
		_, err := ValidateToken(c, authClient)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

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
		// Check if user is admin?
		userValidated, err := ValidateToken(c, authClient)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if userValidated.Role != "admin" {
			return c.JSON(fiber.Map{
				"error": "You don't have permission",
			})
		}

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
		// Check if user is Valid?
		_, err := ValidateToken(c, authClient)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

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
		// Check if user is Valid?
		_, err := ValidateToken(c, authClient)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}
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
		// Check if user is admin?
		userValidated, err := ValidateToken(c, authClient)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if userValidated.Role != "admin" {
			return c.JSON(fiber.Map{
				"error": "You don't have permission",
			})
		}

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
