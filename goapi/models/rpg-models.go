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
	PositionX     *int          `json:"positionX"`
	PositionY     *int          `json:"positionY"`
	LastKey       string        `json:"lastKey"`
	LastDirection string        `json:"lastDirection"`
	LastMoveTime  time.Time     `json:"lastMoveTime"`
	Wins          *[]Win        `json:"wins"`
	Steps         int           `json:"steps"`
	HP            int           `json:"hp"`
	HPTotal       int           `json:"hpTotal"`
}

type GameJoinMessage struct {
	Username      string `json:"username"`
	CharacterType string `json:"characterType,omitempty"`
}

type RPGMessage struct {
	Type WSMessageType   `json:"type"`
	Data json.RawMessage `json:"data"`
}

type RPGBroadcast struct {
	Type         WSMessageType `json:"type"`
	Players      []Player      `json:"players"`
	PlayersCount int           `json:"playersCount"`
}

func HasCollision(startXp1, startYp1, startXp2, startYp2, offset int) (bool, bool) {
	collisionX := false
	collisionY := false

	startXp1 -= offset
	startYp1 -= offset
	startXp2 -= offset
	startYp2 -= offset

	endXp1 := startXp1 + constants.CharacterSize + offset
	endYp1 := startYp1 + constants.CharacterSize + offset
	endXp2 := startXp2 + constants.CharacterSize + offset
	endYp2 := startYp2 + constants.CharacterSize + offset

	if startXp2 < endXp1 && endXp2 > startXp1 {
		collisionY = true
	}
	if startYp2 < endYp1 && endYp2 > startYp1 {
		collisionX = true
	}
	return collisionX, collisionY
}

func (p *Player) GetCollisionsTo(player2 Player, offset int) (bool, bool) {
	return HasCollision(
		*p.PositionX,
		*p.PositionY,
		*player2.PositionX,
		*player2.PositionY,
		offset,
	)
}
