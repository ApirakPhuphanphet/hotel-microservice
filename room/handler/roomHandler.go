package handler

import (
	"context"
	"room/model"
	roompb "room/proto"
	"room/service"
)

type RoomGRPCHandler struct {
	roompb.UnimplementedRoomServiceServer
	service *service.RoomService
}

func NewRoomHandler(service *service.RoomService) *RoomGRPCHandler {
	return &RoomGRPCHandler{
		service: service,
	}
}

func (h *RoomGRPCHandler) CreateRoom(ctx context.Context, req *roompb.CreateRoomRequest) (*roompb.CreateRoomResponse, error) {
	room := model.Room{
		Number:   req.Room.Number,
		Type:     req.Room.Type,
		Price:    req.Room.Price,
		Bookings: []model.Booking{},
	}

	newRoom, err := h.service.CreateRoom(&room)
	if err != nil {
		return nil, err
	}
	var bookingDate []*roompb.Booking
	res := &roompb.Room{
		Id:      newRoom.ID.Hex(),
		Number:  newRoom.Number,
		Type:    newRoom.Type,
		Price:   float32(newRoom.Price),
		Booking: bookingDate,
	}
	return &roompb.CreateRoomResponse{Room: res}, nil
}

func (h *RoomGRPCHandler) GetRoom(ctx context.Context, req *roompb.GetRoomRequest) (*roompb.GetRoomResponse, error) {
	room, err := h.service.GetRoomByNumber(req.Number)
	if err != nil {
		return nil, err
	}

	var bookings []*roompb.Booking
	for _, b := range room.Bookings {
		bookings = append(bookings, &roompb.Booking{Date: b.Date})
	}

	res := &roompb.Room{
		Id:      room.ID.Hex(),
		Number:  room.Number,
		Type:    room.Type,
		Price:   float32(room.Price),
		Booking: bookings,
	}
	return &roompb.GetRoomResponse{Room: res}, nil
}

func (h *RoomGRPCHandler) GetAllRooms(ctx context.Context, req *roompb.GetAllRoomsRequest) (*roompb.GetAllRoomsResponse, error) {
	rooms, err := h.service.GetAllRooms()
	if err != nil {
		return nil, err
	}

	var roomList []*roompb.Room
	for _, room := range rooms {
		var bookings []*roompb.Booking
		for _, b := range room.Bookings {
			bookings = append(bookings, &roompb.Booking{Date: b.Date})
		}
		roomList = append(roomList, &roompb.Room{
			Id:      room.ID.Hex(),
			Number:  room.Number,
			Type:    room.Type,
			Price:   float32(room.Price),
			Booking: bookings,
		})
	}

	return &roompb.GetAllRoomsResponse{Rooms: roomList}, nil
}

func (h *RoomGRPCHandler) UpdatePrice(ctx context.Context, req *roompb.UpdatePriceRequest) (*roompb.UpdatePriceResponse, error) {
	err := h.service.UpdatePrice(req.Price, req.Number)
	if err != nil {
		return nil, err
	}

	room, err := h.service.GetRoomByNumber(req.Number)
	if err != nil {
		return nil, err
	}

	var bookings []*roompb.Booking
	for _, b := range room.Bookings {
		bookings = append(bookings, &roompb.Booking{Date: b.Date})
	}

	res := &roompb.Room{
		Id:      room.ID.Hex(),
		Number:  room.Number,
		Type:    room.Type,
		Price:   float32(room.Price),
		Booking: bookings,
	}
	return &roompb.UpdatePriceResponse{Room: res}, nil
}

func (h *RoomGRPCHandler) AddBooking(ctx context.Context, req *roompb.AddBookingRequest) (*roompb.AddBookingResponse, error) {
	err := h.service.AddBooking(req.Number, req.Booking.Date)
	if err != nil {
		return nil, err
	}

	room, err := h.service.GetRoomByNumber(req.Number)
	if err != nil {
		return nil, err
	}

	var bookings []*roompb.Booking
	for _, b := range room.Bookings {
		bookings = append(bookings, &roompb.Booking{Date: b.Date})
	}

	res := &roompb.Room{
		Id:      room.ID.Hex(),
		Number:  room.Number,
		Type:    room.Type,
		Price:   float32(room.Price),
		Booking: bookings,
	}
	return &roompb.AddBookingResponse{Room: res}, nil
}

func (h RoomGRPCHandler) DeleteBooking(ctx context.Context, req *roompb.DeleteBookingRequest) (*roompb.DeleteBookingResponse, error) {
	err := h.service.CancelBooking(req.Number, req.Date)
	if err != nil {
		return nil, err
	}

	room, err := h.service.GetRoomByNumber(req.Number)
	if err != nil {
		return nil, err
	}

	var bookings []*roompb.Booking
	for _, b := range room.Bookings {
		bookings = append(bookings, &roompb.Booking{Date: b.Date})
	}

	res := &roompb.Room{
		Id:      room.ID.Hex(),
		Number:  room.Number,
		Type:    room.Type,
		Price:   float32(room.Price),
		Booking: bookings,
	}
	return &roompb.DeleteBookingResponse{Room: res}, nil
}

func (h *RoomGRPCHandler) DeleteRoom(ctx context.Context, req *roompb.DeleteRoomRequest) (*roompb.DeleteRoomResponse, error) {
	err := h.service.DeleteRoom(req.Number)
	if err != nil {
		return &roompb.DeleteRoomResponse{
			Status: false,
		}, err
	}

	return &roompb.DeleteRoomResponse{
		Status: true,
	}, nil
}
