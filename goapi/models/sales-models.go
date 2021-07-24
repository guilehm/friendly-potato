package models

import "time"

type Sale struct {
	Region        string
	Country       string
	ItemType      string
	SalesChannel  string
	OrderPriority string
	OrderDate     time.Time
	OrderId       string
	ShipDate      time.Time
}
