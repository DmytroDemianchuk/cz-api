package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type People struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	PhoneNumber string             `json:"phone_number"`
	BirthYear   string             `json:"birth_year"`
	// Watched     bool               `json:"watched,omitempty"`
}
