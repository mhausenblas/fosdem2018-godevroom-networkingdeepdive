package main

import (
	"encoding/json"
	"net/http"
)

// Workshop defines an event.
type Workshop struct {
	Title        string `json:"title"`
	Participants int    `json:"count"`
}

func main() {
	lgworkshop := Workshop{
		Title:        "Let’s Go! Using the Go programming language for system tasks",
		Participants: 22,
	}
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(lgworkshop)
	})
	_ = http.ListenAndServe(":9876", nil)
}
