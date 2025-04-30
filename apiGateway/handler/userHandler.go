package handlers

import (
	"log"

	authpb "github.com/ApirakPhuphanphet/hotel-microservice/authService/proto"
	"github.com/ApirakPhuphanphet/hotel-microservice/userService"
	userpb "github.com/ApirakPhuphanphet/hotel-microservice/userService/proto"
	"github.com/gofiber/fiber/v2"
)

func UserHandler(app *fiber.App, grpcUserClient userpb.UserServiceClient, authClient authpb.AuthServiceClient) {
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
		// Check if user is admin?
		userValidated, err := ValidateToken(c, authClient)
		log.Println(userValidated.Username)
		log.Println(userValidated.Role)
		log.Println(err)
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

		res, err := userService.GetAllUser(grpcUserClient)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		// Check if user is valid?
		_, err := ValidateToken(c, authClient)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		id := c.Params("id")
		log.Print(id)
		res, err := userService.GetUserById(id, grpcUserClient)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})

	app.Put("/user/:id/:role", func(c *fiber.Ctx) error {
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
		// Check if user is valid?
		_, err := ValidateToken(c, authClient)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}
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
		// Check if user is valid?
		_, err := ValidateToken(c, authClient)
		if err != nil {
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		id := c.Params("id")
		log.Printf("ID: %s", id)
		res, err := userService.DeleteUser(id, grpcUserClient)
		if err != nil {
			return err
		}
		return c.JSON(res)
	})
}
