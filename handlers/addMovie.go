package handlers

import (
	"encoding/json"
	"github/chinmay/gocrudmux/database"
	"github/chinmay/gocrudmux/models"
	"net/http"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var movie models.IMovie
	json.NewDecoder(r.Body).Decode(&movie)

	database.DbMovies.AddMovie(movie) // add to database Collection

	json.NewEncoder(w).Encode(movie)

}
