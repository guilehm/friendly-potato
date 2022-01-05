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

type Area struct {
	PosStartX int
	PosEndX   int
	PosStartY int
	PosEndY   int
}

func (h *Hub) FindEmptyAreas(checkStep int) []Area {
	var emptyAreas []Area
	posX := constants.MinPosX
	for posY := constants.MinPosY; posY <= constants.MaxPosY; posY += constants.WalkStep {
		empty := true
		for client := range h.Clients {
			cX, cY := HasCollision(
				posX,
				posY,
				*client.Player.PositionX,
				*client.Player.PositionY,
				constants.SecurityOffset,
			)
			if cX && cY {
				empty = false
			}
		}
		if empty {
			emptyArea := Area{
				PosStartX: posX,
				PosEndX:   posX + constants.CharacterSize,
				PosStartY: posY,
				PosEndY:   posY + constants.CharacterSize,
			}
			emptyAreas = append(emptyAreas, emptyArea)
		}
		if posY == constants.MaxPosY && (posX+checkStep) <= constants.MaxPosX {
			posY = constants.MinPosY
			posX += checkStep
		}
	}
	return emptyAreas
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

			FindWinner:
				for client2, _ := range h.Clients {
					if client.Player.Username == client2.Player.Username {
						continue FindWinner
					}

					cX, cY := client.Player.GetCollisions(*client2.Player, 0)
					if cX && cY {
						if client.Player.LastMoveTime.After(client2.Player.LastMoveTime) {
							*client.Player.Wins = append(*client.Player.Wins, Win{client2.Player.Username})
						} else {
							*client2.Player.Wins = append(*client2.Player.Wins, Win{client.Player.Username})
						}

						emptyAreas := h.FindEmptyAreas(constants.WalkStep + 2)

						if len(emptyAreas) < 2 {
							// TODO: Fix this condition. Max players exceeded.
							continue FindWinner
						}
						// Reposition collided players
						first := emptyAreas[0]
						last := emptyAreas[len(emptyAreas)-1]
						*client.Player.PositionX = first.PosStartX
						*client.Player.PositionY = first.PosStartY
						*client2.Player.PositionX = last.PosStartX
						*client2.Player.PositionY = last.PosStartY
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
