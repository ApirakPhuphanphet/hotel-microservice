package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Room struct {
		ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Number    int64              `json:"number" bson:"number"`
		Type      string             `json:"type" bson:"type"`
		Price     float32            `json:"price" bson:"price"`
		Boockings []Boocking         `json:"boockings" bson:"boockings"`
	}

	Boocking struct {
		Date string `json:"date" bson:"date"`
	}
)
