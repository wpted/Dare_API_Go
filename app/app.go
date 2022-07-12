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
	// Connect to database
	postgresConfig := configs.GetPostgresConfig()
	db, err := gorm.Open(postgresConfig.Dialect(), postgresConfig.GetConnectionURI())
	fmt.Println(postgresConfig.GetConnectionURI())
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("Successfully connect to database!")
	}

	// Migration of structs and the database
	// This should be done only once
	db.AutoMigrate(&model.Dare{})
	defer db.Close()

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
