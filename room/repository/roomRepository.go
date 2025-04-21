package repository

import (
	"context"
	"room/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository struct {
	collection *mongo.Collection
}

func NewRoomRepository(db *mongo.Database, collectionName string) *RoomRepository {
	collection := db.Collection(collectionName)
	return &RoomRepository{
		collection: collection,
	}
}

func (r *RoomRepository) CreateRoom(room *model.Room) (model.Room, error) {
	responseRoom := model.Room{}
	newRoom, err := r.collection.InsertOne(context.Background(), room)
	if err != nil {
		return responseRoom, err
	}

	room.ID = newRoom.InsertedID.(primitive.ObjectID)
	return *room, nil

}

func (r *RoomRepository) GetRoomByNumber(number int64) (model.Room, error) {
	var room model.Room
	err := r.collection.FindOne(context.Background(), bson.M{"number": number}).Decode(&room)
	if err != nil {
		return room, err
	}
	return room, nil
}

func (r *RoomRepository) GetAllRooms() ([]model.Room, error) {
	var rooms []model.Room
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var room model.Room
		if err := cursor.Decode(&room); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *RoomRepository) UpdateRoom(id primitive.ObjectID, room model.Room) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": room.ID}, bson.M{"$set": room})
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepository) DeleteRoom(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
