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
	Clients    *[]Client
	Broadcast  chan bool
	Register   chan *Client
	Unregister chan *Client
}

func (h *Hub) Start() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					fmt.Println("default for select. Closing client.send. Removing client from Clients.")
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
