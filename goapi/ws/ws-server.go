package ws

import (
	"encoding/json"
	"fmt"
	"goapi/models"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: do not allow all origins
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

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
			fmt.Println("Error while reading json:", err)
			break
		}

		switch message.MessageType {
		case models.Login:
			data := models.Tokens{}
			err := json.Unmarshal(message.Data, &data)
			if err != nil {
				fmt.Println("Error during unmarshall:", err)
				break
			}

			err = conn.WriteMessage(websocket.TextMessage, []byte("Authenticated!!!"))
			if err != nil {
				fmt.Println("Error during message writing:", err)
				break
			}
			fmt.Println("successfully unmarshalled", data.RefreshToken, data.Token)
		}
	}
}
