package model

import "github.com/jinzhu/gorm"

type Message struct {
	Message string `json:"Message"`
}
type Dare struct {
	// Model is a struct having default fields where ID is a primary key
	gorm.Model
	DareQuestion string `json:"Dare" gorm:"unique_index;not null"`
}

type DareContainer []Dare
