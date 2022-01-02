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

func (p *Player) GetCollisions(player2 Player) (bool, bool) {
	collisionX := false
	collisionY := false

	startXp1 := *p.PositionX
	endXp1 := startXp1 + constants.CharacterSize
	startYp1 := *p.PositionY
	endYp1 := startYp1 + constants.CharacterSize

	startXp2 := *player2.PositionX
	endXp2 := startXp2 + constants.CharacterSize
	startYp2 := *player2.PositionY
	endYp2 := startYp2 + constants.CharacterSize

	if startXp2 < endXp1 && endXp2 > startXp1 {
		collisionY = true
	}

	if startYp2 < endYp1 && endYp2 > startYp1 {
		collisionX = true
	}
	return collisionX, collisionY
}
