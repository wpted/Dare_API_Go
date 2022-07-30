package app

import (
	"dareAPI/configs"
	"dareAPI/controller"
	"dareAPI/repositories"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"net/http"
)

var (
	CurrentRepo *repositories.DareRepo
	handler     controller.DareHandler
	authHandler controller.AuthHandler
)

func init() {
	// Read the .env file first
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load .env file")
	}

	// Get URI from .env file
	config := configs.NewMongoConfig()
	mongoURI := config.GetMongoURI()
	mongoDatabaseName := "Dare_Mongo"
	mongoCollectionName := "Dares"
	// Establishing connection to database
	CurrentRepo, err = repositories.NewDareRepo(mongoURI, mongoDatabaseName, mongoCollectionName)
	if err != nil {
		log.Fatal(err)
	}

	handler = controller.DareHandler{DareRepo: CurrentRepo}
	authHandler = controller.AuthHandler{}
}

func Run() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to the Dare API"})
	})
	// Dare handler
	router.POST("/Dare", handler.CreateDareHandler)
	router.GET("/Dare/:id", handler.GetDareHandler)
	router.GET("/Dares", handler.GetAllDaresHandler)
	router.PUT("/Dare/:id", handler.UpdateDareHandler)
	router.DELETE("/Dare/:id", handler.DeleteDareHandler)

	// User Handler
	router.GET("/SignIn", authHandler.SignInHandler)
	router.Run()
}
