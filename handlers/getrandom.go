package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func GetRandom(w http.ResponseWriter, r *http.Request) {
	url := "https://api.quotale.io/randofkfms"
	resp, err := http.Get(url)

	if err != nil || strings.HasPrefix(resp.Status, "4") {
		w.Write([]byte("No quotes found"))
		return
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data map[string]string
	json.Unmarshal(body, &data)

	w.Write([]byte(data["content"]))

}
