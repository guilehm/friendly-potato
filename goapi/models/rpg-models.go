package models

import "encoding/json"

type CharacterType string

var (
	GameJoin  WSMessageType = "game-join"
	Broadcast WSMessageType = "broadcast"
	KeyDown   WSMessageType = "key-down"
)

var (
	ArrowLeft  = "ArrowLeft"
	ArrowUp    = "ArrowUp"
	ArrowRight = "ArrowRight"
	ArrowDown  = "ArrowDown"
)

var (
	Human     CharacterType = "human"
	Orc       CharacterType = "orc"
	Skeleton  CharacterType = "skeleton"
	Archer    CharacterType = "archer"
	Vampire   CharacterType = "vampire"
	Berserker CharacterType = "berserker"
)

type Player struct {
	Type          CharacterType `json:"type"`
	Username      string        `json:"username"`
	PositionX     *int          `json:"position_x"`
	PositionY     *int          `json:"position_y"`
	LastKey       string        `json:"last_key"`
	LastDirection string        `json:"last_direction"`
}

type GameJoinMessage struct {
	Username      string `json:"username"`
	CharacterType string `json:"character_type,omitempty"`
}

type RPGMessage struct {
	Type WSMessageType   `json:"type"`
	Data json.RawMessage `json:"data"`
}

type RPGBroadcast struct {
	Type        WSMessageType `json:"type"`
	Players     []Player      `json:"players"`
	PlayerCount int           `json:"player_count"`
}
