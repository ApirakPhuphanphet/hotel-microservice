package userService

import (
	"context"
	"log"

	userpb "github.com/ApirakPhuphanphet/hotel-microservice/userService/proto"
)

func CreateUser(user *userpb.User, client userpb.UserServiceClient) (*userpb.User, error) {
	req := userpb.CreateUserRequest{User: user}
	res, err := client.CreateUser(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling CreateUser RPC: %v", err)
		return nil, err
	}
	return res.User, nil
}

func GetUserById(id string, client userpb.UserServiceClient) (*userpb.User, error) {
	req := userpb.GetUserRequest{Id: id}
	res, err := client.GetUser(context.Background(), &req)

	if err != nil {
		log.Printf("Error while calling GetUser RPC: %v", err)
		return nil, err
	}
	return res.User, nil
}

func GetAllUser(client userpb.UserServiceClient) ([]*userpb.User, error) {
	req := userpb.GetAllUsersRequest{}
	res, err := client.GetAllUsers(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling GetAllUsers RPC: %v", err)
	}
	return res.Users, nil
}

func ChangeRole(id string, role string, client userpb.UserServiceClient) (bool, error) {
	req := userpb.ChangeRoleRequest{Id: id, Role: role}
	res, err := client.ChangeRole(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling UpdateUser RPC: %v", err)
		return false, err
	}
	return res.Changed, nil
}

func UpdateUser(id string, user *userpb.User, client userpb.UserServiceClient) (*userpb.User, error) {
	req := userpb.UpdateUserRequest{Id: id, User: user}
	res, err := client.UpdateUser(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling UpdateUser RPC: %v", err)
		return nil, err
	}
	return res.User, nil
}

func DeleteUser(id string, client userpb.UserServiceClient) (bool, error) {
	req := userpb.DeleteUserRequest{Id: id}
	res, err := client.DeleteUser(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling DeleteUser RPC: %v", err)
		return false, err
	}
	return res.Deleted, nil
}
