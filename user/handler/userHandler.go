package handler

import (
	"context"
	"user/model"
	userpb "user/proto"
	"user/service"
)

type UserGRPCHandler struct {
	userpb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserGRPCHandler {
	return &UserGRPCHandler{
		service: service,
	}
}

func (s *UserGRPCHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := model.User{
		Username: req.User.Username,
		Password: req.User.Password,
		Role:     req.User.Role,
	}

	err := s.service.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	res := &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Role:     user.Role,
		},
	}
	return res, nil

}

func (s *UserGRPCHandler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := s.service.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Role:     user.Role,
		},
	}, nil
}

func (s *UserGRPCHandler) GetAllUsers(ctx context.Context, req *userpb.GetAllUsersRequest) (*userpb.GetAllUsersResponse, error) {
	users, err := s.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var userList []*userpb.User
	for _, user := range users {
		userList = append(userList, &userpb.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Role:     user.Role,
		})
	}

	return &userpb.GetAllUsersResponse{
		Users: userList,
	}, nil
}

func (s *UserGRPCHandler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	ID := req.Id
	user := model.User{
		Username: req.User.Username,
		Password: req.User.Password,
		Role:     req.User.Role,
	}

	err := s.service.UpdateUser(ID, &user)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Role:     user.Role,
		},
	}, nil
}

func (s *UserGRPCHandler) ChangeRole(ctx context.Context, req *userpb.ChangeRoleRequest) (*userpb.ChangeRoleResponse, error) {
	err := s.service.ChangeRole(req.Id, req.Role)

	res := &userpb.ChangeRoleResponse{
		Changed: false,
	}

	if err != nil {
		return res, err
	}

	res = &userpb.ChangeRoleResponse{
		Changed: true,
	}
	return res, nil
}

func (s *UserGRPCHandler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := s.service.DeleteUser(req.Id)
	if err != nil {
		return &userpb.DeleteUserResponse{
			Deleted: false,
		}, nil
	}

	return &userpb.DeleteUserResponse{
		Deleted: true,
	}, nil
}

func (s *UserGRPCHandler) GetUserToLogin(ctx context.Context, req *userpb.GetUserToLoginRequest) (*userpb.GetUserToLoginResponse, error) {
	user, err := s.service.GetUserToLogin(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserToLoginResponse{
		User: &userpb.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Role:     user.Role,
			Password: user.Password,
		},
	}, nil
}
