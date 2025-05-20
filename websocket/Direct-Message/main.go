package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for development
	},
}

type Client struct {
	conn     *websocket.Conn
	send     chan []byte
	username string
}

type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

var (
	clients    = make(map[string]*Client)
	clientsMux sync.Mutex
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	// Read username from query parameter
	username := r.URL.Query().Get("username")
	if username == "" {
		conn.WriteMessage(websocket.CloseMessage, []byte("Username is required"))
		conn.Close()
		return
	}

	client := &Client{
		conn:     conn,
		send:     make(chan []byte, 256),
		username: username,
	}

	// Register client
	clientsMux.Lock()
	clients[username] = client
	clientsMux.Unlock()

	log.Printf("Client connected: %s", username)

	// Start goroutines for reading and writing
	go client.writePump()
	go client.readPump()
}

func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
		clientsMux.Lock()
		delete(clients, c.username)
		clientsMux.Unlock()
		log.Printf("Client disconnected: %s", c.username)
	}()

	for {
		var msg Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Read error: %v", err)
			}
			break
		}

		// Set the sender (in case client tries to spoof)
		msg.Sender = c.username

		log.Printf("Received message from %s to %s: %s", msg.Sender, msg.Recipient, msg.Content)

		// Send to recipient if they're connected
		clientsMux.Lock()
		recipient, ok := clients[msg.Recipient]
		clientsMux.Unlock()

		if ok {
			recipient.send <- []byte(msg.Content)
		} else {
			log.Printf("Recipient %s not found", msg.Recipient)
			// Optionally notify sender that recipient is offline
			c.send <- []byte("User " + msg.Recipient + " is offline")
		}
	}
}

func (c *Client) writePump() {
	defer c.conn.Close()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// Channel closed
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}