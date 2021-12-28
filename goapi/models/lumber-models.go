package models


type Wood struct {
	Name string
	Color string
}

type WoodPile struct {
	Woods []Wood
}

type Player struct {
	ID string
	Coins int
	Sprite string
}
