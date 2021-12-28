package models

import "time"

type Wood struct {
	Name string
	Color string
}

type WoodPile struct {
	Woods []Wood
}

type Player struct {
	ID string
	Username string
	Coins int
	Sprite string
	LastLogin time.Time
}
