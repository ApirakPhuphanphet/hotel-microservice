package service

import (
	"errors"
	"user/model"
	"user/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(user *model.User) error {
	_, err := s.repository.FindUserByUsername(user.Username)
	// If the user does not exist, create a new user
	if err != nil {
		bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(bytes)
		return s.repository.CreateUser(user)
	}
	// If the user exists, return an error
	err = errors.New("username already exists")
	return err
}

func (s *UserService) GetUserByID(id interface{}) (model.User, error) {
	var user model.User
	ID, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return model.User{}, err
	}
	user, err = s.repository.FindUserByID(ID)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	users, err := s.repository.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdateUser(id interface{}, update *model.User) error {
	var user model.User
	ID, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}
	user, err = s.repository.FindUserByID(ID)
	// If the user does not exist, return an error
	if err != nil {
		return err
	} else {
		user, err = s.repository.FindUserByUsername(update.Username)
		if err == nil {
			return errors.New("username already exists")
		}
	}

	if update.Username == "" {
		update.Username = user.Username
	}

	if update.Role == "" {
		update.Role = user.Role
	}

	if update.Password != "" {
		bytes, err := bcrypt.GenerateFromPassword([]byte(update.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		update.Password = string(bytes)
	} else {
		update.Password = user.Password
	}

	err = s.repository.UpdateUser(ID, update)
	return err
}

func (s *UserService) ChangeRole(id interface{}, role string) error {
	ID, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}
	user, err := s.repository.FindUserByID(ID)
	if err != nil {
		return err
	}
	if user.Role == role {
		return errors.New("user role is the same")
	}
	user.Role = role
	if user.Role == "" {
		return errors.New("role is empty")
	}

	err = s.repository.UpdateUser(ID, &user)
	return err
}

func (s *UserService) DeleteUser(id interface{}) error {
	ID, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}
	_, err = s.repository.FindUserByID(ID)
	if err != nil {
		err = errors.New("user not found")
		return err
	}
	err = s.repository.DeleteUser(ID)
	return err
}

func (s *UserService) GetUserToLogin(username, password string) (model.User, error) {
	var user model.User
	user, err := s.repository.FindUserByUsername(username)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
