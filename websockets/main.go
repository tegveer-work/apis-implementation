package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader upgrades HTTP → WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins (for local testing)
		return true
	},
}

var clients []*websocket.Conn

// WebSocket handler
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open WebSocket connection", http.StatusBadRequest)
		return
	}
	clients = append(clients, conn)
	fmt.Println("New client connected")

	for {
		// Wait for messages from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected:", err)
			break
		}
		fmt.Println("Received:", string(msg))

		// Broadcast to all clients
		for _, c := range clients {
			c.WriteMessage(websocket.TextMessage, []byte("Update: "+string(msg)))
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("WebSocket server running on ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
