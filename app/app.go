package app

import (
	"dareAPI/configs"
	"dareAPI/repositories"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
)

var CurrentRepo repositories.Repo

func init() {
	// Read the .env file first
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load .env file")
	}

	// Get URI from .env file
	config := configs.NewMongoConfig()
	mongoURI := config.GetMongoURI()

	// Establishing connection to database
	fmt.Print("Connecting to database...\n\n")
	CurrentRepo, err = repositories.NewDareRepo(mongoURI)

}

func Run() {
	fmt.Printf("%#v", CurrentRepo)
}
