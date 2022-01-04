package models

import (
	"encoding/json"
	"goapi/constants"
	"time"
)

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
	Mage      CharacterType = "mage"
	Orc       CharacterType = "orc"
	Skeleton  CharacterType = "skeleton"
	Archer    CharacterType = "archer"
	Vampire   CharacterType = "vampire"
	Berserker CharacterType = "berserker"
)

type Win struct {
	Defeated string `json:"defeated"`
}

type Player struct {
	Type          CharacterType `json:"type"`
	Username      string        `json:"username"`
	PositionX     *int          `json:"position_x"`
	PositionY     *int          `json:"position_y"`
	LastKey       string        `json:"last_key"`
	LastDirection string        `json:"last_direction"`
	LastMoveTime  time.Time     `json:"last_move_time"`
	Wins          *[]Win        `json:"wins"`
	Steps         int           `json:"steps"`
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

func (p *Player) GetCollisions(player2 Player, offset int) (bool, bool) {
	collisionX := false
	collisionY := false

	startXp1 := *p.PositionX - offset
	endXp1 := startXp1 + constants.CharacterSize + offset
	startYp1 := *p.PositionY - offset
	endYp1 := startYp1 + constants.CharacterSize + offset

	startXp2 := *player2.PositionX - offset
	endXp2 := startXp2 + constants.CharacterSize + offset
	startYp2 := *player2.PositionY - offset
	endYp2 := startYp2 + constants.CharacterSize + offset

	if startXp2 < endXp1 && endXp2 > startXp1 {
		collisionY = true
	}

	if startYp2 < endYp1 && endYp2 > startYp1 {
		collisionX = true
	}
	return collisionX, collisionY
}
