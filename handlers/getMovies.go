package handlers

import (
	"encoding/json"
	"github/chinmay/gocrudmux/database"
	"net/http"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	d, _ := json.Marshal(database.DbMovies)
	w.Write(d)
}
