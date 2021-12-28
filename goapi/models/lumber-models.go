package models

import "time"

type Wood struct {
	Name string
	Color string
	DateCreated time.Time
}

type WoodPile struct {
	Woods []Wood
	DateCreated time.Time
}

type Player struct {
	ID string
	Username string
	Coins int
	Sprite string
	LastLogin time.Time
	WoodPile WoodPile
}
