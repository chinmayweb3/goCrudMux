package main

import (
	"fmt"
	"github/chinmay/gocrudmux/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.GetMovies).Methods("GET")
	router.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	router.HandleFunc("/movie", handlers.GetMovie).Methods("GET")
	router.HandleFunc("/addmovie", handlers.AddMovie).Methods("POST")
	router.HandleFunc("/deletemovie", handlers.DeleteMovie).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	if err == nil {
		fmt.Println("server is listening on :8080")
	} else {
		panic(err)
	}

}
