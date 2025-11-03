package main

import (
	"bytes"
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

// Main service that adds employees and sends webhooks
func main() {
	http.HandleFunc("/add-employee", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Employee added: %+v\n", emp)

		// Simulate sending a webhook
		go sendWebhook(emp)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Employee added successfully, webhook sent",
		})
	})

	log.Println("Main service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Simulate sending a webhook
func sendWebhook(emp Employee) {
	webhookURL := "http://localhost:8081/webhook" // receiver service endpoint

	payload, _ := json.Marshal(emp)
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Failed to send webhook: %v\n", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Webhook sent to %s, status: %s\n", webhookURL, resp.Status)
}
