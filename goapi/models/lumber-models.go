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
	Coins int
	Sprite string
	LastLogin time.Time
}
