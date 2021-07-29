package crawlers

import (
	"context"
	"fmt"
	"goapi/models"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"go.mongodb.org/mongo-driver/bson"
)

func (c UNCrawler) Translate() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var response models.CrawlResponse
	responseCollection.FindOne(
		ctx, bson.M{"translated": false},
	).Decode(&response)

	body := response.Body

	doc, err := htmlquery.Parse(strings.NewReader(body))
	if err != nil {
		fmt.Println("could not parse body for", response.Url)
		return err
	}
	return nil
}
