package main

import (
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn     *websocket.Conn
	username string
	active   bool
}

type Message struct {
	Sender    string    `json:"sender"`
	Recipient string    `json:"recipient"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type UserStatus struct {
	Username     string    `json:"username"`
	Online       bool      `json:"online"`
	LastActivity time.Time `json:"lastActivity"`
}

var (
	clients      = make(map[string]*Client)
	userStatuses = make(map[string]UserStatus)
	messages     = make(map[string][]Message) // key: "user1:user2"
	mutex        = &sync.Mutex{}
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		conn.WriteMessage(websocket.CloseMessage, []byte("Username is required"))
		conn.Close()
		return
	}

	mutex.Lock()
	// Check if username already exists
	if existing, exists := clients[username]; exists {
		existing.conn.Close()
		delete(clients, username)
	}

	client := &Client{
		conn:     conn,
		username: username,
		active:   true,
	}
	clients[username] = client

	// Update user status
	userStatuses[username] = UserStatus{
		Username:     username,
		Online:       true,
		LastActivity: time.Now(),
	}
	mutex.Unlock()

	log.Printf("Client connected: %s", username)

	// Send initial data to client
	sendInitialData(client)

	go handleClient(client)
}

func handleClient(client *Client) {
	defer func() {
		mutex.Lock()
		client.conn.Close()
		delete(clients, client.username)
		// Update user status to offline but keep last activity
		if status, exists := userStatuses[client.username]; exists {
			userStatuses[client.username] = UserStatus{
				Username:     client.username,
				Online:       false,
				LastActivity: status.LastActivity,
			}
		}
		mutex.Unlock()
		broadcastUserStatuses()
		log.Printf("Client disconnected: %s", client.username)
	}()

	for {
		var msg Message
		err := client.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Read error: %v", err)
			}
			break
		}

		msg.Sender = client.username
		msg.Timestamp = time.Now()

		processMessage(msg)
	}
}

func processMessage(msg Message) {
	mutex.Lock()
	defer mutex.Unlock()

	// Update sender's last activity
	if _, exists := userStatuses[msg.Sender]; exists {
		userStatuses[msg.Sender] = UserStatus{
			Username:     msg.Sender,
			Online:       true,
			LastActivity: msg.Timestamp,
		}
	}

	// Store message in both directions for conversation history
	conversationKey1 := getConversationKey(msg.Sender, msg.Recipient)
	conversationKey2 := getConversationKey(msg.Recipient, msg.Sender)

	messages[conversationKey1] = append(messages[conversationKey1], msg)
	messages[conversationKey2] = append(messages[conversationKey2], msg)

	// Send message to recipient if online
	if recipient, exists := clients[msg.Recipient]; exists {
		if err := recipient.conn.WriteJSON(msg); err != nil {
			log.Printf("Error sending message to %s: %v", msg.Recipient, err)
		}
	}

	// Broadcast updated user statuses to all clients
	broadcastUserStatuses()
}

func getConversationKey(user1, user2 string) string {
	if user1 < user2 {
		return user1 + ":" + user2
	}
	return user2 + ":" + user1
}

func sendInitialData(client *Client) {
	mutex.Lock()
	defer mutex.Unlock()

	// Send user's conversation partners with last messages
	partners := make(map[string]Message)
	for key, msgs := range messages {
		user1, user2 := parseConversationKey(key)
		var partner string
		if user1 == client.username {
			partner = user2
		} else if user2 == client.username {
			partner = user1
		} else {
			continue
		}

		if len(msgs) > 0 {
			partners[partner] = msgs[len(msgs)-1]
		}
	}

	// Prepare initial data
	initialData := struct {
		Type     string             `json:"type"`
		User     string             `json:"user"`
		Partners map[string]Message `json:"partners"`
		Statuses []UserStatus       `json:"statuses"`
	}{
		Type:     "initial",
		User:     client.username,
		Partners: partners,
		Statuses: getUserStatusesSorted(),
	}

	if err := client.conn.WriteJSON(initialData); err != nil {
		log.Printf("Error sending initial data to %s: %v", client.username, err)
	}
}

func broadcastUserStatuses() {
	statuses := getUserStatusesSorted()
	for username, client := range clients {
		data := struct {
			Type     string       `json:"type"`
			Statuses []UserStatus `json:"statuses"`
		}{
			Type:     "status_update",
			Statuses: statuses,
		}

		if err := client.conn.WriteJSON(data); err != nil {
			log.Printf("Error sending status update to %s: %v", username, err)
		}
	}
}

func getUserStatusesSorted() []UserStatus {
	statuses := make([]UserStatus, 0, len(userStatuses))
	for _, status := range userStatuses {
		statuses = append(statuses, status)
	}

	// Sort by last activity (newest first), then alphabetically
	sort.Slice(statuses, func(i, j int) bool {
		if statuses[i].LastActivity.Equal(statuses[j].LastActivity) {
			return statuses[i].Username < statuses[j].Username
		}
		return statuses[i].LastActivity.After(statuses[j].LastActivity)
	})

	return statuses
}

func parseConversationKey(key string) (string, string) {
	// Assuming key is in format "user1:user2"
	// Implementation depends on your key format
	// This is a simplified version
	for i := 0; i < len(key); i++ {
		if key[i] == ':' {
			return key[:i], key[i+1:]
		}
	}
	return "", ""
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
