package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dare struct {
	// Model is a struct having default fields where ID is a primary key
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Question string             `json:"question,omitempty" bson:"question,omitempty"`
}

type DareContainer []Dare
