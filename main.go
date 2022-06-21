package main

import (
	"dareAPI/api"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.Home)
	mux.HandleFunc("/drunk/dares", api.ListDares)
	//mux.HandleFunc("/drunk/dares", api.CreateDare)

	http.ListenAndServe("localhost:8080", mux)
}
