package handlers

import (
	"errors"
	"strings"

	"github.com/ApirakPhuphanphet/hotel-microservice/authService"
	authpb "github.com/ApirakPhuphanphet/hotel-microservice/authService/proto"
	"github.com/gofiber/fiber/v2"
)

func AuthHandler(app *fiber.App, authClient authpb.AuthServiceClient) {
	app.Post("/auth", func(c *fiber.Ctx) error {
		var req *authpb.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		token, err := authService.Login(req.Username, req.Password, authClient)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"token": token})
	})
}

func ValidateToken(c *fiber.Ctx, authClient authpb.AuthServiceClient) (authService.UserValidate, error) {
	var user authService.UserValidate
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		err := errors.New("Authorization header missing")
		return user, err
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		err := errors.New("Invalid Authorization header format")
		return user, err
	}

	user, err := authService.TokenValidation(tokenParts[1], authClient)
	if err != nil {
		return user, err
	}

	return user, nil
}
