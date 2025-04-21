package service

import (
	"errors"
	"room/model"
	"room/repository"
)

type RoomService struct {
	repository *repository.RoomRepository
}

func NewRoomService(repo *repository.RoomRepository) *RoomService {
	return &RoomService{
		repository: repo,
	}
}
func (s *RoomService) CreateRoom(room *model.Room) (model.Room, error) {
	newRoom := model.Room{}
	existedroom, err := s.repository.GetRoomByNumber(room.Number)
	// Check if the room number already exists
	if err == nil {
		if existedroom.Number == room.Number {
			return newRoom, errors.New("room already exists")
		}
	}
	// Validate the room number
	if room.Number <= 0 {
		err = errors.New("invalid room number")
		return newRoom, err
	}
	newRoom, err = s.repository.CreateRoom(room)
	return newRoom, err
}

func (s *RoomService) GetRoomByNumber(number int64) (model.Room, error) {
	room, err := s.repository.GetRoomByNumber(number)
	if err != nil {
		return model.Room{}, err
	}
	return room, nil
}

func (s *RoomService) GetAllRooms() ([]model.Room, error) {
	rooms, err := s.repository.GetAllRooms()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *RoomService) UpdatePrice(price float32, number int64) error {
	room, err := s.repository.GetRoomByNumber(number)
	if err != nil {
		return err
	}

	room.Price = price
	err = s.repository.UpdateRoom(room.ID, room)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoomService) AddBooking(number int64, booking string) error {
	room, err := s.repository.GetRoomByNumber(number)
	if err != nil {
		return err
	}
	// Check if the booking date already exists
	for _, b := range room.Boockings {
		if b.Date == booking {
			return errors.New("booking date already exists")
		}
	}
	room.Boockings = append(room.Boockings, model.Boocking{Date: booking})
	err = s.repository.UpdateRoom(room.ID, room)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoomService) CancelBooking(number int64, booking string) error {
	room, err := s.repository.GetRoomByNumber(number)
	if err != nil {
		return err
	}
	for i, b := range room.Boockings {
		if b.Date == booking {
			room.Boockings = append(room.Boockings[:i], room.Boockings[i+1:]...)
			break
		}
	}
	err = s.repository.UpdateRoom(room.ID, room)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoomService) DeleteRoom(number int64) error {
	room, err := s.repository.GetRoomByNumber(number)
	if err != nil {
		return err
	}
	err = s.repository.DeleteRoom(room.ID)
	if err != nil {
		return err
	}
	return nil
}
