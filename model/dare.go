package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Message string `json:"Message"`
}

type Dare struct {
	// Model is a struct having default fields where ID is a primary key
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Dare string             `json:"dare,omitempty" bson:"dare,omitempty"`
}

type DareContainer []Dare
