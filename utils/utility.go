package utils

import (
	"dareAPI/model"
	"encoding/json"
	"io/ioutil"
	"log"
)

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
