package models

type CharacterType string

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
