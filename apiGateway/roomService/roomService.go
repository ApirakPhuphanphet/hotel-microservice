package roomService

import (
	"context"
	"log"

	roompb "github.com/ApirakPhuphanphet/hotel-microservice/roomService/proto"
)

func CreateRoom(room *roompb.Room, client roompb.RoomServiceClient) (*roompb.Room, error) {
	req := roompb.CreateRoomRequest{Room: room}
	res, err := client.CreateRoom(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling CreateRoom RPC: %v", err)
		return nil, err
	}
	return res.Room, nil
}

func GetRoomById(number int64, client roompb.RoomServiceClient) (*roompb.Room, error) {
	req := roompb.GetRoomRequest{Number: number}
	res, err := client.GetRoom(context.Background(), &req)

	if err != nil {
		log.Printf("Error while calling GetRoom RPC: %v", err)
		return nil, err
	}
	return res.Room, nil
}

func GetAllRoom(client roompb.RoomServiceClient) ([]*roompb.Room, error) {
	req := roompb.GetAllRoomsRequest{}
	res, err := client.GetAllRooms(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling GetAllRooms RPC: %v", err)
	}
	return res.Rooms, nil
}

func UpdatePrice(number int64, price float32, client roompb.RoomServiceClient) (*roompb.Room, error) {
	req := roompb.UpdatePriceRequest{Number: number, Price: price}
	res, err := client.UpdatePrice(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling UpdateRoom RPC: %v", err)
		return nil, err
	}
	return res.Room, nil
}

func AddBook(number int64, book *roompb.Booking, client roompb.RoomServiceClient) (*roompb.Room, error) {
	req := roompb.AddBookingRequest{Number: number, Booking: book}
	res, err := client.AddBooking(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling AddBook RPC: %v", err)
		return nil, err
	}
	return res.Room, nil
}

func DeleteBooking(number int64, Date string, client roompb.RoomServiceClient) (*roompb.Room, error) {
	req := roompb.DeleteBookingRequest{Number: number, Date: Date}
	res, err := client.DeleteBooking(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling DeleteBooking RPC: %v", err)
		return nil, err
	}
	return res.Room, nil
}

func DeleteRoom(number int64, client roompb.RoomServiceClient) (bool, error) {
	req := roompb.DeleteRoomRequest{Number: number}
	res, err := client.DeleteRoom(context.Background(), &req)
	if err != nil {
		log.Printf("Error while calling DeleteRoom RPC: %v", err)
		return false, err
	}
	return res.Status, nil
}
