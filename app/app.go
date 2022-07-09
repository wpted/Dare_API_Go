package app

import (
	"dareAPI/controller"
	"net/http"
)

func Run() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/RandomDare", controller.MockDareList.GetRandomDare)
	http.HandleFunc("/Dares", controller.MockDareList.GetAllDare)
	// if we want to use same endpoint names with different http methods, we have to create a different mux (multiplexer)
	// so here we use different endpoint names, and should modify in the future
	http.HandleFunc("/CreateDare", controller.CreateDare)
	//http.HandleFunc("/UpdateDare", controller.UpdateDare)

	http.ListenAndServe(":8080", nil)
}
