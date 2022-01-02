package models

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type WSMessageType string

type WSMessage struct {
	MessageType WSMessageType   `json:"type"`
	Data        json.RawMessage `json:"data"`
}

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
	Player *Player `json:"player"`
}

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan bool
	Register   chan *Client
	Unregister chan *Client
}

func (h *Hub) Start() {
	for {
		select {
		case client := <-h.Register:
			fmt.Println("registering client", client.Player.Username)
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case <-h.Broadcast:
			var players []Player
			for client := range h.Clients {
				players = append(players, *client.Player)
			}

			for client := range h.Clients {
				err := client.Conn.WriteJSON(RPGBroadcast{Broadcast, players})
				if err != nil {
					fmt.Println("Could not send message:", err)
					return
				}
			}
		}
	}
}
