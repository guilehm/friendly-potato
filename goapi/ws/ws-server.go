package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"goapi/db"
	"goapi/models"
	"goapi/utils"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
)

var usersCollection = db.OpenCollection("users", "")
var lumberCollection = db.OpenCollection("lumber", "")
var upgrader = websocket.Upgrader{}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: do not allow all origins
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

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
				quit <- true
				fmt.Println("Connection closed")
				return
			} else {
				fmt.Println("Could not read message:", err)
				continue
			}
		}

		switch message.MessageType {
		case models.Login:
			data := models.Tokens{}
			err := json.Unmarshal(message.Data, &data)
			if err != nil {
				fmt.Println("Error during unmarshall:", err)
				break
			}

			go func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				var user models.User
				if err := usersCollection.FindOne(
					ctx, bson.M{"refresh_token": data.RefreshToken},
				).Decode(&user); err != nil {
					fmt.Println("Could not find user:", err)
					cancel()
					return
				}
				cancel()

				playerData, err := utils.GetPlayerData(user)

				um := models.UpdateMessage{
					Type:       models.Update,
					PlayerData: &playerData,
				}
				err = conn.WriteJSON(um)
				if err != nil {
					fmt.Println("Could not send update message for:", *user.Email)
				}

				go func() {
					time.Sleep(2 * time.Second)
					for {
						select {
						case <-quit:
							err := utils.UpdatePlayerData(&playerData, user)
							if err != nil {
								fmt.Println("Could not save game data for:", *user.Email, err)
								return
							}
							return
						default:
							time.Sleep(1 * time.Second)
							newWood := models.Wood{
								Name:        models.Oak,
								Color:       "",
								DateCreated: time.Now().Local(),
							}
							*playerData.Woods = append(*playerData.Woods, newWood)

							err = conn.WriteJSON(models.UpdateMessage{
								Type:       models.Update,
								PlayerData: &playerData,
							})
							if err != nil {
								fmt.Println("Could not send update message for:", *user.Email, err)
							}
						}
					}
				}()
			}()

		}
	}
}
