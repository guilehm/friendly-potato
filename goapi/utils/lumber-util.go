package utils

import (
	"context"
	"goapi/db"
	"goapi/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lumberCollection = db.OpenCollection("lumber", "")

func UpdatePlayerData(playerData *models.PlayerData, user models.User) error {
	upsert := true
	opt := options.UpdateOptions{Upsert: &upsert}
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := lumberCollection.UpdateOne(
		c,
		bson.M{"user_id": user.UserId},
		bson.M{
			"$set": bson.M{
				"gold":  playerData.Gold,
				"woods": playerData.Woods,
			},
		},
		&opt,
	)
	return err
}
