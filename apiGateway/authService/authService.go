package authService

import (
	"context"
	"encoding/json"
	"log"

	authpb "github.com/ApirakPhuphanphet/hotel-microservice/authService/proto"
)

type UserValidate struct {
	Exp      int64
	Role     string
	Username string
}

func Login(username string, password string, client authpb.AuthServiceClient) (string, error) {
	req := authpb.LoginRequest{
		Username: username,
		Password: password,
	}

	res, err := client.Login(context.Background(), &req)
	if err != nil {
		return res.Token, err
	}

	return res.Token, err
}

func TokenValidation(token string, client authpb.AuthServiceClient) (UserValidate, error) {
	req := authpb.TokenValidationRequest{
		Token: token,
	}
	userValidated := UserValidate{
		Exp:      0,
		Role:     "",
		Username: "",
	}

	res, err := client.TokenValidation(context.Background(), &req)
	if err != nil {
		return userValidated, err
	}
	err = json.Unmarshal([]byte(res.UserValidated), &userValidated)
	if err != nil {
		return userValidated, err
	}
	log.Print(userValidated)
	return userValidated, nil
}
