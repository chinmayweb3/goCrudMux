package handlers

import (
	"encoding/json"
	"github/chinmay/gocrudmux/database"
	"net/http"
	"strings"
)

func GetMovie(w http.ResponseWriter, r *http.Request) {

	names := r.URL.Query()["name"] // comes back in array format
	// ids := r.URL.Query()["id"]     // comes back in array format

	if len(names) <= 0 || strings.TrimSpace(names[0]) == "" {
		w.Write([]byte("no name available found "))
		return
	}

	moves := database.DbMovies.FindByName(names[0])
	b, _ := json.Marshal(moves)

	w.Write(b)
}
