package models

import (
	"time"

	"github.com/gorilla/websocket"
)

var (
	Login  WSMessageType = "login"
	Update WSMessageType = "update"
)

type WoodName string

var (
	Oak WoodName = "oak"
)

type Wood struct {
	Name        WoodName  `bson:"name" json:"name"`
	Color       string    `bson:"color" json:"color"`
	DateCreated time.Time `bson:"date_created" json:"date_created"`
}

type PlayerData struct {
	UserId    string    `bson:"user_id" json:"user_id"`
	Gold      int       `bson:"gold" json:"gold"`
	Sprite    string    `bson:"sprite" json:"sprite"`
	LastLogin time.Time `bson:"last_login" json:"last_login"`
	Woods     *[]Wood   `bson:"woods" json:"woods"`
	conn      *websocket.Conn
}

type UpdateMessage struct {
	Type       WSMessageType `json:"type"`
	PlayerData *PlayerData   `json:"player_data"`
}
