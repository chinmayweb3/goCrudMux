package handlers

import (
	"net/http"
)

func GetRandom(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("running well "))

}
