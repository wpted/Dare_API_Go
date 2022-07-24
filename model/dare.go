package model

type Message struct {
	Message string `json:"Message"`
}
type Dare struct {
	// Model is a struct having default fields where ID is a primary key
	ID   string
	Dare string `json:"dare" "`
}

type DareContainer []Dare
