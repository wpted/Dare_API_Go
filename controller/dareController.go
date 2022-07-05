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
	{2, "Sing hello world out loud"},
	{3, "Dance hello world as if it's the end of the world"},
}

func Home(w http.ResponseWriter, r *http.Request) {
	message := struct {
		Message string `json:"Message"`
	}{"welcome to drunk dares"}
	js, err := json.Marshal(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// getRandomIndex is a helper function getting a random index within [0, len(dareList))
func getRandomIndex(d dareList) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(len(d))
}

func (d *dareList) GetRandomDare(w http.ResponseWriter, r *http.Request) {

	randomDare := (*d)[getRandomIndex(*d)].DareQuestion

	js, err := json.Marshal(randomDare)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (d *dareList) GetAllDare(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(*d)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
