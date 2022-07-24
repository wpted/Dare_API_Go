package app

import (
	configs2 "dareAPI/configs"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
)

// Read the .env file first
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load .env file")
	}
}

func Run() {
	// Get URI from .env file
	config := configs2.NewMongoConfig()
	mongoURI := config.GetMongoURI()
	fmt.Println(mongoURI)

	//
}
