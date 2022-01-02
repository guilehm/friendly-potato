package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"goapi/models"
	"goapi/utils"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const walkStep = 20
const borderOffset = 10
const characterSize = 50
const maxPosX = 1000/2 - characterSize*2 - borderOffset
const maxPosY = 800/2 - characterSize*2 - borderOffset
const minPosX = 0 + borderOffset
const minPosY = 40 + borderOffset

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

		switch message.MessageType {
		case models.GameJoin:
			data := models.GameJoinMessage{}
			err := json.Unmarshal(message.Data, &data)
			if err != nil {
				fmt.Println("Error during unmarshall:", err)
				break
			}

			rand.Seed(time.Now().Unix())
			ctChoices := []models.CharacterType{
				models.Human,
				models.Orc,
				models.Skeleton,
				models.Archer,
				models.Vampire,
				models.Berserker,
			}

			rand.Seed(time.Now().UnixNano())
			posX := rand.Intn(maxPosX-minPosX+1) + minPosX
			posY := rand.Intn(maxPosY-minPosY+1) + minPosY

			np := &models.Player{
				Type:      ctChoices[rand.Int()%len(ctChoices)],
				Username:  data.Username,
				PositionX: &posX,
				PositionY: &posY,
			}
			client = &models.Client{
				Hub:    hub,
				Conn:   conn,
				Send:   make(chan []byte, 256),
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
				*client.Player.PositionX = utils.Max(*client.Player.PositionX-walkStep, minPosX)
				client.Player.LastKey = key
				client.Player.LastDirection = key
			case models.ArrowUp:
				*client.Player.PositionY = utils.Max(*client.Player.PositionY-walkStep, minPosY)
				client.Player.LastKey = key
			case models.ArrowRight:
				*client.Player.PositionX = utils.Min(*client.Player.PositionX+walkStep, maxPosX)
				client.Player.LastKey = key
				client.Player.LastDirection = key
			case models.ArrowDown:
				*client.Player.PositionY = utils.Min(*client.Player.PositionY+walkStep, maxPosY)
				client.Player.LastKey = key
			}
			hub.Broadcast <- true
		}
	}
}
