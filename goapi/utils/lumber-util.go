package utils

import (
	"context"
	"errors"
	"fmt"
	"goapi/db"
	"goapi/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

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

func GetPlayerData(user models.User) (models.PlayerData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	playerData := models.PlayerData{}
	err := lumberCollection.FindOne(ctx, bson.M{"user_id": user.UserId}).Decode(&playerData)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println("Error while trying to find player:", err)
			return playerData, err
		} else {
			playerData = models.PlayerData{
				UserId:    user.UserId,
				Gold:      0,
				Sprite:    "",
				LastLogin: time.Now().Local(),
				Woods:     &[]models.Wood{},
			}
			_, err := lumberCollection.InsertOne(ctx, playerData)
			if err != nil {
				fmt.Println("Could not create user data:", err)
				return playerData, err
			}
			fmt.Println("Player data created for", *user.Email)
		}
	}
	return playerData, err
}
