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
				player := *client.Player

			Validation:
				for client2, _ := range h.Clients {
					player2 := *client2.Player
					if player.Username == player2.Username {
						continue Validation
					}

					cX, cY := player.GetCollisions(player2)
					if cX && cY {
						if player.LastMoveTime.Before(player2.LastMoveTime) {
							*player.Wins = append(*player.Wins, Win{player2.Username})
							x := constants.MinPosX
							x2 := constants.MaxPosX
							player.PositionX = &x
							player2.PositionX = &x2
						}
					}
				}
				players = append(players, player)
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
