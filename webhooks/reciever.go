package webhooks

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// RecieveHandler handles incoming webhook requests
func RecieveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Error decoding JSON payload", http.StatusBadRequest)
		return
	}

	log.Printf("Received webhook payload: %v", payload)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Webhook received successfully")
}