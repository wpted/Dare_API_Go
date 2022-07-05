package app

import (
	"dareAPI/controller"
	"net/http"
)

func Run() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/RandomDare", controller.MockDareList.GetRandomDare)
	http.HandleFunc("/Dares", controller.MockDareList.GetAllDare)

	http.ListenAndServe(":8080", nil)
}
