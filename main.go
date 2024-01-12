package main

import (
	"fmt"
	"github/chinmay/gocrudmux/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.GetRandom).Methods("GET")
	router.HandleFunc("/movies", handlers.GetMovies).Methods("GET")           // done
	router.HandleFunc("/movie", handlers.GetMovie).Methods("GET")             // done
	router.HandleFunc("/addmovie", handlers.AddMovie).Methods("POST")         // done
	router.HandleFunc("/deletemovie", handlers.DeleteMovie).Methods("DELETE") // done

	fmt.Println("server is listening on : 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}
