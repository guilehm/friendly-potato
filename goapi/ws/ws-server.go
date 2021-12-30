package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"goapi/db"
	"goapi/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"

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

				playerData := models.PlayerData{}
				err := lumberCollection.FindOne(ctx, bson.M{"user_id": user.UserId}).Decode(&playerData)
				if err != nil {
					if !errors.Is(err, mongo.ErrNoDocuments) {
						fmt.Println("Error while trying to find player:", err)
					} else {
						playerData = models.PlayerData{
							UserId:    user.UserId,
							Coins:     0,
							Sprite:    "",
							LastLogin: time.Now().Local(),
							Woods:     &[]models.Wood{},
						}
						_, err := lumberCollection.InsertOne(ctx, playerData)
						if err != nil {
							fmt.Println("Could not create user data:", err)
							cancel()
							return
						}
						fmt.Println("Player data created for", *user.Email)
					}
					cancel()
					return
				}
				cancel()
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
							upsert := true
							opt := options.UpdateOptions{Upsert: &upsert}
							c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
							_, err = lumberCollection.UpdateOne(
								c,
								bson.M{"user_id": user.UserId},
								bson.M{
									"$set": bson.M{
										"coins": playerData.Coins,
										"woods": playerData.Woods,
									},
								},
								&opt,
							)
							if err != nil {
								fmt.Println("Could not save game data for:", *user.Email, err)
							} else {
								fmt.Println("Successfully saved game data for:", *user.Email)
							}
							cancel()
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
