package main

import (
	"log"

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

	app := fiber.New()
	userHandler(app, userClient)

	app.Listen(":3000")
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
