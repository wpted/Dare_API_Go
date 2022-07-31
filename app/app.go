package app

import (
	"dareAPI/configs"
	"dareAPI/controller"
	"dareAPI/repositories"
	"dareAPI/utils"
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
	mongoDatabaseName := config.Database
	mongoCollectionName := config.Collection

	// Establishing connection to database
	CurrentRepo, err = repositories.NewDareRepo(mongoURI, mongoDatabaseName, mongoCollectionName)
	if err != nil {
		log.Fatal(err)
	}

	// Check database content status
	if CurrentRepo.IsEmpty() {
		dares := utils.ReadFromJson("dare.json")
		for _, dare := range dares {
			CurrentRepo.CreateDare(&dare)
		}
	}

	// Initialize Administrator
	newAdmin := configs.NewAdmin()
	handler = controller.DareHandler{DareRepo: CurrentRepo}
	authHandler = controller.AuthHandler{Admin: newAdmin}
}

func Run() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to the Dare API"})
	})
	// Dare handler
	router.GET("/Dare/:id", handler.GetDareHandler)
	router.GET("/Dares", handler.GetAllDaresHandler)

	// Group handler func that implements authMiddleware
	needAuth := router.Group("/")
	needAuth.Use(controller.RequireLogin(authHandler.GetSecretKey()))
	{
		needAuth.POST("/Dare", handler.CreateDareHandler)
		needAuth.PUT("/Dare/:id", handler.UpdateDareHandler)
		needAuth.DELETE("/Dare/:id", handler.DeleteDareHandler)
	}
	// User Handler
	router.GET("/SignIn", authHandler.SignInHandler)
	router.Run()
}
