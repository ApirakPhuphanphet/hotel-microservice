package userService

import (
	userServicepb "auth/userService/proto"
	"context"
)

func GetUsertoLogin(username string, password string, client userServicepb.UserServiceClient) (*userServicepb.User, error) {
	// Simulate a user login process
	user, err := client.GetUserToLogin(context.Background(), &userServicepb.GetUserToLoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	res := &userServicepb.User{
		Id:       user.User.Id,
		Username: user.User.Username,
		Password: user.User.Password,
		Role:     user.User.Role,
	}

	return res, nil
}
