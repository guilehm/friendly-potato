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
	UnitsSold     int
	UnitPrice     uint64
	UnitCost      uint64
	TotalRevenue  uint64
	TotalCost     uint64
	TotalProfit   uint64
}
