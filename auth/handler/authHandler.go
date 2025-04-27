package handler

import (
	auth "auth/authentication"
	authpb "auth/proto"
	userServicepb "auth/userService/proto"
	"context"
	"encoding/json"
	"log"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	client userServicepb.UserServiceClient
}

func NewAuthHandler(client userServicepb.UserServiceClient) *AuthHandler {
	return &AuthHandler{
		client: client,
	}
}

func (s *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	reqUser := userServicepb.GetUserToLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}
	user, err := s.client.GetUserToLogin(context.Background(), &reqUser)
	if err != nil {
		return nil, err
	}
	token, err := auth.Login(user.User, req.Password)

	if err != nil {
		return nil, err
	}

	res := authpb.LoginResponse{
		Token: token,
	}

	return &res, nil
}

func (s *AuthHandler) TokenValidation(ctx context.Context, req *authpb.TokenValidationRequest) (*authpb.TokenValidationResponse, error) {
	token := req.Token
	var res *authpb.TokenValidationResponse

	tokenVaidated, err := auth.ValidateJWT(token)

	if err != nil {
		res = &authpb.TokenValidationResponse{
			UserValidated: "",
		}
		return res, err
	}
	claimsJSON, err := json.Marshal(tokenVaidated)
	if err != nil {
		log.Println("Error converting claims to string:", err)
		return nil, err
	}
	res = &authpb.TokenValidationResponse{
		UserValidated: string(claimsJSON),
	}

	return res, nil
}
