package app

import (
	"dareAPI/configs"
	"dareAPI/controller"
	"dareAPI/middlewares"
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
	router.Use(middlewares.CORSMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to the Dare API"})
	})
	// Dare handler
	router.GET("/dare/:id", handler.GetDareHandler)
	router.GET("/dares", handler.GetAllDaresHandler)
	router.GET("/dares/random", handler.GetRandomDareHandler)

	// Group handler func that implements authMiddleware
	needAuth := router.Group("/")
	needAuth.Use(middlewares.RequireLogin(authHandler.GetSecretKey()))
	{
		needAuth.POST("/dare", handler.CreateDareHandler)
		needAuth.PUT("/dare/:id", handler.UpdateDareHandler)
		needAuth.DELETE("/dare/:id", handler.DeleteDareHandler)
	}
	// User Handler
	router.GET("/signin", authHandler.SignInHandler)
	router.Run()
}
