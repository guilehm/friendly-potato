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

type PlayerData struct {
	UserId    string    `bson:"user_id" json:"user_id"`
	Coins     int       `bson:"coins" json:"coins"`
	Sprite    string    `bson:"sprite" json:"sprite"`
	LastLogin time.Time `bson:"last_login" json:"last_login"`
	Woods     *[]Wood   `bson:"woods" json:"woods"`
	conn      *websocket.Conn
}

type UpdateMessage struct {
	Type       WSMessageType `json:"type"`
	PlayerData *PlayerData   `json:"player_data"`
}
