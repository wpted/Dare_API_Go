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

	// Establishing connection to database
	CurrentRepo, err = repositories.NewDareRepo(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	handler = controller.DareHandler{DareRepo: CurrentRepo}
}

func Run() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to the Dare API"})
	})
	router.POST("/Dare", handler.CreateDareHandler)
	router.GET("/Dare/:id", handler.GetDareHandler)
	router.GET("/Dares", handler.GetAllDaresHandler)
	router.PUT("/Dare/:id", handler.UpdateDareHandler)
	router.DELETE("/Dare/:id", handler.DeleteDareHandler)
	router.Run()
}
