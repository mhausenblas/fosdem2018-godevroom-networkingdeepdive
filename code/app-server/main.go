package main

import (
	"encoding/json"
	"net/http"
)

// BEGIN1 OMIT
// Conference defines an event.
type Conference struct {
	Title  string `json:"title"`
	NumPar int    `json:"count"`
}

func main() {
	fosdem18 := Conference{
		Title:  "FOSDEM 2018",
		NumPar: 3000,
	}
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(fosdem18)
	})
	_ = http.ListenAndServe(":9876", nil)
}

// END1 OMIT
