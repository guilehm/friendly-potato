package models


type Wood struct {
	Name string
	Color string
}

type WoodPile struct {
	Woods []Wood
}
