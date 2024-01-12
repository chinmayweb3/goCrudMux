package main_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestGoRandom(t *testing.T) {

	resp, err := http.Get("http://localhost:8080/")

	if err != nil {
		t.Error("error on the request")
	}

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data map[string]string
	json.Unmarshal(body, &data)

	if data["quote"] == "" {
		t.Error("no data for quote found in the response")

	}

}
