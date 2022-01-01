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
			fmt.Println("registering client", client.Player.Username)
			*h.Clients = append(*h.Clients, *client)
		case client := <-h.Unregister:
			// TODO: remove client from hub
			fmt.Println("this client should be removed", client.Player.Username)
		case <-h.Broadcast:
			var players []Player
			for _, client := range *h.Clients {
				players = append(players, *client.Player)
			}

			for _, client := range *h.Clients {
				err := client.Conn.WriteJSON(RPGBroadcast{Broadcast, players})
				if err != nil {
					fmt.Println("Could not send message:", err)
					return
				}
			}
		}
	}
}
