package main

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Work      string `json:"work"`
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allow", http.StatusMethodNotAllowed)
		return
	}

	var client Client

	// Retorna um slice de bytes ([]byte, error).
	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Pega struct e convert para json
	jsonParse, err := json.Marshal(client)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonParse)

}
