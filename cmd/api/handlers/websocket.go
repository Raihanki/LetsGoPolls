package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool)

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error while upgrading connection: %v", err)
		http.Error(w, "could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		// Read message from WebSocket connection (to keep the connection alive)
		_, _, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			break
		}
	}
}

func BroadcastMessage(message []byte) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("error while sending message: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
