package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type message struct {
	Message string `json:"Message"`
}
type dare struct {
	Id           int    `json:"ID"`
	DareQuestion string `json:"Dare"`
}

type dareList []dare

var MockDareList = dareList{
	{1, "Yell hello world out loud"},
	{2, "Sing to someone beside you"},
	{3, "Dance as if it's the end of the world"},
}

type Header interface {
	Set(key, value string)
}

// setHTTPResponseHeader is a helper function setting up the HTTP method
func setHTTPResponseHeader(header Header) {
	// Content-Type is used to indicate the media type of the resource
	header.Set("Content-Type", "application/json")

	// Header set for CORS issue
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Home(w http.ResponseWriter, r *http.Request) {
	// message is a one-time struct only appears here
	message := message{"welcome to drunk dares"}

	// Marshal returns a jsonified message as []byte and an error
	output, err := json.Marshal(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	setHTTPResponseHeader(w.Header())

	// write output []bytes to writer interface, here the http.ResponseWriter
	w.Write(output)

}

// getRandomIndex is a helper function getting a random index within [0, len(dareList))
func getRandomIndex(d dareList) int {
	// set seed related accordingly to the current time
	rand.Seed(time.Now().Unix())
	return rand.Intn(len(d))
}

// getDareByID is a helper to search dare with the given ID
func getDareByID(db dareList, ID int) string {
	return db[ID-1].DareQuestion
}

// GetRandomDare is a dareList method that returns a random jsonified dare
func (d *dareList) GetRandomDare(w http.ResponseWriter, r *http.Request) {

	randomDare := getDareByID(*d, getRandomIndex(*d))

	output, err := json.Marshal(randomDare)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	setHTTPResponseHeader(w.Header())

	w.Write(output)
}

// GetAllDare is a dareList method that returns all jsonified dares
func (d *dareList) GetAllDare(w http.ResponseWriter, r *http.Request) {

	output, err := json.Marshal(*d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	setHTTPResponseHeader(w.Header())

	w.Write(output)
}

// should replace the db to the actual database used
// addToDatabase adds dare to the database
func addToDatabase(d dare, db dareList) {
	db = append(db, d)
}

//
//func
