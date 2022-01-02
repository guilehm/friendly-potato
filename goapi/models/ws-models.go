package models

import (
	"encoding/json"
	"fmt"
	"goapi/constants"

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
			}
		case <-h.Broadcast:
			var players []Player

			for client, _ := range h.Clients {

			Validation:
				for client2, _ := range h.Clients {
					if client.Player.Username == client2.Player.Username {
						continue Validation
					}

					cX, cY := client.Player.GetCollisions(*client2.Player)
					if cX && cY {
						if client.Player.LastMoveTime.Before(client2.Player.LastMoveTime) {
							*client.Player.Wins = append(*client.Player.Wins, Win{client2.Player.Username})

							x := constants.MinPosX
							x2 := constants.MaxPosX
							if *client.Player.PositionX > *client2.Player.PositionX {
								client.Player.PositionX = &x2
								client2.Player.PositionX = &x
							} else {
								client.Player.PositionX = &x
								client2.Player.PositionX = &x2
							}
						}
					}
				}
				players = append(players, *client.Player)
			}

			for client := range h.Clients {
				err := client.Conn.WriteJSON(RPGBroadcast{
					Broadcast,
					players,
					len(players),
				})
				if err != nil {
					fmt.Println("Could not send message:", err)
					return
				}
			}
		}
	}
}
