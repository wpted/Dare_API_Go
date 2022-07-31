package utils

import (
	"dareAPI/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

// ReadFromJson reads from the given json file
func ReadFromJson(fileName string) model.DareContainer {
	var dares model.DareContainer
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &dares)
	if err != nil {
		log.Fatal(err)
	}

	return dares
}

// RandomDare gives a random dare from the dare container
func RandomDare(container model.DareContainer) model.Dare {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(container), func(i, j int) {
		container[i], container[j] = container[j], container[i]
	})
	return container[0]
}
