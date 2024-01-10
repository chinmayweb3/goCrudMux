package handlers

import (
	"encoding/json"
	"fmt"
	"github/chinmay/gocrudmux/database"
	"github/chinmay/gocrudmux/models"
	"net/http"
	"strings"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var movie models.IMovie
	json.NewDecoder(r.Body).Decode(&movie)

	fmt.Println("#Movie", strings.TrimSpace(movie.Name) == "", "id")
	if strings.TrimSpace(movie.Name) == "" || movie.Id <= 0 || movie.Price <= 0 {
		w.Write([]byte("some data are missing"))
		return
	}

	database.DbMovies.AddMovie(movie) // add to database Collection

	json.NewEncoder(w).Encode(movie)

}
