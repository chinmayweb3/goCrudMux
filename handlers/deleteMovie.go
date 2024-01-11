package handlers

import (
	"encoding/json"
	"github/chinmay/gocrudmux/database"
	"net/http"
)

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var reqName map[string]string
	json.NewDecoder(r.Body).Decode(&reqName)

	getMovies := database.DbMovies.DeleteMovie(reqName["name"])

	if getMovies == -1 {
		json.NewEncoder(w).Encode(map[string]string{"name": reqName["name"], "message": "no movie was found"})

	} else {
		json.NewEncoder(w).Encode(map[string]string{"name": reqName["name"], "message": "movie deleted"})

	}
}
