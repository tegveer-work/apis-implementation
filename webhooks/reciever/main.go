package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// Main receiver service that listens for webhooks
func main() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Received webhook: %+v\n", emp)
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Receiver running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
