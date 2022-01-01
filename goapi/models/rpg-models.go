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
	Type      CharacterType
	Username  string
	PositionX int
	PositionY int
}
