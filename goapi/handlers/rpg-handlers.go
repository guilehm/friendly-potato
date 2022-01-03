package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/constants"
	"goapi/models"
	"goapi/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func RPGHandler(hub *models.Hub, w http.ResponseWriter, r *http.Request) {
	// TODO: do not allow all origins
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	client := &models.Client{}

	quit := make(chan bool)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		message := models.WSMessage{}
		err = conn.ReadJSON(&message)
		if err != nil {
			if errors.Is(err.(*websocket.CloseError), err) {
				fmt.Println("Connection closed")
				quit <- true
				// TODO: the client must be removed from hub here
				return
			} else {
				fmt.Println("Could not read message:", err)
				continue
			}
		}

		now := time.Now()
		switch message.MessageType {
		case models.GameJoin:
			data := models.GameJoinMessage{}
			err := json.Unmarshal(message.Data, &data)
			if err != nil {
				fmt.Println("Error during unmarshall:", err)
				break
			}

			ctChoices := []models.CharacterType{
				models.Human,
				models.Orc,
				models.Skeleton,
				models.Archer,
				models.Vampire,
				models.Berserker,
			}

			rand.Seed(time.Now().UnixNano())
			posX := rand.Intn(constants.MaxPosX-constants.MinPosX+1) + constants.MinPosX
			posY := rand.Intn(constants.MaxPosY-constants.MinPosY+1) + constants.MinPosY

			username := data.Username
			for c, _ := range hub.Clients {
				if data.Username == c.Player.Username {
					randomId := strconv.Itoa(rand.Intn(999) + 100)
					username = data.Username + "_" + randomId
				}
			}

			np := &models.Player{
				Type:         ctChoices[rand.Int()%len(ctChoices)],
				Username:     username,
				PositionX:    &posX,
				PositionY:    &posY,
				LastMoveTime: now,
				Wins:         &[]models.Win{},
			}
			client = &models.Client{
				Hub:    hub,
				Conn:   conn,
				Player: np,
			}

			hub.Register <- client
			hub.Broadcast <- true

			// TODO: do I need a go func here?
			go func() {
				fmt.Println("Starting a new go routine to listen quit channel")
				for {
					select {
					case <-quit:
						hub.Unregister <- client
						hub.Broadcast <- true
					}
				}
			}()
		case models.KeyDown:
			key := ""
			err := json.Unmarshal(message.Data, &key)
			if err != nil {
				return
			}

			switch key {
			case models.ArrowLeft:
				oldPosition := *client.Player.PositionX
				*client.Player.PositionX = utils.Max(*client.Player.PositionX-constants.WalkStep, constants.MinPosX)
				if oldPosition != *client.Player.PositionX {
					client.Player.Steps += 1
				}
				client.Player.LastKey = key
				client.Player.LastDirection = key
			case models.ArrowUp:
				oldPosition := *client.Player.PositionY
				*client.Player.PositionY = utils.Max(*client.Player.PositionY-constants.WalkStep, constants.MinPosY)
				if oldPosition != *client.Player.PositionY {
					client.Player.Steps += 1
				}
				client.Player.LastKey = key
			case models.ArrowRight:
				oldPosition := *client.Player.PositionX
				*client.Player.PositionX = utils.Min(*client.Player.PositionX+constants.WalkStep, constants.MaxPosX)
				if oldPosition != *client.Player.PositionX {
					client.Player.Steps += 1
				}
				client.Player.LastKey = key
				client.Player.LastDirection = key
			case models.ArrowDown:
				oldPosition := *client.Player.PositionY
				*client.Player.PositionY = utils.Min(*client.Player.PositionY+constants.WalkStep, constants.MaxPosY)
				if oldPosition != *client.Player.PositionY {
					client.Player.Steps += 1
				}
				client.Player.LastKey = key
			}

			client.Player.LastMoveTime = now
			hub.Broadcast <- true
		}
	}
}
