package repository

import (
	"context"

	"user/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collectionName string) *UserRepository {
	return &UserRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *UserRepository) FindUserByID(id interface{}) (model.User, error) {
	var result model.User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	return result, err
}

func (r *UserRepository) UpdateUser(id interface{}, update *model.User) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (r *UserRepository) DeleteUser(id interface{}) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *UserRepository) FindAllUsers() ([]model.User, error) {
	var users []model.User
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) FindUserByUsername(username string) (model.User, error) {
	var result model.User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	return result, err
}
