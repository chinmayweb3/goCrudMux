package handlers

import "net/http"

func GetMovies(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("sdflk"))
}
