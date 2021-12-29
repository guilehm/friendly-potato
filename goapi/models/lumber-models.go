package models

import (
	"time"

	"github.com/gorilla/websocket"
)

var (
	Login  WSMessageType = "login"
	Update WSMessageType = "update"
)

type Wood struct {
	Name        string
	Color       string
	DateCreated time.Time
}

type WoodPile struct {
	Woods       []Wood
	DateCreated time.Time
}

type PlayerData struct {
	UserId    string    `bson:"id" json:"id"`
	Coins     int       `bson:"coins "json:"coins"`
	Sprite    string    `bson:"sprite" json:"sprite"`
	LastLogin time.Time `bson:"last_login" json:"last_login"`
	WoodPile  WoodPile  `bson:"wood_pile" json:"wood_pile"`
	conn      *websocket.Conn
}

type UpdateMessage struct {
}
