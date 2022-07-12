package app

import (
	"dareAPI/configs"
	"dareAPI/controller"
	"dareAPI/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"net/http"
)

// Read the .env file first
func init() {
	godotenv.Load()
}

func Run() {
	// after reading from the .env in the init function
	// we can successfully get the organized configurations of the database
	dbConfig := configs.GetPostgresConfig()
	db, err := gorm.Open(dbConfig.DbDialect(), dbConfig.GetConnectionURI())
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}
	// defer here closes the database after all threads in the main goroutine is done running
	defer db.Close()

	// AutoMigrate connects the database to the struct we've created
	// It will add fields that don't exist as a table
	// This should only be done once
	db.AutoMigrate(&model.Dare{})

	// Router handling the endpoints
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/RandomDare", controller.MockDareList.GetRandomDare)
	http.HandleFunc("/Dares", controller.MockDareList.GetAllDare)
	// if we want to use same endpoint names with different http methods, we have to create a different mux (multiplexer)
	// so here we use different endpoint names, and should modify in the future
	http.HandleFunc("/CreateDare", controller.CreateDare)
	//http.HandleFunc("/UpdateDare", controller.UpdateDare)

	http.ListenAndServe(":8080", nil)
}
