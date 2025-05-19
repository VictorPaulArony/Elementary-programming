package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", handleWebsocket)
	fmt.Println("Server started at: http://localhost:1234")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

// upgrade http connection to websocket
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all connections
	},
}

// function to handle the websocket connections
func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	// upgrade to websocket
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrading http error")
		return
	}
	defer conn.Close()

	for {
		// reading sms from client
		smstype, sms, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading the sms")
			break
		}
		fmt.Printf("Received: %s\n", sms)

		// Echo sms back to client
		err = conn.WriteMessage(smstype, sms)
		if err != nil {
			fmt.Println("Write Error")
			break
		}
	}
}
