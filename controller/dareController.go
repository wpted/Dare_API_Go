package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

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

func Home(w http.ResponseWriter, r *http.Request) {
	// message is a one-time struct only appears here
	message := struct {
		Message string `json:"Message"`
	}{"welcome to drunk dares"}

	// Marshal returns a jsonified message as []byte and an error
	output, err := json.Marshal(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Content-Type is used to indicate the media type of the resource
	w.Header().Set("Content-Type", "application/json")

	// Header set for CORS issue
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// write output []bytes to writer interface, here the http.ResponseWriter
	w.Write(output)

}

// getRandomIndex is a helper function getting a random index within [0, len(dareList))
func getRandomIndex(d dareList) int {
	// set seed related accordingly to the current time
	rand.Seed(time.Now().Unix())
	return rand.Intn(len(d))
}

// GetRandomDare is a dareList method that returns a random jsonified dare
func (d *dareList) GetRandomDare(w http.ResponseWriter, r *http.Request) {

	randomDare := (*d)[getRandomIndex(*d)].DareQuestion

	output, err := json.Marshal(randomDare)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.Write(output)
}

// GetAllDare is a dareList method that returns all jsonified dares
func (d *dareList) GetAllDare(w http.ResponseWriter, r *http.Request) {
	output, err := json.Marshal(*d)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.Write(output)
}
