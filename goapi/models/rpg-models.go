package models

import "encoding/json"

type CharacterType string

var (
	GameJoin        WSMessageType = "game-join"
	LoginSuccessful WSMessageType = "login-successful"
)

var (
	Human    CharacterType = "human"
	Orc      CharacterType = "orc"
	Skeleton CharacterType = "skeleton"
	Archer   CharacterType = "archer"
)

type Player struct {
	Type      CharacterType `json:"type"`
	Username  string        `json:"username"`
	PositionX int           `json:"position_x"`
	PositionY int           `json:"position_y"`
}

type GameJoinMessage struct {
	Username      string `json:"username"`
	CharacterType string `json:"character_type,omitempty"`
}

type RPGMessage struct {
	Type WSMessageType   `json:"type"`
	Data json.RawMessage `json:"data"`
}
