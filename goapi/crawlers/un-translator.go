package crawlers

import (
	"context"
	"goapi/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (c UNCrawler) Translate() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var response models.CrawlResponse
	responseCollection.FindOne(
		ctx, bson.M{"translated": false},
	).Decode(&response)

	return nil
}
