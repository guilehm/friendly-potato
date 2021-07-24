package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sale struct {
	ID            primitive.ObjectID `bson:"_id"`
	Region        string             `bson:"region"`
	Country       string             `bson:"country"`
	ItemType      string             `bson:"item_type"`
	SalesChannel  string             `bson:"sales_channel"`
	OrderPriority string             `bson:"order_priority"`
	OrderDate     time.Time          `bson:"order_date"`
	OrderId       string             `bson:"order_id"`
	ShipDate      time.Time          `bson:"ship_date"`
	UnitsSold     int                `bson:"units_sold"`
	UnitPrice     uint64             `bson:"unit_price"`
	UnitCost      uint64             `bson:"unit_cost"`
	TotalRevenue  uint64             `bson:"total_revenue"`
	TotalCost     uint64             `bson:"total_cost"`
	TotalProfit   uint64             `bson:"total_profi"`
}
