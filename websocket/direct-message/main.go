package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

type Client struct {
	ID   string
	Conn *websocket.Conn
}

var (
	clients   = make(map[string]*Client)
	clientsMu sync.RWMutex
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Get user ID from query
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	client := &Client{
		ID:   id,
		Conn: conn,
	}

	// Register client
	clientsMu.Lock()
	clients[id] = client
	clientsMu.Unlock()

	fmt.Printf("Client connected: %s\n", id)

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Client %s disconnected\n", id)
			break
		}

		var msg Message
		if err := json.Unmarshal(data, &msg); err != nil {
			fmt.Println("Invalid message format:", err)
			continue
		}

		// Lookup recipient
		clientsMu.RLock()
		recipient, ok := clients[msg.To]
		clientsMu.RUnlock()

		if ok {
			// Forward message to recipient
			outgoing, _ := json.Marshal(msg)
			err := recipient.Conn.WriteMessage(websocket.TextMessage, outgoing)
			if err != nil {
				fmt.Printf("Error sending to %s: %v\n", msg.To, err)
			}
		} else {
			fmt.Printf("User %s not connected\n", msg.To)
		}
	}

	// Cleanup
	clientsMu.Lock()
	delete(clients, id)
	clientsMu.Unlock()
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("Server started at: http://localhost:1234")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
