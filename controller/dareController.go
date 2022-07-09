package controller

import (
	"encoding/json"
	"errors"
	"fmt"
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

// ------------------------------- Get -------------------------------

// Home is a Get handler function being held by the DefaultServerMux
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

// ------------------------------- Post -------------------------------

// this is only a mocker, should replace the db to the actual database used
// addToDatabase adds dare to the database
func (db *dareList) addToDatabase(d dare) (err error) {
	for _, dare := range *db {
		// can't have same id
		// this should automatically increase with primary key in database
		if d.Id == dare.Id {
			err = errors.New("dare with this id already exist")
			return
		}
		// can't have another same dare question
		if d.DareQuestion == dare.DareQuestion {
			err = errors.New("same dare question already exist")
			return
		}
	}
	*db = append(*db, d)
	err = nil
	return err
}

// CreateDare creates a new dare if a dare doesn't exist in the database
func CreateDare(w http.ResponseWriter, r *http.Request) {
	var newDare dare
	// NewDecoder returns a new Decoder(a struct reads and decodes JSON values from an input stream -> io.reader) that reads from r
	// Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
	err := json.NewDecoder(r.Body).Decode(&newDare)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// update request body to database
	err = MockDareList.addToDatabase(newDare)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(200)
	}

	// for debug usage
	fmt.Println(MockDareList)

}

// ------------------------------- Put -------------------------------
// this is only a mocker, should replace the db to the actual database used
// updateDatabase update existing dare in the database
//func (db *dareList) updateDatabase(d dare) (err error) {
//	for _, dare := range *db {
//		if d.Id == dare.Id {
//			dare.DareQuestion = d.DareQuestion
//			return nil
//		} else {
//			err = errors.New("can't find dare id")
//			return err
//		}
//	}
//	return
//}
//
//func UpdateDare(w http.ResponseWriter, r *http.Request) {
//	var newDare dare
//	err := json.NewDecoder(r.Body).Decode(&newDare)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//
//	err = MockDareList.updateDatabase(newDare)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	} else {
//		w.WriteHeader(http.StatusOK)
//	}
//
//	// for debug usage
//	fmt.Println(MockDareList)
//}
