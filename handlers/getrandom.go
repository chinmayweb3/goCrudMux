package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func GetRandom(w http.ResponseWriter, r *http.Request) {
	url := "https://dummyjson.com/quotes/random"

	resp, err := http.Get(url)

	if err != nil || strings.HasPrefix(resp.Status, "4") {
		w.Write([]byte("No quotes found"))
		return
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data map[string]string
	json.Unmarshal(body, &data)

	// data2, _ := json.Marshal(map[string]string{"quotes": data["quote"]})
	// w.Write(data2)

	json.NewEncoder(w).Encode(map[string]string{"quote": data["quote"]})

}
